package inventory

import "math"

// ZScoreForServiceLevel returns the standard normal z-score
// corresponding to the given cycle service level.
//
// The service level must be strictly between 0 and 1.
//
// Examples:
//   - 0.90  -> 1.2816
//   - 0.95  -> 1.6449
//   - 0.975 -> 1.9600
//   - 0.99  -> 2.3263
func ZScoreForServiceLevel(serviceLevel float64) (float64, error) {
	if serviceLevel <= 0 || serviceLevel >= 1 {
		return 0, ErrInvalidServiceLevel
	}

	z := math.Sqrt2 * math.Erfinv(2*serviceLevel-1)
	return z, nil
}
