package inventory

// EOI calculates the economic order interval — EOQ expressed as a time
// interval between orders rather than a quantity:
//
//	EOI = (EOQ / AnnualDemand) × DaysPerYear
//
// AnnualDemand must be greater than zero (unlike EOQ's own AnnualDemand,
// which allows zero — EOI divides by it, so zero is undefined here).
// OrderingCost must be non-negative. HoldingCostPerUnit must be greater
// than zero. DaysPerYear must be greater than zero.
func EOI(in EOIInput) (float64, error) {
	if err := validateFinite(in.DaysPerYear); err != nil {
		return 0, err
	}
	if in.AnnualDemand <= 0 {
		return 0, ErrInvalidAnnualDemand
	}
	if in.DaysPerYear <= 0 {
		return 0, ErrInvalidDaysInPeriod
	}

	eoq, err := EOQ(EOQInput{
		AnnualDemand:       in.AnnualDemand,
		OrderingCost:       in.OrderingCost,
		HoldingCostPerUnit: in.HoldingCostPerUnit,
	})
	if err != nil {
		return 0, err
	}

	return (eoq / in.AnnualDemand) * in.DaysPerYear, nil
}
