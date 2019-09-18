// The server command is a minimal example server-side web app.
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var port = flag.Int("port", 8080, "Port on which server listens.")
var tmpl = template.Must(template.ParseFiles("template.html"))

type params struct {
	Greeting string
}

type app struct{}

func (a *app) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	greeting := "Hola, mundo!"
	query := r.URL.Query()
	greetings, hasG := query["g"]
	if hasG {
		greeting = greetings[0]
	}
	data := params{Greeting: greeting}
	log.Print(tmpl.Execute (w, data))
	log.Printf("%v",query)
}

func main() {
	flag.Parse()
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/favicon.ico", fs)
	http.Handle("/", &app{})
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port),nil))
//	http.ListenAndServe(":8080", nil)
}