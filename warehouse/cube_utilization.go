package warehouse

// CubeUtilization calculates how much of a warehouse's storage volume is
// used:
//
//	Utilization = UsedVolume / TotalVolume
//
// UsedVolume must be non-negative. TotalVolume must be greater than zero.
//
// The result is not clamped: a value above 1 indicates overcommitted
// storage, which is meaningful information, not an error.
func CubeUtilization(in CubeUtilizationInput) (float64, error) {
	if err := validateFinite(in.UsedVolume, in.TotalVolume); err != nil {
		return 0, err
	}
	if in.UsedVolume < 0 {
		return 0, ErrNegativeUsedVolume
	}
	if in.TotalVolume <= 0 {
		return 0, ErrInvalidTotalVolume
	}

	return in.UsedVolume / in.TotalVolume, nil
}
