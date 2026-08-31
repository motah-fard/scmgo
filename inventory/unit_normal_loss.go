package inventory

import "math"

// UnitNormalLoss returns the standard normal unit loss function value:
//
//	L(z) = φ(z) - z × (1 - Φ(z))
//
// where φ is the standard normal probability density function and Φ is the
// standard normal cumulative distribution function. L(z) is the expected
// shortfall, in standard deviations, of a standard normal variable below z;
// it is used to compute expected demand shortage under a normal
// approximation (see ExpectedFillRate, FillRateSafetyStock).
//
// L is strictly decreasing: L(z) → -z as z → -∞ and L(z) → 0 as z → +∞.
func UnitNormalLoss(z float64) (float64, error) {
	if err := validateFinite(z); err != nil {
		return 0, err
	}

	pdf := math.Exp(-z*z/2) / math.Sqrt(2*math.Pi)
	// 1 - Φ(z) computed via Erfc rather than 1 - (1 + Erf(...))/2: for
	// large z, Φ(z) rounds to exactly 1 in float64, so subtracting it from
	// 1 would silently lose the entire (still relevant) tail term. Erfc
	// computes the tail directly, without ever subtracting two
	// nearly-equal numbers.
	survival := 0.5 * math.Erfc(z/math.Sqrt2)

	return pdf - z*survival, nil
}
