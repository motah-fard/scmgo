package inventory

import (
	"math"
	"testing"
)

func almostEqual(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}

func TestEOQ(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		input     EOQInput
		want      float64
		tolerance float64
		wantErr   error
	}{
		{
			name: "valid input",
			input: EOQInput{
				AnnualDemand:       10000,
				OrderingCost:       50,
				HoldingCostPerUnit: 2,
			},
			want:      707.1067811865476,
			tolerance: 1e-9,
		},
		{
			name: "zero annual demand",
			input: EOQInput{
				AnnualDemand:       0,
				OrderingCost:       50,
				HoldingCostPerUnit: 2,
			},
			want:      0,
			tolerance: 1e-9,
		},
		{
			name: "zero ordering cost",
			input: EOQInput{
				AnnualDemand:       10000,
				OrderingCost:       0,
				HoldingCostPerUnit: 2,
			},
			want:      0,
			tolerance: 1e-9,
		},
		{
			name: "negative annual demand",
			input: EOQInput{
				AnnualDemand:       -1,
				OrderingCost:       50,
				HoldingCostPerUnit: 2,
			},
			wantErr: ErrNegativeDemand,
		},
		{
			name: "negative ordering cost",
			input: EOQInput{
				AnnualDemand:       10000,
				OrderingCost:       -1,
				HoldingCostPerUnit: 2,
			},
			wantErr: ErrNegativeOrderingCost,
		},
		{
			name: "zero holding cost",
			input: EOQInput{
				AnnualDemand:       10000,
				OrderingCost:       50,
				HoldingCostPerUnit: 0,
			},
			wantErr: ErrInvalidHoldingCost,
		},
		{
			name: "negative holding cost",
			input: EOQInput{
				AnnualDemand:       10000,
				OrderingCost:       50,
				HoldingCostPerUnit: -1,
			},
			wantErr: ErrInvalidHoldingCost,
		},
		{
			name: "very large values",
			input: EOQInput{
				AnnualDemand:       1_000_000_000,
				OrderingCost:       10_000,
				HoldingCostPerUnit: 25,
			},
			want:      894427.1909999159,
			tolerance: 1e-6,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := EOQ(tt.input)
			if err != tt.wantErr {
				t.Fatalf("EOQ() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}

			if !almostEqual(got, tt.want, tt.tolerance) {
				t.Fatalf("EOQ() = %v, want %v", got, tt.want)
			}
		})
	}
}
