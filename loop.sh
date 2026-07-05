#!/usr/bin/env bash
set -euo pipefail

usage() {
  cat <<'EOF'
Usage:
  ./loop.sh run <queue> [--repo DIR] [--max-ticks N] [--dry-run]

The queue is usually .loop/QUEUE.md in the target repo. Use --repo when the
queue lives outside the repository it should operate on.
EOF
}

die() {
  echo "sliceloop: $*" >&2
  exit 1
}

abs_path() {
  python -c 'import os, sys; print(os.path.abspath(sys.argv[1]))' "$1"
}

first_pending_slice() {
  awk '
    BEGIN { in_block = 0; block = ""; status = ""; found = 0 }
    /^## Slice / {
      if (in_block && status == "pending") { printf "%s", block; found = 1; exit }
      in_block = 1; block = $0 "\n"; status = ""; next
    }
    in_block {
      block = block $0 "\n"
      if ($0 ~ /^Status:[[:space:]]*pending[[:space:]]*$/) status = "pending"
    }
    END { if (!found && in_block && status == "pending") printf "%s", block }
  ' "$1"
}

extract_verify() {
  awk '
    BEGIN { after_verify = 0; in_fence = 0 }
    after_verify && /^```/ { if (!in_fence) { in_fence = 1; next } else exit }
    after_verify && in_fence { print; next }
    /^Verify:[[:space:]]*$/ { after_verify = 1 }
  ' "$1"
}

set_status() {
  python - "$1" "$2" "$3" <<'PY'
import sys
from pathlib import Path

queue, title, status = sys.argv[1:]
path = Path(queue)
lines = path.read_text().splitlines(keepends=True)
heading = f"## {title}"
inside = False
changed = False
out = []
for line in lines:
    stripped = line.strip()
    if stripped.startswith("## Slice "):
        if inside and not changed:
            out.append(f"Status: {status}\n")
            changed = True
        inside = stripped == heading
    if inside and stripped.startswith("Status:") and not changed:
        out.append(f"Status: {status}\n")
        changed = True
        continue
    out.append(line)
if inside and not changed:
    if out and not out[-1].endswith("\n"):
        out[-1] += "\n"
    out.append(f"Status: {status}\n")
    changed = True
if not changed:
    raise SystemExit(f"slice not found or status already changed: {title}")
path.write_text("".join(out))
PY
}

work_snapshot() {
  local repo_dir=$1
  if git -C "$repo_dir" rev-parse --is-inside-work-tree >/dev/null 2>&1; then
    {
      git -C "$repo_dir" diff -- . ':(exclude).loop' || true
      git -C "$repo_dir" status --short --untracked-files=all | awk '$2 !~ /^\.loop\// { print }'
    } | sha256sum | awk '{print $1}'
  else
    echo "__no_git__"
  fi
}

changed_files() {
  local repo_dir=$1
  if git -C "$repo_dir" rev-parse --is-inside-work-tree >/dev/null 2>&1; then
    git -C "$repo_dir" status --short --untracked-files=all || true
  else
    echo "git status unavailable: not inside a git repo"
  fi
}

append_evidence() {
  local evidence=$1 title=$2 status=$3 verify=$4 verify_out=$5 agent_out=$6 repo_dir=$7
  mkdir -p "$(dirname "$evidence")"
  {
    echo
    echo "## $(date -Iseconds) — $title"
    echo
    echo "Status: $status"
    echo
    echo "Files changed:"
    echo '```text'
    changed_files "$repo_dir"
    echo '```'
    echo
    echo "Verify command:"
    echo '```bash'
    echo "$verify"
    echo '```'
    echo
    echo "Verify output:"
    echo '```text'
    cat "$verify_out"
    echo '```'
    echo
    echo "Worker output:"
    echo '```text'
    cat "$agent_out"
    echo '```'
    echo
    echo "What this proves:"
    if [[ "$status" == "done" ]]; then
      echo "- The verify command passed for this slice in the current repo state."
    else
      echo "- The slice is not externally verified."
    fi
    echo
    echo "What remains unverified:"
    echo "- Anything outside the verify command's proof scope."
  } >> "$evidence"
}

[[ $# -ge 1 ]] || { usage; exit 1; }
cmd=$1
shift
[[ "$cmd" == "run" ]] || { usage; exit 1; }
[[ $# -ge 1 ]] || die "missing queue path"

queue_abs=$(abs_path "$1")
shift
[[ -f "$queue_abs" ]] || die "queue not found: $queue_abs"

max_ticks=3
dry_run=0
repo_override=""
while [[ $# -gt 0 ]]; do
  case "$1" in
    --repo)
      [[ $# -ge 2 ]] || die "--repo needs a value"
      repo_override=$(abs_path "$2")
      shift 2
      ;;
    --max-ticks)
      [[ $# -ge 2 ]] || die "--max-ticks needs a value"
      max_ticks=$2
      shift 2
      ;;
    --dry-run)
      dry_run=1
      shift
      ;;
    *) die "unknown argument: $1" ;;
  esac
done

script_dir=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
prompt_file="$script_dir/prompts/worker.md"
[[ -f "$prompt_file" ]] || die "worker prompt not found: $prompt_file"

queue_dir=$(dirname "$queue_abs")
if [[ -n "$repo_override" ]]; then
  [[ -d "$repo_override" ]] || die "repo not found: $repo_override"
  repo_dir="$repo_override"
elif [[ $(basename "$queue_dir") == ".loop" ]]; then
  repo_dir=$(dirname "$queue_dir")
else
  repo_dir=$(pwd)
fi
evidence="$queue_dir/EVIDENCE.md"
no_progress_strikes=0

for ((tick = 1; tick <= max_ticks; tick++)); do
  slice=$(first_pending_slice "$queue_abs")
  [[ -n "$slice" ]] || { echo "sliceloop: no pending slices"; exit 0; }

  slice_file=$(mktemp)
  verify_file=$(mktemp)
  agent_out=$(mktemp)
  verify_out=$(mktemp)
  printf '%s' "$slice" > "$slice_file"

  first_line=$(awk 'NR == 1 { print; exit }' "$slice_file")
  title=${first_line#\#\# }
  verify=$(extract_verify "$slice_file")
  [[ -n "$verify" ]] || die "slice has no Verify fenced block: $title"
  printf '%s\n' "$verify" > "$verify_file"

  if [[ "$dry_run" == 1 ]]; then
    echo "Slice: $title"
    echo "Repo: $repo_dir"
    echo "Verify:"
    cat "$verify_file"
    exit 0
  fi

  echo "sliceloop: tick $tick/$max_ticks — $title"
  set_status "$queue_abs" "$title" "in_progress"
  before=$(work_snapshot "$repo_dir")

  run_prompt=$(mktemp)
  cat > "$run_prompt" <<EOF
$(cat "$prompt_file")

Current slice from $queue_abs:

$(cat "$slice_file")
EOF

  set +e
  if [[ -n "${SLICELOOP_AGENT_CMD:-}" ]]; then
    (cd "$repo_dir" && SLICELOOP_PROMPT_FILE="$run_prompt" bash -lc "$SLICELOOP_AGENT_CMD") > "$agent_out" 2>&1
  else
    (cd "$repo_dir" && pi -p --no-session "$(cat "$run_prompt")") > "$agent_out" 2>&1
  fi
  agent_code=$?
  set -e

  if [[ $agent_code -ne 0 ]]; then
    set_status "$queue_abs" "$title" "blocked"
    : > "$verify_out"
    append_evidence "$evidence" "$title" "blocked" "$verify" "$verify_out" "$agent_out" "$repo_dir"
    cat "$agent_out"
    die "worker exited nonzero for $title"
  fi

  after=$(work_snapshot "$repo_dir")

  set +e
  (cd "$repo_dir" && bash -lc "$verify") > "$verify_out" 2>&1
  verify_code=$?
  set -e

  if [[ $verify_code -eq 0 ]]; then
    set_status "$queue_abs" "$title" "done"
    append_evidence "$evidence" "$title" "done" "$verify" "$verify_out" "$agent_out" "$repo_dir"
    echo "sliceloop: verified — $title"
    continue
  fi

  if [[ "$before" == "$after" ]]; then
    no_progress_strikes=$((no_progress_strikes + 1))
    if [[ $no_progress_strikes -ge 2 ]]; then
      set_status "$queue_abs" "$title" "no_progress"
      append_evidence "$evidence" "$title" "no_progress" "$verify" "$verify_out" "$agent_out" "$repo_dir"
      die "no progress after $no_progress_strikes attempts on $title"
    fi
    set_status "$queue_abs" "$title" "pending"
    append_evidence "$evidence" "$title" "verify_failed" "$verify" "$verify_out" "$agent_out" "$repo_dir"
    echo "sliceloop: verify failed with no progress; retrying once"
    continue
  fi

  set_status "$queue_abs" "$title" "verify_failed"
  append_evidence "$evidence" "$title" "verify_failed" "$verify" "$verify_out" "$agent_out" "$repo_dir"
  cat "$verify_out"
  die "verify failed for $title"
done

if [[ -n "$(first_pending_slice "$queue_abs")" ]]; then
  die "reached max ticks ($max_ticks) with pending work"
fi

echo "sliceloop: reached max ticks ($max_ticks)"
