// Command abc-xyz demonstrates classifying a SKU set two ways —
// by value concentration (ABC) and by demand variability (XYZ) — and
// combining them into the classic ABC-XYZ matrix used to decide which
// SKUs get tight, formula-driven control versus a looser policy.
//
// Run it with:
//
//	go run ./examples/abc-xyz
package main

import (
	"fmt"
	"log"

	"github.com/motah-fard/scmgo/abc"
)

func main() {
	valueItems := []abc.Item{
		{ID: "SKU-001", Value: 82000},
		{ID: "SKU-002", Value: 41000},
		{ID: "SKU-003", Value: 9500},
		{ID: "SKU-004", Value: 3200},
		{ID: "SKU-005", Value: 1100},
	}

	abcResults, err := abc.Classify(abc.ClassifyInput{
		Items:      valueItems,
		AThreshold: 0.80,
		BThreshold: 0.95,
	})
	if err != nil {
		log.Fatalf("ABC classify: %v", err)
	}

	variabilityItems := []abc.VariabilityItem{
		{ID: "SKU-001", MeanDemand: 500, StdDevDemand: 40},  // steady
		{ID: "SKU-002", MeanDemand: 300, StdDevDemand: 210}, // erratic
		{ID: "SKU-003", MeanDemand: 120, StdDevDemand: 15},  // steady
		{ID: "SKU-004", MeanDemand: 60, StdDevDemand: 55},   // erratic
		{ID: "SKU-005", MeanDemand: 20, StdDevDemand: 4},    // steady
	}

	xyzResults, err := abc.ClassifyVariability(abc.ClassifyVariabilityInput{
		Items:      variabilityItems,
		XThreshold: 0.10,
		YThreshold: 0.50,
	})
	if err != nil {
		log.Fatalf("XYZ classify: %v", err)
	}

	combined := abc.Combine(abcResults, xyzResults)

	// Combine only returns ID/class pairs; look the supporting numbers
	// back up by ID (not index — Combine's output isn't guaranteed to
	// stay index-aligned with either input once IDs are filtered out).
	cumPctByID := make(map[string]float64, len(abcResults))
	for _, r := range abcResults {
		cumPctByID[r.ID] = r.CumulativePercent
	}
	cvByID := make(map[string]float64, len(xyzResults))
	for _, r := range xyzResults {
		cvByID[r.ID] = r.CoefficientOfVariation
	}

	fmt.Printf("%-10s %-6s %8s   %-6s %8s   %-6s\n", "SKU", "ABC", "CumPct", "XYZ", "CV", "Matrix")
	for _, c := range combined {
		fmt.Printf("%-10s %-6s %7.1f%%   %-6s %8.2f   %-6s\n",
			c.ID, c.ABCClass, cumPctByID[c.ID]*100,
			c.XYZClass, cvByID[c.ID], c.ABCClass+c.XYZClass)
	}
}
