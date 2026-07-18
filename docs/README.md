# knack documentation

This folder holds the living user documentation for **knack**. Each doc has one job; if you are looking for something specific, use the map below.

## What is knack?

See [`docs/architecture.md`](./architecture.md) for the conceptual shape, human-attention spectrum, and artifact-role model.

## Where to find what

| Question | Owning document |
|---|---|
| What is knack, and why does it exist? | [`README.md`](../README.md) |
| Why was an architectural choice made? | [`decisions/`](../decisions/) |
| How do I install and use the skills? | [`getting-started.md`](./getting-started.md) |
| How do the skills compose? | [`skills.md`](./skills.md) |
| How does `loop.sh` work? | [`loop.md`](./loop.md) |
| What is the `QUEUE.md` format? | [`queue-format.md`](./queue-format.md) |
| What does a knack term mean? | [`glossary.md`](../glossary.md) |
| How do I work on the knack repo? | [`AGENTS.md`](../AGENTS.md) |
| Where is the historical reasoning and wiki grounding? | [`theory.md`](./theory.md) |
| Where are durability and documentation concepts? | [`skills/document/`](../skills/document/SKILL.md) and ADR-0015 |

## Editing these docs

Do not duplicate claims across docs. Put the claim in the authoritative record and project it in a view. If you change a record, update the views that link to it. If you are unsure where something belongs, load the `document` skill.
