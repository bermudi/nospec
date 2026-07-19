---
role: view
---

# Queue format

`QUEUE.md` is a disposable Markdown file containing a queue of work units. The loop parses it mechanically; the rest is for humans and the worker.

## Header

```markdown
# Loop Queue: <short name>
```

## Goal and stop condition

```markdown
Goal:
<one paragraph describing the desired end state>

Stop condition:
`<command that proves the whole packet is done, if one exists>`
```

The stop condition is optional and purely informational; the loop does not execute it.

## Work unit

Each unit is a top-level `##` heading:

```markdown
## <outcome — what changes, observable>
```

The heading is the outcome. There is no `Slice` prefix and no numbering. Avoid `###` headings inside a work unit — they may confuse simple parsers.

## Fields

| Field | Required | Description |
|---|---|---|
| `Agent:` | optional | Overrides `LOOP_AGENT_CMD` for this unit. |
| `Why:` | optional | Non-obvious context only. |
| `Read first:` | recommended | 2–4 bullets of context: ADRs, code areas, or rulings. |
| `Constraints:` | recommended | Boundaries: what must stay true or what is out of bounds. |
| `Done means:` | recommended | Acceptance criteria. |
| `Verify:` | required | A fenced `bash` block with a deterministic command. |
| `Status:` | required | One of `pending`, `in_progress`, `done`, `verify_failed`, `no_progress`, `blocked`. |

## Field rules

- `Read first:` is context, not scope. Prefer areas and rulings over long file lists.
- `Constraints:` state what must stay true or what is out of bounds. They never say what to edit. If a constraint names a file, it is "don't touch X" or "X's public API must not change", not "update X".
- `Done means:` is the acceptance criteria.
- `Verify:` is the mechanically enforceable subset of `Done means:`. The gap between them is the review surface.
- `Status:` starts as `pending`. The loop updates it.

## Example

````markdown
# Loop Queue: parser fix

Goal:
Make the queue parser ignore `###` subheadings.

Stop condition:
`./tests/run.sh` exits 0.

## queue parser ignores `###` subheadings

Read first:
- `knack run` queue parser
- `tests/run.sh` parser tests

Constraints:
- `knack run` behavior remains the source of truth for the parser.
- No other headings are affected.

Done means:
- `knack run` does not treat `###` lines as unit boundaries.
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

When a cycle is complete and verified, `QUEUE.md` and `HANDOFF.md` are disposable. Delete them. Keep `EVIDENCE.md` if you want to trace which ADRs the cycle referenced.
