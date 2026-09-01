package production

// TaktTime calculates the pace of production needed to match customer
// demand:
//
//	TaktTime = AvailableProductionTime / CustomerDemand
//
// AvailableProductionTime and CustomerDemand must both be greater than
// zero.
func TaktTime(in TaktTimeInput) (float64, error) {
	if err := validateFinite(in.AvailableProductionTime, in.CustomerDemand); err != nil {
		return 0, err
	}
	if in.AvailableProductionTime <= 0 {
		return 0, ErrInvalidAvailableTime
	}
	if in.CustomerDemand <= 0 {
		return 0, ErrInvalidCustomerDemand
	}

	return in.AvailableProductionTime / in.CustomerDemand, nil
}
