// Command intermittent-demand demonstrates forecasting a sparse,
// intermittent SKU — most periods have zero demand, which breaks
// standard moving-average/exponential-smoothing methods (they'd forecast
// toward zero). Croston's method and its bias-corrected variant, SBA,
// exist specifically for this.
//
// Run it with:
//
//	go run ./examples/intermittent-demand
package main

import (
	"fmt"
	"log"

	"github.com/motah-fard/scmgo/forecast"
)

func main() {
	// A spare-parts SKU: demand only in 4 of the last 16 periods.
	history := []float64{0, 0, 5, 0, 0, 0, 0, 3, 0, 0, 4, 0, 0, 0, 0, 6}

	// First, check whether this SKU is actually intermittent (as opposed
	// to smooth, erratic, or lumpy) before reaching for Croston at all —
	// classification tells you which forecast family is appropriate.
	class, err := forecast.ClassifyDemandPattern(forecast.DemandClassificationInput{
		History: history,
	})
	if err != nil {
		log.Fatalf("demand classification: %v", err)
	}

	fmt.Printf("Demand pattern: %s (ADI=%.2f, CV²=%.2f)\n", class.Class, class.ADI, class.CV2)

	croston, err := forecast.Croston(forecast.CrostonInput{
		History: history,
		Alpha:   0.2,
	})
	if err != nil {
		log.Fatalf("Croston: %v", err)
	}

	// SBA (Syntetos-Boylan Approximation) corrects a known upward bias in
	// Croston's estimate; prefer it unless you have a specific reason not to.
	sba, err := forecast.SBA(forecast.CrostonInput{
		History: history,
		Alpha:   0.2,
	})
	if err != nil {
		log.Fatalf("SBA: %v", err)
	}

	fmt.Printf("\nCroston forecast (per period): %.3f\n", croston.Forecast)
	fmt.Printf("SBA forecast (per period):     %.3f\n", sba.Forecast)
}
