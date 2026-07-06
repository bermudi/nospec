
## 2026-07-06T09:55:05-06:00 — plan skill writes to `.loop/<name>/QUEUE.md` and skills reference named cycles

Status: verify_failed

Files changed:
```text
 M .agents/skills/build/SKILL.md
 M .agents/skills/explore/SKILL.md
 M .agents/skills/fix/SKILL.md
 M .agents/skills/plan/SKILL.md
 M .agents/skills/review/SKILL.md
 M .loop/named-cycles/QUEUE.md
 M AGENTS.md
 M DESIGN.md
 M cli/embedded/skills/build/SKILL.md
 M cli/embedded/skills/explore/SKILL.md
 M cli/embedded/skills/fix/SKILL.md
 M cli/embedded/skills/plan/SKILL.md
 M cli/embedded/skills/review/SKILL.md
?? .loop/named-cycles/EVIDENCE.md
```

Verify command:
```bash
cd cli && go test ./... && ./tests/run.sh && ! grep -rn '\.loop/QUEUE\.md' .agents/skills/ AGENTS.md DESIGN.md
```

Verify output:
```text
ok  	knack	(cached)
ok  	knack/internal/decisions	(cached)
ok  	knack/internal/glossary	(cached)
ok  	knack/internal/instructions	(cached)
ok  	knack/internal/queue	(cached)
ok  	knack/internal/skills	(cached)
ok  	knack/internal/status	(cached)
bash: line 1: ./tests/run.sh: No such file or directory
```

Worker output:
````text
```text
Unit: plan skill writes to `.loop/<name>/QUEUE.md` and skills reference named cycles
Changed:
  - .agents/skills/plan/SKILL.md — 5 replacements: paths, template, disposability; added name guidance
  - .agents/skills/build/SKILL.md — 2 replacements: description + body
  - .agents/skills/review/SKILL.md — 3 replacements: QUEUE.md + EVIDENCE.md → named
  - .agents/skills/fix/SKILL.md — 5 replacements: paths, new note about appending to same cycle's file
  - .agents/skills/explore/SKILL.md — 1 replacement: handoff + queue references
  - AGENTS.md — 3 replacements: core artifacts, skills section, working conventions
  - DESIGN.md — 17 replacements across diagrams, responsibilities, loop behavior, project layout, disposability, flow, migration
  - cli/embedded/skills/ — synced via sync-skills.sh
Verify expected: cd cli && go test ./... && ./tests/run.sh && ! grep -rn '\.loop/QUEUE\.md' .agents/skills/ AGENTS.md DESIGN.md
Notes: One flat `.loop/HANDOFF.md` remains on AGENTS.md:22 in "Current state" — it describes loop.sh's actual behavior (accurate). Not in the verify scope and not in the work unit's explicit AGENTS.md scope (core artifacts + working conventions only).
```
````

What this proves:
- The work unit is not externally verified.

What remains unverified:
- Anything outside the verify command's proof scope.

## 2026-07-06T09:58:18-06:00 — `knack status` aggregates across all work cycles in `.loop/`

Status: done

Files changed:
```text
 M .agents/skills/build/SKILL.md
 M .agents/skills/explore/SKILL.md
 M .agents/skills/fix/SKILL.md
 M .agents/skills/plan/SKILL.md
 M .agents/skills/review/SKILL.md
 M .loop/named-cycles/QUEUE.md
 M AGENTS.md
 M DESIGN.md
 M cli/embedded/skills/build/SKILL.md
 M cli/embedded/skills/explore/SKILL.md
 M cli/embedded/skills/fix/SKILL.md
 M cli/embedded/skills/plan/SKILL.md
 M cli/embedded/skills/review/SKILL.md
 M cli/internal/status/status.go
 M cli/internal/status/status_test.go
 M cli/main.go
?? .loop/named-cycles/EVIDENCE.md
?? .loop/named-cycles/HANDOFF.md
```

Verify command:
```bash
cd cli && go test ./...
```

Verify output:
```text
ok  	knack	(cached)
ok  	knack/internal/decisions	(cached)
ok  	knack/internal/glossary	(cached)
ok  	knack/internal/instructions	(cached)
ok  	knack/internal/queue	(cached)
ok  	knack/internal/skills	(cached)
ok  	knack/internal/status	(cached)
```

Worker output:
````text
Clean build, no vet issues. All tests pass.

Unit: `knack status` aggregates across all work cycles in `.loop/`
Changed: `cli/internal/status/status.go`, `cli/internal/status/status_test.go`, `cli/main.go` (statusCmd)
Verify expected: `cd cli && go test ./...`
Notes: Report struct changed from flat counts to per-cycle `Cycles []CycleStatus` + `Total CycleStatus` + `ADRs`. Flat `.loop/QUEUE.md` (no subdirectories) is ignored — only subdirectories of `.loop/` are treated as named cycles. ADR count still reads from `decisions/`, independent of cycles.
````

What this proves:
- The verify command passed for this work unit in the current repo state.

What remains unverified:
- Anything outside the verify command's proof scope.
