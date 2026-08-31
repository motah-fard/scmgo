package abc

import "sort"

// Classify performs ABC (Pareto) classification: it ranks items by value
// descending, computes each item's cumulative percentage of total value,
// and assigns a class:
//
//	A: cumulative percent <= AThreshold
//	B: AThreshold < cumulative percent <= BThreshold
//	C: cumulative percent > BThreshold
//
// Ties in value are broken by original input order. The result is returned
// in the same order as Items (not sorted by value), so it stays
// index-aligned with the input and with ClassifyVariability's result for
// the same item set.
//
// Items must contain at least one entry, all with non-negative, finite
// values summing to more than zero. AThreshold must be in (0, 1);
// BThreshold must be greater than AThreshold and at most 1.
func Classify(in ClassifyInput) ([]ClassifiedItem, error) {
	if len(in.Items) == 0 {
		return nil, ErrEmptyItems
	}
	if err := validateFinite(in.AThreshold, in.BThreshold); err != nil {
		return nil, err
	}
	if in.AThreshold <= 0 || in.AThreshold >= 1 {
		return nil, ErrInvalidThreshold
	}
	if in.BThreshold <= in.AThreshold || in.BThreshold > 1 {
		return nil, ErrInvalidThreshold
	}

	total := 0.0
	for _, item := range in.Items {
		if err := validateFinite(item.Value); err != nil {
			return nil, err
		}
		if item.Value < 0 {
			return nil, ErrNegativeValue
		}
		total += item.Value
	}
	if total <= 0 {
		return nil, ErrZeroTotalValue
	}

	rankedIndex := make([]int, len(in.Items))
	for i := range rankedIndex {
		rankedIndex[i] = i
	}
	sort.SliceStable(rankedIndex, func(a, b int) bool {
		return in.Items[rankedIndex[a]].Value > in.Items[rankedIndex[b]].Value
	})

	result := make([]ClassifiedItem, len(in.Items))
	cumulative := 0.0
	for _, idx := range rankedIndex {
		item := in.Items[idx]
		cumulative += item.Value
		cumulativePercent := cumulative / total

		class := "C"
		switch {
		case cumulativePercent <= in.AThreshold:
			class = "A"
		case cumulativePercent <= in.BThreshold:
			class = "B"
		}

		result[idx] = ClassifiedItem{
			ID:                item.ID,
			Value:             item.Value,
			CumulativePercent: cumulativePercent,
			Class:             class,
		}
	}

	return result, nil
}
