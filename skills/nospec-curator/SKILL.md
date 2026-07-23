---
name: nospec-curator
description: Use when deciding where durable knowledge belongs, detecting contradictions across durable artifacts, or maintaining the views that project authoritative records. Triggers on "where does this go", "document this", "are these docs consistent", "this is stale", "update the docs", or when a ruling or interface change may invalidate existing documentation.
---

# Document

Route knowledge to the artifact that should own it, maintain the views that project authoritative records, and spot contradictions across durable context. This is about artifact hygiene, not about implementing features.

Durable knowledge rots when it has no clear owner or when several documents make the same claim independently. The failure mode is not length; it is **overlapping authority**. A build command described in three places will be wrong in at least one of them after the next change.

## Artifact roles

Every durable artifact plays one of these roles. The specific files below are examples — a project may use different names or only some of them. What matters is the role each artifact plays, not the filename.

- **Record** — owns a class of claim.
  - procedural-knowledge files (e.g., `skills/`)
  - architectural rulings (e.g., `decisions/`)
  - domain vocabulary (e.g., `glossary.md`)
  - operational context (e.g., `AGENTS.md`)
  - code and tests — implemented behavior
- **View** — helps a reader understand records together, without becoming an alternate authority.
  - entry point (e.g., `README.md`)
  - conceptual overview (e.g., `docs/architecture.md`)
  - first-use guide (e.g., `docs/getting-started.md`)
  - reference pages (e.g., `docs/<topic>.md`)
- **Ledger** — append-only record of what happened.
  - cycle evidence (e.g., `.loop/<name>/EVIDENCE.md`)
- **Work state** — coordination state consumed then discarded.
  - queues, handoffs, review artifacts, scratch specs (e.g., `.loop/<name>/QUEUE.md`, `HANDOFF.md`, `REVIEW.md`, `specs/`)

The rule is not that each fact appears once. It is that each fact has one owner; everything else is a deliberate projection.

## What document does

These are concepts, not a script. Use the ones that fit the situation.

### Place the claim

When new knowledge appears, ask: what question will a future reader ask, and which record should answer it?

- A new convention → operational-context record (e.g., `AGENTS.md`)
- A new architectural ruling → rulings record (e.g., `decisions/`)
- A new domain term or refined meaning → domain-vocabulary record (e.g., `glossary.md`)
- A lesson the codebase taught about the problem → a durable record (ADR, `glossary.md`, or a new artifact) when one actually appears
- A new workflow concept → the relevant procedural-knowledge file (e.g., `skills/<name>/SKILL.md`)
- A user-facing protocol or interface → a reference page (e.g., `docs/<reference>.md`)
- The current behavior of the code → code/tests

If the project does not have a record for a class of claim, do not invent one preemptively — create it when the first claim of that kind actually appears. Do not place the same claim in two records. If a summary is needed, put the claim in the record and project it in a view.

### Maintain the projection

A view is useful only if it points to the right record and stays aligned.

- When a record changes, update the views that summarize it.
- When a view contradicts its record, the record wins.
- When a view starts containing claims that belong in a record, move the claims and make the view link to them.
- Delete views that no longer serve a distinct purpose; duplication masquerading as help is still duplication.

### Detect contradictions

Coherence is the relationship between durable artifacts. It is not the same as compilation or passing tests. A repo can verify green while its durable docs contradict its rulings.

Some structural drift is mechanically detectable. If the project has a `nospec` CLI (ADR-0017), `./nospec check` catches re-enumerated spine lists, duplicate ownership claims across records, and missing frontmatter — the shapes that drift most reliably. Run it first to surface what's mechanically provable, then apply judgment for the rest.

Look for:

- A view describing a behavior the code no longer has
- A ruling cited as current when it has been superseded
- A domain term whose definition no longer matches usage
- Operational-context instructions that no longer work
- A skill advising a workflow the rulings have retired
- Two documents answering the same question differently

When you find one, route the correction to the owning record and update its projections.

### Retire, don't archive

Superseded or duplicated explanations do not need to be kept for posterity. ADRs remain because rulings have historical value. Repetitive docs and stale guides do not. Delete them. Git preserves the history if it ever matters.

## Reasoned defaults

- **Default to placing knowledge where it is authoritative.** Override: during exploration, a temporary note is fine; move it before the work is done.
- **Default to updating views after a record changes.** Override: a trivial wording change that does not affect meaning can skip a view update.
- **Default to checking coherence after an ADR is superseded or a public interface changes.** Override: skip for purely internal changes that no view or skill describes.
- **Default to scoping from the diff when invoked from a pin alert.** A pin alert says a durable doc changed since a prior cycle pinned it. Scope the coherence check from the diff of the pinned file — what changed, and which other durable docs describe or depend on it? The alert is a trigger, not a finding; it says "something moved," not "coherence broke." Judge whether the move left stale projections. Override: skip if the change is trivial (whitespace, formatting) and no projection references the changed content.
- **Default to deleting duplicated or superseded docs.** Override: keep a migration note only if users of a published version still need it.

## What document is not

- **Not a prose-writing skill.** It does not generate documentation for its own sake. It decides where claims live.
- **Not a semantic validator.** Coherence is judgment, not a command that exits 0. Mechanical checks can catch broken links and missing paths, but they cannot prove that a view correctly projects its record.
- **Not a separate phase.** Invoke it when placement or coherence is a concern, not as ceremony before every change.
- **Not a replacement for `nospec-rule` or `nospec-lexicon`.** `nospec-rule` captures rulings; `nospec-lexicon` captures terms. `nospec-curator` ensures those captures project cleanly into the rest of the durable context.

## Delegation

Other skills call `nospec-curator` inline when durable context changes:

- `nospec-rule` — after superseding a ruling, check which views still cite the old ruling and need projection updates.
- `nospec-lexicon` — after adding or changing a term, check whether the operational context, skills, or other views use the old meaning in a projection.
- `nospec-hew` — when an operational learning surfaces, route it to the operational-context record or the insights ledger and update any view that quotes it.
- `nospec-trial` — under the standards axis, when a change affects a public interface or a ruling, flag projection drift as a standards finding and invoke `nospec-curator` to assess coherence. Also route `Pin alerts:` from `EVIDENCE.md` — each alert is a durable doc that moved since a prior cycle pinned it, and `nospec-curator` scopes the coherence check from the diff.
