---
nospec: true
role: view
---

# Skills guide

Skills are agent-agnostic procedural knowledge stored in `.agents/skills/<name>/SKILL.md`. Any agent that supports agentskills.io can discover them automatically. The loop names the skill explicitly in the worker prompt; the worker's harness auto-loads it by trigger text, same as any skill invocation. See ADR-0007 and ADR-0019.

## Skill catalog

Nospec provides three stances for verified change — Shape the intent, Carve the implementation, and Trial the result — plus an optional Curator that preserves what should outlive the work and an optional Loop that runs shaped work while the human is absent. Each skill is independently invoked; their names do not define a workflow.

| Skill | Stance | Purpose |
|---|---|---|
| `nospec-shape` | Understand and bound the work before implementation | Resolve unclear intent (read, challenge, or run a bounded experiment), decompose into verifiable outcomes, serialize a queue only for handoff or batch. |
| `nospec-carve` | Build or correct production code | Implement one bounded outcome, conversational or queued; resolve an accepted review finding; verify before claiming success. |
| `nospec-trial` | Adversarially challenge an existing change | Run two-axis review (standards + intent) and generate findings. |
| `nospec-curator` *(optional)* | Preserve newly crystallized durable knowledge | Route a lasting claim to its authoritative record (ADR, glossary, operational context, contract record) and reconcile stale projections. |
| `nospec-loop` *(optional)* | Operate unattended execution behind external verification | Judge batch-worthiness and carry the runner that supplies the verify gate while the human is absent. |

Clear work can start directly with `nospec-carve`. Reach for `nospec-shape` when the problem is unclear or outcomes need decomposition, `nospec-trial` when a change needs adversarial review, `nospec-curator` when a lasting claim appears, and `nospec-loop` only for AFK execution.

## Progressive disclosure

Each skill carries its core stance in `SKILL.md` and moves specialized modes, formats, and mechanics into `references/` loaded only when their branch applies. A reference needs an explicit condition; if every invocation needs it, its content belongs in `SKILL.md`. References stay inside their owning skill, preserving self-containment (ADR-0013).

```text
.agents/skills/
└── nospec-shape/
    ├── SKILL.md
    └── references/
        ├── scouting.md       # loaded when an executable experiment is needed
        └── queue-format.md   # loaded when serializing batch work
```

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

`nospec run` does not read skills itself. It prepends the worker prompt (from `skills/nospec-loop/prompts/worker.md`) to the current work unit and runs the worker. The worker prompt tells the worker to load the `nospec-carve` skill by name; the worker's harness auto-loads it by trigger text, same as any skill invocation. No path configuration is needed — the worker is a harness session, and harnesses find their own skills (ADR-0019).

When `--review` is set, the loop also invokes review and fix workers after the build queue drains. The reviewer prompt loads `nospec-trial`; the fixer prompt loads `nospec-shape` and appends one bounded unit per actionable finding. The loop orchestrates the bounded review/fix subloop, reads the actionable count from `REVIEW.md`, and runs another build pass when fix appends pending units. The skills still own judgment: `nospec-trial` decides what the findings are, and the `nospec-shape`-driven fixer decides which findings become work units.

Without `--review`, `nospec-trial` remains a manual skill invocation.

## Customizing skills

After `npx skills add`, the project owns the `.agents/skills/` directory. Edit, override, or delete skills as needed. The repo's `skills/` directory is the source; `npx skills update` refreshes the local copies.

## Composable use

There is no default pipeline. Invoke the narrow skill whose problem is present: shape uncertainty and outcomes, carve one bounded change, trial an existing diff, or curate a lasting claim. Use `nospec-loop` only when you intend to run shaped work while absent.
