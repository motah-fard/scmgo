package forecast

import (
	"math"
	"testing"
)

func TestLinearTrend(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name          string
		input         LinearTrendInput
		wantSlope     float64
		wantIntercept float64
		wantForecast  float64
		tolerance     float64
		wantErr       error
	}{
		{
			name: "valid input",
			input: LinearTrendInput{
				History:      []float64{100, 105, 108, 115, 120},
				PeriodsAhead: 3,
			},
			wantSlope:     5.0,
			wantIntercept: 99.6,
			wantForecast:  134.6,
			tolerance:     1e-9,
		},
		{
			name: "flat history has zero slope",
			input: LinearTrendInput{
				History:      []float64{50, 50, 50, 50},
				PeriodsAhead: 1,
			},
			wantSlope:     0,
			wantIntercept: 50,
			wantForecast:  50,
			tolerance:     1e-9,
		},
		{
			name: "insufficient history",
			input: LinearTrendInput{
				History:      []float64{100},
				PeriodsAhead: 1,
			},
			wantErr: ErrInsufficientHistory,
		},
		{
			name: "zero periods ahead",
			input: LinearTrendInput{
				History:      []float64{100, 105},
				PeriodsAhead: 0,
			},
			wantErr: ErrInvalidPeriods,
		},
		{
			name: "negative demand",
			input: LinearTrendInput{
				History:      []float64{100, -105},
				PeriodsAhead: 1,
			},
			wantErr: ErrNegativeDemand,
		},
		{
			name: "NaN in history",
			input: LinearTrendInput{
				History:      []float64{100, math.NaN()},
				PeriodsAhead: 1,
			},
			wantErr: ErrNonFiniteInput,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := LinearTrend(tt.input)
			if err != tt.wantErr {
				t.Fatalf("LinearTrend() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}

			if !almostEqual(got.Slope, tt.wantSlope, tt.tolerance) {
				t.Fatalf("LinearTrend() slope = %v, want %v", got.Slope, tt.wantSlope)
			}
			if !almostEqual(got.Intercept, tt.wantIntercept, tt.tolerance) {
				t.Fatalf("LinearTrend() intercept = %v, want %v", got.Intercept, tt.wantIntercept)
			}
			if !almostEqual(got.Forecast, tt.wantForecast, tt.tolerance) {
				t.Fatalf("LinearTrend() forecast = %v, want %v", got.Forecast, tt.wantForecast)
			}
		})
	}
}
