## pprof

proto/profiling

- CPU Profiling
- Memory Profiling
- Block Profiling
- Mutex Profiling


```
go tool pprof http://127.0.0.1:6061/debug/pprof/xxx

web
top

```

```
go tool pprof http://127.0.0.1:6061/debug/pprof/profile -seconds 30
go tool pprof -http=:8081 ~/pprof/pprof.samples.cpu.001.pb.gz
```

### Reference

- https://github.com/google/pprof
- https://golang.org/pkg/net/http/pprof/
- https://golang.org/pkg/runtime/pprof/
- https://segmentfault.com/a/1190000016412013
