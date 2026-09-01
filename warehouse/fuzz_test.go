package warehouse

import (
	"math"
	"testing"
)

func FuzzStorageUtilization(f *testing.F) {
	f.Add(800.0, 1000.0)
	f.Fuzz(func(t *testing.T, usedSpace, totalSpace float64) {
		result, err := StorageUtilization(StorageUtilizationInput{UsedSpace: usedSpace, TotalSpace: totalSpace})
		if err == nil && math.IsNaN(result) {
			t.Fatalf("StorageUtilization returned NaN with nil error")
		}
	})
}

func FuzzCubeUtilization(f *testing.F) {
	f.Add(4500.0, 6000.0)
	f.Fuzz(func(t *testing.T, usedVolume, totalVolume float64) {
		result, err := CubeUtilization(CubeUtilizationInput{UsedVolume: usedVolume, TotalVolume: totalVolume})
		if err == nil && math.IsNaN(result) {
			t.Fatalf("CubeUtilization returned NaN with nil error")
		}
	})
}

func FuzzPickRate(f *testing.F) {
	f.Add(480.0, 8.0)
	f.Fuzz(func(t *testing.T, totalPicks, totalHours float64) {
		result, err := PickRate(PickRateInput{TotalPicks: totalPicks, TotalHours: totalHours})
		if err == nil && math.IsNaN(result) {
			t.Fatalf("PickRate returned NaN with nil error")
		}
	})
}
