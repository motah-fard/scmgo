// Command retail is the "everything at once" demo: for a small set of
// SKUs with different demand patterns, it classifies each one's demand
// pattern, picks an appropriate forecast method, ranks the SKUs by value
// (ABC), and turns the result into a reorder point and order quantity —
// the shape of what a real inventory-planning backend does per SKU.
//
// Run it with:
//
//	go run ./examples/retail
package main

import (
	"fmt"
	"log"
	"math"

	"github.com/motah-fard/scmgo/abc"
	"github.com/motah-fard/scmgo/forecast"
	"github.com/motah-fard/scmgo/inventory"
)

// sku is the input data you'd normally pull from an order-history table:
// one demand series (here, daily units for the last 14 days), the
// supplier lead time, and the unit cost used to rank SKUs by value.
type sku struct {
	id           string
	history      []float64
	leadTimeDays float64
	unitCost     float64
}

func main() {
	skus := []sku{
		{
			id:           "SKU-STEADY",
			history:      []float64{48, 52, 50, 47, 53, 49, 51, 50, 48, 52, 49, 51, 50, 52},
			leadTimeDays: 5,
			unitCost:     18,
		},
		{
			id:           "SKU-SPARSE",
			history:      []float64{0, 0, 6, 0, 0, 0, 0, 4, 0, 0, 0, 5, 0, 0},
			leadTimeDays: 10,
			unitCost:     140,
		},
		{
			id:           "SKU-VOLATILE",
			history:      []float64{10, 90, 5, 95, 15, 80, 8, 92, 20, 85, 3, 98, 12, 88},
			leadTimeDays: 7,
			unitCost:     32,
		},
	}

	valueItems := make([]abc.Item, len(skus))
	rows := make([]row, len(skus))

	for i, s := range skus {
		r, err := planSKU(s)
		if err != nil {
			log.Fatalf("%s: %v", s.id, err)
		}
		rows[i] = r
		valueItems[i] = abc.Item{ID: s.id, Value: r.annualValue}
	}

	abcResults, err := abc.Classify(abc.ClassifyInput{
		Items:      valueItems,
		AThreshold: 0.80,
		BThreshold: 0.95,
	})
	if err != nil {
		log.Fatalf("ABC classify: %v", err)
	}
	classByID := make(map[string]string, len(abcResults))
	for _, c := range abcResults {
		classByID[c.ID] = c.Class
	}

	fmt.Printf("%-13s %-13s %10s %6s %8s %8s %6s\n",
		"SKU", "Pattern", "AnnualVal", "Class", "Forecast", "ROP", "EOQ")
	for _, r := range rows {
		fmt.Printf("%-13s %-13s %10.0f %6s %8.1f %8.0f %6.0f\n",
			r.id, r.pattern, r.annualValue, classByID[r.id], r.forecast, r.reorderPoint, r.eoq)
	}
}

type row struct {
	id           string
	pattern      string
	annualValue  float64
	forecast     float64
	reorderPoint float64
	eoq          float64
}

// planSKU runs the forecast -> classify -> policy pipeline for one SKU.
func planSKU(s sku) (row, error) {
	class, err := forecast.ClassifyDemandPattern(forecast.DemandClassificationInput{History: s.history})
	if err != nil {
		return row{}, fmt.Errorf("classify demand pattern: %w", err)
	}

	// Intermittent/lumpy demand needs Croston-family forecasting; smooth
	// or erratic demand is dense enough for exponential smoothing.
	var dailyDemand, stdDevDailyDemand float64
	switch class.Class {
	case "intermittent", "lumpy":
		sba, err := forecast.SBA(forecast.CrostonInput{History: s.history, Alpha: 0.2})
		if err != nil {
			return row{}, fmt.Errorf("SBA: %w", err)
		}
		dailyDemand = sba.Forecast
		stdDevDailyDemand = sba.Forecast // sparse-demand variability is dominated by the zero/non-zero split itself
	default:
		ses, err := forecast.SimpleExponentialSmoothing(forecast.SimpleExponentialSmoothingInput{History: s.history, Alpha: 0.3})
		if err != nil {
			return row{}, fmt.Errorf("simple exponential smoothing: %w", err)
		}
		dailyDemand = ses.Forecast
		stdDevDailyDemand = sampleStdDev(s.history)
	}

	summary, err := inventory.BuildPolicySummaryWithServiceLevel(inventory.PolicySummaryServiceLevelInput{
		DailyDemand:        dailyDemand,
		LeadTimeDays:       s.leadTimeDays,
		ReviewPeriodDays:   s.leadTimeDays, // review as often as the lead time, for this demo
		DemandStdDevPerDay: stdDevDailyDemand,
		ServiceLevel:       0.95,
	})
	if err != nil {
		return row{}, fmt.Errorf("policy summary: %w", err)
	}

	eoq, err := inventory.EOQ(inventory.EOQInput{
		AnnualDemand:       dailyDemand * 365,
		OrderingCost:       50,
		HoldingCostPerUnit: s.unitCost * 0.20, // 20%/year holding cost rate
	})
	if err != nil {
		return row{}, fmt.Errorf("EOQ: %w", err)
	}

	return row{
		id:           s.id,
		pattern:      class.Class,
		annualValue:  dailyDemand * 365 * s.unitCost,
		forecast:     dailyDemand,
		reorderPoint: summary.ReorderPoint,
		eoq:          eoq,
	}, nil
}

// sampleStdDev is a plain population standard deviation over history —
// not part of scmgo's API, just the kind of stat you'd compute upstream
// before calling into it.
func sampleStdDev(history []float64) float64 {
	var sum float64
	for _, v := range history {
		sum += v
	}
	mean := sum / float64(len(history))

	var sq float64
	for _, v := range history {
		sq += (v - mean) * (v - mean)
	}

	return math.Sqrt(sq / float64(len(history)))
}
