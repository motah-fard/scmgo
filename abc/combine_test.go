package abc

import "testing"

func TestCombine(t *testing.T) {
	t.Parallel()

	abcResults := []ClassifiedItem{
		{ID: "sku1", Class: "A"},
		{ID: "sku2", Class: "B"},
		{ID: "sku3", Class: "C"},
	}
	xyzResults := []VariabilityClassifiedItem{
		{ID: "sku1", Class: "X"},
		{ID: "sku3", Class: "Z"},
		{ID: "sku4", Class: "Y"}, // not present in abcResults, should be dropped
	}

	got := Combine(abcResults, xyzResults)

	want := []CombinedClass{
		{ID: "sku1", ABCClass: "A", XYZClass: "X"},
		{ID: "sku3", ABCClass: "C", XYZClass: "Z"},
	}

	if len(got) != len(want) {
		t.Fatalf("Combine() returned %d items, want %d: %+v", len(got), len(want), got)
	}
	for i := range got {
		if got[i] != want[i] {
			t.Fatalf("item %d: got %+v, want %+v", i, got[i], want[i])
		}
	}
}

func TestCombineEmptyInputs(t *testing.T) {
	t.Parallel()

	if got := Combine(nil, nil); len(got) != 0 {
		t.Fatalf("Combine(nil, nil) = %+v, want empty", got)
	}
}
