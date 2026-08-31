package inventory

import "math"

// validateFinite returns ErrNonFiniteInput if v is NaN or an infinity.
//
// Comparisons against NaN are always false in IEEE 754, so a plain "v < 0"
// check silently lets NaN (and, for one-sided checks, +Inf) through. Every
// exported function checks its raw inputs with validateFinite before any
// other validation, so bad upstream data (e.g. a division by zero that
// produced NaN) fails loudly instead of propagating as a silently invalid
// result.
func validateFinite(v float64) error {
	if math.IsNaN(v) || math.IsInf(v, 0) {
		return ErrNonFiniteInput
	}
	return nil
}
