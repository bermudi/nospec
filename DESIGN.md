# Design

Named **knack** (ADR-0003; rename pending). Authoritative rulings live in `decisions/`; this doc is the narrative that ties them together. When they disagree, the ADR wins.

The spine:

- **ADR-0009** — skills are the product; the loop is an optional batch companion.
- **ADR-0010** — skills transmit concepts and reasoning, not rules.
- **ADR-0011** — ship via [skills.sh](https://skills.sh); the Go CLI is deleted.
- **ADR-0012** — orphan-ADR semantics are relevance, not citation.
- **ADR-0013** — wiki links live in docs, not in skill text.

## One sentence

A composable **skills collection** — the procedural encoding of the [AgenticWiki](https://github.com/bermudi/AgenticWiki)'s theory — shipped as plain [agentskills.io](https://agentskills.io) skills via [`npx skills`](https://github.com/vercel-labs/skills) / [skills.sh](https://skills.sh), with an optional bash loop for unattended batch work. Code is the source of truth; specs are disposable; decisions and skills are durable.

## Thesis

Work with agents happens across a spectrum of **human attention**, not a pipeline (ADR-0009):

- **Interactive** — human present, edits land directly, real-time.
- **Plan-then-leave** — human present for the hard thinking, then the agent builds.
- **Batch (AFK)** — human absent; the loop runs units behind a verify gate.

Skills serve all three. The loop serves only batch. Skills are the product; the loop is an optional companion.

The reusable asset is procedural knowledge encoded as skills ([agent-skills](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/agent-skills.md), [procedural-knowledge](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/procedural-knowledge.md)). Code is the source of truth; specs rot ([doc-rot](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/doc-rot.md)); plans are ephemeral coordination state ([plan-disposability](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/plan-disposability.md)); the act of implementing generates decisions the spec didn't anticipate ([code-clarifies-spec](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/code-clarifies-spec.md)).

Skills transmit **concepts and the reasoning behind them**, not rules (ADR-0010). Judgment — decomposition, process choice, depth, when to explore — is concept-guided; the agent decides. Mechanical contracts (the verify gate runs outside the agent and must exit 0; hard stops) stay hard rules. The wiki is the cited source of the *why*; the project's docs link to it rather than redefine its concepts — skills carry the synopsis, docs carry the link (ADR-0013).

## Two layers

```
skills/                       loop.sh                        (no CLI)
the product                   the optional companion         deleted (ADR-0011)
agent-agnostic                agent-agnostic
agentskills.io spec           bash

explore  plan  build  review  fix     │
  │       │     │      │       │      reads QUEUE.md
  │       │     │      │       │      invokes the worker (LOOP_AGENT_CMD)
  │       │     │      │       │      runs the verify gate, outside the agent
  │       │     │      │       │      appends evidence
  │       │     │      │       │      optional --review: orchestrates review/fix
  │       ▼     │      │       │
  │   writes .loop/<name>/QUEUE.md
  │
  └── decide, domain-modeling — shared, used by everyone
```

The two layers are independent. The loop never reads skills — it only reads `QUEUE.md` and invokes the agent. The agent loads skills by name and path (ADR-0007); the loop is skill-agnostic. There is no CLI; distribution, versioning, and discovery are [skills.sh](https://skills.sh)'s job (ADR-0011).

| | Skills | Loop |
|---|---|---|
| **What it is** | Markdown files (agentskills.io spec) | Bash script |
| **Concern** | Procedural knowledge — the workflow | Execution — run the agent, gate on verify |
| **Who reads it** | The coding agent (any of them) | The shell |
| **Agent-agnostic?** | Yes — installed into each agent's native skills path by `npx skills` | Yes — via `LOOP_AGENT_CMD` env var |
| **Lives where** | Source: `skills/` in this repo. Installed: agent-native (`.agents/skills/`, `.claude/skills/`, …) | External (reusable across projects) |
| **Hackable by** | Human + agent ([evolving-context](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/evolving-context.md)) | Human (bash) |

## Skills (the product)

Seven skills, all reworked to ADR-0010 (mode-independent, concept-forward; no external links — ADR-0013):

```
skills/
├── explore/SKILL.md          # read code, grill intent, stress-test ideas
├── plan/SKILL.md             # decompose intent into verifiable work units
├── build/SKILL.md            # implement one work unit, don't self-certify
├── review/SKILL.md           # two-axis parallel review: standards + intent
├── fix/SKILL.md              # address review findings, generate new work units
├── decide/SKILL.md           # shared: capture architectural rulings as ADRs
└── domain-modeling/SKILL.md  # shared: domain vocabulary, glossary management
```

### Design principles for skills

- **Concepts, not rules (ADR-0010).** Each concept a skill teaches carries *what it is*, *why it exists* (the failure mode it prevents), and a reasoned default with its override-reasoning. Defaults are scaffolding, never mandates. A default without override-reasoning is a rule in disguise.
- **Mechanical contracts stay hard.** The verify gate is external and deterministic; hard stops are non-negotiable. These are mechanisms, not prompts.
- **Composable, not monolithic.** The flow is a default path, not a gate. Skills can be invoked independently. `bug → explore → plan → build → done` is as valid as `big feature → explore → plan → build → review → fix → done`.
- **Shared vocabulary skills.** `decide` and `domain-modeling` are shared resources other skills delegate to, rather than duplicating vocabulary. (Pattern from mattpocock/skills.)
- **Decisions captured inline.** ADRs are written during explore/plan/build as decisions crystallize, not as a separate phase. (Pattern from mattpocock's `grill-with-docs`.)
- **No semantic validator.** Skills don't judge plan quality. Verification is distributed: execution (tests), grilling (stress-test), reproduction (bugs), human gates (plan approval). Mechanical checks ride with the loop; judgment lints do not survive as commands (ADR-0011).
- **Durable traces, disposable sessions.** Skills leave durable traces in the repo (ADRs, glossary, code, AGENTS.md) while agent sessions are disposable.
- **Evolving context.** The agent writes what it learns into AGENTS.md and proposes skill updates. ([evolving-context](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/evolving-context.md), [steering-docs](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/steering-docs.md).)
- **Self-contained.** The product travels into other projects via `npx skills add`. External links would be dead weight in a foreign project's context, so the skills carry phrases/synopses only; the wiki links live here, in knack's development docs (ADR-0013).

### Skill responsibilities

**explore** — read the codebase and any existing loop state (`.loop/<name>/HANDOFF.md`, `.loop/<name>/QUEUE.md`), grill the intent, stress-test ideas. No artifacts produced except ADRs and glossary entries written inline when decisions crystallize. Pure conversation. Delegates to `domain-modeling` for vocabulary.

**plan** — decompose intent into deterministic, verifiable work units. Writes `.loop/<name>/QUEUE.md`; reads and discards stale queues if present. Each work unit has an outcome, constraints, done means, and a deterministic verify command. Rejects horizontal phases ("Phase 1: types / Phase 2: wiring") using the [tracer-bullets](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/tracer-bullets.md) heuristic — but this is a heuristic, not a gate. For big/greenfield work, *optionally* produces disposable specs (proposal, design) in `.loop/<name>/specs/` — consumed during build, then discarded, never canonized. Delegates to `domain-modeling` and `decide`.

**build** — implement one work unit, planning around its deterministic `Verify` command and staying within the runner's hard stops. Don't self-certify. In a batch cycle the loop owns the gate; interactively, *you* run the verify before claiming done — the principle (don't claim success without verification) survives across modes, the enforcement mechanism doesn't. If you discover decisions, write ADRs to `decisions/` via the `decide` skill. If you learn operational knowledge, write it to `AGENTS.md` or propose a skill update.

**review** — two-axis parallel review (pattern from mattpocock's `code-review`), starting from the work unit and `.loop/<name>/EVIDENCE.md`:
1. **Standards** — does the change follow the repo's coding conventions and codebase patterns?
2. **Intent** — does the change do what the work unit said it would? (Judgment, not a deterministic gate.)

Both axes run as parallel sub-agents so neither pollutes the other. Review against the actual codebase, not against a spec that may have rotted. Findings become new work units in `.loop/<name>/QUEUE.md` with deterministic `Verify` commands.

**fix** — address review findings. Read the existing `.loop/<name>/QUEUE.md`, append new work units generated from findings, and run another loop pass.

**decide** (shared) — when you make an architectural ruling, capture it as an ADR in `decisions/`. Decisions persist; specs don't. A decision made 6 months ago still explains *why* a choice was made, even if the code moved on. ([decision-extraction](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/decision-extraction.md).) Also carries the orphan-ADR hygiene concept (ADR-0012): an ADR is orphaned when it no longer explains or constrains the system — references are evidence of relevance, not the test.

**domain-modeling** (shared) — manage the project's ubiquitous language. Challenge terms against the glossary, stress-test with edge-case scenarios, update `glossary.md` inline. Used by explore, plan, review. In a user's project this glossary is operational context — the skill must not teach adding wiki links to it (ADR-0013).

## The loop (optional companion, batch only)

A bash script, `loop.sh`. Agent-agnostic via environment variable:

```bash
LOOP_AGENT_CMD="pi -p --no-session"        # Pi
LOOP_AGENT_CMD="claude --print"             # Claude Code
LOOP_AGENT_CMD="codex"                      # Codex
LOOP_AGENT_CMD="opencode run"               # opencode
LOOP_AGENT_CMD="devin --print --prompt-file \"\$LOOP_PROMPT_FILE\" --model kimi-k2.7 --permission-mode dangerous"  # Devin
```

The loop is reached for only when work is AFK and benefits from fresh-context-per-tick plus an external verify gate. Interactive and plan-then-leave modes do not require it; direct edits with durable-trace capture are first-class, not the degenerate case (ADR-0009).

### Per-tick behavior

1. Read the first `Status: pending` work unit from `.loop/<name>/QUEUE.md`.
2. Mark it `in_progress`.
3. Snapshot the repo state (diff + untracked files outside `.loop`).
4. Invoke the coding agent with the worker prompt + the work unit (fresh context per tick — [ralph-loop](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/ralph-loop.md), [smart-zone-dumb-zone](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/smart-zone-dumb-zone.md)).
5. Agent does the work and exits.
6. Loop runs the work unit's `Verify:` command **outside** the agent ([backpressure](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/backpressure.md), [compounding-loops](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/compounding-loops.md)).
7. On success: mark `done`, append to `.loop/<name>/EVIDENCE.md`.
8. On failure: `verify_failed` (repo changed) or `no_progress` (unchanged, retry once); two no-progress strikes stop the loop.
9. Hard stops: max ticks, two no-progress strikes, verify-red after real change ([agent-loop](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/agent-loop.md)).
10. Halts when the queue is empty or a stop condition is met.

### Optional review/fix subloop (ADR-0008)

With `--review`, after the build queue drains the loop orchestrates a bounded `build → review → fix` subloop:

1. Invoke a review worker that writes `.loop/<name>/REVIEW.md`.
2. Read only the `- actionable: N` summary line from `REVIEW.md`.
3. If `N > 0`: invoke a fix worker that appends `Status: pending` units to `QUEUE.md`, then re-run the build pass and review again.
4. Stop when `actionable` is 0, `--max-review-rounds` is hit, `--max-ticks` is exhausted, or fix produces no new units.

The loop owns orchestration and stop conditions only; the `review` and `fix` skills own judgment. Without `--review`, review and fix are manual skill invocations.

### Per-unit model routing

Each work unit in `QUEUE.md` can optionally override the global agent command:

```markdown
## Slice 3: consolidate persistence wrappers
Agent: pi -p --no-session --model glm-5.2
...
```

Default: use `LOOP_AGENT_CMD`. Override: use the work unit's `Agent:` line. Enables a powerful model for hard slices and a fast model for easy ones.

### Handoff file on pause/stop

The loop writes `.loop/<name>/HANDOFF.md` on non-clean exit:

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

This is the [plan-disposability](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/plan-disposability.md) pattern — the handoff is a snapshot of coordination state, not a durable artifact. Delete it when the work resumes.

### What the loop does NOT do

- It does not validate work unit structure — it parses `QUEUE.md` in bash and trusts the format. (There is no CLI to validate; ADR-0011.)
- It does not judge review findings or decide whether a finding is actionable. When `--review` is set, the loop *orchestrates* the review/fix subloop — invoking the workers and reading only the actionable count from `REVIEW.md` — but the judgment stays in the skills (ADR-0008).
- It does not manage ADRs or glossary (that's the `decide` and `domain-modeling` skills' job).
- It does not read skills. The worker prompt (`prompts/worker.md`) names the skill explicitly — name *and path* — and the agent reads the file directly (ADR-0007). Trigger-based agentskills.io discovery was not reliable enough across agents, and the loop cannot confirm a skill loaded.

## Project layout

```
project/
├── AGENTS.md                    # operational context (build/test commands, conventions)
│                                # the agent reads this every session
│                                # updated when operational knowledge is learned (evolving-context)
│
├── skills/                      # THE PRODUCT — procedural knowledge (the workflow)
│   ├── explore/SKILL.md         # source layout; `npx skills add` installs into each
│   ├── plan/SKILL.md            # agent's native skills path (.agents/skills/, .claude/skills/, …)
│   ├── build/SKILL.md
│   ├── review/SKILL.md
│   ├── fix/SKILL.md
│   ├── decide/SKILL.md          # shared: ADR capture
│   └── domain-modeling/SKILL.md # shared: glossary/vocabulary
│
├── .loop/                       # disposable work state (plan-disposability)
│   └── <name>/                  # each work cycle in a named subdirectory
│       ├── QUEUE.md             # bounded queue of verifiable work units (disposable)
│       ├── EVIDENCE.md          # append-only log of what each tick proved (DURABLE)
│       ├── HANDOFF.md           # cross-session handoff (written on pause/stop; disposable)
│       ├── REVIEW.md            # written by the review worker under --review (disposable)
│       └── specs/               # OPTIONAL — only for big/greenfield work
│           ├── proposal.md      # disposable planning artifacts
│           └── design.md        # consumed during build, then discarded
│                                # never merged to a canon, never archived
│
├── decisions/                   # DURABLE — ADRs (decision-extraction)
│   ├── 0001-…                   # architectural rulings, not current behavior
│   └── …
│
├── glossary.md                  # DURABLE — ubiquitous language
│                                # small, curated, doesn't rot the way specs do
│
├── LEARNINGS.md                 # DURABLE — domain and problem insights
│                                # not operational (that's AGENTS.md); not rulings (that's decisions/)
│
└── src/                         # THE SOURCE OF TRUTH
    └── ...                      # code + tests are the real specification
```

## Durable vs disposable

**Durable (maintained records whose value survives the work cycle):**
- `src/` — code and executable tests. Authoritative for current implemented behavior.
- `skills/` — procedural knowledge. The workflow itself. Evolves ([evolving-context](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/evolving-context.md)).
- `decisions/` — ADRs. Authoritative for rulings, not current behavior. A decision made 6 months ago still explains *why* a choice was made, even if the code moved on — and can be superseded when it stops explaining or constraining the system (ADR-0012).
- `glossary.md` — ubiquitous language. Authoritative for domain terms. Terms evolve deliberately; they can stale and need pruning.
- `LEARNINGS.md` — domain and problem insights. Not operational (that's AGENTS.md); not rulings (that's decisions/). It captures what the codebase teaches you about the problem itself.
- `AGENTS.md` — operational context. Build commands, conventions. Updated when things change.
- `.loop/<name>/EVIDENCE.md` — append-only ledger of what each tick proved. Durable so a completed cycle still anchors its ADR references after its `QUEUE.md` is deleted.

**Disposable (consumed then discarded):**
- `.loop/<name>/QUEUE.md` — the work queue. When the work is done, it's done. Delete it.
- `.loop/<name>/HANDOFF.md` — cross-session handoff. Snapshot of coordination state.
- `.loop/<name>/REVIEW.md` — review artifact written under `--review`. Read for its actionable count, then discarded.
- `.loop/<name>/specs/` — planning artifacts for big work. Proposal, design. Consumed during build, then discarded ([doc-rot](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/doc-rot.md), [plan-disposability](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/plan-disposability.md)).

The key inversion from litespec: **specs are disposable, code is durable.** litespec had it backwards — specs were the canon, code was downstream. The wiki says the opposite.

## The flow, by attention mode

The same skills serve all three modes; what changes is who runs verification and whether a queue is involved.

### Interactive — human present

```
1. EXPLORE   you + agent in conversation; agent loads the explore skill
             read code, grill intent, stress-test ideas
             ADRs written inline as decisions crystallize (via decide)
             glossary updated inline (via domain-modeling)
                  │
2. PLAN      agent, guided by the plan skill, decomposes intent into
             verifiable work units — or you skip the queue and just describe
             the outcome to the build skill directly
                  │
3. BUILD     agent loads the build skill and implements the outcome
             the agent runs the Verify: command itself before declaring done
             (no external runner — the discipline of proving it before you
             claim it is the agent's; the principle survives, the enforcement
             mechanism changes)
                  │
4. REVIEW    optional: human or agent invokes the review skill
             two-axis standards + intent review against the actual codebase
                  │
5. FIX       optional: agent invokes the fix skill on the findings
                  │
6. DONE      what remains: code, tests, decisions, glossary, skills,
             AGENTS.md, LEARNINGS.md, and (if a cycle ran) EVIDENCE.md
```

For small work, skip steps. `bug → plan → build → done` is equally valid. `small fix → build → done` is equally valid. The flow is a default path, not a gate.

### Plan-then-leave — human present for the hard thinking, then walks away

```
1. EXPLORE + PLAN together (interactive)
   you and the agent grill intent and produce .loop/<name>/QUEUE.md
   each unit: outcome + constraints + done means + verify command
        │
2. you walk away; the agent builds each unit interactively (delegated),
   running verify itself before declaring each one done
        │
3. DONE — durable traces remain (code, decisions, glossary, AGENTS.md,
   LEARNINGS.md, EVIDENCE.md); the queue and handoff are deleted
```

Plan-then-leave uses the queue as a shared to-do list, not as a runner input. The agent works through it in one session with you absent; verification discipline is still the agent's.

### Batch (AFK) — human absent; the loop runs

```
1. PLAN (ahead of time, interactively or plan-then-leave)
   writes .loop/<name>/QUEUE.md
        │
2. BUILD (the loop runs, agent-agnostic)
   loop.sh reads QUEUE.md
   for each work unit:
     loop invokes the coding agent with a prompt that names the build skill
       (the agent reads skills/build/SKILL.md directly — ADR-0007)
     fresh context per tick
     agent implements, exits
     loop runs the Verify: command OUTSIDE the agent
     on success: mark done, append evidence
     on failure: verify_failed or no_progress, stop or retry
   during implementation the agent may invoke decide/domain-modeling skills
     to capture ADRs and glossary entries — these are durable
   agent may write operational learnings to AGENTS.md (evolving-context)
        │
3. REVIEW + FIX — two paths:

   Manual (default): human invokes the review skill, then re-runs loop.sh on
   the updated QUEUE.md. Same per-tick behavior as step 2.

   Orchestrated (--review): the loop runs a bounded build → review → fix
   subloop after the build queue drains (ADR-0008):
     - loop invokes a review worker that writes .loop/<name>/REVIEW.md
     - loop reads only the actionable count from REVIEW.md
     - if actionable > 0: loop invokes a fix worker that appends pending units
       to QUEUE.md, then re-runs the build pass and reviews again
     - stops when actionable = 0, --max-review-rounds is hit, --max-ticks is
       exhausted, or fix produces no new units
   In both paths the review skill owns judgment; the loop owns only
   orchestration.
        │
4. DONE
   .loop/<name>/QUEUE.md, HANDOFF.md, REVIEW.md, and specs/ are discarded.
   .loop/<name>/EVIDENCE.md remains (durable).
   what else remains: code, tests, decisions, glossary, skills, AGENTS.md,
   LEARNINGS.md
   the code IS the specification now
```

**On skill triggering:** the loop doesn't load skills itself. It invokes the agent with a prompt that names the relevant skill by name *and path* (e.g. "Load and follow the **build** skill in `skills/build/`"). The agent reads the file directly. The loop is skill-agnostic; it only knows the skill *path* to pass in the prompt (ADR-0007). That path resolves against the worker's cwd — the target repo — so it relies on the skills being installed there; making it absolute is a future `loop.sh` hardening, relevant once the project publishes and runs against external repos.

### Composable alternative paths

The flow is a default path, not a gate. Alternative paths are equally valid:

```
bug report → explore → plan → build → done
small fix → build → done                                       # skip explore and plan
architecture review → explore → review → fix → done            # skip plan and build
big feature → explore → plan (with specs) → build (loop, --review) → done
```

Decisions are captured inline throughout (not a separate phase). Glossary is updated inline during explore/plan. The loop only cares about `QUEUE.md` — it doesn't know which skill produced it.

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
- **Constraints close the solution space.** A constraint states what must stay true or what is out of bounds — never what to edit. If it names a file, it is "don't touch X" or "X's public API must not change," not "update X." (ADR-0005.)
- **Done means: is the acceptance criteria; Verify: is the enforceable subset.** The gap between them is the review surface.
- **The verify gate is the load-bearing field.** A work unit without a verify command is not loop-ready. That's the actual gate, not verticality.
- **The verify command must be deterministic and executable by the runner.** Tests, type checks, builds — not an LLM-as-judge. The loop provides mechanical backpressure.
- **"Why" is optional.** Filled in only when there's non-obvious context worth preserving. No padding.
- **No "Why this is vertical" field.** The planner skill treats vertical/horizontal as a warning about decomposition patterns, not a format requirement.

## What we're keeping from litespec

- The flow shape: explore → plan → build → review → fix (now composable, not rigid)
- Skills as procedural knowledge (think/plan/build/review → explore/plan/build/review/fix + shared decide/domain-modeling)
- Decisions (ADRs) — they persist because they're about rulings, not current behavior; they can be superseded but not silently rot
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
- Shared vocabulary skills (domain-modeling, codebase-design pattern → our decide + domain-modeling)
- ADRs captured inline during grilling, not as a separate phase
- Two-axis parallel review (Standards vs Intent, run as parallel sub-agents)
- No semantic validator (verification is distributed: execution, grilling, reproduction, human gates)
- Durable traces, disposable sessions (skills leave durable traces; agent sessions are ephemeral)
- PRDs must be split before execution (plan produces work units; loop only consumes QUEUE.md)

## What we're stealing from opengsd

- Per-work-unit model routing (each work unit can override the global agent command)
- continue-here.md for cross-session handoffs (loop writes HANDOFF.md on pause/stop)

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
- [aiming-problem](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/aiming-problem.md): "the verification signal must capture the actual desired property, not a proxy the loop will learn to game."
- [steering-docs](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/steering-docs.md): AGENTS.md as accumulated learnings, not static configuration.
- [tracer-bullets](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/tracer-bullets.md): thin end-to-end slices for early integration feedback. A heuristic against horizontal phases, not a format requirement.
- [ubiquitous-language](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/ubiquitous-language.md): the shared vocabulary that lets human and agent mean the same thing by the same word.

## Open questions

1. **~~Name.~~** Resolved (ADR-0003): the tool is named **knack**.
2. **~~Skills shipped with the tool vs project-authored skills.~~** Resolved (ADR-0011, superseding ADR-0002): there is no CLI to embed or scaffold skills. The repo's `skills/` is the product; users install via `npx skills add <owner>/<repo>` and own the result.
3. **~~How the loop invokes the agent with the right skill.~~** Resolved (ADR-0007): the loop names the skill explicitly — name *and path* — in the worker prompt (`prompts/worker.md`), which it passes to the agent via `LOOP_PROMPT_FILE`. Trigger-based agentskills.io discovery was not reliable enough across agents, and the loop cannot confirm a skill was loaded; naming the path lets the worker read the skill file directly. No per-agent `--skill` flag is needed. See "On skill triggering" above.
4. **~~Decision coverage check.~~** Resolved (ADR-0011 + ADR-0012): the citation-based `decisions check` CLI command is deleted. Orphan-ADR hygiene is judgment, transmitted as a concept by the `decide` skill — an ADR is orphaned when it no longer explains or constrains the system; references are evidence of relevance, not the test.
