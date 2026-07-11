# knack

A tiny loop-packet runner for agentic development.

Compiles human intent into disposable work unit queues, then runs one unit at a time behind deterministic verification gates. Ships with a read-only CLI that validates structure, tracks decisions, and scaffolds the default skill set into any project.

## Why knack

AI agents eagerly build the wrong thing because they skip the hard part: understanding the problem. The verify gate in the loop catches the symptoms, but the bigger lever is the **explore** stance — grill the intent before writing any code. The `explore` skill is the entry point, not just another phase. The loop is for when you already know what to build.

## Quickstart

Build the CLI:

```bash
cd cli && go build -o ../knack .
```

Dry-run the smoke test:

```bash
./loop.sh run examples/smoke/.loop/smoke/QUEUE.md --dry-run
```

Run a real tick with a fake worker:

```bash
mkdir -p /tmp/smoke/.loop
cp examples/smoke/.loop/smoke/QUEUE.md /tmp/smoke/.loop/QUEUE.md
LOOP_AGENT_CMD='touch smoke.done' ./loop.sh run /tmp/smoke/.loop/QUEUE.md --repo /tmp/smoke --max-ticks 1
```

Scaffold the default skills into a new project:

```bash
cd /path/to/new-project
/path/to/knack skills init
```

Refresh the skills after a CLI upgrade (or `--force` to overwrite local changes):

```bash
/path/to/knack skills update
```

## How it works

1. The runner reads the first `Status: pending` work unit from `QUEUE.md`.
2. It marks the unit `in_progress` and invokes a fresh agent context with the worker prompt and the unit.
3. The worker implements the unit and exits. It does **not** self-certify.
4. The runner executes the unit's `Verify` command outside the worker.
5. On success: the unit is marked `done` and evidence is appended.
6. On failure: the unit is marked `verify_failed` or `no_progress`, evidence is appended, and the loop stops or retries once.
7. The loop halts on: max ticks reached with pending work, two no-progress strikes, or a failed verify after a real change.
8. On any non-clean exit, the runner writes `HANDOFF.md` with completed/in-progress/remaining units and the next action.
9. If `--review` is set, the runner invokes review after the queue drains, reads the actionable count from `REVIEW.md`, invokes fix when actionable findings exist, and runs another build pass for appended units.

Review is opt-in. Without `--review`, the loop runs build ticks only; with it, review/fix are loop-orchestrated but still skill-owned.

Before the loop: run `explore` to grill intent. The loop is for when you already know what to build.

## Queue format

Work units are `## <outcome>` headers with `Read first:`, `Constraints:`, `Done means:`, and `Verify:` fields. The `Verify:` command is the mechanically enforceable subset of `Done means:`; the gap between them is the review surface.

See [docs/queue-format.md](docs/queue-format.md) for the full protocol and an example.

## Agent-agnostic

Override the worker invocation with `LOOP_AGENT_CMD`:

```bash
LOOP_AGENT_CMD='pi -p --no-session --approve "$(cat "$LOOP_PROMPT_FILE")"' ./loop.sh run .loop/<name>/QUEUE.md
LOOP_AGENT_CMD='claude --print --no-session-persistence --dangerously-skip-permissions "$(cat "$LOOP_PROMPT_FILE")"' ./loop.sh run .loop/<name>/QUEUE.md
LOOP_AGENT_CMD='codex exec --dangerously-bypass-approvals-and-sandbox --ephemeral "$(cat "$LOOP_PROMPT_FILE")"' ./loop.sh run .loop/<name>/QUEUE.md
LOOP_AGENT_CMD='opencode run --auto "$(cat "$LOOP_PROMPT_FILE")"' ./loop.sh run .loop/<name>/QUEUE.md
LOOP_AGENT_CMD='devin --print --prompt-file "$LOOP_PROMPT_FILE" --permission-mode dangerous' ./loop.sh run .loop/<name>/QUEUE.md
```

Per-unit override via the `Agent:` field in a work unit:

```markdown
## hard refactor of persistence layer

Agent: pi -p --no-session --approve --model glm-5.2 "$(cat "$LOOP_PROMPT_FILE")"
```

Opt into review/fix orchestration with `--review`:

```bash
LOOP_AGENT_CMD='codex exec --dangerously-bypass-approvals-and-sandbox --ephemeral "$(cat "$LOOP_PROMPT_FILE")"' \
LOOP_REVIEW_CMD='claude --print --no-session-persistence --dangerously-skip-permissions "$(cat "$LOOP_PROMPT_FILE")"' \
LOOP_FIX_CMD='codex exec --dangerously-bypass-approvals-and-sandbox --ephemeral "$(cat "$LOOP_PROMPT_FILE")"' \
./loop.sh run .loop/<name>/QUEUE.md --review --max-review-rounds 2
```

## CLI

The `knack` binary is a read-only validator and context provider. Build it from `cli/`:

```bash
cd cli && go build -o ../knack .
```

### Commands

```
knack skills init [--target DIR]    Scaffold the seven default skills into DIR/.agents/skills/
knack skills check [--dir DIR]      Validate skills and report stale/modified via the manifest
knack skills update [--target DIR] [--force]   Refresh scaffolded skills from the embedded source
knack validate <queue-file>         Validate work-unit structure in a queue file
knack decisions list                List all ADRs in decisions/
knack decisions show NNNN           Print the full text of ADR NNNN
knack decisions check               Flag orphaned ADRs and dangling references
knack status                        Aggregate work-unit counts across all .loop/<name>/ cycles
knack glossary check                Validate glossary.md term references
knack instructions <artifact>       Print a template: work-unit | adr | glossary-entry
```

All commands read from the current directory (run from the repo root). `skills init` and `skills update` are the write operations — `init` scaffolds missing skills and `update` refreshes unmodified ones when a newer embedded version ships (use `--force` to overwrite local changes).

## Documentation

Full docs live in `docs/`:

- [Getting started](docs/getting-started.md)
- [Loop reference](docs/loop.md)
- [CLI reference](docs/cli.md)
- [Skills guide](docs/skills.md)
- [Queue format reference](docs/queue-format.md)
- [FAQ](docs/faq.md)

## Files

- `loop.sh` — the runner.
- `cli/` — the Go CLI (validator, status, decisions, skills, instructions).
- `prompts/worker.md` — one-tick worker instructions.
- `.agents/skills/` — the seven default skills (canonical source; the CLI embeds copies).
- `decisions/` — durable ADRs.
- `glossary.md` — ubiquitous language.
- `LEARNINGS.md` — durable insights (domain/problem learnings).
- `examples/` — sample queues.
- `docs/` — user documentation.
- `tests/run.sh` — test harness.

## Testing

```bash
./tests/run.sh
```

For CLI-only work:

```bash
cd cli && go test ./...
```

Uses `LOOP_AGENT_CMD` to substitute a fake worker, so no real `pi` calls are burned.
