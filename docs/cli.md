# CLI reference

`knack` is a read-only validator and context provider. It does not run agents or the loop. Build it from `cli/`:

```bash
cd cli
go build -o ../knack .
```

## Commands

| Command | Arguments | Description |
|---|---|---|
| `skills init` | `[--target DIR]` | Scaffold the seven default skills into `<DIR>/.agents/skills/` (default `.`). Skips existing skills. Writes the manifest and `.gitignore` patterns. |
| `skills check` | `[--dir DIR]` | Validate skills in `DIR` (default `.agents/skills`). Also reports modified and stale skills via the manifest. |
| `skills update` | `[--target DIR] [--force]` | Refresh scaffolded skills from the embedded source when a newer version ships. With `--force`, overwrite locally modified skills too. |
| `validate` | `<queue-file>` | Validate work-unit structure in a queue. |
| `decisions list` | | List all ADRs in `decisions/`. Superseded ADRs show the number that replaced them. |
| `decisions show` | `NNNN` | Print the full ADR with that number. |
| `decisions check` | | Flag orphaned ADRs, dangling references, and broken/one-sided supersede chains. |
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

`skills init` writes a manifest to `.agents/skills/MANIFEST.json` recording each
scaffolded skill's `name`, `version` (read from its frontmatter), and a SHA-256
content `hash` of the whole skill directory. It also appends `.gitignore`
patterns for disposable loop state (`.loop/**/QUEUE.md`, `.loop/**/HANDOFF.md`,
`.loop/**/REVIEW.md`, `.loop/**/specs/`) into the target directory, idempotently.

### Refresh skills after a CLI upgrade

```bash
/path/to/knack skills update            # refresh only unmodified, stale skills
/path/to/knack skills update --force    # overwrite every skill, modified or not
```

`skills update` compares each on-disk skill against the manifest and the embedded
source. A skill is overwritten if it exists, is unmodified (its current hash matches
the manifest), and the embedded version is newer than the manifest version. If a
skill has local changes (hash differs), it is skipped unless `--force` is given.
Absent skills are scaffolded. After updating, the manifest is rewritten with the
new versions and hashes.

### Decision coverage

```bash
./knack decisions check
```

An ADR is orphaned if it is not referenced by any `QUEUE.md` or `EVIDENCE.md` in `.loop/`. A dangling reference points to an ADR that does not exist. A broken supersede chain points to a missing or non-mutual link between ADRs.

### Status across cycles

```bash
./knack status
```

Output:

```text
cycle go-cli: 1 pending, 2 done, 0 failed, 3 evidence
total: 1 pending, 2 done, 0 failed, 3 evidence
adrs: 7 active (7 total)
```

### Get a template

```bash
./knack instructions work-unit
./knack instructions adr
./knack instructions glossary-entry
```

## Notes

- All commands read from the current directory (run from the repo root unless a path is given).
- `skills init`, `skills update` are the write operations. `init` scaffolds missing skills and writes the manifest plus `.gitignore` patterns; `update` refreshes unmodified skills when a newer embedded version ships, so upgrading the CLI does not overwrite project customizations. `--force` overrides that protection.
- Each skill carries a `version` field in its frontmatter. `skills check` reports `modified:` for locally changed skills and `stale:` for skills whose embedded version is newer than the manifest version; run `skills update` to reconcile.
- The CLI packages the default skills via `go:embed`. After editing `.agents/skills/` in the `knack` repo itself, run `cli/sync-skills.sh` and `diff -r .agents/skills cli/embedded/skills` to verify sync.
