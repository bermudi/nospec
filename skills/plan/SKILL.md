---
name: plan
description: Use when decomposing a software task, bug, cleanup, or vague human intent into bounded, observable, independently-verifiable outcomes. When execution will use loop.sh, serialize them as a `.loop/<name>/QUEUE.md`; the planner picks a short name for the cycle (e.g., `parser-bug`).
---

# Plan

Decompose work into bounded, observable outcomes — each one independently verifiable. That core is the same whether you're planning interactively with the human or producing a packet for unattended batch execution. Only the *serialization* differs: batch work is written as a `.loop/<name>/QUEUE.md` the loop runner consumes; interactive work need not be serialized at all.

A plan is **ephemeral coordination state, not a contract** — when it drifts, throw it away and regenerate from the codebase; salvaging a stale plan costs more than rewriting it (plan-disposability). The goal is verifiable outcomes, not durable specs. When the work is done, the plan is deleted.

For a batch cycle, pick a short, descriptive name (`parser-bug`, `go-cli`) — it scopes the `.loop/<name>/` directory.

## Decomposition concepts

These are leading words — dense phrases that shape how you reason about decomposition. Use the one that fits the work; don't force one shape onto everything.

- **Tracer bullet** — a thin slice that crosses *all* layers end-to-end, to get real integration feedback early. Its purpose is to discover mismatches (schema vs. UI, contract vs. consumer) in hour one rather than week three. It earns its keep when there *is* an end-to-end path to prove; a mechanical migration or a one-line fix has nothing to trace.
- **Vertical slice** — crosses enough layers to produce a user- or system-visible change. The same idea as a tracer bullet, emphasized for the depth of feedback.
- **Horizontal breadth** — one change applied across a layer or family of cases. Efficient when the contract is already understood; its risk is delaying integration feedback when the contract is still uncertain. It can legitimately come first when a shared invariant or mechanical migration *is* the outcome.

The tradeoff each carries is the reason to know them — not a rule for when to deploy them. Read the work, choose the decomposition that fits.

## The work unit

A work unit is one observable outcome that can be verified. It carries:

- **`Read first:`** — context the worker needs (ADRs, code areas, rulings). Context, *not* scope. Naming files-to-edit here smuggles in a script and the worker becomes a script-executor instead of an agent (ADR-0005). Point at the relevant areas; let the worker find the edits.
- **`Constraints:`** — what must stay true or what is out of bounds. If it names a file, it's "don't touch X" or "X's public API must not change," never "update X." Constraints close the solution space; they don't prescribe the path through it.
- **`Done means:`** — the acceptance criteria: what's observably true afterward.
- **`Verify:`** — the mechanically enforceable subset of `Done means:`.

The gap between `Done means:` and `Verify:` is the **review surface** — what the command can't check, review does. A unit whose outcome can't be captured as a deterministic verify plus a short acceptance list isn't ready; vague shape plus weak verify gives the worker nothing to aim at (aiming-problem).

## Verification follows the outcome

Shape the verify from the nature of the change, not from a template:

- A **bug fix** is verified by a regression that *fails before* the fix and *passes after* — a test that already passes proves nothing about the bug.
- A **refactor** is verified by existing behavior staying unchanged (existing tests still pass); if behavior could change, it isn't a refactor.
- An **investigation** produces decision-relevant evidence — a finding that changes what to build next — not merely a file. For batch execution, mechanically verify the *observable* evidence that can be checked (the benchmark ran, the data was collected); the evidence's *relevance* is judgment, so it lives in `Done means:` and the review surface, not in the verify command. An investigation with no observable evidence to check isn't batch-runnable — run it interactively.

These aren't categories to tag units with; they're reminders that the verify must actually probe the claimed outcome.

## Hard rules (mechanical contracts)

Non-negotiable because they're mechanisms, not judgment:

- **Every outcome needs credible verification** — name how it will be checked: a test, a type-check, running the thing, a manual check. This is universal.
- **For batch execution specifically, `Verify:` must be deterministic and runnable by the runner, not an LLM-as-judge.** That is the backpressure — the loop mechanically rejecting wrong output, outside the agent. Tests, type checks, builds; not "ask the model if it looks right." Interactive outcomes have no runner; they name their check but need not use a `Verify:` field or reduce it to a shell command.
- **Batch serialization format** (only when execution uses loop.sh): `.loop/<name>/QUEUE.md`, one unit per `## <outcome>` header (no "Slice" prefix, no numbering), each with a `Status:` line starting at `pending`. The loop's parser requires this; interactive planning has no such constraint.

## Reasoned defaults

Defaults for the common case. Override when the reasoning says to — not because a rule told you to:

- **Keep a batch queue short (2–5 units).** The loop has hard stops (max ticks, no-progress detection); a longer queue means it pauses mid-way and writes a handoff. If the work genuinely needs more, let it — the next session resumes from the handoff.
- **Reach for a tracer bullet when there's an end-to-end path to prove.** When there isn't, don't — there's nothing to trace.
- **If no deterministic verify exists yet, make the first outcome create one.** An outcome without a verify can't be checked.

## When to regenerate

If an existing `.loop/<name>/QUEUE.md` is stale, contradicts the current code, or no longer matches the real goal, discard it and write fresh. Regenerating from the actual codebase is cheaper than salvaging a drifting plan.

## The design note

For cycles where the worker can't recover the reasoning from the codebase alone — an external spec constrains the shape, a trade-off was resolved, a decision spans units — write a one-page `.loop/<name>/DESIGN.md`: the *why* behind the units. **The trigger is not size; it is reasoning the worker cannot recover from the codebase.** That guard prevents both ceremony on every large task and missing context for a small but externally constrained one. The note serves the worker, the reviewer, and the fixer. It is disposable.

## Capture decisions and terms inline

If a ruling crystallizes while you plan, write the ADR now via the `decide` skill — don't queue it. If a domain term is ambiguous or inconsistent, define it now via `domain-modeling`. Decisions and terms are durable; the plan is not.

## Batch queue format

When execution will use loop.sh, serialize the outcomes as:

````markdown
# Loop Queue: <short name>

Goal:
<one paragraph — the desired end state>

Stop condition:
`<command that proves the whole packet is done, if one exists>`

## <outcome — what changes, observable>

Agent: <optional — overrides LOOP_AGENT_CMD for this unit>

Why: <optional — only if non-obvious>

Read first:
- .loop/<name>/DESIGN.md (if it exists — cycle-level reasoning)
- <ADRs, code areas, or rulings — context, not scope>

Constraints:
- <what must stay true, or what is out of bounds>

Done means:
- <observable condition>

Verify:
```bash
<command that exits 0 on success>
```

Status: pending
````

## Disposability

`.loop/<name>/` is scratchpad, not artifact. When the work is done and verified, delete the cycle's subdirectory. What persists: code, tests, `decisions/`, `glossary.md`, `AGENTS.md`.
