---
name: fix
description: Use when addressing review findings and feeding them back into the loop. Converts review findings into new verifiable work units in `.loop/QUEUE.md` and runs another loop pass. Triggers on "fix the review findings", "address the feedback", "rework based on review", "fix what review found", or when review produced findings that need to be resolved.
---

# Fix

Address review findings by triaging them and generating new work units that feed back into the loop. The `fix` skill is the bridge between review and another build pass; it owns the triage and the queue format.

## Procedure

1. **Read the review findings.** Each finding from the `review` skill is a candidate work unit. `review` has already classified findings as trivial / actionable / disputed / deferred; use those classifications or re-triage as needed.

2. **Triage.** Not every finding warrants a work unit:
   - **Actionable** — the finding identifies a real issue with a clear fix. → Create a work unit.
   - **Trivial** — one-line fix, no risk. → Fix it now, don't create a unit.
   - **Disputed** — the finding is wrong or the reviewer is being overly cautious. → Note the disagreement, don't create a unit. Move on.
   - **Deferred** — the finding is valid but not worth fixing now. → Note it in `AGENTS.md` or a backlog, don't create a unit.

3. **Create work units** for actionable findings. Use the standard work unit format from the `plan` skill (single-unit excerpt below):

````markdown
## <fix for the finding — observable outcome>

Why:
<reference to the review finding — which axis, what was found>

Work:
- <what to fix>
- <guardrail: don't break what the review approved>

Verify:
```bash
<deterministic command that proves the fix>
```

Done means:
- <the finding is resolved>
- <no new issue introduced>

Status: pending
````

4. **Append to QUEUE.md.** Read the existing `.loop/QUEUE.md` first, then append the new units. If the queue was completed (all units `done`), the new units extend it. If the queue still has pending units, the new units are added after them. Preserve the existing structure and status.

5. **Run the loop.** After writing the units, run:

```bash
./loop.sh run .loop/QUEUE.md
```

The loop will pick up the first pending unit and proceed.

## What fix is not

- Not a re-review — the findings are already known. Fix turns them into work.
- Not a debate — if a finding is disputed, note it and move on. Don't argue with the review in the queue.
- Not a rewrite — each fix unit should be narrow. If a finding requires a broad rewrite, that's a new explore → plan cycle, not a fix.

## Guardrail

Fix units must not break what review already approved. Each fix unit's `Done means:` should include a no-regression condition. The verify command should cover both the fix and the preservation of what was already working.

## Output

- New work units appended to `.loop/QUEUE.md`
- A summary of what was triaged (actionable / trivial / disputed / deferred)
- Suggestion to run the loop
