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

bash -n "$root/skills/nospec/scripts/nospec"
"$root/skills/nospec/scripts/nospec" run "$root/examples/smoke/.loop/smoke/QUEUE.md" --dry-run >/tmp/loop-dry-run.txt
assert_contains /tmp/loop-dry-run.txt "Verify:"
assert_contains /tmp/loop-dry-run.txt "test -f smoke.done"

repo1="$tmp/repo-pass"
mkdir -p "$repo1"
make_queue "$repo1" "test -f smoke.done"
LOOP_AGENT_CMD='touch smoke.done; echo worker pass' "$root/skills/nospec/scripts/nospec" run "$repo1/.loop/QUEUE.md" --max-ticks 1 >/tmp/loop-pass.txt
assert_contains "$repo1/.loop/QUEUE.md" "Status: done"
assert_contains "$repo1/.loop/EVIDENCE.md" "Status: done"
assert_contains "$repo1/.loop/EVIDENCE.md" "worker pass"

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
echo "# AGENTS v1" > "$repo_pin/AGENTS.md"
make_queue "$repo_pin" "test -f pin1.done"
( cd "$repo_pin" && git add -A && git commit -q -m init )
LOOP_AGENT_CMD='touch pin1.done; echo "# AGENTS v2" > AGENTS.md' \
  "$root/skills/nospec/scripts/nospec" run "$repo_pin/.loop/QUEUE.md" --max-ticks 1 >/dev/null 2>&1
assert_contains "$repo_pin/.loop/EVIDENCE.md" "Pinned: AGENTS.md @"
# No pin alerts on the first cycle
if grep -q 'Pin alert:' "$repo_pin/.loop/EVIDENCE.md"; then
  echo "expected no pin alerts on first cycle" >&2
  exit 1
fi
( cd "$repo_pin" && git add -A && git commit -q -m "cycle 1" )
make_queue "$repo_pin" "test -f pin2.done"
LOOP_AGENT_CMD='touch pin2.done; echo "# AGENTS v3" > AGENTS.md' \
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
  "$root/skills/nospec/scripts/nospec" run "$repo_review/.loop/QUEUE.md" --review --max-ticks 2 >/tmp/loop-review.txt
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

"$root/skills/nospec/scripts/nospec" view --repo "$repo_view" >/tmp/loop-view.txt
assert_contains /tmp/loop-view.txt "Nospec Dashboard"
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
# adrs should list all 20 ADRs (19 + ADR-0020 for the rename)
"$root/skills/nospec/scripts/nospec" --repo "$root" adrs >/tmp/nospec-adrs.txt
adr_count=$(grep -c '^ADR-' /tmp/nospec-adrs.txt)
if [[ "$adr_count" -ne 20 ]]; then
  echo "expected 20 ADRs, got $adr_count" >&2
  cat /tmp/nospec-adrs.txt >&2
  exit 1
fi
# check must pass on the real repo
"$root/skills/nospec/scripts/nospec" --repo "$root" check >/tmp/nospec-check.txt
assert_contains /tmp/nospec-check.txt "all checks passed"

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
