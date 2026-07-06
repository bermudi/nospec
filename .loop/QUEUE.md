# Loop Queue: Go CLI

Goal:
Build the sliceloop CLI in Go — a read-only validator and context provider that also packages the default skill set. The CLI lives in `cli/` as its own Go module, produces a single static binary, and embeds the seven default skills via `go:embed`. It never invokes agents and never runs the loop.

Stop condition:
`cd cli && go test ./...` exits 0, covering all commands: `skills init`, `skills check`, `validate`, `decisions list|show|check`, `status`, `glossary check`, `instructions`.

## Go scaffold and `skills init` scaffolds the seven default skills into a target project

Why:
This is the CLI's one write operation and the user's stated priority — the CLI must package the skills, not just validate. ADR-0002 ruled that the CLI embeds the default skills via `go:embed` and writes them out on `skills init`. This unit also lays the Go project foundation that every subsequent unit builds on.

Work:
- Create `cli/` as a self-contained Go module (`go.mod`, `main.go`).
- Copy `.agents/skills/` into `cli/embedded/skills/` so `go:embed` can reach it. Add a `cli/sync-skills.sh` (or Makefile target) that re-copies from `../.agents/skills/` so the embedded set stays in sync with the canonical source.
- Implement `sliceloop skills init [--target DIR]` — writes `.agents/skills/` into the target directory (default: current dir). Does not overwrite existing skills; only scaffolds missing ones. Prints which skills it wrote and which it skipped.
- Use stdlib `flag` for arg parsing (no external deps). Subcommand dispatch in `main.go`.
- Write tests: init into a temp dir, assert all 7 skill dirs exist with `SKILL.md` inside, assert re-running init skips existing skills.

Verify:
```bash
cd cli && go test ./...
```

Done means:
- `go build` produces a binary that scaffolds all 7 skills into a fresh temp dir.
- Re-running init on the same dir skips existing skills without error.
- `cli/embedded/skills/` is in sync with `../.agents/skills/`.
- No external Go dependencies.

Status: pending

## `skills check` validates that skills follow the agentskills.io spec

Why:
The CLI's core read-only role is structural validation. `skills check` is the first validator — it checks the skills this very project ships, closing the loop on the existing `tests/run.sh` skill validation.

Work:
- Implement `sliceloop skills check [--dir DIR]` — validates every `SKILL.md` under `.agents/skills/` (or the given dir).
- Checks: frontmatter present (YAML between `---` fences), `name` field non-empty, `description` field non-empty, no `[[...]]` or `[...](...)` references that point to nonexistent files.
- Exit 0 if all skills valid, exit 1 if any invalid. Print one line per finding.
- Write tests: valid skill fixture passes, invalid skill fixtures (missing frontmatter, empty description, broken ref) fail with clear messages.

Verify:
```bash
cd cli && go test ./...
```

Done means:
- `skills check` on this project's own `.agents/skills/` exits 0.
- Malformed skills are caught with actionable error messages.
- No regression in `skills init` tests.

Status: pending

## `validate` checks that QUEUE.md work units are well-formed

Why:
The loop (`loop.sh`) parses QUEUE.md but does not validate structure — that's the CLI's job per DESIGN.md. `validate` is the mechanical gate that catches malformed work units before the loop runs them.

Work:
- Implement `sliceloop validate <queue-file>` — parses the QUEUE.md and checks every `## <outcome>` unit has: a `Verify:` section with a fenced code block, and at least one line of outcome text in the header.
- Reports which units pass and which fail, with the specific missing field.
- Exit 0 if all units valid, exit 1 if any invalid.
- Write tests: valid QUEUE fixture passes, invalid fixtures (missing Verify, empty outcome, missing fence) fail.

Verify:
```bash
cd cli && go test ./...
```

Done means:
- `validate` on `examples/smoke/.loop/QUEUE.md` exits 0.
- Malformed queues are caught with unit-level error messages.
- No regression in prior command tests.

Status: pending

## `decisions list|show|check` and `status` provide ADR management and the decision coverage gate

Why:
The decision coverage gate is the user's second stated priority and the one novel CLI feature from DESIGN.md. `status` gives the human a quick snapshot of loop state. This unit delivers both together since they share the ADR/queue parsing layer.

Work:
- `sliceloop decisions list` — lists ADRs in `decisions/` (number, title, status).
- `sliceloop decisions show NNNN` — prints the full ADR file.
- `sliceloop decisions check` — mechanical coverage gate: every ADR in `decisions/` is referenced by at least one work unit in any `QUEUE.md` found under `.loop/`; every ADR referenced by a work unit still exists in `decisions/`. Reports orphaned ADRs and dangling references.
- `sliceloop status` — prints: queue state (pending/done/failed counts from `.loop/QUEUE.md`), evidence entry count (from `.loop/EVIDENCE.md`), ADR count (from `decisions/`).
- Write tests: fixture with ADRs + QUEUE where coverage is clean passes; fixture with orphaned ADR and dangling ref fails with specific messages; `status` reports correct counts.

Verify:
```bash
cd cli && go test ./...
```

Done means:
- `decisions check` on this repo's own `decisions/` + `.loop/QUEUE.md` exits 0 (ADR-0001 and ADR-0002 are both referenced by the CLI queue).
- `decisions list` shows both ADRs.
- `status` reports accurate counts on a fixture.
- No regression in prior command tests.

Status: pending

## `glossary check` and `instructions` round out the CLI command set

Why:
`glossary check` enforces ubiquitous language consistency (terms defined but unused = stale; terms used but undefined = gaps). `instructions` prints templates so the agent doesn't have to memorize work unit / ADR / glossary formats. These are the last two commands in the DESIGN.md spec.

Work:
- `sliceloop glossary check` — parses `glossary.md` for defined terms, scans code + markdown for usage, reports stale definitions (defined but never referenced) and undefined terms (used but not in glossary). Exit 0 if clean, exit 1 if findings.
- `sliceloop instructions <artifact>` — prints the template + guidance for creating a work unit, ADR, or glossary entry. Pure text output. Artifact must be one of: `work-unit`, `adr`, `glossary-entry`. Exit 1 on unknown artifact.
- Write tests: glossary fixture with a stale term and an undefined term produces the right findings; `instructions work-unit` output contains `Verify:` and `Status: pending`; unknown artifact exits 1.

Verify:
```bash
cd cli && go test ./...
```

Done means:
- `glossary check` on a fixture reports stale and undefined terms correctly.
- `instructions work-unit` prints a usable work unit template.
- `instructions adr` and `instructions glossary-entry` print their respective templates.
- All prior command tests still pass.
- The CLI command set is complete per DESIGN.md.

Status: pending
