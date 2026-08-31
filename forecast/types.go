package forecast

// MovingAverageInput contains the inputs required to calculate a simple
// moving average forecast.
type MovingAverageInput struct {
	// History is the chronological demand series, oldest value first.
	History []float64
	// Periods is the number of most recent periods to average.
	Periods int
}

// WeightedMovingAverageInput contains the inputs required to calculate a
// weighted moving average forecast.
type WeightedMovingAverageInput struct {
	// History is the chronological demand series, oldest value first.
	History []float64
	// Weights are applied to the most recent len(Weights) periods of
	// History, in chronological order (Weights[0] applies to the oldest of
	// the selected periods, Weights[len(Weights)-1] to the most recent).
	// Weights must be non-negative and sum to 1.
	Weights []float64
}

// SimpleExponentialSmoothingInput contains the inputs required to calculate
// a simple exponential smoothing forecast.
type SimpleExponentialSmoothingInput struct {
	// History is the chronological demand series, oldest value first.
	History []float64
	// Alpha is the level smoothing constant, in (0, 1].
	Alpha float64
}

// SimpleExponentialSmoothingResult contains the outputs of a simple
// exponential smoothing calculation.
type SimpleExponentialSmoothingResult struct {
	// Forecast is the one-period-ahead forecast.
	Forecast float64
	// Fitted contains the in-sample fitted (one-step-ahead) forecast for
	// each period in History, useful for computing forecast accuracy with
	// Accuracy.
	Fitted []float64
}

// HoltLinearTrendInput contains the inputs required to calculate a Holt's
// linear trend (double exponential smoothing) forecast.
type HoltLinearTrendInput struct {
	// History is the chronological demand series, oldest value first.
	// At least two periods are required to estimate an initial trend.
	History []float64
	// Alpha is the level smoothing constant, in (0, 1].
	Alpha float64
	// Beta is the trend smoothing constant, in (0, 1].
	Beta float64
	// PeriodsAhead is the forecast horizon, in periods, and must be greater
	// than zero.
	PeriodsAhead int
}

// HoltLinearTrendResult contains the outputs of a Holt's linear trend
// calculation.
type HoltLinearTrendResult struct {
	// Level is the smoothed level estimate at the end of History.
	Level float64
	// Trend is the smoothed trend estimate at the end of History.
	Trend float64
	// Forecast is the forecast for PeriodsAhead periods beyond the end of
	// History.
	Forecast float64
}

// CrostonInput contains the inputs required to forecast intermittent demand
// using Croston's method.
type CrostonInput struct {
	// History is the chronological demand series, oldest value first, where
	// zero denotes a period with no demand. At least one non-zero period is
	// required.
	History []float64
	// Alpha is the smoothing constant applied to both demand size and
	// inter-demand interval, in (0, 1].
	Alpha float64
}

// CrostonResult contains the outputs of a Croston's method calculation.
type CrostonResult struct {
	// DemandSize is the smoothed estimate of demand size when demand occurs.
	DemandSize float64
	// Interval is the smoothed estimate of the number of periods between
	// non-zero demand occurrences.
	Interval float64
	// Forecast is the estimated demand per period (DemandSize / Interval).
	Forecast float64
}

// AccuracyInput contains the inputs required to calculate forecast
// accuracy metrics.
type AccuracyInput struct {
	// Actual is the chronological series of actual observed values.
	Actual []float64
	// Forecast is the chronological series of forecast values, aligned
	// period-for-period with Actual. Must be the same length as Actual.
	Forecast []float64
}

// AccuracyResult contains forecast accuracy metrics computed from
// an actual/forecast series pair.
type AccuracyResult struct {
	// MAD is the mean absolute deviation: mean(|actual - forecast|).
	MAD float64
	// MAPE is the mean absolute percentage error, expressed as a fraction
	// (0.05 = 5%). Periods where Actual is zero are excluded. NaN if every
	// actual value is zero.
	MAPE float64
	// Bias is the mean signed error: mean(actual - forecast). Positive
	// means the forecast under-forecasted on average.
	Bias float64
	// RMSE is the root mean squared error.
	RMSE float64
}
