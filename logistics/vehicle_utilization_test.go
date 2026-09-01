package logistics

import (
	"math"
	"testing"
)

func TestVehicleUtilization(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   VehicleUtilizationInput
		want    float64
		wantErr error
	}{
		{
			name:  "valid input",
			input: VehicleUtilizationInput{ActualLoad: 8000, Capacity: 10000},
			want:  0.8,
		},
		{
			name:  "overloaded is not clamped",
			input: VehicleUtilizationInput{ActualLoad: 11000, Capacity: 10000},
			want:  1.1,
		},
		{
			name:    "negative actual load",
			input:   VehicleUtilizationInput{ActualLoad: -1, Capacity: 10000},
			wantErr: ErrNegativeLoad,
		},
		{
			name:    "zero capacity",
			input:   VehicleUtilizationInput{ActualLoad: 8000, Capacity: 0},
			wantErr: ErrInvalidCapacity,
		},
		{
			name:    "NaN actual load",
			input:   VehicleUtilizationInput{ActualLoad: math.NaN(), Capacity: 10000},
			wantErr: ErrNonFiniteInput,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := VehicleUtilization(tt.input)
			if err != tt.wantErr {
				t.Fatalf("VehicleUtilization() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}

			if !almostEqual(got, tt.want, 1e-9) {
				t.Fatalf("VehicleUtilization() = %v, want %v", got, tt.want)
			}
		})
	}
}
