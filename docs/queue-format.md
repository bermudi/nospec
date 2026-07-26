---
nospec: true
role: view
---

# Queue format

`QUEUE.md` is a disposable Markdown file containing a queue of work units. The loop parses it mechanically; the rest is for humans and the worker.

## Header

```markdown
# Loop Queue: <short name>
```

## Goal

```markdown
Goal:
<one paragraph describing the desired end state>
```

Queue completion is derived from unit statuses. Integration checks belong in a unit's `Verify:` block; the format has no decorative queue-level command.

## Work unit

Each unit is a top-level `##` heading:

```markdown
## <outcome — what changes, observable>
```

The heading is a required, nonempty outcome. There is no `Slice` prefix and no numbering. `###` subheadings and heading-shaped lines inside fenced blocks remain part of the current unit.

## Fields

| Field | Required | Description |
|---|---|---|
| `Agent:` | optional | Overrides `LOOP_AGENT_CMD` for this unit. |
| `Why:` | optional | Non-obvious context only. |
| `Read first:` | optional | Nonempty context when needed: ADRs, code areas, or rulings. |
| `Constraints:` | optional | Nonempty boundaries when needed: what must stay true or what is out of bounds. |
| `Done means:` | required | Nonempty acceptance criteria. |
| `Verify:` | required | A fenced `bash` block with a deterministic, discriminating command. |
| `Status:` | required | One of `pending`, `in_progress`, `done`, `verify_failed`, `no_progress`, `blocked`. |

## Field rules

- `Read first:` and `Constraints:` are optional; omit them rather than writing placeholder content. When present, each field must occur once and contain content on following lines.
- `Read first:` is context, not scope. Prefer areas and rulings over long file lists.
- `Constraints:` state what must stay true or what is out of bounds. They never say what to edit. If a constraint names a file, it is "don't touch X" or "X's public API must not change", not "update X".
- `Done means:` is required and nonempty because it defines the acceptance surface against which verification and review can be understood.
- `Verify:` is the mechanically enforceable subset of `Done means:`. It should fail for a plausible state where the unit's central outcome is absent; lint rejects obviously vacuous commands such as `true`, `:`, and `exit 0`. The remaining gap is the review surface.
- `Done means:`, `Verify:`, and `Status:` are required and unique. `Status:` starts as `pending`; the loop updates it.

## Example

````markdown
# Loop Queue: parser fix

Goal:
Make the queue parser ignore `###` subheadings.

## queue parser ignores `###` subheadings

Read first:
- `nospec run` queue parser
- `tests/run.sh` parser tests

Constraints:
- `nospec run` behavior remains the source of truth for the parser.
- No other headings are affected.

Done means:
- `nospec run` does not treat `###` lines as unit boundaries.
- Existing tests still pass.

Verify:
```bash
./tests/run.sh
```

Status: pending
````

## Status values

See [loop.md](./loop.md#work-unit-statuses) for the list and meanings.

## Disposability

When every unit is verified and any review state is clean, `QUEUE.md` and `HANDOFF.md` are disposable. A drained queue with actionable findings in `REVIEW.md` is still review-blocked. Keep `EVIDENCE.md` if you want to trace which ADRs the cycle referenced.
