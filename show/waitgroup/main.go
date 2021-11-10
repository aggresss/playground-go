package main

import "sync"

func main() {

}

type subscriber struct{}

func (s *subscriber) consume(b []byte) {
	return
}

var subscribers []*subscriber

func process(b []byte) error {
	wg := &sync.WaitGroup{}
	for _, s := range subscribers {
		wg.Add(1)
		go func(wg *sync.WaitGroup, s *subscriber) {
			defer wg.Done()
			s.consume(b)
		}(wg, s)
	}
	wg.Wait()

	return nil
}
