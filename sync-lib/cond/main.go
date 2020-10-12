package main

import (
	"fmt"
	"sync"
	"time"
)

// Cond is a condition/signal for announcing the occurance of an event
// Here it announces whenever an item has been removed from queue
func main() {
	// Condition using a standard sync.Mutex as the Locker
	c := sync.NewCond(&sync.Mutex{})
	// A queue of 10 items waiting to be filled
	queue := make([]interface{}, 0, 10)

	removeFromQueue := func(delay time.Duration) {
		time.Sleep(delay)
		c.L.Lock()
		// Dequeueing item by assigning the head of the slice to the next item
		queue = queue[1:]
		fmt.Println("Removed from queue")
		c.L.Unlock()
		// Let the goroutine waiting that something has occured
		c.Signal()
	}

	for i := 0; i < 10; i++ {
		c.L.Lock()
		// We know that an item has been removed from queue
		// but we dont know how many items the queue has
		// So we wait until the queue has 2 items
		// to add another item in the queue
		for len(queue) == 2 {
			// Suspend main goroutine until the "removed from queue" signal has been sent
			c.Wait()
		}
		fmt.Println("Adding to queue")
		queue = append(queue, struct{}{})
		// Dequeue an item after 1 second in goroutine
		go removeFromQueue(1 * time.Second)
		c.L.Unlock()
	}
}
