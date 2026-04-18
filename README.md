# scmgo

[![Go Reference](https://pkg.go.dev/badge/github.com/motah-fard/scmgo/inventory.svg)](https://pkg.go.dev/github.com/motah-fard/scmgo/inventory)
[![License](https://img.shields.io/github/license/motah-fard/scmgo?color=blue)](https://github.com/motah-fard/scmgo/blob/main/LICENSE)
[![GitHub release](https://img.shields.io/github/v/release/motah-fard/scmgo)](https://github.com/motah-fard/scmgo/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/motah-fard/scmgo)](https://goreportcard.com/report/github.com/motah-fard/scmgo)

`scmgo` is a Go library for practical inventory and supply-chain calculations.

The `inventory` package provides clear and reusable functions for common inventory policy calculations such as reorder point, safety stock, EOQ, min/max levels, lead-time demand helpers, service-level-based threshold planning, and policy summary helpers.

The goal is to keep the API:

- simple
- transparent
- practical
- easy to embed in Go applications

## Stability

The `inventory` package is released as `v1.0.0` and is intended to provide a stable public API for practical inventory policy calculations.

## Current Scope

As of `v1.0.0`, the `inventory` package includes:

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
```

## Package

Current package:

- `github.com/motah-fard/scmgo/inventory`

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

## Design Principles

`scmgo` is intentionally designed to be:

- small and focused
- explicit rather than clever
- easy to test
- easy to read
- suitable for both production use and teaching

## Error Handling

The package validates inputs and returns explicit errors for invalid values such as:

- negative demand
- negative expected demand
- negative lead time
- negative review period
- negative safety stock
- invalid service level
- negative standard deviation
- invalid holding cost

This keeps behavior predictable and makes the library easier to integrate into larger systems.

## Assumptions

- Input units must be consistent
- If demand is measured per day, lead time should also be in days
- EOQ uses the classic Wilson EOQ formula
- Service-level calculations assume a normal approximation
- `SafetyStockBasic` uses a simple max/average demand and lead-time formula
- `StdDevDemandDuringLeadTime` assumes independent daily demand variability across lead-time periods
- Policy summary helpers combine lead-time coverage, review-period coverage, safety stock, reorder point, target inventory level, and min/max outputs into a single result

## Versioning

This project follows semantic versioning.

- `v0.1.x` focused on core deterministic inventory formulas
- `v0.2.x` added service-level-based inventory calculations
- `v0.3.x` added lead-time demand and variability helpers
- `v0.4.0` added target inventory level and service-level policy helpers
- `v0.5.0` added policy summary helpers and improved API consistency for inventory planning workflows
- `v0.6.0` focused on documentation tightening, package consistency, and API stabilization ahead of `v1.0.0`
- `v1.0.0` is the first stable release of the `inventory` package

## Documentation

- Go package docs: [pkg.go.dev/github.com/motah-fard/scmgo/inventory](https://pkg.go.dev/github.com/motah-fard/scmgo/inventory)
- Releases: [github.com/motah-fard/scmgo/releases](https://github.com/motah-fard/scmgo/releases)

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
