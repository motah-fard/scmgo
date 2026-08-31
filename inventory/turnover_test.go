package inventory

import (
	"math"
	"testing"
)

func TestTurnover(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		input     TurnoverInput
		want      float64
		tolerance float64
		wantErr   error
	}{
		{
			name:      "valid input",
			input:     TurnoverInput{COGS: 120000, AverageInventoryValue: 20000},
			want:      6,
			tolerance: 1e-9,
		},
		{
			name:      "zero COGS",
			input:     TurnoverInput{COGS: 0, AverageInventoryValue: 20000},
			want:      0,
			tolerance: 1e-9,
		},
		{
			name:    "negative COGS",
			input:   TurnoverInput{COGS: -1, AverageInventoryValue: 20000},
			wantErr: ErrNegativeCOGS,
		},
		{
			name:    "zero average inventory value",
			input:   TurnoverInput{COGS: 120000, AverageInventoryValue: 0},
			wantErr: ErrInvalidAverageInventoryValue,
		},
		{
			name:    "negative average inventory value",
			input:   TurnoverInput{COGS: 120000, AverageInventoryValue: -1},
			wantErr: ErrInvalidAverageInventoryValue,
		},
		{
			name:    "NaN COGS",
			input:   TurnoverInput{COGS: math.NaN(), AverageInventoryValue: 20000},
			wantErr: ErrNonFiniteInput,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := Turnover(tt.input)
			if err != tt.wantErr {
				t.Fatalf("Turnover() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}

			if !almostEqual(got, tt.want, tt.tolerance) {
				t.Fatalf("Turnover() = %v, want %v", got, tt.want)
			}
		})
	}
}
