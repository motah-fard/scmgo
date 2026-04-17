package inventory

// DemandDuringLeadTime calculates expected demand during lead time using:
//
//	demand during lead time = average daily demand × lead time days
//
// All input values must be non-negative.
func DemandDuringLeadTime(in DemandDuringLeadTimeInput) (float64, error) {
	if in.AvgDailyDemand < 0 {
		return 0, ErrNegativeDemand
	}
	if in.LeadTimeDays < 0 {
		return 0, ErrNegativeLeadTime
	}

	result := in.AvgDailyDemand * in.LeadTimeDays
	return result, nil
}
