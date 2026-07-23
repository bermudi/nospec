---
name: nospec-mend
description: Use when resolving review findings — directly when interactive, or as new verifiable work units appended to a `.loop/<name>/QUEUE.md` when running batch. Triggers on "fix the review findings", "address the feedback", "rework based on review", "fix what review found", or when review produced findings that need resolving.
license: MIT
metadata:
  author: bermudi
  version: "1.0.0"
---

# Fix

Resolve review findings. Interactively, fix them directly (or hand them to the human). In a batch cycle, triage the findings and append the actionable ones as new work units to `.loop/<name>/QUEUE.md`; the loop runs the next build pass and review re-checks.

That review → fix → build → review cycle is iterative self-correction, and its known failure mode is overcorrection — the loop can oscillate rather than converge. So fix units stay narrow, guard what review already approved, and don't rewrite: a finding that needs a broad rewrite is a new planning pass, not a fix.

The loop owns orchestration. When the loop invokes fix, stop after appending units and reporting the triage summary; the loop decides whether to run another build pass.

## Inputs

- The findings — from `.loop/<name>/REVIEW.md` in a batch cycle, or from the review/conversation interactively. Read them; don't edit them.
- `.loop/<name>/QUEUE.md` — the existing queue (batch). Read it before appending so new units preserve its structure and statuses.
- `.loop/<name>/DESIGN.md` — if a `Design:` path is provided. It carries the reasoning the work was planned against; use it to resolve ambiguity in a finding's fix direction — the fix must align with the design note's constraints, not contradict them.

`REVIEW.md` carries `## Standards`, `## Intent`, `## Speculative`, and `## Summary` sections; the summary's actionable count is the loop's signal.

## Triage

Not every finding warrants a fix:

- **Actionable** — a real issue with a clear fix. Resolve it (interactively) or create a work unit (batch).
- **Trivial** — one-line, no risk. Fix inline; note it in the summary.
- **Disputed** — the finding is wrong or overcautious. Note the disagreement; don't act. Move on.
- **Deferred** — valid but not now. Note it in the summary.

Treat speculative findings as notes for future work unless review explicitly classified one as actionable.

## Diagnose before you iterate

A fix that patches a symptom without the cause fixes one thing and breaks another; iterate without diagnosis and the loop amplifies error rather than converging — each attempt addresses a surface failure and opens a new one. The signal is blunt: many attempts on the same failing spot, no convergence.

Before the next attempt, isolate the cause. Re-read the code around the failure; find the invariant actually being violated, not the test reporting it. A fix that lands in one or two attempts usually had the cause right; one that needs many almost certainly doesn't — more attempts won't help, a diagnosis will. When the cause is genuinely ambiguous, that's a design question for the `nospec-rule` skill or the human, not a fifth fix attempt.

This is judgment, not a gate on attempt count — the point is to switch from patching to diagnosing the moment iteration stops converging.

## Creating fix units (batch)

For actionable findings in a batch cycle, create work units using the format from the `nospec-shape` skill — one observable outcome, one deterministic verify. Don't restate the format here; load `nospec-shape`. Reference the source finding in `Read first:` (point to `REVIEW.md` plus the finding's evidence paths). Each fix unit's `Done means:` should include a no-regression condition, and its verify should cover both the fix and the preservation of what was already working — fix units must not break what review approved.

Append new units to the end of the existing `.loop/<name>/QUEUE.md`. Don't reorder existing units, edit existing unit bodies, or change existing `Status:` lines — the loop's parser keys on `## ` headers and `Status:` lines (editing them desyncs status tracking), and review's approvals are recorded in those statuses. If there are no actionable findings, append nothing.

## What fix is not

- Not a re-review — the findings are known. Fix turns them into work.
- Not a debate — if a finding is disputed, note it and move on.
- Not a rewrite — each fix is narrow. A broad rewrite is a new planning pass, not a fix.
- Not the orchestrator — in batch, the loop invokes review, invokes fix, and runs the next pass.

## Output

- Interactive: the fixes applied, plus a summary of what was triaged.
- Batch: new `Status: pending` units appended to `QUEUE.md`, plus a triage summary (actionable / trivial / disputed / deferred) and the count of units appended. Stop after the queue edit; the loop decides whether to run another pass.
