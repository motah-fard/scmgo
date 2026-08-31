package abc

// ClassifyVariability performs XYZ classification: it computes each item's
// demand coefficient of variation (StdDevDemand / MeanDemand) and assigns
// a class:
//
//	X: coefficient of variation <= XThreshold (low, most predictable)
//	Y: XThreshold < coefficient of variation <= YThreshold (medium)
//	Z: coefficient of variation > YThreshold (high, least predictable)
//
// The result is returned in the same order as Items.
//
// Items must contain at least one entry. MeanDemand must be strictly
// positive for every item (the coefficient of variation is undefined at
// mean zero). StdDevDemand must be non-negative. XThreshold must be
// greater than zero; YThreshold must be greater than XThreshold.
func ClassifyVariability(in ClassifyVariabilityInput) ([]VariabilityClassifiedItem, error) {
	if len(in.Items) == 0 {
		return nil, ErrEmptyItems
	}
	if err := validateFinite(in.XThreshold, in.YThreshold); err != nil {
		return nil, err
	}
	if in.XThreshold <= 0 {
		return nil, ErrInvalidVariabilityThreshold
	}
	if in.YThreshold <= in.XThreshold {
		return nil, ErrInvalidVariabilityThreshold
	}

	result := make([]VariabilityClassifiedItem, len(in.Items))
	for i, item := range in.Items {
		if err := validateFinite(item.MeanDemand, item.StdDevDemand); err != nil {
			return nil, err
		}
		if item.MeanDemand <= 0 {
			return nil, ErrInvalidMeanDemand
		}
		if item.StdDevDemand < 0 {
			return nil, ErrNegativeStandardDeviation
		}

		cv := item.StdDevDemand / item.MeanDemand

		class := "Z"
		switch {
		case cv <= in.XThreshold:
			class = "X"
		case cv <= in.YThreshold:
			class = "Y"
		}

		result[i] = VariabilityClassifiedItem{
			ID:                     item.ID,
			CoefficientOfVariation: cv,
			Class:                  class,
		}
	}

	return result, nil
}
