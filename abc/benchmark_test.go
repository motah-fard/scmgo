package abc

import "testing"

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
