package procurement

import "testing"

func BenchmarkLandedCost(b *testing.B) {
	in := LandedCostInput{
		Components: []CostComponent{
			{Label: "Unit Cost", Amount: 100},
			{Label: "Freight", Amount: 15},
			{Label: "Duty", Amount: 8},
			{Label: "Insurance", Amount: 2},
		},
	}
	for i := 0; i < b.N; i++ {
		_, _ = LandedCost(in)
	}
}
