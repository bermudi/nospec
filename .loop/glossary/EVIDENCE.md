
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
