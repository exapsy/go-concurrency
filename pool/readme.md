# Pool

A pool is a cache of pre-assigned items. It's used to prevent assigning new variables when there is a huge amount of data, as assigning new items is costly and computationaly inneficient.

## How to benchmark

```
# Test without pool
cd ./without-pool
go test -benchtime=10s -bench=.
```

Should print out something similar to

```
pkg: github.com/exapsy/go-concurrency/pool/without-pool
BenchmarkNetworkRequest-12            10        1000369872 ns/op
PASS
ok      github.com/exapsy/go-concurrency/pool/without-pool      11.007s
```

The result is
**1e9 ns/op**

# Test with pool
cd ./with-pool
go test -benchtime=10s -bench=.
```

Should print out something similar to

```
pkg: github.com/exapsy/go-concurrency/pool/with-pool
BenchmarkNetworkRequest-12    
    2800           6219911 ns/op
PASS
ok      github.com/exapsy/go-concurrency/pool/with-pool 32.429s
```

The result is
**6.2e6 ns/op** ! Magnitudes faster than without using a pool.
Clearly having a pool that has memory allocated variables ready to use helps significantly in efficiency and performance.
