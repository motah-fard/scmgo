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

// PolicySummary contains the main outputs of an inventory policy calculation.
type PolicySummary struct {
	ExpectedDemandDuringLeadTime float64
	SafetyStockUnits             float64
	ReorderPoint                 float64
	TargetInventoryLevel         float64
	MinLevel                     float64
	MaxLevel                     float64
}

// PolicySummaryInput contains deterministic inputs for a policy summary calculation.
type PolicySummaryInput struct {
	DailyDemand      float64
	LeadTimeDays     float64
	ReviewPeriodDays float64
	SafetyStockUnits float64
}

// PolicySummaryServiceLevelInput contains service-level-driven inputs
// for a policy summary calculation.
type PolicySummaryServiceLevelInput struct {
	DailyDemand        float64
	LeadTimeDays       float64
	ReviewPeriodDays   float64
	DemandStdDevPerDay float64
	ServiceLevel       float64
}

// PolicySummaryBatchResult pairs the outcome of one item in a policy
// summary batch with its position in the input slice, so a failure on one
// item (e.g. one bad SKU row) doesn't prevent inspecting the rest.
type PolicySummaryBatchResult struct {
	// Index is the position of this result in the input slice.
	Index int
	// Summary is the computed policy summary. It is the zero value if Err
	// is non-nil.
	Summary PolicySummary
	// Err is the error returned for this item, or nil on success.
	Err error
}

// SafetyStockWithVariableLeadTimeInput contains the inputs for calculating
// safety stock when both demand and lead time vary.
type SafetyStockWithVariableLeadTimeInput struct {
	AvgDailyDemand     float64
	StdDevDailyDemand  float64
	AvgLeadTimeDays    float64
	StdDevLeadTimeDays float64
	ServiceLevel       float64
}

// ReorderPointWithVariableLeadTimeInput contains the inputs for calculating
// reorder point when both demand and lead time vary.
type ReorderPointWithVariableLeadTimeInput struct {
	AvgDailyDemand     float64
	StdDevDailyDemand  float64
	AvgLeadTimeDays    float64
	StdDevLeadTimeDays float64
	ServiceLevel       float64
}

// ExpectedFillRateInput contains the inputs for calculating the item fill
// rate (P2) achieved by a given safety stock and order quantity.
type ExpectedFillRateInput struct {
	SafetyStockUnits           float64
	StdDevDemandDuringLeadTime float64
	OrderQuantity              float64
}

// FillRateSafetyStockInput contains the inputs for calculating the safety
// stock required to achieve a target item fill rate (P2).
type FillRateSafetyStockInput struct {
	TargetFillRate             float64
	StdDevDemandDuringLeadTime float64
	OrderQuantity              float64
}

// EPQInput contains the inputs required to calculate the economic
// production quantity.
type EPQInput struct {
	AnnualDemand         float64
	SetupCost            float64
	HoldingCostPerUnit   float64
	AnnualProductionRate float64
}

// QuantityDiscountTier is one price break in a quantity discount schedule:
// ordering at least MinQuantity units qualifies for UnitPrice.
type QuantityDiscountTier struct {
	MinQuantity float64
	UnitPrice   float64
}

// EOQWithQuantityDiscountsInput contains the inputs required to calculate
// the cost-minimizing order quantity across a quantity discount schedule.
type EOQWithQuantityDiscountsInput struct {
	AnnualDemand float64
	OrderingCost float64
	// HoldingCostRate is the annual holding cost as a fraction of unit
	// price (e.g. 0.2 = 20% of unit price per year).
	HoldingCostRate float64
	// Tiers must be sorted by strictly ascending MinQuantity, the first
	// tier's MinQuantity must be greater than zero, and UnitPrice must be
	// positive and non-increasing across tiers.
	Tiers []QuantityDiscountTier
}

// EOQWithQuantityDiscountsResult contains the cost-minimizing order
// quantity found across a quantity discount schedule.
type EOQWithQuantityDiscountsResult struct {
	OrderQuantity   float64
	UnitPrice       float64
	TotalAnnualCost float64
}

// NewsvendorInput contains the inputs required to calculate the
// newsvendor (single-period) optimal order quantity.
type NewsvendorInput struct {
	MeanDemand          float64
	StdDevDemand        float64
	UnderageCostPerUnit float64
	OverageCostPerUnit  float64
}

// NewsvendorResult contains the outputs of a newsvendor calculation.
type NewsvendorResult struct {
	OrderQuantity float64
	CriticalRatio float64
}

// TurnoverInput contains the inputs required to calculate the
// inventory turnover ratio.
type TurnoverInput struct {
	COGS                  float64
	AverageInventoryValue float64
}

// DaysOfInventoryOnHandInput contains the inputs required to calculate
// days of inventory on hand.
type DaysOfInventoryOnHandInput struct {
	AverageInventoryValue float64
	COGS                  float64
	DaysInPeriod          float64
}

// GMROIInput contains the inputs required to calculate gross margin
// return on inventory investment.
type GMROIInput struct {
	GrossMargin          float64
	AverageInventoryCost float64
}

// EOIInput contains the inputs required to calculate the economic order
// interval.
type EOIInput struct {
	AnnualDemand       float64
	OrderingCost       float64
	HoldingCostPerUnit float64
	// DaysPerYear is the number of days in the period AnnualDemand covers
	// (e.g. 365).
	DaysPerYear float64
}

// SafetyTimeInput contains the inputs required to calculate safety time.
type SafetyTimeInput struct {
	SafetyStockUnits float64
	AvgDailyDemand   float64
}
