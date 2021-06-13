package new

type ElemBetaOptions struct {
	Foo bool
	Bar int
	Zoo string
}

type ElemBeta struct {
	*ElemBetaOptions
	mapT map[string]int
}

func NewElemBeta(opts *ElemBetaOptions) *ElemBeta {
	return &ElemBeta{
		ElemBetaOptions: opts,
		mapT:            make(map[string]int),
	}
}
