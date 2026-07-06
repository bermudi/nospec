# Loop Queue: plan-shape

Goal:
Work units communicate shape, not scripts (ADR-0005). The `Work:` field is gone from the work-unit format, replaced by `Read first:` and `Constraints:`, with `Done means:` promoted above `Verify:`. Every artifact that teaches, consumes, or emits the format — the plan/build/review/fix skills, the worker prompt, DESIGN.md, AGENTS.md, and `knack instructions` — agrees on the new shape.

Stop condition:
`./tests/run.sh && (cd cli && go test ./...) && diff -r .agents/skills cli/embedded/skills && ! grep -rn 'Work:' .agents/skills DESIGN.md prompts/worker.md AGENTS.md cli/internal/instructions/instructions.go`

## Skills and docs teach the shape-not-scripts work-unit format

Why:
The old `Work:` field trained workers to be script executors — the named-cycles cycle proved it (worker saw stale refs outside the listed files and correctly declined to fix them). ADR-0005 rules that units carry shape: what must be true after, what to read, what boundaries hold. The worker determines which files to change and how.

Read first:
- decisions/0005-work-units-carry-shape-not-scripts.md — the ruling this unit implements
- The queue template and field notes in `.agents/skills/plan/SKILL.md`, and the mirrored format section in DESIGN.md
- prompts/worker.md — the worker's contract, especially rule 9 ("stated scope")
- This queue file itself — it is written in the new shape and can serve as the example

Constraints:
- The loop's parser contract must not change: `## <outcome>` headers, `Status:`, `Agent:`, and the `Verify:` fenced block keep their exact syntax. `loop.sh` and `cli/internal/queue/` need no changes.
- `Done means:` precedes `Verify:` in the new template. Verify is the mechanically enforceable subset of the acceptance criteria; the gap between them is the review surface, and the review skill must say so.
- The constraint test goes into the plan skill verbatim: a constraint states what must stay true or what is out of bounds — never what to edit. If it names a file, it is "don't touch X," not "update X."
- `Read first:` guidance: 2–4 entries; areas and rulings preferred over file enumerations; it is context, not scope.
- The worker's scope is the unit's outcome plus its constraints — prompts/worker.md and the build skill must both state this instead of pointing at a "stated scope" list.
- No references to the old `Work:` field may remain in skills, prompts, DESIGN.md, or AGENTS.md — including mentions in the review and fix skills. No historical "this replaces Work:" notes.
- `.agents/skills/` is canonical; `cli/embedded/skills/` must end in sync (`cli/sync-skills.sh`).

Done means:
- The plan skill's template and field notes describe Read first / Constraints / Done means / Verify, carry the constraint test, and warn that a unit whose outcome can't be captured by a deterministic verify plus short acceptance criteria isn't ready.
- The build skill and worker prompt define scope as outcome + constraints.
- The review skill reads `Done means:` and `Constraints:` instead of `Work:`, and names the Done-means/Verify gap as the review surface.
- The fix skill's unit template uses the new shape.
- DESIGN.md's work-unit format section and AGENTS.md's conventions match the skills.
- Embedded skills are in sync; existing tests still pass.

Verify:
```bash
./tests/run.sh && diff -r .agents/skills cli/embedded/skills && ! grep -rn 'Work:' .agents/skills DESIGN.md prompts/worker.md AGENTS.md && grep -q 'Read first:' .agents/skills/plan/SKILL.md && grep -q 'Constraints:' .agents/skills/plan/SKILL.md
```

Status: done

## `knack instructions` emits the new work-unit shape

Why:
The CLI's `instructions` command prints the queue template for agents working in target projects. If it still emits `Work:`, every scaffolded project teaches the anti-pattern.

Read first:
- decisions/0005-work-units-carry-shape-not-scripts.md
- cli/internal/instructions/ — the embedded template and its tests
- The updated queue template in `.agents/skills/plan/SKILL.md` (after the previous unit)

Constraints:
- The emitted template must match the plan skill's queue template — one shape, two places, no drift.
- Queue parsing behavior (`cli/internal/queue/`) must not change; prose fields are invisible to the parser.

Done means:
- `knack instructions` output shows `Read first:`, `Constraints:`, and `Done means:` above `Verify:`, and contains no `Work:` field.
- Tests assert the new template content.
- All existing CLI tests still pass.

Verify:
```bash
(cd cli && go test ./...) && ! grep -n 'Work:' cli/internal/instructions/instructions.go && grep -q 'Read first:' cli/internal/instructions/instructions.go
```

Status: done
