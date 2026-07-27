# Loop Queue: wiki validator cleanup

Goal:
Make one concrete class of AgenticWiki validator failures disappear without broad wiki rewrites.

Stop condition:
`./scripts/validate-page` exits 0, or the queue reaches a deliberately documented blocker.

## the validator reports fewer broken-link failures after one narrow link-fix pass

Read first:
- `./scripts/validate-page` output in the AgenticWiki repo
- The page(s) containing the broken-link pattern

Constraints:
- Fix only one repeated broken-link pattern or one page-local cluster.
- Do not rewrite page content beyond the link fix and required `updated` date changes.

Done means:
- The chosen broken-link failure no longer appears.
- No new validator failure class is introduced by the edits.

Verify:
```bash
./scripts/validate-page
```

Status: pending

## one reported frontmatter failure becomes valid under the real validator

Read first:
- `./scripts/validate-page` output in the AgenticWiki repo
- The page with the frontmatter failure

Constraints:
- Fix only that page's required metadata.
- Update the page's `updated` date if the page content or frontmatter changes.

Done means:
- The selected frontmatter failure no longer appears.
- Remaining failures, if any, are not caused by this edit.

Verify:
```bash
./scripts/validate-page
```

Status: pending
