package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	listenAddr = "0.0.0.0:8082"
)

func main() {
	mux := new(myHandler)
	log.Printf(`Starting httpserver stub, listen %s`, listenAddr)
	fmt.Println("==========================================")
	log.Fatal(http.ListenAndServe(listenAddr, mux))
}

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// skeleton
	log.Printf("[Unix-Timestamp]: %d\n", time.Now().Unix())
	fmt.Printf("Method: %s\nContentLength: %d\n", r.Method, r.ContentLength)
	// header
	fmt.Printf("Header:\n")
	for k, v := range r.Header {
		fmt.Printf("\tField: %q, Value: %q\n", k, v)
	}
	// body
	body, err := ioutil.ReadAll(r.Body)
	if err == nil {
		fmt.Printf("Body:\n%s\n", string(body))
	}
	// response
	w.WriteHeader(http.StatusOK)
	// EOF
	fmt.Println("==========================================")
}
