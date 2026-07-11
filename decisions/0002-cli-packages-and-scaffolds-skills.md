# 0002: CLI packages and scaffolds the default skills

Date: 2026-07-06
Status: superseded
Superseded by: ADR-0011
Grandfathered: enacted before the evidence-ledger convention (ADR-0006); skills are embedded and scaffolded by the CLI.

## Context

Open question #3 in DESIGN.md: do skills ship with the tool, or does each project author its own? The lean was toward shipping defaults as plain markdown that projects can override. But there was no mechanism for getting the default skills into a project — the CLI was specified as validate-only.

The user ruled: the CLI has to "package" the skills, not just validate. This makes the CLI the distribution mechanism for the default skill set.

Alternatives considered:
- **Skills live only in the repo, users copy manually.** Works for this repo but doesn't scale to new projects. No version tracking, no update path.
- **Skills fetched from a remote URL at init time.** Adds a network dependency and a hosting concern. Fragile for offline use.
- **Skills embedded in the CLI binary, written out on `skills init`.** Single binary carries the defaults. `go:embed` makes this trivial. Projects override what they want after scaffolding. Version = CLI version.

## Decision

The CLI will embed the seven default skills (explore, plan, build, review, fix, decide, domain-modeling) via `go:embed` and provide a `skills init` command that scaffolds them into a target project's `.agents/skills/` directory. Projects can then modify or override the scaffolded skills freely — the CLI does not manage them after init.

The CLI remains read-only for everything else (validate, check, status). `skills init` is the one write operation, and it only writes to `.agents/skills/`.

## Consequences

- The Go binary is the source of truth for the default skill set at a given CLI version. The `.agents/skills/` in this repo and the embedded skills in the binary must stay in sync.
- `skills init` will not overwrite existing skills by default — it scaffolds missing ones and leaves existing ones alone (so projects can upgrade the CLI without losing customizations).
- This resolves DESIGN.md open question #3: skills ship with the tool as embedded markdown, projects override after scaffolding.
- The skills in `.agents/skills/` in this repo serve double duty: they are this project's skills AND the canonical source for the embedded defaults. The build copies them into the CLI's embed directory.
