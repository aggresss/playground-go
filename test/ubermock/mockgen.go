//go:build gomock || generate

package ubermock

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -build_flags=\"-tags=gomock\" -package ubermock -destination mock_foo_test.go ubermock Foo"
type Foo = foo
