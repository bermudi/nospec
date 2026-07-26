---
nospec: true
id: 0025
date: 2026-07-26
status: accepted
spine: true
amends: [0009, 0014, 0015, 0017, 0020]
builds_on: [0021]
---

# 0025: Work specs are disposable; contract records are project-owned

## Context

The inversion from litespec was summarized as "specs are disposable; code is the source of truth." That rejected a universal prose canon maintained beside implementation, but the slogan was broad enough to imply that every durable behavioral contract is suspect.

Dogfooding against litespec exposed both sides of the distinction. Its stale canon demonstrated that structural validation cannot keep a second prose model semantically coherent with code. At the same time, public promises such as an API schema, protocol definition, compatibility policy, or executable contract test may be legitimate maintained records with real consumers. Calling those artifacts disposable merely because someone calls them a "spec" would discard useful project authority.

The artifact's role and owner matter more than its filename or label.

## Decision

Nospec distinguishes **work specs** from **contract records**:

- A **work spec** coordinates a current change: a queue, handoff, review artifact, design note, or scratch plan. It describes current intent while being consumed and is disposable afterward.
- A **contract record** is an artifact the host project explicitly designates and maintains as the owner of a public promise or required behavior. Examples include an API schema, protocol definition, compatibility policy, or executable contract test.

Nospec does not create a universal behavioral canon, infer contract authority from a path such as `specs/`, or grant an artifact authority merely by being installed. It also does not forbid a host project's existing contract records.

Designation must be recoverable from an authoritative host-project surface: operational context such as `AGENTS.md`, an accepted ruling, or an established metadata convention that assigns ownership. The `nospec: true` marker adopts Nospec's structural checks only; it does not designate a contract record. If no designation is recoverable, the artifact remains unclassified and the agent asks rather than inferring authority from its name or location.

Authority remains partitioned:

- code and executable tests own current implemented behavior;
- project-owned contract records own the promises or requirements explicitly assigned to them;
- ADRs own architectural rulings;
- glossary and operational context own their existing claim classes; and
- work specs own current implementation intent only while the work is active.

When a contract record, code, and tests disagree, the mismatch is a coherence problem to resolve explicitly; structural validation must not pretend to choose the winner.

## Consequences

- Precise documentation says **work specs** or **coordination state** when describing disposability. The short `nospec` name remains a rejection of spec-rot and universal spec authority, not of every project-owned contract.
- `.loop/<name>/specs/` remains disposable. A host repository's unrelated `specs/` path is not classified by Nospec.
- Shaping and implementation may cite contract records as context or constraints. Verification should exercise them when mechanically possible.
- Projects may use their existing ownership surface; Nospec adds no contract manifest or mandatory metadata schema.
- `nospec-trial` and `nospec-curator` treat designated contracts as records, not as disposable plans or generic views.
- Nospec composes with OpenAPI, schemas, protocol definitions, contract tests, and similar native mechanisms without importing litespec's canonical delta/archive model.
