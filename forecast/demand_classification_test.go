package forecast

import (
	"math"
	"testing"
)

func TestClassifyDemandPattern(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		input     DemandClassificationInput
		wantADI   float64
		wantCV2   float64
		wantClass string
		tolerance float64
		wantErr   error
	}{
		{
			name:      "intermittent",
			input:     DemandClassificationInput{History: []float64{0, 0, 5, 0, 0, 0, 3, 0, 4, 0}},
			wantADI:   3.3333333333333335,
			wantCV2:   0.041666666666666664,
			wantClass: "intermittent",
			tolerance: 1e-9,
		},
		{
			name:      "smooth",
			input:     DemandClassificationInput{History: []float64{10, 12, 9, 11, 10, 13, 8, 12}},
			wantADI:   1.0,
			wantCV2:   0.022006920415224913,
			wantClass: "smooth",
			tolerance: 1e-9,
		},
		{
			name: "lumpy",
			input: DemandClassificationInput{
				History: []float64{0, 0, 0, 0, 0, 0, 0, 0, 50, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0},
			},
			wantADI:   10.0,
			wantCV2:   0.8520710059171598,
			wantClass: "lumpy",
			tolerance: 1e-9,
		},
		{
			name:      "single non-zero period has zero CV2",
			input:     DemandClassificationInput{History: []float64{0, 0, 5}},
			wantADI:   3.0,
			wantCV2:   0,
			wantClass: "intermittent",
			tolerance: 1e-9,
		},
		{
			name:    "empty history",
			input:   DemandClassificationInput{History: nil},
			wantErr: ErrEmptyHistory,
		},
		{
			name:    "no non-zero demand",
			input:   DemandClassificationInput{History: []float64{0, 0, 0}},
			wantErr: ErrNoNonZeroDemand,
		},
		{
			name:    "negative demand",
			input:   DemandClassificationInput{History: []float64{0, -5}},
			wantErr: ErrNegativeDemand,
		},
		{
			name:    "NaN in history",
			input:   DemandClassificationInput{History: []float64{0, math.NaN()}},
			wantErr: ErrNonFiniteInput,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := ClassifyDemandPattern(tt.input)
			if err != tt.wantErr {
				t.Fatalf("ClassifyDemandPattern() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}

			if !almostEqual(got.ADI, tt.wantADI, tt.tolerance) {
				t.Fatalf("ClassifyDemandPattern() ADI = %v, want %v", got.ADI, tt.wantADI)
			}
			if !almostEqual(got.CV2, tt.wantCV2, tt.tolerance) {
				t.Fatalf("ClassifyDemandPattern() CV2 = %v, want %v", got.CV2, tt.wantCV2)
			}
			if got.Class != tt.wantClass {
				t.Fatalf("ClassifyDemandPattern() Class = %v, want %v", got.Class, tt.wantClass)
			}
		})
	}
}
