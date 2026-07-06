---
name: review
description: Use when reviewing completed work units or a finished queue. Two-axis adversarial review — standards (does the change follow codebase conventions?) and intent (does the change do what the work unit said?). Reviews against the actual codebase, not specs. Triggers on "review", "check this", "is this right", "what did we miss", "stress-test the implementation", or when work needs adversarial scrutiny before being accepted.
---

# Review

Adversarial review of completed work. Two axes, run independently so neither pollutes the other:

1. **Standards** — does the change follow the repo's coding conventions and codebase patterns?
2. **Intent** — does the change do what the work unit said it would?

Review against the **actual codebase**, not against specs that may have rotted. The code is the source of truth.

## When to review

- After a work unit is marked `done` and verify passed
- After a full queue is completed
- When the user asks for a sanity check
- Before accepting work as finished

Review is not a gate the loop enforces — it's a skill the user or agent invokes when they want adversarial scrutiny.

## Before you review

Read the work unit from `.loop/<name>/QUEUE.md` and the evidence from `.loop/<name>/EVIDENCE.md` for the unit you're reviewing. The evidence tells you what the verify command actually proved; the work unit tells you what was promised. Review against the actual codebase, not the specs.

## Two-axis review

### Axis 1: Standards

Does the change follow the codebase's existing patterns?

- Read `AGENTS.md` for stated conventions.
- Read neighboring code — does the change look like it belongs?
- Check error handling, naming, file layout, test style.
- Look for regressions — did the change break something nearby?
- Check for dead code, unused imports, leftover debugging.

The question is not "is this good code?" — that's subjective. The question is "does this match the codebase's own standards?"

### Axis 2: Intent

Does the change do what the work unit said it would?

- Read the work unit's `Work:` and `Done means:` fields.
- Read the actual diff.
- Does the diff satisfy the stated outcome?
- Did the change introduce anything the unit didn't ask for?
- Did the change miss anything the unit said it would do?

The verify command passing is necessary but not sufficient. Verify checks one thing. Intent review checks whether the unit's full promise was delivered.

> **Intent review is a judgment check, not a deterministic gate.** Be explicit about its scope and confidence. Any finding that becomes a new work unit must have a deterministic `Verify` command the runner can execute.

## Running the axes

Run both axes. They can be parallel (two passes over the same diff) or sequential. The order doesn't matter — what matters is that each axis is evaluated independently, without the other's conclusions bleeding in.

## Findings become input to the fix skill

Review findings are not just notes — they are the input to the `fix` skill, which triages them and appends actionable ones as new work units in `.loop/<name>/QUEUE.md`.

- **Trivial** findings (a typo, a missing newline) can be fixed inline during review.
- **Actionable** findings are handed to the `fix` skill. Do not write the work units yourself; `fix` owns the triage and formatting.
- **Disputed** or **deferred** findings are recorded in the review summary but not turned into units.

The output of review is a findings summary, not a queue edit.

## What review is not

- Not a lint pass — the verify gate already ran. Review is about what verify *can't* check.
- Not a spec compliance check — specs are disposable and may have rotted. Review against the codebase.
- Not a gate — the loop doesn't enforce review. It's invoked when the user wants it.

## Output

Summarize findings per axis:

- **Standards**: N findings (list them, or "no issues found")
- **Intent**: N findings (list them, or "the change matches the unit's stated outcome")

Classify each finding as trivial / actionable / disputed / deferred.

Then either:
- Hand actionable findings to the `fix` skill, which will triage and append work units to `.loop/<name>/QUEUE.md`, or
- Report "no action needed" if the work is clean.
