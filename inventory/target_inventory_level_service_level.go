package inventory

// TargetInventoryLevelWithServiceLevel calculates target inventory level using
// expected demand coverage and safety stock based on a target cycle service level.
//
// Formula:
//
//	target inventory level = expected demand coverage + safety stock
//
// In this helper, expected demand coverage is calculated over lead time.
// All input values must be non-negative, and service level must be
// strictly between 0 and 1.
func TargetInventoryLevelWithServiceLevel(in TargetInventoryLevelWithServiceLevelInput) (float64, error) {
	expectedDemand, err := DemandDuringLeadTime(DemandDuringLeadTimeInput{
		AvgDailyDemand: in.AvgDailyDemand,
		LeadTimeDays:   in.LeadTimeDays,
	})
	if err != nil {
		return 0, err
	}

	safetyStock, err := SafetyStockWithServiceLevel(SafetyStockWithServiceLevelInput{
		StdDevDailyDemand: in.StdDevDailyDemand,
		LeadTimeDays:      in.LeadTimeDays,
		ServiceLevel:      in.ServiceLevel,
	})
	if err != nil {
		return 0, err
	}

	result, err := TargetInventoryLevel(TargetInventoryLevelInput{
		ExpectedDemandDuringLeadTime: expectedDemand,
		SafetyStockUnits:             safetyStock,
	})
	if err != nil {
		return 0, err
	}

	return result, nil
}
