package finance

import (
	"math"
	"testing"
)

func FuzzDSO(f *testing.F) {
	f.Add(150000.0, 1800000.0, 365.0)
	f.Fuzz(func(t *testing.T, ar, creditSales, days float64) {
		result, err := DSO(DSOInput{AccountsReceivable: ar, TotalCreditSales: creditSales, DaysInPeriod: days})
		if err == nil && math.IsNaN(result) {
			t.Fatalf("DSO returned NaN with nil error")
		}
	})
}

func FuzzDPO(f *testing.F) {
	f.Add(120000.0, 1200000.0, 365.0)
	f.Fuzz(func(t *testing.T, ap, cogs, days float64) {
		result, err := DPO(DPOInput{AccountsPayable: ap, COGS: cogs, DaysInPeriod: days})
		if err == nil && math.IsNaN(result) {
			t.Fatalf("DPO returned NaN with nil error")
		}
	})
}

func FuzzCashToCashCycleTime(f *testing.F) {
	f.Add(60.83, 30.42, 36.5)
	f.Fuzz(func(t *testing.T, dio, dso, dpo float64) {
		result, err := CashToCashCycleTime(CashToCashCycleTimeInput{DIO: dio, DSO: dso, DPO: dpo})
		if err == nil && math.IsNaN(result) {
			t.Fatalf("CashToCashCycleTime returned NaN with nil error")
		}
	})
}

func FuzzPerfectOrderRate(f *testing.F) {
	f.Add(0.95, 0.97, 0.99, 0.98)
	f.Fuzz(func(t *testing.T, onTime, complete, damageFree, accurateDocs float64) {
		result, err := PerfectOrderRate(PerfectOrderRateInput{
			OnTimeRate:                onTime,
			CompleteRate:              complete,
			DamageFreeRate:            damageFree,
			AccurateDocumentationRate: accurateDocs,
		})
		if err == nil && math.IsNaN(result) {
			t.Fatalf("PerfectOrderRate returned NaN with nil error")
		}
	})
}
