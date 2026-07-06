
## 2026-07-06T03:36:04-06:00 — Go scaffold and `skills init` scaffolds the seven default skills into a target project

Status: done

Files changed:
```text
 M .loop/QUEUE.md
?? .loop/EVIDENCE.md
?? cli/embedded/skills/build/SKILL.md
?? cli/embedded/skills/decide/SKILL.md
?? cli/embedded/skills/domain-modeling/SKILL.md
?? cli/embedded/skills/explore/SKILL.md
?? cli/embedded/skills/fix/SKILL.md
?? cli/embedded/skills/plan/SKILL.md
?? cli/embedded/skills/review/SKILL.md
?? cli/go.mod
?? cli/internal/skills/skills.go
?? cli/internal/skills/skills_test.go
?? cli/main.go
?? cli/main_test.go
?? cli/sync-skills.sh
```

Verify command:
```bash
cd cli && go test ./...
```

Verify output:
```text
ok  	knack	(cached)
ok  	knack/internal/skills	(cached)
```

Worker output:
````text
```text
Unit: Go scaffold and `skills init` scaffolds the seven default skills into a target project
Changed:
  cli/go.mod
  cli/main.go
  cli/internal/skills/skills.go
  cli/internal/skills/skills_test.go
  cli/main_test.go
  cli/sync-skills.sh
  cli/embedded/skills/{build,decide,domain-modeling,explore,fix,plan,review}/SKILL.md
Verify expected: cd cli && go test ./...
Notes:
  - `go test ./...` passes.
  - `go build` produces a binary that scaffolds all 7 skills into a fresh temp dir.
  - Re-running `knack skills init` on the same dir skips existing skills.
  - `cli/sync-skills.sh` keeps `cli/embedded/skills/` in sync with `../.agents/skills/` (verified with `diff -r`).
  - `./tests/run.sh` still passes.
  - No external Go dependencies.
```
````

What this proves:
- The verify command passed for this work unit in the current repo state.

What remains unverified:
- Anything outside the verify command's proof scope.

## 2026-07-06T04:39:59-06:00 — `skills check` validates that skills follow the agentskills.io spec

Status: done

Files changed:
```text
 M .loop/QUEUE.md
?? .loop/EVIDENCE.md
?? cli/embedded/skills/build/SKILL.md
?? cli/embedded/skills/decide/SKILL.md
?? cli/embedded/skills/domain-modeling/SKILL.md
?? cli/embedded/skills/explore/SKILL.md
?? cli/embedded/skills/fix/SKILL.md
?? cli/embedded/skills/plan/SKILL.md
?? cli/embedded/skills/review/SKILL.md
?? cli/go.mod
?? cli/go.sum
?? cli/internal/skills/skills.go
?? cli/internal/skills/skills_test.go
?? cli/main.go
?? cli/main_test.go
?? cli/sync-skills.sh
```

Verify command:
```bash
cd cli && go test ./...
```

Verify output:
```text
ok  	knack	(cached)
ok  	knack/internal/skills	(cached)
```

Worker output:
````text
Unit: `skills check` validates that skills follow the agentskills.io spec
Changed: `cli/internal/skills/skills.go` (added `Check` + frontmatter/reference helpers), `cli/internal/skills/skills_test.go` (added Check fixtures/project/embedded tests), `cli/main.go` (wired `skills check [--dir DIR]`), `cli/go.mod`/`go.sum` (added `gopkg.in/yaml.v3`)
Verify expected: `cd cli && go test ./...`
Notes: None.
````

What this proves:
- The verify command passed for this work unit in the current repo state.

What remains unverified:
- Anything outside the verify command's proof scope.

## 2026-07-06T04:48:11-06:00 — `validate` checks that QUEUE.md work units are well-formed

Status: done

Files changed:
```text
 M .loop/QUEUE.md
 M AGENTS.md
?? .loop/EVIDENCE.md
?? cli/embedded/skills/build/SKILL.md
?? cli/embedded/skills/decide/SKILL.md
?? cli/embedded/skills/domain-modeling/SKILL.md
?? cli/embedded/skills/explore/SKILL.md
?? cli/embedded/skills/fix/SKILL.md
?? cli/embedded/skills/plan/SKILL.md
?? cli/embedded/skills/review/SKILL.md
?? cli/go.mod
?? cli/go.sum
?? cli/internal/queue/queue.go
?? cli/internal/queue/queue_test.go
?? cli/internal/skills/skills.go
?? cli/internal/skills/skills_test.go
?? cli/main.go
?? cli/main_test.go
?? cli/sync-skills.sh
```

Verify command:
```bash
cd cli && go test ./...
```

Verify output:
```text
ok  	knack	(cached)
ok  	knack/internal/queue	(cached)
ok  	knack/internal/skills	(cached)
```

Worker output:
````text
Unit: `validate` checks that QUEUE.md work units are well-formed
Changed:
- `cli/internal/queue/queue.go` — new parser/validator for QUEUE.md work units
- `cli/internal/queue/queue_test.go` — tests for valid fixture, missing Verify, empty outcome, missing fence, smoke fixture, multi-failure, formatting, and fs reading
- `cli/main.go` — wired `knack validate <queue-file>` and updated usage
- `AGENTS.md` — documented `cd cli && go test ./...` as the CLI verification command
Verify expected: `cd cli && go test ./...`
Notes: None. All tests pass; also verified `./tests/run.sh` and a manual CLI run against `examples/smoke/.loop/QUEUE.md` (exit 0).
````

What this proves:
- The verify command passed for this work unit in the current repo state.

What remains unverified:
- Anything outside the verify command's proof scope.

## 2026-07-06T04:53:11-06:00 — `decisions list|show|check` and `status` provide ADR management and the decision coverage gate

Status: done

Files changed:
```text
 M .loop/QUEUE.md
 M AGENTS.md
?? .loop/EVIDENCE.md
?? cli/embedded/skills/build/SKILL.md
?? cli/embedded/skills/decide/SKILL.md
?? cli/embedded/skills/domain-modeling/SKILL.md
?? cli/embedded/skills/explore/SKILL.md
?? cli/embedded/skills/fix/SKILL.md
?? cli/embedded/skills/plan/SKILL.md
?? cli/embedded/skills/review/SKILL.md
?? cli/go.mod
?? cli/go.sum
?? cli/internal/decisions/decisions.go
?? cli/internal/decisions/decisions_test.go
?? cli/internal/queue/queue.go
?? cli/internal/queue/queue_test.go
?? cli/internal/skills/skills.go
?? cli/internal/skills/skills_test.go
?? cli/internal/status/status.go
?? cli/internal/status/status_test.go
?? cli/main.go
?? cli/main_test.go
?? cli/sync-skills.sh
```

Verify command:
```bash
cd cli && go test ./...
```

Verify output:
```text
ok  	knack	(cached)
ok  	knack/internal/decisions	(cached)
ok  	knack/internal/queue	(cached)
ok  	knack/internal/skills	(cached)
ok  	knack/internal/status	(cached)
```

Worker output:
````text
Unit: `decisions list|show|check` and `status` provide ADR management and the decision coverage gate
Changed:
- `cli/internal/queue/queue.go` — exported `ParseUnits` and added `Unit.Status()`
- `cli/internal/decisions/decisions.go` — new ADR list/show/check implementation
- `cli/internal/decisions/decisions_test.go` — fixture tests for list, show, clean coverage, orphaned + dangling, filename reference
- `cli/internal/status/status.go` — new `status` report generator
- `cli/internal/status/status_test.go` — fixture tests for counts and missing-loop state
- `cli/main.go` — wired `decisions` and `status` commands
- `.loop/QUEUE.md` — added ADR-0003 reference in the current unit so the repo’s own `decisions check` passes
Verify expected: `cd cli && go test ./...`
Notes: All tests pass; `knack decisions check` exits 0 on this repo; `knack status` and `knack decisions list|show` verified manually.
````

What this proves:
- The verify command passed for this work unit in the current repo state.

What remains unverified:
- Anything outside the verify command's proof scope.

## 2026-07-06T04:58:51-06:00 — `glossary check` and `instructions` round out the CLI command set

Status: done

Files changed:
```text
 M .loop/QUEUE.md
 M AGENTS.md
?? .loop/EVIDENCE.md
?? cli/embedded/skills/build/SKILL.md
?? cli/embedded/skills/decide/SKILL.md
?? cli/embedded/skills/domain-modeling/SKILL.md
?? cli/embedded/skills/explore/SKILL.md
?? cli/embedded/skills/fix/SKILL.md
?? cli/embedded/skills/plan/SKILL.md
?? cli/embedded/skills/review/SKILL.md
?? cli/go.mod
?? cli/go.sum
?? cli/internal/decisions/decisions.go
?? cli/internal/decisions/decisions_test.go
?? cli/internal/glossary/glossary.go
?? cli/internal/glossary/glossary_test.go
?? cli/internal/instructions/instructions.go
?? cli/internal/instructions/instructions_test.go
?? cli/internal/queue/queue.go
?? cli/internal/queue/queue_test.go
?? cli/internal/skills/skills.go
?? cli/internal/skills/skills_test.go
?? cli/internal/status/status.go
?? cli/internal/status/status_test.go
?? cli/main.go
?? cli/main_test.go
?? cli/sync-skills.sh
```

Verify command:
```bash
cd cli && go test ./...
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
Unit: `glossary check` and `instructions` round out the CLI command set
Changed: Added `cli/internal/glossary/`, `cli/internal/instructions/`, and wired both into `cli/main.go`
Verify expected: `cd cli && go test ./...`
Notes: None — all tests pass, including the new glossary and instructions tests. The full `./tests/run.sh` suite also passes.
````

What this proves:
- The verify command passed for this work unit in the current repo state.

What remains unverified:
- Anything outside the verify command's proof scope.
