package inventory

import (
	"math"
	"testing"
)

func almostEqual(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}

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

func TestEOQ(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		input     EOQInput
		want      float64
		tolerance float64
		wantErr   error
	}{
		{
			name: "valid input",
			input: EOQInput{
				AnnualDemand:       10000,
				OrderingCost:       50,
				HoldingCostPerUnit: 2,
			},
			want:      707.1067811865476,
			tolerance: 1e-9,
		},
		{
			name: "zero annual demand",
			input: EOQInput{
				AnnualDemand:       0,
				OrderingCost:       50,
				HoldingCostPerUnit: 2,
			},
			want:      0,
			tolerance: 1e-9,
		},
		{
			name: "zero ordering cost",
			input: EOQInput{
				AnnualDemand:       10000,
				OrderingCost:       0,
				HoldingCostPerUnit: 2,
			},
			want:      0,
			tolerance: 1e-9,
		},
		{
			name: "negative annual demand",
			input: EOQInput{
				AnnualDemand:       -1,
				OrderingCost:       50,
				HoldingCostPerUnit: 2,
			},
			wantErr: ErrNegativeDemand,
		},
		{
			name: "negative ordering cost",
			input: EOQInput{
				AnnualDemand:       10000,
				OrderingCost:       -1,
				HoldingCostPerUnit: 2,
			},
			wantErr: ErrNegativeOrderingCost,
		},
		{
			name: "zero holding cost",
			input: EOQInput{
				AnnualDemand:       10000,
				OrderingCost:       50,
				HoldingCostPerUnit: 0,
			},
			wantErr: ErrInvalidHoldingCost,
		},
		{
			name: "negative holding cost",
			input: EOQInput{
				AnnualDemand:       10000,
				OrderingCost:       50,
				HoldingCostPerUnit: -1,
			},
			wantErr: ErrInvalidHoldingCost,
		},
		{
			name: "very large values",
			input: EOQInput{
				AnnualDemand:       1_000_000_000,
				OrderingCost:       10_000,
				HoldingCostPerUnit: 25,
			},
			want:      894427.1909999159,
			tolerance: 1e-6,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := EOQ(tt.input)
			if err != tt.wantErr {
				t.Fatalf("EOQ() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}

			if !almostEqual(got, tt.want, tt.tolerance) {
				t.Fatalf("EOQ() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinMaxLevels(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   MinMaxInput
		want    MinMaxResult
		wantErr error
	}{
		{
			name: "valid input",
			input: MinMaxInput{
				ReorderPoint:  300,
				OrderQuantity: 200,
			},
			want: MinMaxResult{
				Min: 300,
				Max: 500,
			},
		},
		{
			name: "zero order quantity",
			input: MinMaxInput{
				ReorderPoint:  300,
				OrderQuantity: 0,
			},
			want: MinMaxResult{
				Min: 300,
				Max: 300,
			},
		},
		{
			name: "negative reorder point",
			input: MinMaxInput{
				ReorderPoint:  -1,
				OrderQuantity: 200,
			},
			wantErr: ErrNegativeReorderPoint,
		},
		{
			name: "negative order quantity",
			input: MinMaxInput{
				ReorderPoint:  300,
				OrderQuantity: -1,
			},
			wantErr: ErrNegativeOrderQuantity,
		},
		{
			name: "zero reorder point with positive order quantity",
			input: MinMaxInput{
				ReorderPoint:  0,
				OrderQuantity: 500,
			},
			want: MinMaxResult{
				Min: 0,
				Max: 500,
			},
		},
		{
			name: "very large values",
			input: MinMaxInput{
				ReorderPoint:  10_000_000,
				OrderQuantity: 5_000_000,
			},
			want: MinMaxResult{
				Min: 10_000_000,
				Max: 15_000_000,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := MinMaxLevels(tt.input)
			if err != tt.wantErr {
				t.Fatalf("MinMaxLevels() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}

			if got != tt.want {
				t.Fatalf("MinMaxLevels() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
