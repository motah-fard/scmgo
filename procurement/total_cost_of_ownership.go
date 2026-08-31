package procurement

// TotalCostOfOwnership calculates total cost of ownership by summing
// labeled cost components over an asset or item's lifecycle (e.g.
// acquisition, operating, maintenance, end-of-life — whatever categories
// apply to the caller's situation).
//
// Components must contain at least one entry. Each component's Amount
// must be finite; it may be negative (e.g. expected resale value).
func TotalCostOfOwnership(in TotalCostOfOwnershipInput) (TotalCostOfOwnershipResult, error) {
	total, err := sumCostComponents(in.Components)
	if err != nil {
		return TotalCostOfOwnershipResult{}, err
	}

	return TotalCostOfOwnershipResult{
		Total:      total,
		Components: in.Components,
	}, nil
}
