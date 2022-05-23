package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	addr := ":8888"
	fmt.Printf("Listening on %s\n", addr)
	http.HandleFunc("/", greeting)
	http.ListenAndServe(addr, nil)
}

func greeting(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world, it's %v", time.Now().UTC())
}
