// Package parallel holds a single reusable concurrency primitive: a
// bounded-worker parallel map over a slice, used by the "Concurrent"
// batch variants across scmgo's public packages. It is internal because
// it is an implementation detail, not part of the library's API surface.
package parallel

import (
	"runtime"
	"sync"
)

// Map applies fn to every element of items and returns the results in
// the same order as items, using up to runtime.GOMAXPROCS(0) goroutines.
//
// Work is split into contiguous index ranges, one per goroutine (a
// "fan-out over index ranges" split, not one goroutine per item), so the
// goroutine count is bounded regardless of len(items). Each goroutine
// writes only to the indices in its own range of the pre-allocated
// results slice, so there is no shared mutable state and no lock is
// needed for that -- go test -race confirms this.
//
// If fn panics, the panic is not recovered: it propagates and crashes
// the process, exactly as an unrecovered panic in any other goroutine
// would. Map does not add panic safety on your behalf.
//
// Map is not free: spawning goroutines and splitting work has real
// overhead. For a cheap fn (a single arithmetic formula, say), a plain
// loop beats Map until len(items) reaches the thousands -- see the
// "Concurrent" benchmarks in each package that uses this for the actual
// crossover point on your hardware. Don't reach for this by default;
// reach for it once a sequential loop is measured to be the bottleneck.
func Map[In, Out any](items []In, fn func(index int, item In) Out) []Out {
	return MapWithWorkers(items, runtime.GOMAXPROCS(0), fn)
}

// MapWithWorkers is Map with an explicit worker count instead of
// runtime.GOMAXPROCS(0). workers <= 0 is treated as 1.
func MapWithWorkers[In, Out any](items []In, workers int, fn func(index int, item In) Out) []Out {
	n := len(items)
	results := make([]Out, n)
	if n == 0 {
		return results
	}

	if workers <= 0 {
		workers = 1
	}
	if workers > n {
		workers = n
	}

	chunkSize := (n + workers - 1) / workers

	var wg sync.WaitGroup
	for start := 0; start < n; start += chunkSize {
		end := start + chunkSize
		if end > n {
			end = n
		}

		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			for i := start; i < end; i++ {
				results[i] = fn(i, items[i])
			}
		}(start, end)
	}
	wg.Wait()

	return results
}
