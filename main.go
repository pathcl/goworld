// display homero
// display clock
// expose prometheus metrics
//	> homero
//	> clock covilha
//
// Nice to have: go test

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path[1:] == "covilha" {
		fmt.Fprintf(w, "%s", time.Now().UTC().Format("15:04:05\n"))
	} else if r.URL.Path[1:] == "homerosimpson" {
		fmt.Fprintf(w, "%s", time.Now().Format("15:04:05\n"))
	} else {
		fmt.Fprintf(w, "you tried to access: /%s\n", r.URL.Path[1:])
	}
	log.Println(r.RemoteAddr, r.Method, r.URL)
}
func main() {
	http.HandleFunc("/", handler)
	log.Println("Starting...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
