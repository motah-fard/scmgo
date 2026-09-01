package quality

// CostOfQuality calculates total cost of quality from the standard
// four-category breakdown:
//
//	Total              = PreventionCost + AppraisalCost + InternalFailureCost + ExternalFailureCost
//	ConformanceCost    = PreventionCost + AppraisalCost
//	NonConformanceCost = InternalFailureCost + ExternalFailureCost
//
// All four cost inputs must be non-negative.
func CostOfQuality(in CostOfQualityInput) (CostOfQualityResult, error) {
	if err := validateFinite(in.PreventionCost, in.AppraisalCost, in.InternalFailureCost, in.ExternalFailureCost); err != nil {
		return CostOfQualityResult{}, err
	}
	if in.PreventionCost < 0 || in.AppraisalCost < 0 || in.InternalFailureCost < 0 || in.ExternalFailureCost < 0 {
		return CostOfQualityResult{}, ErrNegativeCost
	}

	conformance := in.PreventionCost + in.AppraisalCost
	nonConformance := in.InternalFailureCost + in.ExternalFailureCost

	return CostOfQualityResult{
		Total:               conformance + nonConformance,
		PreventionCost:      in.PreventionCost,
		AppraisalCost:       in.AppraisalCost,
		InternalFailureCost: in.InternalFailureCost,
		ExternalFailureCost: in.ExternalFailureCost,
		ConformanceCost:     conformance,
		NonConformanceCost:  nonConformance,
	}, nil
}
