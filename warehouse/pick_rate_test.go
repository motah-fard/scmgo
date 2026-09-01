package warehouse

import (
	"math"
	"testing"
)

func TestPickRate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   PickRateInput
		want    float64
		wantErr error
	}{
		{
			name:  "valid input",
			input: PickRateInput{TotalPicks: 480, TotalHours: 8},
			want:  60,
		},
		{
			name:    "negative picks",
			input:   PickRateInput{TotalPicks: -1, TotalHours: 8},
			wantErr: ErrNegativePicks,
		},
		{
			name:    "zero hours",
			input:   PickRateInput{TotalPicks: 480, TotalHours: 0},
			wantErr: ErrInvalidHours,
		},
		{
			name:    "NaN hours",
			input:   PickRateInput{TotalPicks: 480, TotalHours: math.NaN()},
			wantErr: ErrNonFiniteInput,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := PickRate(tt.input)
			if err != tt.wantErr {
				t.Fatalf("PickRate() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}

			if !almostEqual(got, tt.want, 1e-9) {
				t.Fatalf("PickRate() = %v, want %v", got, tt.want)
			}
		})
	}
}
