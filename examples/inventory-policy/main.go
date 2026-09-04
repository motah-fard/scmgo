// Command inventory-policy demonstrates the most common scmgo workflow:
// turn demand, lead time, and a target service level into a complete
// inventory policy (safety stock, reorder point, target level, min/max)
// plus an economic order quantity.
//
// Run it with:
//
//	go run ./examples/inventory-policy
package main

import (
	"fmt"
	"log"

	"github.com/motah-fard/scmgo/inventory"
)

func main() {
	// Inputs you'd typically pull from a demand-planning system: average
	// daily demand and its variability, supplier lead time, how often you
	// review the SKU, and the service level you're targeting.
	summary, err := inventory.BuildPolicySummaryWithServiceLevel(inventory.PolicySummaryServiceLevelInput{
		DailyDemand:        120,
		LeadTimeDays:       7,
		ReviewPeriodDays:   14,
		DemandStdDevPerDay: 25,
		ServiceLevel:       0.95,
	})
	if err != nil {
		log.Fatalf("policy summary: %v", err)
	}

	fmt.Println("Inventory policy (95% service level, 7-day lead time):")
	fmt.Printf("  Expected demand during lead time: %.0f units\n", summary.ExpectedDemandDuringLeadTime)
	fmt.Printf("  Safety stock:                     %.0f units\n", summary.SafetyStockUnits)
	fmt.Printf("  Reorder point:                    %.0f units\n", summary.ReorderPoint)
	fmt.Printf("  Target inventory level:           %.0f units\n", summary.TargetInventoryLevel)
	fmt.Printf("  Min/Max:                          %.0f / %.0f units\n", summary.MinLevel, summary.MaxLevel)

	// How much to order each time you hit the reorder point, given the
	// cost of placing an order and the cost of holding a unit in stock.
	eoq, err := inventory.EOQ(inventory.EOQInput{
		AnnualDemand:       120 * 365,
		OrderingCost:       75,
		HoldingCostPerUnit: 4,
	})
	if err != nil {
		log.Fatalf("EOQ: %v", err)
	}

	fmt.Printf("\nEconomic order quantity: %.0f units\n", eoq)
}
