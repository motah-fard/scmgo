package inventory

import (
	"math"
	"testing"
)

func TestFillRateSafetyStock(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		input     FillRateSafetyStockInput
		want      float64
		tolerance float64
		wantErr   error
	}{
		{
			name: "valid input",
			input: FillRateSafetyStockInput{
				TargetFillRate:             0.98,
				StdDevDemandDuringLeadTime: 50,
				OrderQuantity:              200,
			},
			want:      51.061943877150576,
			tolerance: 1e-6,
		},
		{
			name: "valid input, different parameters",
			input: FillRateSafetyStockInput{
				TargetFillRate:             0.95,
				StdDevDemandDuringLeadTime: 100,
				OrderQuantity:              500,
			},
			want:      34.48674639990238,
			tolerance: 1e-6,
		},
		{
			name: "extreme but achievable target near the bisection bound",
			input: FillRateSafetyStockInput{
				TargetFillRate:             0.999999999999,
				StdDevDemandDuringLeadTime: 100,
				OrderQuantity:              1,
			},
			want:      738.4662067647073,
			tolerance: 1e-3,
		},
		{
			name: "zero target fill rate",
			input: FillRateSafetyStockInput{
				TargetFillRate:             0,
				StdDevDemandDuringLeadTime: 50,
				OrderQuantity:              200,
			},
			wantErr: ErrInvalidFillRate,
		},
		{
			name: "target fill rate of one",
			input: FillRateSafetyStockInput{
				TargetFillRate:             1,
				StdDevDemandDuringLeadTime: 50,
				OrderQuantity:              200,
			},
			wantErr: ErrInvalidFillRate,
		},
		{
			name: "zero std dev demand during lead time",
			input: FillRateSafetyStockInput{
				TargetFillRate:             0.95,
				StdDevDemandDuringLeadTime: 0,
				OrderQuantity:              200,
			},
			wantErr: ErrInvalidStandardDeviation,
		},
		{
			name: "zero order quantity",
			input: FillRateSafetyStockInput{
				TargetFillRate:             0.95,
				StdDevDemandDuringLeadTime: 50,
				OrderQuantity:              0,
			},
			wantErr: ErrInvalidOrderQuantity,
		},
		{
			name: "not achievable within bisection bound",
			input: FillRateSafetyStockInput{
				TargetFillRate:             0.5,
				StdDevDemandDuringLeadTime: 10,
				OrderQuantity:              1000,
			},
			wantErr: ErrFillRateNotAchievable,
		},
		{
			name: "NaN target fill rate",
			input: FillRateSafetyStockInput{
				TargetFillRate:             math.NaN(),
				StdDevDemandDuringLeadTime: 50,
				OrderQuantity:              200,
			},
			wantErr: ErrNonFiniteInput,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := FillRateSafetyStock(tt.input)
			if err != tt.wantErr {
				t.Fatalf("FillRateSafetyStock() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}

			if !almostEqual(got, tt.want, tt.tolerance) {
				t.Fatalf("FillRateSafetyStock() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestFillRateSafetyStockRoundTrip verifies FillRateSafetyStock's result,
// fed back through the forward calculation ExpectedFillRate, recovers
// (within numerical tolerance) the original target fill rate -- a property
// that must hold regardless of the specific formula constants involved.
func TestFillRateSafetyStockRoundTrip(t *testing.T) {
	t.Parallel()

	targets := []float64{0.5, 0.8, 0.9, 0.95, 0.99, 0.999}
	sigmaL := 75.0
	orderQty := 300.0

	for _, target := range targets {
		ss, err := FillRateSafetyStock(FillRateSafetyStockInput{
			TargetFillRate:             target,
			StdDevDemandDuringLeadTime: sigmaL,
			OrderQuantity:              orderQty,
		})
		if err != nil {
			t.Fatalf("FillRateSafetyStock(%v) unexpected error: %v", target, err)
		}

		achieved, err := ExpectedFillRate(ExpectedFillRateInput{
			SafetyStockUnits:           ss,
			StdDevDemandDuringLeadTime: sigmaL,
			OrderQuantity:              orderQty,
		})
		if err != nil {
			t.Fatalf("ExpectedFillRate unexpected error: %v", err)
		}

		if !almostEqual(achieved, target, 1e-6) {
			t.Fatalf("round trip for target=%v: got achieved fill rate %v, safety stock %v", target, achieved, ss)
		}
	}
}
