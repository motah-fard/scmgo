package quality

import "github.com/motah-fard/scmgo/internal/numeric"

// sigmaShift is the conventional long-term/short-term process shift used
// when converting a yield to a "Six Sigma" sigma level.
const sigmaShift = 1.5

// SigmaLevel converts defects per million opportunities to a process
// sigma level, using the conventional 1.5-sigma shift:
//
//	Yield = 1 - DPMO / 1,000,000
//	SigmaLevel = Φ⁻¹(Yield) + 1.5
//
// DPMO must be strictly between 0 and 1,000,000: at either boundary, Yield
// is exactly 1 or 0 and Φ⁻¹ is undefined (an infinity).
//
// Reference points: DPMO = 3.4 -> ~6.0 sigma (the classic "Six Sigma"
// target); DPMO = 66,807 -> ~3.0 sigma.
func SigmaLevel(dpmo float64) (float64, error) {
	if err := validateFinite(dpmo); err != nil {
		return 0, err
	}
	if dpmo <= 0 || dpmo >= 1_000_000 {
		return 0, ErrInvalidDPMO
	}

	yield := 1 - dpmo/1_000_000

	return numeric.InverseNormalCDF(yield) + sigmaShift, nil
}
