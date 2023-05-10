package cascase

import "sync"

type MutexStack struct {
	v  []interface{}
	mu sync.Mutex
}

func NewMutexStack() *MutexStack {
	return &MutexStack{v: make([]interface{}, 0)}
}

func (s *MutexStack) Push(v interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.v = append(s.v, v)
}

func (s *MutexStack) Pop() (v interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.v) > 0 {
		v = s.v[len(s.v)-1]
		s.v = s.v[:len(s.v)-1]
	}
	return v
}
