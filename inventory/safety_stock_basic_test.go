package inventory

import "testing"

func TestSafetyStockBasic(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   SafetyStockInput
		want    float64
		wantErr error
	}{
		{
			name: "valid input",
			input: SafetyStockInput{
				MaxDailyDemand:  120,
				MaxLeadTimeDays: 7,
				AvgDailyDemand:  100,
				AvgLeadTimeDays: 5,
			},
			want: 340, // (120*7) - (100*5) = 840 - 500 = 340
		},
		{
			name: "zero values",
			input: SafetyStockInput{
				MaxDailyDemand:  0,
				MaxLeadTimeDays: 0,
				AvgDailyDemand:  0,
				AvgLeadTimeDays: 0,
			},
			want: 0,
		},
		{
			name: "negative max daily demand",
			input: SafetyStockInput{
				MaxDailyDemand:  -1,
				MaxLeadTimeDays: 7,
				AvgDailyDemand:  100,
				AvgLeadTimeDays: 5,
			},
			wantErr: ErrNegativeDemand,
		},
		{
			name: "negative max lead time",
			input: SafetyStockInput{
				MaxDailyDemand:  120,
				MaxLeadTimeDays: -1,
				AvgDailyDemand:  100,
				AvgLeadTimeDays: 5,
			},
			wantErr: ErrNegativeLeadTime,
		},
		{
			name: "negative average daily demand",
			input: SafetyStockInput{
				MaxDailyDemand:  120,
				MaxLeadTimeDays: 7,
				AvgDailyDemand:  -1,
				AvgLeadTimeDays: 5,
			},
			wantErr: ErrNegativeDemand,
		},
		{
			name: "negative average lead time",
			input: SafetyStockInput{
				MaxDailyDemand:  120,
				MaxLeadTimeDays: 7,
				AvgDailyDemand:  100,
				AvgLeadTimeDays: -1,
			},
			wantErr: ErrNegativeLeadTime,
		},
		{
			name: "negative result is clamped to zero",
			input: SafetyStockInput{
				MaxDailyDemand:  80,
				MaxLeadTimeDays: 4,
				AvgDailyDemand:  100,
				AvgLeadTimeDays: 5,
			},
			want: 0, // (80*4) - (100*5) = -180 => clamp to 0
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := SafetyStockBasic(tt.input)
			if err != tt.wantErr {
				t.Fatalf("SafetyStockBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}

			if got != tt.want {
				t.Fatalf("SafetyStockBasic() = %v, want %v", got, tt.want)
			}
		})
	}
}
