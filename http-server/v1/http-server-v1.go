package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("httpserver v1"))
	})
	http.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("foo"))
	})
	log.Println("Starting httpserver v1 ...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
