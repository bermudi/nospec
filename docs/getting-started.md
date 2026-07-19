---
role: view
---

# Getting started

knack is a skills collection plus an optional bash loop for unattended batch work. Most work is interactive; reach for the loop when you want to leave.

## Install the skills

```bash
npx skills add <owner>/<repo>
```

[`npx skills`](https://github.com/vercel-labs/skills) detects your agent and installs into its native skills path. Update with `npx skills update`; remove with `npx skills remove`.

Once installed, invoke a skill by name: `explore`, `plan`, `build`, `review`, `fix`, `decide`, `domain-modeling`, or `document`.

## Use the skills interactively

```
explore → plan → build → review → fix → done
```

This is a default path, not a gate. `bug → plan → build → done` is equally valid.

## Run the loop (optional)

The loop runs a `QUEUE.md` one work unit at a time while you are away:

```bash
LOOP_AGENT_CMD='pi -p --no-session' ./loop.sh run .loop/<name>/QUEUE.md
```

See [`loop.md`](./loop.md) for flags, environment variables, and the review-fix subloop.

## Write a queue

Create `.loop/<name>/QUEUE.md` using the format in [`queue-format.md`](./queue-format.md). Each work unit needs an outcome, constraints, done criteria, and a deterministic `Verify:` command.

## Test the repo

```bash
./tests/run.sh
```

## Next steps

- Read [`architecture.md`](./architecture.md) for the conceptual overview.
- Read [`skills.md`](./skills.md) for the skill catalog.
- Read [`loop.md`](./loop.md) if you want batch mode.
