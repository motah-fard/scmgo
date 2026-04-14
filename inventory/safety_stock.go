package inventory

// SafetyStockBasic calculates safety stock using the basic formula:
//
//	safety stock = (max daily demand × max lead time days) -
//	               (average daily demand × average lead time days)
//
// All input values must be non-negative.
func SafetyStockBasic(in SafetyStockInput) (float64, error) {
	if in.MaxDailyDemand < 0 {
		return 0, ErrNegativeDemand
	}
	if in.MaxLeadTimeDays < 0 {
		return 0, ErrNegativeLeadTime
	}
	if in.AvgDailyDemand < 0 {
		return 0, ErrNegativeDemand
	}
	if in.AvgLeadTimeDays < 0 {
		return 0, ErrNegativeLeadTime
	}

	result := (in.MaxDailyDemand * in.MaxLeadTimeDays) -
		(in.AvgDailyDemand * in.AvgLeadTimeDays)

	// Safety stock should not be negative in practice for this basic model.
	if result < 0 {
		return 0, nil
	}

	return result, nil
}
