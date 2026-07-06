package skills

import (
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

func TestCheckBrokenLinkReference(t *testing.T) {
	fsys := makeSkillFS(map[string]string{
		"bad/SKILL.md": `---
name: bad
description: Broken link ref.
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

func makeSkillFS(files map[string]string) fs.FS {
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
