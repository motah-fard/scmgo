package forecast

// SBA forecasts intermittent demand using the Syntetos-Boylan
// Approximation, a bias correction to Croston's method:
//
//	SBA forecast = Croston forecast × (1 - Alpha/2)
//
// Croston's method is known to be biased (it systematically over-forecasts
// intermittent demand); SBA is the standard correction. DemandSize and
// Interval are unchanged from the underlying Croston calculation — only
// the final forecast is adjusted.
//
// Same validation as Croston: Alpha must be in (0, 1]; History must
// contain at least one non-zero period.
func SBA(in CrostonInput) (CrostonResult, error) {
	result, err := Croston(in)
	if err != nil {
		return CrostonResult{}, err
	}

	result.Forecast *= 1 - in.Alpha/2

	return result, nil
}
