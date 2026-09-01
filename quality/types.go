package quality

// DPMOInput contains the inputs required to calculate defects per million
// opportunities.
type DPMOInput struct {
	NumberOfDefects      float64
	NumberOfUnits        float64
	OpportunitiesPerUnit float64
}

// CpInput contains the inputs required to calculate the process
// capability index Cp.
type CpInput struct {
	USL   float64
	LSL   float64
	Sigma float64
}

// CpkInput contains the inputs required to calculate the process
// capability index Cpk.
type CpkInput struct {
	USL   float64
	LSL   float64
	Mean  float64
	Sigma float64
}

// CostOfQualityInput contains the inputs required to calculate cost of
// quality.
type CostOfQualityInput struct {
	PreventionCost      float64
	AppraisalCost       float64
	InternalFailureCost float64
	ExternalFailureCost float64
}

// CostOfQualityResult contains the outputs of a cost of quality
// calculation.
type CostOfQualityResult struct {
	Total               float64
	PreventionCost      float64
	AppraisalCost       float64
	InternalFailureCost float64
	ExternalFailureCost float64
	// ConformanceCost is PreventionCost + AppraisalCost: the cost of
	// preventing and catching defects.
	ConformanceCost float64
	// NonConformanceCost is InternalFailureCost + ExternalFailureCost:
	// the cost of defects that occurred.
	NonConformanceCost float64
}
