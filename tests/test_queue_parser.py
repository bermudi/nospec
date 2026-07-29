"""Focused tests for the canonical queue parser.

The shell suite exercises parser behavior through the runner. These tests call the
parser directly so malformed-input failures point at the parser contract rather
than at a later runner phase.
"""

from __future__ import annotations

import contextlib
import importlib.util
import io
import sys
import tempfile
import unittest
from pathlib import Path

PARSER_PATH = (
    Path(__file__).resolve().parents[1]
    / "skills"
    / "nospec-loop"
    / "scripts"
    / "queue_parser.py"
)
SPEC = importlib.util.spec_from_file_location("nospec_queue_parser", PARSER_PATH)
if SPEC is None or SPEC.loader is None:  # pragma: no cover - import setup failure
    raise RuntimeError(f"cannot load queue parser from {PARSER_PATH}")
queue_parser = importlib.util.module_from_spec(SPEC)
sys.modules[SPEC.name] = queue_parser
SPEC.loader.exec_module(queue_parser)


def queue_with(body: str) -> str:
    return f"""# Loop Queue: parser unit test

Goal:
Exercise the parser directly.

{body}"""


class QueueParserTests(unittest.TestCase):
    def setUp(self) -> None:
        self.temp_dir = tempfile.TemporaryDirectory()
        self.addCleanup(self.temp_dir.cleanup)
        self.path = Path(self.temp_dir.name) / "QUEUE.md"

    def write(self, text: str) -> Path:
        self.path.write_text(text)
        return self.path

    def parse(self, text: str, *, validate_shell: bool = False):
        return queue_parser.parse_queue(
            self.write(text), validate_shell=validate_shell
        )

    def test_parses_fields_but_ignores_heading_and_field_examples_in_fences(self) -> None:
        text = queue_with(
            """##   normalized outcome   

Agent: printf 'worker'

Done means:
- The real command is selected.

````markdown
## fake outcome
Status: done
Verify:
```bash
false
```
````

Verify:
~~~bash
printf 'real verify'
~~~

Status: pending
"""
        )

        units, diagnostics = self.parse(text)

        self.assertEqual([], diagnostics)
        self.assertEqual(1, len(units))
        unit = units[0]
        self.assertEqual("normalized outcome", unit.title)
        self.assertEqual("pending", unit.status)
        self.assertEqual("printf 'worker'", unit.agent)
        self.assertEqual("printf 'real verify'\n", unit.verify)
        self.assertIn("## fake outcome", unit.body)

    def test_reports_independent_structural_errors_together(self) -> None:
        text = queue_with(
            """## 

Read first:
Constraints: inline value

Done means:

Verify:
```bash
:
```

Status: mystery
Status: pending
"""
        )

        _, diagnostics = self.parse(text)
        messages = [diagnostic.message for diagnostic in diagnostics]

        self.assertIn("work unit outcome is empty", messages)
        self.assertIn("Read first field is empty", messages)
        self.assertIn(
            "Constraints content must start on the following line", messages
        )
        self.assertIn("Done means field is empty", messages)
        self.assertIn("duplicate Status fields", messages)
        self.assertIn(
            "Verify command is obviously vacuous; assert the unit outcome", messages
        )

    def test_reports_missing_units_and_unclosed_fence(self) -> None:
        _, diagnostics = self.parse("# Queue without units\n\n```markdown\n")

        self.assertEqual(
            [
                "queue has no work units (`## <outcome>` headings)",
                "unclosed Markdown fence",
            ],
            [diagnostic.message for diagnostic in diagnostics],
        )

    def test_shell_validation_maps_error_to_queue_source(self) -> None:
        text = queue_with(
            """## invalid shell

Done means:
- Shell syntax is checked.

Verify:
```bash
if then
  echo broken
fi
```

Status: pending
"""
        )

        _, diagnostics = self.parse(text, validate_shell=True)
        shell_errors = [
            diagnostic
            for diagnostic in diagnostics
            if diagnostic.message.startswith("invalid Verify shell syntax:")
        ]

        self.assertEqual(1, len(shell_errors))
        self.assertGreater(shell_errors[0].line, 5)

    def test_accessors_and_status_mutation_use_the_same_parsed_unit(self) -> None:
        text = queue_with(
            """## first [literal] outcome

Done means:
- The first marker exists.

Verify:
```bash
test -f first.done
```

Status: pending

## second outcome

Done means:
- The second marker exists.

Verify:
```bash
test -f second.done
```

Status: done
"""
        )
        path = self.write(text)

        output = io.StringIO()
        with contextlib.redirect_stdout(output):
            queue_parser.command_first_pending_title(path)
        self.assertEqual("first [literal] outcome", output.getvalue())

        queue_parser.command_set_status(path, "first [literal] outcome", "in_progress")
        units, diagnostics = queue_parser.parse_queue(path, validate_shell=False)
        self.assertEqual([], diagnostics)
        self.assertEqual(["in_progress", "done"], [unit.status for unit in units])
        self.assertEqual(1, path.read_text().count("Status: in_progress"))

    def test_rejects_invalid_status_mutation_and_unknown_title(self) -> None:
        text = queue_with(
            """## known outcome

Done means:
- The queue remains valid.

Verify:
```bash
test -f known.done
```

Status: pending
"""
        )
        path = self.write(text)

        with self.assertRaisesRegex(SystemExit, "invalid status: invented"):
            queue_parser.command_set_status(path, "known outcome", "invented")
        with self.assertRaisesRegex(SystemExit, "unit not found or ambiguous"):
            queue_parser.command_unit_verify(path, "missing outcome")


if __name__ == "__main__":
    unittest.main()
