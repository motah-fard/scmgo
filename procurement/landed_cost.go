package procurement

// LandedCost calculates the total cost of getting a purchased item to its
// destination by summing labeled cost components (e.g. unit cost, freight,
// duty, insurance, handling — whatever categories apply to the caller's
// situation).
//
// Components must contain at least one entry. Each component's Amount
// must be finite; it may be negative (e.g. a rebate or credit).
func LandedCost(in LandedCostInput) (LandedCostResult, error) {
	total, err := sumCostComponents(in.Components)
	if err != nil {
		return LandedCostResult{}, err
	}

	return LandedCostResult{
		Total:      total,
		Components: in.Components,
	}, nil
}
