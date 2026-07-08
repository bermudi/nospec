# 0006: Grandfather pre-ledger ADRs in decisions check

Date: 2026-07-06
Status: accepted
Grandfathered: meta-ruling about the check mechanism itself; implemented and verified directly (go test), not via a loop cycle.

## Context

`decisions check` flags an ADR as orphaned if no QUEUE.md or EVIDENCE.md references it. This is the right rule for ADRs going forward — every new ruling should drive a work unit that lands in the evidence ledger.

But ADRs 0001, 0002, and 0003 were written and enacted before the evidence-ledger convention existed (ADR-0004 introduced named cycles; the durable-ledger semantics landed later). The work they drove — building the CLI, packaging skills, naming the tool — is done and verified in the code itself. There is no EVIDENCE.md that traces back to them, and backfilling one would be fiction: the ledger records what a tick proved, and no tick ran against these ADRs.

Alternatives considered:
- **Backfill a retroactive EVIDENCE.md.** Fabricates evidence for work that was never run through the loop. The ledger would no longer mean "a tick proved this."
- **Date-based heuristic (ADR predates earliest EVIDENCE.md).** Fragile: ADRs 0001–0003 and the first evidence are all dated 2026-07-06. Same-day granularity makes the comparison ambiguous, and ADR dates are day-level, not timestamp-level.
- **Accept the noise.** Three permanent false positives drown out real orphans. The check loses signal.

## Decision

An ADR may carry a `Grandfathered: <reason>` line. `decisions check` treats any ADR with that line as non-orphaned — it is exempt from the reference requirement. The reason field records *why* the ADR has no ledger trace, so the exemption is auditable rather than magical.

This is a one-time acknowledgment for ADRs that predate the ledger convention. New ADRs should not carry it — they should drive work units that produce evidence.

## Consequences

- ADRs 0001, 0002, 0003 get a `Grandfathered:` line and stop appearing as orphaned.
- `decisions check` remains useful: it still catches ADRs written *after* the ledger convention that no work references.
- The exemption is explicit and per-ADR — no global flag, no date heuristic to maintain.
- Future ADRs that genuinely have no work trail must justify themselves with the field, making the exemption a visible choice rather than silent drift.
