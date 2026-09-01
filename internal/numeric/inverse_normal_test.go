package numeric

import (
	"math"
	"testing"
)

func TestInverseNormalCDF(t *testing.T) {
	tests := []struct {
		name      string
		p         float64
		want      float64
		tolerance float64
	}{
		{name: "p=0.5", p: 0.5, want: 0, tolerance: 1e-12},
		{name: "p=0.95", p: 0.95, want: 1.6448536269514724, tolerance: 1e-9},
		{name: "p=0.05", p: 0.05, want: -1.6448536269514724, tolerance: 1e-9},
		{name: "p=0.90", p: 0.90, want: 1.2815515655446004, tolerance: 1e-9},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := InverseNormalCDF(tt.p)
			if math.Abs(got-tt.want) > tt.tolerance {
				t.Fatalf("InverseNormalCDF(%v) = %v, want %v", tt.p, got, tt.want)
			}
		})
	}
}
