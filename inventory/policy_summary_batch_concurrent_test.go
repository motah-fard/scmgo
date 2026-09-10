package inventory

import "testing"

// sameError compares by message rather than by identity: errors.Join
// (used by validatePolicySummaryInput and friends to wrap a sentinel)
// allocates a distinct value on every call, even for identical inputs,
// so two independently computed errors.Join results are never == to
// each other despite being semantically the same error. Comparing the
// message is the right check here, not err != err.
func sameError(a, b error) bool {
	if a == nil || b == nil {
		return a == b
	}
	return a.Error() == b.Error()
}

func TestBuildPolicySummaryBatchConcurrentMatchesSequential(t *testing.T) {
	t.Parallel()

	inputs := make([]PolicySummaryInput, 2000)
	for i := range inputs {
		inputs[i] = PolicySummaryInput{
			DailyDemand:      float64(100 + i),
			LeadTimeDays:     5,
			ReviewPeriodDays: 7,
			SafetyStockUnits: 50,
		}
	}
	// Sprinkle in invalid rows to confirm per-item errors survive
	// concurrent execution the same way they do sequentially.
	inputs[7].DailyDemand = -1
	inputs[999].LeadTimeDays = -1
	inputs[1500].SafetyStockUnits = -1

	sequential := BuildPolicySummaryBatch(inputs)
	concurrent := BuildPolicySummaryBatchConcurrent(inputs)

	if len(concurrent) != len(sequential) {
		t.Fatalf("length = %d, want %d", len(concurrent), len(sequential))
	}
	for i := range sequential {
		if concurrent[i].Index != sequential[i].Index {
			t.Fatalf("index %d: Index = %v, want %v", i, concurrent[i].Index, sequential[i].Index)
		}
		if !sameError(concurrent[i].Err, sequential[i].Err) {
			t.Fatalf("index %d: Err = %v, want %v", i, concurrent[i].Err, sequential[i].Err)
		}
		if concurrent[i].Summary != sequential[i].Summary {
			t.Fatalf("index %d: Summary = %+v, want %+v", i, concurrent[i].Summary, sequential[i].Summary)
		}
	}
}

func TestBuildPolicySummaryWithServiceLevelBatchConcurrentMatchesSequential(t *testing.T) {
	t.Parallel()

	inputs := make([]PolicySummaryServiceLevelInput, 2000)
	for i := range inputs {
		inputs[i] = PolicySummaryServiceLevelInput{
			DailyDemand:        float64(100 + i),
			LeadTimeDays:       5,
			ReviewPeriodDays:   7,
			DemandStdDevPerDay: 20,
			ServiceLevel:       0.95,
		}
	}
	inputs[42].ServiceLevel = 1 // invalid
	inputs[1234].DemandStdDevPerDay = -1

	sequential := BuildPolicySummaryWithServiceLevelBatch(inputs)
	concurrent := BuildPolicySummaryWithServiceLevelBatchConcurrent(inputs)

	if len(concurrent) != len(sequential) {
		t.Fatalf("length = %d, want %d", len(concurrent), len(sequential))
	}
	for i := range sequential {
		if concurrent[i].Index != sequential[i].Index {
			t.Fatalf("index %d: Index = %v, want %v", i, concurrent[i].Index, sequential[i].Index)
		}
		if !sameError(concurrent[i].Err, sequential[i].Err) {
			t.Fatalf("index %d: Err = %v, want %v", i, concurrent[i].Err, sequential[i].Err)
		}
		if concurrent[i].Summary != sequential[i].Summary {
			t.Fatalf("index %d: Summary = %+v, want %+v", i, concurrent[i].Summary, sequential[i].Summary)
		}
	}
}

func TestBuildPolicySummaryBatchConcurrentEmpty(t *testing.T) {
	t.Parallel()

	got := BuildPolicySummaryBatchConcurrent(nil)
	if len(got) != 0 {
		t.Fatalf("BuildPolicySummaryBatchConcurrent(nil) = %v, want empty", got)
	}
}
