package glossary

import (
	"io/fs"
	"strings"
	"testing"
	"testing/fstest"
)

func TestGlossaryCheckStaleAndUndefined(t *testing.T) {
	fsys := makeGlossaryFS(map[string]string{
		"glossary.md": lines(
			"# Glossary",
			"",
			"## ActiveTerm",
			"A term that is used in the project.",
			"",
			"## StaleTerm",
			"A term that is defined but never used.",
		),
		"docs.md": lines(
			"This document uses [[ActiveTerm]] and [[UndefinedTerm]].",
			"It also mentions ActiveTerm in plain text.",
		),
		"code.go": lines(
			"package main",
			"",
			"// ActiveTerm is important.",
			"func main() {}",
		),
	})

	findings, err := Check(fsys, "glossary.md")
	if err != nil {
		t.Fatalf("Check failed: %v", err)
	}
	if len(findings) != 2 {
		t.Fatalf("expected 2 findings, got %d: %v", len(findings), findings)
	}

	var hasStale, hasUndefined bool
	for _, f := range findings {
		if f.Kind == "stale" && f.Term == "StaleTerm" {
			hasStale = true
		}
		if f.Kind == "undefined" && f.Term == "UndefinedTerm" {
			hasUndefined = true
		}
	}
	if !hasStale {
		t.Fatalf("expected stale StaleTerm finding, got: %v", findings)
	}
	if !hasUndefined {
		t.Fatalf("expected undefined UndefinedTerm finding, got: %v", findings)
	}
}

func TestGlossaryCheckClean(t *testing.T) {
	fsys := makeGlossaryFS(map[string]string{
		"glossary.md": lines(
			"# Glossary",
			"",
			"## ActiveTerm",
			"A term that is used.",
		),
		"docs.md": "See [[ActiveTerm]] for details.\n",
	})

	findings, err := Check(fsys, "glossary.md")
	if err != nil {
		t.Fatalf("Check failed: %v", err)
	}
	if len(findings) != 0 {
		t.Fatalf("expected clean glossary, got: %v", findings)
	}
}

func TestGlossaryCheckCaseInsensitive(t *testing.T) {
	fsys := makeGlossaryFS(map[string]string{
		"glossary.md": lines(
			"# Glossary",
			"",
			"## ActiveTerm",
			"A term that is used.",
		),
		"docs.md": "activeterm is used in lowercase.\n",
	})

	findings, err := Check(fsys, "glossary.md")
	if err != nil {
		t.Fatalf("Check failed: %v", err)
	}
	for _, f := range findings {
		if f.Term == "ActiveTerm" {
			t.Fatalf("expected ActiveTerm to be considered used, got: %v", findings)
		}
	}
}

func TestGlossaryCheckSkipsBinary(t *testing.T) {
	fsys := makeGlossaryFS(map[string]string{
		"glossary.md": lines(
			"# Glossary",
			"",
			"## ActiveTerm",
			"A term.",
		),
		"binary.dat": "\x00ActiveTerm\x00",
	})

	findings, err := Check(fsys, "glossary.md")
	if err != nil {
		t.Fatalf("Check failed: %v", err)
	}
	var activeTermStale bool
	for _, f := range findings {
		if f.Kind == "stale" && f.Term == "ActiveTerm" {
			activeTermStale = true
		}
	}
	if !activeTermStale {
		t.Fatalf("expected ActiveTerm to be stale because binary content is ignored, got: %v", findings)
	}
}

func TestGlossaryCheckWholeWord(t *testing.T) {
	fsys := makeGlossaryFS(map[string]string{
		"glossary.md": lines(
			"# Glossary",
			"",
			"## ActiveTerm",
			"A term.",
		),
		"docs.md": "ActiveTermExtra is not a reference to ActiveTerm.\n",
	})

	findings, err := Check(fsys, "glossary.md")
	if err != nil {
		t.Fatalf("Check failed: %v", err)
	}
	for _, f := range findings {
		if f.Term == "ActiveTerm" && f.Kind == "stale" {
			t.Fatalf("expected ActiveTerm to be found as whole word, got: %v", findings)
		}
	}
}

func TestGlossaryParse(t *testing.T) {
	g := Parse(lines(
		"# Glossary",
		"",
		"## One Term",
		"First.",
		"",
		"## Another Term",
		"Second.",
	))
	if len(g.Terms) != 2 {
		t.Fatalf("expected 2 terms, got %d", len(g.Terms))
	}
	if !g.Has("one term") {
		t.Fatalf("expected One Term to be parsed")
	}
	if !g.Has("another term") {
		t.Fatalf("expected Another Term to be parsed")
	}
}

func TestGlossaryFindingString(t *testing.T) {
	f := Finding{Kind: "stale", Term: "StaleTerm"}
	if !strings.Contains(f.String(), "stale") || !strings.Contains(f.String(), "StaleTerm") {
		t.Fatalf("unexpected finding string: %q", f.String())
	}
}

func makeGlossaryFS(files map[string]string) fs.FS {
	fsys := fstest.MapFS{}
	for name, data := range files {
		fsys[name] = &fstest.MapFile{Data: []byte(data)}
	}
	return fsys
}

func lines(ss ...string) string {
	return strings.Join(ss, "\n")
}
