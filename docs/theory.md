# Theory and lineage

This doc is a view, not a record. It preserves the reasoning that led to knack's current shape and links to the [AgenticWiki](https://github.com/bermudi/AgenticWiki) concepts that ground it. The authoritative rulings live in `decisions/`; the current conceptual shape lives in [`architecture.md`](./architecture.md).

## What we're keeping from litespec

- The flow shape: `explore → plan → build → review → fix` (now composable, not rigid)
- Skills as procedural knowledge (`think/plan/build/review` → `explore/plan/build/review/fix` + shared `decide`/`domain-modeling`/`document`)
- Decisions (ADRs) — they persist because they're about rulings, not current behavior; when they stop applying, they're explicitly superseded, not left to silently rot
- Glossary — small, curated, doesn't rot the way specs do
- The patch lane concept (lightweight for small changes — now the default, not a special mode)

## What we're dropping from litespec

- Specs as source of truth → they rot ([doc-rot](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/doc-rot.md))
- Durable canon → accumulates wrong info ([doc-rot](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/doc-rot.md), [plan-disposability](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/plan-disposability.md))
- Archive as delta-merge to canon → ceremony, not evidence
- Unidirectional flow → real work loops back ([code-clarifies-spec](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/code-clarifies-spec.md), [spec-code-triangle](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/spec-code-triangle.md))
- Plan artifacts produce horizontal phases → [tracer-bullets](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/tracer-bullets.md) says don't (heuristic, not gate)
- Specs required for every change → too much ceremony for small work ([spec-driven-development](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/spec-driven-development.md) explicitly says SDD is "not for" simple prototypes and brownfield at scale)
- A semantic validator that pretends to judge plan quality → judgment belongs in skills (ADR-0010), not gate commands (ADR-0011)

## What we're keeping from knack

- Fresh context per tick ([ralph-loop](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/ralph-loop.md), [smart-zone-dumb-zone](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/smart-zone-dumb-zone.md))
- Verify gate outside the worker in batch mode ([backpressure](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/backpressure.md), [compounding-loops](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/compounding-loops.md))
- Bounded queue, hard stops ([agent-loop](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/agent-loop.md))
- Disposable artifacts ([plan-disposability](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/plan-disposability.md))
- Plain files ([compounding-loops](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/compounding-loops.md))

## What we're dropping from knack

- Loop-first framing → skills are the product; the loop is an optional companion (ADR-0009)
- Forces one shape of work (vertical slices) → work units are whatever shape the work is
- Only covers build phase → now covers the full flow via skills
- No decisions, glossary, or evolving context → now first-class
- The Go CLI → deleted; `npx skills` is the package manager (ADR-0011)
- Citation-based orphan-ADR checker → orphan semantics are relevance, transmitted as a concept by `decide` (ADR-0012)

## What we're stealing from mattpocock/skills

- Composable skills, not a monolithic flow (each skill independently invokable)
- Shared vocabulary skills (`domain-modeling`, `codebase-design` pattern → our `decide` + `domain-modeling` + `document`)
- ADRs captured inline during grilling, not as a separate phase
- Two-axis parallel review (Standards vs Intent, run as parallel sub-agents)
- No semantic validator (verification is distributed: execution, grilling, reproduction, human gates)
- Durable traces, disposable sessions (skills leave durable traces; agent sessions are ephemeral)
- PRDs must be split before execution (plan produces work units; loop only consumes `QUEUE.md`)

## What we're stealing from opengsd

- Per-work-unit model routing (each work unit can override the global agent command)
- `continue-here.md` for cross-session handoffs (loop writes `HANDOFF.md` on pause/stop)

## What we're avoiding from opengsd

- SQLite as source of truth (GSD Pi) — code is the source of truth, plain files for state
- Embedded loop in commands — keep the external loop script (swappable, debuggable, agent-agnostic)
- Complex runtime-specific installer — skills are already agent-agnostic via agentskills.io; `npx skills` adapts to each
- Complex capability system with overlays and trust models — keep it simple
- Citation-based decision coverage gates — references are evidence of relevance, not the test (ADR-0012)

## Grounding in the wiki

Every design decision cites a wiki concept. The wiki's position, synthesized:

- [doc-rot](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/doc-rot.md): "documentation can be worse than no documentation when it's stale." Specs are ephemeral destination hints, not living documents.
- [plan-disposability](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/plan-disposability.md): "treat plans as ephemeral coordination state, not contracts. A drifting plan is cheaper to regenerate than to salvage."
- [code-clarifies-spec](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/code-clarifies-spec.md): "no spec is perfect before implementation begins. The act of implementing generates new decisions that weren't in the spec."
- [spec-code-triangle](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/spec-code-triangle.md): spec, tests, and code form a bidirectional feedback loop. But [spec-driven-development](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/spec-driven-development.md) is explicitly "not for brownfield projects at scale."
- [decision-extraction](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/decision-extraction.md): the thing worth keeping from the spec process is the *decisions*, not the spec itself. Decisions persist; specs are consumed.
- [agent-loop](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/agent-loop.md): "cron plus a decision-maker in the body." For-each not while. Hard stops non-negotiable.
- [ralph-loop](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/ralph-loop.md): fresh context per tick, plan file as shared state, one task per iteration.
- [compounding-loops](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/compounding-loops.md): lateral coordination through shared durable files — artifacts, contracts, logs. Plain files as shared memory.
- [backpressure](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/backpressure.md): "engineer the environment so wrong outputs are mechanically rejected." Start with hard gates.
- [verification-loop](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/verification-loop.md) / ContextCov: executable enforcement (88.3%) beats passive instructions (67%) and LLM reflection (50.3%).
- [agent-skills](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/agent-skills.md) / [procedural-knowledge](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/procedural-knowledge.md): "a loop with no reusable skills inside it is just a while-true around a stranger." Skills are the reusable asset.
- [evolving-context](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/evolving-context.md): agents progressively refine their own context — prompts, skills, memories, preferences.
- [context-files](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/context-files.md): empirical evidence is mixed. LLM-generated overview dumps degrade performance. Minimal, developer-written, operational files work.
- [code-as-agent-harness](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/code-as-agent-harness.md): code is the operational substrate — executable, inspectable, stateful.
- [harness-engineering](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/harness-engineering.md): the central challenge is "semantic verification beyond executable feedback" — the green test is not the full specification.
- [aiming-problem](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/aiming-problem.md): the verification signal must capture the actual desired property, not a proxy the loop will learn to game.
- [steering-docs](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/steering-docs.md): `AGENTS.md` as accumulated learnings, not static configuration.
- [tracer-bullets](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/tracer-bullets.md): thin end-to-end slices for early integration feedback. A heuristic against horizontal phases, not a format requirement.
- [ubiquitous-language](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/ubiquitous-language.md): the shared vocabulary that lets human and agent mean the same thing by the same word.
