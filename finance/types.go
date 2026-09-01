package finance

// DSOInput contains the inputs required to calculate days sales
// outstanding.
type DSOInput struct {
	AccountsReceivable float64
	TotalCreditSales   float64
	DaysInPeriod       float64
}

// DPOInput contains the inputs required to calculate days payable
// outstanding.
type DPOInput struct {
	AccountsPayable float64
	COGS            float64
	DaysInPeriod    float64
}

// CashToCashCycleTimeInput contains the inputs required to calculate the
// cash-to-cash cycle time.
type CashToCashCycleTimeInput struct {
	// DIO is days inventory outstanding (see e.g.
	// github.com/motah-fard/scmgo/inventory's DaysOfInventoryOnHand).
	DIO float64
	DSO float64
	DPO float64
}

// PerfectOrderRateInput contains the inputs required to calculate perfect
// order rate.
type PerfectOrderRateInput struct {
	// OnTimeRate is the fraction of orders delivered on time, in [0, 1].
	OnTimeRate float64
	// CompleteRate is the fraction of orders delivered complete, in [0, 1].
	CompleteRate float64
	// DamageFreeRate is the fraction of orders delivered without damage,
	// in [0, 1].
	DamageFreeRate float64
	// AccurateDocumentationRate is the fraction of orders with accurate
	// documentation, in [0, 1].
	AccurateDocumentationRate float64
}
