---
name: decide
description: Use when an architectural ruling crystallizes during any phase — explore, plan, build, or review. Captures decisions as ADRs in `decisions/` so they persist after specs are deleted and code moves on. Triggers on "we decided to", "let's go with", "the ruling is", "why did we choose", "record this decision", or when a design choice with lasting consequences is made.
---

# Decide

Capture architectural rulings as ADRs. Decisions persist; specs don't. A decision made six months ago still explains *why* the code is the way it is, even after the code has moved on.

This is a shared skill — called by explore, plan, build, and review whenever a ruling crystallizes. It is not a separate phase.

## When to write an ADR

Write an ADR when a decision has **lasting consequences** — it shapes future work and won't be obvious from the code alone:

- Choosing a library, framework, or pattern over alternatives
- Picking an architecture shape (monolith vs services, sync vs async, etc.)
- Deciding a convention (naming, file layout, error handling strategy)
- Resolving a tension between competing approaches
- Ruling something out (sometimes "we won't do X" is the decision)

Do **not** write an ADR for:
- Implementation details that are obvious from the code
- Choices with no real alternatives
- Decisions that only matter for the current work unit

## ADR format

````markdown
# NNNN: <title — the ruling, not the topic>

Date: <YYYY-MM-DD>
Status: accepted | superseded
Supersedes: ADR-NNNN      # only if this replaces an earlier ADR
Superseded by: ADR-NNNN   # only if a later ADR replaces this one

## Context

<What problem were we solving? What constraints were in play? What alternatives were considered?>

## Decision

<What we chose, stated as a ruling. "We will X" not "We considered X".>

## Consequences

<What this makes easier, what this makes harder, what to watch out for.>
````

NNNN is a zero-padded sequence number. Look at `decisions/` for the next available number. If the directory doesn't exist, create it and start at `0001`.

An ADR is *active* unless its `Status` is `superseded` (or it carries a `Superseded by:` line). `knack decisions check` treats superseded ADRs as exempt from the orphan gate — their replacement carries coverage.

## Procedure

1. Recognize that a ruling just crystallized. If you're in explore, plan, build, or review and you just resolved a tension with lasting consequences, that's a decision.
2. Check `decisions/` for existing ADRs on the same topic. If one exists and the ruling changed, supersede it. This is a two-step link:
   - Mark the old ADR `Status: superseded` and add `Superseded by: ADR-NNNN`.
   - Write the new ADR with `Supersedes: ADR-NNNN` pointing back to the old one.
   - The link must be mutual — both sides reference each other, or `knack decisions check` flags a broken chain.
3. Pick the next sequence number.
4. Write the ADR using the format above. The title is the ruling itself ("Use SQLite for local state" not "Database choice").
5. Keep it short. One page. The value is in the *why*, not the *what* — the code already shows the what.

## What makes a good ADR

- **Title is the ruling.** "Use SQLite for local state" not "Storage decision."
- **Context names the alternatives.** If there were no real alternatives, it's not a decision.
- **Decision is a statement, not a discussion.** "We will X because Y."
- **Consequences are honest.** Include what gets harder, not just what gets easier.
- **Short enough to read in 30 seconds.** If it's longer, the decision was probably compound — split it.

## Delegation

Other skills call this skill inline — they don't hand off to it. If you're in the middle of planning and a decision crystallizes, write the ADR now, then continue planning. Don't queue it for later.
