---
role: record
owns: operational-context
---

# AGENTS.md

## Project

**knack** (name pending) is a composable **skills collection** for working with coding agents — the procedural encoding of the [AgenticWiki](https://github.com/bermudi/AgenticWiki)'s theory. It ships as plain [agentskills.io](https://agentskills.io) skills, installable into any agent via [`npx skills`](https://github.com/vercel-labs/skills) / [skills.sh](https://skills.sh). An optional bash batch runner (`./knack run`) provides unattended ("AFK") execution; everything else is interactive.

It replaces litespec. Specs are disposable; code is the source of truth; decisions and skills are durable.

`docs/architecture.md` and `docs/theory.md` carry the conceptual shape and lineage; `decisions/` carries the rulings. When they disagree, the ADR wins.

The spine (ADR-0009 onward) is derived from `decisions/` frontmatter — run `./knack spine` to list it. The spine marks the load-bearing rulings that define what knack is and how it works; everything before 0009 is pre-reframe history.

## Thesis

Work with agents happens across a spectrum of **human attention**, not a pipeline:

- **Interactive** — human present, edits land directly.
- **Plan-then-leave** — human present for the hard thinking, then the agent builds.
- **Batch (AFK)** — human absent; the loop runs units behind a verify gate.

Skills serve all three; the loop serves only batch. Skills are the product; the loop is an optional companion.

Skills transmit **concepts and reasoning, not rules** (ADR-0010). Judgment (decomposition, process, depth) is concept-guided — the agent decides. Mechanical contracts (the verify gate runs outside the agent and must exit 0; hard stops) stay hard rules. The wiki is the cited source of the "why"; the project's docs link to it rather than redefine its concepts — skills carry the synopsis, docs carry the link (ADR-0013).

## Current state

Two layers live in this repo. **`skills/` is the product** — what `npx skills add` installs into *other* projects, so it must be self-contained: no external links, which would be dead weight in a foreign project's context (ADR-0013). **Everything else here is knack's own development context** — `AGENTS.md`, `glossary.md`, `decisions/`, `README.md`, `docs/` — which guides working *on* knack, is never installed, and links freely.

- `skills/` — eight skills (`explore`, `plan`, `build`, `review`, `fix`, `decide`, `domain-modeling`, `document`). **These are the product.** Spec-compliant; installable via `npx skills add <owner>/<repo>`. All eight reworked to ADR-0010 (mode-independent, concept-forward); no external links.
- `prompts/` — worker / reviewer / fixer prompts the batch runner sends.
- `decisions/` — durable ADRs with YAML frontmatter (`id`, `date`, `status`, `spine`, optional `supersedes`/`superseded_by`/`amends`/`grandfathered`). `glossary.md` — ubiquitous language: knack-domain terms defined here; wiki concepts linked, not redefined (ADR-0010).
- `knack` — the project's single bash entry point (ADR-0017, ADR-0018). Verbs: `spine` (derive the spine from ADR frontmatter), `adrs` (list all ADRs), `check` (fail on structural drift — re-enumerated spine lists, duplicate ownership, missing frontmatter), `view` (read-only dashboard of cycles, work units, decisions), `run` (the optional AFK batch runner — agent-agnostic via `LOOP_AGENT_CMD`, owns the verify gate, per-unit `Agent:` overrides, opt-in review/fix via `--review`). No Go, no compile step. Distribution of skills is skills.sh's job.

## Core artifacts

**Durable** (maintained records whose value survives the work cycle): `skills/`, `decisions/`, `glossary.md`, `LEARNINGS.md`, this file, and `.loop/<name>/EVIDENCE.md` (the append-only ledger kept after a cycle's `QUEUE.md` is deleted).

**Disposable** (consumed then discarded): `.loop/<name>/QUEUE.md`, `HANDOFF.md`, `REVIEW.md`, `specs/`.

The load-bearing distinction: specs are disposable; code, decisions, and skills are durable. Treating disposable coordination state as durable is the spec-rot failure mode.

## Working conventions

- Ship as skills. `npx skills` is the package manager; skills.sh is the directory. Don't build distribution machinery.
- The AgenticWiki — knack's cited theory — has a local checkout at `~/Documents/AgenticWiki`. When a task needs to read, cite, or enrich the wiki (ingests, concept synopses for the skills, glossary depth), work in that tree — don't clone a duplicate. (A fresh clone once landed in `~/build/` and an ingest went into the wrong tree; the knack docs named only the GitHub URL, never the local path.)
- Skills transmit concepts + reasoning (ADR-0010). Don't prescribe *when* to deploy a concept as a rule. Hard rules are only for mechanical contracts.
- The worker never declares done without a passing verify. In batch the loop owns the gate; interactively the worker runs it — the principle survives across modes, the enforcement mechanism doesn't.
- A `Verify` command must be deterministic and executable by the runner — not an LLM-as-judge.
- `LOOP_AGENT_CMD` overrides the worker invocation (agent-agnostic). `LOOP_REVIEW_CMD` / `LOOP_FIX_CMD` override review/fix. Per-unit `Agent:` overrides for one unit.
- Work units are `## <outcome>` headers with `Read first:`, `Constraints:`, `Done means:`, `Verify:`. `Done means:` is acceptance criteria; `Verify:` is the mechanically enforceable subset. The gap is the review surface.
- Specs are disposable. Decisions are durable. Code is the source of truth.
- Durable-artifact hygiene (orphan ADRs, stale glossary terms, stale projections in docs) is judgment — transmitted as concepts in the `decide`, `domain-modeling`, and `document` skills, not enforced by gate commands. Structural drift (re-enumerated spine lists, duplicate ownership, missing frontmatter) is mechanical — `./knack check` catches it (ADR-0017).
- The evidence ledger (`EVIDENCE.md`) carries a registry-derived proof boundary (mechanical: `knack run` derives it from the verify command) and a pin-state record (mechanical: `knack run` records which durable docs were touched and alerts when a prior pin moves). Pin alerts are triage triggers for `review` → `document`, not coherence gates (ADR-0016).
- Operational gotchas go here; domain/problem insights go in `LEARNINGS.md`.

## Verification

```bash
./tests/run.sh
```

Exercises `knack run` (parsing, verify gate, handoff, review-fix) and `knack view`, validates skill frontmatter via `skills-ref` when available, and runs `knack check` (spine derivation, structural drift, frontmatter validity). There is no Go CLI; `go test` is gone.

## Lessons learned

- **The loop works end-to-end with Devin as the worker.** `LOOP_AGENT_CMD='devin --print --prompt-file "$LOOP_PROMPT_FILE" --model kimi-k2.7 --permission-mode dangerous'` drove units to verified completion. `--permission-mode dangerous` is required for the worker to write code and run verify without hanging on approvals.
- **The worker prompt names the skill explicitly (ADR-0007).** `prompts/worker.md` tells the worker to load the `build` skill by name and path; the loop passes it via `LOOP_PROMPT_FILE`. Trigger-based discovery isn't reliable enough across agents.
- **A verify command run across units compounds.** Each unit's verify runs prior units' tests too, so regressions surface at the next gate.
- **Review catches what verify can't.** The queue-parser regex `^##\s*(.*)$` once matched `###` subheadings as work units — a real bug tests missed because no fixture used `###`. Adversarial review against the actual codebase found it. Fix: exclude `###` in the unit-header check.
- **The verify gate proves the mechanical contract, not the coherence contract.** Tests-green + no-dead-refs-in-code can coexist with a durable doc that contradicts the ADRs: `AGENTS.md` once claimed `glossary.md` held "knack-domain terms only" while the file was 21 wiki-concept redefinitions, and the `README` linked the AgenticWiki as a public backbone while it was a private repo. After changing rulings, separately check that durable docs (`AGENTS.md`, `glossary.md`, `README`, `docs/`) still cohere with them — the grep that proves "no dead CLI refs in code" deliberately excludes docs and will not catch this. Coherence is a separate gate from compilation, handled by the `document` skill. The pin-state record (ADR-0016) catches *direct* drift — a durable doc that a prior cycle pinned has since changed — and routes it to `review` → `document`. It does not catch *indirect* coherence failure (A changed in a way that contradicts unpinned B); that remains judgment.
- **Hand-maintained lists of derivable facts drift; make them derivable.** The "spine" — the curated set of load-bearing ADRs — was listed in three docs (`AGENTS.md`, `docs/architecture.md`, `README.md`) with three different answers (seven, six, and five entries respectively). ADR-0015 says each fact has one owner; the spine-list had three owners and no record. ADR-0016's pin-check couldn't catch it because no pin was set on the list itself — it was indirect coherence failure on an unpinned fact, exactly the case the project's own lessons-learned called out. Fix (ADR-0017): the spine is now a `spine: true` field in ADR frontmatter; `./knack spine` derives it; `./knack check` fails on any doc that re-enumerates it. The lesson generalizes: if a fact is already in the artifacts, derive it — don't re-state it in prose that will drift.
- **Verify commands must be path-correct.** A unit once had `cd subproj && go test ./... && ./tests/run.sh` — but `./tests/run.sh` ran from `subproj/` after the `cd`, not the repo root. The loop correctly caught the failure; the verify command was wrong. Always test verify commands manually before writing them into a queue.
- **Workers scope to the outcome plus constraints, not a file list (ADR-0005).** A constraint states what must stay true or what is out of bounds — never what to edit. Naming files in constraints smuggles scope the same way the old `Work:` field did.
- **Orphan-ADR semantics.** An ADR is orphaned when it no longer explains or constrains the system — references (`QUEUE.md`, `EVIDENCE.md`, code, docs) are evidence of relevance, not the definition; a negative ruling can be alive with no citing work. (Established by ADR-0012, which supersedes the deleted checker's citation model from ADR-0006; transmitted as a concept by the `decide` skill, not a gate — ADR-0011.)
- **Named work cycles enable concurrent work (ADR-0004).** Each cycle gets `.loop/<name>/`; `./knack run .loop/<name>/QUEUE.md` is independent of others.
- **Negative-grep verifies must anchor on field syntax, not bare mentions.** A verify `! grep -rn 'Work:' ...` rotted the moment history documented "ADR-0005 replaced `Work:`." Anchor to the field shape (`^Work:`), not any mention of the word.
- **Design notes prevent fix-direction oscillation.** The work-unit format carries *what* but not *why*, so a fixer once picked the wrong branch of an ambiguous fix direction and reverted a manual fix. `.loop/<name>/DESIGN.md` carries cycle-level reasoning; the review skill's `fix direction` must be a single unambiguous instruction, never a conditional.
