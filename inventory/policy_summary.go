package inventory

import "errors"

// BuildPolicySummary builds a deterministic inventory policy summary
// from average demand, lead time, review period, and fixed safety stock.
//
// It returns expected demand during lead time, safety stock, reorder point,
// target inventory level, and corresponding min/max levels.
func BuildPolicySummary(input PolicySummaryInput) (PolicySummary, error) {
	if err := validatePolicySummaryInput(input); err != nil {
		return PolicySummary{}, err
	}

	expectedLeadTimeDemand, err := DemandDuringLeadTime(DemandDuringLeadTimeInput{
		AvgDailyDemand: input.DailyDemand,
		LeadTimeDays:   input.LeadTimeDays,
	})
	if err != nil {
		return PolicySummary{}, errors.Join(ErrInvalidPolicySummaryInput, err)
	}

	expectedReviewPeriodDemand, err := DemandDuringLeadTime(DemandDuringLeadTimeInput{
		AvgDailyDemand: input.DailyDemand,
		LeadTimeDays:   input.ReviewPeriodDays,
	})
	if err != nil {
		return PolicySummary{}, errors.Join(ErrInvalidPolicySummaryInput, err)
	}

	reorderPoint, err := ReorderPoint(ReorderPointInput{
		AvgDailyDemand:   input.DailyDemand,
		LeadTimeDays:     input.LeadTimeDays,
		SafetyStockUnits: input.SafetyStockUnits,
	})
	if err != nil {
		return PolicySummary{}, errors.Join(ErrInvalidPolicySummaryInput, err)
	}

	targetLevel, err := TargetInventoryLevel(TargetInventoryLevelInput{
		ExpectedDemandDuringLeadTime: expectedLeadTimeDemand + expectedReviewPeriodDemand,
		SafetyStockUnits:             input.SafetyStockUnits,
	})
	if err != nil {
		return PolicySummary{}, errors.Join(ErrInvalidPolicySummaryInput, err)
	}

	minMax, err := MinMaxLevels(MinMaxInput{
		ReorderPoint:  reorderPoint,
		OrderQuantity: targetLevel - reorderPoint,
	})
	if err != nil {
		return PolicySummary{}, errors.Join(ErrInvalidPolicySummaryInput, err)
	}

	return newPolicySummary(
		expectedLeadTimeDemand,
		input.SafetyStockUnits,
		reorderPoint,
		targetLevel,
		minMax.Min,
		minMax.Max,
	), nil
}

// BuildPolicySummaryWithServiceLevel builds an inventory policy summary
// using service-level-based safety stock and reorder point logic.
//
// It returns expected demand during lead time, derived safety stock,
// reorder point, target inventory level, and corresponding min/max levels.
func BuildPolicySummaryWithServiceLevel(input PolicySummaryServiceLevelInput) (PolicySummary, error) {
	if err := validatePolicySummaryServiceLevelInput(input); err != nil {
		return PolicySummary{}, err
	}

	expectedLeadTimeDemand, err := DemandDuringLeadTime(DemandDuringLeadTimeInput{
		AvgDailyDemand: input.DailyDemand,
		LeadTimeDays:   input.LeadTimeDays,
	})
	if err != nil {
		return PolicySummary{}, errors.Join(ErrInvalidPolicySummaryInput, err)
	}

	expectedReviewPeriodDemand, err := DemandDuringLeadTime(DemandDuringLeadTimeInput{
		AvgDailyDemand: input.DailyDemand,
		LeadTimeDays:   input.ReviewPeriodDays,
	})
	if err != nil {
		return PolicySummary{}, errors.Join(ErrInvalidPolicySummaryInput, err)
	}

	reorderPoint, err := ReorderPointWithServiceLevel(ReorderPointWithServiceLevelInput{
		AvgDailyDemand:    input.DailyDemand,
		LeadTimeDays:      input.LeadTimeDays,
		StdDevDailyDemand: input.DemandStdDevPerDay,
		ServiceLevel:      input.ServiceLevel,
	})
	if err != nil {
		return PolicySummary{}, errors.Join(ErrInvalidPolicySummaryInput, err)
	}

	safetyStock := reorderPoint - expectedLeadTimeDemand

	targetLevel, err := TargetInventoryLevel(TargetInventoryLevelInput{
		ExpectedDemandDuringLeadTime: expectedLeadTimeDemand + expectedReviewPeriodDemand,
		SafetyStockUnits:             safetyStock,
	})
	if err != nil {
		return PolicySummary{}, errors.Join(ErrInvalidPolicySummaryInput, err)
	}

	minMax, err := MinMaxLevels(MinMaxInput{
		ReorderPoint:  reorderPoint,
		OrderQuantity: targetLevel - reorderPoint,
	})
	if err != nil {
		return PolicySummary{}, errors.Join(ErrInvalidPolicySummaryInput, err)
	}

	return newPolicySummary(
		expectedLeadTimeDemand,
		safetyStock,
		reorderPoint,
		targetLevel,
		minMax.Min,
		minMax.Max,
	), nil
}
