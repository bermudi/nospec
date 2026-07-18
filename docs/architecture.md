# Architecture

A composable **skills collection** for agentic development — the procedural encoding of the [AgenticWiki](https://github.com/bermudi/AgenticWiki)'s theory — shipped as plain [agentskills.io](https://agentskills.io) skills, with an optional bash loop for unattended batch work.

Authoritative rulings live in `decisions/`; this doc is a view that ties them together. When they disagree, the ADR wins.

## Spine

- **ADR-0009** — skills are the product; the loop is an optional batch companion.
- **ADR-0010** — skills transmit concepts and reasoning, not rules.
- **ADR-0011** — ship via [skills.sh](https://skills.sh); the Go CLI is deleted.
- **ADR-0012** — orphan-ADR semantics are relevance, not citation.
- **ADR-0013** — wiki links live in docs, not in skill text.
- **ADR-0014** — durability is maintenance, not permanence.
- **ADR-0015** — durable knowledge is organized by artifact role with clear ownership.

## One sentence

knack turns intent into a disposable queue of verifiable work units, runs each unit behind a deterministic verify gate, and encodes the resulting know-how as reusable skills — with decisions, glossary, and operational context kept as durable records.

## Human attention spectrum

Work with agents happens across levels of attention, not a pipeline:

- **Interactive** — human present, edits land directly.
- **Plan-then-leave** — human does the hard thinking, then the agent builds.
- **Batch (AFK)** — human absent; the loop runs units behind a verify gate.

Skills serve all three. The loop serves only batch. Skills are the product; the loop is optional.

## Two layers

```
skills/                       loop.sh
product                     optional companion
procedural knowledge        mechanical execution
agent-agnostic              agent-agnostic

explore  plan  build  review  fix
  │       │     │      │       │
  │       │     │      │       └── reads/writes .loop/<name>/QUEUE.md
  │       │     │      │
  └── decide, domain-modeling, document — shared across everyone
```

The loop never reads skills. The worker prompt names the skill explicitly — name and path — so the agent reads the file directly (ADR-0007). The loop only knows `QUEUE.md`, environment variables, and exit status.

## Artifact roles

Durable knowledge is organized by role. Each fact has one owner; other documents are deliberate projections of it.

| Role | Examples | Purpose |
|---|---|---|
| **Record** | `skills/`, `decisions/`, `glossary.md`, `AGENTS.md`, `LEARNINGS.md`, code/tests | Owns a class of claim |
| **View** | `README.md`, `docs/architecture.md`, `docs/getting-started.md`, `docs/skills.md`, `docs/loop.md` | Helps readers understand records together |
| **Ledger** | `.loop/<name>/EVIDENCE.md`, `LEARNINGS.md` | Append-only record of what happened |
| **Work state** | `.loop/<name>/QUEUE.md`, `HANDOFF.md`, `REVIEW.md`, `specs/` | Coordination state consumed then discarded |

Views summarize and link; they do not independently redefine what they project. When a record changes, its projections are reconciled. When a view contradicts its record, the record wins.

## Durable versus disposable

Durable artifacts survive the work cycle because they are maintained records:

- `src/` / code and tests — implemented behavior
- `skills/` — procedural knowledge
- `decisions/` — architectural rulings
- `glossary.md` — domain terms
- `LEARNINGS.md` — domain/problem insights
- `AGENTS.md` — operational context
- `.loop/<name>/EVIDENCE.md` — what each tick proved

Disposable artifacts are consumed then discarded:

- `.loop/<name>/QUEUE.md` — the work queue
- `.loop/<name>/HANDOFF.md` — cross-session handoff
- `.loop/<name>/REVIEW.md` — review artifact under `--review`
- `.loop/<name>/specs/` — planning artifacts for big work

The inversion from litespec: specs are disposable, code is durable.

## The flow, by attention mode

The same skills serve all three modes; only who runs verification and whether a queue is involved changes.

### Interactive

```
explore → plan → build → review → fix → done
```

Skip steps as needed. `bug → plan → build → done` is valid. The agent runs `Verify:` itself before declaring done.

### Plan-then-leave

```
explore + plan (interactive) → queue written → agent builds while human absent → done
```

The queue is a shared to-do list. Verification discipline is still the agent's.

### Batch (AFK)

```
plan (writes QUEUE.md) → loop.sh runs ticks → optional --review subloop → done
```

The loop owns the verify gate: it invokes the agent, runs `Verify:` outside the agent, and appends evidence.

## Coherence is not compilation

Green tests and valid links do not prove that durable docs agree with rulings, terms, or the current code. Coherence is a separate, judgment-based check: do the records and their projections still tell the same story? The `document` skill exists to surface and route these problems.

## Grounding

The cited theory lives in the [AgenticWiki](https://github.com/bermudi/AgenticWiki). See [`theory.md`](./theory.md) for the lineage of what knack keeps, drops, and borrows, plus the full concept map. Key concepts include doc-rot, plan-disposability, code-clarifies-spec, backpressure, compounding-loops, ralph-loop, agent-loop, decision-extraction, ubiquitous-language, evolving-context, and tracer-bullets.
