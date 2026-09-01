package production

import "testing"

func BenchmarkOEE(b *testing.B) {
	in := OEEInput{
		PlannedProductionTime: 480,
		RunTime:               420,
		IdealCycleTime:        1.0,
		TotalCount:            400,
		GoodCount:             385,
	}
	for i := 0; i < b.N; i++ {
		_, _ = OEE(in)
	}
}
