package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", &myHandler{})
	mux.HandleFunc("/bye", sayBye)

	server := &http.Server{
		Addr:         ":8080",
		WriteTimeout: time.Second * 3, //设置3秒的写超时
		Handler:      mux,
	}
	log.Println("Starting v3 httpserver")
	log.Fatal(server.ListenAndServe())
}

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is version 3"))
}

func sayBye(w http.ResponseWriter, r *http.Request) {
	// 睡眠4秒  上面配置了3秒写超时，所以访问 “/bye“路由会出现没有响应的现象
	time.Sleep(4 * time.Second)
	w.Write([]byte("bye bye ,this is v3 httpServer"))
}
