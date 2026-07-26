---
nospec: true
role: view
---

# Architecture

A composable **skills collection** for turning intent into verified code while keeping temporary coordination disposable and durable knowledge explicit. It encodes the [AgenticWiki](https://github.com/bermudi/AgenticWiki)'s theory as plain [agentskills.io](https://agentskills.io) skills, with an optional runner for unattended batch work.

Authoritative rulings live in `decisions/`; this doc is a view that ties them together. When they disagree, the ADR wins.

## Spine

The spine (ADR-0009 onward) is the curated set of load-bearing rulings. It's derived from `decisions/` frontmatter — run `nospec spine` to list it. See `decisions/` for the full set.

## One sentence

Composable capabilities help agents understand, change, challenge, and remember software work across levels of human attention; an optional runner supplies external enforcement when the human leaves.

## Human attention spectrum

Work with agents happens across levels of attention, not a pipeline:

- **Interactive** — human present, edits land directly.
- **Plan-then-leave** — human does the hard thinking, then the agent builds.
- **Batch (AFK)** — human absent; the loop runs units behind a verify gate.

Skills serve all three. The loop serves only batch. Skills are the product; the loop is optional.

## Skills and batch companion

`nospec-carve` is the common implementation skill. Scout, Shape, Trial, and Mend address uncertainty, decomposition, review, and correction only when those problems occur. Rule, Lexicon, and Curator preserve the rulings, language, and context worth keeping after work state is discarded. Each remains independently invokable; their names do not define a workflow.

The ninth skill, `nospec`, transmits when AFK execution is appropriate and what its verify gate can establish, then carries the runner that enforces that concept in batch. The runner never reads skills itself: worker prompts name the relevant skill, and the worker's harness loads it. It knows only its queue, environment, repository state, process results, and evidence contract (ADR-0007, ADR-0019).

## Artifact roles

Durable knowledge is organized by role. Each fact has one owner; other documents are deliberate projections of it.

| Role | Examples | Purpose |
|---|---|---|
| **Record** | `skills/`, `decisions/`, `glossary.md`, `AGENTS.md`, code/tests, project-designated contract records | Owns a class of claim |
| **View** | `README.md`, `docs/architecture.md`, `docs/getting-started.md`, `docs/skills.md`, `docs/loop.md` | Helps readers understand records together |
| **Ledger** | `.loop/<name>/EVIDENCE.md` | Append-only record of what happened |
| **Work state** | `.loop/<name>/QUEUE.md`, `HANDOFF.md`, `REVIEW.md`, scratch work specs | Coordination state consumed then discarded |

Views summarize and link; they do not independently redefine what they project. When a record changes, its projections are reconciled. When a view contradicts its record, the record wins.

## Durable versus disposable

Durable artifacts survive the work cycle because they are maintained records:

- code and tests — current implemented behavior
- project-designated contract records — public promises or required behavior assigned through operational context, accepted rulings, or established metadata conventions
- `skills/` — procedural knowledge
- `decisions/` — architectural rulings
- `glossary.md` — domain terms
- `AGENTS.md` — operational context
- `.loop/<name>/EVIDENCE.md` — accepted cycle baseline, resulting worktree state, exact verify command/result, conservative proof boundary, and adopted-artifact pin state (ADR-0016 and its amendments; baseline: ADR-0024)

Disposable artifacts are consumed then discarded:

- `.loop/<name>/QUEUE.md` — the work queue
- `.loop/<name>/HANDOFF.md` — cross-session handoff
- `.loop/<name>/REVIEW.md` — review artifact under `--review`
- `.loop/<name>/specs/` — planning artifacts for big work

The inversion from litespec is precise: **work specs are disposable**. Nospec does not maintain a universal behavioral canon, but it respects durable contract records explicitly owned by the host project. A filename such as `specs/` grants neither disposability nor authority by itself (ADR-0025).

## The flow, by attention mode

The same skills serve all three modes; only who runs verification and whether a queue is involved changes.

### Interactive

Invoke only the skill whose problem is present: scout uncertainty, shape outcomes, carve a bounded change, trial an existing diff, or mend known findings. No queue is required. The agent runs the relevant verification itself before declaring done.

### Plan-then-leave

The human settles the judgment-heavy parts, then leaves the agent enough bounded context to continue. Use scouting or shaping only if the work needs them. A queue is optional cross-session coordination state; without the runner, verification discipline remains the agent's.

### Batch (AFK)

```
nospec-shape (writes QUEUE.md) → nospec run runs ticks → optional --review subloop → done
```

The loop owns the verify gate: it invokes the agent, runs `Verify:` outside the agent, and appends evidence.

## Coherence is not compilation

Green tests and valid links do not prove that durable docs agree with rulings, terms, or the current code. Coherence is a separate, judgment-based check: do the records and their projections still tell the same story? The `nospec-curator` skill exists to surface and route these problems.

Some structural drift is mechanically detectable: `nospec check` validates metadata and ownership for artifacts that explicitly opt in with `nospec: true`. Ordinary repository Markdown is ignored; installing Nospec does not adopt a document schema. The command does not catch semantic contradiction — that remains judgment.

## Grounding

The cited theory lives in the [AgenticWiki](https://github.com/bermudi/AgenticWiki). See [`theory.md`](./theory.md) for the lineage of what nospec keeps, drops, and borrows, plus the full concept map. Key concepts include doc-rot, plan-disposability, code-clarifies-spec, backpressure, compounding-loops, ralph-loop, agent-loop, decision-extraction, ubiquitous-language, evolving-context, and tracer-bullets.
