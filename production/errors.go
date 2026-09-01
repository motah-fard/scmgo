package production

import "errors"

var (
	ErrInvalidPlannedProductionTime = errors.New("planned production time must be greater than zero")
	ErrInvalidRunTime               = errors.New("run time must be greater than zero and cannot exceed planned production time")
	ErrNegativeIdealCycleTime       = errors.New("ideal cycle time cannot be negative")
	ErrInvalidTotalCount            = errors.New("total count must be greater than zero")
	ErrInvalidGoodCount             = errors.New("good count must be non-negative and cannot exceed total count")
	ErrNegativeWIP                  = errors.New("WIP cannot be negative")
	ErrInvalidThroughput            = errors.New("throughput must be greater than zero")
	ErrNegativeCycleTime            = errors.New("cycle time cannot be negative")
	ErrInvalidCycleTime             = errors.New("cycle time must be greater than zero")
	ErrInvalidAvailableTime         = errors.New("available production time must be greater than zero")
	ErrInvalidCustomerDemand        = errors.New("customer demand must be greater than zero")
	ErrNonFiniteInput               = errors.New("input values must be finite (not NaN or Inf)")
)
