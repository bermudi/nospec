---
name: domain-modeling
description: Use when domain terms surface during explore, plan, or review and need to be defined, challenged, or stress-tested against the project's ubiquitous language. Manages `glossary.md` — the project's shared vocabulary. Triggers on "what does X mean here", "are these the same concept", "let's define our terms", "that term is ambiguous", or when a domain concept is used inconsistently across a conversation or codebase.
---

# Domain Modeling

Manage the project's ubiquitous language. Terms used in code, specs, conversations, and skills should mean the same thing everywhere. The glossary is the single source of truth for that shared meaning, and a durable part of the shared design concept.

This is a shared skill — called by explore, plan, and review. It is not a separate phase.

## The glossary

The glossary lives in `glossary.md` at the project root. It is a small, curated file — not an encyclopedia. Each entry defines one term in one or two sentences.

If `glossary.md` doesn't exist, don't create one preemptively. Create it when the first term actually needs defining.

## Format

````markdown
# Glossary

## <Term>

<One to two sentences. What the term means in this project, not in general.>

## <Another Term>

<...>
````

Alphabetical order. No categories, no nesting. Flat and scannable.

## When to update the glossary

- **A term is used inconsistently** — two parts of the codebase or conversation use the same word for different things.
- **A term is introduced** — a new domain concept surfaces that will recur in future work.
- **A term is challenged** — someone asks "what does X mean here?" and the answer isn't obvious.
- **A term is overloaded** — one word is doing too much work and should be split.

Do **not** add:
- Terms with no project-specific meaning (don't define "database" or "API")
- Terms that appear once and won't recur
- Terms that are obvious from the code itself

## Stress-testing terms

When a term feels slippery, stress-test it with edge-case scenarios:

1. State the term and its current definition.
2. Pose an edge case: "Is X a <term> if it <unusual condition>?"
3. If the answer is unclear, the definition needs sharpening — not the edge case.
4. Rewrite the definition until the edge case has a clear answer.
5. If the term can't be sharpened, it may be two concepts masquerading as one. Split it.

## Procedure

1. Notice that a domain term is being used — either by you, the user, or in the codebase.
2. Check `glossary.md` (if it exists). Is the term already defined?
3. If yes — does the current usage match the definition? If not, either update the definition or correct the usage.
4. If no — does this term warrant a glossary entry? (See "When to update" above.)
5. If yes — add it. Keep it to one or two sentences. State the project-specific meaning, not the dictionary meaning.
6. If the term is slippery, stress-test it before committing to a definition.

## Delegation

Other skills call this skill inline. If you're exploring a codebase and notice a term being used inconsistently, define it now, then continue exploring. Don't queue it.
