package warehouse

// StorageUtilization calculates how much of a warehouse's storage
// capacity is used:
//
//	Utilization = UsedSpace / TotalSpace
//
// UsedSpace must be non-negative. TotalSpace must be greater than zero.
//
// The result is not clamped: a value above 1 indicates overcommitted
// storage, which is meaningful information, not an error.
func StorageUtilization(in StorageUtilizationInput) (float64, error) {
	if err := validateFinite(in.UsedSpace, in.TotalSpace); err != nil {
		return 0, err
	}
	if in.UsedSpace < 0 {
		return 0, ErrNegativeUsedSpace
	}
	if in.TotalSpace <= 0 {
		return 0, ErrInvalidTotalSpace
	}

	return in.UsedSpace / in.TotalSpace, nil
}
