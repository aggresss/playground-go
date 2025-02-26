package main

import (
	"fmt"
	"net/http"
)

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if originHdr := r.Header.Get("Origin"); originHdr != "" {
			w.Header().Set("Access-Control-Allow-Origin", originHdr)
			w.Header().Set("Access-Control-Allow-Credentials", "true")
		}

		if r.Method == http.MethodOptions {
			if reqMethodHdr := r.Header.Get("Access-Control-Request-Method"); reqMethodHdr != "" {
				w.Header().Set("Access-Control-Allow-Methods", reqMethodHdr)
			}
			if reqHeadersHdr := r.Header.Get("Access-Control-Request-Headers"); reqHeadersHdr != "" {
				w.Header().Set("Access-Control-Allow-Headers", reqHeadersHdr)
			}
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)
	corsMux := enableCORS(mux)
	http.ListenAndServe(":8080", corsMux)
}
