package abc

import "fmt"

func ExampleClassify() {
	results, err := Classify(ClassifyInput{
		Items: []Item{
			{ID: "A1", Value: 1000},
			{ID: "A2", Value: 500},
			{ID: "A3", Value: 300},
			{ID: "A4", Value: 100},
			{ID: "A5", Value: 100},
		},
		AThreshold: 0.80,
		BThreshold: 0.95,
	})
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	for _, r := range results {
		fmt.Printf("%s: class=%s cumPct=%.2f\n", r.ID, r.Class, r.CumulativePercent)
	}
	// Output:
	// A1: class=A cumPct=0.50
	// A2: class=A cumPct=0.75
	// A3: class=B cumPct=0.90
	// A4: class=B cumPct=0.95
	// A5: class=C cumPct=1.00
}

func ExampleClassifyVariability() {
	results, err := ClassifyVariability(ClassifyVariabilityInput{
		Items: []VariabilityItem{
			{ID: "steady", MeanDemand: 100, StdDevDemand: 5},
			{ID: "erratic", MeanDemand: 100, StdDevDemand: 80},
		},
		XThreshold: 0.10,
		YThreshold: 0.50,
	})
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	for _, r := range results {
		fmt.Printf("%s: class=%s cv=%.2f\n", r.ID, r.Class, r.CoefficientOfVariation)
	}
	// Output:
	// steady: class=X cv=0.05
	// erratic: class=Z cv=0.80
}

func ExampleCombine() {
	abcResults, _ := Classify(ClassifyInput{
		Items: []Item{
			{ID: "sku1", Value: 8000},
			{ID: "sku2", Value: 2000},
		},
		AThreshold: 0.80,
		BThreshold: 0.95,
	})
	xyzResults, _ := ClassifyVariability(ClassifyVariabilityInput{
		Items: []VariabilityItem{
			{ID: "sku1", MeanDemand: 100, StdDevDemand: 5},
			{ID: "sku2", MeanDemand: 100, StdDevDemand: 80},
		},
		XThreshold: 0.10,
		YThreshold: 0.50,
	})

	for _, c := range Combine(abcResults, xyzResults) {
		fmt.Printf("%s: %s%s\n", c.ID, c.ABCClass, c.XYZClass)
	}
	// Output:
	// sku1: AX
	// sku2: CZ
}
