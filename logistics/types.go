package logistics

// DimensionalWeightInput contains the inputs required to calculate
// dimensional weight.
type DimensionalWeightInput struct {
	Length float64
	Width  float64
	Height float64
	// DimFactor is the carrier's dimensional weight divisor (e.g. 139 or
	// 166 for inches/lb, 5000 for cm/kg).
	DimFactor float64
}

// FreightAllocationItem is one shipment item to allocate a shared freight
// cost across, by Weight (or any other non-negative allocation basis).
type FreightAllocationItem struct {
	ID     string
	Weight float64
}

// AllocateFreightCostInput contains the inputs required to allocate a
// shared freight cost across items.
type AllocateFreightCostInput struct {
	TotalFreightCost float64
	Items            []FreightAllocationItem
}

// FreightAllocationResult is one item's share of an allocated freight
// cost.
type FreightAllocationResult struct {
	ID            string
	Weight        float64
	AllocatedCost float64
}

// VehicleUtilizationInput contains the inputs required to calculate
// vehicle utilization.
type VehicleUtilizationInput struct {
	ActualLoad float64
	Capacity   float64
}

// LocationDemand is one location's coordinates and demand weight, for
// center-of-gravity facility location.
type LocationDemand struct {
	X      float64
	Y      float64
	Demand float64
}

// CenterOfGravityInput contains the inputs required to calculate a
// demand-weighted centroid.
type CenterOfGravityInput struct {
	Locations []LocationDemand
}

// CenterOfGravityResult contains the coordinates of the demand-weighted
// centroid.
type CenterOfGravityResult struct {
	X float64
	Y float64
}
