package inventory

import (
	"math"
	"testing"
)

func TestNewsvendor(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		input     NewsvendorInput
		want      NewsvendorResult
		tolerance float64
		wantErr   error
	}{
		{
			name: "valid input",
			input: NewsvendorInput{
				MeanDemand:          500,
				StdDevDemand:        100,
				UnderageCostPerUnit: 18,
				OverageCostPerUnit:  7,
			},
			want: NewsvendorResult{
				OrderQuantity: 558.2841507271216,
				CriticalRatio: 0.72,
			},
			tolerance: 1e-6,
		},
		{
			name: "balanced costs gives critical ratio 0.5 and Q = mean",
			input: NewsvendorInput{
				MeanDemand:          200,
				StdDevDemand:        30,
				UnderageCostPerUnit: 10,
				OverageCostPerUnit:  10,
			},
			want: NewsvendorResult{
				OrderQuantity: 200,
				CriticalRatio: 0.5,
			},
			tolerance: 1e-6,
		},
		{
			name: "negative mean demand",
			input: NewsvendorInput{
				MeanDemand:          -1,
				StdDevDemand:        100,
				UnderageCostPerUnit: 18,
				OverageCostPerUnit:  7,
			},
			wantErr: ErrNegativeDemand,
		},
		{
			name: "negative std dev demand",
			input: NewsvendorInput{
				MeanDemand:          500,
				StdDevDemand:        -1,
				UnderageCostPerUnit: 18,
				OverageCostPerUnit:  7,
			},
			wantErr: ErrNegativeStandardDeviation,
		},
		{
			name: "zero underage cost",
			input: NewsvendorInput{
				MeanDemand:          500,
				StdDevDemand:        100,
				UnderageCostPerUnit: 0,
				OverageCostPerUnit:  7,
			},
			wantErr: ErrInvalidUnderageCost,
		},
		{
			name: "zero overage cost",
			input: NewsvendorInput{
				MeanDemand:          500,
				StdDevDemand:        100,
				UnderageCostPerUnit: 18,
				OverageCostPerUnit:  0,
			},
			wantErr: ErrInvalidOverageCost,
		},
		{
			name: "NaN underage cost",
			input: NewsvendorInput{
				MeanDemand:          500,
				StdDevDemand:        100,
				UnderageCostPerUnit: math.NaN(),
				OverageCostPerUnit:  7,
			},
			wantErr: ErrNonFiniteInput,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := Newsvendor(tt.input)
			if err != tt.wantErr {
				t.Fatalf("Newsvendor() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}

			if !almostEqual(got.OrderQuantity, tt.want.OrderQuantity, tt.tolerance) {
				t.Fatalf("Newsvendor() OrderQuantity = %v, want %v", got.OrderQuantity, tt.want.OrderQuantity)
			}
			if !almostEqual(got.CriticalRatio, tt.want.CriticalRatio, tt.tolerance) {
				t.Fatalf("Newsvendor() CriticalRatio = %v, want %v", got.CriticalRatio, tt.want.CriticalRatio)
			}
		})
	}
}
