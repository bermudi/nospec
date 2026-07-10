# 0008: The loop orchestrates the review-fix subloop; skills keep the judgment

Date: 2026-07-10
Status: accepted

## Context

The default skill flow is `explore → plan → build → review → fix`, but `loop.sh` only runs `build` ticks. Review and fix are left as manual skills, so the loop cannot autonomously review a completed queue, turn findings into fix units, re-run the build pass, and stop when review is clean. That is the missing piece for a bounded, self-correcting development loop.

The architectural question is *who owns what*. DESIGN.md's "What the loop does NOT do" already states the loop does not run review and does not manage ADRs — those are skill responsibilities. Extending the loop into review/fix risks blurring that boundary: if the loop starts interpreting review content, it stops being the simple, agent-agnostic engine and becomes an LLM-as-judge.

Two roles must stay separate:

- **Orchestration** (the loop) — when to invoke review, when to invoke fix, when to stop. Mechanical and signal-driven.
- **Judgment** (the `review` and `fix` skills) — what counts as a finding, whether it is actionable, how to phrase a fix unit. Adversarial and semantic.

Alternatives considered:

- **LLM-as-judge inside the loop.** Rejected — the loop would interpret review content, breaking agent-agnosticism and the verify-gate principle (the aiming problem: the signal must be the actual property, not a proxy the loop can game).
- **The loop implements review logic.** Rejected — duplicates the `review` skill, drifts from the source of truth, and violates "the loop never reads skills."
- **Manual review/fix only (do nothing).** Rejected — leaves the loop unable to self-correct; the missing piece stays missing.

## Decision

The loop **orchestrates** an optional, bounded `build → review → fix` subloop. It invokes the `review` and `fix` workers and interprets their *structured* outputs as continue/stop signals. It does not implement review or fix logic.

When `--review` is set and the build queue has drained, the loop runs a review worker that writes a structured `REVIEW.md`. The loop reads only the `actionable` count from that file. If it is non-zero, the loop runs a fix worker that appends `Status: pending` units to `QUEUE.md`, then re-runs the build pass and reviews again. It stops when `actionable` is zero, a review-round limit (`--max-review-rounds`, default 2) is hit, the tick budget (`--max-ticks`) is exhausted, or a round generates no new units (no progress).

The boundary, stated as a ruling:

- **The loop owns orchestration and stop conditions** — invoking review/fix, reading the actionable count, enforcing hard stops. Mechanical.
- **The `review` and `fix` skills own judgment** — the two-axis standards/intent review, triage, finding phrasing, and work-unit generation. Semantic.

The loop never parses review content beyond the actionable count and never judges whether a finding is real. Review remains a skill the worker loads; the loop only knows to invoke it and read its summary.

## Consequences

- The loop gains a bounded self-correction capability without becoming an LLM-as-judge. Backpressure stays mechanical: the actionable count is the signal, the hard stops are the guardrails.
- A new contract surface: the loop depends on `REVIEW.md`'s actionable count being honest. A review worker that mislabels findings can loop forever or stop early — the round cap and the no-progress stop are the backstop, not the loop's own judgment.
- The `review` and `fix` skills gain a machine-readable input/output contract (`REVIEW.md`; appended units). Those skill changes land in their own work units; this ADR records only the architectural ruling and the orchestration boundary.
- Review stays opt-in (`--review`); default loop behavior is unchanged.
