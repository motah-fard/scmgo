# Changelog

All notable changes to this project will be documented in this file.

## [Unreleased]
### Added (six follow-up functions)
- `quality.SigmaLevel` — converts DPMO to a Six-Sigma process sigma level
  (the conventional 1.5-sigma shift). Extracted the shared inverse-normal-
  CDF math into `internal/numeric.InverseNormalCDF` so it isn't duplicated
  with `inventory.ZScoreForServiceLevel`, which now calls the same helper
  (no behavior change, verified against its existing tests).
- `forecast.ClassifyDemandPattern` — Syntetos & Boylan (2005) demand
  classification (smooth/intermittent/erratic/lumpy), telling you which
  forecasting method fits a series.
- `forecast.SBA` — Syntetos-Boylan Approximation, a bias correction to
  `Croston` (`Forecast × (1 - Alpha/2)`); takes the same `CrostonInput`.
- `forecast.Naive` and `forecast.SeasonalNaive` — baseline forecasts, the
  natural thing to benchmark `MASE` against.
- `inventory.EOI` — economic order interval, `EOQ` expressed as a time
  interval instead of a quantity.
- `inventory.SafetyTime` — safety stock expressed as a time buffer instead
  of a quantity.

All six verified against independently computed Python reference values
(including cross-checks against known reference points — DPMO=3.4 gives
~6.0 sigma, the classic "Six Sigma" target) before being encoded as Go
tests. Full tests, examples, and fuzz targets.

### Added (finance package)
- New `finance` package: `DSO`, `DPO`, `CashToCashCycleTime` (takes DIO/
  DSO/DPO as direct inputs rather than recomputing them, so DIO can come
  from `inventory.DaysOfInventoryOnHand`; the result itself may be
  negative — a real, favorable outcome — even though each input must be
  individually non-negative), and `PerfectOrderRate` (multiplies its four
  component rates, assuming independence — documented as an upper-bound
  estimate if failure modes are correlated). This is the last of the
  planned domain packages; see the Roadmap section of README.md for what's
  deliberately still excluded and why. Full tests, examples, fuzz targets;
  100% coverage.

### Added (quality package)
- New `quality` package: `DPMO`, `Cp`/`Cpk` (assume a normal process
  distribution; `Cpk`'s `Mean` may fall outside `[LSL, USL]`, and a
  resulting negative value reflects a genuinely out-of-tolerance process
  rather than being rejected), and `CostOfQuality` (the standard four-
  category prevention/appraisal/internal-failure/external-failure
  breakdown, used directly as named fields rather than a caller-supplied
  component list, since this categorization is close to universal in
  quality management). Full tests, examples, fuzz targets; 97.1% coverage.

### Added (warehouse package)
- New `warehouse` package: `StorageUtilization`, `CubeUtilization` (neither
  clamped — a value above 1 signals overcommitted storage), and `PickRate`.
  Deliberately no separate velocity-based slotting function: that's exactly
  what `abc.Classify` already does (classify by pick velocity instead of
  usage value). Full tests, examples, fuzz targets; 100% coverage.

### Added (logistics package)
- New `logistics` package: `DimensionalWeight`/`BillableWeight` (the
  divisor is caller-supplied, not a hardcoded carrier constant, since it
  varies by carrier and changes over time), `AllocateFreightCost`
  (proportional allocation by a caller-supplied basis), `VehicleUtilization`
  (not clamped — a value above 1 signals an overload, which is meaningful),
  and `CenterOfGravity` (demand-weighted centroid for facility location;
  straight-line/Euclidean, not road-network distance). Full tests,
  examples, fuzz targets; 98.2% coverage.

### Added (production package)
- New `production` package: `OEE` (computed from raw availability/
  performance/quality inputs, not precomputed ratios; result is not
  clamped to `[0, 1]`), `WIPFromLittlesLaw`/`CycleTimeFromLittlesLaw`/
  `ThroughputFromLittlesLaw` (Little's Law, one explicitly named function
  per direction rather than an implicit "solve for X" parameter), and
  `TaktTime`. Deliberately does not include BOM explosion or lot-sizing
  algorithms (Wagner-Whitin, Silver-Meal, Part-Period Balancing) — those
  are graph-traversal/dynamic-programming problems, not closed-form
  formulas, and need more design and testing rigor than this pass gave the
  rest of the library's formulas. Full tests, examples, fuzz targets;
  95.7% coverage.

### Added (procurement package)
- New `procurement` package: `LandedCost` and `TotalCostOfOwnership` (sum
  labeled, possibly-negative cost components — deliberately not hardcoding
  a fixed category list, since it varies too much by industry) and
  `PurchasePriceVariance` (standard-costing convention: positive = favorable).
  Full tests, examples, fuzz targets; 100% coverage.

### Added (forecast extensions)
- `HoltWinters` — additive-seasonality triple exponential smoothing; one
  documented initialization convention (see its doc comment), since none is
  universally standard in the literature
- `TrackingSignal` — running per-period signal (cumulative error / MAD) for
  monitoring forecast drift, distinct from `Accuracy`'s single-summary metrics
- `LinearTrend` — OLS trend-line forecast, for a single trend with no seasonality
- `MASE` — mean absolute scaled error (Hyndman & Koehler, 2006), well-defined
  even when `Actual` contains zeros (unlike MAPE)

### Added (abc package)
- New `abc` package: `Classify` (ABC/Pareto classification by cumulative
  value contribution), `ClassifyVariability` (XYZ classification by demand
  coefficient of variation), and `Combine` (joins both into the classic
  ABC-XYZ matrix). `Classify` returns results in input order (not resorted
  by value), so it stays index-aligned with `ClassifyVariability` for the
  same item set. Full tests, examples, and fuzz targets.

### Added (inventory extensions)
- `SafetyStockWithVariableLeadTime` / `ReorderPointWithVariableLeadTime` —
  account for lead-time variability, not just demand variability; reduces
  to the existing fixed-lead-time formulas when lead-time variability is 0
- `UnitNormalLoss`, `ExpectedFillRate`, `FillRateSafetyStock` — fill-rate
  (P2) safety stock in both directions, closing the gap `doc.go` previously
  listed as excluded. `FillRateSafetyStock` inverts the unit normal loss
  function numerically (bisection, since it has no closed-form inverse)
- `EPQ` — economic production quantity
- `EOQWithQuantityDiscounts` — cost-minimizing order quantity across a
  price-break schedule; verified against a known textbook example
- `Newsvendor` — single-period optimal order quantity
- `Turnover`, `DaysOfInventoryOnHand`, `GMROI` — standard inventory ratios

All ship with table-driven tests (values verified independently in Python
before being encoded as Go expectations), runnable examples, and fuzz
targets.

### Fixed (found while building the above)
- **`UnitNormalLoss` catastrophic cancellation for large `z`.** Computing
  `1 - Φ(z)` as `1 - (1 + Erf(...))/2` loses precision once `Φ(z)` rounds to
  exactly 1 in float64 (around `z ≈ 6`), silently dropping the entire tail
  term instead of erroring. Caught by a round-trip test
  (`FillRateSafetyStock` → `ExpectedFillRate` should recover the original
  target) that a plain "not NaN" check would have missed, since the result
  was finite, just wrong. Fixed with `math.Erfc`, which computes the tail
  directly without subtracting two nearly-equal numbers.
- `ExpectedFillRate` rejected negative `SafetyStockUnits`, but
  `FillRateSafetyStock` can legitimately return a negative value for a low
  target fill rate — the two were inconsistent. Removed the restriction;
  negative safety stock is mathematically valid in this model.
- A test-data copy-paste error (wrong expected value carried over from an
  unrelated test case) in `SafetyStockWithVariableLeadTime`'s test table.
- Naming stutter: `InventoryTurnover`/`InventoryTurnoverInput` renamed to
  `Turnover`/`TurnoverInput` (caught by golangci-lint/revive) before this
  shipped in any tagged release.

### Changed
- **Clean Code pass:** eliminated the DRY violation across both packages'
  `validate_internal.go` (near-identical files, both wrapping the same
  NaN/Inf check) by extracting `internal/numeric.AllFinite` and making
  `validateFinite`/`validateSmoothingConstant` variadic, collapsing what was
  one `if` block per field into one call per function. Extracted
  `wrapPolicySummaryErr` to remove 8 repeated `errors.Join(ErrInvalidPolicySummaryInput, err)`
  call sites in `policy_summary.go`/`policy_summary_internal.go`. No
  behavior change; verified with the full fuzz/test/lint suite.

### Added
- `forecast` package: `MovingAverage`, `WeightedMovingAverage`,
  `SimpleExponentialSmoothing`, `HoltLinearTrend`, `Croston`,
  `Accuracy` (MAD, MAPE, Bias, RMSE), with full test coverage and
  runnable examples
- `inventory.BuildPolicySummaryBatch` and
  `inventory.BuildPolicySummaryWithServiceLevelBatch` for computing policy
  summaries across a list of SKUs without one bad row aborting the batch
- CI workflow (build, vet, gofmt check, race-enabled tests, golangci-lint,
  a short fuzzing pass per target) on push/PR
- `CONTRIBUTING.md` documenting the package pattern for new formulas/packages
- `CODE_OF_CONDUCT.md`
- `SECURITY.md`
- GitHub issue templates (bug report, feature request) and PR template
- `.gitignore`
- Fuzz targets (`FuzzEOQ`, `FuzzZScoreForServiceLevel`, `FuzzBuildPolicySummary*`
  in `inventory`; `FuzzMovingAverage`, `FuzzWeightedMovingAverage`,
  `FuzzHoltLinearTrend`, `FuzzCroston`, `FuzzAccuracy` in `forecast`),
  asserting no panic and no undocumented `NaN` leaking through with a `nil`
  error
- Benchmarks for the functions most likely to sit in a hot path (`EOQ`,
  `BuildPolicySummary*`, `MovingAverage`, `SimpleExponentialSmoothing`,
  `HoltLinearTrend`)

### Changed
- **Breaking (pre-release only):** renamed `forecast.ForecastAccuracy` /
  `ForecastAccuracyInput` / `ForecastAccuracyResult` to `Accuracy` /
  `AccuracyInput` / `AccuracyResult` to fix a package-name stutter flagged
  by `golangci-lint`/revive. Done now because `forecast` has not shipped in
  a tagged release yet.

### Fixed
- `go.mod` directive order (`module` before `go`) and relaxed the pinned Go
  version from `1.25.1` to `1.23` so the module builds on older toolchains
- **NaN/Inf silently passing validation in both packages.** Every negative
  and range check (e.g. `v < 0`, `0 < s < 1`) is a no-op against `NaN` in
  IEEE 754, so `EOQ`, `ReorderPointWithServiceLevel`, `MovingAverage`, and
  every other function that took a bad upstream `NaN` or `Inf` returned a
  poisoned `NaN`/`Inf` result with a `nil` error instead of failing. All
  `float64` inputs across both packages are now checked with a new
  `ErrNonFiniteInput` before any other validation.
- `.golangci.yml` was in v1 format; `golangci-lint-action@v6` with
  `version: latest` now installs golangci-lint v2, which refuses a v1
  config outright. The lint CI job would have failed on the first push.
  Migrated to v2 config format.

## [v1.0.0] - 2026-04-18
### Added
- first stable release of the `inventory` package

### Improved
- finalized public API for practical inventory policy calculations
- updated README for stable-release positioning and installation guidance
- clarified package scope, assumptions, and stability language
- completed documentation pass ahead of stable release
- kept examples and package documentation aligned with the `v1.0.0` API

## [v0.6.0] - 2026-04-18
### Improved
- reviewed naming consistency across the `inventory` package
- tightened exported comments and package documentation for clarity and consistency
- aligned README wording with current package semantics and assumptions
- clarified target inventory level wording around expected demand coverage
- refined package documentation for policy summary helpers and service-level-based calculations
- performed package consistency review ahead of `v1.0.0`
- kept the public API stable while improving documentation and overall polish

## [v0.5.0] - 2026-04-18
### Added
- `PolicySummary` for consolidated inventory policy outputs
- `PolicySummaryInput` for deterministic policy summary calculations
- `PolicySummaryServiceLevelInput` for service-level-driven policy summary calculations
- `BuildPolicySummary` for computing expected lead-time demand, safety stock, reorder point, target inventory level, and min/max levels in one call
- `BuildPolicySummaryWithServiceLevel` for computing policy summaries using service-level-based reorder point logic
- examples for deterministic and service-level policy summary workflows
- test coverage for policy summary builders

### Improved
- shared validation helpers for policy summary inputs
- cleaner internal assembly logic for policy summary outputs
- README updated for `v0.5.0` scope and summary-helper usage
- API consistency for higher-level inventory planning workflows

## [v0.4.0] - 2026-04-16
### Added
- `TargetInventoryLevel` for calculating target inventory level from expected demand coverage and safety stock
- `TargetInventoryLevelWithServiceLevel` for calculating target inventory level using service-level-based safety stock
- `MinMaxLevelsWithServiceLevel` for calculating min/max inventory levels using a service-level-based reorder point

### Improved
- reused existing helpers to compose higher-level inventory policy functions
- expanded examples and documentation for target inventory and policy calculations

## [v0.3.0] - 2026-04-16
### Added
- `DemandDuringLeadTime` for calculating expected demand over lead time
- `StdDevDemandDuringLeadTime` for calculating demand variability over lead time

### Improved
- expanded inventory policy support with lead-time demand building blocks
- improved composability for higher-level inventory calculations
- added examples and tests for lead-time demand helpers

## [v0.2.0] - 2026-04-16
### Added
- `ZScoreForServiceLevel` for converting target service levels to z-scores
- `SafetyStockWithServiceLevel` for probabilistic safety stock calculations
- `ReorderPointWithServiceLevel` for reorder point calculations using service level targets

### Improved
- validation coverage for service-level-based inventory functions
- clearer error handling for invalid service level and negative standard deviation inputs

## [v0.1.2] - 2026-04-14
### Fixed
- updated license copyright name

## [v0.1.1] - 2026-04-14
### Added
- included the `inventory` package correctly in the tagged release
- added edge-case test coverage

## [v0.1.0] - 2026-04-14
### Added
- initial release of `scmgo`
- added `inventory` package
- added `ReorderPoint`
- added `SafetyStockBasic`
- added `EOQ`
- added `MinMaxLevels`
- added unit tests and examples
- added README and MIT license
