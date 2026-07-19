---
id: 0018
date: 2026-07-18
status: accepted
spine: false
amends: [0011, 0017]
---

# 0018: One command with verbs; fold the loop into `knack`

## Context

The project's mechanical bash surface is split across two executables:

- `./knack` — `spine`, `adrs`, `check` (derivation and linting).
- `./loop.sh` — `run`, `view` (the optional batch runner and its read-only dashboard).

ADR-0011 put the loop in `loop.sh` and deleted the Go CLI. ADR-0017 restored a bash CLI at `./knack` but narrowed its role to "derives and lints; does not govern, scaffold, package, or version." The `run` verb — which governs the batch loop — was left in `loop.sh` to preserve that narrowing.

The split has two costs. First, the surface is incoherent: `view` is read-only inspection of `.loop/` state and belongs with `spine`/`adrs` as much as with `run`, yet it lives on the other binary. Second, there are two commands to remember, two help texts, and two entry points for what is conceptually one thing — knack's bash tooling. The role-purity line ADR-0017 drew ("derives and lints, does not govern") turned out to be the wrong axis: the real distinction is **skills (the product) vs. bash tooling**, and all the bash tooling belongs in one binary. `run` does not become a skill by moving files; the loop stays optional and stays bash.

Alternatives considered:

- **Move only `view` into `knack`; keep `loop.sh run`.** Rejected — leaves two commands, which is the problem. The role-purity concern (a `run` verb on the derivation CLI) is real but cheaper to pay than the split is to live with.
- **Keep `loop.sh` and add a `loop` verb to `knack` that execs it.** Rejected — a shim that just re-dispatches is indirection without benefit. Either the loop is in `knack` or it isn't.
- **Cut `run` entirely; batch mode is a future harness.** Rejected — ADR-0009 keeps the loop as the batch companion; this decision is about where the binary lives, not whether the loop exists.

## Decision

**`knack` is the single bash entry point.** `loop.sh` is folded into it. The verb surface becomes:

- `knack spine` — derive the spine ADRs (unchanged).
- `knack adrs` — list all ADRs (unchanged).
- `knack check` — lint structural drift (unchanged).
- `knack view` — read-only dashboard of cycles, work units, and decisions (was `loop.sh view`).
- `knack run <queue>` — the optional batch runner behind a verify gate (was `loop.sh run`).

`loop.sh` is deleted. Everything that referenced `./loop.sh ...` references `./knack ...` instead. The loop's optionality is unchanged (ADR-0009): you simply don't run `knack run` if you aren't doing batch work. The verify-gate ownership, `LOOP_AGENT_CMD` / `LOOP_REVIEW_CMD` / `LOOP_FIX_CMD` overrides, per-unit `Agent:` overrides, and `--review` / `--max-review-rounds` / `--max-ticks` / `--dry-run` flags all move with `run` unchanged.

ADR-0011 is amended: "the loop stays as `loop.sh`" becomes "the loop stays as a bash verb on `knack`." ADR-0017 is amended: "derives and lints; does not govern" becomes "derives, lints, inspects, and optionally runs the batch loop; does not scaffold, package, or version." The narrowing that mattered — no Go, no compile step, no packaging, no per-unit governance gates beyond the verify gate the loop already owns — survives. The narrowing that didn't — forbidding a `run` verb on the same binary as `spine` — is dropped.

## Consequences

- One help text, one entry point, one thing to install/symlink. `./knack help` lists all five verbs.
- `knack`'s role broadens from "derives and lints" to "the project's bash tooling." The line against scaffolding/packaging/versioning holds (those are skills.sh's job, per ADR-0011); the line against *governing* softens to "the only governance `knack` does is the verify gate the loop already owned."
- `tests/run.sh`, `AGENTS.md`, `README.md`, `docs/loop.md`, `docs/getting-started.md`, and any skill text that names `./loop.sh` must update to `./knack run` / `./knack view`.
- `knack check`'s own drift checks do not cover command-name references in docs, so the rename must be verified by grep and by running the test suite — same shape as the post-0011 doc-rewrite pass.
- Risk: a single binary with both read-only (`spine`/`adrs`/`check`/`view`) and side-effecting (`run`) verbs is a larger surface to reason about. Mitigation: verbs are disjoint, `run` is the only verb that writes outside `.loop/`, and the help text groups them (inspection vs. batch).
