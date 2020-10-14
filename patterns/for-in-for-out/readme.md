# For-in-for-out

**Forin** is a technique to use many workers for a pipeline.

It's useful when

- The pipeline does not have to wait/rely on values the stage had calculated before
- It takes a long time to run

## Calculating primes

Primes were calculated in the most inneficient way possible for the sake of measuring computational performance.

### With for-in

```go
numFinders := runtime.NumCPU()
randStream := randomGenerator(done, 500000000)
finders := make([]<-chan interface{}, numFinders)
for i := 0; i < numFinders; i++ {
  finders[i] = primeFinder(done, randStream)
}
for i := range take(done, 10, forin(done, finders...)) {
  fmt.Printf("Prime: %v\n", i)
}
```

#### Result

```md
$ go run ./patterns/for-in-for-out/with-forin/main.go  
Prime: 298498063
Prime: 411902059
Prime: 427131847
Prime: 439984019
Prime: 140954381
Prime: 208240433
Prime: 146203261
Prime: 106410691
Prime: 336122539
Prime: 474941317
Done after  15.202641093s
```

### Without for-in

```go
finder := primeFinder(done, randomGenerator(done, 500000000))
for i := range take(done, 10, finder) {
  fmt.Printf("Prime: %v\n", i)
}
```

#### Result

```md
$ go run ./patterns/for-in-for-out/without-forin/main.go
Prime: 427131847
Prime: 336122539
Prime: 208240433
Prime: 146203261
Prime: 106410691
Prime: 247278491
Prime: 460128143
Prime: 317455063
Prime: 183024707
Prime: 6933257
Done after  31.961739102s
```
