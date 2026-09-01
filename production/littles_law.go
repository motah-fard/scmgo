package production

// WIPFromLittlesLaw calculates work in progress from Little's Law:
//
//	WIP = Throughput × CycleTime
//
// Throughput must be greater than zero. CycleTime must be non-negative.
func WIPFromLittlesLaw(in WIPFromLittlesLawInput) (float64, error) {
	if err := validateFinite(in.Throughput, in.CycleTime); err != nil {
		return 0, err
	}
	if in.Throughput <= 0 {
		return 0, ErrInvalidThroughput
	}
	if in.CycleTime < 0 {
		return 0, ErrNegativeCycleTime
	}

	return in.Throughput * in.CycleTime, nil
}

// CycleTimeFromLittlesLaw calculates cycle time from Little's Law:
//
//	CycleTime = WIP / Throughput
//
// WIP must be non-negative. Throughput must be greater than zero.
func CycleTimeFromLittlesLaw(in CycleTimeFromLittlesLawInput) (float64, error) {
	if err := validateFinite(in.WIP, in.Throughput); err != nil {
		return 0, err
	}
	if in.WIP < 0 {
		return 0, ErrNegativeWIP
	}
	if in.Throughput <= 0 {
		return 0, ErrInvalidThroughput
	}

	return in.WIP / in.Throughput, nil
}

// ThroughputFromLittlesLaw calculates throughput from Little's Law:
//
//	Throughput = WIP / CycleTime
//
// WIP must be non-negative. CycleTime must be greater than zero.
func ThroughputFromLittlesLaw(in ThroughputFromLittlesLawInput) (float64, error) {
	if err := validateFinite(in.WIP, in.CycleTime); err != nil {
		return 0, err
	}
	if in.WIP < 0 {
		return 0, ErrNegativeWIP
	}
	if in.CycleTime <= 0 {
		return 0, ErrInvalidCycleTime
	}

	return in.WIP / in.CycleTime, nil
}
