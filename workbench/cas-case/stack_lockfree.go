package cascase

import (
	"sync/atomic"
	"unsafe"
)

type directItem struct {
	next unsafe.Pointer
	v    interface{}
}

type LockFreeStack struct {
	top unsafe.Pointer
	len uint64
}

func NewLockFreeStack() *LockFreeStack {
	return &LockFreeStack{}
}

func (s *LockFreeStack) Push(v interface{}) {
	item := directItem{v: v}
	var top unsafe.Pointer

	for {
		top = atomic.LoadPointer(&s.top)
		item.next = top
		if atomic.CompareAndSwapPointer(&s.top, top, unsafe.Pointer(&item)) {
			atomic.AddUint64(&s.len, 1)
			return
		}
	}
}

func (s *LockFreeStack) Pop() (v interface{}) {
	var top, next unsafe.Pointer

	for {
		top = atomic.LoadPointer(&s.top)
		if top != nil {
			v = (*directItem)(top)
			next = atomic.LoadPointer(&v.(*directItem).next)
			if atomic.CompareAndSwapPointer(&s.top, top, next) {
				atomic.AddUint64(&s.len, ^uint64(0))
			}
		}
		return v
	}
}
