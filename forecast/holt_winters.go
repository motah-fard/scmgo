package forecast

// HoltWinters calculates a forecast using the additive-seasonality
// Holt-Winters method (triple exponential smoothing), extending Holt's
// linear trend with a seasonal component of period m = SeasonLength:
//
//	L[t] = alpha*(A[t] - S[t-m])    + (1-alpha)*(L[t-1] + T[t-1])
//	T[t] = beta*(L[t] - L[t-1])     + (1-beta)*T[t-1]
//	S[t] = gamma*(A[t] - L[t])      + (1-gamma)*S[t-m]
//	Forecast = L[n] + PeriodsAhead*T[n] + S[n+PeriodsAhead-m]
//
// Initialization (using the first two full seasons of History, 0-indexed
// positions 0..m-1 and m..2m-1):
//
//	avg1 = mean(History[0:m])
//	avg2 = mean(History[m:2m])
//	L[m-1] = avg1
//	T[m-1] = (avg2 - avg1) / m
//	S[i] = History[i] - avg1   for i = 0..m-1
//
// The recursion then runs over History[m:], updating one seasonal index
// per step. This is one standard initialization among several found in the
// literature; if you need to match another textbook or tool's specific
// convention, verify against it directly rather than assuming this one
// matches.
//
// Alpha, Beta, and Gamma must be in (0, 1]. SeasonLength must be at least
// 2. History must contain at least 2*SeasonLength periods.
func HoltWinters(in HoltWintersInput) (HoltWintersResult, error) {
	if in.SeasonLength < 2 {
		return HoltWintersResult{}, ErrInvalidSeasonLength
	}
	if len(in.History) < 2*in.SeasonLength {
		return HoltWintersResult{}, ErrInsufficientHistory
	}
	if in.PeriodsAhead <= 0 {
		return HoltWintersResult{}, ErrInvalidPeriods
	}
	if err := validateSmoothingConstant(in.Alpha, in.Beta, in.Gamma); err != nil {
		return HoltWintersResult{}, err
	}
	if err := validateNonNegative(in.History); err != nil {
		return HoltWintersResult{}, err
	}

	m := in.SeasonLength
	n := len(in.History)

	var sum1, sum2 float64
	for i := 0; i < m; i++ {
		sum1 += in.History[i]
		sum2 += in.History[m+i]
	}
	avg1 := sum1 / float64(m)
	avg2 := sum2 / float64(m)

	level := avg1
	trend := (avg2 - avg1) / float64(m)
	seasonal := make([]float64, m)
	for i := 0; i < m; i++ {
		seasonal[i] = in.History[i] - avg1
	}

	for t := m; t < n; t++ {
		prevLevel := level
		idx := t % m
		level = in.Alpha*(in.History[t]-seasonal[idx]) + (1-in.Alpha)*(level+trend)
		trend = in.Beta*(level-prevLevel) + (1-in.Beta)*trend
		seasonal[idx] = in.Gamma*(in.History[t]-level) + (1-in.Gamma)*seasonal[idx]
	}

	h := in.PeriodsAhead
	forecastIdx := (n - 1 + h) % m
	forecast := level + float64(h)*trend + seasonal[forecastIdx]

	return HoltWintersResult{
		Level:    level,
		Trend:    trend,
		Seasonal: seasonal,
		Forecast: forecast,
	}, nil
}
