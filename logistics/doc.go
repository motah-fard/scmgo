// Package logistics provides practical transportation and distribution
// calculations.
//
// It includes:
//
//   - DimensionalWeight and BillableWeight: parcel shipping weight
//     calculations
//   - AllocateFreightCost: proportional allocation of a shared freight
//     cost across shipment items
//   - VehicleUtilization: load vs. capacity ratio
//   - CenterOfGravity: demand-weighted centroid for facility location
//
// Important assumptions:
//
//   - DimensionalWeight takes the divisor (DimFactor) as a caller-supplied
//     input rather than a hardcoded carrier constant — carriers use
//     different divisors (commonly 139 or 166 in inches/lb, 5000 in
//     cm/kg) and these change over time, so hardcoding one would silently
//     go stale.
//   - VehicleUtilization does not clamp its result: a value above 1
//     indicates an overloaded vehicle, which is meaningful information,
//     not an error.
//   - AllocateFreightCost allocates proportionally to a caller-supplied
//     basis (weight, volume, or any other non-negative metric) — it does
//     not decide which basis is appropriate for a given cost.
//   - CenterOfGravity computes a straight-line (Euclidean-plane) weighted
//     centroid; it does not account for road network distance, geographic
//     projection, or the earth's curvature.
//   - This package does not implement vehicle routing (VRP) or network
//     optimization — those require a solver, not a closed-form formula.
package logistics
