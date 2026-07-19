---
id: 0005
date: 2026-07-06
status: accepted
spine: false
---

# 0005: Work units carry shape, not scripts

## Context

The original work unit template had a `Work:` field — a bulleted list of narrow work instructions. In practice, planners filled it with file-by-file edit scripts ("Update `.agents/skills/plan/SKILL.md`: change all `.loop/QUEUE.md` references to..."). The named-cycles cycle proved the failure mode: the worker executed exactly the listed edits, *saw* stale references in files outside the list, noted them in its handoff, and correctly declined to fix them. The script defined the scope, so the worker behaved as a script executor instead of an agent.

The insight (from skill-design practice, "constraints over prescription"): closing off the solution space is more effective than prescribing the path through it. Over-prescriptive instructions override the model's native competence and make output worse. The planner should communicate the *shape* of the work — what must be true after, what to read, what boundaries hold — and let the worker determine the approach.

Alternatives considered:
- **Keep `Work:` but write it less prescriptively.** Rejected — the field's affordance is a to-do list; discipline in prose doesn't survive contact with a planner under pressure.
- **Copy GSD's plan format** (`<task>` XML elements, requirement IDs, `must_haves`). Rejected — our units are smaller in scope than GSD phases; ADR references suffice for traceability, and the structure carries ceremony we don't need.
- **Optional steps field as escape hatch for mechanical work.** Rejected — the named-cycles unit *was* mechanical, and the worker handled it fine; the failure was the script defining the scope, not the absence of steps. Mechanical work gets a precise outcome and a precise verify instead.

## Decision

A work unit communicates shape through four fields, replacing `Work:`:

- **`Read first:`** — 2–4 pointers to the context the worker needs: files, ADRs, or code areas ("ADR-0004", "how loop.sh parses units"). Areas and rulings preferred over file enumerations. It is context, not scope.
- **`Constraints:`** — boundaries and conventions. The test: a constraint states what must stay true or what is out of bounds — never what to edit. If it names a file, it is "don't touch X" or "X's public API must not change," not "update X."
- **`Done means:`** — acceptance criteria; what must be observably true after the unit. Placed *above* `Verify:`.
- **`Verify:`** — the deterministic command; the mechanically enforceable subset of `Done means:`.

The unit's scope is its outcome plus its constraints — not an enumerated file list. The worker determines which files to change and how. The gap between `Done means:` and `Verify:` is the review surface: opt-in review reads what the command cannot check.

`Status:`, `Agent:`, `Why:`, and the `## <outcome>` header are unchanged. The loop's parser is unaffected — it reads only headers, `Status:`, `Agent:`, and the `Verify:` fenced block.

## Consequences

- Workers use judgment: a worker that sees related breakage inside the outcome's scope fixes it, because nothing tells it "your scope is these five files."
- Planners must invest in verify precision. If the outcome can't be captured in a deterministic verify command plus short acceptance criteria, the unit isn't ready — vague shape plus weak verify gives the worker nothing to aim at.
- The `plan` skill (template + valid-unit test), `build` skill, `review` skill (reads unit fields), `fix` skill, `prompts/worker.md` (rule 9: scope = outcome + constraints), `cli/internal/instructions/instructions.go`, DESIGN.md, and AGENTS.md all need updating. Embedded skills must be re-synced.
- Plans stay ephemeral. This changes what the plan carries, not its lifecycle.
