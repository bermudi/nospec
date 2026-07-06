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

Why:
<why this unit matters>

Work:
- <what needs to be done>

Verify:
` + "```bash" + `
<command that proves the unit is done>
` + "```" + `

Done means:
- <list of acceptance criteria>

Status: pending
`

const adrTemplate = `# NNNN: <title>

Date: <YYYY-MM-DD>
Status: proposed

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
