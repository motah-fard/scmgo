package logistics

import (
	"math"
	"testing"
)

func FuzzDimensionalWeight(f *testing.F) {
	f.Add(20.0, 15.0, 10.0, 139.0)
	f.Fuzz(func(t *testing.T, length, width, height, dimFactor float64) {
		result, err := DimensionalWeight(DimensionalWeightInput{
			Length: length, Width: width, Height: height, DimFactor: dimFactor,
		})
		if err == nil && math.IsNaN(result) {
			t.Fatalf("DimensionalWeight returned NaN with nil error")
		}
	})
}

func FuzzAllocateFreightCost(f *testing.F) {
	f.Add(1000.0, 100.0, 300.0, 600.0)
	f.Fuzz(func(t *testing.T, totalCost, w1, w2, w3 float64) {
		results, err := AllocateFreightCost(AllocateFreightCostInput{
			TotalFreightCost: totalCost,
			Items: []FreightAllocationItem{
				{ID: "a", Weight: w1},
				{ID: "b", Weight: w2},
				{ID: "c", Weight: w3},
			},
		})
		if err != nil {
			return
		}
		for _, r := range results {
			if math.IsNaN(r.AllocatedCost) {
				t.Fatalf("AllocateFreightCost returned NaN with nil error: %+v", results)
			}
		}
	})
}

func FuzzCenterOfGravity(f *testing.F) {
	f.Add(10.0, 20.0, 500.0, 30.0, 40.0, 300.0)
	f.Fuzz(func(t *testing.T, x1, y1, d1, x2, y2, d2 float64) {
		result, err := CenterOfGravity(CenterOfGravityInput{
			Locations: []LocationDemand{
				{X: x1, Y: y1, Demand: d1},
				{X: x2, Y: y2, Demand: d2},
			},
		})
		if err == nil && (math.IsNaN(result.X) || math.IsNaN(result.Y)) {
			t.Fatalf("CenterOfGravity returned NaN with nil error")
		}
	})
}
