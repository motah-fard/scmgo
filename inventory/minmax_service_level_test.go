package inventory

import (
	"math"
	"testing"
)

func TestMinMaxLevelsWithServiceLevel(t *testing.T) {
	tests := []struct {
		name    string
		input   MinMaxLevelsWithServiceLevelInput
		want    MinMaxResult
		wantErr error
	}{
		{
			name: "valid 95 percent service level",
			input: MinMaxLevelsWithServiceLevelInput{
				AvgDailyDemand:    50,
				LeadTimeDays:      4,
				StdDevDailyDemand: 10,
				ServiceLevel:      0.95,
				OrderQuantity:     200,
			},
			want: MinMaxResult{
				Min: 232.897072538,
				Max: 432.897072538,
			},
		},
		{
			name: "valid 90 percent service level",
			input: MinMaxLevelsWithServiceLevelInput{
				AvgDailyDemand:    20,
				LeadTimeDays:      9,
				StdDevDailyDemand: 8,
				ServiceLevel:      0.90,
				OrderQuantity:     100,
			},
			want: MinMaxResult{
				Min: 210.757237572,
				Max: 310.757237572,
			},
		},
		{
			name: "zero order quantity",
			input: MinMaxLevelsWithServiceLevelInput{
				AvgDailyDemand:    40,
				LeadTimeDays:      3,
				StdDevDailyDemand: 0,
				ServiceLevel:      0.95,
				OrderQuantity:     0,
			},
			want: MinMaxResult{
				Min: 120,
				Max: 120,
			},
		},
		{
			name: "negative order quantity",
			input: MinMaxLevelsWithServiceLevelInput{
				AvgDailyDemand:    50,
				LeadTimeDays:      4,
				StdDevDailyDemand: 10,
				ServiceLevel:      0.95,
				OrderQuantity:     -1,
			},
			wantErr: ErrNegativeOrderQuantity,
		},
		{
			name: "negative average demand",
			input: MinMaxLevelsWithServiceLevelInput{
				AvgDailyDemand:    -1,
				LeadTimeDays:      4,
				StdDevDailyDemand: 10,
				ServiceLevel:      0.95,
				OrderQuantity:     100,
			},
			wantErr: ErrNegativeDemand,
		},
		{
			name: "negative lead time",
			input: MinMaxLevelsWithServiceLevelInput{
				AvgDailyDemand:    50,
				LeadTimeDays:      -1,
				StdDevDailyDemand: 10,
				ServiceLevel:      0.95,
				OrderQuantity:     100,
			},
			wantErr: ErrNegativeLeadTime,
		},
		{
			name: "negative standard deviation",
			input: MinMaxLevelsWithServiceLevelInput{
				AvgDailyDemand:    50,
				LeadTimeDays:      4,
				StdDevDailyDemand: -2,
				ServiceLevel:      0.95,
				OrderQuantity:     100,
			},
			wantErr: ErrNegativeStandardDeviation,
		},
		{
			name: "invalid service level zero",
			input: MinMaxLevelsWithServiceLevelInput{
				AvgDailyDemand:    50,
				LeadTimeDays:      4,
				StdDevDailyDemand: 10,
				ServiceLevel:      0,
				OrderQuantity:     100,
			},
			wantErr: ErrInvalidServiceLevel,
		},
	}

	const tolerance = 1e-4

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MinMaxLevelsWithServiceLevel(tt.input)

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

			if math.Abs(got.Min-tt.want.Min) > tolerance {
				t.Fatalf("got min %.6f, want %.6f", got.Min, tt.want.Min)
			}
			if math.Abs(got.Max-tt.want.Max) > tolerance {
				t.Fatalf("got max %.6f, want %.6f", got.Max, tt.want.Max)
			}
		})
	}
}
