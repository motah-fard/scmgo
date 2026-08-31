package inventory

// DemandDuringLeadTime calculates expected demand coverage over lead time using:
//
//	demand during lead time = average daily demand × lead time days
//
// All input values must be non-negative.
func DemandDuringLeadTime(in DemandDuringLeadTimeInput) (float64, error) {
	if err := validateFinite(in.AvgDailyDemand); err != nil {
		return 0, err
	}
	if err := validateFinite(in.LeadTimeDays); err != nil {
		return 0, err
	}
	if in.AvgDailyDemand < 0 {
		return 0, ErrNegativeDemand
	}
	if in.LeadTimeDays < 0 {
		return 0, ErrNegativeLeadTime
	}

	result := in.AvgDailyDemand * in.LeadTimeDays
	return result, nil
}
