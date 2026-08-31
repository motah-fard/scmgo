# Contributing to scmgo

Thanks for considering a contribution. `scmgo` stays useful by staying
consistent: every package follows the same shape, so anyone who has read one
function can predict how the next one works. Please follow the pattern below
rather than introducing a new style.

## Before you start

- **New formula?** Open an issue first (use the "Feature request" template)
  with a citation — textbook, paper, or industry standard — for the formula.
  Supply chain formulas have variants (e.g. multiple definitions of "safety
  stock"); the citation avoids merging something subtly different from what
  the name promises.
- **New package?** Propose the package and its function list in an issue
  before sending a PR. Packages are cheap to add and expensive to remove once
  people depend on them.

## The package pattern

Every package (see `inventory/` for the reference implementation) is laid out
as:

```
<package>/
  doc.go              # package doc comment: scope, assumptions, what's excluded
  types.go            # exported Input/Result structs
  errors.go           # sentinel errors (errors.New), one per validation failure
  <formula>.go         # one exported function per file, named after the formula
  <formula>_test.go    # table-driven tests: valid cases + every error branch
  example_test.go      # one runnable Example<Func> per exported function
```

Conventions to follow:

1. **Inputs are structs, not positional args.** `func EOQ(in EOQInput) (float64, error)`,
   not `func EOQ(demand, orderingCost, holdingCost float64)`. This keeps call
   sites self-documenting and lets us add fields without breaking callers.
2. **Validate, then compute.** Check every field for domain validity (e.g.
   non-negative demand, service level strictly between 0 and 1) before doing
   arithmetic. Return the zero value and a sentinel error on failure — never
   panic.
3. **Sentinel errors, reused across functions.** If `ErrNegativeDemand`
   already exists in `errors.go`, reuse it; don't create a near-duplicate.
   Add new sentinels only for genuinely new failure modes. When a function
   composes several validations (see `inventory/policy_summary_internal.go`),
   wrap with `errors.Join(ErrInvalidXInput, ErrSpecificCause)` so callers can
   match on either the general or the specific error via `errors.Is`.
4. **Doc comments cite the formula.** Exported functions get a doc comment
   showing the formula (as a comment-block equation, see `eoq.go`) and, where
   the name is ambiguous (e.g. multiple "safety stock" definitions exist in
   the literature), which definition is used.
5. **`doc.go` states assumptions and exclusions.** Say explicitly what the
   package does *not* do (see `inventory/doc.go`'s closing bullet list). This
   is what lets users trust the package instead of guessing.
6. **Tests are table-driven** with `tolerance`-based float comparisons (see
   `almostEqual` in `inventory/eoq_test.go`), and cover every error branch,
   not just the happy path.
7. **Every exported function gets a runnable `Example`.** These double as
   documentation on pkg.go.dev and as regression tests via `// Output:`.

## Before opening a PR

```bash
go build ./...
go vet ./...
gofmt -l .        # must print nothing
go test ./... -race -cover
```

Then update:
- `README.md` — add the function/package to the relevant section.
- `CHANGELOG.md` — add an entry under `[Unreleased]`.

Use the PR template's checklist; it mirrors this list.

## Scope boundaries

`scmgo` favors deterministic, textbook-verifiable formulas over statistical
estimation or optimization that requires tunable models (e.g. no ML-based
demand forecasting, no generic solver-based network optimization). If your
contribution needs a fitted model or an external solver dependency, raise it
in an issue first — it may be better suited to a separate, related module.

## Code of conduct

Participation in this project is governed by [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md).
