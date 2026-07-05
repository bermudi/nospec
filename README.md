# sliceloop

A tiny loop-packet runner for agentic development.

Compiles human intent into disposable vertical slice queues, then runs one slice at a time behind deterministic verification gates.

## Quickstart

```bash
./loop.sh run <queue> [--repo DIR] [--max-ticks N] [--dry-run]
```

The queue is usually `.loop/QUEUE.md` in the target repo. Use `--repo` when the queue lives outside the repository it should operate on.

Example:

```bash
./loop.sh run examples/smoke/.loop/QUEUE.md --dry-run
```

## How it works

1. The runner reads the first `Status: pending` slice from `QUEUE.md`.
2. It marks the slice `in_progress` and invokes a fresh agent context (`pi -p --no-session` by default) with the worker prompt and the slice.
3. The worker does the work and exits. It does **not** self-certify.
4. The runner executes the slice's `Verify` command outside the worker.
5. On success: the slice is marked `done` and evidence is appended.
6. On failure: the slice is marked `verify_failed` or `no_progress`, evidence is appended, and the loop stops or retries once.
7. The loop halts on: max ticks reached with pending work, two no-progress strikes, or a failed verify after a real change.

## Queue format

```markdown
# Loop Queue: <short name>

Goal:
<one paragraph describing the desired end state>

Stop condition:
`<command that proves the whole packet is done, if one exists>`

## Slice 1: <vertical outcome>

Why this is vertical:
<explain why this crosses enough layers to produce a real observable improvement>

Work:
- <narrow work instruction>
- <guardrail>

Verify:
\`\`\`bash
<command>
\`\`\`

Done means:
- <observable condition>

Status: pending
```

Only three things are mechanically parsed: the `## Slice N:` heading, the `Verify:` fenced block, and the `Status:` line. Everything else is for the agent and the human.

## Planning

Use the `vertical-slice-planner` skill to convert messy intent into a queue. The skill rejects horizontal phase plans ("Phase 1: types / Phase 2: wiring") and forces vertical outcomes.

See `.agents/skills/vertical-slice-planner/SKILL.md`.

## Files

- `loop.sh` — the runner.
- `prompts/worker.md` — one-tick worker instructions.
- `.agents/skills/vertical-slice-planner/SKILL.md` — planner skill.
- `examples/` — sample queues.
- `tests/run.sh` — test harness.

## Testing

```bash
./tests/run.sh
```

Uses `SLICELOOP_AGENT_CMD` to substitute a fake worker, so no real `pi` calls are burned.
