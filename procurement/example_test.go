package procurement

import "fmt"

func ExampleLandedCost() {
	result, err := LandedCost(LandedCostInput{
		Components: []CostComponent{
			{Label: "Unit Cost", Amount: 100},
			{Label: "Freight", Amount: 15},
			{Label: "Duty", Amount: 8},
			{Label: "Insurance", Amount: 2},
		},
	})
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("%.2f\n", result.Total)
	// Output: 125.00
}

func ExampleTotalCostOfOwnership() {
	result, err := TotalCostOfOwnership(TotalCostOfOwnershipInput{
		Components: []CostComponent{
			{Label: "Acquisition", Amount: 50000},
			{Label: "Operating (5yr)", Amount: 20000},
			{Label: "Maintenance (5yr)", Amount: 8000},
			{Label: "Resale Value", Amount: -12000},
		},
	})
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("%.2f\n", result.Total)
	// Output: 66000.00
}

func ExamplePurchasePriceVariance() {
	ppv, err := PurchasePriceVariance(PurchasePriceVarianceInput{
		StandardCost: 10,
		ActualCost:   9,
		Quantity:     1000,
	})
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("%.2f\n", ppv)
	// Output: 1000.00
}
