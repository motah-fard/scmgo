package inventory

// GMROI calculates gross margin return on inventory investment using:
//
//	GMROI = gross margin / average inventory cost
//
// GrossMargin may be negative (a loss). AverageInventoryCost must be
// greater than zero.
func GMROI(in GMROIInput) (float64, error) {
	if err := validateFinite(in.GrossMargin, in.AverageInventoryCost); err != nil {
		return 0, err
	}
	if in.AverageInventoryCost <= 0 {
		return 0, ErrInvalidAverageInventoryCost
	}

	return in.GrossMargin / in.AverageInventoryCost, nil
}
