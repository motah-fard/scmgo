package forecast

import (
	"math"
	"testing"
)

func TestCroston(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		input     CrostonInput
		want      CrostonResult
		tolerance float64
		wantErr   error
	}{
		{
			name: "valid input",
			input: CrostonInput{
				History: []float64{0, 0, 5, 0, 0, 0, 3, 0, 4, 0},
				Alpha:   0.2,
			},
			want: CrostonResult{
				DemandSize: 4.4799999999999995,
				Interval:   2.9600000000000004,
				Forecast:   1.5135135135135132,
			},
			tolerance: 1e-9,
		},
		{
			name: "single non-zero period",
			input: CrostonInput{
				History: []float64{0, 0, 6},
				Alpha:   0.2,
			},
			want: CrostonResult{
				DemandSize: 6,
				Interval:   3,
				Forecast:   2,
			},
			tolerance: 1e-9,
		},
		{
			name: "no non-zero demand",
			input: CrostonInput{
				History: []float64{0, 0, 0},
				Alpha:   0.2,
			},
			wantErr: ErrNoNonZeroDemand,
		},
		{
			name: "empty history",
			input: CrostonInput{
				History: []float64{},
				Alpha:   0.2,
			},
			wantErr: ErrEmptyHistory,
		},
		{
			name: "invalid alpha",
			input: CrostonInput{
				History: []float64{0, 5},
				Alpha:   0,
			},
			wantErr: ErrInvalidSmoothingConstant,
		},
		{
			name: "negative demand",
			input: CrostonInput{
				History: []float64{0, -5},
				Alpha:   0.2,
			},
			wantErr: ErrNegativeDemand,
		},
		{
			name: "NaN in history",
			input: CrostonInput{
				History: []float64{0, math.NaN(), 5},
				Alpha:   0.2,
			},
			wantErr: ErrNonFiniteInput,
		},
		{
			name: "Inf alpha",
			input: CrostonInput{
				History: []float64{0, 5},
				Alpha:   math.Inf(1),
			},
			wantErr: ErrNonFiniteInput,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := Croston(tt.input)
			if err != tt.wantErr {
				t.Fatalf("Croston() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}

			if !almostEqual(got.DemandSize, tt.want.DemandSize, tt.tolerance) {
				t.Fatalf("Croston() demand size = %v, want %v", got.DemandSize, tt.want.DemandSize)
			}
			if !almostEqual(got.Interval, tt.want.Interval, tt.tolerance) {
				t.Fatalf("Croston() interval = %v, want %v", got.Interval, tt.want.Interval)
			}
			if !almostEqual(got.Forecast, tt.want.Forecast, tt.tolerance) {
				t.Fatalf("Croston() forecast = %v, want %v", got.Forecast, tt.want.Forecast)
			}
		})
	}
}
