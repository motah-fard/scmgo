package abc

import (
	"math"
	"testing"
)

func TestClassifyVariability(t *testing.T) {
	t.Parallel()

	items := []VariabilityItem{
		{ID: "steady", MeanDemand: 100, StdDevDemand: 5},    // CV = 0.05
		{ID: "moderate", MeanDemand: 100, StdDevDemand: 25}, // CV = 0.25
		{ID: "erratic", MeanDemand: 100, StdDevDemand: 80},  // CV = 0.80
	}

	tests := []struct {
		name    string
		input   ClassifyVariabilityInput
		want    []VariabilityClassifiedItem
		wantErr error
	}{
		{
			name: "valid input",
			input: ClassifyVariabilityInput{
				Items:      items,
				XThreshold: 0.10,
				YThreshold: 0.50,
			},
			want: []VariabilityClassifiedItem{
				{ID: "steady", CoefficientOfVariation: 0.05, Class: "X"},
				{ID: "moderate", CoefficientOfVariation: 0.25, Class: "Y"},
				{ID: "erratic", CoefficientOfVariation: 0.80, Class: "Z"},
			},
		},
		{
			name: "empty items",
			input: ClassifyVariabilityInput{
				Items:      nil,
				XThreshold: 0.10,
				YThreshold: 0.50,
			},
			wantErr: ErrEmptyItems,
		},
		{
			name: "zero XThreshold",
			input: ClassifyVariabilityInput{
				Items:      items,
				XThreshold: 0,
				YThreshold: 0.50,
			},
			wantErr: ErrInvalidVariabilityThreshold,
		},
		{
			name: "YThreshold not greater than XThreshold",
			input: ClassifyVariabilityInput{
				Items:      items,
				XThreshold: 0.10,
				YThreshold: 0.10,
			},
			wantErr: ErrInvalidVariabilityThreshold,
		},
		{
			name: "zero mean demand",
			input: ClassifyVariabilityInput{
				Items:      []VariabilityItem{{ID: "x", MeanDemand: 0, StdDevDemand: 5}},
				XThreshold: 0.10,
				YThreshold: 0.50,
			},
			wantErr: ErrInvalidMeanDemand,
		},
		{
			name: "negative mean demand",
			input: ClassifyVariabilityInput{
				Items:      []VariabilityItem{{ID: "x", MeanDemand: -1, StdDevDemand: 5}},
				XThreshold: 0.10,
				YThreshold: 0.50,
			},
			wantErr: ErrInvalidMeanDemand,
		},
		{
			name: "negative std dev",
			input: ClassifyVariabilityInput{
				Items:      []VariabilityItem{{ID: "x", MeanDemand: 100, StdDevDemand: -1}},
				XThreshold: 0.10,
				YThreshold: 0.50,
			},
			wantErr: ErrNegativeStandardDeviation,
		},
		{
			name: "NaN mean demand",
			input: ClassifyVariabilityInput{
				Items:      []VariabilityItem{{ID: "x", MeanDemand: math.NaN(), StdDevDemand: 5}},
				XThreshold: 0.10,
				YThreshold: 0.50,
			},
			wantErr: ErrNonFiniteInput,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := ClassifyVariability(tt.input)
			if err != tt.wantErr {
				t.Fatalf("ClassifyVariability() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}

			if len(got) != len(tt.want) {
				t.Fatalf("ClassifyVariability() returned %d items, want %d", len(got), len(tt.want))
			}
			for i := range got {
				if got[i].ID != tt.want[i].ID {
					t.Fatalf("item %d: ID = %v, want %v", i, got[i].ID, tt.want[i].ID)
				}
				if !almostEqual(got[i].CoefficientOfVariation, tt.want[i].CoefficientOfVariation, 1e-9) {
					t.Fatalf("item %d: CoefficientOfVariation = %v, want %v", i, got[i].CoefficientOfVariation, tt.want[i].CoefficientOfVariation)
				}
				if got[i].Class != tt.want[i].Class {
					t.Fatalf("item %d: Class = %v, want %v", i, got[i].Class, tt.want[i].Class)
				}
			}
		})
	}
}
