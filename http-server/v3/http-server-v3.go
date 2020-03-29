package main

import (
	"log"
	"net/http"
)

func main() {
	mux := new(myHandler)
	log.Println("Starting httpserver v3")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/foo":
		w.Write([]byte("foo"))
	default:
		w.Write([]byte("httpserver v3"))
	}
}
