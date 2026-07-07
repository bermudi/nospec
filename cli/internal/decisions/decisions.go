package decisions

import (
	"errors"
	"fmt"
	"io/fs"
	"path"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"knack/internal/queue"
)

// ADR represents a single architectural decision record in decisions/.
type ADR struct {
	Number   string
	Title    string
	Status   string
	Filename string
}

var (
	adrFilenameRe = regexp.MustCompile(`^(\d{4})-(.+)\.md$`)
	adrTitleRe    = regexp.MustCompile(`^#\s*(\d{4}):\s*(.*)$`)
	adrRefRe      = regexp.MustCompile(`ADR-(\d{1,4})`)
)

// List returns all ADRs in the given directory, sorted by number.
func List(fsys fs.FS, dir string) ([]ADR, error) {
	entries, err := fs.ReadDir(fsys, dir)
	if err != nil {
		return nil, err
	}
	var adrs []ADR
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		m := adrFilenameRe.FindStringSubmatch(e.Name())
		if m == nil {
			continue
		}
		data, err := fs.ReadFile(fsys, path.Join(dir, e.Name()))
		if err != nil {
			return nil, err
		}
		number, title := parseADRTitle(string(data), m[1])
		adrs = append(adrs, ADR{
			Number:   number,
			Title:    title,
			Status:   parseADRStatus(string(data)),
			Filename: e.Name(),
		})
	}
	sort.Slice(adrs, func(i, j int) bool { return adrs[i].Number < adrs[j].Number })
	return adrs, nil
}

func parseADRTitle(contents, fallbackNumber string) (string, string) {
	for _, line := range strings.Split(contents, "\n") {
		m := adrTitleRe.FindStringSubmatch(line)
		if m == nil {
			continue
		}
		return m[1], strings.TrimSpace(m[2])
	}
	return fallbackNumber, ""
}

func parseADRStatus(contents string) string {
	for _, line := range strings.Split(contents, "\n") {
		if strings.HasPrefix(strings.TrimSpace(line), "Status:") {
			return strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(line), "Status:"))
		}
	}
	return ""
}

// Show returns the full contents of the ADR identified by number.
// The number may be given with or without leading zeros (e.g., "1" or "0001").
func Show(fsys fs.FS, dir, number string) ([]byte, error) {
	canonical, err := canonicalADRNumber(number)
	if err != nil {
		return nil, err
	}
	entries, err := fs.ReadDir(fsys, dir)
	if err != nil {
		return nil, err
	}
	prefix := canonical + "-"
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		name := e.Name()
		if strings.HasPrefix(name, prefix) && strings.HasSuffix(name, ".md") {
			return fs.ReadFile(fsys, path.Join(dir, name))
		}
	}
	return nil, fmt.Errorf("ADR %s not found", number)
}

func canonicalADRNumber(s string) (string, error) {
	if !regexp.MustCompile(`^\d{1,4}$`).MatchString(s) {
		return "", fmt.Errorf("invalid ADR number %q", s)
	}
	n, _ := strconv.Atoi(s)
	return fmt.Sprintf("%04d", n), nil
}

// Check performs the mechanical decision coverage gate.
// It reports ADRs in decisionsDir that are not referenced by any work unit or
// evidence ledger, and ADR references in work units that do not resolve to an
// existing ADR.
//
// An ADR is "orphaned" if it is not referenced by any QUEUE.md (current work)
// or any EVIDENCE.md (completed work). This avoids false positives when a
// completed cycle's QUEUE.md has been deleted but its EVIDENCE.md ledger
// remains as the durable record.
func Check(fsys fs.FS, decisionsDir string, loopFS fs.FS, loopDir string) ([]string, error) {
	adrs, err := List(fsys, decisionsDir)
	if err != nil {
		return nil, err
	}
	if len(adrs) == 0 {
		return nil, nil
	}

	bodies, err := collectLoopBodies(loopFS, loopDir)
	if err != nil {
		return nil, err
	}
	allBody := strings.Join(bodies, "\n")

	var findings []string
	for _, adr := range adrs {
		if !adrReferenced(adr, allBody) {
			findings = append(findings, fmt.Sprintf("orphaned ADR %s (%s): not referenced by any work unit or evidence ledger", adr.Number, adr.Filename))
		}
	}

	for _, body := range bodies {
		for _, m := range adrRefRe.FindAllStringSubmatch(body, -1) {
			num, _ := canonicalADRNumber(m[1])
			if !adrExists(adrs, num) {
				findings = append(findings, fmt.Sprintf("dangling reference ADR-%s: no matching ADR in %s", num, decisionsDir))
			}
		}
	}

	sort.Strings(findings)
	return findings, nil
}

// collectLoopBodies gathers text from QUEUE.md and EVIDENCE.md files under
// loopDir. QUEUE.md bodies are parsed into work units so that only unit content
// (not queue-level metadata) is checked. EVIDENCE.md files are included whole —
// the ledger is the durable record of completed work and may contain ADR
// references from unit bodies that the loop wrote into it.
func collectLoopBodies(loopFS fs.FS, loopDir string) ([]string, error) {
	var bodies []string
	err := fs.WalkDir(loopFS, loopDir, func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		base := path.Base(p)
		if base != "QUEUE.md" && base != "EVIDENCE.md" {
			return nil
		}
		data, err := fs.ReadFile(loopFS, p)
		if err != nil {
			return err
		}
		if base == "QUEUE.md" {
			for _, u := range queue.ParseUnits(string(data)) {
				bodies = append(bodies, u.Body)
			}
		} else {
			bodies = append(bodies, string(data))
		}
		return nil
	})
	if err != nil && !missingDir(err) {
		return nil, err
	}
	return bodies, nil
}

func missingDir(err error) bool {
	if err == nil {
		return false
	}
	return errors.Is(err, fs.ErrNotExist) || strings.Contains(err.Error(), "file does not exist")
}

func adrReferenced(adr ADR, body string) bool {
	base := strings.TrimSuffix(adr.Filename, ".md")
	candidates := []string{
		"ADR-" + adr.Number,
		adr.Filename,
		base,
		path.Join("decisions", adr.Filename),
		path.Join("decisions", base),
	}
	for _, c := range candidates {
		if strings.Contains(body, c) {
			return true
		}
	}
	return false
}

func adrExists(adrs []ADR, number string) bool {
	for _, adr := range adrs {
		if adr.Number == number {
			return true
		}
	}
	return false
}
