package queue

import (
	"fmt"
	"io/fs"
	"regexp"
	"strings"
)

// Unit represents a single work unit parsed from a QUEUE.md file.
type Unit struct {
	Title string
	Body  string
	Line  int
}

// Status returns the value of the unit's "Status:" line, or an empty string
// if the unit has no such line.
func (u Unit) Status() string {
	for _, line := range strings.Split(u.Body, "\n") {
		if strings.HasPrefix(strings.TrimSpace(line), "Status:") {
			return strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(line), "Status:"))
		}
	}
	return ""
}

// Result reports whether a single work unit is valid and, if not, what is missing.
type Result struct {
	Unit    Unit
	Valid   bool
	Missing string
}

// ValidateFile reads the given QUEUE.md from fsys and validates every work unit.
// It returns a slice of results and any error reading the file. A nil error with
// an empty result slice means the file contains no work units.
func ValidateFile(fsys fs.FS, path string) ([]Result, error) {
	data, err := fs.ReadFile(fsys, path)
	if err != nil {
		return nil, err
	}
	return Validate(string(data)), nil
}

// Validate parses the contents of a QUEUE.md file and validates every work unit.
func Validate(contents string) []Result {
	units := ParseUnits(contents)
	results := make([]Result, 0, len(units))
	for _, u := range units {
		results = append(results, validateUnit(u))
	}
	return results
}

// ParseUnits returns the raw work units parsed from a QUEUE.md file, without
// validating them. This is useful for commands that need to inspect unit bodies
// or status independently of structural validation.
func ParseUnits(contents string) []Unit {
	return parseUnits(contents)
}

// unitHeaderRe matches "## <title>" — the parser additionally excludes
// lines starting with "###" to avoid treating subheadings as work units.
var unitHeaderRe = regexp.MustCompile(`^##\s*(.*)$`)

func isUnitHeader(line string) (string, bool) {
	m := unitHeaderRe.FindStringSubmatch(line)
	if m == nil {
		return "", false
	}
	// Exclude "###", "####", etc. — anything with a third '#'.
	if strings.HasPrefix(line, "###") {
		return "", false
	}
	return strings.TrimSpace(m[1]), true
}

func parseUnits(contents string) []Unit {
	lines := strings.Split(contents, "\n")
	var units []Unit
	var current *Unit

	for i, line := range lines {
		title, ok := isUnitHeader(line)
		if ok {
			if current != nil {
				current.Body = strings.TrimRight(current.Body, "\n")
				units = append(units, *current)
			}
			current = &Unit{
				Title: title,
				Line:  i + 1,
			}
			continue
		}
		if current != nil {
			current.Body += line + "\n"
		}
	}
	if current != nil {
		current.Body = strings.TrimRight(current.Body, "\n")
		units = append(units, *current)
	}
	return units
}

func validateUnit(u Unit) Result {
	if u.Title == "" {
		return Result{Unit: u, Valid: false, Missing: "empty outcome"}
	}

	verifyStart := findVerifyStart(u.Body)
	if verifyStart == -1 {
		return Result{Unit: u, Valid: false, Missing: "Verify section"}
	}

	rest := u.Body[verifyStart:]
	if !containsCodeFence(rest) {
		return Result{Unit: u, Valid: false, Missing: "fenced code block in Verify"}
	}

	return Result{Unit: u, Valid: true}
}

func findVerifyStart(body string) int {
	lines := strings.Split(body, "\n")
	for i, line := range lines {
		if strings.TrimSpace(line) == "Verify:" {
			// Find the start of the line in the original body.
			pos := 0
			for _, l := range lines[:i] {
				pos += len(l) + 1
			}
			return pos
		}
	}
	return -1
}

func containsCodeFence(s string) bool {
	lines := strings.Split(s, "\n")
	for _, line := range lines[1:] {
		if strings.HasPrefix(strings.TrimSpace(line), "```") {
			return true
		}
	}
	return false
}

// Format prints a human-readable result line.
func Format(r Result) string {
	if r.Valid {
		return fmt.Sprintf("PASS: %s (line %d)", r.Unit.Title, r.Unit.Line)
	}
	return fmt.Sprintf("FAIL: %s (line %d): missing %s", r.Unit.Title, r.Unit.Line, r.Missing)
}

// AllValid returns true if every result is valid.
func AllValid(results []Result) bool {
	for _, r := range results {
		if !r.Valid {
			return false
		}
	}
	return true
}
