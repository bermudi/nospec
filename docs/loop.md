# Loop reference

`loop.sh` is the agent-agnostic runner. It reads `QUEUE.md` and executes one work unit per tick, running the verification command outside the worker.

## Usage

```bash
./loop.sh run <queue> [--repo DIR] [--max-ticks N] [--dry-run]
```

- `queue` — path to `QUEUE.md`. Usually `.loop/<name>/QUEUE.md`.
- `--repo DIR` — the working directory for the worker and the verify command. Defaults to the parent of the `.loop/<name>` directory that contains the queue.
- `--max-ticks N` — maximum units to attempt. Default is `3`.
- `--dry-run` — parse the first pending unit and print its title, repo, and verify command, then exit.

## Environment variables

- `LOOP_AGENT_CMD` — optional command used to invoke the worker. If unset, the loop defaults to `pi -p --no-session` with the prompt text as a single argument. If set, the command is evaluated by `bash -lc` in the repo directory, and `LOOP_PROMPT_FILE` is set to a temporary file containing `prompts/worker.md` plus the current work unit.
- `LOOP_PROMPT_FILE` — set by the loop when `LOOP_AGENT_CMD` is used. Points to the generated prompt file. Do not override unless you are calling the worker manually.

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

## Work unit statuses

- `pending` — not yet started.
- `in_progress` — currently being worked.
- `done` — verify passed.
- `verify_failed` — verify failed; may be retried once.
- `no_progress` — verify failed twice with no snapshot change.
- `blocked` — worker exited non-zero.

## Output files

`loop.sh` writes next to the queue:

- `EVIDENCE.md` — append-only ledger. Includes the full unit, changed files, verify command, verify output, and worker output. It is durable; keep it after deleting `QUEUE.md` so `knack decisions check` can still see which ADRs the cycle referenced.
- `HANDOFF.md` — written on non-clean exit. Sections: completed, in progress, remaining, next action. Delete when the work resumes.

## Agent invocation examples

```bash
LOOP_AGENT_CMD='claude --print --no-session-persistence --dangerously-skip-permissions "$(cat "$LOOP_PROMPT_FILE")"' ./loop.sh run .loop/go-cli/QUEUE.md
LOOP_AGENT_CMD='codex exec --dangerously-bypass-approvals-and-sandbox --ephemeral "$(cat "$LOOP_PROMPT_FILE")"' ./loop.sh run .loop/go-cli/QUEUE.md
LOOP_AGENT_CMD='devin --print --prompt-file "$LOOP_PROMPT_FILE" --permission-mode dangerous' ./loop.sh run .loop/go-cli/QUEUE.md
```

The command is passed to `bash -lc` in the repo directory, so the typical pattern is `"$(cat "$LOOP_PROMPT_FILE")"`.

## Verification notes

- `loop.sh` does **not** validate `QUEUE.md` structure. Use `knack validate` first.
- `loop.sh` does **not** run review or manage ADRs/glossary. Those are handled by the skills.
- The `Verify:` command must be deterministic and executable by the runner — tests, builds, type checks, not an LLM-as-judge.
