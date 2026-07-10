# Proposal: Review-fix subloop orchestration

**Status:** disposable planning artifact for the `review-loop` cycle.

## Problem

`knack` today runs the `build` phase well:

```text
pending unit → build worker → Verify: → done
```

But the default skill flow is `explore → plan → build → review → fix`. The `loop.sh` only executes `build` ticks. Review and fix are left as manual skills. The result is that the loop cannot autonomously:

- review a completed queue
- generate fix work units from findings
- re-run the build pass on those fixes
- stop when review finds no actionable issues

This is the missing piece for an autonomous, bounded development loop.

## Decision

The loop will optionally orchestrate a bounded `build → review → fix` subloop. Review remains a skill; the loop does not implement review logic. The loop only invokes the `review` and `fix` workers and interprets their outputs as signals to continue or stop.

See `decisions/0008-loop-orchestrates-review-fix.md` (Unit 1 of this cycle) for the formal ADR.

## Sequence

```text
build pass
  ↓
if --review and no pending units:
  review worker → writes .loop/<name>/REVIEW.md
  ↓
  if actionable findings > 0:
    fix worker → appends fix units to .loop/<name>/QUEUE.md
    ↓
    build pass on new units
    ↓
    review again
  else:
    cycle complete
```

## Hard stops

- `--max-ticks N` — total build ticks across all rounds
- `--max-review-rounds M` — default 2
- no actionable findings in a review pass
- repeated identical actionable findings (no progress)
- review or fix worker exits non-zero
- fix worker produces no new units despite actionable findings

## New loop options

- `--review` — enable review-fix rounds after build pass completes
- `--max-review-rounds N` — default `2`
- `--review-agent-cmd` — optional agent command for review ticks (defaults to `LOOP_AGENT_CMD`)
- `--fix-agent-cmd` — optional agent command for fix ticks (defaults to `LOOP_AGENT_CMD`)

Environment variables for symmetry:

- `LOOP_REVIEW_CMD` — overrides `--review-agent-cmd`
- `LOOP_FIX_CMD` — overrides `--fix-agent-cmd`

Per-unit `Agent:` override does not apply to review/fix ticks; those are orchestrated by the loop.

## New disposable artifact: `REVIEW.md`

The `review` skill will write a structured file next to `QUEUE.md`:

```markdown
# Review: <cycle>
Generated: <timestamp>

## Standards
- [ ] high | medium | low — <finding>
  evidence: path/to/file.go:42 — "quoted line"

## Intent
- [ ] high | medium | low — <finding>
  evidence: path/to/file.go:42 — "quoted line"

## Speculative
- low / uncitable — <finding>

## Summary
- actionable: N
- trivial: N
- disputed: N
- deferred: N
```

The loop checks the `actionable` count to decide whether to run fix. It does not parse the rest of the file.

## Worker prompts

`prompts/worker.md` stays the same for build ticks. We add two optional prompt templates:

- `prompts/reviewer.md` — loaded for review ticks; tells the worker to load the `review` skill and write `REVIEW.md`
- `prompts/fixer.md` — loaded for fix ticks; tells the worker to load the `fix` skill and consume `REVIEW.md`

If these templates are not present, `loop.sh` falls back to `prompts/worker.md` with the skill name prepended.

## Integration with existing skills

- `review/SKILL.md`: update to read `QUEUE.md` and `EVIDENCE.md`, write `REVIEW.md`, and classify findings.
- `fix/SKILL.md`: update to read `REVIEW.md`, triage findings, and append `Status: pending` units to `QUEUE.md`.
- `build/SKILL.md`: unchanged.

## Tests

`tests/run.sh` will gain a test with fake workers:

1. Build worker creates a file with a deliberate bug.
2. Review worker writes `REVIEW.md` with one actionable finding.
3. Fix worker appends a unit to `QUEUE.md`.
4. Build worker on the new unit fixes the bug.
5. Review worker writes `REVIEW.md` with zero actionable findings.
6. Loop stops cleanly.

## Documentation

- `docs/loop.md` — document `--review`, `--max-review-rounds`, `LOOP_REVIEW_CMD`, `LOOP_FIX_CMD`, `REVIEW.md`
- `docs/skills.md` — clarify that review and fix are normally invoked by the loop when `--review` is set
- `README.md` — mention the optional review-fix subloop
- `AGENTS.md` — update current state and working conventions

## Out of scope

- LLM-as-judge review. Review remains a skill output; the loop only interprets the structured `REVIEW.md`.
- Auto-review after every build tick. Review happens after the build queue drains.
- Integration with `knack status` or `knack view`/`knack list`. Those are separate CLI work.
