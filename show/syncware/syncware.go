package syncware

import (
	"sync"
	"sync/atomic"
)

const (
	stateBuilding = iota
	stateReady
	stateClosing
	stateClosed
)

type syncware struct {
	locker    sync.RWMutex
	buildChan chan struct{}
	closeChan chan struct{}
	refCount  int32
	closeFunc func()
	state     int
}

func (c *syncware) getState() int {
	c.locker.RLock()
	defer c.locker.RUnlock()
	return c.state
}

func (c *syncware) setState(state int) bool {
	c.locker.Lock()
	defer c.locker.Unlock()
	// state should be forward
	if state <= c.state {
		return false
	}
	// current state building should close buildChan
	if c.state == stateBuilding {
		close(c.buildChan)
	}
	// next state closed should close closeChan
	if state == stateClosed {
		close(c.closeChan)
	}
	c.state = state
	return true
}

func (c *syncware) addReference() {
	atomic.AddInt32(&c.refCount, 1)
}

func (c *syncware) delReference() {
	atomic.AddInt32(&c.refCount, -1)
	if atomic.LoadInt32(&c.refCount) <= 0 {
		c.close()
	}
}

func (c *syncware) waitReady() <-chan struct{} {
	return c.buildChan
}

func (c *syncware) waitClosed() <-chan struct{} {
	return c.closeChan
}

func (c *syncware) close() error {
	if !c.setState(stateClosing) {
		return nil
	}
	if c.closeFunc != nil {
		c.closeFunc()
	}
	c.setState(stateClosed)
	return nil
}

func newSyncware() *syncware {
	c := &syncware{
		buildChan: make(chan struct{}),
		closeChan: make(chan struct{}),
		state:     stateBuilding,
	}
	return c
}
