---
id: 0011
date: 2026-07-11
status: accepted
spine: true
supersedes: [0001, 0002]
---

# 0011: Ship as a skills collection via skills.sh; delete the Go CLI

## Context

ADR-0009 re-centered the project on skills as the product and questioned the Go CLI's package-manager semantics. Investigation showed the CLI was solving a problem the ecosystem has already standardized on.

[`vercel-labs/skills`](https://github.com/vercel-labs/skills) (`npx skills`) is the de-facto package manager for the open agent skills ecosystem — `add`, `use`, `list`, `find`, `remove`, `update`, `init`. It auto-detects **70+ agents** and installs to each one's native skills path (`.agents/skills/`, `.claude/skills/`, `.codex/skills/`, …), following the [agentskills.io](https://agentskills.io) spec — the same spec knack already follows. It discovers `.agents/skills/` and `skills/` natively, and backs the [skills.sh](https://skills.sh) directory.

The Go CLI's package-manager surface (`skills init` / `check` / `update` via a bespoke manifest + SHA-256 scheme) is a strict subset of what `npx skills` already does — for one agent, behind a compile step, with a duplicate of the very skills it ships. It was reimplementing a solved standard.

This also makes the per-ecosystem plugin-manifest option (considered mid-discussion) moot: `npx skills add` is the universal adapter that lands the same `SKILL.md` in every agent's skills directory. No `.claude-plugin/`, no `.devin-plugin/`, no N manifests.

Alternatives considered:

- **Keep the CLI as a lint-only tool.** Rejected — its lints are either mechanical (work-unit structure, only relevant to the loop) or judgment (orphan decisions, stale glossary terms), and ADR-0010 says judgment belongs in skills as transmitted concepts, not in gate commands.
- **Per-ecosystem plugin manifests.** Rejected — `npx skills` already adapts to every ecosystem via the shared spec; manifests would be redundant adapters.
- **Bash scripts replacing the CLI.** Superseded by this decision — `npx skills` obviates even a bash package manager.

## Decision

**Ship as a skills collection.** The repo's product is `skills/` — plain agentskills.io `SKILL.md` files. That is what replaces `/plan` commands, spec-kit, and openspec.

**`vercel-labs/skills` is the package manager; skills.sh is the directory.** Users run `npx skills add <owner>/<repo>` and pick their agent(s). Install (symlink or copy), update, discovery, and removal are all inherited — knack ships none of that machinery.

**Delete the Go CLI entirely.** ADR-0001 and ADR-0002 are superseded. No `cli/`, no compile step, no Go toolchain, no embedded skill copies, no `sync-skills.sh`, no manifest + SHA versioning.

**The loop stays as an optional companion, not a skill.** `loop.sh` remains a bash script for AFK batch work (ADR-0009). The mechanical `validate` (work-unit structure) rides with it — the loop already parses `QUEUE.md` in bash. The judgment lints (`decisions check`, `glossary check`) do not survive as commands: per ADR-0010, orphan-decision and stale-term hygiene become concepts the `decide` and `domain-modeling` skills transmit, so the agent self-checks from understanding rather than from a gate.

## Consequences

- First supersede in the project: 0001 and 0002 marked superseded with mutual links. ADR-0009's "CLI role narrows, questioned not assumed" resolves to "deleted."
- Repo restructures: `cli/` removed; `.agents/skills/` → `skills/`; the CLI's `MANIFEST.json` removed.
- Removes the Go toolchain dependency, the compile step, the embedded-skill duplication, and the bespoke versioning scheme. Distribution, versioning, and discovery become skills.sh's problem.
- Low coupling: skills are spec-compliant plain files. If `npx skills` or skills.sh fades, the skills survive and any package manager (or a manual copy) works.
- Depends on a third-party PM for the best install UX, but the skills themselves have no such dependency.
- `tests/run.sh` and `AGENTS.md` lose their `go test` / CLI steps; verification is now `./tests/run.sh` (loop behavior) plus skill-spec validity.
- `README.md` and `docs/` reference a now-deleted CLI throughout and need a rewrite pass (separate from this decision).
