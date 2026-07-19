---
id: 0004
date: 2026-07-06
status: accepted
spine: false
grandfathered: "enacted before the ADR-citation convention (ADR-0005); the named-cycles cycle that implemented it predates `Read first:` ADR citations, so its EVIDENCE.md ledger never references this ADR by number. The work is verified in code and the surviving skills/docs."
---

# 0004: Named work cycles under .loop/

## Context

The original design put all loop state in flat files at the root of `.loop/`: `QUEUE.md`, `EVIDENCE.md`, `HANDOFF.md`. This assumes one active work cycle at a time. In practice, work is interleaved — a bug fix lands while a feature is mid-loop, a refactor interrupts a build cycle.

With flat files, starting new work means either finishing or destroying the current cycle's state. `EVIDENCE.md` is append-only and never cleared, so it accumulates across unrelated runs. `decisions check` walks all of `.loop/` for `QUEUE.md` files, so stale queues produce false dangling-reference findings.

The loop itself already takes a queue path argument (`./loop.sh run <queue>`) and derives evidence/handoff paths from the queue's directory. The machinery supports separation; the convention doesn't.

Alternatives considered:
- **Single queue, pause/resume.** Keep flat files, add explicit pause/resume commands. More complex, still one cycle at a time, and evidence still accumulates.
- **Git branches per work cycle.** Overkill — the loop is not a version control system. Branches solve code isolation, not coordination-state isolation.
- **Named subdirectories under `.loop/`.** Each work cycle gets its own directory. The loop already supports this via the queue path argument. Only convention and two CLI commands need to change.

## Decision

Work cycles live in named subdirectories under `.loop/`. The directory name is the work cycle's name (e.g., `go-cli`, `parser-bug`, `refactor-skills`). Each directory contains its own `QUEUE.md`, `EVIDENCE.md`, and `HANDOFF.md`.

```
.loop/
  go-cli/
    QUEUE.md
    EVIDENCE.md
    HANDOFF.md
  urgent-bug/
    QUEUE.md
    EVIDENCE.md
    HANDOFF.md
```

The `plan` skill writes to `.loop/<name>/QUEUE.md` instead of `.loop/QUEUE.md`. The loop is invoked as `./loop.sh run .loop/<name>/QUEUE.md`. Evidence and handoff are scoped to the cycle's directory automatically (the loop already derives them from the queue path).

`knack status` aggregates across all work cycles in `.loop/` — it reports per-cycle counts and a total. `knack decisions check` already walks all of `.loop/` for `QUEUE.md` files, so it works unchanged. *(These commands were removed in [ADR-0011](0011-ship-as-skills-via-skills-sh-delete-cli.md); the status/coverage behavior they described is now loop behavior and self-check concepts in the skills. The named-cycles decision itself is unchanged.)*

When a work cycle is complete and verified, its subdirectory is deleted. The human deletes it; the loop does not.

## Consequences

- Multiple work cycles can coexist without interfering. Starting new work doesn't require finishing or destroying current work.
- Evidence is scoped per cycle — no accumulation across unrelated runs.
- `decisions check` sees only active queues. Completed cycles are deleted, so no stale false positives.
- The `plan` skill and `knack status` need updates. The loop itself needs none.
- The AGENTS.md convention changes from `.loop/QUEUE.md` to `.loop/<name>/QUEUE.md`.
- The plan skill needs a name parameter — the planner picks a short, descriptive name for the cycle.
