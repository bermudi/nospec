package status

import (
	"io/fs"
	"testing"
	"testing/fstest"
)

func TestStatusCounts(t *testing.T) {
	fsys := makeStatusFS(map[string]string{
		".loop/QUEUE.md": lines(
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
		".loop/EVIDENCE.md": lines(
			"## 2026-07-06T00:00:00Z — Unit A",
			"Status: done",
			"",
			"## 2026-07-06T00:00:00Z — Unit B",
			"Status: done",
		),
		"decisions/0001-first.md": "# 0001: First\nStatus: accepted\n",
		"decisions/0002-second.md": "# 0002: Second\nStatus: accepted\n",
	})
	r, err := Generate(fsys)
	if err != nil {
		t.Fatalf("Generate failed: %v", err)
	}
	if r.Pending != 1 || r.Done != 1 || r.Failed != 1 {
		t.Fatalf("unexpected queue counts: %+v", r)
	}
	if r.Evidence != 2 {
		t.Fatalf("expected evidence count 2, got %d", r.Evidence)
	}
	if r.ADRs != 2 {
		t.Fatalf("expected ADR count 2, got %d", r.ADRs)
	}
}

func TestStatusMissingLoop(t *testing.T) {
	fsys := makeStatusFS(map[string]string{})
	r, err := Generate(fsys)
	if err != nil {
		t.Fatalf("Generate failed: %v", err)
	}
	if r.Pending != 0 || r.Done != 0 || r.Failed != 0 || r.Evidence != 0 || r.ADRs != 0 {
		t.Fatalf("expected zero counts for empty repo, got %+v", r)
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
