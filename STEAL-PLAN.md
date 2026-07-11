# Plan: Steal five patterns into knack + review-fix subloop

Five steals from gstack/OpenSpec, plus the larger review-fix subloop orchestration. Executed directly (not dogfooded through the loop). Each becomes a work unit that leaves the repo better if it's the last one landed.

**Verification for every unit:** `./tests/run.sh && cd cli && go test ./...` plus `diff -r .agents/skills cli/embedded/skills` after any skill edit (then run `cli/sync-skills.sh` if drifted).

## Full shape

This is the current map of `knack` — what exists and what is missing. Use it to decide what to work on next.

### Loop (`loop.sh`)

- [x] Build pass: read `QUEUE.md`, invoke worker, run `Verify:`, mark `done`, retry on no-progress
- [x] Agent-agnostic invocation: `LOOP_AGENT_CMD`, `Agent:` override, `LOOP_PROMPT_FILE`
- [x] Default `pi` fallback with `--approve`
- [x] `HANDOFF.md` and `EVIDENCE.md` generation
- [x] Review-fix subloop: after build queue drains, run review, generate fix units, re-build, re-review, bounded by `--max-review-rounds` and `--max-ticks`

### CLI (`knack`)

- [x] `skills init|check`
- [x] `validate <queue-file>`
- [x] `decisions list|show|check`
- [x] `status`
- [x] `glossary check`
- [x] `instructions <template>`
- [ ] `view` / `list` — show next pending unit and all units with status

### Skills

- [x] `explore`
- [x] `plan`
- [x] `build`
- [x] `review`
- [x] `fix`
- [x] `decide`
- [x] `domain-modeling`
- [x] `review` and `fix` loop-orchestrated integration: review writes `REVIEW.md`, fix consumes it

### Durable artifacts

- [x] `AGENTS.md`
- [x] `decisions/` (7 ADRs)
- [x] `glossary.md`
- [x] `.agents/skills/`
- [x] `LEARNINGS.md` (STEAL-PLAN Unit 4)

### Agent support

- [x] `pi` (`zai/glm-5.1`) — tested end-to-end
- [x] `opencode` (`opencode/hy3-free`) — tested end-to-end
- [x] `devin` — tested earlier per `AGENTS.md` lesson
- [ ] `claude` — example in README/docs, not loop-tested
- [ ] `codex` — example in README/docs, not loop-tested

### Patterns to steal

- [x] Unit 1 — Decision supersede-chaining + active-decisions surfacing (CLI)
- [x] Unit 2 — Decision supersede-chaining (decide skill + ADR template)
- [x] Unit 3 — Confidence calibration + pre-emit verification gate (review skill)
- [x] Unit 4 — `LEARNINGS.md` as a durable artifact
- [x] Unit 5 — "Explore first" as positioning
- [x] Unit 6 — Review-fix subloop orchestration

---

## Unit 1 — Decision supersede-chaining + active-decisions surfacing (CLI)

**The steal:** ADR lifecycle — supersede chains and active/superseded status — enforced mechanically.

**Files:**
- `cli/internal/decisions/decisions.go`
- `cli/internal/decisions/decisions_test.go`
- `cli/main.go` (the `decisions list` / `status` output)

**Changes:**

1. **Extend the `ADR` struct** (`decisions.go:17-23`):
   ```go
   type ADR struct {
       Number       string
       Title        string
       Status       string
       Filename     string
       Grandfather  bool
       SupersededBy string // parsed from "Superseded by: ADR-NNNN" (empty if none)
       Supersedes   string // parsed from "Supersedes: ADR-NNNN" (empty if none)
   }
   ```

2. **Add two line-prefix parsers** following the existing `parseADRStatus` / `hasGrandfatherLine` idiom (`decisions.go:74-93`):
   - `parseSupersededBy(contents string) string` — matches `Superseded by:` prefix, returns the trimmed value (e.g. `ADR-0007`).
   - `parseSupersedes(contents string) string` — matches `Supersedes:` prefix.
   Both tolerate the bare number or `ADR-NNNN` form (strip an `ADR-` prefix if present, zero-pad to 4 digits via the existing `canonicalADRNumber`).

3. **Wire them into `List`** (`decisions.go:51-57`) — populate the new fields when building each `ADR`.

4. **Add an `Active()` predicate method** (or a package-level helper):
   ```go
   func (a ADR) Active() bool {
       s := strings.ToLower(strings.TrimSpace(a.Status))
       return s == "" || s == "accepted" || s == "proposed"
   }
   ```
   (Superseded / deprecated / rejected are inactive. Empty defaults to active to preserve current behavior — existing ADRs with no status aren't penalized.)

5. **Make `Check` skip superseded ADRs** (`decisions.go:152-159`), the same way it skips `Grandfather`. Add `if !adr.Active() { continue }` alongside the existing grandfather skip. Rationale: a superseded ADR's replacement carries the coverage; flagging the old one as orphaned is a false positive.

6. **Add chain-integrity validation** to `Check` — a new mechanical gate:
   - If `adr.SupersededBy != ""` but no ADR with that number exists → finding: `"broken supersede chain: ADR %s references missing %s"`.
   - If `adr.SupersededBy != ""` but the target's `Supersedes` doesn't point back → finding: `"one-sided supersede: ADR %s says superseded by %s, but %s does not reference it"`.
   - If `adr.Supersedes != ""` but no ADR with that number exists → same broken-chain finding.
   This is the mechanical equivalent of the dangling-reference check that already exists.

7. **Surface active/superseded in `decisions list`** (`main.go:182-191`): change the format line so superseded ADRs are visually marked — `fmt.Printf("%s: %s (%s)\n", ...)` becomes one that appends ` [superseded by %s]` when `SupersededBy` is non-empty. Active ADRs print unchanged.

8. **Surface an active count in `status`** (`main.go:232`): change `fmt.Printf("adrs: %d\n", r.ADRs)` to print active vs total, e.g. `adrs: %d active (%d total)`. Requires `status.Generate` to compute the active count — it already calls `decisions.List`, so add `r.ActiveADRs` computed by counting `Active()`.

**Tests** (`decisions_test.go`):
- `TestListParsesSupersedeChain` — an ADR with `Superseded by: ADR-0002` and one with `Supersedes: ADR-0001`; assert fields populate.
- `TestCheckSkipsSuperseded` — a superseded ADR unreferenced by any loop body produces no orphan finding.
- `TestCheckBrokenSupersedeChain` — `Superseded by: ADR-0099` (missing) → finding.
- `TestCheckOneSidedSupersede` — A points at B, B doesn't point back → finding.
- `TestActivePredicate` — table test over statuses.

**Verify:**
```bash
cd cli && go test ./... && go vet ./... && cd .. && ./tests/run.sh
```

---

## Unit 2 — Decision supersede-chaining (decide skill + ADR template)

**The steal:** The `decide` skill already mentions superseding in prose; make it a documented, mechanical convention with the chain fields in the template. Pairs with Unit 1 (same feature, skill side).

**Files:**
- `.agents/skills/decide/SKILL.md`
- `cli/internal/instructions/instructions.go` (`adrTemplate`)
- `cli/embedded/skills/decide/SKILL.md` (via `sync-skills.sh`)

**Changes:**

1. **Update the ADR format block** in `decide/SKILL.md` (lines 30-46) to show the optional chain fields and document the Status lifecycle:
   ```
   Status: accepted | superseded
   Supersedes: ADR-NNNN      # only if this replaces an earlier ADR
   Superseded by: ADR-NNNN   # only if a later ADR replaces this one
   ```
   Add a short note: an ADR is *active* unless its Status is `superseded` (or it carries `Superseded by:`). `knack decisions check` treats superseded ADRs as exempt from the orphan gate — their replacement carries coverage.

2. **Flesh out the supersede procedure** (step 2, line 56): keep the existing two-step (mark old `Status: superseded` + add `Superseded by: NNNN`; write new ADR with `Supersedes: NNNN`) and add the invariant: the link must be *mutual* — both sides reference each other, or `decisions check` flags a broken chain.

3. **Update `adrTemplate`** (`instructions.go:54-70`) to include the optional chain lines as comments, matching the work-unit template's commenting style:
   ```
   Status: accepted
   # Supersedes: ADR-NNNN      # if this replaces an earlier ADR
   # Superseded by: ADR-NNNN   # if a later ADR replaces this one
   ```

4. **Sync embedded skills:** run `cli/sync-skills.sh`, verify with `diff -r .agents/skills cli/embedded/skills`.

**Verify:**
```bash
cd cli && go test ./... && cd .. && diff -r .agents/skills cli/embedded/skills && ./tests/run.sh
```

---

## Unit 3 — Confidence calibration + pre-emit verification gate (review skill)

**The steal (highest value, per the gstack session):** Two intertwined review upgrades, expressed compactly in knack's shape (~20 lines added to one file):
- Every finding must cite the specific `file:line` that motivates it and state a confidence level.
- **Pre-emit gate:** if you can't quote the motivating line, the finding is demoted to a speculative appendix rather than promoted to the report. A mechanical quality gate on an otherwise-subjective activity — directly aligned with "verify must be deterministic, worker never self-certifies."

**File:** `.agents/skills/review/SKILL.md` (+ sync to embedded)

**Changes — add a new section `## Confidence and evidence` after the "Two-axis review" section (after line ~60):**

The finding format becomes:
```
[axis] finding: <what's wrong> (confidence: high|medium|low)
  evidence: path/to/file.go:42 — "<the line or block that motivates this>"
```

Three calibration rules, distilled from gstack's 1-10 scale into knack's three-level idiom (matching the trivial/actionable/disputed/deferred taxonomy already in the skill):
- **high** — you read the specific code and can quote the line. Promoted to the report, handed to `fix`.
- **medium** — pattern match, likely but not verified against the actual code. Promoted but flagged; `fix` treats it as worth investigating, not worth acting on blindly.
- **low / uncitable** — you can't point to a specific line. **Not promoted.** Banished to a `## Speculative` appendix at the end of the review. Only surfaces if the user reads the appendix.

The gate (one sentence, the core steal): *Before you write a finding into the report, quote the code that motivates it. If you can't, it goes to the appendix, not the report.*

Update the `## Output` section to reflect the new finding format and the appendix. Keep the existing trivial/actionable/disputed/deferred classification — it's orthogonal to confidence (a finding can be `actionable` + `high confidence`, or `deferred` + `low confidence`).

**Verify:**
```bash
cli/sync-skills.sh && diff -r .agents/skills cli/embedded/skills && ./tests/run.sh
```
(No Go changes; `skills check` validates the skill still parses.)

---

## Unit 4 — LEARNINGS.md as a durable artifact

**The steal:** gstack's `learn` skill, but as a plain markdown file (not Postgres/JSONL). The user chose a new file over formalizing AGENTS.md.

**Files:**
- `LEARNINGS.md` (new, repo root — seeded + format defined)
- `.agents/skills/build/SKILL.md` (route operational learnings here instead of AGENTS.md)
- `.agents/skills/fix/SKILL.md` (resolve the dangling "or a backlog" reference → LEARNINGS.md)
- `.agents/skills/explore/SKILL.md` (mention LEARNINGS.md in "What to read")
- `DESIGN.md` (add to durable list + project layout)
- `AGENTS.md` (Core artifacts list + working conventions)
- `README.md` (Files section)
- `cli/embedded/skills/` (via sync)

**Changes:**

1. **Create `LEARNINGS.md`** with a header explaining its purpose and the entry format:
   ```markdown
   # Learnings

   Durable insights that don't fit in an ADR (not a ruling) or AGENTS.md (not operational). Append-only. Each entry: date, what was learned, where it bites.

   ## 2026-07-10 — <insight>

   <one paragraph. What happened, what you learned, where in the codebase it matters.>
   ```
   AGENTS.md's lessons-learned are tightly operational (verify-command gotchas, the loop+Devin behavior) and should *stay*. LEARNINGS.md draws an explicit line: *AGENTS.md = "how to work in this repo" (operational). LEARNINGS.md = "what we discovered about the problem/domain" (insights).* Seed empty or with a single placeholder explaining that line.

2. **build skill** (`build/SKILL.md:24-28`, "Capturing operational learnings"): split the routing. Operational gotchas (build commands, test conventions) → AGENTS.md (unchanged). Domain/problem insights ("X doesn't work because Y", "the parser has this surprising property") → LEARNINGS.md. Add the distinction in one paragraph.

3. **fix skill** (`fix/SKILL.md:18`): change `Note it in AGENTS.md or a backlog` → `Note it in LEARNINGS.md`. The phantom backlog reference resolves.

4. **explore skill** ("What to read" section, ~line 70): add `LEARNINGS.md` to the list of durable context to read before exploring.

5. **DESIGN.md** `## Durable vs disposable` (L255): add LEARNINGS.md as the 7th durable bullet with the AGENTS.md-vs-LEARNINGS.md distinction. Add to `## Project layout` (L217) tree.

6. **AGENTS.md**: add `LEARNINGS.md` to Core artifacts and a one-line working convention.

7. **README.md** Files section (L108-118): add `LEARNINGS.md — durable insights (domain/problem learnings).`

**Verify:**
```bash
cli/sync-skills.sh && diff -r .agents/skills cli/embedded/skills && cd cli && go test ./... && cd .. && ./tests/run.sh
```

---

## Unit 5 — "Explore first" as positioning

**The steal:** OpenSpec's reframe — explore leads the README/docs, positioned as the cure for "AI eagerly builds the wrong thing." knack already has a strong explore skill; it's just undersold. This is a docs rewrite, no code.

**Files:**
- `README.md`
- `docs/README.md`
- `docs/skills.md` (already the only doc naming explore — reinforce)

**Changes:**

1. **README.md** — add a new `## Why` or `## The problem knack solves` section near the top (after the one-line description, before Quickstart). Frame the upstream pain: AI agents eagerly build before understanding the problem; knack's loop forces a verify gate, but the bigger lever is the *explore* stance — grilling intent before any code is written. Name explore explicitly and link the explore skill. Keep it to ~6 lines. The existing "A tiny loop-packet runner" line stays as the mechanical description; this new section is the *why*.

   Also add a one-line mention of explore in `## How it works` or as a note after it: "Before the loop: run `explore` to grill intent. The loop is for when you already know what to build." (Currently the README never tells you how a QUEUE.md comes into existence — this closes that gap.)

2. **docs/README.md** — reorder so the skills/explore framing leads, not the loop mechanics. Add a "Start here: explore" pointer.

3. **docs/skills.md** — the explore entry in the default-skills table gets a one-line emphasis: explore is the entry point, not just one of seven skills. Reinforce the "stance not workflow" framing (read-only, no artifacts, reaches clarity before plan).

**Verify:**
```bash
./tests/run.sh
```
(No code or skill-structure changes; docs only.)

---

## Unit 6 — Review-fix subloop orchestration

**The missing piece:** The loop only executes `build` ticks today. The intended autonomous flow is `build → review → fix → build → review → fix` until the reviewer finds no actionable issues. The loop should own this bounded subloop.

**Files:**
- `loop.sh` — add `--review`, `--max-review-rounds`, optional `LOOP_REVIEW_CMD`/`LOOP_FIX_CMD`
- `prompts/worker.md` — unchanged; still loads `build` for build ticks, `review` for review ticks, `fix` for fix ticks
- `.agents/skills/review/SKILL.md` — update to write a structured `.loop/<name>/REVIEW.md` artifact
- `.agents/skills/fix/SKILL.md` — update to consume `REVIEW.md` and append fix units to `QUEUE.md`
- `decisions/` — new ADR-0008 capturing the loop-orchestrated review-fix decision
- `docs/loop.md`, `docs/skills.md`, `README.md`, `AGENTS.md` — document the new behavior

**High-level sequence:**

1. Build pass: run all pending units as today.
2. If `--review` is set and the queue reaches `done` state:
   - Run a review worker with the `review` skill, `QUEUE.md`, `EVIDENCE.md`, and the current diff.
   - The review worker writes `.loop/<name>/REVIEW.md` with findings classified as `actionable`, `trivial`, `disputed`, or `deferred`.
   - If no actionable findings: cycle complete.
   - If actionable findings: run a fix worker with the `fix` skill and `REVIEW.md`. The fix worker appends new work units to `QUEUE.md`.
   - Re-run the build pass on the new units.
   - Re-run review.
   - Stop when no actionable findings, `--max-review-rounds` is reached, `--max-ticks` is reached, or a review-round/tick limit is reached (the repeated-identical-actionable-findings stop is deferred; the round cap is the backstop).

**Hard stops:**
- `--max-ticks` total build-tick budget across all review-fix rounds (build ticks only)
- `--max-review-rounds` (default 2)
- No actionable findings in a review pass
- Worker exits non-zero

The following hard stop is deferred in this implementation (the round cap is the backstop):
- Repeated identical actionable findings (no progress)

**Agent routing:**
- `LOOP_AGENT_CMD` (default) for build ticks
- `LOOP_REVIEW_CMD` (optional; defaults to `LOOP_AGENT_CMD`) for review ticks
- `LOOP_FIX_CMD` (optional; defaults to `LOOP_AGENT_CMD`) for fix ticks

**The `REVIEW.md` artifact (disposable):**

```markdown
# Review: <cycle>
Generated: <timestamp>

## Standards
- [ ] high | medium | low — finding with evidence
  evidence: path/to/file.go:42 — "quoted line"

## Intent
- [ ] high | medium | low — finding with evidence
  evidence: path/to/file.go:42 — "quoted line"

## Speculative
Findings that could not cite a specific line.

## Summary
- actionable: N
- trivial: N
- disputed: N
- deferred: N
```

**Implementation cycle:** See `.loop/review-loop/QUEUE.md` and `.loop/review-loop/specs/proposal.md`.

## Out of scope (explicitly not doing)
- `knack doctor` (unified health check) — deferred in the prior session, stays deferred.
- Team-sharing ADR — needs its own ADR draft, separate effort.
- gstack's runtime infra (Postgres/gbrain, JSONL stores, 100KB codegen skills) — rejected as the wrong shape.
- `allowed-tools` frontmatter — rejected in prior session (agent-specific, inert in pi).
