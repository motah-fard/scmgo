package quality

import (
	"math"
	"testing"
)

func FuzzDPMO(f *testing.F) {
	f.Add(45.0, 1000.0, 5.0)
	f.Fuzz(func(t *testing.T, defects, units, opportunities float64) {
		result, err := DPMO(DPMOInput{NumberOfDefects: defects, NumberOfUnits: units, OpportunitiesPerUnit: opportunities})
		if err == nil && math.IsNaN(result) {
			t.Fatalf("DPMO returned NaN with nil error")
		}
	})
}

func FuzzCp(f *testing.F) {
	f.Add(10.5, 9.5, 0.1)
	f.Fuzz(func(t *testing.T, usl, lsl, sigma float64) {
		result, err := Cp(CpInput{USL: usl, LSL: lsl, Sigma: sigma})
		if err == nil && math.IsNaN(result) {
			t.Fatalf("Cp returned NaN with nil error")
		}
	})
}

func FuzzCpk(f *testing.F) {
	f.Add(10.5, 9.5, 10.1, 0.1)
	f.Fuzz(func(t *testing.T, usl, lsl, mean, sigma float64) {
		result, err := Cpk(CpkInput{USL: usl, LSL: lsl, Mean: mean, Sigma: sigma})
		if err == nil && math.IsNaN(result) {
			t.Fatalf("Cpk returned NaN with nil error")
		}
	})
}

func FuzzCostOfQuality(f *testing.F) {
	f.Add(5000.0, 8000.0, 12000.0, 20000.0)
	f.Fuzz(func(t *testing.T, prevention, appraisal, internalFailure, externalFailure float64) {
		result, err := CostOfQuality(CostOfQualityInput{
			PreventionCost:      prevention,
			AppraisalCost:       appraisal,
			InternalFailureCost: internalFailure,
			ExternalFailureCost: externalFailure,
		})
		if err == nil && math.IsNaN(result.Total) {
			t.Fatalf("CostOfQuality returned NaN with nil error")
		}
	})
}
