---
id: 0019
date: 2026-07-20
status: accepted
spine: false
amends: [0007, 0011, 0017, 0018]
builds_on: [0009]
---

# 0019: Bundle the runner as the ninth skill

## Context

ADR-0011 ruled "ship as a skills collection via skills.sh; delete the Go CLI." The runner survived as a bash binary (ADR-0017, ADR-0018), but it lived at the repo root (`./knack`) while the skills lived in `skills/`. `npx skills add` installs skill directories — it does not touch PATH, does not install binaries, does not run postinstall hooks. So the runner had no distribution channel: a user who ran `npx skills add` got the eight skills but not the runner, and a user who cloned the repo got the runner but not the installable skills. Two channels for one product — or worse, one channel and a loose end.

A second gap: ADR-0007 ruled "the worker prompt names the skill explicitly — name and path." The "by path" part was a workaround for non-harness workers (raw shell commands that don't auto-load skills). But `LOOP_AGENT_CMD` drives real harnesses (`pi`, `claude`, `codex`, `devin`...) — and harnesses auto-load skills by trigger text. The path was solving a problem that doesn't exist when the worker is a harness session.

Both gaps have the same root cause: the runner and the skills were not co-located, and the skill-loading model was wrong.

## Decision

**The runner ships as the ninth skill, `knack`, alongside the other eight.** One distribution channel (`npx skills add`), one install, one source of truth for where everything lives. The worker loads skills by name, not by path — harnesses do the resolution.

### 1. Layout

- `skills/knack/SKILL.md` — the skill metadata and the batch-mode concept (when to reach for AFK execution, what the verify gate guarantees, what it doesn't).
- `skills/knack/scripts/knack` — the bash runner (was `./knack` at the repo root).
- `skills/knack/prompts/` — worker, reviewer, fixer prompts (was `prompts/` at the repo root).

The repo-root `knack` and `prompts/` are gone. After `npx skills add`, the runner lives at `.agents/skills/knack/scripts/knack` (project-local) or `~/.agents/skills/knack/scripts/knack` (global, with `-g`).

### 2. Skill loading: by name, not by path (amends ADR-0007)

The worker/reviewer/fixer prompts name the skill (`build`, `review`, `fix`) without a path. The worker is a harness session invoked via `LOOP_AGENT_CMD`; harnesses scan their skills directory and auto-load skills by trigger-text match against the `description` frontmatter field. This is the same mechanism that makes any other skill invocation work — `npx skills add` detects the harness precisely so the harness can find its skills.

ADR-0007's "name and path" was a workaround for non-harness workers that doesn't apply when `LOOP_AGENT_CMD` drives a real agent. The "name" part of ADR-0007 survives (trigger-based discovery isn't reliable enough across agents, so the prompt names the skill explicitly rather than hoping the harness picks it up); the "path" part is dropped.

No path-injection mechanism in the runner. No `@SKILLS_DIR@` placeholder. No skill-directory resolution. The runner prepends the prompt template and runs the worker; the harness does the rest.

### 3. PATH setup: `knack install`, agent-invoked

skills.sh installs skill files but does not touch PATH. The runner gains an `install` verb that symlinks itself onto PATH:

```bash
.agents/skills/knack/scripts/knack install
```

It picks the first writable directory on `PATH` (or `~/.local/bin` as fallback), symlinks `SCRIPT_DIR/knack` there, and reports the result. After that, `knack run ...` works from any directory.

The user doesn't type the long path. When the user asks the agent to set up knack, the agent invokes `scripts/knack install` from the skill's directory (the agent knows that path because it just loaded the skill). The agent does the one mechanical step the user would otherwise have to do. This fits the thesis: the skill transmits the concept (batch mode, when to use it), and the agent executes the mechanical setup.

### 4. REPO_ROOT default

`knack spine/adrs/check/view` operate on a target repo's `decisions/` and `docs/`. Previously `REPO_ROOT` defaulted to `SCRIPT_DIR` (the repo root, because the runner lived there). Now `SCRIPT_DIR` is inside the skill, so `REPO_ROOT` defaults to `pwd` — the repo you're in. The `--repo` flag overrides for when you're not in the target repo's root. This matches `cmd_view`'s existing behavior (it already defaulted to `pwd`).

### 5. The `knack` skill's content

The skill transmits the batch-mode concept — the AFK end of the attention spectrum (ADR-0009): when batch is the right attention level, what the verify gate mechanically guarantees, what it doesn't (coherence is the review surface, not the verify gate), and how to invoke the runner. It carries the runner as `scripts/knack` per the agentskills.io spec's `scripts/` directory. The runner is the mechanical companion; the skill text is the concept. This fits ADR-0010: the concept is transmitted, the mechanical contract (the runner) is bundled.

## Alternatives considered

- **Two channels: skills via `npx skills add`, runner via `npm install -g`.** Rejected — one product, two install commands, two package managers, two update paths. Splits what ADR-0009 keeps together ("skills are the product; the loop is an optional companion" — companion, not separate product).
- **Keep the runner at the repo root; users `git clone` separately for batch mode.** Rejected — same two-channel problem, worse because the runner needs its own distribution story (remote, license, versioning) that skills.sh already solves for the skills.
- **Bundle the runner in an existing skill (e.g., `build`).** Rejected — `build` is mode-independent procedural knowledge; bolting the runner onto it conflates the concept with the mechanism. The runner is its own thing with its own concept (batch mode).
- **Inject absolute skill paths into worker prompts (`@SKILLS_DIR@` substitution).** Rejected — solves a problem that doesn't exist. `LOOP_AGENT_CMD` drives harnesses, and harnesses auto-load skills by trigger text. Adding path-injection machinery would be dead code and would misrepresent the skill-loading model.
- **Document a `find`-based one-liner for the user to symlink the runner.** Rejected — fragile, copy-paste-bait, ages badly. The `install` verb is self-contained and agent-invoked.

## Consequences

- ADR-0007 is amended: "name and path" becomes "name only." The worker prompt names the skill; the harness auto-loads it. The "path" part was a workaround for non-harness workers that doesn't apply when `LOOP_AGENT_CMD` drives a real agent.
- ADR-0011 is amended: "ship via skills.sh" now includes the runner — it ships as the `knack` skill, not as a separate binary at the repo root. The "no Go CLI" narrowing survives; the bash runner is still bash, still no compile step.
- ADR-0017 is amended: the bash CLI's location is `skills/knack/scripts/knack`, not `./knack`. Its role (derives, lints, inspects, runs the batch loop) is unchanged; it gains an `install` verb for PATH setup.
- ADR-0018 is amended: "one command with verbs" survives — still one command, still five verbs (now six with `install`) — but its install path is the skill, not the repo root.
- `tests/run.sh` references the runner at its new path (`$root/skills/knack/scripts/knack`); `spine`/`adrs`/`check` calls use `--repo "$root"` because `REPO_ROOT` now defaults to `pwd`.
- `AGENTS.md`, `README.md`, `docs/`, and ADR text that referenced `./knack` or `prompts/` are updated to the new paths; references to "name and path" in the worker prompt are updated to "name only."
- The `knack` skill is the ninth skill. The first eight are mode-independent procedural concepts; the ninth is the batch-mode concept plus the runner. The skill count in docs ("eight skills") becomes "nine."
- Risk: users who want batch mode must run `knack install` once (directly or via the agent). Mitigation: the `knack` skill's `SKILL.md` and the README document the install verb; the agent runs it when asked to set up knack; the runner is also callable by full path without installing.
- Risk: a worker harness that doesn't auto-load skills by trigger would break. Mitigation: every harness `LOOP_AGENT_CMD` realistically drives (`pi`, `claude`, `codex`, `devin`, `opencode`) does auto-load skills — that's the whole point of `npx skills add` detecting the harness. A non-harness worker (raw shell script) would need the skill content embedded in the prompt directly, which is a different invocation pattern outside `knack run`'s scope.

## Related

- ADR-0007 — name the skill explicitly in the worker prompt (amended: "name only," not "name and path")
- ADR-0009 — skills are the product; the loop is an optional companion (the companion now ships with the product)
- ADR-0010 — skills transmit concepts, not rules (the `knack` skill transmits the batch-mode concept; the runner is the mechanical contract)
- ADR-0011 — ship via skills.sh; delete CLI (amended: the runner ships as a skill, not as a separate binary)
- ADR-0017 — derivable artifact metadata via bash CLI (amended: the CLI's location is inside the skill; gains an `install` verb)
- ADR-0018 — one command with verbs (amended: the command's install path is the skill; `install` is a new verb)
