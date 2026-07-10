package glossary

import (
	"errors"
	"fmt"
	"io/fs"
	"path"
	"regexp"
	"sort"
	"strings"
	"unicode"
)

// Glossary holds the set of terms defined in a glossary file.
type Glossary struct {
	// Terms maps the normalized term to the original term text.
	Terms map[string]string
}

// Parse reads glossary markdown and extracts terms defined by ## headers.
func Parse(data string) *Glossary {
	g := &Glossary{Terms: make(map[string]string)}
	for _, line := range strings.Split(data, "\n") {
		if !strings.HasPrefix(line, "## ") {
			continue
		}
		term := strings.TrimSpace(strings.TrimPrefix(line, "## "))
		if term == "" {
			continue
		}
		g.Terms[normalize(term)] = term
	}
	return g
}

func normalize(s string) string {
	return strings.ToLower(s)
}

// Has reports whether a normalized term is defined in the glossary.
func (g *Glossary) Has(term string) bool {
	_, ok := g.Terms[term]
	return ok
}

// Finding represents a stale or undefined glossary term.
type Finding struct {
	Kind string // "stale" or "undefined"
	Term string
}

func (f Finding) String() string {
	return fmt.Sprintf("%s: %s", f.Kind, f.Term)
}

// Check reads glossary.md at glossaryPath from fsys and walks all other files
// in fsys to report stale (defined but never used) and undefined (referenced
// via [[...]] but not defined) terms.
// If the glossary file is absent, Check reports no findings and no error: a
// project without a glossary has nothing to check.
func Check(fsys fs.FS, glossaryPath string) ([]Finding, error) {
	data, err := fs.ReadFile(fsys, glossaryPath)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return nil, nil
		}
		return nil, fmt.Errorf("read glossary: %w", err)
	}
	g := Parse(string(data))

	var findings []Finding

	stale, err := findStale(fsys, glossaryPath, g)
	if err != nil {
		return nil, err
	}
	for _, term := range stale {
		findings = append(findings, Finding{Kind: "stale", Term: term})
	}

	undefined, err := findUndefined(fsys, glossaryPath, g)
	if err != nil {
		return nil, err
	}
	for _, term := range undefined {
		findings = append(findings, Finding{Kind: "undefined", Term: term})
	}

	return findings, nil
}

func findStale(fsys fs.FS, glossaryPath string, g *Glossary) ([]string, error) {
	used := make(map[string]bool)
	err := fs.WalkDir(fsys, ".", func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			if p == ".loop" {
				return fs.SkipDir
			}
			return nil
		}
		if p == glossaryPath {
			return nil
		}
		data, err := fs.ReadFile(fsys, p)
		if err != nil {
			return err
		}
		if isBinary(data) {
			return nil
		}
		s := string(data)
		for _, term := range g.Terms {
			if termUsed(term, s) {
				used[normalize(term)] = true
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	var stale []string
	for norm, term := range g.Terms {
		if !used[norm] {
			stale = append(stale, term)
		}
	}
	sort.Strings(stale)
	return stale, nil
}

func termUsed(term, s string) bool {
	lowerTerm := strings.ToLower(term)
	lowerS := strings.ToLower(s)
	start := 0
	for {
		idx := strings.Index(lowerS[start:], lowerTerm)
		if idx == -1 {
			break
		}
		idx += start

		before := idx == 0 || !isWordChar(rune(s[idx-1]))
		after := idx+len(term) >= len(s) || !isWordChar(rune(s[idx+len(term)]))
		if before && after {
			return true
		}
		start = idx + 1
	}
	return false
}

func isWordChar(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r) || r == '-' || r == '_' || r == '/'
}

// isMarkdown reports whether p is a markdown file. The [[...]] wiki-ref
// convention is markdown-only, so undefined-term detection scans markdown and
// ignores the [[...]] that appears in shell test syntax and Go source.
func isMarkdown(p string) bool {
	switch strings.ToLower(path.Ext(p)) {
	case ".md", ".markdown":
		return true
	}
	return false
}

func isBinary(data []byte) bool {
	for _, b := range data {
		if b == 0 {
			return true
		}
	}
	return false
}

var wikiRefRe = regexp.MustCompile(`\[\[([^|\]]+)(?:\|[^\]]+)?\]\]`)

func findUndefined(fsys fs.FS, glossaryPath string, g *Glossary) ([]string, error) {
	seen := make(map[string]bool)
	var undefined []string
	err := fs.WalkDir(fsys, ".", func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			if p == ".loop" {
				return fs.SkipDir
			}
			return nil
		}
		if p == glossaryPath {
			return nil
		}
		if !isMarkdown(p) {
			return nil
		}
		data, err := fs.ReadFile(fsys, p)
		if err != nil {
			return err
		}
		if isBinary(data) {
			return nil
		}
		for _, m := range wikiRefRe.FindAllStringSubmatch(string(data), -1) {
			target := normalize(m[1])
			if !g.Has(target) && !seen[target] {
				seen[target] = true
				undefined = append(undefined, m[1])
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	sort.Strings(undefined)
	return undefined, nil
}
