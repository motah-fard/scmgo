// Package forecast provides practical demand-forecasting calculations that
// feed the inputs of the inventory package (AvgDailyDemand, StdDevDailyDemand,
// and friends).
//
// The package includes:
//
//   - Moving average
//   - Weighted moving average
//   - Simple exponential smoothing
//   - Holt's linear trend method (double exponential smoothing)
//   - Croston's method for intermittent demand
//   - Forecast accuracy metrics (MAD, MAPE, Bias, RMSE)
//
// Important assumptions:
//
//   - History series are chronological, oldest value first.
//   - Demand values must be non-negative.
//   - Smoothing constants (Alpha, Beta) must be in the interval (0, 1].
//   - All float64 inputs must be finite: NaN and +/-Inf are rejected with
//     ErrNonFiniteInput (this includes Actual/Forecast in Accuracy,
//     even though its MAPE output can legitimately be NaN when every actual
//     value is zero -- that is a documented output, not a silent input bug).
//   - MovingAverage and WeightedMovingAverage require at least as much
//     history as the requested window.
//   - HoltLinearTrend requires at least two historical periods to estimate
//     an initial trend.
//   - Croston assumes zero values represent "no demand" periods and
//     estimates demand size and inter-demand interval separately; it
//     requires at least one non-zero period.
//   - Accuracy's MAPE excludes periods where the actual value is
//     zero (division is undefined) and reports NaN if every actual value is
//     zero; MAD, Bias, and RMSE are always computed. MAPE and Bias are
//     signed/expressed as fractions, not percentages (e.g. 0.05 = 5%).
//   - This package does not perform model selection, parameter fitting
//     (e.g. optimal alpha/beta search), seasonality decomposition, or
//     ARIMA/ML-based forecasting. It computes the classical formulas given
//     caller-supplied parameters.
package forecast
