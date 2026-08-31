package forecast

import (
	"math"
	"testing"
)

func TestForecastAccuracy(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		input     ForecastAccuracyInput
		want      ForecastAccuracyResult
		wantNaN   bool
		tolerance float64
		wantErr   error
	}{
		{
			name: "valid input",
			input: ForecastAccuracyInput{
				Actual:   []float64{100, 110, 95, 130},
				Forecast: []float64{90, 115, 100, 120},
			},
			want: ForecastAccuracyResult{
				MAD:  7.5,
				MAPE: 0.0687523003312477,
				Bias: 2.5,
				RMSE: 7.905694150420948,
			},
			tolerance: 1e-9,
		},
		{
			name: "zero actual excluded from MAPE",
			input: ForecastAccuracyInput{
				Actual:   []float64{0, 10},
				Forecast: []float64{5, 8},
			},
			want: ForecastAccuracyResult{
				MAD:  3.5,
				MAPE: 0.2,
				Bias: -1.5,
				RMSE: math.Sqrt((25 + 4) / 2.0),
			},
			tolerance: 1e-9,
		},
		{
			name: "all actuals zero yields NaN MAPE",
			input: ForecastAccuracyInput{
				Actual:   []float64{0, 0},
				Forecast: []float64{5, 8},
			},
			want: ForecastAccuracyResult{
				MAD:  6.5,
				Bias: -6.5,
				RMSE: math.Sqrt((25 + 64) / 2.0),
			},
			wantNaN:   true,
			tolerance: 1e-9,
		},
		{
			name: "empty actual",
			input: ForecastAccuracyInput{
				Actual:   []float64{},
				Forecast: []float64{},
			},
			wantErr: ErrEmptyHistory,
		},
		{
			name: "mismatched lengths",
			input: ForecastAccuracyInput{
				Actual:   []float64{1, 2},
				Forecast: []float64{1},
			},
			wantErr: ErrMismatchedLengths,
		},
		{
			name: "NaN in actual",
			input: ForecastAccuracyInput{
				Actual:   []float64{100, math.NaN()},
				Forecast: []float64{90, 100},
			},
			wantErr: ErrNonFiniteInput,
		},
		{
			name: "Inf in forecast",
			input: ForecastAccuracyInput{
				Actual:   []float64{100, 110},
				Forecast: []float64{90, math.Inf(1)},
			},
			wantErr: ErrNonFiniteInput,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := ForecastAccuracy(tt.input)
			if err != tt.wantErr {
				t.Fatalf("ForecastAccuracy() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}

			if !almostEqual(got.MAD, tt.want.MAD, tt.tolerance) {
				t.Fatalf("ForecastAccuracy() MAD = %v, want %v", got.MAD, tt.want.MAD)
			}
			if !almostEqual(got.Bias, tt.want.Bias, tt.tolerance) {
				t.Fatalf("ForecastAccuracy() Bias = %v, want %v", got.Bias, tt.want.Bias)
			}
			if !almostEqual(got.RMSE, tt.want.RMSE, tt.tolerance) {
				t.Fatalf("ForecastAccuracy() RMSE = %v, want %v", got.RMSE, tt.want.RMSE)
			}
			if tt.wantNaN {
				if !math.IsNaN(got.MAPE) {
					t.Fatalf("ForecastAccuracy() MAPE = %v, want NaN", got.MAPE)
				}
				return
			}
			if !almostEqual(got.MAPE, tt.want.MAPE, tt.tolerance) {
				t.Fatalf("ForecastAccuracy() MAPE = %v, want %v", got.MAPE, tt.want.MAPE)
			}
		})
	}
}
