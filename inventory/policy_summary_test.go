package inventory

import (
	"errors"
	"math"
	"testing"
)

func TestBuildPolicySummary(t *testing.T) {
	tests := []struct {
		name    string
		input   PolicySummaryInput
		want    PolicySummary
		wantErr bool
		errs    []error
	}{
		{
			name: "valid full case",
			input: PolicySummaryInput{
				DailyDemand:      100,
				LeadTimeDays:     5,
				ReviewPeriodDays: 7,
				SafetyStockUnits: 50,
			},
			want: PolicySummary{
				ExpectedDemandDuringLeadTime: 500,
				SafetyStockUnits:             50,
				ReorderPoint:                 550,
				TargetInventoryLevel:         1250,
				MinLevel:                     550,
				MaxLevel:                     1250,
			},
		},
		{
			name: "zero safety stock",
			input: PolicySummaryInput{
				DailyDemand:      20,
				LeadTimeDays:     3,
				ReviewPeriodDays: 2,
				SafetyStockUnits: 0,
			},
			want: PolicySummary{
				ExpectedDemandDuringLeadTime: 60,
				SafetyStockUnits:             0,
				ReorderPoint:                 60,
				TargetInventoryLevel:         100,
				MinLevel:                     60,
				MaxLevel:                     100,
			},
		},
		{
			name: "zero review period",
			input: PolicySummaryInput{
				DailyDemand:      10,
				LeadTimeDays:     4,
				ReviewPeriodDays: 0,
				SafetyStockUnits: 5,
			},
			want: PolicySummary{
				ExpectedDemandDuringLeadTime: 40,
				SafetyStockUnits:             5,
				ReorderPoint:                 45,
				TargetInventoryLevel:         45,
				MinLevel:                     45,
				MaxLevel:                     45,
			},
		},
		{
			name: "zero lead time",
			input: PolicySummaryInput{
				DailyDemand:      10,
				LeadTimeDays:     0,
				ReviewPeriodDays: 6,
				SafetyStockUnits: 5,
			},
			want: PolicySummary{
				ExpectedDemandDuringLeadTime: 0,
				SafetyStockUnits:             5,
				ReorderPoint:                 5,
				TargetInventoryLevel:         65,
				MinLevel:                     5,
				MaxLevel:                     65,
			},
		},
		{
			name: "negative demand",
			input: PolicySummaryInput{
				DailyDemand:      -1,
				LeadTimeDays:     5,
				ReviewPeriodDays: 7,
				SafetyStockUnits: 10,
			},
			wantErr: true,
			errs:    []error{ErrInvalidPolicySummaryInput, ErrNegativeDemand},
		},
		{
			name: "negative lead time",
			input: PolicySummaryInput{
				DailyDemand:      10,
				LeadTimeDays:     -5,
				ReviewPeriodDays: 7,
				SafetyStockUnits: 10,
			},
			wantErr: true,
			errs:    []error{ErrInvalidPolicySummaryInput, ErrNegativeLeadTime},
		},
		{
			name: "negative review period",
			input: PolicySummaryInput{
				DailyDemand:      10,
				LeadTimeDays:     5,
				ReviewPeriodDays: -7,
				SafetyStockUnits: 10,
			},
			wantErr: true,
			errs:    []error{ErrInvalidPolicySummaryInput, ErrNegativeReviewPeriod},
		},
		{
			name: "negative safety stock",
			input: PolicySummaryInput{
				DailyDemand:      10,
				LeadTimeDays:     5,
				ReviewPeriodDays: 7,
				SafetyStockUnits: -10,
			},
			wantErr: true,
			errs:    []error{ErrInvalidPolicySummaryInput, ErrNegativeSafetyStock},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BuildPolicySummary(tt.input)

			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				for _, targetErr := range tt.errs {
					if !errors.Is(err, targetErr) {
						t.Fatalf("expected error to match %v, got %v", targetErr, err)
					}
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			assertClose(t, got.ExpectedDemandDuringLeadTime, tt.want.ExpectedDemandDuringLeadTime)
			assertClose(t, got.SafetyStockUnits, tt.want.SafetyStockUnits)
			assertClose(t, got.ReorderPoint, tt.want.ReorderPoint)
			assertClose(t, got.TargetInventoryLevel, tt.want.TargetInventoryLevel)
			assertClose(t, got.MinLevel, tt.want.MinLevel)
			assertClose(t, got.MaxLevel, tt.want.MaxLevel)
		})
	}
}

func TestBuildPolicySummaryConsistency(t *testing.T) {
	input := PolicySummaryInput{
		DailyDemand:      100,
		LeadTimeDays:     5,
		ReviewPeriodDays: 7,
		SafetyStockUnits: 50,
	}

	got, err := BuildPolicySummary(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expectedLeadTimeDemand, err := DemandDuringLeadTime(DemandDuringLeadTimeInput{
		AvgDailyDemand: input.DailyDemand,
		LeadTimeDays:   input.LeadTimeDays,
	})
	if err != nil {
		t.Fatalf("unexpected error from DemandDuringLeadTime: %v", err)
	}

	expectedReviewPeriodDemand, err := DemandDuringLeadTime(DemandDuringLeadTimeInput{
		AvgDailyDemand: input.DailyDemand,
		LeadTimeDays:   input.ReviewPeriodDays,
	})
	if err != nil {
		t.Fatalf("unexpected error from DemandDuringLeadTime for review period: %v", err)
	}

	reorderPoint, err := ReorderPoint(ReorderPointInput{
		AvgDailyDemand:   input.DailyDemand,
		LeadTimeDays:     input.LeadTimeDays,
		SafetyStockUnits: input.SafetyStockUnits,
	})
	if err != nil {
		t.Fatalf("unexpected error from ReorderPoint: %v", err)
	}

	targetLevel, err := TargetInventoryLevel(TargetInventoryLevelInput{
		ExpectedDemandDuringLeadTime: expectedLeadTimeDemand + expectedReviewPeriodDemand,
		SafetyStockUnits:             input.SafetyStockUnits,
	})
	if err != nil {
		t.Fatalf("unexpected error from TargetInventoryLevel: %v", err)
	}

	minMax, err := MinMaxLevels(MinMaxInput{
		ReorderPoint:  reorderPoint,
		OrderQuantity: targetLevel - reorderPoint,
	})
	if err != nil {
		t.Fatalf("unexpected error from MinMaxLevels: %v", err)
	}

	assertClose(t, got.ExpectedDemandDuringLeadTime, expectedLeadTimeDemand)
	assertClose(t, got.SafetyStockUnits, input.SafetyStockUnits)
	assertClose(t, got.ReorderPoint, reorderPoint)
	assertClose(t, got.TargetInventoryLevel, targetLevel)
	assertClose(t, got.MinLevel, minMax.Min)
	assertClose(t, got.MaxLevel, minMax.Max)
}

func assertClose(t *testing.T, got, want float64) {
	t.Helper()
	const tol = 1e-9
	if math.Abs(got-want) > tol {
		t.Fatalf("got %v, want %v", got, want)
	}
}
func TestBuildPolicySummaryWithServiceLevel(t *testing.T) {
	input := PolicySummaryServiceLevelInput{
		DailyDemand:        100,
		LeadTimeDays:       5,
		ReviewPeriodDays:   7,
		DemandStdDevPerDay: 20,
		ServiceLevel:       0.95,
	}

	got, err := BuildPolicySummaryWithServiceLevel(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expectedDLT, err := DemandDuringLeadTime(DemandDuringLeadTimeInput{
		AvgDailyDemand: input.DailyDemand,
		LeadTimeDays:   input.LeadTimeDays,
	})
	if err != nil {
		t.Fatalf("unexpected error computing expected lead-time demand: %v", err)
	}

	reorderPoint, err := ReorderPointWithServiceLevel(ReorderPointWithServiceLevelInput{
		AvgDailyDemand:    input.DailyDemand,
		LeadTimeDays:      input.LeadTimeDays,
		StdDevDailyDemand: input.DemandStdDevPerDay,
		ServiceLevel:      input.ServiceLevel,
	})
	if err != nil {
		t.Fatalf("unexpected error computing reorder point: %v", err)
	}

	if got.ExpectedDemandDuringLeadTime != expectedDLT {
		t.Fatalf("expected lead-time demand = %v, got %v", expectedDLT, got.ExpectedDemandDuringLeadTime)
	}
	if got.ReorderPoint != reorderPoint {
		t.Fatalf("expected reorder point = %v, got %v", reorderPoint, got.ReorderPoint)
	}
	if got.SafetyStockUnits != reorderPoint-expectedDLT {
		t.Fatalf("expected safety stock = %v, got %v", reorderPoint-expectedDLT, got.SafetyStockUnits)
	}
	if got.MinLevel != got.ReorderPoint {
		t.Fatalf("expected min level = reorder point, got min=%v reorder=%v", got.MinLevel, got.ReorderPoint)
	}
	if got.MaxLevel != got.TargetInventoryLevel {
		t.Fatalf("expected max level = target inventory level, got max=%v target=%v", got.MaxLevel, got.TargetInventoryLevel)
	}
}
func TestBuildPolicySummaryWithServiceLevelInvalidServiceLevel(t *testing.T) {
	_, err := BuildPolicySummaryWithServiceLevel(PolicySummaryServiceLevelInput{
		DailyDemand:        100,
		LeadTimeDays:       5,
		ReviewPeriodDays:   7,
		DemandStdDevPerDay: 20,
		ServiceLevel:       1,
	})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !errors.Is(err, ErrInvalidPolicySummaryInput) {
		t.Fatalf("expected ErrInvalidPolicySummaryInput, got %v", err)
	}
	if !errors.Is(err, ErrInvalidServiceLevel) {
		t.Fatalf("expected ErrInvalidServiceLevel, got %v", err)
	}
}
