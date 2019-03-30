// expose prometheus metrics
//	> homero
//	> clock covilha
//
// Nice to have: go test

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

var tmpl = `<html>
<head>
    <title>Hello</title>
</head>
<body>
    <img src="https://i.ytimg.com/vi/yWy2jFg6S1s/hqdefault.jpg"></img>
</body>
</html>
`

func handlerHomero(w http.ResponseWriter, r *http.Request) {
	t := template.New("main")
	t, _ = t.Parse(tmpl)
	log.Println(r.RemoteAddr, r.RequestURI[len("/"):])
	t.Execute(w, r.RequestURI[len("/"):])
}

func handlerClock(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", time.Now().UTC().Format("15:04:05\n"))

}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/homerosimpson", handlerHomero)
	http.HandleFunc("/clock", handlerClock)
	log.Printf("Started")
	server.ListenAndServe()
}
