# knack documentation

This folder holds the living user documentation for **knack**, the agent-agnostic harness for agentic development.

## What is knack?

knack turns human intent into a disposable queue of verifiable work units and runs them one at a time behind deterministic verification gates. It ships with a loop runner, a read-only CLI, and seven agent-agnostic skills.

## Navigating the docs

- [Getting started](./getting-started.md) — build, install, run the smoke test, and scaffold skills.
- [Loop reference](./loop.md) — `loop.sh` command, flags, environment variables, and queue lifecycle.
- [CLI reference](./cli.md) — `knack` binary commands and flags.
- [Skills guide](./skills.md) — the seven default skills and how to author your own.
- [Queue format](./queue-format.md) — the `QUEUE.md` work unit protocol.
- [FAQ](./faq.md) — common questions and short answers.

## Source of truth

When in doubt, read the code: `loop.sh`, `cli/`, `prompts/worker.md`, and `.agents/skills/`. Docs describe behavior, but the code is the source of truth.
