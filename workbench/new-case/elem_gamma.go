package new

// ElemGammaOptions stores options for ElemGamma.
type ElemGammaOptions struct {
	Foo bool
	Bar int
	Zoo string
}

// ElemGammaOption configures a ElemGammaOptions.
type ElemGammaOption interface {
	ApplyToElemGammaOptions(opts *ElemGammaOptions) error
}

// ElemGammaOptionFn configures a ElemGammaOptions.
type ElemGammaOptionFn func(*ElemGammaOptions) error

// ApplyToElemGammaOptions implements ElemGammaOption.
func (o ElemGammaOptionFn) ApplyToElemGammaOptions(opts *ElemGammaOptions) error {
	return o(opts)
}

// WithFoo sets Foo.
func WithFoo(f bool) ElemGammaOptionFn {
	return func(o *ElemGammaOptions) error {
		o.Foo = f
		return nil
	}
}

// WithBar sets Bar.
func WithBar(b int) ElemGammaOptionFn {
	return func(o *ElemGammaOptions) error {
		o.Bar = b
		return nil
	}
}

// WithZoo sets Zoo.
func WithZoo(z string) ElemGammaOptionFn {
	return func(o *ElemGammaOptions) error {
		o.Zoo = z
		return nil
	}
}

// ElemGamma struct
type ElemGamma struct {
	*ElemGammaOptions
	mapT map[string]int
}

// NewElemGamma returns a new ElemGamma
func NewElemGamma(opts ...ElemGammaOption) (*ElemGamma, error) {
	options := &ElemGammaOptions{}
	for _, o := range opts {
		if err := o.ApplyToElemGammaOptions(options); err != nil {
			return nil, err
		}
	}

	return &ElemGamma{
		ElemGammaOptions: options,
		mapT:             make(map[string]int),
	}, nil
}
