package logistics

// AllocateFreightCost allocates a shared freight cost across items in
// proportion to each item's weight:
//
//	AllocatedCost_i = TotalFreightCost × (Weight_i / sum(Weight))
//
// Items must contain at least one entry. TotalFreightCost must be
// non-negative. Each item's Weight must be non-negative, and the weights
// must sum to more than zero.
func AllocateFreightCost(in AllocateFreightCostInput) ([]FreightAllocationResult, error) {
	if len(in.Items) == 0 {
		return nil, ErrEmptyItems
	}
	if err := validateFinite(in.TotalFreightCost); err != nil {
		return nil, err
	}
	if in.TotalFreightCost < 0 {
		return nil, ErrNegativeCost
	}

	var totalWeight float64
	for _, item := range in.Items {
		if err := validateFinite(item.Weight); err != nil {
			return nil, err
		}
		if item.Weight < 0 {
			return nil, ErrNegativeWeight
		}
		totalWeight += item.Weight
	}
	if totalWeight <= 0 {
		return nil, ErrZeroTotalWeight
	}

	result := make([]FreightAllocationResult, len(in.Items))
	for i, item := range in.Items {
		result[i] = FreightAllocationResult{
			ID:            item.ID,
			Weight:        item.Weight,
			AllocatedCost: in.TotalFreightCost * (item.Weight / totalWeight),
		}
	}

	return result, nil
}
