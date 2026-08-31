package forecast

import (
	"math"
	"testing"
)

func TestMASE(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		input     MASEInput
		want      float64
		tolerance float64
		wantErr   error
	}{
		{
			name: "valid input",
			input: MASEInput{
				TrainingHistory: []float64{100, 110, 105, 120, 115, 130},
				Actual:          []float64{125, 135},
				Forecast:        []float64{120, 140},
			},
			want:      0.5,
			tolerance: 1e-9,
		},
		{
			name: "insufficient training history",
			input: MASEInput{
				TrainingHistory: []float64{100},
				Actual:          []float64{100},
				Forecast:        []float64{100},
			},
			wantErr: ErrInsufficientHistory,
		},
		{
			name: "empty actual",
			input: MASEInput{
				TrainingHistory: []float64{100, 110},
				Actual:          []float64{},
				Forecast:        []float64{},
			},
			wantErr: ErrEmptyHistory,
		},
		{
			name: "mismatched lengths",
			input: MASEInput{
				TrainingHistory: []float64{100, 110},
				Actual:          []float64{100, 110},
				Forecast:        []float64{100},
			},
			wantErr: ErrMismatchedLengths,
		},
		{
			name: "constant training history yields zero naive MAE",
			input: MASEInput{
				TrainingHistory: []float64{100, 100, 100},
				Actual:          []float64{100},
				Forecast:        []float64{105},
			},
			wantErr: ErrZeroNaiveMAE,
		},
		{
			name: "NaN in training history",
			input: MASEInput{
				TrainingHistory: []float64{100, math.NaN()},
				Actual:          []float64{100},
				Forecast:        []float64{105},
			},
			wantErr: ErrNonFiniteInput,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := MASE(tt.input)
			if err != tt.wantErr {
				t.Fatalf("MASE() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}

			if !almostEqual(got, tt.want, tt.tolerance) {
				t.Fatalf("MASE() = %v, want %v", got, tt.want)
			}
		})
	}
}
