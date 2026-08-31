# scmgo

[![CI](https://github.com/motah-fard/scmgo/actions/workflows/ci.yml/badge.svg)](https://github.com/motah-fard/scmgo/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/motah-fard/scmgo.svg)](https://pkg.go.dev/github.com/motah-fard/scmgo)
[![License](https://img.shields.io/github/license/motah-fard/scmgo?color=blue)](https://github.com/motah-fard/scmgo/blob/main/LICENSE)
[![GitHub release](https://img.shields.io/github/v/release/motah-fard/scmgo)](https://github.com/motah-fard/scmgo/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/motah-fard/scmgo)](https://goreportcard.com/report/github.com/motah-fard/scmgo)

`scmgo` is a Go library for practical supply-chain calculations, organized as one package per domain.

- **`inventory`** — reorder point, safety stock, EOQ, min/max levels, lead-time demand helpers, service-level-based threshold planning, and policy summary helpers (including batch helpers for SKU lists).
- **`forecast`** — demand forecasting to feed the `inventory` package's inputs: moving average, weighted moving average, simple exponential smoothing, Holt's linear trend, Croston's method for intermittent demand, and forecast accuracy metrics (MAD, MAPE, Bias, RMSE).

See [Roadmap](#roadmap) for planned packages.

The goal is to keep the API:

- simple
- transparent
- practical
- easy to embed in Go applications

## Stability

The `inventory` package's public API is stable as of `v1.0.0`. The `forecast`
package is new and, while tested to the same standard (see [Current Scope](#current-scope)),
has not yet shipped in a tagged release — its API should be considered
provisional until the next version tag.

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
```

## Packages

- `github.com/motah-fard/scmgo/inventory`
- `github.com/motah-fard/scmgo/forecast`

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

Full assumptions and exclusions are documented in each package's `doc.go`.

## Roadmap

Planned packages, each following the same Input-struct/validate/sentinel-error
pattern as `inventory` and `forecast` (see [CONTRIBUTING.md](CONTRIBUTING.md)).
Not yet started — contributions and formula proposals (with a citation) are
welcome via issues.

- **`abc`** — ABC/XYZ classification and Pareto analysis for prioritizing inventory attention
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
- `Unreleased` adds the `forecast` package and batch policy-summary helpers; see [CHANGELOG.md](CHANGELOG.md)

## Documentation

- Go package docs: [pkg.go.dev/github.com/motah-fard/scmgo/inventory](https://pkg.go.dev/github.com/motah-fard/scmgo/inventory), [pkg.go.dev/github.com/motah-fard/scmgo/forecast](https://pkg.go.dev/github.com/motah-fard/scmgo/forecast)
- Releases: [github.com/motah-fard/scmgo/releases](https://github.com/motah-fard/scmgo/releases)
- Contributing: [CONTRIBUTING.md](CONTRIBUTING.md)
- Security policy: [SECURITY.md](SECURITY.md)

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
