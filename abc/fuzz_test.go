package abc

import (
	"math"
	"testing"
)

func FuzzClassify(f *testing.F) {
	f.Add(1000.0, 500.0, 300.0, 100.0, 100.0, 0.80, 0.95)
	f.Fuzz(func(t *testing.T, v1, v2, v3, v4, v5, aThreshold, bThreshold float64) {
		results, err := Classify(ClassifyInput{
			Items: []Item{
				{ID: "1", Value: v1},
				{ID: "2", Value: v2},
				{ID: "3", Value: v3},
				{ID: "4", Value: v4},
				{ID: "5", Value: v5},
			},
			AThreshold: aThreshold,
			BThreshold: bThreshold,
		})
		if err != nil {
			return
		}
		if len(results) != 5 {
			t.Fatalf("Classify() returned %d results, want 5", len(results))
		}
		for _, r := range results {
			if math.IsNaN(r.CumulativePercent) {
				t.Fatalf("Classify() returned NaN CumulativePercent with nil error: %+v", r)
			}
			if r.Class != "A" && r.Class != "B" && r.Class != "C" {
				t.Fatalf("Classify() returned invalid class %q", r.Class)
			}
		}
	})
}

func FuzzClassifyVariability(f *testing.F) {
	f.Add(100.0, 5.0, 0.10, 0.50)
	f.Fuzz(func(t *testing.T, meanDemand, stdDevDemand, xThreshold, yThreshold float64) {
		results, err := ClassifyVariability(ClassifyVariabilityInput{
			Items:      []VariabilityItem{{ID: "1", MeanDemand: meanDemand, StdDevDemand: stdDevDemand}},
			XThreshold: xThreshold,
			YThreshold: yThreshold,
		})
		if err != nil {
			return
		}
		if math.IsNaN(results[0].CoefficientOfVariation) {
			t.Fatalf("ClassifyVariability() returned NaN CoefficientOfVariation with nil error")
		}
		class := results[0].Class
		if class != "X" && class != "Y" && class != "Z" {
			t.Fatalf("ClassifyVariability() returned invalid class %q", class)
		}
	})
}
