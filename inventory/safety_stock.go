package inventory

import "math"

// SafetyStockWithServiceLevel calculates safety stock using
// demand variability, lead time, and a target cycle service level.
//
// Formula:
//
//	safety stock = z × standard deviation of daily demand × sqrt(lead time days)
//
// All input values must be non-negative, and service level must be
// strictly between 0 and 1.
func SafetyStockWithServiceLevel(in SafetyStockWithServiceLevelInput) (float64, error) {
	if in.StdDevDailyDemand < 0 {
		return 0, ErrNegativeStandardDeviation
	}
	if in.LeadTimeDays < 0 {
		return 0, ErrNegativeLeadTime
	}

	z, err := ZScoreForServiceLevel(in.ServiceLevel)
	if err != nil {
		return 0, err
	}

	result := z * in.StdDevDailyDemand * math.Sqrt(in.LeadTimeDays)
	return result, nil
}
