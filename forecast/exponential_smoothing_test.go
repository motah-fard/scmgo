package forecast

import (
	"math"
	"testing"
)

func TestSimpleExponentialSmoothing(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		input        SimpleExponentialSmoothingInput
		wantForecast float64
		wantFitted   []float64
		tolerance    float64
		wantErr      error
	}{
		{
			name: "valid input",
			input: SimpleExponentialSmoothingInput{
				History: []float64{100, 120, 110, 130, 125},
				Alpha:   0.3,
			},
			wantForecast: 117.32799999999999,
			wantFitted:   []float64{100, 100, 106, 107.19999999999999, 114.03999999999999},
			tolerance:    1e-9,
		},
		{
			name: "single period",
			input: SimpleExponentialSmoothingInput{
				History: []float64{50},
				Alpha:   0.5,
			},
			wantForecast: 50,
			wantFitted:   []float64{50},
			tolerance:    1e-9,
		},
		{
			name: "empty history",
			input: SimpleExponentialSmoothingInput{
				History: []float64{},
				Alpha:   0.3,
			},
			wantErr: ErrEmptyHistory,
		},
		{
			name: "zero alpha",
			input: SimpleExponentialSmoothingInput{
				History: []float64{100, 120},
				Alpha:   0,
			},
			wantErr: ErrInvalidSmoothingConstant,
		},
		{
			name: "alpha above one",
			input: SimpleExponentialSmoothingInput{
				History: []float64{100, 120},
				Alpha:   1.1,
			},
			wantErr: ErrInvalidSmoothingConstant,
		},
		{
			name: "negative demand",
			input: SimpleExponentialSmoothingInput{
				History: []float64{100, -120},
				Alpha:   0.3,
			},
			wantErr: ErrNegativeDemand,
		},
		{
			name: "NaN in history",
			input: SimpleExponentialSmoothingInput{
				History: []float64{100, math.NaN()},
				Alpha:   0.3,
			},
			wantErr: ErrNonFiniteInput,
		},
		{
			name: "NaN alpha",
			input: SimpleExponentialSmoothingInput{
				History: []float64{100, 120},
				Alpha:   math.NaN(),
			},
			wantErr: ErrNonFiniteInput,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := SimpleExponentialSmoothing(tt.input)
			if err != tt.wantErr {
				t.Fatalf("SimpleExponentialSmoothing() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}

			if !almostEqual(got.Forecast, tt.wantForecast, tt.tolerance) {
				t.Fatalf("SimpleExponentialSmoothing() forecast = %v, want %v", got.Forecast, tt.wantForecast)
			}
			if len(got.Fitted) != len(tt.wantFitted) {
				t.Fatalf("SimpleExponentialSmoothing() fitted length = %v, want %v", len(got.Fitted), len(tt.wantFitted))
			}
			for i := range got.Fitted {
				if !almostEqual(got.Fitted[i], tt.wantFitted[i], tt.tolerance) {
					t.Fatalf("SimpleExponentialSmoothing() fitted[%d] = %v, want %v", i, got.Fitted[i], tt.wantFitted[i])
				}
			}
		})
	}
}
