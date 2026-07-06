# Loop Queue: validate-shape

Goal:
Prove that a work unit written in the new shape (ADR-0005) — outcome-level constraints, no file lists — produces a worker that uses judgment to find and fix files on its own. This is a real task, not a contrived test.

Stop condition:
`(cd cli && go test ./...) && (cd cli && go run . instructions adr) | grep -q 'Status: accepted'`

## the ADR template emitted by knack instructions matches the project's actual ADR convention

Why:
The `knack instructions adr` template currently emits `Status: proposed`, but every ADR in the project uses `Status: accepted`, and the decide skill teaches `Status: accepted`. A scaffolded project would inherit the wrong convention. This is a real inconsistency, not a manufactured exercise.

Read first:
- The decide skill — it defines the ADR format and the status convention
- The actual ADRs in `decisions/` — they show what the convention looks like in practice

Constraints:
- The template must match what the decide skill teaches and what actual ADRs use — not what seems reasonable in isolation.
- Any test that asserts on the old status value must be updated to match the new one. A test that passes against the wrong convention is worse than no test.

Done means:
- `knack instructions adr` emits `Status: accepted`.
- All existing CLI tests still pass.
- No other template emitted by `knack instructions` is changed.

Verify:
```bash
(cd cli && go test ./...) && (cd cli && go run . instructions adr) | grep -q 'Status: accepted'
```

Status: done
