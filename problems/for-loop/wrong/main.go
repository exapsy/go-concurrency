package main

import (
	"fmt"
	"sync"
)

// Expected output:
// 		hello
// 		greetings
// 		good day

// Actual output:
// 		good day
// 		good day
// 		good day
func main() {
	var wg sync.WaitGroup

	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(salutation)
		}()
	}
	wg.Wait()
}
