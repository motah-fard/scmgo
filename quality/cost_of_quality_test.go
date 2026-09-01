package quality

import (
	"math"
	"testing"
)

func TestCostOfQuality(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   CostOfQualityInput
		want    CostOfQualityResult
		wantErr error
	}{
		{
			name: "valid input",
			input: CostOfQualityInput{
				PreventionCost:      5000,
				AppraisalCost:       8000,
				InternalFailureCost: 12000,
				ExternalFailureCost: 20000,
			},
			want: CostOfQualityResult{
				Total:               45000,
				PreventionCost:      5000,
				AppraisalCost:       8000,
				InternalFailureCost: 12000,
				ExternalFailureCost: 20000,
				ConformanceCost:     13000,
				NonConformanceCost:  32000,
			},
		},
		{
			name: "negative prevention cost",
			input: CostOfQualityInput{
				PreventionCost:      -1,
				AppraisalCost:       8000,
				InternalFailureCost: 12000,
				ExternalFailureCost: 20000,
			},
			wantErr: ErrNegativeCost,
		},
		{
			name: "NaN external failure cost",
			input: CostOfQualityInput{
				PreventionCost:      5000,
				AppraisalCost:       8000,
				InternalFailureCost: 12000,
				ExternalFailureCost: math.NaN(),
			},
			wantErr: ErrNonFiniteInput,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := CostOfQuality(tt.input)
			if err != tt.wantErr {
				t.Fatalf("CostOfQuality() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}

			if got != tt.want {
				t.Fatalf("CostOfQuality() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
