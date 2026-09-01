package production

import (
	"math"
	"testing"
)

func TestTaktTime(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   TaktTimeInput
		want    float64
		wantErr error
	}{
		{
			name:  "valid input",
			input: TaktTimeInput{AvailableProductionTime: 480, CustomerDemand: 120},
			want:  4,
		},
		{
			name:    "zero available time",
			input:   TaktTimeInput{AvailableProductionTime: 0, CustomerDemand: 120},
			wantErr: ErrInvalidAvailableTime,
		},
		{
			name:    "zero customer demand",
			input:   TaktTimeInput{AvailableProductionTime: 480, CustomerDemand: 0},
			wantErr: ErrInvalidCustomerDemand,
		},
		{
			name:    "NaN customer demand",
			input:   TaktTimeInput{AvailableProductionTime: 480, CustomerDemand: math.NaN()},
			wantErr: ErrNonFiniteInput,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := TaktTime(tt.input)
			if err != tt.wantErr {
				t.Fatalf("TaktTime() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}

			if !almostEqual(got, tt.want, 1e-9) {
				t.Fatalf("TaktTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
