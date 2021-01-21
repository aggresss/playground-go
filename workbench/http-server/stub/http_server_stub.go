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
	log.Printf("[Unix-Timestamp]: %d\n", time.Now().UnixNano()/1e6)
	fmt.Printf("RemoteAddr: %s\n", r.RemoteAddr)
	fmt.Printf("Method: %s\nURI: %s\nContentLength: %d\n", r.Method, r.RequestURI, r.ContentLength)
	// header
	fmt.Printf("Header:\n")
	for k, v := range r.Header {
		fmt.Printf("\t%q: %q\n", k, v)
	}
	// body
	body, err := ioutil.ReadAll(r.Body)
	if err == nil {
		fmt.Printf("Body:\n%s\n", string(body))
	}
	// response
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(r.RemoteAddr))
	w.Write([]byte("\r\n"))
	// EOF
	fmt.Println("==========================================")
}
