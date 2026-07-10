# Loop Queue: review-loop

Goal:
Make `loop.sh` optionally orchestrate a bounded `build → review → fix` subloop. After the build queue drains, the loop runs a review worker, writes a structured `REVIEW.md`, and, if there are actionable findings, runs a fix worker that appends new work units to `QUEUE.md`. The loop re-runs build and review until no actionable findings remain, a review-round limit is hit, or no progress is detected.

Stop condition:
`bash -n /home/daniel/build/knack/loop.sh && /home/daniel/build/knack/tests/run.sh && cd /home/daniel/build/knack/cli && go test ./...`

## ADR-0008 captures the loop-orchestrated review-fix decision

Read first:
- `.loop/review-loop/specs/proposal.md`
- `DESIGN.md` (open questions and "What the loop does NOT do")
- `.agents/skills/review/SKILL.md`
- `.agents/skills/fix/SKILL.md`

Constraints:
- The ADR must state the loop's role (orchestrator) vs the review/fix skills' role (judgment).
- Do not change `loop.sh`, `review/SKILL.md`, or `fix/SKILL.md` yet.
- Keep the ADR to one page.

Done means:
- `decisions/0008-loop-orchestrates-review-fix.md` exists with `Status: accepted`.
- `knack decisions check` still passes (the new ADR is not orphaned — it is referenced by this queue).

Verify:
```bash
cd /home/daniel/build/knack && test -f decisions/0008-loop-orchestrates-review-fix.md && grep -q '^Status: accepted' decisions/0008-loop-orchestrates-review-fix.md && /tmp/knack decisions check
```

Status: pending

## review skill writes a structured REVIEW.md

Read first:
- `.agents/skills/review/SKILL.md`
- `.agents/skills/fix/SKILL.md`
- `decisions/0008-loop-orchestrates-review-fix.md`

Constraints:
- Preserve the two-axis (standards + intent) review structure.
- The output format must be machine-parseable enough for `loop.sh` to detect the `actionable` count.
- Do not change `loop.sh` or `fix/SKILL.md` yet.
- Keep the skill file scannable and sync it to `cli/embedded/skills/review/` after editing.

Done means:
- `review/SKILL.md` documents `REVIEW.md` as its output artifact.
- `REVIEW.md` has sections for `## Standards`, `## Intent`, `## Speculative`, and a `## Summary` with counts.
- Each finding has a confidence level (high/medium/low) and evidence (`path/to/file:line` or quoted code).
- `cli/embedded/skills/review/SKILL.md` is in sync with `.agents/skills/review/SKILL.md`.

Verify:
```bash
cd /home/daniel/build/knack && diff -r .agents/skills/review cli/embedded/skills/review && ./tests/run.sh
```

Status: pending

## fix skill consumes REVIEW.md and appends fix units to QUEUE.md

Read first:
- `.agents/skills/fix/SKILL.md`
- `.agents/skills/review/SKILL.md` after Unit 2
- `decisions/0008-loop-orchestrates-review-fix.md`

Constraints:
- The fix skill must not edit `REVIEW.md`.
- It must append new work units to the existing `QUEUE.md` without changing the status of existing units.
- It must follow the existing work unit format (outcome, Read first, Constraints, Done means, Verify, Status: pending).
- Keep the skill file scannable and sync it to `cli/embedded/skills/fix/` after editing.

Done means:
- `fix/SKILL.md` documents how to read `REVIEW.md` and triage findings into actionable units.
- The skill appends `Status: pending` units to `QUEUE.md`.
- `cli/embedded/skills/fix/SKILL.md` is in sync with `.agents/skills/fix/SKILL.md`.

Verify:
```bash
cd /home/daniel/build/knack && diff -r .agents/skills/fix cli/embedded/skills/fix && ./tests/run.sh
```

Status: pending

## loop.sh orchestrates bounded build-review-fix rounds

Read first:
- `loop.sh`
- `decisions/0008-loop-orchestrates-review-fix.md`
- `.loop/review-loop/specs/proposal.md`
- `tests/run.sh`

Constraints:
- Default behavior must be unchanged when `--review` is not set.
- `--max-ticks` remains the total build-tick budget across all rounds.
- Add `--max-review-rounds` (default 2) and `--review` flag.
- Support `LOOP_REVIEW_CMD` and `LOOP_FIX_CMD` env vars with fallback to `LOOP_AGENT_CMD`.
- Do not break the `Agent:` per-unit override for build ticks.
- Preserve `HANDOFF.md` and `EVIDENCE.md` behavior.
- New tests must not require a real model or API key.

Done means:
- `loop.sh --review` runs the build queue, then runs review, then fix if actionable, then loops.
- `loop.sh` stops when `REVIEW.md` reports `actionable: 0`, `--max-review-rounds` is reached, `--max-ticks` is reached, or no new units are generated.
- `tests/run.sh` has a test that uses fake workers to run a full `build → review → fix → review` cycle.
- `./tests/run.sh` and `bash -n loop.sh` pass.

Verify:
```bash
bash -n /home/daniel/build/knack/loop.sh && /home/daniel/build/knack/tests/run.sh
```

Status: pending

## docs and AGENTS.md reflect the new loop behavior

Read first:
- `docs/loop.md`
- `docs/skills.md`
- `README.md`
- `AGENTS.md`
- `decisions/0008-loop-orchestrates-review-fix.md`

Constraints:
- Do not change the core thesis.
- Document that `--review` is opt-in.
- Keep changes scoped to the review-fix subloop.
- Sync embedded skills if any skill files changed.

Done means:
- `docs/loop.md` documents `--review`, `--max-review-rounds`, `LOOP_REVIEW_CMD`, `LOOP_FIX_CMD`, and `REVIEW.md`.
- `docs/skills.md` and `README.md` describe review/fix as loop-orchestrated when `--review` is set.
- `AGENTS.md` current state and working conventions mention the new loop capability.
- Tests pass.

Verify:
```bash
cd /home/daniel/build/knack && ./tests/run.sh && cd cli && go test ./...
```

Status: pending
