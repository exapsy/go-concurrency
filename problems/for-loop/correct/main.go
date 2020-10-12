package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)

		// salutation is now an explicit variable with a separate address
		// that shadows the for-loop's salutation variable
		go func(salutation string) {
			defer wg.Done()
			fmt.Println(salutation)
		}(salutation)
	}
	wg.Wait()
}
