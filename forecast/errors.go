package forecast

import "errors"

var (
	ErrEmptyHistory             = errors.New("history must contain at least one period")
	ErrInsufficientHistory      = errors.New("history does not contain enough periods for the requested window")
	ErrInvalidPeriods           = errors.New("periods must be greater than zero")
	ErrNegativeDemand           = errors.New("demand cannot be negative")
	ErrInvalidWeight            = errors.New("weights cannot be negative")
	ErrWeightsMustSumToOne      = errors.New("weights must sum to 1")
	ErrInvalidSmoothingConstant = errors.New("smoothing constant must be greater than 0 and less than or equal to 1")
	ErrNoNonZeroDemand          = errors.New("history must contain at least one non-zero demand period")
	ErrMismatchedLengths        = errors.New("actual and forecast series must have the same length")
	ErrNonFiniteInput           = errors.New("input values must be finite (not NaN or Inf)")
)
