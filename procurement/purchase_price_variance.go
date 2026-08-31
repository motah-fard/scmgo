package procurement

// PurchasePriceVariance calculates the cost impact of paying a different
// unit price than standard/budgeted:
//
//	PPV = (StandardCost - ActualCost) × Quantity
//
// Following the standard-costing convention, a positive result is a
// favorable variance (paid less than standard); a negative result is
// unfavorable (paid more).
//
// StandardCost and ActualCost must be non-negative. Quantity must be
// non-negative.
func PurchasePriceVariance(in PurchasePriceVarianceInput) (float64, error) {
	if err := validateFinite(in.StandardCost, in.ActualCost, in.Quantity); err != nil {
		return 0, err
	}
	if in.StandardCost < 0 || in.ActualCost < 0 {
		return 0, ErrNegativeCost
	}
	if in.Quantity < 0 {
		return 0, ErrNegativeQuantity
	}

	return (in.StandardCost - in.ActualCost) * in.Quantity, nil
}
