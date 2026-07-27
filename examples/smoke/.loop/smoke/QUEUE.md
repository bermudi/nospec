# Loop Queue: smoke fixture

Goal:
Demonstrate the loop queue protocol with a harmless file-based verification gate.

Stop condition:
`test -f smoke.done` exits 0.

## the smoke fixture creates a file that the verify gate can see

Read first:
- This queue file.

Constraints:
- Do not modify the queue by hand.

Done means:
- `smoke.done` exists.
- The verify command exits 0.

Verify:
```bash
test -f smoke.done
```

Status: pending
