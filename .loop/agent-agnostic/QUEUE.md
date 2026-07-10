# Loop Queue: agent-agnostic

Goal:
Make the loop's agent invocation truly agent-agnostic and tested. Fix the `README` `LOOP_AGENT_CMD` examples so each one consumes the prompt file and includes the trust flag its agent needs for non-interactive use. Update `loop.sh` to default to `pi` with the correct trust flag. Add `tests/run.sh` coverage that exercises the default `pi` path and the `LOOP_PROMPT_FILE` environment variable, then close `DESIGN.md` open question #4.

Stop condition:
`bash -n /home/daniel/build/knack/loop.sh && /home/daniel/build/knack/tests/run.sh`

## README shows correct, runnable agent invocations

Read first:
- `README.md` (Agent-agnostic section)
- `AGENTS.md` (lessons learned)
- `loop.sh` (agent invocation section)
- `pi --help`, `claude --help`, `codex --help`, `opencode --help`, `devin --help`

Constraints:
- Each `LOOP_AGENT_CMD` example must consume the prompt (via `LOOP_PROMPT_FILE` or `$(cat "$LOOP_PROMPT_FILE")`) and include the trust/auto-approve flag the agent needs for automation.
- Do not change `loop.sh` or `tests/run.sh` yet.
- The `Agent:` per-unit override example must also show prompt consumption.

Done means:
- `README.md` has five `LOOP_AGENT_CMD` examples (pi, claude, codex, opencode, devin) and every one references `LOOP_PROMPT_FILE`.
- The `README.md` `Agent:` example references `LOOP_PROMPT_FILE`.
- `AGENTS.md` does not contain any stale `LOOP_AGENT_CMD=` command examples.

Verify:
```bash
cd /home/daniel/build/knack && test $(grep -E '^LOOP_AGENT_CMD=' README.md | grep -E 'LOOP_PROMPT_FILE' | wc -l) -eq 5 && test $(grep -E '^Agent:' README.md | grep -E 'LOOP_PROMPT_FILE' | wc -l) -eq 1 && test $(grep -E 'LOOP_AGENT_CMD=' AGENTS.md | grep -vE 'LOOP_PROMPT_FILE' | wc -l) -eq 0 && bash -n loop.sh
```

Status: done

## loop.sh default pi and tests/run.sh cover agent invocation

Read first:
- `loop.sh`
- `tests/run.sh`
- `prompts/worker.md`

Constraints:
- The default `pi` path must still work when `LOOP_AGENT_CMD` is unset.
- The new tests must not require a real model or API key.
- The `tests/run.sh` harness must continue to use `mktemp` and not modify repo files.
- Do not change the `Agent:` override parsing.

Done means:
- `loop.sh` uses `pi -p --no-session --approve "$(cat "$run_prompt")"` (or equivalent) as its default fallback.
- `tests/run.sh` has a test that runs the loop with a fake `pi` in `PATH` and verifies the prompt is passed.
- `tests/run.sh` has a test that verifies `LOOP_PROMPT_FILE` is set for `LOOP_AGENT_CMD` invocations.
- `./tests/run.sh` passes.

Verify:
```bash
bash -n /home/daniel/build/knack/loop.sh && /home/daniel/build/knack/tests/run.sh
```

Status: pending

## DESIGN.md open question #4 is closed

Read first:
- `DESIGN.md` (open questions and agent-triggering section)
- `AGENTS.md` (lessons learned)

Constraints:
- Only update the status of question #4; do not rewrite the design thesis.
- If the agent-invocation convention becomes a lasting decision, capture it as an ADR using the `decide` skill.

Done means:
- `DESIGN.md` no longer lists question #4 as "Still open".
- `AGENTS.md` reflects the tested invocation convention if it needs updating.
- The loop and tests still pass.

Verify:
```bash
! grep -q '^4\. .*Still open' /home/daniel/build/knack/DESIGN.md && bash -n /home/daniel/build/knack/loop.sh && /home/daniel/build/knack/tests/run.sh
```

Status: pending
