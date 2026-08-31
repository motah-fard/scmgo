package forecast

// LinearTrend calculates a forecast by fitting an ordinary-least-squares
// trend line to History (treating period index 0, 1, 2, ... as X) and
// extrapolating it PeriodsAhead periods beyond the end of History:
//
//	slope     = sum((x_i - x̄)(y_i - ȳ)) / sum((x_i - x̄)²)
//	intercept = ȳ - slope × x̄
//	forecast  = intercept + slope × (len(History) - 1 + PeriodsAhead)
//
// This assumes a single linear trend with no seasonality; use HoltWinters
// for seasonal data.
//
// History must contain at least two periods. PeriodsAhead must be greater
// than zero.
func LinearTrend(in LinearTrendInput) (LinearTrendResult, error) {
	if len(in.History) < 2 {
		return LinearTrendResult{}, ErrInsufficientHistory
	}
	if in.PeriodsAhead <= 0 {
		return LinearTrendResult{}, ErrInvalidPeriods
	}
	if err := validateNonNegative(in.History); err != nil {
		return LinearTrendResult{}, err
	}

	n := len(in.History)
	xBar := float64(n-1) / 2

	var ySum float64
	for _, y := range in.History {
		ySum += y
	}
	yBar := ySum / float64(n)

	var numerator, denominator float64
	for i, y := range in.History {
		dx := float64(i) - xBar
		numerator += dx * (y - yBar)
		denominator += dx * dx
	}

	slope := numerator / denominator
	intercept := yBar - slope*xBar
	forecast := intercept + slope*float64(n-1+in.PeriodsAhead)

	return LinearTrendResult{
		Slope:     slope,
		Intercept: intercept,
		Forecast:  forecast,
	}, nil
}
