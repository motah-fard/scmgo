package logistics

import "math"

// DimensionalWeight calculates a parcel's dimensional (volumetric) weight:
//
//	DimensionalWeight = (Length × Width × Height) / DimFactor
//
// Length, Width, and Height must be greater than zero. DimFactor must be
// greater than zero.
func DimensionalWeight(in DimensionalWeightInput) (float64, error) {
	if err := validateFinite(in.Length, in.Width, in.Height, in.DimFactor); err != nil {
		return 0, err
	}
	if in.Length <= 0 || in.Width <= 0 || in.Height <= 0 {
		return 0, ErrInvalidDimension
	}
	if in.DimFactor <= 0 {
		return 0, ErrInvalidDimFactor
	}

	return (in.Length * in.Width * in.Height) / in.DimFactor, nil
}

// BillableWeight returns the greater of actual and dimensional weight, the
// convention most carriers use to determine what to charge for.
//
// Both inputs must be non-negative.
func BillableWeight(actualWeight, dimensionalWeight float64) (float64, error) {
	if err := validateFinite(actualWeight, dimensionalWeight); err != nil {
		return 0, err
	}
	if actualWeight < 0 || dimensionalWeight < 0 {
		return 0, ErrNegativeWeight
	}

	return math.Max(actualWeight, dimensionalWeight), nil
}
