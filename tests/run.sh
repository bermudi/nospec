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

if command -v skills-ref >/dev/null 2>&1; then
  for skill_dir in "$root/.agents/skills"/*; do
    if [[ -d "$skill_dir" ]]; then
      skills-ref validate "$skill_dir"
    fi
  done
fi

echo "knack tests passed"
