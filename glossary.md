---
role: record
owns: domain-terms
---

# Glossary

Knack's domain-specific terms. For the underlying concepts — doc-rot, backpressure, tracer bullets, and the rest — see the [AgenticWiki](https://github.com/bermudi/AgenticWiki); this file defines only what's specific to this project. Wiki concepts are linked, not redefined (ADR-0010).

## work unit

One chunk of work in a `QUEUE.md`, written as a `## <outcome>` header with `Read first:`, `Constraints:`, `Done means:`, and `Verify:`. The atom the loop processes.

## verify gate

A work unit's deterministic `Verify:` command, executed by the loop *outside* the agent. Exits 0 or the unit fails. The mechanical backpressure — the worker never self-certifies.

## tick

One loop iteration: read a pending unit, mark it in-progress, invoke the worker, run the verify gate, record evidence.

## cycle

A named body of work under `.loop/<name>/` with its own queue, evidence, and handoff. Cycles are independent and can run concurrently.

## queue

`QUEUE.md` — the disposable, ordered list of work units for a cycle. Deleted when the work is done.

## evidence

`EVIDENCE.md` — the durable, append-only ledger of what each tick proved. Each entry carries a registry-derived proof boundary (what the verify command mechanically proves, derived from the command itself) and a pin-state record (which durable docs were touched, with alerts when a prior pin moves). Survives `QUEUE.md` deletion so a completed cycle still anchors its ADR references. Pin alerts are triage triggers for `review` → `document`, not coherence gates (ADR-0016).

## handoff

`HANDOFF.md` — a disposable cross-session snapshot (completed / in-progress / remaining) the loop writes on pause or non-clean exit.

## ADR

A durable architectural ruling in `decisions/`. Records *why* the code is the way it is; survives spec deletion. Active unless marked superseded.

## loop

The optional AFK runner (`loop.sh`): runs a cycle's queue one tick at a time behind the verify gate, agent-agnostic via `LOOP_AGENT_CMD`. Skills serve interactive and plan-then-leave modes; the loop serves batch.

## Concepts (external)

The theory behind these terms lives in the [AgenticWiki](https://github.com/bermudi/AgenticWiki). Load-bearing concepts:

- [tracer bullets](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/tracer-bullets.md) — thin end-to-end slices for early feedback
- [backpressure](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/backpressure.md) — engineer the environment to mechanically reject wrong outputs
- [agent-loop](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/agent-loop.md) — cron plus a decision-maker; non-negotiable hard stops
- [ralph-loop](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/ralph-loop.md) — fresh context per tick, plan file as shared state
- [compounding-loops](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/compounding-loops.md) — coordination through shared durable files
- [doc-rot](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/doc-rot.md) — stale docs are worse than none
- [plan-disposability](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/plan-disposability.md) — plans are ephemeral coordination state, not contracts
- [code-clarifies-spec](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/code-clarifies-spec.md) — implementing surfaces decisions the spec missed
- [decision-extraction](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/decision-extraction.md) — keep the decisions, not the spec
- [ubiquitous-language](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/ubiquitous-language.md) — the shared vocabulary that lets human and agent mean the same thing by the same word
- [evolving-context](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/evolving-context.md) — agents refine their own context over time
- [aiming-problem](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/aiming-problem.md) — verify the real property, not a gameable proxy
- [overcorrection-bias](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/overcorrection-bias.md) — reviewers misclassify correct code as defective when pushed to explain and fix
- [iterative-self-correction](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/iterative-self-correction.md) — feedback loops that can amplify error rather than converge
- [verification-loop](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/verification-loop.md) — close the feedback loop early and often; the rate of feedback is the speed limit
- [infrastructure-blindness](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/infrastructure-blindness.md) — finding the right code but reimplementing its machinery instead of calling it
- [over-engineering](https://github.com/bermudi/AgenticWiki/blob/main/wiki/concepts/over-engineering.md) — patches that add speculative abstraction beyond the requirement; the extra surface is where the bugs live
