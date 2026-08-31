package forecast

import "testing"

func TestWeightedMovingAverage(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		input     WeightedMovingAverageInput
		want      float64
		tolerance float64
		wantErr   error
	}{
		{
			name: "valid input",
			input: WeightedMovingAverageInput{
				History: []float64{100, 120, 110, 130, 125},
				Weights: []float64{0.1, 0.3, 0.6},
			},
			want:      125,
			tolerance: 1e-9,
		},
		{
			name: "empty weights",
			input: WeightedMovingAverageInput{
				History: []float64{100, 120, 110},
				Weights: []float64{},
			},
			wantErr: ErrInvalidPeriods,
		},
		{
			name: "insufficient history",
			input: WeightedMovingAverageInput{
				History: []float64{100},
				Weights: []float64{0.5, 0.5},
			},
			wantErr: ErrInsufficientHistory,
		},
		{
			name: "negative weight",
			input: WeightedMovingAverageInput{
				History: []float64{100, 120, 110},
				Weights: []float64{-0.1, 1.1},
			},
			wantErr: ErrInvalidWeight,
		},
		{
			name: "weights do not sum to one",
			input: WeightedMovingAverageInput{
				History: []float64{100, 120, 110},
				Weights: []float64{0.2, 0.2},
			},
			wantErr: ErrWeightsMustSumToOne,
		},
		{
			name: "negative demand",
			input: WeightedMovingAverageInput{
				History: []float64{100, -120, 110},
				Weights: []float64{0.5, 0.5},
			},
			wantErr: ErrNegativeDemand,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := WeightedMovingAverage(tt.input)
			if err != tt.wantErr {
				t.Fatalf("WeightedMovingAverage() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}

			if !almostEqual(got, tt.want, tt.tolerance) {
				t.Fatalf("WeightedMovingAverage() = %v, want %v", got, tt.want)
			}
		})
	}
}
