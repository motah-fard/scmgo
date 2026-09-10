package abc

import (
	"strconv"
	"testing"
)

func BenchmarkClassify(b *testing.B) {
	items := make([]Item, 1000)
	for i := range items {
		items[i] = Item{ID: "sku", Value: float64(i + 1)}
	}
	in := ClassifyInput{Items: items, AThreshold: 0.80, BThreshold: 0.95}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Classify(in)
	}
}

// BenchmarkClassifyVariabilitySizes and
// BenchmarkClassifyVariabilityConcurrentSizes exist to find the
// crossover point where ClassifyVariabilityConcurrent's goroutine
// overhead stops costing more than it saves. Run both and compare:
//
//	go test ./abc/... -bench 'ClassifyVariability.*Sizes' -benchtime=200x
//
// On this project's development hardware, the crossover landed in the
// tens of thousands of items; below that, ClassifyVariability (the
// plain sequential version) was faster.
func BenchmarkClassifyVariabilitySizes(b *testing.B) {
	for _, n := range []int{100, 1_000, 10_000, 100_000} {
		items := variabilityItemsForBenchmark(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = ClassifyVariability(ClassifyVariabilityInput{
					Items: items, XThreshold: 0.10, YThreshold: 0.50,
				})
			}
		})
	}
}

func BenchmarkClassifyVariabilityConcurrentSizes(b *testing.B) {
	for _, n := range []int{100, 1_000, 10_000, 100_000} {
		items := variabilityItemsForBenchmark(n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = ClassifyVariabilityConcurrent(ClassifyVariabilityInput{
					Items: items, XThreshold: 0.10, YThreshold: 0.50,
				})
			}
		})
	}
}

func variabilityItemsForBenchmark(n int) []VariabilityItem {
	items := make([]VariabilityItem, n)
	for i := range items {
		items[i] = VariabilityItem{ID: "sku", MeanDemand: 50, StdDevDemand: float64(i % 40)}
	}
	return items
}
