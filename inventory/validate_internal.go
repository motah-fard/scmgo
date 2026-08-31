package inventory

import "github.com/motah-fard/scmgo/internal/numeric"

// validateFinite returns ErrNonFiniteInput if any value is NaN or an
// infinity. See internal/numeric.AllFinite for why this check exists and
// runs before every other validation.
func validateFinite(values ...float64) error {
	if !numeric.AllFinite(values...) {
		return ErrNonFiniteInput
	}
	return nil
}
