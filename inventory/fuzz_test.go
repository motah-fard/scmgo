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

func FuzzSafetyStockWithVariableLeadTime(f *testing.F) {
	f.Add(100.0, 10.0, 5.0, 1.0, 0.95)
	f.Fuzz(func(t *testing.T, avgDailyDemand, stdDevDailyDemand, avgLeadTimeDays, stdDevLeadTimeDays, serviceLevel float64) {
		result, err := SafetyStockWithVariableLeadTime(SafetyStockWithVariableLeadTimeInput{
			AvgDailyDemand:     avgDailyDemand,
			StdDevDailyDemand:  stdDevDailyDemand,
			AvgLeadTimeDays:    avgLeadTimeDays,
			StdDevLeadTimeDays: stdDevLeadTimeDays,
			ServiceLevel:       serviceLevel,
		})
		if err == nil && math.IsNaN(result) {
			t.Fatalf("SafetyStockWithVariableLeadTime returned NaN with nil error")
		}
	})
}

func FuzzUnitNormalLoss(f *testing.F) {
	f.Add(1.0)
	f.Fuzz(func(t *testing.T, z float64) {
		result, err := UnitNormalLoss(z)
		if err == nil && math.IsNaN(result) {
			t.Fatalf("UnitNormalLoss returned NaN with nil error for z=%v", z)
		}
	})
}

// FuzzFillRateSafetyStockRoundTrip fuzzes the property that actually caught
// a real numerical bug during development: naive computation of 1-Φ(z)
// loses precision for large z (Φ(z) rounds to exactly 1 in float64), which
// silently produced a wrong-but-finite safety stock -- a "not NaN" check
// alone would not have caught it, since the result was a valid float, just
// the wrong one. Feeding FillRateSafetyStock's result back through
// ExpectedFillRate and checking it recovers the original target is a much
// stronger invariant, and is what actually exercises the fix
// (UnitNormalLoss using math.Erfc instead of 1-math.Erf).
func FuzzFillRateSafetyStockRoundTrip(f *testing.F) {
	f.Add(0.98, 50.0, 200.0)
	f.Add(0.999999999999, 100.0, 1.0)
	f.Fuzz(func(t *testing.T, targetFillRate, sigmaL, orderQty float64) {
		ss, err := FillRateSafetyStock(FillRateSafetyStockInput{
			TargetFillRate:             targetFillRate,
			StdDevDemandDuringLeadTime: sigmaL,
			OrderQuantity:              orderQty,
		})
		if err != nil {
			return
		}
		if math.IsNaN(ss) {
			t.Fatalf("FillRateSafetyStock returned NaN with nil error")
		}

		achieved, err := ExpectedFillRate(ExpectedFillRateInput{
			SafetyStockUnits:           ss,
			StdDevDemandDuringLeadTime: sigmaL,
			OrderQuantity:              orderQty,
		})
		if err != nil {
			t.Fatalf("ExpectedFillRate rejected FillRateSafetyStock's own output (ss=%v, sigmaL=%v, Q=%v): %v", ss, sigmaL, orderQty, err)
		}
		if math.Abs(achieved-targetFillRate) > 1e-6 {
			t.Fatalf("round trip: target=%v achieved=%v (ss=%v, sigmaL=%v, Q=%v)", targetFillRate, achieved, ss, sigmaL, orderQty)
		}
	})
}

func FuzzEPQ(f *testing.F) {
	f.Add(10000.0, 50.0, 2.0, 40000.0)
	f.Fuzz(func(t *testing.T, annualDemand, setupCost, holdingCostPerUnit, annualProductionRate float64) {
		result, err := EPQ(EPQInput{
			AnnualDemand:         annualDemand,
			SetupCost:            setupCost,
			HoldingCostPerUnit:   holdingCostPerUnit,
			AnnualProductionRate: annualProductionRate,
		})
		if err == nil && math.IsNaN(result) {
			t.Fatalf("EPQ returned NaN with nil error")
		}
	})
}

func FuzzNewsvendor(f *testing.F) {
	f.Add(500.0, 100.0, 18.0, 7.0)
	f.Fuzz(func(t *testing.T, meanDemand, stdDevDemand, underageCost, overageCost float64) {
		result, err := Newsvendor(NewsvendorInput{
			MeanDemand:          meanDemand,
			StdDevDemand:        stdDevDemand,
			UnderageCostPerUnit: underageCost,
			OverageCostPerUnit:  overageCost,
		})
		if err != nil {
			return
		}
		if math.IsNaN(result.OrderQuantity) || math.IsNaN(result.CriticalRatio) {
			t.Fatalf("Newsvendor returned NaN field with nil error: %+v", result)
		}
	})
}

func FuzzEOQWithQuantityDiscounts(f *testing.F) {
	f.Add(10000.0, 20.0, 0.20, 1.0, 5.0, 500.0, 4.5, 1000.0, 3.9)
	f.Fuzz(func(t *testing.T, annualDemand, orderingCost, holdingCostRate,
		min1, price1, min2, price2, min3, price3 float64,
	) {
		result, err := EOQWithQuantityDiscounts(EOQWithQuantityDiscountsInput{
			AnnualDemand:    annualDemand,
			OrderingCost:    orderingCost,
			HoldingCostRate: holdingCostRate,
			Tiers: []QuantityDiscountTier{
				{MinQuantity: min1, UnitPrice: price1},
				{MinQuantity: min2, UnitPrice: price2},
				{MinQuantity: min3, UnitPrice: price3},
			},
		})
		if err != nil {
			return
		}
		if math.IsNaN(result.OrderQuantity) || math.IsNaN(result.UnitPrice) || math.IsNaN(result.TotalAnnualCost) {
			t.Fatalf("EOQWithQuantityDiscounts returned NaN field with nil error: %+v", result)
		}
	})
}

func FuzzTurnover(f *testing.F) {
	f.Add(120000.0, 20000.0)
	f.Fuzz(func(t *testing.T, cogs, avgInventoryValue float64) {
		result, err := Turnover(TurnoverInput{COGS: cogs, AverageInventoryValue: avgInventoryValue})
		if err == nil && math.IsNaN(result) {
			t.Fatalf("Turnover returned NaN with nil error")
		}
	})
}

func FuzzDaysOfInventoryOnHand(f *testing.F) {
	f.Add(20000.0, 120000.0, 365.0)
	f.Fuzz(func(t *testing.T, avgInventoryValue, cogs, daysInPeriod float64) {
		result, err := DaysOfInventoryOnHand(DaysOfInventoryOnHandInput{
			AverageInventoryValue: avgInventoryValue,
			COGS:                  cogs,
			DaysInPeriod:          daysInPeriod,
		})
		if err == nil && math.IsNaN(result) {
			t.Fatalf("DaysOfInventoryOnHand returned NaN with nil error")
		}
	})
}

func FuzzGMROI(f *testing.F) {
	f.Add(50000.0, 20000.0)
	f.Fuzz(func(t *testing.T, grossMargin, avgInventoryCost float64) {
		result, err := GMROI(GMROIInput{GrossMargin: grossMargin, AverageInventoryCost: avgInventoryCost})
		if err == nil && math.IsNaN(result) {
			t.Fatalf("GMROI returned NaN with nil error")
		}
	})
}

func FuzzEOI(f *testing.F) {
	f.Add(10000.0, 50.0, 2.0, 365.0)
	f.Fuzz(func(t *testing.T, annualDemand, orderingCost, holdingCostPerUnit, daysPerYear float64) {
		result, err := EOI(EOIInput{
			AnnualDemand:       annualDemand,
			OrderingCost:       orderingCost,
			HoldingCostPerUnit: holdingCostPerUnit,
			DaysPerYear:        daysPerYear,
		})
		if err == nil && math.IsNaN(result) {
			t.Fatalf("EOI returned NaN with nil error")
		}
	})
}

func FuzzSafetyTime(f *testing.F) {
	f.Add(50.0, 20.0)
	f.Fuzz(func(t *testing.T, safetyStockUnits, avgDailyDemand float64) {
		result, err := SafetyTime(SafetyTimeInput{SafetyStockUnits: safetyStockUnits, AvgDailyDemand: avgDailyDemand})
		if err == nil && math.IsNaN(result) {
			t.Fatalf("SafetyTime returned NaN with nil error")
		}
	})
}
