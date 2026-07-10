# Getting started

knack is a loop runner (`loop.sh`) plus a read-only CLI (`knack`). You can use the loop without the CLI, but the CLI makes validating queues, scaffolding skills, and checking decisions much easier.

## Prerequisites

- Bash 4+
- Python 3 (used by `loop.sh` for small path helpers)
- Git (optional; used for work snapshots and changed-files reporting)
- Go 1.21+ (only to build the CLI from source)

## Build the CLI

```bash
cd cli
go build -o ../knack .
```

This writes a `knack` binary in the repo root. You can run it as `./knack`.

## Run the smoke test

Dry-run the bundled smoke queue to see the loop parse a work unit:

```bash
./loop.sh run examples/smoke/.loop/smoke/QUEUE.md --dry-run
```

To run a real tick with a fake worker:

```bash
mkdir -p /tmp/smoke/.loop
cp examples/smoke/.loop/smoke/QUEUE.md /tmp/smoke/.loop/QUEUE.md
LOOP_AGENT_CMD='touch smoke.done' ./loop.sh run /tmp/smoke/.loop/QUEUE.md --repo /tmp/smoke --max-ticks 1
```

The worker creates `smoke.done` in `/tmp/smoke`, the verify command sees it, and the unit is marked `done`.

## Scaffold skills into a new project

```bash
cd /path/to/new-project
/path/to/knack skills init
```

This writes the seven default skills into `.agents/skills/`. The project owns them after that point — edit, override, or delete as needed.

## Write your first queue

Create a file named `.loop/<name>/QUEUE.md` using the format in [queue-format.md](./queue-format.md). Then validate it:

```bash
./knack validate .loop/my-cycle/QUEUE.md
```

`knack instructions work-unit` prints a template you can copy.

## Run the loop

```bash
LOOP_AGENT_CMD='pi -p --no-session --approve "$(cat "$LOOP_PROMPT_FILE")"' ./loop.sh run .loop/my-cycle/QUEUE.md
```

Other common patterns:

```bash
LOOP_AGENT_CMD='claude --print --no-session-persistence --dangerously-skip-permissions "$(cat "$LOOP_PROMPT_FILE")"' ./loop.sh run .loop/my-cycle/QUEUE.md
LOOP_AGENT_CMD='codex exec --dangerously-bypass-approvals-and-sandbox --ephemeral "$(cat "$LOOP_PROMPT_FILE")"' ./loop.sh run .loop/my-cycle/QUEUE.md
LOOP_AGENT_CMD='opencode run --auto "$(cat "$LOOP_PROMPT_FILE")"' ./loop.sh run .loop/my-cycle/QUEUE.md
LOOP_AGENT_CMD='devin --print --prompt-file "$LOOP_PROMPT_FILE" --permission-mode dangerous' ./loop.sh run .loop/my-cycle/QUEUE.md
```

The `Agent:` field in an individual work unit overrides `LOOP_AGENT_CMD` for that unit only.

## Verify the project

```bash
./tests/run.sh
```

For CLI-only work:

```bash
cd cli && go test ./...
```

## Next steps

- [Queue format](./queue-format.md)
- [Loop reference](./loop.md)
- [CLI reference](./cli.md)
- [Skills guide](./skills.md)
