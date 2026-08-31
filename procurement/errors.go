package procurement

import "errors"

var (
	ErrEmptyComponents  = errors.New("components must contain at least one entry")
	ErrNegativeCost     = errors.New("cost cannot be negative")
	ErrNegativeQuantity = errors.New("quantity cannot be negative")
	ErrNonFiniteInput   = errors.New("input values must be finite (not NaN or Inf)")
)
