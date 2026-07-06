# AGENTS.md

## Project

**sliceloop** (placeholder name — will be renamed) is an agent-agnostic harness for agentic development. It replaces litespec.

It is three separate artifacts with three separate concerns:
- **Skills** (`.agents/skills/`) — the workflow as procedural knowledge, agent-agnostic via agentskills.io
- **Loop** (`loop.sh`) — external bash script, agent-agnostic, owns the verify gate
- **CLI** (Go binary, TBD) — read-only validator + context provider

Code is the source of truth. Specs are disposable. Decisions and skills are durable.

## Thesis

Code is the source of truth. Specs rot. Plans are ephemeral coordination state. The reusable asset is procedural knowledge encoded as skills. The loop owns verification; the worker never self-certifies.

See `DESIGN.md` for the full design.

## Current state

The loop (`loop.sh`) and the `plan` skill are migrated to the new work-unit format. The loop now supports per-unit `Agent:` overrides, `LOOP_AGENT_CMD` for agent-agnosticism, and writes `.loop/HANDOFF.md` on non-clean exits. The design in `DESIGN.md` covers the full explore → plan → build → review → fix flow. The remaining skills (explore, build, review, fix, decide, domain-modeling) and the Go CLI are not yet built.

## Core artifacts

- `.loop/QUEUE.md` — disposable queue of verifiable work units.
- `.loop/EVIDENCE.md` — append-only ledger of what each tick proved.
- `.loop/HANDOFF.md` — cross-session handoff (written on pause/stop).
- `decisions/` — durable ADRs (architectural rulings, not current behavior).
- `glossary.md` — durable ubiquitous language.
- `.agents/skills/` — procedural knowledge (plan built; explore, build, review, fix, decide, domain-modeling pending).
- `AGENTS.md` — operational context (this file).

## Working conventions

- Shell first for the loop. Go for the CLI. Markdown for skills.
- Plain Markdown files over stores or schemas.
- The worker never self-certifies. The runner owns verification.
- A work unit must leave the repo better if the loop stops immediately after it.
- Work units are whatever shape the work is — not forced into "vertical slices."
- `LOOP_AGENT_CMD` overrides the worker invocation (agent-agnostic: pi, claude, codex, opencode, etc.).
- Per-unit `Agent:` field in QUEUE.md overrides `LOOP_AGENT_CMD` for one unit (model routing).
- Work units are `## <outcome>` headers — no "Slice" prefix, no numbering. Vertical slice is one type of work unit, not the required format.
- Specs are disposable. Decisions are durable. Code is the source of truth.

## Verification

After meaningful changes, run:

```bash
./tests/run.sh
```
