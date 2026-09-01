package production

import (
	"math"
	"testing"
)

func TestWIPFromLittlesLaw(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   WIPFromLittlesLawInput
		want    float64
		wantErr error
	}{
		{
			name:  "valid input",
			input: WIPFromLittlesLawInput{Throughput: 10, CycleTime: 5},
			want:  50,
		},
		{
			name:    "zero throughput",
			input:   WIPFromLittlesLawInput{Throughput: 0, CycleTime: 5},
			wantErr: ErrInvalidThroughput,
		},
		{
			name:    "negative cycle time",
			input:   WIPFromLittlesLawInput{Throughput: 10, CycleTime: -1},
			wantErr: ErrNegativeCycleTime,
		},
		{
			name:    "NaN throughput",
			input:   WIPFromLittlesLawInput{Throughput: math.NaN(), CycleTime: 5},
			wantErr: ErrNonFiniteInput,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := WIPFromLittlesLaw(tt.input)
			if err != tt.wantErr {
				t.Fatalf("WIPFromLittlesLaw() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}

			if !almostEqual(got, tt.want, 1e-9) {
				t.Fatalf("WIPFromLittlesLaw() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCycleTimeFromLittlesLaw(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   CycleTimeFromLittlesLawInput
		want    float64
		wantErr error
	}{
		{
			name:  "valid input",
			input: CycleTimeFromLittlesLawInput{WIP: 50, Throughput: 10},
			want:  5,
		},
		{
			name:    "negative WIP",
			input:   CycleTimeFromLittlesLawInput{WIP: -1, Throughput: 10},
			wantErr: ErrNegativeWIP,
		},
		{
			name:    "zero throughput",
			input:   CycleTimeFromLittlesLawInput{WIP: 50, Throughput: 0},
			wantErr: ErrInvalidThroughput,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := CycleTimeFromLittlesLaw(tt.input)
			if err != tt.wantErr {
				t.Fatalf("CycleTimeFromLittlesLaw() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}

			if !almostEqual(got, tt.want, 1e-9) {
				t.Fatalf("CycleTimeFromLittlesLaw() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestThroughputFromLittlesLaw(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   ThroughputFromLittlesLawInput
		want    float64
		wantErr error
	}{
		{
			name:  "valid input",
			input: ThroughputFromLittlesLawInput{WIP: 50, CycleTime: 5},
			want:  10,
		},
		{
			name:    "negative WIP",
			input:   ThroughputFromLittlesLawInput{WIP: -1, CycleTime: 5},
			wantErr: ErrNegativeWIP,
		},
		{
			name:    "zero cycle time",
			input:   ThroughputFromLittlesLawInput{WIP: 50, CycleTime: 0},
			wantErr: ErrInvalidCycleTime,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := ThroughputFromLittlesLaw(tt.input)
			if err != tt.wantErr {
				t.Fatalf("ThroughputFromLittlesLaw() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}

			if !almostEqual(got, tt.want, 1e-9) {
				t.Fatalf("ThroughputFromLittlesLaw() = %v, want %v", got, tt.want)
			}
		})
	}
}
