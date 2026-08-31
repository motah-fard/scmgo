package forecast

import (
	"math"
	"testing"
)

func TestTrackingSignal(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   TrackingSignalInput
		want    []float64
		wantErr error
	}{
		{
			name: "valid input",
			input: TrackingSignalInput{
				Actual:   []float64{100, 110, 95, 130},
				Forecast: []float64{90, 115, 100, 120},
			},
			want: []float64{1.0, 0.6666666666666666, 0.0, 1.3333333333333333},
		},
		{
			name: "perfect forecast yields zero throughout",
			input: TrackingSignalInput{
				Actual:   []float64{100, 110, 95},
				Forecast: []float64{100, 110, 95},
			},
			want: []float64{0, 0, 0},
		},
		{
			name: "empty actual",
			input: TrackingSignalInput{
				Actual:   []float64{},
				Forecast: []float64{},
			},
			wantErr: ErrEmptyHistory,
		},
		{
			name: "mismatched lengths",
			input: TrackingSignalInput{
				Actual:   []float64{1, 2},
				Forecast: []float64{1},
			},
			wantErr: ErrMismatchedLengths,
		},
		{
			name: "NaN in actual",
			input: TrackingSignalInput{
				Actual:   []float64{100, math.NaN()},
				Forecast: []float64{90, 100},
			},
			wantErr: ErrNonFiniteInput,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := TrackingSignal(tt.input)
			if err != tt.wantErr {
				t.Fatalf("TrackingSignal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}

			if len(got) != len(tt.want) {
				t.Fatalf("TrackingSignal() length = %v, want %v", len(got), len(tt.want))
			}
			for i := range got {
				if !almostEqual(got[i], tt.want[i], 1e-9) {
					t.Fatalf("TrackingSignal()[%d] = %v, want %v", i, got[i], tt.want[i])
				}
			}
		})
	}
}
