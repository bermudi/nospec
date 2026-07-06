package status

import (
	"io/fs"
	"strings"

	"knack/internal/decisions"
	"knack/internal/queue"
)

// Report captures the snapshot printed by `knack status`.
type Report struct {
	Pending  int
	Done     int
	Failed   int
	Evidence int
	ADRs     int
}

// Generate reads the queue, evidence, and decisions in the given filesystem and
// returns the counts for `knack status`.
func Generate(fsys fs.FS) (Report, error) {
	var r Report
	data, err := fs.ReadFile(fsys, ".loop/QUEUE.md")
	if err == nil {
		for _, u := range queue.ParseUnits(string(data)) {
			switch u.Status() {
			case "pending":
				r.Pending++
			case "done":
				r.Done++
			case "failed":
				r.Failed++
			}
		}
	}
	r.Evidence = countEvidence(fsys)
	adrs, err := decisions.List(fsys, "decisions")
	if err == nil {
		r.ADRs = len(adrs)
	}
	return r, nil
}

func countEvidence(fsys fs.FS) int {
	data, err := fs.ReadFile(fsys, ".loop/EVIDENCE.md")
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
