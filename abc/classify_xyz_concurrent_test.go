package abc

import (
	"errors"
	"math"
	"testing"
)

func TestClassifyVariabilityConcurrentMatchesSequential(t *testing.T) {
	t.Parallel()

	items := make([]VariabilityItem, 3000)
	for i := range items {
		items[i] = VariabilityItem{
			ID:           "sku",
			MeanDemand:   float64(50 + i%20),
			StdDevDemand: float64(i % 40),
		}
	}

	wantSeq, errSeq := ClassifyVariability(ClassifyVariabilityInput{
		Items:      items,
		XThreshold: 0.10,
		YThreshold: 0.50,
	})
	if errSeq != nil {
		t.Fatalf("sequential unexpected error: %v", errSeq)
	}

	got, err := ClassifyVariabilityConcurrent(ClassifyVariabilityInput{
		Items:      items,
		XThreshold: 0.10,
		YThreshold: 0.50,
	})
	if err != nil {
		t.Fatalf("concurrent unexpected error: %v", err)
	}

	if len(got) != len(wantSeq) {
		t.Fatalf("length = %d, want %d", len(got), len(wantSeq))
	}
	for i := range wantSeq {
		if got[i] != wantSeq[i] {
			t.Fatalf("index %d: got %+v, want %+v", i, got[i], wantSeq[i])
		}
	}
}

// TestClassifyVariabilityConcurrentReturnsFirstErrorInInputOrder checks
// the specific behavior the doc comment promises: with multiple invalid
// items scattered through the input, the returned error must be the one
// at the *lowest* index, matching what the sequential fail-fast version
// would have stopped at -- not whichever goroutine finished first.
func TestClassifyVariabilityConcurrentReturnsFirstErrorInInputOrder(t *testing.T) {
	t.Parallel()

	items := make([]VariabilityItem, 5000)
	for i := range items {
		items[i] = VariabilityItem{ID: "sku", MeanDemand: 50, StdDevDemand: 5}
	}
	// Two different invalid items, at very different positions, each
	// producing a different sentinel error.
	items[4321].MeanDemand = -1 // ErrInvalidMeanDemand, late in input order
	items[17].StdDevDemand = -1 // ErrNegativeStandardDeviation, earlier

	_, err := ClassifyVariabilityConcurrent(ClassifyVariabilityInput{
		Items:      items,
		XThreshold: 0.10,
		YThreshold: 0.50,
	})

	if !errors.Is(err, ErrNegativeStandardDeviation) {
		t.Fatalf("error = %v, want ErrNegativeStandardDeviation (the earlier invalid item, index 17)", err)
	}
}

func TestClassifyVariabilityConcurrentEmptyItems(t *testing.T) {
	t.Parallel()

	_, err := ClassifyVariabilityConcurrent(ClassifyVariabilityInput{
		Items:      nil,
		XThreshold: 0.10,
		YThreshold: 0.50,
	})
	if err != ErrEmptyItems {
		t.Fatalf("error = %v, want ErrEmptyItems", err)
	}
}

func TestClassifyVariabilityConcurrentErrorBranches(t *testing.T) {
	t.Parallel()

	validItems := []VariabilityItem{{ID: "sku", MeanDemand: 50, StdDevDemand: 5}}

	tests := []struct {
		name    string
		input   ClassifyVariabilityInput
		wantErr error
	}{
		{
			name:    "zero XThreshold",
			input:   ClassifyVariabilityInput{Items: validItems, XThreshold: 0, YThreshold: 0.50},
			wantErr: ErrInvalidVariabilityThreshold,
		},
		{
			name:    "YThreshold not greater than XThreshold",
			input:   ClassifyVariabilityInput{Items: validItems, XThreshold: 0.50, YThreshold: 0.50},
			wantErr: ErrInvalidVariabilityThreshold,
		},
		{
			name:    "NaN XThreshold",
			input:   ClassifyVariabilityInput{Items: validItems, XThreshold: math.NaN(), YThreshold: 0.50},
			wantErr: ErrNonFiniteInput,
		},
		{
			name: "zero mean demand",
			input: ClassifyVariabilityInput{
				Items:      []VariabilityItem{{ID: "sku", MeanDemand: 0, StdDevDemand: 5}},
				XThreshold: 0.10, YThreshold: 0.50,
			},
			wantErr: ErrInvalidMeanDemand,
		},
		{
			name: "NaN item field",
			input: ClassifyVariabilityInput{
				Items:      []VariabilityItem{{ID: "sku", MeanDemand: math.NaN(), StdDevDemand: 5}},
				XThreshold: 0.10, YThreshold: 0.50,
			},
			wantErr: ErrNonFiniteInput,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			_, err := ClassifyVariabilityConcurrent(tt.input)
			if err != tt.wantErr {
				t.Fatalf("error = %v, want %v", err, tt.wantErr)
			}
		})
	}
}
