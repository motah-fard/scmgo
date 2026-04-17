package inventory

// MinMaxLevelsWithServiceLevel calculates min/max inventory levels using
// a service-level-based reorder point and a fixed order quantity.
//
// Formula:
//
//	min level = reorder point with service level
//	max level = min level + order quantity
//
// All input values must be non-negative, and service level must be
// strictly between 0 and 1.
func MinMaxLevelsWithServiceLevel(in MinMaxLevelsWithServiceLevelInput) (MinMaxResult, error) {
	if in.OrderQuantity < 0 {
		return MinMaxResult{}, ErrNegativeOrderQuantity
	}

	reorderPoint, err := ReorderPointWithServiceLevel(ReorderPointWithServiceLevelInput{
		AvgDailyDemand:    in.AvgDailyDemand,
		LeadTimeDays:      in.LeadTimeDays,
		StdDevDailyDemand: in.StdDevDailyDemand,
		ServiceLevel:      in.ServiceLevel,
	})
	if err != nil {
		return MinMaxResult{}, err
	}

	return MinMaxLevels(MinMaxInput{
		ReorderPoint:  reorderPoint,
		OrderQuantity: in.OrderQuantity,
	})
}
