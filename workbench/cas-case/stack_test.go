package cascase

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"testing"
	"time"
)

func BenchmarkStack(b *testing.B) {
	rand.Seed(time.Now().UnixNano())

	length := 1 << 12
	inputs := make([]int, length)
	for i := 0; i < length; i++ {
		inputs[i] = rand.Int()
	}

	ls, ms := NewLockFreeStack(), NewMutexStack()

	b.ResetTimer()

	for _, s := range [...]StackInterface{ls, ms} {
		b.Run(fmt.Sprintf("%T", s), func(b *testing.B) {
			var c int64
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					i := int(atomic.AddInt64(&c, 1)-1) % length
					v := inputs[i]
					if v >= 0 {
						s.Push(v)
					} else {
						s.Pop()
					}
				}
			})
		})
	}
}
