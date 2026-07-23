---
name: nospec
description: Use when reaching for unattended batch execution — driving a queue of verifiable work units behind a deterministic verify gate while the human is away. Carries the `nospec` runner as `scripts/nospec`. Triggers on "run the loop", "batch mode", "AFK", "run the queue", "nospec run", "install nospec", "set up nospec", or when the human wants to walk away from a planned queue and have the agent work through it.
compatibility: Requires bash and python 3. The runner is invoked as `scripts/nospec` from this skill directory (or via `nospec` on PATH after `nospec install`).
license: MIT
metadata:
  author: bermudi
  version: "1.0.0"
---

# Nospec

The optional batch companion to the skills. Most work is interactive — you're present, edits land directly, you run the verify yourself. Reach for the runner only when you want to leave: you've planned a `.loop/<name>/QUEUE.md` of verifiable units (via the `nospec-shape` skill) and want a mechanical loop to drive them through a verify gate while you're gone.

Skills are the product; the loop is an optional companion. This skill exists to carry the runner and to transmit *when* batch mode is the right attention level — not to push you toward it. If you're present, stay interactive.

## When batch is right

Batch mode trades presence for throughput. It fits when:

- The outcomes are already decomposed into bounded, independently-verifiable units (`nospec-shape` skill).
- Each unit's `Verify:` is deterministic and runnable by the runner — tests, builds, type checks, not an LLM-as-judge.
- The work doesn't need your judgment in the loop. If a unit might require a decision mid-execution, it belongs in interactive mode, or the decision should be made upfront and recorded (ADR via `nospec-rule`) so the worker can apply it.

If those don't hold, batch will pause on a handoff and wait for you anyway — save the round-trip and work interactively.

## What the runner guarantees, and what it doesn't

The runner owns one mechanical contract: it invokes the worker on a unit, then runs that unit's `Verify:` command *outside* the agent, and only marks the unit done if verify exits 0. The worker never self-certifies. This is the backpressure — the mechanism that mechanically rejects wrong output, outside the agent.

What the runner does **not** guarantee:

- **Coherence.** A green verify proves the verify command passed, not that the change coheres with the project's durable docs, rulings, or intent. That's the review surface — the gap between `Verify:` (mechanical) and `Done means:` (acceptance). The `nospec-trial` skill closes it, opt-in via `--review`.
- **Correctness beyond the verify scope.** The runner derives a proof boundary from the verify command (registry of primitives: `go test` → "Go test suite passes", `! grep` → "no matches for pattern", etc.) and records it in `EVIDENCE.md`. Unknown segments fall back to "command exited 0: `<segment>`" — never an interpreted claim. What's outside that scope is explicitly marked unverified.

Misread a green verify as proof of coherence and you've got reward-hacking Case 3 — a structural gate lending false confidence to the gaps that bite you. The runner is honest about this; read `EVIDENCE.md`'s "What this proves" and "What remains unverified" sections.

## Setup: putting the runner on PATH

The runner ships as `scripts/nospec` inside this skill directory. After `npx skills add`, it lives at `.agents/skills/nospec/scripts/nospec` (project-local) or `~/.agents/skills/nospec/scripts/nospec` (global, with `-g`). skills.sh installs skill files but does not touch PATH — so to invoke the runner as `nospec` from anywhere, run the install verb once:

```bash
.agents/skills/nospec/scripts/nospec install
```

That symlinks the runner onto PATH (it picks a writable directory on PATH, or `~/.local/bin` as fallback). After that, `nospec run ...` works from any directory. The user doesn't need to type the long path — when they ask you to set up nospec, run the install command for them.

If the user is using a global install (`-g`), the path is `~/.agents/skills/nospec/scripts/nospec install` instead.

## Invoking the runner

Once installed:

```bash
nospec run .loop/<name>/QUEUE.md
```

The runner drives a worker agent (set via `LOOP_AGENT_CMD`) which loads the `nospec-hew` skill by name — the worker's harness auto-loads skills by trigger text, same as any other skill invocation. No skill-path configuration is needed; the worker is a harness session, and harnesses find their own skills.

## What the runner does per tick

1. Reads the first `Status: pending` unit, marks it `in_progress`.
2. Snapshots the repo (diff + untracked files outside `.loop`).
3. Invokes the worker with the `nospec-hew` skill and the unit.
4. If the worker exits non-zero, marks the unit `blocked`, writes a handoff, stops.
5. Runs the unit's `Verify:` command outside the agent.
6. On success: marks `done`, appends evidence, continues.
7. On failure: if the repo didn't change, retries once; after two no-progress failures, stops. If the repo changed but verify failed, stops.

The runner stops on: max ticks reached (default 3), a blocked worker, verify failure with progress, or no progress after two attempts. On a non-clean exit it writes `HANDOFF.md` so the next session resumes.

## Optional review/fix subloop

With `--review`, after the build queue drains, the runner invokes the `nospec-trial` skill, reads the actionable-finding count from `REVIEW.md`, and if non-zero invokes the `nospec-mend` skill (which appends new pending units), then runs another build pass. Bounded by `--max-review-rounds` (default 2). The runner orchestrates stop conditions only; the `nospec-trial` and `nospec-mend` skills own judgment — what the findings are, which become units.

## Agent-agnostic

The runner doesn't know which agent it's driving. `LOOP_AGENT_CMD` overrides the worker invocation; `LOOP_REVIEW_CMD` / `LOOP_FIX_CMD` override review/fix. Per-unit `Agent:` lines override for one unit. Run `nospec run --help` for the full flag and environment reference.

## Disposability

`.loop/<name>/QUEUE.md`, `HANDOFF.md`, `REVIEW.md`, and `specs/` are disposable — delete them when the cycle completes. `.loop/<name>/EVIDENCE.md` is durable; keep it as the ledger of what each tick proved. The runner records a registry-derived proof boundary and a pin-state (which durable docs were touched) per tick; pin alerts in the ledger are triage triggers for `nospec-trial` → `nospec-curator`, not coherence gates.
