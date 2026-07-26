---
nospec: true
id: 0023
date: 2026-07-26
status: accepted
spine: false
amends: [0016]
---

# 0023: Evidence does not interpret verification commands

> **Starting-state attribution added by [ADR-0024](0024-batch-runs-record-and-guard-the-starting-baseline.md):** evidence records the accepted cycle baseline and labels each tick's repository listing as resulting worktree state rather than changes attributed to that tick.

## Context

ADR-0016 introduced a registry that split a Bash `Verify:` command on `&&` and translated recognized fragments into prose such as "Go test suite passes" or "file exists." Unknown fragments fell back to their command text.

That translation added confidence without adding evidence. Shell composition, wrappers, working-directory changes, quoting, and project-specific commands make syntactic classification unreliable. The exact command and output were already present beside the generated narration, so the registry duplicated the source while sometimes overstating it.

## Decision

`EVIDENCE.md` records:

- the exact work unit;
- changed files;
- the exact `Verify:` command;
- verify output;
- worker output; and
- a conservative proof boundary based only on the external result.

For a completed unit, the proof boundary states that the recorded command exited 0 in that repository state and that no broader behavior is inferred. For any other status, it states that no successful external verification is recorded.

The runner does not split, classify, summarize, or otherwise interpret verification shell syntax. The command itself owns its scope; `Done means:` versus `Verify:` remains the review surface.

## Consequences

- The primitive proof-claim registry and its embedded Python implementation are removed.
- Evidence becomes shorter and cannot turn command names into unsupported behavioral claims.
- Readers must inspect the exact command to understand what was exercised. This is deliberate: a generic runner cannot infer semantic coverage from shell syntax.
- ADR-0016's pin-state ruling is unchanged.
