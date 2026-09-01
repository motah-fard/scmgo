package finance

// DSO calculates days sales outstanding:
//
//	DSO = (AccountsReceivable / TotalCreditSales) × DaysInPeriod
//
// AccountsReceivable must be non-negative. TotalCreditSales and
// DaysInPeriod (both over the same period) must be greater than zero.
func DSO(in DSOInput) (float64, error) {
	if err := validateFinite(in.AccountsReceivable, in.TotalCreditSales, in.DaysInPeriod); err != nil {
		return 0, err
	}
	if in.AccountsReceivable < 0 {
		return 0, ErrNegativeReceivables
	}
	if in.TotalCreditSales <= 0 {
		return 0, ErrInvalidCreditSales
	}
	if in.DaysInPeriod <= 0 {
		return 0, ErrInvalidDaysInPeriod
	}

	return (in.AccountsReceivable / in.TotalCreditSales) * in.DaysInPeriod, nil
}
