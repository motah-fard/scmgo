package logistics

import (
	"math"
	"testing"
)

func TestAllocateFreightCost(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   AllocateFreightCostInput
		want    []FreightAllocationResult
		wantErr error
	}{
		{
			name: "valid input",
			input: AllocateFreightCostInput{
				TotalFreightCost: 1000,
				Items: []FreightAllocationItem{
					{ID: "a", Weight: 100},
					{ID: "b", Weight: 300},
					{ID: "c", Weight: 600},
				},
			},
			want: []FreightAllocationResult{
				{ID: "a", Weight: 100, AllocatedCost: 100},
				{ID: "b", Weight: 300, AllocatedCost: 300},
				{ID: "c", Weight: 600, AllocatedCost: 600},
			},
		},
		{
			name:    "empty items",
			input:   AllocateFreightCostInput{TotalFreightCost: 1000, Items: nil},
			wantErr: ErrEmptyItems,
		},
		{
			name: "negative total freight cost",
			input: AllocateFreightCostInput{
				TotalFreightCost: -1,
				Items:            []FreightAllocationItem{{ID: "a", Weight: 100}},
			},
			wantErr: ErrNegativeCost,
		},
		{
			name: "negative weight",
			input: AllocateFreightCostInput{
				TotalFreightCost: 1000,
				Items:            []FreightAllocationItem{{ID: "a", Weight: -1}},
			},
			wantErr: ErrNegativeWeight,
		},
		{
			name: "zero total weight",
			input: AllocateFreightCostInput{
				TotalFreightCost: 1000,
				Items:            []FreightAllocationItem{{ID: "a", Weight: 0}, {ID: "b", Weight: 0}},
			},
			wantErr: ErrZeroTotalWeight,
		},
		{
			name: "NaN weight",
			input: AllocateFreightCostInput{
				TotalFreightCost: 1000,
				Items:            []FreightAllocationItem{{ID: "a", Weight: math.NaN()}},
			},
			wantErr: ErrNonFiniteInput,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := AllocateFreightCost(tt.input)
			if err != tt.wantErr {
				t.Fatalf("AllocateFreightCost() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}

			if len(got) != len(tt.want) {
				t.Fatalf("AllocateFreightCost() length = %v, want %v", len(got), len(tt.want))
			}
			for i := range got {
				if got[i].ID != tt.want[i].ID {
					t.Fatalf("item %d: ID = %v, want %v", i, got[i].ID, tt.want[i].ID)
				}
				if !almostEqual(got[i].AllocatedCost, tt.want[i].AllocatedCost, 1e-9) {
					t.Fatalf("item %d: AllocatedCost = %v, want %v", i, got[i].AllocatedCost, tt.want[i].AllocatedCost)
				}
			}
		})
	}
}
