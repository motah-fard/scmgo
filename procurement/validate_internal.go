package procurement

import "github.com/motah-fard/scmgo/internal/numeric"

// validateFinite returns ErrNonFiniteInput if any value is NaN or an
// infinity.
func validateFinite(values ...float64) error {
	if !numeric.AllFinite(values...) {
		return ErrNonFiniteInput
	}
	return nil
}

// sumCostComponents validates and sums a labeled cost component list,
// shared by LandedCost and TotalCostOfOwnership: same underlying
// arithmetic, distinct public names for two distinct, recognized business
// concepts.
func sumCostComponents(components []CostComponent) (float64, error) {
	if len(components) == 0 {
		return 0, ErrEmptyComponents
	}

	var total float64
	for _, c := range components {
		if err := validateFinite(c.Amount); err != nil {
			return 0, err
		}
		total += c.Amount
	}

	return total, nil
}
