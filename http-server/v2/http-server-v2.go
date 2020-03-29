package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", &myHandler1{})
	mux.Handle("/foo", &myHandler2{})

	log.Println("Starting httpserver v2")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

type myHandler1 struct{}

func (*myHandler1) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("httpserver v2"))
}

type myHandler2 struct{}

func (*myHandler2) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("foo"))
}
