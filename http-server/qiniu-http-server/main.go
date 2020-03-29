package main

import (
	"net/http"

	"github.com/qiniu/http/restrpc"
)

// Service is a business prototype
type Service struct{}

// GetFoo method
func (p *Service) GetFoo(env *restrpc.Env) (s string, err error) {
	s = "foo"
	return
}

func main() {
	svr := new(Service)
	router := restrpc.Router{}
	http.ListenAndServe(":8080", router.Register(svr))
}
