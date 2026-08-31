package inventory

import (
	"math"
	"testing"
)

func TestEOQWithQuantityDiscounts(t *testing.T) {
	t.Parallel()

	// Classic textbook example: D=10000/yr, S=$20/order, holding rate 20%.
	// Tiers: 1-499 @ $5.00, 500-999 @ $4.50, 1000+ @ $3.90.
	// Known answer: order 1000 units at $3.90, total annual cost $39,590.
	classicTiers := []QuantityDiscountTier{
		{MinQuantity: 1, UnitPrice: 5.00},
		{MinQuantity: 500, UnitPrice: 4.50},
		{MinQuantity: 1000, UnitPrice: 3.90},
	}

	tests := []struct {
		name      string
		input     EOQWithQuantityDiscountsInput
		want      EOQWithQuantityDiscountsResult
		tolerance float64
		wantErr   error
	}{
		{
			name: "classic textbook example",
			input: EOQWithQuantityDiscountsInput{
				AnnualDemand:    10000,
				OrderingCost:    20,
				HoldingCostRate: 0.20,
				Tiers:           classicTiers,
			},
			want: EOQWithQuantityDiscountsResult{
				OrderQuantity:   1000,
				UnitPrice:       3.90,
				TotalAnnualCost: 39590,
			},
			tolerance: 1e-6,
		},
		{
			name: "single tier behaves like plain EOQ",
			input: EOQWithQuantityDiscountsInput{
				AnnualDemand:    10000,
				OrderingCost:    50,
				HoldingCostRate: 0.20,
				Tiers: []QuantityDiscountTier{
					{MinQuantity: 1, UnitPrice: 10},
				},
			},
			want: EOQWithQuantityDiscountsResult{
				// EOQ = sqrt(2*10000*50/(0.2*10)) = sqrt(500000) ≈ 707.1067812
				OrderQuantity:   707.1067811865476,
				UnitPrice:       10,
				TotalAnnualCost: 101414.2135623731,
			},
			tolerance: 1e-6,
		},
		{
			name: "empty tiers",
			input: EOQWithQuantityDiscountsInput{
				AnnualDemand:    10000,
				OrderingCost:    20,
				HoldingCostRate: 0.20,
				Tiers:           nil,
			},
			wantErr: ErrEmptyDiscountTiers,
		},
		{
			name: "first tier min quantity not above zero",
			input: EOQWithQuantityDiscountsInput{
				AnnualDemand:    10000,
				OrderingCost:    20,
				HoldingCostRate: 0.20,
				Tiers: []QuantityDiscountTier{
					{MinQuantity: 0, UnitPrice: 5},
				},
			},
			wantErr: ErrInvalidDiscountTiers,
		},
		{
			name: "tiers not strictly ascending",
			input: EOQWithQuantityDiscountsInput{
				AnnualDemand:    10000,
				OrderingCost:    20,
				HoldingCostRate: 0.20,
				Tiers: []QuantityDiscountTier{
					{MinQuantity: 1, UnitPrice: 5},
					{MinQuantity: 1, UnitPrice: 4.5},
				},
			},
			wantErr: ErrInvalidDiscountTiers,
		},
		{
			name: "price increases with quantity",
			input: EOQWithQuantityDiscountsInput{
				AnnualDemand:    10000,
				OrderingCost:    20,
				HoldingCostRate: 0.20,
				Tiers: []QuantityDiscountTier{
					{MinQuantity: 1, UnitPrice: 5},
					{MinQuantity: 500, UnitPrice: 5.5},
				},
			},
			wantErr: ErrInvalidDiscountTiers,
		},
		{
			name: "zero unit price",
			input: EOQWithQuantityDiscountsInput{
				AnnualDemand:    10000,
				OrderingCost:    20,
				HoldingCostRate: 0.20,
				Tiers: []QuantityDiscountTier{
					{MinQuantity: 1, UnitPrice: 0},
				},
			},
			wantErr: ErrInvalidDiscountTiers,
		},
		{
			name: "negative annual demand",
			input: EOQWithQuantityDiscountsInput{
				AnnualDemand:    -1,
				OrderingCost:    20,
				HoldingCostRate: 0.20,
				Tiers:           classicTiers,
			},
			wantErr: ErrNegativeDemand,
		},
		{
			name: "zero holding cost rate",
			input: EOQWithQuantityDiscountsInput{
				AnnualDemand:    10000,
				OrderingCost:    20,
				HoldingCostRate: 0,
				Tiers:           classicTiers,
			},
			wantErr: ErrInvalidHoldingCost,
		},
		{
			name: "NaN in tier price",
			input: EOQWithQuantityDiscountsInput{
				AnnualDemand:    10000,
				OrderingCost:    20,
				HoldingCostRate: 0.20,
				Tiers: []QuantityDiscountTier{
					{MinQuantity: 1, UnitPrice: math.NaN()},
				},
			},
			wantErr: ErrNonFiniteInput,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := EOQWithQuantityDiscounts(tt.input)
			if err != tt.wantErr {
				t.Fatalf("EOQWithQuantityDiscounts() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}

			if !almostEqual(got.OrderQuantity, tt.want.OrderQuantity, tt.tolerance) {
				t.Fatalf("EOQWithQuantityDiscounts() OrderQuantity = %v, want %v", got.OrderQuantity, tt.want.OrderQuantity)
			}
			if !almostEqual(got.UnitPrice, tt.want.UnitPrice, tt.tolerance) {
				t.Fatalf("EOQWithQuantityDiscounts() UnitPrice = %v, want %v", got.UnitPrice, tt.want.UnitPrice)
			}
			if !almostEqual(got.TotalAnnualCost, tt.want.TotalAnnualCost, tt.tolerance) {
				t.Fatalf("EOQWithQuantityDiscounts() TotalAnnualCost = %v, want %v", got.TotalAnnualCost, tt.want.TotalAnnualCost)
			}
		})
	}
}
