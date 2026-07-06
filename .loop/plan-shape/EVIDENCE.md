
## 2026-07-06T12:37:56-06:00 — Skills and docs teach the shape-not-scripts work-unit format

Status: done

Files changed:
```text
 M .agents/skills/build/SKILL.md
 M .agents/skills/fix/SKILL.md
 M .agents/skills/plan/SKILL.md
 M .agents/skills/review/SKILL.md
 M AGENTS.md
 M DESIGN.md
 M cli/embedded/skills/build/SKILL.md
 M cli/embedded/skills/fix/SKILL.md
 M cli/embedded/skills/plan/SKILL.md
 M cli/embedded/skills/review/SKILL.md
 M prompts/worker.md
?? .loop/plan-shape/EVIDENCE.md
?? .loop/plan-shape/QUEUE.md
?? decisions/0005-work-units-carry-shape-not-scripts.md
```

Verify command:
```bash
./tests/run.sh && diff -r .agents/skills cli/embedded/skills && ! grep -rn 'Work:' .agents/skills DESIGN.md prompts/worker.md AGENTS.md && grep -q 'Read first:' .agents/skills/plan/SKILL.md && grep -q 'Constraints:' .agents/skills/plan/SKILL.md
```

Verify output:
```text
Valid skill: /home/daniel/build/knack/.agents/skills/build
Valid skill: /home/daniel/build/knack/.agents/skills/decide
Valid skill: /home/daniel/build/knack/.agents/skills/domain-modeling
Valid skill: /home/daniel/build/knack/.agents/skills/explore
Valid skill: /home/daniel/build/knack/.agents/skills/fix
Valid skill: /home/daniel/build/knack/.agents/skills/plan
Valid skill: /home/daniel/build/knack/.agents/skills/review
knack tests passed
```

Worker output:
````text
Unit: Skills and docs teach the shape-not-scripts work-unit format
Changed: `.agents/skills/{plan,build,review,fix}/SKILL.md`, `prompts/worker.md`, `DESIGN.md`, `AGENTS.md`, `cli/embedded/skills` (via `cli/sync-skills.sh`)
Verify expected: `./tests/run.sh && diff -r .agents/skills cli/embedded/skills && ! grep -rn 'Work:' .agents/skills DESIGN.md prompts/worker.md AGENTS.md && grep -q 'Read first:' .agents/skills/plan/SKILL.md && grep -q 'Constraints:' .agents/skills/plan/SKILL.md`
Notes: Verify passed. Also ran `cd cli && go test ./...` — all green. Left `cli/internal/instructions/instructions.go` untouched; it is the next pending unit in this cycle.
````

What this proves:
- The verify command passed for this work unit in the current repo state.

What remains unverified:
- Anything outside the verify command's proof scope.

## 2026-07-06T12:38:32-06:00 — `knack instructions` emits the new work-unit shape

Status: done

Files changed:
```text
 M .agents/skills/build/SKILL.md
 M .agents/skills/fix/SKILL.md
 M .agents/skills/plan/SKILL.md
 M .agents/skills/review/SKILL.md
 M AGENTS.md
 M DESIGN.md
 M cli/embedded/skills/build/SKILL.md
 M cli/embedded/skills/fix/SKILL.md
 M cli/embedded/skills/plan/SKILL.md
 M cli/embedded/skills/review/SKILL.md
 M cli/internal/instructions/instructions.go
 M cli/internal/instructions/instructions_test.go
 M prompts/worker.md
?? .loop/plan-shape/EVIDENCE.md
?? .loop/plan-shape/QUEUE.md
?? decisions/0005-work-units-carry-shape-not-scripts.md
```

Verify command:
```bash
(cd cli && go test ./...) && ! grep -n 'Work:' cli/internal/instructions/instructions.go && grep -q 'Read first:' cli/internal/instructions/instructions.go
```

Verify output:
```text
ok  	knack	(cached)
ok  	knack/internal/decisions	(cached)
ok  	knack/internal/glossary	(cached)
ok  	knack/internal/instructions	(cached)
ok  	knack/internal/queue	(cached)
ok  	knack/internal/skills	(cached)
ok  	knack/internal/status	(cached)
```

Worker output:
````text
Unit: `knack instructions` emits the new work-unit shape
Changed: `cli/internal/instructions/instructions.go`, `cli/internal/instructions/instructions_test.go`
Verify expected: `(cd cli && go test ./...) && ! grep -n 'Work:' cli/internal/instructions/instructions.go && grep -q 'Read first:' cli/internal/instructions/instructions.go`
Notes: None. Tests pass and the verify command exits 0.
````

What this proves:
- The verify command passed for this work unit in the current repo state.

What remains unverified:
- Anything outside the verify command's proof scope.
