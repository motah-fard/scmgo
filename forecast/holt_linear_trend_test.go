package forecast

import "testing"

func TestHoltLinearTrend(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		input     HoltLinearTrendInput
		want      HoltLinearTrendResult
		tolerance float64
		wantErr   error
	}{
		{
			name: "valid input",
			input: HoltLinearTrendInput{
				History:      []float64{100, 120, 110, 130, 125},
				Alpha:        0.3,
				Beta:         0.2,
				PeriodsAhead: 3,
			},
			want: HoltLinearTrendResult{
				Level:    149.84159999999997,
				Trend:    14.918719999999997,
				Forecast: 194.59775999999997,
			},
			tolerance: 1e-6,
		},
		{
			name: "insufficient history",
			input: HoltLinearTrendInput{
				History:      []float64{100},
				Alpha:        0.3,
				Beta:         0.2,
				PeriodsAhead: 1,
			},
			wantErr: ErrInsufficientHistory,
		},
		{
			name: "zero periods ahead",
			input: HoltLinearTrendInput{
				History:      []float64{100, 120},
				Alpha:        0.3,
				Beta:         0.2,
				PeriodsAhead: 0,
			},
			wantErr: ErrInvalidPeriods,
		},
		{
			name: "invalid alpha",
			input: HoltLinearTrendInput{
				History:      []float64{100, 120},
				Alpha:        0,
				Beta:         0.2,
				PeriodsAhead: 1,
			},
			wantErr: ErrInvalidSmoothingConstant,
		},
		{
			name: "invalid beta",
			input: HoltLinearTrendInput{
				History:      []float64{100, 120},
				Alpha:        0.3,
				Beta:         1.5,
				PeriodsAhead: 1,
			},
			wantErr: ErrInvalidSmoothingConstant,
		},
		{
			name: "negative demand",
			input: HoltLinearTrendInput{
				History:      []float64{100, -120},
				Alpha:        0.3,
				Beta:         0.2,
				PeriodsAhead: 1,
			},
			wantErr: ErrNegativeDemand,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := HoltLinearTrend(tt.input)
			if err != tt.wantErr {
				t.Fatalf("HoltLinearTrend() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}

			if !almostEqual(got.Level, tt.want.Level, tt.tolerance) {
				t.Fatalf("HoltLinearTrend() level = %v, want %v", got.Level, tt.want.Level)
			}
			if !almostEqual(got.Trend, tt.want.Trend, tt.tolerance) {
				t.Fatalf("HoltLinearTrend() trend = %v, want %v", got.Trend, tt.want.Trend)
			}
			if !almostEqual(got.Forecast, tt.want.Forecast, tt.tolerance) {
				t.Fatalf("HoltLinearTrend() forecast = %v, want %v", got.Forecast, tt.want.Forecast)
			}
		})
	}
}
