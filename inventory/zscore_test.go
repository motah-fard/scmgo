package inventory

import (
	"math"
	"testing"
)

func TestZScoreForServiceLevel(t *testing.T) {
	tests := []struct {
		name         string
		serviceLevel float64
		want         float64
		wantErr      bool
	}{
		{
			name:         "90 percent",
			serviceLevel: 0.90,
			want:         1.2815515655,
			wantErr:      false,
		},
		{
			name:         "95 percent",
			serviceLevel: 0.95,
			want:         1.6448536269,
			wantErr:      false,
		},
		{
			name:         "97.5 percent",
			serviceLevel: 0.975,
			want:         1.9599639845,
			wantErr:      false,
		},
		{
			name:         "99 percent",
			serviceLevel: 0.99,
			want:         2.3263478740,
			wantErr:      false,
		},
		{
			name:         "zero",
			serviceLevel: 0,
			wantErr:      true,
		},
		{
			name:         "one",
			serviceLevel: 1,
			wantErr:      true,
		},
		{
			name:         "negative",
			serviceLevel: -0.10,
			wantErr:      true,
		},
		{
			name:         "greater than one",
			serviceLevel: 1.10,
			wantErr:      true,
		},
	}

	const tolerance = 1e-4

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ZScoreForServiceLevel(tt.serviceLevel)

			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				if err != ErrInvalidServiceLevel {
					t.Fatalf("expected ErrInvalidServiceLevel, got %v", err)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if math.Abs(got-tt.want) > tolerance {
				t.Fatalf("got %.6f, want %.6f", got, tt.want)
			}
		})
	}
}
