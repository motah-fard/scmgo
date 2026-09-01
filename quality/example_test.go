package quality

import "fmt"

func ExampleDPMO() {
	dpmo, err := DPMO(DPMOInput{NumberOfDefects: 45, NumberOfUnits: 1000, OpportunitiesPerUnit: 5})
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("%.0f\n", dpmo)
	// Output: 9000
}

func ExampleCp() {
	cp, err := Cp(CpInput{USL: 10.5, LSL: 9.5, Sigma: 0.1})
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("%.4f\n", cp)
	// Output: 1.6667
}

func ExampleCpk() {
	cpk, err := Cpk(CpkInput{USL: 10.5, LSL: 9.5, Mean: 10.1, Sigma: 0.1})
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("%.4f\n", cpk)
	// Output: 1.3333
}

func ExampleCostOfQuality() {
	result, err := CostOfQuality(CostOfQualityInput{
		PreventionCost:      5000,
		AppraisalCost:       8000,
		InternalFailureCost: 12000,
		ExternalFailureCost: 20000,
	})
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("Total=%.0f Conformance=%.0f NonConformance=%.0f\n", result.Total, result.ConformanceCost, result.NonConformanceCost)
	// Output: Total=45000 Conformance=13000 NonConformance=32000
}
