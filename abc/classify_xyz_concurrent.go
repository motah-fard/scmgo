package abc

import "github.com/motah-fard/scmgo/internal/parallel"

// ClassifyVariabilityConcurrent behaves exactly like ClassifyVariability
// -- same inputs in, same result (or the same error) out -- but computes
// each item's coefficient of variation across multiple goroutines
// (bounded by runtime.GOMAXPROCS(0)) instead of one item at a time.
//
// ClassifyVariability fails fast: on the first invalid item it stops and
// returns that one error for the whole call, not a per-item result. This
// validates every item sequentially first, in input order, before
// parallelizing anything -- so the error returned (if any) is exactly
// the one ClassifyVariability would have stopped at, and no goroutine is
// ever spawned to do work on an input that turns out to be invalid.
// Only once every item is known valid does the (now infallible)
// per-item division run in parallel.
//
// An earlier version of this function computed validation and the
// division in the same parallel pass, carrying a per-item error
// alongside each result. Benchmarking that version (see git history)
// showed it never actually beat ClassifyVariability, even at 100,000
// items -- the bookkeeping to carry a possible error through the
// parallel path cost as much as the parallelism saved. This version's
// separation of "validate" from "compute" removed that cost: on this
// project's development hardware (a multi-core arm64 machine),
// benchmarking showed a real crossover around several tens of thousands
// of items, with ClassifyVariability still faster below that. See
// BenchmarkClassifyVariability and BenchmarkClassifyVariabilityConcurrent
// for the numbers on your own hardware; don't assume them from this doc
// comment, and prefer ClassifyVariability unless you've measured a large
// item set as an actual bottleneck.
func ClassifyVariabilityConcurrent(in ClassifyVariabilityInput) ([]VariabilityClassifiedItem, error) {
	if len(in.Items) == 0 {
		return nil, ErrEmptyItems
	}
	if err := validateFinite(in.XThreshold, in.YThreshold); err != nil {
		return nil, err
	}
	if in.XThreshold <= 0 {
		return nil, ErrInvalidVariabilityThreshold
	}
	if in.YThreshold <= in.XThreshold {
		return nil, ErrInvalidVariabilityThreshold
	}

	for _, item := range in.Items {
		if err := validateFinite(item.MeanDemand, item.StdDevDemand); err != nil {
			return nil, err
		}
		if item.MeanDemand <= 0 {
			return nil, ErrInvalidMeanDemand
		}
		if item.StdDevDemand < 0 {
			return nil, ErrNegativeStandardDeviation
		}
	}

	result := parallel.Map(in.Items, func(_ int, item VariabilityItem) VariabilityClassifiedItem {
		cv := item.StdDevDemand / item.MeanDemand

		class := "Z"
		switch {
		case cv <= in.XThreshold:
			class = "X"
		case cv <= in.YThreshold:
			class = "Y"
		}

		return VariabilityClassifiedItem{
			ID:                     item.ID,
			CoefficientOfVariation: cv,
			Class:                  class,
		}
	})

	return result, nil
}
