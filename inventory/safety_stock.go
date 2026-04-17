package inventory

// SafetyStockWithServiceLevel calculates safety stock using
// demand variability, lead time, and a target cycle service level.
//
// Formula:
//
//	safety stock = z × standard deviation of demand during lead time
//
// where:
//
//	standard deviation of demand during lead time =
//	standard deviation of daily demand × sqrt(lead time days)
//
// All input values must be non-negative, and service level must be
// strictly between 0 and 1.
func SafetyStockWithServiceLevel(in SafetyStockWithServiceLevelInput) (float64, error) {
	z, err := ZScoreForServiceLevel(in.ServiceLevel)
	if err != nil {
		return 0, err
	}

	stdDevLeadTime, err := StdDevDemandDuringLeadTime(StdDevDemandDuringLeadTimeInput{
		StdDevDailyDemand: in.StdDevDailyDemand,
		LeadTimeDays:      in.LeadTimeDays,
	})
	if err != nil {
		return 0, err
	}

	result := z * stdDevLeadTime
	return result, nil
}
