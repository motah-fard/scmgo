package quality

import "testing"

func BenchmarkCpk(b *testing.B) {
	in := CpkInput{USL: 10.5, LSL: 9.5, Mean: 10.1, Sigma: 0.1}
	for i := 0; i < b.N; i++ {
		_, _ = Cpk(in)
	}
}
