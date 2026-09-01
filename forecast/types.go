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

// HoltWintersInput contains the inputs required to calculate a Holt-Winters
// (triple exponential smoothing, additive seasonality) forecast.
type HoltWintersInput struct {
	// History is the chronological demand series, oldest value first. Must
	// contain at least 2*SeasonLength periods (two full seasons), to
	// initialize the level, trend, and seasonal components.
	History []float64
	// Alpha is the level smoothing constant, in (0, 1].
	Alpha float64
	// Beta is the trend smoothing constant, in (0, 1].
	Beta float64
	// Gamma is the seasonal smoothing constant, in (0, 1].
	Gamma float64
	// SeasonLength is the number of periods per season (e.g. 12 for
	// monthly data with annual seasonality). Must be at least 2.
	SeasonLength int
	// PeriodsAhead is the forecast horizon, in periods, and must be
	// greater than zero.
	PeriodsAhead int
}

// HoltWintersResult contains the outputs of a Holt-Winters calculation.
type HoltWintersResult struct {
	// Level is the smoothed level estimate at the end of History.
	Level float64
	// Trend is the smoothed trend estimate at the end of History.
	Trend float64
	// Seasonal holds the SeasonLength most recent seasonal indices,
	// indexed by (period position mod SeasonLength).
	Seasonal []float64
	// Forecast is the forecast for PeriodsAhead periods beyond the end of
	// History.
	Forecast float64
}

// TrackingSignalInput contains the inputs required to calculate a running
// tracking signal.
type TrackingSignalInput struct {
	// Actual is the chronological series of actual observed values.
	Actual []float64
	// Forecast is the chronological series of forecast values, aligned
	// period-for-period with Actual. Must be the same length as Actual.
	Forecast []float64
}

// LinearTrendInput contains the inputs required to calculate a linear
// regression trend forecast.
type LinearTrendInput struct {
	// History is the chronological demand series, oldest value first. At
	// least two periods are required to fit a trend line.
	History []float64
	// PeriodsAhead is the forecast horizon, in periods, and must be
	// greater than zero.
	PeriodsAhead int
}

// LinearTrendResult contains the outputs of a linear regression trend
// calculation.
type LinearTrendResult struct {
	// Slope is the fitted trend's slope (change in value per period).
	Slope float64
	// Intercept is the fitted trend's value at period 0 (the first period
	// in History).
	Intercept float64
	// Forecast is the forecast for PeriodsAhead periods beyond the end of
	// History.
	Forecast float64
}

// MASEInput contains the inputs required to calculate the mean absolute
// scaled error (Hyndman & Koehler, 2006).
type MASEInput struct {
	// TrainingHistory is the in-sample series used to compute the naive
	// one-step-ahead forecast benchmark that scales the error. Must
	// contain at least two periods.
	TrainingHistory []float64
	// Actual is the series of actual values being forecast (in-sample or
	// out-of-sample).
	Actual []float64
	// Forecast is the corresponding forecast values, aligned
	// period-for-period with Actual. Must be the same length as Actual.
	Forecast []float64
}

// DemandClassificationInput contains the inputs required to classify a
// demand pattern.
type DemandClassificationInput struct {
	// History is the chronological demand series, oldest value first,
	// where zero denotes a period with no demand. At least one non-zero
	// period is required.
	History []float64
}

// DemandClassificationResult contains the outputs of a demand pattern
// classification.
type DemandClassificationResult struct {
	// ADI is the average demand interval: the number of periods in
	// History divided by the number of non-zero periods.
	ADI float64
	// CV2 is the squared coefficient of variation of demand size
	// (computed over non-zero periods only): (StdDev/Mean)².
	CV2 float64
	// Class is "smooth", "intermittent", "erratic", or "lumpy".
	Class string
}

// NaiveInput contains the inputs required to calculate a naive forecast.
type NaiveInput struct {
	// History is the chronological demand series, oldest value first.
	History []float64
}

// SeasonalNaiveInput contains the inputs required to calculate a seasonal
// naive forecast.
type SeasonalNaiveInput struct {
	// History is the chronological demand series, oldest value first.
	// Must contain at least SeasonLength periods.
	History []float64
	// SeasonLength is the number of periods per season. Must be at least
	// 2.
	SeasonLength int
	// PeriodsAhead is the forecast horizon, in periods, and must be
	// greater than zero.
	PeriodsAhead int
}
