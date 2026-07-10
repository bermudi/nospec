# CLI reference

`knack` is a read-only validator and context provider. It does not run agents or the loop. Build it from `cli/`:

```bash
cd cli
go build -o ../knack .
```

## Commands

| Command | Arguments | Description |
|---|---|---|
| `skills init` | `[--target DIR]` | Scaffold the seven default skills into `<DIR>/.agents/skills/` (default `.`). Skips existing skills. |
| `skills check` | `[--dir DIR]` | Validate skills in `DIR` (default `.agents/skills`). |
| `validate` | `<queue-file>` | Validate work-unit structure in a queue. |
| `decisions list` | | List all ADRs in `decisions/`. |
| `decisions show` | `NNNN` | Print the full ADR with that number. |
| `decisions check` | | Flag orphaned ADRs and dangling references. |
| `status` | | Aggregate work-unit counts across all `.loop/<name>/` cycles. |
| `glossary check` | `[--file glossary.md]` | Check `glossary.md` term references. |
| `instructions` | `<template>` | Print a template. `<template>` is `work-unit`, `adr`, or `glossary-entry`. |

## Exit codes

- `0` — command succeeded or no findings.
- `1` — validation errors, missing arguments, or CLI failure.

## Examples

### Validate a queue

```bash
./knack validate .loop/my-cycle/QUEUE.md
```

### Scaffold skills into a new project

```bash
cd /path/to/project
/path/to/knack skills init
```

### Decision coverage

```bash
./knack decisions check
```

An ADR is orphaned if it is not referenced by any `QUEUE.md` or `EVIDENCE.md` in `.loop/`. A dangling reference points to an ADR that does not exist.

### Status across cycles

```bash
./knack status
```

Output:

```text
cycle go-cli: 1 pending, 2 done, 0 failed, 3 evidence
total: 1 pending, 2 done, 0 failed, 3 evidence
adrs: 6
```

### Get a template

```bash
./knack instructions work-unit
./knack instructions adr
./knack instructions glossary-entry
```

## Notes

- All commands read from the current directory (run from the repo root unless a path is given).
- `skills init` is the only write operation. It scaffolds missing skills and leaves existing ones alone, so upgrading the CLI will not overwrite project customizations.
- The CLI packages the default skills via `go:embed`. After editing `.agents/skills/` in the `knack` repo itself, run `cli/sync-skills.sh` and `diff -r .agents/skills cli/embedded/skills` to verify sync.
