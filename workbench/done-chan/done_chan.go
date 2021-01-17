package main

// Auto reset event.
type Event chan struct{}

func NewEvent() Event {
	return make(chan struct{}, 1)
}

func (e Event) Set() {
	select {
	case e <- struct{}{}:
	default:
	}
}

func (e Event) R() EventR {
	return EventR((chan struct{})(e))
}

// You can determine whether event setted through EventR.
type EventR <-chan struct{}

// You can notify something done through DoneChan.SetDone().
type DoneChan chan struct{}

func NewDoneChan() DoneChan {
	return make(chan struct{})
}

func (d DoneChan) SetDone() {
	defer func() { recover() }()
	select {
	case <-d:
	default:
		close(d)
	}
}

func (d DoneChan) R() DoneChanR {
	return DoneChanR((chan struct{})(d))
}

// You can determine whether something done through DoneChanR.Done().
type DoneChanR <-chan struct{}

func (d DoneChanR) Done() bool {
	select {
	case <-d:
		return true
	default:
		return false
	}
}
