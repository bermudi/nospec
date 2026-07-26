---
nospec: true
role: view
---

# Getting started

nospec is a skills collection plus an optional Bash loop for unattended batch work. Most work is interactive; reach for the loop when you want to leave. The runner supports macOS and Linux and requires Python 3.10 or newer.

## Install the skills

```bash
skills add -p bermudi/nospec --skill '*'
```

This installs every nospec skill into the current project. Update with `skills update -p`; remove with `skills remove`.

Once installed, invoke a skill by name: `nospec-scout`, `nospec-shape`, `nospec-carve`, `nospec-trial`, `nospec-mend`, `nospec-rule`, `nospec-lexicon`, `nospec-curator`, or `nospec` (the runner skill, for batch mode).

## Use the skills interactively

Invoke only the skill the work needs. A clear change can go directly to `nospec-carve`; unresolved intent may need `nospec-scout`, and decomposition may need `nospec-shape`. No queue or Nospec artifact is required for interactive work.

## Run the loop (optional)

The loop runs a `QUEUE.md` one work unit at a time while you are away. Its first mutating run requires a clean Git baseline outside `.loop/`; use `--accept-dirty-baseline` only for a deliberate dirty handoff, or point `--repo` at a separate worktree:

```bash
LOOP_AGENT_CMD='pi -p --no-session --approve "$(cat "$LOOP_PROMPT_FILE")"' nospec run .loop/<name>/QUEUE.md
```

See [`loop.md`](./loop.md) for flags, environment variables, and the review-fix subloop.

## Write a queue

Create `.loop/<name>/QUEUE.md` using the format in [`queue-format.md`](./queue-format.md). Each work unit needs an outcome, nonempty done criteria, and a deterministic `Verify:` command that meaningfully distinguishes success from failure. Add context and constraints only when they carry information. Before leaving, run:

```bash
nospec lint .loop/<name>/QUEUE.md
```

## Test the repo

```bash
./tests/run.sh
```

## Next steps

- Read [`architecture.md`](./architecture.md) for the conceptual overview.
- Read [`skills.md`](./skills.md) for the skill catalog.
- Read [`loop.md`](./loop.md) if you want batch mode.
