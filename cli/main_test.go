package main

import (
	"io/fs"
	"testing"
)

func TestEmbeddedSkills(t *testing.T) {
	sub, err := fs.Sub(embeddedSkills, "embedded/skills")
	if err != nil {
		t.Fatalf("failed to open embedded skills directory: %v", err)
	}
	entries, err := fs.ReadDir(sub, ".")
	if err != nil {
		t.Fatalf("failed to read embedded skills: %v", err)
	}
	if len(entries) != 7 {
		t.Fatalf("expected 7 embedded skills, got %d", len(entries))
	}
	for _, e := range entries {
		if !e.IsDir() {
			t.Fatalf("expected %s to be a directory", e.Name())
		}
		if _, err := fs.Stat(sub, e.Name()+"/SKILL.md"); err != nil {
			t.Fatalf("expected %s/SKILL.md: %v", e.Name(), err)
		}
	}
}
