package inventory

import (
	"math"
	"testing"
)

func TestSafetyTime(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		input     SafetyTimeInput
		want      float64
		tolerance float64
		wantErr   error
	}{
		{
			name:      "valid input",
			input:     SafetyTimeInput{SafetyStockUnits: 50, AvgDailyDemand: 20},
			want:      2.5,
			tolerance: 1e-9,
		},
		{
			name:      "zero safety stock",
			input:     SafetyTimeInput{SafetyStockUnits: 0, AvgDailyDemand: 20},
			want:      0,
			tolerance: 1e-9,
		},
		{
			name:    "negative safety stock",
			input:   SafetyTimeInput{SafetyStockUnits: -1, AvgDailyDemand: 20},
			wantErr: ErrNegativeSafetyStock,
		},
		{
			name:    "zero avg daily demand",
			input:   SafetyTimeInput{SafetyStockUnits: 50, AvgDailyDemand: 0},
			wantErr: ErrInvalidAvgDailyDemand,
		},
		{
			name:    "negative avg daily demand",
			input:   SafetyTimeInput{SafetyStockUnits: 50, AvgDailyDemand: -1},
			wantErr: ErrInvalidAvgDailyDemand,
		},
		{
			name:    "NaN safety stock",
			input:   SafetyTimeInput{SafetyStockUnits: math.NaN(), AvgDailyDemand: 20},
			wantErr: ErrNonFiniteInput,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := SafetyTime(tt.input)
			if err != tt.wantErr {
				t.Fatalf("SafetyTime() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}

			if !almostEqual(got, tt.want, tt.tolerance) {
				t.Fatalf("SafetyTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
