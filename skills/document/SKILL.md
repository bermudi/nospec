---
name: document
description: Use when deciding where durable knowledge belongs, detecting contradictions across durable artifacts, or maintaining the views that project authoritative records. Triggers on "where does this go", "document this", "are these docs consistent", "this is stale", "update the docs", or when a ruling or interface change may invalidate existing documentation.
---

# Document

Route knowledge to the artifact that should own it, maintain the views that project authoritative records, and spot contradictions across durable context. This is about artifact hygiene, not about implementing features.

Durable knowledge rots when it has no clear owner or when several documents make the same claim independently. The failure mode is not length; it is **overlapping authority**. A build command described in three places will be wrong in at least one of them after the next change.

## Artifact roles

Every durable artifact plays one of these roles:

- **Record** — owns a class of claim.
  - `skills/` — procedural knowledge
  - `decisions/` — architectural rulings
  - `glossary.md` — domain terms
  - `AGENTS.md` — operational context
  - `LEARNINGS.md` — domain/problem observations
  - code and tests — implemented behavior
- **View** — helps a reader understand records together, without becoming an alternate authority.
  - `README.md` — entry point
  - `docs/architecture.md` — conceptual shape
  - `docs/getting-started.md` — first-use guide
  - `docs/skills.md`, `docs/loop.md`, `docs/queue-format.md` — reference views
- **Ledger** — append-only record of what happened.
  - `.loop/<name>/EVIDENCE.md` — cycle evidence
  - `LEARNINGS.md` — durable insights
- **Work state** — coordination state consumed then discarded.
  - `.loop/<name>/QUEUE.md`, `HANDOFF.md`, `REVIEW.md`, `specs/`

The rule is not that each fact appears once. It is that each fact has one owner; everything else is a deliberate projection.

## What document does

These are concepts, not a script. Use the ones that fit the situation.

### Place the claim

When new knowledge appears, ask: what question will a future reader ask, and which record should answer it?

- A new convention → `AGENTS.md`
- A new architectural ruling → `decisions/`
- A new domain term or refined meaning → `glossary.md`
- A lesson the codebase taught about the problem → `LEARNINGS.md`
- A new workflow concept → the relevant `skills/` file
- A user-facing protocol or interface → `docs/<reference>.md`
- The current behavior of the code → code/tests

Do not place the same claim in two records. If a summary is needed, put the claim in the record and project it in a view.

### Maintain the projection

A view is useful only if it points to the right record and stays aligned.

- When a record changes, update the views that summarize it.
- When a view contradicts its record, the record wins.
- When a view starts containing claims that belong in a record, move the claims and make the view link to them.
- Delete views that no longer serve a distinct purpose; duplication masquerading as help is still duplication.

### Detect contradictions

Coherence is the relationship between durable artifacts. It is not the same as compilation or passing tests. A repo can verify green while its durable docs contradict its rulings.

Look for:

- A view describing a behavior the code no longer has
- An ADR cited as current when it has been superseded
- A glossary term whose definition no longer matches usage
- `AGENTS.md` instructions that no longer work
- A skill advising a workflow the ADRs have retired
- Two documents answering the same question differently

When you find one, route the correction to the owning record and update its projections.

### Retire, don't archive

Superseded or duplicated explanations do not need to be kept for posterity. ADRs remain because rulings have historical value. Repetitive docs and stale guides do not. Delete them. Git preserves the history if it ever matters.

## Reasoned defaults

- **Default to placing knowledge where it is authoritative.** Override: during exploration, a temporary note is fine; move it before the work is done.
- **Default to updating views after a record changes.** Override: a trivial wording change that does not affect meaning can skip a view update.
- **Default to checking coherence after an ADR is superseded or a public interface changes.** Override: skip for purely internal changes that no view or skill describes.
- **Default to deleting duplicated or superseded docs.** Override: keep a migration note only if users of a published version still need it.

## What document is not

- **Not a prose-writing skill.** It does not generate documentation for its own sake. It decides where claims live.
- **Not a semantic validator.** Coherence is judgment, not a command that exits 0. Mechanical checks can catch broken links and missing paths, but they cannot prove that a view correctly projects its record.
- **Not a separate phase.** Invoke it when placement or coherence is a concern, not as ceremony before every change.
- **Not a replacement for `decide` or `domain-modeling`.** `decide` captures rulings; `domain-modeling` captures terms. `document` ensures those captures project cleanly into the rest of the durable context.

## Delegation

Other skills call `document` inline when durable context changes:

- `decide` — after superseding an ADR, check which views still cite the old ruling and need projection updates.
- `domain-modeling` — after adding or changing a term, check whether `glossary.md`, `AGENTS.md`, or skills use the old meaning in a projection.
- `build` — when an operational learning surfaces, route it to `AGENTS.md` or `LEARNINGS.md` and update any view that quotes it.
- `review` — under the standards axis, when a change affects a public interface or a ruling, flag projection drift as a standards finding and invoke `document` to assess coherence.
