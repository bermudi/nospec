
## 2026-07-10T17:31:46-06:00 — ADR-0008 captures the loop-orchestrated review-fix decision

Status: done

Unit:
````markdown
## ADR-0008 captures the loop-orchestrated review-fix decision

Read first:
- `.loop/review-loop/specs/proposal.md`
- `DESIGN.md` (open questions and "What the loop does NOT do")
- `.agents/skills/review/SKILL.md`
- `.agents/skills/fix/SKILL.md`

Constraints:
- The ADR must state the loop's role (orchestrator) vs the review/fix skills' role (judgment).
- Do not change `loop.sh`, `review/SKILL.md`, or `fix/SKILL.md` yet.
- Keep the ADR to one page.

Done means:
- `decisions/0008-loop-orchestrates-review-fix.md` exists with `Status: accepted`.
- `knack decisions check` still passes (the new ADR is not orphaned — it is referenced by this queue).

Verify:
```bash
cd /home/daniel/build/knack && test -f decisions/0008-loop-orchestrates-review-fix.md && grep -q '^Status: accepted' decisions/0008-loop-orchestrates-review-fix.md && /tmp/knack decisions check
```

Status: pending````

Files changed:
```text
 M .loop/review-loop/QUEUE.md
?? .loop/review-loop/EVIDENCE.md
?? decisions/0008-loop-orchestrates-review-fix.md
```

Verify command:
```bash
cd /home/daniel/build/knack && test -f decisions/0008-loop-orchestrates-review-fix.md && grep -q '^Status: accepted' decisions/0008-loop-orchestrates-review-fix.md && /tmp/knack decisions check
```

Verify output:
```text
```

Worker output:
````text
Done. The `QUEUE.md` modification (1 line) is the runner's `pending → in_progress` status flip from when it claimed this unit — not my edit. My only change is the new ADR file.

The unit's `Verify` command passes: file exists, `Status: accepted` present, `decisions check` green (the ADR is referenced by five `decisions/0008-...` path citations already in the QUEUE.md unit bodies, so it is not orphaned; no `ADR-0008` token sits in a unit body to dangle).

```text
Unit: ADR-0008 captures the loop-orchestrated review-fix decision
Changed: decisions/0008-loop-orchestrates-review-fix.md (new, 41 lines — one page)
Verify expected: cd /home/daniel/build/knack && test -f decisions/0008-loop-orchestrates-review-fix.md && grep -q '^Status: accepted' decisions/0008-loop-orchestrates-review-fix.md && /tmp/knack decisions check
Notes: ADR follows the decide-skill format (Context/Decision/Consequences), modeled on ADR-0007. The Decision section states the boundary as a ruling — loop owns orchestration + stop conditions (mechanical: invoke review/fix, read REVIEW.md actionable count, enforce hard stops); review/fix skills own judgment (semantic: two-axis review, triage, finding phrasing). Did not touch loop.sh, review/SKILL.md, or fix/SKILL.md per constraints. Confirmed non-orphaned by inspecting adrReferenced() candidates (matches the decisions/0008-...md path already cited in QUEUE.md units 1–5).
```
````

What this proves:
- The verify command passed for this work unit in the current repo state.

What remains unverified:
- Anything outside the verify command's proof scope.
