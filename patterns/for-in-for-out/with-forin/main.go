package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func main() {
	// For in is a stream multiplexer
	// broadcasts/compines many streams (channels) to one stream
	forin := func(
		done <-chan interface{},
		channels ...<-chan interface{},
	) <-chan interface{} {
		var wg sync.WaitGroup
		multiplexedStream := make(chan interface{})

		multiplex := func(channel <-chan interface{}) {
			defer wg.Done()
			for i := range channel {
				select {
				case <-done:
					return
				default:
					multiplexedStream <- i
				}
			}
		}

		wg.Add(len(channels))
		for _, channel := range channels {
			go multiplex(channel)
		}

		// Wait for all the reads to complete
		go func() {
			wg.Wait()
			close(multiplexedStream)
		}()

		return multiplexedStream
	}

	// Finds a prime starting from a big number
	// by dividing all the numbers below it
	// Done intentionally to be slow to demonstrate for-in's performance value
	// Obviously no optimizations should be done
	primeFinder := func(done <-chan interface{}, fromStream <-chan interface{}) <-chan interface{} {
		prime := make(chan interface{})

		go func() {
			defer close(prime)
			for {
				select {
				case <-done:
					return
				case from := <-fromStream:
				primeCandidateLoop:
					for primeCandidate := from.(int); primeCandidate > 3; primeCandidate-- {
						for divider := 2; divider < primeCandidate; divider++ {
							if (primeCandidate % divider) == 0 {
								continue primeCandidateLoop
							}
						}
						prime <- primeCandidate
						break primeCandidateLoop
					}
				}
			}
		}()

		return prime
	}

	randomGenerator := func(done <-chan interface{}, max int) <-chan interface{} {
		randomNumberStream := make(chan interface{})

		go func() {
			defer close(randomNumberStream)
			for {
				select {
				case <-done:
					return
				default:
					randomNumberStream <- rand.Intn(max)
				}
			}
		}()

		return randomNumberStream
	}

	take := func(done <-chan interface{}, times int, inputStream <-chan interface{}) <-chan interface{} {
		outputStream := make(chan interface{})

		go func() {
			defer close(outputStream)
			for i := 0; i < times; i++ {
				select {
				case <-done:
					return
				case i := <-inputStream:
					outputStream <- i
				}
			}
		}()

		return outputStream
	}

	done := make(chan interface{})
	defer close(done)

	start := time.Now()

	numFinders := runtime.NumCPU()
	randStream := randomGenerator(done, 500000000)
	finders := make([]<-chan interface{}, numFinders)
	for i := 0; i < numFinders; i++ {
		finders[i] = primeFinder(done, randStream)
	}
	for i := range take(done, 10, forin(done, finders...)) {
		fmt.Printf("Prime: %v\n", i)
	}

	fmt.Println("Done after ", time.Since(start))
}
