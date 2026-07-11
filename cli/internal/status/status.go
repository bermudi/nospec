package status

import (
	"io/fs"
	"path"
	"strings"

	"knack/internal/decisions"
	"knack/internal/queue"
)

// CycleStatus captures per-cycle counts.
type CycleStatus struct {
	Name     string
	Pending  int
	Done     int
	Failed   int
	Evidence int
}

// Report captures the snapshot printed by `knack status`.
type Report struct {
	Cycles     []CycleStatus
	Total      CycleStatus
	ADRs       int
	ActiveADRs int
}

// Generate walks .loop/ for cycle subdirectories, reads their QUEUE.md and
// EVIDENCE.md files, and aggregates counts across all cycles.
func Generate(fsys fs.FS) (Report, error) {
	var r Report

	entries, err := fs.ReadDir(fsys, ".loop")
	if err != nil {
		// No .loop/ — return zero counts with ADRs.
		if adrs, adrErr := decisions.List(fsys, "decisions"); adrErr == nil {
			r.ADRs = len(adrs)
			for _, a := range adrs {
				if a.Active() {
					r.ActiveADRs++
				}
			}
		}
		return r, nil
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		cycleDir := entry.Name()
		queuePath := path.Join(".loop", cycleDir, "QUEUE.md")
		data, err := fs.ReadFile(fsys, queuePath)
		if err != nil {
			continue
		}
		cs := CycleStatus{Name: cycleDir}
		for _, u := range queue.ParseUnits(string(data)) {
			switch u.Status() {
			case "pending":
				cs.Pending++
			case "done":
				cs.Done++
			case "failed":
				cs.Failed++
			}
		}
		evPath := path.Join(".loop", cycleDir, "EVIDENCE.md")
		cs.Evidence = countEvidenceFile(fsys, evPath)

		r.Cycles = append(r.Cycles, cs)
		r.Total.Pending += cs.Pending
		r.Total.Done += cs.Done
		r.Total.Failed += cs.Failed
		r.Total.Evidence += cs.Evidence
	}

	if adrs, err := decisions.List(fsys, "decisions"); err == nil {
		r.ADRs = len(adrs)
		for _, a := range adrs {
			if a.Active() {
				r.ActiveADRs++
			}
		}
	}
	return r, nil
}

func countEvidenceFile(fsys fs.FS, path string) int {
	data, err := fs.ReadFile(fsys, path)
	if err != nil {
		return 0
	}
	count := 0
	for _, line := range strings.Split(string(data), "\n") {
		if strings.HasPrefix(line, "## ") {
			count++
		}
	}
	return count
}
