package inventory

import (
	"math"
	"testing"
)

func TestSafetyStockWithServiceLevel(t *testing.T) {
	tests := []struct {
		name    string
		input   SafetyStockWithServiceLevelInput
		want    float64
		wantErr error
	}{
		{
			name: "valid 95 percent service level",
			input: SafetyStockWithServiceLevelInput{
				StdDevDailyDemand: 10,
				LeadTimeDays:      4,
				ServiceLevel:      0.95,
			},
			want: 32.897072538,
		},
		{
			name: "valid 90 percent service level",
			input: SafetyStockWithServiceLevelInput{
				StdDevDailyDemand: 8,
				LeadTimeDays:      9,
				ServiceLevel:      0.90,
			},
			want: 30.757237572,
		},
		{
			name: "zero standard deviation",
			input: SafetyStockWithServiceLevelInput{
				StdDevDailyDemand: 0,
				LeadTimeDays:      5,
				ServiceLevel:      0.95,
			},
			want: 0,
		},
		{
			name: "zero lead time",
			input: SafetyStockWithServiceLevelInput{
				StdDevDailyDemand: 12,
				LeadTimeDays:      0,
				ServiceLevel:      0.95,
			},
			want: 0,
		},
		{
			name: "negative standard deviation",
			input: SafetyStockWithServiceLevelInput{
				StdDevDailyDemand: -1,
				LeadTimeDays:      3,
				ServiceLevel:      0.95,
			},
			wantErr: ErrNegativeStandardDeviation,
		},
		{
			name: "negative lead time",
			input: SafetyStockWithServiceLevelInput{
				StdDevDailyDemand: 10,
				LeadTimeDays:      -2,
				ServiceLevel:      0.95,
			},
			wantErr: ErrNegativeLeadTime,
		},
		{
			name: "invalid zero service level",
			input: SafetyStockWithServiceLevelInput{
				StdDevDailyDemand: 10,
				LeadTimeDays:      3,
				ServiceLevel:      0,
			},
			wantErr: ErrInvalidServiceLevel,
		},
		{
			name: "invalid one service level",
			input: SafetyStockWithServiceLevelInput{
				StdDevDailyDemand: 10,
				LeadTimeDays:      3,
				ServiceLevel:      1,
			},
			wantErr: ErrInvalidServiceLevel,
		},
	}

	const tolerance = 1e-4

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SafetyStockWithServiceLevel(tt.input)

			if tt.wantErr != nil {
				if err == nil {
					t.Fatalf("expected error %v, got nil", tt.wantErr)
				}
				if err != tt.wantErr {
					t.Fatalf("expected error %v, got %v", tt.wantErr, err)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if math.Abs(got-tt.want) > tolerance {
				t.Fatalf("got %.6f, want %.6f", got, tt.want)
			}
		})
	}
}
