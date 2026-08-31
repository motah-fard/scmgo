package procurement

import (
	"math"
	"testing"
)

func FuzzLandedCost(f *testing.F) {
	f.Add(100.0, 15.0, 8.0, 2.0)
	f.Fuzz(func(t *testing.T, a, b, c, d float64) {
		result, err := LandedCost(LandedCostInput{
			Components: []CostComponent{
				{Label: "a", Amount: a},
				{Label: "b", Amount: b},
				{Label: "c", Amount: c},
				{Label: "d", Amount: d},
			},
		})
		if err == nil && math.IsNaN(result.Total) {
			t.Fatalf("LandedCost returned NaN with nil error")
		}
	})
}

func FuzzPurchasePriceVariance(f *testing.F) {
	f.Add(10.0, 9.0, 1000.0)
	f.Fuzz(func(t *testing.T, standardCost, actualCost, quantity float64) {
		result, err := PurchasePriceVariance(PurchasePriceVarianceInput{
			StandardCost: standardCost,
			ActualCost:   actualCost,
			Quantity:     quantity,
		})
		if err == nil && math.IsNaN(result) {
			t.Fatalf("PurchasePriceVariance returned NaN with nil error")
		}
	})
}
