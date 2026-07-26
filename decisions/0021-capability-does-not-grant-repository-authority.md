---
nospec: true
id: 0021
date: 2026-07-25
status: accepted
spine: false
amends: [0009, 0016, 0017, 0019]
builds_on: [0010, 0015]
---

# 0021: Nospec capability does not grant repository authority

## Context

Dogfooding exposed the same authority error at two boundaries.

First, the installed `nospec check` applied Nospec's own frontmatter inventory to an unrelated repository's ordinary `README.md` and `AGENTS.md`. Making a command available had been mistaken for adopting its artifact schema.

Second, the skills told an agent to write an accepted ADR whenever a ruling “crystallized.” A recommendation could therefore become durable project policy without the decision owner accepting it, including during unattended work.

Both failures confuse capability with authority. Installation grants tools; agent participation grants labor. Neither grants ownership of repository documents or architectural decisions.

## Decision

Nospec acts only across explicit adoption boundaries.

- `nospec check` validates documents carrying the explicit `nospec: true` frontmatter marker. Generic metadata and untagged Markdown are ignored. Nospec's own required metadata inventory remains a source-test concern, where the repository can state exactly which files it owns.
- Pin-state records follow the same boundary: only changed Markdown artifacts marked `nospec: true` are pinned. Ordinary host-repository documents receive no implicit Nospec provenance semantics.
- An accepted ADR records a ruling made by the decision owner or under explicitly delegated authority.
- Recommendations and unresolved trade-offs remain conversational or disposable design material. A batch worker that discovers a new consequential choice blocks and hands it back instead of accepting its own proposal.

Queue execution remains explicit adoption: supplying a queue to `nospec run` authorizes the runner to mutate that queue's statuses and execute its commands, not to impose unrelated repository conventions.

## Consequences

- Installing Nospec is non-invasive: it does not make conventional repository documents malformed.
- A generic checker cannot detect a completely removed opt-in marker. Repositories that require an exact metadata inventory enforce it in their own tests or an explicit manifest; Nospec does not guess.
- ADR creation may require a human round trip. That is intentional when authority was not delegated.
- ADR-0017's metadata policy remains true for Nospec's own durable artifacts, while its distributed checker is scoped to declared metadata.
- ADR-0019's target-repository commands operate only on explicitly adopted artifacts.
