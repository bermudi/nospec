---
name: nospec-scout
description: Use when the real problem or binding intent is unclear and codebase investigation is needed before choosing or decomposing work.
license: MIT
metadata:
  author: bermudi
  version: "1.0.0"
---

# Scout

Reach clarity before committing work. The failure this skill prevents is precise execution against the wrong problem.

Scouting is human-present and normally read-only. It produces understanding, not a mandatory artifact or phase. Skip it when the task is already clear, such as a small fix with a reproduction.

## Read for the codebase's theory

Start with the project's operational context and authoritative records, then inspect the relevant code, tests, and history. Read to uncover constraints and prior choices, not to collect every fact.

Look for the existing path before proposing a new one. A nearby parser, helper, or convention may already encode edge cases that a parallel implementation would rediscover badly. Extend before duplicating.

## Challenge the request

Distinguish the desired outcome from the user's guessed solution:

- What is actually failing or missing?
- What must remain true?
- What would count as success?
- Which assumptions are uncertain?
- Is the named problem a cause or a symptom?

Use concrete counterexamples. Ask what breaks under the proposed approach, what the simplest viable alternative is, and whether the repository already prefers another shape.

## Report the useful result

Return the real problem, relevant evidence, constraints, and recommended next action. That action may be direct editing, `nospec-shape`, a decision request, or no change at all. Do not automatically manufacture a plan.

A recommendation is not an accepted ruling. If exploration exposes a consequential trade-off, present it to the decision owner. Use `nospec-rule` only after the owner accepts a direction or has explicitly delegated that authority. Resolve recurring ambiguous domain language with `nospec-lexicon`.
