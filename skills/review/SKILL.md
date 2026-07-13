---
name: review
description: Use when reviewing a change against intent and standards — adversarial scrutiny before accepting work. Two axes — standards (does the change follow the codebase's own conventions?) and intent (does it do what it was supposed to?). Reviews against the actual codebase, not specs. Triggers on "review", "check this", "is this right", "what did we miss", "stress-test the implementation", or when work needs adversarial scrutiny before being accepted.
---

# Review

Adversarial review of a change. Two axes, run independently so neither pollutes the other:

1. **Standards** — does the change follow the repo's coding conventions and codebase patterns?
2. **Intent** — does the change do what it was supposed to?

Review against the **actual codebase**, not against specs that may have rotted — stale specs are worse than none ([doc-rot](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/doc-rot.md)); the code is the source of truth.

That core is the same whether the change is an interactive edit or a completed batch work unit. What the batch path adds is artifacts: a `QUEUE.md` unit states what was promised, an `EVIDENCE.md` records what verify proved, and a `REVIEW.md` carries findings to the `fix` skill. Interactively, the "unit" is the stated intent and the change itself; findings go straight to the human or back into the work.

## When to review

- After a change is made and its verify passes
- After a full batch queue is completed
- When the user asks for a sanity check
- Before accepting work as finished

Review is not a gate the loop enforces — it's a skill the user or agent invokes when they want adversarial scrutiny.

## Before you review

Read the change (the diff). If there's a work unit (`.loop/<name>/QUEUE.md`) and evidence (`.loop/<name>/EVIDENCE.md`), read them — the unit states what was promised, the evidence what verify proved. If the prompt includes a `Design:` path, read `.loop/<name>/DESIGN.md` for the reasoning context (external constraints, decisions, trade-offs) the work was planned against; a deviation from its stated constraints is a stronger finding than one based on inference. Interactively, gather the stated intent from the request or conversation.

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

Does the change do what it was supposed to?

- Read the stated outcome and constraints — from the work unit's `Done means:` / `Constraints:` if present, otherwise from the request.
- Read the actual diff.
- Does the diff satisfy the stated outcome? Stay within the constraints?
- Did it introduce anything not asked for, or miss anything it said it would?

The `Verify:` command is the mechanically enforceable subset of the outcome. The gap between the outcome and its verify is the review surface: intent review checks what the verify command cannot.

> **Intent review is a judgment check, not a deterministic gate.** Be explicit about its scope and confidence. Any finding that becomes a new work unit must have a deterministic `Verify` command the runner can execute.

## Running the axes

Run both. They can be parallel (two passes over the same diff) or sequential. The order doesn't matter — what matters is that each axis is evaluated independently, without the other's conclusions bleeding in.

## Confidence and evidence

Every finding must cite the specific `file:line` that motivates it and state a confidence level. Before you write a finding into the report, quote the code that motivates it. If you can't, it goes to the appendix, not the report.

- **high** — you read the specific code and can quote the line. Promoted to the report.
- **medium** — pattern match, likely but not verified against the actual code. Promoted but flagged.
- **low / uncitable** — you can't point to a specific line. **Not promoted.** Banished to `## Speculative`.

Confidence is orthogonal to classification. This discipline exists because reviewers overcorrect — asked to explain and propose fixes, LLMs systematically misclassify correct code as defective ([overcorrection-bias](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/overcorrection-bias.md)). Citing a line you actually read is the filter that keeps false negatives out of the report.

## Findings become input to the fix skill

- **Trivial** findings (a typo, a missing newline) can be fixed inline during review — but only interactively. In batch the reviewer prompt forbids editing source, so hand them to `fix` like any actionable finding.
- **Actionable** findings are handed to `fix`. In a batch cycle they're written to `.loop/<name>/REVIEW.md` for `fix` to triage into new work units; do not write the units yourself, `fix` owns that. Interactively, hand them to the human or act on them directly.
- **Disputed** or **deferred** findings are recorded in the review summary, not turned into units.

The output of review is a structured set of findings, not a queue edit.

## What review is not

- Not a lint pass — the verify gate already ran. Review is about what verify *can't* check.
- Not a spec compliance check — specs are disposable and may have rotted. Review against the codebase.
- Not a gate — the loop doesn't enforce review. It's invoked when the user wants it.

## Output

In a batch cycle, write the structured artifact to the requested review output path (`.loop/<name>/REVIEW.md`, or wherever the prompt's `Review output:` points). Interactively, communicate findings directly — the format below is the batch contract the loop's parser reads; use it when there's a runner, skip it when there isn't.

`REVIEW.md` must have exactly these top-level sections:

1. `## Standards`
2. `## Intent`
3. `## Speculative`
4. `## Summary`

Put standards-axis findings under `## Standards`, intent-axis under `## Intent`. Use `## Speculative` for plausible-but-ungrounded concerns. If a section is clean, write `No issues found.`

Each finding must include:

- A stable id (`S1`, `I1`, `X1`)
- Classification: `trivial`, `actionable`, `disputed`, or `deferred`
- Confidence: `high`, `medium`, or `low`
- Evidence: a `path/to/file:line` reference and a short quoted excerpt
- Finding: the issue in one or two sentences
- Fix direction: a single, unambiguous direction for `fix`, or `None` for non-actionable findings. Do not offer options or conditional branches — the fixer should not have to make a judgment call. If you see two valid approaches, pick one and state it. (Ambiguity invites the fixer to overcorrect; the design note, if provided, should help you decide.)

Finding shape:

```markdown
- S1 | actionable | high
  evidence: `path/to/file:42` — "the quoted line or block that motivates this"
  finding: The change violates the repo's existing queue parser behavior.
  fix direction: Align the parser with the shell loop's unit-header rules.
```

The full shape — with `evidence: file:line` — applies to Standards and Intent findings. Speculative findings relax it: free-form grounding, or omit `evidence:` (they're in `## Speculative` precisely because you couldn't cite a specific line).

The `## Summary` section must include counts:

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

The `- actionable: N` line is the loop's continue/stop signal. Count only findings classified as `actionable`. If there are none, write `- actionable: 0`.
