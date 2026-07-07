# Design: litespec replacement

**Status:** drafted 2026-07-05. Open questions remain (see bottom). Named **knack** (ADR-0003).

## One sentence

litespec's flow + knack's engine + code as source of truth + specs disposable + decisions and skills durable, agent-agnostic throughout.

## Thesis

Code is the source of truth. Specs rot ([[doc-rot]]). Plans are ephemeral coordination state ([[plan-disposability]]). The act of implementing generates decisions the spec didn't anticipate ([[code-clarifies-spec]]). The reusable asset is procedural knowledge encoded as skills ([[agent-skills]], [[procedural-knowledge]]). The loop owns verification; the worker never self-certifies ([[backpressure]], [[compounding-loops]]).

The tool is a harness, not a spec manager. It helps the agent produce and verify disposable work units, then they're gone. What persists: code, tests, decisions, glossary, skills, AGENTS.md.

## Three artifacts, three concerns

```
SKILLS (.agents/skills/)          LOOP (loop.sh)              CLI (Go binary)
the workflow                      the engine                  the validator
agent-agnostic                    agent-agnostic              agent-agnostic
agentskills.io spec               bash                        Go

explore  plan  build  review  fix     │                          │
  │       │      │                   reads QUEUE.md              │
  │       │      │                   invokes agent          validates structure
  │       │      │                   runs verify gate        provides context
  │       │      │                   appends evidence        manages ADR list
  │       ▼      │                          │                     │
  │   writes .loop/<name>/QUEUE.md ───────────────►│                     │
  │                                         │                     │
  │                                         │  may call between ticks
  │                                         └─────────────────────►│
  │                                                               │
  └─────────────────────────────────────────────────────────────────┘
                    skills + CLI are both read by the agent
                    the loop never reads skills directly
```

The three artifacts are independent. The loop never reads skills — it only reads QUEUE.md and invokes the agent. The agent discovers skills via agentskills.io progressive disclosure. The CLI never talks to agents. The loop may call the CLI between ticks for validation, but doesn't have to.

| | Skills | Loop | CLI |
|---|---|---|---|
| **What it is** | Markdown files (agentskills.io spec) | Bash script | Go binary |
| **Concern** | Procedural knowledge — the workflow | Execution — run the agent, gate on verify | Validation — check structure, provide context |
| **Who reads it** | The coding agent (any of them) | The shell | The agent, the loop, the human |
| **Agent-agnostic?** | Yes — all major agents discover `.agents/skills/` | Yes — via `LOOP_AGENT_CMD` env var | Yes — it doesn't talk to agents at all |
| **Lives where** | In the project (per-project workflow) | External (reusable across projects) | External (reusable across projects) |
| **Hackable by** | Human + agent ([[evolving-context]]) | Human (bash) | Human (Go) |

**Don't mix the CLI and the loop.** The loop is bash (hackable, simple). The CLI is Go (structured, reliable). The loop needs to know about agent invocation — that's a mess in a Go binary. The CLI needs to parse markdown and check schemas — that's a mess in bash. The loop can *call* the CLI between ticks if it wants validation, but they're independent. Each stays simple.

## Skills (the workflow)

Agent-agnostic. Every major coding agent (Codex, Claude Code, Devin, Pi, opencode, Cursor, Windsurf, Amazon Q, etc.) discovers `.agents/skills/` natively per the agentskills.io spec. Progressive disclosure: metadata at startup, full body on activation, references on demand.

```
.agents/skills/
├── explore/SKILL.md          # read code, grill intent, stress-test ideas
│                             # captures ADRs inline as decisions crystallize
│                             # delegates to domain-modeling for vocabulary
├── plan/SKILL.md             # decompose intent into verifiable work units
│                             # writes .loop/<name>/QUEUE.md
│                             # for big work: optionally writes .loop/<name>/specs/ (disposable)
│                             # rejects horizontal phases (tracer-bullets heuristic, not gate)
├── build/SKILL.md            # implement one work unit, don't self-certify
│                             # captures decisions as ADRs during implementation
│                             # writes operational learnings to AGENTS.md
├── review/SKILL.md           # two-axis parallel review:
│                             #   1. standards (coding conventions, codebase patterns)
│                             #   2. intent (does the change do what the work unit said)
│                             # review against actual codebase, not specs
├── fix/SKILL.md              # address review findings, generate new work units
├── decide/SKILL.md           # shared: capture architectural rulings as ADRs
│                             # used by explore, plan, build, review
└── domain-modeling/SKILL.md  # shared: domain vocabulary, glossary management
                              # used by explore, plan, review
```

### Design principles for skills

- **Composable, not monolithic.** The flow is a default path, not a gate. Skills can be invoked independently. `bug → explore → plan → build → done` is as valid as `big feature → explore → plan (with specs) → build → review → fix → done`.
- **Shared vocabulary skills.** `decide` and `domain-modeling` are shared resources other skills delegate to, rather than duplicating vocabulary. (Pattern from mattpocock/skills.)
- **Decisions captured inline.** ADRs are written during explore/plan/build as decisions crystallize, not as a separate phase. (Pattern from mattpocock's `grill-with-docs`.)
- **No semantic validator.** Skills don't judge plan quality. Verification is distributed: execution (tests), grilling (stress-test), reproduction (bugs), human gates (plan approval). The CLI does mechanical checks only.
- **Durable traces, disposable sessions.** Skills leave durable traces in the repo (ADRs, glossary, code, AGENTS.md) while agent sessions are disposable.
- **Evolving context.** The agent writes what it learns into AGENTS.md and proposes skill updates. ([[evolving-context]], [[steering-docs]].)

### Skill responsibilities by phase

**explore** — read the codebase and any existing loop state (`.loop/<name>/HANDOFF.md`, `.loop/<name>/QUEUE.md`), grill the intent, stress-test ideas. No artifacts produced except ADRs and glossary entries written inline when decisions crystallize. Pure conversation. Delegates to `domain-modeling` for vocabulary.

**plan** — decompose intent into deterministic, verifiable work units. Writes `.loop/<name>/QUEUE.md`; reads and discards stale queues if present. Each work unit has an outcome, constraints, done means, and a deterministic verify command. Rejects horizontal phases ("Phase 1: types / Phase 2: wiring") using the tracer-bullets heuristic — but this is a heuristic, not a gate. For big/greenfield work, *optionally* produces disposable specs (proposal, design) in `.loop/<name>/specs/` — consumed during build, then discarded, never canonized. Delegates to `domain-modeling` and `decide`.

**build** — implement one work unit, planning around its deterministic `Verify` command and staying within the runner's hard stops. Don't self-certify. The loop owns the gate. If you discover decisions, write ADRs to `decisions/` via the `decide` skill. If you learn operational knowledge, write it to `AGENTS.md` or propose a skill update.

**review** — two-axis parallel review (pattern from mattpocock's `code-review`), starting from the work unit and `.loop/<name>/EVIDENCE.md`:
1. **Standards** — does the change follow the repo's coding conventions and codebase patterns?
2. **Intent** — does the change do what the work unit said it would? (This is a judgment check, not a deterministic gate.)

Both axes run as parallel sub-agents so neither pollutes the other. Review against the actual codebase, not against a spec that may have rotted. Findings become new work units in `.loop/<name>/QUEUE.md` with deterministic `Verify` commands.

**fix** — address review findings. Read the existing `.loop/<name>/QUEUE.md`, append new work units generated from findings, and run another loop pass.

**decide** (shared) — when you make an architectural ruling, capture it as an ADR in `decisions/`. Decisions persist; specs don't. A decision made 6 months ago is still valid even if the code moved on — it explains *why* the code is the way it is. ([[decision-extraction]].)

**domain-modeling** (shared) — manage the project's ubiquitous language. Challenge terms against the glossary, stress-test with edge-case scenarios, update `glossary.md` inline. Used by explore, plan, review.

## The loop (external script, agent-agnostic)

A bash script, evolved from knack's `loop.sh`. Agent-agnostic via environment variable:

```bash
LOOP_AGENT_CMD="pi -p --no-session"        # Pi
LOOP_AGENT_CMD="claude --print"             # Claude Code
LOOP_AGENT_CMD="codex"                      # Codex
LOOP_AGENT_CMD="opencode run"               # opencode
```

### Per-tick behavior

1. Read the first `Status: pending` work unit from `.loop/<name>/QUEUE.md`.
2. Mark it `in_progress`.
3. Invoke the coding agent with the build skill + the work unit (fresh context per tick — [[ralph-loop]], [[smart-zone-dumb-zone]]).
4. Agent does the work and exits.
5. Loop runs the work unit's `Verify:` command **outside** the agent ([[backpressure]], [[compounding-loops]]).
6. On success: mark `done`, append to `.loop/<name>/EVIDENCE.md`.
7. On failure: mark `verify_failed` or `no_progress`, append evidence, stop or retry once.
8. Hard stops: max ticks, two no-progress strikes, verify-red after real change ([[agent-loop]]).
9. Halts when queue is empty or stop condition met.

### Additions from research

**Per-work-unit model routing** (from opengsd). Each work unit in QUEUE.md can optionally override the global agent command:

```markdown
## Slice 3: consolidate persistence wrappers
Agent: pi -p --no-session --model glm-5.2
...
```

Default: use `LOOP_AGENT_CMD`. Override: use the work unit's `Agent:` line. Enables using a powerful model for hard slices and a fast model for easy ones.

**Optional worktree isolation** (from opengsd):

```bash
loop.sh run --worktree .loop/<name>/QUEUE.md
# creates a git worktree, runs all work units there, squash-merges on success
```

Keeps main clean until verification passes. Optional, not default.

**Handoff file on pause/stop** (from opengsd). The loop writes `.loop/<name>/HANDOFF.md` on exit:

````markdown
# Handoff: <queue name>
Generated: <timestamp>

## Completed
- Slice 1: <outcome> (evidence: .loop/<name>/EVIDENCE.md#L<line>)
- Slice 2: <outcome>

## In progress
- Slice 3: <outcome> (status: verify_failed, last error: ...)

## Remaining
- Slice 4: <outcome>
- Slice 5: <outcome>

## Next action
Re-run loop after fixing the verify command for Slice 3.
````

This is the [[plan-disposability]] pattern — the handoff is a snapshot of coordination state, not a durable artifact. Delete it when the work is done.

### What the loop does NOT do

- It does not validate work unit structure (that's the CLI's job).
- It does not run review (that's a skill the agent loads).
- It does not manage ADRs or glossary (that's the `decide` and `domain-modeling` skills' job).
- It does not talk to the CLI unless explicitly configured to (e.g., `loop.sh run --validate-between-ticks`).

## The CLI (external binary, like litespec)

A Go binary. **Read-only context provider + structural validator.** It does NOT run agents. It does NOT run the loop. It validates and provides context.

```
<tool> validate              # check work units are well-formed (have verify, have goal)
<tool> status                # show queue state, evidence count, ADR count
<tool> decisions list|show   # list/show ADRs
<tool> decisions check       # decision coverage gate (see below)
<tool> glossary check        # check glossary terms are referenced consistently
                             #   (term defined in glossary.md but never used in code/docs → stale;
                             #    term used in code/docs but not in glossary → undefined)
<tool> instructions <artifact>  # print the template + guidance for creating a work unit / ADR / glossary entry
                             #   so the agent doesn't have to memorize format; pure text output
<tool> skills check          # validate skills follow agentskills.io spec
                             #   (frontmatter present, description field non-empty, no broken refs)
```

### What the CLI does NOT do (the things litespec did that we're dropping)

- No `new` / `patch` / `archive` — work units are just markdown files the agent writes directly.
- No delta specs / canon / archive-merge — nothing merges to a source of truth.
- No `apply` — that's the loop's job.
- No `review` — that's a skill the agent loads, not a CLI command.
- No semantic validator that pretends to judge plan quality.

### Decision coverage gate (from opengsd)

The one novel CLI feature. It is a **mechanical** check, not a semantic one:

- **Plan-time check:** every ADR in `decisions/` is referenced by at least one work unit in any `QUEUE.md` (current work) or any `EVIDENCE.md` ledger (completed work). The check is: does the ADR's filename or path appear in any work unit's body or evidence ledger? If not, the ADR is orphaned — either it's irrelevant to the current or past work, or the planner forgot to thread it through. EVIDENCE.md is the durable record that survives QUEUE.md deletion, so an ADR that drove a completed cycle is not orphaned as long as its ledger exists.
- **Build-time check:** every ADR referenced by a completed work unit still exists in `decisions/`. If a work unit claims to implement ADR-0007 but there's no `decisions/0007-*`, that's a dangling reference.

We do **not** attempt to verify that code changes are "consistent with" ADR content — that would require semantic understanding the CLI doesn't have. The gate catches reference errors and orphans, not logical contradictions.

## Project layout

```
project/
├── AGENTS.md                    # operational context (build/test commands, conventions)
│                                # the agent reads this every session
│                                # updated when operational knowledge is learned (evolving-context)
│
├── .agents/skills/              # procedural knowledge (the workflow)
│   ├── explore/SKILL.md
│   ├── plan/SKILL.md
│   ├── build/SKILL.md
│   ├── review/SKILL.md
│   ├── fix/SKILL.md
│   ├── decide/SKILL.md          # shared: ADR capture
│   └── domain-modeling/SKILL.md # shared: glossary/vocabulary
│
├── .loop/                       # disposable work state (plan-disposability)
│   └── <name>/                  # each work cycle in a named subdirectory
│       ├── QUEUE.md             # bounded queue of verifiable work units
│       ├── EVIDENCE.md          # append-only log of what each tick proved
│       ├── HANDOFF.md           # cross-session handoff (written on pause/stop)
│       └── specs/               # OPTIONAL — only for big/greenfield work
│           ├── proposal.md      # disposable planning artifacts
│           └── design.md        # consumed during build, then discarded
│                                # never merged to a canon, never archived
│
├── decisions/                   # DURABLE — ADRs (decision-extraction)
│   ├── 0001-use-bun-not-node.md # architectural rulings, not current behavior
│   └── 0002-persistence-pattern.md  # these don't rot because they're about decisions
│
├── glossary.md                  # DURABLE — ubiquitous language
│                                # small, curated, doesn't rot the way specs do
│
└── src/                         # THE SOURCE OF TRUTH
    └── ...                      # code + tests are the real specification
```

## Durable vs disposable

**Durable (persists, doesn't rot):**
- `src/` — the code. The source of truth.
- `decisions/` — ADRs. They're about rulings, not current behavior. A decision made 6 months ago is still valid even if the code moved on — it explains *why* the code is the way it is.
- `glossary.md` — ubiquitous language. Terms don't rot; they evolve deliberately.
- `AGENTS.md` — operational context. Build commands, conventions. Updated when things change.
- `.agents/skills/` — procedural knowledge. The workflow itself. Evolves ([[evolving-context]]).
- `.loop/<name>/EVIDENCE.md` — append-only ledger of what each tick proved. Durable so `decisions check` can trace which ADRs a completed cycle referenced after its QUEUE.md is deleted.

**Disposable (consumed then discarded):**
- `.loop/<name>/QUEUE.md` — the work queue. When the work is done, it's done. Delete it.
- `.loop/<name>/HANDOFF.md` — cross-session handoff. Snapshot of coordination state.
- `.loop/<name>/specs/` — planning artifacts for big work. Proposal, design. Consumed during build, then discarded ([[doc-rot]], [[plan-disposability]]).

The key inversion from litespec: **specs are disposable, code is durable.** litespec had it backwards — specs were the canon, code was downstream. The wiki says the opposite.

## The flow, end to end

```
1. EXPLORE (you + agent in conversation)
   agent loads explore skill (via agentskills.io progressive disclosure)
   read code, grill intent, stress-test ideas
   ADRs written inline as decisions crystallize (via decide skill)
   glossary updated inline (via domain-modeling skill)
   no work artifacts produced
        │
2. PLAN (agent, guided by plan skill)
   agent decomposes intent into verifiable work units
   writes .loop/<name>/QUEUE.md
   each unit: outcome + constraints + done means + verify command
   for big work: optionally writes .loop/<name>/specs/ (disposable)
   CLI validates work units are well-formed
        │
3. BUILD (the loop runs, agent-agnostic)
   loop.sh reads QUEUE.md
   for each work unit:
     loop invokes coding agent (any CLI) with a prompt that names the build skill
       (the agent discovers .agents/skills/build/SKILL.md via agentskills.io)
     fresh context per tick
     agent implements, exits
     loop runs verify command OUTSIDE the agent
     on success: mark done, append evidence
     on failure: mark failed, stop or retry
   during implementation the agent may invoke decide/domain-modeling skills
     to capture ADRs and glossary entries — these are durable
   agent may write operational learnings to AGENTS.md (evolving-context)
        │
4. REVIEW (human re-runs the agent, or runs it themselves, with review skill)
   agent loads review skill
   two-axis parallel review:
     1. standards (coding conventions, codebase patterns)
     2. intent (does the change do what the work unit said)
   review against actual codebase, not against specs
   findings become new work units appended to QUEUE.md
        │
5. FIX (human re-runs loop.sh on the updated QUEUE.md)
   loop runs again on the new work units from review
   same per-tick behavior as step 3
        │
6. DONE
   .loop/<name>/ discarded (queue, evidence, handoff, specs)
   what remains: code, tests, decisions, glossary, skills, AGENTS.md
   the code IS the specification now
```

**On skill triggering:** the loop doesn't load skills itself. It invokes the agent with a prompt that names the relevant skill (e.g. "use the build skill in `.agents/skills/build/`"). The agent then discovers and loads the skill via agentskills.io progressive disclosure — the same mechanism it would use for any skill in any project. The loop is skill-agnostic; it only knows the skill *name* to pass in the prompt. This is the answer to open question #4.

### Composable alternative paths

The flow is a default path, not a gate. Alternative paths are equally valid:

```
bug report → explore → plan → build (loop) → done
architecture review → explore → plan → build (loop) → review → fix (loop) → done
small fix → plan → build (loop) → done                    # skip explore
big feature → explore → plan (with specs) → build (loop) → review → fix (loop) → done
```

Decisions are captured inline throughout (not a separate phase). Glossary is updated inline during explore/plan. The loop only cares about QUEUE.md — it doesn't know which skill produced it.

## Work unit format (QUEUE.md)

````markdown
# Loop Queue: <short name>

Goal:
<one paragraph describing the desired end state>

Stop condition:
`<command that proves the whole packet is done, if one exists>`

## <outcome — what changes, observable>

Agent: <optional — overrides LOOP_AGENT_CMD for this unit only>

Why:
<only if non-obvious — else omit. Vertical/horizontal is context, not a requirement.>

Read first:
- <context the worker needs: ADR, code area, or file>
- <2–4 entries; context, not scope>

Constraints:
- <boundary or guardrail>
- <what must stay true or what is out of bounds>
- <if it names a file, it is "don't touch X" or "X's public API must not change", not "update X">

Done means:
- <observable condition>
- <no regression condition>

Verify:
```bash
<command that exits 0 on success>
```

Status: pending

## <next outcome>
...
````

### Design principles for work units

- **The work unit is whatever shape the work is.** Slice, patch, dedup, move, investigation, bug fix. Not forced into one named shape. "Vertical slice" is a heuristic against horizontal phases, not a required format.
- **Read first: is context, not scope.** Two to four entries: ADRs, code areas, or rulings. Prefer areas and rulings over file enumerations.
- **Constraints close the solution space.** A constraint states what must stay true or what is out of bounds — never what to edit. If it names a file, it is "don't touch X" or "X's public API must not change," not "update X."
- **Done means: is the acceptance criteria; Verify: is the enforceable subset.** The gap between them is the review surface.
- **The verify gate is the load-bearing field.** A work unit without a verify command is not loop-ready. That's the actual gate, not verticality.
- **The verify command must be deterministic and executable by the runner.** Tests, type checks, builds — not an LLM-as-judge. The loop provides mechanical backpressure.
- **"Why" is optional.** Filled in only when there's non-obvious context worth preserving. No padding.
- **No "Why this is vertical" field.** The planner skill treats vertical/horizontal as a warning about decomposition patterns, not a format requirement.

## What we're keeping from litespec

- The flow shape: explore → plan → build → review → fix (now composable, not rigid)
- Skills as procedural knowledge (think/plan/build/review → explore/plan/build/review/fix + shared decide/domain-modeling)
- Decisions (ADRs) — they don't rot because they're about rulings, not current behavior
- Glossary — small, curated, doesn't rot the way specs do
- Validate (structural checks — now mechanical only, no semantic validator)
- The patch lane concept (lightweight for small changes — now the default, not a special mode)

## What we're dropping from litespec

- Specs as source of truth → they rot ([[doc-rot]])
- Durable canon → accumulates wrong info ([[doc-rot]], [[plan-disposability]])
- Archive as delta-merge to canon → ceremony, not evidence
- Unidirectional flow → real work loops back ([[code-clarifies-spec]], [[spec-code-triangle]])
- Plan artifacts produce horizontal phases → [[tracer-bullets]] says don't (heuristic, not gate)
- Specs required for every change → too much ceremony for small work ([[spec-driven-development]] explicitly says SDD is "not for" simple prototypes and brownfield at scale)

## What we're keeping from knack

- Fresh context per tick ([[ralph-loop]], [[smart-zone-dumb-zone]])
- Verify gate outside the worker ([[backpressure]], [[compounding-loops]])
- Bounded queue, hard stops ([[agent-loop]])
- Disposable artifacts ([[plan-disposability]])
- Plain files ([[compounding-loops]])

## What we're dropping from knack

- Only covers build phase → now covers the full flow via skills
- Forces one shape of work (vertical slices) → work units are whatever shape the work is
- No upstream (explore/plan) or downstream (review) → now covered by skills
- No decisions, glossary, or evolving context → now first-class

## What we're stealing from mattpocock/skills

- Composable skills, not a monolithic flow (each skill independently invokable)
- Shared vocabulary skills (domain-modeling, codebase-design pattern → our decide + domain-modeling)
- ADRs captured inline during grilling, not as a separate phase
- Two-axis parallel review (Standards vs Intent, run as parallel sub-agents)
- No semantic validator (verification is distributed: execution, grilling, reproduction, human gates)
- Durable traces, disposable sessions (skills leave durable traces; agent sessions are ephemeral)
- PRDs must be split before execution (plan produces work units; loop only consumes QUEUE.md)

## What we're stealing from opengsd

- Decision coverage gates (every ADR referenced by at least one work unit; code changes consistent with ADRs)
- continue-here.md for cross-session handoffs (loop writes HANDOFF.md on pause/stop)
- Per-work-unit model routing (each work unit can override the global agent command)
- Worktree isolation (optional — git worktree per queue, squash-merge on success)

## What we're avoiding from opengsd

- SQLite as source of truth (GSD Pi) — code is the source of truth, plain files for state
- Embedded loop in commands — keep the external loop script (swappable, debuggable, agent-agnostic)
- Complex runtime-specific installer — skills are already agent-agnostic via agentskills.io
- Complex capability system with overlays and trust models — keep it simple

## Grounding in the wiki

Every design decision cites a wiki concept. The wiki's position, synthesized:

- [[doc-rot]]: "documentation can be worse than no documentation when it's stale." Specs are ephemeral destination hints, not living documents.
- [[plan-disposability]]: "treat plans as ephemeral coordination state, not contracts. A drifting plan is cheaper to regenerate than to salvage."
- [[code-clarifies-spec]]: "no spec is perfect before implementation begins. The act of implementing generates new decisions that weren't in the spec."
- [[spec-code-triangle]]: spec, tests, and code form a bidirectional feedback loop. But [[spec-driven-development]] is explicitly "not for brownfield projects at scale."
- [[decision-extraction]]: the thing worth keeping from the spec process is the *decisions*, not the spec itself. Decisions persist; specs are consumed.
- [[agent-loop]]: "cron plus a decision-maker in the body." For-each not while. Hard stops non-negotiable.
- [[ralph-loop]]: fresh context per tick, plan file as shared state, one task per iteration.
- [[compounding-loops]]: lateral coordination through shared durable files — artifacts, contracts, logs. Plain files as shared memory.
- [[backpressure]]: "engineer the environment so wrong outputs are mechanically rejected." Start with hard gates.
- [[verification-loop]] / ContextCov: executable enforcement (88.3%) beats passive instructions (67%) and LLM reflection (50.3%).
- [[agent-skills]] / [[procedural-knowledge]]: "a loop with no reusable skills inside it is just a while-true around a stranger." Skills are the reusable asset.
- [[evolving-context]]: agents progressively refine their own context — prompts, skills, memories, preferences.
- [[context-files]]: empirical evidence is mixed. LLM-generated overview dumps degrade performance. Minimal, developer-written, operational files work.
- [[code-as-agent-harness]]: code is the operational substrate — executable, inspectable, stateful.
- [[harness-engineering]]: the central challenge is "semantic verification beyond executable feedback" — the green test is not the full specification.
- [[aiming-problem]]: "the verification signal must capture the actual desired property, not a proxy the loop will learn to game."
- [[steering-docs]]: AGENTS.md as accumulated learnings, not static configuration.

## Open questions

1. **~~Name.~~** Resolved (ADR-0003): the tool is named **knack**.
2. **Language for the CLI.** Go (like litespec) is the default assumption. Could be something else, but Go is proven here.
3. **Skills shipped with the tool vs project-authored skills.** Does the tool ship a default set of skills (explore/plan/build/review/fix/decide/domain-modeling) that projects can override? Or does each project author its own? litespec generated skills from Go templates; mattpocock ships skills as plain markdown. Lean toward shipping defaults as plain markdown that projects can override.
4. **~~How the loop invokes the agent with the right skill.~~** Resolved: the loop passes the skill *name* in the prompt; the agent discovers the skill via agentskills.io progressive disclosure. See "On skill triggering" above. Still open: does this work reliably across all target agents (Pi, Claude Code, Codex, opencode, Devin), or do some need an explicit `--skill` flag? Needs testing.
5. **~~Decision coverage check implementation.~~** Resolved: mechanical only — reference existence and orphan detection. See "Decision coverage gate" above. Still open: is this useful enough to justify the implementation cost, or should it be deferred to v2?

## Migration from current state

The repo currently contains:
- `loop.sh` — the first-iteration loop, uses "slice" terminology
- `.agents/skills/vertical-slice-planner/SKILL.md` — a planner skill, also slice-oriented
- `tests/run.sh` — validates the skill

Migration path (not yet executed):
1. **Rename "slice" → "work unit"** in `loop.sh` and the planner skill. The QUEUE.md parser in `loop.sh` reads `## Slice N:` headers — change to `## <outcome>` headers (no numbered prefix, no "Slice" word).
2. **Split the planner skill.** The current `vertical-slice-planner` becomes `plan`. The `explore`, `build`, `review`, `fix`, `decide`, and `domain-modeling` skills are new and need to be authored.
3. **Rename `SLICELOOP_AGENT_CMD` → `LOOP_AGENT_CMD`** in `loop.sh` (already supported, just needs rename for agent-agnosticism).
4. **Add per-work-unit `Agent:` override** parsing to `loop.sh`.
5. **Add handoff file generation** to `loop.sh` (`.loop/<name>/HANDOFF.md` on pause/stop).
6. **Build the CLI** from scratch in Go — no existing CLI code to migrate.
7. **Update `tests/run.sh`** to validate the new skill set, not just the planner.
