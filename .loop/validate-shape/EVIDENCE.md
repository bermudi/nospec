
## 2026-07-06T15:08:56-06:00 — the ADR template emitted by knack instructions matches the project's actual ADR convention

Status: done

Files changed:
```text
 M .agents/skills/build/SKILL.md
 M .agents/skills/fix/SKILL.md
 M .agents/skills/plan/SKILL.md
 M .agents/skills/review/SKILL.md
 M AGENTS.md
 M DESIGN.md
 M README.md
 M cli/embedded/skills/build/SKILL.md
 M cli/embedded/skills/fix/SKILL.md
 M cli/embedded/skills/plan/SKILL.md
 M cli/embedded/skills/review/SKILL.md
 M cli/internal/instructions/instructions.go
 M cli/internal/instructions/instructions_test.go
 M cli/internal/queue/queue_test.go
 M examples/smoke/.loop/QUEUE.md
 M examples/wiki-validator/.loop/QUEUE.md
 M prompts/worker.md
 M tests/run.sh
?? .loop/plan-shape/EVIDENCE.md
?? .loop/plan-shape/QUEUE.md
?? .loop/validate-shape/EVIDENCE.md
?? .loop/validate-shape/QUEUE.md
?? decisions/0005-work-units-carry-shape-not-scripts.md
```

Verify command:
```bash
(cd cli && go test ./...) && (cd cli && go run . instructions adr) | grep -q 'Status: accepted'
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
Unit: the ADR template emitted by knack instructions matches the project's actual ADR convention
Changed: cli/internal/instructions/instructions.go, cli/internal/instructions/instructions_test.go
Verify expected: (cd cli && go test ./...) && (cd cli && go run . instructions adr) | grep -q 'Status: accepted'
Notes: none
````

What this proves:
- The verify command passed for this work unit in the current repo state.

What remains unverified:
- Anything outside the verify command's proof scope.
