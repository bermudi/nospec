# 0003: The tool is named knack

Date: 2026-07-06
Status: accepted

## Context

Open question #1 in DESIGN.md: "sliceloop" was a placeholder. The tool is no longer just a slice loop — it's a full workflow harness (explore, plan, build, review, fix) with a loop engine, a skill set, and a CLI. The name needed to reflect that and not bake in the "slice" framing, which we already dropped when work units stopped being forced into vertical slices.

## Decision

The tool is named **knack**. The CLI binary will be `knack`, the loop's log prefix will be `knack:`, and all documentation references to "sliceloop" become "knack".

## Consequences

- Resolves DESIGN.md open question #1.
- The binary name and Go module path will use `knack` — this must be done before the CLI build starts.
- `SLICELOOP_AGENT_CMD` was already renamed to `LOOP_AGENT_CMD`; no env var rename needed.
- The repo directory name (`sliceloop`) can stay as-is for now — renaming a working directory is disruptive and not load-bearing. The tool's name is what's in the code and docs, not the folder it lives in.
