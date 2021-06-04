## gctrace

### How to use

Use godoc as sample application
```
go get -v -u golang.org/x/tools/cmd/godoc
```

```
export GODEBUG='gctrace=1'
unset GODEBUG
```

```
GODEBUG='gctrace=1' ./xxx
GODEBUG='gctrace=1' go run ./...
```

```
gc 9 @0.491s 3%: 0.033+17+0.014 ms clock, 0.13+0.48/17/6.4+0.056 ms cpu, 60->60->35 MB, 63 MB goal, 4 P
```

```
gctrace: setting gctrace=1 causes the garbage collector to emit a single line to standard
error at each collection, summarizing the amount of memory collected and the
length of the pause. The format of this line is subject to change.
Currently, it is:
	gc # @#s #%: #+#+# ms clock, #+#/#/#+# ms cpu, #->#-># MB, # MB goal, # P
where the fields are as follows:
	gc #        the GC number, incremented at each GC
	@#s         time in seconds since program start
	#%          percentage of time spent in GC since program start
	#+...+#     wall-clock/CPU times for the phases of the GC
	#->#-># MB  heap size at GC start, at GC end, and live heap
	# MB goal   goal heap size
	# P         number of processors used
The phases are stop-the-world (STW) sweep termination, concurrent
mark and scan, and STW mark termination. The CPU times
for mark/scan are broken down in to assist time (GC performed in
line with allocation), background GC time, and idle GC time.
If the line ends with "(forced)", this GC was forced by a
runtime.GC() call.
```

- https://pkg.go.dev/runtime#hdr-Environment_Variables

### Theory

GC流程
- Sweep Termination: 对未清扫的span进行清扫, 只有上一轮的GC的清扫工作完成才可以开始新一轮的GC
- Mark: 扫描所有根对象, 和根对象可以到达的所有对象, 标记它们不被回收
- Mark Termination: 完成标记工作, 重新扫描部分根对象(要求STW)
- Sweep: 按标记结果清扫span

Tgc = Tseq + Tmark + Tsweep( T 表示 time)
- Tseq 表示是停止用户的 goroutine 和做一些准备活动（通常很小）需要的时间
- Tmark 是堆标记时间，标记发生在所有用户 goroutine 停止时，因此可以显著地影响处理的延迟
- Tsweep 是堆清除时间，清除通常与正常的程序运行同时发生，所以对延迟来说是不太关键的

`0.065+0.75+0.003 ms clock` The phases are
- stop-the-world (STW) sweep termination // STWSclock
- concurrent mark and scan // MASclock
- STW mark termination // STWMclock

### Visualization

```
go get -u -v github.com/davecheney/gcvis

tail -f xxx.log | gcvis
GODEBUG=gctrace=1 godoc -index -http=:6060 2>&1 | gcvis
```

- https://dave.cheney.net/2014/07/11/visualising-the-go-garbage-collector
