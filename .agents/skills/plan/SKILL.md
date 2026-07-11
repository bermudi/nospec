---
name: plan
description: Use when converting a software task, bug, cleanup, or vague human intent into a disposable `.loop/<name>/QUEUE.md` loop packet of verifiable work units. The planner picks a short, descriptive name for the cycle (e.g., `go-cli`, `parser-bug`).
metadata:
  version: "1.0.0"
---

# Plan

Convert intent into a `.loop/<name>/QUEUE.md` loop packet: a bounded queue of disposable, independently verifiable work units. Pick a short, descriptive name for the cycle (e.g., `go-cli`, `parser-bug`).

The goal is not to create durable specs. The goal is to give the loop runner work units that can be attempted one at a time and externally verified. When the work is done, the queue is deleted.

## Entry points

You may enter plan from two places:

- **After explore** — the codebase has been read, intent has been grilled, decisions may already be captured as ADRs in `decisions/`, terms may be in `glossary.md`. Use that context. Don't re-derive what explore already established.
- **Directly** — for small work where explore isn't needed. A bug fix, a patch, a small feature. Skip the ceremony.

If you're entering directly and the work is large or greenfield (no existing codebase to ground against), consider producing disposable specs first (see "Big work" below).

## When a queue already exists

If `.loop/<name>/QUEUE.md` already exists, read it before writing. If it is stale, doesn't match the current code, or no longer reflects the real goal, discard it and write a fresh queue. Plans are disposable coordination state; regenerating from the actual codebase is cheaper than salvaging a drifting plan.

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
- Its constraints state what must stay true or what is out of bounds — never what to edit. If a constraint names a file, it is "don't touch X" or "X's public API must not change," not "update X."
- It can leave the repo better if the loop stops immediately afterward.
- It does not depend on future units to have value.

A unit whose outcome cannot be captured by a deterministic `Verify:` command plus a short `Done means:` list is not ready. Vague shape plus weak verify gives the worker nothing to aim at.

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

## Design note

For any cycle where external constraints, non-obvious decisions, or spec compliance applies across multiple units, write `.loop/<name>/DESIGN.md` before decomposing into units. This is reasoning context — the *why* behind the work — that the codebase alone can't provide.

The trigger is not size. The trigger is: **is there context the worker can't get from reading the codebase?** Examples:

- An external spec or validator constrains the shape of the work (e.g., agentskills.io frontmatter rules enforced by `skills-ref`).
- A decision was made during explore that affects multiple units (e.g., "use `metadata.version` not top-level `version` because the spec rejects unknown top-level fields").
- A trade-off was resolved that the worker would otherwise have to re-derive (e.g., "we preserve old manifest hashes for modified skills because recomputing them erases the locally-customized signal").

The design note is one page. It states the constraints, the decisions, and the reasoning. It is **disposable** — consumed during build, then deleted with the rest of `.loop/<name>/`. Code is the source of truth. The note exists to give the worker, reviewer, and fixer the reasoning context they need to make judgment calls correctly.

Skip it for small work. A bug fix, a typo, a one-unit patch doesn't need a design note. If the work is greenfield or large, you may also produce `.loop/<name>/specs/proposal.md` (what and why, one page) for broader context — but the design note is the one that matters for judgment calls during build and fix.

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

Read first:
- .loop/<name>/DESIGN.md (if it exists — cycle-level reasoning context)
- <context the worker needs: ADR, area, or file>
- <2–4 entries; context, not scope>

Constraints:
- <boundary or guardrail>
- <what must stay true or what is out of bounds>
- <if it names a file, it is "don't touch X" or "X's public API must not change", not "update X">

Done means:
- <observable condition>
- <no regression condition>

Verify:
```bash
<command that exits 0 on success>
```

Status: pending

## <next outcome>
...
````

### Field notes

- **Header** is `## <outcome>` — no numbered prefix, no "Slice" word. The outcome itself is the title.
- **Agent:** is optional. Omit unless this unit needs a different model or command than the global `LOOP_AGENT_CMD`.
- **Why:** is optional. Fill in only when there's non-obvious context worth preserving. No padding.
- **Read first:** is context, not scope. Two to four entries: ADRs, code areas, or rulings. Prefer areas and rulings over file enumerations. If `.loop/<name>/DESIGN.md` exists, list it first — it carries the reasoning context the worker needs for judgment calls.
- **Constraints:** are boundaries. A constraint states what must stay true or what is out of bounds — never what to edit. If it names a file, it is "don't touch X" or "X's public API must not change," not "update X."
- **Done means:** is the acceptance criteria — what must be observably true after the unit.
- **Verify:** is the mechanically enforceable subset of `Done means:`. A unit whose outcome can't be captured by a deterministic verify command plus a short `Done means:` list isn't ready.
- **Status:** starts as `pending`. The loop updates it to `in_progress`, `done`, `verify_failed`, `no_progress`, or `blocked`.

## Disposability

`.loop/<name>/QUEUE.md` is disposable coordination state. When the work is done and verified, delete the cycle's `.loop/<name>/` subdirectory. What persists: code, tests, decisions in `decisions/`, glossary entries in `glossary.md`, ADRs, AGENTS.md. The queue is not an artifact — it's a scratchpad.
