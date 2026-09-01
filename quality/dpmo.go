package quality

// DPMO calculates defects per million opportunities:
//
//	DPMO = (NumberOfDefects / (NumberOfUnits × OpportunitiesPerUnit)) × 1,000,000
//
// NumberOfDefects must be non-negative. NumberOfUnits and
// OpportunitiesPerUnit must be greater than zero.
func DPMO(in DPMOInput) (float64, error) {
	if err := validateFinite(in.NumberOfDefects, in.NumberOfUnits, in.OpportunitiesPerUnit); err != nil {
		return 0, err
	}
	if in.NumberOfDefects < 0 {
		return 0, ErrNegativeDefects
	}
	if in.NumberOfUnits <= 0 {
		return 0, ErrInvalidUnits
	}
	if in.OpportunitiesPerUnit <= 0 {
		return 0, ErrInvalidOpportunities
	}

	return (in.NumberOfDefects / (in.NumberOfUnits * in.OpportunitiesPerUnit)) * 1_000_000, nil
}
