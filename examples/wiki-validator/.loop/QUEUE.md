# Loop Queue: wiki validator cleanup

Goal:
Make one concrete class of AgenticWiki validator failures disappear without broad wiki rewrites.

Stop condition:
`./scripts/validate-page` exits 0, or the queue reaches a deliberately documented blocker.

## the validator reports fewer broken-link failures after one narrow link-fix pass

Work:
- Run `./scripts/validate-page` in the AgenticWiki repo.
- Pick one repeated broken-link pattern or one page-local cluster of broken links.
- Fix only that pattern or cluster.
- Do not rewrite page content beyond the link fix and required `updated` date changes.

Verify:
```bash
./scripts/validate-page
```

Done means:
- The chosen broken-link failure no longer appears.
- No new validator failure class is introduced by the edits.

Status: pending

## one reported frontmatter failure becomes valid under the real validator

Work:
- Run `./scripts/validate-page` in the AgenticWiki repo.
- Pick one frontmatter failure reported by the validator.
- Fix only that page's required metadata.
- Update the page's `updated` date if the page content or frontmatter changes.

Verify:
```bash
./scripts/validate-page
```

Done means:
- The selected frontmatter failure no longer appears.
- Remaining failures, if any, are not caused by this edit.

Status: pending
