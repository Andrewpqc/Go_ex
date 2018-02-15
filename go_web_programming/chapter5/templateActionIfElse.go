package main

import (
	"html/template"
	"net/http"
)

func IfElse(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("ifelse.html"))
	t.Execute(w, 0)
	// t.Execute(w,1)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8000",
	}
	http.HandleFunc("/ifelse/", IfElse)
	server.ListenAndServe()
}
