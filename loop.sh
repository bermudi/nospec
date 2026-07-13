#!/usr/bin/env bash
set -euo pipefail

usage() {
  cat <<'EOF'
Usage:
  ./loop.sh run <queue> [--repo DIR] [--max-ticks N] [--review] [--max-review-rounds N] [--dry-run]

The queue is usually .loop/<name>/QUEUE.md in the target repo. Use --repo when the
queue lives outside the repository it should operate on.
EOF
}

die() {
  echo "knack: $*" >&2
  exit 1
}

abs_path() {
  python -c 'import os, sys; print(os.path.abspath(sys.argv[1]))' "$1"
}

first_pending_unit() {
  awk '
    BEGIN { in_block = 0; block = ""; status = ""; found = 0 }
    /^## / {
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

extract_agent() {
  awk '/^Agent:[[:space:]]+/ { sub(/^Agent:[[:space:]]*/, ""); gsub(/[[:space:]]+$/, ""); print; exit }' "$1"
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
    if stripped.startswith("## ") and not stripped.startswith("### "):
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
    raise SystemExit(f"unit not found or status already changed: {title}")
path.write_text("".join(out))
PY
}

write_handoff() {
  python - "$1" "$2" "$3" <<'PY'
import sys, re
from pathlib import Path
from datetime import datetime

queue, evidence, handoff = sys.argv[1:]
lines = Path(queue).read_text().splitlines()

units = []
current = None
for line in lines:
    if re.match(r'^## ', line) and not re.match(r'^###', line):
        if current:
            units.append(current)
        current = {"title": line[3:].strip(), "status": "pending"}
    elif current:
        m = re.match(r'^Status:\s*(\S+)', line)
        if m:
            current["status"] = m.group(1)
if current:
    units.append(current)

pending = [u for u in units if u["status"] != "done"]
if not pending:
    sys.exit(0)

completed = [u for u in units if u["status"] == "done"]
in_progress = [u for u in units if u["status"] in ("in_progress", "verify_failed", "no_progress", "blocked")]
remaining = [u for u in units if u["status"] == "pending"]

out = [
    f"# Handoff: {Path(queue).stem}",
    f"Generated: {datetime.now().isoformat()}",
    "",
    "## Completed",
]
out += [f"- {u['title']}" for u in completed] or ["- (none)"]
out += ["", "## In progress"]
out += [f"- {u['title']} (status: {u['status']})" for u in in_progress] or ["- (none)"]
out += ["", "## Remaining"]
out += [f"- {u['title']}" for u in remaining] or ["- (none)"]
out += ["", "## Next action"]
if in_progress:
    out.append(f"Re-run loop after addressing the {in_progress[0]['status']} state of: {in_progress[0]['title']}.")
elif remaining:
    out.append(f"Re-run loop to continue with: {remaining[0]['title']}.")
else:
    out.append("Queue is complete.")

Path(handoff).write_text("\n".join(out) + "\n")
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
  local evidence=$1 title=$2 status=$3 verify=$4 verify_out=$5 agent_out=$6 repo_dir=$7 unit_file=$8
  mkdir -p "$(dirname "$evidence")"
  {
    echo
    echo "## $(date -Iseconds) — $title"
    echo
    echo "Status: $status"
    echo
    echo "Unit:"
    echo '````markdown'
    cat "$unit_file"
    echo '````'
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
    echo '````text'
    cat "$agent_out"
    echo '````'
    echo
    echo "What this proves:"
    if [[ "$status" == "done" ]]; then
      echo "- The verify command passed for this work unit in the current repo state."
    else
      echo "- The work unit is not externally verified."
    fi
    echo
    echo "What remains unverified:"
    echo "- Anything outside the verify command's proof scope."
  } >> "$evidence"
}

extract_actionable_count() {
  awk '
    /^[[:space:]]*-[[:space:]]*actionable:[[:space:]]*[0-9]+[[:space:]]*$/ {
      sub(/^[[:space:]]*-[[:space:]]*actionable:[[:space:]]*/, "")
      sub(/[[:space:]]*$/, "")
      print
      found = 1
      exit
    }
    END { if (!found) exit 1 }
  ' "$1"
}

write_review_prompt() {
  local out=$1
  local template="$script_dir/prompts/reviewer.md"
  if [[ -f "$template" ]]; then
    cat "$template" > "$out"
  else
    cat > "$out" <<'EOF'
# Knack Reviewer

Load and follow the **review** skill in `skills/review/`.
Review the completed queue against the current repository state. Write the structured review artifact at the path provided below.
EOF
  fi
  cat >> "$out" <<EOF

Queue: $queue_abs
Evidence: $evidence
Review output: $review_file
EOF
  if [[ -f "$design_file" ]]; then
    echo "Design: $design_file" >> "$out"
  fi
}

write_fix_prompt() {
  local out=$1
  local template="$script_dir/prompts/fixer.md"
  if [[ -f "$template" ]]; then
    cat "$template" > "$out"
  else
    cat > "$out" <<'EOF'
# Knack Fixer

Load and follow the **fix** skill in `skills/fix/`.
Read the structured review artifact and append any actionable fix work units to the existing queue. Stop after updating the queue.
EOF
  fi
  cat >> "$out" <<EOF

Queue: $queue_abs
Evidence: $evidence
Review input: $review_file
EOF
  if [[ -f "$design_file" ]]; then
    echo "Design: $design_file" >> "$out"
  fi
}

run_phase_agent() {
  local phase=$1 prompt=$2 output=$3 cmd=${4:-}
  if [[ -n "$cmd" ]]; then
    (
      cd "$repo_dir"
      LOOP_PHASE="$phase" \
        LOOP_PROMPT_FILE="$prompt" \
        LOOP_QUEUE_FILE="$queue_abs" \
        LOOP_EVIDENCE_FILE="$evidence" \
        LOOP_REVIEW_FILE="$review_file" \
        bash -lc "$cmd"
    ) > "$output" 2>&1
  else
    (
      cd "$repo_dir"
      LOOP_PHASE="$phase" \
        LOOP_PROMPT_FILE="$prompt" \
        LOOP_QUEUE_FILE="$queue_abs" \
        LOOP_EVIDENCE_FILE="$evidence" \
        LOOP_REVIEW_FILE="$review_file" \
        pi -p --no-session --approve "$(cat "$prompt")"
    ) > "$output" 2>&1
  fi
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
max_review_rounds=2
review_enabled=0
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
    --review)
      review_enabled=1
      shift
      ;;
    --max-review-rounds)
      [[ $# -ge 2 ]] || die "--max-review-rounds needs a value"
      max_review_rounds=$2
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
handoff="$queue_dir/HANDOFF.md"
review_file="$queue_dir/REVIEW.md"
design_file="$queue_dir/DESIGN.md"
no_progress_strikes=0

write_handoff_on_exit() {
  local rc=$?
  [[ "${dry_run:-0}" == 0 ]] || return $rc
  [[ -f "${queue_abs:-}" ]] || return $rc
  write_handoff "$queue_abs" "$evidence" "$handoff" 2>/dev/null || true
  return $rc
}
trap write_handoff_on_exit EXIT

tick=1
review_round=0

while true; do
while (( tick <= max_ticks )); do
  unit=$(first_pending_unit "$queue_abs")
  [[ -n "$unit" ]] || break

  unit_file=$(mktemp)
  verify_file=$(mktemp)
  agent_out=$(mktemp)
  verify_out=$(mktemp)
  printf '%s' "$unit" > "$unit_file"

  first_line=$(awk 'NR == 1 { print; exit }' "$unit_file")
  title=${first_line#\#\# }
  verify=$(extract_verify "$unit_file")
  [[ -n "$verify" ]] || die "work unit has no Verify fenced block: $title"
  printf '%s\n' "$verify" > "$verify_file"

  if [[ "$dry_run" == 1 ]]; then
    echo "Unit: $title"
    echo "Repo: $repo_dir"
    echo "Verify:"
    cat "$verify_file"
    exit 0
  fi

  echo "knack: tick $tick/$max_ticks — $title"
  set_status "$queue_abs" "$title" "in_progress"
  before=$(work_snapshot "$repo_dir")

  run_prompt=$(mktemp)
  cat > "$run_prompt" <<EOF
$(cat "$prompt_file")

Current work unit from $queue_abs:

$(cat "$unit_file")
EOF

  agent_cmd="${LOOP_AGENT_CMD:-}"
  unit_agent=$(extract_agent "$unit_file")
  if [[ -n "$unit_agent" ]]; then
    agent_cmd="$unit_agent"
  fi

  set +e
  run_phase_agent "build" "$run_prompt" "$agent_out" "$agent_cmd"
  agent_code=$?
  set -e

  if [[ $agent_code -ne 0 ]]; then
    set_status "$queue_abs" "$title" "blocked"
    : > "$verify_out"
    append_evidence "$evidence" "$title" "blocked" "$verify" "$verify_out" "$agent_out" "$repo_dir" "$unit_file"
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
    append_evidence "$evidence" "$title" "done" "$verify" "$verify_out" "$agent_out" "$repo_dir" "$unit_file"
    echo "knack: verified — $title"
    tick=$((tick + 1))
    continue
  fi

  if [[ "$before" == "$after" ]]; then
    no_progress_strikes=$((no_progress_strikes + 1))
    if [[ $no_progress_strikes -ge 2 ]]; then
      set_status "$queue_abs" "$title" "no_progress"
      append_evidence "$evidence" "$title" "no_progress" "$verify" "$verify_out" "$agent_out" "$repo_dir" "$unit_file"
      die "no progress after $no_progress_strikes attempts on $title"
    fi
    set_status "$queue_abs" "$title" "pending"
    append_evidence "$evidence" "$title" "verify_failed" "$verify" "$verify_out" "$agent_out" "$repo_dir" "$unit_file"
    echo "knack: verify failed with no progress; retrying once"
    tick=$((tick + 1))
    continue
  fi

  set_status "$queue_abs" "$title" "verify_failed"
  append_evidence "$evidence" "$title" "verify_failed" "$verify" "$verify_out" "$agent_out" "$repo_dir" "$unit_file"
  cat "$verify_out"
  die "verify failed for $title"
done

if [[ -n "$(first_pending_unit "$queue_abs")" ]]; then
  die "reached max ticks ($max_ticks) with pending work"
fi

if [[ "$review_enabled" == 0 ]]; then
  if (( tick > max_ticks )); then
    echo "knack: reached max ticks ($max_ticks)"
    exit 0
  fi
  echo "knack: no pending work units"
  exit 0
fi

if (( review_round >= max_review_rounds )); then
  die "reached max review rounds ($max_review_rounds)"
fi

review_round=$((review_round + 1))
review_prompt=$(mktemp)
review_out=$(mktemp)
write_review_prompt "$review_prompt"
review_cmd="${LOOP_REVIEW_CMD:-${LOOP_AGENT_CMD:-}}"

echo "knack: review round $review_round/$max_review_rounds"
set +e
run_phase_agent "review" "$review_prompt" "$review_out" "$review_cmd"
review_code=$?
set -e
if [[ $review_code -ne 0 ]]; then
  cat "$review_out"
  die "review worker exited nonzero"
fi
[[ -f "$review_file" ]] || die "review worker did not write $review_file"
actionable=$(extract_actionable_count "$review_file") || die "review file has no actionable summary: $review_file"

if [[ "$actionable" == 0 ]]; then
  echo "knack: review clean"
  exit 0
fi

if (( review_round >= max_review_rounds )); then
  die "review found $actionable actionable issue(s) at max review rounds ($max_review_rounds)"
fi

fix_prompt=$(mktemp)
fix_out=$(mktemp)
write_fix_prompt "$fix_prompt"
fix_cmd="${LOOP_FIX_CMD:-${LOOP_AGENT_CMD:-}}"

echo "knack: fix round $review_round/$max_review_rounds — actionable: $actionable"
set +e
run_phase_agent "fix" "$fix_prompt" "$fix_out" "$fix_cmd"
fix_code=$?
set -e
if [[ $fix_code -ne 0 ]]; then
  cat "$fix_out"
  die "fix worker exited nonzero"
fi

if [[ -z "$(first_pending_unit "$queue_abs")" ]]; then
  cat "$fix_out"
  echo "knack: fix produced no new units — review findings overturned or triaged out; stopping clean"
  exit 0
fi
done
