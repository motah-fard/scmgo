package quality

import "errors"

var (
	ErrNegativeDefects      = errors.New("number of defects cannot be negative")
	ErrInvalidUnits         = errors.New("number of units must be greater than zero")
	ErrInvalidOpportunities = errors.New("opportunities per unit must be greater than zero")
	ErrInvalidSpecLimits    = errors.New("upper spec limit must be greater than lower spec limit")
	ErrInvalidSigma         = errors.New("sigma must be greater than zero")
	ErrNegativeCost         = errors.New("cost cannot be negative")
	ErrNonFiniteInput       = errors.New("input values must be finite (not NaN or Inf)")
)
