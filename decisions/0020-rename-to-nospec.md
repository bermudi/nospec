---
id: 0020
date: 2026-07-20
status: accepted
spine: false
amends: [0003, 0009, 0010, 0011, 0017, 0018, 0019]
---

# 0020: Rename to nospec

## Context

The project was called "knack" from its inception (ADR-0003). The name was a craft metaphor — "a knack for the work" — and it read well as a CLI verb (`knack spine`, `knack run`). But two problems surfaced:

1. **The name didn't carry the thesis.** The project's whole reason for existing is the rejection of spec-rot (ADR-0009: specs are disposable; code is the source of truth). "knack" said nothing about that. It was a pleasant word with no stance.

2. **The skill names collided.** The original skill names — `explore`, `plan`, `build`, `review`, `fix`, `decide`, `document`, `domain-modeling` — are common English words. As the skills.sh ecosystem grew, these names collided with other skill packs: `scout` (scoutos), `shape` (pde-skills), `route` (skill-route), `steward` (skill_steward), `curator` (Skills_Curator), `recon` (recon-skills). A user who installed nospec alongside another pack would get silent skill-name collisions — the wrong skill loading, or both loading and confusing the agent.

## Decision

**Rename the project to `nospec` and all nine skills to a `nospec-` prefixed family.**

### 1. Project name: `nospec`

The name is a stance, not a metaphor. The project exists because specs rot; the name says what it rejects. "nospec" captures ADR-0009's thesis ("specs are disposable; code is the source of truth") in one word. It's a clean CLI verb (`nospec run`, `nospec spine`), a clean skill name (lowercase, no hyphens, 6 chars), and not a common command/package name (one dead 2010-era JS testing lib, no active collisions).

The nuance — "nospec" doesn't mean "never write specs," it means "no spec-rot, no spec-as-contract, no spec-permanence" — is captured in the docs. Specs are disposable coordination state. The name rejects spec *culture*, not the act of writing a queue.

### 2. Skill names: `nospec-` prefixed family

All nine skills are prefixed with `nospec-` to prevent collisions with other skill packs:

| Old name | New name | Role |
|---|---|---|
| `explore` | `nospec-scout` | recon before work |
| `plan` | `nospec-shape` | decompose intent into verifiable units |
| `build` | `nospec-hew` | implement one bounded outcome |
| `review` | `nospec-trial` | adversarial scrutiny |
| `fix` | `nospec-mend` | resolve review findings |
| `decide` | `nospec-rule` | capture ADR rulings |
| `domain-modeling` | `nospec-lexicon` | ubiquitous language |
| `document` | `nospec-curator` | route knowledge to owners |
| `knack` | `nospec` | the batch runner |

The runner skill is just `nospec` (no prefix) — it's the namesake and the mechanical companion. The eight procedural skills are `nospec-<word>`, collision-proof by construction.

The suffix words (scout, shape, hew, trial, mend, rule, lexicon, curator) were chosen for meaning-fit and distinctiveness. Several candidate names were rejected after collision checks against the skills.sh ecosystem: `scout` (scoutos), `recon` (recon-skills), `shape` (pde-skills), `route` (skill-route), `steward` (skill_steward), `curator` (Skills_Curator). The `nospec-` prefix makes these collisions irrelevant — `nospec-scout` won't collide with `scout`.

### 3. CLI binary: `nospec`

The runner is `scripts/nospec` inside the `nospec` skill. All CLI commands change: `nospec spine`, `nospec adrs`, `nospec check`, `nospec view`, `nospec install`, `nospec run`. The `install` verb (ADR-0019) symlinks `nospec` onto PATH.

### 4. Paths

- `skills/knack/scripts/knack` → `skills/nospec/scripts/nospec`
- `skills/knack/prompts/` → `skills/nospec/prompts/`
- `skills/knack/SKILL.md` → `skills/nospec/SKILL.md`
- All eight procedural skill directories renamed: `skills/explore` → `skills/nospec-scout`, etc.

## Alternatives considered

- **Keep the name "knack."** Rejected — the name didn't carry the thesis, and the common-word skill names were collision-bait in a growing ecosystem.
- **Rename only the project, keep generic skill names.** Rejected — the skill names are what actually collide. A user installing multiple packs gets skill-name conflicts, not project-name conflicts.
- **Use distinctive single-word skill names without a prefix (scout, shape, hew, etc.).** Rejected after collision checks showed the obvious-good words were already claimed in the skills.sh ecosystem. The prefix approach is collision-proof by construction.
- **Use a different project name (praxis, loom, helm, canon).** Rejected — "nospec" carries the thesis directly. Metaphor names (loom, helm) are pleasant but don't say what the project rejects. Stance names are honest.

## Consequences

- ADR-0003 ("the tool is named knack") is amended: the tool is now named `nospec`.
- ADR-0009, ADR-0010, ADR-0011, ADR-0017, ADR-0018, ADR-0019 are amended: all references to `knack` as the project/CLI/skill name become `nospec`; all references to old skill names (`explore`, `plan`, `build`, `review`, `fix`, `decide`, `domain-modeling`, `document`) become the new `nospec-` prefixed names.
- The ADR filenames in `decisions/` are NOT renamed — they are historical records. ADR-0003's file is still `0003-the-tool-is-named-knack.md`; its content is amended by this ADR.
- `tests/run.sh` references the runner at `skills/nospec/scripts/nospec`; the stale-CLI-reference check now catches `knack` in docs (the old name) alongside the old Go CLI subcommands.
- The ADR count is now 20 (this ADR added).
- Risk: users who installed the old `knack` skill pack need to reinstall. Mitigation: the project hasn't been published to skills.sh yet (no GitHub remote), so there are no existing users to migrate.
- Risk: the `nospec-` prefix is verbose. Mitigation: it's the conventional approach in skill ecosystems (cf. `shape:design`, `opsx:propose`); the collision-safety is worth the verbosity.

## Related

- ADR-0003 — the tool is named knack (amended: now `nospec`)
- ADR-0009 — skills are the product, loop is optional (the thesis the name captures)
- ADR-0010 — skills transmit concepts, not rules (skill names transmit the concept)
- ADR-0011 — ship via skills.sh (skill names must not collide in the ecosystem)
- ADR-0017 — derivable artifact metadata via bash CLI (CLI name is now `nospec`)
- ADR-0018 — one command with verbs (the command is now `nospec`)
- ADR-0019 — bundle runner as ninth skill (the ninth skill is now `nospec`, not `knack`)
