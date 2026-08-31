package inventory

import (
	"math"
	"testing"
)

func TestSafetyStockWithVariableLeadTime(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		input     SafetyStockWithVariableLeadTimeInput
		want      float64
		tolerance float64
		wantErr   error
	}{
		{
			name: "valid input",
			input: SafetyStockWithVariableLeadTimeInput{
				AvgDailyDemand:     100,
				StdDevDailyDemand:  10,
				AvgLeadTimeDays:    5,
				StdDevLeadTimeDays: 1,
				ServiceLevel:       0.95,
			},
			want:      168.54734132581814,
			tolerance: 1e-6,
		},
		{
			name: "zero lead-time variability reduces to standard formula",
			input: SafetyStockWithVariableLeadTimeInput{
				AvgDailyDemand:     100,
				StdDevDailyDemand:  10,
				AvgLeadTimeDays:    5,
				StdDevLeadTimeDays: 0,
				ServiceLevel:       0.95,
			},
			want:      36.78004522900573,
			tolerance: 1e-9,
		},
		{
			name: "negative avg daily demand",
			input: SafetyStockWithVariableLeadTimeInput{
				AvgDailyDemand:     -1,
				StdDevDailyDemand:  10,
				AvgLeadTimeDays:    5,
				StdDevLeadTimeDays: 1,
				ServiceLevel:       0.95,
			},
			wantErr: ErrNegativeDemand,
		},
		{
			name: "negative std dev daily demand",
			input: SafetyStockWithVariableLeadTimeInput{
				AvgDailyDemand:     100,
				StdDevDailyDemand:  -1,
				AvgLeadTimeDays:    5,
				StdDevLeadTimeDays: 1,
				ServiceLevel:       0.95,
			},
			wantErr: ErrNegativeStandardDeviation,
		},
		{
			name: "negative avg lead time",
			input: SafetyStockWithVariableLeadTimeInput{
				AvgDailyDemand:     100,
				StdDevDailyDemand:  10,
				AvgLeadTimeDays:    -1,
				StdDevLeadTimeDays: 1,
				ServiceLevel:       0.95,
			},
			wantErr: ErrNegativeLeadTime,
		},
		{
			name: "negative std dev lead time",
			input: SafetyStockWithVariableLeadTimeInput{
				AvgDailyDemand:     100,
				StdDevDailyDemand:  10,
				AvgLeadTimeDays:    5,
				StdDevLeadTimeDays: -1,
				ServiceLevel:       0.95,
			},
			wantErr: ErrNegativeStandardDeviation,
		},
		{
			name: "invalid service level",
			input: SafetyStockWithVariableLeadTimeInput{
				AvgDailyDemand:     100,
				StdDevDailyDemand:  10,
				AvgLeadTimeDays:    5,
				StdDevLeadTimeDays: 1,
				ServiceLevel:       1,
			},
			wantErr: ErrInvalidServiceLevel,
		},
		{
			name: "NaN std dev lead time",
			input: SafetyStockWithVariableLeadTimeInput{
				AvgDailyDemand:     100,
				StdDevDailyDemand:  10,
				AvgLeadTimeDays:    5,
				StdDevLeadTimeDays: math.NaN(),
				ServiceLevel:       0.95,
			},
			wantErr: ErrNonFiniteInput,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := SafetyStockWithVariableLeadTime(tt.input)
			if err != tt.wantErr {
				t.Fatalf("SafetyStockWithVariableLeadTime() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}

			if !almostEqual(got, tt.want, tt.tolerance) {
				t.Fatalf("SafetyStockWithVariableLeadTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
