---
name: plan
description: Use when converting a software task, bug, cleanup, or vague human intent into a disposable .loop/QUEUE.md loop packet made of independently verifiable work units for a bounded agent loop. Triggers on planning loop work, creating a work queue, decomposing work for the loop, or replacing phased/horizontal tasks with verifiable units.
---

# Plan

Convert messy intent into a `.loop/QUEUE.md` loop packet: a bounded queue of disposable, independently verifiable work units.

The goal is not to create durable specs. The goal is to give the loop runner work units that can be attempted one at a time and externally verified.

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
6. Keep the queue short enough for bounded execution. Prefer 2-5 units.
7. If no deterministic verification exists, make the first unit create or identify one.

## Valid work unit test

A work unit is valid only if all are true:

- It has one observable outcome.
- It has one verification command.
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
