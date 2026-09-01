package finance

import "fmt"

func ExampleDSO() {
	dso, err := DSO(DSOInput{AccountsReceivable: 150000, TotalCreditSales: 1800000, DaysInPeriod: 365})
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("%.2f\n", dso)
	// Output: 30.42
}

func ExampleDPO() {
	dpo, err := DPO(DPOInput{AccountsPayable: 120000, COGS: 1200000, DaysInPeriod: 365})
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("%.2f\n", dpo)
	// Output: 36.50
}

func ExampleCashToCashCycleTime() {
	c2c, err := CashToCashCycleTime(CashToCashCycleTimeInput{
		DIO: 60.83,
		DSO: 30.42,
		DPO: 36.5,
	})
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("%.2f\n", c2c)
	// Output: 54.75
}

func ExamplePerfectOrderRate() {
	por, err := PerfectOrderRate(PerfectOrderRateInput{
		OnTimeRate:                0.95,
		CompleteRate:              0.97,
		DamageFreeRate:            0.99,
		AccurateDocumentationRate: 0.98,
	})
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("%.4f\n", por)
	// Output: 0.8940
}
