package warehouse

import (
	"math"
	"testing"
)

func almostEqual(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}

func TestStorageUtilization(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   StorageUtilizationInput
		want    float64
		wantErr error
	}{
		{
			name:  "valid input",
			input: StorageUtilizationInput{UsedSpace: 800, TotalSpace: 1000},
			want:  0.8,
		},
		{
			name:  "overcommitted is not clamped",
			input: StorageUtilizationInput{UsedSpace: 1100, TotalSpace: 1000},
			want:  1.1,
		},
		{
			name:    "negative used space",
			input:   StorageUtilizationInput{UsedSpace: -1, TotalSpace: 1000},
			wantErr: ErrNegativeUsedSpace,
		},
		{
			name:    "zero total space",
			input:   StorageUtilizationInput{UsedSpace: 800, TotalSpace: 0},
			wantErr: ErrInvalidTotalSpace,
		},
		{
			name:    "NaN used space",
			input:   StorageUtilizationInput{UsedSpace: math.NaN(), TotalSpace: 1000},
			wantErr: ErrNonFiniteInput,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := StorageUtilization(tt.input)
			if err != tt.wantErr {
				t.Fatalf("StorageUtilization() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}

			if !almostEqual(got, tt.want, 1e-9) {
				t.Fatalf("StorageUtilization() = %v, want %v", got, tt.want)
			}
		})
	}
}
