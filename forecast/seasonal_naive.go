package forecast

// SeasonalNaive forecasts a period as the value observed one full season
// (or the nearest whole number of seasons) earlier:
//
//	Forecast = History[n - m + ((PeriodsAhead - 1) mod m)]
//
// where n = len(History) and m = SeasonLength. For PeriodsAhead in
// [1, m], this is simply the value from exactly one season ago; beyond
// that it wraps around to reuse the same seasonal position.
//
// History must contain at least SeasonLength periods. SeasonLength must
// be at least 2. PeriodsAhead must be greater than zero.
func SeasonalNaive(in SeasonalNaiveInput) (float64, error) {
	if in.SeasonLength < 2 {
		return 0, ErrInvalidSeasonLength
	}
	if len(in.History) < in.SeasonLength {
		return 0, ErrInsufficientHistory
	}
	if in.PeriodsAhead <= 0 {
		return 0, ErrInvalidPeriods
	}
	if err := validateNonNegative(in.History); err != nil {
		return 0, err
	}

	n := len(in.History)
	m := in.SeasonLength
	idx := n - m + ((in.PeriodsAhead - 1) % m)

	return in.History[idx], nil
}
