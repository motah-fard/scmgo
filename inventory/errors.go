package inventory

import "errors"

var (
	ErrNegativeDemand        = errors.New("demand cannot be negative")
	ErrNegativeLeadTime      = errors.New("lead time cannot be negative")
	ErrNegativeSafetyStock   = errors.New("safety stock cannot be negative")
	ErrNegativeOrderingCost  = errors.New("ordering cost cannot be negative")
	ErrInvalidHoldingCost    = errors.New("holding cost must be greater than zero")
	ErrNegativeOrderQuantity = errors.New("order quantity cannot be negative")
	ErrNegativeReorderPoint  = errors.New("reorder point cannot be negative")
)
