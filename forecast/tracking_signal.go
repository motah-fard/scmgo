package forecast

import "math"

// TrackingSignal calculates a running tracking signal, one value per
// period, for monitoring whether a forecast is drifting out of control:
//
//	TS[t] = (cumulative signed error through t) / (cumulative MAD through t)
//
// where cumulative signed error is the running sum of (Actual[i] -
// Forecast[i]) for i = 0..t, and cumulative MAD is the running mean of
// |Actual[i] - Forecast[i]| for i = 0..t. TS[t] is 0 if the cumulative MAD
// through t is 0 (every error so far has been exactly 0, so there is
// nothing to signal).
//
// A tracking signal conventionally outside roughly ±4 to ±8 (depending on
// the control limits chosen by the caller) suggests the forecast model is
// no longer tracking demand well.
//
// Actual and Forecast must be the same non-zero length.
func TrackingSignal(in TrackingSignalInput) ([]float64, error) {
	if len(in.Actual) == 0 {
		return nil, ErrEmptyHistory
	}
	if len(in.Actual) != len(in.Forecast) {
		return nil, ErrMismatchedLengths
	}
	if err := validateFinite(in.Actual...); err != nil {
		return nil, err
	}
	if err := validateFinite(in.Forecast...); err != nil {
		return nil, err
	}

	result := make([]float64, len(in.Actual))
	var cumError, cumAbsError float64
	for i, actual := range in.Actual {
		deviation := actual - in.Forecast[i]
		cumError += deviation
		cumAbsError += math.Abs(deviation)

		mad := cumAbsError / float64(i+1)
		if mad == 0 {
			result[i] = 0
			continue
		}
		result[i] = cumError / mad
	}

	return result, nil
}
