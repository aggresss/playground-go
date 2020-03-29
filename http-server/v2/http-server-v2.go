package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", &myHandler{})
	mux.HandleFunc("/bye", sayBye)

	log.Println("Starting v2 httpserver")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is version 2"))
}
func sayBye(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("bye bye ,this is v2 httpServer"))
}
