package inventory

// ReorderPointWithVariableLeadTime calculates the reorder point using:
//
//	reorder point = average daily demand × average lead time days + safety stock
//
// where safety stock accounts for variability in both demand and lead
// time, via SafetyStockWithVariableLeadTime.
//
// All demand and lead-time inputs must be non-negative. Service level must
// be strictly between 0 and 1.
func ReorderPointWithVariableLeadTime(in ReorderPointWithVariableLeadTimeInput) (float64, error) {
	if err := validateFinite(in.AvgDailyDemand, in.StdDevDailyDemand, in.AvgLeadTimeDays, in.StdDevLeadTimeDays); err != nil {
		return 0, err
	}
	if in.AvgDailyDemand < 0 {
		return 0, ErrNegativeDemand
	}
	if in.AvgLeadTimeDays < 0 {
		return 0, ErrNegativeLeadTime
	}

	// ReorderPointWithVariableLeadTimeInput and
	// SafetyStockWithVariableLeadTimeInput carry identical fields; convert
	// directly rather than re-listing them.
	safetyStock, err := SafetyStockWithVariableLeadTime(SafetyStockWithVariableLeadTimeInput(in))
	if err != nil {
		return 0, err
	}

	return (in.AvgDailyDemand * in.AvgLeadTimeDays) + safetyStock, nil
}
