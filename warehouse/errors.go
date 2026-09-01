package warehouse

import "errors"

var (
	ErrNegativeUsedSpace  = errors.New("used space cannot be negative")
	ErrInvalidTotalSpace  = errors.New("total space must be greater than zero")
	ErrNegativeUsedVolume = errors.New("used volume cannot be negative")
	ErrInvalidTotalVolume = errors.New("total volume must be greater than zero")
	ErrNegativePicks      = errors.New("total picks cannot be negative")
	ErrInvalidHours       = errors.New("total hours must be greater than zero")
	ErrNonFiniteInput     = errors.New("input values must be finite (not NaN or Inf)")
)
