package main

import (
	"bufio"
	"container/ring"
	"context"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	timefly "github.com/aggresss/playground-go/workbench/protobuf-case/time-fly"
)

const (
	laddr = "0.0.0.0:7071"
)

func main() {
	var (
		mutex sync.RWMutex
	)
	ctx, cancel := context.WithCancel(context.TODO())
	pool := &sync.Pool{
		New: func() interface{} {
			return make([]byte, 1500)
		},
	}
	samples := ring.New(100)
	listener, err := net.Listen("tcp", laddr)
	if err != nil {
		fmt.Println("listen error:", err)
		return
	}
	handConn := func(ctx context.Context, conn net.Conn) {
		buff := bufio.NewReaderSize(conn, 1024*64)
		var (
			headerSize uint64
			msg        timefly.Msg
		)
		for {
			headerSize, err = binary.ReadUvarint(buff)
			if err != nil {
				return
			}
			b := pool.Get().([]byte)
			if _, err := io.ReadFull(buff, b[:headerSize]); err != nil {
				return
			}
			if err = msg.Unmarshal(b[:headerSize]); err != nil {
				return
			}

			mutex.Lock()
			samples.Value = time.Now().UnixNano() - msg.GetUnixnano()
			samples = samples.Next()
			mutex.Unlock()

			select {
			case <-ctx.Done():
				return
			default:
			}
		}
	}
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println("accept error", err)
				return
			}
			go handConn(ctx, conn)
		}
	}()

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-time.After(time.Second):
				var sum, count int64
				mutex.RLock()
				samples.Do(func(i interface{}) {
					if i != nil {
						s := i.(int64)
						if s != 0 {
							sum += s
							count++
						}
					}
				})
				mutex.RUnlock()
				if count != 0 {
					fmt.Printf("[%s] time-fly: %d ns\r", time.Now().String(), sum/count)
				}
			}
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	select {
	case <-ch:
		cancel()
	}

}
