package forecast

import "math"

// validateFinite returns ErrNonFiniteInput if v is NaN or an infinity.
//
// Comparisons against NaN are always false in IEEE 754, so a plain "v < 0"
// or range check silently lets NaN (and, for one-sided checks, +Inf)
// through. Every validation helper below checks finiteness first, so bad
// upstream data fails loudly instead of propagating as a silently invalid
// result (e.g. a NaN forecast with a nil error).
func validateFinite(v float64) error {
	if math.IsNaN(v) || math.IsInf(v, 0) {
		return ErrNonFiniteInput
	}
	return nil
}

// validateFiniteSlice returns ErrNonFiniteInput if any value is NaN or an
// infinity.
func validateFiniteSlice(values []float64) error {
	for _, v := range values {
		if err := validateFinite(v); err != nil {
			return err
		}
	}
	return nil
}

// validateNonNegative returns ErrNonFiniteInput or ErrNegativeDemand for the
// first offending value.
func validateNonNegative(values []float64) error {
	if err := validateFiniteSlice(values); err != nil {
		return err
	}
	for _, v := range values {
		if v < 0 {
			return ErrNegativeDemand
		}
	}
	return nil
}

// validateSmoothingConstant returns ErrNonFiniteInput or
// ErrInvalidSmoothingConstant unless the constant is in (0, 1].
func validateSmoothingConstant(c float64) error {
	if err := validateFinite(c); err != nil {
		return err
	}
	if c <= 0 || c > 1 {
		return ErrInvalidSmoothingConstant
	}
	return nil
}
