# knack

A tiny loop-packet runner for agentic development.

Compiles human intent into disposable work unit queues, then runs one unit at a time behind deterministic verification gates.

## Quickstart

```bash
./loop.sh run <queue> [--repo DIR] [--max-ticks N] [--dry-run]
```

The queue is usually `.loop/<name>/QUEUE.md` in the target repo — each work cycle gets its own named subdirectory under `.loop/`. Use `--repo` when the queue lives outside the repository it should operate on.

Example:

```bash
./loop.sh run examples/smoke/.loop/smoke/QUEUE.md --dry-run
```

## How it works

1. The runner reads the first `Status: pending` work unit from `QUEUE.md`.
2. It marks the unit `in_progress` and invokes a fresh agent context (`pi -p --no-session` by default, or `LOOP_AGENT_CMD`, or the unit's `Agent:` override) with the worker prompt and the unit.
3. The worker does the work and exits. It does **not** self-certify.
4. The runner executes the unit's `Verify` command outside the worker.
5. On success: the unit is marked `done` and evidence is appended.
6. On failure: the unit is marked `verify_failed` or `no_progress`, evidence is appended, and the loop stops or retries once.
7. The loop halts on: max ticks reached with pending work, two no-progress strikes, or a failed verify after a real change.
8. On any non-clean exit, the runner writes `.loop/<name>/HANDOFF.md` with completed/in-progress/remaining units and the next action.

## Queue format

````markdown
# Loop Queue: <short name>

Goal:
<one paragraph describing the desired end state>

Stop condition:
`<command that proves the whole packet is done, if one exists>`

## <outcome — what changes, observable>

Agent: <optional — overrides LOOP_AGENT_CMD for this unit only>

Why:
<only if non-obvious — else omit>

Read first:
- <context the worker needs: ADR, area, or file>
- <2–4 entries; context, not scope>

Constraints:
- <boundary or guardrail>
- <what must stay true or what is out of bounds>

Done means:
- <observable condition>
- <no regression condition>

Verify:
```bash
<command that exits 0 on success>
```

Status: pending

## <next outcome>
...
````

Four things are mechanically parsed: the `## ` heading, the `Agent:` line (optional), the `Verify:` fenced block, and the `Status:` line. Everything else is for the agent and the human.

## Planning

Use the `plan` skill to convert messy intent into a queue. The skill prefers vertical slices and rejects horizontal phase plans ("Phase 1: types / Phase 2: wiring") but supports other work unit types: patch, investigation, bug fix, refactor.

See `.agents/skills/plan/SKILL.md`.

## Agent-agnostic

Override the worker invocation with `LOOP_AGENT_CMD`:

```bash
LOOP_AGENT_CMD="pi -p --no-session" ./loop.sh run .loop/<name>/QUEUE.md
LOOP_AGENT_CMD="claude --print" ./loop.sh run .loop/<name>/QUEUE.md
LOOP_AGENT_CMD="codex" ./loop.sh run .loop/<name>/QUEUE.md
```

Per-unit override via the `Agent:` field in a work unit:

```markdown
## hard refactor of persistence layer

Agent: pi -p --no-session --model glm-5.2
```

## Files

- `loop.sh` — the runner.
- `prompts/worker.md` — one-tick worker instructions.
- `.agents/skills/plan/SKILL.md` — planner skill.
- `examples/` — sample queues.
- `tests/run.sh` — test harness.

## Testing

```bash
./tests/run.sh
```

Uses `LOOP_AGENT_CMD` to substitute a fake worker, so no real `pi` calls are burned.
