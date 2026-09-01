package forecast

import (
	"math"
	"testing"
)

func TestNaive(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   NaiveInput
		want    float64
		wantErr error
	}{
		{
			name:  "valid input",
			input: NaiveInput{History: []float64{100, 120, 90, 110, 105, 125, 95, 115}},
			want:  115,
		},
		{
			name:  "single period",
			input: NaiveInput{History: []float64{50}},
			want:  50,
		},
		{
			name:    "empty history",
			input:   NaiveInput{History: nil},
			wantErr: ErrEmptyHistory,
		},
		{
			name:    "negative demand",
			input:   NaiveInput{History: []float64{100, -1}},
			wantErr: ErrNegativeDemand,
		},
		{
			name:    "NaN in history",
			input:   NaiveInput{History: []float64{100, math.NaN()}},
			wantErr: ErrNonFiniteInput,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := Naive(tt.input)
			if err != tt.wantErr {
				t.Fatalf("Naive() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}

			if got != tt.want {
				t.Fatalf("Naive() = %v, want %v", got, tt.want)
			}
		})
	}
}
