
CAS (compare and swap)

- https://mp.weixin.qq.com/s/W3aofshF2qydm759-dvZsA

```
go test -run='^$' -bench=. -count=1 -benchtime=2s
```

```
goos: darwin
goarch: arm64
pkg: github.com/aggresss/playground-go/workbench/cas-case
BenchmarkStack/*cascase.LockFreeStack-8                  6901698               322.6 ns/op
BenchmarkStack/*cascase.MutexStack-8                    15968930               183.7 ns/op
PASS
ok      github.com/aggresss/playground-go/workbench/cas-case    7.549s
```