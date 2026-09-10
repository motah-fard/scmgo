package inventory

import "github.com/motah-fard/scmgo/internal/parallel"

// BuildPolicySummaryBatchConcurrent behaves exactly like
// BuildPolicySummaryBatch -- same inputs in, same results out, in the
// same order -- but computes the batch across multiple goroutines
// (bounded by runtime.GOMAXPROCS(0)) instead of one item at a time.
//
// BuildPolicySummary is cheap (tens of nanoseconds); for small and
// medium batches, the overhead of spawning goroutines costs more than
// it saves, and BuildPolicySummaryBatch is faster. On this project's
// development hardware (a multi-core arm64 machine), the crossover
// landed around 2,000-3,000 items; below that, the plain sequential
// BuildPolicySummaryBatch was faster. That number is not a promise about
// your hardware or your build's optimization level -- see
// BenchmarkBuildPolicySummaryBatchSizes and
// BenchmarkBuildPolicySummaryBatchConcurrentSizes and run them yourself
// before switching over a real batch size.
func BuildPolicySummaryBatchConcurrent(inputs []PolicySummaryInput) []PolicySummaryBatchResult {
	return parallel.Map(inputs, func(index int, input PolicySummaryInput) PolicySummaryBatchResult {
		summary, err := BuildPolicySummary(input)
		return PolicySummaryBatchResult{
			Index:   index,
			Summary: summary,
			Err:     err,
		}
	})
}

// BuildPolicySummaryWithServiceLevelBatchConcurrent is
// BuildPolicySummaryWithServiceLevelBatch computed across multiple
// goroutines instead of one item at a time. See
// BuildPolicySummaryBatchConcurrent's doc comment for when this is (and
// isn't) worth reaching for.
func BuildPolicySummaryWithServiceLevelBatchConcurrent(inputs []PolicySummaryServiceLevelInput) []PolicySummaryBatchResult {
	return parallel.Map(inputs, func(index int, input PolicySummaryServiceLevelInput) PolicySummaryBatchResult {
		summary, err := BuildPolicySummaryWithServiceLevel(input)
		return PolicySummaryBatchResult{
			Index:   index,
			Summary: summary,
			Err:     err,
		}
	})
}
