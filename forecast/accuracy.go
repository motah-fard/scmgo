package forecast

import "math"

// ForecastAccuracy calculates common forecast accuracy metrics (MAD, MAPE,
// Bias, RMSE) from a period-aligned actual/forecast series pair. See
// ForecastAccuracyResult for the definition of each metric, including how
// MAPE handles zero actual values.
func ForecastAccuracy(in ForecastAccuracyInput) (ForecastAccuracyResult, error) {
	if len(in.Actual) == 0 {
		return ForecastAccuracyResult{}, ErrEmptyHistory
	}
	if len(in.Actual) != len(in.Forecast) {
		return ForecastAccuracyResult{}, ErrMismatchedLengths
	}
	if err := validateFiniteSlice(in.Actual); err != nil {
		return ForecastAccuracyResult{}, err
	}
	if err := validateFiniteSlice(in.Forecast); err != nil {
		return ForecastAccuracyResult{}, err
	}

	var (
		absSum     float64
		signedSum  float64
		squaredSum float64
		mapeSum    float64
		mapeCount  int
	)

	n := float64(len(in.Actual))

	for i, actual := range in.Actual {
		deviation := actual - in.Forecast[i]
		absSum += math.Abs(deviation)
		signedSum += deviation
		squaredSum += deviation * deviation

		if actual != 0 {
			mapeSum += math.Abs(deviation) / actual
			mapeCount++
		}
	}

	mape := math.NaN()
	if mapeCount > 0 {
		mape = mapeSum / float64(mapeCount)
	}

	return ForecastAccuracyResult{
		MAD:  absSum / n,
		MAPE: mape,
		Bias: signedSum / n,
		RMSE: math.Sqrt(squaredSum / n),
	}, nil
}
