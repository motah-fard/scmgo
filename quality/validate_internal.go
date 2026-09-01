package quality

import "github.com/motah-fard/scmgo/internal/numeric"

// validateFinite returns ErrNonFiniteInput if any value is NaN or an
// infinity.
func validateFinite(values ...float64) error {
	if !numeric.AllFinite(values...) {
		return ErrNonFiniteInput
	}
	return nil
}
