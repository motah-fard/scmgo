package quality

import (
	"math"
	"testing"
)

func almostEqual(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}

func TestDPMO(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   DPMOInput
		want    float64
		wantErr error
	}{
		{
			name:  "valid input",
			input: DPMOInput{NumberOfDefects: 45, NumberOfUnits: 1000, OpportunitiesPerUnit: 5},
			want:  9000,
		},
		{
			name:    "negative defects",
			input:   DPMOInput{NumberOfDefects: -1, NumberOfUnits: 1000, OpportunitiesPerUnit: 5},
			wantErr: ErrNegativeDefects,
		},
		{
			name:    "zero units",
			input:   DPMOInput{NumberOfDefects: 45, NumberOfUnits: 0, OpportunitiesPerUnit: 5},
			wantErr: ErrInvalidUnits,
		},
		{
			name:    "zero opportunities",
			input:   DPMOInput{NumberOfDefects: 45, NumberOfUnits: 1000, OpportunitiesPerUnit: 0},
			wantErr: ErrInvalidOpportunities,
		},
		{
			name:    "NaN defects",
			input:   DPMOInput{NumberOfDefects: math.NaN(), NumberOfUnits: 1000, OpportunitiesPerUnit: 5},
			wantErr: ErrNonFiniteInput,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := DPMO(tt.input)
			if err != tt.wantErr {
				t.Fatalf("DPMO() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}

			if !almostEqual(got, tt.want, 1e-9) {
				t.Fatalf("DPMO() = %v, want %v", got, tt.want)
			}
		})
	}
}
