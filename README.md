# scmgo

[![CI](https://github.com/motah-fard/scmgo/actions/workflows/ci.yml/badge.svg)](https://github.com/motah-fard/scmgo/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/motah-fard/scmgo.svg)](https://pkg.go.dev/github.com/motah-fard/scmgo)
[![License](https://img.shields.io/github/license/motah-fard/scmgo?color=blue)](https://github.com/motah-fard/scmgo/blob/main/LICENSE)
[![GitHub release](https://img.shields.io/github/v/release/motah-fard/scmgo)](https://github.com/motah-fard/scmgo/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/motah-fard/scmgo)](https://goreportcard.com/report/github.com/motah-fard/scmgo)

`scmgo` is a Go library for practical supply-chain calculations, organized as one package per domain.

- **`inventory`** — reorder point, safety stock (fixed and variable lead time), EOQ (plain, quantity-discount, and production-quantity variants), fill-rate and cycle-service-level models, newsvendor, min/max levels, lead-time demand helpers, inventory ratios, and policy summary helpers (including batch helpers for SKU lists).
- **`forecast`** — demand forecasting to feed the `inventory` package's inputs: moving average, weighted moving average, simple exponential smoothing, Holt's linear trend, Croston's method for intermittent demand, and forecast accuracy metrics (MAD, MAPE, Bias, RMSE).
- **`abc`** — ABC (Pareto) classification by value concentration, XYZ classification by demand variability, and combining the two into the classic ABC-XYZ matrix.
- **`procurement`** — landed cost, purchase price variance, and total cost of ownership.
- **`production`** — OEE, Little's Law (all three directions), and takt time.
- **`logistics`** — dimensional/billable weight, freight cost allocation, vehicle utilization, and center-of-gravity facility location.

See [Roadmap](#roadmap) for planned packages.

The goal is to keep the API:

- simple
- transparent
- practical
- easy to embed in Go applications

## Stability

The `inventory` package's public API is stable as of `v1.0.0`. The `forecast`
and `abc` packages are new and, while tested to the same standard (see
[Current Scope](#current-scope)), have not yet shipped in a tagged release —
their APIs should be considered provisional until the next version tag.
Recent additions to `inventory` itself (everything under
[Beyond Cycle Service Level](#beyond-cycle-service-level)) are likewise
unreleased and could still change before the next tag.

## Current Scope

### `inventory`

- `ReorderPoint`
- `SafetyStockBasic`
- `EOQ`
- `MinMaxLevels`
- `ZScoreForServiceLevel`
- `SafetyStockWithServiceLevel`
- `ReorderPointWithServiceLevel`
- `DemandDuringLeadTime`
- `StdDevDemandDuringLeadTime`
- `TargetInventoryLevel`
- `TargetInventoryLevelWithServiceLevel`
- `MinMaxLevelsWithServiceLevel`
- `BuildPolicySummary`
- `BuildPolicySummaryWithServiceLevel`
- `BuildPolicySummaryBatch`
- `BuildPolicySummaryWithServiceLevelBatch`
- `SafetyStockWithVariableLeadTime`
- `ReorderPointWithVariableLeadTime`
- `UnitNormalLoss`
- `ExpectedFillRate`
- `FillRateSafetyStock`
- `EPQ`
- `EOQWithQuantityDiscounts`
- `Newsvendor`
- `Turnover`
- `DaysOfInventoryOnHand`
- `GMROI`

### `forecast`

- `MovingAverage`
- `WeightedMovingAverage`
- `SimpleExponentialSmoothing`
- `HoltLinearTrend`
- `Croston`
- `Accuracy`
- `HoltWinters`
- `TrackingSignal`
- `LinearTrend`
- `MASE`

### `abc`

- `Classify`
- `ClassifyVariability`
- `Combine`

### `procurement`

- `LandedCost`
- `PurchasePriceVariance`
- `TotalCostOfOwnership`

### `production`

- `OEE`
- `WIPFromLittlesLaw`
- `CycleTimeFromLittlesLaw`
- `ThroughputFromLittlesLaw`
- `TaktTime`

### `logistics`

- `DimensionalWeight`
- `BillableWeight`
- `AllocateFreightCost`
- `VehicleUtilization`
- `CenterOfGravity`

## Why scmgo

Many inventory and supply-chain calculations still live in spreadsheets, internal notes, or one-off scripts. `scmgo` provides a lightweight Go-native alternative for developers building:

- inventory tools
- supply-chain applications
- planning dashboards
- internal operations services
- educational and analytical tools

The package is intentionally small, explicit, and easy to embed.

## Installation

```bash
go get github.com/motah-fard/scmgo/inventory@latest
go get github.com/motah-fard/scmgo/forecast@latest
go get github.com/motah-fard/scmgo/abc@latest
go get github.com/motah-fard/scmgo/procurement@latest
go get github.com/motah-fard/scmgo/production@latest
go get github.com/motah-fard/scmgo/logistics@latest
```

## Packages

- `github.com/motah-fard/scmgo/inventory`
- `github.com/motah-fard/scmgo/forecast`
- `github.com/motah-fard/scmgo/abc`
- `github.com/motah-fard/scmgo/procurement`
- `github.com/motah-fard/scmgo/production`
- `github.com/motah-fard/scmgo/logistics`

## Quick Start

```go
package main

import (
	"fmt"
	"log"

	"github.com/motah-fard/scmgo/inventory"
)

func main() {
	summary, err := inventory.BuildPolicySummary(inventory.PolicySummaryInput{
		DailyDemand:      100,
		LeadTimeDays:     5,
		ReviewPeriodDays: 7,
		SafetyStockUnits: 50,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("reorder point: %.0f\n", summary.ReorderPoint)
	fmt.Printf("target level: %.0f\n", summary.TargetInventoryLevel)
}
```

## Policy Summary Helpers

These helpers provide a higher-level API for common inventory planning workflows. They are useful when you want one call to return the main planning outputs instead of assembling them manually from several lower-level functions.

### Deterministic Policy Summary

Builds a policy summary from average demand, lead time, review period, and fixed safety stock.

```go
summary, err := inventory.BuildPolicySummary(inventory.PolicySummaryInput{
	DailyDemand:      100,
	LeadTimeDays:     5,
	ReviewPeriodDays: 7,
	SafetyStockUnits: 50,
})
```

Returned fields include:

- expected demand during lead time
- safety stock
- reorder point
- target inventory level
- min level
- max level

### Service-Level Policy Summary

Builds a policy summary using service-level-based reorder point logic and demand variability.

```go
summary, err := inventory.BuildPolicySummaryWithServiceLevel(inventory.PolicySummaryServiceLevelInput{
	DailyDemand:        100,
	LeadTimeDays:       5,
	ReviewPeriodDays:   7,
	DemandStdDevPerDay: 20,
	ServiceLevel:       0.95,
})
```

This is useful for dashboards, reorder recommendations, and embedded inventory planning logic where service-level assumptions matter.

### Batch Policy Summaries

`BuildPolicySummaryBatch` and `BuildPolicySummaryWithServiceLevelBatch` run the corresponding single-item helper over a slice of inputs — useful for computing summaries across a full SKU list in one call. Unlike the single-item helpers, one invalid item does not abort the batch: each result carries its own index and error, so you can render the rest of the list and flag the bad row.

```go
results := inventory.BuildPolicySummaryBatch([]inventory.PolicySummaryInput{
	{DailyDemand: 100, LeadTimeDays: 5, ReviewPeriodDays: 7, SafetyStockUnits: 50},
	{DailyDemand: 250, LeadTimeDays: 3, ReviewPeriodDays: 7, SafetyStockUnits: 20},
})

for _, r := range results {
	if r.Err != nil {
		log.Printf("SKU row %d: %v", r.Index, r.Err)
		continue
	}
	fmt.Printf("SKU row %d: ROP=%.0f\n", r.Index, r.Summary.ReorderPoint)
}
```

## Available Functions

### Reorder Point

Calculates reorder point using average daily demand, lead time, and safety stock.

```go
rp, err := inventory.ReorderPoint(inventory.ReorderPointInput{
	AvgDailyDemand:   100,
	LeadTimeDays:     5,
	SafetyStockUnits: 50,
})
```

### Basic Safety Stock

Calculates safety stock using a max-demand and average-demand approach.

```go
ss, err := inventory.SafetyStockBasic(inventory.SafetyStockInput{
	MaxDailyDemand:  120,
	MaxLeadTimeDays: 7,
	AvgDailyDemand:  100,
	AvgLeadTimeDays: 5,
})
```

### EOQ

Calculates economic order quantity.

```go
eoq, err := inventory.EOQ(inventory.EOQInput{
	AnnualDemand:       10000,
	OrderingCost:       50,
	HoldingCostPerUnit: 2,
})
```

### Min/Max Levels

Calculates minimum and maximum inventory levels from reorder point and order quantity.

```go
levels, err := inventory.MinMaxLevels(inventory.MinMaxInput{
	ReorderPoint:  300,
	OrderQuantity: 200,
})
```

### Z-Score for Service Level

Converts a target cycle service level into a standard normal z-score.

```go
z, err := inventory.ZScoreForServiceLevel(0.95)
```

### Safety Stock with Service Level

Calculates safety stock using demand variability, lead time, and a target service level.

```go
ss, err := inventory.SafetyStockWithServiceLevel(inventory.SafetyStockWithServiceLevelInput{
	StdDevDailyDemand: 10,
	LeadTimeDays:      4,
	ServiceLevel:      0.95,
})
```

### Reorder Point with Service Level

Calculates reorder point using average demand, lead time, demand variability, and a target service level.

```go
rp, err := inventory.ReorderPointWithServiceLevel(inventory.ReorderPointWithServiceLevelInput{
	AvgDailyDemand:    50,
	LeadTimeDays:      4,
	StdDevDailyDemand: 10,
	ServiceLevel:      0.95,
})
```

### Demand During Lead Time

Calculates expected demand during lead time.

```go
d, err := inventory.DemandDuringLeadTime(inventory.DemandDuringLeadTimeInput{
	AvgDailyDemand: 100,
	LeadTimeDays:   5,
})
```

### Standard Deviation of Demand During Lead Time

Calculates the standard deviation of demand during lead time.

```go
sd, err := inventory.StdDevDemandDuringLeadTime(inventory.StdDevDemandDuringLeadTimeInput{
	StdDevDailyDemand: 10,
	LeadTimeDays:      4,
})
```

### Target Inventory Level

Calculates target inventory level from expected demand coverage and safety stock.

```go
level, err := inventory.TargetInventoryLevel(inventory.TargetInventoryLevelInput{
	ExpectedDemandDuringLeadTime: 500,
	SafetyStockUnits:             50,
})
```

### Target Inventory Level with Service Level

Calculates target inventory level using average demand, lead time, demand variability, and a target service level.

```go
level, err := inventory.TargetInventoryLevelWithServiceLevel(inventory.TargetInventoryLevelWithServiceLevelInput{
	AvgDailyDemand:    50,
	LeadTimeDays:      4,
	StdDevDailyDemand: 10,
	ServiceLevel:      0.95,
})
```

### Min/Max Levels with Service Level

Calculates min/max inventory levels using a service-level-based reorder point and a fixed order quantity.

```go
levels, err := inventory.MinMaxLevelsWithServiceLevel(inventory.MinMaxLevelsWithServiceLevelInput{
	AvgDailyDemand:    50,
	LeadTimeDays:      4,
	StdDevDailyDemand: 10,
	ServiceLevel:      0.95,
	OrderQuantity:     200,
})
```

## Beyond Cycle Service Level

A few functions extend `inventory` past the fixed-lead-time, cycle-service-level
model used elsewhere in this section. Full parameter details are in each
function's doc comment and runnable `Example` on pkg.go.dev.

```go
// Safety stock and reorder point when lead time varies too, not just demand.
ss, err := inventory.SafetyStockWithVariableLeadTime(inventory.SafetyStockWithVariableLeadTimeInput{
	AvgDailyDemand: 100, StdDevDailyDemand: 10,
	AvgLeadTimeDays: 5, StdDevLeadTimeDays: 1,
	ServiceLevel: 0.95,
})

// Fill rate (P2, % of demand met from stock) instead of cycle service level.
ss, err = inventory.FillRateSafetyStock(inventory.FillRateSafetyStockInput{
	TargetFillRate: 0.98, StdDevDemandDuringLeadTime: 50, OrderQuantity: 200,
})

// Economic order quantity across a price-break schedule.
result, err := inventory.EOQWithQuantityDiscounts(inventory.EOQWithQuantityDiscountsInput{
	AnnualDemand: 10000, OrderingCost: 20, HoldingCostRate: 0.20,
	Tiers: []inventory.QuantityDiscountTier{
		{MinQuantity: 1, UnitPrice: 5.00},
		{MinQuantity: 500, UnitPrice: 4.50},
		{MinQuantity: 1000, UnitPrice: 3.90},
	},
})

// Single-period optimal order quantity (newsvendor model).
nv, err := inventory.Newsvendor(inventory.NewsvendorInput{
	MeanDemand: 500, StdDevDemand: 100,
	UnderageCostPerUnit: 18, OverageCostPerUnit: 7,
})
```

## Forecasting

The `forecast` package estimates demand from historical data — output that
feeds directly into `inventory`'s `AvgDailyDemand` / `StdDevDailyDemand`
inputs.

```go
import "github.com/motah-fard/scmgo/forecast"

result, err := forecast.SimpleExponentialSmoothing(forecast.SimpleExponentialSmoothingInput{
	History: []float64{100, 120, 110, 130, 125},
	Alpha:   0.3,
})
```

### Moving Average

```go
f, err := forecast.MovingAverage(forecast.MovingAverageInput{
	History: []float64{100, 120, 110, 130, 125},
	Periods: 3,
})
```

### Weighted Moving Average

```go
f, err := forecast.WeightedMovingAverage(forecast.WeightedMovingAverageInput{
	History: []float64{100, 120, 110, 130, 125},
	Weights: []float64{0.1, 0.3, 0.6}, // oldest -> newest, must sum to 1
})
```

### Simple Exponential Smoothing

```go
result, err := forecast.SimpleExponentialSmoothing(forecast.SimpleExponentialSmoothingInput{
	History: []float64{100, 120, 110, 130, 125},
	Alpha:   0.3,
})
```

### Holt's Linear Trend (Double Exponential Smoothing)

```go
result, err := forecast.HoltLinearTrend(forecast.HoltLinearTrendInput{
	History:      []float64{100, 120, 110, 130, 125},
	Alpha:        0.3,
	Beta:         0.2,
	PeriodsAhead: 3,
})
```

### Croston's Method (Intermittent Demand)

```go
result, err := forecast.Croston(forecast.CrostonInput{
	History: []float64{0, 0, 5, 0, 0, 0, 3, 0, 4, 0},
	Alpha:   0.2,
})
```

### Forecast Accuracy

```go
result, err := forecast.Accuracy(forecast.AccuracyInput{
	Actual:   []float64{100, 110, 95, 130},
	Forecast: []float64{90, 115, 100, 120},
})
// result.MAD, result.MAPE, result.Bias, result.RMSE
```

### Holt-Winters (Seasonal)

```go
result, err := forecast.HoltWinters(forecast.HoltWintersInput{
	History:      []float64{100, 120, 90, 110, 105, 125, 95, 115}, // 2 full seasons
	Alpha:        0.3, Beta: 0.1, Gamma: 0.2,
	SeasonLength: 4, // e.g. quarterly seasonality
	PeriodsAhead: 1,
})
```

### Tracking Signal

```go
ts, err := forecast.TrackingSignal(forecast.TrackingSignalInput{
	Actual:   []float64{100, 110, 95, 130},
	Forecast: []float64{90, 115, 100, 120},
})
// ts[i] is the running tracking signal through period i; a common rule of
// thumb flags |ts[i]| outside roughly 4-8 as the forecast drifting out of control.
```

### Linear Trend

```go
result, err := forecast.LinearTrend(forecast.LinearTrendInput{
	History:      []float64{100, 105, 108, 115, 120},
	PeriodsAhead: 3,
})
```

### MASE (Mean Absolute Scaled Error)

```go
mase, err := forecast.MASE(forecast.MASEInput{
	TrainingHistory: []float64{100, 110, 105, 120, 115, 130}, // scales the error
	Actual:          []float64{125, 135},
	Forecast:        []float64{120, 140},
})
// mase < 1 means the forecast beats a naive one-step-ahead benchmark
```

## Classification

The `abc` package prioritizes which SKUs deserve tighter inventory control:
ABC classification by value concentration, XYZ classification by demand
variability, and combining both into the classic ABC-XYZ matrix.

```go
import "github.com/motah-fard/scmgo/abc"

results, err := abc.Classify(abc.ClassifyInput{
	Items: []abc.Item{
		{ID: "sku-1", Value: 8000},
		{ID: "sku-2", Value: 2000},
	},
	AThreshold: 0.80, // cumulative value cutoff for class A
	BThreshold: 0.95, // cumulative value cutoff for class B
})
// results[i].Class is "A", "B", or "C"; results[i].CumulativePercent is
// the Pareto curve value at that item -- there's no separate "Pareto
// analysis" function since this cumulative-percent field is exactly that.
```

`Classify`'s result stays in the same order as the input `Items` (not
resorted by value), so it's index-aligned with `ClassifyVariability`'s
result for the same item set and easy to feed into `Combine`.

## Design Principles

`scmgo` is intentionally designed to be:

- small and focused
- explicit rather than clever
- easy to test
- easy to read
- suitable for both production use and teaching

## Error Handling

Every package validates inputs and returns explicit sentinel errors for invalid values instead of panicking — e.g. negative demand, negative lead time, invalid service level, invalid smoothing constant, mismatched series lengths, and non-finite (`NaN`/`Inf`) values. See each package's `errors.go` for the full list. This keeps behavior predictable and makes the library easier to integrate into larger systems.

Validation always checks finiteness first: a plain `v < 0` check is a silent no-op against `NaN` in IEEE 754 (any comparison against `NaN` is `false`), so every function checks `NaN`/`Inf` explicitly before any range check — otherwise a bad upstream value (e.g. from a division by zero elsewhere in a caller's pipeline) would silently produce a poisoned `NaN`/`Inf` result with a `nil` error instead of failing loudly. Both packages have fuzz targets (`go test ./... -fuzz=...`) that assert this — see [CONTRIBUTING.md](CONTRIBUTING.md).

See [SECURITY.md](SECURITY.md) for the project's security scope and how to report an issue.

## Assumptions

**`inventory`**

- Input units must be consistent
- If demand is measured per day, lead time should also be in days
- EOQ uses the classic Wilson EOQ formula
- Service-level calculations assume a normal approximation
- `SafetyStockBasic` uses a simple max/average demand and lead-time formula
- `StdDevDemandDuringLeadTime` assumes independent daily demand variability across lead-time periods
- Policy summary helpers combine lead-time coverage, review-period coverage, safety stock, reorder point, target inventory level, and min/max outputs into a single result
- `SafetyStockWithVariableLeadTime` reduces to `SafetyStockWithServiceLevel` when lead-time variability is zero
- `FillRateSafetyStock`'s result can be negative for a low enough target fill rate — that's not a bug; `ExpectedFillRate` accepts a negative safety stock accordingly
- `EOQWithQuantityDiscounts` evaluates each price tier's own EOQ clamped to that tier's valid range, and returns whichever tier/quantity minimizes total annual cost

**`forecast`**

- History series are chronological, oldest value first, with non-negative demand
- Smoothing constants (Alpha, Beta) must be in `(0, 1]`
- Croston treats zero as "no demand" and requires at least one non-zero period
- `Accuracy`'s MAPE excludes periods where the actual value is zero, and is `NaN` if every actual value is zero
- `HoltWinters` uses additive seasonality only, needs at least two full seasons of history, and documents its exact initialization convention in its doc comment (there's more than one in the literature — verify against your own reference if you need to match a specific one)
- `TrackingSignal` returns one value per period (for monitoring drift over time), not a single summary value
- `MASE`'s `TrainingHistory` must not be perfectly constant (see `ErrZeroNaiveMAE`)

**`abc`**

- `Classify` returns results in input order, not resorted by value, so it stays index-aligned with `ClassifyVariability` for the same item set
- `Classify` requires item values to be non-negative and sum to more than zero (cumulative percent is undefined for an all-zero set)
- `ClassifyVariability` requires `MeanDemand` to be strictly positive (the coefficient of variation is undefined at mean zero)
- `Combine` is an inner join on ID: an item present in only one input is omitted, not defaulted

Full assumptions and exclusions are documented in each package's `doc.go`.

## Roadmap

Planned packages, each following the same Input-struct/validate/sentinel-error
pattern as `inventory`, `forecast`, and `abc` (see [CONTRIBUTING.md](CONTRIBUTING.md)).
Not yet started — contributions and formula proposals (with a citation) are
welcome via issues.

- **`procurement`** — landed cost, purchase price variance, total cost of ownership
- **`production`** — BOM explosion, lot sizing, OEE, Little's Law
- **`logistics`** — freight cost allocation, dimensional weight, center-of-gravity facility location
- **`warehouse`** — storage/cube utilization, pick rate, slotting
- **`quality`** — DPMO, process capability (Cp/Cpk), cost of quality
- **`finance`** — cash-to-cash cycle time, perfect order rate

**Explicitly not planned:** a general multi-echelon safety stock/allocation
package. The guaranteed-service multi-echelon model needs network topology
and NP-hard optimization over a DAG of stocking locations — a fundamentally
different kind of problem from the closed-form/numerically-inverted formulas
elsewhere in this library, not an extension of them. `EOQ` with quantity
discounts and fill-rate safety stock (both closed-form/numerical, not
network problems) already shipped in `inventory` rather than waiting for a
`procurement`/`fillrate` package.

## Versioning

This project follows semantic versioning.

- `v0.1.x` focused on core deterministic inventory formulas
- `v0.2.x` added service-level-based inventory calculations
- `v0.3.x` added lead-time demand and variability helpers
- `v0.4.0` added target inventory level and service-level policy helpers
- `v0.5.0` added policy summary helpers and improved API consistency for inventory planning workflows
- `v0.6.0` focused on documentation tightening, package consistency, and API stabilization ahead of `v1.0.0`
- `v1.0.0` is the first stable release of the `inventory` package
- `Unreleased` adds the `forecast` and `abc` packages, several `inventory` extensions (variable-lead-time and fill-rate safety stock, EPQ, quantity-discount EOQ, newsvendor, inventory ratios), several `forecast` extensions (Holt-Winters, tracking signal, linear trend, MASE), and batch policy-summary helpers; see [CHANGELOG.md](CHANGELOG.md)

## Documentation

- Go package docs: [pkg.go.dev/github.com/motah-fard/scmgo/inventory](https://pkg.go.dev/github.com/motah-fard/scmgo/inventory), [pkg.go.dev/github.com/motah-fard/scmgo/forecast](https://pkg.go.dev/github.com/motah-fard/scmgo/forecast), [pkg.go.dev/github.com/motah-fard/scmgo/abc](https://pkg.go.dev/github.com/motah-fard/scmgo/abc)
- Releases: [github.com/motah-fard/scmgo/releases](https://github.com/motah-fard/scmgo/releases)
- Contributing: [CONTRIBUTING.md](CONTRIBUTING.md)
- Security policy: [SECURITY.md](SECURITY.md)

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
