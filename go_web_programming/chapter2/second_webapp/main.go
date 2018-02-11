package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"templates/layout.html",
		"templates/navbar.html",
		"templates/index.html",
	}

	templates := template.Must(template.ParseFiles(files...))
	if threads, err := data.Threads(); err != nil {
		templates.ExecuteTemplate(w, "layout", threads)
	}
}

func main() {
	mux := http.NewServeMux()

	files := http.FileServer(http.Dir("/public")) //提供静态文件服务
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	//当http://domain.com/static/a.html -> http://domain.com/public/a.html

	mux.HandleFunc("/", index)
	mux.HandleFunc("/err", err)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/signup_account", signupAccount)
	mux.HandleFunc("/authenticate", authenticate)

	mux.HandleFunc("/thread/new",newThread)
	mux.HandleFunc("/thread/create",createThread)
	mux.HandleFunc("/thread/post",postThread)
	mux.HandleFunc("/thread/read",readThread)
	
	server := &http.Server{
		Addr:    "0.0.0.0:8888",
		Handler: mux,
	}
	fmt.Println("server start at 0.0.0.0:8888")
	server.ListenAndServe()

}
