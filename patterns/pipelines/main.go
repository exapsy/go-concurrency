package main

import (
	"fmt"
	"time"
)

func main() {
	generator := func(done <-chan interface{}, arr ...int) <-chan int {
		outputStream := make(chan int, len(arr))
		go func() {
			defer close(outputStream)
			for _, i := range arr {
				outputStream <- i
			}
		}()
		return outputStream
	}
	add := func(done <-chan interface{}, adder int, inputStream <-chan int) <-chan int {
		outputStream := make(chan int)

		go func() {
			defer close(outputStream)
			for i := range inputStream {
				select {
				case <-done:
					return
				case outputStream <- i + adder:
				}
			}
		}()

		return outputStream
	}
	multiply := func(done <-chan interface{}, multiplier int, inputStream <-chan int) <-chan int {
		outputStream := make(chan int)

		go func() {
			defer close(outputStream)
			for i := range inputStream {
				select {
				case <-done:
					return
				case outputStream <- i * multiplier:
				}
			}
		}()

		return outputStream
	}

	done := make(chan interface{})
	defer close(done)

	go func() {
		time.Sleep(1 * time.Millisecond)
		close(done)
	}()

	inputStream := generator(done, 1, 2, 3, 4, 5)
	for output := range multiply(done, 2, add(done, 5, inputStream)) {
		fmt.Println(output)
	}
}
