# knack documentation

This folder holds the living user documentation for **knack**, the agent-agnostic harness for agentic development.

## What is knack?

knack is an agent-agnostic harness for agentic development. It starts with **explore** — reading the codebase and grilling intent before any work is planned. When the problem is clear, `plan` turns intent into a disposable `QUEUE.md`, and the `build` loop runs one unit at a time behind deterministic verification gates. It ships with a loop runner, a read-only CLI, and seven agent-agnostic skills.

## Start here: explore

The biggest failure mode of agentic development is building the wrong thing quickly. The `explore` skill exists to slow that down: read the code, challenge the stated goal, and reach clarity before `plan` produces work units.

## Navigating the docs

- [Getting started](./getting-started.md) — build, install, run the smoke test, and scaffold skills.
- [Skills guide](./skills.md) — the seven default skills. Start with `explore`.
- [Loop reference](./loop.md) — `loop.sh` command, flags, environment variables, and queue lifecycle.
- [CLI reference](./cli.md) — `knack` binary commands and flags.
- [Queue format](./queue-format.md) — the `QUEUE.md` work unit protocol.
- [FAQ](./faq.md) — common questions and short answers.

## Source of truth

When in doubt, read the code: `loop.sh`, `cli/`, `prompts/worker.md`, and `.agents/skills/`. Docs describe behavior, but the code is the source of truth.
