package inventory

import "errors"

// wrapPolicySummaryErr joins err under ErrInvalidPolicySummaryInput so
// callers can match on either the general or the specific cause via
// errors.Is.
func wrapPolicySummaryErr(err error) error {
	return errors.Join(ErrInvalidPolicySummaryInput, err)
}

// validatePolicySummaryInput validates deterministic policy summary inputs.
func validatePolicySummaryInput(input PolicySummaryInput) error {
	if err := validateFinite(input.DailyDemand, input.LeadTimeDays, input.ReviewPeriodDays, input.SafetyStockUnits); err != nil {
		return wrapPolicySummaryErr(err)
	}
	if input.DailyDemand < 0 {
		return wrapPolicySummaryErr(ErrNegativeDemand)
	}
	if input.LeadTimeDays < 0 {
		return wrapPolicySummaryErr(ErrNegativeLeadTime)
	}
	if input.ReviewPeriodDays < 0 {
		return wrapPolicySummaryErr(ErrNegativeReviewPeriod)
	}
	if input.SafetyStockUnits < 0 {
		return wrapPolicySummaryErr(ErrNegativeSafetyStock)
	}
	return nil
}

// validatePolicySummaryServiceLevelInput validates service-level-based
// policy summary inputs.
func validatePolicySummaryServiceLevelInput(input PolicySummaryServiceLevelInput) error {
	if err := validateFinite(input.DailyDemand, input.LeadTimeDays, input.ReviewPeriodDays, input.DemandStdDevPerDay, input.ServiceLevel); err != nil {
		return wrapPolicySummaryErr(err)
	}
	if input.DailyDemand < 0 {
		return wrapPolicySummaryErr(ErrNegativeDemand)
	}
	if input.LeadTimeDays < 0 {
		return wrapPolicySummaryErr(ErrNegativeLeadTime)
	}
	if input.ReviewPeriodDays < 0 {
		return wrapPolicySummaryErr(ErrNegativeReviewPeriod)
	}
	if input.DemandStdDevPerDay < 0 {
		return wrapPolicySummaryErr(ErrNegativeStandardDeviation)
	}
	if input.ServiceLevel <= 0 || input.ServiceLevel >= 1 {
		return wrapPolicySummaryErr(ErrInvalidServiceLevel)
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
