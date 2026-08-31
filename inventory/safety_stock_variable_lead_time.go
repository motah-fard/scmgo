package inventory

import "math"

// SafetyStockWithVariableLeadTime calculates safety stock when both demand
// and lead time vary, using:
//
//	safety stock = z × sqrt(L × σd² + d̄² × σL²)
//
// where z is the service factor for ServiceLevel, L is average lead time,
// σd is the standard deviation of daily demand, d̄ is average daily demand,
// and σL is the standard deviation of lead time.
//
// This reduces to SafetyStockWithServiceLevel when StdDevLeadTimeDays is 0.
// All quantity inputs must be non-negative. Service level must be strictly
// between 0 and 1.
func SafetyStockWithVariableLeadTime(in SafetyStockWithVariableLeadTimeInput) (float64, error) {
	if err := validateFinite(in.AvgDailyDemand, in.StdDevDailyDemand, in.AvgLeadTimeDays, in.StdDevLeadTimeDays); err != nil {
		return 0, err
	}
	if in.AvgDailyDemand < 0 {
		return 0, ErrNegativeDemand
	}
	if in.StdDevDailyDemand < 0 {
		return 0, ErrNegativeStandardDeviation
	}
	if in.AvgLeadTimeDays < 0 {
		return 0, ErrNegativeLeadTime
	}
	if in.StdDevLeadTimeDays < 0 {
		return 0, ErrNegativeStandardDeviation
	}

	z, err := ZScoreForServiceLevel(in.ServiceLevel)
	if err != nil {
		return 0, err
	}

	variance := in.AvgLeadTimeDays*in.StdDevDailyDemand*in.StdDevDailyDemand +
		in.AvgDailyDemand*in.AvgDailyDemand*in.StdDevLeadTimeDays*in.StdDevLeadTimeDays

	return z * math.Sqrt(variance), nil
}
