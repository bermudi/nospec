# AGENTS.md

## Project

**knack** is an agent-agnostic harness for agentic development. It replaces litespec.

It is three separate artifacts with three separate concerns:
- **Skills** (`.agents/skills/`) — the workflow as procedural knowledge, agent-agnostic via agentskills.io
- **Loop** (`loop.sh`) — external bash script, agent-agnostic, owns the verify gate
- **CLI** (Go binary, `cli/`) — read-only validator + context provider. Packages the default skills via `go:embed`.

Code is the source of truth. Specs are disposable. Decisions and skills are durable.

## Thesis

Code is the source of truth. Specs rot. Plans are ephemeral coordination state. The reusable asset is procedural knowledge encoded as skills. The loop owns verification; the worker never self-certifies.

See `DESIGN.md` for the full design.

## Current state

The loop (`loop.sh`), all seven skills, and the Go CLI are built. The loop supports per-unit `Agent:` overrides, `LOOP_AGENT_CMD` for agent-agnosticism, and writes `.loop/HANDOFF.md` on non-clean exits. The skills cover the full explore → plan → build → review → fix flow plus two shared skills (decide, domain-modeling). The CLI (`cli/`) provides `skills init|check`, `validate`, `decisions list|show|check`, `status`, `glossary check`, and `instructions`. It embeds the default skills via `go:embed` and scaffolds them into target projects.

## Core artifacts

- `.loop/<name>/QUEUE.md` — disposable queue of verifiable work units.
- `.loop/<name>/EVIDENCE.md` — append-only ledger of what each tick proved.
- `.loop/<name>/HANDOFF.md` — cross-session handoff (written on pause/stop).
- `decisions/` — durable ADRs (architectural rulings, not current behavior).
- `glossary.md` — durable ubiquitous language.
- `.agents/skills/` — procedural knowledge (explore, plan, build, review, fix, decide, domain-modeling).
- `AGENTS.md` — operational context (this file).

## Skills

- **explore** — read codebase, grill intent, stress-test ideas. No artifacts except ADRs and glossary entries.
- **plan** — decompose intent into verifiable work units. Writes `.loop/<name>/QUEUE.md`. Optionally writes `.loop/<name>/specs/` for big work.
- **build** — implement one work unit. Don't self-certify. The loop owns the verify gate.
- **review** — two-axis adversarial review (standards + intent). Findings become new work units.
- **fix** — address review findings. Generates new work units, feeds back into the loop.
- **decide** (shared) — capture architectural rulings as ADRs in `decisions/`. Used by explore, plan, build, review.
- **domain-modeling** (shared) — manage `glossary.md`. Used by explore, plan, review.

## Working conventions

- Shell first for the loop. Go for the CLI. Markdown for skills.
- Plain Markdown files over stores or schemas.
- The worker never self-certifies. The runner owns verification.
- A `Verify` command must be deterministic and executable by the runner — not an LLM-as-judge. The loop's backpressure is mechanical.
- The runner enforces hard stops (max ticks, no-progress detection). The agent does one work unit per tick and reports what remains if it can't finish.
- Review is opt-in. The loop does not run review automatically; the user invokes it when they want adversarial scrutiny.
- When the queue is complete and verified, the cycle's `.loop/<name>/` subdirectory is disposable. The human deletes it; the loop does not.
- A work unit must leave the repo better if the loop stops immediately after it.
- Work units are whatever shape the work is — not forced into "vertical slices."
- `LOOP_AGENT_CMD` overrides the worker invocation (agent-agnostic: pi, claude, codex, opencode, etc.).
- Per-unit `Agent:` field in QUEUE.md overrides `LOOP_AGENT_CMD` for one unit (model routing).
- Work units are `## <outcome>` headers — no "Slice" prefix, no numbering. Vertical slice is one type of work unit, not the required format.
- Specs are disposable. Decisions are durable. Code is the source of truth.

## Verification

After meaningful changes, run:

```bash
./tests/run.sh
```

For CLI-only work, also run the Go tests directly:

```bash
cd cli && go test ./...
```

## Lessons learned

- **The loop works end-to-end with Devin as the worker.** `LOOP_AGENT_CMD='devin --print --prompt-file "$LOOP_PROMPT_FILE" --model kimi-k2.7 --permission-mode dangerous'` drove all 5 CLI units to verified completion in one run. The `--permission-mode dangerous` flag is required for the worker to actually write code and run `go test` without hanging on approval prompts.
- **The worker prompt must name the skill explicitly.** `prompts/worker.md` now tells the worker to load the `build` skill. Trigger-based discovery is not reliable enough across agents; naming the skill is.
- **`go test ./...` as a verify command compounds.** Each unit's verify runs all prior units' tests too, so regressions across units are caught at the next unit's gate. This makes the verify gate stronger as the queue progresses.
- **Review catches what verify can't.** The queue parser regex `^##\s*(.*)$` matched `###` subheadings as work units — a real bug that diverged from `loop.sh`'s behavior. `go test` passed because no fixture used `###`. Adversarial review against the actual codebase (comparing to `loop.sh`'s parser) found it. The fix: exclude `###` lines explicitly in `isUnitHeader`.
- **Embedded skills must stay in sync with `.agents/skills/`.** `cli/sync-skills.sh` re-copies from `../.agents/skills/`. Run it after editing skills. `diff -r .agents/skills cli/embedded/skills` verifies sync.
- **`go vet` doesn't catch unused test helpers.** `fileExists` in `queue_test.go` was dead code that `go vet` missed. Review caught it.
