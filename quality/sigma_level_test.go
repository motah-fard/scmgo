package quality

import (
	"math"
	"testing"
)

func TestSigmaLevel(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		dpmo      float64
		want      float64
		tolerance float64
		wantErr   error
	}{
		{
			name:      "classic six sigma reference point",
			dpmo:      3.4,
			want:      5.999854470025545,
			tolerance: 1e-6,
		},
		{
			name:      "classic three sigma reference point",
			dpmo:      66807,
			want:      3.000001553990341,
			tolerance: 1e-6,
		},
		{
			name:      "matches earlier DPMO test case",
			dpmo:      9000,
			want:      3.865618126864293,
			tolerance: 1e-6,
		},
		{
			name:    "zero DPMO is out of range",
			dpmo:    0,
			wantErr: ErrInvalidDPMO,
		},
		{
			name:    "1,000,000 DPMO is out of range",
			dpmo:    1_000_000,
			wantErr: ErrInvalidDPMO,
		},
		{
			name:    "negative DPMO",
			dpmo:    -1,
			wantErr: ErrInvalidDPMO,
		},
		{
			name:    "NaN DPMO",
			dpmo:    math.NaN(),
			wantErr: ErrNonFiniteInput,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := SigmaLevel(tt.dpmo)
			if err != tt.wantErr {
				t.Fatalf("SigmaLevel() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}

			if !almostEqual(got, tt.want, tt.tolerance) {
				t.Fatalf("SigmaLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}
