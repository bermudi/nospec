package skills

import (
	"encoding/json"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"
	"testing/fstest"
)

func TestInitScaffoldsAllSkills(t *testing.T) {
	fsys := os.DirFS("../../embedded/skills")
	target := t.TempDir()
	wrote, skipped, err := Init(fsys, target)
	if err != nil {
		t.Fatalf("Init failed: %v", err)
	}
	expected := []string{"build", "decide", "domain-modeling", "explore", "fix", "plan", "review"}
	if len(skipped) != 0 {
		t.Fatalf("expected no skipped skills on first init, got %v", skipped)
	}
	if !slices.Equal(sorted(wrote), expected) {
		t.Fatalf("expected wrote %v, got %v", expected, wrote)
	}
	for _, name := range expected {
		skillPath := filepath.Join(target, ".agents", "skills", name, "SKILL.md")
		if _, err := os.Stat(skillPath); err != nil {
			t.Fatalf("expected skill %s to be scaffolded at %s: %v", name, skillPath, err)
		}
	}
}

func TestInitSkipsExistingSkills(t *testing.T) {
	fsys := os.DirFS("../../embedded/skills")
	target := t.TempDir()
	if _, _, err := Init(fsys, target); err != nil {
		t.Fatalf("first Init failed: %v", err)
	}
	wrote, skipped, err := Init(fsys, target)
	if err != nil {
		t.Fatalf("second Init failed: %v", err)
	}
	if len(wrote) != 0 {
		t.Fatalf("expected no skills written on second init, got %v", wrote)
	}
	if len(skipped) != 7 {
		t.Fatalf("expected 7 skipped skills, got %v", skipped)
	}
}

func TestCheckValidSkill(t *testing.T) {
	fsys := makeSkillFS(map[string]string{
		"valid/SKILL.md": `---
name: valid
description: A valid skill.
metadata:
  version: "1.0.0"
---

See [[notes.md]] and [notes](notes.md).
`,
		"valid/notes.md": "notes",
	})
	findings, err := Check(fsys)
	if err != nil {
		t.Fatalf("Check failed: %v", err)
	}
	if len(findings) != 0 {
		t.Fatalf("expected no findings, got %v", findings)
	}
}

func TestCheckMissingFrontmatter(t *testing.T) {
	fsys := makeSkillFS(map[string]string{
		"bad/SKILL.md": "# No frontmatter\n",
	})
	findings, err := Check(fsys)
	if err != nil {
		t.Fatalf("Check failed: %v", err)
	}
	if len(findings) != 1 || !strings.Contains(findings[0], "missing frontmatter") {
		t.Fatalf("expected missing frontmatter finding, got %v", findings)
	}
}

func TestCheckEmptyDescription(t *testing.T) {
	fsys := makeSkillFS(map[string]string{
		"bad/SKILL.md": `---
name: bad
description:
metadata:
  version: "1.0.0"
---
`,
	})
	findings, err := Check(fsys)
	if err != nil {
		t.Fatalf("Check failed: %v", err)
	}
	if len(findings) != 1 || !strings.Contains(findings[0], `"description" is empty`) {
		t.Fatalf("expected empty description finding, got %v", findings)
	}
}

func TestCheckBrokenWikiReference(t *testing.T) {
	fsys := makeSkillFS(map[string]string{
		"bad/SKILL.md": `---
name: bad
description: Broken wiki ref.
metadata:
  version: "1.0.0"
---

See [[missing.md]].
`,
	})
	findings, err := Check(fsys)
	if err != nil {
		t.Fatalf("Check failed: %v", err)
	}
	if len(findings) != 1 || !strings.Contains(findings[0], "broken reference") || !strings.Contains(findings[0], "missing.md") {
		t.Fatalf("expected broken reference finding, got %v", findings)
	}
}

func TestCheckEmptyVersion(t *testing.T) {
	fsys := makeSkillFS(map[string]string{
		"bad/SKILL.md": `---
name: bad
description: Missing version.
metadata:
  version:
---
`,
	})
	findings, err := Check(fsys)
	if err != nil {
		t.Fatalf("Check failed: %v", err)
	}
	if len(findings) != 1 || !strings.Contains(findings[0], `"version" is empty`) {
		t.Fatalf("expected empty version finding, got %v", findings)
	}
}

func TestCheckBrokenLinkReference(t *testing.T) {
	fsys := makeSkillFS(map[string]string{
		"bad/SKILL.md": `---
name: bad
description: Broken link ref.
metadata:
  version: "1.0.0"
---

See [missing](missing.md).
`,
	})
	findings, err := Check(fsys)
	if err != nil {
		t.Fatalf("Check failed: %v", err)
	}
	if len(findings) != 1 || !strings.Contains(findings[0], "broken reference") || !strings.Contains(findings[0], "missing.md") {
		t.Fatalf("expected broken reference finding, got %v", findings)
	}
}

func TestCheckIgnoresExternalURLs(t *testing.T) {
	fsys := makeSkillFS(map[string]string{
		"ok/SKILL.md": `---
name: ok
description: External links are fine.
metadata:
  version: "1.0.0"
---

See [site](https://example.com) and [mail](mailto:a@b.com).
`,
	})
	findings, err := Check(fsys)
	if err != nil {
		t.Fatalf("Check failed: %v", err)
	}
	if len(findings) != 0 {
		t.Fatalf("expected no findings for external URLs, got %v", findings)
	}
}

func TestCheckRealProjectSkills(t *testing.T) {
	fsys := os.DirFS("../../../.agents/skills")
	findings, err := Check(fsys)
	if err != nil {
		t.Fatalf("Check failed: %v", err)
	}
	if len(findings) != 0 {
		t.Fatalf("project skills have validation findings: %v", findings)
	}
}

func TestCheckEmbeddedSkills(t *testing.T) {
	fsys := os.DirFS("../../embedded/skills")
	findings, err := Check(fsys)
	if err != nil {
		t.Fatalf("Check failed: %v", err)
	}
	if len(findings) != 0 {
		t.Fatalf("embedded skills have validation findings: %v", findings)
	}
}

func makeSkillFS(files map[string]string) fstest.MapFS {
	fsys := fstest.MapFS{}
	for name, data := range files {
		fsys[name] = &fstest.MapFile{Data: []byte(data)}
	}
	return fsys
}

func sorted(ss []string) []string {
	out := append([]string{}, ss...)
	slices.Sort(out)
	return out
}

func readManifestAt(t *testing.T, target string) Manifest {
	t.Helper()
	data, err := os.ReadFile(filepath.Join(target, ".agents", "skills", "MANIFEST.json"))
	if err != nil {
		t.Fatalf("read manifest: %v", err)
	}
	var m Manifest
	if err := json.Unmarshal(data, &m); err != nil {
		t.Fatalf("parse manifest: %v", err)
	}
	return m
}

func findEntry(m Manifest, name string) (SkillManifestEntry, bool) {
	for _, e := range m.Skills {
		if e.Name == name {
			return e, true
		}
	}
	return SkillManifestEntry{}, false
}

func TestInitWritesManifest(t *testing.T) {
	fsys := os.DirFS("../../embedded/skills")
	target := t.TempDir()
	if _, _, err := Init(fsys, target); err != nil {
		t.Fatalf("Init failed: %v", err)
	}
	path := filepath.Join(target, ".agents", "skills", "MANIFEST.json")
	if _, err := os.Stat(path); err != nil {
		t.Fatalf("expected manifest at %s: %v", path, err)
	}
	m := readManifestAt(t, target)
	if len(m.Skills) != 7 {
		t.Fatalf("expected 7 skills in manifest, got %d", len(m.Skills))
	}
}

func TestInitManifestContainsVersionsAndHashes(t *testing.T) {
	fsys := os.DirFS("../../embedded/skills")
	target := t.TempDir()
	if _, _, err := Init(fsys, target); err != nil {
		t.Fatalf("Init failed: %v", err)
	}
	m := readManifestAt(t, target)

	expected := []string{"build", "decide", "domain-modeling", "explore", "fix", "plan", "review"}
	if got := sorted(skillNames(m)); !slices.Equal(got, expected) {
		t.Fatalf("expected skills %v, got %v", expected, got)
	}

	for _, name := range expected {
		entry, ok := findEntry(m, name)
		if !ok {
			t.Fatalf("expected manifest entry for %s", name)
		}
		if entry.Version != "1.0.0" {
			t.Fatalf("skill %s: expected version 1.0.0, got %q", name, entry.Version)
		}
		if entry.Hash == "" {
			t.Fatalf("skill %s: expected non-empty hash", name)
		}
		// Hash must match a fresh computation from the embedded source.
		want, err := buildEntry(fsys, name)
		if err != nil {
			t.Fatalf("buildEntry(%s): %v", name, err)
		}
		if entry.Hash != want.Hash {
			t.Fatalf("skill %s: hash mismatch: got %s, want %s", name, entry.Hash, want.Hash)
		}
	}
}

func TestInitManifestNotClobberedOnSkip(t *testing.T) {
	fsys := os.DirFS("../../embedded/skills")
	target := t.TempDir()
	if _, _, err := Init(fsys, target); err != nil {
		t.Fatalf("first Init failed: %v", err)
	}
	first := readManifestAt(t, target)

	// Simulate the user editing a scaffolded skill on disk. A re-run should
	// preserve the manifest rather than recompute stale hashes for skipped dirs.
	edited := filepath.Join(target, ".agents", "skills", "build", "SKILL.md")
	if err := os.WriteFile(edited, []byte("tampered\n"), 0o644); err != nil {
		t.Fatalf("tamper skill: %v", err)
	}

	if _, _, err := Init(fsys, target); err != nil {
		t.Fatalf("second Init failed: %v", err)
	}
	second := readManifestAt(t, target)

	firstBuild, ok := findEntry(first, "build")
	if !ok {
		t.Fatalf("missing build entry in first manifest")
	}
	secondBuild, ok := findEntry(second, "build")
	if !ok {
		t.Fatalf("missing build entry in second manifest")
	}
	if firstBuild.Hash != secondBuild.Hash {
		t.Fatalf("manifest clobbered on skip: hash changed from %s to %s", firstBuild.Hash, secondBuild.Hash)
	}
	if len(second.Skills) != 7 {
		t.Fatalf("expected 7 skills after re-run, got %d", len(second.Skills))
	}
}

func TestInitWritesGitignorePatterns(t *testing.T) {
	fsys := os.DirFS("../../embedded/skills")
	target := t.TempDir()
	if _, _, err := Init(fsys, target); err != nil {
		t.Fatalf("Init failed: %v", err)
	}
	checkGitignore(t, target, gitignorePatterns)
}

func TestInitGitignoreIdempotent(t *testing.T) {
	fsys := os.DirFS("../../embedded/skills")
	target := t.TempDir()
	if _, _, err := Init(fsys, target); err != nil {
		t.Fatalf("first Init failed: %v", err)
	}
	if _, _, err := Init(fsys, target); err != nil {
		t.Fatalf("second Init failed: %v", err)
	}
	content := readFile(t, filepath.Join(target, ".gitignore"))
	for _, p := range gitignorePatterns {
		if c := strings.Count(content, p); c != 1 {
			t.Fatalf("pattern %q appears %d times, want 1", p, c)
		}
	}
}

func checkGitignore(t *testing.T, target string, patterns []string) {
	t.Helper()
	content := readFile(t, filepath.Join(target, ".gitignore"))
	for _, p := range patterns {
		if !strings.Contains(content, p) {
			t.Fatalf("expected .gitignore to contain %q, got:\n%s", p, content)
		}
	}
}

func readFile(t *testing.T, path string) string {
	t.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read %s: %v", path, err)
	}
	return string(data)
}

func skillNames(m Manifest) []string {
	names := make([]string, 0, len(m.Skills))
	for _, e := range m.Skills {
		names = append(names, e.Name)
	}
	return names
}

const (
	skillManifestV = "---\nname: x\ndescription: a skill\nmetadata:\n  version: \"1.0.0\"\n---\nbody\n"
	skillEmbeddedV = "---\nname: x\ndescription: a skill\nmetadata:\n  version: \"1.1.0\"\n---\nbody\n"
	skillModifiedV = "---\nname: x\ndescription: a skill\nmetadata:\n  version: \"1.0.0\"\n---\nchanged body\n"
)

func addManifest(t *testing.T, disk fstest.MapFS, entries ...SkillManifestEntry) {
	t.Helper()
	data, err := json.MarshalIndent(Manifest{Skills: entries}, "", "  ")
	if err != nil {
		t.Fatalf("marshal manifest: %v", err)
	}
	disk["MANIFEST.json"] = &fstest.MapFile{Data: data}
}

func diskHash(t *testing.T, disk fs.FS, name string) string {
	t.Helper()
	h, err := dirHash(disk, name)
	if err != nil {
		t.Fatalf("dirHash(%s): %v", name, err)
	}
	return h
}

func TestCheckNoManifest(t *testing.T) {
	disk := makeSkillFS(map[string]string{
		"x/SKILL.md": skillManifestV,
	})
	embedded := makeSkillFS(map[string]string{
		"x/SKILL.md": skillEmbeddedV,
	})
	findings, err := CheckVersioning(disk, embedded)
	if err != nil {
		t.Fatalf("CheckVersioning failed: %v", err)
	}
	if len(findings) != 1 || !strings.Contains(findings[0], "no manifest") {
		t.Fatalf("expected no-manifest finding, got %v", findings)
	}
}

func TestCheckModified(t *testing.T) {
	disk := makeSkillFS(map[string]string{
		"x/SKILL.md": skillModifiedV,
	})
	// Manifest records the hash from before the local edit.
	orig := makeSkillFS(map[string]string{"x/SKILL.md": skillManifestV})
	hash := diskHash(t, orig, "x")
	addManifest(t, disk, SkillManifestEntry{Name: "x", Version: "1.0.0", Hash: hash})

	embedded := makeSkillFS(map[string]string{
		"x/SKILL.md": skillEmbeddedV,
	})
	findings, err := CheckVersioning(disk, embedded)
	if err != nil {
		t.Fatalf("CheckVersioning failed: %v", err)
	}
	if len(findings) != 1 || !strings.Contains(findings[0], "modified: skill x has local changes") {
		t.Fatalf("expected modified finding, got %v", findings)
	}
}

func TestCheckStale(t *testing.T) {
	disk := makeSkillFS(map[string]string{
		"x/SKILL.md": skillManifestV,
	})
	hash := diskHash(t, disk, "x")
	addManifest(t, disk, SkillManifestEntry{Name: "x", Version: "1.0.0", Hash: hash})

	embedded := makeSkillFS(map[string]string{
		"x/SKILL.md": skillEmbeddedV,
	})
	findings, err := CheckVersioning(disk, embedded)
	if err != nil {
		t.Fatalf("CheckVersioning failed: %v", err)
	}
	if len(findings) != 1 ||
		!strings.Contains(findings[0], "stale: skill x can be updated from v1.0.0 to v1.1.0") {
		t.Fatalf("expected stale finding, got %v", findings)
	}
}

func TestCheckUpToDate(t *testing.T) {
	disk := makeSkillFS(map[string]string{
		"x/SKILL.md": skillManifestV,
	})
	hash := diskHash(t, disk, "x")
	addManifest(t, disk, SkillManifestEntry{Name: "x", Version: "1.0.0", Hash: hash})

	embedded := makeSkillFS(map[string]string{
		"x/SKILL.md": skillManifestV,
	})
	findings, err := CheckVersioning(disk, embedded)
	if err != nil {
		t.Fatalf("CheckVersioning failed: %v", err)
	}
	if len(findings) != 0 {
		t.Fatalf("expected no findings for up-to-date skill, got %v", findings)
	}
}

func writeSkillOnDisk(t *testing.T, target, name, skillMd string) {
	t.Helper()
	dir := filepath.Join(target, ".agents", "skills", name)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatalf("mkdir %s: %v", dir, err)
	}
	if err := os.WriteFile(filepath.Join(dir, "SKILL.md"), []byte(skillMd), 0o644); err != nil {
		t.Fatalf("write skill %s: %v", name, err)
	}
}

func readSkillOnDisk(t *testing.T, target, name string) string {
	t.Helper()
	data, err := os.ReadFile(filepath.Join(diskSkillsDir(target), name, "SKILL.md"))
	if err != nil {
		t.Fatalf("read skill %s: %v", name, err)
	}
	return string(data)
}

func writeManifestOnDisk(t *testing.T, target string, entries []SkillManifestEntry) {
	t.Helper()
	dir := filepath.Join(target, ".agents", "skills")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatalf("mkdir %s: %v", dir, err)
	}
	data, err := json.MarshalIndent(Manifest{Skills: entries}, "", "  ")
	if err != nil {
		t.Fatalf("marshal manifest: %v", err)
	}
	data = append(data, '\n')
	if err := os.WriteFile(filepath.Join(dir, "MANIFEST.json"), data, 0o644); err != nil {
		t.Fatalf("write manifest: %v", err)
	}
}

func diskSkillsDir(target string) string {
	return filepath.Join(target, ".agents", "skills")
}

func TestUpdateStaleSkill(t *testing.T) {
	embedded := makeSkillFS(map[string]string{"x/SKILL.md": skillEmbeddedV})
	target := t.TempDir()
	writeSkillOnDisk(t, target, "x", skillManifestV)
	hash := diskHash(t, os.DirFS(diskSkillsDir(target)), "x")
	writeManifestOnDisk(t, target, []SkillManifestEntry{{Name: "x", Version: "1.0.0", Hash: hash}})

	report, err := Update(embedded, target, false)
	if err != nil {
		t.Fatalf("Update failed: %v", err)
	}
	if !slices.Contains(report.Updated, "x") {
		t.Fatalf("expected x in Updated, got %v", report.Updated)
	}
	got := readFile(t, filepath.Join(diskSkillsDir(target), "x", "SKILL.md"))
	if got != skillEmbeddedV {
		t.Fatalf("expected x overwritten with embedded content, got %q", got)
	}
	m := readManifestAt(t, target)
	entry, ok := findEntry(m, "x")
	if !ok {
		t.Fatalf("missing manifest entry for x")
	}
	if entry.Version != "1.1.0" {
		t.Fatalf("expected manifest version 1.1.0, got %q", entry.Version)
	}
}

func TestUpdateSkipsModifiedSkill(t *testing.T) {
	embedded := makeSkillFS(map[string]string{"x/SKILL.md": skillEmbeddedV})
	target := t.TempDir()
	writeSkillOnDisk(t, target, "x", skillModifiedV)
	// Manifest records the pre-edit hash (skillManifestV), so the modified
	// on-disk content diverges from the manifest.
	orig := makeSkillFS(map[string]string{"x/SKILL.md": skillManifestV})
	hash := diskHash(t, orig, "x")
	writeManifestOnDisk(t, target, []SkillManifestEntry{{Name: "x", Version: "1.0.0", Hash: hash}})

	report, err := Update(embedded, target, false)
	if err != nil {
		t.Fatalf("Update failed: %v", err)
	}
	if !slices.Contains(report.SkippedMod, "x") {
		t.Fatalf("expected x in SkippedMod, got %v", report.SkippedMod)
	}
	if len(report.Updated) != 0 {
		t.Fatalf("expected no Updated skills, got %v", report.Updated)
	}
	got := readFile(t, filepath.Join(diskSkillsDir(target), "x", "SKILL.md"))
	if got != skillModifiedV {
		t.Fatalf("expected modified content preserved, got %q", got)
	}
}

func TestUpdateForceOverwritesModifiedSkill(t *testing.T) {
	embedded := makeSkillFS(map[string]string{"x/SKILL.md": skillEmbeddedV})
	target := t.TempDir()
	writeSkillOnDisk(t, target, "x", skillModifiedV)
	orig := makeSkillFS(map[string]string{"x/SKILL.md": skillManifestV})
	hash := diskHash(t, orig, "x")
	writeManifestOnDisk(t, target, []SkillManifestEntry{{Name: "x", Version: "1.0.0", Hash: hash}})

	report, err := Update(embedded, target, true)
	if err != nil {
		t.Fatalf("Update failed: %v", err)
	}
	if !slices.Contains(report.Updated, "x") {
		t.Fatalf("expected x in Updated, got %v", report.Updated)
	}
	got := readFile(t, filepath.Join(diskSkillsDir(target), "x", "SKILL.md"))
	if got != skillEmbeddedV {
		t.Fatalf("expected x overwritten with embedded content, got %q", got)
	}
}

func TestUpdateScaffoldsNewSkill(t *testing.T) {
	embedded := makeSkillFS(map[string]string{"x/SKILL.md": skillEmbeddedV})
	target := t.TempDir()
	// No skill on disk, but manifest present (e.g. from a prior init of other skills).
	writeManifestOnDisk(t, target, []SkillManifestEntry{})

	report, err := Update(embedded, target, false)
	if err != nil {
		t.Fatalf("Update failed: %v", err)
	}
	if !slices.Contains(report.Scaffolded, "x") {
		t.Fatalf("expected x in Scaffolded, got %v", report.Scaffolded)
	}
	if _, err := os.Stat(filepath.Join(diskSkillsDir(target), "x", "SKILL.md")); err != nil {
		t.Fatalf("expected scaffolded skill x on disk: %v", err)
	}
}

func TestUpdateRefreshManifest(t *testing.T) {
	embedded := makeSkillFS(map[string]string{"x/SKILL.md": skillEmbeddedV})
	target := t.TempDir()
	writeSkillOnDisk(t, target, "x", skillManifestV)
	hash := diskHash(t, os.DirFS(diskSkillsDir(target)), "x")
	writeManifestOnDisk(t, target, []SkillManifestEntry{{Name: "x", Version: "1.0.0", Hash: hash}})

	if _, err := Update(embedded, target, false); err != nil {
		t.Fatalf("Update failed: %v", err)
	}
	m := readManifestAt(t, target)
	entry, ok := findEntry(m, "x")
	if !ok {
		t.Fatalf("missing manifest entry for x")
	}
	if entry.Version != "1.1.0" {
		t.Fatalf("expected refreshed manifest version 1.1.0, got %q", entry.Version)
	}
	wantHash, err := dirHash(os.DirFS(diskSkillsDir(target)), "x")
	if err != nil {
		t.Fatalf("dirHash: %v", err)
	}
	if entry.Hash != wantHash {
		t.Fatalf("expected refreshed manifest hash %s, got %s", wantHash, entry.Hash)
	}
}

func TestUpdatePreservesManifestForModifiedSkill(t *testing.T) {
	embedded := makeSkillFS(map[string]string{"x/SKILL.md": skillEmbeddedV})
	target := t.TempDir()
	// On-disk skill is modified (different from manifest baseline).
	writeSkillOnDisk(t, target, "x", skillModifiedV)
	origHash := diskHash(t, os.DirFS(diskSkillsDir(target)), "x")
	// Manifest records the original shipped hash (skillManifestV), not the modified hash.
	origManifestHash := diskHash(t, makeSkillFS(map[string]string{"x/SKILL.md": skillManifestV}), "x")
	writeManifestOnDisk(t, target, []SkillManifestEntry{{Name: "x", Version: "1.0.0", Hash: origManifestHash}})

	report, err := Update(embedded, target, false)
	if err != nil {
		t.Fatalf("Update failed: %v", err)
	}
	if !slices.Contains(report.SkippedMod, "x") {
		t.Fatalf("expected x in SkippedMod, got %v", report.SkippedMod)
	}
	// The manifest must still record the original shipped hash, not the modified hash.
	m := readManifestAt(t, target)
	entry, ok := findEntry(m, "x")
	if !ok {
		t.Fatalf("missing manifest entry for x")
	}
	if entry.Hash != origManifestHash {
		t.Fatalf("manifest hash should be preserved as %s (original shipped), got %s (modified content)", origManifestHash, entry.Hash)
	}
	if entry.Version != "1.0.0" {
		t.Fatalf("manifest version should be preserved as 1.0.0, got %q", entry.Version)
	}
	// The on-disk content should be unchanged.
	got := readSkillOnDisk(t, target, "x")
	if got != skillModifiedV {
		t.Fatalf("on-disk skill should be unchanged")
	}
	_ = origHash // current modified hash, should NOT be in manifest
}
