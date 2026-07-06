package queue

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"testing/fstest"
)

func TestValidateValidQueue(t *testing.T) {
	q := lines(
		"# Queue",
		"",
		"Goal: test.",
		"",
		"## first unit does something",
		"",
		"Why:",
		"Because.",
		"",
		"Read first:",
		"- context",
		"",
		"Constraints:",
		"- boundary",
		"",
		"Verify:",
		"```bash",
		"echo first",
		"```",
		"",
		"Done means:",
		"It echoes.",
		"",
		"Status: pending",
		"",
		"## second unit does more",
		"",
		"Verify:",
		"```bash",
		"echo second",
		"```",
		"",
		"Status: pending",
	)

	results := Validate(q)
	if len(results) != 2 {
		t.Fatalf("expected 2 units, got %d", len(results))
	}
	for _, r := range results {
		if !r.Valid {
			t.Fatalf("expected %q to be valid, got missing %s", r.Unit.Title, r.Missing)
		}
	}
	if !AllValid(results) {
		t.Fatalf("expected AllValid true")
	}
}

func TestValidateMissingVerify(t *testing.T) {
	q := lines(
		"## missing verify unit",
		"",
		"Read first:",
		"- context",
		"",
		"Status: pending",
	)

	results := Validate(q)
	if len(results) != 1 {
		t.Fatalf("expected 1 unit, got %d", len(results))
	}
	if results[0].Valid {
		t.Fatalf("expected unit to fail")
	}
	if results[0].Missing != "Verify section" {
		t.Fatalf("expected missing Verify section, got %q", results[0].Missing)
	}
}

func TestValidateEmptyOutcome(t *testing.T) {
	q := lines(
		"## ",
		"",
		"Verify:",
		"```bash",
		"noop",
		"```",
	)

	results := Validate(q)
	if len(results) != 1 {
		t.Fatalf("expected 1 unit, got %d", len(results))
	}
	if results[0].Valid || results[0].Missing != "empty outcome" {
		t.Fatalf("expected empty outcome failure, got missing %q", results[0].Missing)
	}
}

func TestValidateMissingFence(t *testing.T) {
	q := lines(
		"## missing fence unit",
		"",
		"Verify:",
		"this is not a fence",
	)

	results := Validate(q)
	if len(results) != 1 {
		t.Fatalf("expected 1 unit, got %d", len(results))
	}
	if results[0].Valid || results[0].Missing != "fenced code block in Verify" {
		t.Fatalf("expected fenced code block failure, got missing %q", results[0].Missing)
	}
}

func TestValidateSmokeFixture(t *testing.T) {
	cwd, _ := os.Getwd()
	root := filepath.Join(cwd, "../../../")
	fsys := os.DirFS(root)
	results, err := ValidateFile(fsys, "examples/smoke/.loop/QUEUE.md")
	if err != nil {
		t.Fatalf("ValidateFile failed: %v", err)
	}
	if len(results) != 1 {
		t.Fatalf("expected 1 unit, got %d", len(results))
	}
	if !results[0].Valid {
		t.Fatalf("expected smoke fixture to pass, got missing %s", results[0].Missing)
	}
}

func TestValidateReportsMultipleFailures(t *testing.T) {
	q := lines(
		"## first bad",
		"",
		"Read first: none.",
		"",
		"## second bad",
		"",
		"Verify:",
		"no fence",
	)

	results := Validate(q)
	if len(results) != 2 {
		t.Fatalf("expected 2 units, got %d", len(results))
	}
	if results[0].Valid || results[0].Missing != "Verify section" {
		t.Fatalf("expected first unit missing Verify section, got %v", results[0])
	}
	if results[1].Valid || results[1].Missing != "fenced code block in Verify" {
		t.Fatalf("expected second unit missing fence, got %v", results[1])
	}
	if AllValid(results) {
		t.Fatalf("expected AllValid false")
	}
}

func TestValidateFormat(t *testing.T) {
	pass := Result{Unit: Unit{Title: "ok", Line: 5}, Valid: true}
	fail := Result{Unit: Unit{Title: "bad", Line: 10}, Valid: false, Missing: "Verify section"}
	if got := Format(pass); !strings.Contains(got, "PASS") {
		t.Fatalf("expected PASS in formatted result, got %q", got)
	}
	if got := Format(fail); !strings.Contains(got, "FAIL") || !strings.Contains(got, "missing Verify section") {
		t.Fatalf("expected FAIL and missing reason, got %q", got)
	}
}

func TestValidateFileReadsFromFS(t *testing.T) {
	fsys := fstest.MapFS{
		"QUEUE.md": &fstest.MapFile{Data: []byte(lines(
			"## a unit",
			"",
			"Verify:",
			"```bash",
			"true",
			"```",
		))},
	}
	results, err := ValidateFile(fsys, "QUEUE.md")
	if err != nil {
		t.Fatalf("ValidateFile failed: %v", err)
	}
	if len(results) != 1 || !results[0].Valid {
		t.Fatalf("expected one valid unit, got %v", results)
	}
}

func TestValidateIgnoresSubheadings(t *testing.T) {
	q := lines(
		"## unit with subheadings",
		"",
		"### subheading inside unit",
		"",
		"Some content.",
		"",
		"#### deeper subheading",
		"",
		"Verify:",
		"```bash",
		"true",
		"```",
		"",
		"Status: pending",
	)

	results := Validate(q)
	if len(results) != 1 {
		t.Fatalf("expected 1 unit, got %d (subheadings were parsed as units)", len(results))
	}
	if results[0].Unit.Title != "unit with subheadings" {
		t.Fatalf("expected title %q, got %q", "unit with subheadings", results[0].Unit.Title)
	}
	if !results[0].Valid {
		t.Fatalf("expected unit to be valid, got missing %s", results[0].Missing)
	}
}

func lines(ss ...string) string {
	return strings.Join(ss, "\n")
}


