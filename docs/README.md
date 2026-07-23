---
role: view
---

# nospec documentation

This folder holds the living user documentation for **nospec**. Each doc has one job; if you are looking for something specific, use the map below.

## What is nospec?

See [`docs/architecture.md`](./architecture.md) for the conceptual shape, human-attention spectrum, and artifact-role model.

## Where to find what

| Question | Owning document |
|---|---|
| What is nospec, and why does it exist? | [`README.md`](../README.md) |
| Why was an architectural choice made? | [`decisions/`](../decisions/) |
| How do I install and use the skills? | [`getting-started.md`](./getting-started.md) |
| How do the skills compose? | [`skills.md`](./skills.md) |
| How does the batch loop work? | [`loop.md`](./loop.md) |
| What is the `QUEUE.md` format? | [`queue-format.md`](./queue-format.md) |
| What does a nospec term mean? | [`glossary.md`](../glossary.md) |
| How do I work on the nospec repo? | [`AGENTS.md`](../AGENTS.md) |
| Where is the historical reasoning and wiki grounding? | [`theory.md`](./theory.md) |
| Where are durability and documentation concepts? | [`skills/nospec-curator/`](../skills/nospec-curator/SKILL.md) and ADR-0015 |
| How do I check artifact coherence or derive the spine? | `nospec check` / `nospec spine` (ADR-0017) |

## Editing these docs

Do not duplicate claims across docs. Put the claim in the authoritative record and project it in a view. If you change a record, update the views that link to it. If you are unsure where something belongs, load the `nospec-curator` skill.
