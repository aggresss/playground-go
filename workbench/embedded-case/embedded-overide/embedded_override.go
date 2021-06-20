package main

import "fmt"

type rho interface {
	sigma()
	tau()
}

type iota struct{}

func (i *iota) sigma() {
	fmt.Println("siggma by iota")
}

type kappa struct{}

func (k *kappa) tau() {
	fmt.Println("tau by kappa")
}

type lambda struct {
	iota
	kappa
}

func (l *lambda) tau() {
	fmt.Println("tau by lambda")
}

func main() {

	func(r rho) {
		r.sigma()
		r.tau()
	}(&lambda{})

	l := &lambda{}
	l.kappa.tau()
	l.tau()
}
