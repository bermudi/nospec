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

If the prompt includes a `Design:` path, read `.loop/<name>/DESIGN.md` before reviewing. It carries the reasoning context — external constraints, decisions, trade-offs — that the work units were planned against. Use it to ground your findings: a deviation from the design note's stated constraints is a stronger finding than one based on inference alone.

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

- Read the work unit's `Done means:` and `Constraints:` fields.
- Read the actual diff.
- Does the diff satisfy the stated outcome?
- Does the diff stay within the stated constraints?
- Did the change introduce anything the unit didn't ask for?
- Did the change miss anything the unit said it would do?

The `Verify:` command is the mechanically enforceable subset of `Done means:`. The gap between `Done means:` and `Verify:` is the review surface: intent review checks what the verify command cannot.

> **Intent review is a judgment check, not a deterministic gate.** Be explicit about its scope and confidence. Any finding that becomes a new work unit must have a deterministic `Verify` command the runner can execute.

## Running the axes

Run both axes. They can be parallel (two passes over the same diff) or sequential. The order doesn't matter — what matters is that each axis is evaluated independently, without the other's conclusions bleeding in.

## Confidence and evidence

Every finding must cite the specific `file:line` that motivates it and state a confidence level. Before you write a finding into the report, quote the code that motivates it. If you can't, it goes to the appendix, not the report.

Use three confidence levels:

- **high** — you read the specific code and can quote the line. Promoted to the report, handed to `fix`.
- **medium** — pattern match, likely but not verified against the actual code. Promoted but flagged; `fix` treats it as worth investigating, not worth acting on blindly.
- **low / uncitable** — you can't point to a specific line. **Not promoted.** Banished to `## Speculative`. Only surfaces if the user reads the appendix.

Confidence is orthogonal to classification. A finding can be `actionable` + `high confidence`, or `deferred` + `low confidence`.

## Findings become input to the fix skill

Review findings are not just notes — they are written to `.loop/<name>/REVIEW.md` as the input to the `fix` skill, which triages them and appends actionable ones as new work units in `.loop/<name>/QUEUE.md`.

- **Trivial** findings (a typo, a missing newline) can be fixed inline during review.
- **Actionable** findings are handed to the `fix` skill. Do not write the work units yourself; `fix` owns the triage and formatting.
- **Disputed** or **deferred** findings are recorded in the review summary but not turned into units.

The output of review is a structured review artifact, not a queue edit.

## What review is not

- Not a lint pass — the verify gate already ran. Review is about what verify *can't* check.
- Not a spec compliance check — specs are disposable and may have rotted. Review against the codebase.
- Not a gate — the loop doesn't enforce review. It's invoked when the user wants it.

## Output

Write the structured review artifact to the requested review output path. In the loop this is `.loop/<name>/REVIEW.md`; if the prompt provides a different `Review output:` path, write there.

`REVIEW.md` must have exactly these top-level sections:

1. `## Standards`
2. `## Intent`
3. `## Speculative`
4. `## Summary`

Put standards-axis findings under `## Standards` and intent-axis findings under `## Intent`. Use `## Speculative` only for concerns that are plausible but not grounded enough to become a standards or intent finding. If a section is clean, write `No issues found.` under that section.

Each finding must include:

- A stable id, such as `S1`, `I1`, or `X1`
- Classification: `trivial`, `actionable`, `disputed`, or `deferred`
- Confidence: `high`, `medium`, or `low`
- Evidence: a `path/to/file:line` reference and a short quoted code excerpt
- Finding: the issue in one or two sentences
- Fix direction: a single, unambiguous direction for the `fix` skill, or `None` for non-actionable findings. Do not offer options or conditional branches — the fixer should not have to make a judgment call. If you see two valid approaches, pick one and state it. The design note (if provided) should help you decide which is correct.

Use this finding shape:

```markdown
- S1 | actionable | high
  evidence: `path/to/file:42` — "the quoted line or block that motivates this"
  finding: The change violates the repo's existing queue parser behavior.
  fix direction: Align the parser with the shell loop's unit-header rules.
```

Promoted findings go under `## Standards` or `## Intent`. Speculative findings (low confidence or uncitable) go under `## Speculative` so they don't pollute the actionable report.

The `## Summary` section must include counts using this machine-readable shape:

```markdown
## Summary
- standards: 1
- intent: 0
- speculative: 0
- actionable: 1
- trivial: 0
- disputed: 0
- deferred: 0
```

The `- actionable: N` line is the loop's continue/stop signal. Count only findings classified as `actionable`, including actionable speculative findings if you intentionally create one. If there are no actionable findings, write `- actionable: 0`.
