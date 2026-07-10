# Handoff: QUEUE
Generated: 2026-07-10T17:31:46.829245

## Completed
- ADR-0008 captures the loop-orchestrated review-fix decision

## In progress
- (none)

## Remaining
- review skill writes a structured REVIEW.md
- fix skill consumes REVIEW.md and appends fix units to QUEUE.md
- loop.sh orchestrates bounded build-review-fix rounds
- docs and AGENTS.md reflect the new loop behavior

## Next action
Re-run loop to continue with: review skill writes a structured REVIEW.md.
