package finance

import "errors"

var (
	ErrNegativeReceivables = errors.New("accounts receivable cannot be negative")
	ErrInvalidCreditSales  = errors.New("total credit sales must be greater than zero")
	ErrNegativePayables    = errors.New("accounts payable cannot be negative")
	ErrInvalidCOGS         = errors.New("COGS must be greater than zero")
	ErrInvalidDaysInPeriod = errors.New("days in period must be greater than zero")
	ErrNegativeDays        = errors.New("DIO, DSO, and DPO cannot be negative")
	ErrInvalidRate         = errors.New("rate must be between 0 and 1 inclusive")
	ErrNonFiniteInput      = errors.New("input values must be finite (not NaN or Inf)")
)
