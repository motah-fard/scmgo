package inventory

import (
	"math"
	"testing"
)

func TestUnitNormalLoss(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		z         float64
		want      float64
		tolerance float64
		wantErr   error
	}{
		{name: "z=0", z: 0, want: 0.3989422804014327, tolerance: 1e-9},
		{name: "z=1", z: 1, want: 0.08331547058768629, tolerance: 1e-9},
		{name: "z=-1", z: -1, want: 1.0833154705876864, tolerance: 1e-9},
		{name: "z=2", z: 2, want: 0.008490702616829646, tolerance: 1e-9},
		{name: "z=-2", z: -2, want: 2.0084907026168297, tolerance: 1e-9},
		{name: "NaN z", z: math.NaN(), wantErr: ErrNonFiniteInput},
		{name: "+Inf z", z: math.Inf(1), wantErr: ErrNonFiniteInput},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := UnitNormalLoss(tt.z)
			if err != tt.wantErr {
				t.Fatalf("UnitNormalLoss() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}

			if !almostEqual(got, tt.want, tt.tolerance) {
				t.Fatalf("UnitNormalLoss() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnitNormalLossMonotonicallyDecreasing(t *testing.T) {
	t.Parallel()

	prev, err := UnitNormalLoss(-5)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	for z := -4.9; z <= 5; z += 0.1 {
		got, err := UnitNormalLoss(z)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got > prev {
			t.Fatalf("UnitNormalLoss(%v) = %v is greater than UnitNormalLoss at previous z = %v; expected monotonically decreasing", z, got, prev)
		}
		prev = got
	}
}
