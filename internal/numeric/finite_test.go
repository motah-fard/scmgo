package numeric

import (
	"math"
	"testing"
)

func TestAllFinite(t *testing.T) {
	tests := []struct {
		name   string
		values []float64
		want   bool
	}{
		{name: "no values", values: nil, want: true},
		{name: "all finite", values: []float64{1, -2.5, 0}, want: true},
		{name: "contains NaN", values: []float64{1, math.NaN(), 3}, want: false},
		{name: "contains +Inf", values: []float64{1, math.Inf(1)}, want: false},
		{name: "contains -Inf", values: []float64{1, math.Inf(-1)}, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AllFinite(tt.values...); got != tt.want {
				t.Fatalf("AllFinite(%v) = %v, want %v", tt.values, got, tt.want)
			}
		})
	}
}
