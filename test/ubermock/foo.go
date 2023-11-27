package ubermock

type foo interface {
	Bar(x int) int
}

func SUT(f foo) {
	if f.Bar(99) != 101 {
		panic("unexpected return value")
	}
}
