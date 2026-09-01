package finance

import (
	"math"
	"testing"
)

func almostEqual(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}

func TestDSO(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   DSOInput
		want    float64
		wantErr error
	}{
		{
			name:  "valid input",
			input: DSOInput{AccountsReceivable: 150000, TotalCreditSales: 1800000, DaysInPeriod: 365},
			want:  30.416666666666664,
		},
		{
			name:    "negative receivables",
			input:   DSOInput{AccountsReceivable: -1, TotalCreditSales: 1800000, DaysInPeriod: 365},
			wantErr: ErrNegativeReceivables,
		},
		{
			name:    "zero credit sales",
			input:   DSOInput{AccountsReceivable: 150000, TotalCreditSales: 0, DaysInPeriod: 365},
			wantErr: ErrInvalidCreditSales,
		},
		{
			name:    "zero days in period",
			input:   DSOInput{AccountsReceivable: 150000, TotalCreditSales: 1800000, DaysInPeriod: 0},
			wantErr: ErrInvalidDaysInPeriod,
		},
		{
			name:    "NaN accounts receivable",
			input:   DSOInput{AccountsReceivable: math.NaN(), TotalCreditSales: 1800000, DaysInPeriod: 365},
			wantErr: ErrNonFiniteInput,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := DSO(tt.input)
			if err != tt.wantErr {
				t.Fatalf("DSO() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}

			if !almostEqual(got, tt.want, 1e-9) {
				t.Fatalf("DSO() = %v, want %v", got, tt.want)
			}
		})
	}
}
