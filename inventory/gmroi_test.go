package inventory

import (
	"math"
	"testing"
)

func TestGMROI(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		input     GMROIInput
		want      float64
		tolerance float64
		wantErr   error
	}{
		{
			name:      "valid input",
			input:     GMROIInput{GrossMargin: 50000, AverageInventoryCost: 20000},
			want:      2.5,
			tolerance: 1e-9,
		},
		{
			name:      "negative gross margin is a valid loss scenario",
			input:     GMROIInput{GrossMargin: -10000, AverageInventoryCost: 20000},
			want:      -0.5,
			tolerance: 1e-9,
		},
		{
			name:    "zero average inventory cost",
			input:   GMROIInput{GrossMargin: 50000, AverageInventoryCost: 0},
			wantErr: ErrInvalidAverageInventoryCost,
		},
		{
			name:    "negative average inventory cost",
			input:   GMROIInput{GrossMargin: 50000, AverageInventoryCost: -1},
			wantErr: ErrInvalidAverageInventoryCost,
		},
		{
			name:    "NaN gross margin",
			input:   GMROIInput{GrossMargin: math.NaN(), AverageInventoryCost: 20000},
			wantErr: ErrNonFiniteInput,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := GMROI(tt.input)
			if err != tt.wantErr {
				t.Fatalf("GMROI() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}

			if !almostEqual(got, tt.want, tt.tolerance) {
				t.Fatalf("GMROI() = %v, want %v", got, tt.want)
			}
		})
	}
}
