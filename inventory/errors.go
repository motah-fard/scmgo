package inventory

import "errors"

var (
	ErrNegativeDemand            = errors.New("demand cannot be negative")
	ErrNegativeLeadTime          = errors.New("lead time cannot be negative")
	ErrNegativeSafetyStock       = errors.New("safety stock cannot be negative")
	ErrNegativeOrderingCost      = errors.New("ordering cost cannot be negative")
	ErrInvalidOrderingCost       = errors.New("ordering cost must be greater than zero")
	ErrInvalidHoldingCost        = errors.New("holding cost must be greater than zero")
	ErrNegativeOrderQuantity     = errors.New("order quantity cannot be negative")
	ErrNegativeReorderPoint      = errors.New("reorder point cannot be negative")
	ErrInvalidServiceLevel       = errors.New("service level must be between 0 and 1 exclusive")
	ErrNegativeStandardDeviation = errors.New("standard deviation cannot be negative")
	ErrNegativeExpectedDemand    = errors.New("expected demand cannot be negative")
	ErrInvalidPolicySummaryInput = errors.New("invalid policy summary input")
	ErrNegativeReviewPeriod      = errors.New("review period cannot be negative")
)
