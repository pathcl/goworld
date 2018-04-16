package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	datetime := time.Now()
	fmt.Fprintf(w, "Go world %s! time now %s\n version: 0.0.3", r.URL.Path[1:], datetime)
	log.Println(r.RemoteAddr, r.Method, r.URL)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
