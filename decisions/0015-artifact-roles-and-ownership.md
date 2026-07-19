---
id: 0015
date: 2026-07-18
status: accepted
spine: true
---

# 0015: Artifact roles and ownership for durable knowledge

## Context

ADRs 0009, 0010, 0011, 0013, and 0014 established that knack's durable knowledge is partitioned: code owns behavior, ADRs own rulings, glossary owns terms, `AGENTS.md` owns operational context, `LEARNINGS.md` owns domain insights, and skills own procedural knowledge. The repo did not have a corresponding structure for organizing that knowledge, so durable documents kept accumulating overlapping responsibilities.

`DESIGN.md` tried to be thesis, architecture, interface reference, file-format spec, migration history, and wiki index at once. It became stale in places while remaining too large to maintain. `docs/README.md`, `docs/getting-started.md`, `docs/faq.md`, `docs/skills.md`, `docs/loop.md`, and `docs/queue-format.md` duplicated `DESIGN.md` sections and, after ADR-0011 deleted the Go CLI, still referenced the removed CLI in many places. The README declared `DESIGN.md` partially stale even after `DESIGN.md` had been reworked, because ownership of the "is this still accurate?" check was unclear.

The problem is recurring. Any project installing knack will accumulate durable context. Without an artifact model, the same bloat will reappear: READMEs that explain everything, docs that contradict skills, ADRs that no one projects into guides, and instructions that survive the features they describe.

## Decision

Durable knowledge is organized by **artifact role**, not by topic or length:

- **Record** — owns a class of claim.
  - `skills/<name>/SKILL.md` — procedural knowledge
  - `decisions/` — architectural rulings
  - `glossary.md` — domain terms
  - `AGENTS.md` — operational context
  - code/tests — implemented behavior
- **View** — helps a reader understand several records together, but is not itself authoritative.
  - `README.md` — entry-point view
  - `docs/architecture.md` — conceptual shape
  - `docs/getting-started.md` — first-use view
  - `docs/skills.md` — skill catalog
  - `docs/loop.md` and `docs/queue-format.md` — reference views
  - `docs/README.md` — documentation ownership map
- **Ledger** — append-only record of what happened.
  - `.loop/<name>/EVIDENCE.md` — cycle evidence
  - `LEARNINGS.md` — domain/problem insights
- **Work state** — coordination state consumed then discarded.
  - `.loop/<name>/QUEUE.md`, `HANDOFF.md`, `REVIEW.md`, `specs/`

The governing rule:

> One claim has one owner; other documents are deliberate projections of it.

Views may summarize and link, but they do not independently redefine what they project. When a record changes, its projections must be reconciled. When a view contradicts its record, the record wins.

A new shared skill, **`document`**, teaches artifact placement, authority boundaries, lifecycle awareness, and coherence maintenance. It is invoked when knowledge needs a home, when durable artifacts appear to contradict one another, or when a ruling invalidates existing documentation.

## Consequences

- `DESIGN.md` is deleted and its content distributed to the appropriate records or views. `docs/architecture.md` becomes the conceptual overview.
- `docs/faq.md` is deleted; its questions are answered by the owning view or record.
- `docs/README.md` becomes an ownership map: "if you want X, read Y."
- `skills/document/SKILL.md` is added to the product.
- `build`, `review`, `decide`, and `domain-modeling` delegate durable-context placement and coherence checks to `document`.
- Mechanical checks (broken internal links, references to deleted interfaces) may be added to `tests/run.sh`, but they do not claim to prove semantic coherence.

## Related

- ADR-0009 — skills are the product; loop is optional
- ADR-0010 — skills transmit concepts, not rules
- ADR-0011 — ship via skills.sh; delete CLI
- ADR-0013 — wiki links in docs, not skill text
- ADR-0014 — durability is maintenance, not permanence
