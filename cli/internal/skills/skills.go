package skills

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"
)

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
	return wrote, skipped, nil
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
