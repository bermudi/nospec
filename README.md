# knack

A composable **skills collection** for working with coding agents — the procedural encoding of the [AgenticWiki](https://github.com/bermudi/AgenticWiki)'s theory. Shipped as plain [agentskills.io](https://agentskills.io) skills you install into any agent, with an optional bash loop for unattended batch work.

It exists to replace `/plan` commands, spec-kit / openspec, and ad-hoc "ralph loops" with one pick-and-choose workflow you own and edit.

## Why

Most agent tooling prescribes process. These skills do the opposite: they transmit **concepts and the reasoning behind them** ([ADR-0010](decisions/0010-skills-transmit-concepts-not-rules.md)) and let the agent apply judgment. The theory is cited, not redefined — every concept links back to the [AgenticWiki](https://github.com/bermudi/AgenticWiki), which is what distinguishes this from a theory-light skills dump.

Work happens across three levels of human attention, and the same skills serve all three:

- **Interactive** — you're present, edits land directly.
- **Plan-then-leave** — you do the hard thinking, then walk away.
- **Batch (AFK)** — the optional `loop.sh` runs a queue of units behind a verify gate while you're gone.

([ADR-0009](decisions/0009-skills-are-the-product-loop-is-optional.md): skills are the product; the loop is optional.)

## Install

```bash
npx skills add <owner>/<repo>
```

[`npx skills`](https://github.com/vercel-labs/skills) auto-detects your agent(s) — Claude Code, Codex, Cursor, Pi, Gemini, Copilot, opencode, and [70+ more](https://skills.sh) — and installs into each one's native skills path. Update with `npx skills update`; remove with `npx skills remove`.

> **Not on skills.sh yet.** This repo has no GitHub remote and no license, so the command above is a placeholder. Once both are set, it resolves to the real location.

## The skills

| skill | what it transmits |
|---|---|
| **explore** | read the codebase, grill intent, stress-test ideas *before* planning |
| **plan** | decompose intent into a disposable `QUEUE.md` of verifiable work units |
| **build** | implement a bounded observable outcome; verify-first, don't declare done until it passes |
| **review** | two-axis adversarial review — standards + intent — against the actual codebase |
| **fix** | resolve review findings — directly, or as new work units appended to the queue |
| **decide** *(shared)* | capture architectural rulings as ADRs in `decisions/`, inline as they crystallize |
| **domain-modeling** *(shared)* | manage `glossary.md`, the project's ubiquitous language |
| **document** *(shared)* | route knowledge to its authoritative artifact and maintain coherent projections |

`explore → plan → build → review → fix` is a default path, not a gate. `bug → plan → build → done` is equally valid. Skills compose.

## Optional: unattended batch mode

`loop.sh` is for AFK work — run a queue of units behind a deterministic verify gate while you walk away. Agent-agnostic:

```bash
LOOP_AGENT_CMD='pi -p --no-session --approve "$(cat "$LOOP_PROMPT_FILE")"' \
  ./loop.sh run .loop/<name>/QUEUE.md
```

Per-unit model routing (`Agent:`), handoff files on pause, and opt-in review/fix (`--review`) are covered in [docs/loop.md](docs/loop.md). Most work is interactive; reach for the loop when you actually want to leave.

## The thinking

- [AgenticWiki](https://github.com/bermudi/AgenticWiki) — the cited theory behind every concept.
- [decisions/](decisions/) — durable ADRs. The spine: [0009](decisions/0009-skills-are-the-product-loop-is-optional.md) (skills are the product), [0010](decisions/0010-skills-transmit-concepts-not-rules.md) (concepts not rules), [0011](decisions/0011-ship-as-skills-via-skills-sh-delete-cli.md) (ship via skills.sh), [0014](decisions/0014-durability-is-maintenance-not-permanence.md) (durability is maintenance), [0015](decisions/0015-artifact-roles-and-ownership.md) (artifact roles and ownership).
- [docs/architecture.md](docs/architecture.md) — conceptual overview, attention modes, and artifact roles.

## Repo layout

```
skills/        the eight skills — the product
loop.sh        optional AFK batch runner
prompts/       worker / reviewer / fixer prompts the loop sends
decisions/     durable ADRs
glossary.md    ubiquitous language (domain terms; wiki concepts linked, not redefined)
LEARNINGS.md   durable domain/problem insights
docs/          user and architecture docs
tests/run.sh   test harness for loop.sh
```

## Testing

```bash
./tests/run.sh
```

## License

None yet. A license is required before publishing to skills.sh.
