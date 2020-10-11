package main

import (
	"sync"
	"testing"
)

// BenchmarkContextSwitch benchmarks if context switching (eg. channel communication)
// is expensive to be done between goroutines
// It's considered expensive to be done with OS context switching, but
// goroutines context-switch handles it 92% better than OS context switch
func BenchmarkContextSwitch(b *testing.B) {
	var wg sync.WaitGroup
	begin := make(chan struct{})
	c := make(chan struct{})

	var token struct{}
	sender := func() {
		defer wg.Done()
		// Wait until told to begin
		// so the whole setup wont factor into the measurement of context switching
		<-begin
		for i := 0; i < b.N; i++ {
			// A struct{}{} (token) is an empty struct, and takes up no memory
			// Thus, we're only measuring the time it takes to signal a message (to receiver)
			c <- token
		}
	}
	receiver := func() {
		defer wg.Done()
		// Wait until told to begin
		// so the whole setup wont factor into the measurement of context switching
		<-begin
		for i := 0; i < b.N; i++ {
			// Here we receive the message but do nothing with it
			<-c
		}
	}
	wg.Add(2)
	go sender()
	go receiver()
	// Begin performance timer
	b.StartTimer()
	// Tell the two goroutines to begin
	close(begin)
	wg.Wait()
}
