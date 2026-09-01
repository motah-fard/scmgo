package forecast

import (
	"math"
	"testing"
)

func TestSBA(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		input        CrostonInput
		wantForecast float64
		tolerance    float64
		wantErr      error
	}{
		{
			name: "valid input",
			input: CrostonInput{
				History: []float64{0, 0, 5, 0, 0, 0, 3, 0, 4, 0},
				Alpha:   0.2,
			},
			wantForecast: 1.3621621621621618,
			tolerance:    1e-9,
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
			name: "no non-zero demand",
			input: CrostonInput{
				History: []float64{0, 0, 0},
				Alpha:   0.2,
			},
			wantErr: ErrNoNonZeroDemand,
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
			name: "NaN in history",
			input: CrostonInput{
				History: []float64{0, math.NaN(), 5},
				Alpha:   0.2,
			},
			wantErr: ErrNonFiniteInput,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := SBA(tt.input)
			if err != tt.wantErr {
				t.Fatalf("SBA() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}

			if !almostEqual(got.Forecast, tt.wantForecast, tt.tolerance) {
				t.Fatalf("SBA() forecast = %v, want %v", got.Forecast, tt.wantForecast)
			}
		})
	}
}

func TestSBALowerThanCroston(t *testing.T) {
	t.Parallel()

	in := CrostonInput{
		History: []float64{0, 0, 5, 0, 0, 0, 3, 0, 4, 0},
		Alpha:   0.2,
	}

	crostonResult, err := Croston(in)
	if err != nil {
		t.Fatalf("Croston() unexpected error: %v", err)
	}
	sbaResult, err := SBA(in)
	if err != nil {
		t.Fatalf("SBA() unexpected error: %v", err)
	}

	if sbaResult.Forecast >= crostonResult.Forecast {
		t.Fatalf("SBA forecast %v should be strictly less than Croston forecast %v (bias correction)", sbaResult.Forecast, crostonResult.Forecast)
	}
	if sbaResult.DemandSize != crostonResult.DemandSize {
		t.Fatalf("SBA DemandSize = %v, want unchanged from Croston %v", sbaResult.DemandSize, crostonResult.DemandSize)
	}
	if sbaResult.Interval != crostonResult.Interval {
		t.Fatalf("SBA Interval = %v, want unchanged from Croston %v", sbaResult.Interval, crostonResult.Interval)
	}
}
