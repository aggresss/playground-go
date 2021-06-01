package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func Start(addr string) (*http.Server, error) {
	svc := &http.Server{
		Addr:           addr,
		Handler:        http.DefaultServeMux,
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		fmt.Println(svc.ListenAndServe().Error())
	}()

	return svc, nil
}
