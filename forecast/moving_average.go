package forecast

// MovingAverage calculates a simple moving average forecast:
//
//	forecast = (sum of the most recent N periods) / N
//
// where N is Periods. History must contain at least Periods values.
func MovingAverage(in MovingAverageInput) (float64, error) {
	if in.Periods <= 0 {
		return 0, ErrInvalidPeriods
	}
	if len(in.History) < in.Periods {
		return 0, ErrInsufficientHistory
	}
	if err := validateNonNegative(in.History); err != nil {
		return 0, err
	}

	window := in.History[len(in.History)-in.Periods:]

	var sum float64
	for _, v := range window {
		sum += v
	}

	return sum / float64(in.Periods), nil
}
