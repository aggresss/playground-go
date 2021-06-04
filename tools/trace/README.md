## trace

```
go run trace.go 2> trace.out
```

```
wget -O trace.out http://127.0.0.1:6061/debug/pprof/trace?seconds=5
go tool trace trace.out
```

- https://golang.org/pkg/runtime/trace/
- https://golang.org/pkg/net/http/pprof/
- https://github.com/Raffo/go-pprof-trace-example
