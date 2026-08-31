package procurement

import (
	"math"
	"testing"
)

func TestTotalCostOfOwnership(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   TotalCostOfOwnershipInput
		want    float64
		wantErr error
	}{
		{
			name: "valid input",
			input: TotalCostOfOwnershipInput{
				Components: []CostComponent{
					{Label: "Acquisition", Amount: 50000},
					{Label: "Operating (5yr)", Amount: 20000},
					{Label: "Maintenance (5yr)", Amount: 8000},
				},
			},
			want: 78000,
		},
		{
			name: "negative component for resale value",
			input: TotalCostOfOwnershipInput{
				Components: []CostComponent{
					{Label: "Acquisition", Amount: 50000},
					{Label: "Resale Value", Amount: -12000},
				},
			},
			want: 38000,
		},
		{
			name:    "empty components",
			input:   TotalCostOfOwnershipInput{Components: nil},
			wantErr: ErrEmptyComponents,
		},
		{
			name: "Inf component",
			input: TotalCostOfOwnershipInput{
				Components: []CostComponent{{Label: "x", Amount: math.Inf(1)}},
			},
			wantErr: ErrNonFiniteInput,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := TotalCostOfOwnership(tt.input)
			if err != tt.wantErr {
				t.Fatalf("TotalCostOfOwnership() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}

			if !almostEqual(got.Total, tt.want, 1e-9) {
				t.Fatalf("TotalCostOfOwnership() Total = %v, want %v", got.Total, tt.want)
			}
		})
	}
}
