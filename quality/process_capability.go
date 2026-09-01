package quality

import "math"

// Cp calculates the process capability index, comparing the tolerance
// width to process spread:
//
//	Cp = (USL - LSL) / (6 × Sigma)
//
// USL must be greater than LSL. Sigma must be greater than zero.
func Cp(in CpInput) (float64, error) {
	if err := validateFinite(in.USL, in.LSL, in.Sigma); err != nil {
		return 0, err
	}
	if in.USL <= in.LSL {
		return 0, ErrInvalidSpecLimits
	}
	if in.Sigma <= 0 {
		return 0, ErrInvalidSigma
	}

	return (in.USL - in.LSL) / (6 * in.Sigma), nil
}

// Cpk calculates the process capability index accounting for centering,
// not just spread:
//
//	Cpk = min[(USL - Mean) / (3 × Sigma), (Mean - LSL) / (3 × Sigma)]
//
// USL must be greater than LSL. Sigma must be greater than zero. Mean may
// be any finite value, including outside [LSL, USL] — a negative result
// reflects a genuinely out-of-tolerance process, not invalid input.
func Cpk(in CpkInput) (float64, error) {
	if err := validateFinite(in.USL, in.LSL, in.Mean, in.Sigma); err != nil {
		return 0, err
	}
	if in.USL <= in.LSL {
		return 0, ErrInvalidSpecLimits
	}
	if in.Sigma <= 0 {
		return 0, ErrInvalidSigma
	}

	upper := (in.USL - in.Mean) / (3 * in.Sigma)
	lower := (in.Mean - in.LSL) / (3 * in.Sigma)

	return math.Min(upper, lower), nil
}
