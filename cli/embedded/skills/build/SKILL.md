---
name: build
description: Use when implementing one work unit from a `.loop/<name>/QUEUE.md` queue. The worker-side skill — read the unit, do the work, don't self-certify, end with a handoff. Triggers on "build", "implement", "apply this unit", "do the work", "run the loop", or when a work unit needs to be executed. Also use when the loop invokes the worker for a tick.
---

# Build

Implement one work unit from `.loop/<name>/QUEUE.md`. Do the work. Don't self-certify. The loop runner owns the verify gate — your job is to make the repository state satisfy the unit's `Verify` command, not to claim success.

## Core rules

The runner injects `prompts/worker.md` at the start of every tick. If it is not in your context, read it now. Its `Rules` and `Output` sections are the canonical source for this skill; the sections below elaborate on decisions, operational learnings, blockers, and units that are too large for one tick.

> **Scope note:** Updating `AGENTS.md` or writing an ADR during a tick is a durable trace, not part of the unit's `Verify` scope. Do it only when the tick teaches you something that would save the next session time.

## Scope

The unit's scope is its outcome plus its constraints. The worker determines which files to change and how. The `Verify:` command is the mechanically enforceable subset of `Done means:`.

## Capturing decisions during build

If you discover an architectural ruling while implementing — "we need to handle X this way because Y" — capture it as an ADR using the `decide` skill. Do this inline, not after the unit. Decisions made during implementation are the most valuable kind because they come from the code pushing back.

## Capturing operational learnings

If you learn something about how the project works, capture it in the right durable file:

- **Operational gotchas** — build commands, test conventions, loop behavior, how to verify work — go in `AGENTS.md`. That's the living operational context.
- **Domain or problem insights** — "X doesn't work because Y", "the parser has this surprising property" — go in `LEARNINGS.md`. That's the durable learning ledger.

Don't add trivia. Add things that would have saved you time if you'd known them upfront.

## When the unit is too big

If the work unit is larger than what can be done in one tick:

- Do as much as you can while keeping the repo in a working state.
- End with a handoff note explaining what remains.
- The runner will re-queue the unit if verify fails.

Do not try to do everything in one tick if the work genuinely needs more. The loop is designed for multiple ticks.

## When you're blocked

If you hit a blocker you can't resolve within the unit's scope:

1. State the blocker clearly in your final output.
2. Note what would unblock you (a decision, a dependency, a missing file).
3. Stop. Don't thrash.

The runner will mark the unit as `blocked` and write a handoff. The next session can pick up from there.

## Output

The output format is defined in the worker prompt (`prompts/worker.md`). End with a compact terminal handoff there.
