package logistics

import (
	"math"
	"testing"
)

func TestCenterOfGravity(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   CenterOfGravityInput
		want    CenterOfGravityResult
		wantErr error
	}{
		{
			name: "valid input",
			input: CenterOfGravityInput{
				Locations: []LocationDemand{
					{X: 10, Y: 20, Demand: 500},
					{X: 30, Y: 40, Demand: 300},
					{X: 50, Y: 10, Demand: 200},
				},
			},
			want: CenterOfGravityResult{X: 24, Y: 24},
		},
		{
			name:    "empty locations",
			input:   CenterOfGravityInput{Locations: nil},
			wantErr: ErrEmptyLocations,
		},
		{
			name: "negative demand",
			input: CenterOfGravityInput{
				Locations: []LocationDemand{{X: 0, Y: 0, Demand: -1}},
			},
			wantErr: ErrNegativeDemand,
		},
		{
			name: "zero total demand",
			input: CenterOfGravityInput{
				Locations: []LocationDemand{{X: 0, Y: 0, Demand: 0}, {X: 10, Y: 10, Demand: 0}},
			},
			wantErr: ErrZeroTotalDemand,
		},
		{
			name: "NaN X",
			input: CenterOfGravityInput{
				Locations: []LocationDemand{{X: math.NaN(), Y: 0, Demand: 100}},
			},
			wantErr: ErrNonFiniteInput,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := CenterOfGravity(tt.input)
			if err != tt.wantErr {
				t.Fatalf("CenterOfGravity() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}

			if !almostEqual(got.X, tt.want.X, 1e-9) {
				t.Fatalf("CenterOfGravity() X = %v, want %v", got.X, tt.want.X)
			}
			if !almostEqual(got.Y, tt.want.Y, 1e-9) {
				t.Fatalf("CenterOfGravity() Y = %v, want %v", got.Y, tt.want.Y)
			}
		})
	}
}
