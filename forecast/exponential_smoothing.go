package forecast

// SimpleExponentialSmoothing calculates a one-period-ahead forecast using
// simple (single) exponential smoothing:
//
//	F[1] = A[1]
//	F[t] = alpha*A[t-1] + (1-alpha)*F[t-1]   for t = 2..n
//	Forecast = alpha*A[n] + (1-alpha)*F[n]
//
// Alpha must be in (0, 1]. This method assumes no trend or seasonality; use
// HoltLinearTrend when the series has a trend.
func SimpleExponentialSmoothing(in SimpleExponentialSmoothingInput) (SimpleExponentialSmoothingResult, error) {
	if len(in.History) == 0 {
		return SimpleExponentialSmoothingResult{}, ErrEmptyHistory
	}
	if err := validateSmoothingConstant(in.Alpha); err != nil {
		return SimpleExponentialSmoothingResult{}, err
	}
	if err := validateNonNegative(in.History); err != nil {
		return SimpleExponentialSmoothingResult{}, err
	}

	fitted := make([]float64, len(in.History))
	fitted[0] = in.History[0]
	for t := 1; t < len(in.History); t++ {
		fitted[t] = in.Alpha*in.History[t-1] + (1-in.Alpha)*fitted[t-1]
	}

	last := len(in.History) - 1
	forecast := in.Alpha*in.History[last] + (1-in.Alpha)*fitted[last]

	return SimpleExponentialSmoothingResult{
		Forecast: forecast,
		Fitted:   fitted,
	}, nil
}
