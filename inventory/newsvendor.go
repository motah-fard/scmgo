package inventory

// Newsvendor calculates the optimal single-period order quantity under the
// classic newsvendor (single-period, normal demand approximation) model:
//
//	critical ratio = underage cost / (underage cost + overage cost)
//	order quantity = mean demand + z × standard deviation of demand
//
// where z is the standard normal value at the critical ratio (see
// ZScoreForServiceLevel — the critical ratio plays the same role here that
// a target service level plays elsewhere in this package).
//
// MeanDemand and StdDevDemand must be non-negative. UnderageCostPerUnit
// (the cost of ordering one unit too few — typically lost margin) and
// OverageCostPerUnit (the cost of ordering one unit too many — typically
// net of salvage value) must both be greater than zero.
func Newsvendor(in NewsvendorInput) (NewsvendorResult, error) {
	if err := validateFinite(in.MeanDemand, in.StdDevDemand, in.UnderageCostPerUnit, in.OverageCostPerUnit); err != nil {
		return NewsvendorResult{}, err
	}
	if in.MeanDemand < 0 {
		return NewsvendorResult{}, ErrNegativeDemand
	}
	if in.StdDevDemand < 0 {
		return NewsvendorResult{}, ErrNegativeStandardDeviation
	}
	if in.UnderageCostPerUnit <= 0 {
		return NewsvendorResult{}, ErrInvalidUnderageCost
	}
	if in.OverageCostPerUnit <= 0 {
		return NewsvendorResult{}, ErrInvalidOverageCost
	}

	criticalRatio := in.UnderageCostPerUnit / (in.UnderageCostPerUnit + in.OverageCostPerUnit)

	z, err := ZScoreForServiceLevel(criticalRatio)
	if err != nil {
		return NewsvendorResult{}, err
	}

	return NewsvendorResult{
		OrderQuantity: in.MeanDemand + z*in.StdDevDemand,
		CriticalRatio: criticalRatio,
	}, nil
}
