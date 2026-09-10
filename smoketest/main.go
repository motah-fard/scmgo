// Command smoketest is not part of scmgo. It is a separate module (see
// go.mod's replace directive) that imports scmgo the way a real external
// consumer would: by its published module path, not a relative path within
// the same module. That catches a class of bug the in-module examples and
// tests can't: an accidentally-unexported symbol, an internal package
// leaking into a public one, or anything else that only breaks once scmgo
// crosses a module boundary.
//
// It calls one exported function per public package and checks the result
// against a value computed independently, so a wrong answer fails the same
// as a build error.
package main

import (
	"fmt"
	"math"
	"os"

	"github.com/motah-fard/scmgo/abc"
	"github.com/motah-fard/scmgo/finance"
	"github.com/motah-fard/scmgo/forecast"
	"github.com/motah-fard/scmgo/inventory"
	"github.com/motah-fard/scmgo/logistics"
	"github.com/motah-fard/scmgo/procurement"
	"github.com/motah-fard/scmgo/production"
	"github.com/motah-fard/scmgo/quality"
	"github.com/motah-fard/scmgo/warehouse"
)

func main() {
	failed := false

	check := func(pkg string, err error) {
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: unexpected error: %v\n", pkg, err)
			failed = true
		}
	}

	almostEqual := func(pkg string, got, want float64) {
		const tolerance = 1e-9
		if math.Abs(got-want) > tolerance {
			fmt.Fprintf(os.Stderr, "%s: got %v, want %v\n", pkg, got, want)
			failed = true
		}
	}

	classified, err := abc.Classify(abc.ClassifyInput{
		Items:      []abc.Item{{ID: "sku-1", Value: 80}, {ID: "sku-2", Value: 20}},
		AThreshold: 0.8,
		BThreshold: 0.95,
	})
	check("abc", err)
	if len(classified) != 2 {
		fmt.Fprintf(os.Stderr, "abc: got %d classified items, want 2\n", len(classified))
		failed = true
	}

	c2c, err := finance.CashToCashCycleTime(finance.CashToCashCycleTimeInput{DIO: 45, DSO: 30, DPO: 20})
	check("finance", err)
	almostEqual("finance", c2c, 55)

	acc, err := forecast.Accuracy(forecast.AccuracyInput{
		Actual:   []float64{100, 100},
		Forecast: []float64{90, 110},
	})
	check("forecast", err)
	almostEqual("forecast", acc.MAD, 10)

	doh, err := inventory.DaysOfInventoryOnHand(inventory.DaysOfInventoryOnHandInput{
		AverageInventoryValue: 500,
		COGS:                  3650,
		DaysInPeriod:          365,
	})
	check("inventory", err)
	almostEqual("inventory", doh, 50)

	freight, err := logistics.AllocateFreightCost(logistics.AllocateFreightCostInput{
		TotalFreightCost: 100,
		Items: []logistics.FreightAllocationItem{
			{ID: "sku-1", Weight: 75},
			{ID: "sku-2", Weight: 25},
		},
	})
	check("logistics", err)
	if len(freight) == 2 {
		almostEqual("logistics", freight[0].AllocatedCost, 75)
	}

	landed, err := procurement.LandedCost(procurement.LandedCostInput{
		Components: []procurement.CostComponent{
			{Label: "Unit cost", Amount: 100},
			{Label: "Freight", Amount: 15},
		},
	})
	check("procurement", err)
	almostEqual("procurement", landed.Total, 115)

	wip, err := production.WIPFromLittlesLaw(production.WIPFromLittlesLawInput{Throughput: 10, CycleTime: 5})
	check("production", err)
	almostEqual("production", wip, 50)

	coq, err := quality.CostOfQuality(quality.CostOfQualityInput{
		PreventionCost:      10,
		AppraisalCost:       20,
		InternalFailureCost: 30,
		ExternalFailureCost: 40,
	})
	check("quality", err)
	almostEqual("quality", coq.Total, 100)

	cube, err := warehouse.CubeUtilization(warehouse.CubeUtilizationInput{UsedVolume: 750, TotalVolume: 1000})
	check("warehouse", err)
	almostEqual("warehouse", cube, 0.75)

	if failed {
		os.Exit(1)
	}
	fmt.Println("smoketest: all packages imported and called successfully as an external consumer")
}
