package main

import (
	"net/http"
)

func handleRequest(w http.ResponseWriter,r *http.Request){
	var err error
	switch r.Method{
	case "GET":
		err=handleGet(w,r)
	case "POST":
		err=handlePost(w,r)
	case "PUT":
		err=handlePut(w,r)
	case "DELETE":
		err=handleDelete(w,r)
	}
	if err!=nil{
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}
}

func handleGet(w http.ResponseWriter,r *http.Request)(err error){
	w.Write([]byte("GET"))
	w.WriteHeader(200)
	return
}

func handlePost(w http.ResponseWriter,r *http.Request)(err error){
	w.Write([]byte("POST"))
	w.WriteHeader(200)
	return
}

func handlePut(w http.ResponseWriter,r *http.Request)(err error){
	w.Write([]byte("PUT"))
	w.WriteHeader(200)
	return
}

func handleDelete(w http.ResponseWriter,r *http.Request)(err error){
	w.Write([]byte("DELETE"))
	w.WriteHeader(200)
	return
}

func main(){
	server:=http.Server{
		Addr:"127.0.0.1:8000",
	}

	http.HandleFunc("/post/",handleRequest)
	server.ListenAndServe()

}