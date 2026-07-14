# 0010: Skills transmit concepts and reasoning, not rules

Date: 2026-07-11
Status: accepted

## Context

ADR-0005 established that work units should carry *shape* (constraints, outcomes), not *scripts* (file-by-file edit lists), because *"over-prescriptive instructions override the model's native competence and make output worse."* That ruling was about the work-unit format. The same principle applies one level up — to how the skills themselves are written — and the project had not generalized it.

The current skills are largely procedural rulebooks: *"Procedure: 1. Read the codebase. 2. Grill the intent. 3. …"* They prescribe process. Rules are brittle: "always lead with a vertical slice" fails on a refactor, so it gets patched with exceptions, then exceptions to exceptions, producing a decision tree worse than understanding the concept ever was. The agent is strong at contextual application of *understood* concepts and weak at following arbitrary-seeming prescriptions.

The wiki's [prompts-in-code-review](https://github.com/bermudi/AgenticWiki/blob/main/wiki/threads/prompts-in-code-review.md) thread measured this directly: more prescriptive detail increases overcorrection; structured reasoning templates are the reliable fix. Prescribe less, reason more. That generalizes beyond code review to all skill authoring.

Not everything is judgment. Some things are mechanical contracts that exist precisely because the agent must not "judge" them — the verify gate runs outside the agent and must exit 0; that is the backpressure mechanism, not a suggestion. Folding those into "concepts and reasoning" would dissolve the load-bearing backbone.

Alternatives considered:

- **Prescriptive procedural skills (current).** Rejected — produces the ceremony problem at the skill level; brittle to situations the rule-writer didn't anticipate.
- **Pure concept dumps, no procedural scaffolding.** Rejected — risks under-specification for weaker models or common cases; a default that can be overridden is cheaper than rediscovering the common path every time.
- **Rules for everything, including judgment.** Rejected — this is the failure mode being escaped; it is what made litespec feel mandatory.

## Decision

Skills transmit **concepts and the reasoning behind them**, and let the agent apply judgment. They do not prescribe *when* to deploy a concept as a rule.

For each concept a skill teaches, it carries:

- **What it is** — the concept (e.g., a tracer bullet is a thin end-to-end slice to get early integration feedback).
- **Why it exists** — the failure mode it prevents (e.g., discovering the schema doesn't support the UI in week 3 instead of hour 1).
- **A pointer to depth** — the wiki link. The skill carries an operational synopsis (what the concept is, locally, in this project's terms); the wiki carries the full theory, evidence, and elaboration. Summarize briefly; do not duplicate the full theory. *(Refined by [ADR-0013](0013-wiki-links-live-in-docs-not-in-skill-text.md): the link lives in human-facing docs, not in skill text — the skill carries the synopsis/phrase; docs carry the link.)*
- **Optionally, a reasoned default** — a default behavior for the common case, always accompanied by the reasoning for *when to override* it. Defaults are scaffolding, never mandates. A default without its override-reasoning is just a rule in disguise.

The split that keeps the backbone intact:

- **Judgment** (decomposition, process choice, depth, when to explore) → concept + reasoning; the agent decides.
- **Mechanical contracts** (the verify gate is external and deterministic; hard stops) → hard rule, non-negotiable. These are mechanisms, not prompts.

Concretely: the `plan` skill will **not** say "lead with a vertical slice, then go horizontal." It will transmit the decomposition concepts (tracer bullet, vertical slice, horizontal/breadth), link each to the wiki, state the failure mode each prevents, and stop. The agent chooses. This reconciles with — and generalizes — ADR-0005: per-unit *format* is free (0005); *decomposition* is concept-guided (this ADR). They are different axes the old docs conflated.

The glossary is the same ruling applied to vocabulary: knack-domain terms (work unit, verify gate, tick) get defined; wiki concepts (doc-rot, ralph-loop, backpressure) get a one-line local pointer plus a link — the full theory lives in the wiki, not duplicated here. The glossary stops competing with the wiki.

## Consequences

- Skills become lighter and more general — less to maintain, less to rot.
- The ceremony problem is attacked at its root: ceremony *is* prescribed process. Concept-guided judgment produces the right behavior per situation without a rulebook that must be followed-or-rotted.
- Every existing skill becomes a rewrite candidate (most are currently procedure lists). That rework is a separate pass; this ADR records the authoring law they'll be written against.
- `glossary.md` must be stripped of wiki concepts and pointed at the wiki.
- Risk: under-prescription for some models. Mitigation is reasoned defaults (default + override-reasoning), and the wiki evidence says reasoning empirically beats prescription.
- Extends ADR-0005; does not supersede it (0005 governs work-unit format and remains valid).
