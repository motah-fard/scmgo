package inventory

import "errors"

// validatePolicySummaryInput validates deterministic policy summary inputs.
func validatePolicySummaryInput(input PolicySummaryInput) error {
	for _, v := range []float64{input.DailyDemand, input.LeadTimeDays, input.ReviewPeriodDays, input.SafetyStockUnits} {
		if err := validateFinite(v); err != nil {
			return errors.Join(ErrInvalidPolicySummaryInput, err)
		}
	}
	if input.DailyDemand < 0 {
		return errors.Join(ErrInvalidPolicySummaryInput, ErrNegativeDemand)
	}
	if input.LeadTimeDays < 0 {
		return errors.Join(ErrInvalidPolicySummaryInput, ErrNegativeLeadTime)
	}
	if input.ReviewPeriodDays < 0 {
		return errors.Join(ErrInvalidPolicySummaryInput, ErrNegativeReviewPeriod)
	}
	if input.SafetyStockUnits < 0 {
		return errors.Join(ErrInvalidPolicySummaryInput, ErrNegativeSafetyStock)
	}
	return nil
}

// validatePolicySummaryServiceLevelInput validates service-level-based
// policy summary inputs.
func validatePolicySummaryServiceLevelInput(input PolicySummaryServiceLevelInput) error {
	for _, v := range []float64{input.DailyDemand, input.LeadTimeDays, input.ReviewPeriodDays, input.DemandStdDevPerDay, input.ServiceLevel} {
		if err := validateFinite(v); err != nil {
			return errors.Join(ErrInvalidPolicySummaryInput, err)
		}
	}
	if input.DailyDemand < 0 {
		return errors.Join(ErrInvalidPolicySummaryInput, ErrNegativeDemand)
	}
	if input.LeadTimeDays < 0 {
		return errors.Join(ErrInvalidPolicySummaryInput, ErrNegativeLeadTime)
	}
	if input.ReviewPeriodDays < 0 {
		return errors.Join(ErrInvalidPolicySummaryInput, ErrNegativeReviewPeriod)
	}
	if input.DemandStdDevPerDay < 0 {
		return errors.Join(ErrInvalidPolicySummaryInput, ErrNegativeStandardDeviation)
	}
	if input.ServiceLevel <= 0 || input.ServiceLevel >= 1 {
		return errors.Join(ErrInvalidPolicySummaryInput, ErrInvalidServiceLevel)
	}
	return nil
}

// newPolicySummary assembles a PolicySummary from computed values.
func newPolicySummary(
	expectedDemandDuringLeadTime float64,
	safetyStockUnits float64,
	reorderPoint float64,
	targetInventoryLevel float64,
	minLevel float64,
	maxLevel float64,
) PolicySummary {
	return PolicySummary{
		ExpectedDemandDuringLeadTime: expectedDemandDuringLeadTime,
		SafetyStockUnits:             safetyStockUnits,
		ReorderPoint:                 reorderPoint,
		TargetInventoryLevel:         targetInventoryLevel,
		MinLevel:                     minLevel,
		MaxLevel:                     maxLevel,
	}
}
