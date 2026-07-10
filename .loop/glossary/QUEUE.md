# Loop Queue: glossary

Goal:
Create a curated `glossary.md` for the project and fix `knack glossary check` so it is safe and passes on this repository. The CLI's `glossary check` must no longer error when `glossary.md` is missing, and it must not treat bash `[[...]]` syntax or Go test strings as glossary references. The glossary should define the project's domain vocabulary and the `[[...]]` references that `DESIGN.md` uses.

Stop condition:
`cd /home/daniel/build/knack/cli && go build -o /tmp/knack . && cd /home/daniel/build/knack && /tmp/knack glossary check`

## the glossary check is safe and precise

Read first:
- `cli/internal/glossary/glossary.go`
- `cli/internal/glossary/glossary_test.go`
- `DESIGN.md` (the glossary check section)

Constraints:
- Don't change the CLI's public command interface or any command other than `glossary`.
- The `glossary` package must remain fully testable with `go test`.
- The undefined-term matcher must still catch real `[[...]]` references in markdown files.
- Stay inside the CLI; do not touch `loop.sh` or the skills.

Done means:
- `knack glossary check` exits 0 when `glossary.md` is absent.
- `knack glossary check` ignores `[[...]]` in shell scripts and Go source files, but still catches `[[...]]` in markdown files.
- `cd /home/daniel/build/knack/cli && go test ./...` still passes.

Verify:
```bash
cd /home/daniel/build/knack/cli && go test ./... && go build -o /tmp/knack . && cd /home/daniel/build/knack && /tmp/knack glossary check
```

Status: done

## the project has a curated glossary and `knack glossary check` passes

Read first:
- The `domain-modeling` skill
- `DESIGN.md` (the `[[...]]` references it uses)
- `AGENTS.md` and `README.md`
- The `cli/internal/glossary` package after unit 1

Constraints:
- The glossary stays small and curated; each entry is one or two sentences.
- Define a term only if it is used in the project or is a `[[...]]` target in `DESIGN.md`.
- Do not change the CLI code or `DESIGN.md` unless the glossary check proves it is necessary.
- Keep entries in alphabetical order and flat (no categories, no nesting).

Done means:
- `glossary.md` exists at the repo root with the project's domain vocabulary and all `[[...]]` references from `DESIGN.md`.
- `knack glossary check` reports no findings.
- `cd /home/daniel/build/knack/cli && go test ./...` still passes.
- The `./tests/run.sh` harness still passes (run it manually after the queue is done if you want the full loop-level check).

Verify:
```bash
cd /home/daniel/build/knack/cli && go test ./... && go build -o /tmp/knack . && cd /home/daniel/build/knack && /tmp/knack glossary check
```

Status: done
