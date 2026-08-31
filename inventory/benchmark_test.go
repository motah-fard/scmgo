package inventory

import "testing"

func BenchmarkEOQ(b *testing.B) {
	in := EOQInput{
		AnnualDemand:       10000,
		OrderingCost:       50,
		HoldingCostPerUnit: 2,
	}
	for i := 0; i < b.N; i++ {
		_, _ = EOQ(in)
	}
}

func BenchmarkBuildPolicySummary(b *testing.B) {
	in := PolicySummaryInput{
		DailyDemand:      100,
		LeadTimeDays:     5,
		ReviewPeriodDays: 7,
		SafetyStockUnits: 50,
	}
	for i := 0; i < b.N; i++ {
		_, _ = BuildPolicySummary(in)
	}
}

func BenchmarkBuildPolicySummaryWithServiceLevel(b *testing.B) {
	in := PolicySummaryServiceLevelInput{
		DailyDemand:        100,
		LeadTimeDays:       5,
		ReviewPeriodDays:   7,
		DemandStdDevPerDay: 20,
		ServiceLevel:       0.95,
	}
	for i := 0; i < b.N; i++ {
		_, _ = BuildPolicySummaryWithServiceLevel(in)
	}
}

func BenchmarkBuildPolicySummaryBatch(b *testing.B) {
	inputs := make([]PolicySummaryInput, 1000)
	for i := range inputs {
		inputs[i] = PolicySummaryInput{
			DailyDemand:      100,
			LeadTimeDays:     5,
			ReviewPeriodDays: 7,
			SafetyStockUnits: 50,
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = BuildPolicySummaryBatch(inputs)
	}
}
