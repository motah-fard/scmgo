package abc

import (
	"math"
	"testing"
)

func almostEqual(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}

func TestClassify(t *testing.T) {
	t.Parallel()

	items := []Item{
		{ID: "A1", Value: 1000},
		{ID: "A2", Value: 500},
		{ID: "A3", Value: 300},
		{ID: "A4", Value: 100},
		{ID: "A5", Value: 100},
	}

	tests := []struct {
		name    string
		input   ClassifyInput
		want    []ClassifiedItem
		wantErr error
	}{
		{
			name: "valid input, preserves input order",
			input: ClassifyInput{
				Items:      items,
				AThreshold: 0.80,
				BThreshold: 0.95,
			},
			want: []ClassifiedItem{
				{ID: "A1", Value: 1000, CumulativePercent: 0.5, Class: "A"},
				{ID: "A2", Value: 500, CumulativePercent: 0.75, Class: "A"},
				{ID: "A3", Value: 300, CumulativePercent: 0.9, Class: "B"},
				{ID: "A4", Value: 100, CumulativePercent: 0.95, Class: "B"},
				{ID: "A5", Value: 100, CumulativePercent: 1.0, Class: "C"},
			},
		},
		{
			name: "empty items",
			input: ClassifyInput{
				Items:      nil,
				AThreshold: 0.80,
				BThreshold: 0.95,
			},
			wantErr: ErrEmptyItems,
		},
		{
			name: "zero AThreshold",
			input: ClassifyInput{
				Items:      items,
				AThreshold: 0,
				BThreshold: 0.95,
			},
			wantErr: ErrInvalidThreshold,
		},
		{
			name: "AThreshold >= 1",
			input: ClassifyInput{
				Items:      items,
				AThreshold: 1,
				BThreshold: 0.95,
			},
			wantErr: ErrInvalidThreshold,
		},
		{
			name: "BThreshold not greater than AThreshold",
			input: ClassifyInput{
				Items:      items,
				AThreshold: 0.80,
				BThreshold: 0.80,
			},
			wantErr: ErrInvalidThreshold,
		},
		{
			name: "BThreshold above 1",
			input: ClassifyInput{
				Items:      items,
				AThreshold: 0.80,
				BThreshold: 1.01,
			},
			wantErr: ErrInvalidThreshold,
		},
		{
			name: "negative value",
			input: ClassifyInput{
				Items:      []Item{{ID: "x", Value: -1}},
				AThreshold: 0.80,
				BThreshold: 0.95,
			},
			wantErr: ErrNegativeValue,
		},
		{
			name: "zero total value",
			input: ClassifyInput{
				Items:      []Item{{ID: "x", Value: 0}, {ID: "y", Value: 0}},
				AThreshold: 0.80,
				BThreshold: 0.95,
			},
			wantErr: ErrZeroTotalValue,
		},
		{
			name: "NaN value",
			input: ClassifyInput{
				Items:      []Item{{ID: "x", Value: math.NaN()}},
				AThreshold: 0.80,
				BThreshold: 0.95,
			},
			wantErr: ErrNonFiniteInput,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := Classify(tt.input)
			if err != tt.wantErr {
				t.Fatalf("Classify() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}

			if len(got) != len(tt.want) {
				t.Fatalf("Classify() returned %d items, want %d", len(got), len(tt.want))
			}
			for i := range got {
				if got[i].ID != tt.want[i].ID {
					t.Fatalf("item %d: ID = %v, want %v", i, got[i].ID, tt.want[i].ID)
				}
				if !almostEqual(got[i].Value, tt.want[i].Value, 1e-9) {
					t.Fatalf("item %d: Value = %v, want %v", i, got[i].Value, tt.want[i].Value)
				}
				if !almostEqual(got[i].CumulativePercent, tt.want[i].CumulativePercent, 1e-9) {
					t.Fatalf("item %d: CumulativePercent = %v, want %v", i, got[i].CumulativePercent, tt.want[i].CumulativePercent)
				}
				if got[i].Class != tt.want[i].Class {
					t.Fatalf("item %d: Class = %v, want %v", i, got[i].Class, tt.want[i].Class)
				}
			}
		})
	}
}

func TestClassifyTiesBrokenByInputOrder(t *testing.T) {
	t.Parallel()

	// Three items tied at the same value: cumulative order (and thus each
	// item's own CumulativePercent) must follow input order deterministically.
	items := []Item{
		{ID: "first", Value: 100},
		{ID: "second", Value: 100},
		{ID: "third", Value: 100},
	}

	got, err := Classify(ClassifyInput{Items: items, AThreshold: 0.4, BThreshold: 0.7})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	wantCumPercent := []float64{1.0 / 3, 2.0 / 3, 1.0}
	for i, item := range got {
		if !almostEqual(item.CumulativePercent, wantCumPercent[i], 1e-9) {
			t.Fatalf("item %d (%s): CumulativePercent = %v, want %v", i, item.ID, item.CumulativePercent, wantCumPercent[i])
		}
	}
}
