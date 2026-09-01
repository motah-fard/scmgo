package logistics

// VehicleUtilization calculates how much of a vehicle's capacity is used:
//
//	Utilization = ActualLoad / Capacity
//
// ActualLoad must be non-negative. Capacity must be greater than zero.
//
// The result is not clamped: a value above 1 indicates an overloaded
// vehicle, which is meaningful information, not an error.
func VehicleUtilization(in VehicleUtilizationInput) (float64, error) {
	if err := validateFinite(in.ActualLoad, in.Capacity); err != nil {
		return 0, err
	}
	if in.ActualLoad < 0 {
		return 0, ErrNegativeLoad
	}
	if in.Capacity <= 0 {
		return 0, ErrInvalidCapacity
	}

	return in.ActualLoad / in.Capacity, nil
}
