package forecast

import "testing"

func TestHoltWinters(t *testing.T) {
	t.Parallel()

	history := []float64{100, 120, 90, 110, 105, 125, 95, 115}

	tests := []struct {
		name         string
		input        HoltWintersInput
		wantLevel    float64
		wantTrend    float64
		wantSeasonal []float64
		wantForecast float64
		tolerance    float64
		wantErr      error
	}{
		{
			name: "valid input, h=1",
			input: HoltWintersInput{
				History:      history,
				Alpha:        0.3,
				Beta:         0.1,
				Gamma:        0.2,
				SeasonLength: 4,
				PeriodsAhead: 1,
			},
			wantLevel:    111.22255962499999,
			wantTrend:    1.3324808374999992,
			wantSeasonal: []float64{-4.475, 15.176750000000002, -15.072327499999997, 4.755488075000002},
			wantForecast: 108.0800404625,
			tolerance:    1e-6,
		},
		{
			name: "valid input, h=4",
			input: HoltWintersInput{
				History:      history,
				Alpha:        0.3,
				Beta:         0.1,
				Gamma:        0.2,
				SeasonLength: 4,
				PeriodsAhead: 4,
			},
			wantForecast: 121.30797104999999,
			tolerance:    1e-6,
		},
		{
			name: "season length below 2",
			input: HoltWintersInput{
				History:      history,
				Alpha:        0.3,
				Beta:         0.1,
				Gamma:        0.2,
				SeasonLength: 1,
				PeriodsAhead: 1,
			},
			wantErr: ErrInvalidSeasonLength,
		},
		{
			name: "insufficient history for two seasons",
			input: HoltWintersInput{
				History:      []float64{100, 120, 90},
				Alpha:        0.3,
				Beta:         0.1,
				Gamma:        0.2,
				SeasonLength: 4,
				PeriodsAhead: 1,
			},
			wantErr: ErrInsufficientHistory,
		},
		{
			name: "zero periods ahead",
			input: HoltWintersInput{
				History:      history,
				Alpha:        0.3,
				Beta:         0.1,
				Gamma:        0.2,
				SeasonLength: 4,
				PeriodsAhead: 0,
			},
			wantErr: ErrInvalidPeriods,
		},
		{
			name: "invalid gamma",
			input: HoltWintersInput{
				History:      history,
				Alpha:        0.3,
				Beta:         0.1,
				Gamma:        0,
				SeasonLength: 4,
				PeriodsAhead: 1,
			},
			wantErr: ErrInvalidSmoothingConstant,
		},
		{
			name: "negative demand",
			input: HoltWintersInput{
				History:      []float64{100, 120, 90, -110, 105, 125, 95, 115},
				Alpha:        0.3,
				Beta:         0.1,
				Gamma:        0.2,
				SeasonLength: 4,
				PeriodsAhead: 1,
			},
			wantErr: ErrNegativeDemand,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := HoltWinters(tt.input)
			if err != tt.wantErr {
				t.Fatalf("HoltWinters() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}

			if !almostEqual(got.Forecast, tt.wantForecast, tt.tolerance) {
				t.Fatalf("HoltWinters() forecast = %v, want %v", got.Forecast, tt.wantForecast)
			}
			if tt.wantSeasonal == nil {
				return
			}
			if !almostEqual(got.Level, tt.wantLevel, tt.tolerance) {
				t.Fatalf("HoltWinters() level = %v, want %v", got.Level, tt.wantLevel)
			}
			if !almostEqual(got.Trend, tt.wantTrend, tt.tolerance) {
				t.Fatalf("HoltWinters() trend = %v, want %v", got.Trend, tt.wantTrend)
			}
			if len(got.Seasonal) != len(tt.wantSeasonal) {
				t.Fatalf("HoltWinters() seasonal length = %v, want %v", len(got.Seasonal), len(tt.wantSeasonal))
			}
			for i := range got.Seasonal {
				if !almostEqual(got.Seasonal[i], tt.wantSeasonal[i], tt.tolerance) {
					t.Fatalf("HoltWinters() seasonal[%d] = %v, want %v", i, got.Seasonal[i], tt.wantSeasonal[i])
				}
			}
		})
	}
}
