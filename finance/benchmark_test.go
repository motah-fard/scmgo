package finance

import "testing"

func BenchmarkCashToCashCycleTime(b *testing.B) {
	in := CashToCashCycleTimeInput{DIO: 60.83, DSO: 30.42, DPO: 36.5}
	for i := 0; i < b.N; i++ {
		_, _ = CashToCashCycleTime(in)
	}
}
