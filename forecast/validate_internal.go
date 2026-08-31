package forecast

// validateNonNegative returns ErrNegativeDemand if any value is negative.
func validateNonNegative(values []float64) error {
	for _, v := range values {
		if v < 0 {
			return ErrNegativeDemand
		}
	}
	return nil
}

// validateSmoothingConstant returns ErrInvalidSmoothingConstant unless the
// constant is in (0, 1].
func validateSmoothingConstant(c float64) error {
	if c <= 0 || c > 1 {
		return ErrInvalidSmoothingConstant
	}
	return nil
}
