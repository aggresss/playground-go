package main

import "fmt"

// definition

type Printer interface {
	Print()
}

type PrinterFunc func()

func (p PrinterFunc) Print() {
	p()
}

type interceptor interface {
	BindPrinter(Printer) Printer
}

type Chain struct {
	interceptors []interceptor
}

func (i *Chain) BindPrinter(p Printer) Printer {
	for _, s := range i.interceptors {
		p = s.BindPrinter(p)
	}
	return p
}

func NewChain(s []interceptor) interceptor {
	return &Chain{interceptors: s}
}

// implement

type rhoInterceptor struct{}

func (r *rhoInterceptor) BindPrinter(p Printer) Printer {
	return PrinterFunc(func() {
		fmt.Println("rho")
		p.Print()
	})
}

type sigmaInterceptor struct{}

func (r *sigmaInterceptor) BindPrinter(p Printer) Printer {
	return PrinterFunc(func() {
		fmt.Println("sigma")
		p.Print()
	})
}

type tauInterceptor struct{}

func (r *tauInterceptor) BindPrinter(p Printer) Printer {
	return PrinterFunc(func() {
		fmt.Println("tau")
		p.Print()
	})
}

// run
func main() {
	s := NewChain([]interceptor{
		&rhoInterceptor{},
		&sigmaInterceptor{},
		&tauInterceptor{},
	})

	s.BindPrinter(PrinterFunc(func() { fmt.Println("upsilon") })).Print()
}

/*
	tau
	sigma
	rho
	upsilon
*/
