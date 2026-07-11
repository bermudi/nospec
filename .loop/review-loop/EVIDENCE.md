
## 2026-07-10T17:31:46-06:00 — ADR-0008 captures the loop-orchestrated review-fix decision

Status: done

Unit:
````markdown
## ADR-0008 captures the loop-orchestrated review-fix decision

Read first:
- `.loop/review-loop/specs/proposal.md`
- `DESIGN.md` (open questions and "What the loop does NOT do")
- `.agents/skills/review/SKILL.md`
- `.agents/skills/fix/SKILL.md`

Constraints:
- The ADR must state the loop's role (orchestrator) vs the review/fix skills' role (judgment).
- Do not change `loop.sh`, `review/SKILL.md`, or `fix/SKILL.md` yet.
- Keep the ADR to one page.

Done means:
- `decisions/0008-loop-orchestrates-review-fix.md` exists with `Status: accepted`.
- `knack decisions check` still passes (the new ADR is not orphaned — it is referenced by this queue).

Verify:
```bash
cd /home/daniel/build/knack && test -f decisions/0008-loop-orchestrates-review-fix.md && grep -q '^Status: accepted' decisions/0008-loop-orchestrates-review-fix.md && /tmp/knack decisions check
```

Status: pending````

Files changed:
```text
 M .loop/review-loop/QUEUE.md
?? .loop/review-loop/EVIDENCE.md
?? decisions/0008-loop-orchestrates-review-fix.md
```

Verify command:
```bash
cd /home/daniel/build/knack && test -f decisions/0008-loop-orchestrates-review-fix.md && grep -q '^Status: accepted' decisions/0008-loop-orchestrates-review-fix.md && /tmp/knack decisions check
```

Verify output:
```text
```

Worker output:
````text
Done. The `QUEUE.md` modification (1 line) is the runner's `pending → in_progress` status flip from when it claimed this unit — not my edit. My only change is the new ADR file.

The unit's `Verify` command passes: file exists, `Status: accepted` present, `decisions check` green (the ADR is referenced by five `decisions/0008-...` path citations already in the QUEUE.md unit bodies, so it is not orphaned; no `ADR-0008` token sits in a unit body to dangle).

```text
Unit: ADR-0008 captures the loop-orchestrated review-fix decision
Changed: decisions/0008-loop-orchestrates-review-fix.md (new, 41 lines — one page)
Verify expected: cd /home/daniel/build/knack && test -f decisions/0008-loop-orchestrates-review-fix.md && grep -q '^Status: accepted' decisions/0008-loop-orchestrates-review-fix.md && /tmp/knack decisions check
Notes: ADR follows the decide-skill format (Context/Decision/Consequences), modeled on ADR-0007. The Decision section states the boundary as a ruling — loop owns orchestration + stop conditions (mechanical: invoke review/fix, read REVIEW.md actionable count, enforce hard stops); review/fix skills own judgment (semantic: two-axis review, triage, finding phrasing). Did not touch loop.sh, review/SKILL.md, or fix/SKILL.md per constraints. Confirmed non-orphaned by inspecting adrReferenced() candidates (matches the decisions/0008-...md path already cited in QUEUE.md units 1–5).
```
````

What this proves:
- The verify command passed for this work unit in the current repo state.

What remains unverified:
- Anything outside the verify command's proof scope.

## 2026-07-10T19:47:31-06:00 — review skill writes a structured REVIEW.md

Status: blocked

Unit:
````markdown
## review skill writes a structured REVIEW.md

Read first:
- `.agents/skills/review/SKILL.md`
- `.agents/skills/fix/SKILL.md`
- `decisions/0008-loop-orchestrates-review-fix.md`

Constraints:
- Preserve the two-axis (standards + intent) review structure.
- The output format must be machine-parseable enough for `loop.sh` to detect the `actionable` count.
- Do not change `loop.sh` or `fix/SKILL.md` yet.
- Keep the skill file scannable and sync it to `cli/embedded/skills/review/` after editing.

Done means:
- `review/SKILL.md` documents `REVIEW.md` as its output artifact.
- `REVIEW.md` has sections for `## Standards`, `## Intent`, `## Speculative`, and a `## Summary` with counts.
- Each finding has a confidence level (high/medium/low) and evidence (`path/to/file:line` or quoted code).
- `cli/embedded/skills/review/SKILL.md` is in sync with `.agents/skills/review/SKILL.md`.

Verify:
```bash
cd /home/daniel/build/knack && diff -r .agents/skills/review cli/embedded/skills/review && ./tests/run.sh
```

Status: pending````

Files changed:
```text
 M .loop/review-loop/EVIDENCE.md
 M .loop/review-loop/QUEUE.md
```

Verify command:
```bash
cd /home/daniel/build/knack && diff -r .agents/skills/review cli/embedded/skills/review && ./tests/run.sh
```

Verify output:
```text
```

Worker output:
````text
OpenAI Codex v0.144.1
--------
workdir: /home/daniel/build/knack
model: gpt-5.6-luna
provider: openai
approval: never
sandbox: danger-full-access
reasoning effort: high
reasoning summaries: none
session id: 019f4edb-6446-7790-b9b3-a431124f0945
--------
user
# Knack Worker

You are one worker tick in a bounded loop. Complete exactly one work unit, then stop.

Load and follow the **build** skill in `.agents/skills/build/` before doing any work.

## Rules

1. Read `AGENTS.md` first if it exists — it contains operational context.
2. Read the current work unit carefully, especially its `Verify` command.
3. Do only the work needed for that unit.
4. Do not start another work unit.
5. Do not mark the unit complete yourself.
6. Do not edit `.loop/<name>/EVIDENCE.md`; the runner writes evidence after verification.
7. If the unit is blocked, make the smallest useful note in your final response and stop.
8. Stay within the runner's hard stops (max ticks, no-progress detection). If the unit is too large for one tick, do as much as keeps the repo working and report what remains.
9. The worker's scope is the unit's outcome plus its constraints. Keep the diff narrow and aligned with that scope.
10. If verification fails while you are working, fix the cause if it belongs to this unit; otherwise stop and report the blocker.

## Success standard

Your job is not to claim success. Your job is to make the repository state satisfy the unit's `Verify` command. The runner will execute that command after you exit.

## Output

End with a compact terminal handoff. This is not `.loop/<name>/HANDOFF.md` — the runner writes that file. This is your own summary to the runner:

```text
Unit: <title>
Changed: <brief file/area list>
Verify expected: <command from unit>
Notes: <blockers or caveats, if any>
```

Current work unit from /home/daniel/build/knack/.loop/review-loop/QUEUE.md:

## review skill writes a structured REVIEW.md

Read first:
- `.agents/skills/review/SKILL.md`
- `.agents/skills/fix/SKILL.md`
- `decisions/0008-loop-orchestrates-review-fix.md`

Constraints:
- Preserve the two-axis (standards + intent) review structure.
- The output format must be machine-parseable enough for `loop.sh` to detect the `actionable` count.
- Do not change `loop.sh` or `fix/SKILL.md` yet.
- Keep the skill file scannable and sync it to `cli/embedded/skills/review/` after editing.

Done means:
- `review/SKILL.md` documents `REVIEW.md` as its output artifact.
- `REVIEW.md` has sections for `## Standards`, `## Intent`, `## Speculative`, and a `## Summary` with counts.
- Each finding has a confidence level (high/medium/low) and evidence (`path/to/file:line` or quoted code).
- `cli/embedded/skills/review/SKILL.md` is in sync with `.agents/skills/review/SKILL.md`.

Verify:
```bash
cd /home/daniel/build/knack && diff -r .agents/skills/review cli/embedded/skills/review && ./tests/run.sh
```

Status: pending
hook: SessionStart
hook: SessionStart Completed
ERROR: Reconnecting... 2/5
ERROR: Reconnecting... 3/5
ERROR: Reconnecting... 4/5
ERROR: Reconnecting... 5/5
warning: Falling back from WebSockets to HTTPS transport. unexpected status 404 Not Found: Model not found gpt-5.6-luna-free-1p-codexswic-ev3
ERROR: Reconnecting... 1/5
ERROR: Reconnecting... 2/5
ERROR: Reconnecting... 3/5
ERROR: Reconnecting... 4/5
ERROR: Reconnecting... 5/5
ERROR: unexpected status 404 Not Found: Model not found gpt-5.6-luna, url: https://chatgpt.com/backend-api/codex/responses, cf-ray: a19423fd2ae67a65-QRO, request id: 9d2463b8-0b87-4180-956e-42cd7b826ddd
ERROR: unexpected status 404 Not Found: Model not found gpt-5.6-luna, url: https://chatgpt.com/backend-api/codex/responses, cf-ray: a19423fd2ae67a65-QRO, request id: 9d2463b8-0b87-4180-956e-42cd7b826ddd
````

What this proves:
- The work unit is not externally verified.

What remains unverified:
- Anything outside the verify command's proof scope.

## 2026-07-10T19:50:35-06:00 — fix skill consumes REVIEW.md and appends fix units to QUEUE.md

Status: done

Unit:
````markdown
## fix skill consumes REVIEW.md and appends fix units to QUEUE.md

Read first:
- `.agents/skills/fix/SKILL.md`
- `.agents/skills/review/SKILL.md` after Unit 2
- `decisions/0008-loop-orchestrates-review-fix.md`

Constraints:
- The fix skill must not edit `REVIEW.md`.
- It must append new work units to the existing `QUEUE.md` without changing the status of existing units.
- It must follow the existing work unit format (outcome, Read first, Constraints, Done means, Verify, Status: pending).
- Keep the skill file scannable and sync it to `cli/embedded/skills/fix/` after editing.

Done means:
- `fix/SKILL.md` documents how to read `REVIEW.md` and triage findings into actionable units.
- The skill appends `Status: pending` units to `QUEUE.md`.
- `cli/embedded/skills/fix/SKILL.md` is in sync with `.agents/skills/fix/SKILL.md`.

Verify:
```bash
cd /home/daniel/build/knack && diff -r .agents/skills/fix cli/embedded/skills/fix && ./tests/run.sh
```

Status: pending````

Files changed:
```text
 M .agents/skills/fix/SKILL.md
 M .loop/review-loop/EVIDENCE.md
 M .loop/review-loop/HANDOFF.md
 M .loop/review-loop/QUEUE.md
 M cli/embedded/skills/fix/SKILL.md
```

Verify command:
```bash
cd /home/daniel/build/knack && diff -r .agents/skills/fix cli/embedded/skills/fix && ./tests/run.sh
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
OpenAI Codex v0.144.1
--------
workdir: /home/daniel/build/knack
model: gpt-5.5
provider: openai
approval: never
sandbox: danger-full-access
reasoning effort: high
reasoning summaries: none
session id: 019f4edc-d771-7682-8de9-9bfd0f42dadd
--------
user
# Knack Worker

You are one worker tick in a bounded loop. Complete exactly one work unit, then stop.

Load and follow the **build** skill in `.agents/skills/build/` before doing any work.

## Rules

1. Read `AGENTS.md` first if it exists — it contains operational context.
2. Read the current work unit carefully, especially its `Verify` command.
3. Do only the work needed for that unit.
4. Do not start another work unit.
5. Do not mark the unit complete yourself.
6. Do not edit `.loop/<name>/EVIDENCE.md`; the runner writes evidence after verification.
7. If the unit is blocked, make the smallest useful note in your final response and stop.
8. Stay within the runner's hard stops (max ticks, no-progress detection). If the unit is too large for one tick, do as much as keeps the repo working and report what remains.
9. The worker's scope is the unit's outcome plus its constraints. Keep the diff narrow and aligned with that scope.
10. If verification fails while you are working, fix the cause if it belongs to this unit; otherwise stop and report the blocker.

## Success standard

Your job is not to claim success. Your job is to make the repository state satisfy the unit's `Verify` command. The runner will execute that command after you exit.

## Output

End with a compact terminal handoff. This is not `.loop/<name>/HANDOFF.md` — the runner writes that file. This is your own summary to the runner:

```text
Unit: <title>
Changed: <brief file/area list>
Verify expected: <command from unit>
Notes: <blockers or caveats, if any>
```

Current work unit from /home/daniel/build/knack/.loop/review-loop/QUEUE.md:

## fix skill consumes REVIEW.md and appends fix units to QUEUE.md

Read first:
- `.agents/skills/fix/SKILL.md`
- `.agents/skills/review/SKILL.md` after Unit 2
- `decisions/0008-loop-orchestrates-review-fix.md`

Constraints:
- The fix skill must not edit `REVIEW.md`.
- It must append new work units to the existing `QUEUE.md` without changing the status of existing units.
- It must follow the existing work unit format (outcome, Read first, Constraints, Done means, Verify, Status: pending).
- Keep the skill file scannable and sync it to `cli/embedded/skills/fix/` after editing.

Done means:
- `fix/SKILL.md` documents how to read `REVIEW.md` and triage findings into actionable units.
- The skill appends `Status: pending` units to `QUEUE.md`.
- `cli/embedded/skills/fix/SKILL.md` is in sync with `.agents/skills/fix/SKILL.md`.

Verify:
```bash
cd /home/daniel/build/knack && diff -r .agents/skills/fix cli/embedded/skills/fix && ./tests/run.sh
```

Status: pending
hook: SessionStart
hook: SessionStart Completed
codex
I’m using the `build` skill because this is a bounded loop worker tick. I’ll read the project context and the skill instructions first, then make only the unit-scoped edits.
exec
/usr/bin/bash -lc "pwd && sed -n '1,260p' AGENTS.md" in /home/daniel/build/knack
 succeeded in 0ms:
/home/daniel/build/knack
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

The loop (`loop.sh`), all seven skills, and the Go CLI are built. The loop supports per-unit `Agent:` overrides, `LOOP_AGENT_CMD` for agent-agnosticism, and writes `.loop/<name>/HANDOFF.md` on non-clean exits. The skills cover the full explore → plan → build → review → fix flow plus two shared skills (decide, domain-modeling). The CLI (`cli/`) provides `skills init|check`, `validate`, `decisions list|show|check`, `status`, `glossary check`, and `instructions`. It embeds the default skills via `go:embed` and scaffolds them into target projects.

## Core artifacts

- `.loop/<name>/QUEUE.md` — disposable queue of verifiable work units.
- `.loop/<name>/EVIDENCE.md` — append-only ledger of what each tick proved. Durable: keep it after deleting QUEUE.md so `decisions check` can still see which ADRs the completed cycle referenced.
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
- When the queue is complete and verified, the cycle's `.loop/<name>/QUEUE.md` is disposable. The human deletes it; the loop does not. `EVIDENCE.md` is the durable ledger — keep it so `decisions check` can still trace which ADRs the cycle referenced.
- A work unit must leave the repo better if the loop stops immediately after it.
- Work units are whatever shape the work is — not forced into "vertical slices."
- `LOOP_AGENT_CMD` overrides the worker invocation (agent-agnostic: pi, claude, codex, opencode, etc.).
- Per-unit `Agent:` field in QUEUE.md overrides `LOOP_AGENT_CMD` for one unit (model routing).
- Work units are `## <outcome>` headers with `Read first:`, `Constraints:`, `Done means:`, and `Verify:` fields. `Done means:` is the acceptance criteria; `Verify:` is the mechanically enforceable subset. The gap between them is the review surface.
- `Read first:` is context, not scope — 2–4 entries of ADRs, code areas, or rulings.
- `Constraints:` state boundaries. A constraint says what must stay true or what is out of bounds, never what to edit. If it names a file, it is "don't touch X" or "X's public API must not change," not "update X."
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

exec
/usr/bin/bash -lc "sed -n '1,260p' .agents/skills/build/SKILL.md" in /home/daniel/build/knack
 succeeded in 0ms:
---
name: build
description: Use when implementing one work unit from a `.loop/<name>/QUEUE.md` queue. The worker-side skill — read the unit, do the work, don't self-certify, end with a handoff. Triggers on "build", "implement", "apply this unit", "do the work", "run the loop", or when a work unit needs to be executed. Also use when the loop invokes the worker for a tick.
---

# Build

Implement one work unit from `.loop/<name>/QUEUE.md`. Do the work. Don't self-certify. The loop runner owns the verify gate — your job is to make the repository state satisfy the unit's `Verify` command, not to claim success.

## Core rules

The runner injects `prompts/worker.md` at the start of every tick. If it is not in your context, read it now. Its `Rules` and `Output` sections are the canonical source for this skill; the sections below elaborate on decisions, operational learnings, blockers, and units that are too large for one tick.

> **Scope note:** Updating `AGENTS.md` or writing an ADR during a tick is a durable trace, not part of the unit's `Verify` scope. Do it only when the tick teaches you something that would save the next session time.

## Scope

The unit's scope is its outcome plus its constraints. The worker determines which files to change and how. The `Verify:` command is the mechanically enforceable subset of `Done means:`.

## Capturing decisions during build

If you discover an architectural ruling while implementing — "we need to handle X this way because Y" — capture it as an ADR using the `decide` skill. Do this inline, not after the unit. Decisions made during implementation are the most valuable kind because they come from the code pushing back.

## Capturing operational learnings

If you learn something about how the project works — a build command, a test convention, a gotcha — that isn't already in `AGENTS.md`, add it. `AGENTS.md` is the living operational context. If it grows stale, the next agent session starts with worse information.

Don't add trivia. Add things that would have saved you time if you'd known them upfront.

## When the unit is too big

If the work unit is larger than what can be done in one tick:

- Do as much as you can while keeping the repo in a working state.
- End with a handoff note explaining what remains.
- The runner will re-queue the unit if verify fails.

Do not try to do everything in one tick if the work genuinely needs more. The loop is designed for multiple ticks.

## When you're blocked

If you hit a blocker you can't resolve within the unit's scope:

1. State the blocker clearly in your final output.
2. Note what would unblock you (a decision, a dependency, a missing file).
3. Stop. Don't thrash.

The runner will mark the unit as `blocked` and write a handoff. The next session can pick up from there.

## Output

The output format is defined in the worker prompt (`prompts/worker.md`). End with a compact terminal handoff there.

codex
The build instructions point back to `prompts/worker.md` as canonical, so I’m checking that plus the unit’s read-first files. Then I’ll update the fix skill and embedded copy.
exec
/usr/bin/bash -lc "sed -n '1,260p' .agents/skills/fix/SKILL.md" in /home/daniel/build/knack
exec
/usr/bin/bash -lc "sed -n '1,220p' prompts/worker.md" in /home/daniel/build/knack
 succeeded in 0ms:
# Knack Worker

You are one worker tick in a bounded loop. Complete exactly one work unit, then stop.

Load and follow the **build** skill in `.agents/skills/build/` before doing any work.

## Rules

1. Read `AGENTS.md` first if it exists — it contains operational context.
2. Read the current work unit carefully, especially its `Verify` command.
3. Do only the work needed for that unit.
4. Do not start another work unit.
5. Do not mark the unit complete yourself.
6. Do not edit `.loop/<name>/EVIDENCE.md`; the runner writes evidence after verification.
7. If the unit is blocked, make the smallest useful note in your final response and stop.
8. Stay within the runner's hard stops (max ticks, no-progress detection). If the unit is too large for one tick, do as much as keeps the repo working and report what remains.
9. The worker's scope is the unit's outcome plus its constraints. Keep the diff narrow and aligned with that scope.
10. If verification fails while you are working, fix the cause if it belongs to this unit; otherwise stop and report the blocker.

## Success standard

Your job is not to claim success. Your job is to make the repository state satisfy the unit's `Verify` command. The runner will execute that command after you exit.

## Output

End with a compact terminal handoff. This is not `.loop/<name>/HANDOFF.md` — the runner writes that file. This is your own summary to the runner:

```text
Unit: <title>
Changed: <brief file/area list>
Verify expected: <command from unit>
Notes: <blockers or caveats, if any>
```

 succeeded in 0ms:
---
name: fix
description: Use when addressing review findings and feeding them back into the loop. Converts review findings into new verifiable work units in `.loop/<name>/QUEUE.md` and runs another loop pass. Triggers on "fix the review findings", "address the feedback", "rework based on review", "fix what review found", or when review produced findings that need to be resolved.
---

# Fix

Address review findings by triaging them and generating new work units that feed back into the loop. The `fix` skill is the bridge between review and another build pass; it owns the triage and the queue format. It appends to the same `.loop/<name>/QUEUE.md` file as the unit being fixed.

## Procedure

1. **Read the review findings.** Each finding from the `review` skill is a candidate work unit. `review` has already classified findings as trivial / actionable / disputed / deferred; use those classifications or re-triage as needed.

2. **Triage.** Not every finding warrants a work unit:
   - **Actionable** — the finding identifies a real issue with a clear fix. → Create a work unit.
   - **Trivial** — one-line fix, no risk. → Fix it now, don't create a unit.
   - **Disputed** — the finding is wrong or the reviewer is being overly cautious. → Note the disagreement, don't create a unit. Move on.
   - **Deferred** — the finding is valid but not worth fixing now. → Note it in `AGENTS.md` or a backlog, don't create a unit.

3. **Create work units** for actionable findings. Use the standard work unit format from the `plan` skill (single-unit excerpt below):

````markdown
## <fix for the finding — observable outcome>

Why:
<reference to the review finding — which axis, what was found>

Read first:
- <context the worker needs: file, ADR, or area>
- <2–4 entries; context, not scope>

Constraints:
- <boundary>
- <what must stay true or out of bounds; if a file is named, it is "don't touch X" or "X's public API must not change", not "update X">

Done means:
- <the finding is resolved>
- <no new issue introduced>

Verify:
```bash
<deterministic command that proves the fix>
```

Status: pending
````

4. **Append to QUEUE.md.** Read the existing `.loop/<name>/QUEUE.md` first, then append the new units. If the queue was completed (all units `done`), the new units extend it. If the queue still has pending units, the new units are added after them. Preserve the existing structure and status.

5. **Run the loop.** After writing the units, run:

```bash
./loop.sh run .loop/<name>/QUEUE.md
```

The loop will pick up the first pending unit and proceed.

## What fix is not

- Not a re-review — the findings are already known. Fix turns them into work.
- Not a debate — if a finding is disputed, note it and move on. Don't argue with the review in the queue.
- Not a rewrite — each fix unit should be narrow. If a finding requires a broad rewrite, that's a new explore → plan cycle, not a fix.

## Guardrail

Fix units must not break what review already approved. Each fix unit's `Done means:` should include a no-regression condition. The verify command should cover both the fix and the preservation of what was already working.

## Output

- New work units appended to `.loop/<name>/QUEUE.md`
- A summary of what was triaged (actionable / trivial / disputed / deferred)
- Suggestion to run the loop

exec
/usr/bin/bash -lc "sed -n '1,320p' .agents/skills/review/SKILL.md" in /home/daniel/build/knack
exec
/usr/bin/bash -lc "sed -n '1,220p' decisions/0008-loop-orchestrates-review-fix.md" in /home/daniel/build/knack
 succeeded in 0ms:
# 0008: The loop orchestrates the review-fix subloop; skills keep the judgment

Date: 2026-07-10
Status: accepted

## Context

The default skill flow is `explore → plan → build → review → fix`, but `loop.sh` only runs `build` ticks. Review and fix are left as manual skills, so the loop cannot autonomously review a completed queue, turn findings into fix units, re-run the build pass, and stop when review is clean. That is the missing piece for a bounded, self-correcting development loop.

The architectural question is *who owns what*. DESIGN.md's "What the loop does NOT do" already states the loop does not run review and does not manage ADRs — those are skill responsibilities. Extending the loop into review/fix risks blurring that boundary: if the loop starts interpreting review content, it stops being the simple, agent-agnostic engine and becomes an LLM-as-judge.

Two roles must stay separate:

- **Orchestration** (the loop) — when to invoke review, when to invoke fix, when to stop. Mechanical and signal-driven.
- **Judgment** (the `review` and `fix` skills) — what counts as a finding, whether it is actionable, how to phrase a fix unit. Adversarial and semantic.

Alternatives considered:

- **LLM-as-judge inside the loop.** Rejected — the loop would interpret review content, breaking agent-agnosticism and the verify-gate principle (the aiming problem: the signal must be the actual property, not a proxy the loop can game).
- **The loop implements review logic.** Rejected — duplicates the `review` skill, drifts from the source of truth, and violates "the loop never reads skills."
- **Manual review/fix only (do nothing).** Rejected — leaves the loop unable to self-correct; the missing piece stays missing.

## Decision

The loop **orchestrates** an optional, bounded `build → review → fix` subloop. It invokes the `review` and `fix` workers and interprets their *structured* outputs as continue/stop signals. It does not implement review or fix logic.

When `--review` is set and the build queue has drained, the loop runs a review worker that writes a structured `REVIEW.md`. The loop reads only the `actionable` count from that file. If it is non-zero, the loop runs a fix worker that appends `Status: pending` units to `QUEUE.md`, then re-runs the build pass and reviews again. It stops when `actionable` is zero, a review-round limit (`--max-review-rounds`, default 2) is hit, the tick budget (`--max-ticks`) is exhausted, or a round generates no new units (no progress).

The boundary, stated as a ruling:

- **The loop owns orchestration and stop conditions** — invoking review/fix, reading the actionable count, enforcing hard stops. Mechanical.
- **The `review` and `fix` skills own judgment** — the two-axis standards/intent review, triage, finding phrasing, and work-unit generation. Semantic.

The loop never parses review content beyond the actionable count and never judges whether a finding is real. Review remains a skill the worker loads; the loop only knows to invoke it and read its summary.

## Consequences

- The loop gains a bounded self-correction capability without becoming an LLM-as-judge. Backpressure stays mechanical: the actionable count is the signal, the hard stops are the guardrails.
- A new contract surface: the loop depends on `REVIEW.md`'s actionable count being honest. A review worker that mislabels findings can loop forever or stop early — the round cap and the no-progress stop are the backstop, not the loop's own judgment.
- The `review` and `fix` skills gain a machine-readable input/output contract (`REVIEW.md`; appended units). Those skill changes land in their own work units; this ADR records only the architectural ruling and the orchestration boundary.
- Review stays opt-in (`--review`); default loop behavior is unchanged.

 succeeded in 0ms:
---
name: review
description: Use when reviewing completed work units or a finished queue. Two-axis adversarial review — standards (does the change follow codebase conventions?) and intent (does the change do what the work unit said?). Reviews against the actual codebase, not specs. Triggers on "review", "check this", "is this right", "what did we miss", "stress-test the implementation", or when work needs adversarial scrutiny before being accepted.
---

# Review

Adversarial review of completed work. Two axes, run independently so neither pollutes the other:

1. **Standards** — does the change follow the repo's coding conventions and codebase patterns?
2. **Intent** — does the change do what the work unit said it would?

Review against the **actual codebase**, not against specs that may have rotted. The code is the source of truth.

## When to review

- After a work unit is marked `done` and verify passed
- After a full queue is completed
- When the user asks for a sanity check
- Before accepting work as finished

Review is not a gate the loop enforces — it's a skill the user or agent invokes when they want adversarial scrutiny.

## Before you review

Read the work unit from `.loop/<name>/QUEUE.md` and the evidence from `.loop/<name>/EVIDENCE.md` for the unit you're reviewing. The evidence tells you what the verify command actually proved; the work unit tells you what was promised. Review against the actual codebase, not the specs.

## Two-axis review

### Axis 1: Standards

Does the change follow the codebase's existing patterns?

- Read `AGENTS.md` for stated conventions.
- Read neighboring code — does the change look like it belongs?
- Check error handling, naming, file layout, test style.
- Look for regressions — did the change break something nearby?
- Check for dead code, unused imports, leftover debugging.

The question is not "is this good code?" — that's subjective. The question is "does this match the codebase's own standards?"

### Axis 2: Intent

Does the change do what the work unit said it would?

- Read the work unit's `Done means:` and `Constraints:` fields.
- Read the actual diff.
- Does the diff satisfy the stated outcome?
- Does the diff stay within the stated constraints?
- Did the change introduce anything the unit didn't ask for?
- Did the change miss anything the unit said it would do?

The `Verify:` command is the mechanically enforceable subset of `Done means:`. The gap between `Done means:` and `Verify:` is the review surface: intent review checks what the verify command cannot.

> **Intent review is a judgment check, not a deterministic gate.** Be explicit about its scope and confidence. Any finding that becomes a new work unit must have a deterministic `Verify` command the runner can execute.

## Running the axes

Run both axes. They can be parallel (two passes over the same diff) or sequential. The order doesn't matter — what matters is that each axis is evaluated independently, without the other's conclusions bleeding in.

## Findings become input to the fix skill

Review findings are not just notes — they are the input to the `fix` skill, which triages them and appends actionable ones as new work units in `.loop/<name>/QUEUE.md`.

- **Trivial** findings (a typo, a missing newline) can be fixed inline during review.
- **Actionable** findings are handed to the `fix` skill. Do not write the work units yourself; `fix` owns the triage and formatting.
- **Disputed** or **deferred** findings are recorded in the review summary but not turned into units.

The output of review is a findings summary, not a queue edit.

## What review is not

- Not a lint pass — the verify gate already ran. Review is about what verify *can't* check.
- Not a spec compliance check — specs are disposable and may have rotted. Review against the codebase.
- Not a gate — the loop doesn't enforce review. It's invoked when the user wants it.

## Output

Summarize findings per axis:

- **Standards**: N findings (list them, or "no issues found")
- **Intent**: N findings (list them, or "the change matches the unit's stated outcome")

Classify each finding as trivial / actionable / disputed / deferred.

Then either:
- Hand actionable findings to the `fix` skill, which will triage and append work units to `.loop/<name>/QUEUE.md`, or
- Report "no action needed" if the work is clean.

exec
/usr/bin/bash -lc "sed -n '1,260p' .loop/review-loop/QUEUE.md" in /home/daniel/build/knack
 succeeded in 0ms:
# Loop Queue: review-loop

Goal:
Make `loop.sh` optionally orchestrate a bounded `build → review → fix` subloop. After the build queue drains, the loop runs a review worker, writes a structured `REVIEW.md`, and, if there are actionable findings, runs a fix worker that appends new work units to `QUEUE.md`. The loop re-runs build and review until no actionable findings remain, a review-round limit is hit, or no progress is detected.

Stop condition:
`bash -n /home/daniel/build/knack/loop.sh && /home/daniel/build/knack/tests/run.sh && cd /home/daniel/build/knack/cli && go test ./...`

## ADR-0008 captures the loop-orchestrated review-fix decision

Read first:
- `.loop/review-loop/specs/proposal.md`
- `DESIGN.md` (open questions and "What the loop does NOT do")
- `.agents/skills/review/SKILL.md`
- `.agents/skills/fix/SKILL.md`

Constraints:
- The ADR must state the loop's role (orchestrator) vs the review/fix skills' role (judgment).
- Do not change `loop.sh`, `review/SKILL.md`, or `fix/SKILL.md` yet.
- Keep the ADR to one page.

Done means:
- `decisions/0008-loop-orchestrates-review-fix.md` exists with `Status: accepted`.
- `knack decisions check` still passes (the new ADR is not orphaned — it is referenced by this queue).

Verify:
```bash
cd /home/daniel/build/knack && test -f decisions/0008-loop-orchestrates-review-fix.md && grep -q '^Status: accepted' decisions/0008-loop-orchestrates-review-fix.md && /tmp/knack decisions check
```

Status: done

## review skill writes a structured REVIEW.md

Read first:
- `.agents/skills/review/SKILL.md`
- `.agents/skills/fix/SKILL.md`
- `decisions/0008-loop-orchestrates-review-fix.md`

Constraints:
- Preserve the two-axis (standards + intent) review structure.
- The output format must be machine-parseable enough for `loop.sh` to detect the `actionable` count.
- Do not change `loop.sh` or `fix/SKILL.md` yet.
- Keep the skill file scannable and sync it to `cli/embedded/skills/review/` after editing.

Done means:
- `review/SKILL.md` documents `REVIEW.md` as its output artifact.
- `REVIEW.md` has sections for `## Standards`, `## Intent`, `## Speculative`, and a `## Summary` with counts.
- Each finding has a confidence level (high/medium/low) and evidence (`path/to/file:line` or quoted code).
- `cli/embedded/skills/review/SKILL.md` is in sync with `.agents/skills/review/SKILL.md`.

Verify:
```bash
cd /home/daniel/build/knack && diff -r .agents/skills/review cli/embedded/skills/review && ./tests/run.sh
```

Status: blocked

## fix skill consumes REVIEW.md and appends fix units to QUEUE.md

Read first:
- `.agents/skills/fix/SKILL.md`
- `.agents/skills/review/SKILL.md` after Unit 2
- `decisions/0008-loop-orchestrates-review-fix.md`

Constraints:
- The fix skill must not edit `REVIEW.md`.
- It must append new work units to the existing `QUEUE.md` without changing the status of existing units.
- It must follow the existing work unit format (outcome, Read first, Constraints, Done means, Verify, Status: pending).
- Keep the skill file scannable and sync it to `cli/embedded/skills/fix/` after editing.

Done means:
- `fix/SKILL.md` documents how to read `REVIEW.md` and triage findings into actionable units.
- The skill appends `Status: pending` units to `QUEUE.md`.
- `cli/embedded/skills/fix/SKILL.md` is in sync with `.agents/skills/fix/SKILL.md`.

Verify:
```bash
cd /home/daniel/build/knack && diff -r .agents/skills/fix cli/embedded/skills/fix && ./tests/run.sh
```

Status: in_progress

## loop.sh orchestrates bounded build-review-fix rounds

Read first:
- `loop.sh`
- `decisions/0008-loop-orchestrates-review-fix.md`
- `.loop/review-loop/specs/proposal.md`
- `tests/run.sh`

Constraints:
- Default behavior must be unchanged when `--review` is not set.
- `--max-ticks` remains the total build-tick budget across all rounds.
- Add `--max-review-rounds` (default 2) and `--review` flag.
- Support `LOOP_REVIEW_CMD` and `LOOP_FIX_CMD` env vars with fallback to `LOOP_AGENT_CMD`.
- Do not break the `Agent:` per-unit override for build ticks.
- Preserve `HANDOFF.md` and `EVIDENCE.md` behavior.
- New tests must not require a real model or API key.

Done means:
- `loop.sh --review` runs the build queue, then runs review, then fix if actionable, then loops.
- `loop.sh` stops when `REVIEW.md` reports `actionable: 0`, `--max-review-rounds` is reached, `--max-ticks` is reached, or no new units are generated.
- `tests/run.sh` has a test that uses fake workers to run a full `build → review → fix → review` cycle.
- `./tests/run.sh` and `bash -n loop.sh` pass.

Verify:
```bash
bash -n /home/daniel/build/knack/loop.sh && /home/daniel/build/knack/tests/run.sh
```

Status: pending

## docs and AGENTS.md reflect the new loop behavior

Read first:
- `docs/loop.md`
- `docs/skills.md`
- `README.md`
- `AGENTS.md`
- `decisions/0008-loop-orchestrates-review-fix.md`

Constraints:
- Do not change the core thesis.
- Document that `--review` is opt-in.
- Keep changes scoped to the review-fix subloop.
- Sync embedded skills if any skill files changed.

Done means:
- `docs/loop.md` documents `--review`, `--max-review-rounds`, `LOOP_REVIEW_CMD`, `LOOP_FIX_CMD`, and `REVIEW.md`.
- `docs/skills.md` and `README.md` describe review/fix as loop-orchestrated when `--review` is set.
- `AGENTS.md` current state and working conventions mention the new loop capability.
- Tests pass.

Verify:
```bash
cd /home/daniel/build/knack && ./tests/run.sh && cd cli && go test ./...
```

Status: pending

codex
The current review skill still has the pre-`REVIEW.md` output contract, while the ADR already defines the `REVIEW.md` handoff. I can still make the fix skill consume that intended artifact without changing review or queue state.
exec
/usr/bin/bash -lc "sed -n '1,260p' cli/embedded/skills/fix/SKILL.md" in /home/daniel/build/knack
 succeeded in 0ms:
---
name: fix
description: Use when addressing review findings and feeding them back into the loop. Converts review findings into new verifiable work units in `.loop/<name>/QUEUE.md` and runs another loop pass. Triggers on "fix the review findings", "address the feedback", "rework based on review", "fix what review found", or when review produced findings that need to be resolved.
---

# Fix

Address review findings by triaging them and generating new work units that feed back into the loop. The `fix` skill is the bridge between review and another build pass; it owns the triage and the queue format. It appends to the same `.loop/<name>/QUEUE.md` file as the unit being fixed.

## Procedure

1. **Read the review findings.** Each finding from the `review` skill is a candidate work unit. `review` has already classified findings as trivial / actionable / disputed / deferred; use those classifications or re-triage as needed.

2. **Triage.** Not every finding warrants a work unit:
   - **Actionable** — the finding identifies a real issue with a clear fix. → Create a work unit.
   - **Trivial** — one-line fix, no risk. → Fix it now, don't create a unit.
   - **Disputed** — the finding is wrong or the reviewer is being overly cautious. → Note the disagreement, don't create a unit. Move on.
   - **Deferred** — the finding is valid but not worth fixing now. → Note it in `AGENTS.md` or a backlog, don't create a unit.

3. **Create work units** for actionable findings. Use the standard work unit format from the `plan` skill (single-unit excerpt below):

````markdown
## <fix for the finding — observable outcome>

Why:
<reference to the review finding — which axis, what was found>

Read first:
- <context the worker needs: file, ADR, or area>
- <2–4 entries; context, not scope>

Constraints:
- <boundary>
- <what must stay true or out of bounds; if a file is named, it is "don't touch X" or "X's public API must not change", not "update X">

Done means:
- <the finding is resolved>
- <no new issue introduced>

Verify:
```bash
<deterministic command that proves the fix>
```

Status: pending
````

4. **Append to QUEUE.md.** Read the existing `.loop/<name>/QUEUE.md` first, then append the new units. If the queue was completed (all units `done`), the new units extend it. If the queue still has pending units, the new units are added after them. Preserve the existing structure and status.

5. **Run the loop.** After writing the units, run:

```bash
./loop.sh run .loop/<name>/QUEUE.md
```

The loop will pick up the first pending unit and proceed.

## What fix is not

- Not a re-review — the findings are already known. Fix turns them into work.
- Not a debate — if a finding is disputed, note it and move on. Don't argue with the review in the queue.
- Not a rewrite — each fix unit should be narrow. If a finding requires a broad rewrite, that's a new explore → plan cycle, not a fix.

## Guardrail

Fix units must not break what review already approved. Each fix unit's `Done means:` should include a no-regression condition. The verify command should cover both the fix and the preservation of what was already working.

## Output

- New work units appended to `.loop/<name>/QUEUE.md`
- A summary of what was triaged (actionable / trivial / disputed / deferred)
- Suggestion to run the loop

exec
/usr/bin/bash -lc 'git status --short' in /home/daniel/build/knack
 succeeded in 0ms:
 M .loop/review-loop/EVIDENCE.md
 M .loop/review-loop/HANDOFF.md
 M .loop/review-loop/QUEUE.md

exec
/usr/bin/bash -lc "ls -l cli/sync-skills.sh && sed -n '1,180p' cli/sync-skills.sh" in /home/daniel/build/knack
 succeeded in 0ms:
-rwxr-xr-x 1 daniel daniel 151 Jul  6 03:06 cli/sync-skills.sh
#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "$0")"
rm -rf embedded/skills
mkdir -p embedded/skills
cp -r ../.agents/skills/* embedded/skills/

codex
I’m going to remove the stale “fix runs the loop” behavior from the fix instructions and make `REVIEW.md` the explicit input contract. The appended unit template also needs to match the current queue format exactly, without extra fields.
exec
/usr/bin/bash -lc "sed -n '1,260p' .agents/skills/plan/SKILL.md" in /home/daniel/build/knack
 succeeded in 0ms:
---
name: plan
description: Use when converting a software task, bug, cleanup, or vague human intent into a disposable `.loop/<name>/QUEUE.md` loop packet of verifiable work units. The planner picks a short, descriptive name for the cycle (e.g., `go-cli`, `parser-bug`).
---

# Plan

Convert intent into a `.loop/<name>/QUEUE.md` loop packet: a bounded queue of disposable, independently verifiable work units. Pick a short, descriptive name for the cycle (e.g., `go-cli`, `parser-bug`).

The goal is not to create durable specs. The goal is to give the loop runner work units that can be attempted one at a time and externally verified. When the work is done, the queue is deleted.

## Entry points

You may enter plan from two places:

- **After explore** — the codebase has been read, intent has been grilled, decisions may already be captured as ADRs in `decisions/`, terms may be in `glossary.md`. Use that context. Don't re-derive what explore already established.
- **Directly** — for small work where explore isn't needed. A bug fix, a patch, a small feature. Skip the ceremony.

If you're entering directly and the work is large or greenfield (no existing codebase to ground against), consider producing disposable specs first (see "Big work" below).

## When a queue already exists

If `.loop/<name>/QUEUE.md` already exists, read it before writing. If it is stale, doesn't match the current code, or no longer reflects the real goal, discard it and write a fresh queue. Plans are disposable coordination state; regenerating from the actual codebase is cheaper than salvaging a drifting plan.

## Work unit types

A work unit is whatever shape the work is. Pick the right type per unit:

- **vertical slice** — crosses enough layers to produce a user-visible or system-visible improvement. The default preference.
- **patch** — small, localized fix. One change, one verify.
- **investigation** — produces findings or ADRs, not necessarily code. Verify checks that the findings exist and are recorded.
- **bug fix** — reproduce → fix → verify. The verify command must fail before the fix and pass after.
- **refactor** — restructure without behavior change. Verify checks that existing tests still pass.

"Vertical slice" is the preferred default, not a required format. The planner prefers slices and rejects horizontal phases, but a unit can be a patch, investigation, or bug fix when the work genuinely isn't sliceable.

## Planning procedure

1. Restate the user's goal as an observable outcome.
2. Identify the strongest deterministic verification command available now.
3. Split work into units that each leave the repo better if the loop stops immediately after.
4. Prefer vertical slices. Reject horizontal phases (see below).
5. For every unit, pick its type from the taxonomy above.
6. Keep the queue short enough for bounded execution (prefer 2-5 units) and within the runner's hard stops (max ticks, no-progress detection). If the work is larger, the loop will pause and write a handoff; the next session resumes from there.
7. If no deterministic verification exists, make the first unit create or identify one.
8. If a decision crystallizes during decomposition, capture it as an ADR (use the `decide` skill).
9. If domain terms are ambiguous or inconsistent, define them (use the `domain-modeling` skill).

## Valid work unit test

A work unit is valid only if all are true:

- It has one observable outcome.
- It has one verification command.
- That verification command is deterministic and executable by the runner (not an LLM-as-judge).
- It has a narrow scope.
- Its constraints state what must stay true or what is out of bounds — never what to edit. If a constraint names a file, it is "don't touch X" or "X's public API must not change," not "update X."
- It can leave the repo better if the loop stops immediately afterward.
- It does not depend on future units to have value.

A unit whose outcome cannot be captured by a deterministic `Verify:` command plus a short `Done means:` list is not ready. Vague shape plus weak verify gives the worker nothing to aim at.

## Horizontal phase rejection

Reject and rewrite units named after layers or activities:

- "Types"
- "CLI wiring"
- "Backend"
- "Frontend"
- "Tests"
- "Refactor"
- "Verification phase"
- "Implement all X"

Rewrite them as end-to-end outcomes:

- Bad: `Add JSON result types`
- Good: `validate --json reports one existing validation error as machine-readable JSON`

- Bad: `Wire CLI flag`
- Good: `validate --json and text mode report the same broken-link error on the same fixture`

- Bad: `Write tests`
- Good: `the regression fixture fails before the fix and passes after the fix`

This is a heuristic, not a gate. If a unit genuinely can't be vertical (a patch, an investigation), use the right type instead.

## Big work

For greenfield or large work where there's no existing code to ground against, optionally produce disposable planning artifacts before decomposing into units:

- `.loop/<name>/specs/proposal.md` — what we're building and why. One page.
- `.loop/<name>/specs/design.md` — architecture sketch. How the pieces fit. One page.

These are **disposable**. They're consumed during build, then deleted. They are never canonized, never merged, never treated as source of truth. Code is the source of truth. These exist only to help the agent think before the code exists.

Skip this entirely for small work. A bug fix doesn't need a proposal.

## Queue template

Use this exact structure unless the target repo already has a loop convention:

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
- <if it names a file, it is "don't touch X" or "X's public API must not change", not "update X">

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

### Field notes

- **Header** is `## <outcome>` — no numbered prefix, no "Slice" word. The outcome itself is the title.
- **Agent:** is optional. Omit unless this unit needs a different model or command than the global `LOOP_AGENT_CMD`.
- **Why:** is optional. Fill in only when there's non-obvious context worth preserving. No padding.
- **Read first:** is context, not scope. Two to four entries: ADRs, code areas, or rulings. Prefer areas and rulings over file enumerations.
- **Constraints:** are boundaries. A constraint states what must stay true or what is out of bounds — never what to edit. If it names a file, it is "don't touch X" or "X's public API must not change," not "update X."
- **Done means:** is the acceptance criteria — what must be observably true after the unit.
- **Verify:** is the mechanically enforceable subset of `Done means:`. A unit whose outcome can't be captured by a deterministic verify command plus a short `Done means:` list isn't ready.
- **Status:** starts as `pending`. The loop updates it to `in_progress`, `done`, `verify_failed`, `no_progress`, or `blocked`.

## Disposability

`.loop/<name>/QUEUE.md` is disposable coordination state. When the work is done and verified, delete the cycle's `.loop/<name>/` subdirectory. What persists: code, tests, decisions in `decisions/`, glossary entries in `glossary.md`, ADRs, AGENTS.md. The queue is not an artifact — it's a scratchpad.

exec
/usr/bin/bash -lc "rg -n \"REVIEW.md|actionable|Status: pending|Run the loop|Why:\" .agents/skills cli/embedded/skills docs README.md DESIGN.md loop.sh tests -g '"'!cli/embedded/skills/fix/SKILL.md'"'" in /home/daniel/build/knack
 succeeded in 0ms:
DESIGN.md:121:1. Read the first `Status: pending` work unit from `.loop/<name>/QUEUE.md`.
DESIGN.md:351:Why:
DESIGN.md:372:Status: pending
README.md:38:1. The runner reads the first `Status: pending` work unit from `QUEUE.md`.
tests/run.sh:46:Status: pending
tests/run.sh:74:assert_contains "$repo2/.loop/QUEUE.md" "Status: pending"
tests/run.sh:138:Status: pending
docs/loop.md:34:1. Read the first `Status: pending` unit.
docs/getting-started.md:58:## Run the loop
docs/queue-format.md:38:| `Why:` | optional | Non-obvious context only. |
docs/queue-format.md:83:Status: pending
.agents/skills/fix/SKILL.md:12:1. **Read the review findings.** Each finding from the `review` skill is a candidate work unit. `review` has already classified findings as trivial / actionable / disputed / deferred; use those classifications or re-triage as needed.
.agents/skills/fix/SKILL.md:20:3. **Create work units** for actionable findings. Use the standard work unit format from the `plan` skill (single-unit excerpt below):
.agents/skills/fix/SKILL.md:25:Why:
.agents/skills/fix/SKILL.md:45:Status: pending
.agents/skills/fix/SKILL.md:50:5. **Run the loop.** After writing the units, run:
.agents/skills/fix/SKILL.md:71:- A summary of what was triaged (actionable / trivial / disputed / deferred)
cli/embedded/skills/review/SKILL.md:63:Review findings are not just notes — they are the input to the `fix` skill, which triages them and appends actionable ones as new work units in `.loop/<name>/QUEUE.md`.
cli/embedded/skills/review/SKILL.md:84:Classify each finding as trivial / actionable / disputed / deferred.
cli/embedded/skills/review/SKILL.md:87:- Hand actionable findings to the `fix` skill, which will triage and append work units to `.loop/<name>/QUEUE.md`, or
.agents/skills/plan/SKILL.md:117:Why:
.agents/skills/plan/SKILL.md:138:Status: pending
.agents/skills/plan/SKILL.md:148:- **Why:** is optional. Fill in only when there's non-obvious context worth preserving. No padding.
.agents/skills/review/SKILL.md:63:Review findings are not just notes — they are the input to the `fix` skill, which triages them and appends actionable ones as new work units in `.loop/<name>/QUEUE.md`.
.agents/skills/review/SKILL.md:84:Classify each finding as trivial / actionable / disputed / deferred.
.agents/skills/review/SKILL.md:87:- Hand actionable findings to the `fix` skill, which will triage and append work units to `.loop/<name>/QUEUE.md`, or
cli/embedded/skills/plan/SKILL.md:117:Why:
cli/embedded/skills/plan/SKILL.md:138:Status: pending
cli/embedded/skills/plan/SKILL.md:148:- **Why:** is optional. Fill in only when there's non-obvious context worth preserving. No padding.

codex
I’m editing only `.agents/skills/fix/SKILL.md` now, then I’ll use the existing sync script so the embedded copy matches.
apply patch
patch: completed
/home/daniel/build/knack/.agents/skills/fix/SKILL.md
diff --git a/.agents/skills/fix/SKILL.md b/.agents/skills/fix/SKILL.md
index 33985734bd56d197e0a1abcffb2311a17470c91c..790cfd6acb66c765603cd8f0e9b0c27d23320cf5
--- a/.agents/skills/fix/SKILL.md
+++ b/.agents/skills/fix/SKILL.md
@@ -5,28 +5,35 @@
 
 # Fix
 
-Address review findings by triaging them and generating new work units that feed back into the loop. The `fix` skill is the bridge between review and another build pass; it owns the triage and the queue format. It appends to the same `.loop/<name>/QUEUE.md` file as the unit being fixed.
+Address review findings by triaging `REVIEW.md` and generating new work units that feed back into the loop. The `fix` skill is the bridge between review and another build pass; it owns the triage and the queue format. It appends to the existing `.loop/<name>/QUEUE.md` file for the same cycle.
+
+The loop owns orchestration. When `fix` is invoked by the loop, stop after appending units and reporting the triage summary; the loop will run the next build pass.
+
+## Inputs
+
+- `.loop/<name>/REVIEW.md` — structured findings from the `review` skill. Read it, but do not edit it.
+- `.loop/<name>/QUEUE.md` — the existing work queue. Read it before appending so new units preserve the queue's current structure and statuses.
+
+`REVIEW.md` is expected to include `## Standards`, `## Intent`, `## Speculative`, and `## Summary` sections. The summary's actionable count is the loop's signal; the finding entries are the fix skill's input. Use the finding classification, confidence, and evidence when deciding whether to generate a work unit.
 
 ## Procedure
 
-1. **Read the review findings.** Each finding from the `review` skill is a candidate work unit. `review` has already classified findings as trivial / actionable / disputed / deferred; use those classifications or re-triage as needed.
+1. **Read `REVIEW.md`.** Each standards or intent finding is a candidate work unit. `review` has already classified findings as trivial / actionable / disputed / deferred; use those classifications or re-triage as needed. Treat speculative findings as notes for future explore/plan work unless the review explicitly classifies one as actionable.
 
 2. **Triage.** Not every finding warrants a work unit:
-   - **Actionable** — the finding identifies a real issue with a clear fix. → Create a work unit.
-   - **Trivial** — one-line fix, no risk. → Fix it now, don't create a unit.
-   - **Disputed** — the finding is wrong or the reviewer is being overly cautious. → Note the disagreement, don't create a unit. Move on.
-   - **Deferred** — the finding is valid but not worth fixing now. → Note it in `AGENTS.md` or a backlog, don't create a unit.
+   - **Actionable** — the finding identifies a real issue with a clear fix. Create a work unit.
+   - **Trivial** — one-line fix, no risk. Fix it now if that is within the current invocation's scope; otherwise create a narrow work unit.
+   - **Disputed** — the finding is wrong or the reviewer is being overly cautious. Note the disagreement in the summary, don't create a unit. Move on.
+   - **Deferred** — the finding is valid but not worth fixing now. Note it in the summary or a backlog, don't create a unit.
 
-3. **Create work units** for actionable findings. Use the standard work unit format from the `plan` skill (single-unit excerpt below):
+3. **Create work units** for actionable findings. Use the standard work unit format from the `plan` skill. Each unit must have one observable outcome and one deterministic verify command. Reference the source finding in `Read first:` by pointing to `REVIEW.md` plus any evidence paths from the finding.
 
 ````markdown
 ## <fix for the finding — observable outcome>
-
-Why:
-<reference to the review finding — which axis, what was found>
 
 Read first:
-- <context the worker needs: file, ADR, or area>
+- .loop/<name>/REVIEW.md (<finding id or heading>)
+- <evidence path from the finding, if any>
 - <2–4 entries; context, not scope>
 
 Constraints:
@@ -45,21 +52,16 @@
 Status: pending
 ````
 
-4. **Append to QUEUE.md.** Read the existing `.loop/<name>/QUEUE.md` first, then append the new units. If the queue was completed (all units `done`), the new units extend it. If the queue still has pending units, the new units are added after them. Preserve the existing structure and status.
+4. **Append to `QUEUE.md`.** Append new units to the end of the existing `.loop/<name>/QUEUE.md`. Do not reorder existing units, do not edit existing unit bodies, and do not change any existing `Status:` lines. If there are no actionable findings, append nothing.
 
-5. **Run the loop.** After writing the units, run:
+5. **Stop after the queue edit.** Report the triage summary and the number of appended units. The loop will decide whether to run another build pass.
 
-```bash
-./loop.sh run .loop/<name>/QUEUE.md
-```
-
-The loop will pick up the first pending unit and proceed.
-
 ## What fix is not
 
 - Not a re-review — the findings are already known. Fix turns them into work.
 - Not a debate — if a finding is disputed, note it and move on. Don't argue with the review in the queue.
 - Not a rewrite — each fix unit should be narrow. If a finding requires a broad rewrite, that's a new explore → plan cycle, not a fix.
+- Not the orchestrator — the loop invokes review, invokes fix, and runs the next build pass.
 
 ## Guardrail
 
@@ -67,6 +69,6 @@
 
 ## Output
 
-- New work units appended to `.loop/<name>/QUEUE.md`
+- New `Status: pending` work units appended to `.loop/<name>/QUEUE.md`
 - A summary of what was triaged (actionable / trivial / disputed / deferred)
-- Suggestion to run the loop
+- The number of units appended

diff --git a/.agents/skills/fix/SKILL.md b/.agents/skills/fix/SKILL.md
index 33985734bd56d197e0a1abcffb2311a17470c91c..790cfd6acb66c765603cd8f0e9b0c27d23320cf5
--- a/.agents/skills/fix/SKILL.md
+++ b/.agents/skills/fix/SKILL.md
@@ -5,28 +5,35 @@
 
 # Fix
 
-Address review findings by triaging them and generating new work units that feed back into the loop. The `fix` skill is the bridge between review and another build pass; it owns the triage and the queue format. It appends to the same `.loop/<name>/QUEUE.md` file as the unit being fixed.
+Address review findings by triaging `REVIEW.md` and generating new work units that feed back into the loop. The `fix` skill is the bridge between review and another build pass; it owns the triage and the queue format. It appends to the existing `.loop/<name>/QUEUE.md` file for the same cycle.
+
+The loop owns orchestration. When `fix` is invoked by the loop, stop after appending units and reporting the triage summary; the loop will run the next build pass.
+
+## Inputs
+
+- `.loop/<name>/REVIEW.md` — structured findings from the `review` skill. Read it, but do not edit it.
+- `.loop/<name>/QUEUE.md` — the existing work queue. Read it before appending so new units preserve the queue's current structure and statuses.
+
+`REVIEW.md` is expected to include `## Standards`, `## Intent`, `## Speculative`, and `## Summary` sections. The summary's actionable count is the loop's signal; the finding entries are the fix skill's input. Use the finding classification, confidence, and evidence when deciding whether to generate a work unit.
 
 ## Procedure
 
-1. **Read the review findings.** Each finding from the `review` skill is a candidate work unit. `review` has already classified findings as trivial / actionable / disputed / deferred; use those classifications or re-triage as needed.
+1. **Read `REVIEW.md`.** Each standards or intent finding is a candidate work unit. `review` has already classified findings as trivial / actionable / disputed / deferred; use those classifications or re-triage as needed. Treat speculative findings as notes for future explore/plan work unless the review explicitly classifies one as actionable.
 
 2. **Triage.** Not every finding warrants a work unit:
-   - **Actionable** — the finding identifies a real issue with a clear fix. → Create a work unit.
-   - **Trivial** — one-line fix, no risk. → Fix it now, don't create a unit.
-   - **Disputed** — the finding is wrong or the reviewer is being overly cautious. → Note the disagreement, don't create a unit. Move on.
-   - **Deferred** — the finding is valid but not worth fixing now. → Note it in `AGENTS.md` or a backlog, don't create a unit.
+   - **Actionable** — the finding identifies a real issue with a clear fix. Create a work unit.
+   - **Trivial** — one-line fix, no risk. Fix it now if that is within the current invocation's scope; otherwise create a narrow work unit.
+   - **Disputed** — the finding is wrong or the reviewer is being overly cautious. Note the disagreement in the summary, don't create a unit. Move on.
+   - **Deferred** — the finding is valid but not worth fixing now. Note it in the summary or a backlog, don't create a unit.
 
-3. **Create work units** for actionable findings. Use the standard work unit format from the `plan` skill (single-unit excerpt below):
+3. **Create work units** for actionable findings. Use the standard work unit format from the `plan` skill. Each unit must have one observable outcome and one deterministic verify command. Reference the source finding in `Read first:` by pointing to `REVIEW.md` plus any evidence paths from the finding.
 
 ````markdown
 ## <fix for the finding — observable outcome>
-
-Why:
-<reference to the review finding — which axis, what was found>
 
 Read first:
-- <context the worker needs: file, ADR, or area>
+- .loop/<name>/REVIEW.md (<finding id or heading>)
+- <evidence path from the finding, if any>
 - <2–4 entries; context, not scope>
 
 Constraints:
@@ -45,21 +52,16 @@
 Status: pending
 ````
 
-4. **Append to QUEUE.md.** Read the existing `.loop/<name>/QUEUE.md` first, then append the new units. If the queue was completed (all units `done`), the new units extend it. If the queue still has pending units, the new units are added after them. Preserve the existing structure and status.
+4. **Append to `QUEUE.md`.** Append new units to the end of the existing `.loop/<name>/QUEUE.md`. Do not reorder existing units, do not edit existing unit bodies, and do not change any existing `Status:` lines. If there are no actionable findings, append nothing.
 
-5. **Run the loop.** After writing the units, run:
+5. **Stop after the queue edit.** Report the triage summary and the number of appended units. The loop will decide whether to run another build pass.
 
-```bash
-./loop.sh run .loop/<name>/QUEUE.md
-```
-
-The loop will pick up the first pending unit and proceed.
-
 ## What fix is not
 
 - Not a re-review — the findings are already known. Fix turns them into work.
 - Not a debate — if a finding is disputed, note it and move on. Don't argue with the review in the queue.
 - Not a rewrite — each fix unit should be narrow. If a finding requires a broad rewrite, that's a new explore → plan cycle, not a fix.
+- Not the orchestrator — the loop invokes review, invokes fix, and runs the next build pass.
 
 ## Guardrail
 
@@ -67,6 +69,6 @@
 
 ## Output
 
-- New work units appended to `.loop/<name>/QUEUE.md`
+- New `Status: pending` work units appended to `.loop/<name>/QUEUE.md`
 - A summary of what was triaged (actionable / trivial / disputed / deferred)
-- Suggestion to run the loop
+- The number of units appended

codex
One detail in the first pass still left room for a fix worker to patch code inline for trivial findings. I’m tightening that because this unit’s contract is queue generation, not doing the follow-up work itself.
apply patch
patch: completed
/home/daniel/build/knack/.agents/skills/fix/SKILL.md
diff --git a/.agents/skills/fix/SKILL.md b/.agents/skills/fix/SKILL.md
index 33985734bd56d197e0a1abcffb2311a17470c91c..dfdc934ddba779c23b86e0b6fee137ea56e3bd8f
--- a/.agents/skills/fix/SKILL.md
+++ b/.agents/skills/fix/SKILL.md
@@ -5,28 +5,35 @@
 
 # Fix
 
-Address review findings by triaging them and generating new work units that feed back into the loop. The `fix` skill is the bridge between review and another build pass; it owns the triage and the queue format. It appends to the same `.loop/<name>/QUEUE.md` file as the unit being fixed.
+Address review findings by triaging `REVIEW.md` and generating new work units that feed back into the loop. The `fix` skill is the bridge between review and another build pass; it owns the triage and the queue format. It appends to the existing `.loop/<name>/QUEUE.md` file for the same cycle.
+
+The loop owns orchestration. When `fix` is invoked by the loop, stop after appending units and reporting the triage summary; the loop will run the next build pass.
+
+## Inputs
+
+- `.loop/<name>/REVIEW.md` — structured findings from the `review` skill. Read it, but do not edit it.
+- `.loop/<name>/QUEUE.md` — the existing work queue. Read it before appending so new units preserve the queue's current structure and statuses.
+
+`REVIEW.md` is expected to include `## Standards`, `## Intent`, `## Speculative`, and `## Summary` sections. The summary's actionable count is the loop's signal; the finding entries are the fix skill's input. Use the finding classification, confidence, and evidence when deciding whether to generate a work unit.
 
 ## Procedure
 
-1. **Read the review findings.** Each finding from the `review` skill is a candidate work unit. `review` has already classified findings as trivial / actionable / disputed / deferred; use those classifications or re-triage as needed.
+1. **Read `REVIEW.md`.** Each standards or intent finding is a candidate work unit. `review` has already classified findings as trivial / actionable / disputed / deferred; use those classifications or re-triage as needed. Treat speculative findings as notes for future explore/plan work unless the review explicitly classifies one as actionable.
 
 2. **Triage.** Not every finding warrants a work unit:
-   - **Actionable** — the finding identifies a real issue with a clear fix. → Create a work unit.
-   - **Trivial** — one-line fix, no risk. → Fix it now, don't create a unit.
-   - **Disputed** — the finding is wrong or the reviewer is being overly cautious. → Note the disagreement, don't create a unit. Move on.
-   - **Deferred** — the finding is valid but not worth fixing now. → Note it in `AGENTS.md` or a backlog, don't create a unit.
+   - **Actionable** — the finding identifies a real issue with a clear fix. Create a work unit.
+   - **Trivial** — one-line fix, no risk. Note it in the summary, don't create a unit.
+   - **Disputed** — the finding is wrong or the reviewer is being overly cautious. Note the disagreement in the summary, don't create a unit. Move on.
+   - **Deferred** — the finding is valid but not worth fixing now. Note it in the summary or a backlog, don't create a unit.
 
-3. **Create work units** for actionable findings. Use the standard work unit format from the `plan` skill (single-unit excerpt below):
+3. **Create work units** for actionable findings. Use the standard work unit format from the `plan` skill. Each unit must have one observable outcome and one deterministic verify command. Reference the source finding in `Read first:` by pointing to `REVIEW.md` plus any evidence paths from the finding.
 
 ````markdown
 ## <fix for the finding — observable outcome>
-
-Why:
-<reference to the review finding — which axis, what was found>
 
 Read first:
-- <context the worker needs: file, ADR, or area>
+- .loop/<name>/REVIEW.md (<finding id or heading>)
+- <evidence path from the finding, if any>
 - <2–4 entries; context, not scope>
 
 Constraints:
@@ -45,21 +52,16 @@
 Status: pending
 ````
 
-4. **Append to QUEUE.md.** Read the existing `.loop/<name>/QUEUE.md` first, then append the new units. If the queue was completed (all units `done`), the new units extend it. If the queue still has pending units, the new units are added after them. Preserve the existing structure and status.
+4. **Append to `QUEUE.md`.** Append new units to the end of the existing `.loop/<name>/QUEUE.md`. Do not reorder existing units, do not edit existing unit bodies, and do not change any existing `Status:` lines. If there are no actionable findings, append nothing.
 
-5. **Run the loop.** After writing the units, run:
+5. **Stop after the queue edit.** Report the triage summary and the number of appended units. The loop will decide whether to run another build pass.
 
-```bash
-./loop.sh run .loop/<name>/QUEUE.md
-```
-
-The loop will pick up the first pending unit and proceed.
-
 ## What fix is not
 
 - Not a re-review — the findings are already known. Fix turns them into work.
 - Not a debate — if a finding is disputed, note it and move on. Don't argue with the review in the queue.
 - Not a rewrite — each fix unit should be narrow. If a finding requires a broad rewrite, that's a new explore → plan cycle, not a fix.
+- Not the orchestrator — the loop invokes review, invokes fix, and runs the next build pass.
 
 ## Guardrail
 
@@ -67,6 +69,6 @@
 
 ## Output
 
-- New work units appended to `.loop/<name>/QUEUE.md`
+- New `Status: pending` work units appended to `.loop/<name>/QUEUE.md`
 - A summary of what was triaged (actionable / trivial / disputed / deferred)
-- Suggestion to run the loop
+- The number of units appended

diff --git a/.agents/skills/fix/SKILL.md b/.agents/skills/fix/SKILL.md
index 33985734bd56d197e0a1abcffb2311a17470c91c..dfdc934ddba779c23b86e0b6fee137ea56e3bd8f
--- a/.agents/skills/fix/SKILL.md
+++ b/.agents/skills/fix/SKILL.md
@@ -5,28 +5,35 @@
 
 # Fix
 
-Address review findings by triaging them and generating new work units that feed back into the loop. The `fix` skill is the bridge between review and another build pass; it owns the triage and the queue format. It appends to the same `.loop/<name>/QUEUE.md` file as the unit being fixed.
+Address review findings by triaging `REVIEW.md` and generating new work units that feed back into the loop. The `fix` skill is the bridge between review and another build pass; it owns the triage and the queue format. It appends to the existing `.loop/<name>/QUEUE.md` file for the same cycle.
+
+The loop owns orchestration. When `fix` is invoked by the loop, stop after appending units and reporting the triage summary; the loop will run the next build pass.
+
+## Inputs
+
+- `.loop/<name>/REVIEW.md` — structured findings from the `review` skill. Read it, but do not edit it.
+- `.loop/<name>/QUEUE.md` — the existing work queue. Read it before appending so new units preserve the queue's current structure and statuses.
+
+`REVIEW.md` is expected to include `## Standards`, `## Intent`, `## Speculative`, and `## Summary` sections. The summary's actionable count is the loop's signal; the finding entries are the fix skill's input. Use the finding classification, confidence, and evidence when deciding whether to generate a work unit.
 
 ## Procedure
 
-1. **Read the review findings.** Each finding from the `review` skill is a candidate work unit. `review` has already classified findings as trivial / actionable / disputed / deferred; use those classifications or re-triage as needed.
+1. **Read `REVIEW.md`.** Each standards or intent finding is a candidate work unit. `review` has already classified findings as trivial / actionable / disputed / deferred; use those classifications or re-triage as needed. Treat speculative findings as notes for future explore/plan work unless the review explicitly classifies one as actionable.
 
 2. **Triage.** Not every finding warrants a work unit:
-   - **Actionable** — the finding identifies a real issue with a clear fix. → Create a work unit.
-   - **Trivial** — one-line fix, no risk. → Fix it now, don't create a unit.
-   - **Disputed** — the finding is wrong or the reviewer is being overly cautious. → Note the disagreement, don't create a unit. Move on.
-   - **Deferred** — the finding is valid but not worth fixing now. → Note it in `AGENTS.md` or a backlog, don't create a unit.
+   - **Actionable** — the finding identifies a real issue with a clear fix. Create a work unit.
+   - **Trivial** — one-line fix, no risk. Note it in the summary, don't create a unit.
+   - **Disputed** — the finding is wrong or the reviewer is being overly cautious. Note the disagreement in the summary, don't create a unit. Move on.
+   - **Deferred** — the finding is valid but not worth fixing now. Note it in the summary or a backlog, don't create a unit.
 
-3. **Create work units** for actionable findings. Use the standard work unit format from the `plan` skill (single-unit excerpt below):
+3. **Create work units** for actionable findings. Use the standard work unit format from the `plan` skill. Each unit must have one observable outcome and one deterministic verify command. Reference the source finding in `Read first:` by pointing to `REVIEW.md` plus any evidence paths from the finding.
 
 ````markdown
 ## <fix for the finding — observable outcome>
-
-Why:
-<reference to the review finding — which axis, what was found>
 
 Read first:
-- <context the worker needs: file, ADR, or area>
+- .loop/<name>/REVIEW.md (<finding id or heading>)
+- <evidence path from the finding, if any>
 - <2–4 entries; context, not scope>
 
 Constraints:
@@ -45,21 +52,16 @@
 Status: pending
 ````
 
-4. **Append to QUEUE.md.** Read the existing `.loop/<name>/QUEUE.md` first, then append the new units. If the queue was completed (all units `done`), the new units extend it. If the queue still has pending units, the new units are added after them. Preserve the existing structure and status.
+4. **Append to `QUEUE.md`.** Append new units to the end of the existing `.loop/<name>/QUEUE.md`. Do not reorder existing units, do not edit existing unit bodies, and do not change any existing `Status:` lines. If there are no actionable findings, append nothing.
 
-5. **Run the loop.** After writing the units, run:
+5. **Stop after the queue edit.** Report the triage summary and the number of appended units. The loop will decide whether to run another build pass.
 
-```bash
-./loop.sh run .loop/<name>/QUEUE.md
-```
-
-The loop will pick up the first pending unit and proceed.
-
 ## What fix is not
 
 - Not a re-review — the findings are already known. Fix turns them into work.
 - Not a debate — if a finding is disputed, note it and move on. Don't argue with the review in the queue.
 - Not a rewrite — each fix unit should be narrow. If a finding requires a broad rewrite, that's a new explore → plan cycle, not a fix.
+- Not the orchestrator — the loop invokes review, invokes fix, and runs the next build pass.
 
 ## Guardrail
 
@@ -67,6 +69,6 @@
 
 ## Output
 
-- New work units appended to `.loop/<name>/QUEUE.md`
+- New `Status: pending` work units appended to `.loop/<name>/QUEUE.md`
 - A summary of what was triaged (actionable / trivial / disputed / deferred)
-- Suggestion to run the loop
+- The number of units appended

codex
The source fix skill is updated. I’m syncing embedded skills with the repo’s standard script, then I’ll run the unit’s verify command.
exec
/usr/bin/bash -lc ./sync-skills.sh in /home/daniel/build/knack/cli
 succeeded in 0ms:
diff --git a/.agents/skills/fix/SKILL.md b/.agents/skills/fix/SKILL.md
index 33985734bd56d197e0a1abcffb2311a17470c91c..dfdc934ddba779c23b86e0b6fee137ea56e3bd8f
--- a/.agents/skills/fix/SKILL.md
+++ b/.agents/skills/fix/SKILL.md
@@ -5,28 +5,35 @@
 
 # Fix
 
-Address review findings by triaging them and generating new work units that feed back into the loop. The `fix` skill is the bridge between review and another build pass; it owns the triage and the queue format. It appends to the same `.loop/<name>/QUEUE.md` file as the unit being fixed.
+Address review findings by triaging `REVIEW.md` and generating new work units that feed back into the loop. The `fix` skill is the bridge between review and another build pass; it owns the triage and the queue format. It appends to the existing `.loop/<name>/QUEUE.md` file for the same cycle.
+
+The loop owns orchestration. When `fix` is invoked by the loop, stop after appending units and reporting the triage summary; the loop will run the next build pass.
+
+## Inputs
+
+- `.loop/<name>/REVIEW.md` — structured findings from the `review` skill. Read it, but do not edit it.
+- `.loop/<name>/QUEUE.md` — the existing work queue. Read it before appending so new units preserve the queue's current structure and statuses.
+
+`REVIEW.md` is expected to include `## Standards`, `## Intent`, `## Speculative`, and `## Summary` sections. The summary's actionable count is the loop's signal; the finding entries are the fix skill's input. Use the finding classification, confidence, and evidence when deciding whether to generate a work unit.
 
 ## Procedure
 
-1. **Read the review findings.** Each finding from the `review` skill is a candidate work unit. `review` has already classified findings as trivial / actionable / disputed / deferred; use those classifications or re-triage as needed.
+1. **Read `REVIEW.md`.** Each standards or intent finding is a candidate work unit. `review` has already classified findings as trivial / actionable / disputed / deferred; use those classifications or re-triage as needed. Treat speculative findings as notes for future explore/plan work unless the review explicitly classifies one as actionable.
 
 2. **Triage.** Not every finding warrants a work unit:
-   - **Actionable** — the finding identifies a real issue with a clear fix. → Create a work unit.
-   - **Trivial** — one-line fix, no risk. → Fix it now, don't create a unit.
-   - **Disputed** — the finding is wrong or the reviewer is being overly cautious. → Note the disagreement, don't create a unit. Move on.
-   - **Deferred** — the finding is valid but not worth fixing now. → Note it in `AGENTS.md` or a backlog, don't create a unit.
+   - **Actionable** — the finding identifies a real issue with a clear fix. Create a work unit.
+   - **Trivial** — one-line fix, no risk. Note it in the summary, don't create a unit.
+   - **Disputed** — the finding is wrong or the reviewer is being overly cautious. Note the disagreement in the summary, don't create a unit. Move on.
+   - **Deferred** — the finding is valid but not worth fixing now. Note it in the summary or a backlog, don't create a unit.
 
-3. **Create work units** for actionable findings. Use the standard work unit format from the `plan` skill (single-unit excerpt below):
+3. **Create work units** for actionable findings. Use the standard work unit format from the `plan` skill. Each unit must have one observable outcome and one deterministic verify command. Reference the source finding in `Read first:` by pointing to `REVIEW.md` plus any evidence paths from the finding.
 
 ````markdown
 ## <fix for the finding — observable outcome>
-
-Why:
-<reference to the review finding — which axis, what was found>
 
 Read first:
-- <context the worker needs: file, ADR, or area>
+- .loop/<name>/REVIEW.md (<finding id or heading>)
+- <evidence path from the finding, if any>
 - <2–4 entries; context, not scope>
 
 Constraints:
@@ -45,21 +52,16 @@
 Status: pending
 ````
 
-4. **Append to QUEUE.md.** Read the existing `.loop/<name>/QUEUE.md` first, then append the new units. If the queue was completed (all units `done`), the new units extend it. If the queue still has pending units, the new units are added after them. Preserve the existing structure and status.
+4. **Append to `QUEUE.md`.** Append new units to the end of the existing `.loop/<name>/QUEUE.md`. Do not reorder existing units, do not edit existing unit bodies, and do not change any existing `Status:` lines. If there are no actionable findings, append nothing.
 
-5. **Run the loop.** After writing the units, run:
+5. **Stop after the queue edit.** Report the triage summary and the number of appended units. The loop will decide whether to run another build pass.
 
-```bash
-./loop.sh run .loop/<name>/QUEUE.md
-```
-
-The loop will pick up the first pending unit and proceed.
-
 ## What fix is not
 
 - Not a re-review — the findings are already known. Fix turns them into work.
 - Not a debate — if a finding is disputed, note it and move on. Don't argue with the review in the queue.
 - Not a rewrite — each fix unit should be narrow. If a finding requires a broad rewrite, that's a new explore → plan cycle, not a fix.
+- Not the orchestrator — the loop invokes review, invokes fix, and runs the next build pass.
 
 ## Guardrail
 
@@ -67,6 +69,6 @@
 
 ## Output
 
-- New work units appended to `.loop/<name>/QUEUE.md`
+- New `Status: pending` work units appended to `.loop/<name>/QUEUE.md`
 - A summary of what was triaged (actionable / trivial / disputed / deferred)
-- Suggestion to run the loop
+- The number of units appended

codex
Sync completed cleanly. I’m checking the diff shape before verification so I can catch any accidental changes outside the unit scope.
exec
/usr/bin/bash -lc 'git status --short' in /home/daniel/build/knack
 succeeded in 0ms:
 M .agents/skills/fix/SKILL.md
 M .loop/review-loop/EVIDENCE.md
 M .loop/review-loop/HANDOFF.md
 M .loop/review-loop/QUEUE.md
 M cli/embedded/skills/fix/SKILL.md

exec
/usr/bin/bash -lc 'git diff -- .agents/skills/fix/SKILL.md cli/embedded/skills/fix/SKILL.md' in /home/daniel/build/knack
exec
/usr/bin/bash -lc 'diff -r .agents/skills/fix cli/embedded/skills/fix' in /home/daniel/build/knack
 succeeded in 0ms:
diff --git a/.agents/skills/fix/SKILL.md b/.agents/skills/fix/SKILL.md
index 3398573..dfdc934 100644
--- a/.agents/skills/fix/SKILL.md
+++ b/.agents/skills/fix/SKILL.md
@@ -5,28 +5,35 @@ description: Use when addressing review findings and feeding them back into the
 
 # Fix
 
-Address review findings by triaging them and generating new work units that feed back into the loop. The `fix` skill is the bridge between review and another build pass; it owns the triage and the queue format. It appends to the same `.loop/<name>/QUEUE.md` file as the unit being fixed.
+Address review findings by triaging `REVIEW.md` and generating new work units that feed back into the loop. The `fix` skill is the bridge between review and another build pass; it owns the triage and the queue format. It appends to the existing `.loop/<name>/QUEUE.md` file for the same cycle.
+
+The loop owns orchestration. When `fix` is invoked by the loop, stop after appending units and reporting the triage summary; the loop will run the next build pass.
+
+## Inputs
+
+- `.loop/<name>/REVIEW.md` — structured findings from the `review` skill. Read it, but do not edit it.
+- `.loop/<name>/QUEUE.md` — the existing work queue. Read it before appending so new units preserve the queue's current structure and statuses.
+
+`REVIEW.md` is expected to include `## Standards`, `## Intent`, `## Speculative`, and `## Summary` sections. The summary's actionable count is the loop's signal; the finding entries are the fix skill's input. Use the finding classification, confidence, and evidence when deciding whether to generate a work unit.
 
 ## Procedure
 
-1. **Read the review findings.** Each finding from the `review` skill is a candidate work unit. `review` has already classified findings as trivial / actionable / disputed / deferred; use those classifications or re-triage as needed.
+1. **Read `REVIEW.md`.** Each standards or intent finding is a candidate work unit. `review` has already classified findings as trivial / actionable / disputed / deferred; use those classifications or re-triage as needed. Treat speculative findings as notes for future explore/plan work unless the review explicitly classifies one as actionable.
 
 2. **Triage.** Not every finding warrants a work unit:
-   - **Actionable** — the finding identifies a real issue with a clear fix. → Create a work unit.
-   - **Trivial** — one-line fix, no risk. → Fix it now, don't create a unit.
-   - **Disputed** — the finding is wrong or the reviewer is being overly cautious. → Note the disagreement, don't create a unit. Move on.
-   - **Deferred** — the finding is valid but not worth fixing now. → Note it in `AGENTS.md` or a backlog, don't create a unit.
+   - **Actionable** — the finding identifies a real issue with a clear fix. Create a work unit.
+   - **Trivial** — one-line fix, no risk. Note it in the summary, don't create a unit.
+   - **Disputed** — the finding is wrong or the reviewer is being overly cautious. Note the disagreement in the summary, don't create a unit. Move on.
+   - **Deferred** — the finding is valid but not worth fixing now. Note it in the summary or a backlog, don't create a unit.
 
-3. **Create work units** for actionable findings. Use the standard work unit format from the `plan` skill (single-unit excerpt below):
+3. **Create work units** for actionable findings. Use the standard work unit format from the `plan` skill. Each unit must have one observable outcome and one deterministic verify command. Reference the source finding in `Read first:` by pointing to `REVIEW.md` plus any evidence paths from the finding.
 
 ````markdown
 ## <fix for the finding — observable outcome>
 
-Why:
-<reference to the review finding — which axis, what was found>
-
 Read first:
-- <context the worker needs: file, ADR, or area>
+- .loop/<name>/REVIEW.md (<finding id or heading>)
+- <evidence path from the finding, if any>
 - <2–4 entries; context, not scope>
 
 Constraints:
@@ -45,21 +52,16 @@ Verify:
 Status: pending
 ````
 
-4. **Append to QUEUE.md.** Read the existing `.loop/<name>/QUEUE.md` first, then append the new units. If the queue was completed (all units `done`), the new units extend it. If the queue still has pending units, the new units are added after them. Preserve the existing structure and status.
-
-5. **Run the loop.** After writing the units, run:
-
-```bash
-./loop.sh run .loop/<name>/QUEUE.md
-```
+4. **Append to `QUEUE.md`.** Append new units to the end of the existing `.loop/<name>/QUEUE.md`. Do not reorder existing units, do not edit existing unit bodies, and do not change any existing `Status:` lines. If there are no actionable findings, append nothing.
 
-The loop will pick up the first pending unit and proceed.
+5. **Stop after the queue edit.** Report the triage summary and the number of appended units. The loop will decide whether to run another build pass.
 
 ## What fix is not
 
 - Not a re-review — the findings are already known. Fix turns them into work.
 - Not a debate — if a finding is disputed, note it and move on. Don't argue with the review in the queue.
 - Not a rewrite — each fix unit should be narrow. If a finding requires a broad rewrite, that's a new explore → plan cycle, not a fix.
+- Not the orchestrator — the loop invokes review, invokes fix, and runs the next build pass.
 
 ## Guardrail
 
@@ -67,6 +69,6 @@ Fix units must not break what review already approved. Each fix unit's `Done mea
 
 ## Output
 
-- New work units appended to `.loop/<name>/QUEUE.md`
+- New `Status: pending` work units appended to `.loop/<name>/QUEUE.md`
 - A summary of what was triaged (actionable / trivial / disputed / deferred)
-- Suggestion to run the loop
+- The number of units appended
diff --git a/cli/embedded/skills/fix/SKILL.md b/cli/embedded/skills/fix/SKILL.md
index 3398573..dfdc934 100644
--- a/cli/embedded/skills/fix/SKILL.md
+++ b/cli/embedded/skills/fix/SKILL.md
@@ -5,28 +5,35 @@ description: Use when addressing review findings and feeding them back into the
 
 # Fix
 
-Address review findings by triaging them and generating new work units that feed back into the loop. The `fix` skill is the bridge between review and another build pass; it owns the triage and the queue format. It appends to the same `.loop/<name>/QUEUE.md` file as the unit being fixed.
+Address review findings by triaging `REVIEW.md` and generating new work units that feed back into the loop. The `fix` skill is the bridge between review and another build pass; it owns the triage and the queue format. It appends to the existing `.loop/<name>/QUEUE.md` file for the same cycle.
+
+The loop owns orchestration. When `fix` is invoked by the loop, stop after appending units and reporting the triage summary; the loop will run the next build pass.
+
+## Inputs
+
+- `.loop/<name>/REVIEW.md` — structured findings from the `review` skill. Read it, but do not edit it.
+- `.loop/<name>/QUEUE.md` — the existing work queue. Read it before appending so new units preserve the queue's current structure and statuses.
+
+`REVIEW.md` is expected to include `## Standards`, `## Intent`, `## Speculative`, and `## Summary` sections. The summary's actionable count is the loop's signal; the finding entries are the fix skill's input. Use the finding classification, confidence, and evidence when deciding whether to generate a work unit.
 
 ## Procedure
 
-1. **Read the review findings.** Each finding from the `review` skill is a candidate work unit. `review` has already classified findings as trivial / actionable / disputed / deferred; use those classifications or re-triage as needed.
+1. **Read `REVIEW.md`.** Each standards or intent finding is a candidate work unit. `review` has already classified findings as trivial / actionable / disputed / deferred; use those classifications or re-triage as needed. Treat speculative findings as notes for future explore/plan work unless the review explicitly classifies one as actionable.
 
 2. **Triage.** Not every finding warrants a work unit:
-   - **Actionable** — the finding identifies a real issue with a clear fix. → Create a work unit.
-   - **Trivial** — one-line fix, no risk. → Fix it now, don't create a unit.
-   - **Disputed** — the finding is wrong or the reviewer is being overly cautious. → Note the disagreement, don't create a unit. Move on.
-   - **Deferred** — the finding is valid but not worth fixing now. → Note it in `AGENTS.md` or a backlog, don't create a unit.
+   - **Actionable** — the finding identifies a real issue with a clear fix. Create a work unit.
+   - **Trivial** — one-line fix, no risk. Note it in the summary, don't create a unit.
+   - **Disputed** — the finding is wrong or the reviewer is being overly cautious. Note the disagreement in the summary, don't create a unit. Move on.
+   - **Deferred** — the finding is valid but not worth fixing now. Note it in the summary or a backlog, don't create a unit.
 
-3. **Create work units** for actionable findings. Use the standard work unit format from the `plan` skill (single-unit excerpt below):
+3. **Create work units** for actionable findings. Use the standard work unit format from the `plan` skill. Each unit must have one observable outcome and one deterministic verify command. Reference the source finding in `Read first:` by pointing to `REVIEW.md` plus any evidence paths from the finding.
 
 ````markdown
 ## <fix for the finding — observable outcome>
 
-Why:
-<reference to the review finding — which axis, what was found>
-
 Read first:
-- <context the worker needs: file, ADR, or area>
+- .loop/<name>/REVIEW.md (<finding id or heading>)
+- <evidence path from the finding, if any>
 - <2–4 entries; context, not scope>
 
 Constraints:
@@ -45,21 +52,16 @@ Verify:
 Status: pending
 ````
 
-4. **Append to QUEUE.md.** Read the existing `.loop/<name>/QUEUE.md` first, then append the new units. If the queue was completed (all units `done`), the new units extend it. If the queue still has pending units, the new units are added after them. Preserve the existing structure and status.
-
-5. **Run the loop.** After writing the units, run:
-
-```bash
-./loop.sh run .loop/<name>/QUEUE.md
-```
+4. **Append to `QUEUE.md`.** Append new units to the end of the existing `.loop/<name>/QUEUE.md`. Do not reorder existing units, do not edit existing unit bodies, and do not change any existing `Status:` lines. If there are no actionable findings, append nothing.
 
-The loop will pick up the first pending unit and proceed.
+5. **Stop after the queue edit.** Report the triage summary and the number of appended units. The loop will decide whether to run another build pass.
 
 ## What fix is not
 
 - Not a re-review — the findings are already known. Fix turns them into work.
 - Not a debate — if a finding is disputed, note it and move on. Don't argue with the review in the queue.
 - Not a rewrite — each fix unit should be narrow. If a finding requires a broad rewrite, that's a new explore → plan cycle, not a fix.
+- Not the orchestrator — the loop invokes review, invokes fix, and runs the next build pass.
 
 ## Guardrail
 
@@ -67,6 +69,6 @@ Fix units must not break what review already approved. Each fix unit's `Done mea
 
 ## Output
 
-- New work units appended to `.loop/<name>/QUEUE.md`
+- New `Status: pending` work units appended to `.loop/<name>/QUEUE.md`
 - A summary of what was triaged (actionable / trivial / disputed / deferred)
-- Suggestion to run the loop
+- The number of units appended

 succeeded in 0ms:
diff --git a/.agents/skills/fix/SKILL.md b/.agents/skills/fix/SKILL.md
index 33985734bd56d197e0a1abcffb2311a17470c91c..dfdc934ddba779c23b86e0b6fee137ea56e3bd8f
--- a/.agents/skills/fix/SKILL.md
+++ b/.agents/skills/fix/SKILL.md
@@ -5,28 +5,35 @@
 
 # Fix
 
-Address review findings by triaging them and generating new work units that feed back into the loop. The `fix` skill is the bridge between review and another build pass; it owns the triage and the queue format. It appends to the same `.loop/<name>/QUEUE.md` file as the unit being fixed.
+Address review findings by triaging `REVIEW.md` and generating new work units that feed back into the loop. The `fix` skill is the bridge between review and another build pass; it owns the triage and the queue format. It appends to the existing `.loop/<name>/QUEUE.md` file for the same cycle.
+
+The loop owns orchestration. When `fix` is invoked by the loop, stop after appending units and reporting the triage summary; the loop will run the next build pass.
+
+## Inputs
+
+- `.loop/<name>/REVIEW.md` — structured findings from the `review` skill. Read it, but do not edit it.
+- `.loop/<name>/QUEUE.md` — the existing work queue. Read it before appending so new units preserve the queue's current structure and statuses.
+
+`REVIEW.md` is expected to include `## Standards`, `## Intent`, `## Speculative`, and `## Summary` sections. The summary's actionable count is the loop's signal; the finding entries are the fix skill's input. Use the finding classification, confidence, and evidence when deciding whether to generate a work unit.
 
 ## Procedure
 
-1. **Read the review findings.** Each finding from the `review` skill is a candidate work unit. `review` has already classified findings as trivial / actionable / disputed / deferred; use those classifications or re-triage as needed.
+1. **Read `REVIEW.md`.** Each standards or intent finding is a candidate work unit. `review` has already classified findings as trivial / actionable / disputed / deferred; use those classifications or re-triage as needed. Treat speculative findings as notes for future explore/plan work unless the review explicitly classifies one as actionable.
 
 2. **Triage.** Not every finding warrants a work unit:
-   - **Actionable** — the finding identifies a real issue with a clear fix. → Create a work unit.
-   - **Trivial** — one-line fix, no risk. → Fix it now, don't create a unit.
-   - **Disputed** — the finding is wrong or the reviewer is being overly cautious. → Note the disagreement, don't create a unit. Move on.
-   - **Deferred** — the finding is valid but not worth fixing now. → Note it in `AGENTS.md` or a backlog, don't create a unit.
+   - **Actionable** — the finding identifies a real issue with a clear fix. Create a work unit.
+   - **Trivial** — one-line fix, no risk. Note it in the summary, don't create a unit.
+   - **Disputed** — the finding is wrong or the reviewer is being overly cautious. Note the disagreement in the summary, don't create a unit. Move on.
+   - **Deferred** — the finding is valid but not worth fixing now. Note it in the summary or a backlog, don't create a unit.
 
-3. **Create work units** for actionable findings. Use the standard work unit format from the `plan` skill (single-unit excerpt below):
+3. **Create work units** for actionable findings. Use the standard work unit format from the `plan` skill. Each unit must have one observable outcome and one deterministic verify command. Reference the source finding in `Read first:` by pointing to `REVIEW.md` plus any evidence paths from the finding.
 
 ````markdown
 ## <fix for the finding — observable outcome>
-
-Why:
-<reference to the review finding — which axis, what was found>
 
 Read first:
-- <context the worker needs: file, ADR, or area>
+- .loop/<name>/REVIEW.md (<finding id or heading>)
+- <evidence path from the finding, if any>
 - <2–4 entries; context, not scope>
 
 Constraints:
@@ -45,21 +52,16 @@
 Status: pending
 ````
 
-4. **Append to QUEUE.md.** Read the existing `.loop/<name>/QUEUE.md` first, then append the new units. If the queue was completed (all units `done`), the new units extend it. If the queue still has pending units, the new units are added after them. Preserve the existing structure and status.
+4. **Append to `QUEUE.md`.** Append new units to the end of the existing `.loop/<name>/QUEUE.md`. Do not reorder existing units, do not edit existing unit bodies, and do not change any existing `Status:` lines. If there are no actionable findings, append nothing.
 
-5. **Run the loop.** After writing the units, run:
+5. **Stop after the queue edit.** Report the triage summary and the number of appended units. The loop will decide whether to run another build pass.
 
-```bash
-./loop.sh run .loop/<name>/QUEUE.md
-```
-
-The loop will pick up the first pending unit and proceed.
-
 ## What fix is not
 
 - Not a re-review — the findings are already known. Fix turns them into work.
 - Not a debate — if a finding is disputed, note it and move on. Don't argue with the review in the queue.
 - Not a rewrite — each fix unit should be narrow. If a finding requires a broad rewrite, that's a new explore → plan cycle, not a fix.
+- Not the orchestrator — the loop invokes review, invokes fix, and runs the next build pass.
 
 ## Guardrail
 
@@ -67,6 +69,6 @@
 
 ## Output
 
-- New work units appended to `.loop/<name>/QUEUE.md`
+- New `Status: pending` work units appended to `.loop/<name>/QUEUE.md`
 - A summary of what was triaged (actionable / trivial / disputed / deferred)
-- Suggestion to run the loop
+- The number of units appended

codex
The diff is limited to the fix skill and its embedded copy; the pre-existing loop handoff/evidence/queue changes are still untouched by me. I’m running the unit’s verification command now.
exec
/usr/bin/bash -lc 'diff -r .agents/skills/fix cli/embedded/skills/fix && ./tests/run.sh' in /home/daniel/build/knack
diff --git a/.agents/skills/fix/SKILL.md b/.agents/skills/fix/SKILL.md
index 33985734bd56d197e0a1abcffb2311a17470c91c..dfdc934ddba779c23b86e0b6fee137ea56e3bd8f
--- a/.agents/skills/fix/SKILL.md
+++ b/.agents/skills/fix/SKILL.md
@@ -5,28 +5,35 @@
 
 # Fix
 
-Address review findings by triaging them and generating new work units that feed back into the loop. The `fix` skill is the bridge between review and another build pass; it owns the triage and the queue format. It appends to the same `.loop/<name>/QUEUE.md` file as the unit being fixed.
+Address review findings by triaging `REVIEW.md` and generating new work units that feed back into the loop. The `fix` skill is the bridge between review and another build pass; it owns the triage and the queue format. It appends to the existing `.loop/<name>/QUEUE.md` file for the same cycle.
+
+The loop owns orchestration. When `fix` is invoked by the loop, stop after appending units and reporting the triage summary; the loop will run the next build pass.
+
+## Inputs
+
+- `.loop/<name>/REVIEW.md` — structured findings from the `review` skill. Read it, but do not edit it.
+- `.loop/<name>/QUEUE.md` — the existing work queue. Read it before appending so new units preserve the queue's current structure and statuses.
+
+`REVIEW.md` is expected to include `## Standards`, `## Intent`, `## Speculative`, and `## Summary` sections. The summary's actionable count is the loop's signal; the finding entries are the fix skill's input. Use the finding classification, confidence, and evidence when deciding whether to generate a work unit.
 
 ## Procedure
 
-1. **Read the review findings.** Each finding from the `review` skill is a candidate work unit. `review` has already classified findings as trivial / actionable / disputed / deferred; use those classifications or re-triage as needed.
+1. **Read `REVIEW.md`.** Each standards or intent finding is a candidate work unit. `review` has already classified findings as trivial / actionable / disputed / deferred; use those classifications or re-triage as needed. Treat speculative findings as notes for future explore/plan work unless the review explicitly classifies one as actionable.
 
 2. **Triage.** Not every finding warrants a work unit:
-   - **Actionable** — the finding identifies a real issue with a clear fix. → Create a work unit.
-   - **Trivial** — one-line fix, no risk. → Fix it now, don't create a unit.
-   - **Disputed** — the finding is wrong or the reviewer is being overly cautious. → Note the disagreement, don't create a unit. Move on.
-   - **Deferred** — the finding is valid but not worth fixing now. → Note it in `AGENTS.md` or a backlog, don't create a unit.
+   - **Actionable** — the finding identifies a real issue with a clear fix. Create a work unit.
+   - **Trivial** — one-line fix, no risk. Note it in the summary, don't create a unit.
+   - **Disputed** — the finding is wrong or the reviewer is being overly cautious. Note the disagreement in the summary, don't create a unit. Move on.
+   - **Deferred** — the finding is valid but not worth fixing now. Note it in the summary or a backlog, don't create a unit.
 
-3. **Create work units** for actionable findings. Use the standard work unit format from the `plan` skill (single-unit excerpt below):
+3. **Create work units** for actionable findings. Use the standard work unit format from the `plan` skill. Each unit must have one observable outcome and one deterministic verify command. Reference the source finding in `Read first:` by pointing to `REVIEW.md` plus any evidence paths from the finding.
 
 ````markdown
 ## <fix for the finding — observable outcome>
-
-Why:
-<reference to the review finding — which axis, what was found>
 
 Read first:
-- <context the worker needs: file, ADR, or area>
+- .loop/<name>/REVIEW.md (<finding id or heading>)
+- <evidence path from the finding, if any>
 - <2–4 entries; context, not scope>
 
 Constraints:
@@ -45,21 +52,16 @@
 Status: pending
 ````
 
-4. **Append to QUEUE.md.** Read the existing `.loop/<name>/QUEUE.md` first, then append the new units. If the queue was completed (all units `done`), the new units extend it. If the queue still has pending units, the new units are added after them. Preserve the existing structure and status.
+4. **Append to `QUEUE.md`.** Append new units to the end of the existing `.loop/<name>/QUEUE.md`. Do not reorder existing units, do not edit existing unit bodies, and do not change any existing `Status:` lines. If there are no actionable findings, append nothing.
 
-5. **Run the loop.** After writing the units, run:
+5. **Stop after the queue edit.** Report the triage summary and the number of appended units. The loop will decide whether to run another build pass.
 
-```bash
-./loop.sh run .loop/<name>/QUEUE.md
-```
-
-The loop will pick up the first pending unit and proceed.
-
 ## What fix is not
 
 - Not a re-review — the findings are already known. Fix turns them into work.
 - Not a debate — if a finding is disputed, note it and move on. Don't argue with the review in the queue.
 - Not a rewrite — each fix unit should be narrow. If a finding requires a broad rewrite, that's a new explore → plan cycle, not a fix.
+- Not the orchestrator — the loop invokes review, invokes fix, and runs the next build pass.
 
 ## Guardrail
 
@@ -67,6 +69,6 @@
 
 ## Output
 
-- New work units appended to `.loop/<name>/QUEUE.md`
+- New `Status: pending` work units appended to `.loop/<name>/QUEUE.md`
 - A summary of what was triaged (actionable / trivial / disputed / deferred)
-- Suggestion to run the loop
+- The number of units appended

 succeeded in 2874ms:
Valid skill: /home/daniel/build/knack/.agents/skills/build
Valid skill: /home/daniel/build/knack/.agents/skills/decide
Valid skill: /home/daniel/build/knack/.agents/skills/domain-modeling
Valid skill: /home/daniel/build/knack/.agents/skills/explore
Valid skill: /home/daniel/build/knack/.agents/skills/fix
Valid skill: /home/daniel/build/knack/.agents/skills/plan
Valid skill: /home/daniel/build/knack/.agents/skills/review
knack tests passed

codex
```text
Unit: fix skill consumes REVIEW.md and appends fix units to QUEUE.md
Changed: .agents/skills/fix/SKILL.md; cli/embedded/skills/fix/SKILL.md
Verify expected: cd /home/daniel/build/knack && diff -r .agents/skills/fix cli/embedded/skills/fix && ./tests/run.sh
Notes: Local verify command exited 0. Existing .loop/review-loop state files were not edited.
```
diff --git a/.agents/skills/fix/SKILL.md b/.agents/skills/fix/SKILL.md
index 33985734bd56d197e0a1abcffb2311a17470c91c..dfdc934ddba779c23b86e0b6fee137ea56e3bd8f
--- a/.agents/skills/fix/SKILL.md
+++ b/.agents/skills/fix/SKILL.md
@@ -5,28 +5,35 @@
 
 # Fix
 
-Address review findings by triaging them and generating new work units that feed back into the loop. The `fix` skill is the bridge between review and another build pass; it owns the triage and the queue format. It appends to the same `.loop/<name>/QUEUE.md` file as the unit being fixed.
+Address review findings by triaging `REVIEW.md` and generating new work units that feed back into the loop. The `fix` skill is the bridge between review and another build pass; it owns the triage and the queue format. It appends to the existing `.loop/<name>/QUEUE.md` file for the same cycle.
+
+The loop owns orchestration. When `fix` is invoked by the loop, stop after appending units and reporting the triage summary; the loop will run the next build pass.
+
+## Inputs
+
+- `.loop/<name>/REVIEW.md` — structured findings from the `review` skill. Read it, but do not edit it.
+- `.loop/<name>/QUEUE.md` — the existing work queue. Read it before appending so new units preserve the queue's current structure and statuses.
+
+`REVIEW.md` is expected to include `## Standards`, `## Intent`, `## Speculative`, and `## Summary` sections. The summary's actionable count is the loop's signal; the finding entries are the fix skill's input. Use the finding classification, confidence, and evidence when deciding whether to generate a work unit.
 
 ## Procedure
 
-1. **Read the review findings.** Each finding from the `review` skill is a candidate work unit. `review` has already classified findings as trivial / actionable / disputed / deferred; use those classifications or re-triage as needed.
+1. **Read `REVIEW.md`.** Each standards or intent finding is a candidate work unit. `review` has already classified findings as trivial / actionable / disputed / deferred; use those classifications or re-triage as needed. Treat speculative findings as notes for future explore/plan work unless the review explicitly classifies one as actionable.
 
 2. **Triage.** Not every finding warrants a work unit:
-   - **Actionable** — the finding identifies a real issue with a clear fix. → Create a work unit.
-   - **Trivial** — one-line fix, no risk. → Fix it now, don't create a unit.
-   - **Disputed** — the finding is wrong or the reviewer is being overly cautious. → Note the disagreement, don't create a unit. Move on.
-   - **Deferred** — the finding is valid but not worth fixing now. → Note it in `AGENTS.md` or a backlog, don't create a unit.
+   - **Actionable** — the finding identifies a real issue with a clear fix. Create a work unit.
+   - **Trivial** — one-line fix, no risk. Note it in the summary, don't create a unit.
+   - **Disputed** — the finding is wrong or the reviewer is being overly cautious. Note the disagreement in the summary, don't create a unit. Move on.
+   - **Deferred** — the finding is valid but not worth fixing now. Note it in the summary or a backlog, don't create a unit.
 
-3. **Create work units** for actionable findings. Use the standard work unit format from the `plan` skill (single-unit excerpt below):
+3. **Create work units** for actionable findings. Use the standard work unit format from the `plan` skill. Each unit must have one observable outcome and one deterministic verify command. Reference the source finding in `Read first:` by pointing to `REVIEW.md` plus any evidence paths from the finding.
 
 ````markdown
 ## <fix for the finding — observable outcome>
-
-Why:
-<reference to the review finding — which axis, what was found>
 
 Read first:
-- <context the worker needs: file, ADR, or area>
+- .loop/<name>/REVIEW.md (<finding id or heading>)
+- <evidence path from the finding, if any>
 - <2–4 entries; context, not scope>
 
 Constraints:
@@ -45,21 +52,16 @@
 Status: pending
 ````
 
-4. **Append to QUEUE.md.** Read the existing `.loop/<name>/QUEUE.md` first, then append the new units. If the queue was completed (all units `done`), the new units extend it. If the queue still has pending units, the new units are added after them. Preserve the existing structure and status.
+4. **Append to `QUEUE.md`.** Append new units to the end of the existing `.loop/<name>/QUEUE.md`. Do not reorder existing units, do not edit existing unit bodies, and do not change any existing `Status:` lines. If there are no actionable findings, append nothing.
 
-5. **Run the loop.** After writing the units, run:
+5. **Stop after the queue edit.** Report the triage summary and the number of appended units. The loop will decide whether to run another build pass.
 
-```bash
-./loop.sh run .loop/<name>/QUEUE.md
-```
-
-The loop will pick up the first pending unit and proceed.
-
 ## What fix is not
 
 - Not a re-review — the findings are already known. Fix turns them into work.
 - Not a debate — if a finding is disputed, note it and move on. Don't argue with the review in the queue.
 - Not a rewrite — each fix unit should be narrow. If a finding requires a broad rewrite, that's a new explore → plan cycle, not a fix.
+- Not the orchestrator — the loop invokes review, invokes fix, and runs the next build pass.
 
 ## Guardrail
 
@@ -67,6 +69,6 @@
 
 ## Output
 
-- New work units appended to `.loop/<name>/QUEUE.md`
+- New `Status: pending` work units appended to `.loop/<name>/QUEUE.md`
 - A summary of what was triaged (actionable / trivial / disputed / deferred)
-- Suggestion to run the loop
+- The number of units appended

tokens used
45,057
```text
Unit: fix skill consumes REVIEW.md and appends fix units to QUEUE.md
Changed: .agents/skills/fix/SKILL.md; cli/embedded/skills/fix/SKILL.md
Verify expected: cd /home/daniel/build/knack && diff -r .agents/skills/fix cli/embedded/skills/fix && ./tests/run.sh
Notes: Local verify command exited 0. Existing .loop/review-loop state files were not edited.
```
````

What this proves:
- The verify command passed for this work unit in the current repo state.

What remains unverified:
- Anything outside the verify command's proof scope.

## 2026-07-10T19:53:53-06:00 — loop.sh orchestrates bounded build-review-fix rounds

Status: done

Unit:
````markdown
## loop.sh orchestrates bounded build-review-fix rounds

Read first:
- `loop.sh`
- `decisions/0008-loop-orchestrates-review-fix.md`
- `.loop/review-loop/specs/proposal.md`
- `tests/run.sh`

Constraints:
- Default behavior must be unchanged when `--review` is not set.
- `--max-ticks` remains the total build-tick budget across all rounds.
- Add `--max-review-rounds` (default 2) and `--review` flag.
- Support `LOOP_REVIEW_CMD` and `LOOP_FIX_CMD` env vars with fallback to `LOOP_AGENT_CMD`.
- Do not break the `Agent:` per-unit override for build ticks.
- Preserve `HANDOFF.md` and `EVIDENCE.md` behavior.
- New tests must not require a real model or API key.

Done means:
- `loop.sh --review` runs the build queue, then runs review, then fix if actionable, then loops.
- `loop.sh` stops when `REVIEW.md` reports `actionable: 0`, `--max-review-rounds` is reached, `--max-ticks` is reached, or no new units are generated.
- `tests/run.sh` has a test that uses fake workers to run a full `build → review → fix → review` cycle.
- `./tests/run.sh` and `bash -n loop.sh` pass.

Verify:
```bash
bash -n /home/daniel/build/knack/loop.sh && /home/daniel/build/knack/tests/run.sh
```

Status: pending````

Files changed:
```text
 M .agents/skills/fix/SKILL.md
 M .loop/review-loop/EVIDENCE.md
 M .loop/review-loop/HANDOFF.md
 M .loop/review-loop/QUEUE.md
 M cli/embedded/skills/fix/SKILL.md
 M loop.sh
 M tests/run.sh
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
OpenAI Codex v0.144.1
--------
workdir: /home/daniel/build/knack
model: gpt-5.5
provider: openai
approval: never
sandbox: danger-full-access
reasoning effort: high
reasoning summaries: none
session id: 019f4ede-8f15-7bf1-98f0-c51b9d3c035d
--------
user
# Knack Worker

You are one worker tick in a bounded loop. Complete exactly one work unit, then stop.

Load and follow the **build** skill in `.agents/skills/build/` before doing any work.

## Rules

1. Read `AGENTS.md` first if it exists — it contains operational context.
2. Read the current work unit carefully, especially its `Verify` command.
3. Do only the work needed for that unit.
4. Do not start another work unit.
5. Do not mark the unit complete yourself.
6. Do not edit `.loop/<name>/EVIDENCE.md`; the runner writes evidence after verification.
7. If the unit is blocked, make the smallest useful note in your final response and stop.
8. Stay within the runner's hard stops (max ticks, no-progress detection). If the unit is too large for one tick, do as much as keeps the repo working and report what remains.
9. The worker's scope is the unit's outcome plus its constraints. Keep the diff narrow and aligned with that scope.
10. If verification fails while you are working, fix the cause if it belongs to this unit; otherwise stop and report the blocker.

## Success standard

Your job is not to claim success. Your job is to make the repository state satisfy the unit's `Verify` command. The runner will execute that command after you exit.

## Output

End with a compact terminal handoff. This is not `.loop/<name>/HANDOFF.md` — the runner writes that file. This is your own summary to the runner:

```text
Unit: <title>
Changed: <brief file/area list>
Verify expected: <command from unit>
Notes: <blockers or caveats, if any>
```

Current work unit from /home/daniel/build/knack/.loop/review-loop/QUEUE.md:

## loop.sh orchestrates bounded build-review-fix rounds

Read first:
- `loop.sh`
- `decisions/0008-loop-orchestrates-review-fix.md`
- `.loop/review-loop/specs/proposal.md`
- `tests/run.sh`

Constraints:
- Default behavior must be unchanged when `--review` is not set.
- `--max-ticks` remains the total build-tick budget across all rounds.
- Add `--max-review-rounds` (default 2) and `--review` flag.
- Support `LOOP_REVIEW_CMD` and `LOOP_FIX_CMD` env vars with fallback to `LOOP_AGENT_CMD`.
- Do not break the `Agent:` per-unit override for build ticks.
- Preserve `HANDOFF.md` and `EVIDENCE.md` behavior.
- New tests must not require a real model or API key.

Done means:
- `loop.sh --review` runs the build queue, then runs review, then fix if actionable, then loops.
- `loop.sh` stops when `REVIEW.md` reports `actionable: 0`, `--max-review-rounds` is reached, `--max-ticks` is reached, or no new units are generated.
- `tests/run.sh` has a test that uses fake workers to run a full `build → review → fix → review` cycle.
- `./tests/run.sh` and `bash -n loop.sh` pass.

Verify:
```bash
bash -n /home/daniel/build/knack/loop.sh && /home/daniel/build/knack/tests/run.sh
```

Status: pending
hook: SessionStart
hook: SessionStart Completed
codex
I’ll load the required build skill and the local project context first, then I’ll inspect the loop and tests around this unit before editing.
exec
/usr/bin/bash -lc "sed -n '1,240p' AGENTS.md" in /home/daniel/build/knack
 succeeded in 0ms:
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

The loop (`loop.sh`), all seven skills, and the Go CLI are built. The loop supports per-unit `Agent:` overrides, `LOOP_AGENT_CMD` for agent-agnosticism, and writes `.loop/<name>/HANDOFF.md` on non-clean exits. The skills cover the full explore → plan → build → review → fix flow plus two shared skills (decide, domain-modeling). The CLI (`cli/`) provides `skills init|check`, `validate`, `decisions list|show|check`, `status`, `glossary check`, and `instructions`. It embeds the default skills via `go:embed` and scaffolds them into target projects.

## Core artifacts

- `.loop/<name>/QUEUE.md` — disposable queue of verifiable work units.
- `.loop/<name>/EVIDENCE.md` — append-only ledger of what each tick proved. Durable: keep it after deleting QUEUE.md so `decisions check` can still see which ADRs the completed cycle referenced.
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
- When the queue is complete and verified, the cycle's `.loop/<name>/QUEUE.md` is disposable. The human deletes it; the loop does not. `EVIDENCE.md` is the durable ledger — keep it so `decisions check` can still trace which ADRs the cycle referenced.
- A work unit must leave the repo better if the loop stops immediately after it.
- Work units are whatever shape the work is — not forced into "vertical slices."
- `LOOP_AGENT_CMD` overrides the worker invocation (agent-agnostic: pi, claude, codex, opencode, etc.).
- Per-unit `Agent:` field in QUEUE.md overrides `LOOP_AGENT_CMD` for one unit (model routing).
- Work units are `## <outcome>` headers with `Read first:`, `Constraints:`, `Done means:`, and `Verify:` fields. `Done means:` is the acceptance criteria; `Verify:` is the mechanically enforceable subset. The gap between them is the review surface.
- `Read first:` is context, not scope — 2–4 entries of ADRs, code areas, or rulings.
- `Constraints:` state boundaries. A constraint says what must stay true or what is out of bounds, never what to edit. If it names a file, it is "don't touch X" or "X's public API must not change," not "update X."
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

exec
/usr/bin/bash -lc "sed -n '1,240p' .agents/skills/build/SKILL.md" in /home/daniel/build/knack
 succeeded in 0ms:
---
name: build
description: Use when implementing one work unit from a `.loop/<name>/QUEUE.md` queue. The worker-side skill — read the unit, do the work, don't self-certify, end with a handoff. Triggers on "build", "implement", "apply this unit", "do the work", "run the loop", or when a work unit needs to be executed. Also use when the loop invokes the worker for a tick.
---

# Build

Implement one work unit from `.loop/<name>/QUEUE.md`. Do the work. Don't self-certify. The loop runner owns the verify gate — your job is to make the repository state satisfy the unit's `Verify` command, not to claim success.

## Core rules

The runner injects `prompts/worker.md` at the start of every tick. If it is not in your context, read it now. Its `Rules` and `Output` sections are the canonical source for this skill; the sections below elaborate on decisions, operational learnings, blockers, and units that are too large for one tick.

> **Scope note:** Updating `AGENTS.md` or writing an ADR during a tick is a durable trace, not part of the unit's `Verify` scope. Do it only when the tick teaches you something that would save the next session time.

## Scope

The unit's scope is its outcome plus its constraints. The worker determines which files to change and how. The `Verify:` command is the mechanically enforceable subset of `Done means:`.

## Capturing decisions during build

If you discover an architectural ruling while implementing — "we need to handle X this way because Y" — capture it as an ADR using the `decide` skill. Do this inline, not after the unit. Decisions made during implementation are the most valuable kind because they come from the code pushing back.

## Capturing operational learnings

If you learn something about how the project works — a build command, a test convention, a gotcha — that isn't already in `AGENTS.md`, add it. `AGENTS.md` is the living operational context. If it grows stale, the next agent session starts with worse information.

Don't add trivia. Add things that would have saved you time if you'd known them upfront.

## When the unit is too big

If the work unit is larger than what can be done in one tick:

- Do as much as you can while keeping the repo in a working state.
- End with a handoff note explaining what remains.
- The runner will re-queue the unit if verify fails.

Do not try to do everything in one tick if the work genuinely needs more. The loop is designed for multiple ticks.

## When you're blocked

If you hit a blocker you can't resolve within the unit's scope:

1. State the blocker clearly in your final output.
2. Note what would unblock you (a decision, a dependency, a missing file).
3. Stop. Don't thrash.

The runner will mark the unit as `blocked` and write a handoff. The next session can pick up from there.

## Output

The output format is defined in the worker prompt (`prompts/worker.md`). End with a compact terminal handoff there.

codex
The build skill confirms the runner owns verification and this tick should only satisfy the current unit. I’m reading the loop, ADR, proposal, worker prompt, and test harness now to pin down the expected orchestration behavior.
exec
/usr/bin/bash -lc "sed -n '1,320p' tests/run.sh" in /home/daniel/build/knack
 succeeded in 0ms:
#!/usr/bin/env bash
set -euo pipefail

root=$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)
tmp=$(mktemp -d)
trap 'rm -rf "$tmp"' EXIT

assert_contains() {
  local file=$1 pattern=$2
  if ! grep -Fq -- "$pattern" "$file"; then
    echo "expected $file to contain: $pattern" >&2
    echo "--- $file ---" >&2
    cat "$file" >&2
    exit 1
  fi
}

make_queue() {
  local dir=$1 verify=$2
  mkdir -p "$dir/.loop"
  cat > "$dir/.loop/QUEUE.md" <<EOF
# Loop Queue: test

Goal:
Exercise the loop.

Stop condition:
\`$verify\` exits 0.

## the test fixture reaches its verify condition

Read first:
- This queue file.

Constraints:
- Do not modify the queue by hand.

Verify:
\`\`\`bash
$verify
\`\`\`

Done means:
- The verify command exits 0.

Status: pending
EOF
}

bash -n "$root/loop.sh"
"$root/loop.sh" run "$root/examples/smoke/.loop/smoke/QUEUE.md" --dry-run >/tmp/loop-dry-run.txt
assert_contains /tmp/loop-dry-run.txt "Verify:"
assert_contains /tmp/loop-dry-run.txt "test -f smoke.done"

repo1="$tmp/repo-pass"
mkdir -p "$repo1"
make_queue "$repo1" "test -f smoke.done"
LOOP_AGENT_CMD='touch smoke.done; echo worker pass' "$root/loop.sh" run "$repo1/.loop/QUEUE.md" --max-ticks 1 >/tmp/loop-pass.txt
assert_contains "$repo1/.loop/QUEUE.md" "Status: done"
assert_contains "$repo1/.loop/EVIDENCE.md" "Status: done"
assert_contains "$repo1/.loop/EVIDENCE.md" "worker pass"

repo2="$tmp/repo-fail"
mkdir -p "$repo2"
make_queue "$repo2" "test -f never-created"
set +e
LOOP_AGENT_CMD='echo worker failed to create file' "$root/loop.sh" run "$repo2/.loop/QUEUE.md" --max-ticks 1 >/tmp/loop-fail.txt 2>&1
code=$?
set -e
if [[ $code -eq 0 ]]; then
  echo "expected verify failure to exit nonzero" >&2
  exit 1
fi
assert_contains "$repo2/.loop/QUEUE.md" "Status: pending"
assert_contains "$repo2/.loop/EVIDENCE.md" "Status: verify_failed"
assert_contains /tmp/loop-fail.txt "retrying once"

repo3="$tmp/target-repo"
queue_home="$tmp/external-queue"
mkdir -p "$repo3" "$queue_home/.loop"
make_queue "$queue_home" "test -f target.done"
LOOP_AGENT_CMD='pwd > worker.pwd; touch target.done' "$root/loop.sh" run "$queue_home/.loop/QUEUE.md" --repo "$repo3" --max-ticks 1 >/tmp/loop-repo.txt
assert_contains "$queue_home/.loop/QUEUE.md" "Status: done"
test -f "$repo3/target.done"
assert_contains "$repo3/worker.pwd" "$repo3"

# Handoff file is written on non-clean exit (verify failure, max ticks hit)
# Unit was reset to pending for retry, so it appears in Remaining
assert_contains "$repo2/.loop/HANDOFF.md" "## Remaining"
assert_contains "$repo2/.loop/HANDOFF.md" "the test fixture reaches its verify condition"

# Handoff shows blocked unit in In progress when worker exits nonzero
repo5="$tmp/repo-blocked"
mkdir -p "$repo5/.loop"
make_queue "$repo5" "test -f blocked.done"
set +e
LOOP_AGENT_CMD='exit 1' "$root/loop.sh" run "$repo5/.loop/QUEUE.md" --max-ticks 1 >/tmp/loop-blocked.txt 2>&1
code=$?
set -e
if [[ $code -eq 0 ]]; then
  echo "expected blocked worker to exit nonzero" >&2
  exit 1
fi
assert_contains "$repo5/.loop/QUEUE.md" "Status: blocked"
assert_contains "$repo5/.loop/HANDOFF.md" "## In progress"
assert_contains "$repo5/.loop/HANDOFF.md" "blocked"

# Per-unit Agent: override
repo4="$tmp/repo-agent-override"
mkdir -p "$repo4/.loop"
cat > "$repo4/.loop/QUEUE.md" <<EOF
# Loop Queue: agent override

Goal:
Test per-unit Agent override.

Stop condition:
\`test -f override.done\` exits 0.

## the override worker runs instead of LOOP_AGENT_CMD

Agent: touch override.done

Read first:
- This queue file.

Constraints:
- Do not modify the queue by hand.

Verify:
\`\`\`bash
test -f override.done
\`\`\`

Done means:
- The verify command exits 0.

Status: pending
EOF
LOOP_AGENT_CMD='echo should-not-run' "$root/loop.sh" run "$repo4/.loop/QUEUE.md" --max-ticks 1 >/tmp/loop-override.txt
assert_contains "$repo4/.loop/QUEUE.md" "Status: done"
test -f "$repo4/override.done"

# Default fallback: a fake `pi` on PATH receives the prompt body with --approve
repo_pi="$tmp/repo-pi-default"
mkdir -p "$repo_pi"
fake_bin="$tmp/fake-bin"
mkdir -p "$fake_bin"
cat > "$fake_bin/pi" <<'EOF'
#!/usr/bin/env bash
printf '%s\n' "$@" > pi-args.txt
touch smoke.done
EOF
chmod +x "$fake_bin/pi"
make_queue "$repo_pi" "test -f smoke.done"
env -u LOOP_AGENT_CMD PATH="$fake_bin:$PATH" "$root/loop.sh" run "$repo_pi/.loop/QUEUE.md" --max-ticks 1 >/tmp/loop-pi-default.txt
assert_contains "$repo_pi/.loop/QUEUE.md" "Status: done"
assert_contains "$repo_pi/pi-args.txt" "--no-session"
assert_contains "$repo_pi/pi-args.txt" "--approve"
assert_contains "$repo_pi/pi-args.txt" "the test fixture reaches its verify condition"

# LOOP_AGENT_CMD invocations receive LOOP_PROMPT_FILE pointing at the prompt
repo_lpf="$tmp/repo-loop-prompt-file"
mkdir -p "$repo_lpf"
make_queue "$repo_lpf" "test -f lpf.done"
LOOP_AGENT_CMD='test -n "$LOOP_PROMPT_FILE" && test -f "$LOOP_PROMPT_FILE" && cp "$LOOP_PROMPT_FILE" captured-prompt.txt; touch lpf.done' \
  "$root/loop.sh" run "$repo_lpf/.loop/QUEUE.md" --max-ticks 1 >/tmp/loop-lpf.txt
assert_contains "$repo_lpf/.loop/QUEUE.md" "Status: done"
test -f "$repo_lpf/captured-prompt.txt"
assert_contains "$repo_lpf/captured-prompt.txt" "the test fixture reaches its verify condition"

if command -v skills-ref >/dev/null 2>&1; then
  for skill_dir in "$root/.agents/skills"/*; do
    if [[ -d "$skill_dir" ]]; then
      skills-ref validate "$skill_dir"
    fi
  done
fi

echo "knack tests passed"

exec
/usr/bin/bash -lc "sed -n '1,260p' prompts/worker.md" in /home/daniel/build/knack
 succeeded in 0ms:
# Knack Worker

You are one worker tick in a bounded loop. Complete exactly one work unit, then stop.

Load and follow the **build** skill in `.agents/skills/build/` before doing any work.

## Rules

1. Read `AGENTS.md` first if it exists — it contains operational context.
2. Read the current work unit carefully, especially its `Verify` command.
3. Do only the work needed for that unit.
4. Do not start another work unit.
5. Do not mark the unit complete yourself.
6. Do not edit `.loop/<name>/EVIDENCE.md`; the runner writes evidence after verification.
7. If the unit is blocked, make the smallest useful note in your final response and stop.
8. Stay within the runner's hard stops (max ticks, no-progress detection). If the unit is too large for one tick, do as much as keeps the repo working and report what remains.
9. The worker's scope is the unit's outcome plus its constraints. Keep the diff narrow and aligned with that scope.
10. If verification fails while you are working, fix the cause if it belongs to this unit; otherwise stop and report the blocker.

## Success standard

Your job is not to claim success. Your job is to make the repository state satisfy the unit's `Verify` command. The runner will execute that command after you exit.

## Output

End with a compact terminal handoff. This is not `.loop/<name>/HANDOFF.md` — the runner writes that file. This is your own summary to the runner:

```text
Unit: <title>
Changed: <brief file/area list>
Verify expected: <command from unit>
Notes: <blockers or caveats, if any>
```

exec
/usr/bin/bash -lc "sed -n '1,260p' loop.sh" in /home/daniel/build/knack
 succeeded in 0ms:
#!/usr/bin/env bash
set -euo pipefail

usage() {
  cat <<'EOF'
Usage:
  ./loop.sh run <queue> [--repo DIR] [--max-ticks N] [--dry-run]

The queue is usually .loop/<name>/QUEUE.md in the target repo. Use --repo when the
queue lives outside the repository it should operate on.
EOF
}

die() {
  echo "knack: $*" >&2
  exit 1
}

abs_path() {
  python -c 'import os, sys; print(os.path.abspath(sys.argv[1]))' "$1"
}

first_pending_unit() {
  awk '
    BEGIN { in_block = 0; block = ""; status = ""; found = 0 }
    /^## / {
      if (in_block && status == "pending") { printf "%s", block; found = 1; exit }
      in_block = 1; block = $0 "\n"; status = ""; next
    }
    in_block {
      block = block $0 "\n"
      if ($0 ~ /^Status:[[:space:]]*pending[[:space:]]*$/) status = "pending"
    }
    END { if (!found && in_block && status == "pending") printf "%s", block }
  ' "$1"
}

extract_verify() {
  awk '
    BEGIN { after_verify = 0; in_fence = 0 }
    after_verify && /^```/ { if (!in_fence) { in_fence = 1; next } else exit }
    after_verify && in_fence { print; next }
    /^Verify:[[:space:]]*$/ { after_verify = 1 }
  ' "$1"
}

extract_agent() {
  awk '/^Agent:[[:space:]]+/ { sub(/^Agent:[[:space:]]*/, ""); gsub(/[[:space:]]+$/, ""); print; exit }' "$1"
}

set_status() {
  python - "$1" "$2" "$3" <<'PY'
import sys
from pathlib import Path

queue, title, status = sys.argv[1:]
path = Path(queue)
lines = path.read_text().splitlines(keepends=True)
heading = f"## {title}"
inside = False
changed = False
out = []
for line in lines:
    stripped = line.strip()
    if stripped.startswith("## ") and not stripped.startswith("### "):
        if inside and not changed:
            out.append(f"Status: {status}\n")
            changed = True
        inside = stripped == heading
    if inside and stripped.startswith("Status:") and not changed:
        out.append(f"Status: {status}\n")
        changed = True
        continue
    out.append(line)
if inside and not changed:
    if out and not out[-1].endswith("\n"):
        out[-1] += "\n"
    out.append(f"Status: {status}\n")
    changed = True
if not changed:
    raise SystemExit(f"unit not found or status already changed: {title}")
path.write_text("".join(out))
PY
}

write_handoff() {
  python - "$1" "$2" "$3" <<'PY'
import sys, re
from pathlib import Path
from datetime import datetime

queue, evidence, handoff = sys.argv[1:]
lines = Path(queue).read_text().splitlines()

units = []
current = None
for line in lines:
    if re.match(r'^## ', line) and not re.match(r'^###', line):
        if current:
            units.append(current)
        current = {"title": line[3:].strip(), "status": "pending"}
    elif current:
        m = re.match(r'^Status:\s*(\S+)', line)
        if m:
            current["status"] = m.group(1)
if current:
    units.append(current)

pending = [u for u in units if u["status"] != "done"]
if not pending:
    sys.exit(0)

completed = [u for u in units if u["status"] == "done"]
in_progress = [u for u in units if u["status"] in ("in_progress", "verify_failed", "no_progress", "blocked")]
remaining = [u for u in units if u["status"] == "pending"]

out = [
    f"# Handoff: {Path(queue).stem}",
    f"Generated: {datetime.now().isoformat()}",
    "",
    "## Completed",
]
out += [f"- {u['title']}" for u in completed] or ["- (none)"]
out += ["", "## In progress"]
out += [f"- {u['title']} (status: {u['status']})" for u in in_progress] or ["- (none)"]
out += ["", "## Remaining"]
out += [f"- {u['title']}" for u in remaining] or ["- (none)"]
out += ["", "## Next action"]
if in_progress:
    out.append(f"Re-run loop after addressing the {in_progress[0]['status']} state of: {in_progress[0]['title']}.")
elif remaining:
    out.append(f"Re-run loop to continue with: {remaining[0]['title']}.")
else:
    out.append("Queue is complete.")

Path(handoff).write_text("\n".join(out) + "\n")
PY
}

work_snapshot() {
  local repo_dir=$1
  if git -C "$repo_dir" rev-parse --is-inside-work-tree >/dev/null 2>&1; then
    {
      git -C "$repo_dir" diff -- . ':(exclude).loop' || true
      git -C "$repo_dir" status --short --untracked-files=all | awk '$2 !~ /^\.loop\// { print }'
    } | sha256sum | awk '{print $1}'
  else
    echo "__no_git__"
  fi
}

changed_files() {
  local repo_dir=$1
  if git -C "$repo_dir" rev-parse --is-inside-work-tree >/dev/null 2>&1; then
    git -C "$repo_dir" status --short --untracked-files=all || true
  else
    echo "git status unavailable: not inside a git repo"
  fi
}

append_evidence() {
  local evidence=$1 title=$2 status=$3 verify=$4 verify_out=$5 agent_out=$6 repo_dir=$7 unit_file=$8
  mkdir -p "$(dirname "$evidence")"
  {
    echo
    echo "## $(date -Iseconds) — $title"
    echo
    echo "Status: $status"
    echo
    echo "Unit:"
    echo '````markdown'
    cat "$unit_file"
    echo '````'
    echo
    echo "Files changed:"
    echo '```text'
    changed_files "$repo_dir"
    echo '```'
    echo
    echo "Verify command:"
    echo '```bash'
    echo "$verify"
    echo '```'
    echo
    echo "Verify output:"
    echo '```text'
    cat "$verify_out"
    echo '```'
    echo
    echo "Worker output:"
    echo '````text'
    cat "$agent_out"
    echo '````'
    echo
    echo "What this proves:"
    if [[ "$status" == "done" ]]; then
      echo "- The verify command passed for this work unit in the current repo state."
    else
      echo "- The work unit is not externally verified."
    fi
    echo
    echo "What remains unverified:"
    echo "- Anything outside the verify command's proof scope."
  } >> "$evidence"
}

[[ $# -ge 1 ]] || { usage; exit 1; }
cmd=$1
shift
[[ "$cmd" == "run" ]] || { usage; exit 1; }
[[ $# -ge 1 ]] || die "missing queue path"

queue_abs=$(abs_path "$1")
shift
[[ -f "$queue_abs" ]] || die "queue not found: $queue_abs"

max_ticks=3
dry_run=0
repo_override=""
while [[ $# -gt 0 ]]; do
  case "$1" in
    --repo)
      [[ $# -ge 2 ]] || die "--repo needs a value"
      repo_override=$(abs_path "$2")
      shift 2
      ;;
    --max-ticks)
      [[ $# -ge 2 ]] || die "--max-ticks needs a value"
      max_ticks=$2
      shift 2
      ;;
    --dry-run)
      dry_run=1
      shift
      ;;
    *) die "unknown argument: $1" ;;
  esac
done

script_dir=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
prompt_file="$script_dir/prompts/worker.md"
[[ -f "$prompt_file" ]] || die "worker prompt not found: $prompt_file"

queue_dir=$(dirname "$queue_abs")
if [[ -n "$repo_override" ]]; then
  [[ -d "$repo_override" ]] || die "repo not found: $repo_override"
  repo_dir="$repo_override"
elif [[ $(basename "$queue_dir") == ".loop" ]]; then
  repo_dir=$(dirname "$queue_dir")
else
  repo_dir=$(pwd)
fi
evidence="$queue_dir/EVIDENCE.md"
handoff="$queue_dir/HANDOFF.md"
no_progress_strikes=0

write_handoff_on_exit() {
  local rc=$?
  [[ "${dry_run:-0}" == 0 ]] || return $rc
  [[ -f "${queue_abs:-}" ]] || return $rc

exec
/usr/bin/bash -lc "sed -n '1,260p' .loop/review-loop/specs/proposal.md" in /home/daniel/build/knack
 succeeded in 0ms:
# Proposal: Review-fix subloop orchestration

**Status:** disposable planning artifact for the `review-loop` cycle.

## Problem

`knack` today runs the `build` phase well:

```text
pending unit → build worker → Verify: → done
```

But the default skill flow is `explore → plan → build → review → fix`. The `loop.sh` only executes `build` ticks. Review and fix are left as manual skills. The result is that the loop cannot autonomously:

- review a completed queue
- generate fix work units from findings
- re-run the build pass on those fixes
- stop when review finds no actionable issues

This is the missing piece for an autonomous, bounded development loop.

## Decision

The loop will optionally orchestrate a bounded `build → review → fix` subloop. Review remains a skill; the loop does not implement review logic. The loop only invokes the `review` and `fix` workers and interprets their outputs as signals to continue or stop.

See `decisions/0008-loop-orchestrates-review-fix.md` (Unit 1 of this cycle) for the formal ADR.

## Sequence

```text
build pass
  ↓
if --review and no pending units:
  review worker → writes .loop/<name>/REVIEW.md
  ↓
  if actionable findings > 0:
    fix worker → appends fix units to .loop/<name>/QUEUE.md
    ↓
    build pass on new units
    ↓
    review again
  else:
    cycle complete
```

## Hard stops

- `--max-ticks N` — total build ticks across all rounds
- `--max-review-rounds M` — default 2
- no actionable findings in a review pass
- repeated identical actionable findings (no progress)
- review or fix worker exits non-zero
- fix worker produces no new units despite actionable findings

## New loop options

- `--review` — enable review-fix rounds after build pass completes
- `--max-review-rounds N` — default `2`
- `--review-agent-cmd` — optional agent command for review ticks (defaults to `LOOP_AGENT_CMD`)
- `--fix-agent-cmd` — optional agent command for fix ticks (defaults to `LOOP_AGENT_CMD`)

Environment variables for symmetry:

- `LOOP_REVIEW_CMD` — overrides `--review-agent-cmd`
- `LOOP_FIX_CMD` — overrides `--fix-agent-cmd`

Per-unit `Agent:` override does not apply to review/fix ticks; those are orchestrated by the loop.

## New disposable artifact: `REVIEW.md`

The `review` skill will write a structured file next to `QUEUE.md`:

```markdown
# Review: <cycle>
Generated: <timestamp>

## Standards
- [ ] high | medium | low — <finding>
  evidence: path/to/file.go:42 — "quoted line"

## Intent
- [ ] high | medium | low — <finding>
  evidence: path/to/file.go:42 — "quoted line"

## Speculative
- low / uncitable — <finding>

## Summary
- actionable: N
- trivial: N
- disputed: N
- deferred: N
```

The loop checks the `actionable` count to decide whether to run fix. It does not parse the rest of the file.

## Worker prompts

`prompts/worker.md` stays the same for build ticks. We add two optional prompt templates:

- `prompts/reviewer.md` — loaded for review ticks; tells the worker to load the `review` skill and write `REVIEW.md`
- `prompts/fixer.md` — loaded for fix ticks; tells the worker to load the `fix` skill and consume `REVIEW.md`

If these templates are not present, `loop.sh` falls back to `prompts/worker.md` with the skill name prepended.

## Integration with existing skills

- `review/SKILL.md`: update to read `QUEUE.md` and `EVIDENCE.md`, write `REVIEW.md`, and classify findings.
- `fix/SKILL.md`: update to read `REVIEW.md`, triage findings, and append `Status: pending` units to `QUEUE.md`.
- `build/SKILL.md`: unchanged.

## Tests

`tests/run.sh` will gain a test with fake workers:

1. Build worker creates a file with a deliberate bug.
2. Review worker writes `REVIEW.md` with one actionable finding.
3. Fix worker appends a unit to `QUEUE.md`.
4. Build worker on the new unit fixes the bug.
5. Review worker writes `REVIEW.md` with zero actionable findings.
6. Loop stops cleanly.

## Documentation

- `docs/loop.md` — document `--review`, `--max-review-rounds`, `LOOP_REVIEW_CMD`, `LOOP_FIX_CMD`, `REVIEW.md`
- `docs/skills.md` — clarify that review and fix are normally invoked by the loop when `--review` is set
- `README.md` — mention the optional review-fix subloop
- `AGENTS.md` — update current state and working conventions

## Out of scope

- LLM-as-judge review. Review remains a skill output; the loop only interprets the structured `REVIEW.md`.
- Auto-review after every build tick. Review happens after the build queue drains.
- Integration with `knack status` or `knack view`/`knack list`. Those are separate CLI work.

exec
/usr/bin/bash -lc "sed -n '1,240p' decisions/0008-loop-orchestrates-review-fix.md" in /home/daniel/build/knack
 succeeded in 0ms:
# 0008: The loop orchestrates the review-fix subloop; skills keep the judgment

Date: 2026-07-10
Status: accepted

## Context

The default skill flow is `explore → plan → build → review → fix`, but `loop.sh` only runs `build` ticks. Review and fix are left as manual skills, so the loop cannot autonomously review a completed queue, turn findings into fix units, re-run the build pass, and stop when review is clean. That is the missing piece for a bounded, self-correcting development loop.

The architectural question is *who owns what*. DESIGN.md's "What the loop does NOT do" already states the loop does not run review and does not manage ADRs — those are skill responsibilities. Extending the loop into review/fix risks blurring that boundary: if the loop starts interpreting review content, it stops being the simple, agent-agnostic engine and becomes an LLM-as-judge.

Two roles must stay separate:

- **Orchestration** (the loop) — when to invoke review, when to invoke fix, when to stop. Mechanical and signal-driven.
- **Judgment** (the `review` and `fix` skills) — what counts as a finding, whether it is actionable, how to phrase a fix unit. Adversarial and semantic.

Alternatives considered:

- **LLM-as-judge inside the loop.** Rejected — the loop would interpret review content, breaking agent-agnosticism and the verify-gate principle (the aiming problem: the signal must be the actual property, not a proxy the loop can game).
- **The loop implements review logic.** Rejected — duplicates the `review` skill, drifts from the source of truth, and violates "the loop never reads skills."
- **Manual review/fix only (do nothing).** Rejected — leaves the loop unable to self-correct; the missing piece stays missing.

## Decision

The loop **orchestrates** an optional, bounded `build → review → fix` subloop. It invokes the `review` and `fix` workers and interprets their *structured* outputs as continue/stop signals. It does not implement review or fix logic.

When `--review` is set and the build queue has drained, the loop runs a review worker that writes a structured `REVIEW.md`. The loop reads only the `actionable` count from that file. If it is non-zero, the loop runs a fix worker that appends `Status: pending` units to `QUEUE.md`, then re-runs the build pass and reviews again. It stops when `actionable` is zero, a review-round limit (`--max-review-rounds`, default 2) is hit, the tick budget (`--max-ticks`) is exhausted, or a round generates no new units (no progress).

The boundary, stated as a ruling:

- **The loop owns orchestration and stop conditions** — invoking review/fix, reading the actionable count, enforcing hard stops. Mechanical.
- **The `review` and `fix` skills own judgment** — the two-axis standards/intent review, triage, finding phrasing, and work-unit generation. Semantic.

The loop never parses review content beyond the actionable count and never judges whether a finding is real. Review remains a skill the worker loads; the loop only knows to invoke it and read its summary.

## Consequences

- The loop gains a bounded self-correction capability without becoming an LLM-as-judge. Backpressure stays mechanical: the actionable count is the signal, the hard stops are the guardrails.
- A new contract surface: the loop depends on `REVIEW.md`'s actionable count being honest. A review worker that mislabels findings can loop forever or stop early — the round cap and the no-progress stop are the backstop, not the loop's own judgment.
- The `review` and `fix` skills gain a machine-readable input/output contract (`REVIEW.md`; appended units). Those skill changes land in their own work units; this ADR records only the architectural ruling and the orchestration boundary.
- Review stays opt-in (`--review`); default loop behavior is unchanged.

codex
The existing loop already has the build retry, evidence, handoff, per-unit override, and prompt-file plumbing. I’m reading the remaining control flow before patching so the review/fix path reuses those mechanics instead of forking them.
exec
/usr/bin/bash -lc "sed -n '261,620p' loop.sh" in /home/daniel/build/knack
 succeeded in 0ms:
  write_handoff "$queue_abs" "$evidence" "$handoff" 2>/dev/null || true
  return $rc
}
trap write_handoff_on_exit EXIT

for ((tick = 1; tick <= max_ticks; tick++)); do
  unit=$(first_pending_unit "$queue_abs")
  [[ -n "$unit" ]] || { echo "knack: no pending work units"; exit 0; }

  unit_file=$(mktemp)
  verify_file=$(mktemp)
  agent_out=$(mktemp)
  verify_out=$(mktemp)
  printf '%s' "$unit" > "$unit_file"

  first_line=$(awk 'NR == 1 { print; exit }' "$unit_file")
  title=${first_line#\#\# }
  verify=$(extract_verify "$unit_file")
  [[ -n "$verify" ]] || die "work unit has no Verify fenced block: $title"
  printf '%s\n' "$verify" > "$verify_file"

  if [[ "$dry_run" == 1 ]]; then
    echo "Unit: $title"
    echo "Repo: $repo_dir"
    echo "Verify:"
    cat "$verify_file"
    exit 0
  fi

  echo "knack: tick $tick/$max_ticks — $title"
  set_status "$queue_abs" "$title" "in_progress"
  before=$(work_snapshot "$repo_dir")

  run_prompt=$(mktemp)
  cat > "$run_prompt" <<EOF
$(cat "$prompt_file")

Current work unit from $queue_abs:

$(cat "$unit_file")
EOF

  agent_cmd="${LOOP_AGENT_CMD:-}"
  unit_agent=$(extract_agent "$unit_file")
  if [[ -n "$unit_agent" ]]; then
    agent_cmd="$unit_agent"
  fi

  set +e
  if [[ -n "$agent_cmd" ]]; then
    (cd "$repo_dir" && LOOP_PROMPT_FILE="$run_prompt" bash -lc "$agent_cmd") > "$agent_out" 2>&1
  else
    (cd "$repo_dir" && pi -p --no-session --approve "$(cat "$run_prompt")") > "$agent_out" 2>&1
  fi
  agent_code=$?
  set -e

  if [[ $agent_code -ne 0 ]]; then
    set_status "$queue_abs" "$title" "blocked"
    : > "$verify_out"
    append_evidence "$evidence" "$title" "blocked" "$verify" "$verify_out" "$agent_out" "$repo_dir" "$unit_file"
    cat "$agent_out"
    die "worker exited nonzero for $title"
  fi

  after=$(work_snapshot "$repo_dir")

  set +e
  (cd "$repo_dir" && bash -lc "$verify") > "$verify_out" 2>&1
  verify_code=$?
  set -e

  if [[ $verify_code -eq 0 ]]; then
    set_status "$queue_abs" "$title" "done"
    append_evidence "$evidence" "$title" "done" "$verify" "$verify_out" "$agent_out" "$repo_dir" "$unit_file"
    echo "knack: verified — $title"
    continue
  fi

  if [[ "$before" == "$after" ]]; then
    no_progress_strikes=$((no_progress_strikes + 1))
    if [[ $no_progress_strikes -ge 2 ]]; then
      set_status "$queue_abs" "$title" "no_progress"
      append_evidence "$evidence" "$title" "no_progress" "$verify" "$verify_out" "$agent_out" "$repo_dir" "$unit_file"
      die "no progress after $no_progress_strikes attempts on $title"
    fi
    set_status "$queue_abs" "$title" "pending"
    append_evidence "$evidence" "$title" "verify_failed" "$verify" "$verify_out" "$agent_out" "$repo_dir" "$unit_file"
    echo "knack: verify failed with no progress; retrying once"
    continue
  fi

  set_status "$queue_abs" "$title" "verify_failed"
  append_evidence "$evidence" "$title" "verify_failed" "$verify" "$verify_out" "$agent_out" "$repo_dir" "$unit_file"
  cat "$verify_out"
  die "verify failed for $title"
done

if [[ -n "$(first_pending_unit "$queue_abs")" ]]; then
  die "reached max ticks ($max_ticks) with pending work"
fi

echo "knack: reached max ticks ($max_ticks)"

exec
/usr/bin/bash -lc 'git status --short' in /home/daniel/build/knack
 succeeded in 0ms:
 M .agents/skills/fix/SKILL.md
 M .loop/review-loop/EVIDENCE.md
 M .loop/review-loop/HANDOFF.md
 M .loop/review-loop/QUEUE.md
 M cli/embedded/skills/fix/SKILL.md

exec
/usr/bin/bash -lc 'rg -n "REVIEW|review|fix|LOOP_REVIEW|LOOP_FIX|max-review|review-agent|fix-agent" -S .' in /home/daniel/build/knack
 succeeded in 0ms:
./STEAL-PLAN.md:1:# Plan: Steal five patterns into knack + review-fix subloop
./STEAL-PLAN.md:3:Five steals from gstack/OpenSpec, plus the larger review-fix subloop orchestration. Executed directly (not dogfooded through the loop). Each becomes a work unit that leaves the repo better if it's the last one landed.
./STEAL-PLAN.md:17:- [ ] Review-fix subloop: after build queue drains, run review, generate fix units, re-build, re-review, bounded by `--max-review-rounds` and `--max-ticks`
./STEAL-PLAN.md:34:- [x] `review`
./STEAL-PLAN.md:35:- [x] `fix`
./STEAL-PLAN.md:38:- [ ] `review` and `fix` need loop-orchestrated integration: review writes `REVIEW.md`, fix consumes it
./STEAL-PLAN.md:60:- [ ] Unit 3 — Confidence calibration + pre-emit verification gate (review skill)
./STEAL-PLAN.md:63:- [ ] Unit 6 — Review-fix subloop orchestration
./STEAL-PLAN.md:91:2. **Add two line-prefix parsers** following the existing `parseADRStatus` / `hasGrandfatherLine` idiom (`decisions.go:74-93`):
./STEAL-PLAN.md:92:   - `parseSupersededBy(contents string) string` — matches `Superseded by:` prefix, returns the trimmed value (e.g. `ADR-0007`).
./STEAL-PLAN.md:93:   - `parseSupersedes(contents string) string` — matches `Supersedes:` prefix.
./STEAL-PLAN.md:94:   Both tolerate the bare number or `ADR-NNNN` form (strip an `ADR-` prefix if present, zero-pad to 4 digits via the existing `canonicalADRNumber`).
./STEAL-PLAN.md:170:## Unit 3 — Confidence calibration + pre-emit verification gate (review skill)
./STEAL-PLAN.md:172:**The steal (highest value, per the gstack session):** Two intertwined review upgrades, expressed compactly in knack's shape (~20 lines added to one file):
./STEAL-PLAN.md:176:**File:** `.agents/skills/review/SKILL.md` (+ sync to embedded)
./STEAL-PLAN.md:178:**Changes — add a new section `## Confidence and evidence` after the "Two-axis review" section (after line ~60):**
./STEAL-PLAN.md:187:- **high** — you read the specific code and can quote the line. Promoted to the report, handed to `fix`.
./STEAL-PLAN.md:188:- **medium** — pattern match, likely but not verified against the actual code. Promoted but flagged; `fix` treats it as worth investigating, not worth acting on blindly.
./STEAL-PLAN.md:189:- **low / uncitable** — you can't point to a specific line. **Not promoted.** Banished to a `## Speculative` appendix at the end of the review. Only surfaces if the user reads the appendix.
./STEAL-PLAN.md:210:- `.agents/skills/fix/SKILL.md` (resolve the dangling "or a backlog" reference → LEARNINGS.md)
./STEAL-PLAN.md:233:3. **fix skill** (`fix/SKILL.md:18`): change `Note it in AGENTS.md or a backlog` → `Note it in LEARNINGS.md`. The phantom backlog reference resolves.
./STEAL-PLAN.md:277:## Unit 6 — Review-fix subloop orchestration
./STEAL-PLAN.md:279:**The missing piece:** The loop only executes `build` ticks today. The intended autonomous flow is `build → review → fix → build → review → fix` until the reviewer finds no actionable issues. The loop should own this bounded subloop.
./STEAL-PLAN.md:282:- `loop.sh` — add `--review`, `--max-review-rounds`, optional `LOOP_REVIEW_CMD`/`LOOP_FIX_CMD`
./STEAL-PLAN.md:283:- `prompts/worker.md` — unchanged; still loads `build` for build ticks, `review` for review ticks, `fix` for fix ticks
./STEAL-PLAN.md:284:- `.agents/skills/review/SKILL.md` — update to write a structured `.loop/<name>/REVIEW.md` artifact
./STEAL-PLAN.md:285:- `.agents/skills/fix/SKILL.md` — update to consume `REVIEW.md` and append fix units to `QUEUE.md`
./STEAL-PLAN.md:286:- `decisions/` — new ADR-0008 capturing the loop-orchestrated review-fix decision
./STEAL-PLAN.md:292:2. If `--review` is set and the queue reaches `done` state:
./STEAL-PLAN.md:293:   - Run a review worker with the `review` skill, `QUEUE.md`, `EVIDENCE.md`, and the current diff.
./STEAL-PLAN.md:294:   - The review worker writes `.loop/<name>/REVIEW.md` with findings classified as `actionable`, `trivial`, `disputed`, or `deferred`.
./STEAL-PLAN.md:296:   - If actionable findings: run a fix worker with the `fix` skill and `REVIEW.md`. The fix worker appends new work units to `QUEUE.md`.
./STEAL-PLAN.md:298:   - Re-run review.
./STEAL-PLAN.md:299:   - Stop when no actionable findings, `--max-review-rounds` is reached, `--max-ticks` is reached, or two consecutive rounds produce the same actionable findings.
./STEAL-PLAN.md:302:- `--max-ticks` total (build + review + fix rounds)
./STEAL-PLAN.md:303:- `--max-review-rounds` (default 1 or 2)
./STEAL-PLAN.md:304:- No actionable findings in a review pass
./STEAL-PLAN.md:310:- `LOOP_REVIEW_CMD` (optional; defaults to `LOOP_AGENT_CMD`) for review ticks
./STEAL-PLAN.md:311:- `LOOP_FIX_CMD` (optional; defaults to `LOOP_AGENT_CMD`) for fix ticks
./STEAL-PLAN.md:313:**The `REVIEW.md` artifact (disposable):**
./STEAL-PLAN.md:337:**Implementation cycle:** See `.loop/review-loop/QUEUE.md` and `.loop/review-loop/specs/proposal.md`.
./docs/faq.md:11:knack replaces litespec. It keeps the flow (explore → plan → build → review → fix) and the idea of skills, but drops durable specs as source of truth. Code is the source of truth; work units and handoff files are disposable; decisions and glossary are durable.
./docs/faq.md:55:`Done means:` is the acceptance criteria — what must be true when the unit is finished. `Verify:` is the mechanically enforceable subset that the runner can actually execute. The gap between them is the review surface.
./docs/faq.md:59:Yes. "Vertical slice" is the preferred default, but work units can be patches, bug fixes, investigations, or refactors. The only hard requirement is a deterministic `Verify:` command.
./docs/faq.md:97:### Can I skip explore or review?
./docs/queue-format.md:31:The heading is the outcome. There is no `Slice` prefix and no numbering. Avoid `###` headings inside a work unit — they may confuse simple parsers.
./docs/queue-format.md:50:- `Verify:` is the mechanically enforceable subset of `Done means:`. The gap between them is the review surface.
./docs/queue-format.md:56:# Loop Queue: parser fix
./docs/skills.md:12:| `review` | Run two-axis adversarial review (standards + intent) and generate findings. |
./docs/skills.md:13:| `fix` | Convert review findings into new work units. |
./docs/skills.md:40:`loop.sh` does not read skills. It prepends `prompts/worker.md` to the current work unit and runs the worker. `prompts/worker.md` tells the worker to load the `build` skill by name and path (e.g. "Load and follow the **build** skill in `.agents/skills/build/`"). For other phases (explore, plan, review, fix), the human or the agent invokes the skill directly.
./docs/skills.md:51:Skills are not a rigid gate. The default flow is `explore → plan → build → review → fix`, but any valid subset is fine:
./docs/skills.md:54:small fix → plan → build → done
./docs/skills.md:56:big feature → explore → plan → build → review → fix → done
./docs/loop.md:77:- `loop.sh` does **not** run review or manage ADRs/glossary. Those are handled by the skills.
./DESIGN.md:23:explore  plan  build  review  fix     │                          │
./DESIGN.md:68:├── review/SKILL.md           # two-axis parallel review:
./DESIGN.md:71:│                             # review against actual codebase, not specs
./DESIGN.md:72:├── fix/SKILL.md              # address review findings, generate new work units
./DESIGN.md:74:│                             # used by explore, plan, build, review
./DESIGN.md:76:                              # used by explore, plan, review
./DESIGN.md:81:- **Composable, not monolithic.** The flow is a default path, not a gate. Skills can be invoked independently. `bug → explore → plan → build → done` is as valid as `big feature → explore → plan (with specs) → build → review → fix → done`.
./DESIGN.md:96:**review** — two-axis parallel review (pattern from mattpocock's `code-review`), starting from the work unit and `.loop/<name>/EVIDENCE.md`:
./DESIGN.md:102:**fix** — address review findings. Read the existing `.loop/<name>/QUEUE.md`, append new work units generated from findings, and run another loop pass.
./DESIGN.md:106:**domain-modeling** (shared) — manage the project's ubiquitous language. Challenge terms against the glossary, stress-test with edge-case scenarios, update `glossary.md` inline. Used by explore, plan, review.
./DESIGN.md:170:Re-run loop after fixing the verify command for Slice 3.
./DESIGN.md:178:- It does not run review (that's a skill the agent loads).
./DESIGN.md:205:- No `review` — that's a skill the agent loads, not a CLI command.
./DESIGN.md:229:│   ├── review/SKILL.md
./DESIGN.md:230:│   ├── fix/SKILL.md
./DESIGN.md:303:4. REVIEW (human re-runs the agent, or runs it themselves, with review skill)
./DESIGN.md:304:   agent loads review skill
./DESIGN.md:305:   two-axis parallel review:
./DESIGN.md:308:   review against actual codebase, not against specs
./DESIGN.md:312:   loop runs again on the new work units from review
./DESIGN.md:329:architecture review → explore → plan → build (loop) → review → fix (loop) → done
./DESIGN.md:330:small fix → plan → build (loop) → done                    # skip explore
./DESIGN.md:331:big feature → explore → plan (with specs) → build (loop) → review → fix (loop) → done
./DESIGN.md:380:- **The work unit is whatever shape the work is.** Slice, patch, dedup, move, investigation, bug fix. Not forced into one named shape. "Vertical slice" is a heuristic against horizontal phases, not a required format.
./DESIGN.md:383:- **Done means: is the acceptance criteria; Verify: is the enforceable subset.** The gap between them is the review surface.
./DESIGN.md:391:- The flow shape: explore → plan → build → review → fix (now composable, not rigid)
./DESIGN.md:392:- Skills as procedural knowledge (think/plan/build/review → explore/plan/build/review/fix + shared decide/domain-modeling)
./DESIGN.md:419:- No upstream (explore/plan) or downstream (review) → now covered by skills
./DESIGN.md:427:- Two-axis parallel review (Standards vs Intent, run as parallel sub-agents)
./DESIGN.md:484:1. **Rename "slice" → "work unit"** in `loop.sh` and the planner skill. The QUEUE.md parser in `loop.sh` reads `## Slice N:` headers — change to `## <outcome>` headers (no numbered prefix, no "Slice" word).
./DESIGN.md:485:2. **Split the planner skill.** The current `vertical-slice-planner` becomes `plan`. The `explore`, `build`, `review`, `fix`, `decide`, and `domain-modeling` skills are new and need to be authored.
./README.md:49:Work units are `## <outcome>` headers with `Read first:`, `Constraints:`, `Done means:`, and `Verify:` fields. The `Verify:` command is the mechanically enforceable subset of `Done means:`; the gap between them is the review surface.
./tests/run.sh:30:## the test fixture reaches its verify condition
./tests/run.sh:90:assert_contains "$repo2/.loop/HANDOFF.md" "the test fixture reaches its verify condition"
./tests/run.sh:160:assert_contains "$repo_pi/pi-args.txt" "the test fixture reaches its verify condition"
./tests/run.sh:170:assert_contains "$repo_lpf/captured-prompt.txt" "the test fixture reaches its verify condition"
./AGENTS.md:22:The loop (`loop.sh`), all seven skills, and the Go CLI are built. The loop supports per-unit `Agent:` overrides, `LOOP_AGENT_CMD` for agent-agnosticism, and writes `.loop/<name>/HANDOFF.md` on non-clean exits. The skills cover the full explore → plan → build → review → fix flow plus two shared skills (decide, domain-modeling). The CLI (`cli/`) provides `skills init|check`, `validate`, `decisions list|show|check`, `status`, `glossary check`, and `instructions`. It embeds the default skills via `go:embed` and scaffolds them into target projects.
./AGENTS.md:31:- `.agents/skills/` — procedural knowledge (explore, plan, build, review, fix, decide, domain-modeling).
./AGENTS.md:39:- **review** — two-axis adversarial review (standards + intent). Findings become new work units.
./AGENTS.md:40:- **fix** — address review findings. Generates new work units, feeds back into the loop.
./AGENTS.md:41:- **decide** (shared) — capture architectural rulings as ADRs in `decisions/`. Used by explore, plan, build, review.
./AGENTS.md:42:- **domain-modeling** (shared) — manage `glossary.md`. Used by explore, plan, review.
./AGENTS.md:51:- Review is opt-in. The loop does not run review automatically; the user invokes it when they want adversarial scrutiny.
./AGENTS.md:57:- Work units are `## <outcome>` headers with `Read first:`, `Constraints:`, `Done means:`, and `Verify:` fields. `Done means:` is the acceptance criteria; `Verify:` is the mechanically enforceable subset. The gap between them is the review surface.
./AGENTS.md:60:- Work units are `## <outcome>` headers — no "Slice" prefix, no numbering. Vertical slice is one type of work unit, not the required format.
./AGENTS.md:82:- **Review catches what verify can't.** The queue parser regex `^##\s*(.*)$` matched `###` subheadings as work units — a real bug that diverged from `loop.sh`'s behavior. `go test` passed because no fixture used `###`. Adversarial review against the actual codebase (comparing to `loop.sh`'s parser) found it. The fix: exclude `###` lines explicitly in `isUnitHeader`.
./AGENTS.md:86:- **Workers scope to the outcome plus constraints, not to a file list.** ADR-0005 replaced `Work:` with `Read first:` and `Constraints:`. The unit's scope is its outcome plus its constraints — the worker determines which files to change. The old lesson ("name every file in the work unit") is wrong under the new shape: naming files in constraints smuggles scope the same way `Work:` did. The first plan-shape cycle proved this — the constraint said "no `Work:` refs in skills, prompts, DESIGN.md, or AGENTS.md" and the worker touched exactly those files, leaving 9 other files (test fixtures, examples, README) with stale `Work:` fields. Prefer outcome-level constraints ("no artifact that teaches the format may reference `Work:`") over file-enumerated constraints.
./cli/internal/status/status.go:86:		if strings.HasPrefix(line, "## ") {
./cli/internal/queue/queue_test.go:132:		t.Fatalf("expected smoke fixture to pass, got missing %s", results[0].Missing)
./cli/internal/queue/queue.go:21:		if strings.HasPrefix(strings.TrimSpace(line), "Status:") {
./cli/internal/queue/queue.go:22:			return strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(line), "Status:"))
./cli/internal/queue/queue.go:73:	if strings.HasPrefix(line, "###") {
./cli/internal/queue/queue.go:144:		if strings.HasPrefix(strings.TrimSpace(line), "```") {
./cli/internal/glossary/glossary.go:24:		if !strings.HasPrefix(line, "## ") {
./cli/internal/glossary/glossary.go:27:		term := strings.TrimSpace(strings.TrimPrefix(line, "## "))
./prompts/worker.md:18:10. If verification fails while you are working, fix the cause if it belongs to this unit; otherwise stop and report the blocker.
./cli/internal/decisions/decisions.go:76:		if strings.HasPrefix(strings.TrimSpace(line), "Status:") {
./cli/internal/decisions/decisions.go:77:			return strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(line), "Status:"))
./cli/internal/decisions/decisions.go:88:		if strings.HasPrefix(strings.TrimSpace(line), "Grandfathered:") {
./cli/internal/decisions/decisions.go:106:	prefix := canonical + "-"
./cli/internal/decisions/decisions.go:112:		if strings.HasPrefix(name, prefix) && strings.HasSuffix(name, ".md") {
./cli/internal/decisions/decisions.go:219:	base := strings.TrimSuffix(adr.Filename, ".md")
./decisions/0005-work-units-carry-shape-not-scripts.md:8:The original work unit template had a `Work:` field — a bulleted list of narrow work instructions. In practice, planners filled it with file-by-file edit scripts ("Update `.agents/skills/plan/SKILL.md`: change all `.loop/QUEUE.md` references to..."). The named-cycles cycle proved the failure mode: the worker executed exactly the listed edits, *saw* stale references in files outside the list, noted them in its handoff, and correctly declined to fix them. The script defined the scope, so the worker behaved as a script executor instead of an agent.
./decisions/0005-work-units-carry-shape-not-scripts.md:26:The unit's scope is its outcome plus its constraints — not an enumerated file list. The worker determines which files to change and how. The gap between `Done means:` and `Verify:` is the review surface: opt-in review reads what the command cannot check.
./decisions/0005-work-units-carry-shape-not-scripts.md:32:- Workers use judgment: a worker that sees related breakage inside the outcome's scope fixes it, because nothing tells it "your scope is these five files."
./decisions/0005-work-units-carry-shape-not-scripts.md:34:- The `plan` skill (template + valid-unit test), `build` skill, `review` skill (reads unit fields), `fix` skill, `prompts/worker.md` (rule 9: scope = outcome + constraints), `cli/internal/instructions/instructions.go`, DESIGN.md, and AGENTS.md all need updating. Embedded skills must be re-synced.
./decisions/0004-named-work-cycles-under-loop.md:9:The original design put all loop state in flat files at the root of `.loop/`: `QUEUE.md`, `EVIDENCE.md`, `HANDOFF.md`. This assumes one active work cycle at a time. In practice, work is interleaved — a bug fix lands while a feature is mid-loop, a refactor interrupts a build cycle.
./decisions/0003-the-tool-is-named-knack.md:9:Open question #1 in DESIGN.md: "sliceloop" was a placeholder. The tool is no longer just a slice loop — it's a full workflow harness (explore, plan, build, review, fix) with a loop engine, a skill set, and a CLI. The name needed to reflect that and not bake in the "slice" framing, which we already dropped when work units stopped being forced into vertical slices.
./decisions/0003-the-tool-is-named-knack.md:13:The tool is named **knack**. The CLI binary will be `knack`, the loop's log prefix will be `knack:`, and all documentation references to "sliceloop" become "knack".
./decisions/0002-cli-packages-and-scaffolds-skills.md:20:The CLI will embed the seven default skills (explore, plan, build, review, fix, decide, domain-modeling) via `go:embed` and provide a `skills init` command that scaffolds them into a target project's `.agents/skills/` directory. Projects can then modify or override the scaffolded skills freely — the CLI does not manage them after init.
./cli/embedded/skills/review/SKILL.md:2:name: review
./cli/embedded/skills/review/SKILL.md:3:description: Use when reviewing completed work units or a finished queue. Two-axis adversarial review — standards (does the change follow codebase conventions?) and intent (does the change do what the work unit said?). Reviews against the actual codebase, not specs. Triggers on "review", "check this", "is this right", "what did we miss", "stress-test the implementation", or when work needs adversarial scrutiny before being accepted.
./cli/embedded/skills/review/SKILL.md:8:Adversarial review of completed work. Two axes, run independently so neither pollutes the other:
./cli/embedded/skills/review/SKILL.md:15:## When to review
./cli/embedded/skills/review/SKILL.md:24:## Before you review
./cli/embedded/skills/review/SKILL.md:26:Read the work unit from `.loop/<name>/QUEUE.md` and the evidence from `.loop/<name>/EVIDENCE.md` for the unit you're reviewing. The evidence tells you what the verify command actually proved; the work unit tells you what was promised. Review against the actual codebase, not the specs.
./cli/embedded/skills/review/SKILL.md:28:## Two-axis review
./cli/embedded/skills/review/SKILL.md:53:The `Verify:` command is the mechanically enforceable subset of `Done means:`. The gap between `Done means:` and `Verify:` is the review surface: intent review checks what the verify command cannot.
./cli/embedded/skills/review/SKILL.md:55:> **Intent review is a judgment check, not a deterministic gate.** Be explicit about its scope and confidence. Any finding that becomes a new work unit must have a deterministic `Verify` command the runner can execute.
./cli/embedded/skills/review/SKILL.md:61:## Findings become input to the fix skill
./cli/embedded/skills/review/SKILL.md:63:Review findings are not just notes — they are the input to the `fix` skill, which triages them and appends actionable ones as new work units in `.loop/<name>/QUEUE.md`.
./cli/embedded/skills/review/SKILL.md:65:- **Trivial** findings (a typo, a missing newline) can be fixed inline during review.
./cli/embedded/skills/review/SKILL.md:66:- **Actionable** findings are handed to the `fix` skill. Do not write the work units yourself; `fix` owns the triage and formatting.
./cli/embedded/skills/review/SKILL.md:67:- **Disputed** or **deferred** findings are recorded in the review summary but not turned into units.
./cli/embedded/skills/review/SKILL.md:69:The output of review is a findings summary, not a queue edit.
./cli/embedded/skills/review/SKILL.md:71:## What review is not
./cli/embedded/skills/review/SKILL.md:75:- Not a gate — the loop doesn't enforce review. It's invoked when the user wants it.
./cli/embedded/skills/review/SKILL.md:87:- Hand actionable findings to the `fix` skill, which will triage and append work units to `.loop/<name>/QUEUE.md`, or
./decisions/0008-loop-orchestrates-review-fix.md:1:# 0008: The loop orchestrates the review-fix subloop; skills keep the judgment
./decisions/0008-loop-orchestrates-review-fix.md:8:The default skill flow is `explore → plan → build → review → fix`, but `loop.sh` only runs `build` ticks. Review and fix are left as manual skills, so the loop cannot autonomously review a completed queue, turn findings into fix units, re-run the build pass, and stop when review is clean. That is the missing piece for a bounded, self-correcting development loop.
./decisions/0008-loop-orchestrates-review-fix.md:10:The architectural question is *who owns what*. DESIGN.md's "What the loop does NOT do" already states the loop does not run review and does not manage ADRs — those are skill responsibilities. Extending the loop into review/fix risks blurring that boundary: if the loop starts interpreting review content, it stops being the simple, agent-agnostic engine and becomes an LLM-as-judge.
./decisions/0008-loop-orchestrates-review-fix.md:14:- **Orchestration** (the loop) — when to invoke review, when to invoke fix, when to stop. Mechanical and signal-driven.
./decisions/0008-loop-orchestrates-review-fix.md:15:- **Judgment** (the `review` and `fix` skills) — what counts as a finding, whether it is actionable, how to phrase a fix unit. Adversarial and semantic.
./decisions/0008-loop-orchestrates-review-fix.md:19:- **LLM-as-judge inside the loop.** Rejected — the loop would interpret review content, breaking agent-agnosticism and the verify-gate principle (the aiming problem: the signal must be the actual property, not a proxy the loop can game).
./decisions/0008-loop-orchestrates-review-fix.md:20:- **The loop implements review logic.** Rejected — duplicates the `review` skill, drifts from the source of truth, and violates "the loop never reads skills."
./decisions/0008-loop-orchestrates-review-fix.md:21:- **Manual review/fix only (do nothing).** Rejected — leaves the loop unable to self-correct; the missing piece stays missing.
./decisions/0008-loop-orchestrates-review-fix.md:25:The loop **orchestrates** an optional, bounded `build → review → fix` subloop. It invokes the `review` and `fix` workers and interprets their *structured* outputs as continue/stop signals. It does not implement review or fix logic.
./decisions/0008-loop-orchestrates-review-fix.md:27:When `--review` is set and the build queue has drained, the loop runs a review worker that writes a structured `REVIEW.md`. The loop reads only the `actionable` count from that file. If it is non-zero, the loop runs a fix worker that appends `Status: pending` units to `QUEUE.md`, then re-runs the build pass and reviews again. It stops when `actionable` is zero, a review-round limit (`--max-review-rounds`, default 2) is hit, the tick budget (`--max-ticks`) is exhausted, or a round generates no new units (no progress).
./decisions/0008-loop-orchestrates-review-fix.md:31:- **The loop owns orchestration and stop conditions** — invoking review/fix, reading the actionable count, enforcing hard stops. Mechanical.
./decisions/0008-loop-orchestrates-review-fix.md:32:- **The `review` and `fix` skills own judgment** — the two-axis standards/intent review, triage, finding phrasing, and work-unit generation. Semantic.
./decisions/0008-loop-orchestrates-review-fix.md:34:The loop never parses review content beyond the actionable count and never judges whether a finding is real. Review remains a skill the worker loads; the loop only knows to invoke it and read its summary.
./decisions/0008-loop-orchestrates-review-fix.md:39:- A new contract surface: the loop depends on `REVIEW.md`'s actionable count being honest. A review worker that mislabels findings can loop forever or stop early — the round cap and the no-progress stop are the backstop, not the loop's own judgment.
./decisions/0008-loop-orchestrates-review-fix.md:40:- The `review` and `fix` skills gain a machine-readable input/output contract (`REVIEW.md`; appended units). Those skill changes land in their own work units; this ADR records only the architectural ruling and the orchestration boundary.
./decisions/0008-loop-orchestrates-review-fix.md:41:- Review stays opt-in (`--review`); default loop behavior is unchanged.
./cli/embedded/skills/plan/SKILL.md:17:- **Directly** — for small work where explore isn't needed. A bug fix, a patch, a small feature. Skip the ceremony.
./cli/embedded/skills/plan/SKILL.md:30:- **patch** — small, localized fix. One change, one verify.
./cli/embedded/skills/plan/SKILL.md:32:- **bug fix** — reproduce → fix → verify. The verify command must fail before the fix and pass after.
./cli/embedded/skills/plan/SKILL.md:35:"Vertical slice" is the preferred default, not a required format. The planner prefers slices and rejects horizontal phases, but a unit can be a patch, investigation, or bug fix when the work genuinely isn't sliceable.
./cli/embedded/skills/plan/SKILL.md:82:- Good: `validate --json and text mode report the same broken-link error on the same fixture`
./cli/embedded/skills/plan/SKILL.md:85:- Good: `the regression fixture fails before the fix and passes after the fix`
./cli/embedded/skills/plan/SKILL.md:98:Skip this entirely for small work. A bug fix doesn't need a proposal.
./cli/embedded/skills/plan/SKILL.md:146:- **Header** is `## <outcome>` — no numbered prefix, no "Slice" word. The outcome itself is the title.
./cli/embedded/skills/domain-modeling/SKILL.md:3:description: Use when domain terms surface during explore, plan, or review and need to be defined, challenged, or stress-tested against the project's ubiquitous language. Manages `glossary.md` — the project's shared vocabulary. Triggers on "what does X mean here", "are these the same concept", "let's define our terms", "that term is ambiguous", or when a domain concept is used inconsistently across a conversation or codebase.
./cli/embedded/skills/domain-modeling/SKILL.md:10:This is a shared skill — called by explore, plan, and review. It is not a separate phase.
./cli/embedded/skills/explore/SKILL.md:54:- Small fixes where the problem is already clear
./cli/embedded/skills/fix/SKILL.md:2:name: fix
./cli/embedded/skills/fix/SKILL.md:3:description: Use when addressing review findings and feeding them back into the loop. Converts review findings into new verifiable work units in `.loop/<name>/QUEUE.md` and runs another loop pass. Triggers on "fix the review findings", "address the feedback", "rework based on review", "fix what review found", or when review produced findings that need to be resolved.
./cli/embedded/skills/fix/SKILL.md:8:Address review findings by triaging `REVIEW.md` and generating new work units that feed back into the loop. The `fix` skill is the bridge between review and another build pass; it owns the triage and the queue format. It appends to the existing `.loop/<name>/QUEUE.md` file for the same cycle.
./cli/embedded/skills/fix/SKILL.md:10:The loop owns orchestration. When `fix` is invoked by the loop, stop after appending units and reporting the triage summary; the loop will run the next build pass.
./cli/embedded/skills/fix/SKILL.md:14:- `.loop/<name>/REVIEW.md` — structured findings from the `review` skill. Read it, but do not edit it.
./cli/embedded/skills/fix/SKILL.md:17:`REVIEW.md` is expected to include `## Standards`, `## Intent`, `## Speculative`, and `## Summary` sections. The summary's actionable count is the loop's signal; the finding entries are the fix skill's input. Use the finding classification, confidence, and evidence when deciding whether to generate a work unit.
./cli/embedded/skills/fix/SKILL.md:21:1. **Read `REVIEW.md`.** Each standards or intent finding is a candidate work unit. `review` has already classified findings as trivial / actionable / disputed / deferred; use those classifications or re-triage as needed. Treat speculative findings as notes for future explore/plan work unless the review explicitly classifies one as actionable.
./cli/embedded/skills/fix/SKILL.md:24:   - **Actionable** — the finding identifies a real issue with a clear fix. Create a work unit.
./cli/embedded/skills/fix/SKILL.md:25:   - **Trivial** — one-line fix, no risk. Note it in the summary, don't create a unit.
./cli/embedded/skills/fix/SKILL.md:26:   - **Disputed** — the finding is wrong or the reviewer is being overly cautious. Note the disagreement in the summary, don't create a unit. Move on.
./cli/embedded/skills/fix/SKILL.md:27:   - **Deferred** — the finding is valid but not worth fixing now. Note it in the summary or a backlog, don't create a unit.
./cli/embedded/skills/fix/SKILL.md:29:3. **Create work units** for actionable findings. Use the standard work unit format from the `plan` skill. Each unit must have one observable outcome and one deterministic verify command. Reference the source finding in `Read first:` by pointing to `REVIEW.md` plus any evidence paths from the finding.
./cli/embedded/skills/fix/SKILL.md:32:## <fix for the finding — observable outcome>
./cli/embedded/skills/fix/SKILL.md:35:- .loop/<name>/REVIEW.md (<finding id or heading>)
./cli/embedded/skills/fix/SKILL.md:49:<deterministic command that proves the fix>
./cli/embedded/skills/fix/SKILL.md:59:## What fix is not
./cli/embedded/skills/fix/SKILL.md:61:- Not a re-review — the findings are already known. Fix turns them into work.
./cli/embedded/skills/fix/SKILL.md:62:- Not a debate — if a finding is disputed, note it and move on. Don't argue with the review in the queue.
./cli/embedded/skills/fix/SKILL.md:63:- Not a rewrite — each fix unit should be narrow. If a finding requires a broad rewrite, that's a new explore → plan cycle, not a fix.
./cli/embedded/skills/fix/SKILL.md:64:- Not the orchestrator — the loop invokes review, invokes fix, and runs the next build pass.
./cli/embedded/skills/fix/SKILL.md:68:Fix units must not break what review already approved. Each fix unit's `Done means:` should include a no-regression condition. The verify command should cover both the fix and the preservation of what was already working.
./cli/embedded/skills/decide/SKILL.md:3:description: Use when an architectural ruling crystallizes during any phase — explore, plan, build, or review. Captures decisions as ADRs in `decisions/` so they persist after specs are deleted and code moves on. Triggers on "we decided to", "let's go with", "the ruling is", "why did we choose", "record this decision", or when a design choice with lasting consequences is made.
./cli/embedded/skills/decide/SKILL.md:10:This is a shared skill — called by explore, plan, build, and review whenever a ruling crystallizes. It is not a separate phase.
./cli/embedded/skills/decide/SKILL.md:52:1. Recognize that a ruling just crystallized. If you're in explore, plan, build, or review and you just resolved a tension with lasting consequences, that's a decision.
./cli/internal/skills/skills_test.go:20:	expected := []string{"build", "decide", "domain-modeling", "explore", "fix", "plan", "review"}
./cli/internal/skills/skills.go:207:	return strings.HasPrefix(target, "http://") ||
./cli/internal/skills/skills.go:208:		strings.HasPrefix(target, "https://") ||
./cli/internal/skills/skills.go:209:		strings.HasPrefix(target, "mailto:") ||
./cli/internal/skills/skills.go:210:		strings.HasPrefix(target, "ftp://") ||
./cli/internal/skills/skills.go:211:		strings.HasPrefix(target, "file://") ||
./cli/internal/skills/skills.go:212:		strings.HasPrefix(target, "//")

codex
There are unrelated in-progress loop artifacts and skill sync changes already in the worktree, so I’ll leave those alone. The implementation surface for this unit is the runner script and its shell tests.
codex
I’m going to patch `loop.sh` to add review/fix as a second phase after the build queue drains, with max ticks still counting only build ticks. Then I’ll add one fake-worker shell test that exercises the full round trip without any external model.
apply patch
patch: completed
/home/daniel/build/knack/loop.sh
diff --git a/loop.sh b/loop.sh
index 351aca3855b1ce6c9c08b0c586ad838fa0f6e636..f0f1edbe4c62e8c679082fc908ac0a628ef54117
--- a/loop.sh
+++ b/loop.sh
@@ -4,7 +4,7 @@
 usage() {
   cat <<'EOF'
 Usage:
-  ./loop.sh run <queue> [--repo DIR] [--max-ticks N] [--dry-run]
+  ./loop.sh run <queue> [--repo DIR] [--max-ticks N] [--review] [--max-review-rounds N] [--dry-run]
 
 The queue is usually .loop/<name>/QUEUE.md in the target repo. Use --repo when the
 queue lives outside the repository it should operate on.
@@ -204,6 +204,86 @@
   } >> "$evidence"
 }
 
+extract_actionable_count() {
+  awk '
+    /^[[:space:]]*-[[:space:]]*actionable:[[:space:]]*[0-9]+[[:space:]]*$/ {
+      sub(/^[[:space:]]*-[[:space:]]*actionable:[[:space:]]*/, "")
+      sub(/[[:space:]]*$/, "")
+      print
+      found = 1
+      exit
+    }
+    END { if (!found) exit 1 }
+  ' "$1"
+}
+
+write_review_prompt() {
+  local out=$1
+  local template="$script_dir/prompts/reviewer.md"
+  if [[ -f "$template" ]]; then
+    cat "$template" > "$out"
+  else
+    cat > "$out" <<'EOF'
+# Knack Reviewer
+
+Load and follow the **review** skill in `.agents/skills/review/`.
+Review the completed queue against the current repository state. Write the structured review artifact at the path provided below.
+EOF
+  fi
+  cat >> "$out" <<EOF
+
+Queue: $queue_abs
+Evidence: $evidence
+Review output: $review_file
+EOF
+}
+
+write_fix_prompt() {
+  local out=$1
+  local template="$script_dir/prompts/fixer.md"
+  if [[ -f "$template" ]]; then
+    cat "$template" > "$out"
+  else
+    cat > "$out" <<'EOF'
+# Knack Fixer
+
+Load and follow the **fix** skill in `.agents/skills/fix/`.
+Read the structured review artifact and append any actionable fix work units to the existing queue. Stop after updating the queue.
+EOF
+  fi
+  cat >> "$out" <<EOF
+
+Queue: $queue_abs
+Evidence: $evidence
+Review input: $review_file
+EOF
+}
+
+run_phase_agent() {
+  local phase=$1 prompt=$2 output=$3 cmd=${4:-}
+  if [[ -n "$cmd" ]]; then
+    (
+      cd "$repo_dir"
+      LOOP_PHASE="$phase" \
+        LOOP_PROMPT_FILE="$prompt" \
+        LOOP_QUEUE_FILE="$queue_abs" \
+        LOOP_EVIDENCE_FILE="$evidence" \
+        LOOP_REVIEW_FILE="$review_file" \
+        bash -lc "$cmd"
+    ) > "$output" 2>&1
+  else
+    (
+      cd "$repo_dir"
+      LOOP_PHASE="$phase" \
+        LOOP_PROMPT_FILE="$prompt" \
+        LOOP_QUEUE_FILE="$queue_abs" \
+        LOOP_EVIDENCE_FILE="$evidence" \
+        LOOP_REVIEW_FILE="$review_file" \
+        pi -p --no-session --approve "$(cat "$prompt")"
+    ) > "$output" 2>&1
+  fi
+}
+
 [[ $# -ge 1 ]] || { usage; exit 1; }
 cmd=$1
 shift

diff --git a/loop.sh b/loop.sh
index 351aca3855b1ce6c9c08b0c586ad838fa0f6e636..f0f1edbe4c62e8c679082fc908ac0a628ef54117
--- a/loop.sh
+++ b/loop.sh
@@ -4,7 +4,7 @@
 usage() {
   cat <<'EOF'
 Usage:
-  ./loop.sh run <queue> [--repo DIR] [--max-ticks N] [--dry-run]
+  ./loop.sh run <queue> [--repo DIR] [--max-ticks N] [--review] [--max-review-rounds N] [--dry-run]
 
 The queue is usually .loop/<name>/QUEUE.md in the target repo. Use --repo when the
 queue lives outside the repository it should operate on.
@@ -204,6 +204,86 @@
   } >> "$evidence"
 }
 
+extract_actionable_count() {
+  awk '
+    /^[[:space:]]*-[[:space:]]*actionable:[[:space:]]*[0-9]+[[:space:]]*$/ {
+      sub(/^[[:space:]]*-[[:space:]]*actionable:[[:space:]]*/, "")
+      sub(/[[:space:]]*$/, "")
+      print
+      found = 1
+      exit
+    }
+    END { if (!found) exit 1 }
+  ' "$1"
+}
+
+write_review_prompt() {
+  local out=$1
+  local template="$script_dir/prompts/reviewer.md"
+  if [[ -f "$template" ]]; then
+    cat "$template" > "$out"
+  else
+    cat > "$out" <<'EOF'
+# Knack Reviewer
+
+Load and follow the **review** skill in `.agents/skills/review/`.
+Review the completed queue against the current repository state. Write the structured review artifact at the path provided below.
+EOF
+  fi
+  cat >> "$out" <<EOF
+
+Queue: $queue_abs
+Evidence: $evidence
+Review output: $review_file
+EOF
+}
+
+write_fix_prompt() {
+  local out=$1
+  local template="$script_dir/prompts/fixer.md"
+  if [[ -f "$template" ]]; then
+    cat "$template" > "$out"
+  else
+    cat > "$out" <<'EOF'
+# Knack Fixer
+
+Load and follow the **fix** skill in `.agents/skills/fix/`.
+Read the structured review artifact and append any actionable fix work units to the existing queue. Stop after updating the queue.
+EOF
+  fi
+  cat >> "$out" <<EOF
+
+Queue: $queue_abs
+Evidence: $evidence
+Review input: $review_file
+EOF
+}
+
+run_phase_agent() {
+  local phase=$1 prompt=$2 output=$3 cmd=${4:-}
+  if [[ -n "$cmd" ]]; then
+    (
+      cd "$repo_dir"
+      LOOP_PHASE="$phase" \
+        LOOP_PROMPT_FILE="$prompt" \
+        LOOP_QUEUE_FILE="$queue_abs" \
+        LOOP_EVIDENCE_FILE="$evidence" \
+        LOOP_REVIEW_FILE="$review_file" \
+        bash -lc "$cmd"
+    ) > "$output" 2>&1
+  else
+    (
+      cd "$repo_dir"
+      LOOP_PHASE="$phase" \
+        LOOP_PROMPT_FILE="$prompt" \
+        LOOP_QUEUE_FILE="$queue_abs" \
+        LOOP_EVIDENCE_FILE="$evidence" \
+        LOOP_REVIEW_FILE="$review_file" \
+        pi -p --no-session --approve "$(cat "$prompt")"
+    ) > "$output" 2>&1
+  fi
+}
+
 [[ $# -ge 1 ]] || { usage; exit 1; }
 cmd=$1
 shift

codex
The prompt and phase-runner helpers are in place. Next I’m replacing the single build-only loop with a build-pass plus optional review/fix round controller, keeping the existing build tick body intact.
apply patch
patch: completed
/home/daniel/build/knack/loop.sh
diff --git a/loop.sh b/loop.sh
index 351aca3855b1ce6c9c08b0c586ad838fa0f6e636..35066aad26502fbf3f7d5ac0413747bb134cc0b0
--- a/loop.sh
+++ b/loop.sh
@@ -4,7 +4,7 @@
 usage() {
   cat <<'EOF'
 Usage:
-  ./loop.sh run <queue> [--repo DIR] [--max-ticks N] [--dry-run]
+  ./loop.sh run <queue> [--repo DIR] [--max-ticks N] [--review] [--max-review-rounds N] [--dry-run]
 
 The queue is usually .loop/<name>/QUEUE.md in the target repo. Use --repo when the
 queue lives outside the repository it should operate on.
@@ -204,6 +204,86 @@
   } >> "$evidence"
 }
 
+extract_actionable_count() {
+  awk '
+    /^[[:space:]]*-[[:space:]]*actionable:[[:space:]]*[0-9]+[[:space:]]*$/ {
+      sub(/^[[:space:]]*-[[:space:]]*actionable:[[:space:]]*/, "")
+      sub(/[[:space:]]*$/, "")
+      print
+      found = 1
+      exit
+    }
+    END { if (!found) exit 1 }
+  ' "$1"
+}
+
+write_review_prompt() {
+  local out=$1
+  local template="$script_dir/prompts/reviewer.md"
+  if [[ -f "$template" ]]; then
+    cat "$template" > "$out"
+  else
+    cat > "$out" <<'EOF'
+# Knack Reviewer
+
+Load and follow the **review** skill in `.agents/skills/review/`.
+Review the completed queue against the current repository state. Write the structured review artifact at the path provided below.
+EOF
+  fi
+  cat >> "$out" <<EOF
+
+Queue: $queue_abs
+Evidence: $evidence
+Review output: $review_file
+EOF
+}
+
+write_fix_prompt() {
+  local out=$1
+  local template="$script_dir/prompts/fixer.md"
+  if [[ -f "$template" ]]; then
+    cat "$template" > "$out"
+  else
+    cat > "$out" <<'EOF'
+# Knack Fixer
+
+Load and follow the **fix** skill in `.agents/skills/fix/`.
+Read the structured review artifact and append any actionable fix work units to the existing queue. Stop after updating the queue.
+EOF
+  fi
+  cat >> "$out" <<EOF
+
+Queue: $queue_abs
+Evidence: $evidence
+Review input: $review_file
+EOF
+}
+
+run_phase_agent() {
+  local phase=$1 prompt=$2 output=$3 cmd=${4:-}
+  if [[ -n "$cmd" ]]; then
+    (
+      cd "$repo_dir"
+      LOOP_PHASE="$phase" \
+        LOOP_PROMPT_FILE="$prompt" \
+        LOOP_QUEUE_FILE="$queue_abs" \
+        LOOP_EVIDENCE_FILE="$evidence" \
+        LOOP_REVIEW_FILE="$review_file" \
+        bash -lc "$cmd"
+    ) > "$output" 2>&1
+  else
+    (
+      cd "$repo_dir"
+      LOOP_PHASE="$phase" \
+        LOOP_PROMPT_FILE="$prompt" \
+        LOOP_QUEUE_FILE="$queue_abs" \
+        LOOP_EVIDENCE_FILE="$evidence" \
+        LOOP_REVIEW_FILE="$review_file" \
+        pi -p --no-session --approve "$(cat "$prompt")"
+    ) > "$output" 2>&1
+  fi
+}
+
 [[ $# -ge 1 ]] || { usage; exit 1; }
 cmd=$1
 shift
@@ -215,6 +295,8 @@
 [[ -f "$queue_abs" ]] || die "queue not found: $queue_abs"
 
 max_ticks=3
+max_review_rounds=2
+review_enabled=0
 dry_run=0
 repo_override=""
 while [[ $# -gt 0 ]]; do
@@ -229,6 +311,15 @@
       max_ticks=$2
       shift 2
       ;;
+    --review)
+      review_enabled=1
+      shift
+      ;;
+    --max-review-rounds)
+      [[ $# -ge 2 ]] || die "--max-review-rounds needs a value"
+      max_review_rounds=$2
+      shift 2
+      ;;
     --dry-run)
       dry_run=1
       shift
@@ -252,6 +343,7 @@
 fi
 evidence="$queue_dir/EVIDENCE.md"
 handoff="$queue_dir/HANDOFF.md"
+review_file="$queue_dir/REVIEW.md"
 no_progress_strikes=0
 
 write_handoff_on_exit() {
@@ -263,9 +355,13 @@
 }
 trap write_handoff_on_exit EXIT
 
-for ((tick = 1; tick <= max_ticks; tick++)); do
+tick=1
+review_round=0
+
+while true; do
+while (( tick <= max_ticks )); do
   unit=$(first_pending_unit "$queue_abs")
-  [[ -n "$unit" ]] || { echo "knack: no pending work units"; exit 0; }
+  [[ -n "$unit" ]] || break
 
   unit_file=$(mktemp)
   verify_file=$(mktemp)
@@ -307,11 +403,7 @@
   fi
 
   set +e
-  if [[ -n "$agent_cmd" ]]; then
-    (cd "$repo_dir" && LOOP_PROMPT_FILE="$run_prompt" bash -lc "$agent_cmd") > "$agent_out" 2>&1
-  else
-    (cd "$repo_dir" && pi -p --no-session --approve "$(cat "$run_prompt")") > "$agent_out" 2>&1
-  fi
+  run_phase_agent "build" "$run_prompt" "$agent_out" "$agent_cmd"
   agent_code=$?
   set -e
 
@@ -334,6 +426,7 @@
     set_status "$queue_abs" "$title" "done"
     append_evidence "$evidence" "$title" "done" "$verify" "$verify_out" "$agent_out" "$repo_dir" "$unit_file"
     echo "knack: verified — $title"
+    tick=$((tick + 1))
     continue
   fi
 
@@ -347,6 +440,7 @@
     set_status "$queue_abs" "$title" "pending"
     append_evidence "$evidence" "$title" "verify_failed" "$verify" "$verify_out" "$agent_out" "$repo_dir" "$unit_file"
     echo "knack: verify failed with no progress; retrying once"
+    tick=$((tick + 1))
     continue
   fi
 
@@ -360,4 +454,59 @@
   die "reached max ticks ($max_ticks) with pending work"
 fi
 
-echo "knack: reached max ticks ($max_ticks)"
+if [[ "$review_enabled" == 0 ]]; then
+  echo "knack: no pending work units"
+  exit 0
+fi
+
+if (( review_round >= max_review_rounds )); then
+  die "reached max review rounds ($max_review_rounds)"
+fi
+
+review_round=$((review_round + 1))
+review_prompt=$(mktemp)
+review_out=$(mktemp)
+write_review_prompt "$review_prompt"
+review_cmd="${LOOP_REVIEW_CMD:-${LOOP_AGENT_CMD:-}}"
+
+echo "knack: review round $review_round/$max_review_rounds"
+set +e
+run_phase_agent "review" "$review_prompt" "$review_out" "$review_cmd"
+review_code=$?
+set -e
+if [[ $review_code -ne 0 ]]; then
+  cat "$review_out"
+  die "review worker exited nonzero"
+fi
+[[ -f "$review_file" ]] || die "review worker did not write $review_file"
+actionable=$(extract_actionable_count "$review_file") || die "review file has no actionable summary: $review_file"
+
+if [[ "$actionable" == 0 ]]; then
+  echo "knack: review clean"
+  exit 0
+fi
+
+if (( review_round >= max_review_rounds )); then
+  die "review found $actionable actionable issue(s) at max review rounds ($max_review_rounds)"
+fi
+
+fix_prompt=$(mktemp)
+fix_out=$(mktemp)
+write_fix_prompt "$fix_prompt"
+fix_cmd="${LOOP_FIX_CMD:-${LOOP_AGENT_CMD:-}}"
+
+echo "knack: fix round $review_round/$max_review_rounds — actionable: $actionable"
+set +e
+run_phase_agent "fix" "$fix_prompt" "$fix_out" "$fix_cmd"
+fix_code=$?
+set -e
+if [[ $fix_code -ne 0 ]]; then
+  cat "$fix_out"
+  die "fix worker exited nonzero"
+fi
+
+if [[ -z "$(first_pending_unit "$queue_abs")" ]]; then
+  cat "$fix_out"
+  die "fix worker generated no pending work units"
+fi
+done

diff --git a/loop.sh b/loop.sh
index 351aca3855b1ce6c9c08b0c586ad838fa0f6e636..35066aad26502fbf3f7d5ac0413747bb134cc0b0
--- a/loop.sh
+++ b/loop.sh
@@ -4,7 +4,7 @@
 usage() {
   cat <<'EOF'
 Usage:
-  ./loop.sh run <queue> [--repo DIR] [--max-ticks N] [--dry-run]
+  ./loop.sh run <queue> [--repo DIR] [--max-ticks N] [--review] [--max-review-rounds N] [--dry-run]
 
 The queue is usually .loop/<name>/QUEUE.md in the target repo. Use --repo when the
 queue lives outside the repository it should operate on.
@@ -204,6 +204,86 @@
   } >> "$evidence"
 }
 
+extract_actionable_count() {
+  awk '
+    /^[[:space:]]*-[[:space:]]*actionable:[[:space:]]*[0-9]+[[:space:]]*$/ {
+      sub(/^[[:space:]]*-[[:space:]]*actionable:[[:space:]]*/, "")
+      sub(/[[:space:]]*$/, "")
+      print
+      found = 1
+      exit
+    }
+    END { if (!found) exit 1 }
+  ' "$1"
+}
+
+write_review_prompt() {
+  local out=$1
+  local template="$script_dir/prompts/reviewer.md"
+  if [[ -f "$template" ]]; then
+    cat "$template" > "$out"
+  else
+    cat > "$out" <<'EOF'
+# Knack Reviewer
+
+Load and follow the **review** skill in `.agents/skills/review/`.
+Review the completed queue against the current repository state. Write the structured review artifact at the path provided below.
+EOF
+  fi
+  cat >> "$out" <<EOF
+
+Queue: $queue_abs
+Evidence: $evidence
+Review output: $review_file
+EOF
+}
+
+write_fix_prompt() {
+  local out=$1
+  local template="$script_dir/prompts/fixer.md"
+  if [[ -f "$template" ]]; then
+    cat "$template" > "$out"
+  else
+    cat > "$out" <<'EOF'
+# Knack Fixer
+
+Load and follow the **fix** skill in `.agents/skills/fix/`.
+Read the structured review artifact and append any actionable fix work units to the existing queue. Stop after updating the queue.
+EOF
+  fi
+  cat >> "$out" <<EOF
+
+Queue: $queue_abs
+Evidence: $evidence
+Review input: $review_file
+EOF
+}
+
+run_phase_agent() {
+  local phase=$1 prompt=$2 output=$3 cmd=${4:-}
+  if [[ -n "$cmd" ]]; then
+    (
+      cd "$repo_dir"
+      LOOP_PHASE="$phase" \
+        LOOP_PROMPT_FILE="$prompt" \
+        LOOP_QUEUE_FILE="$queue_abs" \
+        LOOP_EVIDENCE_FILE="$evidence" \
+        LOOP_REVIEW_FILE="$review_file" \
+        bash -lc "$cmd"
+    ) > "$output" 2>&1
+  else
+    (
+      cd "$repo_dir"
+      LOOP_PHASE="$phase" \
+        LOOP_PROMPT_FILE="$prompt" \
+        LOOP_QUEUE_FILE="$queue_abs" \
+        LOOP_EVIDENCE_FILE="$evidence" \
+        LOOP_REVIEW_FILE="$review_file" \
+        pi -p --no-session --approve "$(cat "$prompt")"
+    ) > "$output" 2>&1
+  fi
+}
+
 [[ $# -ge 1 ]] || { usage; exit 1; }
 cmd=$1
 shift
@@ -215,6 +295,8 @@
 [[ -f "$queue_abs" ]] || die "queue not found: $queue_abs"
 
 max_ticks=3
+max_review_rounds=2
+review_enabled=0
 dry_run=0
 repo_override=""
 while [[ $# -gt 0 ]]; do
@@ -229,6 +311,15 @@
       max_ticks=$2
       shift 2
       ;;
+    --review)
+      review_enabled=1
+      shift
+      ;;
+    --max-review-rounds)
+      [[ $# -ge 2 ]] || die "--max-review-rounds needs a value"
+      max_review_rounds=$2
+      shift 2
+      ;;
     --dry-run)
       dry_run=1
       shift
@@ -252,6 +343,7 @@
 fi
 evidence="$queue_dir/EVIDENCE.md"
 handoff="$queue_dir/HANDOFF.md"
+review_file="$queue_dir/REVIEW.md"
 no_progress_strikes=0
 
 write_handoff_on_exit() {
@@ -263,9 +355,13 @@
 }
 trap write_handoff_on_exit EXIT
 
-for ((tick = 1; tick <= max_ticks; tick++)); do
+tick=1
+review_round=0
+
+while true; do
+while (( tick <= max_ticks )); do
   unit=$(first_pending_unit "$queue_abs")
-  [[ -n "$unit" ]] || { echo "knack: no pending work units"; exit 0; }
+  [[ -n "$unit" ]] || break
 
   unit_file=$(mktemp)
   verify_file=$(mktemp)
@@ -307,11 +403,7 @@
   fi
 
   set +e
-  if [[ -n "$agent_cmd" ]]; then
-    (cd "$repo_dir" && LOOP_PROMPT_FILE="$run_prompt" bash -lc "$agent_cmd") > "$agent_out" 2>&1
-  else
-    (cd "$repo_dir" && pi -p --no-session --approve "$(cat "$run_prompt")") > "$agent_out" 2>&1
-  fi
+  run_phase_agent "build" "$run_prompt" "$agent_out" "$agent_cmd"
   agent_code=$?
   set -e
 
@@ -334,6 +426,7 @@
     set_status "$queue_abs" "$title" "done"
     append_evidence "$evidence" "$title" "done" "$verify" "$verify_out" "$agent_out" "$repo_dir" "$unit_file"
     echo "knack: verified — $title"
+    tick=$((tick + 1))
     continue
   fi
 
@@ -347,6 +440,7 @@
     set_status "$queue_abs" "$title" "pending"
     append_evidence "$evidence" "$title" "verify_failed" "$verify" "$verify_out" "$agent_out" "$repo_dir" "$unit_file"
     echo "knack: verify failed with no progress; retrying once"
+    tick=$((tick + 1))
     continue
   fi
 
@@ -360,4 +454,59 @@
   die "reached max ticks ($max_ticks) with pending work"
 fi
 
-echo "knack: reached max ticks ($max_ticks)"
+if [[ "$review_enabled" == 0 ]]; then
+  echo "knack: no pending work units"
+  exit 0
+fi
+
+if (( review_round >= max_review_rounds )); then
+  die "reached max review rounds ($max_review_rounds)"
+fi
+
+review_round=$((review_round + 1))
+review_prompt=$(mktemp)
+review_out=$(mktemp)
+write_review_prompt "$review_prompt"
+review_cmd="${LOOP_REVIEW_CMD:-${LOOP_AGENT_CMD:-}}"
+
+echo "knack: review round $review_round/$max_review_rounds"
+set +e
+run_phase_agent "review" "$review_prompt" "$review_out" "$review_cmd"
+review_code=$?
+set -e
+if [[ $review_code -ne 0 ]]; then
+  cat "$review_out"
+  die "review worker exited nonzero"
+fi
+[[ -f "$review_file" ]] || die "review worker did not write $review_file"
+actionable=$(extract_actionable_count "$review_file") || die "review file has no actionable summary: $review_file"
+
+if [[ "$actionable" == 0 ]]; then
+  echo "knack: review clean"
+  exit 0
+fi
+
+if (( review_round >= max_review_rounds )); then
+  die "review found $actionable actionable issue(s) at max review rounds ($max_review_rounds)"
+fi
+
+fix_prompt=$(mktemp)
+fix_out=$(mktemp)
+write_fix_prompt "$fix_prompt"
+fix_cmd="${LOOP_FIX_CMD:-${LOOP_AGENT_CMD:-}}"
+
+echo "knack: fix round $review_round/$max_review_rounds — actionable: $actionable"
+set +e
+run_phase_agent "fix" "$fix_prompt" "$fix_out" "$fix_cmd"
+fix_code=$?
+set -e
+if [[ $fix_code -ne 0 ]]; then
+  cat "$fix_out"
+  die "fix worker exited nonzero"
+fi
+
+if [[ -z "$(first_pending_unit "$queue_abs")" ]]; then
+  cat "$fix_out"
+  die "fix worker generated no pending work units"
+fi
+done

codex
The new controller invokes review only after pending build units are exhausted, and fix only when `REVIEW.md` reports actionable findings. I’m adding the regression test now with three local fake workers so the whole path is deterministic.
apply patch
patch: completed
/home/daniel/build/knack/tests/run.sh
diff --git a/loop.sh b/loop.sh
index 351aca3855b1ce6c9c08b0c586ad838fa0f6e636..35066aad26502fbf3f7d5ac0413747bb134cc0b0
--- a/loop.sh
+++ b/loop.sh
@@ -4,7 +4,7 @@
 usage() {
   cat <<'EOF'
 Usage:
-  ./loop.sh run <queue> [--repo DIR] [--max-ticks N] [--dry-run]
+  ./loop.sh run <queue> [--repo DIR] [--max-ticks N] [--review] [--max-review-rounds N] [--dry-run]
 
 The queue is usually .loop/<name>/QUEUE.md in the target repo. Use --repo when the
 queue lives outside the repository it should operate on.
@@ -204,6 +204,86 @@
   } >> "$evidence"
 }
 
+extract_actionable_count() {
+  awk '
+    /^[[:space:]]*-[[:space:]]*actionable:[[:space:]]*[0-9]+[[:space:]]*$/ {
+      sub(/^[[:space:]]*-[[:space:]]*actionable:[[:space:]]*/, "")
+      sub(/[[:space:]]*$/, "")
+      print
+      found = 1
+      exit
+    }
+    END { if (!found) exit 1 }
+  ' "$1"
+}
+
+write_review_prompt() {
+  local out=$1
+  local template="$script_dir/prompts/reviewer.md"
+  if [[ -f "$template" ]]; then
+    cat "$template" > "$out"
+  else
+    cat > "$out" <<'EOF'
+# Knack Reviewer
+
+Load and follow the **review** skill in `.agents/skills/review/`.
+Review the completed queue against the current repository state. Write the structured review artifact at the path provided below.
+EOF
+  fi
+  cat >> "$out" <<EOF
+
+Queue: $queue_abs
+Evidence: $evidence
+Review output: $review_file
+EOF
+}
+
+write_fix_prompt() {
+  local out=$1
+  local template="$script_dir/prompts/fixer.md"
+  if [[ -f "$template" ]]; then
+    cat "$template" > "$out"
+  else
+    cat > "$out" <<'EOF'
+# Knack Fixer
+
+Load and follow the **fix** skill in `.agents/skills/fix/`.
+Read the structured review artifact and append any actionable fix work units to the existing queue. Stop after updating the queue.
+EOF
+  fi
+  cat >> "$out" <<EOF
+
+Queue: $queue_abs
+Evidence: $evidence
+Review input: $review_file
+EOF
+}
+
+run_phase_agent() {
+  local phase=$1 prompt=$2 output=$3 cmd=${4:-}
+  if [[ -n "$cmd" ]]; then
+    (
+      cd "$repo_dir"
+      LOOP_PHASE="$phase" \
+        LOOP_PROMPT_FILE="$prompt" \
+        LOOP_QUEUE_FILE="$queue_abs" \
+        LOOP_EVIDENCE_FILE="$evidence" \
+        LOOP_REVIEW_FILE="$review_file" \
+        bash -lc "$cmd"
+    ) > "$output" 2>&1
+  else
+    (
+      cd "$repo_dir"
+      LOOP_PHASE="$phase" \
+        LOOP_PROMPT_FILE="$prompt" \
+        LOOP_QUEUE_FILE="$queue_abs" \
+        LOOP_EVIDENCE_FILE="$evidence" \
+        LOOP_REVIEW_FILE="$review_file" \
+        pi -p --no-session --approve "$(cat "$prompt")"
+    ) > "$output" 2>&1
+  fi
+}
+
 [[ $# -ge 1 ]] || { usage; exit 1; }
 cmd=$1
 shift
@@ -215,6 +295,8 @@
 [[ -f "$queue_abs" ]] || die "queue not found: $queue_abs"
 
 max_ticks=3
+max_review_rounds=2
+review_enabled=0
 dry_run=0
 repo_override=""
 while [[ $# -gt 0 ]]; do
@@ -229,6 +311,15 @@
       max_ticks=$2
       shift 2
       ;;
+    --review)
+      review_enabled=1
+      shift
+      ;;
+    --max-review-rounds)
+      [[ $# -ge 2 ]] || die "--max-review-rounds needs a value"
+      max_review_rounds=$2
+      shift 2
+      ;;
     --dry-run)
       dry_run=1
       shift
@@ -252,6 +343,7 @@
 fi
 evidence="$queue_dir/EVIDENCE.md"
 handoff="$queue_dir/HANDOFF.md"
+review_file="$queue_dir/REVIEW.md"
 no_progress_strikes=0
 
 write_handoff_on_exit() {
@@ -263,9 +355,13 @@
 }
 trap write_handoff_on_exit EXIT
 
-for ((tick = 1; tick <= max_ticks; tick++)); do
+tick=1
+review_round=0
+
+while true; do
+while (( tick <= max_ticks )); do
   unit=$(first_pending_unit "$queue_abs")
-  [[ -n "$unit" ]] || { echo "knack: no pending work units"; exit 0; }
+  [[ -n "$unit" ]] || break
 
   unit_file=$(mktemp)
   verify_file=$(mktemp)
@@ -307,11 +403,7 @@
   fi
 
   set +e
-  if [[ -n "$agent_cmd" ]]; then
-    (cd "$repo_dir" && LOOP_PROMPT_FILE="$run_prompt" bash -lc "$agent_cmd") > "$agent_out" 2>&1
-  else
-    (cd "$repo_dir" && pi -p --no-session --approve "$(cat "$run_prompt")") > "$agent_out" 2>&1
-  fi
+  run_phase_agent "build" "$run_prompt" "$agent_out" "$agent_cmd"
   agent_code=$?
   set -e
 
@@ -334,6 +426,7 @@
     set_status "$queue_abs" "$title" "done"
     append_evidence "$evidence" "$title" "done" "$verify" "$verify_out" "$agent_out" "$repo_dir" "$unit_file"
     echo "knack: verified — $title"
+    tick=$((tick + 1))
     continue
   fi
 
@@ -347,6 +440,7 @@
     set_status "$queue_abs" "$title" "pending"
     append_evidence "$evidence" "$title" "verify_failed" "$verify" "$verify_out" "$agent_out" "$repo_dir" "$unit_file"
     echo "knack: verify failed with no progress; retrying once"
+    tick=$((tick + 1))
     continue
   fi
 
@@ -360,4 +454,59 @@
   die "reached max ticks ($max_ticks) with pending work"
 fi
 
-echo "knack: reached max ticks ($max_ticks)"
+if [[ "$review_enabled" == 0 ]]; then
+  echo "knack: no pending work units"
+  exit 0
+fi
+
+if (( review_round >= max_review_rounds )); then
+  die "reached max review rounds ($max_review_rounds)"
+fi
+
+review_round=$((review_round + 1))
+review_prompt=$(mktemp)
+review_out=$(mktemp)
+write_review_prompt "$review_prompt"
+review_cmd="${LOOP_REVIEW_CMD:-${LOOP_AGENT_CMD:-}}"
+
+echo "knack: review round $review_round/$max_review_rounds"
+set +e
+run_phase_agent "review" "$review_prompt" "$review_out" "$review_cmd"
+review_code=$?
+set -e
+if [[ $review_code -ne 0 ]]; then
+  cat "$review_out"
+  die "review worker exited nonzero"
+fi
+[[ -f "$review_file" ]] || die "review worker did not write $review_file"
+actionable=$(extract_actionable_count "$review_file") || die "review file has no actionable summary: $review_file"
+
+if [[ "$actionable" == 0 ]]; then
+  echo "knack: review clean"
+  exit 0
+fi
+
+if (( review_round >= max_review_rounds )); then
+  die "review found $actionable actionable issue(s) at max review rounds ($max_review_rounds)"
+fi
+
+fix_prompt=$(mktemp)
+fix_out=$(mktemp)
+write_fix_prompt "$fix_prompt"
+fix_cmd="${LOOP_FIX_CMD:-${LOOP_AGENT_CMD:-}}"
+
+echo "knack: fix round $review_round/$max_review_rounds — actionable: $actionable"
+set +e
+run_phase_agent "fix" "$fix_prompt" "$fix_out" "$fix_cmd"
+fix_code=$?
+set -e
+if [[ $fix_code -ne 0 ]]; then
+  cat "$fix_out"
+  die "fix worker exited nonzero"
+fi
+
+if [[ -z "$(first_pending_unit "$queue_abs")" ]]; then
+  cat "$fix_out"
+  die "fix worker generated no pending work units"
+fi
+done
diff --git a/tests/run.sh b/tests/run.sh
index eb28d49f637f95de4c6e555bb98b1fe43455e7f2..7a28bc257ecfc567e30366b85efdea03ea9a59c3
--- a/tests/run.sh
+++ b/tests/run.sh
@@ -169,6 +169,123 @@
 test -f "$repo_lpf/captured-prompt.txt"
 assert_contains "$repo_lpf/captured-prompt.txt" "the test fixture reaches its verify condition"
 
+# Review-fix loop with fake build, review, and fix workers.
+repo_review="$tmp/repo-review"
+mkdir -p "$repo_review/.loop"
+cat > "$repo_review/.loop/QUEUE.md" <<'EOF'
+# Loop Queue: review cycle
+
+Goal:
+Exercise build, review, fix, and review again.
+
+Stop condition:
+The generated app is fixed and review reports no actionable issues.
+
+## the initial build creates a reviewable app file
+
+Read first:
+- This queue file.
+
+Constraints:
+- Leave the bug for review to find.
+
+Verify:
+```bash
+test -f app.txt
+```
+
+Done means:
+- app.txt exists.
+
+Status: pending
+EOF
+
+cat > "$repo_review/build-worker.sh" <<'EOF'
+#!/usr/bin/env bash
+set -euo pipefail
+if grep -q "the fix unit repairs the bug" "$LOOP_PROMPT_FILE"; then
+  printf 'fixed\n' > app.txt
+  echo "build fixed app"
+else
+  printf 'bug\n' > app.txt
+  echo "build created app with bug"
+fi
+EOF
+chmod +x "$repo_review/build-worker.sh"
+
+cat > "$repo_review/review-worker.sh" <<'EOF'
+#!/usr/bin/env bash
+set -euo pipefail
+count=0
+if [[ -f review-count.txt ]]; then
+  count=$(cat review-count.txt)
+fi
+count=$((count + 1))
+printf '%s\n' "$count" > review-count.txt
+
+actionable=1
+if [[ -f app.txt ]] && grep -qx 'fixed' app.txt; then
+  actionable=0
+fi
+
+cat > "$LOOP_REVIEW_FILE" <<EOF_REVIEW
+# Review: fake
+
+## Standards
+
+## Intent
+- actionable | high — app.txt must say fixed
+  evidence: app.txt:1
+
+## Speculative
+
+## Summary
+- actionable: $actionable
+- trivial: 0
+- disputed: 0
+- deferred: 0
+EOF_REVIEW
+echo "review actionable: $actionable"
+EOF
+chmod +x "$repo_review/review-worker.sh"
+
+cat > "$repo_review/fix-worker.sh" <<'EOF'
+#!/usr/bin/env bash
+set -euo pipefail
+cat >> "$LOOP_QUEUE_FILE" <<'EOF_QUEUE'
+
+## the fix unit repairs the bug
+
+Read first:
+- .loop/REVIEW.md
+
+Constraints:
+- Preserve the app file created by the first unit.
+
+Verify:
+```bash
+grep -qx fixed app.txt
+```
+
+Done means:
+- app.txt contains fixed.
+
+Status: pending
+EOF_QUEUE
+echo "fix appended unit"
+EOF
+chmod +x "$repo_review/fix-worker.sh"
+
+LOOP_AGENT_CMD="$repo_review/build-worker.sh" \
+  LOOP_REVIEW_CMD="$repo_review/review-worker.sh" \
+  LOOP_FIX_CMD="$repo_review/fix-worker.sh" \
+  "$root/loop.sh" run "$repo_review/.loop/QUEUE.md" --review --max-ticks 2 >/tmp/loop-review.txt
+assert_contains "$repo_review/.loop/QUEUE.md" "## the fix unit repairs the bug"
+assert_contains "$repo_review/.loop/QUEUE.md" "Status: done"
+assert_contains "$repo_review/.loop/REVIEW.md" "- actionable: 0"
+assert_contains "$repo_review/review-count.txt" "2"
+assert_contains "$repo_review/app.txt" "fixed"
+
 if command -v skills-ref >/dev/null 2>&1; then
   for skill_dir in "$root/.agents/skills"/*; do
     if [[ -d "$skill_dir" ]]; then

diff --git a/loop.sh b/loop.sh
index 351aca3855b1ce6c9c08b0c586ad838fa0f6e636..35066aad26502fbf3f7d5ac0413747bb134cc0b0
--- a/loop.sh
+++ b/loop.sh
@@ -4,7 +4,7 @@
 usage() {
   cat <<'EOF'
 Usage:
-  ./loop.sh run <queue> [--repo DIR] [--max-ticks N] [--dry-run]
+  ./loop.sh run <queue> [--repo DIR] [--max-ticks N] [--review] [--max-review-rounds N] [--dry-run]
 
 The queue is usually .loop/<name>/QUEUE.md in the target repo. Use --repo when the
 queue lives outside the repository it should operate on.
@@ -204,6 +204,86 @@
   } >> "$evidence"
 }
 
+extract_actionable_count() {
+  awk '
+    /^[[:space:]]*-[[:space:]]*actionable:[[:space:]]*[0-9]+[[:space:]]*$/ {
+      sub(/^[[:space:]]*-[[:space:]]*actionable:[[:space:]]*/, "")
+      sub(/[[:space:]]*$/, "")
+      print
+      found = 1
+      exit
+    }
+    END { if (!found) exit 1 }
+  ' "$1"
+}
+
+write_review_prompt() {
+  local out=$1
+  local template="$script_dir/prompts/reviewer.md"
+  if [[ -f "$template" ]]; then
+    cat "$template" > "$out"
+  else
+    cat > "$out" <<'EOF'
+# Knack Reviewer
+
+Load and follow the **review** skill in `.agents/skills/review/`.
+Review the completed queue against the current repository state. Write the structured review artifact at the path provided below.
+EOF
+  fi
+  cat >> "$out" <<EOF
+
+Queue: $queue_abs
+Evidence: $evidence
+Review output: $review_file
+EOF
+}
+
+write_fix_prompt() {
+  local out=$1
+  local template="$script_dir/prompts/fixer.md"
+  if [[ -f "$template" ]]; then
+    cat "$template" > "$out"
+  else
+    cat > "$out" <<'EOF'
+# Knack Fixer
+
+Load and follow the **fix** skill in `.agents/skills/fix/`.
+Read the structured review artifact and append any actionable fix work units to the existing queue. Stop after updating the queue.
+EOF
+  fi
+  cat >> "$out" <<EOF
+
+Queue: $queue_abs
+Evidence: $evidence
+Review input: $review_file
+EOF
+}
+
+run_phase_agent() {
+  local phase=$1 prompt=$2 output=$3 cmd=${4:-}
+  if [[ -n "$cmd" ]]; then
+    (
+      cd "$repo_dir"
+      LOOP_PHASE="$phase" \
+        LOOP_PROMPT_FILE="$prompt" \
+        LOOP_QUEUE_FILE="$queue_abs" \
+        LOOP_EVIDENCE_FILE="$evidence" \
+        LOOP_REVIEW_FILE="$review_file" \
+        bash -lc "$cmd"
+    ) > "$output" 2>&1
+  else
+    (
+      cd "$repo_dir"
+      LOOP_PHASE="$phase" \
+        LOOP_PROMPT_FILE="$prompt" \
+        LOOP_QUEUE_FILE="$queue_abs" \
+        LOOP_EVIDENCE_FILE="$evidence" \
+        LOOP_REVIEW_FILE="$review_file" \
+        pi -p --no-session --approve "$(cat "$prompt")"
+    ) > "$output" 2>&1
+  fi
+}
+
 [[ $# -ge 1 ]] || { usage; exit 1; }
 cmd=$1
 shift
@@ -215,6 +295,8 @@
 [[ -f "$queue_abs" ]] || die "queue not found: $queue_abs"
 
 max_ticks=3
+max_review_rounds=2
+review_enabled=0
 dry_run=0
 repo_override=""
 while [[ $# -gt 0 ]]; do
@@ -229,6 +311,15 @@
       max_ticks=$2
       shift 2
       ;;
+    --review)
+      review_enabled=1
+      shift
+      ;;
+    --max-review-rounds)
+      [[ $# -ge 2 ]] || die "--max-review-rounds needs a value"
+      max_review_rounds=$2
+      shift 2
+      ;;
     --dry-run)
       dry_run=1
       shift
@@ -252,6 +343,7 @@
 fi
 evidence="$queue_dir/EVIDENCE.md"
 handoff="$queue_dir/HANDOFF.md"
+review_file="$queue_dir/REVIEW.md"
 no_progress_strikes=0
 
 write_handoff_on_exit() {
@@ -263,9 +355,13 @@
 }
 trap write_handoff_on_exit EXIT
 
-for ((tick = 1; tick <= max_ticks; tick++)); do
+tick=1
+review_round=0
+
+while true; do
+while (( tick <= max_ticks )); do
   unit=$(first_pending_unit "$queue_abs")
-  [[ -n "$unit" ]] || { echo "knack: no pending work units"; exit 0; }
+  [[ -n "$unit" ]] || break
 
   unit_file=$(mktemp)
   verify_file=$(mktemp)
@@ -307,11 +403,7 @@
   fi
 
   set +e
-  if [[ -n "$agent_cmd" ]]; then
-    (cd "$repo_dir" && LOOP_PROMPT_FILE="$run_prompt" bash -lc "$agent_cmd") > "$agent_out" 2>&1
-  else
-    (cd "$repo_dir" && pi -p --no-session --approve "$(cat "$run_prompt")") > "$agent_out" 2>&1
-  fi
+  run_phase_agent "build" "$run_prompt" "$agent_out" "$agent_cmd"
   agent_code=$?
   set -e
 
@@ -334,6 +426,7 @@
     set_status "$queue_abs" "$title" "done"
     append_evidence "$evidence" "$title" "done" "$verify" "$verify_out" "$agent_out" "$repo_dir" "$unit_file"
     echo "knack: verified — $title"
+    tick=$((tick + 1))
     continue
   fi
 
@@ -347,6 +440,7 @@
     set_status "$queue_abs" "$title" "pending"
     append_evidence "$evidence" "$title" "verify_failed" "$verify" "$verify_out" "$agent_out" "$repo_dir" "$unit_file"
     echo "knack: verify failed with no progress; retrying once"
+    tick=$((tick + 1))
     continue
   fi
 
@@ -360,4 +454,59 @@
   die "reached max ticks ($max_ticks) with pending work"
 fi
 
-echo "knack: reached max ticks ($max_ticks)"
+if [[ "$review_enabled" == 0 ]]; then
+  echo "knack: no pending work units"
+  exit 0
+fi
+
+if (( review_round >= max_review_rounds )); then
+  die "reached max review rounds ($max_review_rounds)"
+fi
+
+review_round=$((review_round + 1))
+review_prompt=$(mktemp)
+review_out=$(mktemp)
+write_review_prompt "$review_prompt"
+review_cmd="${LOOP_REVIEW_CMD:-${LOOP_AGENT_CMD:-}}"
+
+echo "knack: review round $review_round/$max_review_rounds"
+set +e
+run_phase_agent "review" "$review_prompt" "$review_out" "$review_cmd"
+review_code=$?
+set -e
+if [[ $review_code -ne 0 ]]; then
+  cat "$review_out"
+  die "review worker exited nonzero"
+fi
+[[ -f "$review_file" ]] || die "review worker did not write $review_file"
+actionable=$(extract_actionable_count "$review_file") || die "review file has no actionable summary: $review_file"
+
+if [[ "$actionable" == 0 ]]; then
+  echo "knack: review clean"
+  exit 0
+fi
+
+if (( review_round >= max_review_rounds )); then
+  die "review found $actionable actionable issue(s) at max review rounds ($max_review_rounds)"
+fi
+
+fix_prompt=$(mktemp)
+fix_out=$(mktemp)
+write_fix_prompt "$fix_prompt"
+fix_cmd="${LOOP_FIX_CMD:-${LOOP_AGENT_CMD:-}}"
+
+echo "knack: fix round $review_round/$max_review_rounds — actionable: $actionable"
+set +e
+run_phase_agent "fix" "$fix_prompt" "$fix_out" "$fix_cmd"
+fix_code=$?
+set -e
+if [[ $fix_code -ne 0 ]]; then
+  cat "$fix_out"
+  die "fix worker exited nonzero"
+fi
+
+if [[ -z "$(first_pending_unit "$queue_abs")" ]]; then
+  cat "$fix_out"
+  die "fix worker generated no pending work units"
+fi
+done
diff --git a/tests/run.sh b/tests/run.sh
index eb28d49f637f95de4c6e555bb98b1fe43455e7f2..7a28bc257ecfc567e30366b85efdea03ea9a59c3
--- a/tests/run.sh
+++ b/tests/run.sh
@@ -169,6 +169,123 @@
 test -f "$repo_lpf/captured-prompt.txt"
 assert_contains "$repo_lpf/captured-prompt.txt" "the test fixture reaches its verify condition"
 
+# Review-fix loop with fake build, review, and fix workers.
+repo_review="$tmp/repo-review"
+mkdir -p "$repo_review/.loop"
+cat > "$repo_review/.loop/QUEUE.md" <<'EOF'
+# Loop Queue: review cycle
+
+Goal:
+Exercise build, review, fix, and review again.
+
+Stop condition:
+The generated app is fixed and review reports no actionable issues.
+
+## the initial build creates a reviewable app file
+
+Read first:
+- This queue file.
+
+Constraints:
+- Leave the bug for review to find.
+
+Verify:
+```bash
+test -f app.txt
+```
+
+Done means:
+- app.txt exists.
+
+Status: pending
+EOF
+
+cat > "$repo_review/build-worker.sh" <<'EOF'
+#!/usr/bin/env bash
+set -euo pipefail
+if grep -q "the fix unit repairs the bug" "$LOOP_PROMPT_FILE"; then
+  printf 'fixed\n' > app.txt
+  echo "build fixed app"
+else
+  printf 'bug\n' > app.txt
+  echo "build created app with bug"
+fi
+EOF
+chmod +x "$repo_review/build-worker.sh"
+
+cat > "$repo_review/review-worker.sh" <<'EOF'
+#!/usr/bin/env bash
+set -euo pipefail
+count=0
+if [[ -f review-count.txt ]]; then
+  count=$(cat review-count.txt)
+fi
+count=$((count + 1))
+printf '%s\n' "$count" > review-count.txt
+
+actionable=1
+if [[ -f app.txt ]] && grep -qx 'fixed' app.txt; then
+  actionable=0
+fi
+
+cat > "$LOOP_REVIEW_FILE" <<EOF_REVIEW
+# Review: fake
+
+## Standards
+
+## Intent
+- actionable | high — app.txt must say fixed
+  evidence: app.txt:1
+
+## Speculative
+
+## Summary
+- actionable: $actionable
+- trivial: 0
+- disputed: 0
+- deferred: 0
+EOF_REVIEW
+echo "review actionable: $actionable"
+EOF
+chmod +x "$repo_review/review-worker.sh"
+
+cat > "$repo_review/fix-worker.sh" <<'EOF'
+#!/usr/bin/env bash
+set -euo pipefail
+cat >> "$LOOP_QUEUE_FILE" <<'EOF_QUEUE'
+
+## the fix unit repairs the bug
+
+Read first:
+- .loop/REVIEW.md
+
+Constraints:
+- Preserve the app file created by the first unit.
+
+Verify:
+```bash
+grep -qx fixed app.txt
+```
+
+Done means:
+- app.txt contains fixed.
+
+Status: pending
+EOF_QUEUE
+echo "fix appended unit"
+EOF
+chmod +x "$repo_review/fix-worker.sh"
+
+LOOP_AGENT_CMD="$repo_review/build-worker.sh" \
+  LOOP_REVIEW_CMD="$repo_review/review-worker.sh" \
+  LOOP_FIX_CMD="$repo_review/fix-worker.sh" \
+  "$root/loop.sh" run "$repo_review/.loop/QUEUE.md" --review --max-ticks 2 >/tmp/loop-review.txt
+assert_contains "$repo_review/.loop/QUEUE.md" "## the fix unit repairs the bug"
+assert_contains "$repo_review/.loop/QUEUE.md" "Status: done"
+assert_contains "$repo_review/.loop/REVIEW.md" "- actionable: 0"
+assert_contains "$repo_review/review-count.txt" "2"
+assert_contains "$repo_review/app.txt" "fixed"
+
 if command -v skills-ref >/dev/null 2>&1; then
   for skill_dir in "$root/.agents/skills"/*; do
     if [[ -d "$skill_dir" ]]; then

codex
The fake cycle test now covers build creating a defect, review reporting one actionable item, fix appending a new pending unit, build fixing it, and review returning clean. I’m running syntax and the project test harness next.
exec
/usr/bin/bash -lc 'bash -n loop.sh' in /home/daniel/build/knack
 succeeded in 0ms:
exec
/usr/bin/bash -lc 'bash -n tests/run.sh' in /home/daniel/build/knack
 succeeded in 0ms:
diff --git a/loop.sh b/loop.sh
index 351aca3855b1ce6c9c08b0c586ad838fa0f6e636..35066aad26502fbf3f7d5ac0413747bb134cc0b0
--- a/loop.sh
+++ b/loop.sh
@@ -4,7 +4,7 @@
 usage() {
   cat <<'EOF'
 Usage:
-  ./loop.sh run <queue> [--repo DIR] [--max-ticks N] [--dry-run]
+  ./loop.sh run <queue> [--repo DIR] [--max-ticks N] [--review] [--max-review-rounds N] [--dry-run]
 
 The queue is usually .loop/<name>/QUEUE.md in the target repo. Use --repo when the
 queue lives outside the repository it should operate on.
@@ -204,6 +204,86 @@
   } >> "$evidence"
 }
 
+extract_actionable_count() {
+  awk '
+    /^[[:space:]]*-[[:space:]]*actionable:[[:space:]]*[0-9]+[[:space:]]*$/ {
+      sub(/^[[:space:]]*-[[:space:]]*actionable:[[:space:]]*/, "")
+      sub(/[[:space:]]*$/, "")
+      print
+      found = 1
+      exit
+    }
+    END { if (!found) exit 1 }
+  ' "$1"
+}
+
+write_review_prompt() {
+  local out=$1
+  local template="$script_dir/prompts/reviewer.md"
+  if [[ -f "$template" ]]; then
+    cat "$template" > "$out"
+  else
+    cat > "$out" <<'EOF'
+# Knack Reviewer
+
+Load and follow the **review** skill in `.agents/skills/review/`.
+Review the completed queue against the current repository state. Write the structured review artifact at the path provided below.
+EOF
+  fi
+  cat >> "$out" <<EOF
+
+Queue: $queue_abs
+Evidence: $evidence
+Review output: $review_file
+EOF
+}
+
+write_fix_prompt() {
+  local out=$1
+  local template="$script_dir/prompts/fixer.md"
+  if [[ -f "$template" ]]; then
+    cat "$template" > "$out"
+  else
+    cat > "$out" <<'EOF'
+# Knack Fixer
+
+Load and follow the **fix** skill in `.agents/skills/fix/`.
+Read the structured review artifact and append any actionable fix work units to the existing queue. Stop after updating the queue.
+EOF
+  fi
+  cat >> "$out" <<EOF
+
+Queue: $queue_abs
+Evidence: $evidence
+Review input: $review_file
+EOF
+}
+
+run_phase_agent() {
+  local phase=$1 prompt=$2 output=$3 cmd=${4:-}
+  if [[ -n "$cmd" ]]; then
+    (
+      cd "$repo_dir"
+      LOOP_PHASE="$phase" \
+        LOOP_PROMPT_FILE="$prompt" \
+        LOOP_QUEUE_FILE="$queue_abs" \
+        LOOP_EVIDENCE_FILE="$evidence" \
+        LOOP_REVIEW_FILE="$review_file" \
+        bash -lc "$cmd"
+    ) > "$output" 2>&1
+  else
+    (
+      cd "$repo_dir"
+      LOOP_PHASE="$phase" \
+        LOOP_PROMPT_FILE="$prompt" \
+        LOOP_QUEUE_FILE="$queue_abs" \
+        LOOP_EVIDENCE_FILE="$evidence" \
+        LOOP_REVIEW_FILE="$review_file" \
+        pi -p --no-session --approve "$(cat "$prompt")"
+    ) > "$output" 2>&1
+  fi
+}
+
 [[ $# -ge 1 ]] || { usage; exit 1; }
 cmd=$1
 shift
@@ -215,6 +295,8 @@
 [[ -f "$queue_abs" ]] || die "queue not found: $queue_abs"
 
 max_ticks=3
+max_review_rounds=2
+review_enabled=0
 dry_run=0
 repo_override=""
 while [[ $# -gt 0 ]]; do
@@ -229,6 +311,15 @@
       max_ticks=$2
       shift 2
       ;;
+    --review)
+      review_enabled=1
+      shift
+      ;;
+    --max-review-rounds)
+      [[ $# -ge 2 ]] || die "--max-review-rounds needs a value"
+      max_review_rounds=$2
+      shift 2
+      ;;
     --dry-run)
       dry_run=1
       shift
@@ -252,6 +343,7 @@
 fi
 evidence="$queue_dir/EVIDENCE.md"
 handoff="$queue_dir/HANDOFF.md"
+review_file="$queue_dir/REVIEW.md"
 no_progress_strikes=0
 
 write_handoff_on_exit() {
@@ -263,9 +355,13 @@
 }
 trap write_handoff_on_exit EXIT
 
-for ((tick = 1; tick <= max_ticks; tick++)); do
+tick=1
+review_round=0
+
+while true; do
+while (( tick <= max_ticks )); do
   unit=$(first_pending_unit "$queue_abs")
-  [[ -n "$unit" ]] || { echo "knack: no pending work units"; exit 0; }
+  [[ -n "$unit" ]] || break
 
   unit_file=$(mktemp)
   verify_file=$(mktemp)
@@ -307,11 +403,7 @@
   fi
 
   set +e
-  if [[ -n "$agent_cmd" ]]; then
-    (cd "$repo_dir" && LOOP_PROMPT_FILE="$run_prompt" bash -lc "$agent_cmd") > "$agent_out" 2>&1
-  else
-    (cd "$repo_dir" && pi -p --no-session --approve "$(cat "$run_prompt")") > "$agent_out" 2>&1
-  fi
+  run_phase_agent "build" "$run_prompt" "$agent_out" "$agent_cmd"
   agent_code=$?
   set -e
 
@@ -334,6 +426,7 @@
     set_status "$queue_abs" "$title" "done"
     append_evidence "$evidence" "$title" "done" "$verify" "$verify_out" "$agent_out" "$repo_dir" "$unit_file"
     echo "knack: verified — $title"
+    tick=$((tick + 1))
     continue
   fi
 
@@ -347,6 +440,7 @@
     set_status "$queue_abs" "$title" "pending"
     append_evidence "$evidence" "$title" "verify_failed" "$verify" "$verify_out" "$agent_out" "$repo_dir" "$unit_file"
     echo "knack: verify failed with no progress; retrying once"
+    tick=$((tick + 1))
     continue
   fi
 
@@ -360,4 +454,59 @@
   die "reached max ticks ($max_ticks) with pending work"
 fi
 
-echo "knack: reached max ticks ($max_ticks)"
+if [[ "$review_enabled" == 0 ]]; then
+  echo "knack: no pending work units"
+  exit 0
+fi
+
+if (( review_round >= max_review_rounds )); then
+  die "reached max review rounds ($max_review_rounds)"
+fi
+
+review_round=$((review_round + 1))
+review_prompt=$(mktemp)
+review_out=$(mktemp)
+write_review_prompt "$review_prompt"
+review_cmd="${LOOP_REVIEW_CMD:-${LOOP_AGENT_CMD:-}}"
+
+echo "knack: review round $review_round/$max_review_rounds"
+set +e
+run_phase_agent "review" "$review_prompt" "$review_out" "$review_cmd"
+review_code=$?
+set -e
+if [[ $review_code -ne 0 ]]; then
+  cat "$review_out"
+  die "review worker exited nonzero"
+fi
+[[ -f "$review_file" ]] || die "review worker did not write $review_file"
+actionable=$(extract_actionable_count "$review_file") || die "review file has no actionable summary: $review_file"
+
+if [[ "$actionable" == 0 ]]; then
+  echo "knack: review clean"
+  exit 0
+fi
+
+if (( review_round >= max_review_rounds )); then
+  die "review found $actionable actionable issue(s) at max review rounds ($max_review_rounds)"
+fi
+
+fix_prompt=$(mktemp)
+fix_out=$(mktemp)
+write_fix_prompt "$fix_prompt"
+fix_cmd="${LOOP_FIX_CMD:-${LOOP_AGENT_CMD:-}}"
+
+echo "knack: fix round $review_round/$max_review_rounds — actionable: $actionable"
+set +e
+run_phase_agent "fix" "$fix_prompt" "$fix_out" "$fix_cmd"
+fix_code=$?
+set -e
+if [[ $fix_code -ne 0 ]]; then
+  cat "$fix_out"
+  die "fix worker exited nonzero"
+fi
+
+if [[ -z "$(first_pending_unit "$queue_abs")" ]]; then
+  cat "$fix_out"
+  die "fix worker generated no pending work units"
+fi
+done
diff --git a/tests/run.sh b/tests/run.sh
index eb28d49f637f95de4c6e555bb98b1fe43455e7f2..7a28bc257ecfc567e30366b85efdea03ea9a59c3
--- a/tests/run.sh
+++ b/tests/run.sh
@@ -169,6 +169,123 @@
 test -f "$repo_lpf/captured-prompt.txt"
 assert_contains "$repo_lpf/captured-prompt.txt" "the test fixture reaches its verify condition"
 
+# Review-fix loop with fake build, review, and fix workers.
+repo_review="$tmp/repo-review"
+mkdir -p "$repo_review/.loop"
+cat > "$repo_review/.loop/QUEUE.md" <<'EOF'
+# Loop Queue: review cycle
+
+Goal:
+Exercise build, review, fix, and review again.
+
+Stop condition:
+The generated app is fixed and review reports no actionable issues.
+
+## the initial build creates a reviewable app file
+
+Read first:
+- This queue file.
+
+Constraints:
+- Leave the bug for review to find.
+
+Verify:
+```bash
+test -f app.txt
+```
+
+Done means:
+- app.txt exists.
+
+Status: pending
+EOF
+
+cat > "$repo_review/build-worker.sh" <<'EOF'
+#!/usr/bin/env bash
+set -euo pipefail
+if grep -q "the fix unit repairs the bug" "$LOOP_PROMPT_FILE"; then
+  printf 'fixed\n' > app.txt
+  echo "build fixed app"
+else
+  printf 'bug\n' > app.txt
+  echo "build created app with bug"
+fi
+EOF
+chmod +x "$repo_review/build-worker.sh"
+
+cat > "$repo_review/review-worker.sh" <<'EOF'
+#!/usr/bin/env bash
+set -euo pipefail
+count=0
+if [[ -f review-count.txt ]]; then
+  count=$(cat review-count.txt)
+fi
+count=$((count + 1))
+printf '%s\n' "$count" > review-count.txt
+
+actionable=1
+if [[ -f app.txt ]] && grep -qx 'fixed' app.txt; then
+  actionable=0
+fi
+
+cat > "$LOOP_REVIEW_FILE" <<EOF_REVIEW
+# Review: fake
+
+## Standards
+
+## Intent
+- actionable | high — app.txt must say fixed
+  evidence: app.txt:1
+
+## Speculative
+
+## Summary
+- actionable: $actionable
+- trivial: 0
+- disputed: 0
+- deferred: 0
+EOF_REVIEW
+echo "review actionable: $actionable"
+EOF
+chmod +x "$repo_review/review-worker.sh"
+
+cat > "$repo_review/fix-worker.sh" <<'EOF'
+#!/usr/bin/env bash
+set -euo pipefail
+cat >> "$LOOP_QUEUE_FILE" <<'EOF_QUEUE'
+
+## the fix unit repairs the bug
+
+Read first:
+- .loop/REVIEW.md
+
+Constraints:
+- Preserve the app file created by the first unit.
+
+Verify:
+```bash
+grep -qx fixed app.txt
+```
+
+Done means:
+- app.txt contains fixed.
+
+Status: pending
+EOF_QUEUE
+echo "fix appended unit"
+EOF
+chmod +x "$repo_review/fix-worker.sh"
+
+LOOP_AGENT_CMD="$repo_review/build-worker.sh" \
+  LOOP_REVIEW_CMD="$repo_review/review-worker.sh" \
+  LOOP_FIX_CMD="$repo_review/fix-worker.sh" \
+  "$root/loop.sh" run "$repo_review/.loop/QUEUE.md" --review --max-ticks 2 >/tmp/loop-review.txt
+assert_contains "$repo_review/.loop/QUEUE.md" "## the fix unit repairs the bug"
+assert_contains "$repo_review/.loop/QUEUE.md" "Status: done"
+assert_contains "$repo_review/.loop/REVIEW.md" "- actionable: 0"
+assert_contains "$repo_review/review-count.txt" "2"
+assert_contains "$repo_review/app.txt" "fixed"
+
 if command -v skills-ref >/dev/null 2>&1; then
   for skill_dir in "$root/.agents/skills"/*; do
     if [[ -d "$skill_dir" ]]; then

codex
Shell syntax is clean. I’m running the full verify command now; if it exposes behavior around the new stop conditions or legacy tests, I’ll tighten that before handing off.
exec
/usr/bin/bash -lc 'bash -n /home/daniel/build/knack/loop.sh && /home/daniel/build/knack/tests/run.sh' in /home/daniel/build/knack
diff --git a/loop.sh b/loop.sh
index 351aca3855b1ce6c9c08b0c586ad838fa0f6e636..35066aad26502fbf3f7d5ac0413747bb134cc0b0
--- a/loop.sh
+++ b/loop.sh
@@ -4,7 +4,7 @@
 usage() {
   cat <<'EOF'
 Usage:
-  ./loop.sh run <queue> [--repo DIR] [--max-ticks N] [--dry-run]
+  ./loop.sh run <queue> [--repo DIR] [--max-ticks N] [--review] [--max-review-rounds N] [--dry-run]
 
 The queue is usually .loop/<name>/QUEUE.md in the target repo. Use --repo when the
 queue lives outside the repository it should operate on.
@@ -204,6 +204,86 @@
   } >> "$evidence"
 }
 
+extract_actionable_count() {
+  awk '
+    /^[[:space:]]*-[[:space:]]*actionable:[[:space:]]*[0-9]+[[:space:]]*$/ {
+      sub(/^[[:space:]]*-[[:space:]]*actionable:[[:space:]]*/, "")
+      sub(/[[:space:]]*$/, "")
+      print
+      found = 1
+      exit
+    }
+    END { if (!found) exit 1 }
+  ' "$1"
+}
+
+write_review_prompt() {
+  local out=$1
+  local template="$script_dir/prompts/reviewer.md"
+  if [[ -f "$template" ]]; then
+    cat "$template" > "$out"
+  else
+    cat > "$out" <<'EOF'
+# Knack Reviewer
+
+Load and follow the **review** skill in `.agents/skills/review/`.
+Review the completed queue against the current repository state. Write the structured review artifact at the path provided below.
+EOF
+  fi
+  cat >> "$out" <<EOF
+
+Queue: $queue_abs
+Evidence: $evidence
+Review output: $review_file
+EOF
+}
+
+write_fix_prompt() {
+  local out=$1
+  local template="$script_dir/prompts/fixer.md"
+  if [[ -f "$template" ]]; then
+    cat "$template" > "$out"
+  else
+    cat > "$out" <<'EOF'
+# Knack Fixer
+
+Load and follow the **fix** skill in `.agents/skills/fix/`.
+Read the structured review artifact and append any actionable fix work units to the existing queue. Stop after updating the queue.
+EOF
+  fi
+  cat >> "$out" <<EOF
+
+Queue: $queue_abs
+Evidence: $evidence
+Review input: $review_file
+EOF
+}
+
+run_phase_agent() {
+  local phase=$1 prompt=$2 output=$3 cmd=${4:-}
+  if [[ -n "$cmd" ]]; then
+    (
+      cd "$repo_dir"
+      LOOP_PHASE="$phase" \
+        LOOP_PROMPT_FILE="$prompt" \
+        LOOP_QUEUE_FILE="$queue_abs" \
+        LOOP_EVIDENCE_FILE="$evidence" \
+        LOOP_REVIEW_FILE="$review_file" \
+        bash -lc "$cmd"
+    ) > "$output" 2>&1
+  else
+    (
+      cd "$repo_dir"
+      LOOP_PHASE="$phase" \
+        LOOP_PROMPT_FILE="$prompt" \
+        LOOP_QUEUE_FILE="$queue_abs" \
+        LOOP_EVIDENCE_FILE="$evidence" \
+        LOOP_REVIEW_FILE="$review_file" \
+        pi -p --no-session --approve "$(cat "$prompt")"
+    ) > "$output" 2>&1
+  fi
+}
+
 [[ $# -ge 1 ]] || { usage; exit 1; }
 cmd=$1
 shift
@@ -215,6 +295,8 @@
 [[ -f "$queue_abs" ]] || die "queue not found: $queue_abs"
 
 max_ticks=3
+max_review_rounds=2
+review_enabled=0
 dry_run=0
 repo_override=""
 while [[ $# -gt 0 ]]; do
@@ -229,6 +311,15 @@
       max_ticks=$2
       shift 2
       ;;
+    --review)
+      review_enabled=1
+      shift
+      ;;
+    --max-review-rounds)
+      [[ $# -ge 2 ]] || die "--max-review-rounds needs a value"
+      max_review_rounds=$2
+      shift 2
+      ;;
     --dry-run)
       dry_run=1
       shift
@@ -252,6 +343,7 @@
 fi
 evidence="$queue_dir/EVIDENCE.md"
 handoff="$queue_dir/HANDOFF.md"
+review_file="$queue_dir/REVIEW.md"
 no_progress_strikes=0
 
 write_handoff_on_exit() {
@@ -263,9 +355,13 @@
 }
 trap write_handoff_on_exit EXIT
 
-for ((tick = 1; tick <= max_ticks; tick++)); do
+tick=1
+review_round=0
+
+while true; do
+while (( tick <= max_ticks )); do
   unit=$(first_pending_unit "$queue_abs")
-  [[ -n "$unit" ]] || { echo "knack: no pending work units"; exit 0; }
+  [[ -n "$unit" ]] || break
 
   unit_file=$(mktemp)
   verify_file=$(mktemp)
@@ -307,11 +403,7 @@
   fi
 
   set +e
-  if [[ -n "$agent_cmd" ]]; then
-    (cd "$repo_dir" && LOOP_PROMPT_FILE="$run_prompt" bash -lc "$agent_cmd") > "$agent_out" 2>&1
-  else
-    (cd "$repo_dir" && pi -p --no-session --approve "$(cat "$run_prompt")") > "$agent_out" 2>&1
-  fi
+  run_phase_agent "build" "$run_prompt" "$agent_out" "$agent_cmd"
   agent_code=$?
   set -e
 
@@ -334,6 +426,7 @@
     set_status "$queue_abs" "$title" "done"
     append_evidence "$evidence" "$title" "done" "$verify" "$verify_out" "$agent_out" "$repo_dir" "$unit_file"
     echo "knack: verified — $title"
+    tick=$((tick + 1))
     continue
   fi
 
@@ -347,6 +440,7 @@
     set_status "$queue_abs" "$title" "pending"
     append_evidence "$evidence" "$title" "verify_failed" "$verify" "$verify_out" "$agent_out" "$repo_dir" "$unit_file"
     echo "knack: verify failed with no progress; retrying once"
+    tick=$((tick + 1))
     continue
   fi
 
@@ -360,4 +454,59 @@
   die "reached max ticks ($max_ticks) with pending work"
 fi
 
-echo "knack: reached max ticks ($max_ticks)"
+if [[ "$review_enabled" == 0 ]]; then
+  echo "knack: no pending work units"
+  exit 0
+fi
+
+if (( review_round >= max_review_rounds )); then
+  die "reached max review rounds ($max_review_rounds)"
+fi
+
+review_round=$((review_round + 1))
+review_prompt=$(mktemp)
+review_out=$(mktemp)
+write_review_prompt "$review_prompt"
+review_cmd="${LOOP_REVIEW_CMD:-${LOOP_AGENT_CMD:-}}"
+
+echo "knack: review round $review_round/$max_review_rounds"
+set +e
+run_phase_agent "review" "$review_prompt" "$review_out" "$review_cmd"
+review_code=$?
+set -e
+if [[ $review_code -ne 0 ]]; then
+  cat "$review_out"
+  die "review worker exited nonzero"
+fi
+[[ -f "$review_file" ]] || die "review worker did not write $review_file"
+actionable=$(extract_actionable_count "$review_file") || die "review file has no actionable summary: $review_file"
+
+if [[ "$actionable" == 0 ]]; then
+  echo "knack: review clean"
+  exit 0
+fi
+
+if (( review_round >= max_review_rounds )); then
+  die "review found $actionable actionable issue(s) at max review rounds ($max_review_rounds)"
+fi
+
+fix_prompt=$(mktemp)
+fix_out=$(mktemp)
+write_fix_prompt "$fix_prompt"
+fix_cmd="${LOOP_FIX_CMD:-${LOOP_AGENT_CMD:-}}"
+
+echo "knack: fix round $review_round/$max_review_rounds — actionable: $actionable"
+set +e
+run_phase_agent "fix" "$fix_prompt" "$fix_out" "$fix_cmd"
+fix_code=$?
+set -e
+if [[ $fix_code -ne 0 ]]; then
+  cat "$fix_out"
+  die "fix worker exited nonzero"
+fi
+
+if [[ -z "$(first_pending_unit "$queue_abs")" ]]; then
+  cat "$fix_out"
+  die "fix worker generated no pending work units"
+fi
+done
diff --git a/tests/run.sh b/tests/run.sh
index eb28d49f637f95de4c6e555bb98b1fe43455e7f2..7a28bc257ecfc567e30366b85efdea03ea9a59c3
--- a/tests/run.sh
+++ b/tests/run.sh
@@ -169,6 +169,123 @@
 test -f "$repo_lpf/captured-prompt.txt"
 assert_contains "$repo_lpf/captured-prompt.txt" "the test fixture reaches its verify condition"
 
+# Review-fix loop with fake build, review, and fix workers.
+repo_review="$tmp/repo-review"
+mkdir -p "$repo_review/.loop"
+cat > "$repo_review/.loop/QUEUE.md" <<'EOF'
+# Loop Queue: review cycle
+
+Goal:
+Exercise build, review, fix, and review again.
+
+Stop condition:
+The generated app is fixed and review reports no actionable issues.
+
+## the initial build creates a reviewable app file
+
+Read first:
+- This queue file.
+
+Constraints:
+- Leave the bug for review to find.
+
+Verify:
+```bash
+test -f app.txt
+```
+
+Done means:
+- app.txt exists.
+
+Status: pending
+EOF
+
+cat > "$repo_review/build-worker.sh" <<'EOF'
+#!/usr/bin/env bash
+set -euo pipefail
+if grep -q "the fix unit repairs the bug" "$LOOP_PROMPT_FILE"; then
+  printf 'fixed\n' > app.txt
+  echo "build fixed app"
+else
+  printf 'bug\n' > app.txt
+  echo "build created app with bug"
+fi
+EOF
+chmod +x "$repo_review/build-worker.sh"
+
+cat > "$repo_review/review-worker.sh" <<'EOF'
+#!/usr/bin/env bash
+set -euo pipefail
+count=0
+if [[ -f review-count.txt ]]; then
+  count=$(cat review-count.txt)
+fi
+count=$((count + 1))
+printf '%s\n' "$count" > review-count.txt
+
+actionable=1
+if [[ -f app.txt ]] && grep -qx 'fixed' app.txt; then
+  actionable=0
+fi
+
+cat > "$LOOP_REVIEW_FILE" <<EOF_REVIEW
+# Review: fake
+
+## Standards
+
+## Intent
+- actionable | high — app.txt must say fixed
+  evidence: app.txt:1
+
+## Speculative
+
+## Summary
+- actionable: $actionable
+- trivial: 0
+- disputed: 0
+- deferred: 0
+EOF_REVIEW
+echo "review actionable: $actionable"
+EOF
+chmod +x "$repo_review/review-worker.sh"
+
+cat > "$repo_review/fix-worker.sh" <<'EOF'
+#!/usr/bin/env bash
+set -euo pipefail
+cat >> "$LOOP_QUEUE_FILE" <<'EOF_QUEUE'
+
+## the fix unit repairs the bug
+
+Read first:
+- .loop/REVIEW.md
+
+Constraints:
+- Preserve the app file created by the first unit.
+
+Verify:
+```bash
+grep -qx fixed app.txt
+```
+
+Done means:
+- app.txt contains fixed.
+
+Status: pending
+EOF_QUEUE
+echo "fix appended unit"
+EOF
+chmod +x "$repo_review/fix-worker.sh"
+
+LOOP_AGENT_CMD="$repo_review/build-worker.sh" \
+  LOOP_REVIEW_CMD="$repo_review/review-worker.sh" \
+  LOOP_FIX_CMD="$repo_review/fix-worker.sh" \
+  "$root/loop.sh" run "$repo_review/.loop/QUEUE.md" --review --max-ticks 2 >/tmp/loop-review.txt
+assert_contains "$repo_review/.loop/QUEUE.md" "## the fix unit repairs the bug"
+assert_contains "$repo_review/.loop/QUEUE.md" "Status: done"
+assert_contains "$repo_review/.loop/REVIEW.md" "- actionable: 0"
+assert_contains "$repo_review/review-count.txt" "2"
+assert_contains "$repo_review/app.txt" "fixed"
+
 if command -v skills-ref >/dev/null 2>&1; then
   for skill_dir in "$root/.agents/skills"/*; do
     if [[ -d "$skill_dir" ]]; then

 succeeded in 3590ms:
Valid skill: /home/daniel/build/knack/.agents/skills/build
Valid skill: /home/daniel/build/knack/.agents/skills/decide
Valid skill: /home/daniel/build/knack/.agents/skills/domain-modeling
Valid skill: /home/daniel/build/knack/.agents/skills/explore
Valid skill: /home/daniel/build/knack/.agents/skills/fix
Valid skill: /home/daniel/build/knack/.agents/skills/plan
Valid skill: /home/daniel/build/knack/.agents/skills/review
knack tests passed

codex
The required verify command passes. I’m doing a quick diff review now to catch any accidental scope creep or interactions with the existing dirty files before the final handoff.
exec
/usr/bin/bash -lc 'git diff -- loop.sh tests/run.sh' in /home/daniel/build/knack
 succeeded in 0ms:
diff --git a/loop.sh b/loop.sh
index 351aca3..35066aa 100755
--- a/loop.sh
+++ b/loop.sh
@@ -4,7 +4,7 @@ set -euo pipefail
 usage() {
   cat <<'EOF'
 Usage:
-  ./loop.sh run <queue> [--repo DIR] [--max-ticks N] [--dry-run]
+  ./loop.sh run <queue> [--repo DIR] [--max-ticks N] [--review] [--max-review-rounds N] [--dry-run]
 
 The queue is usually .loop/<name>/QUEUE.md in the target repo. Use --repo when the
 queue lives outside the repository it should operate on.
@@ -204,6 +204,86 @@ append_evidence() {
   } >> "$evidence"
 }
 
+extract_actionable_count() {
+  awk '
+    /^[[:space:]]*-[[:space:]]*actionable:[[:space:]]*[0-9]+[[:space:]]*$/ {
+      sub(/^[[:space:]]*-[[:space:]]*actionable:[[:space:]]*/, "")
+      sub(/[[:space:]]*$/, "")
+      print
+      found = 1
+      exit
+    }
+    END { if (!found) exit 1 }
+  ' "$1"
+}
+
+write_review_prompt() {
+  local out=$1
+  local template="$script_dir/prompts/reviewer.md"
+  if [[ -f "$template" ]]; then
+    cat "$template" > "$out"
+  else
+    cat > "$out" <<'EOF'
+# Knack Reviewer
+
+Load and follow the **review** skill in `.agents/skills/review/`.
+Review the completed queue against the current repository state. Write the structured review artifact at the path provided below.
+EOF
+  fi
+  cat >> "$out" <<EOF
+
+Queue: $queue_abs
+Evidence: $evidence
+Review output: $review_file
+EOF
+}
+
+write_fix_prompt() {
+  local out=$1
+  local template="$script_dir/prompts/fixer.md"
+  if [[ -f "$template" ]]; then
+    cat "$template" > "$out"
+  else
+    cat > "$out" <<'EOF'
+# Knack Fixer
+
+Load and follow the **fix** skill in `.agents/skills/fix/`.
+Read the structured review artifact and append any actionable fix work units to the existing queue. Stop after updating the queue.
+EOF
+  fi
+  cat >> "$out" <<EOF
+
+Queue: $queue_abs
+Evidence: $evidence
+Review input: $review_file
+EOF
+}
+
+run_phase_agent() {
+  local phase=$1 prompt=$2 output=$3 cmd=${4:-}
+  if [[ -n "$cmd" ]]; then
+    (
+      cd "$repo_dir"
+      LOOP_PHASE="$phase" \
+        LOOP_PROMPT_FILE="$prompt" \
+        LOOP_QUEUE_FILE="$queue_abs" \
+        LOOP_EVIDENCE_FILE="$evidence" \
+        LOOP_REVIEW_FILE="$review_file" \
+        bash -lc "$cmd"
+    ) > "$output" 2>&1
+  else
+    (
+      cd "$repo_dir"
+      LOOP_PHASE="$phase" \
+        LOOP_PROMPT_FILE="$prompt" \
+        LOOP_QUEUE_FILE="$queue_abs" \
+        LOOP_EVIDENCE_FILE="$evidence" \
+        LOOP_REVIEW_FILE="$review_file" \
+        pi -p --no-session --approve "$(cat "$prompt")"
+    ) > "$output" 2>&1
+  fi
+}
+
 [[ $# -ge 1 ]] || { usage; exit 1; }
 cmd=$1
 shift
@@ -215,6 +295,8 @@ shift
 [[ -f "$queue_abs" ]] || die "queue not found: $queue_abs"
 
 max_ticks=3
+max_review_rounds=2
+review_enabled=0
 dry_run=0
 repo_override=""
 while [[ $# -gt 0 ]]; do
@@ -229,6 +311,15 @@ while [[ $# -gt 0 ]]; do
       max_ticks=$2
       shift 2
       ;;
+    --review)
+      review_enabled=1
+      shift
+      ;;
+    --max-review-rounds)
+      [[ $# -ge 2 ]] || die "--max-review-rounds needs a value"
+      max_review_rounds=$2
+      shift 2
+      ;;
     --dry-run)
       dry_run=1
       shift
@@ -252,6 +343,7 @@ else
 fi
 evidence="$queue_dir/EVIDENCE.md"
 handoff="$queue_dir/HANDOFF.md"
+review_file="$queue_dir/REVIEW.md"
 no_progress_strikes=0
 
 write_handoff_on_exit() {
@@ -263,9 +355,13 @@ write_handoff_on_exit() {
 }
 trap write_handoff_on_exit EXIT
 
-for ((tick = 1; tick <= max_ticks; tick++)); do
+tick=1
+review_round=0
+
+while true; do
+while (( tick <= max_ticks )); do
   unit=$(first_pending_unit "$queue_abs")
-  [[ -n "$unit" ]] || { echo "knack: no pending work units"; exit 0; }
+  [[ -n "$unit" ]] || break
 
   unit_file=$(mktemp)
   verify_file=$(mktemp)
@@ -307,11 +403,7 @@ EOF
   fi
 
   set +e
-  if [[ -n "$agent_cmd" ]]; then
-    (cd "$repo_dir" && LOOP_PROMPT_FILE="$run_prompt" bash -lc "$agent_cmd") > "$agent_out" 2>&1
-  else
-    (cd "$repo_dir" && pi -p --no-session --approve "$(cat "$run_prompt")") > "$agent_out" 2>&1
-  fi
+  run_phase_agent "build" "$run_prompt" "$agent_out" "$agent_cmd"
   agent_code=$?
   set -e
 
@@ -334,6 +426,7 @@ EOF
     set_status "$queue_abs" "$title" "done"
     append_evidence "$evidence" "$title" "done" "$verify" "$verify_out" "$agent_out" "$repo_dir" "$unit_file"
     echo "knack: verified — $title"
+    tick=$((tick + 1))
     continue
   fi
 
@@ -347,6 +440,7 @@ EOF
     set_status "$queue_abs" "$title" "pending"
     append_evidence "$evidence" "$title" "verify_failed" "$verify" "$verify_out" "$agent_out" "$repo_dir" "$unit_file"
     echo "knack: verify failed with no progress; retrying once"
+    tick=$((tick + 1))
     continue
   fi
 
@@ -360,4 +454,59 @@ if [[ -n "$(first_pending_unit "$queue_abs")" ]]; then
   die "reached max ticks ($max_ticks) with pending work"
 fi
 
-echo "knack: reached max ticks ($max_ticks)"
+if [[ "$review_enabled" == 0 ]]; then
+  echo "knack: no pending work units"
+  exit 0
+fi
+
+if (( review_round >= max_review_rounds )); then
+  die "reached max review rounds ($max_review_rounds)"
+fi
+
+review_round=$((review_round + 1))
+review_prompt=$(mktemp)
+review_out=$(mktemp)
+write_review_prompt "$review_prompt"
+review_cmd="${LOOP_REVIEW_CMD:-${LOOP_AGENT_CMD:-}}"
+
+echo "knack: review round $review_round/$max_review_rounds"
+set +e
+run_phase_agent "review" "$review_prompt" "$review_out" "$review_cmd"
+review_code=$?
+set -e
+if [[ $review_code -ne 0 ]]; then
+  cat "$review_out"
+  die "review worker exited nonzero"
+fi
+[[ -f "$review_file" ]] || die "review worker did not write $review_file"
+actionable=$(extract_actionable_count "$review_file") || die "review file has no actionable summary: $review_file"
+
+if [[ "$actionable" == 0 ]]; then
+  echo "knack: review clean"
+  exit 0
+fi
+
+if (( review_round >= max_review_rounds )); then
+  die "review found $actionable actionable issue(s) at max review rounds ($max_review_rounds)"
+fi
+
+fix_prompt=$(mktemp)
+fix_out=$(mktemp)
+write_fix_prompt "$fix_prompt"
+fix_cmd="${LOOP_FIX_CMD:-${LOOP_AGENT_CMD:-}}"
+
+echo "knack: fix round $review_round/$max_review_rounds — actionable: $actionable"
+set +e
+run_phase_agent "fix" "$fix_prompt" "$fix_out" "$fix_cmd"
+fix_code=$?
+set -e
+if [[ $fix_code -ne 0 ]]; then
+  cat "$fix_out"
+  die "fix worker exited nonzero"
+fi
+
+if [[ -z "$(first_pending_unit "$queue_abs")" ]]; then
+  cat "$fix_out"
+  die "fix worker generated no pending work units"
+fi
+done
diff --git a/tests/run.sh b/tests/run.sh
index eb28d49..7a28bc2 100755
--- a/tests/run.sh
+++ b/tests/run.sh
@@ -169,6 +169,123 @@ assert_contains "$repo_lpf/.loop/QUEUE.md" "Status: done"
 test -f "$repo_lpf/captured-prompt.txt"
 assert_contains "$repo_lpf/captured-prompt.txt" "the test fixture reaches its verify condition"
 
+# Review-fix loop with fake build, review, and fix workers.
+repo_review="$tmp/repo-review"
+mkdir -p "$repo_review/.loop"
+cat > "$repo_review/.loop/QUEUE.md" <<'EOF'
+# Loop Queue: review cycle
+
+Goal:
+Exercise build, review, fix, and review again.
+
+Stop condition:
+The generated app is fixed and review reports no actionable issues.
+
+## the initial build creates a reviewable app file
+
+Read first:
+- This queue file.
+
+Constraints:
+- Leave the bug for review to find.
+
+Verify:
+```bash
+test -f app.txt
+```
+
+Done means:
+- app.txt exists.
+
+Status: pending
+EOF
+
+cat > "$repo_review/build-worker.sh" <<'EOF'
+#!/usr/bin/env bash
+set -euo pipefail
+if grep -q "the fix unit repairs the bug" "$LOOP_PROMPT_FILE"; then
+  printf 'fixed\n' > app.txt
+  echo "build fixed app"
+else
+  printf 'bug\n' > app.txt
+  echo "build created app with bug"
+fi
+EOF
+chmod +x "$repo_review/build-worker.sh"
+
+cat > "$repo_review/review-worker.sh" <<'EOF'
+#!/usr/bin/env bash
+set -euo pipefail
+count=0
+if [[ -f review-count.txt ]]; then
+  count=$(cat review-count.txt)
+fi
+count=$((count + 1))
+printf '%s\n' "$count" > review-count.txt
+
+actionable=1
+if [[ -f app.txt ]] && grep -qx 'fixed' app.txt; then
+  actionable=0
+fi
+
+cat > "$LOOP_REVIEW_FILE" <<EOF_REVIEW
+# Review: fake
+
+## Standards
+
+## Intent
+- actionable | high — app.txt must say fixed
+  evidence: app.txt:1
+
+## Speculative
+
+## Summary
+- actionable: $actionable
+- trivial: 0
+- disputed: 0
+- deferred: 0
+EOF_REVIEW
+echo "review actionable: $actionable"
+EOF
+chmod +x "$repo_review/review-worker.sh"
+
+cat > "$repo_review/fix-worker.sh" <<'EOF'
+#!/usr/bin/env bash
+set -euo pipefail
+cat >> "$LOOP_QUEUE_FILE" <<'EOF_QUEUE'
+
+## the fix unit repairs the bug
+
+Read first:
+- .loop/REVIEW.md
+
+Constraints:
+- Preserve the app file created by the first unit.
+
+Verify:
+```bash
+grep -qx fixed app.txt
+```
+
+Done means:
+- app.txt contains fixed.
+
+Status: pending
+EOF_QUEUE
+echo "fix appended unit"
+EOF
+chmod +x "$repo_review/fix-worker.sh"
+
+LOOP_AGENT_CMD="$repo_review/build-worker.sh" \
+  LOOP_REVIEW_CMD="$repo_review/review-worker.sh" \
+  LOOP_FIX_CMD="$repo_review/fix-worker.sh" \
+  "$root/loop.sh" run "$repo_review/.loop/QUEUE.md" --review --max-ticks 2 >/tmp/loop-review.txt
+assert_contains "$repo_review/.loop/QUEUE.md" "## the fix unit repairs the bug"
+assert_contains "$repo_review/.loop/QUEUE.md" "Status: done"
+assert_contains "$repo_review/.loop/REVIEW.md" "- actionable: 0"
+assert_contains "$repo_review/review-count.txt" "2"
+assert_contains "$repo_review/app.txt" "fixed"
+
 if command -v skills-ref >/dev/null 2>&1; then
   for skill_dir in "$root/.agents/skills"/*; do
     if [[ -d "$skill_dir" ]]; then

exec
/usr/bin/bash -lc 'git status --short' in /home/daniel/build/knack
 succeeded in 0ms:
 M .agents/skills/fix/SKILL.md
 M .loop/review-loop/EVIDENCE.md
 M .loop/review-loop/HANDOFF.md
 M .loop/review-loop/QUEUE.md
 M cli/embedded/skills/fix/SKILL.md
 M loop.sh
 M tests/run.sh

diff --git a/loop.sh b/loop.sh
index 351aca3855b1ce6c9c08b0c586ad838fa0f6e636..35066aad26502fbf3f7d5ac0413747bb134cc0b0
--- a/loop.sh
+++ b/loop.sh
@@ -4,7 +4,7 @@
 usage() {
   cat <<'EOF'
 Usage:
-  ./loop.sh run <queue> [--repo DIR] [--max-ticks N] [--dry-run]
+  ./loop.sh run <queue> [--repo DIR] [--max-ticks N] [--review] [--max-review-rounds N] [--dry-run]
 
 The queue is usually .loop/<name>/QUEUE.md in the target repo. Use --repo when the
 queue lives outside the repository it should operate on.
@@ -204,6 +204,86 @@
   } >> "$evidence"
 }
 
+extract_actionable_count() {
+  awk '
+    /^[[:space:]]*-[[:space:]]*actionable:[[:space:]]*[0-9]+[[:space:]]*$/ {
+      sub(/^[[:space:]]*-[[:space:]]*actionable:[[:space:]]*/, "")
+      sub(/[[:space:]]*$/, "")
+      print
+      found = 1
+      exit
+    }
+    END { if (!found) exit 1 }
+  ' "$1"
+}
+
+write_review_prompt() {
+  local out=$1
+  local template="$script_dir/prompts/reviewer.md"
+  if [[ -f "$template" ]]; then
+    cat "$template" > "$out"
+  else
+    cat > "$out" <<'EOF'
+# Knack Reviewer
+
+Load and follow the **review** skill in `.agents/skills/review/`.
+Review the completed queue against the current repository state. Write the structured review artifact at the path provided below.
+EOF
+  fi
+  cat >> "$out" <<EOF
+
+Queue: $queue_abs
+Evidence: $evidence
+Review output: $review_file
+EOF
+}
+
+write_fix_prompt() {
+  local out=$1
+  local template="$script_dir/prompts/fixer.md"
+  if [[ -f "$template" ]]; then
+    cat "$template" > "$out"
+  else
+    cat > "$out" <<'EOF'
+# Knack Fixer
+
+Load and follow the **fix** skill in `.agents/skills/fix/`.
+Read the structured review artifact and append any actionable fix work units to the existing queue. Stop after updating the queue.
+EOF
+  fi
+  cat >> "$out" <<EOF
+
+Queue: $queue_abs
+Evidence: $evidence
+Review input: $review_file
+EOF
+}
+
+run_phase_agent() {
+  local phase=$1 prompt=$2 output=$3 cmd=${4:-}
+  if [[ -n "$cmd" ]]; then
+    (
+      cd "$repo_dir"
+      LOOP_PHASE="$phase" \
+        LOOP_PROMPT_FILE="$prompt" \
+        LOOP_QUEUE_FILE="$queue_abs" \
+        LOOP_EVIDENCE_FILE="$evidence" \
+        LOOP_REVIEW_FILE="$review_file" \
+        bash -lc "$cmd"
+    ) > "$output" 2>&1
+  else
+    (
+      cd "$repo_dir"
+      LOOP_PHASE="$phase" \
+        LOOP_PROMPT_FILE="$prompt" \
+        LOOP_QUEUE_FILE="$queue_abs" \
+        LOOP_EVIDENCE_FILE="$evidence" \
+        LOOP_REVIEW_FILE="$review_file" \
+        pi -p --no-session --approve "$(cat "$prompt")"
+    ) > "$output" 2>&1
+  fi
+}
+
 [[ $# -ge 1 ]] || { usage; exit 1; }
 cmd=$1
 shift
@@ -215,6 +295,8 @@
 [[ -f "$queue_abs" ]] || die "queue not found: $queue_abs"
 
 max_ticks=3
+max_review_rounds=2
+review_enabled=0
 dry_run=0
 repo_override=""
 while [[ $# -gt 0 ]]; do
@@ -229,6 +311,15 @@
       max_ticks=$2
       shift 2
       ;;
+    --review)
+      review_enabled=1
+      shift
+      ;;
+    --max-review-rounds)
+      [[ $# -ge 2 ]] || die "--max-review-rounds needs a value"
+      max_review_rounds=$2
+      shift 2
+      ;;
     --dry-run)
       dry_run=1
       shift
@@ -252,6 +343,7 @@
 fi
 evidence="$queue_dir/EVIDENCE.md"
 handoff="$queue_dir/HANDOFF.md"
+review_file="$queue_dir/REVIEW.md"
 no_progress_strikes=0
 
 write_handoff_on_exit() {
@@ -263,9 +355,13 @@
 }
 trap write_handoff_on_exit EXIT
 
-for ((tick = 1; tick <= max_ticks; tick++)); do
+tick=1
+review_round=0
+
+while true; do
+while (( tick <= max_ticks )); do
   unit=$(first_pending_unit "$queue_abs")
-  [[ -n "$unit" ]] || { echo "knack: no pending work units"; exit 0; }
+  [[ -n "$unit" ]] || break
 
   unit_file=$(mktemp)
   verify_file=$(mktemp)
@@ -307,11 +403,7 @@
   fi
 
   set +e
-  if [[ -n "$agent_cmd" ]]; then
-    (cd "$repo_dir" && LOOP_PROMPT_FILE="$run_prompt" bash -lc "$agent_cmd") > "$agent_out" 2>&1
-  else
-    (cd "$repo_dir" && pi -p --no-session --approve "$(cat "$run_prompt")") > "$agent_out" 2>&1
-  fi
+  run_phase_agent "build" "$run_prompt" "$agent_out" "$agent_cmd"
   agent_code=$?
   set -e
 
@@ -334,6 +426,7 @@
     set_status "$queue_abs" "$title" "done"
     append_evidence "$evidence" "$title" "done" "$verify" "$verify_out" "$agent_out" "$repo_dir" "$unit_file"
     echo "knack: verified — $title"
+    tick=$((tick + 1))
     continue
   fi
 
@@ -347,6 +440,7 @@
     set_status "$queue_abs" "$title" "pending"
     append_evidence "$evidence" "$title" "verify_failed" "$verify" "$verify_out" "$agent_out" "$repo_dir" "$unit_file"
     echo "knack: verify failed with no progress; retrying once"
+    tick=$((tick + 1))
     continue
   fi
 
@@ -360,4 +454,59 @@
   die "reached max ticks ($max_ticks) with pending work"
 fi
 
-echo "knack: reached max ticks ($max_ticks)"
+if [[ "$review_enabled" == 0 ]]; then
+  echo "knack: no pending work units"
+  exit 0
+fi
+
+if (( review_round >= max_review_rounds )); then
+  die "reached max review rounds ($max_review_rounds)"
+fi
+
+review_round=$((review_round + 1))
+review_prompt=$(mktemp)
+review_out=$(mktemp)
+write_review_prompt "$review_prompt"
+review_cmd="${LOOP_REVIEW_CMD:-${LOOP_AGENT_CMD:-}}"
+
+echo "knack: review round $review_round/$max_review_rounds"
+set +e
+run_phase_agent "review" "$review_prompt" "$review_out" "$review_cmd"
+review_code=$?
+set -e
+if [[ $review_code -ne 0 ]]; then
+  cat "$review_out"
+  die "review worker exited nonzero"
+fi
+[[ -f "$review_file" ]] || die "review worker did not write $review_file"
+actionable=$(extract_actionable_count "$review_file") || die "review file has no actionable summary: $review_file"
+
+if [[ "$actionable" == 0 ]]; then
+  echo "knack: review clean"
+  exit 0
+fi
+
+if (( review_round >= max_review_rounds )); then
+  die "review found $actionable actionable issue(s) at max review rounds ($max_review_rounds)"
+fi
+
+fix_prompt=$(mktemp)
+fix_out=$(mktemp)
+write_fix_prompt "$fix_prompt"
+fix_cmd="${LOOP_FIX_CMD:-${LOOP_AGENT_CMD:-}}"
+
+echo "knack: fix round $review_round/$max_review_rounds — actionable: $actionable"
+set +e
+run_phase_agent "fix" "$fix_prompt" "$fix_out" "$fix_cmd"
+fix_code=$?
+set -e
+if [[ $fix_code -ne 0 ]]; then
+  cat "$fix_out"
+  die "fix worker exited nonzero"
+fi
+
+if [[ -z "$(first_pending_unit "$queue_abs")" ]]; then
+  cat "$fix_out"
+  die "fix worker generated no pending work units"
+fi
+done
diff --git a/tests/run.sh b/tests/run.sh
index eb28d49f637f95de4c6e555bb98b1fe43455e7f2..7a28bc257ecfc567e30366b85efdea03ea9a59c3
--- a/tests/run.sh
+++ b/tests/run.sh
@@ -169,6 +169,123 @@
 test -f "$repo_lpf/captured-prompt.txt"
 assert_contains "$repo_lpf/captured-prompt.txt" "the test fixture reaches its verify condition"
 
+# Review-fix loop with fake build, review, and fix workers.
+repo_review="$tmp/repo-review"
+mkdir -p "$repo_review/.loop"
+cat > "$repo_review/.loop/QUEUE.md" <<'EOF'
+# Loop Queue: review cycle
+
+Goal:
+Exercise build, review, fix, and review again.
+
+Stop condition:
+The generated app is fixed and review reports no actionable issues.
+
+## the initial build creates a reviewable app file
+
+Read first:
+- This queue file.
+
+Constraints:
+- Leave the bug for review to find.
+
+Verify:
+```bash
+test -f app.txt
+```
+
+Done means:
+- app.txt exists.
+
+Status: pending
+EOF
+
+cat > "$repo_review/build-worker.sh" <<'EOF'
+#!/usr/bin/env bash
+set -euo pipefail
+if grep -q "the fix unit repairs the bug" "$LOOP_PROMPT_FILE"; then
+  printf 'fixed\n' > app.txt
+  echo "build fixed app"
+else
+  printf 'bug\n' > app.txt
+  echo "build created app with bug"
+fi
+EOF
+chmod +x "$repo_review/build-worker.sh"
+
+cat > "$repo_review/review-worker.sh" <<'EOF'
+#!/usr/bin/env bash
+set -euo pipefail
+count=0
+if [[ -f review-count.txt ]]; then
+  count=$(cat review-count.txt)
+fi
+count=$((count + 1))
+printf '%s\n' "$count" > review-count.txt
+
+actionable=1
+if [[ -f app.txt ]] && grep -qx 'fixed' app.txt; then
+  actionable=0
+fi
+
+cat > "$LOOP_REVIEW_FILE" <<EOF_REVIEW
+# Review: fake
+
+## Standards
+
+## Intent
+- actionable | high — app.txt must say fixed
+  evidence: app.txt:1
+
+## Speculative
+
+## Summary
+- actionable: $actionable
+- trivial: 0
+- disputed: 0
+- deferred: 0
+EOF_REVIEW
+echo "review actionable: $actionable"
+EOF
+chmod +x "$repo_review/review-worker.sh"
+
+cat > "$repo_review/fix-worker.sh" <<'EOF'
+#!/usr/bin/env bash
+set -euo pipefail
+cat >> "$LOOP_QUEUE_FILE" <<'EOF_QUEUE'
+
+## the fix unit repairs the bug
+
+Read first:
+- .loop/REVIEW.md
+
+Constraints:
+- Preserve the app file created by the first unit.
+
+Verify:
+```bash
+grep -qx fixed app.txt
+```
+
+Done means:
+- app.txt contains fixed.
+
+Status: pending
+EOF_QUEUE
+echo "fix appended unit"
+EOF
+chmod +x "$repo_review/fix-worker.sh"
+
+LOOP_AGENT_CMD="$repo_review/build-worker.sh" \
+  LOOP_REVIEW_CMD="$repo_review/review-worker.sh" \
+  LOOP_FIX_CMD="$repo_review/fix-worker.sh" \
+  "$root/loop.sh" run "$repo_review/.loop/QUEUE.md" --review --max-ticks 2 >/tmp/loop-review.txt
+assert_contains "$repo_review/.loop/QUEUE.md" "## the fix unit repairs the bug"
+assert_contains "$repo_review/.loop/QUEUE.md" "Status: done"
+assert_contains "$repo_review/.loop/REVIEW.md" "- actionable: 0"
+assert_contains "$repo_review/review-count.txt" "2"
+assert_contains "$repo_review/app.txt" "fixed"
+
 if command -v skills-ref >/dev/null 2>&1; then
   for skill_dir in "$root/.agents/skills"/*; do
     if [[ -d "$skill_dir" ]]; then

codex
The diff is narrow. I spotted one small legacy-behavior mismatch in the no-review completion message when the last allowed build tick drains the queue, so I’m preserving that exact branch before I finish.
apply patch
patch: completed
/home/daniel/build/knack/loop.sh
diff --git a/loop.sh b/loop.sh
index 351aca3855b1ce6c9c08b0c586ad838fa0f6e636..cfa2c6a3b81faf08ea400eff3219310fe24657a6
--- a/loop.sh
+++ b/loop.sh
@@ -4,7 +4,7 @@
 usage() {
   cat <<'EOF'
 Usage:
-  ./loop.sh run <queue> [--repo DIR] [--max-ticks N] [--dry-run]
+  ./loop.sh run <queue> [--repo DIR] [--max-ticks N] [--review] [--max-review-rounds N] [--dry-run]
 
 The queue is usually .loop/<name>/QUEUE.md in the target repo. Use --repo when the
 queue lives outside the repository it should operate on.
@@ -204,6 +204,86 @@
   } >> "$evidence"
 }
 
+extract_actionable_count() {
+  awk '
+    /^[[:space:]]*-[[:space:]]*actionable:[[:space:]]*[0-9]+[[:space:]]*$/ {
+      sub(/^[[:space:]]*-[[:space:]]*actionable:[[:space:]]*/, "")
+      sub(/[[:space:]]*$/, "")
+      print
+      found = 1
+      exit
+    }
+    END { if (!found) exit 1 }
+  ' "$1"
+}
+
+write_review_prompt() {
+  local out=$1
+  local template="$script_dir/prompts/reviewer.md"
+  if [[ -f "$template" ]]; then
+    cat "$template" > "$out"
+  else
+    cat > "$out" <<'EOF'
+# Knack Reviewer
+
+Load and follow the **review** skill in `.agents/skills/review/`.
+Review the completed queue against the current repository state. Write the structured review artifact at the path provided below.
+EOF
+  fi
+  cat >> "$out" <<EOF
+
+Queue: $queue_abs
+Evidence: $evidence
+Review output: $review_file
+EOF
+}
+
+write_fix_prompt() {
+  local out=$1
+  local template="$script_dir/prompts/fixer.md"
+  if [[ -f "$template" ]]; then
+    cat "$template" > "$out"
+  else
+    cat > "$out" <<'EOF'
+# Knack Fixer
+
+Load and follow the **fix** skill in `.agents/skills/fix/`.
+Read the structured review artifact and append any actionable fix work units to the existing queue. Stop after updating the queue.
+EOF
+  fi
+  cat >> "$out" <<EOF
+
+Queue: $queue_abs
+Evidence: $evidence
+Review input: $review_file
+EOF
+}
+
+run_phase_agent() {
+  local phase=$1 prompt=$2 output=$3 cmd=${4:-}
+  if [[ -n "$cmd" ]]; then
+    (
+      cd "$repo_dir"
+      LOOP_PHASE="$phase" \
+        LOOP_PROMPT_FILE="$prompt" \
+        LOOP_QUEUE_FILE="$queue_abs" \
+        LOOP_EVIDENCE_FILE="$evidence" \
+        LOOP_REVIEW_FILE="$review_file" \
+        bash -lc "$cmd"
+    ) > "$output" 2>&1
+  else
+    (
+      cd "$repo_dir"
+      LOOP_PHASE="$phase" \
+        LOOP_PROMPT_FILE="$prompt" \
+        LOOP_QUEUE_FILE="$queue_abs" \
+        LOOP_EVIDENCE_FILE="$evidence" \
+        LOOP_REVIEW_FILE="$review_file" \
+        pi -p --no-session --approve "$(cat "$prompt")"
+    ) > "$output" 2>&1
+  fi
+}
+
 [[ $# -ge 1 ]] || { usage; exit 1; }
 cmd=$1
 shift
@@ -215,6 +295,8 @@
 [[ -f "$queue_abs" ]] || die "queue not found: $queue_abs"
 
 max_ticks=3
+max_review_rounds=2
+review_enabled=0
 dry_run=0
 repo_override=""
 while [[ $# -gt 0 ]]; do
@@ -229,6 +311,15 @@
       max_ticks=$2
       shift 2
       ;;
+    --review)
+      review_enabled=1
+      shift
+      ;;
+    --max-review-rounds)
+      [[ $# -ge 2 ]] || die "--max-review-rounds needs a value"
+      max_review_rounds=$2
+      shift 2
+      ;;
     --dry-run)
       dry_run=1
       shift
@@ -252,6 +343,7 @@
 fi
 evidence="$queue_dir/EVIDENCE.md"
 handoff="$queue_dir/HANDOFF.md"
+review_file="$queue_dir/REVIEW.md"
 no_progress_strikes=0
 
 write_handoff_on_exit() {
@@ -263,9 +355,13 @@
 }
 trap write_handoff_on_exit EXIT
 
-for ((tick = 1; tick <= max_ticks; tick++)); do
+tick=1
+review_round=0
+
+while true; do
+while (( tick <= max_ticks )); do
   unit=$(first_pending_unit "$queue_abs")
-  [[ -n "$unit" ]] || { echo "knack: no pending work units"; exit 0; }
+  [[ -n "$unit" ]] || break
 
   unit_file=$(mktemp)
   verify_file=$(mktemp)
@@ -307,11 +403,7 @@
   fi
 
   set +e
-  if [[ -n "$agent_cmd" ]]; then
-    (cd "$repo_dir" && LOOP_PROMPT_FILE="$run_prompt" bash -lc "$agent_cmd") > "$agent_out" 2>&1
-  else
-    (cd "$repo_dir" && pi -p --no-session --approve "$(cat "$run_prompt")") > "$agent_out" 2>&1
-  fi
+  run_phase_agent "build" "$run_prompt" "$agent_out" "$agent_cmd"
   agent_code=$?
   set -e
 
@@ -334,6 +426,7 @@
     set_status "$queue_abs" "$title" "done"
     append_evidence "$evidence" "$title" "done" "$verify" "$verify_out" "$agent_out" "$repo_dir" "$unit_file"
     echo "knack: verified — $title"
+    tick=$((tick + 1))
     continue
   fi
 
@@ -347,6 +440,7 @@
     set_status "$queue_abs" "$title" "pending"
     append_evidence "$evidence" "$title" "verify_failed" "$verify" "$verify_out" "$agent_out" "$repo_dir" "$unit_file"
     echo "knack: verify failed with no progress; retrying once"
+    tick=$((tick + 1))
     continue
   fi
 
@@ -360,4 +454,63 @@
   die "reached max ticks ($max_ticks) with pending work"
 fi
 
-echo "knack: reached max ticks ($max_ticks)"
+if [[ "$review_enabled" == 0 ]]; then
+  if (( tick > max_ticks )); then
+    echo "knack: reached max ticks ($max_ticks)"
+    exit 0
+  fi
+  echo "knack: no pending work units"
+  exit 0
+fi
+
+if (( review_round >= max_review_rounds )); then
+  die "reached max review rounds ($max_review_rounds)"
+fi
+
+review_round=$((review_round + 1))
+review_prompt=$(mktemp)
+review_out=$(mktemp)
+write_review_prompt "$review_prompt"
+review_cmd="${LOOP_REVIEW_CMD:-${LOOP_AGENT_CMD:-}}"
+
+echo "knack: review round $review_round/$max_review_rounds"
+set +e
+run_phase_agent "review" "$review_prompt" "$review_out" "$review_cmd"
+review_code=$?
+set -e
+if [[ $review_code -ne 0 ]]; then
+  cat "$review_out"
+  die "review worker exited nonzero"
+fi
+[[ -f "$review_file" ]] || die "review worker did not write $review_file"
+actionable=$(extract_actionable_count "$review_file") || die "review file has no actionable summary: $review_file"
+
+if [[ "$actionable" == 0 ]]; then
+  echo "knack: review clean"
+  exit 0
+fi
+
+if (( review_round >= max_review_rounds )); then
+  die "review found $actionable actionable issue(s) at max review rounds ($max_review_rounds)"
+fi
+
+fix_prompt=$(mktemp)
+fix_out=$(mktemp)
+write_fix_prompt "$fix_prompt"
+fix_cmd="${LOOP_FIX_CMD:-${LOOP_AGENT_CMD:-}}"
+
+echo "knack: fix round $review_round/$max_review_rounds — actionable: $actionable"
+set +e
+run_phase_agent "fix" "$fix_prompt" "$fix_out" "$fix_cmd"
+fix_code=$?
+set -e
+if [[ $fix_code -ne 0 ]]; then
+  cat "$fix_out"
+  die "fix worker exited nonzero"
+fi
+
+if [[ -z "$(first_pending_unit "$queue_abs")" ]]; then
+  cat "$fix_out"
+  die "fix worker generated no pending work units"
+fi
+done
diff --git a/tests/run.sh b/tests/run.sh
index eb28d49f637f95de4c6e555bb98b1fe43455e7f2..7a28bc257ecfc567e30366b85efdea03ea9a59c3
--- a/tests/run.sh
+++ b/tests/run.sh
@@ -169,6 +169,123 @@
 test -f "$repo_lpf/captured-prompt.txt"
 assert_contains "$repo_lpf/captured-prompt.txt" "the test fixture reaches its verify condition"
 
+# Review-fix loop with fake build, review, and fix workers.
+repo_review="$tmp/repo-review"
+mkdir -p "$repo_review/.loop"
+cat > "$repo_review/.loop/QUEUE.md" <<'EOF'
+# Loop Queue: review cycle
+
+Goal:
+Exercise build, review, fix, and review again.
+
+Stop condition:
+The generated app is fixed and review reports no actionable issues.
+
+## the initial build creates a reviewable app file
+
+Read first:
+- This queue file.
+
+Constraints:
+- Leave the bug for review to find.
+
+Verify:
+```bash
+test -f app.txt
+```
+
+Done means:
+- app.txt exists.
+
+Status: pending
+EOF
+
+cat > "$repo_review/build-worker.sh" <<'EOF'
+#!/usr/bin/env bash
+set -euo pipefail
+if grep -q "the fix unit repairs the bug" "$LOOP_PROMPT_FILE"; then
+  printf 'fixed\n' > app.txt
+  echo "build fixed app"
+else
+  printf 'bug\n' > app.txt
+  echo "build created app with bug"
+fi
+EOF
+chmod +x "$repo_review/build-worker.sh"
+
+cat > "$repo_review/review-worker.sh" <<'EOF'
+#!/usr/bin/env bash
+set -euo pipefail
+count=0
+if [[ -f review-count.txt ]]; then
+  count=$(cat review-count.txt)
+fi
+count=$((count + 1))
+printf '%s\n' "$count" > review-count.txt
+
+actionable=1
+if [[ -f app.txt ]] && grep -qx 'fixed' app.txt; then
+  actionable=0
+fi
+
+cat > "$LOOP_REVIEW_FILE" <<EOF_REVIEW
+# Review: fake
+
+## Standards
+
+## Intent
+- actionable | high — app.txt must say fixed
+  evidence: app.txt:1
+
+## Speculative
+
+## Summary
+- actionable: $actionable
+- trivial: 0
+- disputed: 0
+- deferred: 0
+EOF_REVIEW
+echo "review actionable: $actionable"
+EOF
+chmod +x "$repo_review/review-worker.sh"
+
+cat > "$repo_review/fix-worker.sh" <<'EOF'
+#!/usr/bin/env bash
+set -euo pipefail
+cat >> "$LOOP_QUEUE_FILE" <<'EOF_QUEUE'
+
+## the fix unit repairs the bug
+
+Read first:
+- .loop/REVIEW.md
+
+Constraints:
+- Preserve the app file created by the first unit.
+
+Verify:
+```bash
+grep -qx fixed app.txt
+```
+
+Done means:
+- app.txt contains fixed.
+
+Status: pending
+EOF_QUEUE
+echo "fix appended unit"
+EOF
+chmod +x "$repo_review/fix-worker.sh"
+
+LOOP_AGENT_CMD="$repo_review/build-worker.sh" \
+  LOOP_REVIEW_CMD="$repo_review/review-worker.sh" \
+  LOOP_FIX_CMD="$repo_review/fix-worker.sh" \
+  "$root/loop.sh" run "$repo_review/.loop/QUEUE.md" --review --max-ticks 2 >/tmp/loop-review.txt
+assert_contains "$repo_review/.loop/QUEUE.md" "## the fix unit repairs the bug"
+assert_contains "$repo_review/.loop/QUEUE.md" "Status: done"
+assert_contains "$repo_review/.loop/REVIEW.md" "- actionable: 0"
+assert_contains "$repo_review/review-count.txt" "2"
+assert_contains "$repo_review/app.txt" "fixed"
+
 if command -v skills-ref >/dev/null 2>&1; then
   for skill_dir in "$root/.agents/skills"/*; do
     if [[ -d "$skill_dir" ]]; then

diff --git a/loop.sh b/loop.sh
index 351aca3855b1ce6c9c08b0c586ad838fa0f6e636..cfa2c6a3b81faf08ea400eff3219310fe24657a6
--- a/loop.sh
+++ b/loop.sh
@@ -4,7 +4,7 @@
 usage() {
   cat <<'EOF'
 Usage:
-  ./loop.sh run <queue> [--repo DIR] [--max-ticks N] [--dry-run]
+  ./loop.sh run <queue> [--repo DIR] [--max-ticks N] [--review] [--max-review-rounds N] [--dry-run]
 
 The queue is usually .loop/<name>/QUEUE.md in the target repo. Use --repo when the
 queue lives outside the repository it should operate on.
@@ -204,6 +204,86 @@
   } >> "$evidence"
 }
 
+extract_actionable_count() {
+  awk '
+    /^[[:space:]]*-[[:space:]]*actionable:[[:space:]]*[0-9]+[[:space:]]*$/ {
+      sub(/^[[:space:]]*-[[:space:]]*actionable:[[:space:]]*/, "")
+      sub(/[[:space:]]*$/, "")
+      print
+      found = 1
+      exit
+    }
+    END { if (!found) exit 1 }
+  ' "$1"
+}
+
+write_review_prompt() {
+  local out=$1
+  local template="$script_dir/prompts/reviewer.md"
+  if [[ -f "$template" ]]; then
+    cat "$template" > "$out"
+  else
+    cat > "$out" <<'EOF'
+# Knack Reviewer
+
+Load and follow the **review** skill in `.agents/skills/review/`.
+Review the completed queue against the current repository state. Write the structured review artifact at the path provided below.
+EOF
+  fi
+  cat >> "$out" <<EOF
+
+Queue: $queue_abs
+Evidence: $evidence
+Review output: $review_file
+EOF
+}
+
+write_fix_prompt() {
+  local out=$1
+  local template="$script_dir/prompts/fixer.md"
+  if [[ -f "$template" ]]; then
+    cat "$template" > "$out"
+  else
+    cat > "$out" <<'EOF'
+# Knack Fixer
+
+Load and follow the **fix** skill in `.agents/skills/fix/`.
+Read the structured review artifact and append any actionable fix work units to the existing queue. Stop after updating the queue.
+EOF
+  fi
+  cat >> "$out" <<EOF
+
+Queue: $queue_abs
+Evidence: $evidence
+Review input: $review_file
+EOF
+}
+
+run_phase_agent() {
+  local phase=$1 prompt=$2 output=$3 cmd=${4:-}
+  if [[ -n "$cmd" ]]; then
+    (
+      cd "$repo_dir"
+      LOOP_PHASE="$phase" \
+        LOOP_PROMPT_FILE="$prompt" \
+        LOOP_QUEUE_FILE="$queue_abs" \
+        LOOP_EVIDENCE_FILE="$evidence" \
+        LOOP_REVIEW_FILE="$review_file" \
+        bash -lc "$cmd"
+    ) > "$output" 2>&1
+  else
+    (
+      cd "$repo_dir"
+      LOOP_PHASE="$phase" \
+        LOOP_PROMPT_FILE="$prompt" \
+        LOOP_QUEUE_FILE="$queue_abs" \
+        LOOP_EVIDENCE_FILE="$evidence" \
+        LOOP_REVIEW_FILE="$review_file" \
+        pi -p --no-session --approve "$(cat "$prompt")"
+    ) > "$output" 2>&1
+  fi
+}
+
 [[ $# -ge 1 ]] || { usage; exit 1; }
 cmd=$1
 shift
@@ -215,6 +295,8 @@
 [[ -f "$queue_abs" ]] || die "queue not found: $queue_abs"
 
 max_ticks=3
+max_review_rounds=2
+review_enabled=0
 dry_run=0
 repo_override=""
 while [[ $# -gt 0 ]]; do
@@ -229,6 +311,15 @@
       max_ticks=$2
       shift 2
       ;;
+    --review)
+      review_enabled=1
+      shift
+      ;;
+    --max-review-rounds)
+      [[ $# -ge 2 ]] || die "--max-review-rounds needs a value"
+      max_review_rounds=$2
+      shift 2
+      ;;
     --dry-run)
       dry_run=1
       shift
@@ -252,6 +343,7 @@
 fi
 evidence="$queue_dir/EVIDENCE.md"
 handoff="$queue_dir/HANDOFF.md"
+review_file="$queue_dir/REVIEW.md"
 no_progress_strikes=0
 
 write_handoff_on_exit() {
@@ -263,9 +355,13 @@
 }
 trap write_handoff_on_exit EXIT
 
-for ((tick = 1; tick <= max_ticks; tick++)); do
+tick=1
+review_round=0
+
+while true; do
+while (( tick <= max_ticks )); do
   unit=$(first_pending_unit "$queue_abs")
-  [[ -n "$unit" ]] || { echo "knack: no pending work units"; exit 0; }
+  [[ -n "$unit" ]] || break
 
   unit_file=$(mktemp)
   verify_file=$(mktemp)
@@ -307,11 +403,7 @@
   fi
 
   set +e
-  if [[ -n "$agent_cmd" ]]; then
-    (cd "$repo_dir" && LOOP_PROMPT_FILE="$run_prompt" bash -lc "$agent_cmd") > "$agent_out" 2>&1
-  else
-    (cd "$repo_dir" && pi -p --no-session --approve "$(cat "$run_prompt")") > "$agent_out" 2>&1
-  fi
+  run_phase_agent "build" "$run_prompt" "$agent_out" "$agent_cmd"
   agent_code=$?
   set -e
 
@@ -334,6 +426,7 @@
     set_status "$queue_abs" "$title" "done"
     append_evidence "$evidence" "$title" "done" "$verify" "$verify_out" "$agent_out" "$repo_dir" "$unit_file"
     echo "knack: verified — $title"
+    tick=$((tick + 1))
     continue
   fi
 
@@ -347,6 +440,7 @@
     set_status "$queue_abs" "$title" "pending"
     append_evidence "$evidence" "$title" "verify_failed" "$verify" "$verify_out" "$agent_out" "$repo_dir" "$unit_file"
     echo "knack: verify failed with no progress; retrying once"
+    tick=$((tick + 1))
     continue
   fi
 
@@ -360,4 +454,63 @@
   die "reached max ticks ($max_ticks) with pending work"
 fi
 
-echo "knack: reached max ticks ($max_ticks)"
+if [[ "$review_enabled" == 0 ]]; then
+  if (( tick > max_ticks )); then
+    echo "knack: reached max ticks ($max_ticks)"
+    exit 0
+  fi
+  echo "knack: no pending work units"
+  exit 0
+fi
+
+if (( review_round >= max_review_rounds )); then
+  die "reached max review rounds ($max_review_rounds)"
+fi
+
+review_round=$((review_round + 1))
+review_prompt=$(mktemp)
+review_out=$(mktemp)
+write_review_prompt "$review_prompt"
+review_cmd="${LOOP_REVIEW_CMD:-${LOOP_AGENT_CMD:-}}"
+
+echo "knack: review round $review_round/$max_review_rounds"
+set +e
+run_phase_agent "review" "$review_prompt" "$review_out" "$review_cmd"
+review_code=$?
+set -e
+if [[ $review_code -ne 0 ]]; then
+  cat "$review_out"
+  die "review worker exited nonzero"
+fi
+[[ -f "$review_file" ]] || die "review worker did not write $review_file"
+actionable=$(extract_actionable_count "$review_file") || die "review file has no actionable summary: $review_file"
+
+if [[ "$actionable" == 0 ]]; then
+  echo "knack: review clean"
+  exit 0
+fi
+
+if (( review_round >= max_review_rounds )); then
+  die "review found $actionable actionable issue(s) at max review rounds ($max_review_rounds)"
+fi
+
+fix_prompt=$(mktemp)
+fix_out=$(mktemp)
+write_fix_prompt "$fix_prompt"
+fix_cmd="${LOOP_FIX_CMD:-${LOOP_AGENT_CMD:-}}"
+
+echo "knack: fix round $review_round/$max_review_rounds — actionable: $actionable"
+set +e
+run_phase_agent "fix" "$fix_prompt" "$fix_out" "$fix_cmd"
+fix_code=$?
+set -e
+if [[ $fix_code -ne 0 ]]; then
+  cat "$fix_out"
+  die "fix worker exited nonzero"
+fi
+
+if [[ -z "$(first_pending_unit "$queue_abs")" ]]; then
+  cat "$fix_out"
+  die "fix worker generated no pending work units"
+fi
+done
diff --git a/tests/run.sh b/tests/run.sh
index eb28d49f637f95de4c6e555bb98b1fe43455e7f2..7a28bc257ecfc567e30366b85efdea03ea9a59c3
--- a/tests/run.sh
+++ b/tests/run.sh
@@ -169,6 +169,123 @@
 test -f "$repo_lpf/captured-prompt.txt"
 assert_contains "$repo_lpf/captured-prompt.txt" "the test fixture reaches its verify condition"
 
+# Review-fix loop with fake build, review, and fix workers.
+repo_review="$tmp/repo-review"
+mkdir -p "$repo_review/.loop"
+cat > "$repo_review/.loop/QUEUE.md" <<'EOF'
+# Loop Queue: review cycle
+
+Goal:
+Exercise build, review, fix, and review again.
+
+Stop condition:
+The generated app is fixed and review reports no actionable issues.
+
+## the initial build creates a reviewable app file
+
+Read first:
+- This queue file.
+
+Constraints:
+- Leave the bug for review to find.
+
+Verify:
+```bash
+test -f app.txt
+```
+
+Done means:
+- app.txt exists.
+
+Status: pending
+EOF
+
+cat > "$repo_review/build-worker.sh" <<'EOF'
+#!/usr/bin/env bash
+set -euo pipefail
+if grep -q "the fix unit repairs the bug" "$LOOP_PROMPT_FILE"; then
+  printf 'fixed\n' > app.txt
+  echo "build fixed app"
+else
+  printf 'bug\n' > app.txt
+  echo "build created app with bug"
+fi
+EOF
+chmod +x "$repo_review/build-worker.sh"
+
+cat > "$repo_review/review-worker.sh" <<'EOF'
+#!/usr/bin/env bash
+set -euo pipefail
+count=0
+if [[ -f review-count.txt ]]; then
+  count=$(cat review-count.txt)
+fi
+count=$((count + 1))
+printf '%s\n' "$count" > review-count.txt
+
+actionable=1
+if [[ -f app.txt ]] && grep -qx 'fixed' app.txt; then
+  actionable=0
+fi
+
+cat > "$LOOP_REVIEW_FILE" <<EOF_REVIEW
+# Review: fake
+
+## Standards
+
+## Intent
+- actionable | high — app.txt must say fixed
+  evidence: app.txt:1
+
+## Speculative
+
+## Summary
+- actionable: $actionable
+- trivial: 0
+- disputed: 0
+- deferred: 0
+EOF_REVIEW
+echo "review actionable: $actionable"
+EOF
+chmod +x "$repo_review/review-worker.sh"
+
+cat > "$repo_review/fix-worker.sh" <<'EOF'
+#!/usr/bin/env bash
+set -euo pipefail
+cat >> "$LOOP_QUEUE_FILE" <<'EOF_QUEUE'
+
+## the fix unit repairs the bug
+
+Read first:
+- .loop/REVIEW.md
+
+Constraints:
+- Preserve the app file created by the first unit.
+
+Verify:
+```bash
+grep -qx fixed app.txt
+```
+
+Done means:
+- app.txt contains fixed.
+
+Status: pending
+EOF_QUEUE
+echo "fix appended unit"
+EOF
+chmod +x "$repo_review/fix-worker.sh"
+
+LOOP_AGENT_CMD="$repo_review/build-worker.sh" \
+  LOOP_REVIEW_CMD="$repo_review/review-worker.sh" \
+  LOOP_FIX_CMD="$repo_review/fix-worker.sh" \
+  "$root/loop.sh" run "$repo_review/.loop/QUEUE.md" --review --max-ticks 2 >/tmp/loop-review.txt
+assert_contains "$repo_review/.loop/QUEUE.md" "## the fix unit repairs the bug"
+assert_contains "$repo_review/.loop/QUEUE.md" "Status: done"
+assert_contains "$repo_review/.loop/REVIEW.md" "- actionable: 0"
+assert_contains "$repo_review/review-count.txt" "2"
+assert_contains "$repo_review/app.txt" "fixed"
+
 if command -v skills-ref >/dev/null 2>&1; then
   for skill_dir in "$root/.agents/skills"/*; do
     if [[ -d "$skill_dir" ]]; then

codex
That compatibility branch is restored. I’m rerunning the unit’s verify command once more after the final patch.
exec
/usr/bin/bash -lc 'bash -n /home/daniel/build/knack/loop.sh && /home/daniel/build/knack/tests/run.sh' in /home/daniel/build/knack
diff --git a/loop.sh b/loop.sh
index 351aca3855b1ce6c9c08b0c586ad838fa0f6e636..cfa2c6a3b81faf08ea400eff3219310fe24657a6
--- a/loop.sh
+++ b/loop.sh
@@ -4,7 +4,7 @@
 usage() {
   cat <<'EOF'
 Usage:
-  ./loop.sh run <queue> [--repo DIR] [--max-ticks N] [--dry-run]
+  ./loop.sh run <queue> [--repo DIR] [--max-ticks N] [--review] [--max-review-rounds N] [--dry-run]
 
 The queue is usually .loop/<name>/QUEUE.md in the target repo. Use --repo when the
 queue lives outside the repository it should operate on.
@@ -204,6 +204,86 @@
   } >> "$evidence"
 }
 
+extract_actionable_count() {
+  awk '
+    /^[[:space:]]*-[[:space:]]*actionable:[[:space:]]*[0-9]+[[:space:]]*$/ {
+      sub(/^[[:space:]]*-[[:space:]]*actionable:[[:space:]]*/, "")
+      sub(/[[:space:]]*$/, "")
+      print
+      found = 1
+      exit
+    }
+    END { if (!found) exit 1 }
+  ' "$1"
+}
+
+write_review_prompt() {
+  local out=$1
+  local template="$script_dir/prompts/reviewer.md"
+  if [[ -f "$template" ]]; then
+    cat "$template" > "$out"
+  else
+    cat > "$out" <<'EOF'
+# Knack Reviewer
+
+Load and follow the **review** skill in `.agents/skills/review/`.
+Review the completed queue against the current repository state. Write the structured review artifact at the path provided below.
+EOF
+  fi
+  cat >> "$out" <<EOF
+
+Queue: $queue_abs
+Evidence: $evidence
+Review output: $review_file
+EOF
+}
+
+write_fix_prompt() {
+  local out=$1
+  local template="$script_dir/prompts/fixer.md"
+  if [[ -f "$template" ]]; then
+    cat "$template" > "$out"
+  else
+    cat > "$out" <<'EOF'
+# Knack Fixer
+
+Load and follow the **fix** skill in `.agents/skills/fix/`.
+Read the structured review artifact and append any actionable fix work units to the existing queue. Stop after updating the queue.
+EOF
+  fi
+  cat >> "$out" <<EOF
+
+Queue: $queue_abs
+Evidence: $evidence
+Review input: $review_file
+EOF
+}
+
+run_phase_agent() {
+  local phase=$1 prompt=$2 output=$3 cmd=${4:-}
+  if [[ -n "$cmd" ]]; then
+    (
+      cd "$repo_dir"
+      LOOP_PHASE="$phase" \
+        LOOP_PROMPT_FILE="$prompt" \
+        LOOP_QUEUE_FILE="$queue_abs" \
+        LOOP_EVIDENCE_FILE="$evidence" \
+        LOOP_REVIEW_FILE="$review_file" \
+        bash -lc "$cmd"
+    ) > "$output" 2>&1
+  else
+    (
+      cd "$repo_dir"
+      LOOP_PHASE="$phase" \
+        LOOP_PROMPT_FILE="$prompt" \
+        LOOP_QUEUE_FILE="$queue_abs" \
+        LOOP_EVIDENCE_FILE="$evidence" \
+        LOOP_REVIEW_FILE="$review_file" \
+        pi -p --no-session --approve "$(cat "$prompt")"
+    ) > "$output" 2>&1
+  fi
+}
+
 [[ $# -ge 1 ]] || { usage; exit 1; }
 cmd=$1
 shift
@@ -215,6 +295,8 @@
 [[ -f "$queue_abs" ]] || die "queue not found: $queue_abs"
 
 max_ticks=3
+max_review_rounds=2
+review_enabled=0
 dry_run=0
 repo_override=""
 while [[ $# -gt 0 ]]; do
@@ -229,6 +311,15 @@
       max_ticks=$2
       shift 2
       ;;
+    --review)
+      review_enabled=1
+      shift
+      ;;
+    --max-review-rounds)
+      [[ $# -ge 2 ]] || die "--max-review-rounds needs a value"
+      max_review_rounds=$2
+      shift 2
+      ;;
     --dry-run)
       dry_run=1
       shift
@@ -252,6 +343,7 @@
 fi
 evidence="$queue_dir/EVIDENCE.md"
 handoff="$queue_dir/HANDOFF.md"
+review_file="$queue_dir/REVIEW.md"
 no_progress_strikes=0
 
 write_handoff_on_exit() {
@@ -263,9 +355,13 @@
 }
 trap write_handoff_on_exit EXIT
 
-for ((tick = 1; tick <= max_ticks; tick++)); do
+tick=1
+review_round=0
+
+while true; do
+while (( tick <= max_ticks )); do
   unit=$(first_pending_unit "$queue_abs")
-  [[ -n "$unit" ]] || { echo "knack: no pending work units"; exit 0; }
+  [[ -n "$unit" ]] || break
 
   unit_file=$(mktemp)
   verify_file=$(mktemp)
@@ -307,11 +403,7 @@
   fi
 
   set +e
-  if [[ -n "$agent_cmd" ]]; then
-    (cd "$repo_dir" && LOOP_PROMPT_FILE="$run_prompt" bash -lc "$agent_cmd") > "$agent_out" 2>&1
-  else
-    (cd "$repo_dir" && pi -p --no-session --approve "$(cat "$run_prompt")") > "$agent_out" 2>&1
-  fi
+  run_phase_agent "build" "$run_prompt" "$agent_out" "$agent_cmd"
   agent_code=$?
   set -e
 
@@ -334,6 +426,7 @@
     set_status "$queue_abs" "$title" "done"
     append_evidence "$evidence" "$title" "done" "$verify" "$verify_out" "$agent_out" "$repo_dir" "$unit_file"
     echo "knack: verified — $title"
+    tick=$((tick + 1))
     continue
   fi
 
@@ -347,6 +440,7 @@
     set_status "$queue_abs" "$title" "pending"
     append_evidence "$evidence" "$title" "verify_failed" "$verify" "$verify_out" "$agent_out" "$repo_dir" "$unit_file"
     echo "knack: verify failed with no progress; retrying once"
+    tick=$((tick + 1))
     continue
   fi
 
@@ -360,4 +454,63 @@
   die "reached max ticks ($max_ticks) with pending work"
 fi
 
-echo "knack: reached max ticks ($max_ticks)"
+if [[ "$review_enabled" == 0 ]]; then
+  if (( tick > max_ticks )); then
+    echo "knack: reached max ticks ($max_ticks)"
+    exit 0
+  fi
+  echo "knack: no pending work units"
+  exit 0
+fi
+
+if (( review_round >= max_review_rounds )); then
+  die "reached max review rounds ($max_review_rounds)"
+fi
+
+review_round=$((review_round + 1))
+review_prompt=$(mktemp)
+review_out=$(mktemp)
+write_review_prompt "$review_prompt"
+review_cmd="${LOOP_REVIEW_CMD:-${LOOP_AGENT_CMD:-}}"
+
+echo "knack: review round $review_round/$max_review_rounds"
+set +e
+run_phase_agent "review" "$review_prompt" "$review_out" "$review_cmd"
+review_code=$?
+set -e
+if [[ $review_code -ne 0 ]]; then
+  cat "$review_out"
+  die "review worker exited nonzero"
+fi
+[[ -f "$review_file" ]] || die "review worker did not write $review_file"
+actionable=$(extract_actionable_count "$review_file") || die "review file has no actionable summary: $review_file"
+
+if [[ "$actionable" == 0 ]]; then
+  echo "knack: review clean"
+  exit 0
+fi
+
+if (( review_round >= max_review_rounds )); then
+  die "review found $actionable actionable issue(s) at max review rounds ($max_review_rounds)"
+fi
+
+fix_prompt=$(mktemp)
+fix_out=$(mktemp)
+write_fix_prompt "$fix_prompt"
+fix_cmd="${LOOP_FIX_CMD:-${LOOP_AGENT_CMD:-}}"
+
+echo "knack: fix round $review_round/$max_review_rounds — actionable: $actionable"
+set +e
+run_phase_agent "fix" "$fix_prompt" "$fix_out" "$fix_cmd"
+fix_code=$?
+set -e
+if [[ $fix_code -ne 0 ]]; then
+  cat "$fix_out"
+  die "fix worker exited nonzero"
+fi
+
+if [[ -z "$(first_pending_unit "$queue_abs")" ]]; then
+  cat "$fix_out"
+  die "fix worker generated no pending work units"
+fi
+done
diff --git a/tests/run.sh b/tests/run.sh
index eb28d49f637f95de4c6e555bb98b1fe43455e7f2..7a28bc257ecfc567e30366b85efdea03ea9a59c3
--- a/tests/run.sh
+++ b/tests/run.sh
@@ -169,6 +169,123 @@
 test -f "$repo_lpf/captured-prompt.txt"
 assert_contains "$repo_lpf/captured-prompt.txt" "the test fixture reaches its verify condition"
 
+# Review-fix loop with fake build, review, and fix workers.
+repo_review="$tmp/repo-review"
+mkdir -p "$repo_review/.loop"
+cat > "$repo_review/.loop/QUEUE.md" <<'EOF'
+# Loop Queue: review cycle
+
+Goal:
+Exercise build, review, fix, and review again.
+
+Stop condition:
+The generated app is fixed and review reports no actionable issues.
+
+## the initial build creates a reviewable app file
+
+Read first:
+- This queue file.
+
+Constraints:
+- Leave the bug for review to find.
+
+Verify:
+```bash
+test -f app.txt
+```
+
+Done means:
+- app.txt exists.
+
+Status: pending
+EOF
+
+cat > "$repo_review/build-worker.sh" <<'EOF'
+#!/usr/bin/env bash
+set -euo pipefail
+if grep -q "the fix unit repairs the bug" "$LOOP_PROMPT_FILE"; then
+  printf 'fixed\n' > app.txt
+  echo "build fixed app"
+else
+  printf 'bug\n' > app.txt
+  echo "build created app with bug"
+fi
+EOF
+chmod +x "$repo_review/build-worker.sh"
+
+cat > "$repo_review/review-worker.sh" <<'EOF'
+#!/usr/bin/env bash
+set -euo pipefail
+count=0
+if [[ -f review-count.txt ]]; then
+  count=$(cat review-count.txt)
+fi
+count=$((count + 1))
+printf '%s\n' "$count" > review-count.txt
+
+actionable=1
+if [[ -f app.txt ]] && grep -qx 'fixed' app.txt; then
+  actionable=0
+fi
+
+cat > "$LOOP_REVIEW_FILE" <<EOF_REVIEW
+# Review: fake
+
+## Standards
+
+## Intent
+- actionable | high — app.txt must say fixed
+  evidence: app.txt:1
+
+## Speculative
+
+## Summary
+- actionable: $actionable
+- trivial: 0
+- disputed: 0
+- deferred: 0
+EOF_REVIEW
+echo "review actionable: $actionable"
+EOF
+chmod +x "$repo_review/review-worker.sh"
+
+cat > "$repo_review/fix-worker.sh" <<'EOF'
+#!/usr/bin/env bash
+set -euo pipefail
+cat >> "$LOOP_QUEUE_FILE" <<'EOF_QUEUE'
+
+## the fix unit repairs the bug
+
+Read first:
+- .loop/REVIEW.md
+
+Constraints:
+- Preserve the app file created by the first unit.
+
+Verify:
+```bash
+grep -qx fixed app.txt
+```
+
+Done means:
+- app.txt contains fixed.
+
+Status: pending
+EOF_QUEUE
+echo "fix appended unit"
+EOF
+chmod +x "$repo_review/fix-worker.sh"
+
+LOOP_AGENT_CMD="$repo_review/build-worker.sh" \
+  LOOP_REVIEW_CMD="$repo_review/review-worker.sh" \
+  LOOP_FIX_CMD="$repo_review/fix-worker.sh" \
+  "$root/loop.sh" run "$repo_review/.loop/QUEUE.md" --review --max-ticks 2 >/tmp/loop-review.txt
+assert_contains "$repo_review/.loop/QUEUE.md" "## the fix unit repairs the bug"
+assert_contains "$repo_review/.loop/QUEUE.md" "Status: done"
+assert_contains "$repo_review/.loop/REVIEW.md" "- actionable: 0"
+assert_contains "$repo_review/review-count.txt" "2"
+assert_contains "$repo_review/app.txt" "fixed"
+
 if command -v skills-ref >/dev/null 2>&1; then
   for skill_dir in "$root/.agents/skills"/*; do
     if [[ -d "$skill_dir" ]]; then

 succeeded in 3595ms:
Valid skill: /home/daniel/build/knack/.agents/skills/build
Valid skill: /home/daniel/build/knack/.agents/skills/decide
Valid skill: /home/daniel/build/knack/.agents/skills/domain-modeling
Valid skill: /home/daniel/build/knack/.agents/skills/explore
Valid skill: /home/daniel/build/knack/.agents/skills/fix
Valid skill: /home/daniel/build/knack/.agents/skills/plan
Valid skill: /home/daniel/build/knack/.agents/skills/review
knack tests passed

codex
Verification is still green. I’m checking the final status so the handoff separates my changes from the pre-existing dirty files.
exec
/usr/bin/bash -lc 'git status --short' in /home/daniel/build/knack
 succeeded in 0ms:
 M .agents/skills/fix/SKILL.md
 M .loop/review-loop/EVIDENCE.md
 M .loop/review-loop/HANDOFF.md
 M .loop/review-loop/QUEUE.md
 M cli/embedded/skills/fix/SKILL.md
 M loop.sh
 M tests/run.sh

exec
/usr/bin/bash -lc 'git diff --stat -- loop.sh tests/run.sh' in /home/daniel/build/knack
 succeeded in 0ms:
 loop.sh      | 171 +++++++++++++++++++++++++++++++++++++++++++++++++++++++----
 tests/run.sh | 117 ++++++++++++++++++++++++++++++++++++++++
 2 files changed, 279 insertions(+), 9 deletions(-)

diff --git a/loop.sh b/loop.sh
index 351aca3855b1ce6c9c08b0c586ad838fa0f6e636..cfa2c6a3b81faf08ea400eff3219310fe24657a6
--- a/loop.sh
+++ b/loop.sh
@@ -4,7 +4,7 @@
 usage() {
   cat <<'EOF'
 Usage:
-  ./loop.sh run <queue> [--repo DIR] [--max-ticks N] [--dry-run]
+  ./loop.sh run <queue> [--repo DIR] [--max-ticks N] [--review] [--max-review-rounds N] [--dry-run]
 
 The queue is usually .loop/<name>/QUEUE.md in the target repo. Use --repo when the
 queue lives outside the repository it should operate on.
@@ -204,6 +204,86 @@
   } >> "$evidence"
 }
 
+extract_actionable_count() {
+  awk '
+    /^[[:space:]]*-[[:space:]]*actionable:[[:space:]]*[0-9]+[[:space:]]*$/ {
+      sub(/^[[:space:]]*-[[:space:]]*actionable:[[:space:]]*/, "")
+      sub(/[[:space:]]*$/, "")
+      print
+      found = 1
+      exit
+    }
+    END { if (!found) exit 1 }
+  ' "$1"
+}
+
+write_review_prompt() {
+  local out=$1
+  local template="$script_dir/prompts/reviewer.md"
+  if [[ -f "$template" ]]; then
+    cat "$template" > "$out"
+  else
+    cat > "$out" <<'EOF'
+# Knack Reviewer
+
+Load and follow the **review** skill in `.agents/skills/review/`.
+Review the completed queue against the current repository state. Write the structured review artifact at the path provided below.
+EOF
+  fi
+  cat >> "$out" <<EOF
+
+Queue: $queue_abs
+Evidence: $evidence
+Review output: $review_file
+EOF
+}
+
+write_fix_prompt() {
+  local out=$1
+  local template="$script_dir/prompts/fixer.md"
+  if [[ -f "$template" ]]; then
+    cat "$template" > "$out"
+  else
+    cat > "$out" <<'EOF'
+# Knack Fixer
+
+Load and follow the **fix** skill in `.agents/skills/fix/`.
+Read the structured review artifact and append any actionable fix work units to the existing queue. Stop after updating the queue.
+EOF
+  fi
+  cat >> "$out" <<EOF
+
+Queue: $queue_abs
+Evidence: $evidence
+Review input: $review_file
+EOF
+}
+
+run_phase_agent() {
+  local phase=$1 prompt=$2 output=$3 cmd=${4:-}
+  if [[ -n "$cmd" ]]; then
+    (
+      cd "$repo_dir"
+      LOOP_PHASE="$phase" \
+        LOOP_PROMPT_FILE="$prompt" \
+        LOOP_QUEUE_FILE="$queue_abs" \
+        LOOP_EVIDENCE_FILE="$evidence" \
+        LOOP_REVIEW_FILE="$review_file" \
+        bash -lc "$cmd"
+    ) > "$output" 2>&1
+  else
+    (
+      cd "$repo_dir"
+      LOOP_PHASE="$phase" \
+        LOOP_PROMPT_FILE="$prompt" \
+        LOOP_QUEUE_FILE="$queue_abs" \
+        LOOP_EVIDENCE_FILE="$evidence" \
+        LOOP_REVIEW_FILE="$review_file" \
+        pi -p --no-session --approve "$(cat "$prompt")"
+    ) > "$output" 2>&1
+  fi
+}
+
 [[ $# -ge 1 ]] || { usage; exit 1; }
 cmd=$1
 shift
@@ -215,6 +295,8 @@
 [[ -f "$queue_abs" ]] || die "queue not found: $queue_abs"
 
 max_ticks=3
+max_review_rounds=2
+review_enabled=0
 dry_run=0
 repo_override=""
 while [[ $# -gt 0 ]]; do
@@ -229,6 +311,15 @@
       max_ticks=$2
       shift 2
       ;;
+    --review)
+      review_enabled=1
+      shift
+      ;;
+    --max-review-rounds)
+      [[ $# -ge 2 ]] || die "--max-review-rounds needs a value"
+      max_review_rounds=$2
+      shift 2
+      ;;
     --dry-run)
       dry_run=1
       shift
@@ -252,6 +343,7 @@
 fi
 evidence="$queue_dir/EVIDENCE.md"
 handoff="$queue_dir/HANDOFF.md"
+review_file="$queue_dir/REVIEW.md"
 no_progress_strikes=0
 
 write_handoff_on_exit() {
@@ -263,9 +355,13 @@
 }
 trap write_handoff_on_exit EXIT
 
-for ((tick = 1; tick <= max_ticks; tick++)); do
+tick=1
+review_round=0
+
+while true; do
+while (( tick <= max_ticks )); do
   unit=$(first_pending_unit "$queue_abs")
-  [[ -n "$unit" ]] || { echo "knack: no pending work units"; exit 0; }
+  [[ -n "$unit" ]] || break
 
   unit_file=$(mktemp)
   verify_file=$(mktemp)
@@ -307,11 +403,7 @@
   fi
 
   set +e
-  if [[ -n "$agent_cmd" ]]; then
-    (cd "$repo_dir" && LOOP_PROMPT_FILE="$run_prompt" bash -lc "$agent_cmd") > "$agent_out" 2>&1
-  else
-    (cd "$repo_dir" && pi -p --no-session --approve "$(cat "$run_prompt")") > "$agent_out" 2>&1
-  fi
+  run_phase_agent "build" "$run_prompt" "$agent_out" "$agent_cmd"
   agent_code=$?
   set -e
 
@@ -334,6 +426,7 @@
     set_status "$queue_abs" "$title" "done"
     append_evidence "$evidence" "$title" "done" "$verify" "$verify_out" "$agent_out" "$repo_dir" "$unit_file"
     echo "knack: verified — $title"
+    tick=$((tick + 1))
     continue
   fi
 
@@ -347,6 +440,7 @@
     set_status "$queue_abs" "$title" "pending"
     append_evidence "$evidence" "$title" "verify_failed" "$verify" "$verify_out" "$agent_out" "$repo_dir" "$unit_file"
     echo "knack: verify failed with no progress; retrying once"
+    tick=$((tick + 1))
     continue
   fi
 
@@ -360,4 +454,63 @@
   die "reached max ticks ($max_ticks) with pending work"
 fi
 
-echo "knack: reached max ticks ($max_ticks)"
+if [[ "$review_enabled" == 0 ]]; then
+  if (( tick > max_ticks )); then
+    echo "knack: reached max ticks ($max_ticks)"
+    exit 0
+  fi
+  echo "knack: no pending work units"
+  exit 0
+fi
+
+if (( review_round >= max_review_rounds )); then
+  die "reached max review rounds ($max_review_rounds)"
+fi
+
+review_round=$((review_round + 1))
+review_prompt=$(mktemp)
+review_out=$(mktemp)
+write_review_prompt "$review_prompt"
+review_cmd="${LOOP_REVIEW_CMD:-${LOOP_AGENT_CMD:-}}"
+
+echo "knack: review round $review_round/$max_review_rounds"
+set +e
+run_phase_agent "review" "$review_prompt" "$review_out" "$review_cmd"
+review_code=$?
+set -e
+if [[ $review_code -ne 0 ]]; then
+  cat "$review_out"
+  die "review worker exited nonzero"
+fi
+[[ -f "$review_file" ]] || die "review worker did not write $review_file"
+actionable=$(extract_actionable_count "$review_file") || die "review file has no actionable summary: $review_file"
+
+if [[ "$actionable" == 0 ]]; then
+  echo "knack: review clean"
+  exit 0
+fi
+
+if (( review_round >= max_review_rounds )); then
+  die "review found $actionable actionable issue(s) at max review rounds ($max_review_rounds)"
+fi
+
+fix_prompt=$(mktemp)
+fix_out=$(mktemp)
+write_fix_prompt "$fix_prompt"
+fix_cmd="${LOOP_FIX_CMD:-${LOOP_AGENT_CMD:-}}"
+
+echo "knack: fix round $review_round/$max_review_rounds — actionable: $actionable"
+set +e
+run_phase_agent "fix" "$fix_prompt" "$fix_out" "$fix_cmd"
+fix_code=$?
+set -e
+if [[ $fix_code -ne 0 ]]; then
+  cat "$fix_out"
+  die "fix worker exited nonzero"
+fi
+
+if [[ -z "$(first_pending_unit "$queue_abs")" ]]; then
+  cat "$fix_out"
+  die "fix worker generated no pending work units"
+fi
+done
diff --git a/tests/run.sh b/tests/run.sh
index eb28d49f637f95de4c6e555bb98b1fe43455e7f2..7a28bc257ecfc567e30366b85efdea03ea9a59c3
--- a/tests/run.sh
+++ b/tests/run.sh
@@ -169,6 +169,123 @@
 test -f "$repo_lpf/captured-prompt.txt"
 assert_contains "$repo_lpf/captured-prompt.txt" "the test fixture reaches its verify condition"
 
+# Review-fix loop with fake build, review, and fix workers.
+repo_review="$tmp/repo-review"
+mkdir -p "$repo_review/.loop"
+cat > "$repo_review/.loop/QUEUE.md" <<'EOF'
+# Loop Queue: review cycle
+
+Goal:
+Exercise build, review, fix, and review again.
+
+Stop condition:
+The generated app is fixed and review reports no actionable issues.
+
+## the initial build creates a reviewable app file
+
+Read first:
+- This queue file.
+
+Constraints:
+- Leave the bug for review to find.
+
+Verify:
+```bash
+test -f app.txt
+```
+
+Done means:
+- app.txt exists.
+
+Status: pending
+EOF
+
+cat > "$repo_review/build-worker.sh" <<'EOF'
+#!/usr/bin/env bash
+set -euo pipefail
+if grep -q "the fix unit repairs the bug" "$LOOP_PROMPT_FILE"; then
+  printf 'fixed\n' > app.txt
+  echo "build fixed app"
+else
+  printf 'bug\n' > app.txt
+  echo "build created app with bug"
+fi
+EOF
+chmod +x "$repo_review/build-worker.sh"
+
+cat > "$repo_review/review-worker.sh" <<'EOF'
+#!/usr/bin/env bash
+set -euo pipefail
+count=0
+if [[ -f review-count.txt ]]; then
+  count=$(cat review-count.txt)
+fi
+count=$((count + 1))
+printf '%s\n' "$count" > review-count.txt
+
+actionable=1
+if [[ -f app.txt ]] && grep -qx 'fixed' app.txt; then
+  actionable=0
+fi
+
+cat > "$LOOP_REVIEW_FILE" <<EOF_REVIEW
+# Review: fake
+
+## Standards
+
+## Intent
+- actionable | high — app.txt must say fixed
+  evidence: app.txt:1
+
+## Speculative
+
+## Summary
+- actionable: $actionable
+- trivial: 0
+- disputed: 0
+- deferred: 0
+EOF_REVIEW
+echo "review actionable: $actionable"
+EOF
+chmod +x "$repo_review/review-worker.sh"
+
+cat > "$repo_review/fix-worker.sh" <<'EOF'
+#!/usr/bin/env bash
+set -euo pipefail
+cat >> "$LOOP_QUEUE_FILE" <<'EOF_QUEUE'
+
+## the fix unit repairs the bug
+
+Read first:
+- .loop/REVIEW.md
+
+Constraints:
+- Preserve the app file created by the first unit.
+
+Verify:
+```bash
+grep -qx fixed app.txt
+```
+
+Done means:
+- app.txt contains fixed.
+
+Status: pending
+EOF_QUEUE
+echo "fix appended unit"
+EOF
+chmod +x "$repo_review/fix-worker.sh"
+
+LOOP_AGENT_CMD="$repo_review/build-worker.sh" \
+  LOOP_REVIEW_CMD="$repo_review/review-worker.sh" \
+  LOOP_FIX_CMD="$repo_review/fix-worker.sh" \
+  "$root/loop.sh" run "$repo_review/.loop/QUEUE.md" --review --max-ticks 2 >/tmp/loop-review.txt
+assert_contains "$repo_review/.loop/QUEUE.md" "## the fix unit repairs the bug"
+assert_contains "$repo_review/.loop/QUEUE.md" "Status: done"
+assert_contains "$repo_review/.loop/REVIEW.md" "- actionable: 0"
+assert_contains "$repo_review/review-count.txt" "2"
+assert_contains "$repo_review/app.txt" "fixed"
+
 if command -v skills-ref >/dev/null 2>&1; then
   for skill_dir in "$root/.agents/skills"/*; do
     if [[ -d "$skill_dir" ]]; then

codex
```text
Unit: loop.sh orchestrates bounded build-review-fix rounds
Changed: loop.sh review/fix orchestration; tests/run.sh fake-worker review-fix cycle
Verify expected: bash -n /home/daniel/build/knack/loop.sh && /home/daniel/build/knack/tests/run.sh
Notes: Verify command passed locally. Pre-existing dirty files in .agents/skills/fix, cli/embedded/skills/fix, and .loop/review-loop artifacts were left untouched.
```
diff --git a/loop.sh b/loop.sh
index 351aca3855b1ce6c9c08b0c586ad838fa0f6e636..cfa2c6a3b81faf08ea400eff3219310fe24657a6
--- a/loop.sh
+++ b/loop.sh
@@ -4,7 +4,7 @@
 usage() {
   cat <<'EOF'
 Usage:
-  ./loop.sh run <queue> [--repo DIR] [--max-ticks N] [--dry-run]
+  ./loop.sh run <queue> [--repo DIR] [--max-ticks N] [--review] [--max-review-rounds N] [--dry-run]
 
 The queue is usually .loop/<name>/QUEUE.md in the target repo. Use --repo when the
 queue lives outside the repository it should operate on.
@@ -204,6 +204,86 @@
   } >> "$evidence"
 }
 
+extract_actionable_count() {
+  awk '
+    /^[[:space:]]*-[[:space:]]*actionable:[[:space:]]*[0-9]+[[:space:]]*$/ {
+      sub(/^[[:space:]]*-[[:space:]]*actionable:[[:space:]]*/, "")
+      sub(/[[:space:]]*$/, "")
+      print
+      found = 1
+      exit
+    }
+    END { if (!found) exit 1 }
+  ' "$1"
+}
+
+write_review_prompt() {
+  local out=$1
+  local template="$script_dir/prompts/reviewer.md"
+  if [[ -f "$template" ]]; then
+    cat "$template" > "$out"
+  else
+    cat > "$out" <<'EOF'
+# Knack Reviewer
+
+Load and follow the **review** skill in `.agents/skills/review/`.
+Review the completed queue against the current repository state. Write the structured review artifact at the path provided below.
+EOF
+  fi
+  cat >> "$out" <<EOF
+
+Queue: $queue_abs
+Evidence: $evidence
+Review output: $review_file
+EOF
+}
+
+write_fix_prompt() {
+  local out=$1
+  local template="$script_dir/prompts/fixer.md"
+  if [[ -f "$template" ]]; then
+    cat "$template" > "$out"
+  else
+    cat > "$out" <<'EOF'
+# Knack Fixer
+
+Load and follow the **fix** skill in `.agents/skills/fix/`.
+Read the structured review artifact and append any actionable fix work units to the existing queue. Stop after updating the queue.
+EOF
+  fi
+  cat >> "$out" <<EOF
+
+Queue: $queue_abs
+Evidence: $evidence
+Review input: $review_file
+EOF
+}
+
+run_phase_agent() {
+  local phase=$1 prompt=$2 output=$3 cmd=${4:-}
+  if [[ -n "$cmd" ]]; then
+    (
+      cd "$repo_dir"
+      LOOP_PHASE="$phase" \
+        LOOP_PROMPT_FILE="$prompt" \
+        LOOP_QUEUE_FILE="$queue_abs" \
+        LOOP_EVIDENCE_FILE="$evidence" \
+        LOOP_REVIEW_FILE="$review_file" \
+        bash -lc "$cmd"
+    ) > "$output" 2>&1
+  else
+    (
+      cd "$repo_dir"
+      LOOP_PHASE="$phase" \
+        LOOP_PROMPT_FILE="$prompt" \
+        LOOP_QUEUE_FILE="$queue_abs" \
+        LOOP_EVIDENCE_FILE="$evidence" \
+        LOOP_REVIEW_FILE="$review_file" \
+        pi -p --no-session --approve "$(cat "$prompt")"
+    ) > "$output" 2>&1
+  fi
+}
+
 [[ $# -ge 1 ]] || { usage; exit 1; }
 cmd=$1
 shift
@@ -215,6 +295,8 @@
 [[ -f "$queue_abs" ]] || die "queue not found: $queue_abs"
 
 max_ticks=3
+max_review_rounds=2
+review_enabled=0
 dry_run=0
 repo_override=""
 while [[ $# -gt 0 ]]; do
@@ -229,6 +311,15 @@
       max_ticks=$2
       shift 2
       ;;
+    --review)
+      review_enabled=1
+      shift
+      ;;
+    --max-review-rounds)
+      [[ $# -ge 2 ]] || die "--max-review-rounds needs a value"
+      max_review_rounds=$2
+      shift 2
+      ;;
     --dry-run)
       dry_run=1
       shift
@@ -252,6 +343,7 @@
 fi
 evidence="$queue_dir/EVIDENCE.md"
 handoff="$queue_dir/HANDOFF.md"
+review_file="$queue_dir/REVIEW.md"
 no_progress_strikes=0
 
 write_handoff_on_exit() {
@@ -263,9 +355,13 @@
 }
 trap write_handoff_on_exit EXIT
 
-for ((tick = 1; tick <= max_ticks; tick++)); do
+tick=1
+review_round=0
+
+while true; do
+while (( tick <= max_ticks )); do
   unit=$(first_pending_unit "$queue_abs")
-  [[ -n "$unit" ]] || { echo "knack: no pending work units"; exit 0; }
+  [[ -n "$unit" ]] || break
 
   unit_file=$(mktemp)
   verify_file=$(mktemp)
@@ -307,11 +403,7 @@
   fi
 
   set +e
-  if [[ -n "$agent_cmd" ]]; then
-    (cd "$repo_dir" && LOOP_PROMPT_FILE="$run_prompt" bash -lc "$agent_cmd") > "$agent_out" 2>&1
-  else
-    (cd "$repo_dir" && pi -p --no-session --approve "$(cat "$run_prompt")") > "$agent_out" 2>&1
-  fi
+  run_phase_agent "build" "$run_prompt" "$agent_out" "$agent_cmd"
   agent_code=$?
   set -e
 
@@ -334,6 +426,7 @@
     set_status "$queue_abs" "$title" "done"
     append_evidence "$evidence" "$title" "done" "$verify" "$verify_out" "$agent_out" "$repo_dir" "$unit_file"
     echo "knack: verified — $title"
+    tick=$((tick + 1))
     continue
   fi
 
@@ -347,6 +440,7 @@
     set_status "$queue_abs" "$title" "pending"
     append_evidence "$evidence" "$title" "verify_failed" "$verify" "$verify_out" "$agent_out" "$repo_dir" "$unit_file"
     echo "knack: verify failed with no progress; retrying once"
+    tick=$((tick + 1))
     continue
   fi
 
@@ -360,4 +454,63 @@
   die "reached max ticks ($max_ticks) with pending work"
 fi
 
-echo "knack: reached max ticks ($max_ticks)"
+if [[ "$review_enabled" == 0 ]]; then
+  if (( tick > max_ticks )); then
+    echo "knack: reached max ticks ($max_ticks)"
+    exit 0
+  fi
+  echo "knack: no pending work units"
+  exit 0
+fi
+
+if (( review_round >= max_review_rounds )); then
+  die "reached max review rounds ($max_review_rounds)"
+fi
+
+review_round=$((review_round + 1))
+review_prompt=$(mktemp)
+review_out=$(mktemp)
+write_review_prompt "$review_prompt"
+review_cmd="${LOOP_REVIEW_CMD:-${LOOP_AGENT_CMD:-}}"
+
+echo "knack: review round $review_round/$max_review_rounds"
+set +e
+run_phase_agent "review" "$review_prompt" "$review_out" "$review_cmd"
+review_code=$?
+set -e
+if [[ $review_code -ne 0 ]]; then
+  cat "$review_out"
+  die "review worker exited nonzero"
+fi
+[[ -f "$review_file" ]] || die "review worker did not write $review_file"
+actionable=$(extract_actionable_count "$review_file") || die "review file has no actionable summary: $review_file"
+
+if [[ "$actionable" == 0 ]]; then
+  echo "knack: review clean"
+  exit 0
+fi
+
+if (( review_round >= max_review_rounds )); then
+  die "review found $actionable actionable issue(s) at max review rounds ($max_review_rounds)"
+fi
+
+fix_prompt=$(mktemp)
+fix_out=$(mktemp)
+write_fix_prompt "$fix_prompt"
+fix_cmd="${LOOP_FIX_CMD:-${LOOP_AGENT_CMD:-}}"
+
+echo "knack: fix round $review_round/$max_review_rounds — actionable: $actionable"
+set +e
+run_phase_agent "fix" "$fix_prompt" "$fix_out" "$fix_cmd"
+fix_code=$?
+set -e
+if [[ $fix_code -ne 0 ]]; then
+  cat "$fix_out"
+  die "fix worker exited nonzero"
+fi
+
+if [[ -z "$(first_pending_unit "$queue_abs")" ]]; then
+  cat "$fix_out"
+  die "fix worker generated no pending work units"
+fi
+done
diff --git a/tests/run.sh b/tests/run.sh
index eb28d49f637f95de4c6e555bb98b1fe43455e7f2..7a28bc257ecfc567e30366b85efdea03ea9a59c3
--- a/tests/run.sh
+++ b/tests/run.sh
@@ -169,6 +169,123 @@
 test -f "$repo_lpf/captured-prompt.txt"
 assert_contains "$repo_lpf/captured-prompt.txt" "the test fixture reaches its verify condition"
 
+# Review-fix loop with fake build, review, and fix workers.
+repo_review="$tmp/repo-review"
+mkdir -p "$repo_review/.loop"
+cat > "$repo_review/.loop/QUEUE.md" <<'EOF'
+# Loop Queue: review cycle
+
+Goal:
+Exercise build, review, fix, and review again.
+
+Stop condition:
+The generated app is fixed and review reports no actionable issues.
+
+## the initial build creates a reviewable app file
+
+Read first:
+- This queue file.
+
+Constraints:
+- Leave the bug for review to find.
+
+Verify:
+```bash
+test -f app.txt
+```
+
+Done means:
+- app.txt exists.
+
+Status: pending
+EOF
+
+cat > "$repo_review/build-worker.sh" <<'EOF'
+#!/usr/bin/env bash
+set -euo pipefail
+if grep -q "the fix unit repairs the bug" "$LOOP_PROMPT_FILE"; then
+  printf 'fixed\n' > app.txt
+  echo "build fixed app"
+else
+  printf 'bug\n' > app.txt
+  echo "build created app with bug"
+fi
+EOF
+chmod +x "$repo_review/build-worker.sh"
+
+cat > "$repo_review/review-worker.sh" <<'EOF'
+#!/usr/bin/env bash
+set -euo pipefail
+count=0
+if [[ -f review-count.txt ]]; then
+  count=$(cat review-count.txt)
+fi
+count=$((count + 1))
+printf '%s\n' "$count" > review-count.txt
+
+actionable=1
+if [[ -f app.txt ]] && grep -qx 'fixed' app.txt; then
+  actionable=0
+fi
+
+cat > "$LOOP_REVIEW_FILE" <<EOF_REVIEW
+# Review: fake
+
+## Standards
+
+## Intent
+- actionable | high — app.txt must say fixed
+  evidence: app.txt:1
+
+## Speculative
+
+## Summary
+- actionable: $actionable
+- trivial: 0
+- disputed: 0
+- deferred: 0
+EOF_REVIEW
+echo "review actionable: $actionable"
+EOF
+chmod +x "$repo_review/review-worker.sh"
+
+cat > "$repo_review/fix-worker.sh" <<'EOF'
+#!/usr/bin/env bash
+set -euo pipefail
+cat >> "$LOOP_QUEUE_FILE" <<'EOF_QUEUE'
+
+## the fix unit repairs the bug
+
+Read first:
+- .loop/REVIEW.md
+
+Constraints:
+- Preserve the app file created by the first unit.
+
+Verify:
+```bash
+grep -qx fixed app.txt
+```
+
+Done means:
+- app.txt contains fixed.
+
+Status: pending
+EOF_QUEUE
+echo "fix appended unit"
+EOF
+chmod +x "$repo_review/fix-worker.sh"
+
+LOOP_AGENT_CMD="$repo_review/build-worker.sh" \
+  LOOP_REVIEW_CMD="$repo_review/review-worker.sh" \
+  LOOP_FIX_CMD="$repo_review/fix-worker.sh" \
+  "$root/loop.sh" run "$repo_review/.loop/QUEUE.md" --review --max-ticks 2 >/tmp/loop-review.txt
+assert_contains "$repo_review/.loop/QUEUE.md" "## the fix unit repairs the bug"
+assert_contains "$repo_review/.loop/QUEUE.md" "Status: done"
+assert_contains "$repo_review/.loop/REVIEW.md" "- actionable: 0"
+assert_contains "$repo_review/review-count.txt" "2"
+assert_contains "$repo_review/app.txt" "fixed"
+
 if command -v skills-ref >/dev/null 2>&1; then
   for skill_dir in "$root/.agents/skills"/*; do
     if [[ -d "$skill_dir" ]]; then

tokens used
59,132
```text
Unit: loop.sh orchestrates bounded build-review-fix rounds
Changed: loop.sh review/fix orchestration; tests/run.sh fake-worker review-fix cycle
Verify expected: bash -n /home/daniel/build/knack/loop.sh && /home/daniel/build/knack/tests/run.sh
Notes: Verify command passed locally. Pre-existing dirty files in .agents/skills/fix, cli/embedded/skills/fix, and .loop/review-loop artifacts were left untouched.
```
````

What this proves:
- The verify command passed for this work unit in the current repo state.

What remains unverified:
- Anything outside the verify command's proof scope.

## 2026-07-10T19:55:56-06:00 — docs and AGENTS.md reflect the new loop behavior

Status: done

Unit:
````markdown
## docs and AGENTS.md reflect the new loop behavior

Read first:
- `docs/loop.md`
- `docs/skills.md`
- `README.md`
- `AGENTS.md`
- `decisions/0008-loop-orchestrates-review-fix.md`

Constraints:
- Do not change the core thesis.
- Document that `--review` is opt-in.
- Keep changes scoped to the review-fix subloop.
- Sync embedded skills if any skill files changed.

Done means:
- `docs/loop.md` documents `--review`, `--max-review-rounds`, `LOOP_REVIEW_CMD`, `LOOP_FIX_CMD`, and `REVIEW.md`.
- `docs/skills.md` and `README.md` describe review/fix as loop-orchestrated when `--review` is set.
- `AGENTS.md` current state and working conventions mention the new loop capability.
- Tests pass.

Verify:
```bash
cd /home/daniel/build/knack && ./tests/run.sh && cd cli && go test ./...
```

Status: pending````

Files changed:
```text
 M .agents/skills/fix/SKILL.md
 M .loop/review-loop/EVIDENCE.md
 M .loop/review-loop/HANDOFF.md
 M .loop/review-loop/QUEUE.md
 M AGENTS.md
 M README.md
 M cli/embedded/skills/fix/SKILL.md
 M docs/loop.md
 M docs/skills.md
 M loop.sh
 M tests/run.sh
```

Verify command:
```bash
cd /home/daniel/build/knack && ./tests/run.sh && cd cli && go test ./...
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
ok  	knack	(cached)
ok  	knack/internal/decisions	(cached)
ok  	knack/internal/glossary	(cached)
ok  	knack/internal/instructions	(cached)
ok  	knack/internal/queue	(cached)
ok  	knack/internal/skills	(cached)
ok  	knack/internal/status	(cached)
```

Worker output:
````text
OpenAI Codex v0.144.1
--------
workdir: /home/daniel/build/knack
model: gpt-5.5
provider: openai
approval: never
sandbox: danger-full-access
reasoning effort: high
reasoning summaries: none
session id: 019f4ee1-92e8-7190-8ddd-68d87801a7a0
--------
user
# Knack Worker

You are one worker tick in a bounded loop. Complete exactly one work unit, then stop.

Load and follow the **build** skill in `.agents/skills/build/` before doing any work.

## Rules

1. Read `AGENTS.md` first if it exists — it contains operational context.
2. Read the current work unit carefully, especially its `Verify` command.
3. Do only the work needed for that unit.
4. Do not start another work unit.
5. Do not mark the unit complete yourself.
6. Do not edit `.loop/<name>/EVIDENCE.md`; the runner writes evidence after verification.
7. If the unit is blocked, make the smallest useful note in your final response and stop.
8. Stay within the runner's hard stops (max ticks, no-progress detection). If the unit is too large for one tick, do as much as keeps the repo working and report what remains.
9. The worker's scope is the unit's outcome plus its constraints. Keep the diff narrow and aligned with that scope.
10. If verification fails while you are working, fix the cause if it belongs to this unit; otherwise stop and report the blocker.

## Success standard

Your job is not to claim success. Your job is to make the repository state satisfy the unit's `Verify` command. The runner will execute that command after you exit.

## Output

End with a compact terminal handoff. This is not `.loop/<name>/HANDOFF.md` — the runner writes that file. This is your own summary to the runner:

```text
Unit: <title>
Changed: <brief file/area list>
Verify expected: <command from unit>
Notes: <blockers or caveats, if any>
```

Current work unit from /home/daniel/build/knack/.loop/review-loop/QUEUE.md:

## docs and AGENTS.md reflect the new loop behavior

Read first:
- `docs/loop.md`
- `docs/skills.md`
- `README.md`
- `AGENTS.md`
- `decisions/0008-loop-orchestrates-review-fix.md`

Constraints:
- Do not change the core thesis.
- Document that `--review` is opt-in.
- Keep changes scoped to the review-fix subloop.
- Sync embedded skills if any skill files changed.

Done means:
- `docs/loop.md` documents `--review`, `--max-review-rounds`, `LOOP_REVIEW_CMD`, `LOOP_FIX_CMD`, and `REVIEW.md`.
- `docs/skills.md` and `README.md` describe review/fix as loop-orchestrated when `--review` is set.
- `AGENTS.md` current state and working conventions mention the new loop capability.
- Tests pass.

Verify:
```bash
cd /home/daniel/build/knack && ./tests/run.sh && cd cli && go test ./...
```

Status: pending
hook: SessionStart
hook: SessionStart Completed
codex
I’ll load the local build skill first, then read the unit’s referenced docs and ADR so the doc changes stay tightly scoped to the review/fix loop behavior.
exec
/usr/bin/bash -lc "sed -n '1,240p' AGENTS.md" in /home/daniel/build/knack
exec
/usr/bin/bash -lc "sed -n '1,240p' .agents/skills/build/SKILL.md" in /home/daniel/build/knack
 succeeded in 0ms:
---
name: build
description: Use when implementing one work unit from a `.loop/<name>/QUEUE.md` queue. The worker-side skill — read the unit, do the work, don't self-certify, end with a handoff. Triggers on "build", "implement", "apply this unit", "do the work", "run the loop", or when a work unit needs to be executed. Also use when the loop invokes the worker for a tick.
---

# Build

Implement one work unit from `.loop/<name>/QUEUE.md`. Do the work. Don't self-certify. The loop runner owns the verify gate — your job is to make the repository state satisfy the unit's `Verify` command, not to claim success.

## Core rules

The runner injects `prompts/worker.md` at the start of every tick. If it is not in your context, read it now. Its `Rules` and `Output` sections are the canonical source for this skill; the sections below elaborate on decisions, operational learnings, blockers, and units that are too large for one tick.

> **Scope note:** Updating `AGENTS.md` or writing an ADR during a tick is a durable trace, not part of the unit's `Verify` scope. Do it only when the tick teaches you something that would save the next session time.

## Scope

The unit's scope is its outcome plus its constraints. The worker determines which files to change and how. The `Verify:` command is the mechanically enforceable subset of `Done means:`.

## Capturing decisions during build

If you discover an architectural ruling while implementing — "we need to handle X this way because Y" — capture it as an ADR using the `decide` skill. Do this inline, not after the unit. Decisions made during implementation are the most valuable kind because they come from the code pushing back.

## Capturing operational learnings

If you learn something about how the project works — a build command, a test convention, a gotcha — that isn't already in `AGENTS.md`, add it. `AGENTS.md` is the living operational context. If it grows stale, the next agent session starts with worse information.

Don't add trivia. Add things that would have saved you time if you'd known them upfront.

## When the unit is too big

If the work unit is larger than what can be done in one tick:

- Do as much as you can while keeping the repo in a working state.
- End with a handoff note explaining what remains.
- The runner will re-queue the unit if verify fails.

Do not try to do everything in one tick if the work genuinely needs more. The loop is designed for multiple ticks.

## When you're blocked

If you hit a blocker you can't resolve within the unit's scope:

1. State the blocker clearly in your final output.
2. Note what would unblock you (a decision, a dependency, a missing file).
3. Stop. Don't thrash.

The runner will mark the unit as `blocked` and write a handoff. The next session can pick up from there.

## Output

The output format is defined in the worker prompt (`prompts/worker.md`). End with a compact terminal handoff there.

 succeeded in 0ms:
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

The loop (`loop.sh`), all seven skills, and the Go CLI are built. The loop supports per-unit `Agent:` overrides, `LOOP_AGENT_CMD` for agent-agnosticism, and writes `.loop/<name>/HANDOFF.md` on non-clean exits. The skills cover the full explore → plan → build → review → fix flow plus two shared skills (decide, domain-modeling). The CLI (`cli/`) provides `skills init|check`, `validate`, `decisions list|show|check`, `status`, `glossary check`, and `instructions`. It embeds the default skills via `go:embed` and scaffolds them into target projects.

## Core artifacts

- `.loop/<name>/QUEUE.md` — disposable queue of verifiable work units.
- `.loop/<name>/EVIDENCE.md` — append-only ledger of what each tick proved. Durable: keep it after deleting QUEUE.md so `decisions check` can still see which ADRs the completed cycle referenced.
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
- When the queue is complete and verified, the cycle's `.loop/<name>/QUEUE.md` is disposable. The human deletes it; the loop does not. `EVIDENCE.md` is the durable ledger — keep it so `decisions check` can still trace which ADRs the cycle referenced.
- A work unit must leave the repo better if the loop stops immediately after it.
- Work units are whatever shape the work is — not forced into "vertical slices."
- `LOOP_AGENT_CMD` overrides the worker invocation (agent-agnostic: pi, claude, codex, opencode, etc.).
- Per-unit `Agent:` field in QUEUE.md overrides `LOOP_AGENT_CMD` for one unit (model routing).
- Work units are `## <outcome>` headers with `Read first:`, `Constraints:`, `Done means:`, and `Verify:` fields. `Done means:` is the acceptance criteria; `Verify:` is the mechanically enforceable subset. The gap between them is the review surface.
- `Read first:` is context, not scope — 2–4 entries of ADRs, code areas, or rulings.
- `Constraints:` state boundaries. A constraint says what must stay true or what is out of bounds, never what to edit. If it names a file, it is "don't touch X" or "X's public API must not change," not "update X."
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

codex
The build skill confirms this tick should only update the current unit and leave verification to the runner. I’m going to compare the docs against ADR-0008 and the actual loop flags before editing.
exec
/usr/bin/bash -lc "sed -n '1,220p' prompts/worker.md" in /home/daniel/build/knack
 succeeded in 0ms:
# Knack Worker

You are one worker tick in a bounded loop. Complete exactly one work unit, then stop.

Load and follow the **build** skill in `.agents/skills/build/` before doing any work.

## Rules

1. Read `AGENTS.md` first if it exists — it contains operational context.
2. Read the current work unit carefully, especially its `Verify` command.
3. Do only the work needed for that unit.
4. Do not start another work unit.
5. Do not mark the unit complete yourself.
6. Do not edit `.loop/<name>/EVIDENCE.md`; the runner writes evidence after verification.
7. If the unit is blocked, make the smallest useful note in your final response and stop.
8. Stay within the runner's hard stops (max ticks, no-progress detection). If the unit is too large for one tick, do as much as keeps the repo working and report what remains.
9. The worker's scope is the unit's outcome plus its constraints. Keep the diff narrow and aligned with that scope.
10. If verification fails while you are working, fix the cause if it belongs to this unit; otherwise stop and report the blocker.

## Success standard

Your job is not to claim success. Your job is to make the repository state satisfy the unit's `Verify` command. The runner will execute that command after you exit.

## Output

End with a compact terminal handoff. This is not `.loop/<name>/HANDOFF.md` — the runner writes that file. This is your own summary to the runner:

```text
Unit: <title>
Changed: <brief file/area list>
Verify expected: <command from unit>
Notes: <blockers or caveats, if any>
```

exec
/usr/bin/bash -lc "sed -n '1,260p' docs/loop.md" in /home/daniel/build/knack
 succeeded in 0ms:
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

- `LOOP_AGENT_CMD` — optional command used to invoke the worker. If unset, the loop defaults to `pi -p --no-session --approve` with the prompt text as a single argument. If set, the command is evaluated by `bash -lc` in the repo directory, and `LOOP_PROMPT_FILE` is set to a temporary file containing `prompts/worker.md` plus the current work unit.
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
LOOP_AGENT_CMD='pi -p --no-session --approve "$(cat "$LOOP_PROMPT_FILE")"' ./loop.sh run .loop/<name>/QUEUE.md
LOOP_AGENT_CMD='claude --print --no-session-persistence --dangerously-skip-permissions "$(cat "$LOOP_PROMPT_FILE")"' ./loop.sh run .loop/<name>/QUEUE.md
LOOP_AGENT_CMD='codex exec --dangerously-bypass-approvals-and-sandbox --ephemeral "$(cat "$LOOP_PROMPT_FILE")"' ./loop.sh run .loop/<name>/QUEUE.md
LOOP_AGENT_CMD='opencode run --auto "$(cat "$LOOP_PROMPT_FILE")"' ./loop.sh run .loop/<name>/QUEUE.md
LOOP_AGENT_CMD='devin --print --prompt-file "$LOOP_PROMPT_FILE" --permission-mode dangerous' ./loop.sh run .loop/<name>/QUEUE.md
```

The command is passed to `bash -lc` in the repo directory, so the typical pattern is `"$(cat "$LOOP_PROMPT_FILE")"`.

## Verification notes

- `loop.sh` does **not** validate `QUEUE.md` structure. Use `knack validate` first.
- `loop.sh` does **not** run review or manage ADRs/glossary. Those are handled by the skills.
- The `Verify:` command must be deterministic and executable by the runner — tests, builds, type checks, not an LLM-as-judge.

exec
/usr/bin/bash -lc "sed -n '1,260p' README.md" in /home/daniel/build/knack
 succeeded in 0ms:
# knack

A tiny loop-packet runner for agentic development.

Compiles human intent into disposable work unit queues, then runs one unit at a time behind deterministic verification gates. Ships with a read-only CLI that validates structure, tracks decisions, and scaffolds the default skill set into any project.

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

## How it works

1. The runner reads the first `Status: pending` work unit from `QUEUE.md`.
2. It marks the unit `in_progress` and invokes a fresh agent context with the worker prompt and the unit.
3. The worker implements the unit and exits. It does **not** self-certify.
4. The runner executes the unit's `Verify` command outside the worker.
5. On success: the unit is marked `done` and evidence is appended.
6. On failure: the unit is marked `verify_failed` or `no_progress`, evidence is appended, and the loop stops or retries once.
7. The loop halts on: max ticks reached with pending work, two no-progress strikes, or a failed verify after a real change.
8. On any non-clean exit, the runner writes `HANDOFF.md` with completed/in-progress/remaining units and the next action.

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

## CLI

The `knack` binary is a read-only validator and context provider. Build it from `cli/`:

```bash
cd cli && go build -o ../knack .
```

### Commands

```
knack skills init [--target DIR]    Scaffold the seven default skills into DIR/.agents/skills/
knack skills check [--dir DIR]      Validate skills in DIR (default: .agents/skills)
knack validate <queue-file>         Validate work-unit structure in a queue file
knack decisions list                List all ADRs in decisions/
knack decisions show NNNN           Print the full text of ADR NNNN
knack decisions check               Flag orphaned ADRs and dangling references
knack status                        Aggregate work-unit counts across all .loop/<name>/ cycles
knack glossary check                Validate glossary.md term references
knack instructions <artifact>       Print a template: work-unit | adr | glossary-entry
```

All commands read from the current directory (run from the repo root). `skills init` is the only write operation — it scaffolds missing skills and leaves existing ones alone.

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

exec
/usr/bin/bash -lc "sed -n '1,240p' decisions/0008-loop-orchestrates-review-fix.md" in /home/daniel/build/knack
exec
/usr/bin/bash -lc "sed -n '1,260p' docs/skills.md" in /home/daniel/build/knack
 succeeded in 0ms:
# 0008: The loop orchestrates the review-fix subloop; skills keep the judgment

Date: 2026-07-10
Status: accepted

## Context

The default skill flow is `explore → plan → build → review → fix`, but `loop.sh` only runs `build` ticks. Review and fix are left as manual skills, so the loop cannot autonomously review a completed queue, turn findings into fix units, re-run the build pass, and stop when review is clean. That is the missing piece for a bounded, self-correcting development loop.

The architectural question is *who owns what*. DESIGN.md's "What the loop does NOT do" already states the loop does not run review and does not manage ADRs — those are skill responsibilities. Extending the loop into review/fix risks blurring that boundary: if the loop starts interpreting review content, it stops being the simple, agent-agnostic engine and becomes an LLM-as-judge.

Two roles must stay separate:

- **Orchestration** (the loop) — when to invoke review, when to invoke fix, when to stop. Mechanical and signal-driven.
- **Judgment** (the `review` and `fix` skills) — what counts as a finding, whether it is actionable, how to phrase a fix unit. Adversarial and semantic.

Alternatives considered:

- **LLM-as-judge inside the loop.** Rejected — the loop would interpret review content, breaking agent-agnosticism and the verify-gate principle (the aiming problem: the signal must be the actual property, not a proxy the loop can game).
- **The loop implements review logic.** Rejected — duplicates the `review` skill, drifts from the source of truth, and violates "the loop never reads skills."
- **Manual review/fix only (do nothing).** Rejected — leaves the loop unable to self-correct; the missing piece stays missing.

## Decision

The loop **orchestrates** an optional, bounded `build → review → fix` subloop. It invokes the `review` and `fix` workers and interprets their *structured* outputs as continue/stop signals. It does not implement review or fix logic.

When `--review` is set and the build queue has drained, the loop runs a review worker that writes a structured `REVIEW.md`. The loop reads only the `actionable` count from that file. If it is non-zero, the loop runs a fix worker that appends `Status: pending` units to `QUEUE.md`, then re-runs the build pass and reviews again. It stops when `actionable` is zero, a review-round limit (`--max-review-rounds`, default 2) is hit, the tick budget (`--max-ticks`) is exhausted, or a round generates no new units (no progress).

The boundary, stated as a ruling:

- **The loop owns orchestration and stop conditions** — invoking review/fix, reading the actionable count, enforcing hard stops. Mechanical.
- **The `review` and `fix` skills own judgment** — the two-axis standards/intent review, triage, finding phrasing, and work-unit generation. Semantic.

The loop never parses review content beyond the actionable count and never judges whether a finding is real. Review remains a skill the worker loads; the loop only knows to invoke it and read its summary.

## Consequences

- The loop gains a bounded self-correction capability without becoming an LLM-as-judge. Backpressure stays mechanical: the actionable count is the signal, the hard stops are the guardrails.
- A new contract surface: the loop depends on `REVIEW.md`'s actionable count being honest. A review worker that mislabels findings can loop forever or stop early — the round cap and the no-progress stop are the backstop, not the loop's own judgment.
- The `review` and `fix` skills gain a machine-readable input/output contract (`REVIEW.md`; appended units). Those skill changes land in their own work units; this ADR records only the architectural ruling and the orchestration boundary.
- Review stays opt-in (`--review`); default loop behavior is unchanged.

 succeeded in 0ms:
# Skills guide

Skills are agent-agnostic procedural knowledge stored in `.agents/skills/<name>/SKILL.md`. Any agent that supports agentskills.io can discover them automatically. The loop names the skill explicitly — name and path — in the worker prompt (`prompts/worker.md`); the agent reads the skill file directly. See ADR-0007.

## Default skills

| Skill | Purpose |
|---|---|
| `explore` | Investigate a codebase, grill intent, and stress-test ideas before planning. |
| `plan` | Convert intent into a disposable `QUEUE.md` of verifiable work units. |
| `build` | Implement one work unit from `QUEUE.md`; do not self-certify. |
| `review` | Run two-axis adversarial review (standards + intent) and generate findings. |
| `fix` | Convert review findings into new work units. |
| `decide` | Capture architectural rulings as ADRs in `decisions/`. |
| `domain-modeling` | Define and update `glossary.md` terms. |

## Skill format

A skill is a Markdown file named `SKILL.md` inside a directory named after the skill:

```text
.agents/skills/
└── build/
    └── SKILL.md
```

Required frontmatter:

```yaml
---
name: build
description: Use when implementing one work unit...
---
```

The `name` must match the directory name. The `description` is the trigger text used by agents to decide when to invoke the skill.

## How the loop uses skills

`loop.sh` does not read skills. It prepends `prompts/worker.md` to the current work unit and runs the worker. `prompts/worker.md` tells the worker to load the `build` skill by name and path (e.g. "Load and follow the **build** skill in `.agents/skills/build/`"). For other phases (explore, plan, review, fix), the human or the agent invokes the skill directly.

## Customizing skills

- After `knack skills init`, the project owns the `.agents/skills/` directory.
- Edit, override, or delete skills as needed.
- The CLI embeds the default skills. If you edit the defaults in the `knack` repo, run `cli/sync-skills.sh` to copy them into `cli/embedded/skills` and `diff -r .agents/skills cli/embedded/skills` to verify sync.
- Use `knack skills check` to validate your local skills.

## Composable flows

Skills are not a rigid gate. The default flow is `explore → plan → build → review → fix`, but any valid subset is fine:

```text
small fix → plan → build → done
bug report → explore → plan → build → done
big feature → explore → plan → build → review → fix → done
```

Decisions are captured inline throughout the flow using the `decide` skill, and terms are updated using `domain-modeling`.

codex
The existing docs still describe review as manual only. I’m checking `loop.sh` for the current CLI/env names and `REVIEW.md` contract so the references match implementation, not just the ADR.
exec
/usr/bin/bash -lc "sed -n '1,280p' loop.sh" in /home/daniel/build/knack
 succeeded in 0ms:
#!/usr/bin/env bash
set -euo pipefail

usage() {
  cat <<'EOF'
Usage:
  ./loop.sh run <queue> [--repo DIR] [--max-ticks N] [--review] [--max-review-rounds N] [--dry-run]

The queue is usually .loop/<name>/QUEUE.md in the target repo. Use --repo when the
queue lives outside the repository it should operate on.
EOF
}

die() {
  echo "knack: $*" >&2
  exit 1
}

abs_path() {
  python -c 'import os, sys; print(os.path.abspath(sys.argv[1]))' "$1"
}

first_pending_unit() {
  awk '
    BEGIN { in_block = 0; block = ""; status = ""; found = 0 }
    /^## / {
      if (in_block && status == "pending") { printf "%s", block; found = 1; exit }
      in_block = 1; block = $0 "\n"; status = ""; next
    }
    in_block {
      block = block $0 "\n"
      if ($0 ~ /^Status:[[:space:]]*pending[[:space:]]*$/) status = "pending"
    }
    END { if (!found && in_block && status == "pending") printf "%s", block }
  ' "$1"
}

extract_verify() {
  awk '
    BEGIN { after_verify = 0; in_fence = 0 }
    after_verify && /^```/ { if (!in_fence) { in_fence = 1; next } else exit }
    after_verify && in_fence { print; next }
    /^Verify:[[:space:]]*$/ { after_verify = 1 }
  ' "$1"
}

extract_agent() {
  awk '/^Agent:[[:space:]]+/ { sub(/^Agent:[[:space:]]*/, ""); gsub(/[[:space:]]+$/, ""); print; exit }' "$1"
}

set_status() {
  python - "$1" "$2" "$3" <<'PY'
import sys
from pathlib import Path

queue, title, status = sys.argv[1:]
path = Path(queue)
lines = path.read_text().splitlines(keepends=True)
heading = f"## {title}"
inside = False
changed = False
out = []
for line in lines:
    stripped = line.strip()
    if stripped.startswith("## ") and not stripped.startswith("### "):
        if inside and not changed:
            out.append(f"Status: {status}\n")
            changed = True
        inside = stripped == heading
    if inside and stripped.startswith("Status:") and not changed:
        out.append(f"Status: {status}\n")
        changed = True
        continue
    out.append(line)
if inside and not changed:
    if out and not out[-1].endswith("\n"):
        out[-1] += "\n"
    out.append(f"Status: {status}\n")
    changed = True
if not changed:
    raise SystemExit(f"unit not found or status already changed: {title}")
path.write_text("".join(out))
PY
}

write_handoff() {
  python - "$1" "$2" "$3" <<'PY'
import sys, re
from pathlib import Path
from datetime import datetime

queue, evidence, handoff = sys.argv[1:]
lines = Path(queue).read_text().splitlines()

units = []
current = None
for line in lines:
    if re.match(r'^## ', line) and not re.match(r'^###', line):
        if current:
            units.append(current)
        current = {"title": line[3:].strip(), "status": "pending"}
    elif current:
        m = re.match(r'^Status:\s*(\S+)', line)
        if m:
            current["status"] = m.group(1)
if current:
    units.append(current)

pending = [u for u in units if u["status"] != "done"]
if not pending:
    sys.exit(0)

completed = [u for u in units if u["status"] == "done"]
in_progress = [u for u in units if u["status"] in ("in_progress", "verify_failed", "no_progress", "blocked")]
remaining = [u for u in units if u["status"] == "pending"]

out = [
    f"# Handoff: {Path(queue).stem}",
    f"Generated: {datetime.now().isoformat()}",
    "",
    "## Completed",
]
out += [f"- {u['title']}" for u in completed] or ["- (none)"]
out += ["", "## In progress"]
out += [f"- {u['title']} (status: {u['status']})" for u in in_progress] or ["- (none)"]
out += ["", "## Remaining"]
out += [f"- {u['title']}" for u in remaining] or ["- (none)"]
out += ["", "## Next action"]
if in_progress:
    out.append(f"Re-run loop after addressing the {in_progress[0]['status']} state of: {in_progress[0]['title']}.")
elif remaining:
    out.append(f"Re-run loop to continue with: {remaining[0]['title']}.")
else:
    out.append("Queue is complete.")

Path(handoff).write_text("\n".join(out) + "\n")
PY
}

work_snapshot() {
  local repo_dir=$1
  if git -C "$repo_dir" rev-parse --is-inside-work-tree >/dev/null 2>&1; then
    {
      git -C "$repo_dir" diff -- . ':(exclude).loop' || true
      git -C "$repo_dir" status --short --untracked-files=all | awk '$2 !~ /^\.loop\// { print }'
    } | sha256sum | awk '{print $1}'
  else
    echo "__no_git__"
  fi
}

changed_files() {
  local repo_dir=$1
  if git -C "$repo_dir" rev-parse --is-inside-work-tree >/dev/null 2>&1; then
    git -C "$repo_dir" status --short --untracked-files=all || true
  else
    echo "git status unavailable: not inside a git repo"
  fi
}

append_evidence() {
  local evidence=$1 title=$2 status=$3 verify=$4 verify_out=$5 agent_out=$6 repo_dir=$7 unit_file=$8
  mkdir -p "$(dirname "$evidence")"
  {
    echo
    echo "## $(date -Iseconds) — $title"
    echo
    echo "Status: $status"
    echo
    echo "Unit:"
    echo '````markdown'
    cat "$unit_file"
    echo '````'
    echo
    echo "Files changed:"
    echo '```text'
    changed_files "$repo_dir"
    echo '```'
    echo
    echo "Verify command:"
    echo '```bash'
    echo "$verify"
    echo '```'
    echo
    echo "Verify output:"
    echo '```text'
    cat "$verify_out"
    echo '```'
    echo
    echo "Worker output:"
    echo '````text'
    cat "$agent_out"
    echo '````'
    echo
    echo "What this proves:"
    if [[ "$status" == "done" ]]; then
      echo "- The verify command passed for this work unit in the current repo state."
    else
      echo "- The work unit is not externally verified."
    fi
    echo
    echo "What remains unverified:"
    echo "- Anything outside the verify command's proof scope."
  } >> "$evidence"
}

extract_actionable_count() {
  awk '
    /^[[:space:]]*-[[:space:]]*actionable:[[:space:]]*[0-9]+[[:space:]]*$/ {
      sub(/^[[:space:]]*-[[:space:]]*actionable:[[:space:]]*/, "")
      sub(/[[:space:]]*$/, "")
      print
      found = 1
      exit
    }
    END { if (!found) exit 1 }
  ' "$1"
}

write_review_prompt() {
  local out=$1
  local template="$script_dir/prompts/reviewer.md"
  if [[ -f "$template" ]]; then
    cat "$template" > "$out"
  else
    cat > "$out" <<'EOF'
# Knack Reviewer

Load and follow the **review** skill in `.agents/skills/review/`.
Review the completed queue against the current repository state. Write the structured review artifact at the path provided below.
EOF
  fi
  cat >> "$out" <<EOF

Queue: $queue_abs
Evidence: $evidence
Review output: $review_file
EOF
}

write_fix_prompt() {
  local out=$1
  local template="$script_dir/prompts/fixer.md"
  if [[ -f "$template" ]]; then
    cat "$template" > "$out"
  else
    cat > "$out" <<'EOF'
# Knack Fixer

Load and follow the **fix** skill in `.agents/skills/fix/`.
Read the structured review artifact and append any actionable fix work units to the existing queue. Stop after updating the queue.
EOF
  fi
  cat >> "$out" <<EOF

Queue: $queue_abs
Evidence: $evidence
Review input: $review_file
EOF
}

run_phase_agent() {
  local phase=$1 prompt=$2 output=$3 cmd=${4:-}
  if [[ -n "$cmd" ]]; then
    (
      cd "$repo_dir"
      LOOP_PHASE="$phase" \
        LOOP_PROMPT_FILE="$prompt" \
        LOOP_QUEUE_FILE="$queue_abs" \
        LOOP_EVIDENCE_FILE="$evidence" \
        LOOP_REVIEW_FILE="$review_file" \
        bash -lc "$cmd"
    ) > "$output" 2>&1
  else
    (
      cd "$repo_dir"
      LOOP_PHASE="$phase" \
        LOOP_PROMPT_FILE="$prompt" \
        LOOP_QUEUE_FILE="$queue_abs" \
        LOOP_EVIDENCE_FILE="$evidence" \

exec
/usr/bin/bash -lc "sed -n '280,620p' loop.sh" in /home/daniel/build/knack
 succeeded in 0ms:
        LOOP_EVIDENCE_FILE="$evidence" \
        LOOP_REVIEW_FILE="$review_file" \
        pi -p --no-session --approve "$(cat "$prompt")"
    ) > "$output" 2>&1
  fi
}

[[ $# -ge 1 ]] || { usage; exit 1; }
cmd=$1
shift
[[ "$cmd" == "run" ]] || { usage; exit 1; }
[[ $# -ge 1 ]] || die "missing queue path"

queue_abs=$(abs_path "$1")
shift
[[ -f "$queue_abs" ]] || die "queue not found: $queue_abs"

max_ticks=3
max_review_rounds=2
review_enabled=0
dry_run=0
repo_override=""
while [[ $# -gt 0 ]]; do
  case "$1" in
    --repo)
      [[ $# -ge 2 ]] || die "--repo needs a value"
      repo_override=$(abs_path "$2")
      shift 2
      ;;
    --max-ticks)
      [[ $# -ge 2 ]] || die "--max-ticks needs a value"
      max_ticks=$2
      shift 2
      ;;
    --review)
      review_enabled=1
      shift
      ;;
    --max-review-rounds)
      [[ $# -ge 2 ]] || die "--max-review-rounds needs a value"
      max_review_rounds=$2
      shift 2
      ;;
    --dry-run)
      dry_run=1
      shift
      ;;
    *) die "unknown argument: $1" ;;
  esac
done

script_dir=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
prompt_file="$script_dir/prompts/worker.md"
[[ -f "$prompt_file" ]] || die "worker prompt not found: $prompt_file"

queue_dir=$(dirname "$queue_abs")
if [[ -n "$repo_override" ]]; then
  [[ -d "$repo_override" ]] || die "repo not found: $repo_override"
  repo_dir="$repo_override"
elif [[ $(basename "$queue_dir") == ".loop" ]]; then
  repo_dir=$(dirname "$queue_dir")
else
  repo_dir=$(pwd)
fi
evidence="$queue_dir/EVIDENCE.md"
handoff="$queue_dir/HANDOFF.md"
review_file="$queue_dir/REVIEW.md"
no_progress_strikes=0

write_handoff_on_exit() {
  local rc=$?
  [[ "${dry_run:-0}" == 0 ]] || return $rc
  [[ -f "${queue_abs:-}" ]] || return $rc
  write_handoff "$queue_abs" "$evidence" "$handoff" 2>/dev/null || true
  return $rc
}
trap write_handoff_on_exit EXIT

tick=1
review_round=0

while true; do
while (( tick <= max_ticks )); do
  unit=$(first_pending_unit "$queue_abs")
  [[ -n "$unit" ]] || break

  unit_file=$(mktemp)
  verify_file=$(mktemp)
  agent_out=$(mktemp)
  verify_out=$(mktemp)
  printf '%s' "$unit" > "$unit_file"

  first_line=$(awk 'NR == 1 { print; exit }' "$unit_file")
  title=${first_line#\#\# }
  verify=$(extract_verify "$unit_file")
  [[ -n "$verify" ]] || die "work unit has no Verify fenced block: $title"
  printf '%s\n' "$verify" > "$verify_file"

  if [[ "$dry_run" == 1 ]]; then
    echo "Unit: $title"
    echo "Repo: $repo_dir"
    echo "Verify:"
    cat "$verify_file"
    exit 0
  fi

  echo "knack: tick $tick/$max_ticks — $title"
  set_status "$queue_abs" "$title" "in_progress"
  before=$(work_snapshot "$repo_dir")

  run_prompt=$(mktemp)
  cat > "$run_prompt" <<EOF
$(cat "$prompt_file")

Current work unit from $queue_abs:

$(cat "$unit_file")
EOF

  agent_cmd="${LOOP_AGENT_CMD:-}"
  unit_agent=$(extract_agent "$unit_file")
  if [[ -n "$unit_agent" ]]; then
    agent_cmd="$unit_agent"
  fi

  set +e
  run_phase_agent "build" "$run_prompt" "$agent_out" "$agent_cmd"
  agent_code=$?
  set -e

  if [[ $agent_code -ne 0 ]]; then
    set_status "$queue_abs" "$title" "blocked"
    : > "$verify_out"
    append_evidence "$evidence" "$title" "blocked" "$verify" "$verify_out" "$agent_out" "$repo_dir" "$unit_file"
    cat "$agent_out"
    die "worker exited nonzero for $title"
  fi

  after=$(work_snapshot "$repo_dir")

  set +e
  (cd "$repo_dir" && bash -lc "$verify") > "$verify_out" 2>&1
  verify_code=$?
  set -e

  if [[ $verify_code -eq 0 ]]; then
    set_status "$queue_abs" "$title" "done"
    append_evidence "$evidence" "$title" "done" "$verify" "$verify_out" "$agent_out" "$repo_dir" "$unit_file"
    echo "knack: verified — $title"
    tick=$((tick + 1))
    continue
  fi

  if [[ "$before" == "$after" ]]; then
    no_progress_strikes=$((no_progress_strikes + 1))
    if [[ $no_progress_strikes -ge 2 ]]; then
      set_status "$queue_abs" "$title" "no_progress"
      append_evidence "$evidence" "$title" "no_progress" "$verify" "$verify_out" "$agent_out" "$repo_dir" "$unit_file"
      die "no progress after $no_progress_strikes attempts on $title"
    fi
    set_status "$queue_abs" "$title" "pending"
    append_evidence "$evidence" "$title" "verify_failed" "$verify" "$verify_out" "$agent_out" "$repo_dir" "$unit_file"
    echo "knack: verify failed with no progress; retrying once"
    tick=$((tick + 1))
    continue
  fi

  set_status "$queue_abs" "$title" "verify_failed"
  append_evidence "$evidence" "$title" "verify_failed" "$verify" "$verify_out" "$agent_out" "$repo_dir" "$unit_file"
  cat "$verify_out"
  die "verify failed for $title"
done

if [[ -n "$(first_pending_unit "$queue_abs")" ]]; then
  die "reached max ticks ($max_ticks) with pending work"
fi

if [[ "$review_enabled" == 0 ]]; then
  if (( tick > max_ticks )); then
    echo "knack: reached max ticks ($max_ticks)"
    exit 0
  fi
  echo "knack: no pending work units"
  exit 0
fi

if (( review_round >= max_review_rounds )); then
  die "reached max review rounds ($max_review_rounds)"
fi

review_round=$((review_round + 1))
review_prompt=$(mktemp)
review_out=$(mktemp)
write_review_prompt "$review_prompt"
review_cmd="${LOOP_REVIEW_CMD:-${LOOP_AGENT_CMD:-}}"

echo "knack: review round $review_round/$max_review_rounds"
set +e
run_phase_agent "review" "$review_prompt" "$review_out" "$review_cmd"
review_code=$?
set -e
if [[ $review_code -ne 0 ]]; then
  cat "$review_out"
  die "review worker exited nonzero"
fi
[[ -f "$review_file" ]] || die "review worker did not write $review_file"
actionable=$(extract_actionable_count "$review_file") || die "review file has no actionable summary: $review_file"

if [[ "$actionable" == 0 ]]; then
  echo "knack: review clean"
  exit 0
fi

if (( review_round >= max_review_rounds )); then
  die "review found $actionable actionable issue(s) at max review rounds ($max_review_rounds)"
fi

fix_prompt=$(mktemp)
fix_out=$(mktemp)
write_fix_prompt "$fix_prompt"
fix_cmd="${LOOP_FIX_CMD:-${LOOP_AGENT_CMD:-}}"

echo "knack: fix round $review_round/$max_review_rounds — actionable: $actionable"
set +e
run_phase_agent "fix" "$fix_prompt" "$fix_out" "$fix_cmd"
fix_code=$?
set -e
if [[ $fix_code -ne 0 ]]; then
  cat "$fix_out"
  die "fix worker exited nonzero"
fi

if [[ -z "$(first_pending_unit "$queue_abs")" ]]; then
  cat "$fix_out"
  die "fix worker generated no pending work units"
fi
done

exec
/usr/bin/bash -lc "rg -n \"review|max-review|LOOP_REVIEW|LOOP_FIX|REVIEW.md|fix\" loop.sh tests docs README.md AGENTS.md .agents/skills -g '"'!cli/embedded/**'"'" in /home/daniel/build/knack
 succeeded in 0ms:
.agents/skills/fix/SKILL.md:2:name: fix
.agents/skills/fix/SKILL.md:3:description: Use when addressing review findings and feeding them back into the loop. Converts review findings into new verifiable work units in `.loop/<name>/QUEUE.md` and runs another loop pass. Triggers on "fix the review findings", "address the feedback", "rework based on review", "fix what review found", or when review produced findings that need to be resolved.
.agents/skills/fix/SKILL.md:8:Address review findings by triaging `REVIEW.md` and generating new work units that feed back into the loop. The `fix` skill is the bridge between review and another build pass; it owns the triage and the queue format. It appends to the existing `.loop/<name>/QUEUE.md` file for the same cycle.
.agents/skills/fix/SKILL.md:10:The loop owns orchestration. When `fix` is invoked by the loop, stop after appending units and reporting the triage summary; the loop will run the next build pass.
.agents/skills/fix/SKILL.md:14:- `.loop/<name>/REVIEW.md` — structured findings from the `review` skill. Read it, but do not edit it.
.agents/skills/fix/SKILL.md:17:`REVIEW.md` is expected to include `## Standards`, `## Intent`, `## Speculative`, and `## Summary` sections. The summary's actionable count is the loop's signal; the finding entries are the fix skill's input. Use the finding classification, confidence, and evidence when deciding whether to generate a work unit.
.agents/skills/fix/SKILL.md:21:1. **Read `REVIEW.md`.** Each standards or intent finding is a candidate work unit. `review` has already classified findings as trivial / actionable / disputed / deferred; use those classifications or re-triage as needed. Treat speculative findings as notes for future explore/plan work unless the review explicitly classifies one as actionable.
.agents/skills/fix/SKILL.md:24:   - **Actionable** — the finding identifies a real issue with a clear fix. Create a work unit.
.agents/skills/fix/SKILL.md:25:   - **Trivial** — one-line fix, no risk. Note it in the summary, don't create a unit.
.agents/skills/fix/SKILL.md:26:   - **Disputed** — the finding is wrong or the reviewer is being overly cautious. Note the disagreement in the summary, don't create a unit. Move on.
.agents/skills/fix/SKILL.md:27:   - **Deferred** — the finding is valid but not worth fixing now. Note it in the summary or a backlog, don't create a unit.
.agents/skills/fix/SKILL.md:29:3. **Create work units** for actionable findings. Use the standard work unit format from the `plan` skill. Each unit must have one observable outcome and one deterministic verify command. Reference the source finding in `Read first:` by pointing to `REVIEW.md` plus any evidence paths from the finding.
.agents/skills/fix/SKILL.md:32:## <fix for the finding — observable outcome>
.agents/skills/fix/SKILL.md:35:- .loop/<name>/REVIEW.md (<finding id or heading>)
.agents/skills/fix/SKILL.md:49:<deterministic command that proves the fix>
.agents/skills/fix/SKILL.md:59:## What fix is not
.agents/skills/fix/SKILL.md:61:- Not a re-review — the findings are already known. Fix turns them into work.
.agents/skills/fix/SKILL.md:62:- Not a debate — if a finding is disputed, note it and move on. Don't argue with the review in the queue.
.agents/skills/fix/SKILL.md:63:- Not a rewrite — each fix unit should be narrow. If a finding requires a broad rewrite, that's a new explore → plan cycle, not a fix.
.agents/skills/fix/SKILL.md:64:- Not the orchestrator — the loop invokes review, invokes fix, and runs the next build pass.
.agents/skills/fix/SKILL.md:68:Fix units must not break what review already approved. Each fix unit's `Done means:` should include a no-regression condition. The verify command should cover both the fix and the preservation of what was already working.
.agents/skills/review/SKILL.md:2:name: review
.agents/skills/review/SKILL.md:3:description: Use when reviewing completed work units or a finished queue. Two-axis adversarial review — standards (does the change follow codebase conventions?) and intent (does the change do what the work unit said?). Reviews against the actual codebase, not specs. Triggers on "review", "check this", "is this right", "what did we miss", "stress-test the implementation", or when work needs adversarial scrutiny before being accepted.
.agents/skills/review/SKILL.md:8:Adversarial review of completed work. Two axes, run independently so neither pollutes the other:
.agents/skills/review/SKILL.md:15:## When to review
.agents/skills/review/SKILL.md:24:## Before you review
.agents/skills/review/SKILL.md:26:Read the work unit from `.loop/<name>/QUEUE.md` and the evidence from `.loop/<name>/EVIDENCE.md` for the unit you're reviewing. The evidence tells you what the verify command actually proved; the work unit tells you what was promised. Review against the actual codebase, not the specs.
.agents/skills/review/SKILL.md:28:## Two-axis review
.agents/skills/review/SKILL.md:53:The `Verify:` command is the mechanically enforceable subset of `Done means:`. The gap between `Done means:` and `Verify:` is the review surface: intent review checks what the verify command cannot.
.agents/skills/review/SKILL.md:55:> **Intent review is a judgment check, not a deterministic gate.** Be explicit about its scope and confidence. Any finding that becomes a new work unit must have a deterministic `Verify` command the runner can execute.
.agents/skills/review/SKILL.md:61:## Findings become input to the fix skill
.agents/skills/review/SKILL.md:63:Review findings are not just notes — they are the input to the `fix` skill, which triages them and appends actionable ones as new work units in `.loop/<name>/QUEUE.md`.
.agents/skills/review/SKILL.md:65:- **Trivial** findings (a typo, a missing newline) can be fixed inline during review.
.agents/skills/review/SKILL.md:66:- **Actionable** findings are handed to the `fix` skill. Do not write the work units yourself; `fix` owns the triage and formatting.
.agents/skills/review/SKILL.md:67:- **Disputed** or **deferred** findings are recorded in the review summary but not turned into units.
.agents/skills/review/SKILL.md:69:The output of review is a findings summary, not a queue edit.
.agents/skills/review/SKILL.md:71:## What review is not
.agents/skills/review/SKILL.md:75:- Not a gate — the loop doesn't enforce review. It's invoked when the user wants it.
.agents/skills/review/SKILL.md:87:- Hand actionable findings to the `fix` skill, which will triage and append work units to `.loop/<name>/QUEUE.md`, or
.agents/skills/explore/SKILL.md:54:- Small fixes where the problem is already clear
.agents/skills/domain-modeling/SKILL.md:3:description: Use when domain terms surface during explore, plan, or review and need to be defined, challenged, or stress-tested against the project's ubiquitous language. Manages `glossary.md` — the project's shared vocabulary. Triggers on "what does X mean here", "are these the same concept", "let's define our terms", "that term is ambiguous", or when a domain concept is used inconsistently across a conversation or codebase.
.agents/skills/domain-modeling/SKILL.md:10:This is a shared skill — called by explore, plan, and review. It is not a separate phase.
.agents/skills/decide/SKILL.md:3:description: Use when an architectural ruling crystallizes during any phase — explore, plan, build, or review. Captures decisions as ADRs in `decisions/` so they persist after specs are deleted and code moves on. Triggers on "we decided to", "let's go with", "the ruling is", "why did we choose", "record this decision", or when a design choice with lasting consequences is made.
.agents/skills/decide/SKILL.md:10:This is a shared skill — called by explore, plan, build, and review whenever a ruling crystallizes. It is not a separate phase.
.agents/skills/decide/SKILL.md:52:1. Recognize that a ruling just crystallized. If you're in explore, plan, build, or review and you just resolved a tension with lasting consequences, that's a decision.
AGENTS.md:22:The loop (`loop.sh`), all seven skills, and the Go CLI are built. The loop supports per-unit `Agent:` overrides, `LOOP_AGENT_CMD` for agent-agnosticism, and writes `.loop/<name>/HANDOFF.md` on non-clean exits. The skills cover the full explore → plan → build → review → fix flow plus two shared skills (decide, domain-modeling). The CLI (`cli/`) provides `skills init|check`, `validate`, `decisions list|show|check`, `status`, `glossary check`, and `instructions`. It embeds the default skills via `go:embed` and scaffolds them into target projects.
AGENTS.md:31:- `.agents/skills/` — procedural knowledge (explore, plan, build, review, fix, decide, domain-modeling).
AGENTS.md:39:- **review** — two-axis adversarial review (standards + intent). Findings become new work units.
AGENTS.md:40:- **fix** — address review findings. Generates new work units, feeds back into the loop.
AGENTS.md:41:- **decide** (shared) — capture architectural rulings as ADRs in `decisions/`. Used by explore, plan, build, review.
AGENTS.md:42:- **domain-modeling** (shared) — manage `glossary.md`. Used by explore, plan, review.
AGENTS.md:51:- Review is opt-in. The loop does not run review automatically; the user invokes it when they want adversarial scrutiny.
AGENTS.md:57:- Work units are `## <outcome>` headers with `Read first:`, `Constraints:`, `Done means:`, and `Verify:` fields. `Done means:` is the acceptance criteria; `Verify:` is the mechanically enforceable subset. The gap between them is the review surface.
AGENTS.md:60:- Work units are `## <outcome>` headers — no "Slice" prefix, no numbering. Vertical slice is one type of work unit, not the required format.
AGENTS.md:82:- **Review catches what verify can't.** The queue parser regex `^##\s*(.*)$` matched `###` subheadings as work units — a real bug that diverged from `loop.sh`'s behavior. `go test` passed because no fixture used `###`. Adversarial review against the actual codebase (comparing to `loop.sh`'s parser) found it. The fix: exclude `###` lines explicitly in `isUnitHeader`.
AGENTS.md:86:- **Workers scope to the outcome plus constraints, not to a file list.** ADR-0005 replaced `Work:` with `Read first:` and `Constraints:`. The unit's scope is its outcome plus its constraints — the worker determines which files to change. The old lesson ("name every file in the work unit") is wrong under the new shape: naming files in constraints smuggles scope the same way `Work:` did. The first plan-shape cycle proved this — the constraint said "no `Work:` refs in skills, prompts, DESIGN.md, or AGENTS.md" and the worker touched exactly those files, leaving 9 other files (test fixtures, examples, README) with stale `Work:` fields. Prefer outcome-level constraints ("no artifact that teaches the format may reference `Work:`") over file-enumerated constraints.
README.md:49:Work units are `## <outcome>` headers with `Read first:`, `Constraints:`, `Done means:`, and `Verify:` fields. The `Verify:` command is the mechanically enforceable subset of `Done means:`; the gap between them is the review surface.
.agents/skills/plan/SKILL.md:17:- **Directly** — for small work where explore isn't needed. A bug fix, a patch, a small feature. Skip the ceremony.
.agents/skills/plan/SKILL.md:30:- **patch** — small, localized fix. One change, one verify.
.agents/skills/plan/SKILL.md:32:- **bug fix** — reproduce → fix → verify. The verify command must fail before the fix and pass after.
.agents/skills/plan/SKILL.md:35:"Vertical slice" is the preferred default, not a required format. The planner prefers slices and rejects horizontal phases, but a unit can be a patch, investigation, or bug fix when the work genuinely isn't sliceable.
.agents/skills/plan/SKILL.md:82:- Good: `validate --json and text mode report the same broken-link error on the same fixture`
.agents/skills/plan/SKILL.md:85:- Good: `the regression fixture fails before the fix and passes after the fix`
.agents/skills/plan/SKILL.md:98:Skip this entirely for small work. A bug fix doesn't need a proposal.
.agents/skills/plan/SKILL.md:146:- **Header** is `## <outcome>` — no numbered prefix, no "Slice" word. The outcome itself is the title.
loop.sh:7:  ./loop.sh run <queue> [--repo DIR] [--max-ticks N] [--review] [--max-review-rounds N] [--dry-run]
loop.sh:220:write_review_prompt() {
loop.sh:222:  local template="$script_dir/prompts/reviewer.md"
loop.sh:229:Load and follow the **review** skill in `.agents/skills/review/`.
loop.sh:230:Review the completed queue against the current repository state. Write the structured review artifact at the path provided below.
loop.sh:237:Review output: $review_file
loop.sh:241:write_fix_prompt() {
loop.sh:243:  local template="$script_dir/prompts/fixer.md"
loop.sh:250:Load and follow the **fix** skill in `.agents/skills/fix/`.
loop.sh:251:Read the structured review artifact and append any actionable fix work units to the existing queue. Stop after updating the queue.
loop.sh:258:Review input: $review_file
loop.sh:271:        LOOP_REVIEW_FILE="$review_file" \
loop.sh:281:        LOOP_REVIEW_FILE="$review_file" \
loop.sh:298:max_review_rounds=2
loop.sh:299:review_enabled=0
loop.sh:314:    --review)
loop.sh:315:      review_enabled=1
loop.sh:318:    --max-review-rounds)
loop.sh:319:      [[ $# -ge 2 ]] || die "--max-review-rounds needs a value"
loop.sh:320:      max_review_rounds=$2
loop.sh:346:review_file="$queue_dir/REVIEW.md"
loop.sh:359:review_round=0
loop.sh:457:if [[ "$review_enabled" == 0 ]]; then
loop.sh:466:if (( review_round >= max_review_rounds )); then
loop.sh:467:  die "reached max review rounds ($max_review_rounds)"
loop.sh:470:review_round=$((review_round + 1))
loop.sh:471:review_prompt=$(mktemp)
loop.sh:472:review_out=$(mktemp)
loop.sh:473:write_review_prompt "$review_prompt"
loop.sh:474:review_cmd="${LOOP_REVIEW_CMD:-${LOOP_AGENT_CMD:-}}"
loop.sh:476:echo "knack: review round $review_round/$max_review_rounds"
loop.sh:478:run_phase_agent "review" "$review_prompt" "$review_out" "$review_cmd"
loop.sh:479:review_code=$?
loop.sh:481:if [[ $review_code -ne 0 ]]; then
loop.sh:482:  cat "$review_out"
loop.sh:483:  die "review worker exited nonzero"
loop.sh:485:[[ -f "$review_file" ]] || die "review worker did not write $review_file"
loop.sh:486:actionable=$(extract_actionable_count "$review_file") || die "review file has no actionable summary: $review_file"
loop.sh:489:  echo "knack: review clean"
loop.sh:493:if (( review_round >= max_review_rounds )); then
loop.sh:494:  die "review found $actionable actionable issue(s) at max review rounds ($max_review_rounds)"
loop.sh:497:fix_prompt=$(mktemp)
loop.sh:498:fix_out=$(mktemp)
loop.sh:499:write_fix_prompt "$fix_prompt"
loop.sh:500:fix_cmd="${LOOP_FIX_CMD:-${LOOP_AGENT_CMD:-}}"
loop.sh:502:echo "knack: fix round $review_round/$max_review_rounds — actionable: $actionable"
loop.sh:504:run_phase_agent "fix" "$fix_prompt" "$fix_out" "$fix_cmd"
loop.sh:505:fix_code=$?
loop.sh:507:if [[ $fix_code -ne 0 ]]; then
loop.sh:508:  cat "$fix_out"
loop.sh:509:  die "fix worker exited nonzero"
loop.sh:513:  cat "$fix_out"
loop.sh:514:  die "fix worker generated no pending work units"
docs/loop.md:77:- `loop.sh` does **not** run review or manage ADRs/glossary. Those are handled by the skills.
tests/run.sh:30:## the test fixture reaches its verify condition
tests/run.sh:90:assert_contains "$repo2/.loop/HANDOFF.md" "the test fixture reaches its verify condition"
tests/run.sh:160:assert_contains "$repo_pi/pi-args.txt" "the test fixture reaches its verify condition"
tests/run.sh:170:assert_contains "$repo_lpf/captured-prompt.txt" "the test fixture reaches its verify condition"
tests/run.sh:172:# Review-fix loop with fake build, review, and fix workers.
tests/run.sh:173:repo_review="$tmp/repo-review"
tests/run.sh:174:mkdir -p "$repo_review/.loop"
tests/run.sh:175:cat > "$repo_review/.loop/QUEUE.md" <<'EOF'
tests/run.sh:176:# Loop Queue: review cycle
tests/run.sh:179:Exercise build, review, fix, and review again.
tests/run.sh:182:The generated app is fixed and review reports no actionable issues.
tests/run.sh:184:## the initial build creates a reviewable app file
tests/run.sh:190:- Leave the bug for review to find.
tests/run.sh:203:cat > "$repo_review/build-worker.sh" <<'EOF'
tests/run.sh:206:if grep -q "the fix unit repairs the bug" "$LOOP_PROMPT_FILE"; then
tests/run.sh:207:  printf 'fixed\n' > app.txt
tests/run.sh:208:  echo "build fixed app"
tests/run.sh:214:chmod +x "$repo_review/build-worker.sh"
tests/run.sh:216:cat > "$repo_review/review-worker.sh" <<'EOF'
tests/run.sh:220:if [[ -f review-count.txt ]]; then
tests/run.sh:221:  count=$(cat review-count.txt)
tests/run.sh:224:printf '%s\n' "$count" > review-count.txt
tests/run.sh:227:if [[ -f app.txt ]] && grep -qx 'fixed' app.txt; then
tests/run.sh:231:cat > "$LOOP_REVIEW_FILE" <<EOF_REVIEW
tests/run.sh:237:- actionable | high — app.txt must say fixed
tests/run.sh:248:echo "review actionable: $actionable"
tests/run.sh:250:chmod +x "$repo_review/review-worker.sh"
tests/run.sh:252:cat > "$repo_review/fix-worker.sh" <<'EOF'
tests/run.sh:257:## the fix unit repairs the bug
tests/run.sh:260:- .loop/REVIEW.md
tests/run.sh:267:grep -qx fixed app.txt
tests/run.sh:271:- app.txt contains fixed.
tests/run.sh:275:echo "fix appended unit"
tests/run.sh:277:chmod +x "$repo_review/fix-worker.sh"
tests/run.sh:279:LOOP_AGENT_CMD="$repo_review/build-worker.sh" \
tests/run.sh:280:  LOOP_REVIEW_CMD="$repo_review/review-worker.sh" \
tests/run.sh:281:  LOOP_FIX_CMD="$repo_review/fix-worker.sh" \
tests/run.sh:282:  "$root/loop.sh" run "$repo_review/.loop/QUEUE.md" --review --max-ticks 2 >/tmp/loop-review.txt
tests/run.sh:283:assert_contains "$repo_review/.loop/QUEUE.md" "## the fix unit repairs the bug"
tests/run.sh:284:assert_contains "$repo_review/.loop/QUEUE.md" "Status: done"
tests/run.sh:285:assert_contains "$repo_review/.loop/REVIEW.md" "- actionable: 0"
tests/run.sh:286:assert_contains "$repo_review/review-count.txt" "2"
tests/run.sh:287:assert_contains "$repo_review/app.txt" "fixed"
docs/queue-format.md:31:The heading is the outcome. There is no `Slice` prefix and no numbering. Avoid `###` headings inside a work unit — they may confuse simple parsers.
docs/queue-format.md:50:- `Verify:` is the mechanically enforceable subset of `Done means:`. The gap between them is the review surface.
docs/queue-format.md:56:# Loop Queue: parser fix
docs/faq.md:11:knack replaces litespec. It keeps the flow (explore → plan → build → review → fix) and the idea of skills, but drops durable specs as source of truth. Code is the source of truth; work units and handoff files are disposable; decisions and glossary are durable.
docs/faq.md:55:`Done means:` is the acceptance criteria — what must be true when the unit is finished. `Verify:` is the mechanically enforceable subset that the runner can actually execute. The gap between them is the review surface.
docs/faq.md:59:Yes. "Vertical slice" is the preferred default, but work units can be patches, bug fixes, investigations, or refactors. The only hard requirement is a deterministic `Verify:` command.
docs/faq.md:97:### Can I skip explore or review?
docs/skills.md:12:| `review` | Run two-axis adversarial review (standards + intent) and generate findings. |
docs/skills.md:13:| `fix` | Convert review findings into new work units. |
docs/skills.md:40:`loop.sh` does not read skills. It prepends `prompts/worker.md` to the current work unit and runs the worker. `prompts/worker.md` tells the worker to load the `build` skill by name and path (e.g. "Load and follow the **build** skill in `.agents/skills/build/`"). For other phases (explore, plan, review, fix), the human or the agent invokes the skill directly.
docs/skills.md:51:Skills are not a rigid gate. The default flow is `explore → plan → build → review → fix`, but any valid subset is fine:
docs/skills.md:54:small fix → plan → build → done
docs/skills.md:56:big feature → explore → plan → build → review → fix → done

codex
I’ve got the implementation contract: `--review` adds bounded review rounds, `LOOP_REVIEW_CMD` and `LOOP_FIX_CMD` override phase agents, and `.loop/<name>/REVIEW.md` carries the actionable count. I’m editing only the docs and AGENTS.md now; no skill files need syncing for this unit.
apply patch
patch: completed
/home/daniel/build/knack/docs/loop.md
diff --git a/docs/loop.md b/docs/loop.md
index a5cd2cd28883321ca006489e2c630fbdc8d6bb8f..ef47822f4007c3fbcb5ee094f39164166ca63646
--- a/docs/loop.md
+++ b/docs/loop.md
@@ -5,18 +5,24 @@
 ## Usage
 
 ```bash
-./loop.sh run <queue> [--repo DIR] [--max-ticks N] [--dry-run]
+./loop.sh run <queue> [--repo DIR] [--max-ticks N] [--review] [--max-review-rounds N] [--dry-run]
 ```
 
 - `queue` — path to `QUEUE.md`. Usually `.loop/<name>/QUEUE.md`.
 - `--repo DIR` — the working directory for the worker and the verify command. Defaults to the parent of the `.loop/<name>` directory that contains the queue.
 - `--max-ticks N` — maximum units to attempt. Default is `3`.
+- `--review` — opt into the bounded review/fix subloop after pending build units drain. Default behavior does not run review.
+- `--max-review-rounds N` — maximum review/fix rounds when `--review` is enabled. Default is `2`.
 - `--dry-run` — parse the first pending unit and print its title, repo, and verify command, then exit.
 
 ## Environment variables
 
 - `LOOP_AGENT_CMD` — optional command used to invoke the worker. If unset, the loop defaults to `pi -p --no-session --approve` with the prompt text as a single argument. If set, the command is evaluated by `bash -lc` in the repo directory, and `LOOP_PROMPT_FILE` is set to a temporary file containing `prompts/worker.md` plus the current work unit.
 - `LOOP_PROMPT_FILE` — set by the loop when `LOOP_AGENT_CMD` is used. Points to the generated prompt file. Do not override unless you are calling the worker manually.
+- `LOOP_REVIEW_CMD` — optional command used for the review phase when `--review` is enabled. Defaults to `LOOP_AGENT_CMD`, or to the loop's default `pi` invocation if neither is set.
+- `LOOP_FIX_CMD` — optional command used for the fix phase when `--review` is enabled. Defaults to `LOOP_AGENT_CMD`, or to the loop's default `pi` invocation if neither is set.
+
+During every phase, the loop also sets `LOOP_PHASE`, `LOOP_QUEUE_FILE`, `LOOP_EVIDENCE_FILE`, and `LOOP_REVIEW_FILE` for the worker process.
 
 ## Per-unit agent override
 
@@ -43,6 +49,20 @@
    - If the repo changed, mark `verify_failed` and stop.
 9. If max ticks is reached with pending work, stop.
 
+## Optional review/fix subloop
+
+By default, `loop.sh` stops when the build queue has no pending work. With `--review`, it then runs a bounded review/fix subloop:
+
+1. Invoke a review worker with the review prompt and the completed queue/evidence.
+2. Require the review worker to write `.loop/<name>/REVIEW.md`.
+3. Read only the `- actionable: N` summary line from `REVIEW.md`.
+4. If `N` is `0`, stop cleanly.
+5. If `N` is non-zero, invoke a fix worker.
+6. Require the fix worker to append new `Status: pending` work units to the same `QUEUE.md`.
+7. Run the build pass again, then review again.
+
+The subloop stops when review is clean, `--max-review-rounds` is reached, `--max-ticks` is exhausted, or fix produces no pending work. The loop owns orchestration and stop conditions only; the `review` and `fix` skills own judgment and work-unit generation.
+
 ## Work unit statuses
 
 - `pending` — not yet started.
@@ -58,6 +78,7 @@
 
 - `EVIDENCE.md` — append-only ledger. Includes the full unit, changed files, verify command, verify output, and worker output. It is durable; keep it after deleting `QUEUE.md` so `knack decisions check` can still see which ADRs the cycle referenced.
 - `HANDOFF.md` — written on non-clean exit. Sections: completed, in progress, remaining, next action. Delete when the work resumes.
+- `REVIEW.md` — structured review artifact written by the review worker when `--review` is enabled. The loop reads only its `- actionable: N` summary line.
 
 ## Agent invocation examples
 
@@ -71,8 +92,17 @@
 
 The command is passed to `bash -lc` in the repo directory, so the typical pattern is `"$(cat "$LOOP_PROMPT_FILE")"`.
 
+Review and fix can use separate agents:
+
+```bash
+LOOP_AGENT_CMD='codex exec --dangerously-bypass-approvals-and-sandbox --ephemeral "$(cat "$LOOP_PROMPT_FILE")"' \
+LOOP_REVIEW_CMD='claude --print --no-session-persistence --dangerously-skip-permissions "$(cat "$LOOP_PROMPT_FILE")"' \
+LOOP_FIX_CMD='codex exec --dangerously-bypass-approvals-and-sandbox --ephemeral "$(cat "$LOOP_PROMPT_FILE")"' \
+./loop.sh run .loop/<name>/QUEUE.md --review
+```
+
 ## Verification notes
 
 - `loop.sh` does **not** validate `QUEUE.md` structure. Use `knack validate` first.
-- `loop.sh` does **not** run review or manage ADRs/glossary. Those are handled by the skills.
+- `loop.sh` runs review and fix only when `--review` is set. It invokes the skills and reads the actionable count from `REVIEW.md`; it does not judge findings or manage ADRs/glossary.
 - The `Verify:` command must be deterministic and executable by the runner — tests, builds, type checks, not an LLM-as-judge.

diff --git a/docs/loop.md b/docs/loop.md
index a5cd2cd28883321ca006489e2c630fbdc8d6bb8f..ef47822f4007c3fbcb5ee094f39164166ca63646
--- a/docs/loop.md
+++ b/docs/loop.md
@@ -5,18 +5,24 @@
 ## Usage
 
 ```bash
-./loop.sh run <queue> [--repo DIR] [--max-ticks N] [--dry-run]
+./loop.sh run <queue> [--repo DIR] [--max-ticks N] [--review] [--max-review-rounds N] [--dry-run]
 ```
 
 - `queue` — path to `QUEUE.md`. Usually `.loop/<name>/QUEUE.md`.
 - `--repo DIR` — the working directory for the worker and the verify command. Defaults to the parent of the `.loop/<name>` directory that contains the queue.
 - `--max-ticks N` — maximum units to attempt. Default is `3`.
+- `--review` — opt into the bounded review/fix subloop after pending build units drain. Default behavior does not run review.
+- `--max-review-rounds N` — maximum review/fix rounds when `--review` is enabled. Default is `2`.
 - `--dry-run` — parse the first pending unit and print its title, repo, and verify command, then exit.
 
 ## Environment variables
 
 - `LOOP_AGENT_CMD` — optional command used to invoke the worker. If unset, the loop defaults to `pi -p --no-session --approve` with the prompt text as a single argument. If set, the command is evaluated by `bash -lc` in the repo directory, and `LOOP_PROMPT_FILE` is set to a temporary file containing `prompts/worker.md` plus the current work unit.
 - `LOOP_PROMPT_FILE` — set by the loop when `LOOP_AGENT_CMD` is used. Points to the generated prompt file. Do not override unless you are calling the worker manually.
+- `LOOP_REVIEW_CMD` — optional command used for the review phase when `--review` is enabled. Defaults to `LOOP_AGENT_CMD`, or to the loop's default `pi` invocation if neither is set.
+- `LOOP_FIX_CMD` — optional command used for the fix phase when `--review` is enabled. Defaults to `LOOP_AGENT_CMD`, or to the loop's default `pi` invocation if neither is set.
+
+During every phase, the loop also sets `LOOP_PHASE`, `LOOP_QUEUE_FILE`, `LOOP_EVIDENCE_FILE`, and `LOOP_REVIEW_FILE` for the worker process.
 
 ## Per-unit agent override
 
@@ -43,6 +49,20 @@
    - If the repo changed, mark `verify_failed` and stop.
 9. If max ticks is reached with pending work, stop.
 
+## Optional review/fix subloop
+
+By default, `loop.sh` stops when the build queue has no pending work. With `--review`, it then runs a bounded review/fix subloop:
+
+1. Invoke a review worker with the review prompt and the completed queue/evidence.
+2. Require the review worker to write `.loop/<name>/REVIEW.md`.
+3. Read only the `- actionable: N` summary line from `REVIEW.md`.
+4. If `N` is `0`, stop cleanly.
+5. If `N` is non-zero, invoke a fix worker.
+6. Require the fix worker to append new `Status: pending` work units to the same `QUEUE.md`.
+7. Run the build pass again, then review again.
+
+The subloop stops when review is clean, `--max-review-rounds` is reached, `--max-ticks` is exhausted, or fix produces no pending work. The loop owns orchestration and stop conditions only; the `review` and `fix` skills own judgment and work-unit generation.
+
 ## Work unit statuses
 
 - `pending` — not yet started.
@@ -58,6 +78,7 @@
 
 - `EVIDENCE.md` — append-only ledger. Includes the full unit, changed files, verify command, verify output, and worker output. It is durable; keep it after deleting `QUEUE.md` so `knack decisions check` can still see which ADRs the cycle referenced.
 - `HANDOFF.md` — written on non-clean exit. Sections: completed, in progress, remaining, next action. Delete when the work resumes.
+- `REVIEW.md` — structured review artifact written by the review worker when `--review` is enabled. The loop reads only its `- actionable: N` summary line.
 
 ## Agent invocation examples
 
@@ -71,8 +92,17 @@
 
 The command is passed to `bash -lc` in the repo directory, so the typical pattern is `"$(cat "$LOOP_PROMPT_FILE")"`.
 
+Review and fix can use separate agents:
+
+```bash
+LOOP_AGENT_CMD='codex exec --dangerously-bypass-approvals-and-sandbox --ephemeral "$(cat "$LOOP_PROMPT_FILE")"' \
+LOOP_REVIEW_CMD='claude --print --no-session-persistence --dangerously-skip-permissions "$(cat "$LOOP_PROMPT_FILE")"' \
+LOOP_FIX_CMD='codex exec --dangerously-bypass-approvals-and-sandbox --ephemeral "$(cat "$LOOP_PROMPT_FILE")"' \
+./loop.sh run .loop/<name>/QUEUE.md --review
+```
+
 ## Verification notes
 
 - `loop.sh` does **not** validate `QUEUE.md` structure. Use `knack validate` first.
-- `loop.sh` does **not** run review or manage ADRs/glossary. Those are handled by the skills.
+- `loop.sh` runs review and fix only when `--review` is set. It invokes the skills and reads the actionable count from `REVIEW.md`; it does not judge findings or manage ADRs/glossary.
 - The `Verify:` command must be deterministic and executable by the runner — tests, builds, type checks, not an LLM-as-judge.

apply patch
patch: completed
/home/daniel/build/knack/docs/skills.md
diff --git a/docs/loop.md b/docs/loop.md
index a5cd2cd28883321ca006489e2c630fbdc8d6bb8f..ef47822f4007c3fbcb5ee094f39164166ca63646
--- a/docs/loop.md
+++ b/docs/loop.md
@@ -5,18 +5,24 @@
 ## Usage
 
 ```bash
-./loop.sh run <queue> [--repo DIR] [--max-ticks N] [--dry-run]
+./loop.sh run <queue> [--repo DIR] [--max-ticks N] [--review] [--max-review-rounds N] [--dry-run]
 ```
 
 - `queue` — path to `QUEUE.md`. Usually `.loop/<name>/QUEUE.md`.
 - `--repo DIR` — the working directory for the worker and the verify command. Defaults to the parent of the `.loop/<name>` directory that contains the queue.
 - `--max-ticks N` — maximum units to attempt. Default is `3`.
+- `--review` — opt into the bounded review/fix subloop after pending build units drain. Default behavior does not run review.
+- `--max-review-rounds N` — maximum review/fix rounds when `--review` is enabled. Default is `2`.
 - `--dry-run` — parse the first pending unit and print its title, repo, and verify command, then exit.
 
 ## Environment variables
 
 - `LOOP_AGENT_CMD` — optional command used to invoke the worker. If unset, the loop defaults to `pi -p --no-session --approve` with the prompt text as a single argument. If set, the command is evaluated by `bash -lc` in the repo directory, and `LOOP_PROMPT_FILE` is set to a temporary file containing `prompts/worker.md` plus the current work unit.
 - `LOOP_PROMPT_FILE` — set by the loop when `LOOP_AGENT_CMD` is used. Points to the generated prompt file. Do not override unless you are calling the worker manually.
+- `LOOP_REVIEW_CMD` — optional command used for the review phase when `--review` is enabled. Defaults to `LOOP_AGENT_CMD`, or to the loop's default `pi` invocation if neither is set.
+- `LOOP_FIX_CMD` — optional command used for the fix phase when `--review` is enabled. Defaults to `LOOP_AGENT_CMD`, or to the loop's default `pi` invocation if neither is set.
+
+During every phase, the loop also sets `LOOP_PHASE`, `LOOP_QUEUE_FILE`, `LOOP_EVIDENCE_FILE`, and `LOOP_REVIEW_FILE` for the worker process.
 
 ## Per-unit agent override
 
@@ -43,6 +49,20 @@
    - If the repo changed, mark `verify_failed` and stop.
 9. If max ticks is reached with pending work, stop.
 
+## Optional review/fix subloop
+
+By default, `loop.sh` stops when the build queue has no pending work. With `--review`, it then runs a bounded review/fix subloop:
+
+1. Invoke a review worker with the review prompt and the completed queue/evidence.
+2. Require the review worker to write `.loop/<name>/REVIEW.md`.
+3. Read only the `- actionable: N` summary line from `REVIEW.md`.
+4. If `N` is `0`, stop cleanly.
+5. If `N` is non-zero, invoke a fix worker.
+6. Require the fix worker to append new `Status: pending` work units to the same `QUEUE.md`.
+7. Run the build pass again, then review again.
+
+The subloop stops when review is clean, `--max-review-rounds` is reached, `--max-ticks` is exhausted, or fix produces no pending work. The loop owns orchestration and stop conditions only; the `review` and `fix` skills own judgment and work-unit generation.
+
 ## Work unit statuses
 
 - `pending` — not yet started.
@@ -58,6 +78,7 @@
 
 - `EVIDENCE.md` — append-only ledger. Includes the full unit, changed files, verify command, verify output, and worker output. It is durable; keep it after deleting `QUEUE.md` so `knack decisions check` can still see which ADRs the cycle referenced.
 - `HANDOFF.md` — written on non-clean exit. Sections: completed, in progress, remaining, next action. Delete when the work resumes.
+- `REVIEW.md` — structured review artifact written by the review worker when `--review` is enabled. The loop reads only its `- actionable: N` summary line.
 
 ## Agent invocation examples
 
@@ -71,8 +92,17 @@
 
 The command is passed to `bash -lc` in the repo directory, so the typical pattern is `"$(cat "$LOOP_PROMPT_FILE")"`.
 
+Review and fix can use separate agents:
+
+```bash
+LOOP_AGENT_CMD='codex exec --dangerously-bypass-approvals-and-sandbox --ephemeral "$(cat "$LOOP_PROMPT_FILE")"' \
+LOOP_REVIEW_CMD='claude --print --no-session-persistence --dangerously-skip-permissions "$(cat "$LOOP_PROMPT_FILE")"' \
+LOOP_FIX_CMD='codex exec --dangerously-bypass-approvals-and-sandbox --ephemeral "$(cat "$LOOP_PROMPT_FILE")"' \
+./loop.sh run .loop/<name>/QUEUE.md --review
+```
+
 ## Verification notes
 
 - `loop.sh` does **not** validate `QUEUE.md` structure. Use `knack validate` first.
-- `loop.sh` does **not** run review or manage ADRs/glossary. Those are handled by the skills.
+- `loop.sh` runs review and fix only when `--review` is set. It invokes the skills and reads the actionable count from `REVIEW.md`; it does not judge findings or manage ADRs/glossary.
 - The `Verify:` command must be deterministic and executable by the runner — tests, builds, type checks, not an LLM-as-judge.
diff --git a/docs/skills.md b/docs/skills.md
index fb1ec98119e1193859c0a1612f580d2da5afe95d..476d494e3c88962825e751b84151a0ff48ac2d41
--- a/docs/skills.md
+++ b/docs/skills.md
@@ -37,7 +37,11 @@
 
 ## How the loop uses skills
 
-`loop.sh` does not read skills. It prepends `prompts/worker.md` to the current work unit and runs the worker. `prompts/worker.md` tells the worker to load the `build` skill by name and path (e.g. "Load and follow the **build** skill in `.agents/skills/build/`"). For other phases (explore, plan, review, fix), the human or the agent invokes the skill directly.
+`loop.sh` does not read skills. It prepends `prompts/worker.md` to the current work unit and runs the worker. `prompts/worker.md` tells the worker to load the `build` skill by name and path (e.g. "Load and follow the **build** skill in `.agents/skills/build/`").
+
+When `--review` is set, the loop also invokes review and fix workers after the build queue drains. Those prompts tell the worker to load the `review` or `fix` skill directly. The loop orchestrates the bounded review/fix subloop, reads the actionable count from `REVIEW.md`, and runs another build pass when fix appends pending units. The skills still own judgment: review decides what the findings are, and fix decides which findings become work units.
+
+Without `--review`, review and fix remain manual skill invocations.
 
 ## Customizing skills
 
@@ -53,7 +57,7 @@
 ```text
 small fix → plan → build → done
 bug report → explore → plan → build → done
-big feature → explore → plan → build → review → fix → done
+big feature → explore → plan → build --review → review → fix → build → done
 ```
 
 Decisions are captured inline throughout the flow using the `decide` skill, and terms are updated using `domain-modeling`.

diff --git a/docs/loop.md b/docs/loop.md
index a5cd2cd28883321ca006489e2c630fbdc8d6bb8f..ef47822f4007c3fbcb5ee094f39164166ca63646
--- a/docs/loop.md
+++ b/docs/loop.md
@@ -5,18 +5,24 @@
 ## Usage
 
 ```bash
-./loop.sh run <queue> [--repo DIR] [--max-ticks N] [--dry-run]
+./loop.sh run <queue> [--repo DIR] [--max-ticks N] [--review] [--max-review-rounds N] [--dry-run]
 ```
 
 - `queue` — path to `QUEUE.md`. Usually `.loop/<name>/QUEUE.md`.
 - `--repo DIR` — the working directory for the worker and the verify command. Defaults to the parent of the `.loop/<name>` directory that contains the queue.
 - `--max-ticks N` — maximum units to attempt. Default is `3`.
+- `--review` — opt into the bounded review/fix subloop after pending build units drain. Default behavior does not run review.
+- `--max-review-rounds N` — maximum review/fix rounds when `--review` is enabled. Default is `2`.
 - `--dry-run` — parse the first pending unit and print its title, repo, and verify command, then exit.
 
 ## Environment variables
 
 - `LOOP_AGENT_CMD` — optional command used to invoke the worker. If unset, the loop defaults to `pi -p --no-session --approve` with the prompt text as a single argument. If set, the command is evaluated by `bash -lc` in the repo directory, and `LOOP_PROMPT_FILE` is set to a temporary file containing `prompts/worker.md` plus the current work unit.
 - `LOOP_PROMPT_FILE` — set by the loop when `LOOP_AGENT_CMD` is used. Points to the generated prompt file. Do not override unless you are calling the worker manually.
+- `LOOP_REVIEW_CMD` — optional command used for the review phase when `--review` is enabled. Defaults to `LOOP_AGENT_CMD`, or to the loop's default `pi` invocation if neither is set.
+- `LOOP_FIX_CMD` — optional command used for the fix phase when `--review` is enabled. Defaults to `LOOP_AGENT_CMD`, or to the loop's default `pi` invocation if neither is set.
+
+During every phase, the loop also sets `LOOP_PHASE`, `LOOP_QUEUE_FILE`, `LOOP_EVIDENCE_FILE`, and `LOOP_REVIEW_FILE` for the worker process.
 
 ## Per-unit agent override
 
@@ -43,6 +49,20 @@
    - If the repo changed, mark `verify_failed` and stop.
 9. If max ticks is reached with pending work, stop.
 
+## Optional review/fix subloop
+
+By default, `loop.sh` stops when the build queue has no pending work. With `--review`, it then runs a bounded review/fix subloop:
+
+1. Invoke a review worker with the review prompt and the completed queue/evidence.
+2. Require the review worker to write `.loop/<name>/REVIEW.md`.
+3. Read only the `- actionable: N` summary line from `REVIEW.md`.
+4. If `N` is `0`, stop cleanly.
+5. If `N` is non-zero, invoke a fix worker.
+6. Require the fix worker to append new `Status: pending` work units to the same `QUEUE.md`.
+7. Run the build pass again, then review again.
+
+The subloop stops when review is clean, `--max-review-rounds` is reached, `--max-ticks` is exhausted, or fix produces no pending work. The loop owns orchestration and stop conditions only; the `review` and `fix` skills own judgment and work-unit generation.
+
 ## Work unit statuses
 
 - `pending` — not yet started.
@@ -58,6 +78,7 @@
 
 - `EVIDENCE.md` — append-only ledger. Includes the full unit, changed files, verify command, verify output, and worker output. It is durable; keep it after deleting `QUEUE.md` so `knack decisions check` can still see which ADRs the cycle referenced.
 - `HANDOFF.md` — written on non-clean exit. Sections: completed, in progress, remaining, next action. Delete when the work resumes.
+- `REVIEW.md` — structured review artifact written by the review worker when `--review` is enabled. The loop reads only its `- actionable: N` summary line.
 
 ## Agent invocation examples
 
@@ -71,8 +92,17 @@
 
 The command is passed to `bash -lc` in the repo directory, so the typical pattern is `"$(cat "$LOOP_PROMPT_FILE")"`.
 
+Review and fix can use separate agents:
+
+```bash
+LOOP_AGENT_CMD='codex exec --dangerously-bypass-approvals-and-sandbox --ephemeral "$(cat "$LOOP_PROMPT_FILE")"' \
+LOOP_REVIEW_CMD='claude --print --no-session-persistence --dangerously-skip-permissions "$(cat "$LOOP_PROMPT_FILE")"' \
+LOOP_FIX_CMD='codex exec --dangerously-bypass-approvals-and-sandbox --ephemeral "$(cat "$LOOP_PROMPT_FILE")"' \
+./loop.sh run .loop/<name>/QUEUE.md --review
+```
+
 ## Verification notes
 
 - `loop.sh` does **not** validate `QUEUE.md` structure. Use `knack validate` first.
-- `loop.sh` does **not** run review or manage ADRs/glossary. Those are handled by the skills.
+- `loop.sh` runs review and fix only when `--review` is set. It invokes the skills and reads the actionable count from `REVIEW.md`; it does not judge findings or manage ADRs/glossary.
 - The `Verify:` command must be deterministic and executable by the runner — tests, builds, type checks, not an LLM-as-judge.
diff --git a/docs/skills.md b/docs/skills.md
index fb1ec98119e1193859c0a1612f580d2da5afe95d..476d494e3c88962825e751b84151a0ff48ac2d41
--- a/docs/skills.md
+++ b/docs/skills.md
@@ -37,7 +37,11 @@
 
 ## How the loop uses skills
 
-`loop.sh` does not read skills. It prepends `prompts/worker.md` to the current work unit and runs the worker. `prompts/worker.md` tells the worker to load the `build` skill by name and path (e.g. "Load and follow the **build** skill in `.agents/skills/build/`"). For other phases (explore, plan, review, fix), the human or the agent invokes the skill directly.
+`loop.sh` does not read skills. It prepends `prompts/worker.md` to the current work unit and runs the worker. `prompts/worker.md` tells the worker to load the `build` skill by name and path (e.g. "Load and follow the **build** skill in `.agents/skills/build/`").
+
+When `--review` is set, the loop also invokes review and fix workers after the build queue drains. Those prompts tell the worker to load the `review` or `fix` skill directly. The loop orchestrates the bounded review/fix subloop, reads the actionable count from `REVIEW.md`, and runs another build pass when fix appends pending units. The skills still own judgment: review decides what the findings are, and fix decides which findings become work units.
+
+Without `--review`, review and fix remain manual skill invocations.
 
 ## Customizing skills
 
@@ -53,7 +57,7 @@
 ```text
 small fix → plan → build → done
 bug report → explore → plan → build → done
-big feature → explore → plan → build → review → fix → done
+big feature → explore → plan → build --review → review → fix → build → done
 ```
 
 Decisions are captured inline throughout the flow using the `decide` skill, and terms are updated using `domain-modeling`.

apply patch
patch: completed
/home/daniel/build/knack/README.md
diff --git a/README.md b/README.md
index 1300a1bc12102f97cd89ace07333f179f8700cee..c1541df3127dce0ba5ba94a9cc4d8ba97c5a6b6a
--- a/README.md
+++ b/README.md
@@ -43,6 +43,9 @@
 6. On failure: the unit is marked `verify_failed` or `no_progress`, evidence is appended, and the loop stops or retries once.
 7. The loop halts on: max ticks reached with pending work, two no-progress strikes, or a failed verify after a real change.
 8. On any non-clean exit, the runner writes `HANDOFF.md` with completed/in-progress/remaining units and the next action.
+9. If `--review` is set, the runner invokes review after the queue drains, reads the actionable count from `REVIEW.md`, invokes fix when actionable findings exist, and runs another build pass for appended units.
+
+Review is opt-in. Without `--review`, the loop runs build ticks only; with it, review/fix are loop-orchestrated but still skill-owned.
 
 ## Queue format
 
@@ -70,6 +73,15 @@
 Agent: pi -p --no-session --approve --model glm-5.2 "$(cat "$LOOP_PROMPT_FILE")"
 ```
 
+Opt into review/fix orchestration with `--review`:
+
+```bash
+LOOP_AGENT_CMD='codex exec --dangerously-bypass-approvals-and-sandbox --ephemeral "$(cat "$LOOP_PROMPT_FILE")"' \
+LOOP_REVIEW_CMD='claude --print --no-session-persistence --dangerously-skip-permissions "$(cat "$LOOP_PROMPT_FILE")"' \
+LOOP_FIX_CMD='codex exec --dangerously-bypass-approvals-and-sandbox --ephemeral "$(cat "$LOOP_PROMPT_FILE")"' \
+./loop.sh run .loop/<name>/QUEUE.md --review --max-review-rounds 2
+```
+
 ## CLI
 
 The `knack` binary is a read-only validator and context provider. Build it from `cli/`:
diff --git a/docs/loop.md b/docs/loop.md
index a5cd2cd28883321ca006489e2c630fbdc8d6bb8f..ef47822f4007c3fbcb5ee094f39164166ca63646
--- a/docs/loop.md
+++ b/docs/loop.md
@@ -5,18 +5,24 @@
 ## Usage
 
 ```bash
-./loop.sh run <queue> [--repo DIR] [--max-ticks N] [--dry-run]
+./loop.sh run <queue> [--repo DIR] [--max-ticks N] [--review] [--max-review-rounds N] [--dry-run]
 ```
 
 - `queue` — path to `QUEUE.md`. Usually `.loop/<name>/QUEUE.md`.
 - `--repo DIR` — the working directory for the worker and the verify command. Defaults to the parent of the `.loop/<name>` directory that contains the queue.
 - `--max-ticks N` — maximum units to attempt. Default is `3`.
+- `--review` — opt into the bounded review/fix subloop after pending build units drain. Default behavior does not run review.
+- `--max-review-rounds N` — maximum review/fix rounds when `--review` is enabled. Default is `2`.
 - `--dry-run` — parse the first pending unit and print its title, repo, and verify command, then exit.
 
 ## Environment variables
 
 - `LOOP_AGENT_CMD` — optional command used to invoke the worker. If unset, the loop defaults to `pi -p --no-session --approve` with the prompt text as a single argument. If set, the command is evaluated by `bash -lc` in the repo directory, and `LOOP_PROMPT_FILE` is set to a temporary file containing `prompts/worker.md` plus the current work unit.
 - `LOOP_PROMPT_FILE` — set by the loop when `LOOP_AGENT_CMD` is used. Points to the generated prompt file. Do not override unless you are calling the worker manually.
+- `LOOP_REVIEW_CMD` — optional command used for the review phase when `--review` is enabled. Defaults to `LOOP_AGENT_CMD`, or to the loop's default `pi` invocation if neither is set.
+- `LOOP_FIX_CMD` — optional command used for the fix phase when `--review` is enabled. Defaults to `LOOP_AGENT_CMD`, or to the loop's default `pi` invocation if neither is set.
+
+During every phase, the loop also sets `LOOP_PHASE`, `LOOP_QUEUE_FILE`, `LOOP_EVIDENCE_FILE`, and `LOOP_REVIEW_FILE` for the worker process.
 
 ## Per-unit agent override
 
@@ -43,6 +49,20 @@
    - If the repo changed, mark `verify_failed` and stop.
 9. If max ticks is reached with pending work, stop.
 
+## Optional review/fix subloop
+
+By default, `loop.sh` stops when the build queue has no pending work. With `--review`, it then runs a bounded review/fix subloop:
+
+1. Invoke a review worker with the review prompt and the completed queue/evidence.
+2. Require the review worker to write `.loop/<name>/REVIEW.md`.
+3. Read only the `- actionable: N` summary line from `REVIEW.md`.
+4. If `N` is `0`, stop cleanly.
+5. If `N` is non-zero, invoke a fix worker.
+6. Require the fix worker to append new `Status: pending` work units to the same `QUEUE.md`.
+7. Run the build pass again, then review again.
+
+The subloop stops when review is clean, `--max-review-rounds` is reached, `--max-ticks` is exhausted, or fix produces no pending work. The loop owns orchestration and stop conditions only; the `review` and `fix` skills own judgment and work-unit generation.
+
 ## Work unit statuses
 
 - `pending` — not yet started.
@@ -58,6 +78,7 @@
 
 - `EVIDENCE.md` — append-only ledger. Includes the full unit, changed files, verify command, verify output, and worker output. It is durable; keep it after deleting `QUEUE.md` so `knack decisions check` can still see which ADRs the cycle referenced.
 - `HANDOFF.md` — written on non-clean exit. Sections: completed, in progress, remaining, next action. Delete when the work resumes.
+- `REVIEW.md` — structured review artifact written by the review worker when `--review` is enabled. The loop reads only its `- actionable: N` summary line.
 
 ## Agent invocation examples
 
@@ -71,8 +92,17 @@
 
 The command is passed to `bash -lc` in the repo directory, so the typical pattern is `"$(cat "$LOOP_PROMPT_FILE")"`.
 
+Review and fix can use separate agents:
+
+```bash
+LOOP_AGENT_CMD='codex exec --dangerously-bypass-approvals-and-sandbox --ephemeral "$(cat "$LOOP_PROMPT_FILE")"' \
+LOOP_REVIEW_CMD='claude --print --no-session-persistence --dangerously-skip-permissions "$(cat "$LOOP_PROMPT_FILE")"' \
+LOOP_FIX_CMD='codex exec --dangerously-bypass-approvals-and-sandbox --ephemeral "$(cat "$LOOP_PROMPT_FILE")"' \
+./loop.sh run .loop/<name>/QUEUE.md --review
+```
+
 ## Verification notes
 
 - `loop.sh` does **not** validate `QUEUE.md` structure. Use `knack validate` first.
-- `loop.sh` does **not** run review or manage ADRs/glossary. Those are handled by the skills.
+- `loop.sh` runs review and fix only when `--review` is set. It invokes the skills and reads the actionable count from `REVIEW.md`; it does not judge findings or manage ADRs/glossary.
 - The `Verify:` command must be deterministic and executable by the runner — tests, builds, type checks, not an LLM-as-judge.
diff --git a/docs/skills.md b/docs/skills.md
index fb1ec98119e1193859c0a1612f580d2da5afe95d..476d494e3c88962825e751b84151a0ff48ac2d41
--- a/docs/skills.md
+++ b/docs/skills.md
@@ -37,7 +37,11 @@
 
 ## How the loop uses skills
 
-`loop.sh` does not read skills. It prepends `prompts/worker.md` to the current work unit and runs the worker. `prompts/worker.md` tells the worker to load the `build` skill by name and path (e.g. "Load and follow the **build** skill in `.agents/skills/build/`"). For other phases (explore, plan, review, fix), the human or the agent invokes the skill directly.
+`loop.sh` does not read skills. It prepends `prompts/worker.md` to the current work unit and runs the worker. `prompts/worker.md` tells the worker to load the `build` skill by name and path (e.g. "Load and follow the **build** skill in `.agents/skills/build/`").
+
+When `--review` is set, the loop also invokes review and fix workers after the build queue drains. Those prompts tell the worker to load the `review` or `fix` skill directly. The loop orchestrates the bounded review/fix subloop, reads the actionable count from `REVIEW.md`, and runs another build pass when fix appends pending units. The skills still own judgment: review decides what the findings are, and fix decides which findings become work units.
+
+Without `--review`, review and fix remain manual skill invocations.
 
 ## Customizing skills
 
@@ -53,7 +57,7 @@
 ```text
 small fix → plan → build → done
 bug report → explore → plan → build → done
-big feature → explore → plan → build → review → fix → done
+big feature → explore → plan → build --review → review → fix → build → done
 ```
 
 Decisions are captured inline throughout the flow using the `decide` skill, and terms are updated using `domain-modeling`.

diff --git a/README.md b/README.md
index 1300a1bc12102f97cd89ace07333f179f8700cee..c1541df3127dce0ba5ba94a9cc4d8ba97c5a6b6a
--- a/README.md
+++ b/README.md
@@ -43,6 +43,9 @@
 6. On failure: the unit is marked `verify_failed` or `no_progress`, evidence is appended, and the loop stops or retries once.
 7. The loop halts on: max ticks reached with pending work, two no-progress strikes, or a failed verify after a real change.
 8. On any non-clean exit, the runner writes `HANDOFF.md` with completed/in-progress/remaining units and the next action.
+9. If `--review` is set, the runner invokes review after the queue drains, reads the actionable count from `REVIEW.md`, invokes fix when actionable findings exist, and runs another build pass for appended units.
+
+Review is opt-in. Without `--review`, the loop runs build ticks only; with it, review/fix are loop-orchestrated but still skill-owned.
 
 ## Queue format
 
@@ -70,6 +73,15 @@
 Agent: pi -p --no-session --approve --model glm-5.2 "$(cat "$LOOP_PROMPT_FILE")"
 ```
 
+Opt into review/fix orchestration with `--review`:
+
+```bash
+LOOP_AGENT_CMD='codex exec --dangerously-bypass-approvals-and-sandbox --ephemeral "$(cat "$LOOP_PROMPT_FILE")"' \
+LOOP_REVIEW_CMD='claude --print --no-session-persistence --dangerously-skip-permissions "$(cat "$LOOP_PROMPT_FILE")"' \
+LOOP_FIX_CMD='codex exec --dangerously-bypass-approvals-and-sandbox --ephemeral "$(cat "$LOOP_PROMPT_FILE")"' \
+./loop.sh run .loop/<name>/QUEUE.md --review --max-review-rounds 2
+```
+
 ## CLI
 
 The `knack` binary is a read-only validator and context provider. Build it from `cli/`:
diff --git a/docs/loop.md b/docs/loop.md
index a5cd2cd28883321ca006489e2c630fbdc8d6bb8f..ef47822f4007c3fbcb5ee094f39164166ca63646
--- a/docs/loop.md
+++ b/docs/loop.md
@@ -5,18 +5,24 @@
 ## Usage
 
 ```bash
-./loop.sh run <queue> [--repo DIR] [--max-ticks N] [--dry-run]
+./loop.sh run <queue> [--repo DIR] [--max-ticks N] [--review] [--max-review-rounds N] [--dry-run]
 ```
 
 - `queue` — path to `QUEUE.md`. Usually `.loop/<name>/QUEUE.md`.
 - `--repo DIR` — the working directory for the worker and the verify command. Defaults to the parent of the `.loop/<name>` directory that contains the queue.
 - `--max-ticks N` — maximum units to attempt. Default is `3`.
+- `--review` — opt into the bounded review/fix subloop after pending build units drain. Default behavior does not run review.
+- `--max-review-rounds N` — maximum review/fix rounds when `--review` is enabled. Default is `2`.
 - `--dry-run` — parse the first pending unit and print its title, repo, and verify command, then exit.
 
 ## Environment variables
 
 - `LOOP_AGENT_CMD` — optional command used to invoke the worker. If unset, the loop defaults to `pi -p --no-session --approve` with the prompt text as a single argument. If set, the command is evaluated by `bash -lc` in the repo directory, and `LOOP_PROMPT_FILE` is set to a temporary file containing `prompts/worker.md` plus the current work unit.
 - `LOOP_PROMPT_FILE` — set by the loop when `LOOP_AGENT_CMD` is used. Points to the generated prompt file. Do not override unless you are calling the worker manually.
+- `LOOP_REVIEW_CMD` — optional command used for the review phase when `--review` is enabled. Defaults to `LOOP_AGENT_CMD`, or to the loop's default `pi` invocation if neither is set.
+- `LOOP_FIX_CMD` — optional command used for the fix phase when `--review` is enabled. Defaults to `LOOP_AGENT_CMD`, or to the loop's default `pi` invocation if neither is set.
+
+During every phase, the loop also sets `LOOP_PHASE`, `LOOP_QUEUE_FILE`, `LOOP_EVIDENCE_FILE`, and `LOOP_REVIEW_FILE` for the worker process.
 
 ## Per-unit agent override
 
@@ -43,6 +49,20 @@
    - If the repo changed, mark `verify_failed` and stop.
 9. If max ticks is reached with pending work, stop.
 
+## Optional review/fix subloop
+
+By default, `loop.sh` stops when the build queue has no pending work. With `--review`, it then runs a bounded review/fix subloop:
+
+1. Invoke a review worker with the review prompt and the completed queue/evidence.
+2. Require the review worker to write `.loop/<name>/REVIEW.md`.
+3. Read only the `- actionable: N` summary line from `REVIEW.md`.
+4. If `N` is `0`, stop cleanly.
+5. If `N` is non-zero, invoke a fix worker.
+6. Require the fix worker to append new `Status: pending` work units to the same `QUEUE.md`.
+7. Run the build pass again, then review again.
+
+The subloop stops when review is clean, `--max-review-rounds` is reached, `--max-ticks` is exhausted, or fix produces no pending work. The loop owns orchestration and stop conditions only; the `review` and `fix` skills own judgment and work-unit generation.
+
 ## Work unit statuses
 
 - `pending` — not yet started.
@@ -58,6 +78,7 @@
 
 - `EVIDENCE.md` — append-only ledger. Includes the full unit, changed files, verify command, verify output, and worker output. It is durable; keep it after deleting `QUEUE.md` so `knack decisions check` can still see which ADRs the cycle referenced.
 - `HANDOFF.md` — written on non-clean exit. Sections: completed, in progress, remaining, next action. Delete when the work resumes.
+- `REVIEW.md` — structured review artifact written by the review worker when `--review` is enabled. The loop reads only its `- actionable: N` summary line.
 
 ## Agent invocation examples
 
@@ -71,8 +92,17 @@
 
 The command is passed to `bash -lc` in the repo directory, so the typical pattern is `"$(cat "$LOOP_PROMPT_FILE")"`.
 
+Review and fix can use separate agents:
+
+```bash
+LOOP_AGENT_CMD='codex exec --dangerously-bypass-approvals-and-sandbox --ephemeral "$(cat "$LOOP_PROMPT_FILE")"' \
+LOOP_REVIEW_CMD='claude --print --no-session-persistence --dangerously-skip-permissions "$(cat "$LOOP_PROMPT_FILE")"' \
+LOOP_FIX_CMD='codex exec --dangerously-bypass-approvals-and-sandbox --ephemeral "$(cat "$LOOP_PROMPT_FILE")"' \
+./loop.sh run .loop/<name>/QUEUE.md --review
+```
+
 ## Verification notes
 
 - `loop.sh` does **not** validate `QUEUE.md` structure. Use `knack validate` first.
-- `loop.sh` does **not** run review or manage ADRs/glossary. Those are handled by the skills.
+- `loop.sh` runs review and fix only when `--review` is set. It invokes the skills and reads the actionable count from `REVIEW.md`; it does not judge findings or manage ADRs/glossary.
 - The `Verify:` command must be deterministic and executable by the runner — tests, builds, type checks, not an LLM-as-judge.
diff --git a/docs/skills.md b/docs/skills.md
index fb1ec98119e1193859c0a1612f580d2da5afe95d..476d494e3c88962825e751b84151a0ff48ac2d41
--- a/docs/skills.md
+++ b/docs/skills.md
@@ -37,7 +37,11 @@
 
 ## How the loop uses skills
 
-`loop.sh` does not read skills. It prepends `prompts/worker.md` to the current work unit and runs the worker. `prompts/worker.md` tells the worker to load the `build` skill by name and path (e.g. "Load and follow the **build** skill in `.agents/skills/build/`"). For other phases (explore, plan, review, fix), the human or the agent invokes the skill directly.
+`loop.sh` does not read skills. It prepends `prompts/worker.md` to the current work unit and runs the worker. `prompts/worker.md` tells the worker to load the `build` skill by name and path (e.g. "Load and follow the **build** skill in `.agents/skills/build/`").
+
+When `--review` is set, the loop also invokes review and fix workers after the build queue drains. Those prompts tell the worker to load the `review` or `fix` skill directly. The loop orchestrates the bounded review/fix subloop, reads the actionable count from `REVIEW.md`, and runs another build pass when fix appends pending units. The skills still own judgment: review decides what the findings are, and fix decides which findings become work units.
+
+Without `--review`, review and fix remain manual skill invocations.
 
 ## Customizing skills
 
@@ -53,7 +57,7 @@
 ```text
 small fix → plan → build → done
 bug report → explore → plan → build → done
-big feature → explore → plan → build → review → fix → done
+big feature → explore → plan → build --review → review → fix → build → done
 ```
 
 Decisions are captured inline throughout the flow using the `decide` skill, and terms are updated using `domain-modeling`.

codex
The user-facing docs now describe opt-in review/fix orchestration without changing the thesis. I’m updating AGENTS.md to keep the project’s operational context consistent with the new loop behavior.
apply patch
patch: completed
/home/daniel/build/knack/AGENTS.md
diff --git a/AGENTS.md b/AGENTS.md
index 50f568dd4d4d4e617b4be191116bb8fea556720d..019d79a3530e1a7621ecda0e6bfcdf1bd9e53ffa
--- a/AGENTS.md
+++ b/AGENTS.md
@@ -17,15 +17,16 @@
 
 See `DESIGN.md` for the full design.
 
-## Current state
+## Current State
 
-The loop (`loop.sh`), all seven skills, and the Go CLI are built. The loop supports per-unit `Agent:` overrides, `LOOP_AGENT_CMD` for agent-agnosticism, and writes `.loop/<name>/HANDOFF.md` on non-clean exits. The skills cover the full explore → plan → build → review → fix flow plus two shared skills (decide, domain-modeling). The CLI (`cli/`) provides `skills init|check`, `validate`, `decisions list|show|check`, `status`, `glossary check`, and `instructions`. It embeds the default skills via `go:embed` and scaffolds them into target projects.
+The loop (`loop.sh`), all seven skills, and the Go CLI are built. The loop supports per-unit `Agent:` overrides, `LOOP_AGENT_CMD` for agent-agnosticism, opt-in review/fix orchestration via `--review`, and `.loop/<name>/HANDOFF.md` on non-clean exits. The skills cover the full explore → plan → build → review → fix flow plus two shared skills (decide, domain-modeling). The CLI (`cli/`) provides `skills init|check`, `validate`, `decisions list|show|check`, `status`, `glossary check`, and `instructions`. It embeds the default skills via `go:embed` and scaffolds them into target projects.
 
 ## Core artifacts
 
 - `.loop/<name>/QUEUE.md` — disposable queue of verifiable work units.
 - `.loop/<name>/EVIDENCE.md` — append-only ledger of what each tick proved. Durable: keep it after deleting QUEUE.md so `decisions check` can still see which ADRs the completed cycle referenced.
 - `.loop/<name>/HANDOFF.md` — cross-session handoff (written on pause/stop).
+- `.loop/<name>/REVIEW.md` — structured review artifact written when `--review` is enabled. The loop reads only the actionable count.
 - `decisions/` — durable ADRs (architectural rulings, not current behavior).
 - `glossary.md` — durable ubiquitous language.
 - `.agents/skills/` — procedural knowledge (explore, plan, build, review, fix, decide, domain-modeling).
@@ -48,11 +49,13 @@
 - The worker never self-certifies. The runner owns verification.
 - A `Verify` command must be deterministic and executable by the runner — not an LLM-as-judge. The loop's backpressure is mechanical.
 - The runner enforces hard stops (max ticks, no-progress detection). The agent does one work unit per tick and reports what remains if it can't finish.
-- Review is opt-in. The loop does not run review automatically; the user invokes it when they want adversarial scrutiny.
+- Review is opt-in. The loop runs review/fix only when invoked with `--review`; otherwise it runs build ticks only.
+- With `--review`, the loop invokes review after pending build units drain, reads only `REVIEW.md`'s actionable count, invokes fix when actionable findings exist, then runs another build pass for appended units. Judgment stays in the `review` and `fix` skills.
 - When the queue is complete and verified, the cycle's `.loop/<name>/QUEUE.md` is disposable. The human deletes it; the loop does not. `EVIDENCE.md` is the durable ledger — keep it so `decisions check` can still trace which ADRs the cycle referenced.
 - A work unit must leave the repo better if the loop stops immediately after it.
 - Work units are whatever shape the work is — not forced into "vertical slices."
 - `LOOP_AGENT_CMD` overrides the worker invocation (agent-agnostic: pi, claude, codex, opencode, etc.).
+- `LOOP_REVIEW_CMD` and `LOOP_FIX_CMD` override review/fix phase invocations when `--review` is enabled. They default to `LOOP_AGENT_CMD`.
 - Per-unit `Agent:` field in QUEUE.md overrides `LOOP_AGENT_CMD` for one unit (model routing).
 - Work units are `## <outcome>` headers with `Read first:`, `Constraints:`, `Done means:`, and `Verify:` fields. `Done means:` is the acceptance criteria; `Verify:` is the mechanically enforceable subset. The gap between them is the review surface.
 - `Read first:` is context, not scope — 2–4 entries of ADRs, code areas, or rulings.
diff --git a/README.md b/README.md
index 1300a1bc12102f97cd89ace07333f179f8700cee..c1541df3127dce0ba5ba94a9cc4d8ba97c5a6b6a
--- a/README.md
+++ b/README.md
@@ -43,6 +43,9 @@
 6. On failure: the unit is marked `verify_failed` or `no_progress`, evidence is appended, and the loop stops or retries once.
 7. The loop halts on: max ticks reached with pending work, two no-progress strikes, or a failed verify after a real change.
 8. On any non-clean exit, the runner writes `HANDOFF.md` with completed/in-progress/remaining units and the next action.
+9. If `--review` is set, the runner invokes review after the queue drains, reads the actionable count from `REVIEW.md`, invokes fix when actionable findings exist, and runs another build pass for appended units.
+
+Review is opt-in. Without `--review`, the loop runs build ticks only; with it, review/fix are loop-orchestrated but still skill-owned.
 
 ## Queue format
 
@@ -70,6 +73,15 @@
 Agent: pi -p --no-session --approve --model glm-5.2 "$(cat "$LOOP_PROMPT_FILE")"
 ```
 
+Opt into review/fix orchestration with `--review`:
+
+```bash
+LOOP_AGENT_CMD='codex exec --dangerously-bypass-approvals-and-sandbox --ephemeral "$(cat "$LOOP_PROMPT_FILE")"' \
+LOOP_REVIEW_CMD='claude --print --no-session-persistence --dangerously-skip-permissions "$(cat "$LOOP_PROMPT_FILE")"' \
+LOOP_FIX_CMD='codex exec --dangerously-bypass-approvals-and-sandbox --ephemeral "$(cat "$LOOP_PROMPT_FILE")"' \
+./loop.sh run .loop/<name>/QUEUE.md --review --max-review-rounds 2
+```
+
 ## CLI
 
 The `knack` binary is a read-only validator and context provider. Build it from `cli/`:
diff --git a/docs/loop.md b/docs/loop.md
index a5cd2cd28883321ca006489e2c630fbdc8d6bb8f..ef47822f4007c3fbcb5ee094f39164166ca63646
--- a/docs/loop.md
+++ b/docs/loop.md
@@ -5,18 +5,24 @@
 ## Usage
 
 ```bash
-./loop.sh run <queue> [--repo DIR] [--max-ticks N] [--dry-run]
+./loop.sh run <queue> [--repo DIR] [--max-ticks N] [--review] [--max-review-rounds N] [--dry-run]
 ```
 
 - `queue` — path to `QUEUE.md`. Usually `.loop/<name>/QUEUE.md`.
 - `--repo DIR` — the working directory for the worker and the verify command. Defaults to the parent of the `.loop/<name>` directory that contains the queue.
 - `--max-ticks N` — maximum units to attempt. Default is `3`.
+- `--review` — opt into the bounded review/fix subloop after pending build units drain. Default behavior does not run review.
+- `--max-review-rounds N` — maximum review/fix rounds when `--review` is enabled. Default is `2`.
 - `--dry-run` — parse the first pending unit and print its title, repo, and verify command, then exit.
 
 ## Environment variables
 
 - `LOOP_AGENT_CMD` — optional command used to invoke the worker. If unset, the loop defaults to `pi -p --no-session --approve` with the prompt text as a single argument. If set, the command is evaluated by `bash -lc` in the repo directory, and `LOOP_PROMPT_FILE` is set to a temporary file containing `prompts/worker.md` plus the current work unit.
 - `LOOP_PROMPT_FILE` — set by the loop when `LOOP_AGENT_CMD` is used. Points to the generated prompt file. Do not override unless you are calling the worker manually.
+- `LOOP_REVIEW_CMD` — optional command used for the review phase when `--review` is enabled. Defaults to `LOOP_AGENT_CMD`, or to the loop's default `pi` invocation if neither is set.
+- `LOOP_FIX_CMD` — optional command used for the fix phase when `--review` is enabled. Defaults to `LOOP_AGENT_CMD`, or to the loop's default `pi` invocation if neither is set.
+
+During every phase, the loop also sets `LOOP_PHASE`, `LOOP_QUEUE_FILE`, `LOOP_EVIDENCE_FILE`, and `LOOP_REVIEW_FILE` for the worker process.
 
 ## Per-unit agent override
 
@@ -43,6 +49,20 @@
    - If the repo changed, mark `verify_failed` and stop.
 9. If max ticks is reached with pending work, stop.
 
+## Optional review/fix subloop
+
+By default, `loop.sh` stops when the build queue has no pending work. With `--review`, it then runs a bounded review/fix subloop:
+
+1. Invoke a review worker with the review prompt and the completed queue/evidence.
+2. Require the review worker to write `.loop/<name>/REVIEW.md`.
+3. Read only the `- actionable: N` summary line from `REVIEW.md`.
+4. If `N` is `0`, stop cleanly.
+5. If `N` is non-zero, invoke a fix worker.
+6. Require the fix worker to append new `Status: pending` work units to the same `QUEUE.md`.
+7. Run the build pass again, then review again.
+
+The subloop stops when review is clean, `--max-review-rounds` is reached, `--max-ticks` is exhausted, or fix produces no pending work. The loop owns orchestration and stop conditions only; the `review` and `fix` skills own judgment and work-unit generation.
+
 ## Work unit statuses
 
 - `pending` — not yet started.
@@ -58,6 +78,7 @@
 
 - `EVIDENCE.md` — append-only ledger. Includes the full unit, changed files, verify command, verify output, and worker output. It is durable; keep it after deleting `QUEUE.md` so `knack decisions check` can still see which ADRs the cycle referenced.
 - `HANDOFF.md` — written on non-clean exit. Sections: completed, in progress, remaining, next action. Delete when the work resumes.
+- `REVIEW.md` — structured review artifact written by the review worker when `--review` is enabled. The loop reads only its `- actionable: N` summary line.
 
 ## Agent invocation examples
 
@@ -71,8 +92,17 @@
 
 The command is passed to `bash -lc` in the repo directory, so the typical pattern is `"$(cat "$LOOP_PROMPT_FILE")"`.
 
+Review and fix can use separate agents:
+
+```bash
+LOOP_AGENT_CMD='codex exec --dangerously-bypass-approvals-and-sandbox --ephemeral "$(cat "$LOOP_PROMPT_FILE")"' \
+LOOP_REVIEW_CMD='claude --print --no-session-persistence --dangerously-skip-permissions "$(cat "$LOOP_PROMPT_FILE")"' \
+LOOP_FIX_CMD='codex exec --dangerously-bypass-approvals-and-sandbox --ephemeral "$(cat "$LOOP_PROMPT_FILE")"' \
+./loop.sh run .loop/<name>/QUEUE.md --review
+```
+
 ## Verification notes
 
 - `loop.sh` does **not** validate `QUEUE.md` structure. Use `knack validate` first.
-- `loop.sh` does **not** run review or manage ADRs/glossary. Those are handled by the skills.
+- `loop.sh` runs review and fix only when `--review` is set. It invokes the skills and reads the actionable count from `REVIEW.md`; it does not judge findings or manage ADRs/glossary.
 - The `Verify:` command must be deterministic and executable by the runner — tests, builds, type checks, not an LLM-as-judge.
diff --git a/docs/skills.md b/docs/skills.md
index fb1ec98119e1193859c0a1612f580d2da5afe95d..476d494e3c88962825e751b84151a0ff48ac2d41
--- a/docs/skills.md
+++ b/docs/skills.md
@@ -37,7 +37,11 @@
 
 ## How the loop uses skills
 
-`loop.sh` does not read skills. It prepends `prompts/worker.md` to the current work unit and runs the worker. `prompts/worker.md` tells the worker to load the `build` skill by name and path (e.g. "Load and follow the **build** skill in `.agents/skills/build/`"). For other phases (explore, plan, review, fix), the human or the agent invokes the skill directly.
+`loop.sh` does not read skills. It prepends `prompts/worker.md` to the current work unit and runs the worker. `prompts/worker.md` tells the worker to load the `build` skill by name and path (e.g. "Load and follow the **build** skill in `.agents/skills/build/`").
+
+When `--review` is set, the loop also invokes review and fix workers after the build queue drains. Those prompts tell the worker to load the `review` or `fix` skill directly. The loop orchestrates the bounded review/fix subloop, reads the actionable count from `REVIEW.md`, and runs another build pass when fix appends pending units. The skills still own judgment: review decides what the findings are, and fix decides which findings become work units.
+
+Without `--review`, review and fix remain manual skill invocations.
 
 ## Customizing skills
 
@@ -53,7 +57,7 @@
 ```text
 small fix → plan → build → done
 bug report → explore → plan → build → done
-big feature → explore → plan → build → review → fix → done
+big feature → explore → plan → build --review → review → fix → build → done
 ```
 
 Decisions are captured inline throughout the flow using the `decide` skill, and terms are updated using `domain-modeling`.

diff --git a/AGENTS.md b/AGENTS.md
index 50f568dd4d4d4e617b4be191116bb8fea556720d..019d79a3530e1a7621ecda0e6bfcdf1bd9e53ffa
--- a/AGENTS.md
+++ b/AGENTS.md
@@ -17,15 +17,16 @@
 
 See `DESIGN.md` for the full design.
 
-## Current state
+## Current State
 
-The loop (`loop.sh`), all seven skills, and the Go CLI are built. The loop supports per-unit `Agent:` overrides, `LOOP_AGENT_CMD` for agent-agnosticism, and writes `.loop/<name>/HANDOFF.md` on non-clean exits. The skills cover the full explore → plan → build → review → fix flow plus two shared skills (decide, domain-modeling). The CLI (`cli/`) provides `skills init|check`, `validate`, `decisions list|show|check`, `status`, `glossary check`, and `instructions`. It embeds the default skills via `go:embed` and scaffolds them into target projects.
+The loop (`loop.sh`), all seven skills, and the Go CLI are built. The loop supports per-unit `Agent:` overrides, `LOOP_AGENT_CMD` for agent-agnosticism, opt-in review/fix orchestration via `--review`, and `.loop/<name>/HANDOFF.md` on non-clean exits. The skills cover the full explore → plan → build → review → fix flow plus two shared skills (decide, domain-modeling). The CLI (`cli/`) provides `skills init|check`, `validate`, `decisions list|show|check`, `status`, `glossary check`, and `instructions`. It embeds the default skills via `go:embed` and scaffolds them into target projects.
 
 ## Core artifacts
 
 - `.loop/<name>/QUEUE.md` — disposable queue of verifiable work units.
 - `.loop/<name>/EVIDENCE.md` — append-only ledger of what each tick proved. Durable: keep it after deleting QUEUE.md so `decisions check` can still see which ADRs the completed cycle referenced.
 - `.loop/<name>/HANDOFF.md` — cross-session handoff (written on pause/stop).
+- `.loop/<name>/REVIEW.md` — structured review artifact written when `--review` is enabled. The loop reads only the actionable count.
 - `decisions/` — durable ADRs (architectural rulings, not current behavior).
 - `glossary.md` — durable ubiquitous language.
 - `.agents/skills/` — procedural knowledge (explore, plan, build, review, fix, decide, domain-modeling).
@@ -48,11 +49,13 @@
 - The worker never self-certifies. The runner owns verification.
 - A `Verify` command must be deterministic and executable by the runner — not an LLM-as-judge. The loop's backpressure is mechanical.
 - The runner enforces hard stops (max ticks, no-progress detection). The agent does one work unit per tick and reports what remains if it can't finish.
-- Review is opt-in. The loop does not run review automatically; the user invokes it when they want adversarial scrutiny.
+- Review is opt-in. The loop runs review/fix only when invoked with `--review`; otherwise it runs build ticks only.
+- With `--review`, the loop invokes review after pending build units drain, reads only `REVIEW.md`'s actionable count, invokes fix when actionable findings exist, then runs another build pass for appended units. Judgment stays in the `review` and `fix` skills.
 - When the queue is complete and verified, the cycle's `.loop/<name>/QUEUE.md` is disposable. The human deletes it; the loop does not. `EVIDENCE.md` is the durable ledger — keep it so `decisions check` can still trace which ADRs the cycle referenced.
 - A work unit must leave the repo better if the loop stops immediately after it.
 - Work units are whatever shape the work is — not forced into "vertical slices."
 - `LOOP_AGENT_CMD` overrides the worker invocation (agent-agnostic: pi, claude, codex, opencode, etc.).
+- `LOOP_REVIEW_CMD` and `LOOP_FIX_CMD` override review/fix phase invocations when `--review` is enabled. They default to `LOOP_AGENT_CMD`.
 - Per-unit `Agent:` field in QUEUE.md overrides `LOOP_AGENT_CMD` for one unit (model routing).
 - Work units are `## <outcome>` headers with `Read first:`, `Constraints:`, `Done means:`, and `Verify:` fields. `Done means:` is the acceptance criteria; `Verify:` is the mechanically enforceable subset. The gap between them is the review surface.
 - `Read first:` is context, not scope — 2–4 entries of ADRs, code areas, or rulings.
diff --git a/README.md b/README.md
index 1300a1bc12102f97cd89ace07333f179f8700cee..c1541df3127dce0ba5ba94a9cc4d8ba97c5a6b6a
--- a/README.md
+++ b/README.md
@@ -43,6 +43,9 @@
 6. On failure: the unit is marked `verify_failed` or `no_progress`, evidence is appended, and the loop stops or retries once.
 7. The loop halts on: max ticks reached with pending work, two no-progress strikes, or a failed verify after a real change.
 8. On any non-clean exit, the runner writes `HANDOFF.md` with completed/in-progress/remaining units and the next action.
+9. If `--review` is set, the runner invokes review after the queue drains, reads the actionable count from `REVIEW.md`, invokes fix when actionable findings exist, and runs another build pass for appended units.
+
+Review is opt-in. Without `--review`, the loop runs build ticks only; with it, review/fix are loop-orchestrated but still skill-owned.
 
 ## Queue format
 
@@ -70,6 +73,15 @@
 Agent: pi -p --no-session --approve --model glm-5.2 "$(cat "$LOOP_PROMPT_FILE")"
 ```
 
+Opt into review/fix orchestration with `--review`:
+
+```bash
+LOOP_AGENT_CMD='codex exec --dangerously-bypass-approvals-and-sandbox --ephemeral "$(cat "$LOOP_PROMPT_FILE")"' \
+LOOP_REVIEW_CMD='claude --print --no-session-persistence --dangerously-skip-permissions "$(cat "$LOOP_PROMPT_FILE")"' \
+LOOP_FIX_CMD='codex exec --dangerously-bypass-approvals-and-sandbox --ephemeral "$(cat "$LOOP_PROMPT_FILE")"' \
+./loop.sh run .loop/<name>/QUEUE.md --review --max-review-rounds 2
+```
+
 ## CLI
 
 The `knack` binary is a read-only validator and context provider. Build it from `cli/`:
diff --git a/docs/loop.md b/docs/loop.md
index a5cd2cd28883321ca006489e2c630fbdc8d6bb8f..ef47822f4007c3fbcb5ee094f39164166ca63646
--- a/docs/loop.md
+++ b/docs/loop.md
@@ -5,18 +5,24 @@
 ## Usage
 
 ```bash
-./loop.sh run <queue> [--repo DIR] [--max-ticks N] [--dry-run]
+./loop.sh run <queue> [--repo DIR] [--max-ticks N] [--review] [--max-review-rounds N] [--dry-run]
 ```
 
 - `queue` — path to `QUEUE.md`. Usually `.loop/<name>/QUEUE.md`.
 - `--repo DIR` — the working directory for the worker and the verify command. Defaults to the parent of the `.loop/<name>` directory that contains the queue.
 - `--max-ticks N` — maximum units to attempt. Default is `3`.
+- `--review` — opt into the bounded review/fix subloop after pending build units drain. Default behavior does not run review.
+- `--max-review-rounds N` — maximum review/fix rounds when `--review` is enabled. Default is `2`.
 - `--dry-run` — parse the first pending unit and print its title, repo, and verify command, then exit.
 
 ## Environment variables
 
 - `LOOP_AGENT_CMD` — optional command used to invoke the worker. If unset, the loop defaults to `pi -p --no-session --approve` with the prompt text as a single argument. If set, the command is evaluated by `bash -lc` in the repo directory, and `LOOP_PROMPT_FILE` is set to a temporary file containing `prompts/worker.md` plus the current work unit.
 - `LOOP_PROMPT_FILE` — set by the loop when `LOOP_AGENT_CMD` is used. Points to the generated prompt file. Do not override unless you are calling the worker manually.
+- `LOOP_REVIEW_CMD` — optional command used for the review phase when `--review` is enabled. Defaults to `LOOP_AGENT_CMD`, or to the loop's default `pi` invocation if neither is set.
+- `LOOP_FIX_CMD` — optional command used for the fix phase when `--review` is enabled. Defaults to `LOOP_AGENT_CMD`, or to the loop's default `pi` invocation if neither is set.
+
+During every phase, the loop also sets `LOOP_PHASE`, `LOOP_QUEUE_FILE`, `LOOP_EVIDENCE_FILE`, and `LOOP_REVIEW_FILE` for the worker process.
 
 ## Per-unit agent override
 
@@ -43,6 +49,20 @@
    - If the repo changed, mark `verify_failed` and stop.
 9. If max ticks is reached with pending work, stop.
 
+## Optional review/fix subloop
+
+By default, `loop.sh` stops when the build queue has no pending work. With `--review`, it then runs a bounded review/fix subloop:
+
+1. Invoke a review worker with the review prompt and the completed queue/evidence.
+2. Require the review worker to write `.loop/<name>/REVIEW.md`.
+3. Read only the `- actionable: N` summary line from `REVIEW.md`.
+4. If `N` is `0`, stop cleanly.
+5. If `N` is non-zero, invoke a fix worker.
+6. Require the fix worker to append new `Status: pending` work units to the same `QUEUE.md`.
+7. Run the build pass again, then review again.
+
+The subloop stops when review is clean, `--max-review-rounds` is reached, `--max-ticks` is exhausted, or fix produces no pending work. The loop owns orchestration and stop conditions only; the `review` and `fix` skills own judgment and work-unit generation.
+
 ## Work unit statuses
 
 - `pending` — not yet started.
@@ -58,6 +78,7 @@
 
 - `EVIDENCE.md` — append-only ledger. Includes the full unit, changed files, verify command, verify output, and worker output. It is durable; keep it after deleting `QUEUE.md` so `knack decisions check` can still see which ADRs the cycle referenced.
 - `HANDOFF.md` — written on non-clean exit. Sections: completed, in progress, remaining, next action. Delete when the work resumes.
+- `REVIEW.md` — structured review artifact written by the review worker when `--review` is enabled. The loop reads only its `- actionable: N` summary line.
 
 ## Agent invocation examples
 
@@ -71,8 +92,17 @@
 
 The command is passed to `bash -lc` in the repo directory, so the typical pattern is `"$(cat "$LOOP_PROMPT_FILE")"`.
 
+Review and fix can use separate agents:
+
+```bash
+LOOP_AGENT_CMD='codex exec --dangerously-bypass-approvals-and-sandbox --ephemeral "$(cat "$LOOP_PROMPT_FILE")"' \
+LOOP_REVIEW_CMD='claude --print --no-session-persistence --dangerously-skip-permissions "$(cat "$LOOP_PROMPT_FILE")"' \
+LOOP_FIX_CMD='codex exec --dangerously-bypass-approvals-and-sandbox --ephemeral "$(cat "$LOOP_PROMPT_FILE")"' \
+./loop.sh run .loop/<name>/QUEUE.md --review
+```
+
 ## Verification notes
 
 - `loop.sh` does **not** validate `QUEUE.md` structure. Use `knack validate` first.
-- `loop.sh` does **not** run review or manage ADRs/glossary. Those are handled by the skills.
+- `loop.sh` runs review and fix only when `--review` is set. It invokes the skills and reads the actionable count from `REVIEW.md`; it does not judge findings or manage ADRs/glossary.
 - The `Verify:` command must be deterministic and executable by the runner — tests, builds, type checks, not an LLM-as-judge.
diff --git a/docs/skills.md b/docs/skills.md
index fb1ec98119e1193859c0a1612f580d2da5afe95d..476d494e3c88962825e751b84151a0ff48ac2d41
--- a/docs/skills.md
+++ b/docs/skills.md
@@ -37,7 +37,11 @@
 
 ## How the loop uses skills
 
-`loop.sh` does not read skills. It prepends `prompts/worker.md` to the current work unit and runs the worker. `prompts/worker.md` tells the worker to load the `build` skill by name and path (e.g. "Load and follow the **build** skill in `.agents/skills/build/`"). For other phases (explore, plan, review, fix), the human or the agent invokes the skill directly.
+`loop.sh` does not read skills. It prepends `prompts/worker.md` to the current work unit and runs the worker. `prompts/worker.md` tells the worker to load the `build` skill by name and path (e.g. "Load and follow the **build** skill in `.agents/skills/build/`").
+
+When `--review` is set, the loop also invokes review and fix workers after the build queue drains. Those prompts tell the worker to load the `review` or `fix` skill directly. The loop orchestrates the bounded review/fix subloop, reads the actionable count from `REVIEW.md`, and runs another build pass when fix appends pending units. The skills still own judgment: review decides what the findings are, and fix decides which findings become work units.
+
+Without `--review`, review and fix remain manual skill invocations.
 
 ## Customizing skills
 
@@ -53,7 +57,7 @@
 ```text
 small fix → plan → build → done
 bug report → explore → plan → build → done
-big feature → explore → plan → build → review → fix → done
+big feature → explore → plan → build --review → review → fix → build → done
 ```
 
 Decisions are captured inline throughout the flow using the `decide` skill, and terms are updated using `domain-modeling`.

codex
I noticed one incidental heading case change in AGENTS.md while patching. I’m cleaning that up before checking the diff and running the unit’s verify command.
apply patch
patch: completed
/home/daniel/build/knack/AGENTS.md
diff --git a/AGENTS.md b/AGENTS.md
index 50f568dd4d4d4e617b4be191116bb8fea556720d..f005f47ed72aa4f97f47324d98c00d8c79b12921
--- a/AGENTS.md
+++ b/AGENTS.md
@@ -19,13 +19,14 @@
 
 ## Current state
 
-The loop (`loop.sh`), all seven skills, and the Go CLI are built. The loop supports per-unit `Agent:` overrides, `LOOP_AGENT_CMD` for agent-agnosticism, and writes `.loop/<name>/HANDOFF.md` on non-clean exits. The skills cover the full explore → plan → build → review → fix flow plus two shared skills (decide, domain-modeling). The CLI (`cli/`) provides `skills init|check`, `validate`, `decisions list|show|check`, `status`, `glossary check`, and `instructions`. It embeds the default skills via `go:embed` and scaffolds them into target projects.
+The loop (`loop.sh`), all seven skills, and the Go CLI are built. The loop supports per-unit `Agent:` overrides, `LOOP_AGENT_CMD` for agent-agnosticism, opt-in review/fix orchestration via `--review`, and `.loop/<name>/HANDOFF.md` on non-clean exits. The skills cover the full explore → plan → build → review → fix flow plus two shared skills (decide, domain-modeling). The CLI (`cli/`) provides `skills init|check`, `validate`, `decisions list|show|check`, `status`, `glossary check`, and `instructions`. It embeds the default skills via `go:embed` and scaffolds them into target projects.
 
 ## Core artifacts
 
 - `.loop/<name>/QUEUE.md` — disposable queue of verifiable work units.
 - `.loop/<name>/EVIDENCE.md` — append-only ledger of what each tick proved. Durable: keep it after deleting QUEUE.md so `decisions check` can still see which ADRs the completed cycle referenced.
 - `.loop/<name>/HANDOFF.md` — cross-session handoff (written on pause/stop).
+- `.loop/<name>/REVIEW.md` — structured review artifact written when `--review` is enabled. The loop reads only the actionable count.
 - `decisions/` — durable ADRs (architectural rulings, not current behavior).
 - `glossary.md` — durable ubiquitous language.
 - `.agents/skills/` — procedural knowledge (explore, plan, build, review, fix, decide, domain-modeling).
@@ -48,11 +49,13 @@
 - The worker never self-certifies. The runner owns verification.
 - A `Verify` command must be deterministic and executable by the runner — not an LLM-as-judge. The loop's backpressure is mechanical.
 - The runner enforces hard stops (max ticks, no-progress detection). The agent does one work unit per tick and reports what remains if it can't finish.
-- Review is opt-in. The loop does not run review automatically; the user invokes it when they want adversarial scrutiny.
+- Review is opt-in. The loop runs review/fix only when invoked with `--review`; otherwise it runs build ticks only.
+- With `--review`, the loop invokes review after pending build units drain, reads only `REVIEW.md`'s actionable count, invokes fix when actionable findings exist, then runs another build pass for appended units. Judgment stays in the `review` and `fix` skills.
 - When the queue is complete and verified, the cycle's `.loop/<name>/QUEUE.md` is disposable. The human deletes it; the loop does not. `EVIDENCE.md` is the durable ledger — keep it so `decisions check` can still trace which ADRs the cycle referenced.
 - A work unit must leave the repo better if the loop stops immediately after it.
 - Work units are whatever shape the work is — not forced into "vertical slices."
 - `LOOP_AGENT_CMD` overrides the worker invocation (agent-agnostic: pi, claude, codex, opencode, etc.).
+- `LOOP_REVIEW_CMD` and `LOOP_FIX_CMD` override review/fix phase invocations when `--review` is enabled. They default to `LOOP_AGENT_CMD`.
 - Per-unit `Agent:` field in QUEUE.md overrides `LOOP_AGENT_CMD` for one unit (model routing).
 - Work units are `## <outcome>` headers with `Read first:`, `Constraints:`, `Done means:`, and `Verify:` fields. `Done means:` is the acceptance criteria; `Verify:` is the mechanically enforceable subset. The gap between them is the review surface.
 - `Read first:` is context, not scope — 2–4 entries of ADRs, code areas, or rulings.
diff --git a/README.md b/README.md
index 1300a1bc12102f97cd89ace07333f179f8700cee..c1541df3127dce0ba5ba94a9cc4d8ba97c5a6b6a
--- a/README.md
+++ b/README.md
@@ -43,6 +43,9 @@
 6. On failure: the unit is marked `verify_failed` or `no_progress`, evidence is appended, and the loop stops or retries once.
 7. The loop halts on: max ticks reached with pending work, two no-progress strikes, or a failed verify after a real change.
 8. On any non-clean exit, the runner writes `HANDOFF.md` with completed/in-progress/remaining units and the next action.
+9. If `--review` is set, the runner invokes review after the queue drains, reads the actionable count from `REVIEW.md`, invokes fix when actionable findings exist, and runs another build pass for appended units.
+
+Review is opt-in. Without `--review`, the loop runs build ticks only; with it, review/fix are loop-orchestrated but still skill-owned.
 
 ## Queue format
 
@@ -70,6 +73,15 @@
 Agent: pi -p --no-session --approve --model glm-5.2 "$(cat "$LOOP_PROMPT_FILE")"
 ```
 
+Opt into review/fix orchestration with `--review`:
+
+```bash
+LOOP_AGENT_CMD='codex exec --dangerously-bypass-approvals-and-sandbox --ephemeral "$(cat "$LOOP_PROMPT_FILE")"' \
+LOOP_REVIEW_CMD='claude --print --no-session-persistence --dangerously-skip-permissions "$(cat "$LOOP_PROMPT_FILE")"' \
+LOOP_FIX_CMD='codex exec --dangerously-bypass-approvals-and-sandbox --ephemeral "$(cat "$LOOP_PROMPT_FILE")"' \
+./loop.sh run .loop/<name>/QUEUE.md --review --max-review-rounds 2
+```
+
 ## CLI
 
 The `knack` binary is a read-only validator and context provider. Build it from `cli/`:
diff --git a/docs/loop.md b/docs/loop.md
index a5cd2cd28883321ca006489e2c630fbdc8d6bb8f..ef47822f4007c3fbcb5ee094f39164166ca63646
--- a/docs/loop.md
+++ b/docs/loop.md
@@ -5,18 +5,24 @@
 ## Usage
 
 ```bash
-./loop.sh run <queue> [--repo DIR] [--max-ticks N] [--dry-run]
+./loop.sh run <queue> [--repo DIR] [--max-ticks N] [--review] [--max-review-rounds N] [--dry-run]
 ```
 
 - `queue` — path to `QUEUE.md`. Usually `.loop/<name>/QUEUE.md`.
 - `--repo DIR` — the working directory for the worker and the verify command. Defaults to the parent of the `.loop/<name>` directory that contains the queue.
 - `--max-ticks N` — maximum units to attempt. Default is `3`.
+- `--review` — opt into the bounded review/fix subloop after pending build units drain. Default behavior does not run review.
+- `--max-review-rounds N` — maximum review/fix rounds when `--review` is enabled. Default is `2`.
 - `--dry-run` — parse the first pending unit and print its title, repo, and verify command, then exit.
 
 ## Environment variables
 
 - `LOOP_AGENT_CMD` — optional command used to invoke the worker. If unset, the loop defaults to `pi -p --no-session --approve` with the prompt text as a single argument. If set, the command is evaluated by `bash -lc` in the repo directory, and `LOOP_PROMPT_FILE` is set to a temporary file containing `prompts/worker.md` plus the current work unit.
 - `LOOP_PROMPT_FILE` — set by the loop when `LOOP_AGENT_CMD` is used. Points to the generated prompt file. Do not override unless you are calling the worker manually.
+- `LOOP_REVIEW_CMD` — optional command used for the review phase when `--review` is enabled. Defaults to `LOOP_AGENT_CMD`, or to the loop's default `pi` invocation if neither is set.
+- `LOOP_FIX_CMD` — optional command used for the fix phase when `--review` is enabled. Defaults to `LOOP_AGENT_CMD`, or to the loop's default `pi` invocation if neither is set.
+
+During every phase, the loop also sets `LOOP_PHASE`, `LOOP_QUEUE_FILE`, `LOOP_EVIDENCE_FILE`, and `LOOP_REVIEW_FILE` for the worker process.
 
 ## Per-unit agent override
 
@@ -43,6 +49,20 @@
    - If the repo changed, mark `verify_failed` and stop.
 9. If max ticks is reached with pending work, stop.
 
+## Optional review/fix subloop
+
+By default, `loop.sh` stops when the build queue has no pending work. With `--review`, it then runs a bounded review/fix subloop:
+
+1. Invoke a review worker with the review prompt and the completed queue/evidence.
+2. Require the review worker to write `.loop/<name>/REVIEW.md`.
+3. Read only the `- actionable: N` summary line from `REVIEW.md`.
+4. If `N` is `0`, stop cleanly.
+5. If `N` is non-zero, invoke a fix worker.
+6. Require the fix worker to append new `Status: pending` work units to the same `QUEUE.md`.
+7. Run the build pass again, then review again.
+
+The subloop stops when review is clean, `--max-review-rounds` is reached, `--max-ticks` is exhausted, or fix produces no pending work. The loop owns orchestration and stop conditions only; the `review` and `fix` skills own judgment and work-unit generation.
+
 ## Work unit statuses
 
 - `pending` — not yet started.
@@ -58,6 +78,7 @@
 
 - `EVIDENCE.md` — append-only ledger. Includes the full unit, changed files, verify command, verify output, and worker output. It is durable; keep it after deleting `QUEUE.md` so `knack decisions check` can still see which ADRs the cycle referenced.
 - `HANDOFF.md` — written on non-clean exit. Sections: completed, in progress, remaining, next action. Delete when the work resumes.
+- `REVIEW.md` — structured review artifact written by the review worker when `--review` is enabled. The loop reads only its `- actionable: N` summary line.
 
 ## Agent invocation examples
 
@@ -71,8 +92,17 @@
 
 The command is passed to `bash -lc` in the repo directory, so the typical pattern is `"$(cat "$LOOP_PROMPT_FILE")"`.
 
+Review and fix can use separate agents:
+
+```bash
+LOOP_AGENT_CMD='codex exec --dangerously-bypass-approvals-and-sandbox --ephemeral "$(cat "$LOOP_PROMPT_FILE")"' \
+LOOP_REVIEW_CMD='claude --print --no-session-persistence --dangerously-skip-permissions "$(cat "$LOOP_PROMPT_FILE")"' \
+LOOP_FIX_CMD='codex exec --dangerously-bypass-approvals-and-sandbox --ephemeral "$(cat "$LOOP_PROMPT_FILE")"' \
+./loop.sh run .loop/<name>/QUEUE.md --review
+```
+
 ## Verification notes
 
 - `loop.sh` does **not** validate `QUEUE.md` structure. Use `knack validate` first.
-- `loop.sh` does **not** run review or manage ADRs/glossary. Those are handled by the skills.
+- `loop.sh` runs review and fix only when `--review` is set. It invokes the skills and reads the actionable count from `REVIEW.md`; it does not judge findings or manage ADRs/glossary.
 - The `Verify:` command must be deterministic and executable by the runner — tests, builds, type checks, not an LLM-as-judge.
diff --git a/docs/skills.md b/docs/skills.md
index fb1ec98119e1193859c0a1612f580d2da5afe95d..476d494e3c88962825e751b84151a0ff48ac2d41
--- a/docs/skills.md
+++ b/docs/skills.md
@@ -37,7 +37,11 @@
 
 ## How the loop uses skills
 
-`loop.sh` does not read skills. It prepends `prompts/worker.md` to the current work unit and runs the worker. `prompts/worker.md` tells the worker to load the `build` skill by name and path (e.g. "Load and follow the **build** skill in `.agents/skills/build/`"). For other phases (explore, plan, review, fix), the human or the agent invokes the skill directly.
+`loop.sh` does not read skills. It prepends `prompts/worker.md` to the current work unit and runs the worker. `prompts/worker.md` tells the worker to load the `build` skill by name and path (e.g. "Load and follow the **build** skill in `.agents/skills/build/`").
+
+When `--review` is set, the loop also invokes review and fix workers after the build queue drains. Those prompts tell the worker to load the `review` or `fix` skill directly. The loop orchestrates the bounded review/fix subloop, reads the actionable count from `REVIEW.md`, and runs another build pass when fix appends pending units. The skills still own judgment: review decides what the findings are, and fix decides which findings become work units.
+
+Without `--review`, review and fix remain manual skill invocations.
 
 ## Customizing skills
 
@@ -53,7 +57,7 @@
 ```text
 small fix → plan → build → done
 bug report → explore → plan → build → done
-big feature → explore → plan → build → review → fix → done
+big feature → explore → plan → build --review → review → fix → build → done
 ```
 
 Decisions are captured inline throughout the flow using the `decide` skill, and terms are updated using `domain-modeling`.

diff --git a/AGENTS.md b/AGENTS.md
index 50f568dd4d4d4e617b4be191116bb8fea556720d..f005f47ed72aa4f97f47324d98c00d8c79b12921
--- a/AGENTS.md
+++ b/AGENTS.md
@@ -19,13 +19,14 @@
 
 ## Current state
 
-The loop (`loop.sh`), all seven skills, and the Go CLI are built. The loop supports per-unit `Agent:` overrides, `LOOP_AGENT_CMD` for agent-agnosticism, and writes `.loop/<name>/HANDOFF.md` on non-clean exits. The skills cover the full explore → plan → build → review → fix flow plus two shared skills (decide, domain-modeling). The CLI (`cli/`) provides `skills init|check`, `validate`, `decisions list|show|check`, `status`, `glossary check`, and `instructions`. It embeds the default skills via `go:embed` and scaffolds them into target projects.
+The loop (`loop.sh`), all seven skills, and the Go CLI are built. The loop supports per-unit `Agent:` overrides, `LOOP_AGENT_CMD` for agent-agnosticism, opt-in review/fix orchestration via `--review`, and `.loop/<name>/HANDOFF.md` on non-clean exits. The skills cover the full explore → plan → build → review → fix flow plus two shared skills (decide, domain-modeling). The CLI (`cli/`) provides `skills init|check`, `validate`, `decisions list|show|check`, `status`, `glossary check`, and `instructions`. It embeds the default skills via `go:embed` and scaffolds them into target projects.
 
 ## Core artifacts
 
 - `.loop/<name>/QUEUE.md` — disposable queue of verifiable work units.
 - `.loop/<name>/EVIDENCE.md` — append-only ledger of what each tick proved. Durable: keep it after deleting QUEUE.md so `decisions check` can still see which ADRs the completed cycle referenced.
 - `.loop/<name>/HANDOFF.md` — cross-session handoff (written on pause/stop).
+- `.loop/<name>/REVIEW.md` — structured review artifact written when `--review` is enabled. The loop reads only the actionable count.
 - `decisions/` — durable ADRs (architectural rulings, not current behavior).
 - `glossary.md` — durable ubiquitous language.
 - `.agents/skills/` — procedural knowledge (explore, plan, build, review, fix, decide, domain-modeling).
@@ -48,11 +49,13 @@
 - The worker never self-certifies. The runner owns verification.
 - A `Verify` command must be deterministic and executable by the runner — not an LLM-as-judge. The loop's backpressure is mechanical.
 - The runner enforces hard stops (max ticks, no-progress detection). The agent does one work unit per tick and reports what remains if it can't finish.
-- Review is opt-in. The loop does not run review automatically; the user invokes it when they want adversarial scrutiny.
+- Review is opt-in. The loop runs review/fix only when invoked with `--review`; otherwise it runs build ticks only.
+- With `--review`, the loop invokes review after pending build units drain, reads only `REVIEW.md`'s actionable count, invokes fix when actionable findings exist, then runs another build pass for appended units. Judgment stays in the `review` and `fix` skills.
 - When the queue is complete and verified, the cycle's `.loop/<name>/QUEUE.md` is disposable. The human deletes it; the loop does not. `EVIDENCE.md` is the durable ledger — keep it so `decisions check` can still trace which ADRs the cycle referenced.
 - A work unit must leave the repo better if the loop stops immediately after it.
 - Work units are whatever shape the work is — not forced into "vertical slices."
 - `LOOP_AGENT_CMD` overrides the worker invocation (agent-agnostic: pi, claude, codex, opencode, etc.).
+- `LOOP_REVIEW_CMD` and `LOOP_FIX_CMD` override review/fix phase invocations when `--review` is enabled. They default to `LOOP_AGENT_CMD`.
 - Per-unit `Agent:` field in QUEUE.md overrides `LOOP_AGENT_CMD` for one unit (model routing).
 - Work units are `## <outcome>` headers with `Read first:`, `Constraints:`, `Done means:`, and `Verify:` fields. `Done means:` is the acceptance criteria; `Verify:` is the mechanically enforceable subset. The gap between them is the review surface.
 - `Read first:` is context, not scope — 2–4 entries of ADRs, code areas, or rulings.
diff --git a/README.md b/README.md
index 1300a1bc12102f97cd89ace07333f179f8700cee..c1541df3127dce0ba5ba94a9cc4d8ba97c5a6b6a
--- a/README.md
+++ b/README.md
@@ -43,6 +43,9 @@
 6. On failure: the unit is marked `verify_failed` or `no_progress`, evidence is appended, and the loop stops or retries once.
 7. The loop halts on: max ticks reached with pending work, two no-progress strikes, or a failed verify after a real change.
 8. On any non-clean exit, the runner writes `HANDOFF.md` with completed/in-progress/remaining units and the next action.
+9. If `--review` is set, the runner invokes review after the queue drains, reads the actionable count from `REVIEW.md`, invokes fix when actionable findings exist, and runs another build pass for appended units.
+
+Review is opt-in. Without `--review`, the loop runs build ticks only; with it, review/fix are loop-orchestrated but still skill-owned.
 
 ## Queue format
 
@@ -70,6 +73,15 @@
 Agent: pi -p --no-session --approve --model glm-5.2 "$(cat "$LOOP_PROMPT_FILE")"
 ```
 
+Opt into review/fix orchestration with `--review`:
+
+```bash
+LOOP_AGENT_CMD='codex exec --dangerously-bypass-approvals-and-sandbox --ephemeral "$(cat "$LOOP_PROMPT_FILE")"' \
+LOOP_REVIEW_CMD='claude --print --no-session-persistence --dangerously-skip-permissions "$(cat "$LOOP_PROMPT_FILE")"' \
+LOOP_FIX_CMD='codex exec --dangerously-bypass-approvals-and-sandbox --ephemeral "$(cat "$LOOP_PROMPT_FILE")"' \
+./loop.sh run .loop/<name>/QUEUE.md --review --max-review-rounds 2
+```
+
 ## CLI
 
 The `knack` binary is a read-only validator and context provider. Build it from `cli/`:
diff --git a/docs/loop.md b/docs/loop.md
index a5cd2cd28883321ca006489e2c630fbdc8d6bb8f..ef47822f4007c3fbcb5ee094f39164166ca63646
--- a/docs/loop.md
+++ b/docs/loop.md
@@ -5,18 +5,24 @@
 ## Usage
 
 ```bash
-./loop.sh run <queue> [--repo DIR] [--max-ticks N] [--dry-run]
+./loop.sh run <queue> [--repo DIR] [--max-ticks N] [--review] [--max-review-rounds N] [--dry-run]
 ```
 
 - `queue` — path to `QUEUE.md`. Usually `.loop/<name>/QUEUE.md`.
 - `--repo DIR` — the working directory for the worker and the verify command. Defaults to the parent of the `.loop/<name>` directory that contains the queue.
 - `--max-ticks N` — maximum units to attempt. Default is `3`.
+- `--review` — opt into the bounded review/fix subloop after pending build units drain. Default behavior does not run review.
+- `--max-review-rounds N` — maximum review/fix rounds when `--review` is enabled. Default is `2`.
 - `--dry-run` — parse the first pending unit and print its title, repo, and verify command, then exit.
 
 ## Environment variables
 
 - `LOOP_AGENT_CMD` — optional command used to invoke the worker. If unset, the loop defaults to `pi -p --no-session --approve` with the prompt text as a single argument. If set, the command is evaluated by `bash -lc` in the repo directory, and `LOOP_PROMPT_FILE` is set to a temporary file containing `prompts/worker.md` plus the current work unit.
 - `LOOP_PROMPT_FILE` — set by the loop when `LOOP_AGENT_CMD` is used. Points to the generated prompt file. Do not override unless you are calling the worker manually.
+- `LOOP_REVIEW_CMD` — optional command used for the review phase when `--review` is enabled. Defaults to `LOOP_AGENT_CMD`, or to the loop's default `pi` invocation if neither is set.
+- `LOOP_FIX_CMD` — optional command used for the fix phase when `--review` is enabled. Defaults to `LOOP_AGENT_CMD`, or to the loop's default `pi` invocation if neither is set.
+
+During every phase, the loop also sets `LOOP_PHASE`, `LOOP_QUEUE_FILE`, `LOOP_EVIDENCE_FILE`, and `LOOP_REVIEW_FILE` for the worker process.
 
 ## Per-unit agent override
 
@@ -43,6 +49,20 @@
    - If the repo changed, mark `verify_failed` and stop.
 9. If max ticks is reached with pending work, stop.
 
+## Optional review/fix subloop
+
+By default, `loop.sh` stops when the build queue has no pending work. With `--review`, it then runs a bounded review/fix subloop:
+
+1. Invoke a review worker with the review prompt and the completed queue/evidence.
+2. Require the review worker to write `.loop/<name>/REVIEW.md`.
+3. Read only the `- actionable: N` summary line from `REVIEW.md`.
+4. If `N` is `0`, stop cleanly.
+5. If `N` is non-zero, invoke a fix worker.
+6. Require the fix worker to append new `Status: pending` work units to the same `QUEUE.md`.
+7. Run the build pass again, then review again.
+
+The subloop stops when review is clean, `--max-review-rounds` is reached, `--max-ticks` is exhausted, or fix produces no pending work. The loop owns orchestration and stop conditions only; the `review` and `fix` skills own judgment and work-unit generation.
+
 ## Work unit statuses
 
 - `pending` — not yet started.
@@ -58,6 +78,7 @@
 
 - `EVIDENCE.md` — append-only ledger. Includes the full unit, changed files, verify command, verify output, and worker output. It is durable; keep it after deleting `QUEUE.md` so `knack decisions check` can still see which ADRs the cycle referenced.
 - `HANDOFF.md` — written on non-clean exit. Sections: completed, in progress, remaining, next action. Delete when the work resumes.
+- `REVIEW.md` — structured review artifact written by the review worker when `--review` is enabled. The loop reads only its `- actionable: N` summary line.
 
 ## Agent invocation examples
 
@@ -71,8 +92,17 @@
 
 The command is passed to `bash -lc` in the repo directory, so the typical pattern is `"$(cat "$LOOP_PROMPT_FILE")"`.
 
+Review and fix can use separate agents:
+
+```bash
+LOOP_AGENT_CMD='codex exec --dangerously-bypass-approvals-and-sandbox --ephemeral "$(cat "$LOOP_PROMPT_FILE")"' \
+LOOP_REVIEW_CMD='claude --print --no-session-persistence --dangerously-skip-permissions "$(cat "$LOOP_PROMPT_FILE")"' \
+LOOP_FIX_CMD='codex exec --dangerously-bypass-approvals-and-sandbox --ephemeral "$(cat "$LOOP_PROMPT_FILE")"' \
+./loop.sh run .loop/<name>/QUEUE.md --review
+```
+
 ## Verification notes
 
 - `loop.sh` does **not** validate `QUEUE.md` structure. Use `knack validate` first.
-- `loop.sh` does **not** run review or manage ADRs/glossary. Those are handled by the skills.
+- `loop.sh` runs review and fix only when `--review` is set. It invokes the skills and reads the actionable count from `REVIEW.md`; it does not judge findings or manage ADRs/glossary.
 - The `Verify:` command must be deterministic and executable by the runner — tests, builds, type checks, not an LLM-as-judge.
diff --git a/docs/skills.md b/docs/skills.md
index fb1ec98119e1193859c0a1612f580d2da5afe95d..476d494e3c88962825e751b84151a0ff48ac2d41
--- a/docs/skills.md
+++ b/docs/skills.md
@@ -37,7 +37,11 @@
 
 ## How the loop uses skills
 
-`loop.sh` does not read skills. It prepends `prompts/worker.md` to the current work unit and runs the worker. `prompts/worker.md` tells the worker to load the `build` skill by name and path (e.g. "Load and follow the **build** skill in `.agents/skills/build/`"). For other phases (explore, plan, review, fix), the human or the agent invokes the skill directly.
+`loop.sh` does not read skills. It prepends `prompts/worker.md` to the current work unit and runs the worker. `prompts/worker.md` tells the worker to load the `build` skill by name and path (e.g. "Load and follow the **build** skill in `.agents/skills/build/`").
+
+When `--review` is set, the loop also invokes review and fix workers after the build queue drains. Those prompts tell the worker to load the `review` or `fix` skill directly. The loop orchestrates the bounded review/fix subloop, reads the actionable count from `REVIEW.md`, and runs another build pass when fix appends pending units. The skills still own judgment: review decides what the findings are, and fix decides which findings become work units.
+
+Without `--review`, review and fix remain manual skill invocations.
 
 ## Customizing skills
 
@@ -53,7 +57,7 @@
 ```text
 small fix → plan → build → done
 bug report → explore → plan → build → done
-big feature → explore → plan → build → review → fix → done
+big feature → explore → plan → build --review → review → fix → build → done
 ```
 
 Decisions are captured inline throughout the flow using the `decide` skill, and terms are updated using `domain-modeling`.

exec
/usr/bin/bash -lc 'git status --short' in /home/daniel/build/knack
exec
/usr/bin/bash -lc 'git diff -- docs/loop.md docs/skills.md README.md AGENTS.md' in /home/daniel/build/knack
 succeeded in 0ms:
 M .agents/skills/fix/SKILL.md
 M .loop/review-loop/EVIDENCE.md
 M .loop/review-loop/HANDOFF.md
 M .loop/review-loop/QUEUE.md
 M AGENTS.md
 M README.md
 M cli/embedded/skills/fix/SKILL.md
 M docs/loop.md
 M docs/skills.md
 M loop.sh
 M tests/run.sh

 succeeded in 0ms:
diff --git a/AGENTS.md b/AGENTS.md
index 50f568d..f005f47 100644
--- a/AGENTS.md
+++ b/AGENTS.md
@@ -19,13 +19,14 @@ See `DESIGN.md` for the full design.
 
 ## Current state
 
-The loop (`loop.sh`), all seven skills, and the Go CLI are built. The loop supports per-unit `Agent:` overrides, `LOOP_AGENT_CMD` for agent-agnosticism, and writes `.loop/<name>/HANDOFF.md` on non-clean exits. The skills cover the full explore → plan → build → review → fix flow plus two shared skills (decide, domain-modeling). The CLI (`cli/`) provides `skills init|check`, `validate`, `decisions list|show|check`, `status`, `glossary check`, and `instructions`. It embeds the default skills via `go:embed` and scaffolds them into target projects.
+The loop (`loop.sh`), all seven skills, and the Go CLI are built. The loop supports per-unit `Agent:` overrides, `LOOP_AGENT_CMD` for agent-agnosticism, opt-in review/fix orchestration via `--review`, and `.loop/<name>/HANDOFF.md` on non-clean exits. The skills cover the full explore → plan → build → review → fix flow plus two shared skills (decide, domain-modeling). The CLI (`cli/`) provides `skills init|check`, `validate`, `decisions list|show|check`, `status`, `glossary check`, and `instructions`. It embeds the default skills via `go:embed` and scaffolds them into target projects.
 
 ## Core artifacts
 
 - `.loop/<name>/QUEUE.md` — disposable queue of verifiable work units.
 - `.loop/<name>/EVIDENCE.md` — append-only ledger of what each tick proved. Durable: keep it after deleting QUEUE.md so `decisions check` can still see which ADRs the completed cycle referenced.
 - `.loop/<name>/HANDOFF.md` — cross-session handoff (written on pause/stop).
+- `.loop/<name>/REVIEW.md` — structured review artifact written when `--review` is enabled. The loop reads only the actionable count.
 - `decisions/` — durable ADRs (architectural rulings, not current behavior).
 - `glossary.md` — durable ubiquitous language.
 - `.agents/skills/` — procedural knowledge (explore, plan, build, review, fix, decide, domain-modeling).
@@ -48,11 +49,13 @@ The loop (`loop.sh`), all seven skills, and the Go CLI are built. The loop suppo
 - The worker never self-certifies. The runner owns verification.
 - A `Verify` command must be deterministic and executable by the runner — not an LLM-as-judge. The loop's backpressure is mechanical.
 - The runner enforces hard stops (max ticks, no-progress detection). The agent does one work unit per tick and reports what remains if it can't finish.
-- Review is opt-in. The loop does not run review automatically; the user invokes it when they want adversarial scrutiny.
+- Review is opt-in. The loop runs review/fix only when invoked with `--review`; otherwise it runs build ticks only.
+- With `--review`, the loop invokes review after pending build units drain, reads only `REVIEW.md`'s actionable count, invokes fix when actionable findings exist, then runs another build pass for appended units. Judgment stays in the `review` and `fix` skills.
 - When the queue is complete and verified, the cycle's `.loop/<name>/QUEUE.md` is disposable. The human deletes it; the loop does not. `EVIDENCE.md` is the durable ledger — keep it so `decisions check` can still trace which ADRs the cycle referenced.
 - A work unit must leave the repo better if the loop stops immediately after it.
 - Work units are whatever shape the work is — not forced into "vertical slices."
 - `LOOP_AGENT_CMD` overrides the worker invocation (agent-agnostic: pi, claude, codex, opencode, etc.).
+- `LOOP_REVIEW_CMD` and `LOOP_FIX_CMD` override review/fix phase invocations when `--review` is enabled. They default to `LOOP_AGENT_CMD`.
 - Per-unit `Agent:` field in QUEUE.md overrides `LOOP_AGENT_CMD` for one unit (model routing).
 - Work units are `## <outcome>` headers with `Read first:`, `Constraints:`, `Done means:`, and `Verify:` fields. `Done means:` is the acceptance criteria; `Verify:` is the mechanically enforceable subset. The gap between them is the review surface.
 - `Read first:` is context, not scope — 2–4 entries of ADRs, code areas, or rulings.
diff --git a/README.md b/README.md
index 1300a1b..c1541df 100644
--- a/README.md
+++ b/README.md
@@ -43,6 +43,9 @@ cd /path/to/new-project
 6. On failure: the unit is marked `verify_failed` or `no_progress`, evidence is appended, and the loop stops or retries once.
 7. The loop halts on: max ticks reached with pending work, two no-progress strikes, or a failed verify after a real change.
 8. On any non-clean exit, the runner writes `HANDOFF.md` with completed/in-progress/remaining units and the next action.
+9. If `--review` is set, the runner invokes review after the queue drains, reads the actionable count from `REVIEW.md`, invokes fix when actionable findings exist, and runs another build pass for appended units.
+
+Review is opt-in. Without `--review`, the loop runs build ticks only; with it, review/fix are loop-orchestrated but still skill-owned.
 
 ## Queue format
 
@@ -70,6 +73,15 @@ Per-unit override via the `Agent:` field in a work unit:
 Agent: pi -p --no-session --approve --model glm-5.2 "$(cat "$LOOP_PROMPT_FILE")"
 ```
 
+Opt into review/fix orchestration with `--review`:
+
+```bash
+LOOP_AGENT_CMD='codex exec --dangerously-bypass-approvals-and-sandbox --ephemeral "$(cat "$LOOP_PROMPT_FILE")"' \
+LOOP_REVIEW_CMD='claude --print --no-session-persistence --dangerously-skip-permissions "$(cat "$LOOP_PROMPT_FILE")"' \
+LOOP_FIX_CMD='codex exec --dangerously-bypass-approvals-and-sandbox --ephemeral "$(cat "$LOOP_PROMPT_FILE")"' \
+./loop.sh run .loop/<name>/QUEUE.md --review --max-review-rounds 2
+```
+
 ## CLI
 
 The `knack` binary is a read-only validator and context provider. Build it from `cli/`:
diff --git a/docs/loop.md b/docs/loop.md
index a5cd2cd..ef47822 100644
--- a/docs/loop.md
+++ b/docs/loop.md
@@ -5,18 +5,24 @@
 ## Usage
 
 ```bash
-./loop.sh run <queue> [--repo DIR] [--max-ticks N] [--dry-run]
+./loop.sh run <queue> [--repo DIR] [--max-ticks N] [--review] [--max-review-rounds N] [--dry-run]
 ```
 
 - `queue` — path to `QUEUE.md`. Usually `.loop/<name>/QUEUE.md`.
 - `--repo DIR` — the working directory for the worker and the verify command. Defaults to the parent of the `.loop/<name>` directory that contains the queue.
 - `--max-ticks N` — maximum units to attempt. Default is `3`.
+- `--review` — opt into the bounded review/fix subloop after pending build units drain. Default behavior does not run review.
+- `--max-review-rounds N` — maximum review/fix rounds when `--review` is enabled. Default is `2`.
 - `--dry-run` — parse the first pending unit and print its title, repo, and verify command, then exit.
 
 ## Environment variables
 
 - `LOOP_AGENT_CMD` — optional command used to invoke the worker. If unset, the loop defaults to `pi -p --no-session --approve` with the prompt text as a single argument. If set, the command is evaluated by `bash -lc` in the repo directory, and `LOOP_PROMPT_FILE` is set to a temporary file containing `prompts/worker.md` plus the current work unit.
 - `LOOP_PROMPT_FILE` — set by the loop when `LOOP_AGENT_CMD` is used. Points to the generated prompt file. Do not override unless you are calling the worker manually.
+- `LOOP_REVIEW_CMD` — optional command used for the review phase when `--review` is enabled. Defaults to `LOOP_AGENT_CMD`, or to the loop's default `pi` invocation if neither is set.
+- `LOOP_FIX_CMD` — optional command used for the fix phase when `--review` is enabled. Defaults to `LOOP_AGENT_CMD`, or to the loop's default `pi` invocation if neither is set.
+
+During every phase, the loop also sets `LOOP_PHASE`, `LOOP_QUEUE_FILE`, `LOOP_EVIDENCE_FILE`, and `LOOP_REVIEW_FILE` for the worker process.
 
 ## Per-unit agent override
 
@@ -43,6 +49,20 @@ Agent: claude --print --no-session-persistence --dangerously-skip-permissions "$
    - If the repo changed, mark `verify_failed` and stop.
 9. If max ticks is reached with pending work, stop.
 
+## Optional review/fix subloop
+
+By default, `loop.sh` stops when the build queue has no pending work. With `--review`, it then runs a bounded review/fix subloop:
+
+1. Invoke a review worker with the review prompt and the completed queue/evidence.
+2. Require the review worker to write `.loop/<name>/REVIEW.md`.
+3. Read only the `- actionable: N` summary line from `REVIEW.md`.
+4. If `N` is `0`, stop cleanly.
+5. If `N` is non-zero, invoke a fix worker.
+6. Require the fix worker to append new `Status: pending` work units to the same `QUEUE.md`.
+7. Run the build pass again, then review again.
+
+The subloop stops when review is clean, `--max-review-rounds` is reached, `--max-ticks` is exhausted, or fix produces no pending work. The loop owns orchestration and stop conditions only; the `review` and `fix` skills own judgment and work-unit generation.
+
 ## Work unit statuses
 
 - `pending` — not yet started.
@@ -58,6 +78,7 @@ Agent: claude --print --no-session-persistence --dangerously-skip-permissions "$
 
 - `EVIDENCE.md` — append-only ledger. Includes the full unit, changed files, verify command, verify output, and worker output. It is durable; keep it after deleting `QUEUE.md` so `knack decisions check` can still see which ADRs the cycle referenced.
 - `HANDOFF.md` — written on non-clean exit. Sections: completed, in progress, remaining, next action. Delete when the work resumes.
+- `REVIEW.md` — structured review artifact written by the review worker when `--review` is enabled. The loop reads only its `- actionable: N` summary line.
 
 ## Agent invocation examples
 
@@ -71,8 +92,17 @@ LOOP_AGENT_CMD='devin --print --prompt-file "$LOOP_PROMPT_FILE" --permission-mod
 
 The command is passed to `bash -lc` in the repo directory, so the typical pattern is `"$(cat "$LOOP_PROMPT_FILE")"`.
 
+Review and fix can use separate agents:
+
+```bash
+LOOP_AGENT_CMD='codex exec --dangerously-bypass-approvals-and-sandbox --ephemeral "$(cat "$LOOP_PROMPT_FILE")"' \
+LOOP_REVIEW_CMD='claude --print --no-session-persistence --dangerously-skip-permissions "$(cat "$LOOP_PROMPT_FILE")"' \
+LOOP_FIX_CMD='codex exec --dangerously-bypass-approvals-and-sandbox --ephemeral "$(cat "$LOOP_PROMPT_FILE")"' \
+./loop.sh run .loop/<name>/QUEUE.md --review
+```
+
 ## Verification notes
 
 - `loop.sh` does **not** validate `QUEUE.md` structure. Use `knack validate` first.
-- `loop.sh` does **not** run review or manage ADRs/glossary. Those are handled by the skills.
+- `loop.sh` runs review and fix only when `--review` is set. It invokes the skills and reads the actionable count from `REVIEW.md`; it does not judge findings or manage ADRs/glossary.
 - The `Verify:` command must be deterministic and executable by the runner — tests, builds, type checks, not an LLM-as-judge.
diff --git a/docs/skills.md b/docs/skills.md
index fb1ec98..476d494 100644
--- a/docs/skills.md
+++ b/docs/skills.md
@@ -37,7 +37,11 @@ The `name` must match the directory name. The `description` is the trigger text
 
 ## How the loop uses skills
 
-`loop.sh` does not read skills. It prepends `prompts/worker.md` to the current work unit and runs the worker. `prompts/worker.md` tells the worker to load the `build` skill by name and path (e.g. "Load and follow the **build** skill in `.agents/skills/build/`"). For other phases (explore, plan, review, fix), the human or the agent invokes the skill directly.
+`loop.sh` does not read skills. It prepends `prompts/worker.md` to the current work unit and runs the worker. `prompts/worker.md` tells the worker to load the `build` skill by name and path (e.g. "Load and follow the **build** skill in `.agents/skills/build/`").
+
+When `--review` is set, the loop also invokes review and fix workers after the build queue drains. Those prompts tell the worker to load the `review` or `fix` skill directly. The loop orchestrates the bounded review/fix subloop, reads the actionable count from `REVIEW.md`, and runs another build pass when fix appends pending units. The skills still own judgment: review decides what the findings are, and fix decides which findings become work units.
+
+Without `--review`, review and fix remain manual skill invocations.
 
 ## Customizing skills
 
@@ -53,7 +57,7 @@ Skills are not a rigid gate. The default flow is `explore → plan → build →
 ```text
 small fix → plan → build → done
 bug report → explore → plan → build → done
-big feature → explore → plan → build → review → fix → done
+big feature → explore → plan → build --review → review → fix → build → done
 ```
 
 Decisions are captured inline throughout the flow using the `decide` skill, and terms are updated using `domain-modeling`.

diff --git a/AGENTS.md b/AGENTS.md
index 50f568dd4d4d4e617b4be191116bb8fea556720d..f005f47ed72aa4f97f47324d98c00d8c79b12921
--- a/AGENTS.md
+++ b/AGENTS.md
@@ -19,13 +19,14 @@
 
 ## Current state
 
-The loop (`loop.sh`), all seven skills, and the Go CLI are built. The loop supports per-unit `Agent:` overrides, `LOOP_AGENT_CMD` for agent-agnosticism, and writes `.loop/<name>/HANDOFF.md` on non-clean exits. The skills cover the full explore → plan → build → review → fix flow plus two shared skills (decide, domain-modeling). The CLI (`cli/`) provides `skills init|check`, `validate`, `decisions list|show|check`, `status`, `glossary check`, and `instructions`. It embeds the default skills via `go:embed` and scaffolds them into target projects.
+The loop (`loop.sh`), all seven skills, and the Go CLI are built. The loop supports per-unit `Agent:` overrides, `LOOP_AGENT_CMD` for agent-agnosticism, opt-in review/fix orchestration via `--review`, and `.loop/<name>/HANDOFF.md` on non-clean exits. The skills cover the full explore → plan → build → review → fix flow plus two shared skills (decide, domain-modeling). The CLI (`cli/`) provides `skills init|check`, `validate`, `decisions list|show|check`, `status`, `glossary check`, and `instructions`. It embeds the default skills via `go:embed` and scaffolds them into target projects.
 
 ## Core artifacts
 
 - `.loop/<name>/QUEUE.md` — disposable queue of verifiable work units.
 - `.loop/<name>/EVIDENCE.md` — append-only ledger of what each tick proved. Durable: keep it after deleting QUEUE.md so `decisions check` can still see which ADRs the completed cycle referenced.
 - `.loop/<name>/HANDOFF.md` — cross-session handoff (written on pause/stop).
+- `.loop/<name>/REVIEW.md` — structured review artifact written when `--review` is enabled. The loop reads only the actionable count.
 - `decisions/` — durable ADRs (architectural rulings, not current behavior).
 - `glossary.md` — durable ubiquitous language.
 - `.agents/skills/` — procedural knowledge (explore, plan, build, review, fix, decide, domain-modeling).
@@ -48,11 +49,13 @@
 - The worker never self-certifies. The runner owns verification.
 - A `Verify` command must be deterministic and executable by the runner — not an LLM-as-judge. The loop's backpressure is mechanical.
 - The runner enforces hard stops (max ticks, no-progress detection). The agent does one work unit per tick and reports what remains if it can't finish.
-- Review is opt-in. The loop does not run review automatically; the user invokes it when they want adversarial scrutiny.
+- Review is opt-in. The loop runs review/fix only when invoked with `--review`; otherwise it runs build ticks only.
+- With `--review`, the loop invokes review after pending build units drain, reads only `REVIEW.md`'s actionable count, invokes fix when actionable findings exist, then runs another build pass for appended units. Judgment stays in the `review` and `fix` skills.
 - When the queue is complete and verified, the cycle's `.loop/<name>/QUEUE.md` is disposable. The human deletes it; the loop does not. `EVIDENCE.md` is the durable ledger — keep it so `decisions check` can still trace which ADRs the cycle referenced.
 - A work unit must leave the repo better if the loop stops immediately after it.
 - Work units are whatever shape the work is — not forced into "vertical slices."
 - `LOOP_AGENT_CMD` overrides the worker invocation (agent-agnostic: pi, claude, codex, opencode, etc.).
+- `LOOP_REVIEW_CMD` and `LOOP_FIX_CMD` override review/fix phase invocations when `--review` is enabled. They default to `LOOP_AGENT_CMD`.
 - Per-unit `Agent:` field in QUEUE.md overrides `LOOP_AGENT_CMD` for one unit (model routing).
 - Work units are `## <outcome>` headers with `Read first:`, `Constraints:`, `Done means:`, and `Verify:` fields. `Done means:` is the acceptance criteria; `Verify:` is the mechanically enforceable subset. The gap between them is the review surface.
 - `Read first:` is context, not scope — 2–4 entries of ADRs, code areas, or rulings.
diff --git a/README.md b/README.md
index 1300a1bc12102f97cd89ace07333f179f8700cee..c1541df3127dce0ba5ba94a9cc4d8ba97c5a6b6a
--- a/README.md
+++ b/README.md
@@ -43,6 +43,9 @@
 6. On failure: the unit is marked `verify_failed` or `no_progress`, evidence is appended, and the loop stops or retries once.
 7. The loop halts on: max ticks reached with pending work, two no-progress strikes, or a failed verify after a real change.
 8. On any non-clean exit, the runner writes `HANDOFF.md` with completed/in-progress/remaining units and the next action.
+9. If `--review` is set, the runner invokes review after the queue drains, reads the actionable count from `REVIEW.md`, invokes fix when actionable findings exist, and runs another build pass for appended units.
+
+Review is opt-in. Without `--review`, the loop runs build ticks only; with it, review/fix are loop-orchestrated but still skill-owned.
 
 ## Queue format
 
@@ -70,6 +73,15 @@
 Agent: pi -p --no-session --approve --model glm-5.2 "$(cat "$LOOP_PROMPT_FILE")"
 ```
 
+Opt into review/fix orchestration with `--review`:
+
+```bash
+LOOP_AGENT_CMD='codex exec --dangerously-bypass-approvals-and-sandbox --ephemeral "$(cat "$LOOP_PROMPT_FILE")"' \
+LOOP_REVIEW_CMD='claude --print --no-session-persistence --dangerously-skip-permissions "$(cat "$LOOP_PROMPT_FILE")"' \
+LOOP_FIX_CMD='codex exec --dangerously-bypass-approvals-and-sandbox --ephemeral "$(cat "$LOOP_PROMPT_FILE")"' \
+./loop.sh run .loop/<name>/QUEUE.md --review --max-review-rounds 2
+```
+
 ## CLI
 
 The `knack` binary is a read-only validator and context provider. Build it from `cli/`:
diff --git a/docs/loop.md b/docs/loop.md
index a5cd2cd28883321ca006489e2c630fbdc8d6bb8f..ef47822f4007c3fbcb5ee094f39164166ca63646
--- a/docs/loop.md
+++ b/docs/loop.md
@@ -5,18 +5,24 @@
 ## Usage
 
 ```bash
-./loop.sh run <queue> [--repo DIR] [--max-ticks N] [--dry-run]
+./loop.sh run <queue> [--repo DIR] [--max-ticks N] [--review] [--max-review-rounds N] [--dry-run]
 ```
 
 - `queue` — path to `QUEUE.md`. Usually `.loop/<name>/QUEUE.md`.
 - `--repo DIR` — the working directory for the worker and the verify command. Defaults to the parent of the `.loop/<name>` directory that contains the queue.
 - `--max-ticks N` — maximum units to attempt. Default is `3`.
+- `--review` — opt into the bounded review/fix subloop after pending build units drain. Default behavior does not run review.
+- `--max-review-rounds N` — maximum review/fix rounds when `--review` is enabled. Default is `2`.
 - `--dry-run` — parse the first pending unit and print its title, repo, and verify command, then exit.
 
 ## Environment variables
 
 - `LOOP_AGENT_CMD` — optional command used to invoke the worker. If unset, the loop defaults to `pi -p --no-session --approve` with the prompt text as a single argument. If set, the command is evaluated by `bash -lc` in the repo directory, and `LOOP_PROMPT_FILE` is set to a temporary file containing `prompts/worker.md` plus the current work unit.
 - `LOOP_PROMPT_FILE` — set by the loop when `LOOP_AGENT_CMD` is used. Points to the generated prompt file. Do not override unless you are calling the worker manually.
+- `LOOP_REVIEW_CMD` — optional command used for the review phase when `--review` is enabled. Defaults to `LOOP_AGENT_CMD`, or to the loop's default `pi` invocation if neither is set.
+- `LOOP_FIX_CMD` — optional command used for the fix phase when `--review` is enabled. Defaults to `LOOP_AGENT_CMD`, or to the loop's default `pi` invocation if neither is set.
+
+During every phase, the loop also sets `LOOP_PHASE`, `LOOP_QUEUE_FILE`, `LOOP_EVIDENCE_FILE`, and `LOOP_REVIEW_FILE` for the worker process.
 
 ## Per-unit agent override
 
@@ -43,6 +49,20 @@
    - If the repo changed, mark `verify_failed` and stop.
 9. If max ticks is reached with pending work, stop.
 
+## Optional review/fix subloop
+
+By default, `loop.sh` stops when the build queue has no pending work. With `--review`, it then runs a bounded review/fix subloop:
+
+1. Invoke a review worker with the review prompt and the completed queue/evidence.
+2. Require the review worker to write `.loop/<name>/REVIEW.md`.
+3. Read only the `- actionable: N` summary line from `REVIEW.md`.
+4. If `N` is `0`, stop cleanly.
+5. If `N` is non-zero, invoke a fix worker.
+6. Require the fix worker to append new `Status: pending` work units to the same `QUEUE.md`.
+7. Run the build pass again, then review again.
+
+The subloop stops when review is clean, `--max-review-rounds` is reached, `--max-ticks` is exhausted, or fix produces no pending work. The loop owns orchestration and stop conditions only; the `review` and `fix` skills own judgment and work-unit generation.
+
 ## Work unit statuses
 
 - `pending` — not yet started.
@@ -58,6 +78,7 @@
 
 - `EVIDENCE.md` — append-only ledger. Includes the full unit, changed files, verify command, verify output, and worker output. It is durable; keep it after deleting `QUEUE.md` so `knack decisions check` can still see which ADRs the cycle referenced.
 - `HANDOFF.md` — written on non-clean exit. Sections: completed, in progress, remaining, next action. Delete when the work resumes.
+- `REVIEW.md` — structured review artifact written by the review worker when `--review` is enabled. The loop reads only its `- actionable: N` summary line.
 
 ## Agent invocation examples
 
@@ -71,8 +92,17 @@
 
 The command is passed to `bash -lc` in the repo directory, so the typical pattern is `"$(cat "$LOOP_PROMPT_FILE")"`.
 
+Review and fix can use separate agents:
+
+```bash
+LOOP_AGENT_CMD='codex exec --dangerously-bypass-approvals-and-sandbox --ephemeral "$(cat "$LOOP_PROMPT_FILE")"' \
+LOOP_REVIEW_CMD='claude --print --no-session-persistence --dangerously-skip-permissions "$(cat "$LOOP_PROMPT_FILE")"' \
+LOOP_FIX_CMD='codex exec --dangerously-bypass-approvals-and-sandbox --ephemeral "$(cat "$LOOP_PROMPT_FILE")"' \
+./loop.sh run .loop/<name>/QUEUE.md --review
+```
+
 ## Verification notes
 
 - `loop.sh` does **not** validate `QUEUE.md` structure. Use `knack validate` first.
-- `loop.sh` does **not** run review or manage ADRs/glossary. Those are handled by the skills.
+- `loop.sh` runs review and fix only when `--review` is set. It invokes the skills and reads the actionable count from `REVIEW.md`; it does not judge findings or manage ADRs/glossary.
 - The `Verify:` command must be deterministic and executable by the runner — tests, builds, type checks, not an LLM-as-judge.
diff --git a/docs/skills.md b/docs/skills.md
index fb1ec98119e1193859c0a1612f580d2da5afe95d..476d494e3c88962825e751b84151a0ff48ac2d41
--- a/docs/skills.md
+++ b/docs/skills.md
@@ -37,7 +37,11 @@
 
 ## How the loop uses skills
 
-`loop.sh` does not read skills. It prepends `prompts/worker.md` to the current work unit and runs the worker. `prompts/worker.md` tells the worker to load the `build` skill by name and path (e.g. "Load and follow the **build** skill in `.agents/skills/build/`"). For other phases (explore, plan, review, fix), the human or the agent invokes the skill directly.
+`loop.sh` does not read skills. It prepends `prompts/worker.md` to the current work unit and runs the worker. `prompts/worker.md` tells the worker to load the `build` skill by name and path (e.g. "Load and follow the **build** skill in `.agents/skills/build/`").
+
+When `--review` is set, the loop also invokes review and fix workers after the build queue drains. Those prompts tell the worker to load the `review` or `fix` skill directly. The loop orchestrates the bounded review/fix subloop, reads the actionable count from `REVIEW.md`, and runs another build pass when fix appends pending units. The skills still own judgment: review decides what the findings are, and fix decides which findings become work units.
+
+Without `--review`, review and fix remain manual skill invocations.
 
 ## Customizing skills
 
@@ -53,7 +57,7 @@
 ```text
 small fix → plan → build → done
 bug report → explore → plan → build → done
-big feature → explore → plan → build → review → fix → done
+big feature → explore → plan → build --review → review → fix → build → done
 ```
 
 Decisions are captured inline throughout the flow using the `decide` skill, and terms are updated using `domain-modeling`.

codex
The diff for this tick is limited to AGENTS.md, README.md, and the two docs. There are existing modified loop/skill files in the worktree from earlier work; I’m leaving them alone and only checking whether the embedded skills are already in sync.
exec
/usr/bin/bash -lc 'diff -r .agents/skills cli/embedded/skills' in /home/daniel/build/knack
 succeeded in 0ms:
exec
/usr/bin/bash -lc "rg -n \"does \\*\\*not\\*\\* run review|does not run review|human or the agent invokes|--review|max-review-rounds|LOOP_REVIEW_CMD|LOOP_FIX_CMD|REVIEW.md\" docs/loop.md docs/skills.md README.md AGENTS.md" in /home/daniel/build/knack
 succeeded in 0ms:
AGENTS.md:22:The loop (`loop.sh`), all seven skills, and the Go CLI are built. The loop supports per-unit `Agent:` overrides, `LOOP_AGENT_CMD` for agent-agnosticism, opt-in review/fix orchestration via `--review`, and `.loop/<name>/HANDOFF.md` on non-clean exits. The skills cover the full explore → plan → build → review → fix flow plus two shared skills (decide, domain-modeling). The CLI (`cli/`) provides `skills init|check`, `validate`, `decisions list|show|check`, `status`, `glossary check`, and `instructions`. It embeds the default skills via `go:embed` and scaffolds them into target projects.
AGENTS.md:29:- `.loop/<name>/REVIEW.md` — structured review artifact written when `--review` is enabled. The loop reads only the actionable count.
AGENTS.md:52:- Review is opt-in. The loop runs review/fix only when invoked with `--review`; otherwise it runs build ticks only.
AGENTS.md:53:- With `--review`, the loop invokes review after pending build units drain, reads only `REVIEW.md`'s actionable count, invokes fix when actionable findings exist, then runs another build pass for appended units. Judgment stays in the `review` and `fix` skills.
AGENTS.md:58:- `LOOP_REVIEW_CMD` and `LOOP_FIX_CMD` override review/fix phase invocations when `--review` is enabled. They default to `LOOP_AGENT_CMD`.
README.md:46:9. If `--review` is set, the runner invokes review after the queue drains, reads the actionable count from `REVIEW.md`, invokes fix when actionable findings exist, and runs another build pass for appended units.
README.md:48:Review is opt-in. Without `--review`, the loop runs build ticks only; with it, review/fix are loop-orchestrated but still skill-owned.
README.md:76:Opt into review/fix orchestration with `--review`:
README.md:80:LOOP_REVIEW_CMD='claude --print --no-session-persistence --dangerously-skip-permissions "$(cat "$LOOP_PROMPT_FILE")"' \
README.md:81:LOOP_FIX_CMD='codex exec --dangerously-bypass-approvals-and-sandbox --ephemeral "$(cat "$LOOP_PROMPT_FILE")"' \
README.md:82:./loop.sh run .loop/<name>/QUEUE.md --review --max-review-rounds 2
docs/skills.md:42:When `--review` is set, the loop also invokes review and fix workers after the build queue drains. Those prompts tell the worker to load the `review` or `fix` skill directly. The loop orchestrates the bounded review/fix subloop, reads the actionable count from `REVIEW.md`, and runs another build pass when fix appends pending units. The skills still own judgment: review decides what the findings are, and fix decides which findings become work units.
docs/skills.md:44:Without `--review`, review and fix remain manual skill invocations.
docs/skills.md:60:big feature → explore → plan → build --review → review → fix → build → done
docs/loop.md:8:./loop.sh run <queue> [--repo DIR] [--max-ticks N] [--review] [--max-review-rounds N] [--dry-run]
docs/loop.md:14:- `--review` — opt into the bounded review/fix subloop after pending build units drain. Default behavior does not run review.
docs/loop.md:15:- `--max-review-rounds N` — maximum review/fix rounds when `--review` is enabled. Default is `2`.
docs/loop.md:22:- `LOOP_REVIEW_CMD` — optional command used for the review phase when `--review` is enabled. Defaults to `LOOP_AGENT_CMD`, or to the loop's default `pi` invocation if neither is set.
docs/loop.md:23:- `LOOP_FIX_CMD` — optional command used for the fix phase when `--review` is enabled. Defaults to `LOOP_AGENT_CMD`, or to the loop's default `pi` invocation if neither is set.
docs/loop.md:54:By default, `loop.sh` stops when the build queue has no pending work. With `--review`, it then runs a bounded review/fix subloop:
docs/loop.md:57:2. Require the review worker to write `.loop/<name>/REVIEW.md`.
docs/loop.md:58:3. Read only the `- actionable: N` summary line from `REVIEW.md`.
docs/loop.md:64:The subloop stops when review is clean, `--max-review-rounds` is reached, `--max-ticks` is exhausted, or fix produces no pending work. The loop owns orchestration and stop conditions only; the `review` and `fix` skills own judgment and work-unit generation.
docs/loop.md:81:- `REVIEW.md` — structured review artifact written by the review worker when `--review` is enabled. The loop reads only its `- actionable: N` summary line.
docs/loop.md:99:LOOP_REVIEW_CMD='claude --print --no-session-persistence --dangerously-skip-permissions "$(cat "$LOOP_PROMPT_FILE")"' \
docs/loop.md:100:LOOP_FIX_CMD='codex exec --dangerously-bypass-approvals-and-sandbox --ephemeral "$(cat "$LOOP_PROMPT_FILE")"' \
docs/loop.md:101:./loop.sh run .loop/<name>/QUEUE.md --review
docs/loop.md:107:- `loop.sh` runs review and fix only when `--review` is set. It invokes the skills and reads the actionable count from `REVIEW.md`; it does not judge findings or manage ADRs/glossary.

diff --git a/AGENTS.md b/AGENTS.md
index 50f568dd4d4d4e617b4be191116bb8fea556720d..f005f47ed72aa4f97f47324d98c00d8c79b12921
--- a/AGENTS.md
+++ b/AGENTS.md
@@ -19,13 +19,14 @@
 
 ## Current state
 
-The loop (`loop.sh`), all seven skills, and the Go CLI are built. The loop supports per-unit `Agent:` overrides, `LOOP_AGENT_CMD` for agent-agnosticism, and writes `.loop/<name>/HANDOFF.md` on non-clean exits. The skills cover the full explore → plan → build → review → fix flow plus two shared skills (decide, domain-modeling). The CLI (`cli/`) provides `skills init|check`, `validate`, `decisions list|show|check`, `status`, `glossary check`, and `instructions`. It embeds the default skills via `go:embed` and scaffolds them into target projects.
+The loop (`loop.sh`), all seven skills, and the Go CLI are built. The loop supports per-unit `Agent:` overrides, `LOOP_AGENT_CMD` for agent-agnosticism, opt-in review/fix orchestration via `--review`, and `.loop/<name>/HANDOFF.md` on non-clean exits. The skills cover the full explore → plan → build → review → fix flow plus two shared skills (decide, domain-modeling). The CLI (`cli/`) provides `skills init|check`, `validate`, `decisions list|show|check`, `status`, `glossary check`, and `instructions`. It embeds the default skills via `go:embed` and scaffolds them into target projects.
 
 ## Core artifacts
 
 - `.loop/<name>/QUEUE.md` — disposable queue of verifiable work units.
 - `.loop/<name>/EVIDENCE.md` — append-only ledger of what each tick proved. Durable: keep it after deleting QUEUE.md so `decisions check` can still see which ADRs the completed cycle referenced.
 - `.loop/<name>/HANDOFF.md` — cross-session handoff (written on pause/stop).
+- `.loop/<name>/REVIEW.md` — structured review artifact written when `--review` is enabled. The loop reads only the actionable count.
 - `decisions/` — durable ADRs (architectural rulings, not current behavior).
 - `glossary.md` — durable ubiquitous language.
 - `.agents/skills/` — procedural knowledge (explore, plan, build, review, fix, decide, domain-modeling).
@@ -48,11 +49,13 @@
 - The worker never self-certifies. The runner owns verification.
 - A `Verify` command must be deterministic and executable by the runner — not an LLM-as-judge. The loop's backpressure is mechanical.
 - The runner enforces hard stops (max ticks, no-progress detection). The agent does one work unit per tick and reports what remains if it can't finish.
-- Review is opt-in. The loop does not run review automatically; the user invokes it when they want adversarial scrutiny.
+- Review is opt-in. The loop runs review/fix only when invoked with `--review`; otherwise it runs build ticks only.
+- With `--review`, the loop invokes review after pending build units drain, reads only `REVIEW.md`'s actionable count, invokes fix when actionable findings exist, then runs another build pass for appended units. Judgment stays in the `review` and `fix` skills.
 - When the queue is complete and verified, the cycle's `.loop/<name>/QUEUE.md` is disposable. The human deletes it; the loop does not. `EVIDENCE.md` is the durable ledger — keep it so `decisions check` can still trace which ADRs the cycle referenced.
 - A work unit must leave the repo better if the loop stops immediately after it.
 - Work units are whatever shape the work is — not forced into "vertical slices."
 - `LOOP_AGENT_CMD` overrides the worker invocation (agent-agnostic: pi, claude, codex, opencode, etc.).
+- `LOOP_REVIEW_CMD` and `LOOP_FIX_CMD` override review/fix phase invocations when `--review` is enabled. They default to `LOOP_AGENT_CMD`.
 - Per-unit `Agent:` field in QUEUE.md overrides `LOOP_AGENT_CMD` for one unit (model routing).
 - Work units are `## <outcome>` headers with `Read first:`, `Constraints:`, `Done means:`, and `Verify:` fields. `Done means:` is the acceptance criteria; `Verify:` is the mechanically enforceable subset. The gap between them is the review surface.
 - `Read first:` is context, not scope — 2–4 entries of ADRs, code areas, or rulings.
diff --git a/README.md b/README.md
index 1300a1bc12102f97cd89ace07333f179f8700cee..c1541df3127dce0ba5ba94a9cc4d8ba97c5a6b6a
--- a/README.md
+++ b/README.md
@@ -43,6 +43,9 @@
 6. On failure: the unit is marked `verify_failed` or `no_progress`, evidence is appended, and the loop stops or retries once.
 7. The loop halts on: max ticks reached with pending work, two no-progress strikes, or a failed verify after a real change.
 8. On any non-clean exit, the runner writes `HANDOFF.md` with completed/in-progress/remaining units and the next action.
+9. If `--review` is set, the runner invokes review after the queue drains, reads the actionable count from `REVIEW.md`, invokes fix when actionable findings exist, and runs another build pass for appended units.
+
+Review is opt-in. Without `--review`, the loop runs build ticks only; with it, review/fix are loop-orchestrated but still skill-owned.
 
 ## Queue format
 
@@ -70,6 +73,15 @@
 Agent: pi -p --no-session --approve --model glm-5.2 "$(cat "$LOOP_PROMPT_FILE")"
 ```
 
+Opt into review/fix orchestration with `--review`:
+
+```bash
+LOOP_AGENT_CMD='codex exec --dangerously-bypass-approvals-and-sandbox --ephemeral "$(cat "$LOOP_PROMPT_FILE")"' \
+LOOP_REVIEW_CMD='claude --print --no-session-persistence --dangerously-skip-permissions "$(cat "$LOOP_PROMPT_FILE")"' \
+LOOP_FIX_CMD='codex exec --dangerously-bypass-approvals-and-sandbox --ephemeral "$(cat "$LOOP_PROMPT_FILE")"' \
+./loop.sh run .loop/<name>/QUEUE.md --review --max-review-rounds 2
+```
+
 ## CLI
 
 The `knack` binary is a read-only validator and context provider. Build it from `cli/`:
diff --git a/docs/loop.md b/docs/loop.md
index a5cd2cd28883321ca006489e2c630fbdc8d6bb8f..ef47822f4007c3fbcb5ee094f39164166ca63646
--- a/docs/loop.md
+++ b/docs/loop.md
@@ -5,18 +5,24 @@
 ## Usage
 
 ```bash
-./loop.sh run <queue> [--repo DIR] [--max-ticks N] [--dry-run]
+./loop.sh run <queue> [--repo DIR] [--max-ticks N] [--review] [--max-review-rounds N] [--dry-run]
 ```
 
 - `queue` — path to `QUEUE.md`. Usually `.loop/<name>/QUEUE.md`.
 - `--repo DIR` — the working directory for the worker and the verify command. Defaults to the parent of the `.loop/<name>` directory that contains the queue.
 - `--max-ticks N` — maximum units to attempt. Default is `3`.
+- `--review` — opt into the bounded review/fix subloop after pending build units drain. Default behavior does not run review.
+- `--max-review-rounds N` — maximum review/fix rounds when `--review` is enabled. Default is `2`.
 - `--dry-run` — parse the first pending unit and print its title, repo, and verify command, then exit.
 
 ## Environment variables
 
 - `LOOP_AGENT_CMD` — optional command used to invoke the worker. If unset, the loop defaults to `pi -p --no-session --approve` with the prompt text as a single argument. If set, the command is evaluated by `bash -lc` in the repo directory, and `LOOP_PROMPT_FILE` is set to a temporary file containing `prompts/worker.md` plus the current work unit.
 - `LOOP_PROMPT_FILE` — set by the loop when `LOOP_AGENT_CMD` is used. Points to the generated prompt file. Do not override unless you are calling the worker manually.
+- `LOOP_REVIEW_CMD` — optional command used for the review phase when `--review` is enabled. Defaults to `LOOP_AGENT_CMD`, or to the loop's default `pi` invocation if neither is set.
+- `LOOP_FIX_CMD` — optional command used for the fix phase when `--review` is enabled. Defaults to `LOOP_AGENT_CMD`, or to the loop's default `pi` invocation if neither is set.
+
+During every phase, the loop also sets `LOOP_PHASE`, `LOOP_QUEUE_FILE`, `LOOP_EVIDENCE_FILE`, and `LOOP_REVIEW_FILE` for the worker process.
 
 ## Per-unit agent override
 
@@ -43,6 +49,20 @@
    - If the repo changed, mark `verify_failed` and stop.
 9. If max ticks is reached with pending work, stop.
 
+## Optional review/fix subloop
+
+By default, `loop.sh` stops when the build queue has no pending work. With `--review`, it then runs a bounded review/fix subloop:
+
+1. Invoke a review worker with the review prompt and the completed queue/evidence.
+2. Require the review worker to write `.loop/<name>/REVIEW.md`.
+3. Read only the `- actionable: N` summary line from `REVIEW.md`.
+4. If `N` is `0`, stop cleanly.
+5. If `N` is non-zero, invoke a fix worker.
+6. Require the fix worker to append new `Status: pending` work units to the same `QUEUE.md`.
+7. Run the build pass again, then review again.
+
+The subloop stops when review is clean, `--max-review-rounds` is reached, `--max-ticks` is exhausted, or fix produces no pending work. The loop owns orchestration and stop conditions only; the `review` and `fix` skills own judgment and work-unit generation.
+
 ## Work unit statuses
 
 - `pending` — not yet started.
@@ -58,6 +78,7 @@
 
 - `EVIDENCE.md` — append-only ledger. Includes the full unit, changed files, verify command, verify output, and worker output. It is durable; keep it after deleting `QUEUE.md` so `knack decisions check` can still see which ADRs the cycle referenced.
 - `HANDOFF.md` — written on non-clean exit. Sections: completed, in progress, remaining, next action. Delete when the work resumes.
+- `REVIEW.md` — structured review artifact written by the review worker when `--review` is enabled. The loop reads only its `- actionable: N` summary line.
 
 ## Agent invocation examples
 
@@ -71,8 +92,17 @@
 
 The command is passed to `bash -lc` in the repo directory, so the typical pattern is `"$(cat "$LOOP_PROMPT_FILE")"`.
 
+Review and fix can use separate agents:
+
+```bash
+LOOP_AGENT_CMD='codex exec --dangerously-bypass-approvals-and-sandbox --ephemeral "$(cat "$LOOP_PROMPT_FILE")"' \
+LOOP_REVIEW_CMD='claude --print --no-session-persistence --dangerously-skip-permissions "$(cat "$LOOP_PROMPT_FILE")"' \
+LOOP_FIX_CMD='codex exec --dangerously-bypass-approvals-and-sandbox --ephemeral "$(cat "$LOOP_PROMPT_FILE")"' \
+./loop.sh run .loop/<name>/QUEUE.md --review
+```
+
 ## Verification notes
 
 - `loop.sh` does **not** validate `QUEUE.md` structure. Use `knack validate` first.
-- `loop.sh` does **not** run review or manage ADRs/glossary. Those are handled by the skills.
+- `loop.sh` runs review and fix only when `--review` is set. It invokes the skills and reads the actionable count from `REVIEW.md`; it does not judge findings or manage ADRs/glossary.
 - The `Verify:` command must be deterministic and executable by the runner — tests, builds, type checks, not an LLM-as-judge.
diff --git a/docs/skills.md b/docs/skills.md
index fb1ec98119e1193859c0a1612f580d2da5afe95d..476d494e3c88962825e751b84151a0ff48ac2d41
--- a/docs/skills.md
+++ b/docs/skills.md
@@ -37,7 +37,11 @@
 
 ## How the loop uses skills
 
-`loop.sh` does not read skills. It prepends `prompts/worker.md` to the current work unit and runs the worker. `prompts/worker.md` tells the worker to load the `build` skill by name and path (e.g. "Load and follow the **build** skill in `.agents/skills/build/`"). For other phases (explore, plan, review, fix), the human or the agent invokes the skill directly.
+`loop.sh` does not read skills. It prepends `prompts/worker.md` to the current work unit and runs the worker. `prompts/worker.md` tells the worker to load the `build` skill by name and path (e.g. "Load and follow the **build** skill in `.agents/skills/build/`").
+
+When `--review` is set, the loop also invokes review and fix workers after the build queue drains. Those prompts tell the worker to load the `review` or `fix` skill directly. The loop orchestrates the bounded review/fix subloop, reads the actionable count from `REVIEW.md`, and runs another build pass when fix appends pending units. The skills still own judgment: review decides what the findings are, and fix decides which findings become work units.
+
+Without `--review`, review and fix remain manual skill invocations.
 
 ## Customizing skills
 
@@ -53,7 +57,7 @@
 ```text
 small fix → plan → build → done
 bug report → explore → plan → build → done
-big feature → explore → plan → build → review → fix → done
+big feature → explore → plan → build --review → review → fix → build → done
 ```
 
 Decisions are captured inline throughout the flow using the `decide` skill, and terms are updated using `domain-modeling`.

codex
Embedded skills are already synchronized (`diff -r .agents/skills cli/embedded/skills` produced no output). I’m running the unit’s verify command now.
exec
/usr/bin/bash -lc './tests/run.sh && cd cli && go test ./...' in /home/daniel/build/knack
diff --git a/AGENTS.md b/AGENTS.md
index 50f568dd4d4d4e617b4be191116bb8fea556720d..f005f47ed72aa4f97f47324d98c00d8c79b12921
--- a/AGENTS.md
+++ b/AGENTS.md
@@ -19,13 +19,14 @@
 
 ## Current state
 
-The loop (`loop.sh`), all seven skills, and the Go CLI are built. The loop supports per-unit `Agent:` overrides, `LOOP_AGENT_CMD` for agent-agnosticism, and writes `.loop/<name>/HANDOFF.md` on non-clean exits. The skills cover the full explore → plan → build → review → fix flow plus two shared skills (decide, domain-modeling). The CLI (`cli/`) provides `skills init|check`, `validate`, `decisions list|show|check`, `status`, `glossary check`, and `instructions`. It embeds the default skills via `go:embed` and scaffolds them into target projects.
+The loop (`loop.sh`), all seven skills, and the Go CLI are built. The loop supports per-unit `Agent:` overrides, `LOOP_AGENT_CMD` for agent-agnosticism, opt-in review/fix orchestration via `--review`, and `.loop/<name>/HANDOFF.md` on non-clean exits. The skills cover the full explore → plan → build → review → fix flow plus two shared skills (decide, domain-modeling). The CLI (`cli/`) provides `skills init|check`, `validate`, `decisions list|show|check`, `status`, `glossary check`, and `instructions`. It embeds the default skills via `go:embed` and scaffolds them into target projects.
 
 ## Core artifacts
 
 - `.loop/<name>/QUEUE.md` — disposable queue of verifiable work units.
 - `.loop/<name>/EVIDENCE.md` — append-only ledger of what each tick proved. Durable: keep it after deleting QUEUE.md so `decisions check` can still see which ADRs the completed cycle referenced.
 - `.loop/<name>/HANDOFF.md` — cross-session handoff (written on pause/stop).
+- `.loop/<name>/REVIEW.md` — structured review artifact written when `--review` is enabled. The loop reads only the actionable count.
 - `decisions/` — durable ADRs (architectural rulings, not current behavior).
 - `glossary.md` — durable ubiquitous language.
 - `.agents/skills/` — procedural knowledge (explore, plan, build, review, fix, decide, domain-modeling).
@@ -48,11 +49,13 @@
 - The worker never self-certifies. The runner owns verification.
 - A `Verify` command must be deterministic and executable by the runner — not an LLM-as-judge. The loop's backpressure is mechanical.
 - The runner enforces hard stops (max ticks, no-progress detection). The agent does one work unit per tick and reports what remains if it can't finish.
-- Review is opt-in. The loop does not run review automatically; the user invokes it when they want adversarial scrutiny.
+- Review is opt-in. The loop runs review/fix only when invoked with `--review`; otherwise it runs build ticks only.
+- With `--review`, the loop invokes review after pending build units drain, reads only `REVIEW.md`'s actionable count, invokes fix when actionable findings exist, then runs another build pass for appended units. Judgment stays in the `review` and `fix` skills.
 - When the queue is complete and verified, the cycle's `.loop/<name>/QUEUE.md` is disposable. The human deletes it; the loop does not. `EVIDENCE.md` is the durable ledger — keep it so `decisions check` can still trace which ADRs the cycle referenced.
 - A work unit must leave the repo better if the loop stops immediately after it.
 - Work units are whatever shape the work is — not forced into "vertical slices."
 - `LOOP_AGENT_CMD` overrides the worker invocation (agent-agnostic: pi, claude, codex, opencode, etc.).
+- `LOOP_REVIEW_CMD` and `LOOP_FIX_CMD` override review/fix phase invocations when `--review` is enabled. They default to `LOOP_AGENT_CMD`.
 - Per-unit `Agent:` field in QUEUE.md overrides `LOOP_AGENT_CMD` for one unit (model routing).
 - Work units are `## <outcome>` headers with `Read first:`, `Constraints:`, `Done means:`, and `Verify:` fields. `Done means:` is the acceptance criteria; `Verify:` is the mechanically enforceable subset. The gap between them is the review surface.
 - `Read first:` is context, not scope — 2–4 entries of ADRs, code areas, or rulings.
diff --git a/README.md b/README.md
index 1300a1bc12102f97cd89ace07333f179f8700cee..c1541df3127dce0ba5ba94a9cc4d8ba97c5a6b6a
--- a/README.md
+++ b/README.md
@@ -43,6 +43,9 @@
 6. On failure: the unit is marked `verify_failed` or `no_progress`, evidence is appended, and the loop stops or retries once.
 7. The loop halts on: max ticks reached with pending work, two no-progress strikes, or a failed verify after a real change.
 8. On any non-clean exit, the runner writes `HANDOFF.md` with completed/in-progress/remaining units and the next action.
+9. If `--review` is set, the runner invokes review after the queue drains, reads the actionable count from `REVIEW.md`, invokes fix when actionable findings exist, and runs another build pass for appended units.
+
+Review is opt-in. Without `--review`, the loop runs build ticks only; with it, review/fix are loop-orchestrated but still skill-owned.
 
 ## Queue format
 
@@ -70,6 +73,15 @@
 Agent: pi -p --no-session --approve --model glm-5.2 "$(cat "$LOOP_PROMPT_FILE")"
 ```
 
+Opt into review/fix orchestration with `--review`:
+
+```bash
+LOOP_AGENT_CMD='codex exec --dangerously-bypass-approvals-and-sandbox --ephemeral "$(cat "$LOOP_PROMPT_FILE")"' \
+LOOP_REVIEW_CMD='claude --print --no-session-persistence --dangerously-skip-permissions "$(cat "$LOOP_PROMPT_FILE")"' \
+LOOP_FIX_CMD='codex exec --dangerously-bypass-approvals-and-sandbox --ephemeral "$(cat "$LOOP_PROMPT_FILE")"' \
+./loop.sh run .loop/<name>/QUEUE.md --review --max-review-rounds 2
+```
+
 ## CLI
 
 The `knack` binary is a read-only validator and context provider. Build it from `cli/`:
diff --git a/docs/loop.md b/docs/loop.md
index a5cd2cd28883321ca006489e2c630fbdc8d6bb8f..ef47822f4007c3fbcb5ee094f39164166ca63646
--- a/docs/loop.md
+++ b/docs/loop.md
@@ -5,18 +5,24 @@
 ## Usage
 
 ```bash
-./loop.sh run <queue> [--repo DIR] [--max-ticks N] [--dry-run]
+./loop.sh run <queue> [--repo DIR] [--max-ticks N] [--review] [--max-review-rounds N] [--dry-run]
 ```
 
 - `queue` — path to `QUEUE.md`. Usually `.loop/<name>/QUEUE.md`.
 - `--repo DIR` — the working directory for the worker and the verify command. Defaults to the parent of the `.loop/<name>` directory that contains the queue.
 - `--max-ticks N` — maximum units to attempt. Default is `3`.
+- `--review` — opt into the bounded review/fix subloop after pending build units drain. Default behavior does not run review.
+- `--max-review-rounds N` — maximum review/fix rounds when `--review` is enabled. Default is `2`.
 - `--dry-run` — parse the first pending unit and print its title, repo, and verify command, then exit.
 
 ## Environment variables
 
 - `LOOP_AGENT_CMD` — optional command used to invoke the worker. If unset, the loop defaults to `pi -p --no-session --approve` with the prompt text as a single argument. If set, the command is evaluated by `bash -lc` in the repo directory, and `LOOP_PROMPT_FILE` is set to a temporary file containing `prompts/worker.md` plus the current work unit.
 - `LOOP_PROMPT_FILE` — set by the loop when `LOOP_AGENT_CMD` is used. Points to the generated prompt file. Do not override unless you are calling the worker manually.
+- `LOOP_REVIEW_CMD` — optional command used for the review phase when `--review` is enabled. Defaults to `LOOP_AGENT_CMD`, or to the loop's default `pi` invocation if neither is set.
+- `LOOP_FIX_CMD` — optional command used for the fix phase when `--review` is enabled. Defaults to `LOOP_AGENT_CMD`, or to the loop's default `pi` invocation if neither is set.
+
+During every phase, the loop also sets `LOOP_PHASE`, `LOOP_QUEUE_FILE`, `LOOP_EVIDENCE_FILE`, and `LOOP_REVIEW_FILE` for the worker process.
 
 ## Per-unit agent override
 
@@ -43,6 +49,20 @@
    - If the repo changed, mark `verify_failed` and stop.
 9. If max ticks is reached with pending work, stop.
 
+## Optional review/fix subloop
+
+By default, `loop.sh` stops when the build queue has no pending work. With `--review`, it then runs a bounded review/fix subloop:
+
+1. Invoke a review worker with the review prompt and the completed queue/evidence.
+2. Require the review worker to write `.loop/<name>/REVIEW.md`.
+3. Read only the `- actionable: N` summary line from `REVIEW.md`.
+4. If `N` is `0`, stop cleanly.
+5. If `N` is non-zero, invoke a fix worker.
+6. Require the fix worker to append new `Status: pending` work units to the same `QUEUE.md`.
+7. Run the build pass again, then review again.
+
+The subloop stops when review is clean, `--max-review-rounds` is reached, `--max-ticks` is exhausted, or fix produces no pending work. The loop owns orchestration and stop conditions only; the `review` and `fix` skills own judgment and work-unit generation.
+
 ## Work unit statuses
 
 - `pending` — not yet started.
@@ -58,6 +78,7 @@
 
 - `EVIDENCE.md` — append-only ledger. Includes the full unit, changed files, verify command, verify output, and worker output. It is durable; keep it after deleting `QUEUE.md` so `knack decisions check` can still see which ADRs the cycle referenced.
 - `HANDOFF.md` — written on non-clean exit. Sections: completed, in progress, remaining, next action. Delete when the work resumes.
+- `REVIEW.md` — structured review artifact written by the review worker when `--review` is enabled. The loop reads only its `- actionable: N` summary line.
 
 ## Agent invocation examples
 
@@ -71,8 +92,17 @@
 
 The command is passed to `bash -lc` in the repo directory, so the typical pattern is `"$(cat "$LOOP_PROMPT_FILE")"`.
 
+Review and fix can use separate agents:
+
+```bash
+LOOP_AGENT_CMD='codex exec --dangerously-bypass-approvals-and-sandbox --ephemeral "$(cat "$LOOP_PROMPT_FILE")"' \
+LOOP_REVIEW_CMD='claude --print --no-session-persistence --dangerously-skip-permissions "$(cat "$LOOP_PROMPT_FILE")"' \
+LOOP_FIX_CMD='codex exec --dangerously-bypass-approvals-and-sandbox --ephemeral "$(cat "$LOOP_PROMPT_FILE")"' \
+./loop.sh run .loop/<name>/QUEUE.md --review
+```
+
 ## Verification notes
 
 - `loop.sh` does **not** validate `QUEUE.md` structure. Use `knack validate` first.
-- `loop.sh` does **not** run review or manage ADRs/glossary. Those are handled by the skills.
+- `loop.sh` runs review and fix only when `--review` is set. It invokes the skills and reads the actionable count from `REVIEW.md`; it does not judge findings or manage ADRs/glossary.
 - The `Verify:` command must be deterministic and executable by the runner — tests, builds, type checks, not an LLM-as-judge.
diff --git a/docs/skills.md b/docs/skills.md
index fb1ec98119e1193859c0a1612f580d2da5afe95d..476d494e3c88962825e751b84151a0ff48ac2d41
--- a/docs/skills.md
+++ b/docs/skills.md
@@ -37,7 +37,11 @@
 
 ## How the loop uses skills
 
-`loop.sh` does not read skills. It prepends `prompts/worker.md` to the current work unit and runs the worker. `prompts/worker.md` tells the worker to load the `build` skill by name and path (e.g. "Load and follow the **build** skill in `.agents/skills/build/`"). For other phases (explore, plan, review, fix), the human or the agent invokes the skill directly.
+`loop.sh` does not read skills. It prepends `prompts/worker.md` to the current work unit and runs the worker. `prompts/worker.md` tells the worker to load the `build` skill by name and path (e.g. "Load and follow the **build** skill in `.agents/skills/build/`").
+
+When `--review` is set, the loop also invokes review and fix workers after the build queue drains. Those prompts tell the worker to load the `review` or `fix` skill directly. The loop orchestrates the bounded review/fix subloop, reads the actionable count from `REVIEW.md`, and runs another build pass when fix appends pending units. The skills still own judgment: review decides what the findings are, and fix decides which findings become work units.
+
+Without `--review`, review and fix remain manual skill invocations.
 
 ## Customizing skills
 
@@ -53,7 +57,7 @@
 ```text
 small fix → plan → build → done
 bug report → explore → plan → build → done
-big feature → explore → plan → build → review → fix → done
+big feature → explore → plan → build --review → review → fix → build → done
 ```
 
 Decisions are captured inline throughout the flow using the `decide` skill, and terms are updated using `domain-modeling`.

 succeeded in 3926ms:
Valid skill: /home/daniel/build/knack/.agents/skills/build
Valid skill: /home/daniel/build/knack/.agents/skills/decide
Valid skill: /home/daniel/build/knack/.agents/skills/domain-modeling
Valid skill: /home/daniel/build/knack/.agents/skills/explore
Valid skill: /home/daniel/build/knack/.agents/skills/fix
Valid skill: /home/daniel/build/knack/.agents/skills/plan
Valid skill: /home/daniel/build/knack/.agents/skills/review
knack tests passed
ok  	knack	0.002s
ok  	knack/internal/decisions	(cached)
ok  	knack/internal/glossary	(cached)
ok  	knack/internal/instructions	(cached)
ok  	knack/internal/queue	(cached)
ok  	knack/internal/skills	0.004s
ok  	knack/internal/status	(cached)

codex
Unit: docs and AGENTS.md reflect the new loop behavior
Changed: docs/loop.md, docs/skills.md, README.md, AGENTS.md
Verify expected: cd /home/daniel/build/knack && ./tests/run.sh && cd cli && go test ./...
Notes: Local verify passed. Embedded skills are already in sync; no skill files changed in this tick.
diff --git a/AGENTS.md b/AGENTS.md
index 50f568dd4d4d4e617b4be191116bb8fea556720d..f005f47ed72aa4f97f47324d98c00d8c79b12921
--- a/AGENTS.md
+++ b/AGENTS.md
@@ -19,13 +19,14 @@
 
 ## Current state
 
-The loop (`loop.sh`), all seven skills, and the Go CLI are built. The loop supports per-unit `Agent:` overrides, `LOOP_AGENT_CMD` for agent-agnosticism, and writes `.loop/<name>/HANDOFF.md` on non-clean exits. The skills cover the full explore → plan → build → review → fix flow plus two shared skills (decide, domain-modeling). The CLI (`cli/`) provides `skills init|check`, `validate`, `decisions list|show|check`, `status`, `glossary check`, and `instructions`. It embeds the default skills via `go:embed` and scaffolds them into target projects.
+The loop (`loop.sh`), all seven skills, and the Go CLI are built. The loop supports per-unit `Agent:` overrides, `LOOP_AGENT_CMD` for agent-agnosticism, opt-in review/fix orchestration via `--review`, and `.loop/<name>/HANDOFF.md` on non-clean exits. The skills cover the full explore → plan → build → review → fix flow plus two shared skills (decide, domain-modeling). The CLI (`cli/`) provides `skills init|check`, `validate`, `decisions list|show|check`, `status`, `glossary check`, and `instructions`. It embeds the default skills via `go:embed` and scaffolds them into target projects.
 
 ## Core artifacts
 
 - `.loop/<name>/QUEUE.md` — disposable queue of verifiable work units.
 - `.loop/<name>/EVIDENCE.md` — append-only ledger of what each tick proved. Durable: keep it after deleting QUEUE.md so `decisions check` can still see which ADRs the completed cycle referenced.
 - `.loop/<name>/HANDOFF.md` — cross-session handoff (written on pause/stop).
+- `.loop/<name>/REVIEW.md` — structured review artifact written when `--review` is enabled. The loop reads only the actionable count.
 - `decisions/` — durable ADRs (architectural rulings, not current behavior).
 - `glossary.md` — durable ubiquitous language.
 - `.agents/skills/` — procedural knowledge (explore, plan, build, review, fix, decide, domain-modeling).
@@ -48,11 +49,13 @@
 - The worker never self-certifies. The runner owns verification.
 - A `Verify` command must be deterministic and executable by the runner — not an LLM-as-judge. The loop's backpressure is mechanical.
 - The runner enforces hard stops (max ticks, no-progress detection). The agent does one work unit per tick and reports what remains if it can't finish.
-- Review is opt-in. The loop does not run review automatically; the user invokes it when they want adversarial scrutiny.
+- Review is opt-in. The loop runs review/fix only when invoked with `--review`; otherwise it runs build ticks only.
+- With `--review`, the loop invokes review after pending build units drain, reads only `REVIEW.md`'s actionable count, invokes fix when actionable findings exist, then runs another build pass for appended units. Judgment stays in the `review` and `fix` skills.
 - When the queue is complete and verified, the cycle's `.loop/<name>/QUEUE.md` is disposable. The human deletes it; the loop does not. `EVIDENCE.md` is the durable ledger — keep it so `decisions check` can still trace which ADRs the cycle referenced.
 - A work unit must leave the repo better if the loop stops immediately after it.
 - Work units are whatever shape the work is — not forced into "vertical slices."
 - `LOOP_AGENT_CMD` overrides the worker invocation (agent-agnostic: pi, claude, codex, opencode, etc.).
+- `LOOP_REVIEW_CMD` and `LOOP_FIX_CMD` override review/fix phase invocations when `--review` is enabled. They default to `LOOP_AGENT_CMD`.
 - Per-unit `Agent:` field in QUEUE.md overrides `LOOP_AGENT_CMD` for one unit (model routing).
 - Work units are `## <outcome>` headers with `Read first:`, `Constraints:`, `Done means:`, and `Verify:` fields. `Done means:` is the acceptance criteria; `Verify:` is the mechanically enforceable subset. The gap between them is the review surface.
 - `Read first:` is context, not scope — 2–4 entries of ADRs, code areas, or rulings.
diff --git a/README.md b/README.md
index 1300a1bc12102f97cd89ace07333f179f8700cee..c1541df3127dce0ba5ba94a9cc4d8ba97c5a6b6a
--- a/README.md
+++ b/README.md
@@ -43,6 +43,9 @@
 6. On failure: the unit is marked `verify_failed` or `no_progress`, evidence is appended, and the loop stops or retries once.
 7. The loop halts on: max ticks reached with pending work, two no-progress strikes, or a failed verify after a real change.
 8. On any non-clean exit, the runner writes `HANDOFF.md` with completed/in-progress/remaining units and the next action.
+9. If `--review` is set, the runner invokes review after the queue drains, reads the actionable count from `REVIEW.md`, invokes fix when actionable findings exist, and runs another build pass for appended units.
+
+Review is opt-in. Without `--review`, the loop runs build ticks only; with it, review/fix are loop-orchestrated but still skill-owned.
 
 ## Queue format
 
@@ -70,6 +73,15 @@
 Agent: pi -p --no-session --approve --model glm-5.2 "$(cat "$LOOP_PROMPT_FILE")"
 ```
 
+Opt into review/fix orchestration with `--review`:
+
+```bash
+LOOP_AGENT_CMD='codex exec --dangerously-bypass-approvals-and-sandbox --ephemeral "$(cat "$LOOP_PROMPT_FILE")"' \
+LOOP_REVIEW_CMD='claude --print --no-session-persistence --dangerously-skip-permissions "$(cat "$LOOP_PROMPT_FILE")"' \
+LOOP_FIX_CMD='codex exec --dangerously-bypass-approvals-and-sandbox --ephemeral "$(cat "$LOOP_PROMPT_FILE")"' \
+./loop.sh run .loop/<name>/QUEUE.md --review --max-review-rounds 2
+```
+
 ## CLI
 
 The `knack` binary is a read-only validator and context provider. Build it from `cli/`:
diff --git a/docs/loop.md b/docs/loop.md
index a5cd2cd28883321ca006489e2c630fbdc8d6bb8f..ef47822f4007c3fbcb5ee094f39164166ca63646
--- a/docs/loop.md
+++ b/docs/loop.md
@@ -5,18 +5,24 @@
 ## Usage
 
 ```bash
-./loop.sh run <queue> [--repo DIR] [--max-ticks N] [--dry-run]
+./loop.sh run <queue> [--repo DIR] [--max-ticks N] [--review] [--max-review-rounds N] [--dry-run]
 ```
 
 - `queue` — path to `QUEUE.md`. Usually `.loop/<name>/QUEUE.md`.
 - `--repo DIR` — the working directory for the worker and the verify command. Defaults to the parent of the `.loop/<name>` directory that contains the queue.
 - `--max-ticks N` — maximum units to attempt. Default is `3`.
+- `--review` — opt into the bounded review/fix subloop after pending build units drain. Default behavior does not run review.
+- `--max-review-rounds N` — maximum review/fix rounds when `--review` is enabled. Default is `2`.
 - `--dry-run` — parse the first pending unit and print its title, repo, and verify command, then exit.
 
 ## Environment variables
 
 - `LOOP_AGENT_CMD` — optional command used to invoke the worker. If unset, the loop defaults to `pi -p --no-session --approve` with the prompt text as a single argument. If set, the command is evaluated by `bash -lc` in the repo directory, and `LOOP_PROMPT_FILE` is set to a temporary file containing `prompts/worker.md` plus the current work unit.
 - `LOOP_PROMPT_FILE` — set by the loop when `LOOP_AGENT_CMD` is used. Points to the generated prompt file. Do not override unless you are calling the worker manually.
+- `LOOP_REVIEW_CMD` — optional command used for the review phase when `--review` is enabled. Defaults to `LOOP_AGENT_CMD`, or to the loop's default `pi` invocation if neither is set.
+- `LOOP_FIX_CMD` — optional command used for the fix phase when `--review` is enabled. Defaults to `LOOP_AGENT_CMD`, or to the loop's default `pi` invocation if neither is set.
+
+During every phase, the loop also sets `LOOP_PHASE`, `LOOP_QUEUE_FILE`, `LOOP_EVIDENCE_FILE`, and `LOOP_REVIEW_FILE` for the worker process.
 
 ## Per-unit agent override
 
@@ -43,6 +49,20 @@
    - If the repo changed, mark `verify_failed` and stop.
 9. If max ticks is reached with pending work, stop.
 
+## Optional review/fix subloop
+
+By default, `loop.sh` stops when the build queue has no pending work. With `--review`, it then runs a bounded review/fix subloop:
+
+1. Invoke a review worker with the review prompt and the completed queue/evidence.
+2. Require the review worker to write `.loop/<name>/REVIEW.md`.
+3. Read only the `- actionable: N` summary line from `REVIEW.md`.
+4. If `N` is `0`, stop cleanly.
+5. If `N` is non-zero, invoke a fix worker.
+6. Require the fix worker to append new `Status: pending` work units to the same `QUEUE.md`.
+7. Run the build pass again, then review again.
+
+The subloop stops when review is clean, `--max-review-rounds` is reached, `--max-ticks` is exhausted, or fix produces no pending work. The loop owns orchestration and stop conditions only; the `review` and `fix` skills own judgment and work-unit generation.
+
 ## Work unit statuses
 
 - `pending` — not yet started.
@@ -58,6 +78,7 @@
 
 - `EVIDENCE.md` — append-only ledger. Includes the full unit, changed files, verify command, verify output, and worker output. It is durable; keep it after deleting `QUEUE.md` so `knack decisions check` can still see which ADRs the cycle referenced.
 - `HANDOFF.md` — written on non-clean exit. Sections: completed, in progress, remaining, next action. Delete when the work resumes.
+- `REVIEW.md` — structured review artifact written by the review worker when `--review` is enabled. The loop reads only its `- actionable: N` summary line.
 
 ## Agent invocation examples
 
@@ -71,8 +92,17 @@
 
 The command is passed to `bash -lc` in the repo directory, so the typical pattern is `"$(cat "$LOOP_PROMPT_FILE")"`.
 
+Review and fix can use separate agents:
+
+```bash
+LOOP_AGENT_CMD='codex exec --dangerously-bypass-approvals-and-sandbox --ephemeral "$(cat "$LOOP_PROMPT_FILE")"' \
+LOOP_REVIEW_CMD='claude --print --no-session-persistence --dangerously-skip-permissions "$(cat "$LOOP_PROMPT_FILE")"' \
+LOOP_FIX_CMD='codex exec --dangerously-bypass-approvals-and-sandbox --ephemeral "$(cat "$LOOP_PROMPT_FILE")"' \
+./loop.sh run .loop/<name>/QUEUE.md --review
+```
+
 ## Verification notes
 
 - `loop.sh` does **not** validate `QUEUE.md` structure. Use `knack validate` first.
-- `loop.sh` does **not** run review or manage ADRs/glossary. Those are handled by the skills.
+- `loop.sh` runs review and fix only when `--review` is set. It invokes the skills and reads the actionable count from `REVIEW.md`; it does not judge findings or manage ADRs/glossary.
 - The `Verify:` command must be deterministic and executable by the runner — tests, builds, type checks, not an LLM-as-judge.
diff --git a/docs/skills.md b/docs/skills.md
index fb1ec98119e1193859c0a1612f580d2da5afe95d..476d494e3c88962825e751b84151a0ff48ac2d41
--- a/docs/skills.md
+++ b/docs/skills.md
@@ -37,7 +37,11 @@
 
 ## How the loop uses skills
 
-`loop.sh` does not read skills. It prepends `prompts/worker.md` to the current work unit and runs the worker. `prompts/worker.md` tells the worker to load the `build` skill by name and path (e.g. "Load and follow the **build** skill in `.agents/skills/build/`"). For other phases (explore, plan, review, fix), the human or the agent invokes the skill directly.
+`loop.sh` does not read skills. It prepends `prompts/worker.md` to the current work unit and runs the worker. `prompts/worker.md` tells the worker to load the `build` skill by name and path (e.g. "Load and follow the **build** skill in `.agents/skills/build/`").
+
+When `--review` is set, the loop also invokes review and fix workers after the build queue drains. Those prompts tell the worker to load the `review` or `fix` skill directly. The loop orchestrates the bounded review/fix subloop, reads the actionable count from `REVIEW.md`, and runs another build pass when fix appends pending units. The skills still own judgment: review decides what the findings are, and fix decides which findings become work units.
+
+Without `--review`, review and fix remain manual skill invocations.
 
 ## Customizing skills
 
@@ -53,7 +57,7 @@
 ```text
 small fix → plan → build → done
 bug report → explore → plan → build → done
-big feature → explore → plan → build → review → fix → done
+big feature → explore → plan → build --review → review → fix → build → done
 ```
 
 Decisions are captured inline throughout the flow using the `decide` skill, and terms are updated using `domain-modeling`.

tokens used
38,672
Unit: docs and AGENTS.md reflect the new loop behavior
Changed: docs/loop.md, docs/skills.md, README.md, AGENTS.md
Verify expected: cd /home/daniel/build/knack && ./tests/run.sh && cd cli && go test ./...
Notes: Local verify passed. Embedded skills are already in sync; no skill files changed in this tick.
````

What this proves:
- The verify command passed for this work unit in the current repo state.

What remains unverified:
- Anything outside the verify command's proof scope.

## 2026-07-10T19:57:46-06:00 — review skill writes a structured REVIEW.md

Status: done

Unit:
````markdown
## review skill writes a structured REVIEW.md

Read first:
- `.agents/skills/review/SKILL.md`
- `.agents/skills/fix/SKILL.md`
- `decisions/0008-loop-orchestrates-review-fix.md`

Constraints:
- Preserve the two-axis (standards + intent) review structure.
- The output format must be machine-parseable enough for `loop.sh` to detect the `actionable` count.
- Do not change `loop.sh` or `fix/SKILL.md` yet.
- Keep the skill file scannable and sync it to `cli/embedded/skills/review/` after editing.

Done means:
- `review/SKILL.md` documents `REVIEW.md` as its output artifact.
- `REVIEW.md` has sections for `## Standards`, `## Intent`, `## Speculative`, and a `## Summary` with counts.
- Each finding has a confidence level (high/medium/low) and evidence (`path/to/file:line` or quoted code).
- `cli/embedded/skills/review/SKILL.md` is in sync with `.agents/skills/review/SKILL.md`.

Verify:
```bash
cd /home/daniel/build/knack && diff -r .agents/skills/review cli/embedded/skills/review && ./tests/run.sh
```

Status: pending````

Files changed:
```text
 M .agents/skills/fix/SKILL.md
 M .agents/skills/review/SKILL.md
 M .loop/review-loop/EVIDENCE.md
 M .loop/review-loop/HANDOFF.md
 M .loop/review-loop/QUEUE.md
 M AGENTS.md
 M README.md
 M cli/embedded/skills/fix/SKILL.md
 M cli/embedded/skills/review/SKILL.md
 M docs/loop.md
 M docs/skills.md
 M loop.sh
 M tests/run.sh
```

Verify command:
```bash
cd /home/daniel/build/knack && diff -r .agents/skills/review cli/embedded/skills/review && ./tests/run.sh
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
OpenAI Codex v0.144.1
--------
workdir: /home/daniel/build/knack
model: gpt-5.5
provider: openai
approval: never
sandbox: danger-full-access
reasoning effort: high
reasoning summaries: none
session id: 019f4ee4-1ba3-7e10-81ef-247bdd81df72
--------
user
# Knack Worker

You are one worker tick in a bounded loop. Complete exactly one work unit, then stop.

Load and follow the **build** skill in `.agents/skills/build/` before doing any work.

## Rules

1. Read `AGENTS.md` first if it exists — it contains operational context.
2. Read the current work unit carefully, especially its `Verify` command.
3. Do only the work needed for that unit.
4. Do not start another work unit.
5. Do not mark the unit complete yourself.
6. Do not edit `.loop/<name>/EVIDENCE.md`; the runner writes evidence after verification.
7. If the unit is blocked, make the smallest useful note in your final response and stop.
8. Stay within the runner's hard stops (max ticks, no-progress detection). If the unit is too large for one tick, do as much as keeps the repo working and report what remains.
9. The worker's scope is the unit's outcome plus its constraints. Keep the diff narrow and aligned with that scope.
10. If verification fails while you are working, fix the cause if it belongs to this unit; otherwise stop and report the blocker.

## Success standard

Your job is not to claim success. Your job is to make the repository state satisfy the unit's `Verify` command. The runner will execute that command after you exit.

## Output

End with a compact terminal handoff. This is not `.loop/<name>/HANDOFF.md` — the runner writes that file. This is your own summary to the runner:

```text
Unit: <title>
Changed: <brief file/area list>
Verify expected: <command from unit>
Notes: <blockers or caveats, if any>
```

Current work unit from /home/daniel/build/knack/.loop/review-loop/QUEUE.md:

## review skill writes a structured REVIEW.md

Read first:
- `.agents/skills/review/SKILL.md`
- `.agents/skills/fix/SKILL.md`
- `decisions/0008-loop-orchestrates-review-fix.md`

Constraints:
- Preserve the two-axis (standards + intent) review structure.
- The output format must be machine-parseable enough for `loop.sh` to detect the `actionable` count.
- Do not change `loop.sh` or `fix/SKILL.md` yet.
- Keep the skill file scannable and sync it to `cli/embedded/skills/review/` after editing.

Done means:
- `review/SKILL.md` documents `REVIEW.md` as its output artifact.
- `REVIEW.md` has sections for `## Standards`, `## Intent`, `## Speculative`, and a `## Summary` with counts.
- Each finding has a confidence level (high/medium/low) and evidence (`path/to/file:line` or quoted code).
- `cli/embedded/skills/review/SKILL.md` is in sync with `.agents/skills/review/SKILL.md`.

Verify:
```bash
cd /home/daniel/build/knack && diff -r .agents/skills/review cli/embedded/skills/review && ./tests/run.sh
```

Status: pending
hook: SessionStart
hook: SessionStart Completed
codex
I’ll load the build skill and the referenced review/fix/ADR context, then make the narrow skill-doc change and sync the embedded copy.
exec
/usr/bin/bash -lc "sed -n '1,240p' .agents/skills/build/SKILL.md" in /home/daniel/build/knack
 succeeded in 0ms:
---
name: build
description: Use when implementing one work unit from a `.loop/<name>/QUEUE.md` queue. The worker-side skill — read the unit, do the work, don't self-certify, end with a handoff. Triggers on "build", "implement", "apply this unit", "do the work", "run the loop", or when a work unit needs to be executed. Also use when the loop invokes the worker for a tick.
---

# Build

Implement one work unit from `.loop/<name>/QUEUE.md`. Do the work. Don't self-certify. The loop runner owns the verify gate — your job is to make the repository state satisfy the unit's `Verify` command, not to claim success.

## Core rules

The runner injects `prompts/worker.md` at the start of every tick. If it is not in your context, read it now. Its `Rules` and `Output` sections are the canonical source for this skill; the sections below elaborate on decisions, operational learnings, blockers, and units that are too large for one tick.

> **Scope note:** Updating `AGENTS.md` or writing an ADR during a tick is a durable trace, not part of the unit's `Verify` scope. Do it only when the tick teaches you something that would save the next session time.

## Scope

The unit's scope is its outcome plus its constraints. The worker determines which files to change and how. The `Verify:` command is the mechanically enforceable subset of `Done means:`.

## Capturing decisions during build

If you discover an architectural ruling while implementing — "we need to handle X this way because Y" — capture it as an ADR using the `decide` skill. Do this inline, not after the unit. Decisions made during implementation are the most valuable kind because they come from the code pushing back.

## Capturing operational learnings

If you learn something about how the project works — a build command, a test convention, a gotcha — that isn't already in `AGENTS.md`, add it. `AGENTS.md` is the living operational context. If it grows stale, the next agent session starts with worse information.

Don't add trivia. Add things that would have saved you time if you'd known them upfront.

## When the unit is too big

If the work unit is larger than what can be done in one tick:

- Do as much as you can while keeping the repo in a working state.
- End with a handoff note explaining what remains.
- The runner will re-queue the unit if verify fails.

Do not try to do everything in one tick if the work genuinely needs more. The loop is designed for multiple ticks.

## When you're blocked

If you hit a blocker you can't resolve within the unit's scope:

1. State the blocker clearly in your final output.
2. Note what would unblock you (a decision, a dependency, a missing file).
3. Stop. Don't thrash.

The runner will mark the unit as `blocked` and write a handoff. The next session can pick up from there.

## Output

The output format is defined in the worker prompt (`prompts/worker.md`). End with a compact terminal handoff there.

exec
/usr/bin/bash -lc "sed -n '1,260p' .agents/skills/review/SKILL.md" in /home/daniel/build/knack
 succeeded in 0ms:
---
name: review
description: Use when reviewing completed work units or a finished queue. Two-axis adversarial review — standards (does the change follow codebase conventions?) and intent (does the change do what the work unit said?). Reviews against the actual codebase, not specs. Triggers on "review", "check this", "is this right", "what did we miss", "stress-test the implementation", or when work needs adversarial scrutiny before being accepted.
---

# Review

Adversarial review of completed work. Two axes, run independently so neither pollutes the other:

1. **Standards** — does the change follow the repo's coding conventions and codebase patterns?
2. **Intent** — does the change do what the work unit said it would?

Review against the **actual codebase**, not against specs that may have rotted. The code is the source of truth.

## When to review

- After a work unit is marked `done` and verify passed
- After a full queue is completed
- When the user asks for a sanity check
- Before accepting work as finished

Review is not a gate the loop enforces — it's a skill the user or agent invokes when they want adversarial scrutiny.

## Before you review

Read the work unit from `.loop/<name>/QUEUE.md` and the evidence from `.loop/<name>/EVIDENCE.md` for the unit you're reviewing. The evidence tells you what the verify command actually proved; the work unit tells you what was promised. Review against the actual codebase, not the specs.

## Two-axis review

### Axis 1: Standards

Does the change follow the codebase's existing patterns?

- Read `AGENTS.md` for stated conventions.
- Read neighboring code — does the change look like it belongs?
- Check error handling, naming, file layout, test style.
- Look for regressions — did the change break something nearby?
- Check for dead code, unused imports, leftover debugging.

The question is not "is this good code?" — that's subjective. The question is "does this match the codebase's own standards?"

### Axis 2: Intent

Does the change do what the work unit said it would?

- Read the work unit's `Done means:` and `Constraints:` fields.
- Read the actual diff.
- Does the diff satisfy the stated outcome?
- Does the diff stay within the stated constraints?
- Did the change introduce anything the unit didn't ask for?
- Did the change miss anything the unit said it would do?

The `Verify:` command is the mechanically enforceable subset of `Done means:`. The gap between `Done means:` and `Verify:` is the review surface: intent review checks what the verify command cannot.

> **Intent review is a judgment check, not a deterministic gate.** Be explicit about its scope and confidence. Any finding that becomes a new work unit must have a deterministic `Verify` command the runner can execute.

## Running the axes

Run both axes. They can be parallel (two passes over the same diff) or sequential. The order doesn't matter — what matters is that each axis is evaluated independently, without the other's conclusions bleeding in.

## Findings become input to the fix skill

Review findings are not just notes — they are the input to the `fix` skill, which triages them and appends actionable ones as new work units in `.loop/<name>/QUEUE.md`.

- **Trivial** findings (a typo, a missing newline) can be fixed inline during review.
- **Actionable** findings are handed to the `fix` skill. Do not write the work units yourself; `fix` owns the triage and formatting.
- **Disputed** or **deferred** findings are recorded in the review summary but not turned into units.

The output of review is a findings summary, not a queue edit.

## What review is not

- Not a lint pass — the verify gate already ran. Review is about what verify *can't* check.
- Not a spec compliance check — specs are disposable and may have rotted. Review against the codebase.
- Not a gate — the loop doesn't enforce review. It's invoked when the user wants it.

## Output

Summarize findings per axis:

- **Standards**: N findings (list them, or "no issues found")
- **Intent**: N findings (list them, or "the change matches the unit's stated outcome")

Classify each finding as trivial / actionable / disputed / deferred.

Then either:
- Hand actionable findings to the `fix` skill, which will triage and append work units to `.loop/<name>/QUEUE.md`, or
- Report "no action needed" if the work is clean.

exec
/usr/bin/bash -lc "sed -n '1,260p' .agents/skills/fix/SKILL.md" in /home/daniel/build/knack
 succeeded in 0ms:
---
name: fix
description: Use when addressing review findings and feeding them back into the loop. Converts review findings into new verifiable work units in `.loop/<name>/QUEUE.md` and runs another loop pass. Triggers on "fix the review findings", "address the feedback", "rework based on review", "fix what review found", or when review produced findings that need to be resolved.
---

# Fix

Address review findings by triaging `REVIEW.md` and generating new work units that feed back into the loop. The `fix` skill is the bridge between review and another build pass; it owns the triage and the queue format. It appends to the existing `.loop/<name>/QUEUE.md` file for the same cycle.

The loop owns orchestration. When `fix` is invoked by the loop, stop after appending units and reporting the triage summary; the loop will run the next build pass.

## Inputs

- `.loop/<name>/REVIEW.md` — structured findings from the `review` skill. Read it, but do not edit it.
- `.loop/<name>/QUEUE.md` — the existing work queue. Read it before appending so new units preserve the queue's current structure and statuses.

`REVIEW.md` is expected to include `## Standards`, `## Intent`, `## Speculative`, and `## Summary` sections. The summary's actionable count is the loop's signal; the finding entries are the fix skill's input. Use the finding classification, confidence, and evidence when deciding whether to generate a work unit.

## Procedure

1. **Read `REVIEW.md`.** Each standards or intent finding is a candidate work unit. `review` has already classified findings as trivial / actionable / disputed / deferred; use those classifications or re-triage as needed. Treat speculative findings as notes for future explore/plan work unless the review explicitly classifies one as actionable.

2. **Triage.** Not every finding warrants a work unit:
   - **Actionable** — the finding identifies a real issue with a clear fix. Create a work unit.
   - **Trivial** — one-line fix, no risk. Note it in the summary, don't create a unit.
   - **Disputed** — the finding is wrong or the reviewer is being overly cautious. Note the disagreement in the summary, don't create a unit. Move on.
   - **Deferred** — the finding is valid but not worth fixing now. Note it in the summary or a backlog, don't create a unit.

3. **Create work units** for actionable findings. Use the standard work unit format from the `plan` skill. Each unit must have one observable outcome and one deterministic verify command. Reference the source finding in `Read first:` by pointing to `REVIEW.md` plus any evidence paths from the finding.

````markdown
## <fix for the finding — observable outcome>

Read first:
- .loop/<name>/REVIEW.md (<finding id or heading>)
- <evidence path from the finding, if any>
- <2–4 entries; context, not scope>

Constraints:
- <boundary>
- <what must stay true or out of bounds; if a file is named, it is "don't touch X" or "X's public API must not change", not "update X">

Done means:
- <the finding is resolved>
- <no new issue introduced>

Verify:
```bash
<deterministic command that proves the fix>
```

Status: pending
````

4. **Append to `QUEUE.md`.** Append new units to the end of the existing `.loop/<name>/QUEUE.md`. Do not reorder existing units, do not edit existing unit bodies, and do not change any existing `Status:` lines. If there are no actionable findings, append nothing.

5. **Stop after the queue edit.** Report the triage summary and the number of appended units. The loop will decide whether to run another build pass.

## What fix is not

- Not a re-review — the findings are already known. Fix turns them into work.
- Not a debate — if a finding is disputed, note it and move on. Don't argue with the review in the queue.
- Not a rewrite — each fix unit should be narrow. If a finding requires a broad rewrite, that's a new explore → plan cycle, not a fix.
- Not the orchestrator — the loop invokes review, invokes fix, and runs the next build pass.

## Guardrail

Fix units must not break what review already approved. Each fix unit's `Done means:` should include a no-regression condition. The verify command should cover both the fix and the preservation of what was already working.

## Output

- New `Status: pending` work units appended to `.loop/<name>/QUEUE.md`
- A summary of what was triaged (actionable / trivial / disputed / deferred)
- The number of units appended

exec
/usr/bin/bash -lc "sed -n '1,220p' decisions/0008-loop-orchestrates-review-fix.md" in /home/daniel/build/knack
 succeeded in 0ms:
# 0008: The loop orchestrates the review-fix subloop; skills keep the judgment

Date: 2026-07-10
Status: accepted

## Context

The default skill flow is `explore → plan → build → review → fix`, but `loop.sh` only runs `build` ticks. Review and fix are left as manual skills, so the loop cannot autonomously review a completed queue, turn findings into fix units, re-run the build pass, and stop when review is clean. That is the missing piece for a bounded, self-correcting development loop.

The architectural question is *who owns what*. DESIGN.md's "What the loop does NOT do" already states the loop does not run review and does not manage ADRs — those are skill responsibilities. Extending the loop into review/fix risks blurring that boundary: if the loop starts interpreting review content, it stops being the simple, agent-agnostic engine and becomes an LLM-as-judge.

Two roles must stay separate:

- **Orchestration** (the loop) — when to invoke review, when to invoke fix, when to stop. Mechanical and signal-driven.
- **Judgment** (the `review` and `fix` skills) — what counts as a finding, whether it is actionable, how to phrase a fix unit. Adversarial and semantic.

Alternatives considered:

- **LLM-as-judge inside the loop.** Rejected — the loop would interpret review content, breaking agent-agnosticism and the verify-gate principle (the aiming problem: the signal must be the actual property, not a proxy the loop can game).
- **The loop implements review logic.** Rejected — duplicates the `review` skill, drifts from the source of truth, and violates "the loop never reads skills."
- **Manual review/fix only (do nothing).** Rejected — leaves the loop unable to self-correct; the missing piece stays missing.

## Decision

The loop **orchestrates** an optional, bounded `build → review → fix` subloop. It invokes the `review` and `fix` workers and interprets their *structured* outputs as continue/stop signals. It does not implement review or fix logic.

When `--review` is set and the build queue has drained, the loop runs a review worker that writes a structured `REVIEW.md`. The loop reads only the `actionable` count from that file. If it is non-zero, the loop runs a fix worker that appends `Status: pending` units to `QUEUE.md`, then re-runs the build pass and reviews again. It stops when `actionable` is zero, a review-round limit (`--max-review-rounds`, default 2) is hit, the tick budget (`--max-ticks`) is exhausted, or a round generates no new units (no progress).

The boundary, stated as a ruling:

- **The loop owns orchestration and stop conditions** — invoking review/fix, reading the actionable count, enforcing hard stops. Mechanical.
- **The `review` and `fix` skills own judgment** — the two-axis standards/intent review, triage, finding phrasing, and work-unit generation. Semantic.

The loop never parses review content beyond the actionable count and never judges whether a finding is real. Review remains a skill the worker loads; the loop only knows to invoke it and read its summary.

## Consequences

- The loop gains a bounded self-correction capability without becoming an LLM-as-judge. Backpressure stays mechanical: the actionable count is the signal, the hard stops are the guardrails.
- A new contract surface: the loop depends on `REVIEW.md`'s actionable count being honest. A review worker that mislabels findings can loop forever or stop early — the round cap and the no-progress stop are the backstop, not the loop's own judgment.
- The `review` and `fix` skills gain a machine-readable input/output contract (`REVIEW.md`; appended units). Those skill changes land in their own work units; this ADR records only the architectural ruling and the orchestration boundary.
- Review stays opt-in (`--review`); default loop behavior is unchanged.

codex
The current review skill already has the two-axis content, but its output contract is still prose-oriented. I’m going to tighten only that section so it writes `.loop/<name>/REVIEW.md` with parseable counts and evidence-bearing findings, then copy it into the embedded skill directory.
exec
/usr/bin/bash -lc "sed -n '1,260p' cli/embedded/skills/review/SKILL.md" in /home/daniel/build/knack
 succeeded in 0ms:
---
name: review
description: Use when reviewing completed work units or a finished queue. Two-axis adversarial review — standards (does the change follow codebase conventions?) and intent (does the change do what the work unit said?). Reviews against the actual codebase, not specs. Triggers on "review", "check this", "is this right", "what did we miss", "stress-test the implementation", or when work needs adversarial scrutiny before being accepted.
---

# Review

Adversarial review of completed work. Two axes, run independently so neither pollutes the other:

1. **Standards** — does the change follow the repo's coding conventions and codebase patterns?
2. **Intent** — does the change do what the work unit said it would?

Review against the **actual codebase**, not against specs that may have rotted. The code is the source of truth.

## When to review

- After a work unit is marked `done` and verify passed
- After a full queue is completed
- When the user asks for a sanity check
- Before accepting work as finished

Review is not a gate the loop enforces — it's a skill the user or agent invokes when they want adversarial scrutiny.

## Before you review

Read the work unit from `.loop/<name>/QUEUE.md` and the evidence from `.loop/<name>/EVIDENCE.md` for the unit you're reviewing. The evidence tells you what the verify command actually proved; the work unit tells you what was promised. Review against the actual codebase, not the specs.

## Two-axis review

### Axis 1: Standards

Does the change follow the codebase's existing patterns?

- Read `AGENTS.md` for stated conventions.
- Read neighboring code — does the change look like it belongs?
- Check error handling, naming, file layout, test style.
- Look for regressions — did the change break something nearby?
- Check for dead code, unused imports, leftover debugging.

The question is not "is this good code?" — that's subjective. The question is "does this match the codebase's own standards?"

### Axis 2: Intent

Does the change do what the work unit said it would?

- Read the work unit's `Done means:` and `Constraints:` fields.
- Read the actual diff.
- Does the diff satisfy the stated outcome?
- Does the diff stay within the stated constraints?
- Did the change introduce anything the unit didn't ask for?
- Did the change miss anything the unit said it would do?

The `Verify:` command is the mechanically enforceable subset of `Done means:`. The gap between `Done means:` and `Verify:` is the review surface: intent review checks what the verify command cannot.

> **Intent review is a judgment check, not a deterministic gate.** Be explicit about its scope and confidence. Any finding that becomes a new work unit must have a deterministic `Verify` command the runner can execute.

## Running the axes

Run both axes. They can be parallel (two passes over the same diff) or sequential. The order doesn't matter — what matters is that each axis is evaluated independently, without the other's conclusions bleeding in.

## Findings become input to the fix skill

Review findings are not just notes — they are the input to the `fix` skill, which triages them and appends actionable ones as new work units in `.loop/<name>/QUEUE.md`.

- **Trivial** findings (a typo, a missing newline) can be fixed inline during review.
- **Actionable** findings are handed to the `fix` skill. Do not write the work units yourself; `fix` owns the triage and formatting.
- **Disputed** or **deferred** findings are recorded in the review summary but not turned into units.

The output of review is a findings summary, not a queue edit.

## What review is not

- Not a lint pass — the verify gate already ran. Review is about what verify *can't* check.
- Not a spec compliance check — specs are disposable and may have rotted. Review against the codebase.
- Not a gate — the loop doesn't enforce review. It's invoked when the user wants it.

## Output

Summarize findings per axis:

- **Standards**: N findings (list them, or "no issues found")
- **Intent**: N findings (list them, or "the change matches the unit's stated outcome")

Classify each finding as trivial / actionable / disputed / deferred.

Then either:
- Hand actionable findings to the `fix` skill, which will triage and append work units to `.loop/<name>/QUEUE.md`, or
- Report "no action needed" if the work is clean.

exec
/usr/bin/bash -lc 'git status --short' in /home/daniel/build/knack
 succeeded in 0ms:
 M .agents/skills/fix/SKILL.md
 M .loop/review-loop/EVIDENCE.md
 M .loop/review-loop/HANDOFF.md
 M .loop/review-loop/QUEUE.md
 M AGENTS.md
 M README.md
 M cli/embedded/skills/fix/SKILL.md
 M docs/loop.md
 M docs/skills.md
 M loop.sh
 M tests/run.sh

exec
/usr/bin/bash -lc "rg -n \"REVIEW.md|actionable|Summary|review\" .agents cli tests loop.sh DESIGN.md prompts -g '"'!cli/embedded/skills/review/SKILL.md'"'" in /home/daniel/build/knack
 succeeded in 0ms:
DESIGN.md:23:explore  plan  build  review  fix     │                          │
DESIGN.md:68:├── review/SKILL.md           # two-axis parallel review:
DESIGN.md:71:│                             # review against actual codebase, not specs
DESIGN.md:72:├── fix/SKILL.md              # address review findings, generate new work units
DESIGN.md:74:│                             # used by explore, plan, build, review
DESIGN.md:76:                              # used by explore, plan, review
DESIGN.md:81:- **Composable, not monolithic.** The flow is a default path, not a gate. Skills can be invoked independently. `bug → explore → plan → build → done` is as valid as `big feature → explore → plan (with specs) → build → review → fix → done`.
DESIGN.md:96:**review** — two-axis parallel review (pattern from mattpocock's `code-review`), starting from the work unit and `.loop/<name>/EVIDENCE.md`:
DESIGN.md:102:**fix** — address review findings. Read the existing `.loop/<name>/QUEUE.md`, append new work units generated from findings, and run another loop pass.
DESIGN.md:106:**domain-modeling** (shared) — manage the project's ubiquitous language. Challenge terms against the glossary, stress-test with edge-case scenarios, update `glossary.md` inline. Used by explore, plan, review.
DESIGN.md:178:- It does not run review (that's a skill the agent loads).
DESIGN.md:205:- No `review` — that's a skill the agent loads, not a CLI command.
DESIGN.md:229:│   ├── review/SKILL.md
DESIGN.md:303:4. REVIEW (human re-runs the agent, or runs it themselves, with review skill)
DESIGN.md:304:   agent loads review skill
DESIGN.md:305:   two-axis parallel review:
DESIGN.md:308:   review against actual codebase, not against specs
DESIGN.md:312:   loop runs again on the new work units from review
DESIGN.md:329:architecture review → explore → plan → build (loop) → review → fix (loop) → done
DESIGN.md:331:big feature → explore → plan (with specs) → build (loop) → review → fix (loop) → done
DESIGN.md:383:- **Done means: is the acceptance criteria; Verify: is the enforceable subset.** The gap between them is the review surface.
DESIGN.md:391:- The flow shape: explore → plan → build → review → fix (now composable, not rigid)
DESIGN.md:392:- Skills as procedural knowledge (think/plan/build/review → explore/plan/build/review/fix + shared decide/domain-modeling)
DESIGN.md:419:- No upstream (explore/plan) or downstream (review) → now covered by skills
DESIGN.md:427:- Two-axis parallel review (Standards vs Intent, run as parallel sub-agents)
DESIGN.md:485:2. **Split the planner skill.** The current `vertical-slice-planner` becomes `plan`. The `explore`, `build`, `review`, `fix`, `decide`, and `domain-modeling` skills are new and need to be authored.
loop.sh:7:  ./loop.sh run <queue> [--repo DIR] [--max-ticks N] [--review] [--max-review-rounds N] [--dry-run]
loop.sh:207:extract_actionable_count() {
loop.sh:209:    /^[[:space:]]*-[[:space:]]*actionable:[[:space:]]*[0-9]+[[:space:]]*$/ {
loop.sh:210:      sub(/^[[:space:]]*-[[:space:]]*actionable:[[:space:]]*/, "")
loop.sh:220:write_review_prompt() {
loop.sh:222:  local template="$script_dir/prompts/reviewer.md"
loop.sh:229:Load and follow the **review** skill in `.agents/skills/review/`.
loop.sh:230:Review the completed queue against the current repository state. Write the structured review artifact at the path provided below.
loop.sh:237:Review output: $review_file
loop.sh:251:Read the structured review artifact and append any actionable fix work units to the existing queue. Stop after updating the queue.
loop.sh:258:Review input: $review_file
loop.sh:271:        LOOP_REVIEW_FILE="$review_file" \
loop.sh:281:        LOOP_REVIEW_FILE="$review_file" \
loop.sh:298:max_review_rounds=2
loop.sh:299:review_enabled=0
loop.sh:314:    --review)
loop.sh:315:      review_enabled=1
loop.sh:318:    --max-review-rounds)
loop.sh:319:      [[ $# -ge 2 ]] || die "--max-review-rounds needs a value"
loop.sh:320:      max_review_rounds=$2
loop.sh:346:review_file="$queue_dir/REVIEW.md"
loop.sh:359:review_round=0
loop.sh:457:if [[ "$review_enabled" == 0 ]]; then
loop.sh:466:if (( review_round >= max_review_rounds )); then
loop.sh:467:  die "reached max review rounds ($max_review_rounds)"
loop.sh:470:review_round=$((review_round + 1))
loop.sh:471:review_prompt=$(mktemp)
loop.sh:472:review_out=$(mktemp)
loop.sh:473:write_review_prompt "$review_prompt"
loop.sh:474:review_cmd="${LOOP_REVIEW_CMD:-${LOOP_AGENT_CMD:-}}"
loop.sh:476:echo "knack: review round $review_round/$max_review_rounds"
loop.sh:478:run_phase_agent "review" "$review_prompt" "$review_out" "$review_cmd"
loop.sh:479:review_code=$?
loop.sh:481:if [[ $review_code -ne 0 ]]; then
loop.sh:482:  cat "$review_out"
loop.sh:483:  die "review worker exited nonzero"
loop.sh:485:[[ -f "$review_file" ]] || die "review worker did not write $review_file"
loop.sh:486:actionable=$(extract_actionable_count "$review_file") || die "review file has no actionable summary: $review_file"
loop.sh:488:if [[ "$actionable" == 0 ]]; then
loop.sh:489:  echo "knack: review clean"
loop.sh:493:if (( review_round >= max_review_rounds )); then
loop.sh:494:  die "review found $actionable actionable issue(s) at max review rounds ($max_review_rounds)"
loop.sh:502:echo "knack: fix round $review_round/$max_review_rounds — actionable: $actionable"
tests/run.sh:172:# Review-fix loop with fake build, review, and fix workers.
tests/run.sh:173:repo_review="$tmp/repo-review"
tests/run.sh:174:mkdir -p "$repo_review/.loop"
tests/run.sh:175:cat > "$repo_review/.loop/QUEUE.md" <<'EOF'
tests/run.sh:176:# Loop Queue: review cycle
tests/run.sh:179:Exercise build, review, fix, and review again.
tests/run.sh:182:The generated app is fixed and review reports no actionable issues.
tests/run.sh:184:## the initial build creates a reviewable app file
tests/run.sh:190:- Leave the bug for review to find.
tests/run.sh:203:cat > "$repo_review/build-worker.sh" <<'EOF'
tests/run.sh:214:chmod +x "$repo_review/build-worker.sh"
tests/run.sh:216:cat > "$repo_review/review-worker.sh" <<'EOF'
tests/run.sh:220:if [[ -f review-count.txt ]]; then
tests/run.sh:221:  count=$(cat review-count.txt)
tests/run.sh:224:printf '%s\n' "$count" > review-count.txt
tests/run.sh:226:actionable=1
tests/run.sh:228:  actionable=0
tests/run.sh:237:- actionable | high — app.txt must say fixed
tests/run.sh:242:## Summary
tests/run.sh:243:- actionable: $actionable
tests/run.sh:248:echo "review actionable: $actionable"
tests/run.sh:250:chmod +x "$repo_review/review-worker.sh"
tests/run.sh:252:cat > "$repo_review/fix-worker.sh" <<'EOF'
tests/run.sh:260:- .loop/REVIEW.md
tests/run.sh:277:chmod +x "$repo_review/fix-worker.sh"
tests/run.sh:279:LOOP_AGENT_CMD="$repo_review/build-worker.sh" \
tests/run.sh:280:  LOOP_REVIEW_CMD="$repo_review/review-worker.sh" \
tests/run.sh:281:  LOOP_FIX_CMD="$repo_review/fix-worker.sh" \
tests/run.sh:282:  "$root/loop.sh" run "$repo_review/.loop/QUEUE.md" --review --max-ticks 2 >/tmp/loop-review.txt
tests/run.sh:283:assert_contains "$repo_review/.loop/QUEUE.md" "## the fix unit repairs the bug"
tests/run.sh:284:assert_contains "$repo_review/.loop/QUEUE.md" "Status: done"
tests/run.sh:285:assert_contains "$repo_review/.loop/REVIEW.md" "- actionable: 0"
tests/run.sh:286:assert_contains "$repo_review/review-count.txt" "2"
tests/run.sh:287:assert_contains "$repo_review/app.txt" "fixed"
.agents/skills/fix/SKILL.md:3:description: Use when addressing review findings and feeding them back into the loop. Converts review findings into new verifiable work units in `.loop/<name>/QUEUE.md` and runs another loop pass. Triggers on "fix the review findings", "address the feedback", "rework based on review", "fix what review found", or when review produced findings that need to be resolved.
.agents/skills/fix/SKILL.md:8:Address review findings by triaging `REVIEW.md` and generating new work units that feed back into the loop. The `fix` skill is the bridge between review and another build pass; it owns the triage and the queue format. It appends to the existing `.loop/<name>/QUEUE.md` file for the same cycle.
.agents/skills/fix/SKILL.md:14:- `.loop/<name>/REVIEW.md` — structured findings from the `review` skill. Read it, but do not edit it.
.agents/skills/fix/SKILL.md:17:`REVIEW.md` is expected to include `## Standards`, `## Intent`, `## Speculative`, and `## Summary` sections. The summary's actionable count is the loop's signal; the finding entries are the fix skill's input. Use the finding classification, confidence, and evidence when deciding whether to generate a work unit.
.agents/skills/fix/SKILL.md:21:1. **Read `REVIEW.md`.** Each standards or intent finding is a candidate work unit. `review` has already classified findings as trivial / actionable / disputed / deferred; use those classifications or re-triage as needed. Treat speculative findings as notes for future explore/plan work unless the review explicitly classifies one as actionable.
.agents/skills/fix/SKILL.md:26:   - **Disputed** — the finding is wrong or the reviewer is being overly cautious. Note the disagreement in the summary, don't create a unit. Move on.
.agents/skills/fix/SKILL.md:29:3. **Create work units** for actionable findings. Use the standard work unit format from the `plan` skill. Each unit must have one observable outcome and one deterministic verify command. Reference the source finding in `Read first:` by pointing to `REVIEW.md` plus any evidence paths from the finding.
.agents/skills/fix/SKILL.md:35:- .loop/<name>/REVIEW.md (<finding id or heading>)
.agents/skills/fix/SKILL.md:55:4. **Append to `QUEUE.md`.** Append new units to the end of the existing `.loop/<name>/QUEUE.md`. Do not reorder existing units, do not edit existing unit bodies, and do not change any existing `Status:` lines. If there are no actionable findings, append nothing.
.agents/skills/fix/SKILL.md:61:- Not a re-review — the findings are already known. Fix turns them into work.
.agents/skills/fix/SKILL.md:62:- Not a debate — if a finding is disputed, note it and move on. Don't argue with the review in the queue.
.agents/skills/fix/SKILL.md:64:- Not the orchestrator — the loop invokes review, invokes fix, and runs the next build pass.
.agents/skills/fix/SKILL.md:68:Fix units must not break what review already approved. Each fix unit's `Done means:` should include a no-regression condition. The verify command should cover both the fix and the preservation of what was already working.
.agents/skills/fix/SKILL.md:73:- A summary of what was triaged (actionable / trivial / disputed / deferred)
.agents/skills/review/SKILL.md:2:name: review
.agents/skills/review/SKILL.md:3:description: Use when reviewing completed work units or a finished queue. Two-axis adversarial review — standards (does the change follow codebase conventions?) and intent (does the change do what the work unit said?). Reviews against the actual codebase, not specs. Triggers on "review", "check this", "is this right", "what did we miss", "stress-test the implementation", or when work needs adversarial scrutiny before being accepted.
.agents/skills/review/SKILL.md:8:Adversarial review of completed work. Two axes, run independently so neither pollutes the other:
.agents/skills/review/SKILL.md:15:## When to review
.agents/skills/review/SKILL.md:24:## Before you review
.agents/skills/review/SKILL.md:26:Read the work unit from `.loop/<name>/QUEUE.md` and the evidence from `.loop/<name>/EVIDENCE.md` for the unit you're reviewing. The evidence tells you what the verify command actually proved; the work unit tells you what was promised. Review against the actual codebase, not the specs.
.agents/skills/review/SKILL.md:28:## Two-axis review
.agents/skills/review/SKILL.md:53:The `Verify:` command is the mechanically enforceable subset of `Done means:`. The gap between `Done means:` and `Verify:` is the review surface: intent review checks what the verify command cannot.
.agents/skills/review/SKILL.md:55:> **Intent review is a judgment check, not a deterministic gate.** Be explicit about its scope and confidence. Any finding that becomes a new work unit must have a deterministic `Verify` command the runner can execute.
.agents/skills/review/SKILL.md:63:Review findings are not just notes — they are the input to the `fix` skill, which triages them and appends actionable ones as new work units in `.loop/<name>/QUEUE.md`.
.agents/skills/review/SKILL.md:65:- **Trivial** findings (a typo, a missing newline) can be fixed inline during review.
.agents/skills/review/SKILL.md:67:- **Disputed** or **deferred** findings are recorded in the review summary but not turned into units.
.agents/skills/review/SKILL.md:69:The output of review is a findings summary, not a queue edit.
.agents/skills/review/SKILL.md:71:## What review is not
.agents/skills/review/SKILL.md:75:- Not a gate — the loop doesn't enforce review. It's invoked when the user wants it.
.agents/skills/review/SKILL.md:84:Classify each finding as trivial / actionable / disputed / deferred.
.agents/skills/review/SKILL.md:87:- Hand actionable findings to the `fix` skill, which will triage and append work units to `.loop/<name>/QUEUE.md`, or
cli/internal/skills/skills_test.go:20:	expected := []string{"build", "decide", "domain-modeling", "explore", "fix", "plan", "review"}
.agents/skills/domain-modeling/SKILL.md:3:description: Use when domain terms surface during explore, plan, or review and need to be defined, challenged, or stress-tested against the project's ubiquitous language. Manages `glossary.md` — the project's shared vocabulary. Triggers on "what does X mean here", "are these the same concept", "let's define our terms", "that term is ambiguous", or when a domain concept is used inconsistently across a conversation or codebase.
.agents/skills/domain-modeling/SKILL.md:10:This is a shared skill — called by explore, plan, and review. It is not a separate phase.
.agents/skills/decide/SKILL.md:3:description: Use when an architectural ruling crystallizes during any phase — explore, plan, build, or review. Captures decisions as ADRs in `decisions/` so they persist after specs are deleted and code moves on. Triggers on "we decided to", "let's go with", "the ruling is", "why did we choose", "record this decision", or when a design choice with lasting consequences is made.
.agents/skills/decide/SKILL.md:10:This is a shared skill — called by explore, plan, build, and review whenever a ruling crystallizes. It is not a separate phase.
.agents/skills/decide/SKILL.md:52:1. Recognize that a ruling just crystallized. If you're in explore, plan, build, or review and you just resolved a tension with lasting consequences, that's a decision.
cli/embedded/skills/domain-modeling/SKILL.md:3:description: Use when domain terms surface during explore, plan, or review and need to be defined, challenged, or stress-tested against the project's ubiquitous language. Manages `glossary.md` — the project's shared vocabulary. Triggers on "what does X mean here", "are these the same concept", "let's define our terms", "that term is ambiguous", or when a domain concept is used inconsistently across a conversation or codebase.
cli/embedded/skills/domain-modeling/SKILL.md:10:This is a shared skill — called by explore, plan, and review. It is not a separate phase.
cli/embedded/skills/decide/SKILL.md:3:description: Use when an architectural ruling crystallizes during any phase — explore, plan, build, or review. Captures decisions as ADRs in `decisions/` so they persist after specs are deleted and code moves on. Triggers on "we decided to", "let's go with", "the ruling is", "why did we choose", "record this decision", or when a design choice with lasting consequences is made.
cli/embedded/skills/decide/SKILL.md:10:This is a shared skill — called by explore, plan, build, and review whenever a ruling crystallizes. It is not a separate phase.
cli/embedded/skills/decide/SKILL.md:52:1. Recognize that a ruling just crystallized. If you're in explore, plan, build, or review and you just resolved a tension with lasting consequences, that's a decision.
cli/embedded/skills/fix/SKILL.md:3:description: Use when addressing review findings and feeding them back into the loop. Converts review findings into new verifiable work units in `.loop/<name>/QUEUE.md` and runs another loop pass. Triggers on "fix the review findings", "address the feedback", "rework based on review", "fix what review found", or when review produced findings that need to be resolved.
cli/embedded/skills/fix/SKILL.md:8:Address review findings by triaging `REVIEW.md` and generating new work units that feed back into the loop. The `fix` skill is the bridge between review and another build pass; it owns the triage and the queue format. It appends to the existing `.loop/<name>/QUEUE.md` file for the same cycle.
cli/embedded/skills/fix/SKILL.md:14:- `.loop/<name>/REVIEW.md` — structured findings from the `review` skill. Read it, but do not edit it.
cli/embedded/skills/fix/SKILL.md:17:`REVIEW.md` is expected to include `## Standards`, `## Intent`, `## Speculative`, and `## Summary` sections. The summary's actionable count is the loop's signal; the finding entries are the fix skill's input. Use the finding classification, confidence, and evidence when deciding whether to generate a work unit.
cli/embedded/skills/fix/SKILL.md:21:1. **Read `REVIEW.md`.** Each standards or intent finding is a candidate work unit. `review` has already classified findings as trivial / actionable / disputed / deferred; use those classifications or re-triage as needed. Treat speculative findings as notes for future explore/plan work unless the review explicitly classifies one as actionable.
cli/embedded/skills/fix/SKILL.md:26:   - **Disputed** — the finding is wrong or the reviewer is being overly cautious. Note the disagreement in the summary, don't create a unit. Move on.
cli/embedded/skills/fix/SKILL.md:29:3. **Create work units** for actionable findings. Use the standard work unit format from the `plan` skill. Each unit must have one observable outcome and one deterministic verify command. Reference the source finding in `Read first:` by pointing to `REVIEW.md` plus any evidence paths from the finding.
cli/embedded/skills/fix/SKILL.md:35:- .loop/<name>/REVIEW.md (<finding id or heading>)
cli/embedded/skills/fix/SKILL.md:55:4. **Append to `QUEUE.md`.** Append new units to the end of the existing `.loop/<name>/QUEUE.md`. Do not reorder existing units, do not edit existing unit bodies, and do not change any existing `Status:` lines. If there are no actionable findings, append nothing.
cli/embedded/skills/fix/SKILL.md:61:- Not a re-review — the findings are already known. Fix turns them into work.
cli/embedded/skills/fix/SKILL.md:62:- Not a debate — if a finding is disputed, note it and move on. Don't argue with the review in the queue.
cli/embedded/skills/fix/SKILL.md:64:- Not the orchestrator — the loop invokes review, invokes fix, and runs the next build pass.
cli/embedded/skills/fix/SKILL.md:68:Fix units must not break what review already approved. Each fix unit's `Done means:` should include a no-regression condition. The verify command should cover both the fix and the preservation of what was already working.
cli/embedded/skills/fix/SKILL.md:73:- A summary of what was triaged (actionable / trivial / disputed / deferred)

codex
There are unrelated in-progress changes already in the tree, including loop and fix-skill work from earlier units. I’ll leave those alone and touch only the review skill plus its embedded mirror.
exec
/usr/bin/bash -lc "sed -n '1,220p' prompts/reviewer.md" in /home/daniel/build/knack
 exited 2 in 0ms:
sed: can't read prompts/reviewer.md: No such file or directory

exec
/usr/bin/bash -lc "sed -n '190,250p' loop.sh" in /home/daniel/build/knack
 succeeded in 0ms:
    echo "Worker output:"
    echo '````text'
    cat "$agent_out"
    echo '````'
    echo
    echo "What this proves:"
    if [[ "$status" == "done" ]]; then
      echo "- The verify command passed for this work unit in the current repo state."
    else
      echo "- The work unit is not externally verified."
    fi
    echo
    echo "What remains unverified:"
    echo "- Anything outside the verify command's proof scope."
  } >> "$evidence"
}

extract_actionable_count() {
  awk '
    /^[[:space:]]*-[[:space:]]*actionable:[[:space:]]*[0-9]+[[:space:]]*$/ {
      sub(/^[[:space:]]*-[[:space:]]*actionable:[[:space:]]*/, "")
      sub(/[[:space:]]*$/, "")
      print
      found = 1
      exit
    }
    END { if (!found) exit 1 }
  ' "$1"
}

write_review_prompt() {
  local out=$1
  local template="$script_dir/prompts/reviewer.md"
  if [[ -f "$template" ]]; then
    cat "$template" > "$out"
  else
    cat > "$out" <<'EOF'
# Knack Reviewer

Load and follow the **review** skill in `.agents/skills/review/`.
Review the completed queue against the current repository state. Write the structured review artifact at the path provided below.
EOF
  fi
  cat >> "$out" <<EOF

Queue: $queue_abs
Evidence: $evidence
Review output: $review_file
EOF
}

write_fix_prompt() {
  local out=$1
  local template="$script_dir/prompts/fixer.md"
  if [[ -f "$template" ]]; then
    cat "$template" > "$out"
  else
    cat > "$out" <<'EOF'
# Knack Fixer

Load and follow the **fix** skill in `.agents/skills/fix/`.

codex
I’m editing the output contract now: the substance of the review stays the same, but the artifact shape will match what `loop.sh` and `fix` expect.
apply patch
patch: completed
/home/daniel/build/knack/.agents/skills/review/SKILL.md
diff --git a/.agents/skills/review/SKILL.md b/.agents/skills/review/SKILL.md
index 01f38c754c076bd9aacfb31ef10cd0d2627bfb2c..c0ae6eac9010e35d7c3a8272cb2301645aaee0d3
--- a/.agents/skills/review/SKILL.md
+++ b/.agents/skills/review/SKILL.md
@@ -60,13 +60,13 @@
 
 ## Findings become input to the fix skill
 
-Review findings are not just notes — they are the input to the `fix` skill, which triages them and appends actionable ones as new work units in `.loop/<name>/QUEUE.md`.
+Review findings are not just notes — they are written to `.loop/<name>/REVIEW.md` as the input to the `fix` skill, which triages them and appends actionable ones as new work units in `.loop/<name>/QUEUE.md`.
 
 - **Trivial** findings (a typo, a missing newline) can be fixed inline during review.
 - **Actionable** findings are handed to the `fix` skill. Do not write the work units yourself; `fix` owns the triage and formatting.
 - **Disputed** or **deferred** findings are recorded in the review summary but not turned into units.
 
-The output of review is a findings summary, not a queue edit.
+The output of review is a structured review artifact, not a queue edit.
 
 ## What review is not
 
@@ -76,13 +76,46 @@
 
 ## Output
 
-Summarize findings per axis:
+Write the structured review artifact to the requested review output path. In the loop this is `.loop/<name>/REVIEW.md`; if the prompt provides a different `Review output:` path, write there.
+
+`REVIEW.md` must have exactly these top-level sections:
+
+1. `## Standards`
+2. `## Intent`
+3. `## Speculative`
+4. `## Summary`
 
-- **Standards**: N findings (list them, or "no issues found")
-- **Intent**: N findings (list them, or "the change matches the unit's stated outcome")
+Put standards-axis findings under `## Standards` and intent-axis findings under `## Intent`. Use `## Speculative` only for concerns that are plausible but not grounded enough to become a standards or intent finding. If a section is clean, write `No issues found.` under that section.
 
-Classify each finding as trivial / actionable / disputed / deferred.
+Each finding must include:
 
-Then either:
-- Hand actionable findings to the `fix` skill, which will triage and append work units to `.loop/<name>/QUEUE.md`, or
-- Report "no action needed" if the work is clean.
+- A stable id, such as `S1`, `I1`, or `X1`
+- Classification: `trivial`, `actionable`, `disputed`, or `deferred`
+- Confidence: `high`, `medium`, or `low`
+- Evidence: a `path/to/file:line` reference or a short quoted code excerpt
+- Finding: the issue in one or two sentences
+- Fix direction: the smallest useful direction for the `fix` skill, or `None` for non-actionable findings
+
+Use this finding shape:
+
+```markdown
+- S1 | actionable | high
+  Evidence: `path/to/file:42`
+  Finding: The change violates the repo's existing queue parser behavior.
+  Fix direction: Align the parser with the shell loop's unit-header rules.
+```
+
+The `## Summary` section must include counts using this machine-readable shape:
+
+```markdown
+## Summary
+- standards: 1
+- intent: 0
+- speculative: 0
+- actionable: 1
+- trivial: 0
+- disputed: 0
+- deferred: 0
+```
+
+The `- actionable: N` line is the loop's continue/stop signal. Count only findings classified as `actionable`, including actionable speculative findings if you intentionally create one. If there are no actionable findings, write `- actionable: 0`.

diff --git a/.agents/skills/review/SKILL.md b/.agents/skills/review/SKILL.md
index 01f38c754c076bd9aacfb31ef10cd0d2627bfb2c..c0ae6eac9010e35d7c3a8272cb2301645aaee0d3
--- a/.agents/skills/review/SKILL.md
+++ b/.agents/skills/review/SKILL.md
@@ -60,13 +60,13 @@
 
 ## Findings become input to the fix skill
 
-Review findings are not just notes — they are the input to the `fix` skill, which triages them and appends actionable ones as new work units in `.loop/<name>/QUEUE.md`.
+Review findings are not just notes — they are written to `.loop/<name>/REVIEW.md` as the input to the `fix` skill, which triages them and appends actionable ones as new work units in `.loop/<name>/QUEUE.md`.
 
 - **Trivial** findings (a typo, a missing newline) can be fixed inline during review.
 - **Actionable** findings are handed to the `fix` skill. Do not write the work units yourself; `fix` owns the triage and formatting.
 - **Disputed** or **deferred** findings are recorded in the review summary but not turned into units.
 
-The output of review is a findings summary, not a queue edit.
+The output of review is a structured review artifact, not a queue edit.
 
 ## What review is not
 
@@ -76,13 +76,46 @@
 
 ## Output
 
-Summarize findings per axis:
+Write the structured review artifact to the requested review output path. In the loop this is `.loop/<name>/REVIEW.md`; if the prompt provides a different `Review output:` path, write there.
+
+`REVIEW.md` must have exactly these top-level sections:
+
+1. `## Standards`
+2. `## Intent`
+3. `## Speculative`
+4. `## Summary`
 
-- **Standards**: N findings (list them, or "no issues found")
-- **Intent**: N findings (list them, or "the change matches the unit's stated outcome")
+Put standards-axis findings under `## Standards` and intent-axis findings under `## Intent`. Use `## Speculative` only for concerns that are plausible but not grounded enough to become a standards or intent finding. If a section is clean, write `No issues found.` under that section.
 
-Classify each finding as trivial / actionable / disputed / deferred.
+Each finding must include:
 
-Then either:
-- Hand actionable findings to the `fix` skill, which will triage and append work units to `.loop/<name>/QUEUE.md`, or
-- Report "no action needed" if the work is clean.
+- A stable id, such as `S1`, `I1`, or `X1`
+- Classification: `trivial`, `actionable`, `disputed`, or `deferred`
+- Confidence: `high`, `medium`, or `low`
+- Evidence: a `path/to/file:line` reference or a short quoted code excerpt
+- Finding: the issue in one or two sentences
+- Fix direction: the smallest useful direction for the `fix` skill, or `None` for non-actionable findings
+
+Use this finding shape:
+
+```markdown
+- S1 | actionable | high
+  Evidence: `path/to/file:42`
+  Finding: The change violates the repo's existing queue parser behavior.
+  Fix direction: Align the parser with the shell loop's unit-header rules.
+```
+
+The `## Summary` section must include counts using this machine-readable shape:
+
+```markdown
+## Summary
+- standards: 1
+- intent: 0
+- speculative: 0
+- actionable: 1
+- trivial: 0
+- disputed: 0
+- deferred: 0
+```
+
+The `- actionable: N` line is the loop's continue/stop signal. Count only findings classified as `actionable`, including actionable speculative findings if you intentionally create one. If there are no actionable findings, write `- actionable: 0`.

codex
The source skill now names `REVIEW.md`, required sections, finding metadata, and the exact `- actionable: N` signal. I’m syncing that single skill into the embedded copy and then running the unit’s verify command.
exec
/usr/bin/bash -lc 'cp .agents/skills/review/SKILL.md cli/embedded/skills/review/SKILL.md' in /home/daniel/build/knack
 succeeded in 0ms:
exec
/usr/bin/bash -lc 'git diff -- .agents/skills/review/SKILL.md' in /home/daniel/build/knack
 succeeded in 0ms:
diff --git a/.agents/skills/review/SKILL.md b/.agents/skills/review/SKILL.md
index 01f38c7..c0ae6ea 100644
--- a/.agents/skills/review/SKILL.md
+++ b/.agents/skills/review/SKILL.md
@@ -60,13 +60,13 @@ Run both axes. They can be parallel (two passes over the same diff) or sequentia
 
 ## Findings become input to the fix skill
 
-Review findings are not just notes — they are the input to the `fix` skill, which triages them and appends actionable ones as new work units in `.loop/<name>/QUEUE.md`.
+Review findings are not just notes — they are written to `.loop/<name>/REVIEW.md` as the input to the `fix` skill, which triages them and appends actionable ones as new work units in `.loop/<name>/QUEUE.md`.
 
 - **Trivial** findings (a typo, a missing newline) can be fixed inline during review.
 - **Actionable** findings are handed to the `fix` skill. Do not write the work units yourself; `fix` owns the triage and formatting.
 - **Disputed** or **deferred** findings are recorded in the review summary but not turned into units.
 
-The output of review is a findings summary, not a queue edit.
+The output of review is a structured review artifact, not a queue edit.
 
 ## What review is not
 
@@ -76,13 +76,46 @@ The output of review is a findings summary, not a queue edit.
 
 ## Output
 
-Summarize findings per axis:
+Write the structured review artifact to the requested review output path. In the loop this is `.loop/<name>/REVIEW.md`; if the prompt provides a different `Review output:` path, write there.
 
-- **Standards**: N findings (list them, or "no issues found")
-- **Intent**: N findings (list them, or "the change matches the unit's stated outcome")
+`REVIEW.md` must have exactly these top-level sections:
 
-Classify each finding as trivial / actionable / disputed / deferred.
+1. `## Standards`
+2. `## Intent`
+3. `## Speculative`
+4. `## Summary`
 
-Then either:
-- Hand actionable findings to the `fix` skill, which will triage and append work units to `.loop/<name>/QUEUE.md`, or
-- Report "no action needed" if the work is clean.
+Put standards-axis findings under `## Standards` and intent-axis findings under `## Intent`. Use `## Speculative` only for concerns that are plausible but not grounded enough to become a standards or intent finding. If a section is clean, write `No issues found.` under that section.
+
+Each finding must include:
+
+- A stable id, such as `S1`, `I1`, or `X1`
+- Classification: `trivial`, `actionable`, `disputed`, or `deferred`
+- Confidence: `high`, `medium`, or `low`
+- Evidence: a `path/to/file:line` reference or a short quoted code excerpt
+- Finding: the issue in one or two sentences
+- Fix direction: the smallest useful direction for the `fix` skill, or `None` for non-actionable findings
+
+Use this finding shape:
+
+```markdown
+- S1 | actionable | high
+  Evidence: `path/to/file:42`
+  Finding: The change violates the repo's existing queue parser behavior.
+  Fix direction: Align the parser with the shell loop's unit-header rules.
+```
+
+The `## Summary` section must include counts using this machine-readable shape:
+
+```markdown
+## Summary
+- standards: 1
+- intent: 0
+- speculative: 0
+- actionable: 1
+- trivial: 0
+- disputed: 0
+- deferred: 0
+```
+
+The `- actionable: N` line is the loop's continue/stop signal. Count only findings classified as `actionable`, including actionable speculative findings if you intentionally create one. If there are no actionable findings, write `- actionable: 0`.

diff --git a/.agents/skills/review/SKILL.md b/.agents/skills/review/SKILL.md
index 01f38c754c076bd9aacfb31ef10cd0d2627bfb2c..c0ae6eac9010e35d7c3a8272cb2301645aaee0d3
--- a/.agents/skills/review/SKILL.md
+++ b/.agents/skills/review/SKILL.md
@@ -60,13 +60,13 @@
 
 ## Findings become input to the fix skill
 
-Review findings are not just notes — they are the input to the `fix` skill, which triages them and appends actionable ones as new work units in `.loop/<name>/QUEUE.md`.
+Review findings are not just notes — they are written to `.loop/<name>/REVIEW.md` as the input to the `fix` skill, which triages them and appends actionable ones as new work units in `.loop/<name>/QUEUE.md`.
 
 - **Trivial** findings (a typo, a missing newline) can be fixed inline during review.
 - **Actionable** findings are handed to the `fix` skill. Do not write the work units yourself; `fix` owns the triage and formatting.
 - **Disputed** or **deferred** findings are recorded in the review summary but not turned into units.
 
-The output of review is a findings summary, not a queue edit.
+The output of review is a structured review artifact, not a queue edit.
 
 ## What review is not
 
@@ -76,13 +76,46 @@
 
 ## Output
 
-Summarize findings per axis:
+Write the structured review artifact to the requested review output path. In the loop this is `.loop/<name>/REVIEW.md`; if the prompt provides a different `Review output:` path, write there.
+
+`REVIEW.md` must have exactly these top-level sections:
+
+1. `## Standards`
+2. `## Intent`
+3. `## Speculative`
+4. `## Summary`
 
-- **Standards**: N findings (list them, or "no issues found")
-- **Intent**: N findings (list them, or "the change matches the unit's stated outcome")
+Put standards-axis findings under `## Standards` and intent-axis findings under `## Intent`. Use `## Speculative` only for concerns that are plausible but not grounded enough to become a standards or intent finding. If a section is clean, write `No issues found.` under that section.
 
-Classify each finding as trivial / actionable / disputed / deferred.
+Each finding must include:
 
-Then either:
-- Hand actionable findings to the `fix` skill, which will triage and append work units to `.loop/<name>/QUEUE.md`, or
-- Report "no action needed" if the work is clean.
+- A stable id, such as `S1`, `I1`, or `X1`
+- Classification: `trivial`, `actionable`, `disputed`, or `deferred`
+- Confidence: `high`, `medium`, or `low`
+- Evidence: a `path/to/file:line` reference or a short quoted code excerpt
+- Finding: the issue in one or two sentences
+- Fix direction: the smallest useful direction for the `fix` skill, or `None` for non-actionable findings
+
+Use this finding shape:
+
+```markdown
+- S1 | actionable | high
+  Evidence: `path/to/file:42`
+  Finding: The change violates the repo's existing queue parser behavior.
+  Fix direction: Align the parser with the shell loop's unit-header rules.
+```
+
+The `## Summary` section must include counts using this machine-readable shape:
+
+```markdown
+## Summary
+- standards: 1
+- intent: 0
+- speculative: 0
+- actionable: 1
+- trivial: 0
+- disputed: 0
+- deferred: 0
+```
+
+The `- actionable: N` line is the loop's continue/stop signal. Count only findings classified as `actionable`, including actionable speculative findings if you intentionally create one. If there are no actionable findings, write `- actionable: 0`.

exec
/usr/bin/bash -lc 'cd /home/daniel/build/knack && diff -r .agents/skills/review cli/embedded/skills/review && ./tests/run.sh' in /home/daniel/build/knack
diff --git a/.agents/skills/review/SKILL.md b/.agents/skills/review/SKILL.md
index 01f38c754c076bd9aacfb31ef10cd0d2627bfb2c..c0ae6eac9010e35d7c3a8272cb2301645aaee0d3
--- a/.agents/skills/review/SKILL.md
+++ b/.agents/skills/review/SKILL.md
@@ -60,13 +60,13 @@
 
 ## Findings become input to the fix skill
 
-Review findings are not just notes — they are the input to the `fix` skill, which triages them and appends actionable ones as new work units in `.loop/<name>/QUEUE.md`.
+Review findings are not just notes — they are written to `.loop/<name>/REVIEW.md` as the input to the `fix` skill, which triages them and appends actionable ones as new work units in `.loop/<name>/QUEUE.md`.
 
 - **Trivial** findings (a typo, a missing newline) can be fixed inline during review.
 - **Actionable** findings are handed to the `fix` skill. Do not write the work units yourself; `fix` owns the triage and formatting.
 - **Disputed** or **deferred** findings are recorded in the review summary but not turned into units.
 
-The output of review is a findings summary, not a queue edit.
+The output of review is a structured review artifact, not a queue edit.
 
 ## What review is not
 
@@ -76,13 +76,46 @@
 
 ## Output
 
-Summarize findings per axis:
+Write the structured review artifact to the requested review output path. In the loop this is `.loop/<name>/REVIEW.md`; if the prompt provides a different `Review output:` path, write there.
+
+`REVIEW.md` must have exactly these top-level sections:
+
+1. `## Standards`
+2. `## Intent`
+3. `## Speculative`
+4. `## Summary`
 
-- **Standards**: N findings (list them, or "no issues found")
-- **Intent**: N findings (list them, or "the change matches the unit's stated outcome")
+Put standards-axis findings under `## Standards` and intent-axis findings under `## Intent`. Use `## Speculative` only for concerns that are plausible but not grounded enough to become a standards or intent finding. If a section is clean, write `No issues found.` under that section.
 
-Classify each finding as trivial / actionable / disputed / deferred.
+Each finding must include:
 
-Then either:
-- Hand actionable findings to the `fix` skill, which will triage and append work units to `.loop/<name>/QUEUE.md`, or
-- Report "no action needed" if the work is clean.
+- A stable id, such as `S1`, `I1`, or `X1`
+- Classification: `trivial`, `actionable`, `disputed`, or `deferred`
+- Confidence: `high`, `medium`, or `low`
+- Evidence: a `path/to/file:line` reference or a short quoted code excerpt
+- Finding: the issue in one or two sentences
+- Fix direction: the smallest useful direction for the `fix` skill, or `None` for non-actionable findings
+
+Use this finding shape:
+
+```markdown
+- S1 | actionable | high
+  Evidence: `path/to/file:42`
+  Finding: The change violates the repo's existing queue parser behavior.
+  Fix direction: Align the parser with the shell loop's unit-header rules.
+```
+
+The `## Summary` section must include counts using this machine-readable shape:
+
+```markdown
+## Summary
+- standards: 1
+- intent: 0
+- speculative: 0
+- actionable: 1
+- trivial: 0
+- disputed: 0
+- deferred: 0
+```
+
+The `- actionable: N` line is the loop's continue/stop signal. Count only findings classified as `actionable`, including actionable speculative findings if you intentionally create one. If there are no actionable findings, write `- actionable: 0`.

 succeeded in 3717ms:
Valid skill: /home/daniel/build/knack/.agents/skills/build
Valid skill: /home/daniel/build/knack/.agents/skills/decide
Valid skill: /home/daniel/build/knack/.agents/skills/domain-modeling
Valid skill: /home/daniel/build/knack/.agents/skills/explore
Valid skill: /home/daniel/build/knack/.agents/skills/fix
Valid skill: /home/daniel/build/knack/.agents/skills/plan
Valid skill: /home/daniel/build/knack/.agents/skills/review
knack tests passed

codex
```text
Unit: review skill writes a structured REVIEW.md
Changed: .agents/skills/review/SKILL.md; cli/embedded/skills/review/SKILL.md
Verify expected: cd /home/daniel/build/knack && diff -r .agents/skills/review cli/embedded/skills/review && ./tests/run.sh
Notes: No blockers.
```
diff --git a/.agents/skills/review/SKILL.md b/.agents/skills/review/SKILL.md
index 01f38c754c076bd9aacfb31ef10cd0d2627bfb2c..c0ae6eac9010e35d7c3a8272cb2301645aaee0d3
--- a/.agents/skills/review/SKILL.md
+++ b/.agents/skills/review/SKILL.md
@@ -60,13 +60,13 @@
 
 ## Findings become input to the fix skill
 
-Review findings are not just notes — they are the input to the `fix` skill, which triages them and appends actionable ones as new work units in `.loop/<name>/QUEUE.md`.
+Review findings are not just notes — they are written to `.loop/<name>/REVIEW.md` as the input to the `fix` skill, which triages them and appends actionable ones as new work units in `.loop/<name>/QUEUE.md`.
 
 - **Trivial** findings (a typo, a missing newline) can be fixed inline during review.
 - **Actionable** findings are handed to the `fix` skill. Do not write the work units yourself; `fix` owns the triage and formatting.
 - **Disputed** or **deferred** findings are recorded in the review summary but not turned into units.
 
-The output of review is a findings summary, not a queue edit.
+The output of review is a structured review artifact, not a queue edit.
 
 ## What review is not
 
@@ -76,13 +76,46 @@
 
 ## Output
 
-Summarize findings per axis:
+Write the structured review artifact to the requested review output path. In the loop this is `.loop/<name>/REVIEW.md`; if the prompt provides a different `Review output:` path, write there.
+
+`REVIEW.md` must have exactly these top-level sections:
+
+1. `## Standards`
+2. `## Intent`
+3. `## Speculative`
+4. `## Summary`
 
-- **Standards**: N findings (list them, or "no issues found")
-- **Intent**: N findings (list them, or "the change matches the unit's stated outcome")
+Put standards-axis findings under `## Standards` and intent-axis findings under `## Intent`. Use `## Speculative` only for concerns that are plausible but not grounded enough to become a standards or intent finding. If a section is clean, write `No issues found.` under that section.
 
-Classify each finding as trivial / actionable / disputed / deferred.
+Each finding must include:
 
-Then either:
-- Hand actionable findings to the `fix` skill, which will triage and append work units to `.loop/<name>/QUEUE.md`, or
-- Report "no action needed" if the work is clean.
+- A stable id, such as `S1`, `I1`, or `X1`
+- Classification: `trivial`, `actionable`, `disputed`, or `deferred`
+- Confidence: `high`, `medium`, or `low`
+- Evidence: a `path/to/file:line` reference or a short quoted code excerpt
+- Finding: the issue in one or two sentences
+- Fix direction: the smallest useful direction for the `fix` skill, or `None` for non-actionable findings
+
+Use this finding shape:
+
+```markdown
+- S1 | actionable | high
+  Evidence: `path/to/file:42`
+  Finding: The change violates the repo's existing queue parser behavior.
+  Fix direction: Align the parser with the shell loop's unit-header rules.
+```
+
+The `## Summary` section must include counts using this machine-readable shape:
+
+```markdown
+## Summary
+- standards: 1
+- intent: 0
+- speculative: 0
+- actionable: 1
+- trivial: 0
+- disputed: 0
+- deferred: 0
+```
+
+The `- actionable: N` line is the loop's continue/stop signal. Count only findings classified as `actionable`, including actionable speculative findings if you intentionally create one. If there are no actionable findings, write `- actionable: 0`.

tokens used
21,885
```text
Unit: review skill writes a structured REVIEW.md
Changed: .agents/skills/review/SKILL.md; cli/embedded/skills/review/SKILL.md
Verify expected: cd /home/daniel/build/knack && diff -r .agents/skills/review cli/embedded/skills/review && ./tests/run.sh
Notes: No blockers.
```
````

What this proves:
- The verify command passed for this work unit in the current repo state.

What remains unverified:
- Anything outside the verify command's proof scope.
