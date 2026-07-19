---
id: 0017
date: 2026-07-18
status: accepted
spine: false
amends: [0011]
builds_on: [0015]
---

# 0017: Derivable artifact metadata via a bash CLI

## Context

ADR-0015 established that each fact has one owner and other documents are projections. ADR-0016 conceded that the pin-state record catches *direct* drift (a pinned doc changed) but not *indirect* coherence failure (A changed in a way that contradicts unpinned B). The project's own headline concept — the "spine" of load-bearing ADRs — was exactly that failure:

- `AGENTS.md` listed the spine as ADR-0009 through 0016 (seven entries).
- `docs/architecture.md` listed 0009 through 0015 (six, missing 0016) — then cited 0016 two screens later.
- `README.md` listed 0009, 0010, 0011, 0014, 0015 (five, missing 0012, 0013, 0016).

Three views, three different answers to "what is the spine." The spine-list is one fact with three owners and no record — a direct ADR-0015 violation that ADR-0016's pin-check cannot catch because no pin was set on the list itself. The `document` skill is the judgment-based remedy, but judgment failed to fire while the lists silently diverged.

ADR-0011 deleted the Go CLI and moved coherence-hygiene into skills as transmitted concepts. That ruling was correct for the *governance* apparatus (scaffolding, packaging, versioning, per-unit gates). But it left the project with no mechanical way to derive facts that are *already in the artifacts* — the spine subset, the ownership map, the role of each doc. These are not judgment; they are metadata that the artifacts already carry in ad-hoc prose (`Status:`, `Supersedes:`, `Grandfathered:`) or not at all (`AGENTS.md`, `glossary.md`, `README.md`).

The problem is recurring. Any fact stated in prose across multiple docs will drift. The fix is not better vigilance; it is to make the fact derivable from a single source.

## Decision

### 1. Frontmatter standard

All durable artifacts carry YAML frontmatter between `---` fences at the top of the file, before the first heading. Disposable artifacts (`.loop/<name>/QUEUE.md`, `HANDOFF.md`, `REVIEW.md`, `specs/`) do not carry frontmatter.

**ADRs** (`decisions/*.md`):

```yaml
---
id: 0016
date: 2026-07-18
status: accepted
spine: true
supersedes: [0006]        # optional
superseded_by: [0011]     # optional
amends: [0010]            # optional
grandfathered: <reason>   # optional
---
```

- `id` — the four-digit ADR number, zero-padded.
- `date` — ISO date.
- `status` — `accepted` | `superseded`.
- `spine` — `true` if this ADR is load-bearing (part of the curated spine); `false` otherwise. Default `false`. This is the single source of truth for spine membership.
- `supersedes` / `superseded_by` / `amends` / `builds_on` — lists of ADR ids. `supersedes`/`superseded_by` mark replacement (the old ADR's status becomes `superseded`); `amends` marks a narrowing or extension of an earlier ruling (the amended ADR's status stays `accepted`); `builds_on` marks a non-amending derivation (the new ruling builds on the old but doesn't change it). Replace the prose `Supersedes:` / `Superseded by:` / `Amends:` lines.
- `grandfathered` — free-text reason, replacing the prose `Grandfathered:` line.

The existing prose fields (`Date:`, `Status:`, etc.) are removed; their content moves into frontmatter.

**Skills** (`skills/<name>/SKILL.md`): keep existing `name` and `description` frontmatter. No changes required; skills are implicitly records (ADR-0015).

**Other durable docs** (`AGENTS.md`, `glossary.md`, `README.md`, `docs/*.md`):

```yaml
---
role: record
owns: operational-context
---
```

- `role` — `record` | `view` | `ledger`. Mechanizes the ADR-0015 role table.
- `owns` — (records only) the claim-class this artifact owns. Used by `knack check` to enforce uniqueness. Two records with the same `owns` value is a failure.

### 2. The `knack` CLI

A bash-only executable at the repo root. No Go, no compile step, no dependencies beyond `bash`, `grep`, `awk`, `sed`. It derives facts from frontmatter; it does not govern, scaffold, package, or version.

Subcommands:

- `knack spine` — prints the spine ADRs (id, title, one-line synopsis) derived from `decisions/*.md` where `spine: true`. This is the single source of truth for the spine list. Docs that want to present the spine reference this output or `decisions/`, not a re-typed list.
- `knack check` — exits non-zero if any of:
  - A durable doc outside `decisions/` re-enumerates the spine: two or more ADR numbers in a single line, or three or more lines containing ADR-numbered list entries in a block. Individual ADR references ("see ADR-0010") are fine; re-listing the spine as a curated subset is not.
  - Two records declare the same `owns` claim-class.
  - A durable doc is missing required frontmatter fields for its role.
- `knack adrs` — prints all ADRs (id, status, title). For browsing.

The CLI is agent-agnostic and has no opinions about process. It is a derivation and linting tool, not a gate that the loop runs per-tick. `knack check` may be added to `tests/run.sh` as a mechanical check, but it does not replace the judgment-based `document` skill for semantic coherence.

### 3. Views do not re-enumerate the spine

`README.md`, `docs/architecture.md`, and any other view that previously listed ADR numbers as a curated subset replaces that list with a pointer to `decisions/` and/or `knack spine`. Individual ADR references ("see ADR-0010") are fine; re-listing the spine as a subset is not.

## What this does not do

- Does not restore the deleted Go CLI (ADR-0011). No scaffolding, packaging, versioning, manifests, coverage gates, or per-unit gates.
- Does not make coherence mechanical. `knack check` catches the *structural* failure (re-enumerated lists, duplicate ownership) — not semantic contradiction. The `document` skill remains the judgment-based remedy for indirect coherence failure (ADR-0016).
- Does not replace the `document` skill. The CLI is a tool the skill (and the human) can invoke; it does not transmit the concept of when to invoke it.

## Consequences

- ADR-0011 is amended: "no CLI" becomes "no Go CLI; a bash derivation CLI is in scope." The amendment is narrow — the CLI derives and lints, it does not govern.
- All 16 existing ADRs get YAML frontmatter; their prose `Date:`/`Status:`/`Supersedes:`/`Grandfathered:` lines are removed.
- `AGENTS.md`, `glossary.md`, `README.md`, and `docs/*.md` get `role` (+ `owns` for records) frontmatter.
- The three drifting spine lists are consolidated: one source (`knack spine` / `decisions/` frontmatter), zero re-enumerations.
- `knack check` is added to `tests/run.sh`.
- Future ADRs must include `spine: true|false` in frontmatter. The `decide` skill is updated to reflect this.

## Related

- ADR-0011 — ship via skills.sh; delete CLI (amended: bash derivation CLI is in scope)
- ADR-0015 — artifact roles and ownership (mechanized: `role` and `owns` frontmatter)
- ADR-0016 — proof-boundary is mechanical; pin-state is provenance (complemented: `knack check` catches structural drift the pin-check cannot)
