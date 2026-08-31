package forecast

// Croston forecasts intermittent (sparse, lumpy) demand using Croston's
// method. It separates the demand series into non-zero demand sizes and the
// intervals between them, smooths each independently, and combines them
// into a per-period forecast:
//
//	Z[i] = alpha*demand_i    + (1-alpha)*Z[i-1]   (smoothed demand size)
//	P[i] = alpha*interval_i  + (1-alpha)*P[i-1]   (smoothed inter-demand interval)
//	Forecast = Z / P
//
// The first non-zero demand initializes Z, and the number of periods from
// the start of History up to and including it initializes P. Alpha must be
// in (0, 1]. History must contain at least one non-zero value.
func Croston(in CrostonInput) (CrostonResult, error) {
	if len(in.History) == 0 {
		return CrostonResult{}, ErrEmptyHistory
	}
	if err := validateSmoothingConstant(in.Alpha); err != nil {
		return CrostonResult{}, err
	}
	if err := validateNonNegative(in.History); err != nil {
		return CrostonResult{}, err
	}

	var (
		demandSize  float64
		interval    float64
		havePrior   bool
		lastNonZero int
	)

	for i, v := range in.History {
		if v == 0 {
			continue
		}

		if !havePrior {
			demandSize = v
			interval = float64(i + 1)
			havePrior = true
		} else {
			gap := float64(i - lastNonZero)
			demandSize = in.Alpha*v + (1-in.Alpha)*demandSize
			interval = in.Alpha*gap + (1-in.Alpha)*interval
		}
		lastNonZero = i
	}

	if !havePrior {
		return CrostonResult{}, ErrNoNonZeroDemand
	}

	return CrostonResult{
		DemandSize: demandSize,
		Interval:   interval,
		Forecast:   demandSize / interval,
	}, nil
}
