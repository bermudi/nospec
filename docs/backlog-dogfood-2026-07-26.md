---
nospec: true
role: view
---

# Backlog: Dogfood Findings Not Yet Implemented

Source: Two dogfood sessions on 2026-07-26 where nospec was installed into an
external project (`pi-session-search`) and evaluated by both the worker (Pi/GPT)
and a reviewer (Opus). The worker produced two detailed verdicts; the reviewer
produced a code-review-style findings list and a litespec comparison.

Session: `~/.pi/agent/sessions/--home-daniel-build-nospec--/2026-07-26T01-16-03-302Z`

Cross-referenced with git history: the 2026-07-28 consolidation commit
(`d45062d`, "Consolidate skills around five behavioral stances") and the
preceding ADRs (0021–0025, 0022, 0023) resolved findings #1–#7, #9, #11, and
#13. This file now tracks only the **partial** and **open** findings from the
original list: #8, #10, #12, #14, #15, #16, #17.

---

## 8. `nospec-shape` drops ADR-0010 reasoning

**Severity:** P2 — skill quality  
**Sources:** Opus review (finding #5)  
**Category:** Skill content

### Problem

ADR-0010 says skills must transmit concepts AND reasoning (why, failure modes,
trade-offs). The current nospec-shape skill describes vertical slices and
horizontal breadth as definitions/labels without explaining why each cut helps
or when it becomes risky.

### Current state

Shape skill is 102 lines. It does include some trade-off language ("A slice
too thin... gives false confidence; one too broad delays the feedback") but
the Opus review found it insufficient compared to ADR-0010's requirements.

### Proposed direction

Restore the "why" and "failure modes" for each cut type. One paragraph per
cut explaining: when to use it, what it proves, what it misses, and what
goes wrong when misapplied.

---

## 10. Add runner/parser tests

**Severity:** P2 — quality  
**Sources:** Opus review (finding #4)  
**Category:** Testing

### Problem

The runner is ~1,494 lines of Bash and the queue parser is ~482 lines of
Python. Neither has unit test coverage. The existing `tests/run.sh` exercises
the runner end-to-end but doesn't cover parser edge cases, signal handling,
or baseline recording in isolation.

### What was said

> "The installed artifact contains no runner tests." — Opus review

### Proposed direction

- Python: add pytest tests for queue_parser.py (field validation, shell syntax
  checking, accessor correctness, malformed-input handling).
- Bash: add integration tests for runner verbs (lint, spine, adrs, check, view)
  with fixture queues. The existing `tests/run.sh` is a start but exercises only
  the happy path.

---

## 12. Worker prompt: don't escalate reversible choices

**Severity:** P3 — prompting  
**Sources:** Opus review (finding #7)  
**Category:** Skill content

### Problem

Agents loaded with nospec skills tend to ask permission for every reversible
implementation detail, treating minor decisions as if they need ADR-level
confirmation. The worker prompt should explicitly counter this.

### What was said

> "Reversible local choices are worker-owned; only consequential trade-offs
> block execution." — Opus review

### Current state

The worker prompt (`skills/nospec-loop/prompts/worker.md`) does not contain this
line. The nospec-shape skill mentions it implicitly but not with the directness
the reviewer recommended.

### Proposed direction

Add to worker.md (and possibly shape's prompt):

> "Do not ask permission for reversible local choices — variable naming, file
> organization within the outcome, implementation technique, test structure.
> Only escalate when the choice is consequential, hard to reverse, or affects
> an external contract."

---

## 14. Family taxonomy reintroduces ADR-0017's violation

**Severity:** P2 — design consistency  
**Sources:** Opus review of ADR-0025  
**Category:** ADR consistency

### Problem

The family taxonomy (Understand/Change/Challenge/Remember/Leave) is a derivable
fact from ADR frontmatter, but it was re-enumerated by hand in README.md,
AGENTS.md, docs/architecture.md, docs/skills.md, and docs/getting-started.md.
This is exactly the spine-list failure that ADR-0017 was written to kill.

Additionally:
- The ordering implies a pipeline but every table has to say "not phases" (5
  disclaimers across docs is a tell).
- `metadata.family` with a closed five-value set is enforced by tests but has
  no ADR — it's a new taxonomy with no ruling.

### What was said

> "Either add a nospec verb (and let the docs point at it, as architecture.md's
> spine section already does), or accept the projections and add a check that
> they match frontmatter." — Opus review

### Proposed direction

1. Add an ADR for the family taxonomy (or explicitly reject it).
2. Derive family membership from frontmatter via `nospec family` (like `nospec
   spine`).
3. Add a `nospec check` rule that fails on re-enumerated family lists in docs.
4. Break the ordering appearance (alphabetize, or lead with Change).

---

## 15. ADR threshold: distinguish proposed vs. accepted

**Severity:** P2 — process design  
**Sources:** First worker review  
**Category:** Skill content / process

### Problem

The shaping instructions pushed the worker to record an *accepted* ADR while
the architectural choice was only implicitly accepted. The lazy-indexing ruling
was definite; the catalog architecture was still a recommendation.

Durable records should have a stronger confirmation threshold than ephemeral
plans.

### What was said

> "Nospec should distinguish: observed requirement, proposed ruling, explicitly
> accepted ruling." — Worker review

### Current state

ADR frontmatter has `status: accepted` but no mechanism for "proposed" or
"under review." The nospec-curator skill says agents must not launder
recommendations into accepted ADRs, but there's no structural enforcement.

### Proposed direction

Either:
- Add a `status: proposed` state to ADR frontmatter with its own lifecycle, or
- Make nospec-curator explicitly warn when a worker is writing an ADR without
  human confirmation.

---

## 16. AFK safety: isolated worktrees by default

**Severity:** P2 — safety  
**Sources:** Opus review  
**Category:** Runner safety

### Problem

The default worker command runs with full permissions in the current worktree.
There is no clean-tree requirement, no isolated worktree, no checkpoint, and no
rollback. The external verify proves its command passed but doesn't prove the
worker avoided collateral damage.

### Current state

`--repo DIR` exists for pointing at a different directory. `--accept-dirty-
baseline` records and accepts dirty-state risk. But isolation is opt-in, not
default.

### What was said

> "I'd want isolated worktrees by default, or at minimum a dirty-tree refusal
> requiring --allow-dirty." — Opus review

### Proposed direction

Default to `git worktree add` for each cycle. Refuse to run on the main
worktree without `--allow-dirty` or `--repo`. This is a significant UX change
but aligns with the safety framing.

---

## 17. Combine best of nospec and litespec

**Severity:** P3 — vision  
**Sources:** Litespec comparison  
**Category:** Product vision

### Problem

The dogfood comparison identified that the best system would combine:

- Litespec's behavioral canon and delta engine for durable public capabilities
- Nospec's disposable work queues and external verification for execution
- Code/tests as authority for implemented behavior
- Explicit amendment (not "never go backward")
- Human-owned archive/ruling decisions

### Status

This was identified as a direction, not a specific task. No work has been done.
The two projects remain separate.

### What to decide

Is this worth pursuing as a combined project, or are nospec and litespec
better as separate tools with complementary strengths?

---

## Summary

| # | Item | Severity | Effort |
|---|------|----------|--------|
| 8 | Shape drops ADR-0010 reasoning | P2 quality | Low |
| 10 | Runner/parser tests | P2 quality | High |
| 12 | Worker: don't escalate reversible | P3 prompting | Low |
| 14 | Family taxonomy: derive, don't re-enum | P2 consistency | Medium |
| 15 | ADR threshold: proposed vs accepted | P2 process | Medium |
| 16 | AFK safety: isolated worktrees | P2 safety | High |
| 17 | Combine nospec + litespec | P3 vision | Unknown |
