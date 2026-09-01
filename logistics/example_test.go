package logistics

import "fmt"

func ExampleDimensionalWeight() {
	dimWeight, err := DimensionalWeight(DimensionalWeightInput{
		Length: 20, Width: 15, Height: 10, DimFactor: 139,
	})
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("%.2f\n", dimWeight)
	// Output: 21.58
}

func ExampleAllocateFreightCost() {
	results, err := AllocateFreightCost(AllocateFreightCostInput{
		TotalFreightCost: 1000,
		Items: []FreightAllocationItem{
			{ID: "a", Weight: 100},
			{ID: "b", Weight: 300},
			{ID: "c", Weight: 600},
		},
	})
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	for _, r := range results {
		fmt.Printf("%s: %.2f\n", r.ID, r.AllocatedCost)
	}
	// Output:
	// a: 100.00
	// b: 300.00
	// c: 600.00
}

func ExampleVehicleUtilization() {
	util, err := VehicleUtilization(VehicleUtilizationInput{ActualLoad: 8000, Capacity: 10000})
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("%.0f%%\n", util*100)
	// Output: 80%
}

func ExampleCenterOfGravity() {
	result, err := CenterOfGravity(CenterOfGravityInput{
		Locations: []LocationDemand{
			{X: 10, Y: 20, Demand: 500},
			{X: 30, Y: 40, Demand: 300},
			{X: 50, Y: 10, Demand: 200},
		},
	})
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("(%.0f, %.0f)\n", result.X, result.Y)
	// Output: (24, 24)
}
