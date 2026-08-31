package forecast

import "math"

// MASE calculates the mean absolute scaled error (Hyndman & Koehler,
// 2006):
//
//	naive MAE = mean(|TrainingHistory[i] - TrainingHistory[i-1]|)  for i = 1..len(TrainingHistory)-1
//	MASE      = mean(|Actual[i] - Forecast[i]|) / naive MAE
//
// MASE expresses forecast error relative to a naive one-step-ahead
// forecast on the training series: MASE < 1 means the forecast
// outperforms the naive benchmark; MASE > 1 means it underperforms.
// Unlike MAPE, MASE is well-defined even when Actual contains zeros.
//
// TrainingHistory must contain at least two periods and must not be
// perfectly constant (a constant series has a naive MAE of zero, making
// MASE's denominator zero — see ErrZeroNaiveMAE). Actual and Forecast must
// be the same non-zero length.
func MASE(in MASEInput) (float64, error) {
	if len(in.TrainingHistory) < 2 {
		return 0, ErrInsufficientHistory
	}
	if len(in.Actual) == 0 {
		return 0, ErrEmptyHistory
	}
	if len(in.Actual) != len(in.Forecast) {
		return 0, ErrMismatchedLengths
	}
	if err := validateFinite(in.TrainingHistory...); err != nil {
		return 0, err
	}
	if err := validateFinite(in.Actual...); err != nil {
		return 0, err
	}
	if err := validateFinite(in.Forecast...); err != nil {
		return 0, err
	}

	var naiveSum float64
	for i := 1; i < len(in.TrainingHistory); i++ {
		naiveSum += math.Abs(in.TrainingHistory[i] - in.TrainingHistory[i-1])
	}
	naiveMAE := naiveSum / float64(len(in.TrainingHistory)-1)
	if naiveMAE == 0 {
		return 0, ErrZeroNaiveMAE
	}

	var errSum float64
	for i, actual := range in.Actual {
		errSum += math.Abs(actual - in.Forecast[i])
	}
	mae := errSum / float64(len(in.Actual))

	return mae / naiveMAE, nil
}
