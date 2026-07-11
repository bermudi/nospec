package instructions

import (
	"strings"
	"testing"
)

func TestInstructionsWorkUnit(t *testing.T) {
	var b strings.Builder
	if err := Print(&b, "work-unit"); err != nil {
		t.Fatalf("Print failed: %v", err)
	}
	out := b.String()
	if !strings.Contains(out, "Read first:") {
		t.Fatalf("expected Read first: in work-unit template, got:\n%s", out)
	}
	if !strings.Contains(out, "Constraints:") {
		t.Fatalf("expected Constraints: in work-unit template, got:\n%s", out)
	}
	if !strings.Contains(out, "Done means:") {
		t.Fatalf("expected Done means: in work-unit template, got:\n%s", out)
	}
	if !strings.Contains(out, "Verify:") {
		t.Fatalf("expected Verify: in work-unit template, got:\n%s", out)
	}
	if strings.Contains(out, "Work:") {
		t.Fatalf("expected no Work: field in work-unit template, got:\n%s", out)
	}
	if !strings.Contains(out, "Status: pending") {
		t.Fatalf("expected Status: pending in work-unit template, got:\n%s", out)
	}
}

func TestInstructionsAdr(t *testing.T) {
	var b strings.Builder
	if err := Print(&b, "adr"); err != nil {
		t.Fatalf("Print failed: %v", err)
	}
	out := b.String()
	if !strings.Contains(out, "# NNNN: <title>") {
		t.Fatalf("expected ADR heading in adr template, got:\n%s", out)
	}
	if !strings.Contains(out, "Status: accepted") {
		t.Fatalf("expected Status: accepted in adr template, got:\n%s", out)
	}
	if !strings.Contains(out, "# Supersedes: ADR-NNNN") {
		t.Fatalf("expected commented Supersedes field in adr template, got:\n%s", out)
	}
	if !strings.Contains(out, "# Superseded by: ADR-NNNN") {
		t.Fatalf("expected commented Superseded by field in adr template, got:\n%s", out)
	}
	if strings.Contains(out, "\nSupersedes: ADR-NNNN") || strings.Contains(out, "\nSuperseded by: ADR-NNNN") {
		t.Fatalf("ADR template has Supersedes/Superseded by as active fields, got:\n%s", out)
	}
}

func TestInstructionsGlossaryEntry(t *testing.T) {
	var b strings.Builder
	if err := Print(&b, "glossary-entry"); err != nil {
		t.Fatalf("Print failed: %v", err)
	}
	out := b.String()
	if !strings.Contains(out, "## <term>") {
		t.Fatalf("expected glossary entry heading in template, got:\n%s", out)
	}
	if !strings.Contains(out, "Used in:") {
		t.Fatalf("expected Used in: in glossary entry template, got:\n%s", out)
	}
}

func TestInstructionsUnknown(t *testing.T) {
	var b strings.Builder
	err := Print(&b, "nope")
	if err == nil {
		t.Fatalf("expected error for unknown artifact")
	}
}
