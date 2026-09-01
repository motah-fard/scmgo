package production

import "fmt"

func ExampleOEE() {
	result, err := OEE(OEEInput{
		PlannedProductionTime: 480,
		RunTime:               420,
		IdealCycleTime:        1.0,
		TotalCount:            400,
		GoodCount:             385,
	})
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("OEE=%.4f (A=%.4f P=%.4f Q=%.4f)\n", result.OEE, result.Availability, result.Performance, result.Quality)
	// Output: OEE=0.8021 (A=0.8750 P=0.9524 Q=0.9625)
}

func ExampleWIPFromLittlesLaw() {
	wip, err := WIPFromLittlesLaw(WIPFromLittlesLawInput{Throughput: 10, CycleTime: 5})
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("%.0f\n", wip)
	// Output: 50
}

func ExampleTaktTime() {
	takt, err := TaktTime(TaktTimeInput{AvailableProductionTime: 480, CustomerDemand: 120})
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("%.0f\n", takt)
	// Output: 4
}
