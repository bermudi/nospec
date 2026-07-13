# 0007: Name the skill explicitly in the worker prompt

Date: 2026-07-09
Status: accepted

## Context

DESIGN.md open question #4 asked how the loop gets the right skill into the agent: does the agent discover `.agents/skills/build/SKILL.md` on its own via agentskills.io progressive disclosure (trigger-based), or does the loop have to point at it? The loop is skill-agnostic — it never reads skill files and only knows a skill name to inject into a prompt.

Testing across the target agents (the Devin-driven CLI cycle and the default-`pi` path) showed trigger-based discovery is not reliable enough: discovery depends on each agent's own metadata-loading behavior and timing, which varies, and the loop has no way to confirm the agent actually loaded the skill. A worker that silently fails to load `build` does the wrong work, and nothing in the loop catches a skill that was never loaded.

Alternatives considered:
- **Rely on trigger-based discovery.** Rejected — inconsistent across agents; the loop cannot verify the skill loaded.
- **Inline the full skill body in the prompt.** Rejected — duplicates the on-disk skill, drifts from the source of truth, bloats the prompt.
- **A per-agent `--skill` flag in `LOOP_AGENT_CMD`.** Rejected — not every agent has one, and it couples the loop to each agent's CLI surface.

## Decision

The worker prompt (`prompts/worker.md`) names the skill explicitly — name *and* on-disk path (e.g. "Load and follow the **build** skill in `.agents/skills/build/`"). The loop passes this prompt to the agent via the `LOOP_PROMPT_FILE` environment variable (or inline) and trusts the agent's file-reading to load it. The loop never reads skill files and needs no per-agent `--skill` flag.

Naming the path — not just the name — removes dependence on each agent's discovery timing: the worker can read the skill file directly even if progressive disclosure never surfaces it.

## Consequences

- `prompts/worker.md` is the single place the worker skill is named; changing it means editing one prompt.
- The `.agents/skills/build/` path in the example above was the layout at decision time; [ADR-0011](0011-ship-as-skills-via-skills-sh-delete-cli.md) relocated skills to `skills/` at the repo root, and the prompt now reads `skills/build/`. (That path resolves against the worker's cwd — the target repo — so it relies on the skills being installed there; making it absolute is a future loop.sh hardening, relevant once the project publishes and runs against external repos.)
- Agent integrations (`LOOP_AGENT_CMD`) only need to consume the prompt file; they don't implement skill discovery. Agent-agnosticism holds.
- A worker that ignores the prompt and never loads the skill is still possible. The backpressure is the verify gate, not skill loading — and that is the right thing to enforce mechanically (the gate catches wrong output, which is what matters).
