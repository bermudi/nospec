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

The loop (`loop.sh`), all seven skills, and the Go CLI are built. The loop supports per-unit `Agent:` overrides, `LOOP_AGENT_CMD` for agent-agnosticism, opt-in review/fix orchestration via `--review`, and `.loop/<name>/HANDOFF.md` on non-clean exits. The skills cover the full explore → plan → build → review → fix flow plus two shared skills (decide, domain-modeling). The CLI (`cli/`) provides `skills init|check`, `validate`, `decisions list|show|check`, `status`, `glossary check`, and `instructions`. It embeds the default skills via `go:embed` and scaffolds them into target projects.

## Core artifacts

- `.loop/<name>/QUEUE.md` — disposable queue of verifiable work units.
- `.loop/<name>/EVIDENCE.md` — append-only ledger of what each tick proved. Durable: keep it after deleting QUEUE.md so `decisions check` can still see which ADRs the completed cycle referenced.
- `.loop/<name>/HANDOFF.md` — cross-session handoff (written on pause/stop).
- `.loop/<name>/REVIEW.md` — structured review artifact written when `--review` is enabled. The loop reads only the actionable count.
- `decisions/` — durable ADRs (architectural rulings, not current behavior).
- `glossary.md` — durable ubiquitous language.
- `LEARNINGS.md` — durable domain and problem insights (not operational, not rulings).
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
- Review is opt-in. The loop runs review/fix only when invoked with `--review`; otherwise it runs build ticks only.
- With `--review`, the loop invokes review after pending build units drain, reads only `REVIEW.md`'s actionable count, invokes fix when actionable findings exist, then runs another build pass for appended units. Judgment stays in the `review` and `fix` skills.
- When the queue is complete and verified, the cycle's `.loop/<name>/QUEUE.md` is disposable. The human deletes it; the loop does not. `EVIDENCE.md` is the durable ledger — keep it so `decisions check` can still trace which ADRs the cycle referenced.
- A work unit must leave the repo better if the loop stops immediately after it.
- Work units are whatever shape the work is — not forced into "vertical slices."
- `LOOP_AGENT_CMD` overrides the worker invocation (agent-agnostic: pi, claude, codex, opencode, etc.).
- `LOOP_REVIEW_CMD` and `LOOP_FIX_CMD` override review/fix phase invocations when `--review` is enabled. They default to `LOOP_AGENT_CMD`.
- Per-unit `Agent:` field in QUEUE.md overrides `LOOP_AGENT_CMD` for one unit (model routing).
- Work units are `## <outcome>` headers with `Read first:`, `Constraints:`, `Done means:`, and `Verify:` fields. `Done means:` is the acceptance criteria; `Verify:` is the mechanically enforceable subset. The gap between them is the review surface.
- `Read first:` is context, not scope — 2–4 entries of ADRs, code areas, or rulings.
- `Constraints:` state boundaries. A constraint says what must stay true or what is out of bounds, never what to edit. If it names a file, it is "don't touch X" or "X's public API must not change," not "update X."
- Work units are `## <outcome>` headers — no "Slice" prefix, no numbering. Vertical slice is one type of work unit, not the required format.
- Specs are disposable. Decisions are durable. Code is the source of truth.
- Operational gotchas go in `AGENTS.md`; domain/problem insights go in `LEARNINGS.md`.

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
- **The worker prompt names the skill explicitly (ADR-0007).** `prompts/worker.md` tells the worker to load the `build` skill by name and path; the loop passes the prompt via `LOOP_PROMPT_FILE`. Trigger-based discovery is not reliable enough across agents, so the loop does not rely on it.
- **`go test ./...` as a verify command compounds.** Each unit's verify runs all prior units' tests too, so regressions across units are caught at the next unit's gate. This makes the verify gate stronger as the queue progresses.
- **Review catches what verify can't.** The queue parser regex `^##\s*(.*)$` matched `###` subheadings as work units — a real bug that diverged from `loop.sh`'s behavior. `go test` passed because no fixture used `###`. Adversarial review against the actual codebase (comparing to `loop.sh`'s parser) found it. The fix: exclude `###` lines explicitly in `isUnitHeader`.
- **Embedded skills must stay in sync with `.agents/skills/`.** `cli/sync-skills.sh` re-copies from `../.agents/skills/`. Run it after editing skills. `diff -r .agents/skills cli/embedded/skills` verifies sync.
- **`go vet` doesn't catch unused test helpers.** `fileExists` in `queue_test.go` was dead code that `go vet` missed. Review caught it.
- **Verify commands must be path-correct.** Unit 1 of the named-cycles queue had `cd cli && go test ./... && ./tests/run.sh` — but `./tests/run.sh` ran from `cli/` after the `cd`, not from the repo root. The worker did the work correctly; the verify command was wrong. The loop correctly caught the failure (mechanical gate working), but it was a false negative. Always test verify commands manually before writing them into a queue.
- **Workers scope to the outcome plus constraints, not to a file list.** ADR-0005 replaced `Work:` with `Read first:` and `Constraints:`. The unit's scope is its outcome plus its constraints — the worker determines which files to change. The old lesson ("name every file in the work unit") is wrong under the new shape: naming files in constraints smuggles scope the same way `Work:` did. The first plan-shape cycle proved this — the constraint said "no `Work:` refs in skills, prompts, DESIGN.md, or AGENTS.md" and the worker touched exactly those files, leaving 9 other files (test fixtures, examples, README) with stale `Work:` fields. Prefer outcome-level constraints ("no artifact that teaches the format may reference `Work:`") over file-enumerated constraints.
- **`decisions check` orphaned-ADR semantics resolved.** An ADR is orphaned if it is not referenced by any QUEUE.md (current work) or any EVIDENCE.md (completed work). The loop now writes the full unit body into EVIDENCE.md so ADR references survive QUEUE.md deletion. EVIDENCE.md is the durable ledger; QUEUE.md is disposable. An ADR that drove a completed cycle is not orphaned as long as its EVIDENCE.md ledger exists.
- **Named work cycles enable concurrent work.** ADR-0004 gave each work cycle its own subdirectory under `.loop/`. The loop already supported this via the queue path argument — only convention (skills, docs) and `knack status` needed to change. Running `./loop.sh run .loop/<name>/QUEUE.md` is fully independent of other cycles.
- **Negative-grep verifies must anchor on field syntax, not bare mentions.** The plan-shape cycle's verify used `! grep -rn 'Work:' ...` to prove the old field was gone. The same commit that landed the work also added this lessons-learned entry, which says "ADR-0005 replaced `Work:`..." — so the verify rotted the moment history got documented. Anchor negative greps to the field shape (`^Work:`) or to a syntax-specific pattern, not to any mention of the word. Otherwise the verify forbids the project from ever recording why the change was made.
