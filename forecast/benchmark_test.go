package forecast

import "testing"

func BenchmarkMovingAverage(b *testing.B) {
	in := MovingAverageInput{
		History: []float64{100, 120, 110, 130, 125},
		Periods: 3,
	}
	for i := 0; i < b.N; i++ {
		_, _ = MovingAverage(in)
	}
}

func BenchmarkSimpleExponentialSmoothing(b *testing.B) {
	in := SimpleExponentialSmoothingInput{
		History: []float64{100, 120, 110, 130, 125},
		Alpha:   0.3,
	}
	for i := 0; i < b.N; i++ {
		_, _ = SimpleExponentialSmoothing(in)
	}
}

func BenchmarkHoltLinearTrend(b *testing.B) {
	in := HoltLinearTrendInput{
		History:      []float64{100, 120, 110, 130, 125},
		Alpha:        0.3,
		Beta:         0.2,
		PeriodsAhead: 3,
	}
	for i := 0; i < b.N; i++ {
		_, _ = HoltLinearTrend(in)
	}
}
