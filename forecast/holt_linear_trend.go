package forecast

// HoltLinearTrend calculates a forecast using Holt's linear trend method
// (double exponential smoothing), which extends simple exponential
// smoothing with a trend component:
//
//	L[1] = A[1]
//	T[1] = A[2] - A[1]
//	L[t] = alpha*A[t]       + (1-alpha)*(L[t-1] + T[t-1])   for t = 2..n
//	T[t] = beta*(L[t]-L[t-1]) + (1-beta)*T[t-1]              for t = 2..n
//	Forecast = L[n] + PeriodsAhead*T[n]
//
// Alpha and Beta must be in (0, 1]. History must contain at least two
// periods.
func HoltLinearTrend(in HoltLinearTrendInput) (HoltLinearTrendResult, error) {
	if len(in.History) < 2 {
		return HoltLinearTrendResult{}, ErrInsufficientHistory
	}
	if in.PeriodsAhead <= 0 {
		return HoltLinearTrendResult{}, ErrInvalidPeriods
	}
	if err := validateSmoothingConstant(in.Alpha); err != nil {
		return HoltLinearTrendResult{}, err
	}
	if err := validateSmoothingConstant(in.Beta); err != nil {
		return HoltLinearTrendResult{}, err
	}
	if err := validateNonNegative(in.History); err != nil {
		return HoltLinearTrendResult{}, err
	}

	level := in.History[0]
	trend := in.History[1] - in.History[0]

	for t := 1; t < len(in.History); t++ {
		prevLevel := level
		level = in.Alpha*in.History[t] + (1-in.Alpha)*(level+trend)
		trend = in.Beta*(level-prevLevel) + (1-in.Beta)*trend
	}

	forecast := level + float64(in.PeriodsAhead)*trend

	return HoltLinearTrendResult{
		Level:    level,
		Trend:    trend,
		Forecast: forecast,
	}, nil
}
