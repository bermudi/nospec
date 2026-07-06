# Loop Queue: named work cycles

Goal:
Support multiple concurrent work cycles under `.loop/` by giving each cycle its own named subdirectory. The loop already takes a queue path argument and derives evidence/handoff from the queue's directory — the machinery is fine. The work is updating the convention (skills, AGENTS.md, DESIGN.md) and the one CLI command that hardcodes a flat path (`knack status`).

Stop condition:
`cd cli && go test ./...` exits 0, and `./tests/run.sh` exits 0, and `grep -rn '\.loop/QUEUE.md' .agents/skills/ AGENTS.md DESIGN.md` returns no matches (all references updated to `.loop/<name>/QUEUE.md`).

## plan skill writes to `.loop/<name>/QUEUE.md` and skills reference named cycles

Why:
The plan skill currently hardcodes `.loop/QUEUE.md`. Per ADR-0004, each work cycle gets a named subdirectory. The plan, build, review, and fix skills all reference the flat path and need to use `.loop/<name>/QUEUE.md` instead. This is the convention change that makes multiple concurrent cycles possible.

Work:
- Update `.agents/skills/plan/SKILL.md`: change all `.loop/QUEUE.md` references to `.loop/<name>/QUEUE.md`. Add a note that the planner picks a short, descriptive name for the cycle (e.g., `go-cli`, `parser-bug`). Change `.loop/specs/` to `.loop/<name>/specs/`. Update the disposability section to say delete the cycle's subdirectory, not all of `.loop/`.
- Update `.agents/skills/build/SKILL.md`: change `.loop/QUEUE.md` references to `.loop/<name>/QUEUE.md`.
- Update `.agents/skills/review/SKILL.md`: change `.loop/QUEUE.md` and `.loop/EVIDENCE.md` references to `.loop/<name>/QUEUE.md` and `.loop/<name>/EVIDENCE.md`.
- Update `.agents/skills/fix/SKILL.md`: change `.loop/QUEUE.md` references to `.loop/<name>/QUEUE.md`.
- Update `AGENTS.md`: change the core artifacts section and working conventions to reference `.loop/<name>/` instead of flat `.loop/` files.
- Update `DESIGN.md`: change all `.loop/QUEUE.md`, `.loop/EVIDENCE.md`, `.loop/HANDOFF.md`, `.loop/specs/` references to the named-subdirectory form.
- Run `cli/sync-skills.sh` to keep embedded skills in sync.

Verify:
```bash
cd cli && go test ./... && ./tests/run.sh && ! grep -rn '\.loop/QUEUE\.md' .agents/skills/ AGENTS.md DESIGN.md
```

Done means:
- All skill files reference `.loop/<name>/QUEUE.md`, not `.loop/QUEUE.md`.
- AGENTS.md and DESIGN.md reference named subdirectories.
- Embedded skills are in sync with `.agents/skills/`.
- Existing tests still pass.

Status: pending

## `knack status` aggregates across all work cycles in `.loop/`

Why:
`knack status` currently reads `.loop/QUEUE.md` and `.loop/EVIDENCE.md` as flat files. With named work cycles, it needs to walk `.loop/` and aggregate counts across all cycle subdirectories. This is the only CLI code change needed — `decisions check` already walks all of `.loop/`.

Work:
- Update `cli/internal/status/status.go`: instead of reading `.loop/QUEUE.md` directly, walk `.loop/` for subdirectories containing `QUEUE.md`. For each cycle, parse units and count statuses. Aggregate evidence counts across all `EVIDENCE.md` files found.
- Change the `Report` struct to include per-cycle breakdowns (cycle name + counts) in addition to totals.
- Update `cli/main.go` `statusCmd` to print per-cycle lines and a total line.
- Update `cli/internal/status/status_test.go`: test with multiple cycle directories, verify aggregation works. Test with no `.loop/` directory (zero counts). Test with a flat `.loop/QUEUE.md` for backward compat (treat as a cycle named after the parent — or skip if no subdirectories).
- Keep the ADR count as-is (reads `decisions/` which is not per-cycle).

Verify:
```bash
cd cli && go test ./...
```

Done means:
- `knack status` on a repo with multiple `.loop/<name>/` cycles prints per-cycle counts and totals.
- `knack status` on a repo with no `.loop/` prints zero counts.
- Existing `decisions check`, `validate`, `skills` tests still pass.

Status: pending
