package inventory

import (
	"math"
	"testing"
)

// These fuzz targets assert two invariants that must hold for any float64
// input whatsoever:
//
//  1. The function must never panic.
//  2. If it returns a nil error, the result must never contain NaN.
//
// (2) exists because the NaN/Inf validation gap fixed in this package (see
// validateFinite) was exactly the kind of bug fuzzing is built to catch —
// a plain "v < 0" check silently passing NaN through instead of failing,
// and returning a NaN result with a nil error. Every function fuzzed here
// only ever combines validated non-negative/positive finite values with
// addition, multiplication, division by a strictly-positive quantity, or
// Erfinv over an argument strictly inside (-1, 1) — none of which can
// produce NaN from finite input, so "not NaN" is a safe invariant to
// assert (unlike "not Inf", which a sufficiently large but individually
// valid input could legitimately overflow to).

func FuzzEOQ(f *testing.F) {
	f.Add(10000.0, 50.0, 2.0)
	f.Fuzz(func(t *testing.T, annualDemand, orderingCost, holdingCostPerUnit float64) {
		result, err := EOQ(EOQInput{
			AnnualDemand:       annualDemand,
			OrderingCost:       orderingCost,
			HoldingCostPerUnit: holdingCostPerUnit,
		})
		if err == nil && math.IsNaN(result) {
			t.Fatalf("EOQ returned NaN with nil error for input %+v", EOQInput{annualDemand, orderingCost, holdingCostPerUnit})
		}
	})
}

func FuzzZScoreForServiceLevel(f *testing.F) {
	f.Add(0.95)
	f.Fuzz(func(t *testing.T, serviceLevel float64) {
		result, err := ZScoreForServiceLevel(serviceLevel)
		if err == nil && math.IsNaN(result) {
			t.Fatalf("ZScoreForServiceLevel returned NaN with nil error for serviceLevel=%v", serviceLevel)
		}
	})
}

func FuzzBuildPolicySummary(f *testing.F) {
	f.Add(100.0, 5.0, 7.0, 50.0)
	f.Fuzz(func(t *testing.T, dailyDemand, leadTimeDays, reviewPeriodDays, safetyStockUnits float64) {
		summary, err := BuildPolicySummary(PolicySummaryInput{
			DailyDemand:      dailyDemand,
			LeadTimeDays:     leadTimeDays,
			ReviewPeriodDays: reviewPeriodDays,
			SafetyStockUnits: safetyStockUnits,
		})
		if err != nil {
			return
		}
		for _, v := range []float64{
			summary.ExpectedDemandDuringLeadTime, summary.SafetyStockUnits,
			summary.ReorderPoint, summary.TargetInventoryLevel,
			summary.MinLevel, summary.MaxLevel,
		} {
			if math.IsNaN(v) {
				t.Fatalf("BuildPolicySummary returned NaN field with nil error: %+v", summary)
			}
		}
	})
}

func FuzzBuildPolicySummaryWithServiceLevel(f *testing.F) {
	f.Add(100.0, 5.0, 7.0, 20.0, 0.95)
	f.Fuzz(func(t *testing.T, dailyDemand, leadTimeDays, reviewPeriodDays, demandStdDevPerDay, serviceLevel float64) {
		summary, err := BuildPolicySummaryWithServiceLevel(PolicySummaryServiceLevelInput{
			DailyDemand:        dailyDemand,
			LeadTimeDays:       leadTimeDays,
			ReviewPeriodDays:   reviewPeriodDays,
			DemandStdDevPerDay: demandStdDevPerDay,
			ServiceLevel:       serviceLevel,
		})
		if err != nil {
			return
		}
		for _, v := range []float64{
			summary.ExpectedDemandDuringLeadTime, summary.SafetyStockUnits,
			summary.ReorderPoint, summary.TargetInventoryLevel,
			summary.MinLevel, summary.MaxLevel,
		} {
			if math.IsNaN(v) {
				t.Fatalf("BuildPolicySummaryWithServiceLevel returned NaN field with nil error: %+v", summary)
			}
		}
	})
}
