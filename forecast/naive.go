package forecast

// Naive forecasts the next period as simply the most recent observed
// value. It exists mainly as a baseline to benchmark other methods
// against (e.g. via MASE, which scales error against exactly this naive
// one-step benchmark).
//
// History must contain at least one period.
func Naive(in NaiveInput) (float64, error) {
	if len(in.History) == 0 {
		return 0, ErrEmptyHistory
	}
	if err := validateNonNegative(in.History); err != nil {
		return 0, err
	}

	return in.History[len(in.History)-1], nil
}
