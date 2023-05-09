package main

import "time"

type transport struct {
	closeChan chan struct{}
}

func newTransport() *transport {
	t := &transport{
		closeChan: make(chan struct{}),
	}
	go t.monitor()
	return t
}

func (t *transport) monitor() {
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-t.closeChan:
			ticker.Stop()
			return
		case <-ticker.C:
			// TODO
		}
	}
}

func main() {
	newTransport()
}
