---
nospec: true
id: 0022
date: 2026-07-26
status: accepted
spine: false
amends: [0005]
---

# 0022: Batch lint enforces the minimum work-unit contract

## Context

ADR-0005 gave work units their outcome, context, constraints, acceptance criteria, and verify shape, but deliberately left the runner parser unchanged. Dogfooding exposed the resulting mismatch: `nospec lint` accepted a unit with only a vague heading, `Verify: true`, and a status. Such a unit was structurally executable but did not carry enough acceptance information for safe unattended work.

Requiring placeholder context and constraints would replace that gap with ceremony. Some bounded outcomes need no special reading or boundary beyond the repository's existing operational context.

## Decision

Batch queue preflight enforces this minimum contract for every work unit:

- `## <outcome>` is nonempty.
- `Done means:`, `Verify:`, and `Status:` occur exactly once and are nonempty or valid for their field type.
- `Read first:` and `Constraints:` are optional. When present, each occurs once and contains nonempty content on following lines.
- `Verify:` remains a fenced Bash command with valid shell syntax. Lint rejects only mechanically obvious vacuity: a comment-only command or a standalone `true`, `:`, or `exit 0`.

Lint does not claim to understand whether an otherwise valid command discriminates the outcome. That remains shaping judgment and review surface.

## Consequences

- `nospec lint`, `run`, and `view` reject queues that omit acceptance criteria.
- Optional fields can be omitted instead of filled with `none` placeholders.
- The parser now reads shape fields in addition to the header, agent, verify, and status, amending ADR-0005's original parser boundary.
- Existing queues that predate this contract must add `Done means:` before they can run.
