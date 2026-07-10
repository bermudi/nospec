
## 2026-07-10T13:21:23-06:00 — the glossary check is safe and precise

Status: done

Unit:
````markdown
## the glossary check is safe and precise

Read first:
- `cli/internal/glossary/glossary.go`
- `cli/internal/glossary/glossary_test.go`
- `DESIGN.md` (the glossary check section)

Constraints:
- Don't change the CLI's public command interface or any command other than `glossary`.
- The `glossary` package must remain fully testable with `go test`.
- The undefined-term matcher must still catch real `[[...]]` references in markdown files.
- Stay inside the CLI; do not touch `loop.sh` or the skills.

Done means:
- `knack glossary check` exits 0 when `glossary.md` is absent.
- `knack glossary check` ignores `[[...]]` in shell scripts and Go source files, but still catches `[[...]]` in markdown files.
- `cd /home/daniel/build/knack/cli && go test ./...` still passes.

Verify:
```bash
cd /home/daniel/build/knack/cli && go test ./... && go build -o /tmp/knack . && cd /home/daniel/build/knack && /tmp/knack glossary check
```

Status: pending````

Files changed:
```text
 M .loop/glossary/QUEUE.md
 M cli/internal/glossary/glossary.go
 M cli/internal/glossary/glossary_test.go
?? .loop/glossary/EVIDENCE.md
```

Verify command:
```bash
cd /home/daniel/build/knack/cli && go test ./... && go build -o /tmp/knack . && cd /home/daniel/build/knack && /tmp/knack glossary check
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
[0m
> build · hy3-free
[0m
[0m→ [0mSkill "build"
[0m→ [0mRead cli/internal/glossary/glossary.go
[0m→ [0mRead cli/internal/glossary/glossary_test.go
[0m✱ [0mGrep "glossary"[90m in cli · 51 matches[0m
[0m→ [0mRead cli/main.go[90m [offset=256, limit=30][0m
[0m✱ [0mGrep "glossary|\[\["[90m in . · 54 matches[0m
[0m
[0m$ [0mls -la /home/daniel/build/knack/glossary.md 2>&1; echo "---"; cd /home/daniel/build/knack/cli && go build -o /tmp/knack-old . 2>&1 && cd /home/daniel/build/knack && /tmp/knack-old glossary check; echo "exit: $?"
ls: cannot access '/home/daniel/build/knack/glossary.md': No such file or directory
---
glossary check: read glossary: open glossary.md: no such file or directory
exit: 1
[0m
Now I understand the requirements. Let me make the changes:
[0m← [0mEdit cli/internal/glossary/glossary.go
Index: /home/daniel/build/knack/cli/internal/glossary/glossary.go
===================================================================
--- /home/daniel/build/knack/cli/internal/glossary/glossary.go
+++ /home/daniel/build/knack/cli/internal/glossary/glossary.go
@@ -1,9 +1,11 @@
 package glossary
 
 import (
+	"errors"
 	"fmt"
 	"io/fs"
+	"path"
 	"regexp"
 	"sort"
 	"strings"
 	"unicode"

[0m
[0m← [0mEdit cli/internal/glossary/glossary.go
Index: /home/daniel/build/knack/cli/internal/glossary/glossary.go
===================================================================
--- /home/daniel/build/knack/cli/internal/glossary/glossary.go
+++ /home/daniel/build/knack/cli/internal/glossary/glossary.go
@@ -55,11 +55,16 @@
 
 // Check reads glossary.md at glossaryPath from fsys and walks all other files
 // in fsys to report stale (defined but never used) and undefined (referenced
 // via [[...]] but not defined) terms.
+// If the glossary file is absent, Check reports no findings and no error: a
+// project without a glossary has nothing to check.
 func Check(fsys fs.FS, glossaryPath string) ([]Finding, error) {
 	data, err := fs.ReadFile(fsys, glossaryPath)
 	if err != nil {
+		if errors.Is(err, fs.ErrNotExist) {
+			return nil, nil
+		}
 		return nil, fmt.Errorf("read glossary: %w", err)
 	}
 	g := Parse(string(data))
 

[0m
[0m← [0mEdit cli/internal/glossary/glossary.go
Index: /home/daniel/build/knack/cli/internal/glossary/glossary.go
===================================================================
--- /home/daniel/build/knack/cli/internal/glossary/glossary.go
+++ /home/daniel/build/knack/cli/internal/glossary/glossary.go
@@ -172,8 +172,11 @@
 }
 if d.IsDir() || p == glossaryPath {
 	return nil
 }
+if !isMarkdown(p) {
+	return nil
+}
 data, err := fs.ReadFile(fsys, p)
 if err != nil {
 	return err
 }

[0m
Now restrict the `[[...]]` matcher to markdown files:
[0m← [0mEdit cli/internal/glossary/glossary.go
Index: /home/daniel/build/knack/cli/internal/glossary/glossary.go
===================================================================
--- /home/daniel/build/knack/cli/internal/glossary/glossary.go
+++ /home/daniel/build/knack/cli/internal/glossary/glossary.go
@@ -151,8 +151,19 @@
 func isWordChar(r rune) bool {
 	return unicode.IsLetter(r) || unicode.IsDigit(r) || r == '-' || r == '_' || r == '/'
 }
 
+// isMarkdown reports whether p is a markdown file. The [[...]] wiki-ref
+// convention is markdown-only, so undefined-term detection scans markdown and
+// ignores the [[...]] that appears in shell test syntax and Go source.
+func isMarkdown(p string) bool {
+	switch strings.ToLower(path.Ext(p)) {
+	case ".md", ".markdown":
+		return true
+	}
+	return false
+}
+
 func isBinary(data []byte) bool {
 	for _, b := range data {
 		if b == 0 {
 			return true

[0m
[0m← [0mEdit cli/internal/glossary/glossary_test.go
Index: /home/daniel/build/knack/cli/internal/glossary/glossary_test.go
===================================================================
--- /home/daniel/build/knack/cli/internal/glossary/glossary_test.go
+++ /home/daniel/build/knack/cli/internal/glossary/glossary_test.go
@@ -144,8 +144,72 @@
 		}
 	}
 }
 
+func TestGlossaryCheckAbsentGlossary(t *testing.T) {
+	fsys := makeGlossaryFS(map[string]string{
+		"docs.md": "Some content with [[SomeTerm]].\n",
+	})
+
+	findings, err := Check(fsys, "glossary.md")
+	if err != nil {
+		t.Fatalf("expected no error for absent glossary, got: %v", err)
+	}
+	if len(findings) != 0 {
+		t.Fatalf("expected no findings for absent glossary, got: %v", findings)
+	}
+}
+
+func TestGlossaryCheckIgnoresWikiRefsOutsideMarkdown(t *testing.T) {
+	fsys := makeGlossaryFS(map[string]string{
+		"glossary.md": lines(
+			"# Glossary",
+			"",
+			"## ActiveTerm",
+			"A term that is used.",
+		),
+		"docs.md":  "See [[ActiveTerm]] for details.\n",
+		"script.sh": "if [[ -f foo ]]; then echo hi; fi\n",
+		"code.go":  "package main\n// arr[[0]] style [[NotATerm]]\n",
+	})
+
+	findings, err := Check(fsys, "glossary.md")
+	if err != nil {
+		t.Fatalf("Check failed: %v", err)
+	}
+	for _, f := range findings {
+		if f.Kind == "undefined" {
+			t.Fatalf("expected no undefined findings from shell/go files, got: %v", findings)
+		}
+	}
+}
+
+func TestGlossaryCheckCatchesWikiRefsInMarkdown(t *testing.T) {
+	fsys := makeGlossaryFS(map[string]string{
+		"glossary.md": lines(
+			"# Glossary",
+			"",
+			"## ActiveTerm",
+			"A term that is used.",
+		),
+		"docs.md": "See [[ActiveTerm]] and [[MissingTerm]].\n",
+	})
+
+	findings, err := Check(fsys, "glossary.md")
+	if err != nil {
+		t.Fatalf("Check failed: %v", err)
+	}
+	var hasUndefined bool
+	for _, f := range findings {
+		if f.Kind == "undefined" && f.Term == "MissingTerm" {
+			hasUndefined = true
+		}
+	}
+	if !hasUndefined {
+		t.Fatalf("expected undefined MissingTerm from markdown, got: %v", findings)
+	}
+}
+
 func TestGlossaryParse(t *testing.T) {
 	g := Parse(lines(
 		"# Glossary",
 		"",

[0m
Now let me add tests for the new behavior:
[0m$ [0mcd /home/daniel/build/knack/cli && gofmt -w internal/glossary/glossary.go internal/glossary/glossary_test.go && go test ./... && go build -o /tmp/knack . && cd /home/daniel/build/knack && /tmp/knack glossary check; echo "exit: $?"
ok  	knack	0.003s
ok  	knack/internal/decisions	(cached)
ok  	knack/internal/glossary	0.002s
ok  	knack/internal/instructions	(cached)
ok  	knack/internal/queue	(cached)
ok  	knack/internal/skills	(cached)
ok  	knack/internal/status	(cached)
exit: 0
[0m
All tests pass and verify command exits 0.

```text
Unit: the glossary check is safe and precise
Changed: cli/internal/glossary/glossary.go (Check returns nil on absent glossary; findUndefined restricted to markdown via new isMarkdown helper), cli/internal/glossary/glossary_test.go (3 new tests: absent glossary, ignores .sh/.go wiki-refs, catches markdown wiki-refs)
Verify expected: cd /home/daniel/build/knack/cli && go test ./... && go build -o /tmp/knack . && cd /home/daniel/build/knack && /tmp/knack glossary check
Notes: Restricted [[...]] scanning to .md/.markdown files rather than blacklisting .sh/.go — the wiki-ref convention is markdown-only, so this is the precise fix and ignores shell test syntax and Go source cleanly.
```
````

What this proves:
- The verify command passed for this work unit in the current repo state.

What remains unverified:
- Anything outside the verify command's proof scope.

## 2026-07-10T13:23:54-06:00 — the project has a curated glossary and `knack glossary check` passes

Status: done

Unit:
````markdown
## the project has a curated glossary and `knack glossary check` passes

Read first:
- The `domain-modeling` skill
- `DESIGN.md` (the `[[...]]` references it uses)
- `AGENTS.md` and `README.md`
- The `cli/internal/glossary` package after unit 1

Constraints:
- The glossary stays small and curated; each entry is one or two sentences.
- Define a term only if it is used in the project or is a `[[...]]` target in `DESIGN.md`.
- Do not change the CLI code or `DESIGN.md` unless the glossary check proves it is necessary.
- Keep entries in alphabetical order and flat (no categories, no nesting).

Done means:
- `glossary.md` exists at the repo root with the project's domain vocabulary and all `[[...]]` references from `DESIGN.md`.
- `knack glossary check` reports no findings.
- `cd /home/daniel/build/knack/cli && go test ./...` still passes.
- The `./tests/run.sh` harness still passes (run it manually after the queue is done if you want the full loop-level check).

Verify:
```bash
cd /home/daniel/build/knack/cli && go test ./... && go build -o /tmp/knack . && cd /home/daniel/build/knack && /tmp/knack glossary check
```

Status: pending````

Files changed:
```text
 M .loop/glossary/EVIDENCE.md
 M .loop/glossary/QUEUE.md
 M cli/internal/glossary/glossary.go
 M cli/internal/glossary/glossary_test.go
?? glossary.md
```

Verify command:
```bash
cd /home/daniel/build/knack/cli && go test ./... && go build -o /tmp/knack . && cd /home/daniel/build/knack && /tmp/knack glossary check
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
[0m
> build · hy3-free
[0m
[0m→ [0mSkill "build"
[0m→ [0mRead DESIGN.md
[0m→ [0mRead AGENTS.md
[0m→ [0mRead README.md
Now let me examine the glossary CLI package and the domain-modeling skill.
[0m
[0m$ [0mls -R /home/daniel/build/knack/cli/internal/glossary 2>/dev/null; echo "---SKILL---"; ls /home/daniel/build/knack/.agents/skills/domain-modeling/
/home/daniel/build/knack/cli/internal/glossary:
glossary.go
glossary_test.go
---SKILL---
SKILL.md
[0m
[0m→ [0mRead cli/internal/glossary/glossary.go
[0m→ [0mRead .agents/skills/domain-modeling/SKILL.md
[0m
[0m$ [0mcd /home/daniel/build/knack && grep -rhoE '\[\[[^|\]]+\|?[^\]]*\]\]' --include=*.md . | sed -E 's/^\[\[//; s/\]\]$//; s/\|.*//' | sort -u
(no output)
[0m
Let me extract all `[[...]]` references from the markdown files to know which terms must be defined.
[0m$ [0mcd /home/daniel/build/knack && grep -rohE '\[\[[^]]+\]\]' --include=*.md . | sort -u | head -100
[[...]]
[[0]]
[[ActiveTerm]]
[[agent-loop]]
[[agent-skills]]
[[aiming-problem]]
[[backpressure]]
[[code-as-agent-harness]]
[[code-clarifies-spec]]
[[compounding-loops]]
[[context-files]]
[[decision-extraction]]
[[doc-rot]]
[[evolving-context]]
[[ -f foo ]]
[[harness-engineering]]
[[MissingTerm]]
[[NotATerm]]
[[plan-disposability]]
[[procedural-knowledge]]
[[ralph-loop]]
[[smart-zone-dumb-zone]]
[[SomeTerm]]
[[spec-code-triangle]]
[[spec-driven-development]]
[[steering-docs]]
[[tracer-bullets]]
[[verification-loop]]
[0m
[0m$ [0mcd /home/daniel/build/knack && grep -rlE '\[\[[^]]+\]\]' --include=*.md . 
./DESIGN.md
./.loop/glossary/QUEUE.md
./.loop/glossary/EVIDENCE.md
[0m
[0m$ [0mcd /home/daniel/build/knack && grep -oE '\[\[[^]]+\]\]' .loop/glossary/QUEUE.md .loop/glossary/EVIDENCE.md | sort -u
.loop/glossary/EVIDENCE.md:[[...]]
.loop/glossary/EVIDENCE.md:[[0]]
.loop/glossary/EVIDENCE.md:[[ActiveTerm]]
.loop/glossary/EVIDENCE.md:[[ -f foo ]]
.loop/glossary/EVIDENCE.md:[[MissingTerm]]
.loop/glossary/EVIDENCE.md:[[NotATerm]]
.loop/glossary/EVIDENCE.md:[[SomeTerm]]
.loop/glossary/QUEUE.md:[[...]]
[0m
[0m$ [0mcd /home/daniel/build/knack && cat .loop/glossary/EVIDENCE.md; echo "=== QUEUE ==="; cat .loop/glossary/QUEUE.md
## 2026-07-10T13:21:23-06:00 — the glossary check is safe and precise

Status: done

Unit:
````markdown
## the glossary check is safe and precise

Read first:
- `cli/internal/glossary/glossary.go`
- `cli/internal/glossary/glossary_test.go`
- `DESIGN.md` (the glossary check section)

Constraints:
- Don't change the CLI's public command interface or any command other than `glossary`.
- The `glossary` package must remain fully testable with `go test`.
- The undefined-term matcher must still catch real `[[...]]` references in markdown files.
- Stay inside the CLI; do not touch `loop.sh` or the skills.

Done means:
- `knack glossary check` exits 0 when `glossary.md` is absent.
- `knack glossary check` ignores `[[...]]` in shell scripts and Go source files, but still catches `[[...]]` in markdown files.
- `cd /home/daniel/build/knack/cli && go test ./...` still passes.

Verify:
```bash
cd /home/daniel/build/knack/cli && go test ./... && go build -o /tmp/knack . && cd /home/daniel/build/knack && /tmp/knack glossary check
```

Status: pending````

Files changed:
```text
 M .loop/glossary/QUEUE.md
 M cli/internal/glossary/glossary.go
 M cli/internal/glossary/glossary_test.go
?? .loop/glossary/EVIDENCE.md
```

Verify command:
```bash
cd /home/daniel/build/knack/cli && go test ./... && go build -o /tmp/knack . && cd /home/daniel/build/knack && /tmp/knack glossary check
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
[0m
> build · hy3-free
[0m
[0m→ [0mSkill "build"
[0m→ [0mRead cli/internal/glossary/glossary.go
[0m→ [0mRead cli/internal/glossary/glossary_test.go
[0m✱ [0mGrep "glossary"[90m in cli · 51 matches[0m
[0m→ [0mRead cli/main.go[90m [offset=256, limit=30][0m
[0m✱ [0mGrep "glossary|\[\["[90m in . · 54 matches[0m
[0m
[0m$ [0mls -la /home/daniel/build/knack/glossary.md 2>&1; echo "---"; cd /home/daniel/build/knack/cli && go build -o /tmp/knack-old . 2>&1 && cd /home/daniel/build/knack && /tmp/knack-old glossary check; echo "exit: $?"
ls: cannot access '/home/daniel/build/knack/glossary.md': No such file or directory
---
glossary check: read glossary: open glossary.md: no such file or directory
exit: 1
[0m
Now I understand the requirements. Let me make the changes:
[0m← [0mEdit cli/internal/glossary/glossary.go
Index: /home/daniel/build/knack/cli/internal/glossary/glossary.go
===================================================================
--- /home/daniel/build/knack/cli/internal/glossary/glossary.go
+++ /home/daniel/build/knack/cli/internal/glossary/glossary.go
@@ -1,9 +1,11 @@
 package glossary
 
 import (
+	"errors"
 	"fmt"
 	"io/fs"
+	"path"
 	"regexp"
 	"sort"
 	"strings"
 	"unicode"

[0m
[0m← [0mEdit cli/internal/glossary/glossary.go
Index: /home/daniel/build/knack/cli/internal/glossary/glossary.go
===================================================================
--- /home/daniel/build/knack/cli/internal/glossary/glossary.go
+++ /home/daniel/build/knack/cli/internal/glossary/glossary.go
@@ -55,11 +55,16 @@
 
 // Check reads glossary.md at glossaryPath from fsys and walks all other files
 // in fsys to report stale (defined but never used) and undefined (referenced
 // via [[...]] but not defined) terms.
+// If the glossary file is absent, Check reports no findings and no error: a
+// project without a glossary has nothing to check.
 func Check(fsys fs.FS, glossaryPath string) ([]Finding, error) {
 	data, err := fs.ReadFile(fsys, glossaryPath)
 	if err != nil {
+		if errors.Is(err, fs.ErrNotExist) {
+			return nil, nil
+		}
 		return nil, fmt.Errorf("read glossary: %w", err)
 	}
 	g := Parse(string(data))
 

[0m
[0m← [0mEdit cli/internal/glossary/glossary.go
Index: /home/daniel/build/knack/cli/internal/glossary/glossary.go
===================================================================
--- /home/daniel/build/knack/cli/internal/glossary/glossary.go
+++ /home/daniel/build/knack/cli/internal/glossary/glossary.go
@@ -172,8 +172,11 @@
 }
 if d.IsDir() || p == glossaryPath {
 	return nil
 }
+if !isMarkdown(p) {
+	return nil
+}
 data, err := fs.ReadFile(fsys, p)
 if err != nil {
 	return err
 }

[0m
Now restrict the `[[...]]` matcher to markdown files:
[0m← [0mEdit cli/internal/glossary/glossary.go
Index: /home/daniel/build/knack/cli/internal/glossary/glossary.go
===================================================================
--- /home/daniel/build/knack/cli/internal/glossary/glossary.go
+++ /home/daniel/build/knack/cli/internal/glossary/glossary.go
@@ -151,8 +151,19 @@
 func isWordChar(r rune) bool {
 	return unicode.IsLetter(r) || unicode.IsDigit(r) || r == '-' || r == '_' || r == '/'
 }
 
+// isMarkdown reports whether p is a markdown file. The [[...]] wiki-ref
+// convention is markdown-only, so undefined-term detection scans markdown and
+// ignores the [[...]] that appears in shell test syntax and Go source.
+func isMarkdown(p string) bool {
+	switch strings.ToLower(path.Ext(p)) {
+	case ".md", ".markdown":
+		return true
+	}
+	return false
+}
+
 func isBinary(data []byte) bool {
 	for _, b := range data {
 		if b == 0 {
 			return true

[0m
[0m← [0mEdit cli/internal/glossary/glossary_test.go
Index: /home/daniel/build/knack/cli/internal/glossary/glossary_test.go
===================================================================
--- /home/daniel/build/knack/cli/internal/glossary/glossary_test.go
+++ /home/daniel/build/knack/cli/internal/glossary/glossary_test.go
@@ -144,8 +144,72 @@
 		}
 	}
 }
 
+func TestGlossaryCheckAbsentGlossary(t *testing.T) {
+	fsys := makeGlossaryFS(map[string]string{
+		"docs.md": "Some content with [[SomeTerm]].\n",
+	})
+
+	findings, err := Check(fsys, "glossary.md")
+	if err != nil {
+		t.Fatalf("expected no error for absent glossary, got: %v", err)
+	}
+	if len(findings) != 0 {
+		t.Fatalf("expected no findings for absent glossary, got: %v", findings)
+	}
+}
+
+func TestGlossaryCheckIgnoresWikiRefsOutsideMarkdown(t *testing.T) {
+	fsys := makeGlossaryFS(map[string]string{
+		"glossary.md": lines(
+			"# Glossary",
+			"",
+			"## ActiveTerm",
+			"A term that is used.",
+		),
+		"docs.md":  "See [[ActiveTerm]] for details.\n",
+		"script.sh": "if [[ -f foo ]]; then echo hi; fi\n",
+		"code.go":  "package main\n// arr[[0]] style [[NotATerm]]\n",
+	})
+
+	findings, err := Check(fsys, "glossary.md")
+	if err != nil {
+		t.Fatalf("Check failed: %v", err)
+	}
+	for _, f := range findings {
+		if f.Kind == "undefined" {
+			t.Fatalf("expected no undefined findings from shell/go files, got: %v", findings)
+		}
+	}
+}
+
+func TestGlossaryCheckCatchesWikiRefsInMarkdown(t *testing.T) {
+	fsys := makeGlossaryFS(map[string]string{
+		"glossary.md": lines(
+			"# Glossary",
+			"",
+			"## ActiveTerm",
+			"A term that is used.",
+		),
+		"docs.md": "See [[ActiveTerm]] and [[MissingTerm]].\n",
+	})
+
+	findings, err := Check(fsys, "glossary.md")
+	if err != nil {
+		t.Fatalf("Check failed: %v", err)
+	}
+	var hasUndefined bool
+	for _, f := range findings {
+		if f.Kind == "undefined" && f.Term == "MissingTerm" {
+			hasUndefined = true
+		}
+	}
+	if !hasUndefined {
+		t.Fatalf("expected undefined MissingTerm from markdown, got: %v", findings)
+	}
+}
+
 func TestGlossaryParse(t *testing.T) {
 	g := Parse(lines(
 		"# Glossary",
 		"",

[0m
Now let me add tests for the new behavior:
[0m$ [0mcd /home/daniel/build/knack/cli && gofmt -w internal/glossary/glossary.go internal/glossary/glossary_test.go && go test ./... && go build -o /tmp/knack . && cd /home/daniel/build/knack && /tmp/knack glossary check; echo "exit: $?"
ok  	knack	0.003s
ok  	knack/internal/decisions	(cached)
ok  	knack/internal/glossary	0.002s
ok  	knack/internal/instructions	(cached)
ok  	knack/internal/queue	(cached)
ok  	knack/internal/skills	(cached)
ok  	knack/internal/status	(cached)
exit: 0
[0m
All tests pass and verify command exits 0.

```text
Unit: the glossary check is safe and precise
Changed: cli/internal/glossary/glossary.go (Check returns nil on absent glossary; findUndefined restricted to markdown via new isMarkdown helper), cli/internal/glossary/glossary_test.go (3 new tests: absent glossary, ignores .sh/.go wiki-refs, catches markdown wiki-refs)
Verify expected: cd /home/daniel/build/knack/cli && go test ./... && go build -o /tmp/knack . && cd /home/daniel/build/knack && /tmp/knack glossary check
Notes: Restricted [[...]] scanning to .md/.markdown files rather than blacklisting .sh/.go — the wiki-ref convention is markdown-only, so this is the precise fix and ignores shell test syntax and Go source cleanly.
```
````

What this proves:
- The verify command passed for this work unit in the current repo state.

What remains unverified:
- Anything outside the verify command's proof scope.
=== QUEUE ===
# Loop Queue: glossary

Goal:
Create a curated `glossary.md` for the project and fix `knack glossary check` so it is safe and passes on this repository. The CLI's `glossary check` must no longer error when `glossary.md` is missing, and it must not treat bash `[[...]]` syntax or Go test strings as glossary references. The glossary should define the project's domain vocabulary and the `[[...]]` references that `DESIGN.md` uses.

Stop condition:
`cd /home/daniel/build/knack/cli && go build -o /tmp/knack . && cd /home/daniel/build/knack && /tmp/knack glossary check`

## the glossary check is safe and precise

Read first:
- `cli/internal/glossary/glossary.go`
- `cli/internal/glossary/glossary_test.go`
- `DESIGN.md` (the glossary check section)

Constraints:
- Don't change the CLI's public command interface or any command other than `glossary`.
- The `glossary` package must remain fully testable with `go test`.
- The undefined-term matcher must still catch real `[[...]]` references in markdown files.
- Stay inside the CLI; do not touch `loop.sh` or the skills.

Done means:
- `knack glossary check` exits 0 when `glossary.md` is absent.
- `knack glossary check` ignores `[[...]]` in shell scripts and Go source files, but still catches `[[...]]` in markdown files.
- `cd /home/daniel/build/knack/cli && go test ./...` still passes.

Verify:
```bash
cd /home/daniel/build/knack/cli && go test ./... && go build -o /tmp/knack . && cd /home/daniel/build/knack && /tmp/knack glossary check
```

Status: done

## the project has a curated glossary and `knack glossary check` passes

Read first:
- The `domain-modeling` skill
- `DESIGN.md` (the `[[...]]` references it uses)
- `AGENTS.md` and `README.md`
- The `cli/internal/glossary` package after unit 1

Constraints:
- The glossary stays small and curated; each entry is one or two sentences.
- Define a term only if it is used in the project or is a `[[...]]` target in `DESIGN.md`.
- Do not change the CLI code or `DESIGN.md` unless the glossary check proves it is necessary.
- Keep entries in alphabetical order and flat (no categories, no nesting).

Done means:
- `glossary.md` exists at the repo root with the project's domain vocabulary and all `[[...]]` references from `DESIGN.md`.
- `knack glossary check` reports no findings.
- `cd /home/daniel/build/knack/cli && go test ./...` still passes.
- The `./tests/run.sh` harness still passes (run it manually after the queue is done if you want the full loop-level check).

Verify:
```bash
cd /home/daniel/build/knack/cli && go test ./... && go build -o /tmp/knack . && cd /home/daniel/build/knack && /tmp/knack glossary check
```

Status: in_progress
[0m
[0m$ [0mcd /home/daniel/build/knack && grep -oE '\[\[[^]|]+\]\]' DESIGN.md | sed -E 's/^\[\[//; s/\]\]$//' | sort -u
agent-loop
agent-skills
aiming-problem
backpressure
code-as-agent-harness
code-clarifies-spec
compounding-loops
context-files
decision-extraction
doc-rot
evolving-context
harness-engineering
plan-disposability
procedural-knowledge
ralph-loop
smart-zone-dumb-zone
spec-code-triangle
spec-driven-development
steering-docs
tracer-bullets
verification-loop
[0m
This is unit 2. Unit 1 already landed. Let me get the exact `[[...]]` targets from DESIGN.md (the real ones) and check how the decisions check handles `.loop` for comparison.
[0m$ [0mcd /home/daniel/build/knack && grep -rnE 'WalkDir|\.loop|SkipDir|fs.FS' cli/ --include=*.go | grep -v _test | head -40
cli/internal/skills/skills.go:20:func Init(fsys fs.FS, targetDir string) (wrote []string, skipped []string, err error) {
cli/internal/skills/skills.go:50:func Check(fsys fs.FS) ([]string, error) {
cli/internal/skills/skills.go:52:	err := fs.WalkDir(fsys, ".", func(p string, d fs.DirEntry, err error) error {
cli/internal/skills/skills.go:77:func copyDir(fsys fs.FS, src, dst string) error {
cli/internal/skills/skills.go:78:	return fs.WalkDir(fsys, src, func(path string, d fs.DirEntry, err error) error {
cli/internal/skills/skills.go:115:func checkSkill(fsys fs.FS, skillDir string, data []byte) []string {
cli/internal/skills/skills.go:163:func checkReferences(fsys fs.FS, skillDir string, body []byte) []string {
cli/internal/skills/skills.go:215:func refExists(fsys fs.FS, skillDir, target string) bool {
cli/internal/queue/queue.go:38:func ValidateFile(fsys fs.FS, path string) ([]Result, error) {
cli/internal/decisions/decisions.go:32:func List(fsys fs.FS, dir string) ([]ADR, error) {
cli/internal/decisions/decisions.go:97:func Show(fsys fs.FS, dir, number string) ([]byte, error) {
cli/internal/decisions/decisions.go:136:func Check(fsys fs.FS, decisionsDir string, loopFS fs.FS, loopDir string) ([]string, error) {
cli/internal/decisions/decisions.go:179:func collectLoopBodies(loopFS fs.FS, loopDir string) ([]string, error) {
cli/internal/decisions/decisions.go:181:	err := fs.WalkDir(loopFS, loopDir, func(p string, d fs.DirEntry, err error) error {
cli/internal/status/status.go:28:// Generate walks .loop/ for cycle subdirectories, reads their QUEUE.md and
cli/internal/status/status.go:30:func Generate(fsys fs.FS) (Report, error) {
cli/internal/status/status.go:33:	entries, err := fs.ReadDir(fsys, ".loop")
cli/internal/status/status.go:35:		// No .loop/ — return zero counts with ADRs.
cli/internal/status/status.go:47:		queuePath := path.Join(".loop", cycleDir, "QUEUE.md")
cli/internal/status/status.go:63:		evPath := path.Join(".loop", cycleDir, "EVIDENCE.md")
cli/internal/status/status.go:79:func countEvidenceFile(fsys fs.FS, path string) int {
cli/internal/glossary/glossary.go:61:func Check(fsys fs.FS, glossaryPath string) ([]Finding, error) {
cli/internal/glossary/glossary.go:92:func findStale(fsys fs.FS, glossaryPath string, g *Glossary) ([]string, error) {
cli/internal/glossary/glossary.go:94:	err := fs.WalkDir(fsys, ".", func(p string, d fs.DirEntry, err error) error {
cli/internal/glossary/glossary.go:177:func findUndefined(fsys fs.FS, glossaryPath string, g *Glossary) ([]string, error) {
cli/internal/glossary/glossary.go:180:	err := fs.WalkDir(fsys, ".", func(p string, d fs.DirEntry, err error) error {
cli/main.go:207:	findings, err := decisions.Check(os.DirFS("."), "decisions", os.DirFS("."), ".loop")
[0m
[0m← [0mWrite glossary.md
Wrote file successfully.
[0m
Let me create the glossary from DESIGN.md's wiki-ref targets, then run the verify to see what the check actually reports (especially regarding `.loop/` noise).
[0m$ [0mcd /home/daniel/build/knack/cli && go build -o /tmp/knack . && cd /home/daniel/build/knack && /tmp/knack glossary check; echo "exit: $?"
undefined:  -f foo 
undefined: ...
undefined: 0
undefined: ActiveTerm
undefined: MissingTerm
undefined: NotATerm
undefined: SomeTerm
exit: 1
[0m
[0m← [0mEdit cli/internal/glossary/glossary.go
Index: /home/daniel/build/knack/cli/internal/glossary/glossary.go
===================================================================
--- /home/daniel/build/knack/cli/internal/glossary/glossary.go
+++ /home/daniel/build/knack/cli/internal/glossary/glossary.go
@@ -94,11 +94,17 @@
 err := fs.WalkDir(fsys, ".", func(p string, d fs.DirEntry, err error) error {
 	if err != nil {
 		return err
 	}
-	if d.IsDir() || p == glossaryPath {
+	if d.IsDir() {
+		if p == ".loop" {
+			return fs.SkipDir
+		}
 		return nil
 	}
+	if p == glossaryPath {
+		return nil
+	}
 	data, err := fs.ReadFile(fsys, p)
 	if err != nil {
 		return err
 	}

[0m
The disposable `.loop/glossary/EVIDENCE.md` (and `QUEUE.md`) contain pasted worker output with test `[[...]]` references, which the check flags as undefined. These are disposable coordination-state files, not project prose — the check shouldn't scan `.loop/`. This is the case the constraint's escape hatch covers ("unless the glossary check proves it is necessary"). Let me add a `.loop` skip to the walks.
[0m← [0mEdit cli/internal/glossary/glossary.go
Index: /home/daniel/build/knack/cli/internal/glossary/glossary.go
===================================================================
--- /home/daniel/build/knack/cli/internal/glossary/glossary.go
+++ /home/daniel/build/knack/cli/internal/glossary/glossary.go
@@ -186,11 +186,17 @@
 err := fs.WalkDir(fsys, ".", func(p string, d fs.DirEntry, err error) error {
 	if err != nil {
 		return err
 	}
-	if d.IsDir() || p == glossaryPath {
+	if d.IsDir() {
+		if p == ".loop" {
+			return fs.SkipDir
+		}
 		return nil
 	}
+	if p == glossaryPath {
+		return nil
+	}
 	if !isMarkdown(p) {
 		return nil
 	}
 	data, err := fs.ReadFile(fsys, p)

[0m
[0m$ [0mcd /home/daniel/build/knack/cli && go test ./... && go build -o /tmp/knack . && cd /home/daniel/build/knack && /tmp/knack glossary check; echo "exit: $?"
ok  	knack	0.003s
ok  	knack/internal/decisions	(cached)
ok  	knack/internal/glossary	0.004s
ok  	knack/internal/instructions	(cached)
ok  	knack/internal/queue	(cached)
ok  	knack/internal/skills	(cached)
ok  	knack/internal/status	(cached)
exit: 0
[0m
[0m→ [0mRead cli/internal/glossary/glossary_test.go[90m [limit=60][0m
The check passes. Let me add a regression test for the `.loop` skip, matching the existing test style, since the project's review culture relies on tests to catch this class of bug.
[0m→ [0mRead cli/internal/glossary/glossary_test.go[90m [offset=205, limit=45][0m
[0m
[0m← [0mEdit cli/internal/glossary/glossary_test.go
Index: /home/daniel/build/knack/cli/internal/glossary/glossary_test.go
===================================================================
--- /home/daniel/build/knack/cli/internal/glossary/glossary_test.go
+++ /home/daniel/build/knack/cli/internal/glossary/glossary_test.go
@@ -208,8 +208,35 @@
 		t.Fatalf("expected undefined MissingTerm from markdown, got: %v", findings)
 	}
 }
 
+func TestGlossaryCheckIgnoresLoopDir(t *testing.T) {
+	fsys := makeGlossaryFS(map[string]string{
+		"glossary.md": lines(
+			"# Glossary",
+			"",
+			"## ActiveTerm",
+			"A term that is used in the project.",
+		),
+		".loop/glossary/EVIDENCE.md": lines(
+			"Worker pasted test code with [[ActiveTerm]] and [[MissingTerm]].",
+		),
+		".loop/glossary/QUEUE.md": lines(
+			"[[...]] placeholder reference.",
+		),
+	})
+
+	findings, err := Check(fsys, "glossary.md")
+	if err != nil {
+		t.Fatalf("Check failed: %v", err)
+	}
+	for _, f := range findings {
+		if f.Kind == "undefined" {
+			t.Fatalf("expected .loop references to be ignored, got: %v", findings)
+		}
+	}
+}
+
 func TestGlossaryParse(t *testing.T) {
 	g := Parse(lines(
 		"# Glossary",
 		"",

[0m
[0m$ [0mcd /home/daniel/build/knack/cli && go test ./... && go build -o /tmp/knack . && cd /home/daniel/build/knack && /tmp/knack glossary check; echo "exit: $?"
ok  	knack	(cached)
ok  	knack/internal/decisions	(cached)
ok  	knack/internal/glossary	0.003s
ok  	knack/internal/instructions	(cached)
ok  	knack/internal/queue	(cached)
ok  	knack/internal/skills	(cached)
ok  	knack/internal/status	(cached)
exit: 0
[0m
All green. Let me run the full harness as `Done means` requires it to still pass.
[0m$ [0mcd /home/daniel/build/knack && ./tests/run.sh 2>&1 | tail -20; echo "exit: ${PIPESTATUS[0]}"
Valid skill: /home/daniel/build/knack/.agents/skills/build
Valid skill: /home/daniel/build/knack/.agents/skills/decide
Valid skill: /home/daniel/build/knack/.agents/skills/domain-modeling
Valid skill: /home/daniel/build/knack/.agents/skills/explore
Valid skill: /home/daniel/build/knack/.agents/skills/fix
Valid skill: /home/daniel/build/knack/.agents/skills/plan
Valid skill: /home/daniel/build/knack/.agents/skills/review
knack tests passed
exit: 0
[0m
All checks pass. 

```text
Unit: the project has a curated glossary and `knack glossary check` passes
Changed: glossary.md (new, 21 DESIGN.md wiki-ref terms, alphabetical/flat); cli/internal/glossary/glossary.go (skip `.loop/` in both walks); cli/internal/glossary/glossary_test.go (+TestGlossaryCheckIgnoresLoopDir)
Verify expected: cd /home/daniel/build/knack/cli && go test ./... && go build -o /tmp/knack . && cd /home/daniel/build/knack && /tmp/knack glossary check
Notes: `.loop/glossary/EVIDENCE.md` contained pasted test `[[...]]` refs that falsely flagged undefined findings; skipping disposable `.loop/` in the walk was the necessary CLI change (constraint escape hatch). ./tests/run.sh also passes.
```
````

What this proves:
- The verify command passed for this work unit in the current repo state.

What remains unverified:
- Anything outside the verify command's proof scope.
