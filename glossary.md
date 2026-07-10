# Glossary

## agent-loop

A loop architecture described as "cron plus a decision-maker in the body" — for-each, not while, with non-negotiable hard stops.

## agent-skills

Reusable procedural knowledge packaged as skills; "a loop with no reusable skills inside it is just a while-true around a stranger."

## aiming-problem

The risk that a verification signal measures a proxy the loop learns to game rather than the actual desired property.

## backpressure

Engineering the environment so wrong outputs are mechanically rejected, starting with hard gates.

## code-as-agent-harness

Code as the operational substrate — executable, inspectable, stateful.

## code-clarifies-spec

The principle that implementing a spec surfaces decisions the spec did not anticipate.

## compounding-loops

Lateral coordination through shared durable files (artifacts, contracts, logs) used as shared memory.

## context-files

Empirically mixed evidence that LLM-generated overview dumps degrade performance while minimal developer-written operational files work.

## decision-extraction

The idea that the decisions a spec process produces are worth keeping, not the spec itself.

## doc-rot

Documentation worse than no documentation when it is stale; specs are ephemeral destination hints, not living documents.

## evolving-context

Agents progressively refine their own context — prompts, skills, memories, preferences.

## harness-engineering

The central challenge of semantic verification beyond executable feedback — the green test is not the full specification.

## plan-disposability

Treating plans as ephemeral coordination state, not contracts; a drifting plan is cheaper to regenerate than to salvage.

## procedural-knowledge

Knowledge encoded as reusable skills rather than one-off instructions.

## ralph-loop

A loop pattern with fresh context per tick, a plan file as shared state, and one task per iteration.

## smart-zone-dumb-zone

A loop design separating a thin "dumb" execution zone from a "smart" decision zone.

## spec-code-triangle

The bidirectional feedback loop between spec, tests, and code.

## spec-driven-development

A methodology explicitly "not for" simple prototypes and brownfield projects at scale.

## steering-docs

AGENTS.md as accumulated learnings rather than static configuration.

## tracer-bullets

A heuristic against horizontal implementation phases; build end-to-end thin slices instead.

## verification-loop

Executable enforcement (e.g., ContextCov at 88.3%) beating passive instructions (67%) and LLM reflection (50.3%).
