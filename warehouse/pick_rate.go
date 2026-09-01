package warehouse

// PickRate calculates picks completed per hour:
//
//	PickRate = TotalPicks / TotalHours
//
// TotalPicks must be non-negative. TotalHours must be greater than zero.
func PickRate(in PickRateInput) (float64, error) {
	if err := validateFinite(in.TotalPicks, in.TotalHours); err != nil {
		return 0, err
	}
	if in.TotalPicks < 0 {
		return 0, ErrNegativePicks
	}
	if in.TotalHours <= 0 {
		return 0, ErrInvalidHours
	}

	return in.TotalPicks / in.TotalHours, nil
}
