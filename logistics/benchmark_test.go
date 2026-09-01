package logistics

import "testing"

func BenchmarkAllocateFreightCost(b *testing.B) {
	in := AllocateFreightCostInput{
		TotalFreightCost: 1000,
		Items: []FreightAllocationItem{
			{ID: "a", Weight: 100},
			{ID: "b", Weight: 300},
			{ID: "c", Weight: 600},
		},
	}
	for i := 0; i < b.N; i++ {
		_, _ = AllocateFreightCost(in)
	}
}
