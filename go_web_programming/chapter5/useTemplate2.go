package main

import (
	"html/template"
	"net/http"
)

func useTemplate1(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("tmp1.html", "tmp2.html"))
	t.Execute(w, "hello,world")
}

func useTemplate2(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("tmp1.html", "tmp2.html"))
	t.ExecuteTemplate(w, "tmp2.html", "hello,world")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8000",
	}

	http.HandleFunc("/tmp1/", useTemplate1)
	http.HandleFunc("/tmp2/", useTemplate2)

	server.ListenAndServe()
}
