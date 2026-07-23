---
name: nospec-lexicon
description: Use when domain terms surface during exploration, planning, building, or review and need to be defined, challenged, or stress-tested against the project's ubiquitous language. Manages `glossary.md` — the project's shared vocabulary. Triggers on "what does X mean here", "are these the same concept", "let's define our terms", "that term is ambiguous", or when a domain concept is used inconsistently across a conversation or codebase.
license: MIT
metadata:
  author: bermudi
  version: "1.0.0"
---

# Domain Modeling

Manage the project's ubiquitous language — the shared vocabulary that lets human and agent mean the same thing by the same word, in code, specs, conversation, and skills. A shared glossary is the protocol for precision between human and agent; without it, the agent drifts toward competing terminology and slop.

This is a shared skill — called inline by `nospec-scout`, `nospec-shape`, `nospec-hew`, and `nospec-trial` whenever a term needs defining, in any attention-mode. It is not a separate phase. If a term is being used inconsistently, define it now, then continue.

## The glossary

The glossary lives in `glossary.md` at the project root. It is small and curated — not an encyclopedia. Each entry defines one term in one or two sentences, stating what it means *in this project*, not in general. If `glossary.md` doesn't exist, don't create one preemptively — create it when the first term actually needs defining.

## Format

The glossary is the project's ubiquitous language — one `## <Term>` per entry, one or two sentences stating what it means *in this project*, not in general:

````markdown
## <Term>

<One to two sentences. What the term means in this project, not in general.>
````

Keep it scannable. A flat list of definitions is usually enough; split into sections only if the project's vocabulary genuinely falls into distinct kinds.

## When to update the glossary

- **A term is used inconsistently** — two parts of the codebase or conversation use the same word for different things.
- **A term is introduced** — a new domain concept surfaces that will recur.
- **A term is challenged** — someone asks "what does X mean here?" and the answer isn't obvious.
- **A term is overloaded** — one word is doing too much work and should be split.
- **A term has gone stale** — it no longer belongs to the project's current domain model, or its definition no longer matches current usage. Remove the entry (or sharpen it). Search results (code, docs, conversation) are evidence of staleness, not the test — a term can appear everywhere and still have a stale definition.

Don't add terms with no project-specific meaning (don't define "database" or "API"), terms that appear once and won't recur, or terms obvious from the code itself.

After adding, removing, or redefining a term, invoke the `nospec-curator` skill to check whether `AGENTS.md`, the relevant skills, or other durable docs use the old meaning in a projection.

## Stress-testing terms

When a term feels slippery, stress-test it with edge cases — the way ubiquitous-language practice road-tests definitions against concrete scenarios. Pose an edge case: "Is X still a `<term>` if it `<unusual condition>`?" If the answer is unclear, the definition needs sharpening, not the edge case. Rewrite it until the edge case has a clear answer. If it can't be sharpened, the term may be two concepts masquerading as one — split it.

## Delegation

Other skills call this skill inline. A term that needs defining is defined now, not deferred.
