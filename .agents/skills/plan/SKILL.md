---
name: plan
description: Use when converting a software task, bug, cleanup, or vague human intent into a disposable `.loop/QUEUE.md` loop packet of verifiable work units. Triggers on "plan this", "break this down", "how should we approach", "split this into steps", "decompose", "create a queue", "what's the work plan", or when intent needs to become independently verifiable units for a bounded agent loop.
---

# Plan

Convert intent into a `.loop/QUEUE.md` loop packet: a bounded queue of disposable, independently verifiable work units.

The goal is not to create durable specs. The goal is to give the loop runner work units that can be attempted one at a time and externally verified. When the work is done, the queue is deleted.

## Entry points

You may enter plan from two places:

- **After explore** — the codebase has been read, intent has been grilled, decisions may already be captured as ADRs in `decisions/`, terms may be in `glossary.md`. Use that context. Don't re-derive what explore already established.
- **Directly** — for small work where explore isn't needed. A bug fix, a patch, a small feature. Skip the ceremony.

If you're entering directly and the work is large or greenfield (no existing codebase to ground against), consider producing disposable specs first (see "Big work" below).

## When a queue already exists

If `.loop/QUEUE.md` already exists, read it before writing. If it is stale, doesn't match the current code, or no longer reflects the real goal, discard it and write a fresh queue. Plans are disposable coordination state; regenerating from the actual codebase is cheaper than salvaging a drifting plan.

## Work unit types

A work unit is whatever shape the work is. Pick the right type per unit:

- **vertical slice** — crosses enough layers to produce a user-visible or system-visible improvement. The default preference.
- **patch** — small, localized fix. One change, one verify.
- **investigation** — produces findings or ADRs, not necessarily code. Verify checks that the findings exist and are recorded.
- **bug fix** — reproduce → fix → verify. The verify command must fail before the fix and pass after.
- **refactor** — restructure without behavior change. Verify checks that existing tests still pass.

"Vertical slice" is the preferred default, not a required format. The planner prefers slices and rejects horizontal phases, but a unit can be a patch, investigation, or bug fix when the work genuinely isn't sliceable.

## Planning procedure

1. Restate the user's goal as an observable outcome.
2. Identify the strongest deterministic verification command available now.
3. Split work into units that each leave the repo better if the loop stops immediately after.
4. Prefer vertical slices. Reject horizontal phases (see below).
5. For every unit, pick its type from the taxonomy above.
6. Keep the queue short enough for bounded execution (prefer 2-5 units) and within the runner's hard stops (max ticks, no-progress detection). If the work is larger, the loop will pause and write a handoff; the next session resumes from there.
7. If no deterministic verification exists, make the first unit create or identify one.
8. If a decision crystallizes during decomposition, capture it as an ADR (use the `decide` skill).
9. If domain terms are ambiguous or inconsistent, define them (use the `domain-modeling` skill).

## Valid work unit test

A work unit is valid only if all are true:

- It has one observable outcome.
- It has one verification command.
- That verification command is deterministic and executable by the runner (not an LLM-as-judge).
- It has a narrow scope.
- It can leave the repo better if the loop stops immediately afterward.
- It does not depend on future units to have value.

## Horizontal phase rejection

Reject and rewrite units named after layers or activities:

- "Types"
- "CLI wiring"
- "Backend"
- "Frontend"
- "Tests"
- "Refactor"
- "Verification phase"
- "Implement all X"

Rewrite them as end-to-end outcomes:

- Bad: `Add JSON result types`
- Good: `validate --json reports one existing validation error as machine-readable JSON`

- Bad: `Wire CLI flag`
- Good: `validate --json and text mode report the same broken-link error on the same fixture`

- Bad: `Write tests`
- Good: `the regression fixture fails before the fix and passes after the fix`

This is a heuristic, not a gate. If a unit genuinely can't be vertical (a patch, an investigation), use the right type instead.

## Big work

For greenfield or large work where there's no existing code to ground against, optionally produce disposable planning artifacts before decomposing into units:

- `.loop/specs/proposal.md` — what we're building and why. One page.
- `.loop/specs/design.md` — architecture sketch. How the pieces fit. One page.

These are **disposable**. They're consumed during build, then deleted. They are never canonized, never merged, never treated as source of truth. Code is the source of truth. These exist only to help the agent think before the code exists.

Skip this entirely for small work. A bug fix doesn't need a proposal.

## Queue template

Use this exact structure unless the target repo already has a loop convention:

````markdown
# Loop Queue: <short name>

Goal:
<one paragraph describing the desired end state>

Stop condition:
`<command that proves the whole packet is done, if one exists>`

## <outcome — what changes, observable>

Agent: <optional — overrides LOOP_AGENT_CMD for this unit only>

Why:
<only if non-obvious — else omit>

Work:
- <narrow work instruction>
- <guardrail>

Verify:
```bash
<command that exits 0 on success>
```

Done means:
- <observable condition>
- <no regression condition>

Status: pending

## <next outcome>
...
````

### Field notes

- **Header** is `## <outcome>` — no numbered prefix, no "Slice" word. The outcome itself is the title.
- **Agent:** is optional. Omit unless this unit needs a different model or command than the global `LOOP_AGENT_CMD`.
- **Why:** is optional. Fill in only when there's non-obvious context worth preserving. No padding.
- **Verify:** is the load-bearing field. A unit without a verify command is not loop-ready.
- **Status:** starts as `pending`. The loop updates it to `in_progress`, `done`, `verify_failed`, `no_progress`, or `blocked`.

## Disposability

`.loop/QUEUE.md` is disposable coordination state. When the work is done and verified, delete the `.loop/` directory. What persists: code, tests, decisions in `decisions/`, glossary entries in `glossary.md`, ADRs, AGENTS.md. The queue is not an artifact — it's a scratchpad.
