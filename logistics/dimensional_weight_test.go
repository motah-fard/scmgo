package logistics

import (
	"math"
	"testing"
)

func almostEqual(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}

func TestDimensionalWeight(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   DimensionalWeightInput
		want    float64
		wantErr error
	}{
		{
			name:  "valid input",
			input: DimensionalWeightInput{Length: 20, Width: 15, Height: 10, DimFactor: 139},
			want:  21.58273381294964,
		},
		{
			name:    "zero length",
			input:   DimensionalWeightInput{Length: 0, Width: 15, Height: 10, DimFactor: 139},
			wantErr: ErrInvalidDimension,
		},
		{
			name:    "zero dim factor",
			input:   DimensionalWeightInput{Length: 20, Width: 15, Height: 10, DimFactor: 0},
			wantErr: ErrInvalidDimFactor,
		},
		{
			name:    "NaN width",
			input:   DimensionalWeightInput{Length: 20, Width: math.NaN(), Height: 10, DimFactor: 139},
			wantErr: ErrNonFiniteInput,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := DimensionalWeight(tt.input)
			if err != tt.wantErr {
				t.Fatalf("DimensionalWeight() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}

			if !almostEqual(got, tt.want, 1e-9) {
				t.Fatalf("DimensionalWeight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBillableWeight(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name              string
		actualWeight      float64
		dimensionalWeight float64
		want              float64
		wantErr           error
	}{
		{name: "actual heavier", actualWeight: 30, dimensionalWeight: 21.58, want: 30},
		{name: "dimensional heavier", actualWeight: 15, dimensionalWeight: 21.58, want: 21.58},
		{name: "equal", actualWeight: 20, dimensionalWeight: 20, want: 20},
		{name: "negative actual weight", actualWeight: -1, dimensionalWeight: 20, wantErr: ErrNegativeWeight},
		{name: "NaN dimensional weight", actualWeight: 20, dimensionalWeight: math.NaN(), wantErr: ErrNonFiniteInput},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := BillableWeight(tt.actualWeight, tt.dimensionalWeight)
			if err != tt.wantErr {
				t.Fatalf("BillableWeight() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}

			if !almostEqual(got, tt.want, 1e-9) {
				t.Fatalf("BillableWeight() = %v, want %v", got, tt.want)
			}
		})
	}
}
