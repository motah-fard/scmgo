package parallel

import "testing"

// FuzzMapWithWorkers exercises arbitrary combinations of item count and
// worker count -- including workers > n, workers <= 0, and n = 0 -- to
// catch any off-by-one in the chunk-splitting logic. Run with -race
// (this project's CI fuzz job does) since the property that actually
// matters here is "no two goroutines ever write the same index", not
// just "no panic".
func FuzzMapWithWorkers(f *testing.F) {
	f.Add(0, 4)
	f.Add(1, 4)
	f.Add(7, 3)
	f.Add(100, 1)
	f.Add(5, 100)
	f.Fuzz(func(t *testing.T, n, workers int) {
		if n < 0 || n > 1_000_000 {
			return
		}

		items := make([]int, n)
		for i := range items {
			items[i] = i
		}

		got := MapWithWorkers(items, workers, func(index int, v int) int {
			if v != index {
				t.Fatalf("fn received item %d at index %d -- items and indices are misaligned", v, index)
			}
			return v * 2
		})

		if len(got) != n {
			t.Fatalf("result length = %d, want %d", len(got), n)
		}
		for i, v := range got {
			if v != i*2 {
				t.Fatalf("result[%d] = %d, want %d", i, v, i*2)
			}
		}
	})
}
