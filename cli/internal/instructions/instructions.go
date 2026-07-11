package instructions

import (
	"fmt"
	"io"
)

// Print writes the template and guidance for the requested artifact to w.
// It returns an error if the artifact is not recognized.
func Print(w io.Writer, artifact string) error {
	switch artifact {
	case "work-unit":
		_, err := fmt.Fprint(w, workUnitTemplate)
		return err
	case "adr":
		_, err := fmt.Fprint(w, adrTemplate)
		return err
	case "glossary-entry":
		_, err := fmt.Fprint(w, glossaryEntryTemplate)
		return err
	default:
		return fmt.Errorf("unknown artifact: %s", artifact)
	}
}

const workUnitTemplate = `## <outcome>

Agent: <optional — overrides LOOP_AGENT_CMD for this unit only>

Why:
<only if non-obvious — else omit>

Read first:
- <context the worker needs: ADR, area, or file>
- <2–4 entries; context, not scope>

Constraints:
- <boundary or guardrail>
- <what must stay true or what is out of bounds>
- <if it names a file, it is "don't touch X" or "X's public API must not change", not "update X">

Done means:
- <observable condition>
- <no regression condition>

Verify:
` + "```bash" + `
<command that exits 0 on success>
` + "```" + `

Status: pending
`

const adrTemplate = `# NNNN: <title>

Date: <YYYY-MM-DD>
Status: accepted
# Supersedes: ADR-NNNN      # if this replaces an earlier ADR
# Superseded by: ADR-NNNN   # if a later ADR replaces this one

## Context

<what situation led to this decision>

## Decision

<the architectural ruling>

## Consequences

<what this choice makes easy or hard>
`

const glossaryEntryTemplate = `## <term>

<definition>

Used in: <where this term appears in the codebase>
`
