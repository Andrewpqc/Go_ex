package main

import (
	// "fmt"
	"html/template"
	"net/http"
)

func useTemplate(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmp1.html")
	t.Execute(w, "hello,world")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8000",
	}

	http.HandleFunc("/template/", useTemplate)

	server.ListenAndServe()

}
