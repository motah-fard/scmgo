package quality

import (
	"math"
	"testing"
)

func TestCp(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   CpInput
		want    float64
		wantErr error
	}{
		{
			name:  "valid input",
			input: CpInput{USL: 10.5, LSL: 9.5, Sigma: 0.1},
			want:  1.6666666666666665,
		},
		{
			name:    "USL not greater than LSL",
			input:   CpInput{USL: 9.5, LSL: 9.5, Sigma: 0.1},
			wantErr: ErrInvalidSpecLimits,
		},
		{
			name:    "zero sigma",
			input:   CpInput{USL: 10.5, LSL: 9.5, Sigma: 0},
			wantErr: ErrInvalidSigma,
		},
		{
			name:    "NaN sigma",
			input:   CpInput{USL: 10.5, LSL: 9.5, Sigma: math.NaN()},
			wantErr: ErrNonFiniteInput,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := Cp(tt.input)
			if err != tt.wantErr {
				t.Fatalf("Cp() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}

			if !almostEqual(got, tt.want, 1e-9) {
				t.Fatalf("Cp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCpk(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   CpkInput
		want    float64
		wantErr error
	}{
		{
			name:  "valid input",
			input: CpkInput{USL: 10.5, LSL: 9.5, Mean: 10.1, Sigma: 0.1},
			want:  1.3333333333333344,
		},
		{
			name:  "mean outside spec limits yields negative Cpk",
			input: CpkInput{USL: 10.5, LSL: 9.5, Mean: 10.8, Sigma: 0.1},
			want:  -1,
		},
		{
			name:    "USL not greater than LSL",
			input:   CpkInput{USL: 9.5, LSL: 9.5, Mean: 9.5, Sigma: 0.1},
			wantErr: ErrInvalidSpecLimits,
		},
		{
			name:    "zero sigma",
			input:   CpkInput{USL: 10.5, LSL: 9.5, Mean: 10.1, Sigma: 0},
			wantErr: ErrInvalidSigma,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := Cpk(tt.input)
			if err != tt.wantErr {
				t.Fatalf("Cpk() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}

			if !almostEqual(got, tt.want, 1e-9) {
				t.Fatalf("Cpk() = %v, want %v", got, tt.want)
			}
		})
	}
}
