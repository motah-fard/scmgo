package forecast

import "math"

// weightSumTolerance is the allowed floating-point slack when checking that
// weights sum to 1.
const weightSumTolerance = 1e-9

// WeightedMovingAverage calculates a weighted moving average forecast:
//
//	forecast = sum(weight_i * demand_i)
//
// over the most recent len(Weights) periods of History, with Weights[0]
// applied to the oldest of those periods. Weights must be non-negative and
// sum to 1.
func WeightedMovingAverage(in WeightedMovingAverageInput) (float64, error) {
	if len(in.Weights) == 0 {
		return 0, ErrInvalidPeriods
	}
	if len(in.History) < len(in.Weights) {
		return 0, ErrInsufficientHistory
	}
	if err := validateNonNegative(in.History); err != nil {
		return 0, err
	}
	if err := validateFinite(in.Weights...); err != nil {
		return 0, err
	}

	var weightSum float64
	for _, w := range in.Weights {
		if w < 0 {
			return 0, ErrInvalidWeight
		}
		weightSum += w
	}
	if math.Abs(weightSum-1) > weightSumTolerance {
		return 0, ErrWeightsMustSumToOne
	}

	window := in.History[len(in.History)-len(in.Weights):]

	var result float64
	for i, w := range in.Weights {
		result += w * window[i]
	}

	return result, nil
}
