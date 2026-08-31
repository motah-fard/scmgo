package inventory

// BuildPolicySummaryBatch runs BuildPolicySummary over a slice of inputs,
// useful for computing policy summaries across a list of SKUs in one call.
// Unlike BuildPolicySummary, an invalid item does not prevent the rest of
// the batch from being computed: each result carries its own error.
func BuildPolicySummaryBatch(inputs []PolicySummaryInput) []PolicySummaryBatchResult {
	results := make([]PolicySummaryBatchResult, len(inputs))
	for i, input := range inputs {
		summary, err := BuildPolicySummary(input)
		results[i] = PolicySummaryBatchResult{
			Index:   i,
			Summary: summary,
			Err:     err,
		}
	}
	return results
}

// BuildPolicySummaryWithServiceLevelBatch runs BuildPolicySummaryWithServiceLevel
// over a slice of inputs, useful for computing service-level-based policy
// summaries across a list of SKUs in one call. Unlike
// BuildPolicySummaryWithServiceLevel, an invalid item does not prevent the
// rest of the batch from being computed: each result carries its own error.
func BuildPolicySummaryWithServiceLevelBatch(inputs []PolicySummaryServiceLevelInput) []PolicySummaryBatchResult {
	results := make([]PolicySummaryBatchResult, len(inputs))
	for i, input := range inputs {
		summary, err := BuildPolicySummaryWithServiceLevel(input)
		results[i] = PolicySummaryBatchResult{
			Index:   i,
			Summary: summary,
			Err:     err,
		}
	}
	return results
}
