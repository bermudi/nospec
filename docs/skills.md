---
nospec: true
role: view
---

# Skills guide

Skills are agent-agnostic procedural knowledge stored in `.agents/skills/<name>/SKILL.md`. Any agent that supports agentskills.io can discover them automatically. The loop names the skill explicitly in the worker prompt; the worker's harness auto-loads it by trigger text, same as any skill invocation. See ADR-0007 and ADR-0019.

## Skill catalog

Clear work can start directly with `nospec-carve`; the other skills are selected only when their specific problem is present. Agents discover and load each skill independently.

| Skill | Purpose |
|---|---|
| `nospec-carve` | Implement one bounded outcome, conversational or queued; verify before claiming success. |
| `nospec-scout` | Investigate when the real problem or binding intent is unclear. Read-only; no mandatory artifact. |
| `nospec-shape` | Decompose intent into verifiable outcomes; serialize a queue only for handoff or batch. |
| `nospec-trial` | Run two-axis adversarial review (standards + intent) and generate findings. |
| `nospec-mend` | Resolve supported findings directly or as new work units. |
| `nospec-rule` | Record an architectural ruling after explicit acceptance or delegated authority. |
| `nospec-lexicon` | Define and update `glossary.md` terms. |
| `nospec-curator` | Route lasting knowledge to its authoritative record and maintain coherent projections. |
| `nospec` *(optional)* | Explain batch-worthiness and carry the runner that supplies an external verify gate while the human is absent. |

## Skill format

A skill is a Markdown file named `SKILL.md` inside a directory named after the skill:

```text
.agents/skills/
└── nospec-carve/
    └── SKILL.md
```

Required frontmatter:

```yaml
---
name: nospec-carve
description: Use when implementing one work unit...
---
```

The `name` must match the directory name. The `description` is the trigger text used by agents to decide when to invoke the skill.

## How the loop uses skills

`nospec run` does not read skills itself. It prepends the worker prompt (from `skills/nospec/prompts/worker.md`) to the current work unit and runs the worker. The worker prompt tells the worker to load the `nospec-carve` skill by name; the worker's harness auto-loads it by trigger text, same as any skill invocation. No path configuration is needed — the worker is a harness session, and harnesses find their own skills (ADR-0019).

When `--review` is set, the loop also invokes review and fix workers after the build queue drains. Those prompts tell the worker to load the `nospec-trial` or `nospec-mend` skill directly. The loop orchestrates the bounded review/fix subloop, reads the actionable count from `REVIEW.md`, and runs another build pass when fix appends pending units. The skills still own judgment: nospec-trial decides what the findings are, and nospec-mend decides which findings become work units.

Without `--review`, nospec-trial and nospec-mend remain manual skill invocations.

## Customizing skills

After `npx skills add`, the project owns the `.agents/skills/` directory. Edit, override, or delete skills as needed. The repo's `skills/` directory is the source; `npx skills update` refreshes the local copies.

## Composable use

There is no default pipeline. Invoke the narrow skill whose problem is present: scout uncertainty, shape outcomes, carve one bounded change, trial an existing diff, or mend known findings. Use `nospec-rule` only for an accepted ruling, `nospec-lexicon` for unresolved domain language, and `nospec-curator` for durable-context ownership or coherence.
