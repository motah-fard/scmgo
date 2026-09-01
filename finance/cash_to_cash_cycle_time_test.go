package finance

import (
	"math"
	"testing"
)

func TestCashToCashCycleTime(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   CashToCashCycleTimeInput
		want    float64
		wantErr error
	}{
		{
			name:  "valid input",
			input: CashToCashCycleTimeInput{DIO: 60.83, DSO: 30.416666666666664, DPO: 36.5},
			want:  54.74666666666667,
		},
		{
			name:  "negative cycle time is a valid favorable outcome",
			input: CashToCashCycleTimeInput{DIO: 10, DSO: 5, DPO: 40},
			want:  -25,
		},
		{
			name:    "negative DIO",
			input:   CashToCashCycleTimeInput{DIO: -1, DSO: 30, DPO: 36},
			wantErr: ErrNegativeDays,
		},
		{
			name:    "negative DPO",
			input:   CashToCashCycleTimeInput{DIO: 60, DSO: 30, DPO: -1},
			wantErr: ErrNegativeDays,
		},
		{
			name:    "NaN DSO",
			input:   CashToCashCycleTimeInput{DIO: 60, DSO: math.NaN(), DPO: 36},
			wantErr: ErrNonFiniteInput,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := CashToCashCycleTime(tt.input)
			if err != tt.wantErr {
				t.Fatalf("CashToCashCycleTime() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}

			if !almostEqual(got, tt.want, 1e-9) {
				t.Fatalf("CashToCashCycleTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
