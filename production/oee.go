package production

// OEE calculates overall equipment effectiveness from raw inputs:
//
//	Availability = RunTime / PlannedProductionTime
//	Performance  = (IdealCycleTime × TotalCount) / RunTime
//	Quality      = GoodCount / TotalCount
//	OEE          = Availability × Performance × Quality
//
// PlannedProductionTime must be greater than zero. RunTime must be greater
// than zero and cannot exceed PlannedProductionTime. IdealCycleTime must
// be non-negative. TotalCount must be greater than zero. GoodCount must be
// non-negative and cannot exceed TotalCount.
//
// The result is not clamped to [0, 1]: a value above 1 indicates the
// inputs are inconsistent (e.g. an overstated TotalCount), not a bug.
func OEE(in OEEInput) (OEEResult, error) {
	if err := validateFinite(in.PlannedProductionTime, in.RunTime, in.IdealCycleTime, in.TotalCount, in.GoodCount); err != nil {
		return OEEResult{}, err
	}
	if in.PlannedProductionTime <= 0 {
		return OEEResult{}, ErrInvalidPlannedProductionTime
	}
	if in.RunTime <= 0 || in.RunTime > in.PlannedProductionTime {
		return OEEResult{}, ErrInvalidRunTime
	}
	if in.IdealCycleTime < 0 {
		return OEEResult{}, ErrNegativeIdealCycleTime
	}
	if in.TotalCount <= 0 {
		return OEEResult{}, ErrInvalidTotalCount
	}
	if in.GoodCount < 0 || in.GoodCount > in.TotalCount {
		return OEEResult{}, ErrInvalidGoodCount
	}

	availability := in.RunTime / in.PlannedProductionTime
	performance := (in.IdealCycleTime * in.TotalCount) / in.RunTime
	quality := in.GoodCount / in.TotalCount

	return OEEResult{
		Availability: availability,
		Performance:  performance,
		Quality:      quality,
		OEE:          availability * performance * quality,
	}, nil
}
