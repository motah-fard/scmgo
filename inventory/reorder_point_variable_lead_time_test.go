package inventory

import (
	"math"
	"testing"
)

func TestReorderPointWithVariableLeadTime(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		input     ReorderPointWithVariableLeadTimeInput
		want      float64
		tolerance float64
		wantErr   error
	}{
		{
			name: "valid input",
			input: ReorderPointWithVariableLeadTimeInput{
				AvgDailyDemand:     100,
				StdDevDailyDemand:  10,
				AvgLeadTimeDays:    5,
				StdDevLeadTimeDays: 1,
				ServiceLevel:       0.95,
			},
			want:      668.5473413258181,
			tolerance: 1e-6,
		},
		{
			name: "negative avg daily demand",
			input: ReorderPointWithVariableLeadTimeInput{
				AvgDailyDemand:     -1,
				StdDevDailyDemand:  10,
				AvgLeadTimeDays:    5,
				StdDevLeadTimeDays: 1,
				ServiceLevel:       0.95,
			},
			wantErr: ErrNegativeDemand,
		},
		{
			name: "negative avg lead time",
			input: ReorderPointWithVariableLeadTimeInput{
				AvgDailyDemand:     100,
				StdDevDailyDemand:  10,
				AvgLeadTimeDays:    -1,
				StdDevLeadTimeDays: 1,
				ServiceLevel:       0.95,
			},
			wantErr: ErrNegativeLeadTime,
		},
		{
			name: "invalid service level propagated from safety stock",
			input: ReorderPointWithVariableLeadTimeInput{
				AvgDailyDemand:     100,
				StdDevDailyDemand:  10,
				AvgLeadTimeDays:    5,
				StdDevLeadTimeDays: 1,
				ServiceLevel:       0,
			},
			wantErr: ErrInvalidServiceLevel,
		},
		{
			name: "NaN service level",
			input: ReorderPointWithVariableLeadTimeInput{
				AvgDailyDemand:     100,
				StdDevDailyDemand:  10,
				AvgLeadTimeDays:    5,
				StdDevLeadTimeDays: 1,
				ServiceLevel:       math.NaN(),
			},
			wantErr: ErrNonFiniteInput,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := ReorderPointWithVariableLeadTime(tt.input)
			if err != tt.wantErr {
				t.Fatalf("ReorderPointWithVariableLeadTime() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}

			if !almostEqual(got, tt.want, tt.tolerance) {
				t.Fatalf("ReorderPointWithVariableLeadTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
