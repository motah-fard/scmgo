package inventory

// SafetyTime calculates safety stock expressed as a time buffer instead of
// a quantity:
//
//	SafetyTime = SafetyStockUnits / AvgDailyDemand
//
// SafetyStockUnits must be non-negative. AvgDailyDemand must be greater
// than zero.
func SafetyTime(in SafetyTimeInput) (float64, error) {
	if err := validateFinite(in.SafetyStockUnits, in.AvgDailyDemand); err != nil {
		return 0, err
	}
	if in.SafetyStockUnits < 0 {
		return 0, ErrNegativeSafetyStock
	}
	if in.AvgDailyDemand <= 0 {
		return 0, ErrInvalidAvgDailyDemand
	}

	return in.SafetyStockUnits / in.AvgDailyDemand, nil
}
