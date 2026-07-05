# AGENTS.md

## Project

**sliceloop** is a tiny loop-packet runner for agentic development.

It compiles human intent into disposable vertical slice queues, then runs one slice at a time behind deterministic verification gates.

## Thesis

The reusable asset is the vertical slicing procedure. The runner should stay boring: bounded queue iteration, fresh agent context per tick, verification outside the worker, and append-only evidence.

## Non-goals

- No spec lifecycle.
- No proposal/design/tasks/archive flow.
- No durable slice canon.
- No semantic validator that pretends to judge plan quality.
- No broad CLI surface until the file protocol proves itself.

## Core artifacts

- `.loop/QUEUE.md` — disposable loop packet: bounded queue of vertical slices.
- `.loop/EVIDENCE.md` — append-only ledger of what each tick actually proved.
- `prompts/worker.md` — one-tick worker instructions.
- `.agents/skills/vertical-slice-planner/SKILL.md` — planner skill that turns messy intent into verified vertical slices.

## Working conventions

- Shell first. Add another language only when shell becomes the bottleneck.
- Plain Markdown files over stores or schemas.
- One command first: `./loop.sh run <queue>`.
- The worker never self-certifies. The runner owns verification.
- A slice must leave the repo better if the loop stops immediately after it.
- Prefer vertical outcomes over horizontal layers.
- `SLICELOOP_AGENT_CMD` overrides the worker invocation (used by tests; set to a shell command that reads `SLICELOOP_PROMPT_FILE` and acts on the repo).

## Verification

After meaningful changes, run:

```bash
./tests/run.sh
```
