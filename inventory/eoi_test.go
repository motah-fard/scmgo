package inventory

import (
	"math"
	"testing"
)

func TestEOI(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		input     EOIInput
		want      float64
		tolerance float64
		wantErr   error
	}{
		{
			name: "valid input",
			input: EOIInput{
				AnnualDemand:       10000,
				OrderingCost:       50,
				HoldingCostPerUnit: 2,
				DaysPerYear:        365,
			},
			want:      25.809397513308983,
			tolerance: 1e-6,
		},
		{
			name: "zero annual demand",
			input: EOIInput{
				AnnualDemand:       0,
				OrderingCost:       50,
				HoldingCostPerUnit: 2,
				DaysPerYear:        365,
			},
			wantErr: ErrInvalidAnnualDemand,
		},
		{
			name: "negative annual demand",
			input: EOIInput{
				AnnualDemand:       -1,
				OrderingCost:       50,
				HoldingCostPerUnit: 2,
				DaysPerYear:        365,
			},
			wantErr: ErrInvalidAnnualDemand,
		},
		{
			name: "zero days per year",
			input: EOIInput{
				AnnualDemand:       10000,
				OrderingCost:       50,
				HoldingCostPerUnit: 2,
				DaysPerYear:        0,
			},
			wantErr: ErrInvalidDaysInPeriod,
		},
		{
			name: "zero holding cost propagates from EOQ",
			input: EOIInput{
				AnnualDemand:       10000,
				OrderingCost:       50,
				HoldingCostPerUnit: 0,
				DaysPerYear:        365,
			},
			wantErr: ErrInvalidHoldingCost,
		},
		{
			name: "NaN ordering cost",
			input: EOIInput{
				AnnualDemand:       10000,
				OrderingCost:       math.NaN(),
				HoldingCostPerUnit: 2,
				DaysPerYear:        365,
			},
			wantErr: ErrNonFiniteInput,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := EOI(tt.input)
			if err != tt.wantErr {
				t.Fatalf("EOI() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}

			if !almostEqual(got, tt.want, tt.tolerance) {
				t.Fatalf("EOI() = %v, want %v", got, tt.want)
			}
		})
	}
}
