package ubermock

//go:generate sh -c "go run go.uber.org/mock/mockgen -typed -package ubermock -destination mock_foo_test.go ubermock Foo"
type Foo = foo
