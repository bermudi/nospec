---
name: explore
description: Use when investigating a codebase, grilling intent, or stress-testing ideas before planning or building work. Read code, challenge assumptions, surface the real problem, and capture decisions as they crystallize. Triggers on "explore", "investigate", "let me understand", "what's going on with", "grill this", "stress-test", "help me decide", "what should I do next", or when the real problem isn't clear yet.
---

# Explore

Read the codebase, grill the intent, stress-test ideas. The goal is clarity before work is planned or built — reaching the same theory as the codebase so the work that follows aims at the real problem, not a symptom of it. The output of explore is a clearer head, not a file on disk.

That core is the same whether you're exploring interactively with the human or preparing to leave a plan for unattended batch execution. Explore is a human-present activity: it produces understanding, not an artifact, so it has no mechanical verify and isn't a batch unit — it runs *before* batch, not in it. No edits land during explore except durable traces (an ADR, a glossary entry) written inline when a decision crystallizes.

The failure mode explore prevents is working on the wrong problem. A precisely executed solution to the wrong problem is still wrong, and most wasted effort comes from solving the wrong problem, not from solving the right one badly. The user's stated goal is often a guessed-at solution to an unspoken problem; explore settles the target before work is committed.

## What explore does

These are concepts, not a script — they interleave, not sequence. Read the work and decide how much of each it needs.

- **Read the codebase.** Not to gather facts, but to surface the decisions the code already embodies. Code carries rulings and constraints the surface — README, stated intent, even the user's description — doesn't state; reading it is how you reach the codebase's actual theory. Code clarifies spec runs in reverse here: if implementing code clarifies a spec, reading existing code clarifies the real intent. Read the durable context first — `AGENTS.md`, `decisions/`, `glossary.md`, `LEARNINGS.md`, and `.loop/<name>/EVIDENCE.md` if a cycle has run (it survives the queue's deletion, so it's the record of what was already proven) — then the code the task touches and the tests that show what it's supposed to do. If a cycle is in flight, read its `.loop/<name>/HANDOFF.md` and `QUEUE.md` too.

- **Grill the intent.** The user's stated goal may not be the real goal. Ask what outcome they're after, what's actually binding, whether the stated problem is the real one or a symptom, what "done" looks like to them. Settling the right problem is cheapest here, before any code moves.

- **Stress-test ideas.** Before committing to an approach, poke holes. What breaks if we do it this way? What's the simplest thing that could work? What are we assuming that might be wrong? Is there an existing pattern in the codebase to follow?

- **Capture decisions inline.** If a ruling crystallizes during exploration — "we'll use X because Y" — write the ADR now via the `decide` skill, not queued for later. If a domain term is ambiguous or inconsistent, define it now via `domain-modeling`. Decisions made during exploration are recorded during exploration.

## What explore is not

- Not planning — that's `plan`. Explore doesn't write `QUEUE.md`.
- Not a spec phase — no proposals or designs. Those are disposable and belong in `plan` if they're needed at all.
- Not a code phase — no edits. Read-only, except durable traces captured inline.

Explore is where you figure out what the problem actually is.

## Reasoned defaults

- **Default to exploring when the real problem isn't clear yet.** Override: skip it when the problem is already clear — a small fix, code you already know, a bug report with a reproduction. The skills are composable; reach for explore when clarity is the missing ingredient, not as a mandatory first step.
- **Communicate what you found.** The real problem if it differs from the stated one, the approach you'd take and why, the decisions and terms you captured. Then the next step is whatever the clarity points to — `plan` to decompose, direct edits, or nothing — not automatically `plan`.
