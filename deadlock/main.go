package main

import (
	"fmt"
	"sync"
	"time"
)

type value struct {
	mu    sync.Mutex
	value int
}

func main() {
	var wg sync.WaitGroup
	printSum := func(v1, v2 *value) {
		defer wg.Done()

		v1.mu.Lock()
		defer v1.mu.Unlock()

		time.Sleep(1 * time.Second) // Introduces deadlock
		v2.mu.Lock()
		defer v2.mu.Unlock()

		fmt.Printf("Sum: %v\n", v1.value+v2.value)
	}

	var a, b value
	wg.Add(2)
	// B will be locked loaded after 1 second of A
	go printSum(&a, &b)
	// A will be locked and loaded after 1 second of B
	go printSum(&b, &a)
	// So now b "belongs" to 2nd printSum, and a "belong" to 1st printSum
	// But both require both variables to continue
	// which makes them stuck in a deadlock
	wg.Wait()
}
