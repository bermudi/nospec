---
nospec: true
id: 0024
date: 2026-07-26
status: accepted
spine: false
amends: [0023]
builds_on: [0009, 0016]
---

# 0024: Batch runs record and guard the starting baseline

## Context

`nospec run` invoked workers directly in the selected repository without examining its starting state. Existing staged, unstaged, and untracked files were exposed to overwrite; verification could depend on pre-existing work; and evidence labelled the repository's whole status as files changed by a tick. The runner therefore provided an external verify gate without an honest attribution boundary.

Automatically creating a Git worktree would separate ordinary edits, but it would also require Nospec to own queue transfer, branch naming, result promotion, conflict handling, and cleanup. It would still not sandbox an unrestricted agent from the rest of the filesystem.

Plan-then-leave work may deliberately begin with human edits, so a blanket clean-tree requirement would also reject a legitimate handoff.

## Decision

Before a cycle's first mutating run, after queue preflight and before status mutation or worker invocation, the runner inspects the selected repository's Git status:

- paths under `.loop/` are operational cycle state and do not make the baseline dirty;
- staged, unstaged, or untracked paths elsewhere make the baseline unsafe;
- absence of a Git worktree also makes the baseline unsafe because Nospec cannot establish attribution or recoverability;
- an unsafe baseline stops the run by default;
- `--accept-dirty-baseline` explicitly accepts and records the unsafe state without claiming to protect it; and
- `--dry-run` reports baseline safety without writing evidence.

The accepted baseline is appended once to the cycle's `EVIDENCE.md`, including Git root, HEAD, and pre-existing paths. Tick entries describe the resulting **worktree state**, not files attributed to that tick.

Nospec does not automatically create worktrees. Operators can create one themselves and select it with the existing `--repo` option while leaving queue state in the original checkout. A worktree is workspace separation, not a security boundary; containment remains the responsibility of the harness, container, or operating system.

## Consequences

- AFK runs no longer silently begin on unrelated human changes.
- Deliberate dirty handoffs remain possible, but the risk and starting paths are durable evidence.
- Git remains optional only through explicit unsafe-baseline acceptance.
- Existing cycle dirt is allowed after the baseline marker is recorded so verified units can accumulate changes and failed units can resume.
- Nospec does not acquire branch promotion, merge, rollback, or worktree-cleanup machinery.
