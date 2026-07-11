# Skills guide

Skills are agent-agnostic procedural knowledge stored in `.agents/skills/<name>/SKILL.md`. Any agent that supports agentskills.io can discover them automatically. The loop names the skill explicitly — name and path — in the worker prompt (`prompts/worker.md`); the agent reads the skill file directly. See ADR-0007.

## Default skills

| Skill | Purpose |
|---|---|
| `explore` | **Entry point.** Investigate a codebase, grill intent, and stress-test ideas before planning. Read-only, no artifacts, reaches clarity before any `QUEUE.md` is written. |
| `plan` | Convert intent into a disposable `QUEUE.md` of verifiable work units. |
| `build` | Implement one work unit from `QUEUE.md`; do not self-certify. |
| `review` | Run two-axis adversarial review (standards + intent) and generate findings. |
| `fix` | Convert review findings into new work units. |
| `decide` | Capture architectural rulings as ADRs in `decisions/`. |
| `domain-modeling` | Define and update `glossary.md` terms. |

## Skill format

A skill is a Markdown file named `SKILL.md` inside a directory named after the skill:

```text
.agents/skills/
└── build/
    └── SKILL.md
```

Required frontmatter:

```yaml
---
name: build
description: Use when implementing one work unit...
---
```

The `name` must match the directory name. The `description` is the trigger text used by agents to decide when to invoke the skill.

## How the loop uses skills

`loop.sh` does not read skills. It prepends `prompts/worker.md` to the current work unit and runs the worker. `prompts/worker.md` tells the worker to load the `build` skill by name and path (e.g. "Load and follow the **build** skill in `.agents/skills/build/`").

When `--review` is set, the loop also invokes review and fix workers after the build queue drains. Those prompts tell the worker to load the `review` or `fix` skill directly. The loop orchestrates the bounded review/fix subloop, reads the actionable count from `REVIEW.md`, and runs another build pass when fix appends pending units. The skills still own judgment: review decides what the findings are, and fix decides which findings become work units.

Without `--review`, review and fix remain manual skill invocations.

## Customizing skills

- After `knack skills init`, the project owns the `.agents/skills/` directory.
- Edit, override, or delete skills as needed.
- The CLI embeds the default skills. If you edit the defaults in the `knack` repo, run `cli/sync-skills.sh` to copy them into `cli/embedded/skills` and `diff -r .agents/skills cli/embedded/skills` to verify sync.
- Use `knack skills check` to validate your local skills.

## Composable flows

Skills are not a rigid gate. The default flow is `explore → plan → build → review → fix`, but any valid subset is fine:

```text
small fix → plan → build → done
bug report → explore → plan → build → done
big feature → explore → plan → build --review → review → fix → build → done
```

Decisions are captured inline throughout the flow using the `decide` skill, and terms are updated using `domain-modeling`.
