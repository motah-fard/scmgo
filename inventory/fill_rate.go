package inventory

// ExpectedFillRate calculates the item fill rate (P2) — the fraction of
// demand met directly from stock — achieved by a given safety stock and
// order quantity, using:
//
//	fill rate = 1 - (σL × L(z)) / Q
//
// where z = SafetyStockUnits / σL, σL is StdDevDemandDuringLeadTime, L is
// the standard normal unit loss function (see UnitNormalLoss), and Q is
// OrderQuantity.
//
// SafetyStockUnits may be negative: for a low enough target fill rate, the
// quantity required to hit it (see FillRateSafetyStock) can be less than
// zero — the order quantity alone already exceeds the target without any
// additional buffer. StdDevDemandDuringLeadTime and OrderQuantity must be
// greater than zero.
func ExpectedFillRate(in ExpectedFillRateInput) (float64, error) {
	if err := validateFinite(in.SafetyStockUnits, in.StdDevDemandDuringLeadTime, in.OrderQuantity); err != nil {
		return 0, err
	}
	if in.StdDevDemandDuringLeadTime <= 0 {
		return 0, ErrInvalidStandardDeviation
	}
	if in.OrderQuantity <= 0 {
		return 0, ErrInvalidOrderQuantity
	}

	z := in.SafetyStockUnits / in.StdDevDemandDuringLeadTime
	loss, err := UnitNormalLoss(z)
	if err != nil {
		return 0, err
	}

	return 1 - (in.StdDevDemandDuringLeadTime*loss)/in.OrderQuantity, nil
}
