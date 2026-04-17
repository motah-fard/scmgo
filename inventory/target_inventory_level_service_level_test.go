package inventory

import (
	"math"
	"testing"
)

func TestTargetInventoryLevelWithServiceLevel(t *testing.T) {
	tests := []struct {
		name    string
		input   TargetInventoryLevelWithServiceLevelInput
		want    float64
		wantErr error
	}{
		{
			name: "valid 95 percent service level",
			input: TargetInventoryLevelWithServiceLevelInput{
				AvgDailyDemand:    50,
				LeadTimeDays:      4,
				StdDevDailyDemand: 10,
				ServiceLevel:      0.95,
			},
			want: 232.897072538,
		},
		{
			name: "valid 90 percent service level",
			input: TargetInventoryLevelWithServiceLevelInput{
				AvgDailyDemand:    20,
				LeadTimeDays:      9,
				StdDevDailyDemand: 8,
				ServiceLevel:      0.90,
			},
			want: 210.757237572,
		},
		{
			name: "zero average demand",
			input: TargetInventoryLevelWithServiceLevelInput{
				AvgDailyDemand:    0,
				LeadTimeDays:      5,
				StdDevDailyDemand: 10,
				ServiceLevel:      0.95,
			},
			want: 36.780044795,
		},
		{
			name: "zero lead time",
			input: TargetInventoryLevelWithServiceLevelInput{
				AvgDailyDemand:    50,
				LeadTimeDays:      0,
				StdDevDailyDemand: 10,
				ServiceLevel:      0.95,
			},
			want: 0,
		},
		{
			name: "zero standard deviation",
			input: TargetInventoryLevelWithServiceLevelInput{
				AvgDailyDemand:    40,
				LeadTimeDays:      3,
				StdDevDailyDemand: 0,
				ServiceLevel:      0.95,
			},
			want: 120,
		},
		{
			name: "negative average demand",
			input: TargetInventoryLevelWithServiceLevelInput{
				AvgDailyDemand:    -1,
				LeadTimeDays:      4,
				StdDevDailyDemand: 10,
				ServiceLevel:      0.95,
			},
			wantErr: ErrNegativeDemand,
		},
		{
			name: "negative lead time",
			input: TargetInventoryLevelWithServiceLevelInput{
				AvgDailyDemand:    50,
				LeadTimeDays:      -1,
				StdDevDailyDemand: 10,
				ServiceLevel:      0.95,
			},
			wantErr: ErrNegativeLeadTime,
		},
		{
			name: "negative standard deviation",
			input: TargetInventoryLevelWithServiceLevelInput{
				AvgDailyDemand:    50,
				LeadTimeDays:      4,
				StdDevDailyDemand: -2,
				ServiceLevel:      0.95,
			},
			wantErr: ErrNegativeStandardDeviation,
		},
		{
			name: "invalid service level zero",
			input: TargetInventoryLevelWithServiceLevelInput{
				AvgDailyDemand:    50,
				LeadTimeDays:      4,
				StdDevDailyDemand: 10,
				ServiceLevel:      0,
			},
			wantErr: ErrInvalidServiceLevel,
		},
		{
			name: "invalid service level one",
			input: TargetInventoryLevelWithServiceLevelInput{
				AvgDailyDemand:    50,
				LeadTimeDays:      4,
				StdDevDailyDemand: 10,
				ServiceLevel:      1,
			},
			wantErr: ErrInvalidServiceLevel,
		},
	}

	const tolerance = 1e-4

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TargetInventoryLevelWithServiceLevel(tt.input)

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
