package warehouse

import "testing"

func BenchmarkPickRate(b *testing.B) {
	in := PickRateInput{TotalPicks: 480, TotalHours: 8}
	for i := 0; i < b.N; i++ {
		_, _ = PickRate(in)
	}
}
