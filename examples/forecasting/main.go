// Command forecasting demonstrates running several forecasting methods
// over the same historical series and comparing them with accuracy
// metrics — the usual first step in picking a forecast method for a SKU.
//
// Run it with:
//
//	go run ./examples/forecasting
package main

import (
	"fmt"
	"log"

	"github.com/motah-fard/scmgo/forecast"
)

func main() {
	// 8 weeks of actual demand for one SKU.
	history := []float64{210, 225, 198, 240, 233, 219, 251, 244}

	moving, err := forecast.MovingAverage(forecast.MovingAverageInput{
		History: history,
		Periods: 4,
	})
	if err != nil {
		log.Fatalf("moving average: %v", err)
	}

	ses, err := forecast.SimpleExponentialSmoothing(forecast.SimpleExponentialSmoothingInput{
		History: history,
		Alpha:   0.3,
	})
	if err != nil {
		log.Fatalf("simple exponential smoothing: %v", err)
	}

	trend, err := forecast.HoltLinearTrend(forecast.HoltLinearTrendInput{
		History:      history,
		Alpha:        0.3,
		Beta:         0.2,
		PeriodsAhead: 1,
	})
	if err != nil {
		log.Fatalf("Holt linear trend: %v", err)
	}

	fmt.Println("One-week-ahead forecasts from the same 8-week history:")
	fmt.Printf("  Moving average (4wk):   %.1f\n", moving)
	fmt.Printf("  Simple exp. smoothing:  %.1f\n", ses.Forecast)
	fmt.Printf("  Holt linear trend:      %.1f\n", trend.Forecast)

	// Fitted holds the one-step-ahead in-sample forecast for each period;
	// comparing it against the actual history at the same periods is the
	// standard way to check how well a method already fit the data.
	accuracy, err := forecast.Accuracy(forecast.AccuracyInput{
		Actual:   history[4:],
		Forecast: ses.Fitted[4:],
	})
	if err != nil {
		log.Fatalf("accuracy: %v", err)
	}

	fmt.Println("\nIn-sample accuracy of the SES fit over the last 4 periods:")
	fmt.Printf("  MAD=%.2f  MAPE=%.2f%%  Bias=%.2f  RMSE=%.2f\n",
		accuracy.MAD, accuracy.MAPE*100, accuracy.Bias, accuracy.RMSE)
}
