#!/usr/bin/env bash
set -euo pipefail

root=$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)
tmp=$(mktemp -d)
trap 'rm -rf "$tmp"' EXIT

assert_contains() {
  local file=$1 pattern=$2
  if ! grep -Fq "$pattern" "$file"; then
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
Exercise sliceloop.

Stop condition:
\`$verify\` exits 0.

## Slice 1: the test fixture reaches its verify condition

Why this is vertical:
The worker changes repo state and the runner verifies the visible result.

Work:
- Satisfy the verify command.

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
"$root/loop.sh" run "$root/examples/smoke/.loop/QUEUE.md" --dry-run >/tmp/sliceloop-dry-run.txt
assert_contains /tmp/sliceloop-dry-run.txt "Verify:"
assert_contains /tmp/sliceloop-dry-run.txt "test -f smoke.done"

repo1="$tmp/repo-pass"
mkdir -p "$repo1"
make_queue "$repo1" "test -f smoke.done"
SLICELOOP_AGENT_CMD='touch smoke.done; echo worker pass' "$root/loop.sh" run "$repo1/.loop/QUEUE.md" --max-ticks 1 >/tmp/sliceloop-pass.txt
assert_contains "$repo1/.loop/QUEUE.md" "Status: done"
assert_contains "$repo1/.loop/EVIDENCE.md" "Status: done"
assert_contains "$repo1/.loop/EVIDENCE.md" "worker pass"

repo2="$tmp/repo-fail"
mkdir -p "$repo2"
make_queue "$repo2" "test -f never-created"
set +e
SLICELOOP_AGENT_CMD='echo worker failed to create file' "$root/loop.sh" run "$repo2/.loop/QUEUE.md" --max-ticks 1 >/tmp/sliceloop-fail.txt 2>&1
code=$?
set -e
if [[ $code -eq 0 ]]; then
  echo "expected verify failure to exit nonzero" >&2
  exit 1
fi
assert_contains "$repo2/.loop/QUEUE.md" "Status: pending"
assert_contains "$repo2/.loop/EVIDENCE.md" "Status: verify_failed"
assert_contains /tmp/sliceloop-fail.txt "retrying once"

repo3="$tmp/target-repo"
queue_home="$tmp/external-queue"
mkdir -p "$repo3" "$queue_home/.loop"
make_queue "$queue_home" "test -f target.done"
SLICELOOP_AGENT_CMD='pwd > worker.pwd; touch target.done' "$root/loop.sh" run "$queue_home/.loop/QUEUE.md" --repo "$repo3" --max-ticks 1 >/tmp/sliceloop-repo.txt
assert_contains "$queue_home/.loop/QUEUE.md" "Status: done"
test -f "$repo3/target.done"
assert_contains "$repo3/worker.pwd" "$repo3"

if command -v skills-ref >/dev/null 2>&1; then
  skills-ref validate "$root/.agents/skills/vertical-slice-planner"
fi

echo "sliceloop tests passed"
