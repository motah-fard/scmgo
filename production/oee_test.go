package production

import (
	"math"
	"testing"
)

func almostEqual(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}

func TestOEE(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   OEEInput
		want    OEEResult
		wantErr error
	}{
		{
			name: "valid input",
			input: OEEInput{
				PlannedProductionTime: 480,
				RunTime:               420,
				IdealCycleTime:        1.0,
				TotalCount:            400,
				GoodCount:             385,
			},
			want: OEEResult{
				Availability: 0.875,
				Performance:  0.9523809523809523,
				Quality:      0.9625,
				OEE:          0.8020833333333333,
			},
		},
		{
			name: "zero planned production time",
			input: OEEInput{
				PlannedProductionTime: 0,
				RunTime:               420,
				IdealCycleTime:        1.0,
				TotalCount:            400,
				GoodCount:             385,
			},
			wantErr: ErrInvalidPlannedProductionTime,
		},
		{
			name: "run time exceeds planned production time",
			input: OEEInput{
				PlannedProductionTime: 480,
				RunTime:               500,
				IdealCycleTime:        1.0,
				TotalCount:            400,
				GoodCount:             385,
			},
			wantErr: ErrInvalidRunTime,
		},
		{
			name: "zero run time",
			input: OEEInput{
				PlannedProductionTime: 480,
				RunTime:               0,
				IdealCycleTime:        1.0,
				TotalCount:            400,
				GoodCount:             385,
			},
			wantErr: ErrInvalidRunTime,
		},
		{
			name: "negative ideal cycle time",
			input: OEEInput{
				PlannedProductionTime: 480,
				RunTime:               420,
				IdealCycleTime:        -1,
				TotalCount:            400,
				GoodCount:             385,
			},
			wantErr: ErrNegativeIdealCycleTime,
		},
		{
			name: "zero total count",
			input: OEEInput{
				PlannedProductionTime: 480,
				RunTime:               420,
				IdealCycleTime:        1.0,
				TotalCount:            0,
				GoodCount:             0,
			},
			wantErr: ErrInvalidTotalCount,
		},
		{
			name: "good count exceeds total count",
			input: OEEInput{
				PlannedProductionTime: 480,
				RunTime:               420,
				IdealCycleTime:        1.0,
				TotalCount:            400,
				GoodCount:             450,
			},
			wantErr: ErrInvalidGoodCount,
		},
		{
			name: "negative good count",
			input: OEEInput{
				PlannedProductionTime: 480,
				RunTime:               420,
				IdealCycleTime:        1.0,
				TotalCount:            400,
				GoodCount:             -1,
			},
			wantErr: ErrInvalidGoodCount,
		},
		{
			name: "NaN ideal cycle time",
			input: OEEInput{
				PlannedProductionTime: 480,
				RunTime:               420,
				IdealCycleTime:        math.NaN(),
				TotalCount:            400,
				GoodCount:             385,
			},
			wantErr: ErrNonFiniteInput,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := OEE(tt.input)
			if err != tt.wantErr {
				t.Fatalf("OEE() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}

			if !almostEqual(got.Availability, tt.want.Availability, 1e-9) {
				t.Fatalf("OEE() Availability = %v, want %v", got.Availability, tt.want.Availability)
			}
			if !almostEqual(got.Performance, tt.want.Performance, 1e-9) {
				t.Fatalf("OEE() Performance = %v, want %v", got.Performance, tt.want.Performance)
			}
			if !almostEqual(got.Quality, tt.want.Quality, 1e-9) {
				t.Fatalf("OEE() Quality = %v, want %v", got.Quality, tt.want.Quality)
			}
			if !almostEqual(got.OEE, tt.want.OEE, 1e-9) {
				t.Fatalf("OEE() OEE = %v, want %v", got.OEE, tt.want.OEE)
			}
		})
	}
}
