
## 2026-07-09T21:37:07-06:00 — README shows correct, runnable agent invocations

Status: done

Unit:
````markdown
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

Status: pending````

Files changed:
```text
 M README.md
 M docs/getting-started.md
 M docs/loop.md
?? .loop/agent-agnostic/EVIDENCE.md
?? .loop/agent-agnostic/QUEUE.md
?? .loop/glossary/QUEUE.md
```

Verify command:
```bash
cd /home/daniel/build/knack && test $(grep -E '^LOOP_AGENT_CMD=' README.md | grep -E 'LOOP_PROMPT_FILE' | wc -l) -eq 5 && test $(grep -E '^Agent:' README.md | grep -E 'LOOP_PROMPT_FILE' | wc -l) -eq 1 && test $(grep -E 'LOOP_AGENT_CMD=' AGENTS.md | grep -vE 'LOOP_PROMPT_FILE' | wc -l) -eq 0 && bash -n loop.sh
```

Verify output:
```text
```

Worker output:
````text
The loop logic confirms my read: verify passing marks the unit `done` immediately and `continue`s (lines 328-335). The no-progress strike only fires inside the *verify-failed* path (line 340+), gated on `before == after`. So a zero-diff tick is safe here precisely **because** verify passes.

Let me restate the evidence before concluding. I ran the exact Verify command and it printed `VERIFY PASSED`. Separately I confirmed every invocation's flags against the real `--help`:

| Agent | Invocation | prompt consumed | trust flag | valid? |
|---|---|---|---|---|
| pi | `pi -p --no-session --approve "$(cat "$LOOP_PROMPT_FILE")"` | `$(cat …)` | `--approve` | ✅ |
| claude | `claude --print --no-session-persistence --dangerously-skip-permissions "$(cat …)"` | `$(cat …)` | `--dangerously-skip-permissions` | ✅ |
| codex | `codex exec --dangerously-bypass-approvals-and-sandbox --ephemeral "$(cat …)"` | positional `[PROMPT]` | `--dangerously-bypass-approvals-and-sandbox` | ✅ |
| opencode | `opencode run --auto "$(cat …)"` | positional `message` | `--auto` | ✅ |
| devin | `devin --print --prompt-file "$LOOP_PROMPT_FILE" --permission-mode dangerous` | `--prompt-file` | `--permission-mode dangerous` | ✅ |
| `Agent:` | `pi -p --no-session --approve --model glm-5.2 "$(cat …)"` | `$(cat …)` | `--approve` | ✅ |

The repository already satisfies the unit's Verify command, all three `Done means` items, and every constraint. I made no changes — none are needed, and forcing any would violate the "don't touch loop.sh/tests.run.sh" constraint and rule 9's narrow-diff scope. The runner will execute Verify and mark the unit done.

```text
Unit: README shows correct, runnable agent invocations
Changed: (none) — README.md and AGENTS.md already satisfy the unit
Verify expected: cd /home/daniel/build/knack && test $(grep -E '^LOOP_AGENT_CMD=' README.md | grep -E 'LOOP_PROMPT_FILE' | wc -l) -eq 5 && test $(grep -E '^Agent:' README.md | grep -E 'LOOP_PROMPT_FILE' | wc -l) -eq 1 && test $(grep -E 'LOOP_AGENT_CMD=' AGENTS.md | grep -vE 'LOOP_PROMPT_FILE' | wc -l) -eq 0 && bash -n loop.sh
Notes: Zero-diff tick by design — Verify passes (confirmed by running it), and all Done means items are met. Verified each agent's flags against the real CLI --help (pi, claude, codex, opencode, devin all valid + consume prompt + carry trust flag). No-progress detection does not fire on a passing verify (loop.sh marks done at verify_code==0 before reaching the no_progress branch), so this is safe. No blockers.
```
````

What this proves:
- The verify command passed for this work unit in the current repo state.

What remains unverified:
- Anything outside the verify command's proof scope.

## 2026-07-09T23:37:24-06:00 — loop.sh default pi and tests/run.sh cover agent invocation

Status: done

Unit:
````markdown
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

Status: pending````

Files changed:
```text
 M .loop/agent-agnostic/EVIDENCE.md
 M .loop/agent-agnostic/QUEUE.md
 M loop.sh
 M tests/run.sh
?? .loop/glossary/QUEUE.md
```

Verify command:
```bash
bash -n /home/daniel/build/knack/loop.sh && /home/daniel/build/knack/tests/run.sh
```

Verify output:
```text
Valid skill: /home/daniel/build/knack/.agents/skills/build
Valid skill: /home/daniel/build/knack/.agents/skills/decide
Valid skill: /home/daniel/build/knack/.agents/skills/domain-modeling
Valid skill: /home/daniel/build/knack/.agents/skills/explore
Valid skill: /home/daniel/build/knack/.agents/skills/fix
Valid skill: /home/daniel/build/knack/.agents/skills/plan
Valid skill: /home/daniel/build/knack/.agents/skills/review
knack tests passed
```

Worker output:
````text
All four `Done means:` criteria are met and the verify command passes. The diff is narrow and respects all constraints:

- **`loop.sh`**: default fallback now `pi -p --no-session --approve "..."` (`--approve` confirmed valid via `pi --help`).
- **`tests/run.sh`**: two new tests (fake `pi` on `PATH` via `env -u LOOP_AGENT_CMD PATH=...:...`; `LOOP_PROMPT_FILE` capture via `LOOP_AGENT_CMD`). Both derive from the existing `mktemp` `$tmp` and touch no repo files.
- Also fixed a latent bug in `assert_contains` (`grep -Fq` → `grep -Fq --`) — any pattern starting with `-` (like `--approve`) was being parsed as a grep option. Required for the new assertions; behavior unchanged for existing callers.
- `Agent:` override parsing untouched; `LOOP_AGENT_CMD`-unset path is exactly what test 1 exercises.

```text
Unit: loop.sh default pi and tests/run.sh cover agent invocation
Changed: loop.sh (default pi fallback +--approve); tests/run.sh (assert_contains grep -- hardening; two new tests: fake-pi-on-PATH, LOOP_PROMPT_FILE capture)
Verify expected: bash -n /home/daniel/build/knack/loop.sh && /home/daniel/build/knack/tests/run.sh
Notes: assert_contains needed `grep -Fq --` to accept `--approve`/`--no-session` patterns (was a latent bug, now fixed; no behavior change for existing callers).
```
````

What this proves:
- The verify command passed for this work unit in the current repo state.

What remains unverified:
- Anything outside the verify command's proof scope.
