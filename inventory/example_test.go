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
