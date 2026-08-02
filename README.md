---
nospec: true
role: view
---

# nospec

A composable **skills collection** for turning intent into verified code while keeping temporary coordination disposable and durable knowledge explicit. It encodes the [AgenticWiki](https://github.com/bermudi/AgenticWiki)'s theory as plain [agentskills.io](https://agentskills.io) skills, with an optional batch runner for unattended work.

It replaces `/plan` commands, spec-kit / openspec, and ad-hoc "ralph loops" with capabilities you pick only when their problem is present. It is not a pipeline.

Your harness increasingly ships native plan mode, autopilot-style loops, and cross-model review — on their own, enough for interactive feature work. Nospec doesn't reimplement those mechanics. It supplies what the harness doesn't: a **discriminating verify gate the worker cannot self-certify**, **durable-vs-disposable contract discipline**, **AFK batch** behind an evidence ledger and baseline guard, and **multi-session coordination**. The harness gives you the loop; nospec is the judgment layer that makes it produce verified, durable work instead of plausible-looking output.

## Why

Most agent tooling prescribes process. These skills do the opposite: they transmit **concepts and the reasoning behind them** ([ADR-0010](decisions/0010-skills-transmit-concepts-not-rules.md)) and let the agent apply judgment. The theory is cited, not redefined — every concept links back to the [AgenticWiki](https://github.com/bermudi/AgenticWiki), which is what distinguishes this from a theory-light skills dump.

Work happens across three levels of human attention, and the same skills serve all three:

- **Interactive** — you're present, edits land directly.
- **Plan-then-leave** — you do the hard thinking, then walk away.
- **Batch (AFK)** — `nospec run` drives a queue of units behind a verify gate while you're gone.

([ADR-0009](decisions/0009-skills-are-the-product-loop-is-optional.md): skills are the product; the loop is optional.)

## Install

Install the three core stances:

```bash
skills add bermudi/nospec --skill nospec-shape nospec-carve nospec-trial
```

Add durable-knowledge curation or AFK execution only when needed:

```bash
skills add bermudi/nospec --skill nospec-curator
skills add bermudi/nospec --skill nospec-loop
```

Use `--skill '*'` to install all five. Update project skills with `skills update -p`; remove with `skills remove`.

Existing nine-skill installations should remove retired names before reinstalling, because an update may not delete old directories:

```bash
skills remove nospec-scout nospec-mend nospec-rule nospec-lexicon nospec -y
```

The runner ships as the optional `nospec-loop` skill's `scripts/nospec`. skills.sh installs skill files but does not touch PATH, so to invoke the runner as `nospec` from anywhere, run the install verb once (your agent will do this for you when you ask it to set up nospec):

```bash
.agents/skills/nospec-loop/scripts/nospec install    # project-local install
# or: ~/.agents/skills/nospec-loop/scripts/nospec install  # global install (-g)
```

That symlinks the runner onto PATH. Then `nospec run ...` works from any directory.

> **On skills.sh.** Skill selection is explicit: install the three core names above, either optional companion, or `--skill '*'` for all five.

## The skills

Nospec provides three stances for verified change — Shape the intent, Carve the implementation, and Trial the result — plus an optional Curator that preserves what should outlive the work and an optional Loop that runs shaped work while the human is absent. Clear work can go directly to `nospec-carve`, the common implementation skill. Reach for `nospec-shape` when the problem is unclear, outcomes need decomposition, or shaping itself spans multiple sessions; reach for `nospec-trial` when a change needs adversarial review. `nospec-curator` preserves the decisions, shared language, and durable context worth keeping after disposable planning is gone. The optional `nospec-loop` skill judges batch-worthiness and carries the runner used when the human leaves. These are independently invoked capabilities, not a prescribed workflow. See the [skills guide](docs/skills.md) for the full catalog.

## Work specs and durable contracts

Nospec discards **work specs**—queues, handoffs, scratch designs, and multi-session shaping maps used to coordinate a current change. It does not discard contracts designated by the project's operational context, accepted rulings, or established metadata conventions, such as OpenAPI schemas, protocol definitions, compatibility policies, or executable contract tests. `nospec: true` opts into structural checks; it does not confer contract authority. Nospec creates no universal prose canon beside the code.

## Optional: unattended batch mode

`nospec run` is for AFK work — drive a queue of units behind a deterministic verify gate while you walk away. Agent-agnostic:

```bash
LOOP_AGENT_CMD='pi -p --no-session --approve "$(cat "$LOOP_PROMPT_FILE")"' \
  nospec run .loop/<name>/QUEUE.md
```

Run `nospec lint .loop/<name>/QUEUE.md` before leaving; `nospec run` repeats that whole-queue preflight and requires a clean starting worktree before mutation. Deliberate dirty handoffs require `--accept-dirty-baseline`; a separate worktree can be selected with `--repo`. Per-unit model routing (`Agent:`), handoffs, and opt-in review/fix (`--review`) are covered in [docs/loop.md](docs/loop.md). Most work is interactive.

## The thinking

- [AgenticWiki](https://github.com/bermudi/AgenticWiki) — the cited theory behind every concept.
- [decisions/](decisions/) — durable ADRs. The spine (ADR-0009 onward) is derived from frontmatter; run `nospec spine` to list it.
- [docs/architecture.md](docs/architecture.md) — conceptual overview, attention modes, and artifact roles.

## Repo layout

```
skills/        the five skills — the product (incl. the nospec-loop runner skill)
decisions/     durable ADRs (YAML frontmatter: nospec, id, date, status, spine, ...)
glossary.md    ubiquitous language (domain terms; wiki concepts linked, not redefined)
docs/          user and architecture docs
tests/run.sh   test harness for nospec
```

The runner lives at `skills/nospec-loop/scripts/nospec`; the worker/reviewer/fixer prompts at `skills/nospec-loop/prompts/`. Both ship inside the `nospec-loop` skill.

## Testing

```bash
./tests/run.sh
```

## License

[MIT](LICENSE).
