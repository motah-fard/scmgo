package forecast

import (
	"math"
	"testing"
)

func TestSeasonalNaive(t *testing.T) {
	t.Parallel()

	history := []float64{100, 120, 90, 110, 105, 125, 95, 115}

	tests := []struct {
		name    string
		input   SeasonalNaiveInput
		want    float64
		wantErr error
	}{
		{
			name:  "h=1",
			input: SeasonalNaiveInput{History: history, SeasonLength: 4, PeriodsAhead: 1},
			want:  105,
		},
		{
			name:  "h=4",
			input: SeasonalNaiveInput{History: history, SeasonLength: 4, PeriodsAhead: 4},
			want:  115,
		},
		{
			name:  "h=5 wraps to same seasonal position as h=1",
			input: SeasonalNaiveInput{History: history, SeasonLength: 4, PeriodsAhead: 5},
			want:  105,
		},
		{
			name:    "season length below 2",
			input:   SeasonalNaiveInput{History: history, SeasonLength: 1, PeriodsAhead: 1},
			wantErr: ErrInvalidSeasonLength,
		},
		{
			name:    "insufficient history",
			input:   SeasonalNaiveInput{History: []float64{100, 120}, SeasonLength: 4, PeriodsAhead: 1},
			wantErr: ErrInsufficientHistory,
		},
		{
			name:    "zero periods ahead",
			input:   SeasonalNaiveInput{History: history, SeasonLength: 4, PeriodsAhead: 0},
			wantErr: ErrInvalidPeriods,
		},
		{
			name:    "negative demand",
			input:   SeasonalNaiveInput{History: []float64{100, -120, 90, 110}, SeasonLength: 4, PeriodsAhead: 1},
			wantErr: ErrNegativeDemand,
		},
		{
			name:    "NaN in history",
			input:   SeasonalNaiveInput{History: []float64{100, math.NaN(), 90, 110}, SeasonLength: 4, PeriodsAhead: 1},
			wantErr: ErrNonFiniteInput,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := SeasonalNaive(tt.input)
			if err != tt.wantErr {
				t.Fatalf("SeasonalNaive() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}

			if got != tt.want {
				t.Fatalf("SeasonalNaive() = %v, want %v", got, tt.want)
			}
		})
	}
}
