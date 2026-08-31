package forecast

import "fmt"

func ExampleMovingAverage() {
	forecast, err := MovingAverage(MovingAverageInput{
		History: []float64{100, 120, 110, 130, 125},
		Periods: 3,
	})
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("%.2f\n", forecast)
	// Output: 121.67
}

func ExampleWeightedMovingAverage() {
	forecast, err := WeightedMovingAverage(WeightedMovingAverageInput{
		History: []float64{100, 120, 110, 130, 125},
		Weights: []float64{0.1, 0.3, 0.6},
	})
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("%.2f\n", forecast)
	// Output: 125.00
}

func ExampleSimpleExponentialSmoothing() {
	result, err := SimpleExponentialSmoothing(SimpleExponentialSmoothingInput{
		History: []float64{100, 120, 110, 130, 125},
		Alpha:   0.3,
	})
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("%.2f\n", result.Forecast)
	// Output: 117.33
}

func ExampleHoltLinearTrend() {
	result, err := HoltLinearTrend(HoltLinearTrendInput{
		History:      []float64{100, 120, 110, 130, 125},
		Alpha:        0.3,
		Beta:         0.2,
		PeriodsAhead: 3,
	})
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("%.2f\n", result.Forecast)
	// Output: 194.60
}

func ExampleCroston() {
	result, err := Croston(CrostonInput{
		History: []float64{0, 0, 5, 0, 0, 0, 3, 0, 4, 0},
		Alpha:   0.2,
	})
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("%.2f\n", result.Forecast)
	// Output: 1.51
}

func ExampleAccuracy() {
	result, err := Accuracy(AccuracyInput{
		Actual:   []float64{100, 110, 95, 130},
		Forecast: []float64{90, 115, 100, 120},
	})
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("MAD=%.2f MAPE=%.4f Bias=%.2f RMSE=%.2f\n", result.MAD, result.MAPE, result.Bias, result.RMSE)
	// Output: MAD=7.50 MAPE=0.0688 Bias=2.50 RMSE=7.91
}
