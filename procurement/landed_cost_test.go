package procurement

import (
	"math"
	"testing"
)

func almostEqual(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}

func TestLandedCost(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   LandedCostInput
		want    float64
		wantErr error
	}{
		{
			name: "valid input",
			input: LandedCostInput{
				Components: []CostComponent{
					{Label: "Unit Cost", Amount: 100},
					{Label: "Freight", Amount: 15},
					{Label: "Duty", Amount: 8},
					{Label: "Insurance", Amount: 2},
				},
			},
			want: 125,
		},
		{
			name: "negative component is a valid credit",
			input: LandedCostInput{
				Components: []CostComponent{
					{Label: "Unit Cost", Amount: 100},
					{Label: "Volume Rebate", Amount: -10},
				},
			},
			want: 90,
		},
		{
			name:    "empty components",
			input:   LandedCostInput{Components: nil},
			wantErr: ErrEmptyComponents,
		},
		{
			name: "NaN component",
			input: LandedCostInput{
				Components: []CostComponent{{Label: "x", Amount: math.NaN()}},
			},
			wantErr: ErrNonFiniteInput,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := LandedCost(tt.input)
			if err != tt.wantErr {
				t.Fatalf("LandedCost() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}

			if !almostEqual(got.Total, tt.want, 1e-9) {
				t.Fatalf("LandedCost() Total = %v, want %v", got.Total, tt.want)
			}
			if len(got.Components) != len(tt.input.Components) {
				t.Fatalf("LandedCost() Components length = %v, want %v", len(got.Components), len(tt.input.Components))
			}
		})
	}
}
