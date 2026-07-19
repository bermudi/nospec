---
name: build
description: Use when implementing a bounded, observable outcome — whether supplied conversationally or as a work unit from a `.loop/<name>/QUEUE.md`. Verify-first — read the verify before changing code, and don't declare done until it actually passes. Triggers on "build", "implement", "apply this unit", "do the work", "run the loop", or when work needs executing. Also used when the loop invokes the worker for a tick.
---

# Build

Implement one bounded, observable outcome. Do the work. Don't declare done on vibes — the verify must actually pass before you stop.

That core is the same whether the outcome arrived in conversation or as a unit in a `.loop/<name>/QUEUE.md`. What changes across modes is who runs the verify. Interactively, you run it yourself — there's no external runner, so the discipline of proving it before you claim it is yours. In a batch cycle, the loop owns the verify gate: it runs the unit's `Verify:` command after you exit, so you never self-certify; your job is to make the repository state satisfy that command. The principle — don't claim success without verification — survives across modes; the enforcement mechanism doesn't.

## Scope

The scope is the outcome plus its constraints — never a file list. You decide which files to change and how; constraints close the solution space, they don't prescribe the path (ADR-0005). The `Verify:` command is the mechanically enforceable subset of `Done means:`; the gap between them is the review surface.

If the unit's `Read first:` cites `.loop/<name>/DESIGN.md`, read it first — it carries cycle-level reasoning you can't recover from the codebase alone. The planner's `Read first:` is the worker's channel to it (review and fix get `Design:` injected directly, because they span the whole queue).

## Execution discipline

How you reach the outcome matters as much as reaching it. Three concepts, each preventing a concrete failure mode:

**Close the loop early.** The strongest predictor of a clean outcome is *when* you first run something executable against the work — not whether you verify at the end, but how soon you close a feedback loop after the first change. Bugs bake into unverified code; the longer the gap between writing and first testing, the more error accumulates blind and the costlier the eventual fix. Default: as soon as a sliver of the outcome works end-to-end, run it — don't accumulate a large diff you've never exercised. Override: a change small enough to hold in your head — but notice when "trivial" is a rationalization for not closing the loop.

**Edit structurally.** Prefer replacing a coherent block over chaining fragile inline patches. Each inline string-surgery edit is a chance to break syntax, imports, or formatting, and failed edits cascade into repair cycles that consume the budget without advancing the outcome. Default: replace whole functions or blocks; reach for inline edits only on genuine one-liners, and validate after each. Override: a harness where inline edits are the sane primitive — but never batch fragile edits and hope.

**Keep it minimal.** The smallest change that satisfies the outcome is the default; abstraction is earned, not assumed. Every speculative layer, wrapper, or parallel path is an invariant you now have to get right *and* verify — surface area is bug surface, and the bugs land in the layers the outcome never asked for. Default: extend the existing path before building a parallel one; add abstraction only when the outcome demands it. Override: a change whose purpose *is* to introduce an abstraction — but then the abstraction is the outcome, not decoration.

These are judgment, not gates. The verify gate below is the mechanical contract; these are how you reach it without baking in avoidable error.

## Verification

Verify-first: read the `Verify:` command before you change code, so you know what state the repo must reach. Then make it pass.

- **Interactive** — run the verify yourself before you stop. If it fails, fix the cause if it belongs to this outcome; otherwise report the blocker. Don't claim success the command didn't confirm.
- **Batch (under a runner)** — the runner runs `Verify:` after you exit. You don't self-certify and you don't mark the unit done; you make the repository state satisfy the command. Don't edit `.loop/<name>/EVIDENCE.md` — the runner writes the evidence ledger after verification.

The verify gate is the backpressure — the mechanism that mechanically rejects wrong output, outside the agent. Your relationship to it is to aim the work at making it pass, not to assert that it would.

## Capturing decisions during build

Implementation surfaces decisions the spec missed — the code pushes back, and that's when the most valuable rulings crystallize (code-clarifies-spec). If you discover one — "we need to handle X this way because Y" — write the ADR now via the `decide` skill, inline, not after the outcome. It's a durable trace, not part of the verify scope.

## Capturing operational learnings

If you learn how the project works, capture it in the right durable file:

- **Operational gotchas** — build commands, test conventions, how to verify — go in `AGENTS.md`.
- **Domain or problem insights** — "X doesn't work because Y", "the parser has this surprising property" — go in `LEARNINGS.md`.

If you discover that a durable doc contradicts a ruling, a term, or the current code, invoke the `document` skill to route the correction to the owning record and update its projections. Don't patch the contradiction in place without checking whether the same claim lives elsewhere.

Don't add trivia. Add what would have saved you time upfront.

## When the work is too big for one pass

Do as much as keeps the repo in a working state, then hand off what remains — interactively to the human or the next session; in a batch cycle, as a handoff note. If verify fails in batch, the runner writes a handoff and stops; a later session resumes from it. Don't cram work that genuinely needs more into one pass; the loop is designed for multiple ticks, and interactive work resumes from where you left it.

## When you're blocked

State the blocker clearly, note what would unblock you (a decision, a dependency, a missing file), and stop. Don't thrash — and recognize thrashing's signature: if verify keeps failing on the same spot, you're almost certainly treating a symptom, not the cause. Stop patching and re-read the code to isolate the invariant actually being violated; a fix that needs many attempts is a diagnosis problem, not an attempt-count problem. Interactively that's a message to the human; in a batch cycle the runner marks the unit `blocked` and writes a handoff the next session picks up.

## Batch behavior (under a runner)

When the loop invokes you for a tick, `prompts/worker.md` governs the tick — one unit only, the handoff output format, the don't-self-certify posture. The sections above are the skill's core; the worker prompt is the batch-tick protocol. If it's in your context, follow it.
