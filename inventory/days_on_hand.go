package inventory

// DaysOfInventoryOnHand calculates how many days of inventory are on hand
// using:
//
//	days on hand = (average inventory value / COGS) × days in period
//
// AverageInventoryValue must be non-negative. COGS and DaysInPeriod (both
// over the same period, e.g. 365 for a year) must be greater than zero.
func DaysOfInventoryOnHand(in DaysOfInventoryOnHandInput) (float64, error) {
	if err := validateFinite(in.AverageInventoryValue, in.COGS, in.DaysInPeriod); err != nil {
		return 0, err
	}
	if in.AverageInventoryValue < 0 {
		return 0, ErrNegativeAverageInventoryValue
	}
	if in.COGS <= 0 {
		return 0, ErrInvalidCOGS
	}
	if in.DaysInPeriod <= 0 {
		return 0, ErrInvalidDaysInPeriod
	}

	return (in.AverageInventoryValue / in.COGS) * in.DaysInPeriod, nil
}
