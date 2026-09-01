package finance

// DPO calculates days payable outstanding:
//
//	DPO = (AccountsPayable / COGS) × DaysInPeriod
//
// AccountsPayable must be non-negative. COGS and DaysInPeriod (both over
// the same period) must be greater than zero.
func DPO(in DPOInput) (float64, error) {
	if err := validateFinite(in.AccountsPayable, in.COGS, in.DaysInPeriod); err != nil {
		return 0, err
	}
	if in.AccountsPayable < 0 {
		return 0, ErrNegativePayables
	}
	if in.COGS <= 0 {
		return 0, ErrInvalidCOGS
	}
	if in.DaysInPeriod <= 0 {
		return 0, ErrInvalidDaysInPeriod
	}

	return (in.AccountsPayable / in.COGS) * in.DaysInPeriod, nil
}
