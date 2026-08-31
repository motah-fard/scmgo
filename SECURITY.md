# Security Policy

`scmgo` is a pure computation library (no network I/O, no filesystem access,
no external dependencies) operating on caller-supplied numeric input. Its
main attack surface is malformed input causing a panic, an infinite loop, or
a silently incorrect result in a caller's process.

## Supported Versions

Only the latest tagged release is supported. Security fixes are not
backported to older tags.

## Reporting a Vulnerability

Please report suspected vulnerabilities privately using [GitHub's private
vulnerability reporting](https://github.com/motah-fard/scmgo/security/advisories/new)
rather than a public issue. Include:

- The affected function(s) and version/commit
- Input that triggers the issue (panic, hang, or silently wrong result with
  a security-relevant consequence — e.g. a validation bypass)
- What you expected vs. what happened

We aim to acknowledge reports within a few days. If you don't hear back,
a public issue is a reasonable fallback.

## What's in scope

- Panics or hangs on any input, including adversarial `float64` values
  (`NaN`, `Inf`, subnormals) or malformed slices
- Validation bypasses: a documented invalid input (negative demand, a
  service level outside `(0, 1)`, non-finite values, etc.) that produces a
  result instead of an error
- Integer/floating-point overflow that silently produces an incorrect
  result without an error, where the inputs themselves were valid

## What's out of scope

- Correctness of the underlying supply-chain formulas themselves (report
  those as regular bugs — see [CONTRIBUTING.md](CONTRIBUTING.md))
- Resource exhaustion from a caller passing an enormous slice — the
  library does no bounds-limiting on collection sizes by design; that's the
  caller's responsibility, the same as any Go slice-processing function
