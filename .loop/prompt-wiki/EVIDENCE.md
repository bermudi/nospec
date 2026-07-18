# Evidence: prompt-wiki

> This cycle was completed manually (not via the loop). EVIDENCE.md was
> reconstructed after the fact from QUEUE.md to keep the durable ledger
> complete. File-change lists are approximate.

## reviewer prompt uses evidence-first reasoning certificate and overcorrection guardrail

Status: done

Unit:
````markdown
## reviewer prompt uses evidence-first reasoning certificate and overcorrection guardrail

Read first:
- prompts/reviewer.md
- .agents/skills/review/SKILL.md
- wiki/threads/prompts-in-code-review.md
- wiki/concepts/semi-formal-reasoning.md
- wiki/concepts/overcorrection-bias.md

Constraints:
- Do not change the review skill's output format or the four top-level sections.
- Do not change loop.sh's parsing of the review artifact.
- Do not ask the reviewer to propose full patches or write code.

Done means:
- reviewer.md loads the review skill and preserves the compact terminal handoff.
- reviewer.md has an explicit anti-overcorrection guardrail (no patches, no implementation, no extended explanations).
- reviewer.md has an evidence-first reasoning certificate (premise, quoted file:line evidence, deviation, classification, confidence, fix direction).

Verify:
```bash
grep -q "Load and follow" prompts/reviewer.md && \
grep -q "Review: <cycle name>" prompts/reviewer.md && \
grep -qi "do not propose\|do not implement\|do not write patches" prompts/reviewer.md && \
grep -qi "premise" prompts/reviewer.md && \
grep -qi "evidence" prompts/reviewer.md
```

Status: done````

Files changed:
```text
M prompts/reviewer.md
```

---

## worker prompt is denser and uses leading words

Status: done

Unit:
````markdown
## worker prompt is denser and uses leading words

Read first:
- prompts/worker.md
- .agents/skills/build/SKILL.md
- wiki/concepts/leading-words.md
- wiki/concepts/system-prompt-effects.md
- wiki/concepts/context-engineering.md

Constraints:
- Do not remove the explicit skill-load instruction (ADR-0007).
- Do not change the output template shape.
- Keep the prompt under 50 lines.

Done means:
- worker.md has a reduced or grouped rule list and a de-emphasized role preamble.
- worker.md repeats a leading word/phrase that shapes the worker's posture (e.g., "one unit", "verify gate").
- worker.md still directs the worker to load the build skill and end with the compact terminal handoff.

Verify:
```bash
test "$(wc -l < prompts/worker.md)" -le 50 && \
grep -q "Load and follow" prompts/worker.md && \
grep -q "compact terminal handoff" prompts/worker.md && \
grep -qi "one unit" prompts/worker.md && \
grep -q "## " prompts/worker.md
```

Status: done````

Files changed:
```text
M prompts/worker.md
```

---

## fixer prompt strengthens triage and non-implementation guardrail

Status: done

Unit:
````markdown
## fixer prompt strengthens triage and non-implementation guardrail

Read first:
- prompts/fixer.md
- .agents/skills/fix/SKILL.md
- wiki/concepts/overcorrection-bias.md

Constraints:
- Do not change the fix skill's output format.
- Do not change the loop's queue format or parsing.

Done means:
- fixer.md has a dedicated triage section that explicitly classifies findings before queueing.
- fixer.md restates the non-implementation guardrail at the top of the triage section.
- fixer.md still loads the fix skill and preserves the compact terminal handoff.

Verify:
```bash
grep -q "Load and follow" prompts/fixer.md && \
grep -q "Fix: <cycle name>" prompts/fixer.md && \
grep -q "## Triage" prompts/fixer.md && \
grep -qi "do not implement" prompts/fixer.md && \
grep -qi "actionable\|trivial\|disputed\|deferred" prompts/fixer.md
```

Status: done````

Files changed:
```text
M prompts/fixer.md
```

---

## all prompt changes pass the test suite

Status: done

Unit:
````markdown
## all prompt changes pass the test suite

Read first:
- prompts/worker.md
- prompts/reviewer.md
- prompts/fixer.md
- tests/run.sh

Constraints:
- No code changes.
- No skill content changes.

Done means:
- The loop test suite passes.
- The CLI Go tests pass.

Verify:
```bash
./tests/run.sh && cd cli && go test ./...
```

Status: done````

Files changed:
```text
(none — verification only)
```
