# 0009: Skills are the product; the loop is an optional batch companion

Date: 2026-07-11
Status: accepted

## Context

knack was built loop-first: `loop.sh` is the engine, the CLI scaffolds and versions skills, and the seven skills are embedded assets the loop drives per tick. A week of dogfooding showed the cost: nearly every work cycle (glossary, skill-versioning, review-loop, prompt-wiki, …) was the tool improving the tool. The reusable asset — the skills — stayed thin while the delivery mechanism for one mode accumulated governance apparatus (manifests, coverage gates, glossary checkers, a review subloop).

The reframe: work with agents happens across a spectrum of **human attention**, not a pipeline:

- **Interactive** — human present, edits land directly, real-time.
- **Plan-then-leave** — human present for the hard thinking, then the agent builds.
- **Batch (AFK)** — human absent; summaries, approvals, return to results.

The loop only serves the batch end. Skills serve all three. The original DESIGN.md thesis already said *"the reusable asset is procedural knowledge encoded as skills"* — but the project built the engine as the product and the skills as an embedded afterthought.

The things we want to replace (`/plan` commands, ralph loops, "gas towns", spec-kit/openspec) map almost entirely to **skills**, not to a runner. Only ralph loops / gas towns are the loop. Building the runner as the center serves a minority of the ambition.

Alternatives considered:

- **Keep loop-first; document the other modes as "composable paths."** Rejected — this is the current state. Composable paths in the docs still route every change through `plan → build (loop)`; the no-loop escape hatch was never made first-class, so in practice every change drifts toward ceremony. That is the litespec failure mode recurring.
- **Pure skills collection (a fork of obra/superpowers or mattpocock/skills).** Rejected — loses two things we want: (a) a lightweight way to capture durable traces (decisions, tangents) mid-flow, and (b) the batch runner for unattended work. The distinctive contribution is the combination, not skills alone.
- **Cut the loop entirely.** Rejected — the batch/AFK mode ("close 50 issues overnight") genuinely needs an unattended runner behind a verify gate. The loop isn't cruft; it's mis-centered.

## Decision

**Skills are the product.** The composable skills collection is the headline artifact — what works across all three attention-modes, what replaces `/plan` and spec-kit, what you point someone at.

**The loop is an optional companion for batch mode only.** It is demoted from "the engine" to "the unattended runner." It is reached for only when work is AFK and benefits from fresh-context-per-tick plus an external verify gate. Interactive and plan-then-leave modes do not require it; direct edits with a durable-trace capture are first-class, not the degenerate case.

**Durable traces stay thin and are written only when something crystallizes** — never per change, never as ceremony. `decisions/` when a ruling crystallizes. Glossary and LEARNINGS are optional and written-when-needed. The gate-style CLI checks (`decisions check`, `glossary check`) become occasional lints, not per-unit gates — the moment a durable-artifact check becomes a per-unit gate, it recreates litespec's archive ceremony.

The CLI's role narrows: skill scaffolding plus optional linting. Its package-manager-style versioning semantics are now questioned, not assumed.

## Consequences

- Re-frames ADR-0001/0002: the CLI still exists, but its scope shrinks from "read-only validator + skill packager" toward "scaffolder + lints." A future ADR may cut or slim it further. *(Resolved by [ADR-0011](0011-ship-as-skills-via-skills-sh-delete-cli.md): the CLI was cut entirely, not slimmed, and ADR-0001/0002 superseded.)*
- Contextualizes ADR-0008: the review-fix subloop is a batch-mode feature — valuable in AFK work, irrelevant to interactive mode.
- DESIGN.md and AGENTS.md must be re-centered: lead with skills and the attention-level axis, demote the loop, make the interactive/direct-edit path first-class. The "three artifacts" framing stays accurate but is no longer the headline.
- The meta-work treadmill loses a gear: changes to the tool no longer *must* route through the loop, so small improvements can land directly.
- Risk: "the loop is optional" can drift to "the loop is never used." Mitigation lives in the skills and docs naming *when* batch mode pays off — not in forcing it.
