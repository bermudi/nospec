# Loop Queue: smoke fixture

Goal:
Demonstrate the loop queue protocol with a harmless file-based verification gate.

Stop condition:
`test -f smoke.done` exits 0.

## the smoke fixture creates a file that the verify gate can see

Work:
- Create `smoke.done` in this example directory.
- Do not modify the queue by hand.

Verify:
```bash
test -f smoke.done
```

Done means:
- `smoke.done` exists.
- The verify command exits 0.

Status: pending
