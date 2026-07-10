# FAQ

## General

### What is knack?

knack is a small harness for agentic development. It turns intent into a disposable queue of verifiable work units, runs one unit at a time behind a deterministic verification gate, and ships with a read-only CLI and a set of skills.

### How does knack differ from litespec?

knack replaces litespec. It keeps the flow (explore → plan → build → review → fix) and the idea of skills, but drops durable specs as source of truth. Code is the source of truth; work units and handoff files are disposable; decisions and glossary are durable.

### Do I have to use the CLI?

No. `loop.sh` works on its own. The CLI is a convenience for validating queues, scaffolding skills, checking ADRs, and getting templates.

## The loop

### What is the worker? What is the runner?

The **runner** is `loop.sh`. It reads `QUEUE.md`, invokes the **worker** (any coding agent), and runs the `Verify:` command after the worker exits. The worker does not self-certify — the runner owns the gate.

### What can I use for the worker?

Any agent with a CLI wrapper. Set `LOOP_AGENT_CMD` to the command, or use the `Agent:` field per work unit. Examples include `pi`, `claude`, `codex`, `opencode`, and `devin`.

### Why does `LOOP_AGENT_CMD` need `$(cat "$LOOP_PROMPT_FILE")`?

When `LOOP_AGENT_CMD` is set, `loop.sh` writes the worker prompt into a temporary file and sets `LOOP_PROMPT_FILE` to that file. Your command must read the prompt from that file and pass it to the agent. The default `loop.sh` invocation (when `LOOP_AGENT_CMD` is unset) does this for `pi` automatically.

### What does `Verify:` have to be?

A deterministic, executable command — tests, builds, type checks, a script. Not an LLM-as-judge. The runner executes it and treats exit code `0` as success.

### What happens if the loop stops before finishing?

`loop.sh` writes `HANDOFF.md` on any non-clean exit. It lists completed, in-progress, and remaining units plus a next action. Resume by re-running the loop.

### What is `EVIDENCE.md`?

An append-only ledger next to `QUEUE.md`. It records the unit, changed files, verify command, verify output, and worker output. It is durable — keep it after deleting `QUEUE.md` so `knack decisions check` can still trace which ADRs the cycle referenced.

### Why must `Status:` start as `pending`?

The loop updates it. The runner marks it `in_progress`, `done`, `verify_failed`, `no_progress`, or `blocked`. Do not edit it by hand.

## Queue format

### Why no `###` subheadings inside a work unit?

`loop.sh` parses work units by looking for `## ` lines at the top level. `###` subheadings can confuse simple parsers and are best avoided inside a work unit. Put extra detail in the fields or in a disposable spec under `.loop/<name>/specs/`.

### What is the difference between `Done means:` and `Verify:`?

`Done means:` is the acceptance criteria — what must be true when the unit is finished. `Verify:` is the mechanically enforceable subset that the runner can actually execute. The gap between them is the review surface.

### Can I use a work unit that is not a vertical slice?

Yes. "Vertical slice" is the preferred default, but work units can be patches, bug fixes, investigations, or refactors. The only hard requirement is a deterministic `Verify:` command.

## Skills

### What are `.agents/skills/`?

Markdown files that encode procedural knowledge. Each file is a skill. The worker prompt names the skill by name and path; the agent reads the skill file directly. agentskills.io discovery is not relied upon (ADR-0007).

### Can I change the default skills after scaffolding?

Yes. `knack skills init` writes them into your project. After that, the project owns the `.agents/skills/` directory. Edit, override, or delete as needed.

### How do I keep the CLI's embedded skills in sync with the repo defaults?

If you edit `.agents/skills/` in the `knack` repo itself, run `cli/sync-skills.sh` and then `diff -r .agents/skills cli/embedded/skills` to verify.

## CLI

### What is `knack validate` for?

It checks that a `QUEUE.md` has valid work units (each has a `##` outcome, a `Verify:` block, and a `Status:`). It does not run the loop or the verify command.

### What does `knack decisions check` do?

It flags two problems:
- **Orphaned ADRs** — an ADR in `decisions/` is not referenced by any current `QUEUE.md` or completed `EVIDENCE.md`.
- **Dangling references** — a work unit references an ADR that does not exist.

### What is `knack status`?

It shows aggregate counts across all `.loop/<name>/` cycles: pending, done, failed, evidence entries, and total ADRs.

### Why is `knack` called read-only?

`knack` validates, lists, and provides context. It does not run agents, edit the queue, or execute the loop. The only write operation is `knack skills init`, which scaffolds skills into a new project.

## Workflows

### Can I skip explore or review?

Yes. The flow is composable. A small bug can be `plan → build → done`. Review is opt-in.

### Can I run multiple work cycles at once?

Yes. Each cycle lives in its own `.loop/<name>/` directory. Run `loop.sh` independently for each one. The cycle name is the directory name.

### Where do I put architectural decisions?

In `decisions/` as ADRs. Use the `decide` skill to capture them inline while exploring, planning, or building. Decisions are durable; the queue is disposable.

### Where do I put domain terminology?

In `glossary.md` at the repo root. Use the `domain-modeling` skill to define and update it.
