package decisions

import (
	"io/fs"
	"strings"
	"testing"
	"testing/fstest"
)

func TestList(t *testing.T) {
	fsys := makeADRFS(map[string]string{
		"decisions/0001-first.md": `# 0001: First Decision
Date: 2026-07-06
Status: accepted

## Context
One.
`,
		"decisions/0002-second.md": `# 0002: Second Decision
Status: proposed

## Context
Two.
`,
	})
	adrs, err := List(fsys, "decisions")
	if err != nil {
		t.Fatalf("List failed: %v", err)
	}
	if len(adrs) != 2 {
		t.Fatalf("expected 2 ADRs, got %d", len(adrs))
	}
	if adrs[0].Number != "0001" || adrs[0].Title != "First Decision" || adrs[0].Status != "accepted" {
		t.Fatalf("unexpected first ADR: %+v", adrs[0])
	}
	if adrs[1].Number != "0002" || adrs[1].Title != "Second Decision" || adrs[1].Status != "proposed" {
		t.Fatalf("unexpected second ADR: %+v", adrs[1])
	}
}

func TestShow(t *testing.T) {
	contents := `# 0001: First Decision
Status: accepted

Body.
`
	fsys := makeADRFS(map[string]string{
		"decisions/0001-first.md": contents,
	})
	data, err := Show(fsys, "decisions", "1")
	if err != nil {
		t.Fatalf("Show failed: %v", err)
	}
	if string(data) != contents {
		t.Fatalf("unexpected show output:\n%s", string(data))
	}
}

func TestShowNotFound(t *testing.T) {
	fsys := makeADRFS(map[string]string{
		"decisions/0001-first.md": "# 0001\n",
	})
	_, err := Show(fsys, "decisions", "9")
	if err == nil {
		t.Fatalf("expected error for missing ADR")
	}
}

func TestCheckCleanCoverage(t *testing.T) {
	fsys := makeADRFS(map[string]string{
		"decisions/0001-first.md":  "# 0001: First\nStatus: accepted\n",
		"decisions/0002-second.md": "# 0002: Second\nStatus: accepted\n",
		".loop/QUEUE.md": lines(
			"## First unit",
			"",
			"Implements ADR-0001.",
			"",
			"Verify:",
			"```bash",
			"true",
			"```",
			"",
			"Status: pending",
			"",
			"## Second unit",
			"",
			"Implements ADR-0002.",
			"",
			"Verify:",
			"```bash",
			"true",
			"```",
			"",
			"Status: pending",
		),
	})
	findings, err := Check(fsys, "decisions", fsys, ".loop")
	if err != nil {
		t.Fatalf("Check failed: %v", err)
	}
	if len(findings) != 0 {
		t.Fatalf("expected clean coverage, got: %v", findings)
	}
}

func TestCheckOrphanedAndDangling(t *testing.T) {
	fsys := makeADRFS(map[string]string{
		"decisions/0001-first.md":  "# 0001: First\nStatus: accepted\n",
		"decisions/0002-second.md": "# 0002: Second\nStatus: accepted\n",
		".loop/QUEUE.md": lines(
			"## Only unit",
			"",
			"Implements ADR-0001 and ADR-0003.",
			"",
			"Verify:",
			"```bash",
			"true",
			"```",
			"",
			"Status: pending",
		),
	})
	findings, err := Check(fsys, "decisions", fsys, ".loop")
	if err != nil {
		t.Fatalf("Check failed: %v", err)
	}
	if len(findings) != 2 {
		t.Fatalf("expected 2 findings, got %d: %v", len(findings), findings)
	}
	var hasOrphan, hasDangle bool
	for _, f := range findings {
		if strings.Contains(f, "orphaned ADR 0002") {
			hasOrphan = true
		}
		if strings.Contains(f, "dangling reference ADR-0003") {
			hasDangle = true
		}
	}
	if !hasOrphan {
		t.Fatalf("expected orphaned ADR 0002 finding, got: %v", findings)
	}
	if !hasDangle {
		t.Fatalf("expected dangling ADR-0003 finding, got: %v", findings)
	}
}

func TestCheckEvidenceCoversCompletedCycle(t *testing.T) {
	fsys := makeADRFS(map[string]string{
		"decisions/0001-first.md": "# 0001: First\nStatus: accepted\n",
		".loop/done-cycle/EVIDENCE.md": lines(
			"## 2026-07-06T12:00:00 — Some unit",
			"",
			"Status: done",
			"",
			"Unit:",
			"````markdown",
			"## Some unit",
			"",
			"Read first:",
			"- decisions/0001-first.md",
			"",
			"Verify:",
			"```bash",
			"true",
			"```",
			"````",
			"",
			"Verify output:",
			"```text",
			"ok",
			"```",
		),
	})
	findings, err := Check(fsys, "decisions", fsys, ".loop")
	if err != nil {
		t.Fatalf("Check failed: %v", err)
	}
	if len(findings) != 0 {
		t.Fatalf("expected no findings (ADR covered by evidence), got: %v", findings)
	}
}

func TestCheckByFilename(t *testing.T) {
	fsys := makeADRFS(map[string]string{
		"decisions/0001-build-the-cli-in-go.md": "# 0001: Build the CLI in Go\nStatus: accepted\n",
		".loop/QUEUE.md": lines(
			"## Unit",
			"",
			"See decisions/0001-build-the-cli-in-go.md for context.",
			"",
			"Verify:",
			"```bash",
			"true",
			"```",
			"",
			"Status: pending",
		),
	})
	findings, err := Check(fsys, "decisions", fsys, ".loop")
	if err != nil {
		t.Fatalf("Check failed: %v", err)
	}
	if len(findings) != 0 {
		t.Fatalf("expected coverage by filename, got: %v", findings)
	}
}

func TestCheckGrandfatheredADRNotOrphaned(t *testing.T) {
	fsys := makeADRFS(map[string]string{
		"decisions/0001-first.md": "# 0001: First\nStatus: accepted\nGrandfathered: predates the ledger (ADR-0006).\n",
		"decisions/0002-second.md": "# 0002: Second\nStatus: accepted\n",
		".loop/QUEUE.md": lines(
			"## Unit",
			"",
			"Implements ADR-0002.",
			"",
			"Verify:",
			"```bash",
			"true",
			"```",
			"",
			"Status: pending",
		),
	})
	findings, err := Check(fsys, "decisions", fsys, ".loop")
	if err != nil {
		t.Fatalf("Check failed: %v", err)
	}
	if len(findings) != 0 {
		t.Fatalf("expected no findings (0001 grandfathered, 0002 referenced), got: %v", findings)
	}
}

func TestCheckGrandfatheredADRStillParsedByList(t *testing.T) {
	fsys := makeADRFS(map[string]string{
		"decisions/0001-first.md": "# 0001: First\nStatus: accepted\nGrandfathered: predates the ledger.\n",
	})
	adrs, err := List(fsys, "decisions")
	if err != nil {
		t.Fatalf("List failed: %v", err)
	}
	if len(adrs) != 1 {
		t.Fatalf("expected 1 ADR, got %d", len(adrs))
	}
	if !adrs[0].Grandfather {
		t.Fatalf("expected Grandfather=true, got false")
	}
}

func makeADRFS(files map[string]string) fs.FS {
	fsys := fstest.MapFS{}
	for name, data := range files {
		fsys[name] = &fstest.MapFile{Data: []byte(data)}
	}
	return fsys
}

func lines(ss ...string) string {
	return strings.Join(ss, "\n")
}
