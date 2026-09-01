package finance

// PerfectOrderRate calculates the probability that an order is delivered
// on time, complete, damage-free, and accurately documented:
//
//	PerfectOrderRate = OnTimeRate × CompleteRate × DamageFreeRate × AccurateDocumentationRate
//
// This multiplies the four component rates, which assumes they are
// independent — a standard simplification. If failure modes are
// correlated in your operation (e.g. a late order is also more likely to
// arrive damaged), treat the result as an upper-bound estimate.
//
// Each rate must be in [0, 1].
func PerfectOrderRate(in PerfectOrderRateInput) (float64, error) {
	if err := validateFinite(in.OnTimeRate, in.CompleteRate, in.DamageFreeRate, in.AccurateDocumentationRate); err != nil {
		return 0, err
	}
	for _, rate := range []float64{in.OnTimeRate, in.CompleteRate, in.DamageFreeRate, in.AccurateDocumentationRate} {
		if rate < 0 || rate > 1 {
			return 0, ErrInvalidRate
		}
	}

	return in.OnTimeRate * in.CompleteRate * in.DamageFreeRate * in.AccurateDocumentationRate, nil
}
