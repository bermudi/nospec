# 0016: Proof-boundary is mechanical; pin-state is provenance, not coherence

Date: 2026-07-18
Status: accepted

## Context

knack's `EVIDENCE.md` ledger records what verify proved per cycle. Two gaps in the format surfaced during a comparison with the abandoned litespec-v2 prototype, which stated the thesis "a claim becomes usable when it has provenance, current-state validation, and an explicit verification scope":

1. **The proof-boundary was vacuous.** `loop.sh`'s `append_evidence` wrote "What remains unverified: Anything outside the verify command's proof scope" — a tautology true for any verify command, carrying no information the command's own name didn't imply. litespec-v2's rich "What This Does Not Prove" section was judgment-filled and reward-hackable (the worker filling it in has the same blind spots that caused the gap); the generic version swung too far the other way — unfalsifiable, zero information.

2. **No provenance pin.** `EVIDENCE.md` recorded what verify ran and what changed, but not which durable docs the cycle touched. When a later cycle edited a durable doc (e.g., `glossary.md`) without updating the docs that describe it (e.g., `AGENTS.md`), the indirect coherence failure was invisible until a manual `document` pass caught it. litespec-v2 had a "Promotion Applied" section naming which durable docs were touched; knack had no equivalent.

The question was whether to pull litespec-v2's evidence sections across wholesale. Three rounds of analysis (against the AgenticWiki's `reward-hacking`, `aiming-problem`, `satisfaction-of-search`, `harness-engineering`, `MemRefine`, `babysitter-agent`, and `proactive-service` concepts) produced a narrower ruling.

## Decision

Three changes to `EVIDENCE.md`'s format and the loop's coherence-triggering story, each scoped to its ADR-0010 layer:

### 1. Proof-boundary: registry-derived positive (mechanical contract)

Replace the vacuous negative with a **registry-derived positive**. `loop.sh` decomposes the verify command on `&&`, classifies each segment against a primitive registry (`go test` → "Go test suite passes"; `! grep` → "no matches for pattern in scope"; `diff -r` → "directories match"; etc.), and emits a bullet per segment. Unknown segments fall back to "command exited 0: `<segment>`" — never an interpreted claim. The claims stay **syntactic** (what the command mechanically checks), never **semantic** (what the feature does). The raw verify command remains adjacent in the evidence for cross-check.

The negative stays but is honest: "Anything outside the above proof scope; see the verify command for the exact check." This is the "explicit verification scope" leg of litespec-v2's invariant, realized mechanically instead of as judgment — ADR-0010-clean.

This is a mechanical contract (`loop.sh` derives it deterministically), not a prompt. Misclassification is bounded: it is systematic (a registry bug is found once and fixed for all future evidence), not per-instance (the agent's blind spot is re-rolled fresh every bundle). Bounded by `harness-engineering`'s "the green test is not the full specification" — emit syntactic claims, never semantic ones.

### 2. Pin-state: provenance + pin, not coherence closure

litespec-v2 called this "Promotion Applied"; we call it **pin-state** because the implementation records git blob SHAs (pins) for durable docs touched in the cycle, not a boolean "was this promoted." The vocabulary change is deliberate: "promotion" suggests a one-time act, while "pin" captures the ongoing comparison (has this doc moved since the last cycle pinned it?) that makes the alert mechanism work.

Add a `Pinned durable docs:` section to each `EVIDENCE.md` entry. For each file in the cycle's `changed_files` that matches the durable-doc set (`decisions/*.md`, `glossary.md`, `AGENTS.md`, `README.md`, `LEARNINGS.md`, `docs/**/*.md`, `skills/**/SKILL.md`), record its git blob SHA at evidence-write time. This is **provenance** — the state of the durable doc when the claim was made — not coherence validation.

A `Pin alerts:` section follows when a prior cycle pinned the same path and its SHA has since changed. This is a **triage trigger**, not a coherence gate. It converts silent drift into a flagged signal for the next review pass; it does not close indirect coherence failure (A changed in a way that contradicts unpinned B — the `AGENTS.md` ↔ `glossary.md` description-drift case). That remains judgment territory, handled by the `document` skill.

The framing discipline is load-bearing: the section is labeled `Pinned durable docs:` and `Pin alerts:`, never `Coherence validated:`. A green pin is executable feedback about provenance; misread as feedback about coherence, it is `reward-hacking` Case 3 (structural gate lending false confidence to the gaps that bite you).

### 3. Pin-check location: loop.sh records, review judges

The pin-moved comparison is **mechanical** → `loop.sh` computes it and writes the alert into `EVIDENCE.md`. The response to a pin alert is **judgment** → the `review` skill reads `Pin alerts:` lines and routes them to `document`, which scopes the coherence check from the diff of the pinned file. `loop.sh` never invokes `document` (that would violate ADR-0008: the loop owns orchestration, not judgment). The pin alert is invisible to the build worker (`prompts/worker.md` stays build-only) — the worker never sees a coherence warning and never fixes coherence inline, avoiding the review/fix responsibility blur (commit c222676) at a new layer.

This mirrors the existing verify-gate pattern: `loop.sh` runs verify (mechanical) and records pass/fail; `review` (judgment) interprets. The pin-check is the same pattern at the coherence layer — `loop.sh` computes + records, `review` judges. Bounded by `babysitter-agent` ("invisible to the master agent") and `proactive-service` ("executable constraint over typed state… surfaces an alert in the manifest").

## What this does not close

Neither mechanism closes indirect or semantic coherence. A change to `glossary.md` that invalidates `AGENTS.md`'s description of it — without an ADR being superseded or a pinned ref moving in the promotion sense — passes both the verify gate and the pin-check. Coherence is relational (the wiki's `MemRefine`: "single-entry scoring cannot tell when an edit removes a fact not redundantly covered elsewhere; pairwise judgment is the minimal unit"). It remains judgment — the `document` skill, triggered by the existing named conditions (ADR superseded, public interface changed) **plus** a pinned durable doc moving.

## Consequences

- `loop.sh`'s `append_evidence` gains a `derive_proof_claims` helper (python, embedded) and a `record_pin_state` helper (python, embedded). The primitive registry is code with tests.
- `EVIDENCE.md` entries gain `What this proves:` (registry-derived bullets), `What remains unverified:` (honest negative), `Pinned durable docs:`, and `Pin alerts:` sections. The `What this proves:` header is retained; its content changes from a single generic line to multiple specific bullets.
- The `review` skill adds: read `Pin alerts:` from `EVIDENCE.md`; route each alert to `document` for coherence assessment.
- The `document` skill adds: when invoked from a pin alert, scope the coherence check from the diff of the pinned file — what changed, and which other durable docs describe or depend on it.
- The `build` skill is unchanged — the proof-boundary concept is already transmitted ("the verify gate is the mechanical contract; the gap between `Verify:` and `Done means:` is the review surface"). The format change is mechanical, in `loop.sh`.
- The durable-doc set is knack-specific (this is knack's own development tool, not a shipped skill). Projects installing knack's skills do not inherit the pin-check; they get the concept via the `document` skill.
- `tests/run.sh` gains assertions for the new `EVIDENCE.md` sections.
- Risk: the primitive registry may misclassify compound segments (e.g., `test $(grep ... | wc -l) -eq 5`). Mitigation: fallback to "command exited 0: `<segment>`" for undecomposable segments; the raw command is always adjacent for cross-check; registry bugs are systematic and fixed once.
- Risk: the durable-doc set may miss project-specific durable docs. Mitigation: the set is extensible; this ADR records the default, not a closed list.

## Related

- ADR-0008 — loop orchestrates review-fix (the loop owns orchestration, not judgment; the pin-check respects this boundary)
- ADR-0010 — skills transmit concepts, not rules (the proof-boundary is a mechanical contract in `loop.sh`; the coherence response is a concept in `review`/`document`)
- ADR-0014 — durability is maintenance, not permanence (pins record state, not permanence)
- ADR-0015 — artifact roles and ownership (the durable-doc set is derived from the artifact-role model)
