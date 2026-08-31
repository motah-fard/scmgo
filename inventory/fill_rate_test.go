package inventory

import (
	"math"
	"testing"
)

func TestExpectedFillRate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		input     ExpectedFillRateInput
		want      float64
		tolerance float64
		wantErr   error
	}{
		{
			name: "valid input",
			input: ExpectedFillRateInput{
				SafetyStockUnits:           50,
				StdDevDemandDuringLeadTime: 50,
				OrderQuantity:              200,
			},
			want:      0.9791711323530784,
			tolerance: 1e-9,
		},
		{
			name: "zero safety stock",
			input: ExpectedFillRateInput{
				SafetyStockUnits:           0,
				StdDevDemandDuringLeadTime: 50,
				OrderQuantity:              200,
			},
			want:      0.9002644298996418,
			tolerance: 1e-9,
		},
		{
			name: "negative safety stock is valid and lowers fill rate",
			input: ExpectedFillRateInput{
				SafetyStockUnits:           -10,
				StdDevDemandDuringLeadTime: 50,
				OrderQuantity:              200,
			},
			want:      0.8732763410341808,
			tolerance: 1e-9,
		},
		{
			name: "zero std dev demand during lead time",
			input: ExpectedFillRateInput{
				SafetyStockUnits:           50,
				StdDevDemandDuringLeadTime: 0,
				OrderQuantity:              200,
			},
			wantErr: ErrInvalidStandardDeviation,
		},
		{
			name: "zero order quantity",
			input: ExpectedFillRateInput{
				SafetyStockUnits:           50,
				StdDevDemandDuringLeadTime: 50,
				OrderQuantity:              0,
			},
			wantErr: ErrInvalidOrderQuantity,
		},
		{
			name: "NaN order quantity",
			input: ExpectedFillRateInput{
				SafetyStockUnits:           50,
				StdDevDemandDuringLeadTime: 50,
				OrderQuantity:              math.NaN(),
			},
			wantErr: ErrNonFiniteInput,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := ExpectedFillRate(tt.input)
			if err != tt.wantErr {
				t.Fatalf("ExpectedFillRate() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}

			if !almostEqual(got, tt.want, tt.tolerance) {
				t.Fatalf("ExpectedFillRate() = %v, want %v", got, tt.want)
			}
		})
	}
}
