package main

import (
	"fmt"
	"runtime"
	"sync"
)

// The point to prove is that go routines should be treated as a free resource
// The amount of memory they consume, even when there are thousands of them, is minimal
func main() {
	memConsumed := func() uint64 {
		runtime.GC()
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		return s.Sys
	}

	var c <-chan interface{}
	var wg sync.WaitGroup

	// Go routine that won't exit until the process is finished
	noop := func() { wg.Done(); <-c }

	// Number of goroutines to create
	const numGoroutines = 1e4
	wg.Add(numGoroutines)

	// Measure the amount of memory consumed before creating our goroutines
	before := memConsumed()
	for i := numGoroutines; i > 0; i-- {
		go noop()
	}
	wg.Wait()

	// Amount of consumed memory after creating the goroutines
	after := memConsumed()
	fmt.Printf("%.3fkb\n", float64(after-before)/numGoroutines/1000)
}
