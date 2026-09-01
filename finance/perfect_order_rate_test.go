package finance

import (
	"math"
	"testing"
)

func TestPerfectOrderRate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   PerfectOrderRateInput
		want    float64
		wantErr error
	}{
		{
			name: "valid input",
			input: PerfectOrderRateInput{
				OnTimeRate:                0.95,
				CompleteRate:              0.97,
				DamageFreeRate:            0.99,
				AccurateDocumentationRate: 0.98,
			},
			want: 0.8940393,
		},
		{
			name: "all perfect",
			input: PerfectOrderRateInput{
				OnTimeRate:                1,
				CompleteRate:              1,
				DamageFreeRate:            1,
				AccurateDocumentationRate: 1,
			},
			want: 1,
		},
		{
			name: "rate above one",
			input: PerfectOrderRateInput{
				OnTimeRate:                1.1,
				CompleteRate:              0.97,
				DamageFreeRate:            0.99,
				AccurateDocumentationRate: 0.98,
			},
			wantErr: ErrInvalidRate,
		},
		{
			name: "negative rate",
			input: PerfectOrderRateInput{
				OnTimeRate:                -0.1,
				CompleteRate:              0.97,
				DamageFreeRate:            0.99,
				AccurateDocumentationRate: 0.98,
			},
			wantErr: ErrInvalidRate,
		},
		{
			name: "NaN rate",
			input: PerfectOrderRateInput{
				OnTimeRate:                math.NaN(),
				CompleteRate:              0.97,
				DamageFreeRate:            0.99,
				AccurateDocumentationRate: 0.98,
			},
			wantErr: ErrNonFiniteInput,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := PerfectOrderRate(tt.input)
			if err != tt.wantErr {
				t.Fatalf("PerfectOrderRate() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}

			if !almostEqual(got, tt.want, 1e-9) {
				t.Fatalf("PerfectOrderRate() = %v, want %v", got, tt.want)
			}
		})
	}
}
