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
