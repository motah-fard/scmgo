package finance

import (
	"math"
	"testing"
)

func TestDPO(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   DPOInput
		want    float64
		wantErr error
	}{
		{
			name:  "valid input",
			input: DPOInput{AccountsPayable: 120000, COGS: 1200000, DaysInPeriod: 365},
			want:  36.5,
		},
		{
			name:    "negative payables",
			input:   DPOInput{AccountsPayable: -1, COGS: 1200000, DaysInPeriod: 365},
			wantErr: ErrNegativePayables,
		},
		{
			name:    "zero COGS",
			input:   DPOInput{AccountsPayable: 120000, COGS: 0, DaysInPeriod: 365},
			wantErr: ErrInvalidCOGS,
		},
		{
			name:    "zero days in period",
			input:   DPOInput{AccountsPayable: 120000, COGS: 1200000, DaysInPeriod: 0},
			wantErr: ErrInvalidDaysInPeriod,
		},
		{
			name:    "NaN COGS",
			input:   DPOInput{AccountsPayable: 120000, COGS: math.NaN(), DaysInPeriod: 365},
			wantErr: ErrNonFiniteInput,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := DPO(tt.input)
			if err != tt.wantErr {
				t.Fatalf("DPO() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}

			if !almostEqual(got, tt.want, 1e-9) {
				t.Fatalf("DPO() = %v, want %v", got, tt.want)
			}
		})
	}
}
