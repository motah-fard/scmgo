package production

import (
	"math"
	"testing"
)

func FuzzOEE(f *testing.F) {
	f.Add(480.0, 420.0, 1.0, 400.0, 385.0)
	f.Fuzz(func(t *testing.T, plannedTime, runTime, idealCycleTime, totalCount, goodCount float64) {
		result, err := OEE(OEEInput{
			PlannedProductionTime: plannedTime,
			RunTime:               runTime,
			IdealCycleTime:        idealCycleTime,
			TotalCount:            totalCount,
			GoodCount:             goodCount,
		})
		if err == nil && math.IsNaN(result.OEE) {
			t.Fatalf("OEE returned NaN with nil error")
		}
	})
}

func FuzzWIPFromLittlesLaw(f *testing.F) {
	f.Add(10.0, 5.0)
	f.Fuzz(func(t *testing.T, throughput, cycleTime float64) {
		result, err := WIPFromLittlesLaw(WIPFromLittlesLawInput{Throughput: throughput, CycleTime: cycleTime})
		if err == nil && math.IsNaN(result) {
			t.Fatalf("WIPFromLittlesLaw returned NaN with nil error")
		}
	})
}

func FuzzTaktTime(f *testing.F) {
	f.Add(480.0, 120.0)
	f.Fuzz(func(t *testing.T, availableTime, demand float64) {
		result, err := TaktTime(TaktTimeInput{AvailableProductionTime: availableTime, CustomerDemand: demand})
		if err == nil && math.IsNaN(result) {
			t.Fatalf("TaktTime returned NaN with nil error")
		}
	})
}
