---
name: vertical-slice-planner
description: Use when converting a software task, bug, cleanup, or vague human intent into a disposable .loop/QUEUE.md loop packet made of independently verifiable vertical slices for a bounded agent loop. Triggers on planning loop work, creating a slice queue, decomposing work for sliceloop, or replacing phased/horizontal tasks with vertical slices.
---

# Vertical Slice Planner

Convert messy intent into a `.loop/QUEUE.md` loop packet: a bounded queue of disposable, independently verifiable vertical slices.

The goal is not to create durable specs. The goal is to give a loop runner work units that can be attempted one at a time and externally verified.

## Planning procedure

1. Restate the user's goal as an observable outcome.
2. Identify the strongest deterministic verification command available now.
3. Split work into slices that each cross enough layers to produce a user-visible or system-visible improvement.
4. For every slice, write why it is vertical.
5. Reject horizontal phases before writing the queue.
6. Keep the queue short enough for bounded execution. Prefer 2-5 slices.
7. If no deterministic verification exists, make the first slice create or identify one.

## Valid slice test

A slice is valid only if all are true:

- It has one observable outcome.
- It has one verification command.
- It has a narrow scope.
- It explains why it is vertical.
- It can leave the repo better if the loop stops immediately afterward.
- It does not depend on future slices to have value.

## Invalid slice patterns

Reject and rewrite slices named after layers or activities:

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
- Good: ``validate --json reports one existing validation error as machine-readable JSON``

- Bad: `Wire CLI flag`
- Good: ``validate --json and text mode report the same broken-link error on the same fixture``

- Bad: `Write tests`
- Good: `the regression fixture fails before the fix and passes after the fix`

## Queue template

Use this exact structure unless the target repo already has a sliceloop convention:

````markdown
# Loop Queue: <short name>

Goal:
<one paragraph describing the desired end state>

Stop condition:
`<command that proves the whole packet is done, if one exists>`

## Slice 1: <vertical outcome>

Why this is vertical:
<explain why this crosses enough layers to produce a real observable improvement>

Work:
- <narrow work instruction>
- <guardrail>
- <what not to broaden into>

Verify:
```bash
<command>
```

Done means:
- <observable condition>
- <no regression condition>

Status: pending

## Slice 2: <vertical outcome>

Why this is vertical:
<explanation>

Work:
- <instruction>

Verify:
```bash
<command>
```

Done means:
- <condition>

Status: pending
````

## Guardrails

- Do not create proposal/design/tasks/archive artifacts.
- Do not invent a durable canon for slices.
- Do not make a slice whose only output is scaffolding unless that scaffolding is immediately verified and useful.
- Do not rely on the worker to judge success.
- Do not include more process than the runner can act on.
- Prefer one strong verify command over many weak checks.

## Final answer when planning in chat

Present:

1. The proposed queue.
2. The verification command(s) it relies on.
3. Any weak point where verification is missing or judgment-based.
4. Ask whether to write it to `.loop/QUEUE.md`.
