package procurement

import (
	"math"
	"testing"
)

func TestPurchasePriceVariance(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   PurchasePriceVarianceInput
		want    float64
		wantErr error
	}{
		{
			name: "favorable variance",
			input: PurchasePriceVarianceInput{
				StandardCost: 10,
				ActualCost:   9,
				Quantity:     1000,
			},
			want: 1000,
		},
		{
			name: "unfavorable variance",
			input: PurchasePriceVarianceInput{
				StandardCost: 10,
				ActualCost:   12,
				Quantity:     1000,
			},
			want: -2000,
		},
		{
			name: "zero variance",
			input: PurchasePriceVarianceInput{
				StandardCost: 10,
				ActualCost:   10,
				Quantity:     1000,
			},
			want: 0,
		},
		{
			name: "negative standard cost",
			input: PurchasePriceVarianceInput{
				StandardCost: -1,
				ActualCost:   10,
				Quantity:     1000,
			},
			wantErr: ErrNegativeCost,
		},
		{
			name: "negative actual cost",
			input: PurchasePriceVarianceInput{
				StandardCost: 10,
				ActualCost:   -1,
				Quantity:     1000,
			},
			wantErr: ErrNegativeCost,
		},
		{
			name: "negative quantity",
			input: PurchasePriceVarianceInput{
				StandardCost: 10,
				ActualCost:   9,
				Quantity:     -1,
			},
			wantErr: ErrNegativeQuantity,
		},
		{
			name: "NaN quantity",
			input: PurchasePriceVarianceInput{
				StandardCost: 10,
				ActualCost:   9,
				Quantity:     math.NaN(),
			},
			wantErr: ErrNonFiniteInput,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := PurchasePriceVariance(tt.input)
			if err != tt.wantErr {
				t.Fatalf("PurchasePriceVariance() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}

			if !almostEqual(got, tt.want, 1e-9) {
				t.Fatalf("PurchasePriceVariance() = %v, want %v", got, tt.want)
			}
		})
	}
}
