package finance

// CashToCashCycleTime calculates the cash-to-cash (cash conversion) cycle
// time:
//
//	CashToCashCycleTime = DIO + DSO - DPO
//
// DIO, DSO, and DPO must each be non-negative. The result itself is not
// restricted in sign: a negative cash-to-cash cycle time is a real,
// favorable outcome (paying suppliers after collecting from customers),
// not an error.
func CashToCashCycleTime(in CashToCashCycleTimeInput) (float64, error) {
	if err := validateFinite(in.DIO, in.DSO, in.DPO); err != nil {
		return 0, err
	}
	if in.DIO < 0 || in.DSO < 0 || in.DPO < 0 {
		return 0, ErrNegativeDays
	}

	return in.DIO + in.DSO - in.DPO, nil
}
