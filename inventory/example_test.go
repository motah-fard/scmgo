package inventory

import (
	"fmt"
)

func ExampleReorderPoint() {
	rp, err := ReorderPoint(ReorderPointInput{
		AvgDailyDemand:   100,
		LeadTimeDays:     5,
		SafetyStockUnits: 50,
	})
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("%.0f\n", rp)
	// Output: 550
}

func ExampleSafetyStockBasic() {
	ss, err := SafetyStockBasic(SafetyStockInput{
		MaxDailyDemand:  120,
		MaxLeadTimeDays: 7,
		AvgDailyDemand:  100,
		AvgLeadTimeDays: 5,
	})
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("%.0f\n", ss)
	// Output: 340
}

func ExampleEOQ() {
	eoq, err := EOQ(EOQInput{
		AnnualDemand:       10000,
		OrderingCost:       50,
		HoldingCostPerUnit: 2,
	})
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("%.2f\n", eoq)
	// Output: 707.11
}

func ExampleMinMaxLevels() {
	levels, err := MinMaxLevels(MinMaxInput{
		ReorderPoint:  300,
		OrderQuantity: 200,
	})
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("min=%.0f max=%.0f\n", levels.Min, levels.Max)
	// Output: min=300 max=500
}

func ExampleZScoreForServiceLevel() {
	z, err := ZScoreForServiceLevel(0.95)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("%.4f\n", z)
	// Output: 1.6449
}

func ExampleSafetyStockWithServiceLevel() {
	ss, err := SafetyStockWithServiceLevel(SafetyStockWithServiceLevelInput{
		StdDevDailyDemand: 10,
		LeadTimeDays:      4,
		ServiceLevel:      0.95,
	})
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("%.2f\n", ss)
	// Output: 32.90
}

func ExampleReorderPointWithServiceLevel() {
	rp, err := ReorderPointWithServiceLevel(ReorderPointWithServiceLevelInput{
		AvgDailyDemand:    50,
		LeadTimeDays:      4,
		StdDevDailyDemand: 10,
		ServiceLevel:      0.95,
	})
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("%.2f\n", rp)
	// Output: 232.90
}

func ExampleDemandDuringLeadTime() {
	d, err := DemandDuringLeadTime(DemandDuringLeadTimeInput{
		AvgDailyDemand: 100,
		LeadTimeDays:   5,
	})
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("%.0f\n", d)
	// Output: 500
}

func ExampleStdDevDemandDuringLeadTime() {
	sd, err := StdDevDemandDuringLeadTime(StdDevDemandDuringLeadTimeInput{
		StdDevDailyDemand: 10,
		LeadTimeDays:      4,
	})
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("%.0f\n", sd)
	// Output: 20
}
func ExampleTargetInventoryLevel() {
	level, err := TargetInventoryLevel(TargetInventoryLevelInput{
		ExpectedDemandDuringLeadTime: 500,
		SafetyStockUnits:             50,
	})
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("%.0f\n", level)
	// Output: 550
}
func ExampleTargetInventoryLevelWithServiceLevel() {
	level, err := TargetInventoryLevelWithServiceLevel(TargetInventoryLevelWithServiceLevelInput{
		AvgDailyDemand:    50,
		LeadTimeDays:      4,
		StdDevDailyDemand: 10,
		ServiceLevel:      0.95,
	})
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("%.2f\n", level)
	// Output: 232.90
}

func ExampleMinMaxLevelsWithServiceLevel() {
	levels, err := MinMaxLevelsWithServiceLevel(MinMaxLevelsWithServiceLevelInput{
		AvgDailyDemand:    50,
		LeadTimeDays:      4,
		StdDevDailyDemand: 10,
		ServiceLevel:      0.95,
		OrderQuantity:     200,
	})
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("min=%.2f max=%.2f\n", levels.Min, levels.Max)
	// Output: min=232.90 max=432.90
}

func ExampleBuildPolicySummary() {
	summary, err := BuildPolicySummary(PolicySummaryInput{
		DailyDemand:      100,
		LeadTimeDays:     5,
		ReviewPeriodDays: 7,
		SafetyStockUnits: 50,
	})
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("ROP: %.0f, Target: %.0f\n", summary.ReorderPoint, summary.TargetInventoryLevel)
	// Output:
	// ROP: 550, Target: 1250
}

func ExampleBuildPolicySummaryWithServiceLevel() {
	summary, err := BuildPolicySummaryWithServiceLevel(PolicySummaryServiceLevelInput{
		DailyDemand:        100,
		LeadTimeDays:       5,
		ReviewPeriodDays:   7,
		DemandStdDevPerDay: 20,
		ServiceLevel:       0.95,
	})
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf(
		"Expected lead-time demand: %.0f, Safety stock: %.2f, ROP: %.2f, Target: %.2f, Min: %.2f, Max: %.2f\n",
		summary.ExpectedDemandDuringLeadTime,
		summary.SafetyStockUnits,
		summary.ReorderPoint,
		summary.TargetInventoryLevel,
		summary.MinLevel,
		summary.MaxLevel,
	)
	// Output:
	// Expected lead-time demand: 500, Safety stock: 73.56, ROP: 573.56, Target: 1273.56, Min: 573.56, Max: 1273.56
}

func ExampleBuildPolicySummaryBatch() {
	results := BuildPolicySummaryBatch([]PolicySummaryInput{
		{DailyDemand: 100, LeadTimeDays: 5, ReviewPeriodDays: 7, SafetyStockUnits: 50},
		{DailyDemand: -1, LeadTimeDays: 5, ReviewPeriodDays: 7, SafetyStockUnits: 50},
	})

	for _, r := range results {
		if r.Err != nil {
			fmt.Printf("item %d: error: %v\n", r.Index, r.Err)
			continue
		}
		fmt.Printf("item %d: ROP=%.0f\n", r.Index, r.Summary.ReorderPoint)
	}
	// Output:
	// item 0: ROP=550
	// item 1: error: invalid policy summary input
	// demand cannot be negative
}

func ExampleSafetyStockWithVariableLeadTime() {
	ss, err := SafetyStockWithVariableLeadTime(SafetyStockWithVariableLeadTimeInput{
		AvgDailyDemand:     100,
		StdDevDailyDemand:  10,
		AvgLeadTimeDays:    5,
		StdDevLeadTimeDays: 1,
		ServiceLevel:       0.95,
	})
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("%.2f\n", ss)
	// Output: 168.55
}

func ExampleReorderPointWithVariableLeadTime() {
	rp, err := ReorderPointWithVariableLeadTime(ReorderPointWithVariableLeadTimeInput{
		AvgDailyDemand:     100,
		StdDevDailyDemand:  10,
		AvgLeadTimeDays:    5,
		StdDevLeadTimeDays: 1,
		ServiceLevel:       0.95,
	})
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("%.2f\n", rp)
	// Output: 668.55
}

func ExampleUnitNormalLoss() {
	loss, err := UnitNormalLoss(1)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("%.4f\n", loss)
	// Output: 0.0833
}

func ExampleExpectedFillRate() {
	fillRate, err := ExpectedFillRate(ExpectedFillRateInput{
		SafetyStockUnits:           50,
		StdDevDemandDuringLeadTime: 50,
		OrderQuantity:              200,
	})
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("%.4f\n", fillRate)
	// Output: 0.9792
}

func ExampleFillRateSafetyStock() {
	ss, err := FillRateSafetyStock(FillRateSafetyStockInput{
		TargetFillRate:             0.98,
		StdDevDemandDuringLeadTime: 50,
		OrderQuantity:              200,
	})
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("%.2f\n", ss)
	// Output: 51.06
}

func ExampleEPQ() {
	epq, err := EPQ(EPQInput{
		AnnualDemand:         10000,
		SetupCost:            50,
		HoldingCostPerUnit:   2,
		AnnualProductionRate: 40000,
	})
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("%.2f\n", epq)
	// Output: 816.50
}

func ExampleEOQWithQuantityDiscounts() {
	result, err := EOQWithQuantityDiscounts(EOQWithQuantityDiscountsInput{
		AnnualDemand:    10000,
		OrderingCost:    20,
		HoldingCostRate: 0.20,
		Tiers: []QuantityDiscountTier{
			{MinQuantity: 1, UnitPrice: 5.00},
			{MinQuantity: 500, UnitPrice: 4.50},
			{MinQuantity: 1000, UnitPrice: 3.90},
		},
	})
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("qty=%.0f price=%.2f totalCost=%.2f\n", result.OrderQuantity, result.UnitPrice, result.TotalAnnualCost)
	// Output: qty=1000 price=3.90 totalCost=39590.00
}

func ExampleNewsvendor() {
	result, err := Newsvendor(NewsvendorInput{
		MeanDemand:          500,
		StdDevDemand:        100,
		UnderageCostPerUnit: 18,
		OverageCostPerUnit:  7,
	})
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("Q=%.2f criticalRatio=%.2f\n", result.OrderQuantity, result.CriticalRatio)
	// Output: Q=558.28 criticalRatio=0.72
}

func ExampleTurnover() {
	turnover, err := Turnover(TurnoverInput{
		COGS:                  120000,
		AverageInventoryValue: 20000,
	})
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("%.1f\n", turnover)
	// Output: 6.0
}

func ExampleDaysOfInventoryOnHand() {
	days, err := DaysOfInventoryOnHand(DaysOfInventoryOnHandInput{
		AverageInventoryValue: 20000,
		COGS:                  120000,
		DaysInPeriod:          365,
	})
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("%.1f\n", days)
	// Output: 60.8
}

func ExampleGMROI() {
	gmroi, err := GMROI(GMROIInput{
		GrossMargin:          50000,
		AverageInventoryCost: 20000,
	})
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("%.2f\n", gmroi)
	// Output: 2.50
}

func ExampleEOI() {
	eoi, err := EOI(EOIInput{
		AnnualDemand:       10000,
		OrderingCost:       50,
		HoldingCostPerUnit: 2,
		DaysPerYear:        365,
	})
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("%.2f\n", eoi)
	// Output: 25.81
}

func ExampleSafetyTime() {
	safetyTime, err := SafetyTime(SafetyTimeInput{SafetyStockUnits: 50, AvgDailyDemand: 20})
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("%.1f\n", safetyTime)
	// Output: 2.5
}
