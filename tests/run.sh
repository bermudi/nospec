#!/usr/bin/env bash
set -euo pipefail

root=$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)
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

Stop condition:
\`$verify\` exits 0.

## the test fixture reaches its verify condition

Read first:
- This queue file.

Constraints:
- Do not modify the queue by hand.

Verify:
\`\`\`bash
$verify
\`\`\`

Done means:
- The verify command exits 0.

Status: pending
EOF
}

bash -n "$root/loop.sh"
"$root/loop.sh" run "$root/examples/smoke/.loop/smoke/QUEUE.md" --dry-run >/tmp/loop-dry-run.txt
assert_contains /tmp/loop-dry-run.txt "Verify:"
assert_contains /tmp/loop-dry-run.txt "test -f smoke.done"

repo1="$tmp/repo-pass"
mkdir -p "$repo1"
make_queue "$repo1" "test -f smoke.done"
LOOP_AGENT_CMD='touch smoke.done; echo worker pass' "$root/loop.sh" run "$repo1/.loop/QUEUE.md" --max-ticks 1 >/tmp/loop-pass.txt
assert_contains "$repo1/.loop/QUEUE.md" "Status: done"
assert_contains "$repo1/.loop/EVIDENCE.md" "Status: done"
assert_contains "$repo1/.loop/EVIDENCE.md" "worker pass"

repo2="$tmp/repo-fail"
mkdir -p "$repo2"
make_queue "$repo2" "test -f never-created"
set +e
LOOP_AGENT_CMD='echo worker failed to create file' "$root/loop.sh" run "$repo2/.loop/QUEUE.md" --max-ticks 1 >/tmp/loop-fail.txt 2>&1
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
LOOP_AGENT_CMD='pwd > worker.pwd; touch target.done' "$root/loop.sh" run "$queue_home/.loop/QUEUE.md" --repo "$repo3" --max-ticks 1 >/tmp/loop-repo.txt
assert_contains "$queue_home/.loop/QUEUE.md" "Status: done"
test -f "$repo3/target.done"
assert_contains "$repo3/worker.pwd" "$repo3"

# Handoff file is written on non-clean exit (verify failure, max ticks hit)
# Unit was reset to pending for retry, so it appears in Remaining
assert_contains "$repo2/.loop/HANDOFF.md" "## Remaining"
assert_contains "$repo2/.loop/HANDOFF.md" "the test fixture reaches its verify condition"

# Handoff shows blocked unit in In progress when worker exits nonzero
repo5="$tmp/repo-blocked"
mkdir -p "$repo5/.loop"
make_queue "$repo5" "test -f blocked.done"
set +e
LOOP_AGENT_CMD='exit 1' "$root/loop.sh" run "$repo5/.loop/QUEUE.md" --max-ticks 1 >/tmp/loop-blocked.txt 2>&1
code=$?
set -e
if [[ $code -eq 0 ]]; then
  echo "expected blocked worker to exit nonzero" >&2
  exit 1
fi
assert_contains "$repo5/.loop/QUEUE.md" "Status: blocked"
assert_contains "$repo5/.loop/HANDOFF.md" "## In progress"
assert_contains "$repo5/.loop/HANDOFF.md" "blocked"

# Per-unit Agent: override
repo4="$tmp/repo-agent-override"
mkdir -p "$repo4/.loop"
cat > "$repo4/.loop/QUEUE.md" <<EOF
# Loop Queue: agent override

Goal:
Test per-unit Agent override.

Stop condition:
\`test -f override.done\` exits 0.

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
LOOP_AGENT_CMD='echo should-not-run' "$root/loop.sh" run "$repo4/.loop/QUEUE.md" --max-ticks 1 >/tmp/loop-override.txt
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
env -u LOOP_AGENT_CMD PATH="$fake_bin:$PATH" "$root/loop.sh" run "$repo_pi/.loop/QUEUE.md" --max-ticks 1 >/tmp/loop-pi-default.txt
assert_contains "$repo_pi/.loop/QUEUE.md" "Status: done"
assert_contains "$repo_pi/pi-args.txt" "--no-session"
assert_contains "$repo_pi/pi-args.txt" "--approve"
assert_contains "$repo_pi/pi-args.txt" "the test fixture reaches its verify condition"

# LOOP_AGENT_CMD invocations receive LOOP_PROMPT_FILE pointing at the prompt
repo_lpf="$tmp/repo-loop-prompt-file"
mkdir -p "$repo_lpf"
make_queue "$repo_lpf" "test -f lpf.done"
LOOP_AGENT_CMD='test -n "$LOOP_PROMPT_FILE" && test -f "$LOOP_PROMPT_FILE" && cp "$LOOP_PROMPT_FILE" captured-prompt.txt; touch lpf.done' \
  "$root/loop.sh" run "$repo_lpf/.loop/QUEUE.md" --max-ticks 1 >/tmp/loop-lpf.txt
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

Stop condition:
The generated app is fixed and review reports no actionable issues.

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
  "$root/loop.sh" run "$repo_review/.loop/QUEUE.md" --review --max-ticks 2 >/tmp/loop-review.txt
assert_contains "$repo_review/.loop/QUEUE.md" "## the fix unit repairs the bug"
assert_contains "$repo_review/.loop/QUEUE.md" "Status: done"
assert_contains "$repo_review/.loop/REVIEW.md" "- actionable: 0"
assert_contains "$repo_review/review-count.txt" "2"
assert_contains "$repo_review/app.txt" "fixed"

# view: read-only dashboard of cycles, work units, and decisions
repo_view="$tmp/repo-view"
mkdir -p "$repo_view/.loop/feature-a" "$repo_view/decisions"
cat > "$repo_view/.loop/feature-a/QUEUE.md" <<'EOF'
# Loop Queue: feature-a

Goal:
Test the view dashboard.

Stop condition:
All units done.

## first unit is done

Verify:
```bash
true
```

Status: done

## second unit is pending

Verify:
```bash
true
```

Status: pending

## third unit is in progress

Verify:
```bash
true
```

Status: in_progress
EOF

cat > "$repo_view/decisions/0001-first-ruling.md" <<'EOF'
# 0001: First ruling

Date: 2026-07-17
Status: accepted

## Context
Test.
EOF

cat > "$repo_view/decisions/0002-proposed-ruling.md" <<'EOF'
# 0002: Proposed ruling

Date: 2026-07-17
Status: proposed

## Context
Test.
EOF

"$root/loop.sh" view --repo "$repo_view" >/tmp/loop-view.txt
assert_contains /tmp/loop-view.txt "Knack Dashboard"
assert_contains /tmp/loop-view.txt "Active Cycles: 1"
assert_contains /tmp/loop-view.txt "feature-a"
assert_contains /tmp/loop-view.txt "1/3 done"
assert_contains /tmp/loop-view.txt "Decisions"
assert_contains /tmp/loop-view.txt "0001"
assert_contains /tmp/loop-view.txt "accepted"
assert_contains /tmp/loop-view.txt "0002"
assert_contains /tmp/loop-view.txt "proposed"

# view with no cycles and no decisions is not an error
repo_empty="$tmp/repo-empty"
mkdir -p "$repo_empty"
"$root/loop.sh" view --repo "$repo_empty" >/tmp/loop-view-empty.txt
assert_contains /tmp/loop-view-empty.txt "Knack Dashboard"
assert_contains /tmp/loop-view-empty.txt "Active Cycles: 0"

if command -v skills-ref >/dev/null 2>&1; then
  for skill_dir in "$root/skills"/*; do
    if [[ -d "$skill_dir" ]]; then
      skills-ref validate "$skill_dir"
    fi
  done
fi

echo "knack tests passed"
