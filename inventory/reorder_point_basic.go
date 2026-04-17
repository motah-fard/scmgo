package inventory

// ReorderPoint calculates the reorder point using:
//
//	reorder point = average daily demand × lead time days + safety stock
//
// All input values must be non-negative.
func ReorderPoint(in ReorderPointInput) (float64, error) {
	if in.AvgDailyDemand < 0 {
		return 0, ErrNegativeDemand
	}
	if in.LeadTimeDays < 0 {
		return 0, ErrNegativeLeadTime
	}
	if in.SafetyStockUnits < 0 {
		return 0, ErrNegativeSafetyStock
	}

	result := (in.AvgDailyDemand * in.LeadTimeDays) + in.SafetyStockUnits
	return result, nil
}
