# 0014: Durability is maintenance, not permanence

Date: 2026-07-18
Status: accepted

## Context

The durable-vs-disposable distinction is the spine of the project's anti-spec-rot argument. But the durable side was described in language stronger than the project's own rulings support:

- DESIGN.md said durable artifacts "persist, don't rot," glossary "Terms don't rot," and ADRs "are still valid even if the code moved on."
- AGENTS.md mirrored with "Durable (persist, don't rot)."
- Yet ADR-0012 *supersedes* ADR-0006 — an ADR literally stopped being valid. The orphan-ADR model is built on the premise that ADRs can stop explaining or constraining the system.
- AGENTS.md itself lists "stale glossary terms" as a hygiene concern. DESIGN.md denied what AGENTS.md admitted.

Two overclaims were operating:

1. **"Code is the source of truth" read as a complete epistemology.** The doc also establishes that intent lives in `Done means`, rationale in ADRs, domain meaning in the glossary, and operational knowledge in AGENTS.md — none reconstructible from code. Green verify explicitly does not prove semantic correctness (the coherence-vs-compilation gap, AGENTS.md "Lessons learned").
2. **Durability presented as immunity to rot.** "Doesn't rot" / "still valid" implied permanence or automatic correctness, contradicting ADR-0012's relevance-based orphan model and the project's own admission that glossary terms stale.

The qualified comparative ("doesn't rot *the way specs do*") was already in the doc and is correct — durability is a matter of degree and maintenance, not a binary against rot. The unqualified forms were the outliers.

Alternatives considered:

- **Leave the slogan, fix only the durable list.** Rejected — the same overclaim appeared in DESIGN.md:105 (decide synopsis), DESIGN.md:417 ("they don't rot because..."), and AGENTS.md:39. A partial fix would re-introduce the coherence gap AGENTS.md warns about (green verify ≠ durable docs cohere with ADRs).
- **Drop the "code is the source of truth" slogan entirely.** Rejected — it carries the anti-litespec inversion (specs disposable, code durable) that the thesis depends on. The slogan is a flag, not a definition; the problem was reading it as one.

## Decision

Durable artifacts are **maintained records whose value survives the work cycle** — not artifacts immune to rot. The contrast with disposable coordination state (consumed then discarded) is about *disposability*, not about *permanence*.

Authority is partitioned, not collapsed into code:

- **Code and executable tests** are authoritative for *current implemented behavior*.
- **ADRs** are authoritative for *rulings*. A decision made 6 months ago still explains *why* a choice was made, even if the code moved on — and can be superseded when it stops explaining or constraining the system (per ADR-0012).
- **Glossary** is authoritative for *domain terms*. Terms evolve deliberately; they can stale and need pruning.
- **Work artifacts** (`QUEUE.md`, `HANDOFF.md`, `REVIEW.md`, `specs/`) describe intent only while being consumed.

The slogan "code is the source of truth" is retained in the thesis lines (DESIGN.md:15, 27, 250) as the rhetorical anchor for the anti-litespec inversion. It is *not* used as a definition in the durable-artifact list, where the partitioned-authority formulation replaces it.

## Consequences

- DESIGN.md:105, 235–239, 417 and AGENTS.md:39 now use the maintained-records framing and the partitioned-authority formulation. The qualified comparative ("doesn't rot the way specs do") is retained where it already appeared.
- The slogan survives as a tagline only; readers of the durable list get the precise model.
- "Still valid" → "still explains why a choice was made" aligns the durable docs with ADR-0012's supersession model.
- Glossary-staleness is now consistent across DESIGN.md and AGENTS.md (both admit it).
- No skill text needed changing — `skills/decide/SKILL.md` already framed decisions as "explaining why the code is the way it is," not as "still valid." The overclaim was a docs-only drift.
- Risk: "maintained records whose value survives the work cycle" is longer than "doesn't rot." Mitigation: the slogan still carries the short form where rhetoric matters; the durable list carries the precise form where definition matters.
