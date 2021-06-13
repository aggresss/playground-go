package new

type ElemAlpha struct {
	foo  bool
	bar  int
	zoo  string
	mapT map[string]int
}

func NewElemAlpha(f bool, b int, z string) *ElemAlpha {
	return &ElemAlpha{
		foo:  f,
		bar:  b,
		zoo:  z,
		mapT: make(map[string]int),
	}
}
