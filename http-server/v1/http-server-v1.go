package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("httpserver v1"))
	})
	http.HandleFunc("/bye", sayBye)
	log.Println("Starting v1 server ...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func sayBye(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("bye bye ,this is v1 httpServer"))
}
