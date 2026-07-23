---
role: view
---

# nospec

A composable **skills collection** for working with coding agents — the procedural encoding of the [AgenticWiki](https://github.com/bermudi/AgenticWiki)'s theory. Shipped as plain [agentskills.io](https://agentskills.io) skills you install into any agent, with an optional bash loop for unattended batch work.

It exists to replace `/plan` commands, spec-kit / openspec, and ad-hoc "ralph loops" with one pick-and-choose workflow you own and edit.

## Why

Most agent tooling prescribes process. These skills do the opposite: they transmit **concepts and the reasoning behind them** ([ADR-0010](decisions/0010-skills-transmit-concepts-not-rules.md)) and let the agent apply judgment. The theory is cited, not redefined — every concept links back to the [AgenticWiki](https://github.com/bermudi/AgenticWiki), which is what distinguishes this from a theory-light skills dump.

Work happens across three levels of human attention, and the same skills serve all three:

- **Interactive** — you're present, edits land directly.
- **Plan-then-leave** — you do the hard thinking, then walk away.
- **Batch (AFK)** — `nospec run` drives a queue of units behind a verify gate while you're gone.

([ADR-0009](decisions/0009-skills-are-the-product-loop-is-optional.md): skills are the product; the loop is optional.)

## Install

```bash
npx skills add bermudi/nospec
```

[`npx skills`](https://github.com/vercel-labs/skills) auto-detects your agent(s) — Claude Code, Codex, Cursor, Pi, Gemini, Copilot, opencode, and [70+ more](https://skills.sh) — and installs into each one's native skills path. Update with `npx skills update`; remove with `npx skills remove`.

The runner ships as the `nospec` skill's `scripts/nospec`. skills.sh installs skill files but does not touch PATH, so to invoke the runner as `nospec` from anywhere, run the install verb once (your agent will do this for you when you ask it to set up nospec):

```bash
.agents/skills/nospec/scripts/nospec install    # project-local install
# or: ~/.agents/skills/nospec/scripts/nospec install  # global install (-g)
```

That symlinks the runner onto PATH. Then `nospec run ...` works from any directory.

> **On skills.sh.** Install with `npx skills add bermudi/nospec`. Update with `npx skills update`; remove with `npx skills remove`.

## The skills

| skill | what it transmits |
|---|---|
| **nospec-scout** | read the codebase, grill intent, stress-test ideas *before* planning |
| **nospec-shape** | decompose intent into a disposable `QUEUE.md` of verifiable work units |
| **nospec-hew** | implement a bounded observable outcome; verify-first, don't declare done until it passes |
| **nospec-trial** | two-axis adversarial review — standards + intent — against the actual codebase |
| **nospec-mend** | resolve review findings — directly, or as new work units appended to the queue |
| **nospec-rule** *(shared)* | capture architectural rulings as ADRs in `decisions/`, inline as they crystallize |
| **nospec-lexicon** *(shared)* | manage `glossary.md`, the project's ubiquitous language |
| **nospec-curator** *(shared)* | route knowledge to its authoritative artifact and maintain coherent projections |
| **nospec** *(optional)* | the batch runner — drives a `QUEUE.md` behind a verify gate while you're away |

`nospec-scout → nospec-shape → nospec-hew → nospec-trial → nospec-mend` is a default path, not a gate. `bug → nospec-shape → nospec-hew → done` is equally valid. Skills compose.

## Optional: unattended batch mode

`nospec run` is for AFK work — drive a queue of units behind a deterministic verify gate while you walk away. Agent-agnostic:

```bash
LOOP_AGENT_CMD='pi -p --no-session --approve "$(cat "$LOOP_PROMPT_FILE")"' \
  nospec run .loop/<name>/QUEUE.md
```

Per-unit model routing (`Agent:`), handoff files on pause, and opt-in review/fix (`--review`) are covered in [docs/loop.md](docs/loop.md). Most work is interactive; reach for the loop when you actually want to leave.

## The thinking

- [AgenticWiki](https://github.com/bermudi/AgenticWiki) — the cited theory behind every concept.
- [decisions/](decisions/) — durable ADRs. The spine (ADR-0009 onward) is derived from frontmatter; run `nospec spine` to list it.
- [docs/architecture.md](docs/architecture.md) — conceptual overview, attention modes, and artifact roles.

## Repo layout

```
skills/        the nine skills — the product (incl. the nospec runner skill)
decisions/     durable ADRs (YAML frontmatter: id, date, status, spine, ...)
glossary.md    ubiquitous language (domain terms; wiki concepts linked, not redefined)
docs/          user and architecture docs
tests/run.sh   test harness for nospec
```

The runner lives at `skills/nospec/scripts/nospec`; the worker/reviewer/fixer prompts at `skills/nospec/prompts/`. Both ship inside the `nospec` skill.

## Testing

```bash
./tests/run.sh
```

## License

[MIT](LICENSE).
