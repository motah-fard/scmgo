package forecast

import (
	"math"
	"testing"
)

func almostEqual(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}

func TestMovingAverage(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		input     MovingAverageInput
		want      float64
		tolerance float64
		wantErr   error
	}{
		{
			name: "valid input",
			input: MovingAverageInput{
				History: []float64{100, 120, 110, 130, 125},
				Periods: 3,
			},
			want:      121.66666666666667,
			tolerance: 1e-9,
		},
		{
			name: "window equals full history",
			input: MovingAverageInput{
				History: []float64{10, 20, 30},
				Periods: 3,
			},
			want:      20,
			tolerance: 1e-9,
		},
		{
			name: "zero periods",
			input: MovingAverageInput{
				History: []float64{10, 20, 30},
				Periods: 0,
			},
			wantErr: ErrInvalidPeriods,
		},
		{
			name: "negative periods",
			input: MovingAverageInput{
				History: []float64{10, 20, 30},
				Periods: -1,
			},
			wantErr: ErrInvalidPeriods,
		},
		{
			name: "insufficient history",
			input: MovingAverageInput{
				History: []float64{10, 20},
				Periods: 3,
			},
			wantErr: ErrInsufficientHistory,
		},
		{
			name: "negative demand",
			input: MovingAverageInput{
				History: []float64{10, -5, 30},
				Periods: 3,
			},
			wantErr: ErrNegativeDemand,
		},
		{
			name: "NaN in history",
			input: MovingAverageInput{
				History: []float64{10, math.NaN(), 30},
				Periods: 3,
			},
			wantErr: ErrNonFiniteInput,
		},
		{
			name: "Inf in history",
			input: MovingAverageInput{
				History: []float64{10, math.Inf(1), 30},
				Periods: 3,
			},
			wantErr: ErrNonFiniteInput,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := MovingAverage(tt.input)
			if err != tt.wantErr {
				t.Fatalf("MovingAverage() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}

			if !almostEqual(got, tt.want, tt.tolerance) {
				t.Fatalf("MovingAverage() = %v, want %v", got, tt.want)
			}
		})
	}
}
