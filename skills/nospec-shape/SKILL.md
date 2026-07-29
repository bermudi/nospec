---
name: nospec-shape
description: Use when work must be understood or bounded before implementation — clarifying intent, investigating uncertainty, decomposing outcomes, or preparing a cross-session or batch queue.
license: MIT
metadata:
  author: bermudi
  version: "1.0.0"
---

# Shape

Turn intent into bounded, observable work. Scouting and decomposition are levels of one evidence-seeking stance, not required phases. Skip them when a clear small change can go directly to `nospec-carve`.

## Raise fidelity only as needed

Start with project operational context and authoritative records, then inspect relevant code, tests, and history. Separate the desired outcome from the proposed solution, find existing implementation paths, and identify what must remain true.

When reading and discussion cannot resolve a consequential question about how something should look, behave, or fit the system, read [`references/scouting.md`](references/scouting.md). A runnable experiment is evidence, not automatically production code.

A recommendation is not an accepted ruling. Present consequential choices to the decision owner; after acceptance, use the project's ADR practice or `nospec-curator` when installed.

## Cut for useful evidence

One work unit produces one observable outcome. Choose the cut that exposes the relevant risk earliest:

- A **tracer bullet** crosses an uncertain path end to end; it is waste when no integration assumption needs proving.
- A **vertical slice** exposes cross-layer contract mismatches; too thin gives false confidence, too broad delays feedback.
- **Horizontal breadth** efficiently applies an already-proven contract; used too early, it multiplies the same wrong assumption.

A unit may carry `Read first:` context and `Constraints:` boundaries. It must carry nonempty `Done means:` criteria and a mechanically checkable `Verify:` subset. The worker owns the implementation path; context and constraints must not become an edit script.

Verification must fail for a believable state where the central outcome is absent. If it cannot, strengthen the assertion, split the outcome, establish a better verification seam first, or keep the work interactive. Review may inspect the remaining judgment surface but cannot turn it into deterministic proof.

## Disclose the applicable mode

Interactive shaping can remain conversational. Work specs are disposable coordination state; project-designated contracts are durable regardless of filename.

- When execution needs a cross-session or batch handoff, read [`references/queue-format.md`](references/queue-format.md). Batch requires every unit to be bounded, runner-verifiable, and free of unresolved decisions.
- When the loop supplies `REVIEW.md` for conversion into new units, also read [`references/review-findings.md`](references/review-findings.md). This fixer mode appends queue units and never edits source.

Return clarified intent and bounded outcomes, an optional queue, a decision request, or a justified no-change conclusion. Do not manufacture a plan merely because this skill was invoked.
