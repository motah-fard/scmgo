package inventory

import "math"

// EPQ calculates the economic production quantity — EOQ's counterpart for
// items produced in-house at a finite rate rather than purchased in a
// single delivery:
//
//	EPQ = sqrt((2 × annual demand × setup cost) / (holding cost per unit × (1 - annual demand / annual production rate)))
//
// Annual demand and setup cost must be non-negative. Holding cost per unit
// must be greater than zero. Annual production rate must be strictly
// greater than annual demand (production must be able to keep up with
// demand, and the (1 - D/P) factor is undefined or non-positive otherwise).
func EPQ(in EPQInput) (float64, error) {
	if err := validateFinite(in.AnnualDemand, in.SetupCost, in.HoldingCostPerUnit, in.AnnualProductionRate); err != nil {
		return 0, err
	}
	if in.AnnualDemand < 0 {
		return 0, ErrNegativeDemand
	}
	if in.SetupCost < 0 {
		return 0, ErrNegativeOrderingCost
	}
	if in.HoldingCostPerUnit <= 0 {
		return 0, ErrInvalidHoldingCost
	}
	if in.AnnualProductionRate <= in.AnnualDemand {
		return 0, ErrInvalidProductionRate
	}

	factor := 1 - in.AnnualDemand/in.AnnualProductionRate
	result := math.Sqrt((2 * in.AnnualDemand * in.SetupCost) / (in.HoldingCostPerUnit * factor))

	return result, nil
}
