## trace

```
go run trace.go 2> trace.out
```

```
curl -o trace.out "http://127.0.0.1:6061/debug/pprof/trace?seconds=30"
go tool trace trace.out
```

> 有多种方法可以描述停顿时间(STW)的分布，最简单的方法是统计停顿时间的变化，例如标准差或者图形表示等。更有效的方法包括最小赋值器使用率（minimum mutator utilization，MMU）和界限赋值器使用率（bounded mutator utilization，BMU）。MMU 和 BMU 都简明地展示了任意给定时间窗内赋值器占用的（最小）时间比例。x轴表示程序从开始到结束的整体执行时间，y轴表示赋值器占用的CPU时间比例（使用率）。MMU 和 BMU曲线不仅反映了垃圾回收过程占整个执行时间的比例（y轴截距，即曲线最右侧的点，代表了赋值器占用整体处理时间的比例，用100%减去该值即可得到回收过程的时间占比），也反应了垃圾回收的最长停顿时间（x轴截距，即赋值器CPU使用率为0%的***时间窗口）。一般来说，曲线在y轴上越高，表示赋值器的CPU占用率越高，在x轴上越靠左，表示垃圾回收的***停顿时间越小。MMU曲线反映了程序在任意时间窗（x）内赋值器的最小使用率（y），但是较大时间窗的MMU可能会比较小时间窗的更低，因此MMU曲线会出现下降。

- https://golang.org/pkg/runtime/trace/
- https://golang.org/pkg/net/http/pprof/
- https://github.com/Raffo/go-pprof-trace-example
