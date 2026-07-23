---
role: view
---

# Getting started

nospec is a skills collection plus an optional bash loop for unattended batch work. Most work is interactive; reach for the loop when you want to leave.

## Install the skills

```bash
npx skills add bermudi/nospec
```

[`npx skills`](https://github.com/vercel-labs/skills) detects your agent and installs into its native skills path. Update with `npx skills update`; remove with `npx skills remove`.

Once installed, invoke a skill by name: `nospec-scout`, `nospec-shape`, `nospec-hew`, `nospec-trial`, `nospec-mend`, `nospec-rule`, `nospec-lexicon`, `nospec-curator`, or `nospec` (the runner skill, for batch mode).

## Use the skills interactively

```
nospec-scout → nospec-shape → nospec-hew → nospec-trial → nospec-mend → done
```

This is a default path, not a gate. `bug → nospec-shape → nospec-hew → done` is equally valid.

## Run the loop (optional)

The loop runs a `QUEUE.md` one work unit at a time while you are away:

```bash
LOOP_AGENT_CMD='pi -p --no-session --approve "$(cat "$LOOP_PROMPT_FILE")"' nospec run .loop/<name>/QUEUE.md
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
