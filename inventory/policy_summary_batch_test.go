package inventory

import "testing"

func TestBuildPolicySummaryBatch(t *testing.T) {
	t.Parallel()

	inputs := []PolicySummaryInput{
		{DailyDemand: 100, LeadTimeDays: 5, ReviewPeriodDays: 7, SafetyStockUnits: 50},
		{DailyDemand: -1, LeadTimeDays: 5, ReviewPeriodDays: 7, SafetyStockUnits: 50},
		{DailyDemand: 200, LeadTimeDays: 3, ReviewPeriodDays: 7, SafetyStockUnits: 20},
	}

	results := BuildPolicySummaryBatch(inputs)

	if len(results) != len(inputs) {
		t.Fatalf("got %d results, want %d", len(results), len(inputs))
	}

	if results[0].Err != nil {
		t.Fatalf("results[0].Err = %v, want nil", results[0].Err)
	}
	if results[0].Index != 0 {
		t.Fatalf("results[0].Index = %d, want 0", results[0].Index)
	}
	if results[0].Summary.ReorderPoint != 550 {
		t.Fatalf("results[0].Summary.ReorderPoint = %v, want 550", results[0].Summary.ReorderPoint)
	}

	if results[1].Err == nil {
		t.Fatalf("results[1].Err = nil, want an error")
	}
	if results[1].Index != 1 {
		t.Fatalf("results[1].Index = %d, want 1", results[1].Index)
	}

	if results[2].Err != nil {
		t.Fatalf("results[2].Err = %v, want nil", results[2].Err)
	}
	if results[2].Index != 2 {
		t.Fatalf("results[2].Index = %d, want 2", results[2].Index)
	}
}

func TestBuildPolicySummaryWithServiceLevelBatch(t *testing.T) {
	t.Parallel()

	inputs := []PolicySummaryServiceLevelInput{
		{DailyDemand: 100, LeadTimeDays: 5, ReviewPeriodDays: 7, DemandStdDevPerDay: 20, ServiceLevel: 0.95},
		{DailyDemand: 100, LeadTimeDays: 5, ReviewPeriodDays: 7, DemandStdDevPerDay: 20, ServiceLevel: 1.5},
	}

	results := BuildPolicySummaryWithServiceLevelBatch(inputs)

	if len(results) != len(inputs) {
		t.Fatalf("got %d results, want %d", len(results), len(inputs))
	}

	if results[0].Err != nil {
		t.Fatalf("results[0].Err = %v, want nil", results[0].Err)
	}
	if !almostEqual(results[0].Summary.ReorderPoint, 573.5600904580115, 1e-6) {
		t.Fatalf("results[0].Summary.ReorderPoint = %v, want ~573.56", results[0].Summary.ReorderPoint)
	}

	if results[1].Err == nil {
		t.Fatalf("results[1].Err = nil, want an error")
	}
}
