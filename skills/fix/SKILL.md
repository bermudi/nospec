---
name: fix
description: Use when addressing review findings and feeding them back into the loop. Converts review findings into new verifiable work units in `.loop/<name>/QUEUE.md` and runs another loop pass. Triggers on "fix the review findings", "address the feedback", "rework based on review", "fix what review found", or when review produced findings that need to be resolved.
metadata:
  version: "1.0.0"
---

# Fix

Address review findings by triaging `REVIEW.md` and generating new work units that feed back into the loop. The `fix` skill is the bridge between review and another build pass; it owns the triage and the queue format. It appends to the existing `.loop/<name>/QUEUE.md` file for the same cycle.

The loop owns orchestration. When `fix` is invoked by the loop, stop after appending units and reporting the triage summary; the loop will run the next build pass.

## Inputs

- `.loop/<name>/REVIEW.md` — structured findings from the `review` skill. Read it, but do not edit it.
- `.loop/<name>/QUEUE.md` — the existing work queue. Read it before appending so new units preserve the queue's current structure and statuses.
- `.loop/<name>/DESIGN.md` — if the prompt includes a `Design:` path, read it. It carries the reasoning context the work units were planned against. Use it to resolve ambiguity in a finding's `fix direction`: if the design note states a constraint or decision, the fix direction must align with it, not contradict it.

`REVIEW.md` is expected to include `## Standards`, `## Intent`, `## Speculative`, and `## Summary` sections. The summary's actionable count is the loop's signal; the finding entries are the fix skill's input. Use the finding classification, confidence, and evidence when deciding whether to generate a work unit.

## Procedure

1. **Read `REVIEW.md`.** Each standards or intent finding is a candidate work unit. `review` has already classified findings as trivial / actionable / disputed / deferred; use those classifications or re-triage as needed. Treat speculative findings as notes for future explore/plan work unless the review explicitly classifies one as actionable.

2. **Triage.** Not every finding warrants a work unit:
   - **Actionable** — the finding identifies a real issue with a clear fix. Create a work unit.
   - **Trivial** — one-line fix, no risk. Note it in the summary, don't create a unit.
   - **Disputed** — the finding is wrong or the reviewer is being overly cautious. Note the disagreement in the summary, don't create a unit. Move on.
   - **Deferred** — the finding is valid but not worth fixing now. Note it in the summary or `LEARNINGS.md`, don't create a unit.

3. **Create work units** for actionable findings. Use the standard work unit format from the `plan` skill. Each unit must have one observable outcome and one deterministic verify command. Reference the source finding in `Read first:` by pointing to `REVIEW.md` plus any evidence paths from the finding.

````markdown
## <fix for the finding — observable outcome>

Read first:
- .loop/<name>/REVIEW.md (<finding id or heading>)
- <evidence path from the finding, if any>
- <2–4 entries; context, not scope>

Constraints:
- <boundary>
- <what must stay true or out of bounds; if a file is named, it is "don't touch X" or "X's public API must not change", not "update X">

Done means:
- <the finding is resolved>
- <no new issue introduced>

Verify:
```bash
<deterministic command that proves the fix>
```

Status: pending
````

4. **Append to `QUEUE.md`.** Append new units to the end of the existing `.loop/<name>/QUEUE.md`. Do not reorder existing units, do not edit existing unit bodies, and do not change any existing `Status:` lines. If there are no actionable findings, append nothing.

5. **Stop after the queue edit.** Report the triage summary and the number of appended units. The loop will decide whether to run another build pass.

## What fix is not

- Not a re-review — the findings are already known. Fix turns them into work.
- Not a debate — if a finding is disputed, note it and move on. Don't argue with the review in the queue.
- Not a rewrite — each fix unit should be narrow. If a finding requires a broad rewrite, that's a new explore → plan cycle, not a fix.
- Not the orchestrator — the loop invokes review, invokes fix, and runs the next build pass.

## Guardrail

Fix units must not break what review already approved. Each fix unit's `Done means:` should include a no-regression condition. The verify command should cover both the fix and the preservation of what was already working.

## Output

- New `Status: pending` work units appended to `.loop/<name>/QUEUE.md`
- A summary of what was triaged (actionable / trivial / disputed / deferred)
- The number of units appended
