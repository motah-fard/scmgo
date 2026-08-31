package inventory

import (
	"math"
	"testing"
)

func TestDaysOfInventoryOnHand(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		input     DaysOfInventoryOnHandInput
		want      float64
		tolerance float64
		wantErr   error
	}{
		{
			name: "valid input",
			input: DaysOfInventoryOnHandInput{
				AverageInventoryValue: 20000,
				COGS:                  120000,
				DaysInPeriod:          365,
			},
			want:      60.83333333333333,
			tolerance: 1e-9,
		},
		{
			name: "zero average inventory value",
			input: DaysOfInventoryOnHandInput{
				AverageInventoryValue: 0,
				COGS:                  120000,
				DaysInPeriod:          365,
			},
			want:      0,
			tolerance: 1e-9,
		},
		{
			name: "negative average inventory value",
			input: DaysOfInventoryOnHandInput{
				AverageInventoryValue: -1,
				COGS:                  120000,
				DaysInPeriod:          365,
			},
			wantErr: ErrNegativeAverageInventoryValue,
		},
		{
			name: "zero COGS",
			input: DaysOfInventoryOnHandInput{
				AverageInventoryValue: 20000,
				COGS:                  0,
				DaysInPeriod:          365,
			},
			wantErr: ErrInvalidCOGS,
		},
		{
			name: "zero days in period",
			input: DaysOfInventoryOnHandInput{
				AverageInventoryValue: 20000,
				COGS:                  120000,
				DaysInPeriod:          0,
			},
			wantErr: ErrInvalidDaysInPeriod,
		},
		{
			name: "NaN days in period",
			input: DaysOfInventoryOnHandInput{
				AverageInventoryValue: 20000,
				COGS:                  120000,
				DaysInPeriod:          math.NaN(),
			},
			wantErr: ErrNonFiniteInput,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := DaysOfInventoryOnHand(tt.input)
			if err != tt.wantErr {
				t.Fatalf("DaysOfInventoryOnHand() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}

			if !almostEqual(got, tt.want, tt.tolerance) {
				t.Fatalf("DaysOfInventoryOnHand() = %v, want %v", got, tt.want)
			}
		})
	}
}
