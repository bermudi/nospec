---
nospec: true
id: 0026
date: 2026-07-28
status: accepted
spine: true
amends: [0008, 0019, 0020]
builds_on: [0010, 0013]
---

# 0026: Consolidate nine skills into five stance-based skills with progressive disclosure

## Context

ADR-0020 shipped nine skills — `nospec-scout`, `nospec-shape`, `nospec-carve`, `nospec-trial`, `nospec-mend`, `nospec-rule`, `nospec-lexicon`, `nospec-curator`, and `nospec`. Several pairs answered nearly the same user intent and forced the model to make a routing decision *before* it understood the problem:

- `nospec-scout` and `nospec-shape` both answer "I need to understand what should be built before building it." Scouting and decomposition are different levels of fidelity, not different stances. Requiring the model to choose between them before it understands the uncertainty creates the routing problem it then has to solve.
- `nospec-rule`, `nospec-lexicon`, and `nospec-curator` differ by *record format*, not behavioral stance. All begin with "a lasting claim appeared — does it deserve maintenance, who owns it, which projections need reconciliation?" The model should not have to decide "ADR skill or glossary skill?" before diagnosing the claim.
- `nospec-mend`'s builder behavior (diagnose the violated invariant, make the smallest coherent correction, verify) is the same stance as `nospec-carve`. A separate skill forced the model to distinguish "implement requested behavior" from "implement behavior requested by review," even though both require the same builder context. Trial's finding/classification behavior and Shape's queue-conversion behavior do not move.

Consolidation risks replacing routing cost with context cost: simply concatenating Scout into Shape or Rule and Lexicon into Curator would make every invocation load unrelated branches. Progressive disclosure applies ADR-0010 at the skill-structure level so merged stances remain compact while specialized formats and modes load only when needed.

## Decision

**Consolidate the nine skills into five, each organized around a behavioral stance, with progressive disclosure moving specialized modes, formats, and mechanics into `references/` loaded only when their branch applies.**

### 1. The five skills

| Skill | Stance | Merges |
|---|---|---|
| `nospec-shape` | Understand and bound the work before implementation | Scout + Shape |
| `nospec-carve` | Build or correct production code | Carve + the builder part of Mend |
| `nospec-trial` | Adversarially challenge an existing change | Trial (unchanged stance) |
| `nospec-curator` *(optional)* | Preserve newly crystallized durable knowledge | Rule + Lexicon + Curator |
| `nospec-loop` *(optional)* | Operate unattended execution behind external verification | the former `nospec` runner skill, renamed |

The runner skill is renamed `nospec` → `nospec-loop`. The runner binary stays `scripts/nospec` (the command on PATH is still `nospec`); only the skill directory and frontmatter `name` change. The `nospec` CLI verbs (`spine`, `adrs`, `check`, `lint`, `view`, `install`, `run`) remain the project's single bash entry point (ADR-0017, ADR-0018) and stay in the one script under `skills/nospec-loop/scripts/nospec`. No second bash entry point is created; `nospec-curator` references the existing record verbs rather than introducing a `nospec-records` script.

### 2. Why each merger

**Shape = Scout + Shape.** Scouting and decomposition are different levels of an evidence ladder, not different stances: inspect when existing evidence can resolve uncertainty; run a bounded experiment when it cannot; decompose once the outcome is understood; serialize a queue only when execution needs a handoff. The merged skill makes the fidelity choice internally. `references/scouting.md` loads only when an executable experiment is needed; `references/queue-format.md` loads only when serializing batch work.

**Carve = Carve + the builder part of Mend.** Resolving an accepted review finding is implementation, not a distinct stance. The builder still diagnoses the violated invariant, makes the smallest coherent correction, verifies, and avoids unrelated cleanup and fix oscillation. `references/review-corrections.md` loads only when applying accepted findings directly. Trial owns finding evidence and classification; Shape's `references/review-findings.md` owns conversion into queued outcomes; the batch fixer prompt enforces its mechanical "queue only, do not edit source" restriction; Carve owns actual source correction.

**Trial remains separate.** Review's adversarial stance conflicts with Carve's builder stance. Combining them primes the same context to defend and attack the implementation, and increases the chance a reviewer silently edits source or rationalizes the implementation it just produced. `references/batch-output.md` loads only for runner-backed review; interactive review gets the standards/intent axes, evidence rules, and classifications without the machine-readable artifact format.

**Curator = Rule + Lexicon + Curator.** These differed by record format, not stance. Curator routes internally by claim class: accepted consequential ruling → ADR; ambiguous recurring domain term → glossary; repository practice → operational context; designated promise → contract record; implemented behavior → code/tests; stale summary → repair or remove the view. `references/adr.md`, `references/glossary.md`, and `references/coherence.md` load only for their claim class, so an ADR task does not load glossary mechanics and vice versa. This creates one deep interface — "preserve this lasting knowledge" — rather than three shallow file-oriented interfaces.

**Loop remains separate and becomes `nospec-loop`.** Loop is an operator stance backed by a mechanical runner. Merging it into Shape would load queue execution, baseline, resume, and review-loop mechanics during ordinary interactive thinking, and would make batch execution appear mandatory again, contrary to ADR-0009. `references/operations.md` and `references/review-loop.md` load only for their mode; the runner scripts execute mechanics without being loaded as prose; the runtime prompts remain thin adapters.

### 3. Progressive-disclosure rule

Every skill follows three levels:

1. **Frontmatter `description`:** trigger and stance only.
2. **`SKILL.md`:** core reasoning shared by every invocation of that stance.
3. **`references/` (and `scripts/`, `prompts/`):** specialized modes, formats, and mechanics loaded only when their branch applies.

A reference needs an explicit condition ("If unresolved uncertainty requires an executable experiment, read `references/scouting.md`"). If every invocation needs a reference, its content belongs in `SKILL.md`. Splitting files merely to hit a line-count target would be fake progressive disclosure. References remain inside their owning skill, preserving ADR-0013 self-containment. There is no shared cross-skill preamble.

### 4. Runner composition after removing Mend

The batch roles become cleaner:

- worker prompt → `nospec-carve`;
- reviewer prompt → `nospec-trial`;
- fixer prompt → `nospec-shape`, appending one bounded unit per actionable finding;
- subsequent worker tick → `nospec-carve`.

The fixer is a restricted Shape invocation, not its own behavioral skill. This amends ADR-0008: the fix skill is now `nospec-shape` rather than `nospec-mend`, but the orchestration boundary (the loop owns stop conditions; skills own judgment) is unchanged.

## Alternatives considered

- **Keep nine skills; add cross-references between near-duplicate pairs.** Rejected — cross-references do not remove the routing decision the model must make before it understands the problem. The routing problem is the symptom; the cause is splitting one stance across two skills.
- **Merge by concatenating existing files under a new name.** Rejected — that produces progressive disclosure as a way to move templates, not a first-class design constraint. The skills must be re-authored around behavioral stances, not formed by concatenating existing files.
- **Merge Trial into Carve (one build+review skill).** Rejected — review's adversarial stance conflicts with the builder stance. Combining them primes the same context to defend and attack, and increases the chance a reviewer silently edits source or rationalizes the implementation it just produced.
- **Merge Loop into Shape.** Rejected — would load queue execution, baseline, resume, and review-loop mechanics during ordinary interactive thinking, and make batch execution appear mandatory again (contrary to ADR-0009).
- **Create a `nospec-records` script under Curator for the record verbs.** Rejected — would contradict ADR-0017 (single bash entry point). The record verbs stay in the one `nospec` CLI; Curator references them.
- **Keep the runner skill named `nospec`.** Rejected — the namesake skill carried the runner, but the stance-based model names skills by what they do. `nospec-loop` says "operate unattended execution"; `nospec` said nothing about the stance and collided with the project name and the CLI command. The CLI command stays `nospec`; only the skill directory and frontmatter `name` change.

## Consequences

- ADR-0008 is amended: the fix skill is `nospec-shape` (was `nospec-mend`). The orchestration boundary is unchanged.
- ADR-0019 is amended: the runner ships as the `nospec-loop` skill (was `nospec`), at `skills/nospec-loop/scripts/nospec` and `skills/nospec-loop/prompts/`. The "name only" skill-loading model and the `install` verb are unchanged. The skill count is no longer "nine."
- ADR-0020 is amended: the `nospec-` prefixed family is now five skills, not nine. The runner skill is `nospec-loop`, not `nospec`. The project name and CLI binary remain `nospec`. The `nospec-` prefix collision-safety rationale survives.
- The four removed skill directories (`nospec-scout`, `nospec-mend`, `nospec-rule`, `nospec-lexicon`) are deleted; their concepts survive inside `nospec-shape`, `nospec-carve`, and `nospec-curator`.
- `tests/run.sh` references the runner at `skills/nospec-loop/scripts/nospec`; the symlink-target assertion points there.
- `AGENTS.md`, `README.md`, `docs/`, and `glossary.md` are updated to the five-skill model and the `nospec-loop` path. Historical ADR bodies remain historical; inline notes on ADR-0008, ADR-0019, and ADR-0020 point cold readers to this ruling.
- Risk: an update may leave removed skill directories in an existing installation. Migration documentation names the retired skills explicitly so users can remove them before reinstalling the selected five-skill profile.
- Risk: a model that learned the old nine-skill routing may try to invoke a removed skill by name. Mitigation: the removed names are not installed; the surviving skills' `description` frontmatter covers the merged triggers, and the harness loads by trigger text, not by the old names.

## Related

- ADR-0008 — loop orchestrates review-fix (amended: fix skill is `nospec-shape`, not `nospec-mend`)
- ADR-0010 — skills transmit concepts, not rules (progressive disclosure is the concepts-not-rules mechanism applied to skill structure)
- ADR-0013 — wiki links live in docs, not in skill text (references stay inside their owning skill; self-containment preserved)
- ADR-0017 — derivable artifact metadata via bash CLI (the single bash entry point and its record verbs are unchanged; Curator references them)
- ADR-0019 — bundle runner as ninth skill (amended: the runner is the `nospec-loop` skill; skill count is no longer nine)
- ADR-0020 — rename to nospec (amended: five `nospec-` skills, not nine; runner skill is `nospec-loop`)
