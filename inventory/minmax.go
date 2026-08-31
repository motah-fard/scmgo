package inventory

// MinMaxLevels calculates the minimum and maximum inventory levels using:
//
//	min level = reorder point
//	max level = reorder point + order quantity
//
// Both input values must be non-negative.
func MinMaxLevels(in MinMaxInput) (MinMaxResult, error) {
	if err := validateFinite(in.ReorderPoint, in.OrderQuantity); err != nil {
		return MinMaxResult{}, err
	}
	if in.ReorderPoint < 0 {
		return MinMaxResult{}, ErrNegativeReorderPoint
	}
	if in.OrderQuantity < 0 {
		return MinMaxResult{}, ErrNegativeOrderQuantity
	}

	result := MinMaxResult{
		Min: in.ReorderPoint,
		Max: in.ReorderPoint + in.OrderQuantity,
	}

	return result, nil
}
