

## Table-driven Tests

https://github.com/golang/go/wiki/TableDrivenTests

## Operation

```shell
go test ./...
go test -benchmem -bench=.
```

## Benchmark Illustration
```
go test -benchmem -bench=.
goos: darwin
goarch: amd64
pkg: github.com/aggresss/playground-go/test/unit

benchmark                                   iter              time/iter    bytes alloc            allocs
---------                                   ----              ---------    -----------            ------
BenchmarkAppendFloat/10-4              184696033             6.35 ns/op         0 B/op       0 allocs/op
BenchmarkAppendFloat/100-4              18421809             66.1 ns/op         0 B/op       0 allocs/op
BenchmarkAppendFloat/1000-4              1450104              822 ns/op         0 B/op       0 allocs/op
BenchmarkAppendFloat/10000-4              141331             8715 ns/op         0 B/op       0 allocs/op

PASS
ok      github.com/aggresss/playground-go/test/unit     6.693s
```