package inventory

import "math"

// EOQ calculates the economic order quantity using the classic Wilson EOQ formula:
//
//	EOQ = sqrt((2 × annual demand × ordering cost) / holding cost per unit)
//
// Annual demand and ordering cost must be non-negative.
// Holding cost per unit must be greater than zero.
func EOQ(in EOQInput) (float64, error) {
	if in.AnnualDemand < 0 {
		return 0, ErrNegativeDemand
	}
	if in.OrderingCost < 0 {
		return 0, ErrNegativeOrderingCost
	}
	if in.HoldingCostPerUnit <= 0 {
		return 0, ErrInvalidHoldingCost
	}

	result := math.Sqrt((2 * in.AnnualDemand * in.OrderingCost) / in.HoldingCostPerUnit)
	return result, nil
}
