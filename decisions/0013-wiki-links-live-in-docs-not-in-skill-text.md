---
id: 0013
date: 2026-07-13
status: accepted
spine: true
amends: [0010]
---

# 0013: Wiki links live in docs, not in skill text

## Context

ADR-0010 said each concept a skill teaches carries "a pointer to depth — the wiki link." Skills were written with inline markdown links to the AgenticWiki, and over time the same convention spread to other repo artifacts without a clear rule for where links belong.

The root question is what a link *does*. The **synopsis is the payload** — it primes the model's prior on the concept (the wiki's leading-words mechanism) and is self-sufficient. The **link is provenance** — a citation for a human or reviewer, and an optional path to depth; the agent is not expected to fetch the page. (Confirmed empirically this session: two reviews caught a misused link and a too-thin synopsis precisely because the link made the synopsis auditable against the wiki.)

The first cut at this ruling — "links are human-facing provenance, so keep them out of anything an LLM reads" — proved too broad: it stripped `AGENTS.md`, which is a *development* doc agents genuinely benefit from linking. The right line is **product vs development**. This repo holds two layers:

- **The product (`skills/`)** travels — `npx skills add` installs it into *other* projects, where it becomes operational LLM context. In a foreign project knack's AgenticWiki links are dead weight: no surrounding context, never fetched. So the product must be self-contained.
- **knack's own development context** (`AGENTS.md`, `glossary.md`, `decisions/`, `README.md`, `DESIGN.md`, `docs/`) *stays* in this repo. It guides anyone working *on* knack, where the links are load-bearing — knack's theory *is* the AgenticWiki. These are never installed; they link freely. *(Subsequently amended by ADR-0015: `DESIGN.md` was deleted and its content redistributed to `docs/architecture.md` and `docs/theory.md`. The product-vs-development split this ruling establishes is unchanged.)*

This also disambiguates two files that share a name: **knack's own `glossary.md`** (a development doc, links freely) vs **the `glossary.md` the `domain-modeling` skill manages in a user's project** (operational context there — no links; the skill must not instruct adding them).

## Decision

External links belong in knack's **development docs** — not in the **product** (`skills/`, or anything else that travels into a user's project). A skill carries the concept as a phrase/synopsis (the payload); the link (provenance + depth) lives in the development docs that stay in the knack repo.

This is a product rule, not a blanket "no URLs in LLM context" rule. An agent working *on* knack benefits from links in `AGENTS.md`/`glossary.md`/`decisions/`; an agent using knack's skills *in another project* never sees those docs, only the self-contained skills.

Amends ADR-0010's "pointer to depth — the wiki link" bullet: the link is not part of what a skill carries inline. The skill carries the synopsis; the link is maintained in knack's development docs.

## Consequences

- `skills/` (the product) is self-contained — all 7 stripped of inline wiki links; phrases/synopses retained (leading-words unaffected).
- knack's development docs (`AGENTS.md`, `glossary.md`, `README.md`, `decisions/`, `docs/`) keep their links — they stay in the repo and guide development. (`AGENTS.md` was briefly over-stripped under the too-broad reading; restored.)
- The `domain-modeling` skill must not teach users to add wiki links to *their* project glossary — that glossary is operational context in a foreign project. knack's own `glossary.md` is a different file (a development doc) and links freely.
- Risk: an agent using a skill in another project can't self-serve wiki depth. Mitigation: the synopsis is self-sufficient by design (ADR-0010); a skill that genuinely needs the source names it explicitly as the rare excepted case.
- Amends ADR-0010; does not supersede it.
