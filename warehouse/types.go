package warehouse

// StorageUtilizationInput contains the inputs required to calculate
// storage utilization.
type StorageUtilizationInput struct {
	UsedSpace  float64
	TotalSpace float64
}

// CubeUtilizationInput contains the inputs required to calculate cube
// (volumetric) utilization.
type CubeUtilizationInput struct {
	UsedVolume  float64
	TotalVolume float64
}

// PickRateInput contains the inputs required to calculate pick rate.
type PickRateInput struct {
	TotalPicks float64
	TotalHours float64
}
