package production

// OEEInput contains the inputs required to calculate overall equipment
// effectiveness.
type OEEInput struct {
	// PlannedProductionTime is the scheduled production time (e.g.
	// minutes), excluding planned downtime such as breaks.
	PlannedProductionTime float64
	// RunTime is the actual time equipment was running, at most
	// PlannedProductionTime.
	RunTime float64
	// IdealCycleTime is the theoretical minimum time to produce one unit.
	IdealCycleTime float64
	// TotalCount is the total number of units produced (good and bad).
	TotalCount float64
	// GoodCount is the number of units meeting quality standards, at most
	// TotalCount.
	GoodCount float64
}

// OEEResult contains the outputs of an OEE calculation.
type OEEResult struct {
	// Availability is RunTime / PlannedProductionTime.
	Availability float64
	// Performance is (IdealCycleTime * TotalCount) / RunTime.
	Performance float64
	// Quality is GoodCount / TotalCount.
	Quality float64
	// OEE is Availability * Performance * Quality.
	OEE float64
}

// WIPFromLittlesLawInput contains the inputs required to calculate work in
// progress from Little's Law.
type WIPFromLittlesLawInput struct {
	Throughput float64
	CycleTime  float64
}

// CycleTimeFromLittlesLawInput contains the inputs required to calculate
// cycle time from Little's Law.
type CycleTimeFromLittlesLawInput struct {
	WIP        float64
	Throughput float64
}

// ThroughputFromLittlesLawInput contains the inputs required to calculate
// throughput from Little's Law.
type ThroughputFromLittlesLawInput struct {
	WIP       float64
	CycleTime float64
}

// TaktTimeInput contains the inputs required to calculate takt time.
type TaktTimeInput struct {
	AvailableProductionTime float64
	CustomerDemand          float64
}
