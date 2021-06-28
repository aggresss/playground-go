package main

import (
	"context"
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	timefly "github.com/aggresss/playground-go/workbench/protobuf-case/time-fly"
)

const (
	raddr = "127.0.0.1:7071"
)

func dial(raddr string) (conn net.Conn, err error) {
	conn, err = net.Dial("tcp", raddr)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func send(c net.Conn, p *sync.Pool) (err error) {
	m := timefly.Msg{
		Unixnano: time.Now().UnixNano(),
	}
	b := p.Get().([]byte)
	defer p.Put(b)

	bodySize := m.Size()
	headerSize := binary.PutUvarint(b, uint64(bodySize))
	if _, err = m.MarshalTo(b[headerSize:]); err != nil {
		return err
	}
	if _, err := c.Write(b[:headerSize+bodySize]); err != nil {
		return err
	}

	return nil
}

func main() {
	var (
		conn net.Conn
		err  error
		pool *sync.Pool
	)
	pool = &sync.Pool{
		New: func() interface{} {
			return make([]byte, 1500)
		},
	}
	if conn, err = dial(raddr); err != nil {
		fmt.Println("dial error:", err)
		return
	}
	defer conn.Close()

	ctx, cancel := context.WithCancel(context.TODO())
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-time.After(time.Millisecond * 100):
				if conn != nil {
					send(conn, pool)
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
