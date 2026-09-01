package logistics

// CenterOfGravity calculates the demand-weighted centroid of a set of
// locations, a common first-pass estimate for facility location:
//
//	X = sum(Demand_i × X_i) / sum(Demand_i)
//	Y = sum(Demand_i × Y_i) / sum(Demand_i)
//
// Locations must contain at least one entry. Each location's Demand must
// be non-negative, and the demands must sum to more than zero. X and Y
// may be any finite value (e.g. negative coordinates in a local
// coordinate system).
//
// This computes a straight-line (Euclidean-plane) centroid; it does not
// account for road network distance or the earth's curvature.
func CenterOfGravity(in CenterOfGravityInput) (CenterOfGravityResult, error) {
	if len(in.Locations) == 0 {
		return CenterOfGravityResult{}, ErrEmptyLocations
	}

	var totalDemand, weightedX, weightedY float64
	for _, loc := range in.Locations {
		if err := validateFinite(loc.X, loc.Y, loc.Demand); err != nil {
			return CenterOfGravityResult{}, err
		}
		if loc.Demand < 0 {
			return CenterOfGravityResult{}, ErrNegativeDemand
		}
		totalDemand += loc.Demand
		weightedX += loc.Demand * loc.X
		weightedY += loc.Demand * loc.Y
	}
	if totalDemand <= 0 {
		return CenterOfGravityResult{}, ErrZeroTotalDemand
	}

	return CenterOfGravityResult{
		X: weightedX / totalDemand,
		Y: weightedY / totalDemand,
	}, nil
}
