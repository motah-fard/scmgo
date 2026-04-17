package inventory

// ReorderPointInput contains the inputs required to calculate a reorder point.
type ReorderPointInput struct {
	AvgDailyDemand   float64
	LeadTimeDays     float64
	SafetyStockUnits float64
}

// SafetyStockInput contains the inputs required to calculate basic safety stock.
type SafetyStockInput struct {
	MaxDailyDemand  float64
	MaxLeadTimeDays float64
	AvgDailyDemand  float64
	AvgLeadTimeDays float64
}

// EOQInput contains the inputs required to calculate economic order quantity.
type EOQInput struct {
	AnnualDemand       float64
	OrderingCost       float64
	HoldingCostPerUnit float64
}

// MinMaxInput contains the inputs required to calculate min/max inventory levels.
type MinMaxInput struct {
	ReorderPoint  float64
	OrderQuantity float64
}

// MinMaxResult contains the calculated minimum and maximum inventory levels.
type MinMaxResult struct {
	Min float64
	Max float64
}
type ReorderPointWithServiceLevelInput struct {
	AvgDailyDemand    float64
	LeadTimeDays      float64
	StdDevDailyDemand float64
	ServiceLevel      float64
}
type SafetyStockWithServiceLevelInput struct {
	StdDevDailyDemand float64
	LeadTimeDays      float64
	ServiceLevel      float64
}
