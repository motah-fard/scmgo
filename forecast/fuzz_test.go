package forecast

import (
	"math"
	"testing"
)

// These fuzz targets assert two invariants that must hold for any float64
// input whatsoever:
//
//  1. The function must never panic.
//  2. If it returns a nil error, the result must never contain an
//     undocumented NaN.
//
// (2) exists because the NaN/Inf validation gap fixed in this package (see
// validateFinite) was exactly the kind of bug fuzzing is built to catch —
// a plain "v < 0" check silently passing NaN through instead of failing,
// and returning a NaN result with a nil error. Every function fuzzed here
// only ever combines validated finite values with addition, multiplication,
// or division by a strictly-positive quantity, none of which can produce
// NaN from finite input, so "not NaN" is a safe invariant to assert
// (unlike "not Inf", which a sufficiently large but individually valid
// input could legitimately overflow to). Accuracy's MAPE is the one
// documented exception (NaN when every actual value is zero), so it's
// checked against that precondition instead of asserted non-NaN outright.

func FuzzMovingAverage(f *testing.F) {
	f.Add(100.0, 120.0, 110.0, 130.0, 125.0, 3)
	f.Fuzz(func(t *testing.T, a, b, c, d, e float64, periods int) {
		result, err := MovingAverage(MovingAverageInput{
			History: []float64{a, b, c, d, e},
			Periods: periods,
		})
		if err == nil && math.IsNaN(result) {
			t.Fatalf("MovingAverage returned NaN with nil error")
		}
	})
}

func FuzzWeightedMovingAverage(f *testing.F) {
	f.Add(100.0, 120.0, 110.0, 0.3, 0.7)
	f.Fuzz(func(t *testing.T, a, b, c, w1, w2 float64) {
		result, err := WeightedMovingAverage(WeightedMovingAverageInput{
			History: []float64{a, b, c},
			Weights: []float64{w1, w2},
		})
		if err == nil && math.IsNaN(result) {
			t.Fatalf("WeightedMovingAverage returned NaN with nil error")
		}
	})
}

func FuzzHoltLinearTrend(f *testing.F) {
	f.Add(100.0, 120.0, 110.0, 0.3, 0.2, 3)
	f.Fuzz(func(t *testing.T, a, b, c, alpha, beta float64, periodsAhead int) {
		result, err := HoltLinearTrend(HoltLinearTrendInput{
			History:      []float64{a, b, c},
			Alpha:        alpha,
			Beta:         beta,
			PeriodsAhead: periodsAhead,
		})
		if err != nil {
			return
		}
		if math.IsNaN(result.Level) || math.IsNaN(result.Trend) || math.IsNaN(result.Forecast) {
			t.Fatalf("HoltLinearTrend returned NaN field with nil error: %+v", result)
		}
	})
}

func FuzzCroston(f *testing.F) {
	f.Add(0.0, 5.0, 0.0, 3.0, 0.2)
	f.Fuzz(func(t *testing.T, a, b, c, d, alpha float64) {
		result, err := Croston(CrostonInput{
			History: []float64{a, b, c, d},
			Alpha:   alpha,
		})
		if err != nil {
			return
		}
		if math.IsNaN(result.DemandSize) || math.IsNaN(result.Interval) || math.IsNaN(result.Forecast) {
			t.Fatalf("Croston returned NaN field with nil error: %+v", result)
		}
	})
}

func FuzzAccuracy(f *testing.F) {
	f.Add(100.0, 110.0, 90.0, 115.0)
	f.Fuzz(func(t *testing.T, actual1, actual2, forecast1, forecast2 float64) {
		result, err := Accuracy(AccuracyInput{
			Actual:   []float64{actual1, actual2},
			Forecast: []float64{forecast1, forecast2},
		})
		if err != nil {
			return
		}
		if math.IsNaN(result.MAD) || math.IsNaN(result.Bias) || math.IsNaN(result.RMSE) {
			t.Fatalf("Accuracy returned NaN MAD/Bias/RMSE with nil error: %+v", result)
		}
		if math.IsNaN(result.MAPE) && (actual1 != 0 || actual2 != 0) {
			t.Fatalf("Accuracy returned NaN MAPE despite a non-zero actual value: %+v", result)
		}
	})
}
