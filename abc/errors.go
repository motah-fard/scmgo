package abc

import "errors"

var (
	ErrEmptyItems                  = errors.New("items must contain at least one entry")
	ErrNegativeValue               = errors.New("value cannot be negative")
	ErrZeroTotalValue              = errors.New("total value must be greater than zero")
	ErrInvalidThreshold            = errors.New("AThreshold must be in (0, 1) and BThreshold must be in (AThreshold, 1]")
	ErrInvalidVariabilityThreshold = errors.New("XThreshold must be greater than zero and YThreshold must be greater than XThreshold")
	ErrInvalidMeanDemand           = errors.New("mean demand must be greater than zero")
	ErrNegativeStandardDeviation   = errors.New("standard deviation cannot be negative")
	ErrNonFiniteInput              = errors.New("input values must be finite (not NaN or Inf)")
)
