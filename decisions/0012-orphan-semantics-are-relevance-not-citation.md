---
id: 0012
date: 2026-07-12
status: accepted
spine: true
supersedes: [0006]
---

# 0012: Orphan-ADR semantics are relevance, not citation

## Context

ADR-0006 established the orphan rule as the deleted `decisions check` CLI enforced it: an ADR is orphaned if no `QUEUE.md` or `EVIDENCE.md` references it, with a `Grandfathered:` exemption for ADRs that predate the evidence ledger. ADR-0011 deleted `decisions check` and moved orphan-hygiene into the `decide` skill as a transmitted concept ("the agent self-checks from understanding rather than from a gate"). But the *semantics* survived — `decide` and `AGENTS.md` carried the citation-based definition mode-independently, generalizing the deleted checker's proxy rather than replacing it.

Citation is evidence of relevance, not relevance itself. The citation model misfires on:

- **Negative rulings** — "we will not build X" constrains future work but drives no work unit, so it has no citing work; under the citation model it reads as orphaned.
- **Conventions** — a naming or layout ruling constrains future work without an immediate change to cite it.
- **ADRs obvious in code** — a ruling whose implementation is visible in the codebase but has no explicit citation reads as orphaned, though it still explains the code.
- **Interactive decisions** — there is no `QUEUE.md` or `EVIDENCE.md` at all; the citation model has nothing to grep.

This is the same proxy-vs-judgment distinction ADR-0010 draws (judgment belongs in skills as transmitted concepts, not in mechanical gates) and that ADR-0011 applied to delete the checker. Carrying the checker's citation proxy into a mode-independent skill re-imports the gate the project already rejected.

Alternatives considered:

- **Keep citation-based, generalized mode-independently.** Rejected — still the deleted checker's proxy; fails every case above; contradicts ADR-0010's judgment-over-mechanical-proxy principle.
- **Relevance-based with no reference guidance.** Rejected — references are useful evidence of relevance; the skill should teach using them as signals, just not as the definition.

## Decision

An ADR is **orphaned when it no longer explains or constrains the system** — when nothing about the current code or ongoing work depends on its ruling. References (`QUEUE.md`, `EVIDENCE.md`, code, docs, the change in flight) are **evidence of relevance, not the definition**. A negative ruling can be alive with no citing work at all.

This is judgment, transmitted by the `decide` skill, not a gate. No mechanical checker can decide "does this ADR still explain or constrain the system" — that is the point: the proxy was wrong, and the judgment is the skill's job.

Supersedes ADR-0006: the citation-based `decisions check` rule and its `Grandfathered:` exemption described the deleted checker (ADR-0011) and are replaced by this relevance concept. The `Grandfathered:` fields on ADRs 0001–0004 remain as historical records of why those ADRs predate the evidence ledger; they no longer reference an active rule.

## Consequences

- `decide` and `AGENTS.md` carry the relevance-based definition; references are signals, not the test.
- ADR-0006 is superseded (mutual link). The `Grandfathered:` fields on 0001–0004 are historical, not active.
- Orphan-hygiene is fully judgment — consistent with ADR-0011 having deleted the gate. There is no path back to a mechanical orphan checker without re-importing the proxy this ADR rejects.
- Risk: relevance is harder to assess than grepping for citations. Mitigation: that is the intended trade — the proxy produced false positives (negative rulings, conventions) and false negatives (uncited-but-live ADRs); judgment is the skill's job, per ADR-0010.
