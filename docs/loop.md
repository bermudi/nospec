---
role: view
---

# Loop reference

`nospec run` is the agent-agnostic runner. It reads `QUEUE.md` and executes one work unit per tick, running the verification command outside the worker.

## Usage

```bash
nospec run <queue> [--repo DIR] [--max-ticks N] [--review] [--max-review-rounds N] [--dry-run]
```

- `queue` — path to `QUEUE.md`. Usually `.loop/<name>/QUEUE.md`.
- `--repo DIR` — the working directory for the worker and the verify command. Defaults to the parent of the `.loop/<name>` directory that contains the queue.
- `--max-ticks N` — maximum units to attempt. Default is `3`.
- `--review` — opt into the bounded review/fix subloop after pending build units drain. Default behavior does not run review.
- `--max-review-rounds N` — maximum review/fix rounds when `--review` is enabled. Default is `2`.
- `--dry-run` — parse the first pending unit and print its title, repo, and verify command, then exit.

## Environment variables

- `LOOP_AGENT_CMD` — optional command used to invoke the worker. If unset, the loop defaults to `pi -p --no-session --approve` with the prompt text as a single argument. If set, the command is evaluated by `bash -lc` in the repo directory, and `LOOP_PROMPT_FILE` is set to a temporary file containing the worker prompt (from `skills/nospec/prompts/worker.md`) plus the current work unit. The worker's harness auto-loads the `nospec-hew` skill by trigger text; no path configuration is needed.
- `LOOP_PROMPT_FILE` — set by the loop when `LOOP_AGENT_CMD` is used. Points to the generated prompt file. Do not override unless you are calling the worker manually.
- `LOOP_REVIEW_CMD` — optional command used for the review phase when `--review` is enabled. Defaults to `LOOP_AGENT_CMD`, or to the loop's default `pi` invocation if neither is set.
- `LOOP_FIX_CMD` — optional command used for the fix phase when `--review` is enabled. Defaults to `LOOP_AGENT_CMD`, or to the loop's default `pi` invocation if neither is set.

During every phase, the loop also sets `LOOP_PHASE`, `LOOP_QUEUE_FILE`, `LOOP_EVIDENCE_FILE`, and `LOOP_REVIEW_FILE` for the worker process.

## Per-unit agent override

A work unit can include an `Agent:` line to override `LOOP_AGENT_CMD` for that unit only. The override command is evaluated the same way as `LOOP_AGENT_CMD`:

```markdown
## rewrite the parser in PEG

Agent: claude --print --no-session-persistence --dangerously-skip-permissions "$(cat "$LOOP_PROMPT_FILE")"
...
```

## Per-tick behavior

1. Read the first `Status: pending` unit.
2. Mark it `in_progress`.
3. Snapshot the repo state (diff + untracked files outside `.loop`).
4. Invoke the agent with the worker prompt and the unit.
5. If the agent exits non-zero, mark the unit `blocked`, append evidence, and stop.
6. Run the unit's `Verify:` command.
7. On success: mark `done`, append evidence, continue.
8. On failure:
   - If the repo snapshot is unchanged, retry once (`verify_failed` then `pending`). After two consecutive no-progress failures, mark `no_progress` and stop.
   - If the repo changed, mark `verify_failed` and stop.
9. If max ticks is reached with pending work, stop.

## Optional review/fix subloop

By default, `nospec run` stops when the build queue has no pending work. With `--review`, it then runs a bounded review/fix subloop:

1. Invoke a review worker with the review prompt and the completed queue/evidence.
2. Require the review worker to write `.loop/<name>/REVIEW.md`.
3. Read only the `- actionable: N` summary line from `REVIEW.md`.
4. If `N` is `0`, stop cleanly.
5. If `N` is non-zero, invoke a fix worker.
6. Require the fix worker to append new `Status: pending` work units to the same `QUEUE.md`.
7. Run the build pass again, then review again.

The subloop stops when review is clean, `--max-review-rounds` is reached, `--max-ticks` is exhausted, or fix produces no pending work. The loop owns orchestration and stop conditions only; the `nospec-trial` and `nospec-mend` skills own judgment and work-unit generation.

## Work unit statuses

- `pending` — not yet started.
- `in_progress` — currently being worked.
- `done` — verify passed.
- `verify_failed` — verify failed; may be retried once.
- `no_progress` — verify failed twice with no snapshot change.
- `blocked` — worker exited non-zero.

## Output files

`nospec run` writes next to the queue:

- `EVIDENCE.md` — append-only ledger. Includes the full unit, changed files, verify command, verify output, worker output, a registry-derived proof boundary (what this verify mechanically proves, derived from the command), and a pin-state record (which durable docs were touched and whether any prior pins have moved). It is durable; keep it after deleting `QUEUE.md` so completed work still anchors its ADR references. Pin alerts in the ledger are triage triggers for the `nospec-trial` skill, not coherence gates (ADR-0016).
- `HANDOFF.md` — written on non-clean exit. Sections: completed, in progress, remaining, next action. Delete when the work resumes.
- `REVIEW.md` — structured review artifact written by the review worker when `--review` is enabled. The loop reads only its `- actionable: N` summary line.

## Agent invocation examples

```bash
LOOP_AGENT_CMD='pi -p --no-session --approve "$(cat "$LOOP_PROMPT_FILE")"' nospec run .loop/<name>/QUEUE.md
LOOP_AGENT_CMD='claude --print --no-session-persistence --dangerously-skip-permissions "$(cat "$LOOP_PROMPT_FILE")"' nospec run .loop/<name>/QUEUE.md
LOOP_AGENT_CMD='codex exec --dangerously-bypass-approvals-and-sandbox --ephemeral "$(cat "$LOOP_PROMPT_FILE")"' nospec run .loop/<name>/QUEUE.md
LOOP_AGENT_CMD='opencode run --auto "$(cat "$LOOP_PROMPT_FILE")"' nospec run .loop/<name>/QUEUE.md
LOOP_AGENT_CMD='devin --print --prompt-file "$LOOP_PROMPT_FILE" --permission-mode dangerous' nospec run .loop/<name>/QUEUE.md
```

The command is passed to `bash -lc` in the repo directory, so the typical pattern is `"$(cat "$LOOP_PROMPT_FILE")"`.

Review and fix can use separate agents:

```bash
LOOP_AGENT_CMD='codex exec --dangerously-bypass-approvals-and-sandbox --ephemeral "$(cat "$LOOP_PROMPT_FILE")"' \
LOOP_REVIEW_CMD='claude --print --no-session-persistence --dangerously-skip-permissions "$(cat "$LOOP_PROMPT_FILE")"' \
LOOP_FIX_CMD='codex exec --dangerously-bypass-approvals-and-sandbox --ephemeral "$(cat "$LOOP_PROMPT_FILE")"' \
nospec run .loop/<name>/QUEUE.md --review
```

## Verification notes

- `nospec run` does **not** validate `QUEUE.md` structure. The loop trusts the format; use the `nospec-shape` skill or inspect the file directly before running.
- `nospec run` runs review and fix only when `--review` is set. It invokes the skills and reads the actionable count from `REVIEW.md`; it does not judge findings or manage ADRs/glossary.
- The `Verify:` command must be deterministic and executable by the runner — tests, builds, type checks, not an LLM-as-judge.
