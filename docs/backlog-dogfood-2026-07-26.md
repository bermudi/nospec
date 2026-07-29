---
nospec: true
role: view
---

# Closed Dogfood Findings: 2026-07-26

Source: Two dogfood sessions where nospec was installed into an external project
(`pi-session-search`) and evaluated by both the worker (Pi/GPT) and a reviewer
(Opus). The worker produced two detailed verdicts; the reviewer produced a
code-review-style findings list and a litespec comparison.

Session: `~/.pi/agent/sessions/--home-daniel-build-nospec--/2026-07-26T01-16-03-302Z`

The 2026-07-28 consolidation commit (`d45062d`, “Consolidate skills around five
behavioral stances”) and preceding ADRs resolved findings #1–#7, #9, #11, and
#13. The remaining findings are now closed below. This file is a historical
view, not an active backlog.

---

## 8. `nospec-shape` reasoning

**Disposition: resolved.**

The cut guidance now explains, for tracer bullets, vertical slices, and
horizontal breadth:

- when the cut is useful;
- what evidence it produces;
- what remains unproven; and
- how the cut fails when misapplied.

This restores ADR-0010’s concepts-and-reasoning requirement without turning the
cuts into mandatory phases.

## 10. Runner/parser tests

**Disposition: resolved.**

The original finding understated the existing shell suite. `tests/run.sh`
already exercises whole-queue parser edge cases, every public runner verb,
baseline refusal and recording, blocker signals, verify failure and recovery,
handoff invalidation, review/fix limits, pin state, adoption boundaries, and
foreign-repository behavior.

Focused standard-library unit tests now call `queue_parser.py` directly for
field/fence parsing, combined malformed-input diagnostics, shell-error source
mapping, accessors, status mutation, and invalid mutations. The shell suite also
checks that an operating-system interruption exits nonzero while preserving the
in-progress queue state and writing a recoverable handoff. No pytest dependency
or second test command was introduced; `./tests/run.sh` remains the project’s
verification entry point.

## 12. Worker ownership of reversible choices

**Disposition: resolved by the current skill boundary.**

The worker adapter explicitly loads `nospec-carve`. Carve says reversible local
implementation choices are worker-owned and distinguishes them from novel,
consequential trade-offs, which block unattended execution unless authority was
delegated. Repeating the concept in the adapter prompt would create another
owner for the same instruction.

## 14. Family taxonomy

**Disposition: superseded by ADR-0026.**

The reviewed Understand/Change/Challenge/Remember/Leave taxonomy and
`metadata.family` field no longer exist. ADR-0026 owns the current five-skill,
stance-based topology; the source suite enforces the exact skill inventory and
runner role mapping. A `nospec family` verb would derive a deleted taxonomy, so
none is added.

The public catalog remains a deliberate view of ADR-0026 and the installed
skill records, not a second taxonomy.

## 15. Proposed versus accepted ADRs

**Disposition: resolved by ADR-0021.**

Nospec deliberately does not create durable proposed ADRs. Recommendations and
unresolved trade-offs remain conversational or in disposable design material.
Only a ruling accepted by the decision owner, made under delegated authority,
or covered by an established project convention becomes an accepted record.
`nospec-shape`, `nospec-carve`, and `nospec-curator` all transmit this boundary;
batch workers report a blocker rather than self-authorizing a ruling.

## 16. AFK safety and worktrees

**Disposition: resolved by ADR-0024.**

The runner now refuses a first mutating run with staged, unstaged, or untracked
paths outside `.loop/`, and refuses when Git cannot establish a baseline.
`--accept-dirty-baseline` explicitly records and accepts that risk; `--repo`
selects an operator-created worktree.

ADR-0024 explicitly rejects automatic worktree creation. It would make Nospec
own branch naming, queue transfer, result promotion, conflict handling, and
cleanup while still providing no security sandbox. Isolation beyond workspace
separation belongs to the harness, container, or operating system.

## 17. Combining nospec and litespec

**Disposition: rejected by the current architecture.**

ADR-0025 preserves project-designated contracts while rejecting a universal
behavioral canon and any inference of authority from a spec-like filename.
`docs/theory.md` correspondingly rejects importing litespec’s
canon/delta/archive model. Nospec keeps disposable work queues, external
verification, code/tests as the authority for implemented behavior, and
human-owned accepted rulings without becoming a combined litespec successor.
Reopening that direction would require superseding the accepted architecture,
not another backlog task.

---

## Final status

| # | Item | Disposition |
|---|---|---|
| 8 | Shape cut reasoning | Resolved |
| 10 | Runner/parser tests | Resolved |
| 12 | Reversible worker choices | Resolved by Carve ownership |
| 14 | Family taxonomy | Superseded by ADR-0026 |
| 15 | Proposed vs. accepted ADRs | Resolved by ADR-0021 |
| 16 | AFK worktree safety | Resolved by ADR-0024 |
| 17 | Nospec/litespec combination | Rejected by ADR-0025 |

**Open findings: none.**
