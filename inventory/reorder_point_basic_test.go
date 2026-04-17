package inventory

import "testing"

func TestReorderPoint(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   ReorderPointInput
		want    float64
		wantErr error
	}{
		{
			name: "valid input",
			input: ReorderPointInput{
				AvgDailyDemand:   100,
				LeadTimeDays:     5,
				SafetyStockUnits: 50,
			},
			want: 550,
		},
		{
			name: "zero values",
			input: ReorderPointInput{
				AvgDailyDemand:   0,
				LeadTimeDays:     0,
				SafetyStockUnits: 0,
			},
			want: 0,
		},
		{
			name: "negative demand",
			input: ReorderPointInput{
				AvgDailyDemand:   -1,
				LeadTimeDays:     5,
				SafetyStockUnits: 10,
			},
			wantErr: ErrNegativeDemand,
		},
		{
			name: "negative lead time",
			input: ReorderPointInput{
				AvgDailyDemand:   10,
				LeadTimeDays:     -5,
				SafetyStockUnits: 10,
			},
			wantErr: ErrNegativeLeadTime,
		},
		{
			name: "negative safety stock",
			input: ReorderPointInput{
				AvgDailyDemand:   10,
				LeadTimeDays:     5,
				SafetyStockUnits: -1,
			},
			wantErr: ErrNegativeSafetyStock,
		},
		{
			name: "zero demand with nonzero safety stock",
			input: ReorderPointInput{
				AvgDailyDemand:   0,
				LeadTimeDays:     5,
				SafetyStockUnits: 25,
			},
			want: 25,
		},
		{
			name: "zero lead time with nonzero demand and safety stock",
			input: ReorderPointInput{
				AvgDailyDemand:   100,
				LeadTimeDays:     0,
				SafetyStockUnits: 40,
			},
			want: 40,
		},
		{
			name: "very large values",
			input: ReorderPointInput{
				AvgDailyDemand:   1_000_000,
				LeadTimeDays:     365,
				SafetyStockUnits: 10_000_000,
			},
			want: 375_000_000,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := ReorderPoint(tt.input)
			if err != tt.wantErr {
				t.Fatalf("ReorderPoint() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}

			if got != tt.want {
				t.Fatalf("ReorderPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}
