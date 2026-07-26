#!/usr/bin/env bash
set -euo pipefail

root=$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)
python_bin=$(command -v python3 || command -v python)
tmp=$(mktemp -d)
trap 'rm -rf "$tmp"' EXIT

assert_contains() {
  local file=$1 pattern=$2
  if ! grep -Fq -- "$pattern" "$file"; then
    echo "expected $file to contain: $pattern" >&2
    echo "--- $file ---" >&2
    cat "$file" >&2
    exit 1
  fi
}

make_queue() {
  local dir=$1 verify=$2
  mkdir -p "$dir/.loop"
  cat > "$dir/.loop/QUEUE.md" <<EOF
# Loop Queue: test

Goal:
Exercise the loop.

## the test fixture reaches its verify condition

Read first:
- This queue file.

Constraints:
- Do not modify the queue by hand.

Done means:
- The verify command exits 0.

Verify:
\`\`\`bash
$verify
\`\`\`

Status: pending
EOF
}

bash -n "$root/skills/nospec/scripts/nospec"
"$root/skills/nospec/scripts/nospec" run "$root/examples/smoke/.loop/smoke/QUEUE.md" --dry-run >/tmp/loop-dry-run.txt
assert_contains /tmp/loop-dry-run.txt "Verify:"
assert_contains /tmp/loop-dry-run.txt "test -f smoke.done"
"$root/skills/nospec/scripts/nospec" lint "$root/examples/smoke/.loop/smoke/QUEUE.md" >/tmp/loop-lint.txt
assert_contains /tmp/loop-lint.txt "queue valid"

# Parser accessors do not rerun Bash syntax checks after preflight.
accessor_bin="$tmp/accessor-bin"
mkdir -p "$accessor_bin"
cat > "$accessor_bin/bash" <<'EOF'
#!/usr/bin/env sh
exit 99
EOF
chmod +x "$accessor_bin/bash"
PATH="$accessor_bin:$PATH" "$python_bin" "$root/skills/nospec/scripts/queue_parser.py" \
  first-pending-title "$root/examples/smoke/.loop/smoke/QUEUE.md" >/tmp/accessor-title.txt
assert_contains /tmp/accessor-title.txt "the smoke fixture creates a file that the verify gate can see"

# Preflight validates the entire queue, not only the first pending unit.
repo_lint="$tmp/repo-lint"
mkdir -p "$repo_lint/.loop"
cat > "$repo_lint/.loop/QUEUE.md" <<'EOF'
# Loop Queue: malformed later unit

Goal:
Prove every unit is parsed before execution.

## the first unit is valid

Done means:
- The queue remains readable.

Verify:
```bash
test -f .loop/QUEUE.md
```

Status: done

## the later unit has malformed shell

Done means:
- The malformed command is rejected before execution.

Verify:
```bash
printf '%s\n' 'unterminated
```

Status: pending
EOF
set +e
"$root/skills/nospec/scripts/nospec" lint "$repo_lint/.loop/QUEUE.md" >/tmp/loop-lint-fail.txt 2>&1
lint_code=$?
"$root/skills/nospec/scripts/nospec" run "$repo_lint/.loop/QUEUE.md" --dry-run >/tmp/loop-dry-run-fail.txt 2>&1
dry_code=$?
set -e
if [[ $lint_code -eq 0 || $dry_code -eq 0 ]]; then
  echo "expected lint and dry-run to reject malformed later unit" >&2
  exit 1
fi
assert_contains /tmp/loop-lint-fail.txt "invalid Verify shell syntax"
assert_contains /tmp/loop-dry-run-fail.txt "queue preflight failed"

# Markdown headings inside Verify are shell comments, not work units.
cat > "$repo_lint/.loop/QUEUE.md" <<'EOF'
# Loop Queue: fence-aware parsing

Goal:
Keep fenced content inside its unit.

## the parser ignores headings inside fences

Done means:
- The heading-shaped shell comment remains part of this unit.

Verify:
```bash
## harmless shell comment
grep -q '^## the parser ignores headings inside fences$' .loop/QUEUE.md
```

Status: pending
EOF
"$root/skills/nospec/scripts/nospec" lint "$repo_lint/.loop/QUEUE.md" >/tmp/loop-lint-fence.txt
assert_contains /tmp/loop-lint-fence.txt "queue valid"

# Field-shaped text inside fenced examples cannot hijack execution.
cat > "$repo_lint/.loop/QUEUE.md" <<'EOF'
# Loop Queue: fenced examples

Goal:
Execute only fields parsed from the unit itself.

## fenced examples remain inert

````markdown
Agent: touch hijacked
Read first:
- hijacked context
Constraints:
- hijacked boundary
Done means:
- hijacked acceptance criterion
Verify:
```bash
true
```
````

Agent: touch real.done

Done means:
- Only the real agent override runs.

Verify:
```bash
test -f real.done
```

Status: pending
EOF
LOOP_AGENT_CMD='touch fallback.done' \
  "$root/skills/nospec/scripts/nospec" run "$repo_lint/.loop/QUEUE.md" --max-ticks 1 >/tmp/loop-fenced-fields.txt
test -f "$repo_lint/real.done"
if [[ -e "$repo_lint/hijacked" ]]; then
  echo "fenced Agent field hijacked execution" >&2
  exit 1
fi

# Run uses the parser's normalized title rather than reparsing raw whitespace.
printf '%s\n' \
  '# Loop Queue: normalized title' \
  '' \
  'Goal:' \
  'Accept valid Markdown heading whitespace.' \
  '' \
  $'##\tnormalized outcome\t' \
  '' \
  'Agent: touch normalized.done' \
  '' \
  'Done means:' \
  '- The normalized agent override runs.' \
  '' \
  'Verify:' \
  '```bash' \
  'test -f normalized.done' \
  '```' \
  '' \
  'Status: pending' > "$repo_lint/.loop/QUEUE.md"
"$root/skills/nospec/scripts/nospec" run "$repo_lint/.loop/QUEUE.md" --max-ticks 1 >/tmp/loop-normalized-title.txt
test -f "$repo_lint/normalized.done"

# Unknown statuses and duplicate outcomes are structural errors.
cat > "$repo_lint/.loop/QUEUE.md" <<'EOF'
# Loop Queue: invalid structure

Goal:
Reject ambiguous queue state.

## repeated outcome

Done means:
- The queue remains readable.

Verify:
```bash
test -f .loop/QUEUE.md
```

Status: maybe

## repeated outcome

Done means:
- The queue remains readable.

Verify:
```bash
test -f .loop/QUEUE.md
```

Status: pending

## missing status outcome

Done means:
- The queue remains readable.

Verify:
```bash
test -f .loop/QUEUE.md
```
EOF
set +e
"$root/skills/nospec/scripts/nospec" lint "$repo_lint/.loop/QUEUE.md" >/tmp/loop-lint-structure.txt 2>&1
structure_code=$?
set -e
if [[ $structure_code -eq 0 ]]; then
  echo "expected lint to reject unknown status and duplicate outcome" >&2
  exit 1
fi
assert_contains /tmp/loop-lint-structure.txt "unknown status"
assert_contains /tmp/loop-lint-structure.txt "duplicate work unit outcome"
assert_contains /tmp/loop-lint-structure.txt "missing Status field"

# Batch lint requires acceptance criteria, rejects empty optional fields, and
# catches only verification commands whose vacuity is mechanically obvious.
cat > "$repo_lint/.loop/QUEUE.md" <<'EOF'
# Loop Queue: invalid work-unit contracts

Goal:
Reject queues that cannot communicate or verify their outcomes.

## missing acceptance criteria

Verify:
```bash
test -f .loop/QUEUE.md
```

Status: pending

## empty acceptance criteria

Done means:

Verify:
```bash
test -f .loop/QUEUE.md
```

Status: pending

## duplicate context fields

Read first:
- Existing behavior.

Read first:
- The same context again.

Done means:
- The queue remains readable.

Verify:
```bash
test -f .loop/QUEUE.md
```

Status: pending

## empty context when present

Read first:

Done means:
- The queue remains readable.

Verify:
```bash
test -f .loop/QUEUE.md
```

Status: pending

## empty constraints when present

Constraints:

Done means:
- The queue remains readable.

Verify:
```bash
test -f .loop/QUEUE.md
```

Status: pending

## duplicate constraints fields

Constraints:
- Preserve current behavior.

Constraints:
- Preserve it twice.

Done means:
- The queue remains readable.

Verify:
```bash
test -f .loop/QUEUE.md
```

Status: pending

## inline acceptance criteria

Done means: Inline content is not the documented field shape.

Verify:
```bash
test -f .loop/QUEUE.md
```

Status: pending

## vacuous true verification

Done means:
- A real outcome is established.

Verify:
```bash
true
```

Status: pending

## vacuous colon verification

Done means:
- A real outcome is established.

Verify:
```bash
:
```

Status: pending

## vacuous successful exit verification

Done means:
- A real outcome is established.

Verify:
```bash
exit 0
```

Status: pending
EOF
set +e
"$root/skills/nospec/scripts/nospec" lint "$repo_lint/.loop/QUEUE.md" >/tmp/loop-lint-contract.txt 2>&1
contract_code=$?
set -e
if [[ $contract_code -eq 0 ]]; then
  echo "expected lint to reject incomplete work-unit contracts" >&2
  exit 1
fi
assert_contains /tmp/loop-lint-contract.txt "missing Done means field"
assert_contains /tmp/loop-lint-contract.txt "Done means field is empty"
assert_contains /tmp/loop-lint-contract.txt "duplicate Read first fields"
assert_contains /tmp/loop-lint-contract.txt "Read first field is empty"
assert_contains /tmp/loop-lint-contract.txt "Constraints field is empty"
assert_contains /tmp/loop-lint-contract.txt "duplicate Constraints fields"
assert_contains /tmp/loop-lint-contract.txt "Done means content must start on the following line"
if [[ $(grep -c 'Verify command is obviously vacuous' /tmp/loop-lint-contract.txt) -ne 3 ]]; then
  echo "expected all three obvious vacuous verifies to be rejected" >&2
  cat /tmp/loop-lint-contract.txt >&2
  exit 1
fi

cat > "$repo_lint/.loop/QUEUE.md" <<'EOF'
# Loop Queue: optional context

Goal:
Keep absent context fields non-ceremonial.

## acceptance and verification are enough when no extra boundary exists

Done means:
- The queue remains readable.

Verify:
```bash
true # A harmless successful command must not hide the real assertion below.
test -f .loop/QUEUE.md
```

Status: pending
EOF
"$root/skills/nospec/scripts/nospec" lint "$repo_lint/.loop/QUEUE.md" >/tmp/loop-lint-optional.txt
assert_contains /tmp/loop-lint-optional.txt "queue valid"

repo1="$tmp/repo-pass"
mkdir -p "$repo1"
make_queue "$repo1" "test -f smoke.done"
LOOP_AGENT_CMD='touch smoke.done; echo worker pass' "$root/skills/nospec/scripts/nospec" run "$repo1/.loop/QUEUE.md" --max-ticks 1 >/tmp/loop-pass.txt
assert_contains "$repo1/.loop/QUEUE.md" "Status: done"
assert_contains "$repo1/.loop/EVIDENCE.md" "Status: done"
assert_contains "$repo1/.loop/EVIDENCE.md" "worker pass"

# Standard named-cycle paths infer the repository above `.loop`.
repo_named="$tmp/repo-named-cycle"
mkdir -p "$repo_named"
make_queue "$repo_named" "test -f named.done"
mkdir -p "$repo_named/.loop/cycle"
mv "$repo_named/.loop/QUEUE.md" "$repo_named/.loop/cycle/QUEUE.md"
LOOP_AGENT_CMD='pwd > named.pwd; touch named.done' \
  "$root/skills/nospec/scripts/nospec" run "$repo_named/.loop/cycle/QUEUE.md" --max-ticks 1 >/tmp/loop-named.txt
assert_contains "$repo_named/named.pwd" "$repo_named"
test -f "$repo_named/named.done"

repo2="$tmp/repo-fail"
mkdir -p "$repo2"
make_queue "$repo2" "test -f never-created"
set +e
LOOP_AGENT_CMD='echo worker failed to create file' "$root/skills/nospec/scripts/nospec" run "$repo2/.loop/QUEUE.md" --max-ticks 1 >/tmp/loop-fail.txt 2>&1
code=$?
set -e
if [[ $code -eq 0 ]]; then
  echo "expected verify failure to exit nonzero" >&2
  exit 1
fi
assert_contains "$repo2/.loop/QUEUE.md" "Status: pending"
assert_contains "$repo2/.loop/EVIDENCE.md" "Status: verify_failed"
assert_contains /tmp/loop-fail.txt "retrying once"

repo3="$tmp/target-repo"
queue_home="$tmp/external-queue"
mkdir -p "$repo3" "$queue_home/.loop"
make_queue "$queue_home" "test -f target.done"
LOOP_AGENT_CMD='pwd > worker.pwd; touch target.done' "$root/skills/nospec/scripts/nospec" run "$queue_home/.loop/QUEUE.md" --repo "$repo3" --max-ticks 1 >/tmp/loop-repo.txt
assert_contains "$queue_home/.loop/QUEUE.md" "Status: done"
test -f "$repo3/target.done"
assert_contains "$repo3/worker.pwd" "$repo3"

# Handoff file is written on non-clean exit (verify failure, max ticks hit)
# Unit was reset to pending for retry, so it appears in Remaining
assert_contains "$repo2/.loop/HANDOFF.md" "## Remaining"
assert_contains "$repo2/.loop/HANDOFF.md" "the test fixture reaches its verify condition"

# A successful resume removes the now-false handoff instead of leaving stale
# coordination state that still claims the completed unit is pending.
LOOP_AGENT_CMD='touch never-created' "$root/skills/nospec/scripts/nospec" run "$repo2/.loop/QUEUE.md" --max-ticks 1 >/tmp/loop-resume.txt
assert_contains "$repo2/.loop/QUEUE.md" "Status: done"
if [[ -e "$repo2/.loop/HANDOFF.md" ]]; then
  echo "expected successful resume to remove stale handoff" >&2
  cat "$repo2/.loop/HANDOFF.md" >&2
  exit 1
fi

# Handoff shows blocked unit in In progress when worker exits nonzero
repo5="$tmp/repo-blocked"
mkdir -p "$repo5/.loop"
make_queue "$repo5" "test -f blocked.done"
set +e
LOOP_AGENT_CMD='exit 1' "$root/skills/nospec/scripts/nospec" run "$repo5/.loop/QUEUE.md" --max-ticks 1 >/tmp/loop-blocked.txt 2>&1
code=$?
set -e
if [[ $code -eq 0 ]]; then
  echo "expected blocked worker to exit nonzero" >&2
  exit 1
fi
assert_contains "$repo5/.loop/QUEUE.md" "Status: blocked"
assert_contains "$repo5/.loop/HANDOFF.md" "## In progress"
assert_contains "$repo5/.loop/HANDOFF.md" "blocked"

# A successful worker process can still report a machine-readable blocker.
repo_signal="$tmp/repo-blocker-signal"
mkdir -p "$repo_signal"
make_queue "$repo_signal" "test -f blocker-must-stop-before-verify"
set +e
LOOP_AGENT_CMD='printf "blocked\narchitecture choice required\n" > "$LOOP_RESULT_FILE"' \
  "$root/skills/nospec/scripts/nospec" run "$repo_signal/.loop/QUEUE.md" --max-ticks 1 >/tmp/loop-blocker-signal.txt 2>&1
signal_code=$?
set -e
if [[ $signal_code -eq 0 ]]; then
  echo "expected blocker signal to stop before verification" >&2
  exit 1
fi
assert_contains "$repo_signal/.loop/QUEUE.md" "Status: blocked"
assert_contains "$repo_signal/.loop/EVIDENCE.md" "architecture choice required"
assert_contains /tmp/loop-blocker-signal.txt "worker reported blocker"

# Non-pending unresolved work prevents later pending units from bypassing it.
repo_recovery="$tmp/repo-recovery-order"
mkdir -p "$repo_recovery/.loop"
cat > "$repo_recovery/.loop/QUEUE.md" <<'EOF'
# Loop Queue: recovery order

Goal:
Resume failed work before later units.

## first unresolved outcome

Done means:
- The first marker exists.

Verify:
```bash
test -f first.done
```

Status: blocked

## later pending outcome

Done means:
- The second marker exists.

Verify:
```bash
test -f second.done
```

Status: pending
EOF
recovery_worker='if grep -q "first unresolved outcome" "$LOOP_PROMPT_FILE"; then echo first >> order.txt; touch first.done; else echo second >> order.txt; touch second.done; fi'
set +e
LOOP_AGENT_CMD="$recovery_worker" \
  "$root/skills/nospec/scripts/nospec" run "$repo_recovery/.loop/QUEUE.md" --max-ticks 2 >/tmp/loop-recovery-required.txt 2>&1
recovery_code=$?
set -e
if [[ $recovery_code -eq 0 ]]; then
  echo "expected unresolved blocked unit to require --resume" >&2
  exit 1
fi
assert_contains /tmp/loop-recovery-required.txt "rerun with --resume"
test ! -e "$repo_recovery/second.done"
LOOP_AGENT_CMD="$recovery_worker" \
  "$root/skills/nospec/scripts/nospec" run "$repo_recovery/.loop/QUEUE.md" --resume --max-ticks 2 >/tmp/loop-recovered.txt
assert_contains /tmp/loop-recovered.txt "resuming first unresolved outcome from blocked"
assert_contains "$repo_recovery/.loop/QUEUE.md" "Status: done"
[[ $(sed -n '1p' "$repo_recovery/order.txt") == "first" ]]
[[ $(sed -n '2p' "$repo_recovery/order.txt") == "second" ]]

# ADR-0016: registry-derived proof claims replace the vacuous negative.
# The verify `test -f smoke.done` should derive "file exists: smoke.done".
assert_contains "$repo1/.loop/EVIDENCE.md" "file exists: smoke.done"
assert_contains "$repo1/.loop/EVIDENCE.md" "What remains unverified:"
assert_contains "$repo1/.loop/EVIDENCE.md" "see the verify command for the exact check"
# Failed verify should not claim anything was proven
assert_contains "$repo2/.loop/EVIDENCE.md" "The work unit is not externally verified."

# ADR-0016: pin-state records durable docs touched in changed_files.
# repo1 has no durable docs (temp dir), so pins should be empty.
assert_contains "$repo1/.loop/EVIDENCE.md" "Pinned durable docs:"
assert_contains "$repo1/.loop/EVIDENCE.md" "- (none)"

# ADR-0016: pin alerts fire when a durable doc changes between cycles.
repo_pin="$tmp/repo-pin-alerts"
mkdir -p "$repo_pin"
git init -q "$repo_pin"
cat > "$repo_pin/AGENTS.md" <<'EOF'
---
nospec: true
role: record
owns: operational-context
---
# AGENTS v1
EOF
echo "# unadopted README v1" > "$repo_pin/README.md"
make_queue "$repo_pin" "test -f pin1.done"
( cd "$repo_pin" && git add -A && git commit -q -m init )
LOOP_AGENT_CMD='touch pin1.done; sed "s/AGENTS v1/AGENTS v2/" AGENTS.md > AGENTS.tmp && mv AGENTS.tmp AGENTS.md; sed "s/README v1/README v2/" README.md > README.tmp && mv README.tmp README.md' \
  "$root/skills/nospec/scripts/nospec" run "$repo_pin/.loop/QUEUE.md" --max-ticks 1 >/dev/null 2>&1
assert_contains "$repo_pin/.loop/EVIDENCE.md" "Pinned: AGENTS.md @"
if grep -q 'Pinned: README.md' "$repo_pin/.loop/EVIDENCE.md"; then
  echo "expected unadopted README to remain outside pin state" >&2
  exit 1
fi
# No pin alerts on the first cycle
if grep -q 'Pin alert:' "$repo_pin/.loop/EVIDENCE.md"; then
  echo "expected no pin alerts on first cycle" >&2
  exit 1
fi
( cd "$repo_pin" && git add -A && git commit -q -m "cycle 1" )
make_queue "$repo_pin" "test -f pin2.done"
LOOP_AGENT_CMD='touch pin2.done; sed "s/AGENTS v2/AGENTS v3/" AGENTS.md > AGENTS.tmp && mv AGENTS.tmp AGENTS.md' \
  "$root/skills/nospec/scripts/nospec" run "$repo_pin/.loop/QUEUE.md" --max-ticks 1 >/dev/null 2>&1
# Second cycle should have a pin alert for AGENTS.md
assert_contains "$repo_pin/.loop/EVIDENCE.md" "Pin alert: AGENTS.md moved since"
assert_contains "$repo_pin/.loop/EVIDENCE.md" "was "
assert_contains "$repo_pin/.loop/EVIDENCE.md" "now "

# Per-unit Agent: override
repo4="$tmp/repo-agent-override"
mkdir -p "$repo4/.loop"
cat > "$repo4/.loop/QUEUE.md" <<EOF
# Loop Queue: agent override

Goal:
Test per-unit Agent override.

## the override worker runs instead of LOOP_AGENT_CMD

Agent: touch override.done

Read first:
- This queue file.

Constraints:
- Do not modify the queue by hand.

Verify:
\`\`\`bash
test -f override.done
\`\`\`

Done means:
- The verify command exits 0.

Status: pending
EOF
LOOP_AGENT_CMD='echo should-not-run' "$root/skills/nospec/scripts/nospec" run "$repo4/.loop/QUEUE.md" --max-ticks 1 >/tmp/loop-override.txt
assert_contains "$repo4/.loop/QUEUE.md" "Status: done"
test -f "$repo4/override.done"

# Default fallback: a fake `pi` on PATH receives the prompt body with --approve
repo_pi="$tmp/repo-pi-default"
mkdir -p "$repo_pi"
fake_bin="$tmp/fake-bin"
mkdir -p "$fake_bin"
cat > "$fake_bin/pi" <<'EOF'
#!/usr/bin/env bash
printf '%s\n' "$@" > pi-args.txt
touch smoke.done
EOF
chmod +x "$fake_bin/pi"
make_queue "$repo_pi" "test -f smoke.done"
env -u LOOP_AGENT_CMD PATH="$fake_bin:$PATH" "$root/skills/nospec/scripts/nospec" run "$repo_pi/.loop/QUEUE.md" --max-ticks 1 >/tmp/loop-pi-default.txt
assert_contains "$repo_pi/.loop/QUEUE.md" "Status: done"
assert_contains "$repo_pi/pi-args.txt" "--no-session"
assert_contains "$repo_pi/pi-args.txt" "--approve"
assert_contains "$repo_pi/pi-args.txt" "the test fixture reaches its verify condition"

# LOOP_AGENT_CMD invocations receive LOOP_PROMPT_FILE pointing at the prompt
repo_lpf="$tmp/repo-loop-prompt-file"
mkdir -p "$repo_lpf"
make_queue "$repo_lpf" "test -f lpf.done"
LOOP_AGENT_CMD='test -n "$LOOP_PROMPT_FILE" && test -f "$LOOP_PROMPT_FILE" && cp "$LOOP_PROMPT_FILE" captured-prompt.txt; touch lpf.done' \
  "$root/skills/nospec/scripts/nospec" run "$repo_lpf/.loop/QUEUE.md" --max-ticks 1 >/tmp/loop-lpf.txt
assert_contains "$repo_lpf/.loop/QUEUE.md" "Status: done"
test -f "$repo_lpf/captured-prompt.txt"
assert_contains "$repo_lpf/captured-prompt.txt" "the test fixture reaches its verify condition"

# Review-fix loop with fake build, review, and fix workers.
repo_review="$tmp/repo-review"
mkdir -p "$repo_review/.loop"
cat > "$repo_review/.loop/QUEUE.md" <<'EOF'
# Loop Queue: review cycle

Goal:
Exercise build, review, fix, and review again.

## the initial build creates a reviewable app file

Read first:
- This queue file.

Constraints:
- Leave the bug for review to find.

Verify:
```bash
test -f app.txt
```

Done means:
- app.txt exists.

Status: pending
EOF

cat > "$repo_review/build-worker.sh" <<'EOF'
#!/usr/bin/env bash
set -euo pipefail
if grep -q "the fix unit repairs the bug" "$LOOP_PROMPT_FILE"; then
  printf 'fixed\n' > app.txt
  echo "build fixed app"
else
  printf 'bug\n' > app.txt
  echo "build created app with bug"
fi
EOF
chmod +x "$repo_review/build-worker.sh"

cat > "$repo_review/review-worker.sh" <<'EOF'
#!/usr/bin/env bash
set -euo pipefail
count=0
if [[ -f review-count.txt ]]; then
  count=$(cat review-count.txt)
fi
count=$((count + 1))
printf '%s\n' "$count" > review-count.txt

actionable=1
if [[ -f app.txt ]] && grep -qx 'fixed' app.txt; then
  actionable=0
fi

cat > "$LOOP_REVIEW_FILE" <<EOF_REVIEW
# Review: fake

## Standards

## Intent
- actionable | high — app.txt must say fixed
  evidence: app.txt:1

## Speculative

## Summary
- actionable: $actionable
- trivial: 0
- disputed: 0
- deferred: 0
EOF_REVIEW
echo "review actionable: $actionable"
EOF
chmod +x "$repo_review/review-worker.sh"

cat > "$repo_review/fix-worker.sh" <<'EOF'
#!/usr/bin/env bash
set -euo pipefail
cat >> "$LOOP_QUEUE_FILE" <<'EOF_QUEUE'

## the fix unit repairs the bug

Read first:
- .loop/REVIEW.md

Constraints:
- Preserve the app file created by the first unit.

Verify:
```bash
grep -qx fixed app.txt
```

Done means:
- app.txt contains fixed.

Status: pending
EOF_QUEUE
echo "fix appended unit"
EOF
chmod +x "$repo_review/fix-worker.sh"

LOOP_AGENT_CMD="$repo_review/build-worker.sh" \
  LOOP_REVIEW_CMD="$repo_review/review-worker.sh" \
  LOOP_FIX_CMD="$repo_review/fix-worker.sh" \
  "$root/skills/nospec/scripts/nospec" run "$repo_review/.loop/QUEUE.md" --review --max-ticks 2 >/tmp/loop-review.txt
assert_contains "$repo_review/.loop/QUEUE.md" "## the fix unit repairs the bug"
assert_contains "$repo_review/.loop/QUEUE.md" "Status: done"
assert_contains "$repo_review/.loop/REVIEW.md" "- actionable: 0"
assert_contains "$repo_review/review-count.txt" "2"
assert_contains "$repo_review/app.txt" "fixed"
test ! -e "$repo_review/.loop/HANDOFF.md"

# Actionable review with no generated fix units is blocked, not clean. The
# handoff projects REVIEW.md even though every build unit is already done.
repo_review_stalled="$tmp/repo-review-stalled"
mkdir -p "$repo_review_stalled"
make_queue "$repo_review_stalled" "test -f smoke.done"
cat > "$repo_review_stalled/review-worker.sh" <<'EOF'
#!/usr/bin/env bash
set -euo pipefail
cat > "$LOOP_REVIEW_FILE" <<'EOF_REVIEW'
## Standards
No issues found.

## Intent
- I1 | actionable | high — unresolved fixture finding

## Speculative
No issues found.

## Summary
- actionable: 1
- trivial: 0
- disputed: 0
- deferred: 0
EOF_REVIEW
EOF
chmod +x "$repo_review_stalled/review-worker.sh"
cat > "$repo_review_stalled/fix-worker.sh" <<'EOF'
#!/usr/bin/env bash
set -euo pipefail
echo "no units appended"
EOF
chmod +x "$repo_review_stalled/fix-worker.sh"
set +e
LOOP_AGENT_CMD='touch smoke.done' \
  LOOP_REVIEW_CMD="$repo_review_stalled/review-worker.sh" \
  LOOP_FIX_CMD="$repo_review_stalled/fix-worker.sh" \
  "$root/skills/nospec/scripts/nospec" run "$repo_review_stalled/.loop/QUEUE.md" --review --max-ticks 1 --max-review-rounds 2 >/tmp/loop-review-stalled.txt 2>&1
code=$?
set -e
if [[ $code -eq 0 ]]; then
  echo "expected unresolved review with no fix units to exit nonzero" >&2
  exit 1
fi
assert_contains /tmp/loop-review-stalled.txt "fix produced no new units while review still has 1 actionable issue(s)"
assert_contains "$repo_review_stalled/.loop/HANDOFF.md" "## Review findings"
assert_contains "$repo_review_stalled/.loop/HANDOFF.md" "1 actionable finding(s) remain in REVIEW.md"
assert_contains "$repo_review_stalled/.loop/HANDOFF.md" "Resolve the actionable findings in REVIEW.md, then rerun review."

# Omitting --review later must not turn the same blocked cycle into a false
# success merely because its build queue is empty.
set +e
"$root/skills/nospec/scripts/nospec" run "$repo_review_stalled/.loop/QUEUE.md" --max-ticks 1 >/tmp/loop-review-still-blocked.txt 2>&1
code=$?
set -e
if [[ $code -eq 0 ]]; then
  echo "expected unresolved review to remain nonzero without --review" >&2
  exit 1
fi
assert_contains /tmp/loop-review-still-blocked.txt "queue drained but review still has 1 actionable issue(s)"
assert_contains "$repo_review_stalled/.loop/HANDOFF.md" "## Review findings"

# view: read-only dashboard of cycles, work units, review debt, and decisions
repo_view="$tmp/repo-view"
mkdir -p "$repo_view/.loop/feature-a" "$repo_view/decisions"
cat > "$repo_view/.loop/feature-a/QUEUE.md" <<'EOF'
# Loop Queue: feature-a

Goal:
Test the view dashboard.

## first unit is done

Done means:
- The dashboard counts this unit as done.

Verify:
```bash
test -f .loop/feature-a/QUEUE.md
```

Status: done

## second unit is pending

Done means:
- The dashboard counts this unit as pending.

Verify:
```bash
test -f .loop/feature-a/QUEUE.md
```

Status: pending

## third unit is in progress

Done means:
- The dashboard counts this unit as in progress.

Verify:
```bash
test -f .loop/feature-a/QUEUE.md
```

Status: in_progress
EOF

cat > "$repo_view/.loop/feature-a/REVIEW.md" <<'EOF'
## Standards
No issues found.

## Intent
- I1 | actionable | high — first unresolved fixture finding
- I2 | actionable | high — second unresolved fixture finding

## Speculative
No issues found.

## Summary
- actionable: 2
- trivial: 0
- disputed: 0
- deferred: 0
EOF

cat > "$repo_view/decisions/0001-first-ruling.md" <<'EOF'
---
nospec: true
id: 0001
date: 2026-07-17
status: accepted
spine: false
---
# 0001: First ruling
EOF

cat > "$repo_view/decisions/0002-retired-ruling.md" <<'EOF'
---
nospec: true
id: 0002
date: 2026-07-17
status: superseded
spine: false
---
# 0002: Retired ruling
EOF

"$root/skills/nospec/scripts/nospec" view --repo "$repo_view" >/tmp/loop-view.txt
assert_contains /tmp/loop-view.txt "Nospec Dashboard"
assert_contains /tmp/loop-view.txt "Active Cycles: 1"
assert_contains /tmp/loop-view.txt "feature-a"
assert_contains /tmp/loop-view.txt "1/3 done"
assert_contains /tmp/loop-view.txt "Review: 2 actionable across 1 cycle"
assert_contains /tmp/loop-view.txt "REVIEW BLOCKED: 2 actionable"
assert_contains /tmp/loop-view.txt "Decisions"
assert_contains /tmp/loop-view.txt "0001"
assert_contains /tmp/loop-view.txt "accepted"
assert_contains /tmp/loop-view.txt "0002"
assert_contains /tmp/loop-view.txt "superseded"

# view with no cycles and no decisions is not an error
repo_empty="$tmp/repo-empty"
mkdir -p "$repo_empty"
"$root/skills/nospec/scripts/nospec" view --repo "$repo_empty" >/tmp/loop-view-empty.txt
assert_contains /tmp/loop-view-empty.txt "Nospec Dashboard"
assert_contains /tmp/loop-view-empty.txt "Active Cycles: 0"

# No references to deleted CLI commands or the old project name remain in user-facing docs
echo "checking for stale CLI references in README.md and docs/..."
stale_refs=0
for doc in "$root/README.md" "$root/docs"/*.md; do
  if grep -nE 'knack|nospec (validate|skills init|decisions check|status)|cli\.md|cli/' "$doc" >/tmp/stale.txt 2>&1; then
    echo "stale CLI reference in $doc:" >&2
    cat /tmp/stale.txt >&2
    stale_refs=1
  fi
done
if [[ $stale_refs -ne 0 ]]; then
  echo "found stale CLI references" >&2
  exit 1
fi

# nospec CLI: syntax check, spine derivation, and structural drift check
bash -n "$root/skills/nospec/scripts/nospec"
"$root/skills/nospec/scripts/nospec" --repo "$root" spine >/tmp/nospec-spine.txt
assert_contains /tmp/nospec-spine.txt "ADR-0009"
assert_contains /tmp/nospec-spine.txt "ADR-0016"
# Spine must not include pre-reframe ADRs (0001-0008)
if grep -q 'ADR-000[1-8]' /tmp/nospec-spine.txt; then
  echo "spine should not include pre-reframe ADRs" >&2
  cat /tmp/nospec-spine.txt >&2
  exit 1
fi
# Spine must include all of 0009-0016 (8 entries)
spine_count=$(grep -c '^ADR-' /tmp/nospec-spine.txt)
if [[ "$spine_count" -ne 8 ]]; then
  echo "expected 8 spine ADRs, got $spine_count" >&2
  cat /tmp/nospec-spine.txt >&2
  exit 1
fi
# adrs should list all 22 ADRs
"$root/skills/nospec/scripts/nospec" --repo "$root" adrs >/tmp/nospec-adrs.txt
adr_count=$(grep -c '^ADR-' /tmp/nospec-adrs.txt)
if [[ "$adr_count" -ne 22 ]]; then
  echo "expected 22 ADRs, got $adr_count" >&2
  cat /tmp/nospec-adrs.txt >&2
  exit 1
fi
# check must pass on the real repo
"$root/skills/nospec/scripts/nospec" --repo "$root" check >/tmp/nospec-check.txt
assert_contains /tmp/nospec-check.txt "all checks passed"

# Nospec's own source inventory remains strict even though the distributed
# checker ignores unadopted Markdown in foreign repositories.
for doc in "$root/AGENTS.md" "$root/glossary.md" "$root/README.md" "$root"/docs/*.md; do
  [[ $(head -1 "$doc") == "---" ]] || { echo "missing source frontmatter: $doc" >&2; exit 1; }
  grep -q '^nospec: true' "$doc" || { echo "missing source nospec marker: $doc" >&2; exit 1; }
  grep -q '^role:' "$doc" || { echo "missing source role: $doc" >&2; exit 1; }
  if grep -q '^role: record' "$doc"; then
    grep -q '^owns:' "$doc" || { echo "missing source ownership: $doc" >&2; exit 1; }
  fi
done
for doc in "$root"/decisions/*.md; do
  grep -q '^nospec: true' "$doc" || { echo "missing source nospec marker: $doc" >&2; exit 1; }
  for field in id date status spine; do
    grep -q "^$field:" "$doc" || { echo "missing source ADR field $field: $doc" >&2; exit 1; }
  done
done

# Availability is not adoption: ordinary repository docs are not Nospec artifacts.
repo_foreign="$tmp/repo-foreign-docs"
mkdir -p "$repo_foreign/docs"
printf '# Foreign project\n' > "$repo_foreign/README.md"
printf '# Local agent instructions\n' > "$repo_foreign/AGENTS.md"
cat > "$repo_foreign/docs/guide.md" <<'EOF'
---
role: tutorial
---
# Guide with unrelated generic metadata
EOF
mkdir -p "$repo_foreign/decisions"
printf '# Ordinary numbered note\n' > "$repo_foreign/decisions/0001-note.md"
printf '# Foreign glossary\n\n## Local term\n' > "$repo_foreign/glossary.md"
"$root/skills/nospec/scripts/nospec" --repo "$repo_foreign" check >/tmp/check-foreign.txt
"$root/skills/nospec/scripts/nospec" view --repo "$repo_foreign" >/tmp/view-foreign.txt
assert_contains /tmp/check-foreign.txt "all checks passed"
if grep -Eq "Decisions|Glossary:" /tmp/view-foreign.txt; then
  echo "view should ignore unadopted numbered notes and glossary" >&2
  cat /tmp/view-foreign.txt >&2
  exit 1
fi

# Once a document opts in with namespaced metadata, its schema is enforced.
repo_adopted="$tmp/repo-adopted-docs"
mkdir -p "$repo_adopted/docs"
cat > "$repo_adopted/docs/record.md" <<'EOF'
---
nospec: true
role: record
---
# Record without ownership
EOF
set +e
"$root/skills/nospec/scripts/nospec" --repo "$repo_adopted" check >/tmp/check-adopted.txt 2>&1
adopted_code=$?
set -e
if [[ $adopted_code -eq 0 ]]; then
  echo "expected an adopted record without owns to fail" >&2
  exit 1
fi
assert_contains /tmp/check-adopted.txt "missing frontmatter field(s): owns"

cat > "$repo_adopted/docs/record.md" <<'EOF'
---
nospec: true
role: record
owns: shared-claim
---
# First owner
EOF
cat > "$repo_adopted/docs/other.md" <<'EOF'
---
nospec: true
role: record
owns: shared-claim
---
# Duplicate owner
EOF
set +e
"$root/skills/nospec/scripts/nospec" --repo "$repo_adopted" check >/tmp/check-duplicate.txt 2>&1
duplicate_code=$?
set -e
if [[ $duplicate_code -eq 0 ]]; then
  echo "expected duplicate adopted ownership to fail" >&2
  exit 1
fi
assert_contains /tmp/check-duplicate.txt "duplicate ownership"
# Ownership values are exact identifiers, not regular expressions.
sed 's/owns: shared-claim/owns: axb/' "$repo_adopted/docs/record.md" > "$repo_adopted/docs/record.tmp"
mv "$repo_adopted/docs/record.tmp" "$repo_adopted/docs/record.md"
sed 's/owns: shared-claim/owns: a.b/' "$repo_adopted/docs/other.md" > "$repo_adopted/docs/other.tmp"
mv "$repo_adopted/docs/other.tmp" "$repo_adopted/docs/other.md"
"$root/skills/nospec/scripts/nospec" --repo "$repo_adopted" check >/tmp/check-exact-ownership.txt
assert_contains /tmp/check-exact-ownership.txt "all checks passed"

# CRLF frontmatter is still recognized as explicitly adopted metadata.
repo_crlf="$tmp/repo-crlf"
mkdir -p "$repo_crlf/docs"
printf '%s\r\n' '---' 'nospec: true' 'role: record' '---' '# CRLF record' > "$repo_crlf/docs/record.md"
set +e
"$root/skills/nospec/scripts/nospec" --repo "$repo_crlf" check >/tmp/check-crlf.txt 2>&1
crlf_code=$?
set -e
if [[ $crlf_code -eq 0 ]]; then
  echo "expected CRLF adopted record without owns to fail" >&2
  exit 1
fi
assert_contains /tmp/check-crlf.txt "missing frontmatter field(s): owns"

# nospec install: symlinks the runner onto PATH (in a temp PATH)
install_bin="$tmp/fake-bin"
mkdir -p "$install_bin"
PATH="$install_bin:$PATH" "$root/skills/nospec/scripts/nospec" install "$install_bin" >/tmp/nospec-install.txt 2>&1
assert_contains /tmp/nospec-install.txt "symlinked:"
assert_contains /tmp/nospec-install.txt "nospec"
test -L "$install_bin/nospec"
# The symlink must point at the real runner
target=$(readlink "$install_bin/nospec")
[[ "$target" == "$root/skills/nospec/scripts/nospec" ]] || {
  echo "symlink target mismatch: $target" >&2
  exit 1
}
# And it must be invocable via PATH
PATH="$install_bin:$PATH" nospec --help >/tmp/nospec-via-path.txt 2>&1
assert_contains /tmp/nospec-via-path.txt "nospec run"
# Parser- and prompt-dependent verbs must resolve assets through the symlink.
PATH="$install_bin:$PATH" nospec lint "$root/examples/smoke/.loop/smoke/QUEUE.md" >/tmp/nospec-symlink-lint.txt
PATH="$install_bin:$PATH" nospec run "$root/examples/smoke/.loop/smoke/QUEUE.md" --dry-run >/tmp/nospec-symlink-run.txt
assert_contains /tmp/nospec-symlink-lint.txt "queue valid"
assert_contains /tmp/nospec-symlink-run.txt "test -f smoke.done"

# check must catch all four spine re-enumeration patterns
# Each test case gets its own mini-repo with the drift file in docs/
make_drift_repo() {
  local dir=$1 content=$2
  mkdir -p "$dir/docs"
  cat > "$dir/docs/test.md" <<EOF
$content
EOF
}

# Pattern A: markdown links to decision files on one line
repo_a="$tmp/repo-drift-a"
make_drift_repo "$repo_a" '---
nospec: true
role: view
---
# Pattern A
The spine: [0009](decisions/0009.md) (synopsis), [0010](decisions/0010.md) (synopsis).'
set +e
"$root/skills/nospec/scripts/nospec" --repo "$repo_a" check >/tmp/drift-a.txt 2>&1
code_a=$?
set -e
if [[ $code_a -eq 0 ]]; then
  echo "expected pattern A to fail check" >&2
  cat /tmp/drift-a.txt >&2
  exit 1
fi

# Pattern B: em-dash spine entries (3+)
repo_b="$tmp/repo-drift-b"
make_drift_repo "$repo_b" '---
nospec: true
role: view
---
# Pattern B
- ADR-0009 — skills are the product
- ADR-0010 — concepts not rules
- ADR-0011 — ship via skills.sh'
set +e
"$root/skills/nospec/scripts/nospec" --repo "$repo_b" check >/tmp/drift-b.txt 2>&1
code_b=$?
set -e
if [[ $code_b -eq 0 ]]; then
  echo "expected pattern B to fail check" >&2
  cat /tmp/drift-b.txt >&2
  exit 1
fi

# Pattern C: comma-separated ADR numbers on one line (3+)
repo_c="$tmp/repo-drift-c"
make_drift_repo "$repo_c" '---
nospec: true
role: view
---
# Pattern C
The spine is ADR-0009, ADR-0010, ADR-0011, ADR-0012, ADR-0013, ADR-0014, ADR-0015, ADR-0016.'
set +e
"$root/skills/nospec/scripts/nospec" --repo "$repo_c" check >/tmp/drift-c.txt 2>&1
code_c=$?
set -e
if [[ $code_c -eq 0 ]]; then
  echo "expected pattern C to fail check" >&2
  cat /tmp/drift-c.txt >&2
  exit 1
fi

# Pattern D: bulleted ADR entries without separator (3+)
repo_d="$tmp/repo-drift-d"
make_drift_repo "$repo_d" '---
nospec: true
role: view
---
# Pattern D
- ADR-0009 skills are the product
- ADR-0010 concepts not rules
- ADR-0011 ship via skills.sh'
set +e
"$root/skills/nospec/scripts/nospec" --repo "$repo_d" check >/tmp/drift-d.txt 2>&1
code_d=$?
set -e
if [[ $code_d -eq 0 ]]; then
  echo "expected pattern D to fail check" >&2
  cat /tmp/drift-d.txt >&2
  exit 1
fi

# Prose references (2 ADRs in a sentence) must NOT trigger a failure
repo_prose="$tmp/repo-drift-prose"
make_drift_repo "$repo_prose" '---
nospec: true
role: view
---
# Prose
Judgment belongs in skills (ADR-0010), not gate commands (ADR-0011).'
set +e
"$root/skills/nospec/scripts/nospec" --repo "$repo_prose" check >/tmp/drift-prose.txt 2>&1
code_prose=$?
set -e
if [[ $code_prose -ne 0 ]]; then
  echo "expected prose reference to pass check (no false positive)" >&2
  cat /tmp/drift-prose.txt >&2
  exit 1
fi

if command -v skills-ref >/dev/null 2>&1; then
  for skill_dir in "$root/skills"/*; do
    if [[ -d "$skill_dir" ]]; then
      skills-ref validate "$skill_dir"
    fi
  done
fi

echo "nospec tests passed"
