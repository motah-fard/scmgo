package inventory

// ReorderPointWithServiceLevel calculates the reorder point using:
//
//	reorder point = average daily demand × lead time days + safety stock
//
// where safety stock is computed from demand variability, lead time,
// and the target cycle service level.
//
// All input values must be non-negative, and service level must be
// strictly between 0 and 1.
func ReorderPointWithServiceLevel(in ReorderPointWithServiceLevelInput) (float64, error) {
	if in.AvgDailyDemand < 0 {
		return 0, ErrNegativeDemand
	}
	if in.LeadTimeDays < 0 {
		return 0, ErrNegativeLeadTime
	}
	if in.StdDevDailyDemand < 0 {
		return 0, ErrNegativeStandardDeviation
	}

	safetyStock, err := SafetyStockWithServiceLevel(SafetyStockWithServiceLevelInput{
		StdDevDailyDemand: in.StdDevDailyDemand,
		LeadTimeDays:      in.LeadTimeDays,
		ServiceLevel:      in.ServiceLevel,
	})
	if err != nil {
		return 0, err
	}

	result := (in.AvgDailyDemand * in.LeadTimeDays) + safetyStock
	return result, nil
}
