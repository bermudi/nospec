---
name: explore
description: Use when investigating a codebase, grilling intent, or stress-testing ideas before planning work. The upstream phase of the loop — read code, challenge assumptions, surface constraints, and capture decisions as they crystallize. Triggers on "explore", "investigate", "let me understand", "what's going on with", "grill this", "stress-test", "help me decide", "what should I do next", or when the real problem isn't clear yet.
---

# Explore

Read the codebase, grill the intent, stress-test ideas. The goal is to reach clarity before planning work units. No artifacts produced except ADRs and glossary entries written inline when decisions crystallize.

Pure conversation. The output of explore is a clearer head, not a file on disk.

## What explore is not

- Not a planning phase — that's `plan`. Explore doesn't write QUEUE.md.
- Not a spec phase — no proposals, no designs. Those are disposable and belong in `plan` if needed at all.
- Not a code phase — no edits. Read-only.

Explore is the phase where you figure out what the problem actually is.

## Before you start

If the project already has loop state, read it first:

- `.loop/HANDOFF.md` — what the previous session left behind
- `.loop/QUEUE.md` — any pending work units
- `AGENTS.md`, `decisions/`, `glossary.md` — durable project context

These are part of the shared design concept. The goal of explore is to get you and the codebase on the same theory before any work is planned.

## Procedure

1. **Read the codebase.** Understand the current state. What exists? What patterns are in use? What conventions does the repo follow? Read `AGENTS.md` if it exists — it's operational context.

2. **Grill the intent.** The user's stated goal may not be the real goal. Ask:
   - What outcome are they trying to achieve?
   - What's the constraint that's actually binding?
   - Is the stated problem the real problem, or a symptom?
   - What would "done" look like to them?

3. **Stress-test ideas.** Before committing to an approach, poke holes:
   - What breaks if we do it this way?
   - What's the simplest thing that could work?
   - What are we assuming that might be wrong?
   - Is there an existing pattern in the codebase we should follow?

4. **Capture decisions inline.** If a ruling crystallizes during exploration — "we'll use X because Y" — write an ADR now using the `decide` skill. Don't queue it for later. The decision was made during exploration; record it during exploration.

5. **Define terms inline.** If domain terms are ambiguous or inconsistent, define them now using the `domain-modeling` skill.

6. **Hand off to plan.** When the intent is clear and the codebase is understood, the next step is `plan` — decompose into verifiable work units. Don't do that here.

## When to skip explore

- Small fixes where the problem is already clear
- Work where you already know the codebase
- Bug reports that already specify the reproduction

Don't force exploration on work that doesn't need it. The flow is composable: `bug → plan → build → done` is valid.

## What to read

- `AGENTS.md` — operational context, working conventions
- `decisions/` — existing ADRs that constrain the solution space
- `glossary.md` — shared vocabulary, if it exists
- The codebase itself — focus on the areas the task touches
- Existing tests — they show what the code is supposed to do

## Output

The output of explore is clarity, communicated to the user. Summarize:

- What loop state you read (handoff, queue, existing ADRs/glossary)
- What you found in the codebase
- What the real problem is (if different from the stated one)
- What approach you'd take and why
- What decisions you captured as ADRs
- What terms you defined in the glossary

Then suggest `plan` as the next step.
