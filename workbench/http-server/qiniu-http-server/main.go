package main

import (
	"github.com/qiniu/http/restrpc"
)

// Service is a business prototype
type Service struct{}

// GetFoo method
func (*Service) GetFoo(env *restrpc.Env) (s string, err error) {
	s = "foo"
	return
}

func main() {
	svr := new(Service)
	router := restrpc.Router{}
	router.ListenAndServe(":8080", svr)
}
