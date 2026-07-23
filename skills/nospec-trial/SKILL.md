---
name: nospec-trial
description: Use when reviewing a change against intent and standards — adversarial scrutiny before accepting work. Two axes — standards (does the change follow the codebase's own conventions?) and intent (does it do what it was supposed to?). Reviews against the actual codebase, not specs. Triggers on "review", "check this", "is this right", "what did we miss", "stress-test the implementation", or when work needs adversarial scrutiny before being accepted.
---

# Review

Adversarial review of a change. Two axes, run independently so neither pollutes the other:

1. **Standards** — does the change follow the repo's coding conventions and codebase patterns?
2. **Intent** — does the change do what it was supposed to?

Review against the **actual codebase**, not against specs that may have rotted — stale specs are worse than none (doc-rot); the code is the source of truth.

That core is the same whether the change is an interactive edit or a completed batch work unit. What the batch path adds is artifacts: a `QUEUE.md` unit states what was promised, an `EVIDENCE.md` records what verify proved, and a `REVIEW.md` carries findings to the `nospec-mend` skill. Interactively, the "unit" is the stated intent and the change itself; findings go straight to the human or back into the work.

## When to review

- After a change is made and its verify passes
- After a full batch queue is completed
- When the user asks for a sanity check
- Before accepting work as finished

Review is not a gate the loop enforces — it's a skill the user or agent invokes when they want adversarial scrutiny.

## Before you review

Read the change (the diff). If there's a work unit (`.loop/<name>/QUEUE.md`) and evidence (`.loop/<name>/EVIDENCE.md`), read them — the unit states what was promised, the evidence what verify proved. Read the `Pin alerts:` section of `EVIDENCE.md` if present — each alert means a durable doc that a prior cycle pinned has since changed, which is a coherence signal the verify gate cannot see. If the prompt includes a `Design:` path, read `.loop/<name>/DESIGN.md` for the reasoning context (external constraints, decisions, trade-offs) the work was planned against; a deviation from its stated constraints is a stronger finding than one based on inference. Interactively, gather the stated intent from the request or conversation.

## Two-axis review

### Axis 1: Standards

Does the change follow the codebase's existing patterns?

- Read `AGENTS.md` for stated conventions.
- Read neighboring code — does the change look like it belongs?
- Check error handling, naming, file layout, test style.
- Look for regressions — did the change break something nearby?
- Check for dead code, unused imports, leftover debugging.
- Check for over-reach — does the change add abstraction, wrappers, or parallel paths the outcome never asked for? Speculative structure is a standards finding even when it works: surface area is bug surface, and the bugs land in the layers nobody required.
- Check for durable-context drift — if the change alters a public interface, a convention, a ruling, or a domain term, does it leave stale projections behind? If so, invoke the `document` skill to assess coherence and route corrections to the owning record.
- Check `Pin alerts:` in `EVIDENCE.md` — each alert is a durable doc that a prior cycle pinned and has since changed. Route each alert to the `document` skill to assess whether the change left stale projections in other durable docs that describe or depend on it. A pin alert is a triage trigger, not a coherence finding by itself — it says "something moved," not "something is wrong." `document` judges whether the move broke coherence.

The question is not "is this good code?" — that's subjective. The question is "does this match the codebase's own standards?"

### Axis 2: Intent

Does the change do what it was supposed to?

- Read the stated outcome and constraints — from the work unit's `Done means:` / `Constraints:` if present, otherwise from the request.
- Read the actual diff.
- Does the diff satisfy the stated outcome? Stay within the constraints?
- Did it introduce anything not asked for, or miss anything it said it would?

The `Verify:` command is the mechanically enforceable subset of the outcome. The gap between the outcome and its verify is the review surface: intent review checks what the verify command cannot.

Verify passing is a proxy, not proof — it means the mechanically enforceable subset held, not that the outcome is genuinely satisfied. Intent review asks the harder question: does the change *generalize*? Does it hold up against behavior the verify doesn't exercise, or did it satisfy the visible surface while leaving the real property unmet? A change that passes its verify only by fitting the visible tests — not by implementing the underlying behavior — is the signature intent finding, and it's the one the gate is structurally blind to.

The technique that surfaces it: probe *how* the result is produced, not just *what* was returned. A verify that checks return values can pass while the underlying behavior is wrong — a `traverse` that eagerly maps every element then checks for failure returns the right `Nothing` while still calling the callback on elements past the failure. The intent review asks whether the verify exercises the side effects that distinguish the real property from its visible shadow: invocation counts, processing order, variant exhaustion (does each sum-type variant do what its case demands, or did one branch get copy-pasted from its neighbor?). A counter test — assert the callback ran exactly twice, not three times — is the shape. When the verify only observes the outcome, the intent review observes the trace.

Know the limit before you lean on this. Most of these failures are omission — the worker knew how and didn't check, so a verification pause catches them. Some are comprehension — the worker misread the contract and will re-confirm its wrong reading under any pause that asks it to "verify the signature." A checklist entrenches comprehension failures; it doesn't fix them. The intent review pays off when the worker *could have* done it right and skipped; it doesn't when the worker's model of the problem is itself wrong, which is a `fix`-direction or a `decide`-direction problem, not a review finding.

> **Intent review is a judgment check, not a deterministic gate.** Be explicit about its scope and confidence. Any finding that becomes a new work unit must have a deterministic `Verify` command the runner can execute.

## Running the axes

Run both. They can be parallel (two passes over the same diff) or sequential. The order doesn't matter — what matters is that each axis is evaluated independently, without the other's conclusions bleeding in.

## Confidence and evidence

Every finding must cite the specific `file:line` that motivates it and state a confidence level. Before you write a finding into the report, quote the code that motivates it. If you can't, it goes to the appendix, not the report.

- **high** — you read the specific code and can quote the line. Promoted to the report.
- **medium** — pattern match, likely but not verified against the actual code. Promoted but flagged.
- **low / uncitable** — you can't point to a specific line. **Not promoted.** Banished to `## Speculative`.

Confidence is orthogonal to classification. This discipline exists because reviewers overcorrect — asked to explain and propose fixes, LLMs systematically misclassify correct code as defective (overcorrection-bias). Citing a line you actually read is the filter that keeps false negatives out of the report.

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
