// Command concurrent-batch measures BuildPolicySummaryBatch against
// BuildPolicySummaryBatchConcurrent at a few batch sizes on your own
// machine, rather than asking you to trust a number in a doc comment.
//
// A formula this cheap doesn't automatically get faster from goroutines:
// below some size, scheduling overhead costs more than the parallelism
// saves. This prints the actual crossover on the hardware it runs on.
//
// Run it with:
//
//	go run ./examples/concurrent-batch
package main

import (
	"fmt"
	"time"

	"github.com/motah-fard/scmgo/inventory"
)

func main() {
	sizes := []int{100, 1_000, 10_000, 100_000}

	fmt.Printf("%10s %15s %15s %10s\n", "Batch size", "Sequential", "Concurrent", "Faster?")
	for _, n := range sizes {
		inputs := make([]inventory.PolicySummaryInput, n)
		for i := range inputs {
			inputs[i] = inventory.PolicySummaryInput{
				DailyDemand:      100,
				LeadTimeDays:     5,
				ReviewPeriodDays: 7,
				SafetyStockUnits: 50,
			}
		}

		seqStart := time.Now()
		_ = inventory.BuildPolicySummaryBatch(inputs)
		seqElapsed := time.Since(seqStart)

		concStart := time.Now()
		_ = inventory.BuildPolicySummaryBatchConcurrent(inputs)
		concElapsed := time.Since(concStart)

		winner := "sequential"
		if concElapsed < seqElapsed {
			winner = "concurrent"
		}

		fmt.Printf("%10d %15s %15s %10s\n", n, seqElapsed, concElapsed, winner)
	}

	fmt.Println("\nA single run of each size is noisy -- see")
	fmt.Println("BenchmarkBuildPolicySummaryBatchSizes / *ConcurrentSizes in the")
	fmt.Println("inventory package for a properly repeated measurement.")
}
