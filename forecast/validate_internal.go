package forecast

import "github.com/motah-fard/scmgo/internal/numeric"

// validateFinite returns ErrNonFiniteInput if any value is NaN or an
// infinity. See internal/numeric.AllFinite for why this check exists and
// runs before every other validation. Accepts either individual values
// (validateFinite(a, b, c)) or a slice spread (validateFinite(values...)).
func validateFinite(values ...float64) error {
	if !numeric.AllFinite(values...) {
		return ErrNonFiniteInput
	}
	return nil
}

// validateNonNegative returns ErrNonFiniteInput or ErrNegativeDemand for the
// first offending value.
func validateNonNegative(values []float64) error {
	if err := validateFinite(values...); err != nil {
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
// ErrInvalidSmoothingConstant unless every constant is in (0, 1].
func validateSmoothingConstant(cs ...float64) error {
	if err := validateFinite(cs...); err != nil {
		return err
	}
	for _, c := range cs {
		if c <= 0 || c > 1 {
			return ErrInvalidSmoothingConstant
		}
	}
	return nil
}
