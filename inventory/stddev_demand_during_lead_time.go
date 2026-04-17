package inventory

import "math"

// StdDevDemandDuringLeadTime calculates the standard deviation of demand
// during lead time using:
//
//	standard deviation during lead time = standard deviation of daily demand × sqrt(lead time days)
//
// All input values must be non-negative.
func StdDevDemandDuringLeadTime(in StdDevDemandDuringLeadTimeInput) (float64, error) {
	if in.StdDevDailyDemand < 0 {
		return 0, ErrNegativeStandardDeviation
	}
	if in.LeadTimeDays < 0 {
		return 0, ErrNegativeLeadTime
	}

	result := in.StdDevDailyDemand * math.Sqrt(in.LeadTimeDays)
	return result, nil
}
