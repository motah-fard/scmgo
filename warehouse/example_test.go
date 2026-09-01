package warehouse

import "fmt"

func ExampleStorageUtilization() {
	util, err := StorageUtilization(StorageUtilizationInput{UsedSpace: 800, TotalSpace: 1000})
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("%.0f%%\n", util*100)
	// Output: 80%
}

func ExampleCubeUtilization() {
	util, err := CubeUtilization(CubeUtilizationInput{UsedVolume: 4500, TotalVolume: 6000})
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("%.0f%%\n", util*100)
	// Output: 75%
}

func ExamplePickRate() {
	rate, err := PickRate(PickRateInput{TotalPicks: 480, TotalHours: 8})
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("%.0f\n", rate)
	// Output: 60
}
