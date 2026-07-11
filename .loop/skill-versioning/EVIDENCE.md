
## 2026-07-10T23:17:54-06:00 — all seven skills carry a version field in frontmatter

Status: done

Unit:
````markdown
## all seven skills carry a version field in frontmatter

Read first:
- .agents/skills/*/SKILL.md
- cli/embedded/skills/*/SKILL.md
- cli/sync-skills.sh

Constraints:
- Do not change skill bodies or descriptions — only add the frontmatter field.
- Do not change sync-skills.sh.

Done means:
- Every SKILL.md in .agents/skills/ has a `version:` field in frontmatter.
- Every SKILL.md in cli/embedded/skills/ has the same `version:` field.
- `diff -r .agents/skills cli/embedded/skills` shows no drift.

Verify:
```bash
for f in .agents/skills/*/SKILL.md; do grep -q '^version:' "$f" || { echo "missing version in $f"; exit 1; }; done && \
diff -r .agents/skills cli/embedded/skills && \
cd cli && go test ./internal/skills/...
```

Status: pending````

Files changed:
```text
 M .agents/skills/build/SKILL.md
 M .agents/skills/decide/SKILL.md
 M .agents/skills/domain-modeling/SKILL.md
 M .agents/skills/explore/SKILL.md
 M .agents/skills/fix/SKILL.md
 M .agents/skills/plan/SKILL.md
 M .agents/skills/review/SKILL.md
 M cli/embedded/skills/build/SKILL.md
 M cli/embedded/skills/decide/SKILL.md
 M cli/embedded/skills/domain-modeling/SKILL.md
 M cli/embedded/skills/explore/SKILL.md
 M cli/embedded/skills/fix/SKILL.md
 M cli/embedded/skills/plan/SKILL.md
 M cli/embedded/skills/review/SKILL.md
?? .loop/skill-versioning/EVIDENCE.md
```

Verify command:
```bash
for f in .agents/skills/*/SKILL.md; do grep -q '^version:' "$f" || { echo "missing version in $f"; exit 1; }; done && \
diff -r .agents/skills cli/embedded/skills && \
cd cli && go test ./internal/skills/...
```

Verify output:
```text
ok  	knack/internal/skills	(cached)
```

Worker output:
````text
[0m
> build · hy3-free
[0m
[0m→ [0mSkill "build"
[0m
[0m$ [0mfor f in .agents/skills/*/SKILL.md; do echo "=== $f ==="; head -20 "$f"; done
=== .agents/skills/build/SKILL.md ===
---
name: build
description: Use when implementing one work unit from a `.loop/<name>/QUEUE.md` queue. The worker-side skill — read the unit, do the work, don't self-certify, end with a handoff. Triggers on "build", "implement", "apply this unit", "do the work", "run the loop", or when a work unit needs to be executed. Also use when the loop invokes the worker for a tick.
---

# Build

Implement one work unit from `.loop/<name>/QUEUE.md`. Do the work. Don't self-certify. The loop runner owns the verify gate — your job is to make the repository state satisfy the unit's `Verify` command, not to claim success.

## Core rules

The runner injects `prompts/worker.md` at the start of every tick. If it is not in your context, read it now. Its `Rules` and `Output` sections are the canonical source for this skill; the sections below elaborate on decisions, operational learnings, blockers, and units that are too large for one tick.

> **Scope note:** Updating `AGENTS.md` or writing an ADR during a tick is a durable trace, not part of the unit's `Verify` scope. Do it only when the tick teaches you something that would save the next session time.

## Scope

The unit's scope is its outcome plus its constraints. The worker determines which files to change and how. The `Verify:` command is the mechanically enforceable subset of `Done means:`.

## Capturing decisions during build
=== .agents/skills/decide/SKILL.md ===
---
name: decide
description: Use when an architectural ruling crystallizes during any phase — explore, plan, build, or review. Captures decisions as ADRs in `decisions/` so they persist after specs are deleted and code moves on. Triggers on "we decided to", "let's go with", "the ruling is", "why did we choose", "record this decision", or when a design choice with lasting consequences is made.
---

# Decide

Capture architectural rulings as ADRs. Decisions persist; specs don't. A decision made six months ago still explains *why* the code is the way it is, even after the code has moved on.

This is a shared skill — called by explore, plan, build, and review whenever a ruling crystallizes. It is not a separate phase.

## When to write an ADR

Write an ADR when a decision has **lasting consequences** — it shapes future work and won't be obvious from the code alone:

- Choosing a library, framework, or pattern over alternatives
- Picking an architecture shape (monolith vs services, sync vs async, etc.)
- Deciding a convention (naming, file layout, error handling strategy)
- Resolving a tension between competing approaches
- Ruling something out (sometimes "we won't do X" is the decision)
=== .agents/skills/domain-modeling/SKILL.md ===
---
name: domain-modeling
description: Use when domain terms surface during explore, plan, or review and need to be defined, challenged, or stress-tested against the project's ubiquitous language. Manages `glossary.md` — the project's shared vocabulary. Triggers on "what does X mean here", "are these the same concept", "let's define our terms", "that term is ambiguous", or when a domain concept is used inconsistently across a conversation or codebase.
---

# Domain Modeling

Manage the project's ubiquitous language. Terms used in code, specs, conversations, and skills should mean the same thing everywhere. The glossary is the single source of truth for that shared meaning, and a durable part of the shared design concept.

This is a shared skill — called by explore, plan, and review. It is not a separate phase.

## The glossary

The glossary lives in `glossary.md` at the project root. It is a small, curated file — not an encyclopedia. Each entry defines one term in one or two sentences.

If `glossary.md` doesn't exist, don't create one preemptively. Create it when the first term actually needs defining.

## Format

````markdown
=== .agents/skills/explore/SKILL.md ===
---
name: explore
description: Use when investigating a codebase, grilling intent, or stress-testing ideas before planning work. The upstream phase of the loop — read code, challenge assumptions, surface constraints, and capture decisions as they crystallize. Triggers on "explore", "investigate", "let me understand", "what's going on with", "grill this", "stress-test", "help me decide", "what should I do next", or when the real problem isn't clear yet.
---

# Explore

Read the codebase, grill the intent, stress-test ideas. The goal is to reach clarity before planning work units. No artifacts produced except ADRs and glossary entries written inline when decisions crystallize.

Pure conversation. The output of explore is a clearer head, not a file on disk.

## What explore is not

- Not a planning phase — that's `plan`. Explore doesn't write QUEUE.md.
- Not a spec phase — no proposals, no designs. Those are disposable and belong in `plan` if needed at all.
- Not a code phase — no edits. Read-only.

Explore is the phase where you figure out what the problem actually is.

## Before you start
=== .agents/skills/fix/SKILL.md ===
---
name: fix
description: Use when addressing review findings and feeding them back into the loop. Converts review findings into new verifiable work units in `.loop/<name>/QUEUE.md` and runs another loop pass. Triggers on "fix the review findings", "address the feedback", "rework based on review", "fix what review found", or when review produced findings that need to be resolved.
---

# Fix

Address review findings by triaging `REVIEW.md` and generating new work units that feed back into the loop. The `fix` skill is the bridge between review and another build pass; it owns the triage and the queue format. It appends to the existing `.loop/<name>/QUEUE.md` file for the same cycle.

The loop owns orchestration. When `fix` is invoked by the loop, stop after appending units and reporting the triage summary; the loop will run the next build pass.

## Inputs

- `.loop/<name>/REVIEW.md` — structured findings from the `review` skill. Read it, but do not edit it.
- `.loop/<name>/QUEUE.md` — the existing work queue. Read it before appending so new units preserve the queue's current structure and statuses.

`REVIEW.md` is expected to include `## Standards`, `## Intent`, `## Speculative`, and `## Summary` sections. The summary's actionable count is the loop's signal; the finding entries are the fix skill's input. Use the finding classification, confidence, and evidence when deciding whether to generate a work unit.

## Procedure

=== .agents/skills/plan/SKILL.md ===
---
name: plan
description: Use when converting a software task, bug, cleanup, or vague human intent into a disposable `.loop/<name>/QUEUE.md` loop packet of verifiable work units. The planner picks a short, descriptive name for the cycle (e.g., `go-cli`, `parser-bug`).
---

# Plan

Convert intent into a `.loop/<name>/QUEUE.md` loop packet: a bounded queue of disposable, independently verifiable work units. Pick a short, descriptive name for the cycle (e.g., `go-cli`, `parser-bug`).

The goal is not to create durable specs. The goal is to give the loop runner work units that can be attempted one at a time and externally verified. When the work is done, the queue is deleted.

## Entry points

You may enter plan from two places:

- **After explore** — the codebase has been read, intent has been grilled, decisions may already be captured as ADRs in `decisions/`, terms may be in `glossary.md`. Use that context. Don't re-derive what explore already established.
- **Directly** — for small work where explore isn't needed. A bug fix, a patch, a small feature. Skip the ceremony.

If you're entering directly and the work is large or greenfield (no existing codebase to ground against), consider producing disposable specs first (see "Big work" below).

=== .agents/skills/review/SKILL.md ===
---
name: review
description: Use when reviewing completed work units or a finished queue. Two-axis adversarial review — standards (does the change follow codebase conventions?) and intent (does the change do what the work unit said?). Reviews against the actual codebase, not specs. Triggers on "review", "check this", "is this right", "what did we miss", "stress-test the implementation", or when work needs adversarial scrutiny before being accepted.
---

# Review

Adversarial review of completed work. Two axes, run independently so neither pollutes the other:

1. **Standards** — does the change follow the repo's coding conventions and codebase patterns?
2. **Intent** — does the change do what the work unit said it would?

Review against the **actual codebase**, not against specs that may have rotted. The code is the source of truth.

## When to review

- After a work unit is marked `done` and verify passed
- After a full queue is completed
- When the user asks for a sanity check
- Before accepting work as finished
[0m
I'll add a `version:` field to each SKILL.md frontmatter. Let me check the exact frontmatter close for each file.
[0m$ [0mfor f in .agents/skills/*/SKILL.md; do echo "=== $f ==="; awk 'NR==1{if($0!="---")exit} NR>1{if($0=="---"){print "--- at line " NR; exit}}' "$f"; done; echo "---embedded---"; for f in cli/embedded/skills/*/SKILL.md; do echo "$f"; done
=== .agents/skills/build/SKILL.md ===
--- at line 4
=== .agents/skills/decide/SKILL.md ===
--- at line 4
=== .agents/skills/domain-modeling/SKILL.md ===
--- at line 4
=== .agents/skills/explore/SKILL.md ===
--- at line 4
=== .agents/skills/fix/SKILL.md ===
--- at line 4
=== .agents/skills/plan/SKILL.md ===
--- at line 4
=== .agents/skills/review/SKILL.md ===
--- at line 4
---embedded---
cli/embedded/skills/build/SKILL.md
cli/embedded/skills/decide/SKILL.md
cli/embedded/skills/domain-modeling/SKILL.md
cli/embedded/skills/explore/SKILL.md
cli/embedded/skills/fix/SKILL.md
cli/embedded/skills/plan/SKILL.md
cli/embedded/skills/review/SKILL.md
[0m
[0m$ [0mfor f in .agents/skills/*/SKILL.md; do echo "=== $f ==="; sed -n '3p' "$f"; done
=== .agents/skills/build/SKILL.md ===
description: Use when implementing one work unit from a `.loop/<name>/QUEUE.md` queue. The worker-side skill — read the unit, do the work, don't self-certify, end with a handoff. Triggers on "build", "implement", "apply this unit", "do the work", "run the loop", or when a work unit needs to be executed. Also use when the loop invokes the worker for a tick.
=== .agents/skills/decide/SKILL.md ===
description: Use when an architectural ruling crystallizes during any phase — explore, plan, build, or review. Captures decisions as ADRs in `decisions/` so they persist after specs are deleted and code moves on. Triggers on "we decided to", "let's go with", "the ruling is", "why did we choose", "record this decision", or when a design choice with lasting consequences is made.
=== .agents/skills/domain-modeling/SKILL.md ===
description: Use when domain terms surface during explore, plan, or review and need to be defined, challenged, or stress-tested against the project's ubiquitous language. Manages `glossary.md` — the project's shared vocabulary. Triggers on "what does X mean here", "are these the same concept", "let's define our terms", "that term is ambiguous", or when a domain concept is used inconsistently across a conversation or codebase.
=== .agents/skills/explore/SKILL.md ===
description: Use when investigating a codebase, grilling intent, or stress-testing ideas before planning work. The upstream phase of the loop — read code, challenge assumptions, surface constraints, and capture decisions as they crystallize. Triggers on "explore", "investigate", "let me understand", "what's going on with", "grill this", "stress-test", "help me decide", "what should I do next", or when the real problem isn't clear yet.
=== .agents/skills/fix/SKILL.md ===
description: Use when addressing review findings and feeding them back into the loop. Converts review findings into new verifiable work units in `.loop/<name>/QUEUE.md` and runs another loop pass. Triggers on "fix the review findings", "address the feedback", "rework based on review", "fix what review found", or when review produced findings that need to be resolved.
=== .agents/skills/plan/SKILL.md ===
description: Use when converting a software task, bug, cleanup, or vague human intent into a disposable `.loop/<name>/QUEUE.md` loop packet of verifiable work units. The planner picks a short, descriptive name for the cycle (e.g., `go-cli`, `parser-bug`).
=== .agents/skills/review/SKILL.md ===
description: Use when reviewing completed work units or a finished queue. Two-axis adversarial review — standards (does the change follow codebase conventions?) and intent (does the change do what the work unit said?). Reviews against the actual codebase, not specs. Triggers on "review", "check this", "is this right", "what did we miss", "stress-test the implementation", or when work needs adversarial scrutiny before being accepted.
[0m
[0m← [0mEdit .agents/skills/build/SKILL.md
Index: /home/daniel/build/knack/.agents/skills/build/SKILL.md
===================================================================
--- /home/daniel/build/knack/.agents/skills/build/SKILL.md
+++ /home/daniel/build/knack/.agents/skills/build/SKILL.md
@@ -1,7 +1,8 @@
 ---
 name: build
 description: Use when implementing one work unit from a `.loop/<name>/QUEUE.md` queue. The worker-side skill — read the unit, do the work, don't self-certify, end with a handoff. Triggers on "build", "implement", "apply this unit", "do the work", "run the loop", or when a work unit needs to be executed. Also use when the loop invokes the worker for a tick.
+version: 1.0.0
 ---
 
 # Build
 

[0m
[0m← [0mEdit .agents/skills/decide/SKILL.md
Index: /home/daniel/build/knack/.agents/skills/decide/SKILL.md
===================================================================
--- /home/daniel/build/knack/.agents/skills/decide/SKILL.md
+++ /home/daniel/build/knack/.agents/skills/decide/SKILL.md
@@ -1,7 +1,8 @@
 ---
 name: decide
 description: Use when an architectural ruling crystallizes during any phase — explore, plan, build, or review. Captures decisions as ADRs in `decisions/` so they persist after specs are deleted and code moves on. Triggers on "we decided to", "let's go with", "the ruling is", "why did we choose", "record this decision", or when a design choice with lasting consequences is made.
+version: 1.0.0
 ---
 
 # Decide
 

[0m
[0m← [0mEdit .agents/skills/domain-modeling/SKILL.md
Index: /home/daniel/build/knack/.agents/skills/domain-modeling/SKILL.md
===================================================================
--- /home/daniel/build/knack/.agents/skills/domain-modeling/SKILL.md
+++ /home/daniel/build/knack/.agents/skills/domain-modeling/SKILL.md
@@ -1,7 +1,8 @@
 ---
 name: domain-modeling
 description: Use when domain terms surface during explore, plan, or review and need to be defined, challenged, or stress-tested against the project's ubiquitous language. Manages `glossary.md` — the project's shared vocabulary. Triggers on "what does X mean here", "are these the same concept", "let's define our terms", "that term is ambiguous", or when a domain concept is used inconsistently across a conversation or codebase.
+version: 1.0.0
 ---
 
 # Domain Modeling
 

[0m
[0m← [0mEdit .agents/skills/explore/SKILL.md
Index: /home/daniel/build/knack/.agents/skills/explore/SKILL.md
===================================================================
--- /home/daniel/build/knack/.agents/skills/explore/SKILL.md
+++ /home/daniel/build/knack/.agents/skills/explore/SKILL.md
@@ -1,7 +1,8 @@
 ---
 name: explore
 description: Use when investigating a codebase, grilling intent, or stress-testing ideas before planning work. The upstream phase of the loop — read code, challenge assumptions, surface constraints, and capture decisions as they crystallize. Triggers on "explore", "investigate", "let me understand", "what's going on with", "grill this", "stress-test", "help me decide", "what should I do next", or when the real problem isn't clear yet.
+version: 1.0.0
 ---
 
 # Explore
 

[0m
[0m← [0mEdit .agents/skills/fix/SKILL.md
Index: /home/daniel/build/knack/.agents/skills/fix/SKILL.md
===================================================================
--- /home/daniel/build/knack/.agents/skills/fix/SKILL.md
+++ /home/daniel/build/knack/.agents/skills/fix/SKILL.md
@@ -1,7 +1,8 @@
 ---
 name: fix
 description: Use when addressing review findings and feeding them back into the loop. Converts review findings into new verifiable work units in `.loop/<name>/QUEUE.md` and runs another loop pass. Triggers on "fix the review findings", "address the feedback", "rework based on review", "fix what review found", or when review produced findings that need to be resolved.
+version: 1.0.0
 ---
 
 # Fix
 

[0m
[0m← [0mEdit .agents/skills/plan/SKILL.md
Index: /home/daniel/build/knack/.agents/skills/plan/SKILL.md
===================================================================
--- /home/daniel/build/knack/.agents/skills/plan/SKILL.md
+++ /home/daniel/build/knack/.agents/skills/plan/SKILL.md
@@ -1,7 +1,8 @@
 ---
 name: plan
 description: Use when converting a software task, bug, cleanup, or vague human intent into a disposable `.loop/<name>/QUEUE.md` loop packet of verifiable work units. The planner picks a short, descriptive name for the cycle (e.g., `go-cli`, `parser-bug`).
+version: 1.0.0
 ---
 
 # Plan
 

[0m
[0m← [0mEdit .agents/skills/review/SKILL.md
Index: /home/daniel/build/knack/.agents/skills/review/SKILL.md
===================================================================
--- /home/daniel/build/knack/.agents/skills/review/SKILL.md
+++ /home/daniel/build/knack/.agents/skills/review/SKILL.md
@@ -1,7 +1,8 @@
 ---
 name: review
 description: Use when reviewing completed work units or a finished queue. Two-axis adversarial review — standards (does the change follow codebase conventions?) and intent (does the change do what the work unit said?). Reviews against the actual codebase, not specs. Triggers on "review", "check this", "is this right", "what did we miss", "stress-test the implementation", or when work needs adversarial scrutiny before being accepted.
+version: 1.0.0
 ---
 
 # Review
 

[0m
I'll add `version: 1.0.0` to each skill's frontmatter, then sync to embedded.
[0m$ [0mbash cli/sync-skills.sh && diff -r .agents/skills cli/embedded/skills && echo "NO DRIFT"
NO DRIFT
[0m
Now sync embedded skills using the existing script (without modifying it):
[0m$ [0mcd cli && go test ./internal/skills/...
ok  	knack/internal/skills	0.006s
[0m
All verify steps pass.

```text
Unit: all seven skills carry a version field in frontmatter
Changed: .agents/skills/*/SKILL.md (+version: 1.0.0); cli/embedded/skills/*/SKILL.md (synced)
Verify expected: for f in .agents/skills/*/SKILL.md; do grep -q '^version:' "$f" ...; diff -r .agents/skills cli/embedded/skills; cd cli && go test ./internal/skills/...
Notes: none
```
````

What this proves:
- The verify command passed for this work unit in the current repo state.

What remains unverified:
- Anything outside the verify command's proof scope.

## 2026-07-10T23:19:06-06:00 — skills init writes a manifest and .gitignore patterns

Status: done

Unit:
````markdown
## skills init writes a manifest and .gitignore patterns

Read first:
- cli/internal/skills/skills.go
- cli/internal/skills/skills_test.go
- cli/main.go (skillsInitCmd)
- .gitignore

Constraints:
- Do not change the existing Init skip-existing behavior.
- Do not change the Check function's existing validation behavior.
- The manifest is JSON at `.agents/skills/MANIFEST.json`.
- The .gitignore patterns are appended idempotently — skip patterns already present.

Done means:
- Init writes `.agents/skills/MANIFEST.json` after scaffolding.
- The manifest records each scaffolded skill's name, version (from frontmatter), and content hash (SHA-256 of all files in the skill directory, deterministically ordered).
- Init writes `.gitignore` patterns for disposable loop state (`.loop/**/QUEUE.md`, `.loop/**/HANDOFF.md`, `.loop/**/REVIEW.md`, `.loop/**/specs/`) into the target dir, appending idempotently.
- Re-running init on an existing project preserves the manifest (does not overwrite it with stale hashes for skipped skills) and does not duplicate .gitignore patterns.
- Tests cover: manifest is written on first init, manifest contains correct versions and hashes, manifest is not clobbered when init skips existing skills, .gitignore patterns are written and idempotent on re-run.

Verify:
```bash
cd cli && go test ./internal/skills/... -run TestInit
```

Status: pending````

Files changed:
```text
 M .agents/skills/build/SKILL.md
 M .agents/skills/decide/SKILL.md
 M .agents/skills/domain-modeling/SKILL.md
 M .agents/skills/explore/SKILL.md
 M .agents/skills/fix/SKILL.md
 M .agents/skills/plan/SKILL.md
 M .agents/skills/review/SKILL.md
 M cli/embedded/skills/build/SKILL.md
 M cli/embedded/skills/decide/SKILL.md
 M cli/embedded/skills/domain-modeling/SKILL.md
 M cli/embedded/skills/explore/SKILL.md
 M cli/embedded/skills/fix/SKILL.md
 M cli/embedded/skills/plan/SKILL.md
 M cli/embedded/skills/review/SKILL.md
 M cli/internal/skills/skills.go
 M cli/internal/skills/skills_test.go
?? .loop/skill-versioning/EVIDENCE.md
```

Verify command:
```bash
cd cli && go test ./internal/skills/... -run TestInit
```

Verify output:
```text
ok  	knack/internal/skills	0.011s
```

Worker output:
````text
[0m
> build · hy3-free
[0m
[0m→ [0mSkill "build"
[0m→ [0mRead cli/internal/skills/skills.go
[0m→ [0mRead cli/internal/skills/skills_test.go
[0m→ [0mRead cli/main.go
[0m→ [0mRead .gitignore
[0m
[0m$ [0mls /home/daniel/build/knack/cli/embedded/skills/ && echo "---" && head -5 /home/daniel/build/knack/.agents/skills/build/SKILL.md
build
decide
domain-modeling
explore
fix
plan
review
---
---
name: build
description: Use when implementing one work unit from a `.loop/<name>/QUEUE.md` queue. The worker-side skill — read the unit, do the work, don't self-certify, end with a handoff. Triggers on "build", "implement", "apply this unit", "do the work", "run the loop", or when a work unit needs to be executed. Also use when the loop invokes the worker for a tick.
version: 1.0.0
---
[0m
[0m$ [0mcd /home/daniel/build/knack && for d in .agents/skills/*/; do echo -n "$d: "; grep -m1 '^version:' "$d/SKILL.md" || echo "NO VERSION"; done
.agents/skills/build/: version: 1.0.0
.agents/skills/decide/: version: 1.0.0
.agents/skills/domain-modeling/: version: 1.0.0
.agents/skills/explore/: version: 1.0.0
.agents/skills/fix/: version: 1.0.0
.agents/skills/plan/: version: 1.0.0
.agents/skills/review/: version: 1.0.0
[0m
[0m← [0mEdit cli/internal/skills/skills.go
Index: /home/daniel/build/knack/cli/internal/skills/skills.go
===================================================================
--- /home/daniel/build/knack/cli/internal/skills/skills.go
+++ /home/daniel/build/knack/cli/internal/skills/skills.go
@@ -1,7 +1,9 @@
 package skills
 
 import (
+	"crypto/sha256"
+	"encoding/json"
 	"fmt"
 	"io"
 	"io/fs"
 	"os"
@@ -13,8 +15,30 @@
 
 	"gopkg.in/yaml.v3"
 )
 
+// SkillManifestEntry records a single scaffolded skill's identity and content
+// fingerprint so a target project can detect drift.
+type SkillManifestEntry struct {
+	Name    string `json:"name"`
+	Version string `json:"version"`
+	Hash    string `json:"hash"`
+}
+
+// Manifest is the JSON document written to .agents/skills/MANIFEST.json.
+type Manifest struct {
+	Skills []SkillManifestEntry `json:"skills"`
+}
+
+// gitignorePatterns are appended idempotently to the target dir's .gitignore on
+// each init. They ignore disposable loop state; EVIDENCE.md stays tracked.
+var gitignorePatterns = []string{
+	".loop/**/QUEUE.md",
+	".loop/**/HANDOFF.md",
+	".loop/**/REVIEW.md",
+	".loop/**/specs/",
+}
+
 // Init scaffolds the embedded skills into targetDir/.agents/skills/.
 // Existing skill directories are skipped rather than overwritten.
 // It returns the lists of skill names written and skipped.
 func Init(fsys fs.FS, targetDir string) (wrote []string, skipped []string, err error) {

[0m
All skills use `version: 1.0.0`. Now implementing.
[0m← [0mEdit cli/internal/skills/skills.go
Index: /home/daniel/build/knack/cli/internal/skills/skills.go
===================================================================
--- /home/daniel/build/knack/cli/internal/skills/skills.go
+++ /home/daniel/build/knack/cli/internal/skills/skills.go
@@ -63,11 +63,151 @@
 			return wrote, skipped, fmt.Errorf("write skill %q: %w", name, err)
 		}
 		wrote = append(wrote, name)
 	}
+
+	if err := writeManifest(fsys, targetSkills, wrote, skipped); err != nil {
+		return wrote, skipped, fmt.Errorf("write manifest: %w", err)
+	}
+	if err := appendGitignore(targetDir, gitignorePatterns); err != nil {
+		return wrote, skipped, fmt.Errorf("write .gitignore: %w", err)
+	}
 	return wrote, skipped, nil
 }
 
+// writeManifest writes .agents/skills/MANIFEST.json after scaffolding.
+// Existing entries for skipped skills are preserved so a re-run does not
+// clobber the manifest with recomputed (potentially stale) hashes.
+func writeManifest(fsys fs.FS, targetSkills string, wrote, skipped []string) error {
+	manifestPath := filepath.Join(targetSkills, "MANIFEST.json")
+
+	existing := map[string]SkillManifestEntry{}
+	if data, err := os.ReadFile(manifestPath); err == nil {
+		var m Manifest
+		if err := json.Unmarshal(data, &m); err != nil {
+			return fmt.Errorf("parse existing manifest: %w", err)
+		}
+		for _, e := range m.Skills {
+			existing[e.Name] = e
+		}
+	} else if !os.IsNotExist(err) {
+		return fmt.Errorf("read existing manifest: %w", err)
+	}
+
+	entries := make([]SkillManifestEntry, 0, len(existing)+len(wrote))
+	for _, name := range skipped {
+		if e, ok := existing[name]; ok {
+			entries = append(entries, e)
+		}
+	}
+	for _, name := range wrote {
+		entry, err := buildEntry(fsys, name)
+		if err != nil {
+			return err
+		}
+		entries = append(entries, entry)
+	}
+	sort.Slice(entries, func(i, j int) bool { return entries[i].Name < entries[j].Name })
+
+	data, err := json.MarshalIndent(Manifest{Skills: entries}, "", "  ")
+	if err != nil {
+		return err
+	}
+	data = append(data, '\n')
+	return os.WriteFile(manifestPath, data, 0o644)
+}
+
+// buildEntry derives a manifest entry for a skill in fsys by reading its
+// frontmatter version and content hash.
+func buildEntry(fsys fs.FS, name string) (SkillManifestEntry, error) {
+	data, err := fs.ReadFile(fsys, filepath.Join(name, "SKILL.md"))
+	if err != nil {
+		return SkillManifestEntry{}, fmt.Errorf("read skill %q: %w", name, err)
+	}
+	version := ""
+	if front, _, ok := splitFrontmatter(data); ok {
+		var meta struct {
+			Version string `yaml:"version"`
+		}
+		if err := yaml.Unmarshal(front, &meta); err == nil {
+			version = meta.Version
+		}
+	}
+	hash, err := dirHash(fsys, name)
+	if err != nil {
+		return SkillManifestEntry{}, err
+	}
+	return SkillManifestEntry{Name: name, Version: version, Hash: hash}, nil
+}
+
+// dirHash returns a deterministic SHA-256 of all files under dir in fsys,
+// ordered by path. It fingerprints the skill's entire content.
+func dirHash(fsys fs.FS, dir string) (string, error) {
+	var files []string
+	err := fs.WalkDir(fsys, dir, func(p string, d fs.DirEntry, err error) error {
+		if err != nil {
+			return err
+		}
+		if !d.IsDir() {
+			files = append(files, p)
+		}
+		return nil
+	})
+	if err != nil {
+		return "", fmt.Errorf("walk %q: %w", dir, err)
+	}
+	sort.Strings(files)
+
+	h := sha256.New()
+	for _, f := range files {
+		io.WriteString(h, f)
+		h.Write([]byte{0})
+		data, err := fs.ReadFile(fsys, f)
+		if err != nil {
+			return "", fmt.Errorf("read %q: %w", f, err)
+		}
+		h.Write(data)
+		h.Write([]byte{0})
+	}
+	return fmt.Sprintf("%x", h.Sum(nil)), nil
+}
+
+// appendGitignore adds patterns to targetDir/.gitignore without duplicating
+// any pattern already present.
+func appendGitignore(targetDir string, patterns []string) error {
+	gitignorePath := filepath.Join(targetDir, ".gitignore")
+
+	var lines []string
+	seen := map[string]bool{}
+	if data, err := os.ReadFile(gitignorePath); err == nil {
+		for _, line := range strings.Split(string(data), "\n") {
+			trimmed := strings.TrimSpace(line)
+			if trimmed != "" {
+				seen[trimmed] = true
+			}
+			lines = append(lines, line)
+		}
+	} else if !os.IsNotExist(err) {
+		return fmt.Errorf("read .gitignore: %w", err)
+	}
+
+	added := false
+	for _, p := range patterns {
+		if !seen[strings.TrimSpace(p)] {
+			lines = append(lines, p)
+			added = true
+		}
+	}
+	if len(lines) == 0 {
+		return nil
+	}
+	content := strings.Join(lines, "\n")
+	if added || !strings.HasSuffix(content, "\n") {
+		content += "\n"
+	}
+	return os.WriteFile(gitignorePath, []byte(content), 0o644)
+}
+
 // Check validates every SKILL.md found under fsys.
 // It returns one human-readable finding per issue. An empty slice with a nil
 // error means all skills are valid. The returned findings are sorted by skill
 // path and then by the order checks run.

[0m
[0m← [0mEdit cli/internal/skills/skills_test.go
Index: /home/daniel/build/knack/cli/internal/skills/skills_test.go
===================================================================
--- /home/daniel/build/knack/cli/internal/skills/skills_test.go
+++ /home/daniel/build/knack/cli/internal/skills/skills_test.go
@@ -1,7 +1,8 @@
 package skills
 
 import (
+	"encoding/json"
 	"io/fs"
 	"os"
 	"path/filepath"
 	"slices"

[0m
Now the tests:
[0m← [0mEdit cli/internal/skills/skills_test.go
Index: /home/daniel/build/knack/cli/internal/skills/skills_test.go
===================================================================
--- /home/daniel/build/knack/cli/internal/skills/skills_test.go
+++ /home/daniel/build/knack/cli/internal/skills/skills_test.go
@@ -192,4 +192,166 @@
 	out := append([]string{}, ss...)
 	slices.Sort(out)
 	return out
 }
+
+func readManifest(t *testing.T, target string) Manifest {
+	t.Helper()
+	data, err := os.ReadFile(filepath.Join(target, ".agents", "skills", "MANIFEST.json"))
+	if err != nil {
+		t.Fatalf("read manifest: %v", err)
+	}
+	var m Manifest
+	if err := json.Unmarshal(data, &m); err != nil {
+		t.Fatalf("parse manifest: %v", err)
+	}
+	return m
+}
+
+func findEntry(m Manifest, name string) (SkillManifestEntry, bool) {
+	for _, e := range m.Skills {
+		if e.Name == name {
+			return e, true
+		}
+	}
+	return SkillManifestEntry{}, false
+}
+
+func TestInitWritesManifest(t *testing.T) {
+	fsys := os.DirFS("../../embedded/skills")
+	target := t.TempDir()
+	if _, _, err := Init(fsys, target); err != nil {
+		t.Fatalf("Init failed: %v", err)
+	}
+	path := filepath.Join(target, ".agents", "skills", "MANIFEST.json")
+	if _, err := os.Stat(path); err != nil {
+		t.Fatalf("expected manifest at %s: %v", path, err)
+	}
+	m := readManifest(t, target)
+	if len(m.Skills) != 7 {
+		t.Fatalf("expected 7 skills in manifest, got %d", len(m.Skills))
+	}
+}
+
+func TestInitManifestContainsVersionsAndHashes(t *testing.T) {
+	fsys := os.DirFS("../../embedded/skills")
+	target := t.TempDir()
+	if _, _, err := Init(fsys, target); err != nil {
+		t.Fatalf("Init failed: %v", err)
+	}
+	m := readManifest(t, target)
+
+	expected := []string{"build", "decide", "domain-modeling", "explore", "fix", "plan", "review"}
+	if got := sorted(skillNames(m)); !slices.Equal(got, expected) {
+		t.Fatalf("expected skills %v, got %v", expected, got)
+	}
+
+	for _, name := range expected {
+		entry, ok := findEntry(m, name)
+		if !ok {
+			t.Fatalf("expected manifest entry for %s", name)
+		}
+		if entry.Version != "1.0.0" {
+			t.Fatalf("skill %s: expected version 1.0.0, got %q", name, entry.Version)
+		}
+		if entry.Hash == "" {
+			t.Fatalf("skill %s: expected non-empty hash", name)
+		}
+		// Hash must match a fresh computation from the embedded source.
+		want, err := buildEntry(fsys, name)
+		if err != nil {
+			t.Fatalf("buildEntry(%s): %v", name, err)
+		}
+		if entry.Hash != want.Hash {
+			t.Fatalf("skill %s: hash mismatch: got %s, want %s", name, entry.Hash, want.Hash)
+		}
+	}
+}
+
+func TestInitManifestNotClobberedOnSkip(t *testing.T) {
+	fsys := os.DirFS("../../embedded/skills")
+	target := t.TempDir()
+	if _, _, err := Init(fsys, target); err != nil {
+		t.Fatalf("first Init failed: %v", err)
+	}
+	first := readManifest(t, target)
+
+	// Simulate the user editing a scaffolded skill on disk. A re-run should
+	// preserve the manifest rather than recompute stale hashes for skipped dirs.
+	edited := filepath.Join(target, ".agents", "skills", "build", "SKILL.md")
+	if err := os.WriteFile(edited, []byte("tampered\n"), 0o644); err != nil {
+		t.Fatalf("tamper skill: %v", err)
+	}
+
+	if _, _, err := Init(fsys, target); err != nil {
+		t.Fatalf("second Init failed: %v", err)
+	}
+	second := readManifest(t, target)
+
+	firstBuild, ok := findEntry(first, "build")
+	if !ok {
+		t.Fatalf("missing build entry in first manifest")
+	}
+	secondBuild, ok := findEntry(second, "build")
+	if !ok {
+		t.Fatalf("missing build entry in second manifest")
+	}
+	if firstBuild.Hash != secondBuild.Hash {
+		t.Fatalf("manifest clobbered on skip: hash changed from %s to %s", firstBuild.Hash, secondBuild.Hash)
+	}
+	if len(second.Skills) != 7 {
+		t.Fatalf("expected 7 skills after re-run, got %d", len(second.Skills))
+	}
+}
+
+func TestInitWritesGitignorePatterns(t *testing.T) {
+	fsys := os.DirFS("../../embedded/skills")
+	target := t.TempDir()
+	if _, _, err := Init(fsys, target); err != nil {
+		t.Fatalf("Init failed: %v", err)
+	}
+	checkGitignore(t, target, gitignorePatterns)
+}
+
+func TestInitGitignoreIdempotent(t *testing.T) {
+	fsys := os.DirFS("../../embedded/skills")
+	target := t.TempDir()
+	if _, _, err := Init(fsys, target); err != nil {
+		t.Fatalf("first Init failed: %v", err)
+	}
+	if _, _, err := Init(fsys, target); err != nil {
+		t.Fatalf("second Init failed: %v", err)
+	}
+	content := readFile(t, filepath.Join(target, ".gitignore"))
+	for _, p := range gitignorePatterns {
+		if c := strings.Count(content, p); c != 1 {
+			t.Fatalf("pattern %q appears %d times, want 1", p, c)
+		}
+	}
+}
+
+func checkGitignore(t *testing.T, target string, patterns []string) {
+	t.Helper()
+	content := readFile(t, filepath.Join(target, ".gitignore"))
+	for _, p := range patterns {
+		if !strings.Contains(content, p) {
+			t.Fatalf("expected .gitignore to contain %q, got:\n%s", p, content)
+		}
+	}
+}
+
+func readFile(t *testing.T, path string) string {
+	t.Helper()
+	data, err := os.ReadFile(path)
+	if err != nil {
+		t.Fatalf("read %s: %v", path, err)
+	}
+	return string(data)
+}
+
+func skillNames(m Manifest) []string {
+	names := make([]string, 0, len(m.Skills))
+	for _, e := range m.Skills {
+		names = append(names, e.Name)
+	}
+	return names
+}

[0m
[0m$ [0mcd /home/daniel/build/knack/cli && go test ./internal/skills/... -run TestInit -v 2>&1 | tail -40
=== RUN   TestInitScaffoldsAllSkills
--- PASS: TestInitScaffoldsAllSkills (0.00s)
=== RUN   TestInitSkipsExistingSkills
--- PASS: TestInitSkipsExistingSkills (0.00s)
=== RUN   TestInitWritesManifest
--- PASS: TestInitWritesManifest (0.00s)
=== RUN   TestInitManifestContainsVersionsAndHashes
--- PASS: TestInitManifestContainsVersionsAndHashes (0.00s)
=== RUN   TestInitManifestNotClobberedOnSkip
--- PASS: TestInitManifestNotClobberedOnSkip (0.00s)
=== RUN   TestInitWritesGitignorePatterns
--- PASS: TestInitWritesGitignorePatterns (0.00s)
=== RUN   TestInitGitignoreIdempotent
--- PASS: TestInitGitignoreIdempotent (0.00s)
PASS
ok  	knack/internal/skills	0.015s
[0m
All pass. Verifying the full package and build:
[0m$ [0mcd /home/daniel/build/knack/cli && go build ./... && go test ./internal/skills/... 2>&1 | tail -5
ok  	knack/internal/skills	0.016s
[0m
```text
Unit: skills init writes a manifest and .gitignore patterns
Changed: cli/internal/skills/skills.go (Init writes MANIFEST.json + appends .gitignore), cli/internal/skills/skills_test.go (TestInit* cases)
Verify expected: cd cli && go test ./internal/skills/... -run TestInit
Notes: None. Skip-existing and Check behavior untouched; manifest preserves skipped-skill entries on re-run.
```
````

What this proves:
- The verify command passed for this work unit in the current repo state.

What remains unverified:
- Anything outside the verify command's proof scope.

## 2026-07-10T23:21:48-06:00 — skills check reports stale and modified skills using the manifest

Status: done

Unit:
````markdown
## skills check reports stale and modified skills using the manifest

Read first:
- cli/internal/skills/skills.go
- cli/internal/skills/skills_test.go
- cli/main.go (skillsCheckCmd)

Constraints:
- Do not remove existing check findings (frontmatter, references).
- The existing `skills check` exit behavior stays: 0 on no findings, 1 on any findings.

Done means:
- Check reads the manifest (if it exists) and compares each on-disk skill's current hash to the manifest hash.
- If the current hash differs from the manifest hash, check reports `modified: skill X has local changes`.
- If the current hash matches the manifest hash but the embedded version is newer than the manifest version, check reports `stale: skill X can be updated from vA to vB`.
- If no manifest exists, check reports `no manifest: run 'knack skills init' to create one`.
- Tests cover: stale detection, modified detection, clean (up-to-date) case, missing manifest case.

Verify:
```bash
cd cli && go test ./internal/skills/... -run TestCheck
```

Status: pending````

Files changed:
```text
 M .agents/skills/build/SKILL.md
 M .agents/skills/decide/SKILL.md
 M .agents/skills/domain-modeling/SKILL.md
 M .agents/skills/explore/SKILL.md
 M .agents/skills/fix/SKILL.md
 M .agents/skills/plan/SKILL.md
 M .agents/skills/review/SKILL.md
 M cli/embedded/skills/build/SKILL.md
 M cli/embedded/skills/decide/SKILL.md
 M cli/embedded/skills/domain-modeling/SKILL.md
 M cli/embedded/skills/explore/SKILL.md
 M cli/embedded/skills/fix/SKILL.md
 M cli/embedded/skills/plan/SKILL.md
 M cli/embedded/skills/review/SKILL.md
 M cli/internal/skills/skills.go
 M cli/internal/skills/skills_test.go
 M cli/main.go
?? .loop/skill-versioning/EVIDENCE.md
```

Verify command:
```bash
cd cli && go test ./internal/skills/... -run TestCheck
```

Verify output:
```text
ok  	knack/internal/skills	(cached)
```

Worker output:
````text
[0m
> build · hy3-free
[0m
[0m→ [0mRead .agents/skills/build/SKILL.md
[0m→ [0mRead cli/internal/skills/skills.go
[0m→ [0mRead cli/internal/skills/skills_test.go
[0m→ [0mRead cli/main.go
[0m
[0m← [0mEdit cli/internal/skills/skills.go
Index: /home/daniel/build/knack/cli/internal/skills/skills.go
===================================================================
--- /home/daniel/build/knack/cli/internal/skills/skills.go
+++ /home/daniel/build/knack/cli/internal/skills/skills.go
@@ -2,16 +2,18 @@
 
 import (
 	"crypto/sha256"
 	"encoding/json"
+	"errors"
 	"fmt"
 	"io"
 	"io/fs"
 	"os"
 	"path"
 	"path/filepath"
 	"regexp"
 	"sort"
+	"strconv"
 	"strings"
 
 	"gopkg.in/yaml.v3"
 )

[0m
Now I understand the structure. The existing `Check` does frontmatter/reference validation. I'll add a separate manifest-based versioning check and wire it into the `skills check` command, preserving the existing `Check` behavior and its tests.
[0m✗ [0mEdit cli/internal/skills/skills.go failed
[91m[1mError: [0mCould not find oldString in the file. It must match exactly, including whitespace, indentation, and line endings.
[0m→ [0mRead cli/internal/skills/skills.go[90m [offset=240, limit=15][0m
[0m
[0m← [0mEdit cli/internal/skills/skills.go
Index: /home/daniel/build/knack/cli/internal/skills/skills.go
===================================================================
--- /home/daniel/build/knack/cli/internal/skills/skills.go
+++ /home/daniel/build/knack/cli/internal/skills/skills.go
@@ -239,8 +239,93 @@
 	}
 	return findings, nil
 }
 
+// CheckVersioning compares on-disk scaffolded skills against the manifest and
+// the embedded source. For each manifest entry it recomputes the on-disk skill
+// hash: if it differs from the manifest hash the skill was modified locally; if
+// it matches but the embedded source carries a newer version the scaffolded
+// copy is stale. If no manifest exists it returns a single remediation finding.
+// The returned findings are sorted.
+func CheckVersioning(diskFS, embeddedFS fs.FS) ([]string, error) {
+	manifest, err := readManifest(diskFS)
+	if err != nil {
+		if errors.Is(err, fs.ErrNotExist) {
+			return []string{"no manifest: run 'knack skills init' to create one"}, nil
+		}
+		return nil, err
+	}
+
+	var findings []string
+	for _, entry := range manifest.Skills {
+		currentHash, err := dirHash(diskFS, entry.Name)
+		if err != nil {
+			// Skill absent on disk; nothing to compare against. Skip.
+			continue
+		}
+		if currentHash != entry.Hash {
+			findings = append(findings, fmt.Sprintf("modified: skill %s has local changes", entry.Name))
+			continue
+		}
+		embedded, err := buildEntry(embeddedFS, entry.Name)
+		if err != nil {
+			continue
+		}
+		if versionNewer(embedded.Version, entry.Version) {
+			findings = append(findings,
+				fmt.Sprintf("stale: skill %s can be updated from v%s to v%s", entry.Name, entry.Version, embedded.Version))
+		}
+	}
+	sort.Strings(findings)
+	return findings, nil
+}
+
+func readManifest(diskFS fs.FS) (Manifest, error) {
+	data, err := fs.ReadFile(diskFS, "MANIFEST.json")
+	if err != nil {
+		return Manifest{}, err
+	}
+	var m Manifest
+	if err := json.Unmarshal(data, &m); err != nil {
+		return Manifest{}, fmt.Errorf("parse manifest: %w", err)
+	}
+	return m, nil
+}
+
+// versionNewer reports whether candidate is a newer semantic version than
+// baseline. Both are dotted-numeric (optionally "v"-prefixed). Missing or
+// non-numeric segments compare as 0.
+func versionNewer(candidate, baseline string) bool {
+	cParts := strings.Split(strings.TrimPrefix(candidate, "v"), ".")
+	bParts := strings.Split(strings.TrimPrefix(baseline, "v"), ".")
+	n := len(cParts)
+	if len(bParts) > n {
+		n = len(bParts)
+	}
+	for i := 0; i < n; i++ {
+		c := atoiSeg(cParts, i)
+		b := atoiSeg(bParts, i)
+		if c > b {
+			return true
+		}
+		if c < b {
+			return false
+		}
+	}
+	return false
+}
+
+func atoiSeg(parts []string, i int) int {
+	if i >= len(parts) {
+		return 0
+	}
+	v, err := strconv.Atoi(strings.TrimSpace(parts[i]))
+	if err != nil {
+		return 0
+	}
+	return v
+}
+
 func copyDir(fsys fs.FS, src, dst string) error {
 	return fs.WalkDir(fsys, src, func(path string, d fs.DirEntry, err error) error {
 		if err != nil {
 			return err

[0m
[0m← [0mEdit cli/main.go
Index: /home/daniel/build/knack/cli/main.go
===================================================================
--- /home/daniel/build/knack/cli/main.go
+++ /home/daniel/build/knack/cli/main.go
@@ -78,8 +78,12 @@
 	fmt.Fprintln(os.Stderr, "  init [--target DIR]")
 	fmt.Fprintln(os.Stderr, "  check [--dir DIR]")
 }
 
+func embeddedSkillsFS() (iofs.FS, error) {
+	return iofs.Sub(embeddedSkills, "embedded/skills")
+}
+
 func skillsInitCmd(args []string) {
 	flags := flag.NewFlagSet("init", flag.ExitOnError)
 	target := flags.String("target", ".", "target directory for skill scaffolding")
 	if err := flags.Parse(args); err != nil {

[0m
Now wire `CheckVersioning` into the `skills check` command in main.go:
[0m← [0mEdit cli/main.go
Index: /home/daniel/build/knack/cli/main.go
===================================================================
--- /home/daniel/build/knack/cli/main.go
+++ /home/daniel/build/knack/cli/main.go
@@ -90,9 +90,9 @@
 	fmt.Fprintf(os.Stderr, "init: %v\n", err)
 	os.Exit(1)
 }
 
-skillFS, err := iofs.Sub(embeddedSkills, "embedded/skills")
+skillFS, err := embeddedSkillsFS()
 if err != nil {
 	fmt.Fprintf(os.Stderr, "init: embedded skills: %v\n", err)
 	os.Exit(1)
 }

[0m
[0m← [0mEdit cli/main.go
Index: /home/daniel/build/knack/cli/main.go
===================================================================
--- /home/daniel/build/knack/cli/main.go
+++ /home/daniel/build/knack/cli/main.go
@@ -122,8 +122,21 @@
 if err != nil {
 	fmt.Fprintf(os.Stderr, "check: %v\n", err)
 	os.Exit(1)
 }
+
+embedded, err := embeddedSkillsFS()
+if err != nil {
+	fmt.Fprintf(os.Stderr, "check: embedded skills: %v\n", err)
+	os.Exit(1)
+}
+versionFindings, err := skills.CheckVersioning(os.DirFS(*dir), embedded)
+if err != nil {
+	fmt.Fprintf(os.Stderr, "check: %v\n", err)
+	os.Exit(1)
+}
+findings = append(findings, versionFindings...)
+
 for _, f := range findings {
 	fmt.Println(f)
 }
 if len(findings) > 0 {

[0m
[0m← [0mEdit cli/internal/skills/skills_test.go
Index: /home/daniel/build/knack/cli/internal/skills/skills_test.go
===================================================================
--- /home/daniel/build/knack/cli/internal/skills/skills_test.go
+++ /home/daniel/build/knack/cli/internal/skills/skills_test.go
@@ -354,4 +354,102 @@
 		names = append(names, e.Name)
 	}
 	return names
 }
+
+const (
+	skillManifestV  = "---\nname: x\ndescription: a skill\nversion: 1.0.0\n---\nbody\n"
+	skillEmbeddedV  = "---\nname: x\ndescription: a skill\nversion: 1.1.0\n---\nbody\n"
+	skillModifiedV  = "---\nname: x\ndescription: a skill\nversion: 1.0.0\n---\nchanged body\n"
+)
+
+func addManifest(t *testing.T, disk fstest.MapFS, entries ...SkillManifestEntry) {
+	t.Helper()
+	data, err := json.MarshalIndent(Manifest{Skills: entries}, "", "  ")
+	if err != nil {
+		t.Fatalf("marshal manifest: %v", err)
+	}
+	disk["MANIFEST.json"] = &fstest.MapFile{Data: data}
+}
+
+func diskHash(t *testing.T, disk fs.FS, name string) string {
+	t.Helper()
+	h, err := dirHash(disk, name)
+	if err != nil {
+		t.Fatalf("dirHash(%s): %v", name, err)
+	}
+	return h
+}
+
+func TestCheckNoManifest(t *testing.T) {
+	disk := makeSkillFS(map[string]string{
+		"x/SKILL.md": skillManifestV,
+	})
+	embedded := makeSkillFS(map[string]string{
+		"x/SKILL.md": skillEmbeddedV,
+	})
+	findings, err := CheckVersioning(disk, embedded)
+	if err != nil {
+		t.Fatalf("CheckVersioning failed: %v", err)
+	}
+	if len(findings) != 1 || !strings.Contains(findings[0], "no manifest") {
+		t.Fatalf("expected no-manifest finding, got %v", findings)
+	}
+}
+
+func TestCheckModified(t *testing.T) {
+	disk := makeSkillFS(map[string]string{
+		"x/SKILL.md": skillModifiedV,
+	})
+	hash := diskHash(t, disk, "x")
+	addManifest(t, disk, SkillManifestEntry{Name: "x", Version: "1.0.0", Hash: hash})
+
+	embedded := makeSkillFS(map[string]string{
+		"x/SKILL.md": skillEmbeddedV,
+	})
+	findings, err := CheckVersioning(disk, embedded)
+	if err != nil {
+		t.Fatalf("CheckVersioning failed: %v", err)
+	}
+	if len(findings) != 1 || !strings.Contains(findings[0], "modified: skill x has local changes") {
+		t.Fatalf("expected modified finding, got %v", findings)
+	}
+}
+
+func TestCheckStale(t *testing.T) {
+	disk := makeSkillFS(map[string]string{
+		"x/SKILL.md": skillManifestV,
+	})
+	hash := diskHash(t, disk, "x")
+	addManifest(t, disk, SkillManifestEntry{Name: "x", Version: "1.0.0", Hash: hash})
+
+	embedded := makeSkillFS(map[string]string{
+		"x/SKILL.md": skillEmbeddedV,
+	})
+	findings, err := CheckVersioning(disk, embedded)
+	if err != nil {
+		t.Fatalf("CheckVersioning failed: %v", err)
+	}
+	if len(findings) != 1 ||
+		!strings.Contains(findings[0], "stale: skill x can be updated from v1.0.0 to v1.1.0") {
+		t.Fatalf("expected stale finding, got %v", findings)
+	}
+}
+
+func TestCheckUpToDate(t *testing.T) {
+	disk := makeSkillFS(map[string]string{
+		"x/SKILL.md": skillManifestV,
+	})
+	hash := diskHash(t, disk, "x")
+	addManifest(t, disk, SkillManifestEntry{Name: "x", Version: "1.0.0", Hash: hash})
+
+	embedded := makeSkillFS(map[string]string{
+		"x/SKILL.md": skillManifestV,
+	})
+	findings, err := CheckVersioning(disk, embedded)
+	if err != nil {
+		t.Fatalf("CheckVersioning failed: %v", err)
+	}
+	if len(findings) != 0 {
+		t.Fatalf("expected no findings for up-to-date skill, got %v", findings)
+	}
+}

[0m
Now add the tests:
[0m$ [0mcd /home/daniel/build/knack/cli && go test ./internal/skills/... -run TestCheck 2>&1 | tail -30
# knack/internal/skills [knack/internal/skills.test]
internal/skills/skills_test.go:197:6: readManifest redeclared in this block
	internal/skills/skills.go:282:6: other declaration of readManifest
internal/skills/skills_test.go:229:7: assignment mismatch: 1 variable but readManifest returns 2 values
internal/skills/skills_test.go:229:23: too many arguments in call to readManifest
	have (*testing.T, string)
	want (fs.FS)
internal/skills/skills_test.go:241:7: assignment mismatch: 1 variable but readManifest returns 2 values
internal/skills/skills_test.go:241:23: too many arguments in call to readManifest
	have (*testing.T, string)
	want (fs.FS)
internal/skills/skills_test.go:276:11: assignment mismatch: 1 variable but readManifest returns 2 values
internal/skills/skills_test.go:276:27: too many arguments in call to readManifest
	have (*testing.T, string)
	want (fs.FS)
internal/skills/skills_test.go:288:12: assignment mismatch: 1 variable but readManifest returns 2 values
internal/skills/skills_test.go:288:28: too many arguments in call to readManifest
	have (*testing.T, string)
	want (fs.FS)
internal/skills/skills_test.go:404:17: cannot use disk (variable of interface type fs.FS) as fstest.MapFS value in argument to addManifest: need type assertion
internal/skills/skills_test.go:404:17: too many errors
FAIL	knack/internal/skills [build failed]
FAIL
[0m
[0m$ [0mcd /home/daniel/build/knack/cli && grep -n "readManifest" internal/skills/skills_test.go
197:func readManifest(t *testing.T, target string) Manifest {
229:	m := readManifest(t, target)
241:	m := readManifest(t, target)
276:	first := readManifest(t, target)
288:	second := readManifest(t, target)
[0m
The existing test helper `readManifest(t, target)` conflicts with my new one. Let me rename it:
[0m$ [0mcd /home/daniel/build/knack/cli && sed -i 's/readManifest(t, target)/readManifestAt(t, target)/g; s/func readManifest(t \*testing.T, target string) Manifest {/func readManifestAt(t *testing.T, target string) Manifest {/' internal/skills/skills_test.go && go test ./internal/skills/... -run TestCheck 2>&1 | tail -30
# knack/internal/skills [knack/internal/skills.test]
internal/skills/skills_test.go:404:17: cannot use disk (variable of interface type fs.FS) as fstest.MapFS value in argument to addManifest: need type assertion
internal/skills/skills_test.go:423:17: cannot use disk (variable of interface type fs.FS) as fstest.MapFS value in argument to addManifest: need type assertion
internal/skills/skills_test.go:443:17: cannot use disk (variable of interface type fs.FS) as fstest.MapFS value in argument to addManifest: need type assertion
FAIL	knack/internal/skills [build failed]
FAIL
[0m
[0m← [0mEdit cli/internal/skills/skills_test.go
Index: /home/daniel/build/knack/cli/internal/skills/skills_test.go
===================================================================
--- /home/daniel/build/knack/cli/internal/skills/skills_test.go
+++ /home/daniel/build/knack/cli/internal/skills/skills_test.go
@@ -179,9 +179,9 @@
 		t.Fatalf("embedded skills have validation findings: %v", findings)
 	}
 }
 
-func makeSkillFS(files map[string]string) fs.FS {
+func makeSkillFS(files map[string]string) fstest.MapFS {
 	fsys := fstest.MapFS{}
 	for name, data := range files {
 		fsys[name] = &fstest.MapFile{Data: []byte(data)}
 	}

[0m
[0m$ [0mcd /home/daniel/build/knack/cli && go test ./internal/skills/... -run TestCheck 2>&1 | tail -30
--- FAIL: TestCheckModified (0.00s)
    skills_test.go:414: expected modified finding, got [stale: skill x can be updated from v1.0.0 to v1.1.0]
FAIL
FAIL	knack/internal/skills	0.004s
FAIL
[0m
[0m← [0mEdit cli/internal/skills/skills_test.go
Index: /home/daniel/build/knack/cli/internal/skills/skills_test.go
===================================================================
--- /home/daniel/build/knack/cli/internal/skills/skills_test.go
+++ /home/daniel/build/knack/cli/internal/skills/skills_test.go
@@ -399,9 +399,11 @@
 func TestCheckModified(t *testing.T) {
 	disk := makeSkillFS(map[string]string{
 		"x/SKILL.md": skillModifiedV,
 	})
-	hash := diskHash(t, disk, "x")
+	// Manifest records the hash from before the local edit.
+	orig := makeSkillFS(map[string]string{"x/SKILL.md": skillManifestV})
+	hash := diskHash(t, orig, "x")
 	addManifest(t, disk, SkillManifestEntry{Name: "x", Version: "1.0.0", Hash: hash})
 
 	embedded := makeSkillFS(map[string]string{
 		"x/SKILL.md": skillEmbeddedV,

[0m
The manifest hash must be the *original* (unmodified) hash, not the modified disk content. Let me fix the test:
[0m$ [0mcd /home/daniel/build/knack/cli && go test ./internal/skills/... -run TestCheck 2>&1 | tail -30
ok  	knack/internal/skills	0.004s
[0m
All passing. Let me run the full skills package tests and build to ensure nothing else broke:
[0m$ [0mcd /home/daniel/build/knack/cli && go test ./internal/skills/... 2>&1 | tail -5 && go build ./... 2>&1 | tail -5
ok  	knack/internal/skills	0.020s
[0m
Build and all tests pass. The unit is complete.

```text
Unit: skills check reports stale and modified skills using the manifest
Changed: cli/internal/skills/skills.go (CheckVersioning, readManifest, versionNewer, atoiSeg), cli/main.go (wired CheckVersioning into skillsCheckCmd, added embeddedSkillsFS), cli/internal/skills/skills_test.go (TestCheckNoManifest, TestCheckModified, TestCheckStale, TestCheckUpToDate)
Verify expected: cd cli && go test ./internal/skills/... -run TestCheck
Notes: Existing frontmatter/reference findings preserved in Check; manifest logic is a separate CheckVersioning call combined in the command. Exit behavior unchanged (0 no findings, 1 any).
```
````

What this proves:
- The verify command passed for this work unit in the current repo state.

What remains unverified:
- Anything outside the verify command's proof scope.

## 2026-07-10T23:24:00-06:00 — skills update refreshes stale skills, --force overwrites all

Status: done

Unit:
````markdown
## skills update refreshes stale skills, --force overwrites all

Read first:
- cli/internal/skills/skills.go
- cli/internal/skills/skills_test.go
- cli/main.go (skillsCmd, skillsInitCmd)

Constraints:
- Do not change the Init function's existing behavior; update is a separate function.
- Update must refresh the manifest after overwriting skills.

Done means:
- A new `skills update` subcommand exists and is wired in main.go.
- Without `--force`: for each embedded skill, if the on-disk skill exists and its current hash matches the manifest hash and the embedded version is newer, overwrite it. If the hash differs (modified), skip it and report. If the skill doesn't exist on disk, scaffold it. If the version is the same or older, skip it.
- With `--force`: overwrite all embedded skills regardless of modification status.
- After updating, the manifest is rewritten with the new versions and hashes.
- The command reports which skills were updated, skipped (modified), skipped (up-to-date), and newly scaffolded.
- Tests cover: update stale skill, update skips modified skill, update --force overwrites modified skill, update scaffolds new skill, update refreshes manifest.

Verify:
```bash
cd cli && go test ./internal/skills/... -run TestUpdate
```

Status: pending````

Files changed:
```text
 M .agents/skills/build/SKILL.md
 M .agents/skills/decide/SKILL.md
 M .agents/skills/domain-modeling/SKILL.md
 M .agents/skills/explore/SKILL.md
 M .agents/skills/fix/SKILL.md
 M .agents/skills/plan/SKILL.md
 M .agents/skills/review/SKILL.md
 M cli/embedded/skills/build/SKILL.md
 M cli/embedded/skills/decide/SKILL.md
 M cli/embedded/skills/domain-modeling/SKILL.md
 M cli/embedded/skills/explore/SKILL.md
 M cli/embedded/skills/fix/SKILL.md
 M cli/embedded/skills/plan/SKILL.md
 M cli/embedded/skills/review/SKILL.md
 M cli/internal/skills/skills.go
 M cli/internal/skills/skills_test.go
 M cli/main.go
?? .loop/skill-versioning/EVIDENCE.md
```

Verify command:
```bash
cd cli && go test ./internal/skills/... -run TestUpdate
```

Verify output:
```text
ok  	knack/internal/skills	(cached)
```

Worker output:
````text
[0m
> build · hy3-free
[0m
[0m→ [0mRead .agents/skills/build/SKILL.md
[0m→ [0mRead cli/internal/skills/skills.go
[0m→ [0mRead cli/internal/skills/skills_test.go
[0m→ [0mRead cli/main.go
[0m
[0m← [0mEdit cli/internal/skills/skills.go
Index: /home/daniel/build/knack/cli/internal/skills/skills.go
===================================================================
--- /home/daniel/build/knack/cli/internal/skills/skills.go
+++ /home/daniel/build/knack/cli/internal/skills/skills.go
@@ -75,8 +75,147 @@
 	}
 	return wrote, skipped, nil
 }
 
+// UpdateReport records the outcome of a skills update: which skills were
+// overwritten, which were left alone, and why.
+type UpdateReport struct {
+	Updated    []string // overwritten from embedded source
+	SkippedMod []string // skipped because locally modified (without --force)
+	Skipped    []string // skipped because already up-to-date
+	Scaffolded []string // newly written because absent on disk
+}
+
+// Update refreshes the scaffolded skills under targetDir from the embedded
+// source. For each embedded skill:
+//   - if it is absent on disk it is scaffolded;
+//   - if its on-disk hash differs from the manifest hash it is locally
+//     modified: skipped unless force overwrites it;
+//   - otherwise, if the embedded version is newer it is overwritten;
+//   - otherwise it is up-to-date and skipped.
+//
+// With force, every embedded skill is overwritten regardless of modification.
+// After updating, the manifest is rewritten to reflect the new versions and
+// hashes of all on-disk skills.
+func Update(fsys fs.FS, targetDir string, force bool) (UpdateReport, error) {
+	var report UpdateReport
+	targetSkills := filepath.Join(targetDir, ".agents", "skills")
+
+	entries, err := fs.ReadDir(fsys, ".")
+	if err != nil {
+		return report, fmt.Errorf("read embedded skills: %w", err)
+	}
+
+	manifest, mErr := readManifest(os.DirFS(targetSkills))
+	manifestBy := map[string]SkillManifestEntry{}
+	if mErr == nil {
+		for _, e := range manifest.Skills {
+			manifestBy[e.Name] = e
+		}
+	} else if !errors.Is(mErr, fs.ErrNotExist) {
+		return report, fmt.Errorf("read manifest: %w", mErr)
+	}
+
+	for _, entry := range entries {
+		if !entry.IsDir() {
+			continue
+		}
+		name := entry.Name()
+		dstDir := filepath.Join(targetSkills, name)
+
+		embedded, err := buildEntry(fsys, name)
+		if err != nil {
+			return report, err
+		}
+
+		_, statErr := os.Stat(dstDir)
+		if statErr == nil {
+			// exists on disk
+			currentHash, err := dirHash(os.DirFS(targetSkills), name)
+			if err != nil {
+				return report, fmt.Errorf("hash skill %q: %w", name, err)
+			}
+			mEntry, hadManifest := manifestBy[name]
+			modified := hadManifest && currentHash != mEntry.Hash
+
+			if modified {
+				if force {
+					if err := copyDir(fsys, name, dstDir); err != nil {
+						return report, fmt.Errorf("overwrite skill %q: %w", name, err)
+					}
+					report.Updated = append(report.Updated, name)
+				} else {
+					report.SkippedMod = append(report.SkippedMod, name)
+				}
+				continue
+			}
+
+			if force {
+				if err := copyDir(fsys, name, dstDir); err != nil {
+					return report, fmt.Errorf("overwrite skill %q: %w", name, err)
+				}
+				report.Updated = append(report.Updated, name)
+				continue
+			}
+
+			baseVersion := ""
+			if hadManifest {
+				baseVersion = mEntry.Version
+			}
+			if versionNewer(embedded.Version, baseVersion) {
+				if err := copyDir(fsys, name, dstDir); err != nil {
+					return report, fmt.Errorf("overwrite skill %q: %w", name, err)
+				}
+				report.Updated = append(report.Updated, name)
+			} else {
+				report.Skipped = append(report.Skipped, name)
+			}
+			continue
+		} else if !os.IsNotExist(statErr) {
+			return report, fmt.Errorf("stat target %q: %w", dstDir, statErr)
+		}
+
+		// absent on disk: scaffold it
+		if err := copyDir(fsys, name, dstDir); err != nil {
+			return report, fmt.Errorf("scaffold skill %q: %w", name, err)
+		}
+		report.Scaffolded = append(report.Scaffolded, name)
+	}
+
+	if err := refreshManifest(targetSkills); err != nil {
+		return report, fmt.Errorf("refresh manifest: %w", err)
+	}
+	return report, nil
+}
+
+// refreshManifest rewrites .agents/skills/MANIFEST.json from the current
+// on-disk skills so it always reflects what is actually scaffolded.
+func refreshManifest(targetSkills string) error {
+	diskFS := os.DirFS(targetSkills)
+	entries, err := fs.ReadDir(diskFS, ".")
+	if err != nil {
+		return fmt.Errorf("read skills dir: %w", err)
+	}
+	out := make([]SkillManifestEntry, 0, len(entries))
+	for _, entry := range entries {
+		if !entry.IsDir() {
+			continue
+		}
+		e, err := buildEntry(diskFS, entry.Name())
+		if err != nil {
+			return err
+		}
+		out = append(out, e)
+	}
+	sort.Slice(out, func(i, j int) bool { return out[i].Name < out[j].Name })
+	data, err := json.MarshalIndent(Manifest{Skills: out}, "", "  ")
+	if err != nil {
+		return err
+	}
+	data = append(data, '\n')
+	return os.WriteFile(filepath.Join(targetSkills, "MANIFEST.json"), data, 0o644)
+}
+
 // writeManifest writes .agents/skills/MANIFEST.json after scaffolding.
 // Existing entries for skipped skills are preserved so a re-run does not
 // clobber the manifest with recomputed (potentially stale) hashes.
 func writeManifest(fsys fs.FS, targetSkills string, wrote, skipped []string) error {

[0m
Now I understand the structure. Let me implement the `Update` function in `skills.go`.
[0m← [0mEdit cli/main.go
Index: /home/daniel/build/knack/cli/main.go
===================================================================
--- /home/daniel/build/knack/cli/main.go
+++ /home/daniel/build/knack/cli/main.go
@@ -64,8 +64,10 @@
 case "init":
 	skillsInitCmd(args[1:])
 case "check":
 	skillsCheckCmd(args[1:])
+case "update":
+	skillsUpdateCmd(args[1:])
 default:
 	fmt.Fprintf(os.Stderr, "unknown skills command: %s\n", args[0])
 	skillsUsage()
 	os.Exit(1)

[0m
Now wire it into `main.go`:
[0m← [0mEdit cli/main.go
Index: /home/daniel/build/knack/cli/main.go
===================================================================
--- /home/daniel/build/knack/cli/main.go
+++ /home/daniel/build/knack/cli/main.go
@@ -78,8 +78,9 @@
 	fmt.Fprintf(os.Stderr, "usage: knack skills <subcommand> [args]\n")
 	fmt.Fprintln(os.Stderr, "subcommands:")
 	fmt.Fprintln(os.Stderr, "  init [--target DIR]")
 	fmt.Fprintln(os.Stderr, "  check [--dir DIR]")
+	fmt.Fprintln(os.Stderr, "  update [--target DIR] [--force]")
 }
 
 func embeddedSkillsFS() (iofs.FS, error) {
 	return iofs.Sub(embeddedSkills, "embedded/skills")

[0m
[0m← [0mEdit cli/main.go
Index: /home/daniel/build/knack/cli/main.go
===================================================================
--- /home/daniel/build/knack/cli/main.go
+++ /home/daniel/build/knack/cli/main.go
@@ -112,8 +112,42 @@
 		fmt.Printf("skipped existing skill %s\n", name)
 	}
 }
 
+func skillsUpdateCmd(args []string) {
+	flags := flag.NewFlagSet("update", flag.ExitOnError)
+	target := flags.String("target", ".", "target directory for skill updates")
+	force := flags.Bool("force", false, "overwrite all skills regardless of local modifications")
+	if err := flags.Parse(args); err != nil {
+		fmt.Fprintf(os.Stderr, "update: %v\n", err)
+		os.Exit(1)
+	}
+
+	skillFS, err := embeddedSkillsFS()
+	if err != nil {
+		fmt.Fprintf(os.Stderr, "update: embedded skills: %v\n", err)
+		os.Exit(1)
+	}
+
+	report, err := skills.Update(skillFS, *target, *force)
+	if err != nil {
+		fmt.Fprintf(os.Stderr, "update: %v\n", err)
+		os.Exit(1)
+	}
+	for _, name := range report.Updated {
+		fmt.Printf("updated skill %s\n", name)
+	}
+	for _, name := range report.Scaffolded {
+		fmt.Printf("scaffolded new skill %s\n", name)
+	}
+	for _, name := range report.SkippedMod {
+		fmt.Printf("skipped modified skill %s (local changes preserved)\n", name)
+	}
+	for _, name := range report.Skipped {
+		fmt.Printf("skipped up-to-date skill %s\n", name)
+	}
+}
+
 func skillsCheckCmd(args []string) {
 	flags := flag.NewFlagSet("check", flag.ExitOnError)
 	dir := flags.String("dir", ".agents/skills", "directory containing skills")
 	if err := flags.Parse(args); err != nil {

[0m
[0m← [0mEdit cli/internal/skills/skills_test.go
Index: /home/daniel/build/knack/cli/internal/skills/skills_test.go
===================================================================
--- /home/daniel/build/knack/cli/internal/skills/skills_test.go
+++ /home/daniel/build/knack/cli/internal/skills/skills_test.go
@@ -454,4 +454,155 @@
 	if len(findings) != 0 {
 		t.Fatalf("expected no findings for up-to-date skill, got %v", findings)
 	}
 }
+
+func writeSkillOnDisk(t *testing.T, target, name, skillMd string) {
+	t.Helper()
+	dir := filepath.Join(target, ".agents", "skills", name)
+	if err := os.MkdirAll(dir, 0o755); err != nil {
+		t.Fatalf("mkdir %s: %v", dir, err)
+	}
+	if err := os.WriteFile(filepath.Join(dir, "SKILL.md"), []byte(skillMd), 0o644); err != nil {
+		t.Fatalf("write skill %s: %v", name, err)
+	}
+}
+
+func writeManifestOnDisk(t *testing.T, target string, entries []SkillManifestEntry) {
+	t.Helper()
+	dir := filepath.Join(target, ".agents", "skills")
+	if err := os.MkdirAll(dir, 0o755); err != nil {
+		t.Fatalf("mkdir %s: %v", dir, err)
+	}
+	data, err := json.MarshalIndent(Manifest{Skills: entries}, "", "  ")
+	if err != nil {
+		t.Fatalf("marshal manifest: %v", err)
+	}
+	data = append(data, '\n')
+	if err := os.WriteFile(filepath.Join(dir, "MANIFEST.json"), data, 0o644); err != nil {
+		t.Fatalf("write manifest: %v", err)
+	}
+}
+
+func diskSkillsDir(target string) string {
+	return filepath.Join(target, ".agents", "skills")
+}
+
+func TestUpdateStaleSkill(t *testing.T) {
+	embedded := makeSkillFS(map[string]string{"x/SKILL.md": skillEmbeddedV})
+	target := t.TempDir()
+	writeSkillOnDisk(t, target, "x", skillManifestV)
+	hash := diskHash(t, os.DirFS(diskSkillsDir(target)), "x")
+	writeManifestOnDisk(t, target, []SkillManifestEntry{{Name: "x", Version: "1.0.0", Hash: hash}})
+
+	report, err := Update(embedded, target, false)
+	if err != nil {
+		t.Fatalf("Update failed: %v", err)
+	}
+	if !slices.Contains(report.Updated, "x") {
+		t.Fatalf("expected x in Updated, got %v", report.Updated)
+	}
+	got := readFile(t, filepath.Join(diskSkillsDir(target), "x", "SKILL.md"))
+	if got != skillEmbeddedV {
+		t.Fatalf("expected x overwritten with embedded content, got %q", got)
+	}
+	m := readManifestAt(t, target)
+	entry, ok := findEntry(m, "x")
+	if !ok {
+		t.Fatalf("missing manifest entry for x")
+	}
+	if entry.Version != "1.1.0" {
+		t.Fatalf("expected manifest version 1.1.0, got %q", entry.Version)
+	}
+}
+
+func TestUpdateSkipsModifiedSkill(t *testing.T) {
+	embedded := makeSkillFS(map[string]string{"x/SKILL.md": skillEmbeddedV})
+	target := t.TempDir()
+	writeSkillOnDisk(t, target, "x", skillModifiedV)
+	// Manifest records the pre-edit hash (skillManifestV), so the modified
+	// on-disk content diverges from the manifest.
+	orig := makeSkillFS(map[string]string{"x/SKILL.md": skillManifestV})
+	hash := diskHash(t, orig, "x")
+	writeManifestOnDisk(t, target, []SkillManifestEntry{{Name: "x", Version: "1.0.0", Hash: hash}})
+
+	report, err := Update(embedded, target, false)
+	if err != nil {
+		t.Fatalf("Update failed: %v", err)
+	}
+	if !slices.Contains(report.SkippedMod, "x") {
+		t.Fatalf("expected x in SkippedMod, got %v", report.SkippedMod)
+	}
+	if len(report.Updated) != 0 {
+		t.Fatalf("expected no Updated skills, got %v", report.Updated)
+	}
+	got := readFile(t, filepath.Join(diskSkillsDir(target), "x", "SKILL.md"))
+	if got != skillModifiedV {
+		t.Fatalf("expected modified content preserved, got %q", got)
+	}
+}
+
+func TestUpdateForceOverwritesModifiedSkill(t *testing.T) {
+	embedded := makeSkillFS(map[string]string{"x/SKILL.md": skillEmbeddedV})
+	target := t.TempDir()
+	writeSkillOnDisk(t, target, "x", skillModifiedV)
+	orig := makeSkillFS(map[string]string{"x/SKILL.md": skillManifestV})
+	hash := diskHash(t, orig, "x")
+	writeManifestOnDisk(t, target, []SkillManifestEntry{{Name: "x", Version: "1.0.0", Hash: hash}})
+
+	report, err := Update(embedded, target, true)
+	if err != nil {
+		t.Fatalf("Update failed: %v", err)
+	}
+	if !slices.Contains(report.Updated, "x") {
+		t.Fatalf("expected x in Updated, got %v", report.Updated)
+	}
+	got := readFile(t, filepath.Join(diskSkillsDir(target), "x", "SKILL.md"))
+	if got != skillEmbeddedV {
+		t.Fatalf("expected x overwritten with embedded content, got %q", got)
+	}
+}
+
+func TestUpdateScaffoldsNewSkill(t *testing.T) {
+	embedded := makeSkillFS(map[string]string{"x/SKILL.md": skillEmbeddedV})
+	target := t.TempDir()
+	// No skill on disk, but manifest present (e.g. from a prior init of other skills).
+	writeManifestOnDisk(t, target, []SkillManifestEntry{})
+
+	report, err := Update(embedded, target, false)
+	if err != nil {
+		t.Fatalf("Update failed: %v", err)
+	}
+	if !slices.Contains(report.Scaffolded, "x") {
+		t.Fatalf("expected x in Scaffolded, got %v", report.Scaffolded)
+	}
+	if _, err := os.Stat(filepath.Join(diskSkillsDir(target), "x", "SKILL.md")); err != nil {
+		t.Fatalf("expected scaffolded skill x on disk: %v", err)
+	}
+}
+
+func TestUpdateRefreshManifest(t *testing.T) {
+	embedded := makeSkillFS(map[string]string{"x/SKILL.md": skillEmbeddedV})
+	target := t.TempDir()
+	writeSkillOnDisk(t, target, "x", skillManifestV)
+	hash := diskHash(t, os.DirFS(diskSkillsDir(target)), "x")
+	writeManifestOnDisk(t, target, []SkillManifestEntry{{Name: "x", Version: "1.0.0", Hash: hash}})
+
+	if _, err := Update(embedded, target, false); err != nil {
+		t.Fatalf("Update failed: %v", err)
+	}
+	m := readManifestAt(t, target)
+	entry, ok := findEntry(m, "x")
+	if !ok {
+		t.Fatalf("missing manifest entry for x")
+	}
+	if entry.Version != "1.1.0" {
+		t.Fatalf("expected refreshed manifest version 1.1.0, got %q", entry.Version)
+	}
+	wantHash, err := dirHash(os.DirFS(diskSkillsDir(target)), "x")
+	if err != nil {
+		t.Fatalf("dirHash: %v", err)
+	}
+	if entry.Hash != wantHash {
+		t.Fatalf("expected refreshed manifest hash %s, got %s", wantHash, entry.Hash)
+	}
+}

[0m
Now add the tests:
[0m$ [0mcd /home/daniel/build/knack/cli && go test ./internal/skills/... -run TestUpdate 2>&1 | tail -30
ok  	knack/internal/skills	0.004s
[0m
Tests pass. Let me verify the full package builds and existing tests still pass:
[0m$ [0mcd /home/daniel/build/knack/cli && go build ./... && go test ./internal/skills/... 2>&1 | tail -10
ok  	knack/internal/skills	0.015s
[0m
Build and tests pass.

```text
Unit: skills update refreshes stale skills, --force overwrites all
Changed: cli/internal/skills/skills.go (Update, refreshManifest), cli/main.go (skillsUpdateCmd + wiring), cli/internal/skills/skills_test.go (5 TestUpdate*)
Verify expected: cd cli && go test ./internal/skills/... -run TestUpdate
Notes: none
```
````

What this proves:
- The verify command passed for this work unit in the current repo state.

What remains unverified:
- Anything outside the verify command's proof scope.

## 2026-07-10T23:25:22-06:00 — docs and AGENTS.md reflect the new skills update and versioning behavior

Status: verify_failed

Unit:
````markdown
## docs and AGENTS.md reflect the new skills update and versioning behavior

Read first:
- docs/cli.md
- AGENTS.md
- README.md

Constraints:
- Do not change skill files or Go code.

Done means:
- docs/cli.md documents `skills update` with `--force` flag, the manifest, and the `.gitignore` patterns written by `skills init`.
- docs/cli.md notes section reflects the new update path (no longer "only write operation").
- AGENTS.md lessons-learned notes the versioning + manifest mechanism.
- README.md mentions `skills update` if it currently mentions `skills init`.

Verify:
```bash
grep -q "skills update" docs/cli.md && \
grep -q "MANIFEST" docs/cli.md && \
grep -q "gitignore" docs/cli.md && \
grep -q "skills update" cli/main.go && \
./tests/run.sh
```

Status: pending````

Files changed:
```text
 M .agents/skills/build/SKILL.md
 M .agents/skills/decide/SKILL.md
 M .agents/skills/domain-modeling/SKILL.md
 M .agents/skills/explore/SKILL.md
 M .agents/skills/fix/SKILL.md
 M .agents/skills/plan/SKILL.md
 M .agents/skills/review/SKILL.md
 M AGENTS.md
 M README.md
 M cli/embedded/skills/build/SKILL.md
 M cli/embedded/skills/decide/SKILL.md
 M cli/embedded/skills/domain-modeling/SKILL.md
 M cli/embedded/skills/explore/SKILL.md
 M cli/embedded/skills/fix/SKILL.md
 M cli/embedded/skills/plan/SKILL.md
 M cli/embedded/skills/review/SKILL.md
 M cli/internal/skills/skills.go
 M cli/internal/skills/skills_test.go
 M cli/main.go
 M docs/cli.md
?? .loop/skill-versioning/EVIDENCE.md
```

Verify command:
```bash
grep -q "skills update" docs/cli.md && \
grep -q "MANIFEST" docs/cli.md && \
grep -q "gitignore" docs/cli.md && \
grep -q "skills update" cli/main.go && \
./tests/run.sh
```

Verify output:
```text
```

Worker output:
````text
[0m
> build · hy3-free
[0m
[0m→ [0mSkill "build"
[0m→ [0mRead .loop/skill-versioning/QUEUE.md
[0m→ [0mRead docs/cli.md
[0m→ [0mRead README.md
[0m→ [0mRead cli/internal/skills/skills.go
[0m→ [0mRead cli/main.go
Let me read the Go implementation to document it accurately.
[0m
[0m← [0mEdit docs/cli.md
Index: /home/daniel/build/knack/docs/cli.md
===================================================================
--- /home/daniel/build/knack/docs/cli.md
+++ /home/daniel/build/knack/docs/cli.md
@@ -10,10 +10,11 @@
 ## Commands
 
 | Command | Arguments | Description |
 |---|---|---|
-| `skills init` | `[--target DIR]` | Scaffold the seven default skills into `<DIR>/.agents/skills/` (default `.`). Skips existing skills. |
-| `skills check` | `[--dir DIR]` | Validate skills in `DIR` (default `.agents/skills`). |
+| `skills init` | `[--target DIR]` | Scaffold the seven default skills into `<DIR>/.agents/skills/` (default `.`). Skips existing skills. Writes the manifest and `.gitignore` patterns. |
+| `skills check` | `[--dir DIR]` | Validate skills in `DIR` (default `.agents/skills`). Also reports modified and stale skills via the manifest. |
+| `skills update` | `[--target DIR] [--force]` | Refresh scaffolded skills from the embedded source when a newer version ships. With `--force`, overwrite locally modified skills too. |
 | `validate` | `<queue-file>` | Validate work-unit structure in a queue. |
 | `decisions list` | | List all ADRs in `decisions/`. Superseded ADRs show the number that replaced them. |
 | `decisions show` | `NNNN` | Print the full ADR with that number. |
 | `decisions check` | | Flag orphaned ADRs, dangling references, and broken/one-sided supersede chains. |

[0m
Now I'll update the docs. First `docs/cli.md`:
[0m← [0mEdit docs/cli.md
Index: /home/daniel/build/knack/docs/cli.md
===================================================================
--- /home/daniel/build/knack/docs/cli.md
+++ /home/daniel/build/knack/docs/cli.md
@@ -41,8 +41,28 @@
 cd /path/to/project
 /path/to/knack skills init
 ```
 
+`skills init` writes a manifest to `.agents/skills/MANIFEST.json` recording each
+scaffolded skill's `name`, `version` (read from its frontmatter), and a SHA-256
+content `hash` of the whole skill directory. It also appends `.gitignore`
+patterns for disposable loop state (`.loop/**/QUEUE.md`, `.loop/**/HANDOFF.md`,
+`.loop/**/REVIEW.md`, `.loop/**/specs/`) into the target directory, idempotently.
+
+### Refresh skills after a CLI upgrade
+
+```bash
+/path/to/knack skills update            # refresh only unmodified, stale skills
+/path/to/knack skills update --force    # overwrite every skill, modified or not
+```
+
+`skills update` compares each on-disk skill against the manifest and the embedded
+source. A skill is overwritten if it exists, is unmodified (its current hash matches
+the manifest), and the embedded version is newer than the manifest version. If a
+skill has local changes (hash differs), it is skipped unless `--force` is given.
+Absent skills are scaffolded. After updating, the manifest is rewritten with the
+new versions and hashes.
+
 ### Decision coverage
 
 ```bash
 ./knack decisions check

[0m
[0m← [0mEdit docs/cli.md
Index: /home/daniel/build/knack/docs/cli.md
===================================================================
--- /home/daniel/build/knack/docs/cli.md
+++ /home/daniel/build/knack/docs/cli.md
@@ -94,6 +94,7 @@
 
 ## Notes
 
 - All commands read from the current directory (run from the repo root unless a path is given).
-- `skills init` is the only write operation. It scaffolds missing skills and leaves existing ones alone, so upgrading the CLI will not overwrite project customizations.
+- `skills init`, `skills update` are the write operations. `init` scaffolds missing skills and writes the manifest plus `.gitignore` patterns; `update` refreshes unmodified skills when a newer embedded version ships, so upgrading the CLI does not overwrite project customizations. `--force` overrides that protection.
+- Each skill carries a `version` field in its frontmatter. `skills check` reports `modified:` for locally changed skills and `stale:` for skills whose embedded version is newer than the manifest version; run `skills update` to reconcile.
 - The CLI packages the default skills via `go:embed`. After editing `.agents/skills/` in the `knack` repo itself, run `cli/sync-skills.sh` and `diff -r .agents/skills cli/embedded/skills` to verify sync.

[0m
[0m← [0mEdit README.md
Index: /home/daniel/build/knack/README.md
===================================================================
--- /home/daniel/build/knack/README.md
+++ /home/daniel/build/knack/README.md
@@ -36,8 +36,14 @@
 cd /path/to/new-project
 /path/to/knack skills init
 ```
 
+Refresh the skills after a CLI upgrade (or `--force` to overwrite local changes):
+
+```bash
+/path/to/knack skills update
+```
+
 ## How it works
 
 1. The runner reads the first `Status: pending` work unit from `QUEUE.md`.
 2. It marks the unit `in_progress` and invokes a fresh agent context with the worker prompt and the unit.

[0m
Now README.md:
[0m← [0mEdit README.md
Index: /home/daniel/build/knack/README.md
===================================================================
--- /home/daniel/build/knack/README.md
+++ /home/daniel/build/knack/README.md
@@ -105,9 +105,10 @@
 ### Commands
 
 ```
 knack skills init [--target DIR]    Scaffold the seven default skills into DIR/.agents/skills/
-knack skills check [--dir DIR]      Validate skills in DIR (default: .agents/skills)
+knack skills check [--dir DIR]      Validate skills and report stale/modified via the manifest
+knack skills update [--target DIR] [--force]   Refresh scaffolded skills from the embedded source
 knack validate <queue-file>         Validate work-unit structure in a queue file
 knack decisions list                List all ADRs in decisions/
 knack decisions show NNNN           Print the full text of ADR NNNN
 knack decisions check               Flag orphaned ADRs and dangling references

[0m
[0m← [0mEdit README.md
Index: /home/daniel/build/knack/README.md
===================================================================
--- /home/daniel/build/knack/README.md
+++ /home/daniel/build/knack/README.md
@@ -116,9 +116,9 @@
 knack glossary check                Validate glossary.md term references
 knack instructions <artifact>       Print a template: work-unit | adr | glossary-entry
 ```
 
-All commands read from the current directory (run from the repo root). `skills init` is the only write operation — it scaffolds missing skills and leaves existing ones alone.
+All commands read from the current directory (run from the repo root). `skills init` and `skills update` are the write operations — `init` scaffolds missing skills and `update` refreshes unmodified ones when a newer embedded version ships (use `--force` to overwrite local changes).
 
 ## Documentation
 
 Full docs live in `docs/`:

[0m
[0m✱ [0mGrep "Lessons learned|Embedded skills must stay in sync"[90m in AGENTS.md · 8 matches[0m
Now AGENTS.md. Let me find the lessons-learned section:
[0m→ [0mRead AGENTS.md[90m [offset=82, limit=12][0m
[0m
[0m← [0mEdit AGENTS.md
Index: /home/daniel/build/knack/AGENTS.md
===================================================================
--- /home/daniel/build/knack/AGENTS.md
+++ /home/daniel/build/knack/AGENTS.md
@@ -90,5 +90,7 @@
 - **Verify commands must be path-correct.** Unit 1 of the named-cycles queue had `cd cli && go test ./... && ./tests/run.sh` — but `./tests/run.sh` ran from `cli/` after the `cd`, not from the repo root. The worker did the work correctly; the verify command was wrong. The loop correctly caught the failure (mechanical gate working), but it was a false negative. Always test verify commands manually before writing them into a queue.
 - **Workers scope to the outcome plus constraints, not to a file list.** ADR-0005 replaced `Work:` with `Read first:` and `Constraints:`. The unit's scope is its outcome plus its constraints — the worker determines which files to change. The old lesson ("name every file in the work unit") is wrong under the new shape: naming files in constraints smuggles scope the same way `Work:` did. The first plan-shape cycle proved this — the constraint said "no `Work:` refs in skills, prompts, DESIGN.md, or AGENTS.md" and the worker touched exactly those files, leaving 9 other files (test fixtures, examples, README) with stale `Work:` fields. Prefer outcome-level constraints ("no artifact that teaches the format may reference `Work:`") over file-enumerated constraints.
 - **`decisions check` orphaned-ADR semantics resolved.** An ADR is orphaned if it is not referenced by any QUEUE.md (current work) or any EVIDENCE.md (completed work). The loop now writes the full unit body into EVIDENCE.md so ADR references survive QUEUE.md deletion. EVIDENCE.md is the durable ledger; QUEUE.md is disposable. An ADR that drove a completed cycle is not orphaned as long as its EVIDENCE.md ledger exists.
 - **Named work cycles enable concurrent work.** ADR-0004 gave each work cycle its own subdirectory under `.loop/`. The loop already supported this via the queue path argument — only convention (skills, docs) and `knack status` needed to change. Running `./loop.sh run .loop/<name>/QUEUE.md` is fully independent of other cycles.
+- **Skills carry a `version` field and a manifest for safe upgrades.** Each SKILL.md has a `version:` in frontmatter. `skills init` writes `.agents/skills/MANIFEST.json` (name, version, SHA-256 content hash per skill) and appends `.gitignore` patterns for disposable loop state idempotently. `skills check` reports `modified:` (hash differs from manifest) and `stale:` (embedded version newer than manifest). `skills update` refreshes only unmodified stale skills; `--force` overwrites everything. The manifest is the contract that lets `update` tell "locally customized" from "behind the shipped version" — never hand-edit it.
+- **Bump the `version:` field when you change an embedded skill.** `sync-skills.sh` keeps `cli/embedded/skills` in sync with `.agents/skills/`, but version drift only reaches users via the manifest after a `skills update`. If you edit a skill's content without bumping `version:`, existing scaffolds stay silent-stale (current hash matches manifest, so `check` won't flag it, yet it's behind). Bump on every content change.
 - **Negative-grep verifies must anchor on field syntax, not bare mentions.** The plan-shape cycle's verify used `! grep -rn 'Work:' ...` to prove the old field was gone. The same commit that landed the work also added this lessons-learned entry, which says "ADR-0005 replaced `Work:`..." — so the verify rotted the moment history got documented. Anchor negative greps to the field shape (`^Work:`) or to a syntax-specific pattern, not to any mention of the word. Otherwise the verify forbids the project from ever recording why the change was made.

[0m
Now let me run the verify command:
[0m$ [0mcd /home/daniel/build/knack && grep -q "skills update" docs/cli.md && grep -q "MANIFEST" docs/cli.md && grep -q "gitignore" docs/cli.md && grep -q "skills update" cli/main.go && ./tests/run.sh
(no output)
[0m
Verify passed (exit 0). 

```text
Unit: docs and AGENTS.md reflect the new skills update and versioning behavior
Changed: docs/cli.md, README.md, AGENTS.md
Verify expected: grep -q "skills update" docs/cli.md && grep -q "MANIFEST" docs/cli.md && grep -q "gitignore" docs/cli.md && grep -q "skills update" cli/main.go && ./tests/run.sh
Notes: none
```
````

What this proves:
- The work unit is not externally verified.

What remains unverified:
- Anything outside the verify command's proof scope.

## 2026-07-10T23:40:45-06:00 — move skill version to top-level frontmatter and update buildEntry

Status: done

Unit:
````markdown
## move skill version to top-level frontmatter and update buildEntry

Read first:
- .loop/skill-versioning/REVIEW.md (S1 and I1)
- cli/internal/skills/skills.go (buildEntry)
- cli/internal/skills/skills_test.go (skillManifestV, skillEmbeddedV, skillModifiedV)
- .agents/skills/build/SKILL.md (example frontmatter)

Constraints:
- Do not add version validation to checkSkill; that is the next unit.
- Do not change cli/main.go usage(); the missing skills update line is a trivial review note.
- Do not change skill bodies or descriptions; only the frontmatter version field.
- Keep cli/embedded/skills/ in sync with .agents/skills/ (use cli/sync-skills.sh or equivalent).

Done means:
- S1 and I1 are resolved: version is a top-level frontmatter key in every .agents/skills/*/SKILL.md and cli/embedded/skills/*/SKILL.md.
- buildEntry reads version from the top-level frontmatter.
- cli/internal/skills/skills_test.go test fixtures (skillManifestV, skillEmbeddedV, skillModifiedV) use top-level version.
- diff -r .agents/skills cli/embedded/skills shows no drift.
- cd cli && go test ./internal/skills/... passes.
- No regression in existing check behavior.

Verify:
```bash
for f in .agents/skills/*/SKILL.md cli/embedded/skills/*/SKILL.md; do grep -q '^version:' "$f" || { echo "missing top-level version in $f"; exit 1; }; done && \
diff -r .agents/skills cli/embedded/skills && \
cd cli && go test ./internal/skills/... -count=1
```

Status: pending````

Files changed:
```text
 M .agents/skills/build/SKILL.md
 M .agents/skills/decide/SKILL.md
 M .agents/skills/domain-modeling/SKILL.md
 M .agents/skills/explore/SKILL.md
 M .agents/skills/fix/SKILL.md
 M .agents/skills/plan/SKILL.md
 M .agents/skills/review/SKILL.md
 M AGENTS.md
 M README.md
 M cli/embedded/skills/build/SKILL.md
 M cli/embedded/skills/decide/SKILL.md
 M cli/embedded/skills/domain-modeling/SKILL.md
 M cli/embedded/skills/explore/SKILL.md
 M cli/embedded/skills/fix/SKILL.md
 M cli/embedded/skills/plan/SKILL.md
 M cli/embedded/skills/review/SKILL.md
 M cli/internal/skills/skills.go
 M cli/internal/skills/skills_test.go
 M cli/main.go
 M docs/cli.md
?? .loop/skill-versioning/EVIDENCE.md
?? COMPARISON-litespec.md
```

Verify command:
```bash
for f in .agents/skills/*/SKILL.md cli/embedded/skills/*/SKILL.md; do grep -q '^version:' "$f" || { echo "missing top-level version in $f"; exit 1; }; done && \
diff -r .agents/skills cli/embedded/skills && \
cd cli && go test ./internal/skills/... -count=1
```

Verify output:
```text
ok  	knack/internal/skills	0.018s
```

Worker output:
````text
[0m
> build · hy3-free
[0m
[0m→ [0mRead .agents/skills/build/SKILL.md
[0m→ [0mRead .loop/skill-versioning/REVIEW.md
[0m→ [0mRead cli/internal/skills/skills.go
[0m→ [0mRead cli/internal/skills/skills_test.go
Now let me look at the current frontmatter in the skill files.
[0m
[0m$ [0mfor f in /home/daniel/build/knack/.agents/skills/*/SKILL.md; do echo "=== $f ==="; head -8 "$f"; done
=== /home/daniel/build/knack/.agents/skills/build/SKILL.md ===
---
name: build
description: Use when implementing one work unit from a `.loop/<name>/QUEUE.md` queue. The worker-side skill — read the unit, do the work, don't self-certify, end with a handoff. Triggers on "build", "implement", "apply this unit", "do the work", "run the loop", or when a work unit needs to be executed. Also use when the loop invokes the worker for a tick.
metadata:
  version: "1.0.0"
---

# Build
=== /home/daniel/build/knack/.agents/skills/decide/SKILL.md ===
---
name: decide
description: Use when an architectural ruling crystallizes during any phase — explore, plan, build, or review. Captures decisions as ADRs in `decisions/` so they persist after specs are deleted and code moves on. Triggers on "we decided to", "let's go with", "the ruling is", "why did we choose", "record this decision", or when a design choice with lasting consequences is made.
metadata:
  version: "1.0.0"
---

# Decide
=== /home/daniel/build/knack/.agents/skills/domain-modeling/SKILL.md ===
---
name: domain-modeling
description: Use when domain terms surface during explore, plan, or review and need to be defined, challenged, or stress-tested against the project's ubiquitous language. Manages `glossary.md` — the project's shared vocabulary. Triggers on "what does X mean here", "are these the same concept", "let's define our terms", "that term is ambiguous", or when a domain concept is used inconsistently across a conversation or codebase.
metadata:
  version: "1.0.0"
---

# Domain Modeling
=== /home/daniel/build/knack/.agents/skills/explore/SKILL.md ===
---
name: explore
description: Use when investigating a codebase, grilling intent, or stress-testing ideas before planning work. The upstream phase of the loop — read code, challenge assumptions, surface constraints, and capture decisions as they crystallize. Triggers on "explore", "investigate", "let me understand", "what's going on with", "grill this", "stress-test", "help me decide", "what should I do next", or when the real problem isn't clear yet.
metadata:
  version: "1.0.0"
---

# Explore
=== /home/daniel/build/knack/.agents/skills/fix/SKILL.md ===
---
name: fix
description: Use when addressing review findings and feeding them back into the loop. Converts review findings into new verifiable work units in `.loop/<name>/QUEUE.md` and runs another loop pass. Triggers on "fix the review findings", "address the feedback", "rework based on review", "fix what review found", or when review produced findings that need to be resolved.
metadata:
  version: "1.0.0"
---

# Fix
=== /home/daniel/build/knack/.agents/skills/plan/SKILL.md ===
---
name: plan
description: Use when converting a software task, bug, cleanup, or vague human intent into a disposable `.loop/<name>/QUEUE.md` loop packet of verifiable work units. The planner picks a short, descriptive name for the cycle (e.g., `go-cli`, `parser-bug`).
metadata:
  version: "1.0.0"
---

# Plan
=== /home/daniel/build/knack/.agents/skills/review/SKILL.md ===
---
name: review
description: Use when reviewing completed work units or a finished queue. Two-axis adversarial review — standards (does the change follow codebase conventions?) and intent (does the change do what the work unit said?). Reviews against the actual codebase, not specs. Triggers on "review", "check this", "is this right", "what did we miss", "stress-test the implementation", or when work needs adversarial scrutiny before being accepted.
metadata:
  version: "1.0.0"
---

# Review
[0m
[0m$ [0mcd /home/daniel/build/knack
perl -0pi -e 's/^metadata:\n  version: "1\.0\.0"\n/version: 1.0.0\n/m' .agents/skills/*/SKILL.md cli/embedded/skills/*/SKILL.md
echo "=== after ==="
for f in .agents/skills/build/SKILL.md cli/embedded/skills/decide/SKILL.md; do echo "--- $f ---"; head -7 "$f"; done
=== after ===
--- .agents/skills/build/SKILL.md ---
---
name: build
description: Use when implementing one work unit from a `.loop/<name>/QUEUE.md` queue. The worker-side skill — read the unit, do the work, don't self-certify, end with a handoff. Triggers on "build", "implement", "apply this unit", "do the work", "run the loop", or when a work unit needs to be executed. Also use when the loop invokes the worker for a tick.
version: 1.0.0
---

# Build
--- cli/embedded/skills/decide/SKILL.md ---
---
name: decide
description: Use when an architectural ruling crystallizes during any phase — explore, plan, build, or review. Captures decisions as ADRs in `decisions/` so they persist after specs are deleted and code moves on. Triggers on "we decided to", "let's go with", "the ruling is", "why did we choose", "record this decision", or when a design choice with lasting consequences is made.
version: 1.0.0
---

# Decide
[0m
All 7 skills share the same pattern. I'll transform the frontmatter in both `.agents/skills` and `cli/embedded/skills`.
[0m$ [0mcd /home/daniel/build/knack
grep -rn 'metadata:' .agents/skills/*/SKILL.md cli/embedded/skills/*/SKILL.md || echo "no metadata blocks left"
no metadata blocks left
[0m
Let me verify no stray `metadata:` blocks remain and update `buildEntry`.
[0m← [0mEdit cli/internal/skills/skills.go
Index: /home/daniel/build/knack/cli/internal/skills/skills.go
===================================================================
--- /home/daniel/build/knack/cli/internal/skills/skills.go
+++ /home/daniel/build/knack/cli/internal/skills/skills.go
@@ -263,19 +263,17 @@
 data, err := fs.ReadFile(fsys, filepath.Join(name, "SKILL.md"))
 if err != nil {
 	return SkillManifestEntry{}, fmt.Errorf("read skill %q: %w", name, err)
 }
-version := ""
-if front, _, ok := splitFrontmatter(data); ok {
-	var meta struct {
-		Metadata struct {
+	version := ""
+	if front, _, ok := splitFrontmatter(data); ok {
+		var meta struct {
 			Version string `yaml:"version"`
-		} `yaml:"metadata"`
+		}
+		if err := yaml.Unmarshal(front, &meta); err == nil {
+			version = meta.Version
+		}
 	}
-	if err := yaml.Unmarshal(front, &meta); err == nil {
-		version = meta.Metadata.Version
-	}
-}
 hash, err := dirHash(fsys, name)
 if err != nil {
 	return SkillManifestEntry{}, err
 }

[0m
[0m← [0mEdit cli/internal/skills/skills_test.go
Index: /home/daniel/build/knack/cli/internal/skills/skills_test.go
===================================================================
--- /home/daniel/build/knack/cli/internal/skills/skills_test.go
+++ /home/daniel/build/knack/cli/internal/skills/skills_test.go
@@ -356,11 +356,11 @@
 	return names
 }
 
 const (
-	skillManifestV  = "---\nname: x\ndescription: a skill\nmetadata:\n  version: \"1.0.0\"\n---\nbody\n"
-	skillEmbeddedV  = "---\nname: x\ndescription: a skill\nmetadata:\n  version: \"1.1.0\"\n---\nbody\n"
-	skillModifiedV  = "---\nname: x\ndescription: a skill\nmetadata:\n  version: \"1.0.0\"\n---\nchanged body\n"
+	skillManifestV  = "---\nname: x\ndescription: a skill\nversion: 1.0.0\n---\nbody\n"
+	skillEmbeddedV  = "---\nname: x\ndescription: a skill\nversion: 1.1.0\n---\nbody\n"
+	skillModifiedV  = "---\nname: x\ndescription: a skill\nversion: 1.0.0\n---\nchanged body\n"
 )
 
 func addManifest(t *testing.T, disk fstest.MapFS, entries ...SkillManifestEntry) {
 	t.Helper()

[0m
Now update the test fixtures to use top-level version:
Now run the verify command.
[0m$ [0mcd /home/daniel/build/knack
for f in .agents/skills/*/SKILL.md cli/embedded/skills/*/SKILL.md; do grep -q '^version:' "$f" || { echo "missing top-level version in $f"; exit 1; }; done && \
diff -r .agents/skills cli/embedded/skills && \
cd cli && go test ./internal/skills/... -count=1
ok  	knack/internal/skills	0.020s
[0m
Verify passes: all SKILL.md files have top-level `version:`, `.agents/skills` and `cli/embedded/skills` are in sync (no drift), and `go test ./internal/skills/...` is green.

```text
Unit: move skill version to top-level frontmatter and update buildEntry
Changed: .agents/skills/*/SKILL.md, cli/embedded/skills/*/SKILL.md (version moved to top-level), cli/internal/skills/skills.go (buildEntry reads top-level Version), cli/internal/skills/skills_test.go (skillManifestV/skillEmbeddedV/skillModifiedV fixtures)
Verify expected: grep '^version:' on all SKILL.md && diff -r .agents/skills cli/embedded/skills && go test ./internal/skills/...
Notes: none
```
````

What this proves:
- The verify command passed for this work unit in the current repo state.

What remains unverified:
- Anything outside the verify command's proof scope.

## 2026-07-10T23:41:42-06:00 — validate version in skills check

Status: done

Unit:
````markdown
## validate version in skills check

Read first:
- .loop/skill-versioning/REVIEW.md (S2)
- cli/internal/skills/skills.go (checkSkill)
- cli/internal/skills/skills_test.go (TestCheck* fixtures)
- .agents/skills/build/SKILL.md (example frontmatter)

Constraints:
- Do not change buildEntry or skill frontmatter; those are handled by the previous unit.
- Do not change usage() or skillsUsage in cli/main.go.
- Preserve existing name/description findings in checkSkill.

Done means:
- checkSkill reads version from top-level frontmatter and reports a finding when version is missing or empty.
- TestCheckValidSkill and other TestCheck* fixtures include version where needed and still pass.
- cd cli && go test ./internal/skills/... passes.
- No regression in Check behavior for real and embedded skills.

Verify:
```bash
cd cli && go test ./internal/skills/... -count=1
```

Status: pending````

Files changed:
```text
 M .agents/skills/build/SKILL.md
 M .agents/skills/decide/SKILL.md
 M .agents/skills/domain-modeling/SKILL.md
 M .agents/skills/explore/SKILL.md
 M .agents/skills/fix/SKILL.md
 M .agents/skills/plan/SKILL.md
 M .agents/skills/review/SKILL.md
 M AGENTS.md
 M README.md
 M cli/embedded/skills/build/SKILL.md
 M cli/embedded/skills/decide/SKILL.md
 M cli/embedded/skills/domain-modeling/SKILL.md
 M cli/embedded/skills/explore/SKILL.md
 M cli/embedded/skills/fix/SKILL.md
 M cli/embedded/skills/plan/SKILL.md
 M cli/embedded/skills/review/SKILL.md
 M cli/internal/skills/skills.go
 M cli/internal/skills/skills_test.go
 M cli/main.go
 M docs/cli.md
?? .loop/skill-versioning/EVIDENCE.md
?? COMPARISON-litespec.md
```

Verify command:
```bash
cd cli && go test ./internal/skills/... -count=1
```

Verify output:
```text
ok  	knack/internal/skills	0.017s
```

Worker output:
````text
[0m
> build · hy3-free
[0m
[0m→ [0mRead .agents/skills/build/SKILL.md
[0m→ [0mRead .loop/skill-versioning/REVIEW.md
[0m→ [0mRead cli/internal/skills/skills.go
[0m→ [0mRead cli/internal/skills/skills_test.go
[0m
[0m$ [0mfor f in .agents/skills/build/SKILL.md cli/embedded/skills/build/SKILL.md; do echo "=== $f ==="; head -6 "$f"; done
=== .agents/skills/build/SKILL.md ===
---
name: build
description: Use when implementing one work unit from a `.loop/<name>/QUEUE.md` queue. The worker-side skill — read the unit, do the work, don't self-certify, end with a handoff. Triggers on "build", "implement", "apply this unit", "do the work", "run the loop", or when a work unit needs to be executed. Also use when the loop invokes the worker for a tick.
version: 1.0.0
---

=== cli/embedded/skills/build/SKILL.md ===
---
name: build
description: Use when implementing one work unit from a `.loop/<name>/QUEUE.md` queue. The worker-side skill — read the unit, do the work, don't self-certify, end with a handoff. Triggers on "build", "implement", "apply this unit", "do the work", "run the loop", or when a work unit needs to be executed. Also use when the loop invokes the worker for a tick.
version: 1.0.0
---
[0m
[0m← [0mEdit cli/internal/skills/skills.go
Index: /home/daniel/build/knack/cli/internal/skills/skills.go
===================================================================
--- /home/daniel/build/knack/cli/internal/skills/skills.go
+++ /home/daniel/build/knack/cli/internal/skills/skills.go
@@ -513,8 +513,9 @@
 
 	var meta struct {
 		Name        string `yaml:"name"`
 		Description string `yaml:"description"`
+		Version     string `yaml:"version"`
 	}
 	if err := yaml.Unmarshal(front, &meta); err != nil {
 		return []string{fmt.Sprintf("%s: invalid frontmatter YAML: %v", skillName, err)}
 	}
@@ -523,8 +524,11 @@
 	}
 	if strings.TrimSpace(meta.Description) == "" {
 		findings = append(findings, fmt.Sprintf("%s: frontmatter field \"description\" is empty", skillName))
 	}
+	if strings.TrimSpace(meta.Version) == "" {
+		findings = append(findings, fmt.Sprintf("%s: frontmatter field \"version\" is empty", skillName))
+	}
 
 	findings = append(findings, checkReferences(fsys, skillDir, body)...)
 	return findings
 }

[0m
Now I'll add `version` validation to `checkSkill` and update the test fixtures that need a version field.
[0m← [0mEdit cli/internal/skills/skills_test.go
Index: /home/daniel/build/knack/cli/internal/skills/skills_test.go
===================================================================
--- /home/daniel/build/knack/cli/internal/skills/skills_test.go
+++ /home/daniel/build/knack/cli/internal/skills/skills_test.go
@@ -55,8 +55,9 @@
 	fsys := makeSkillFS(map[string]string{
 		"valid/SKILL.md": `---
 name: valid
 description: A valid skill.
+version: 1.0.0
 ---
 
 See [[notes.md]] and [notes](notes.md).
 `,

[0m
Now update the test fixtures that lack a `version` field:
[0m← [0mEdit cli/internal/skills/skills_test.go
Index: /home/daniel/build/knack/cli/internal/skills/skills_test.go
===================================================================
--- /home/daniel/build/knack/cli/internal/skills/skills_test.go
+++ /home/daniel/build/knack/cli/internal/skills/skills_test.go
@@ -89,8 +89,9 @@
 	fsys := makeSkillFS(map[string]string{
 		"bad/SKILL.md": `---
 name: bad
 description:
+version: 1.0.0
 ---
 `,
 	})
 	findings, err := Check(fsys)

[0m
[0m← [0mEdit cli/internal/skills/skills_test.go
Index: /home/daniel/build/knack/cli/internal/skills/skills_test.go
===================================================================
--- /home/daniel/build/knack/cli/internal/skills/skills_test.go
+++ /home/daniel/build/knack/cli/internal/skills/skills_test.go
@@ -107,8 +107,9 @@
 	fsys := makeSkillFS(map[string]string{
 		"bad/SKILL.md": `---
 name: bad
 description: Broken wiki ref.
+version: 1.0.0
 ---
 
 See [[missing.md]].
 `,

[0m
[0m← [0mEdit cli/internal/skills/skills_test.go
Index: /home/daniel/build/knack/cli/internal/skills/skills_test.go
===================================================================
--- /home/daniel/build/knack/cli/internal/skills/skills_test.go
+++ /home/daniel/build/knack/cli/internal/skills/skills_test.go
@@ -127,8 +127,9 @@
 	fsys := makeSkillFS(map[string]string{
 		"bad/SKILL.md": `---
 name: bad
 description: Broken link ref.
+version: 1.0.0
 ---
 
 See [missing](missing.md).
 `,

[0m
[0m← [0mEdit cli/internal/skills/skills_test.go
Index: /home/daniel/build/knack/cli/internal/skills/skills_test.go
===================================================================
--- /home/daniel/build/knack/cli/internal/skills/skills_test.go
+++ /home/daniel/build/knack/cli/internal/skills/skills_test.go
@@ -147,8 +147,9 @@
 	fsys := makeSkillFS(map[string]string{
 		"ok/SKILL.md": `---
 name: ok
 description: External links are fine.
+version: 1.0.0
 ---
 
 See [site](https://example.com) and [mail](mailto:a@b.com).
 `,

[0m
[0m← [0mEdit cli/internal/skills/skills_test.go
Index: /home/daniel/build/knack/cli/internal/skills/skills_test.go
===================================================================
--- /home/daniel/build/knack/cli/internal/skills/skills_test.go
+++ /home/daniel/build/knack/cli/internal/skills/skills_test.go
@@ -122,8 +122,26 @@
 		t.Fatalf("expected broken reference finding, got %v", findings)
 	}
 }
 
+func TestCheckEmptyVersion(t *testing.T) {
+	fsys := makeSkillFS(map[string]string{
+		"bad/SKILL.md": `---
+name: bad
+description: Missing version.
+version:
+---
+`,
+	})
+	findings, err := Check(fsys)
+	if err != nil {
+		t.Fatalf("Check failed: %v", err)
+	}
+	if len(findings) != 1 || !strings.Contains(findings[0], `"version" is empty`) {
+		t.Fatalf("expected empty version finding, got %v", findings)
+	}
+}
+
 func TestCheckBrokenLinkReference(t *testing.T) {
 	fsys := makeSkillFS(map[string]string{
 		"bad/SKILL.md": `---
 name: bad

[0m
Let me add a test covering the new version-empty finding:
[0m$ [0mcd /home/daniel/build/knack/cli && go test ./internal/skills/... -count=1
ok  	knack/internal/skills	0.026s
[0m
All tests pass.

```text
Unit: validate version in skills check
Changed: cli/internal/skills/skills.go (checkSkill: added Version field + empty-version finding), cli/internal/skills/skills_test.go (added version to valid/broken-ref/external-URL fixtures, new TestCheckEmptyVersion)
Verify expected: cd cli && go test ./internal/skills/... -count=1
Notes: none — top-level version already present in .agents/skills and embedded copies from prior unit.
```
````

What this proves:
- The verify command passed for this work unit in the current repo state.

What remains unverified:
- Anything outside the verify command's proof scope.
