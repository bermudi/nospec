---
name: decide
description: Use when an architectural ruling crystallizes — during exploration, planning, building, or review. Captures decisions as ADRs in `decisions/` so they persist after specs are deleted and code moves on. Triggers on "we decided to", "let's go with", "the ruling is", "why did we choose", "record this decision", or when a design choice with lasting consequences is made.
---

# Decide

Capture architectural rulings as ADRs. Decisions persist; specs and plans don't. A plan is disposable coordination state — throw it away when it drifts and regenerate from the codebase ([plan-disposability](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/plan-disposability.md)) — but a decision stays, explaining why the code is the way it is long after the code has moved on.

This is a shared skill — called inline by `explore`, `plan`, `build`, and `review` whenever a ruling crystallizes, in any attention-mode. It is not a separate phase. If you're mid-plan and a decision crystallizes, write the ADR now, then continue planning; don't queue it for later.

## When to write an ADR

Write an ADR when a decision has **lasting consequences** — it shapes future work and won't be obvious from the code alone:

- Choosing a library, framework, or pattern over alternatives
- Picking an architecture shape (monolith vs services, sync vs async, etc.)
- Deciding a convention (naming, file layout, error handling strategy)
- Resolving a tension between competing approaches
- Ruling something out — sometimes "we won't do X" is the decision

Don't write one for implementation details obvious from the code, choices with no real alternatives, or decisions that only matter for the current work unit.

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

<What we chose, stated as a ruling. "We will X" not "we considered X".>

## Consequences

<What this makes easier, what it makes harder, what to watch out for.>
````

NNNN is a zero-padded sequence number — look at `decisions/` for the next available one; if the directory doesn't exist, create it and start at `0001`.

An ADR is *active* unless its `Status` is `superseded`. Superseding is a two-step mutual link: mark the old ADR `Status: superseded` and add `Superseded by: ADR-NNNN`; write the new ADR with `Supersedes: ADR-NNNN` pointing back. Both sides must reference each other — a one-sided link is a broken chain; fix it before moving on.

## Orphan-ADR hygiene

An ADR is *orphaned* if nothing references it — neither current work nor a record of completed work. In a batch cycle that's `QUEUE.md` (in flight) and `EVIDENCE.md` (completed); interactively, judge it against the change in flight and the codebase as it stands. When you write an ADR, thread a reference to it through the work so it stays covered; when you supersede one, carry that reference into the replacement.

This is judgment, not a gate — read `decisions/` against what's actually being done and notice orphans yourself.

## What makes a good ADR

- **Title is the ruling.** "Use SQLite for local state" not "Storage decision."
- **Context names the alternatives.** If there were no real alternatives, it's not a decision.
- **Decision is a statement, not a discussion.** "We will X because Y."
- **Consequences are honest.** Include what gets harder, not just what gets easier.
- **Short enough to read in 30 seconds.** If it's longer, the decision was probably compound — split it.

## Delegation

Other skills call this skill inline — they don't hand off to it. Record the ruling where it crystallized; don't defer it.
