package logistics

import "errors"

var (
	ErrInvalidDimension = errors.New("length, width, and height must be greater than zero")
	ErrInvalidDimFactor = errors.New("dim factor must be greater than zero")
	ErrNegativeWeight   = errors.New("weight cannot be negative")
	ErrEmptyItems       = errors.New("items must contain at least one entry")
	ErrNegativeCost     = errors.New("total freight cost cannot be negative")
	ErrZeroTotalWeight  = errors.New("total weight across items must be greater than zero")
	ErrNegativeLoad     = errors.New("actual load cannot be negative")
	ErrInvalidCapacity  = errors.New("capacity must be greater than zero")
	ErrEmptyLocations   = errors.New("locations must contain at least one entry")
	ErrNegativeDemand   = errors.New("demand cannot be negative")
	ErrZeroTotalDemand  = errors.New("total demand across locations must be greater than zero")
	ErrNonFiniteInput   = errors.New("input values must be finite (not NaN or Inf)")
)
