package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"sync"
	"time"
)

type pkt struct {
	seq int
	b   []byte
}

type queue struct {
	pkts []pkt
	lock sync.RWMutex
}

func main() {
	q := &queue{}

	go func() {
		cnt := 0
		for {
			select {
			case <-time.After(time.Millisecond * 10):
				q.lock.Lock()
				q.pkts = append(q.pkts, pkt{
					seq: cnt,
					b:   make([]byte, 256*1024, 256*1024),
				})
				cnt++
				q.lock.Unlock()
			}
		}

	}()

	go func() {
		for {
			select {
			case <-time.After(time.Millisecond * 10):
				q.lock.Lock()
				if len(q.pkts) > 0 {
					q.pkts = q.pkts[1:]
				}
				q.lock.Unlock()
			}
		}
	}()

	go func() {
		for {
			select {
			case <-time.After(time.Second):
				m := &runtime.MemStats{}
				runtime.ReadMemStats(m)
				fmt.Println("Alloc", m.Alloc)
			}
		}
	}()

	log.Fatal(http.ListenAndServe("0.0.0.0:8012", nil))

}
