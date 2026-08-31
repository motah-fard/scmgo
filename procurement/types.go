package procurement

// CostComponent is one labeled amount contributing to a total cost
// (e.g. "Freight", "Duty", "Insurance"). Amount may be negative to
// represent a credit, rebate, or offsetting value.
type CostComponent struct {
	Label  string
	Amount float64
}

// LandedCostInput contains the inputs required to calculate landed cost.
type LandedCostInput struct {
	Components []CostComponent
}

// LandedCostResult contains the outputs of a landed cost calculation.
type LandedCostResult struct {
	Total      float64
	Components []CostComponent
}

// PurchasePriceVarianceInput contains the inputs required to calculate
// purchase price variance.
type PurchasePriceVarianceInput struct {
	StandardCost float64
	ActualCost   float64
	Quantity     float64
}

// TotalCostOfOwnershipInput contains the inputs required to calculate
// total cost of ownership.
type TotalCostOfOwnershipInput struct {
	Components []CostComponent
}

// TotalCostOfOwnershipResult contains the outputs of a total cost of
// ownership calculation.
type TotalCostOfOwnershipResult struct {
	Total      float64
	Components []CostComponent
}
