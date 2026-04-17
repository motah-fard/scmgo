package inventory

// TargetInventoryLevel calculates target inventory level using:
//
//	target inventory level = expected demand during lead time + safety stock
//
// All input values must be non-negative.
func TargetInventoryLevel(in TargetInventoryLevelInput) (float64, error) {
	if in.ExpectedDemandDuringLeadTime < 0 {
		return 0, ErrNegativeExpectedDemand
	}
	if in.SafetyStockUnits < 0 {
		return 0, ErrNegativeSafetyStock
	}

	result := in.ExpectedDemandDuringLeadTime + in.SafetyStockUnits
	return result, nil
}
