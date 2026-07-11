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
		"decisions/0001-first.md":  "# 0001: First\nStatus: accepted\nGrandfathered: predates the ledger (ADR-0006).\n",
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

func TestListParsesSupersedeChain(t *testing.T) {
	fsys := makeADRFS(map[string]string{
		"decisions/0001-first.md":  "# 0001: First\nStatus: superseded\nSuperseded by: ADR-0002\n",
		"decisions/0002-second.md": "# 0002: Second\nStatus: accepted\nSupersedes: ADR-0001\n",
	})
	adrs, err := List(fsys, "decisions")
	if err != nil {
		t.Fatalf("List failed: %v", err)
	}
	if len(adrs) != 2 {
		t.Fatalf("expected 2 ADRs, got %d", len(adrs))
	}
	if adrs[0].SupersededBy != "0002" {
		t.Fatalf("expected 0001.SupersededBy=0002, got %q", adrs[0].SupersededBy)
	}
	if adrs[1].Supersedes != "0001" {
		t.Fatalf("expected 0002.Supersedes=0001, got %q", adrs[1].Supersedes)
	}
}

func TestCheckSkipsSuperseded(t *testing.T) {
	fsys := makeADRFS(map[string]string{
		"decisions/0001-old.md": "# 0001: Old\nStatus: superseded\nSuperseded by: ADR-0002\n",
		"decisions/0002-new.md": "# 0002: New\nStatus: accepted\nSupersedes: ADR-0001\n",
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
		t.Fatalf("expected no findings (superseded ADR skipped), got: %v", findings)
	}
}

func TestCheckBrokenSupersedeChain(t *testing.T) {
	fsys := makeADRFS(map[string]string{
		"decisions/0001-old.md": "# 0001: Old\nStatus: superseded\nSuperseded by: ADR-0099\n",
	})
	findings, err := Check(fsys, "decisions", fsys, ".loop")
	if err != nil {
		t.Fatalf("Check failed: %v", err)
	}
	if len(findings) != 1 {
		t.Fatalf("expected 1 finding, got %d: %v", len(findings), findings)
	}
	if !strings.Contains(findings[0], "broken supersede chain") {
		t.Fatalf("expected broken supersede chain finding, got: %v", findings)
	}
}

func TestCheckOneSidedSupersede(t *testing.T) {
	fsys := makeADRFS(map[string]string{
		"decisions/0001-old.md": "# 0001: Old\nStatus: superseded\nSuperseded by: ADR-0002\n",
		"decisions/0002-new.md": "# 0002: New\nStatus: accepted\n",
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
	if len(findings) != 1 {
		t.Fatalf("expected 1 finding, got %d: %v", len(findings), findings)
	}
	if !strings.Contains(findings[0], "one-sided supersede") {
		t.Fatalf("expected one-sided supersede finding, got: %v", findings)
	}
}

func TestActivePredicate(t *testing.T) {
	cases := []struct {
		status string
		active bool
	}{
		{"", true},
		{"accepted", true},
		{"proposed", true},
		{"superseded", false},
		{"deprecated", false},
		{"rejected", false},
		{"Superseded", false},
	}
	for _, c := range cases {
		adr := ADR{Status: c.status}
		if got := adr.Active(); got != c.active {
			t.Fatalf("Active(%q) = %v, want %v", c.status, got, c.active)
		}
	}
}

func TestActiveHonorsSupersededByLine(t *testing.T) {
	if adr := (ADR{Status: "accepted", SupersededBy: "0002"}); adr.Active() {
		t.Fatalf("accepted ADR with SupersededBy should be inactive")
	}
	if adr := (ADR{Status: "superseded", SupersededBy: ""}); adr.Active() {
		t.Fatalf("superseded ADR should be inactive")
	}
	if adr := (ADR{Status: "accepted", SupersededBy: ""}); !adr.Active() {
		t.Fatalf("accepted ADR with no SupersededBy should be active")
	}
}

func TestCheckSupersededByLineSkipsOrphan(t *testing.T) {
	fsys := makeADRFS(map[string]string{
		"decisions/0001-old.md": "# 0001: Old\nStatus: accepted\nSuperseded by: ADR-0002\n",
		"decisions/0002-new.md": "# 0002: New\nStatus: accepted\nSupersedes: ADR-0001\n",
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
		t.Fatalf("expected no findings (0001 skipped by SupersededBy line), got: %v", findings)
	}
}

func TestCheckOneSidedSupersedeFromNew(t *testing.T) {
	fsys := makeADRFS(map[string]string{
		"decisions/0001-old.md": "# 0001: Old\nStatus: accepted\n",
		"decisions/0002-new.md": "# 0002: New\nStatus: accepted\nSupersedes: ADR-0001\n",
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
		".loop/old-cycle/EVIDENCE.md": lines(
			"## 2026-07-06T12:00:00 — Old unit",
			"",
			"Status: done",
			"",
			"Unit:",
			"````markdown",
			"## Old unit",
			"",
			"Read first:",
			"- decisions/0001-old.md",
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
	if len(findings) != 1 {
		t.Fatalf("expected 1 finding, got %d: %v", len(findings), findings)
	}
	if !strings.Contains(findings[0], "one-sided supersede") {
		t.Fatalf("expected one-sided supersede finding, got: %v", findings)
	}
}

func TestListParsesSupersedeCaseSensitive(t *testing.T) {
	fsys := makeADRFS(map[string]string{
		"decisions/0001-first.md":  "# 0001: First\nStatus: superseded\nSuperseded by: ADR-0002\n",
		"decisions/0002-second.md": "# 0002: Second\nStatus: accepted\nSupersedes: ADR-0001\n",
		"decisions/0003-lower.md":  "# 0003: Lower\nStatus: superseded\nsuperseded by: ADR-0002\n",
	})
	adrs, err := List(fsys, "decisions")
	if err != nil {
		t.Fatalf("List failed: %v", err)
	}
	want := map[string]struct {
		supersededBy string
		supersedes   string
	}{
		"0001": {"0002", ""},
		"0002": {"", "0001"},
		"0003": {"", ""},
	}
	if len(adrs) != len(want) {
		t.Fatalf("expected %d ADRs, got %d", len(want), len(adrs))
	}
	for _, a := range adrs {
		w, ok := want[a.Number]
		if !ok {
			t.Fatalf("unexpected ADR %s", a.Number)
		}
		if a.SupersededBy != w.supersededBy || a.Supersedes != w.supersedes {
			t.Fatalf("ADR %s: SupersededBy=%q Supersedes=%q, want %q/%q", a.Number, a.SupersededBy, a.Supersedes, w.supersededBy, w.supersedes)
		}
	}
}

func TestListParsesSupersedeWithInlineComment(t *testing.T) {
	fsys := makeADRFS(map[string]string{
		"decisions/0001-old.md": lines("# 0001: Old", "Status: superseded", "Superseded by: ADR-0002   # replaced by second"),
		"decisions/0002-new.md": lines("# 0002: New", "Status: accepted", "Supersedes: ADR-0001   # replaces first"),
	})
	adrs, err := List(fsys, "decisions")
	if err != nil {
		t.Fatalf("List failed: %v", err)
	}
	if len(adrs) != 2 {
		t.Fatalf("expected 2 ADRs, got %d", len(adrs))
	}
	if adrs[0].SupersededBy != "0002" {
		t.Fatalf("expected 0001.SupersededBy=0002, got %q", adrs[0].SupersededBy)
	}
	if adrs[1].Supersedes != "0001" {
		t.Fatalf("expected 0002.Supersedes=0001, got %q", adrs[1].Supersedes)
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
