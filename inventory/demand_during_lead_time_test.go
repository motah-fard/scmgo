package inventory

import "testing"

func TestDemandDuringLeadTime(t *testing.T) {
	tests := []struct {
		name    string
		input   DemandDuringLeadTimeInput
		want    float64
		wantErr error
	}{
		{
			name: "valid inputs",
			input: DemandDuringLeadTimeInput{
				AvgDailyDemand: 100,
				LeadTimeDays:   5,
			},
			want: 500,
		},
		{
			name: "zero demand",
			input: DemandDuringLeadTimeInput{
				AvgDailyDemand: 0,
				LeadTimeDays:   5,
			},
			want: 0,
		},
		{
			name: "zero lead time",
			input: DemandDuringLeadTimeInput{
				AvgDailyDemand: 100,
				LeadTimeDays:   0,
			},
			want: 0,
		},
		{
			name: "zero demand and zero lead time",
			input: DemandDuringLeadTimeInput{
				AvgDailyDemand: 0,
				LeadTimeDays:   0,
			},
			want: 0,
		},
		{
			name: "negative demand",
			input: DemandDuringLeadTimeInput{
				AvgDailyDemand: -1,
				LeadTimeDays:   5,
			},
			wantErr: ErrNegativeDemand,
		},
		{
			name: "negative lead time",
			input: DemandDuringLeadTimeInput{
				AvgDailyDemand: 100,
				LeadTimeDays:   -2,
			},
			wantErr: ErrNegativeLeadTime,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DemandDuringLeadTime(tt.input)

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

			if got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}
