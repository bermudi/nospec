---
role: view
---

# Skills guide

Skills are agent-agnostic procedural knowledge stored in `.agents/skills/<name>/SKILL.md`. Any agent that supports agentskills.io can discover them automatically. The loop names the skill explicitly in the worker prompt; the worker's harness auto-loads it by trigger text, same as any skill invocation. See ADR-0007 and ADR-0019.

## Default skills

| Skill | Purpose |
|---|---|
| `nospec-scout` | **Entry point.** Investigate a codebase, grill intent, and stress-test ideas before planning. Read-only, no artifacts, reaches clarity before any `QUEUE.md` is written. |
| `nospec-shape` | Convert intent into a disposable `QUEUE.md` of verifiable work units. |
| `nospec-hew` | Implement one work unit from `QUEUE.md`; do not self-certify. |
| `nospec-trial` | Run two-axis adversarial review (standards + intent) and generate findings. |
| `nospec-mend` | Convert review findings into new work units. |
| `nospec-rule` | Capture architectural rulings as ADRs in `decisions/`. |
| `nospec-lexicon` | Define and update `glossary.md` terms. |
| `nospec-curator` | Route knowledge to its authoritative artifact and maintain coherent projections. |
| `nospec` *(optional)* | The batch runner. Carries `scripts/nospec` and transmits the batch-mode concept — when to reach for AFK execution and what the verify gate guarantees. |

## Skill format

A skill is a Markdown file named `SKILL.md` inside a directory named after the skill:

```text
.agents/skills/
└── nospec-hew/
    └── SKILL.md
```

Required frontmatter:

```yaml
---
name: nospec-hew
description: Use when implementing one work unit...
---
```

The `name` must match the directory name. The `description` is the trigger text used by agents to decide when to invoke the skill.

## How the loop uses skills

`nospec run` does not read skills itself. It prepends the worker prompt (from `skills/nospec/prompts/worker.md`) to the current work unit and runs the worker. The worker prompt tells the worker to load the `nospec-hew` skill by name; the worker's harness auto-loads it by trigger text, same as any skill invocation. No path configuration is needed — the worker is a harness session, and harnesses find their own skills (ADR-0019).

When `--review` is set, the loop also invokes review and fix workers after the build queue drains. Those prompts tell the worker to load the `nospec-trial` or `nospec-mend` skill directly. The loop orchestrates the bounded review/fix subloop, reads the actionable count from `REVIEW.md`, and runs another build pass when fix appends pending units. The skills still own judgment: nospec-trial decides what the findings are, and nospec-mend decides which findings become work units.

Without `--review`, nospec-trial and nospec-mend remain manual skill invocations.

## Customizing skills

After `npx skills add`, the project owns the `.agents/skills/` directory. Edit, override, or delete skills as needed. The repo's `skills/` directory is the source; `npx skills update` refreshes the local copies.

## Composable flows

Skills are not a rigid gate. The default flow is `nospec-scout → nospec-shape → nospec-hew → nospec-trial → nospec-mend`, but any valid subset is fine:

```text
small fix → nospec-shape → nospec-hew → done
bug report → nospec-scout → nospec-shape → nospec-hew → done
big feature → nospec-scout → nospec-shape → nospec-hew --review → nospec-trial → nospec-mend → nospec-hew → done
```

Decisions are captured inline throughout the flow using the `nospec-rule` skill, terms are updated using `nospec-lexicon`, and durable-context placement is checked using `nospec-curator`.
