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

// ReorderPointWithServiceLevelInput contains the inputs for calculating
// reorder point using average daily demand, lead time, demand variability,
// and a target cycle service level.
type ReorderPointWithServiceLevelInput struct {
	AvgDailyDemand    float64
	LeadTimeDays      float64
	StdDevDailyDemand float64
	ServiceLevel      float64
}

// SafetyStockWithServiceLevelInput contains the inputs for calculating
// safety stock using demand variability, lead time, and a target cycle
// service level.
type SafetyStockWithServiceLevelInput struct {
	StdDevDailyDemand float64
	LeadTimeDays      float64
	ServiceLevel      float64
}

// DemandDuringLeadTimeInput contains the inputs for calculating
// expected demand during lead time.
type DemandDuringLeadTimeInput struct {
	AvgDailyDemand float64
	LeadTimeDays   float64
}

// StdDevDemandDuringLeadTimeInput contains the inputs for calculating
// the standard deviation of demand during lead time.
type StdDevDemandDuringLeadTimeInput struct {
	StdDevDailyDemand float64
	LeadTimeDays      float64
}

// TargetInventoryLevelInput contains the inputs for calculating
// target inventory level from expected demand during lead time and safety stock.
type TargetInventoryLevelInput struct {
	ExpectedDemandDuringLeadTime float64
	SafetyStockUnits             float64
}

// TargetInventoryLevelWithServiceLevelInput contains the inputs for calculating
// target inventory level using demand, lead time, demand variability,
// and a target cycle service level.
type TargetInventoryLevelWithServiceLevelInput struct {
	AvgDailyDemand    float64
	LeadTimeDays      float64
	StdDevDailyDemand float64
	ServiceLevel      float64
}

// MinMaxLevelsWithServiceLevelInput contains the inputs for calculating
// min/max inventory levels using a service-level-based reorder point
// and a fixed order quantity.
type MinMaxLevelsWithServiceLevelInput struct {
	AvgDailyDemand    float64
	LeadTimeDays      float64
	StdDevDailyDemand float64
	ServiceLevel      float64
	OrderQuantity     float64
}
