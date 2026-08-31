package forecast

import (
	"math"
	"testing"
)

func TestAccuracy(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		input     AccuracyInput
		want      AccuracyResult
		wantNaN   bool
		tolerance float64
		wantErr   error
	}{
		{
			name: "valid input",
			input: AccuracyInput{
				Actual:   []float64{100, 110, 95, 130},
				Forecast: []float64{90, 115, 100, 120},
			},
			want: AccuracyResult{
				MAD:  7.5,
				MAPE: 0.0687523003312477,
				Bias: 2.5,
				RMSE: 7.905694150420948,
			},
			tolerance: 1e-9,
		},
		{
			name: "zero actual excluded from MAPE",
			input: AccuracyInput{
				Actual:   []float64{0, 10},
				Forecast: []float64{5, 8},
			},
			want: AccuracyResult{
				MAD:  3.5,
				MAPE: 0.2,
				Bias: -1.5,
				RMSE: math.Sqrt((25 + 4) / 2.0),
			},
			tolerance: 1e-9,
		},
		{
			name: "all actuals zero yields NaN MAPE",
			input: AccuracyInput{
				Actual:   []float64{0, 0},
				Forecast: []float64{5, 8},
			},
			want: AccuracyResult{
				MAD:  6.5,
				Bias: -6.5,
				RMSE: math.Sqrt((25 + 64) / 2.0),
			},
			wantNaN:   true,
			tolerance: 1e-9,
		},
		{
			name: "empty actual",
			input: AccuracyInput{
				Actual:   []float64{},
				Forecast: []float64{},
			},
			wantErr: ErrEmptyHistory,
		},
		{
			name: "mismatched lengths",
			input: AccuracyInput{
				Actual:   []float64{1, 2},
				Forecast: []float64{1},
			},
			wantErr: ErrMismatchedLengths,
		},
		{
			name: "NaN in actual",
			input: AccuracyInput{
				Actual:   []float64{100, math.NaN()},
				Forecast: []float64{90, 100},
			},
			wantErr: ErrNonFiniteInput,
		},
		{
			name: "Inf in forecast",
			input: AccuracyInput{
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

			got, err := Accuracy(tt.input)
			if err != tt.wantErr {
				t.Fatalf("Accuracy() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}

			if !almostEqual(got.MAD, tt.want.MAD, tt.tolerance) {
				t.Fatalf("Accuracy() MAD = %v, want %v", got.MAD, tt.want.MAD)
			}
			if !almostEqual(got.Bias, tt.want.Bias, tt.tolerance) {
				t.Fatalf("Accuracy() Bias = %v, want %v", got.Bias, tt.want.Bias)
			}
			if !almostEqual(got.RMSE, tt.want.RMSE, tt.tolerance) {
				t.Fatalf("Accuracy() RMSE = %v, want %v", got.RMSE, tt.want.RMSE)
			}
			if tt.wantNaN {
				if !math.IsNaN(got.MAPE) {
					t.Fatalf("Accuracy() MAPE = %v, want NaN", got.MAPE)
				}
				return
			}
			if !almostEqual(got.MAPE, tt.want.MAPE, tt.tolerance) {
				t.Fatalf("Accuracy() MAPE = %v, want %v", got.MAPE, tt.want.MAPE)
			}
		})
	}
}
