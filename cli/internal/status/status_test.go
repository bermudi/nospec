package status

import (
	"io/fs"
	"testing"
	"testing/fstest"
)

func TestStatusCounts(t *testing.T) {
	fsys := makeStatusFS(map[string]string{
		".loop/my-cycle/QUEUE.md": lines(
			"## Unit A",
			"",
			"Verify:",
			"```bash",
			"true",
			"```",
			"",
			"Status: pending",
			"",
			"## Unit B",
			"",
			"Verify:",
			"```bash",
			"true",
			"```",
			"",
			"Status: done",
			"",
			"## Unit C",
			"",
			"Verify:",
			"```bash",
			"true",
			"```",
			"",
			"Status: failed",
		),
		".loop/my-cycle/EVIDENCE.md": lines(
			"## 2026-07-06T00:00:00Z — Unit A",
			"Status: done",
			"",
			"## 2026-07-06T00:00:00Z — Unit B",
			"Status: done",
		),
		"decisions/0001-first.md":  "# 0001: First\nStatus: accepted\n",
		"decisions/0002-second.md": "# 0002: Second\nStatus: accepted\n",
	})
	r, err := Generate(fsys)
	if err != nil {
		t.Fatalf("Generate failed: %v", err)
	}
	if len(r.Cycles) != 1 {
		t.Fatalf("expected 1 cycle, got %d", len(r.Cycles))
	}
	c := r.Cycles[0]
	if c.Name != "my-cycle" {
		t.Fatalf("expected cycle name 'my-cycle', got %q", c.Name)
	}
	if c.Pending != 1 || c.Done != 1 || c.Failed != 1 {
		t.Fatalf("unexpected cycle counts: %+v", c)
	}
	if c.Evidence != 2 {
		t.Fatalf("expected evidence count 2, got %d", c.Evidence)
	}
	if r.Total.Pending != 1 || r.Total.Done != 1 || r.Total.Failed != 1 || r.Total.Evidence != 2 {
		t.Fatalf("unexpected total counts: %+v", r.Total)
	}
	if r.ADRs != 2 {
		t.Fatalf("expected ADR count 2, got %d", r.ADRs)
	}
	if r.ActiveADRs != 2 {
		t.Fatalf("expected active ADR count 2, got %d", r.ActiveADRs)
	}
}

func TestStatusMissingLoop(t *testing.T) {
	fsys := makeStatusFS(map[string]string{})
	r, err := Generate(fsys)
	if err != nil {
		t.Fatalf("Generate failed: %v", err)
	}
	if len(r.Cycles) != 0 {
		t.Fatalf("expected 0 cycles, got %d", len(r.Cycles))
	}
	if r.Total.Pending != 0 || r.Total.Done != 0 || r.Total.Failed != 0 || r.Total.Evidence != 0 {
		t.Fatalf("expected zero total counts, got %+v", r.Total)
	}
	if r.ADRs != 0 {
		t.Fatalf("expected 0 ADRs, got %d", r.ADRs)
	}
	if r.ActiveADRs != 0 {
		t.Fatalf("expected 0 active ADRs, got %d", r.ActiveADRs)
	}
}

func TestStatusMultipleCycles(t *testing.T) {
	fsys := makeStatusFS(map[string]string{
		".loop/cycle-a/QUEUE.md": lines(
			"## Unit A1",
			"",
			"Verify:",
			"```bash",
			"true",
			"```",
			"",
			"Status: pending",
			"",
			"## Unit A2",
			"",
			"Verify:",
			"```bash",
			"true",
			"```",
			"",
			"Status: done",
		),
		".loop/cycle-a/EVIDENCE.md": lines(
			"## Entry 1",
			"Status: done",
		),
		".loop/cycle-b/QUEUE.md": lines(
			"## Unit B1",
			"",
			"Verify:",
			"```bash",
			"true",
			"```",
			"",
			"Status: failed",
		),
		".loop/cycle-b/EVIDENCE.md": lines(
			"## Entry 1",
			"Status: done",
			"",
			"## Entry 2",
			"Status: done",
		),
		"decisions/0001-one.md": "# 0001: One\nStatus: accepted\n",
	})
	r, err := Generate(fsys)
	if err != nil {
		t.Fatalf("Generate failed: %v", err)
	}
	if len(r.Cycles) != 2 {
		t.Fatalf("expected 2 cycles, got %d", len(r.Cycles))
	}
	ca := r.Cycles[0]
	if ca.Name != "cycle-a" || ca.Pending != 1 || ca.Done != 1 || ca.Failed != 0 || ca.Evidence != 1 {
		t.Fatalf("unexpected cycle-a: %+v", ca)
	}
	cb := r.Cycles[1]
	if cb.Name != "cycle-b" || cb.Pending != 0 || cb.Done != 0 || cb.Failed != 1 || cb.Evidence != 2 {
		t.Fatalf("unexpected cycle-b: %+v", cb)
	}
	if r.Total.Pending != 1 || r.Total.Done != 1 || r.Total.Failed != 1 || r.Total.Evidence != 3 {
		t.Fatalf("unexpected totals: %+v", r.Total)
	}
	if r.ADRs != 1 {
		t.Fatalf("expected ADR count 1, got %d", r.ADRs)
	}
	if r.ActiveADRs != 1 {
		t.Fatalf("expected active ADR count 1, got %d", r.ActiveADRs)
	}
}

func TestStatusFlatQueue(t *testing.T) {
	// Flat .loop/QUEUE.md without subdirectories is ignored — only
	// subdirectories of .loop/ are treated as named cycles.
	fsys := makeStatusFS(map[string]string{
		".loop/QUEUE.md": lines(
			"## Old unit",
			"",
			"Verify:",
			"```bash",
			"true",
			"```",
			"",
			"Status: done",
		),
		".loop/EVIDENCE.md": lines(
			"## Old evidence",
			"Status: done",
		),
		"decisions/0001-one.md": "# 0001: One\nStatus: accepted\n",
	})
	r, err := Generate(fsys)
	if err != nil {
		t.Fatalf("Generate failed: %v", err)
	}
	if len(r.Cycles) != 0 {
		t.Fatalf("expected 0 cycles for flat .loop/, got %d", len(r.Cycles))
	}
	if r.Total.Pending != 0 || r.Total.Done != 0 || r.Total.Failed != 0 || r.Total.Evidence != 0 {
		t.Fatalf("expected zero total counts, got %+v", r.Total)
	}
	if r.ADRs != 1 {
		t.Fatalf("expected ADR count 1, got %d", r.ADRs)
	}
	if r.ActiveADRs != 1 {
		t.Fatalf("expected active ADR count 1, got %d", r.ActiveADRs)
	}
}

func TestStatusLoopWithNoSubdirs(t *testing.T) {
	// .loop/ exists but contains no directories — only flat files.
	fsys := makeStatusFS(map[string]string{
		".loop/QUEUE.md": lines(
			"## old unit",
			"",
			"Verify:",
			"```bash",
			"true",
			"```",
			"",
			"Status: done",
		),
		"decisions/0001-one.md": "# 0001: One\nStatus: accepted\n",
	})
	r, err := Generate(fsys)
	if err != nil {
		t.Fatalf("Generate failed: %v", err)
	}
	if len(r.Cycles) != 0 {
		t.Fatalf("expected 0 cycles, got %d", len(r.Cycles))
	}
	if r.ADRs != 1 {
		t.Fatalf("expected ADR count 1, got %d", r.ADRs)
	}
	if r.ActiveADRs != 1 {
		t.Fatalf("expected active ADR count 1, got %d", r.ActiveADRs)
	}
}

func TestStatusMissingLoopWithDecisions(t *testing.T) {
	fsys := makeStatusFS(map[string]string{
		"decisions/0001-one.md":   "# 0001: One\nStatus: accepted\n",
		"decisions/0002-two.md":   "# 0002: Two\nStatus: superseded\nSuperseded by: ADR-0001\n",
		"decisions/0003-three.md": "# 0003: Three\nStatus: accepted\n",
	})
	r, err := Generate(fsys)
	if err != nil {
		t.Fatalf("Generate failed: %v", err)
	}
	if len(r.Cycles) != 0 {
		t.Fatalf("expected 0 cycles, got %d", len(r.Cycles))
	}
	if r.ADRs != 3 {
		t.Fatalf("expected 3 ADRs, got %d", r.ADRs)
	}
	if r.ActiveADRs != 2 {
		t.Fatalf("expected 2 active ADRs, got %d", r.ActiveADRs)
	}
}

func makeStatusFS(files map[string]string) fs.FS {
	fsys := fstest.MapFS{}
	for name, data := range files {
		fsys[name] = &fstest.MapFile{Data: []byte(data)}
	}
	return fsys
}

func lines(ss ...string) string {
	out := ""
	for i, s := range ss {
		if i > 0 {
			out += "\n"
		}
		out += s
	}
	return out
}
