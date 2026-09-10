package parallel

import (
	"strconv"
	"sync/atomic"
	"testing"
)

func sequentialMap[In, Out any](items []In, fn func(int, In) Out) []Out {
	results := make([]Out, len(items))
	for i, item := range items {
		results[i] = fn(i, item)
	}
	return results
}

func TestMapMatchesSequentialMap(t *testing.T) {
	t.Parallel()

	sizes := []int{0, 1, 2, 3, 7, 100, 10_000}

	for _, n := range sizes {
		n := n
		t.Run(strconv.Itoa(n), func(t *testing.T) {
			t.Parallel()

			items := make([]int, n)
			for i := range items {
				items[i] = i * 3
			}

			square := func(_ int, v int) int { return v * v }

			got := Map(items, square)
			want := sequentialMap(items, square)

			if len(got) != len(want) {
				t.Fatalf("length = %d, want %d", len(got), len(want))
			}
			for i := range got {
				if got[i] != want[i] {
					t.Fatalf("index %d: got %v, want %v", i, got[i], want[i])
				}
			}
		})
	}
}

func TestMapPassesCorrectIndex(t *testing.T) {
	t.Parallel()

	items := make([]string, 500)
	for i := range items {
		items[i] = "x"
	}

	got := Map(items, func(index int, _ string) int { return index })

	for i, v := range got {
		if v != i {
			t.Fatalf("result[%d] = %d, want %d (Map must report each item's own index, not a racy shared counter)", i, v, i)
		}
	}
}

func TestMapWithWorkersEdgeCases(t *testing.T) {
	t.Parallel()

	items := []int{1, 2, 3, 4, 5}
	double := func(_ int, v int) int { return v * 2 }
	want := []int{2, 4, 6, 8, 10}

	for _, workers := range []int{-1, 0, 1, 2, 5, 100} {
		workers := workers
		t.Run(strconv.Itoa(workers), func(t *testing.T) {
			t.Parallel()

			got := MapWithWorkers(items, workers, double)
			if len(got) != len(want) {
				t.Fatalf("workers=%d: length = %d, want %d", workers, len(got), len(want))
			}
			for i := range got {
				if got[i] != want[i] {
					t.Fatalf("workers=%d: index %d = %v, want %v", workers, i, got[i], want[i])
				}
			}
		})
	}
}

func TestMapEmptyInput(t *testing.T) {
	t.Parallel()

	got := Map([]int{}, func(_ int, v int) int { return v })
	if len(got) != 0 {
		t.Fatalf("Map on empty input = %v, want empty", got)
	}
}

// TestMapNoConcurrentWriteRace exercises the actual concurrency: many
// goroutines writing distinct indices of a shared slice with no explicit
// synchronization inside fn. This is meaningless without -race, but under
// -race it fails loudly if Map's index-range partitioning ever overlaps.
func TestMapNoConcurrentWriteRace(t *testing.T) {
	t.Parallel()

	const n = 50_000
	items := make([]int, n)
	for i := range items {
		items[i] = i
	}

	var calls atomic.Int64
	got := Map(items, func(index int, v int) int {
		calls.Add(1)
		return v + index
	})

	if calls.Load() != n {
		t.Fatalf("fn called %d times, want %d", calls.Load(), n)
	}
	for i, v := range got {
		if v != i+i {
			t.Fatalf("result[%d] = %d, want %d", i, v, i+i)
		}
	}
}
