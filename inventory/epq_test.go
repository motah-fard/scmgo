package inventory

import (
	"math"
	"testing"
)

func TestEPQ(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		input     EPQInput
		want      float64
		tolerance float64
		wantErr   error
	}{
		{
			name: "valid input",
			input: EPQInput{
				AnnualDemand:         10000,
				SetupCost:            50,
				HoldingCostPerUnit:   2,
				AnnualProductionRate: 40000,
			},
			want:      816.496580927726,
			tolerance: 1e-6,
		},
		{
			name: "negative annual demand",
			input: EPQInput{
				AnnualDemand:         -1,
				SetupCost:            50,
				HoldingCostPerUnit:   2,
				AnnualProductionRate: 40000,
			},
			wantErr: ErrNegativeDemand,
		},
		{
			name: "negative setup cost",
			input: EPQInput{
				AnnualDemand:         10000,
				SetupCost:            -1,
				HoldingCostPerUnit:   2,
				AnnualProductionRate: 40000,
			},
			wantErr: ErrNegativeOrderingCost,
		},
		{
			name: "zero holding cost",
			input: EPQInput{
				AnnualDemand:         10000,
				SetupCost:            50,
				HoldingCostPerUnit:   0,
				AnnualProductionRate: 40000,
			},
			wantErr: ErrInvalidHoldingCost,
		},
		{
			name: "production rate equal to demand",
			input: EPQInput{
				AnnualDemand:         10000,
				SetupCost:            50,
				HoldingCostPerUnit:   2,
				AnnualProductionRate: 10000,
			},
			wantErr: ErrInvalidProductionRate,
		},
		{
			name: "production rate below demand",
			input: EPQInput{
				AnnualDemand:         10000,
				SetupCost:            50,
				HoldingCostPerUnit:   2,
				AnnualProductionRate: 5000,
			},
			wantErr: ErrInvalidProductionRate,
		},
		{
			name: "NaN production rate",
			input: EPQInput{
				AnnualDemand:         10000,
				SetupCost:            50,
				HoldingCostPerUnit:   2,
				AnnualProductionRate: math.NaN(),
			},
			wantErr: ErrNonFiniteInput,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := EPQ(tt.input)
			if err != tt.wantErr {
				t.Fatalf("EPQ() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}

			if !almostEqual(got, tt.want, tt.tolerance) {
				t.Fatalf("EPQ() = %v, want %v", got, tt.want)
			}
		})
	}
}
