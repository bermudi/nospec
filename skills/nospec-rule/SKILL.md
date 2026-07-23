---
name: nospec-rule
description: Use when an architectural ruling crystallizes — during exploration, planning, building, or review. Captures decisions as ADRs in `decisions/` so they persist after specs are deleted and code moves on. Triggers on "we decided to", "let's go with", "the ruling is", "why did we choose", "record this decision", or when a design choice with lasting consequences is made.
---

# Decide

Capture architectural rulings as ADRs. Decisions persist; specs and plans don't. A plan is disposable coordination state — throw it away when it drifts and regenerate from the codebase (plan-disposability) — but a decision stays, explaining why the code is the way it is long after the code has moved on.

This is a shared skill — called inline by `nospec-scout`, `nospec-shape`, `nospec-hew`, and `nospec-trial` whenever a ruling crystallizes, in any attention-mode. It is not a separate phase. If you're mid-plan and a decision crystallizes, write the ADR now, then continue planning; don't queue it for later.

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
---
id: NNNN
date: <YYYY-MM-DD>
status: accepted | superseded
spine: false
supersedes: [NNNN]          # only if this replaces an earlier ADR
superseded_by: [NNNN]       # only if a later ADR replaces this one
amends: [NNNN]              # only if this narrows or extends an earlier ADR
builds_on: [NNNN]           # only if this derives from but doesn't change an earlier ADR
grandfathered: "<reason>"   # only for pre-ledger ADRs (rare)
---

# NNNN: <title — the ruling, not the topic>

## Context

<What problem were we solving? What constraints were in play? What alternatives were considered?>

## Decision

<What we chose, stated as a ruling. "We will X" not "we considered X".>

## Consequences

<What this makes easier, what it makes harder, what to watch out for.>
````

NNNN is a zero-padded sequence number — look at `decisions/` for the next available one; if the directory doesn't exist, create it and start at `0001`.

The `spine` field marks whether this ADR is load-bearing — part of the curated spine that defines what the project is and how it works. Most ADRs are `spine: false`; only rulings that reshape the project's thesis get `spine: true`. The spine is derived from this field by `./nospec spine`; never re-list the spine in prose (ADR-0017).

An ADR is *active* unless its `status` is `superseded`. Superseding is a two-step mutual link: mark the old ADR `status: superseded` and add `superseded_by: [NNNN]`; write the new ADR with `supersedes: [NNNN]` pointing back. Both sides must reference each other — a one-sided link is a broken chain; fix it before moving on. The bilateral-link rule applies to `supersedes`/`superseded_by` only; `amends` and `builds_on` are one-directional (the amended ADR's status stays `accepted` and needs no back-reference).

## Orphan-ADR hygiene

An ADR is *orphaned* when it no longer explains or constrains the system — when nothing about the current code or ongoing work depends on its ruling. A negative ruling ("we will not build X") can be alive with no citing work at all; a convention can constrain future work without driving an immediate change. References — a `QUEUE.md` or `EVIDENCE.md` in a batch cycle, the change in flight, the codebase, the docs — are *evidence* of relevance, not the definition (ADR-0012).

When you write an ADR, thread a reference to it through the work; that creates evidence of its relevance. When you supersede one, carry that reference into the replacement. This is judgment, not a gate — read `decisions/` against the system as it stands and notice orphans yourself.

After superseding an ADR, invoke the `document` skill to check which views still cite the old ruling and need projection updates. A superseded ADR is a common source of stale documentation.

## What makes a good ADR

- **Title is the ruling.** "Use SQLite for local state" not "Storage decision."
- **Context names the alternatives.** If there were no real alternatives, it's not a decision.
- **Decision is a statement, not a discussion.** "We will X because Y."
- **Consequences are honest.** Include what gets harder, not just what gets easier.
- **Bias toward ADRs that can be understood quickly.** Split one only when it contains independently changeable rulings — not to hit a length target. Some single rulings genuinely need substantial context and consequences.

## Delegation

Other skills call this skill inline — they don't hand off to it. Record the ruling where it crystallized; don't defer it.
