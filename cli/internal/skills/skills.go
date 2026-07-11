package skills

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

// SkillManifestEntry records a single scaffolded skill's identity and content
// fingerprint so a target project can detect drift.
type SkillManifestEntry struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Hash    string `json:"hash"`
}

// Manifest is the JSON document written to .agents/skills/MANIFEST.json.
type Manifest struct {
	Skills []SkillManifestEntry `json:"skills"`
}

// gitignorePatterns are appended idempotently to the target dir's .gitignore on
// each init. They ignore disposable loop state; EVIDENCE.md stays tracked.
var gitignorePatterns = []string{
	".loop/**/QUEUE.md",
	".loop/**/HANDOFF.md",
	".loop/**/REVIEW.md",
	".loop/**/specs/",
}

// Init scaffolds the embedded skills into targetDir/.agents/skills/.
// Existing skill directories are skipped rather than overwritten.
// It returns the lists of skill names written and skipped.
func Init(fsys fs.FS, targetDir string) (wrote []string, skipped []string, err error) {
	targetSkills := filepath.Join(targetDir, ".agents", "skills")
	entries, err := fs.ReadDir(fsys, ".")
	if err != nil {
		return nil, nil, fmt.Errorf("read embedded skills: %w", err)
	}
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		name := entry.Name()
		dstDir := filepath.Join(targetSkills, name)
		if _, statErr := os.Stat(dstDir); statErr == nil {
			skipped = append(skipped, name)
			continue
		} else if !os.IsNotExist(statErr) {
			return wrote, skipped, fmt.Errorf("stat target %q: %w", dstDir, statErr)
		}
		if err := copyDir(fsys, name, dstDir); err != nil {
			return wrote, skipped, fmt.Errorf("write skill %q: %w", name, err)
		}
		wrote = append(wrote, name)
	}

	if err := writeManifest(fsys, targetSkills, wrote, skipped); err != nil {
		return wrote, skipped, fmt.Errorf("write manifest: %w", err)
	}
	if err := appendGitignore(targetDir, gitignorePatterns); err != nil {
		return wrote, skipped, fmt.Errorf("write .gitignore: %w", err)
	}
	return wrote, skipped, nil
}

// UpdateReport records the outcome of a skills update: which skills were
// overwritten, which were left alone, and why.
type UpdateReport struct {
	Updated    []string // overwritten from embedded source
	SkippedMod []string // skipped because locally modified (without --force)
	Skipped    []string // skipped because already up-to-date
	Scaffolded []string // newly written because absent on disk
}

// Update refreshes the scaffolded skills under targetDir from the embedded
// source. For each embedded skill:
//   - if it is absent on disk it is scaffolded;
//   - if its on-disk hash differs from the manifest hash it is locally
//     modified: skipped unless force overwrites it;
//   - otherwise, if the embedded version is newer it is overwritten;
//   - otherwise it is up-to-date and skipped.
//
// With force, every embedded skill is overwritten regardless of modification.
// After updating, the manifest is rewritten to reflect the new versions and
// hashes of all on-disk skills.
func Update(fsys fs.FS, targetDir string, force bool) (UpdateReport, error) {
	var report UpdateReport
	targetSkills := filepath.Join(targetDir, ".agents", "skills")

	entries, err := fs.ReadDir(fsys, ".")
	if err != nil {
		return report, fmt.Errorf("read embedded skills: %w", err)
	}

	manifest, mErr := readManifest(os.DirFS(targetSkills))
	manifestBy := map[string]SkillManifestEntry{}
	if mErr == nil {
		for _, e := range manifest.Skills {
			manifestBy[e.Name] = e
		}
	} else if !errors.Is(mErr, fs.ErrNotExist) {
		return report, fmt.Errorf("read manifest: %w", mErr)
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		name := entry.Name()
		dstDir := filepath.Join(targetSkills, name)

		embedded, err := buildEntry(fsys, name)
		if err != nil {
			return report, err
		}

		_, statErr := os.Stat(dstDir)
		if statErr == nil {
			// exists on disk
			currentHash, err := dirHash(os.DirFS(targetSkills), name)
			if err != nil {
				return report, fmt.Errorf("hash skill %q: %w", name, err)
			}
			mEntry, hadManifest := manifestBy[name]
			modified := hadManifest && currentHash != mEntry.Hash

			if modified {
				if force {
					if err := copyDir(fsys, name, dstDir); err != nil {
						return report, fmt.Errorf("overwrite skill %q: %w", name, err)
					}
					report.Updated = append(report.Updated, name)
				} else {
					report.SkippedMod = append(report.SkippedMod, name)
				}
				continue
			}

			if force {
				if err := copyDir(fsys, name, dstDir); err != nil {
					return report, fmt.Errorf("overwrite skill %q: %w", name, err)
				}
				report.Updated = append(report.Updated, name)
				continue
			}

			baseVersion := ""
			if hadManifest {
				baseVersion = mEntry.Version
			}
			if versionNewer(embedded.Version, baseVersion) {
				if err := copyDir(fsys, name, dstDir); err != nil {
					return report, fmt.Errorf("overwrite skill %q: %w", name, err)
				}
				report.Updated = append(report.Updated, name)
			} else {
				report.Skipped = append(report.Skipped, name)
			}
			continue
		} else if !os.IsNotExist(statErr) {
			return report, fmt.Errorf("stat target %q: %w", dstDir, statErr)
		}

		// absent on disk: scaffold it
		if err := copyDir(fsys, name, dstDir); err != nil {
			return report, fmt.Errorf("scaffold skill %q: %w", name, err)
		}
		report.Scaffolded = append(report.Scaffolded, name)
	}

	written := append(append([]string{}, report.Updated...), report.Scaffolded...)
	if err := refreshManifest(targetSkills, manifestBy, written); err != nil {
		return report, fmt.Errorf("refresh manifest: %w", err)
	}
	return report, nil
}

// refreshManifest rewrites .agents/skills/MANIFEST.json from the current
// on-disk skills. Only skills in the written set get a fresh entry
// recomputed from disk; all others preserve their old manifest entry so
// that locally modified (SkippedMod) skills keep their original shipped
// hash and remain detectable as modified on future update calls.
func refreshManifest(targetSkills string, oldManifest map[string]SkillManifestEntry, written []string) error {
	writtenSet := map[string]bool{}
	for _, n := range written {
		writtenSet[n] = true
	}
	diskFS := os.DirFS(targetSkills)
	entries, err := fs.ReadDir(diskFS, ".")
	if err != nil {
		return fmt.Errorf("read skills dir: %w", err)
	}
	out := make([]SkillManifestEntry, 0, len(entries))
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		name := entry.Name()
		if writtenSet[name] || oldManifest[name] == (SkillManifestEntry{}) {
			// Skill was just written or has no prior manifest entry:
			// recompute from disk.
			e, err := buildEntry(diskFS, name)
			if err != nil {
				return err
			}
			out = append(out, e)
		} else {
			// Skill was not written: preserve the old entry.
			out = append(out, oldManifest[name])
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Name < out[j].Name })
	data, err := json.MarshalIndent(Manifest{Skills: out}, "", "  ")
	if err != nil {
		return err
	}
	data = append(data, '\n')
	return os.WriteFile(filepath.Join(targetSkills, "MANIFEST.json"), data, 0o644)
}

// writeManifest writes .agents/skills/MANIFEST.json after scaffolding.
// Existing entries for skipped skills are preserved so a re-run does not
// clobber the manifest with recomputed (potentially stale) hashes.
func writeManifest(fsys fs.FS, targetSkills string, wrote, skipped []string) error {
	manifestPath := filepath.Join(targetSkills, "MANIFEST.json")

	existing := map[string]SkillManifestEntry{}
	if data, err := os.ReadFile(manifestPath); err == nil {
		var m Manifest
		if err := json.Unmarshal(data, &m); err != nil {
			return fmt.Errorf("parse existing manifest: %w", err)
		}
		for _, e := range m.Skills {
			existing[e.Name] = e
		}
	} else if !os.IsNotExist(err) {
		return fmt.Errorf("read existing manifest: %w", err)
	}

	entries := make([]SkillManifestEntry, 0, len(existing)+len(wrote))
	for _, name := range skipped {
		if e, ok := existing[name]; ok {
			entries = append(entries, e)
		}
	}
	for _, name := range wrote {
		entry, err := buildEntry(fsys, name)
		if err != nil {
			return err
		}
		entries = append(entries, entry)
	}
	sort.Slice(entries, func(i, j int) bool { return entries[i].Name < entries[j].Name })

	data, err := json.MarshalIndent(Manifest{Skills: entries}, "", "  ")
	if err != nil {
		return err
	}
	data = append(data, '\n')
	return os.WriteFile(manifestPath, data, 0o644)
}

// buildEntry derives a manifest entry for a skill in fsys by reading its
// frontmatter version and content hash.
func buildEntry(fsys fs.FS, name string) (SkillManifestEntry, error) {
	data, err := fs.ReadFile(fsys, filepath.Join(name, "SKILL.md"))
	if err != nil {
		return SkillManifestEntry{}, fmt.Errorf("read skill %q: %w", name, err)
	}
	version := ""
	if front, _, ok := splitFrontmatter(data); ok {
		var meta struct {
			Metadata struct {
				Version string `yaml:"version"`
			} `yaml:"metadata"`
		}
		if err := yaml.Unmarshal(front, &meta); err == nil {
			version = meta.Metadata.Version
		}
	}
	hash, err := dirHash(fsys, name)
	if err != nil {
		return SkillManifestEntry{}, err
	}
	return SkillManifestEntry{Name: name, Version: version, Hash: hash}, nil
}

// dirHash returns a deterministic SHA-256 of all files under dir in fsys,
// ordered by path. It fingerprints the skill's entire content.
func dirHash(fsys fs.FS, dir string) (string, error) {
	var files []string
	err := fs.WalkDir(fsys, dir, func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			files = append(files, p)
		}
		return nil
	})
	if err != nil {
		return "", fmt.Errorf("walk %q: %w", dir, err)
	}
	sort.Strings(files)

	h := sha256.New()
	for _, f := range files {
		io.WriteString(h, f)
		h.Write([]byte{0})
		data, err := fs.ReadFile(fsys, f)
		if err != nil {
			return "", fmt.Errorf("read %q: %w", f, err)
		}
		h.Write(data)
		h.Write([]byte{0})
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

// appendGitignore adds patterns to targetDir/.gitignore without duplicating
// any pattern already present.
func appendGitignore(targetDir string, patterns []string) error {
	gitignorePath := filepath.Join(targetDir, ".gitignore")

	var lines []string
	seen := map[string]bool{}
	if data, err := os.ReadFile(gitignorePath); err == nil {
		for _, line := range strings.Split(string(data), "\n") {
			trimmed := strings.TrimSpace(line)
			if trimmed != "" {
				seen[trimmed] = true
			}
			lines = append(lines, line)
		}
	} else if !os.IsNotExist(err) {
		return fmt.Errorf("read .gitignore: %w", err)
	}

	added := false
	for _, p := range patterns {
		if !seen[strings.TrimSpace(p)] {
			lines = append(lines, p)
			added = true
		}
	}
	if len(lines) == 0 {
		return nil
	}
	content := strings.Join(lines, "\n")
	if added || !strings.HasSuffix(content, "\n") {
		content += "\n"
	}
	return os.WriteFile(gitignorePath, []byte(content), 0o644)
}

// Check validates every SKILL.md found under fsys.
// It returns one human-readable finding per issue. An empty slice with a nil
// error means all skills are valid. The returned findings are sorted by skill
// path and then by the order checks run.
func Check(fsys fs.FS) ([]string, error) {
	var skillPaths []string
	err := fs.WalkDir(fsys, ".", func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.Name() == "SKILL.md" && !d.IsDir() {
			skillPaths = append(skillPaths, p)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	sort.Strings(skillPaths)

	var findings []string
	for _, skillPath := range skillPaths {
		data, err := fs.ReadFile(fsys, skillPath)
		if err != nil {
			return nil, fmt.Errorf("read %s: %w", skillPath, err)
		}
		findings = append(findings, checkSkill(fsys, path.Dir(skillPath), data)...)
	}
	return findings, nil
}

// CheckVersioning compares on-disk scaffolded skills against the manifest and
// the embedded source. For each manifest entry it recomputes the on-disk skill
// hash: if it differs from the manifest hash the skill was modified locally; if
// it matches but the embedded source carries a newer version the scaffolded
// copy is stale. If no manifest exists it returns a single remediation finding.
// The returned findings are sorted.
func CheckVersioning(diskFS, embeddedFS fs.FS) ([]string, error) {
	manifest, err := readManifest(diskFS)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return []string{"no manifest: run 'knack skills init' to create one"}, nil
		}
		return nil, err
	}

	var findings []string
	for _, entry := range manifest.Skills {
		currentHash, err := dirHash(diskFS, entry.Name)
		if err != nil {
			// Skill absent on disk; nothing to compare against. Skip.
			continue
		}
		if currentHash != entry.Hash {
			findings = append(findings, fmt.Sprintf("modified: skill %s has local changes", entry.Name))
			continue
		}
		embedded, err := buildEntry(embeddedFS, entry.Name)
		if err != nil {
			continue
		}
		if versionNewer(embedded.Version, entry.Version) {
			findings = append(findings,
				fmt.Sprintf("stale: skill %s can be updated from v%s to v%s", entry.Name, entry.Version, embedded.Version))
		}
	}
	sort.Strings(findings)
	return findings, nil
}

func readManifest(diskFS fs.FS) (Manifest, error) {
	data, err := fs.ReadFile(diskFS, "MANIFEST.json")
	if err != nil {
		return Manifest{}, err
	}
	var m Manifest
	if err := json.Unmarshal(data, &m); err != nil {
		return Manifest{}, fmt.Errorf("parse manifest: %w", err)
	}
	return m, nil
}

// versionNewer reports whether candidate is a newer semantic version than
// baseline. Both are dotted-numeric (optionally "v"-prefixed). Missing or
// non-numeric segments compare as 0.
func versionNewer(candidate, baseline string) bool {
	cParts := strings.Split(strings.TrimPrefix(candidate, "v"), ".")
	bParts := strings.Split(strings.TrimPrefix(baseline, "v"), ".")
	n := len(cParts)
	if len(bParts) > n {
		n = len(bParts)
	}
	for i := 0; i < n; i++ {
		c := atoiSeg(cParts, i)
		b := atoiSeg(bParts, i)
		if c > b {
			return true
		}
		if c < b {
			return false
		}
	}
	return false
}

func atoiSeg(parts []string, i int) int {
	if i >= len(parts) {
		return 0
	}
	v, err := strconv.Atoi(strings.TrimSpace(parts[i]))
	if err != nil {
		return 0
	}
	return v
}

func copyDir(fsys fs.FS, src, dst string) error {
	return fs.WalkDir(fsys, src, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		dstPath := filepath.Join(dst, rel)
		if d.IsDir() {
			return os.MkdirAll(dstPath, 0o755)
		}
		if err := os.MkdirAll(filepath.Dir(dstPath), 0o755); err != nil {
			return err
		}
		srcFile, err := fsys.Open(path)
		if err != nil {
			return err
		}
		defer srcFile.Close()
		dstFile, err := os.Create(dstPath)
		if err != nil {
			return err
		}
		defer dstFile.Close()
		if _, err := io.Copy(dstFile, srcFile); err != nil {
			return err
		}
		return nil
	})
}

var (
	wikiRefRe = regexp.MustCompile(`\[\[([^\]]+)\]\]`)
	linkRefRe = regexp.MustCompile(`\[([^\]]+)\]\(([^)]+)\)`)
)

func checkSkill(fsys fs.FS, skillDir string, data []byte) []string {
	var findings []string
	skillName := path.Base(skillDir)

	front, body, ok := splitFrontmatter(data)
	if !ok {
		return []string{fmt.Sprintf("%s: missing frontmatter", skillName)}
	}

	var meta struct {
		Name        string `yaml:"name"`
		Description string `yaml:"description"`
		Metadata    struct {
			Version string `yaml:"version"`
		} `yaml:"metadata"`
	}
	if err := yaml.Unmarshal(front, &meta); err != nil {
		return []string{fmt.Sprintf("%s: invalid frontmatter YAML: %v", skillName, err)}
	}
	if strings.TrimSpace(meta.Name) == "" {
		findings = append(findings, fmt.Sprintf("%s: frontmatter field \"name\" is empty", skillName))
	}
	if strings.TrimSpace(meta.Description) == "" {
		findings = append(findings, fmt.Sprintf("%s: frontmatter field \"description\" is empty", skillName))
	}
	if strings.TrimSpace(meta.Metadata.Version) == "" {
		findings = append(findings, fmt.Sprintf("%s: frontmatter field \"version\" is empty", skillName))
	}

	findings = append(findings, checkReferences(fsys, skillDir, body)...)
	return findings
}

func splitFrontmatter(data []byte) (front, body []byte, ok bool) {
	s := string(data)
	lines := strings.Split(s, "\n")
	if len(lines) < 2 || strings.TrimRight(lines[0], "\r") != "---" {
		return nil, nil, false
	}
	endIdx := -1
	for i := 1; i < len(lines); i++ {
		if strings.TrimRight(lines[i], "\r") == "---" {
			endIdx = i
			break
		}
	}
	if endIdx == -1 {
		return nil, nil, false
	}
	front = []byte(strings.Join(lines[1:endIdx], "\n"))
	body = []byte(strings.Join(lines[endIdx+1:], "\n"))
	return front, body, true
}

func checkReferences(fsys fs.FS, skillDir string, body []byte) []string {
	var findings []string
	s := string(body)
	skillName := path.Base(skillDir)

	for _, m := range wikiRefRe.FindAllStringSubmatch(s, -1) {
		target := stripAnchor(m[1])
		if target == "" {
			findings = append(findings, fmt.Sprintf("%s: broken reference [[%s]]: empty target", skillName, m[1]))
			continue
		}
		if isExternal(target) {
			continue
		}
		if !refExists(fsys, skillDir, target) {
			findings = append(findings, fmt.Sprintf("%s: broken reference [[%s]]: file not found", skillName, m[1]))
		}
	}

	for _, m := range linkRefRe.FindAllStringSubmatch(s, -1) {
		target := stripAnchor(m[2])
		if target == "" {
			findings = append(findings, fmt.Sprintf("%s: broken reference [%s](): empty target", skillName, m[1]))
			continue
		}
		if isExternal(target) {
			continue
		}
		if !refExists(fsys, skillDir, target) {
			findings = append(findings, fmt.Sprintf("%s: broken reference [%s](%s): file not found", skillName, m[1], m[2]))
		}
	}

	return findings
}

func stripAnchor(target string) string {
	if i := strings.IndexByte(target, '#'); i >= 0 {
		return target[:i]
	}
	return target
}

func isExternal(target string) bool {
	return strings.HasPrefix(target, "http://") ||
		strings.HasPrefix(target, "https://") ||
		strings.HasPrefix(target, "mailto:") ||
		strings.HasPrefix(target, "ftp://") ||
		strings.HasPrefix(target, "file://") ||
		strings.HasPrefix(target, "//")
}

func refExists(fsys fs.FS, skillDir, target string) bool {
	_, err := fs.Stat(fsys, path.Join(skillDir, target))
	return err == nil
}
