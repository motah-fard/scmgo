package warehouse

import (
	"math"
	"testing"
)

func TestCubeUtilization(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   CubeUtilizationInput
		want    float64
		wantErr error
	}{
		{
			name:  "valid input",
			input: CubeUtilizationInput{UsedVolume: 4500, TotalVolume: 6000},
			want:  0.75,
		},
		{
			name:    "negative used volume",
			input:   CubeUtilizationInput{UsedVolume: -1, TotalVolume: 6000},
			wantErr: ErrNegativeUsedVolume,
		},
		{
			name:    "zero total volume",
			input:   CubeUtilizationInput{UsedVolume: 4500, TotalVolume: 0},
			wantErr: ErrInvalidTotalVolume,
		},
		{
			name:    "NaN total volume",
			input:   CubeUtilizationInput{UsedVolume: 4500, TotalVolume: math.NaN()},
			wantErr: ErrNonFiniteInput,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := CubeUtilization(tt.input)
			if err != tt.wantErr {
				t.Fatalf("CubeUtilization() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}

			if !almostEqual(got, tt.want, 1e-9) {
				t.Fatalf("CubeUtilization() = %v, want %v", got, tt.want)
			}
		})
	}
}
