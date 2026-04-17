package inventory

import (
	"math"
	"testing"
)

func TestStdDevDemandDuringLeadTime(t *testing.T) {
	tests := []struct {
		name    string
		input   StdDevDemandDuringLeadTimeInput
		want    float64
		wantErr error
	}{
		{
			name: "valid inputs",
			input: StdDevDemandDuringLeadTimeInput{
				StdDevDailyDemand: 10,
				LeadTimeDays:      4,
			},
			want: 20,
		},
		{
			name: "fractional lead time",
			input: StdDevDemandDuringLeadTimeInput{
				StdDevDailyDemand: 12,
				LeadTimeDays:      2.25,
			},
			want: 18,
		},
		{
			name: "zero standard deviation",
			input: StdDevDemandDuringLeadTimeInput{
				StdDevDailyDemand: 0,
				LeadTimeDays:      5,
			},
			want: 0,
		},
		{
			name: "zero lead time",
			input: StdDevDemandDuringLeadTimeInput{
				StdDevDailyDemand: 10,
				LeadTimeDays:      0,
			},
			want: 0,
		},
		{
			name: "negative standard deviation",
			input: StdDevDemandDuringLeadTimeInput{
				StdDevDailyDemand: -1,
				LeadTimeDays:      5,
			},
			wantErr: ErrNegativeStandardDeviation,
		},
		{
			name: "negative lead time",
			input: StdDevDemandDuringLeadTimeInput{
				StdDevDailyDemand: 10,
				LeadTimeDays:      -1,
			},
			wantErr: ErrNegativeLeadTime,
		},
	}

	const tolerance = 1e-9

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StdDevDemandDuringLeadTime(tt.input)

			if tt.wantErr != nil {
				if err == nil {
					t.Fatalf("expected error %v, got nil", tt.wantErr)
				}
				if err != tt.wantErr {
					t.Fatalf("expected error %v, got %v", tt.wantErr, err)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if math.Abs(got-tt.want) > tolerance {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}
