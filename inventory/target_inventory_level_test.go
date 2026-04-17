package inventory

import "testing"

func TestTargetInventoryLevel(t *testing.T) {
	tests := []struct {
		name    string
		input   TargetInventoryLevelInput
		want    float64
		wantErr error
	}{
		{
			name: "valid inputs",
			input: TargetInventoryLevelInput{
				ExpectedDemandDuringLeadTime: 500,
				SafetyStockUnits:             50,
			},
			want: 550,
		},
		{
			name: "zero expected demand",
			input: TargetInventoryLevelInput{
				ExpectedDemandDuringLeadTime: 0,
				SafetyStockUnits:             50,
			},
			want: 50,
		},
		{
			name: "zero safety stock",
			input: TargetInventoryLevelInput{
				ExpectedDemandDuringLeadTime: 500,
				SafetyStockUnits:             0,
			},
			want: 500,
		},
		{
			name: "both zero",
			input: TargetInventoryLevelInput{
				ExpectedDemandDuringLeadTime: 0,
				SafetyStockUnits:             0,
			},
			want: 0,
		},
		{
			name: "negative expected demand",
			input: TargetInventoryLevelInput{
				ExpectedDemandDuringLeadTime: -1,
				SafetyStockUnits:             50,
			},
			wantErr: ErrNegativeExpectedDemand,
		},
		{
			name: "negative safety stock",
			input: TargetInventoryLevelInput{
				ExpectedDemandDuringLeadTime: 500,
				SafetyStockUnits:             -1,
			},
			wantErr: ErrNegativeSafetyStock,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TargetInventoryLevel(tt.input)

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

			if got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}
