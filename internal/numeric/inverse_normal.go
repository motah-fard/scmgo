package numeric

import "math"

// InverseNormalCDF returns the standard normal inverse cumulative
// distribution function (probit) at p: the z such that Φ(z) = p.
//
// It is the shared building block behind inventory.ZScoreForServiceLevel
// and quality.SigmaLevel. Callers are responsible for validating p is
// finite and strictly inside (0, 1) before calling this — outside that
// range, math.Erfinv (which this is built on) returns NaN or an infinity,
// and this function does not guard against that itself.
func InverseNormalCDF(p float64) float64 {
	return math.Sqrt2 * math.Erfinv(2*p-1)
}
