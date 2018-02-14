package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// func upload(w http.ResponseWriter,r *http.Request){
// 	r.ParseMultipartForm(1024)
// 	fileHeader:=r.MultipartForm.File["upload"][0]
// 	file,err:=fileHeader.Open()
// 	if err==nil{
// 		data,err:=ioutil.ReadAll(file)
// 		if err==nil{
// 			fmt.Fprint(w,string(data))
// 		}
// 	}
// }

func upload2(w http.ResponseWriter,r *http.Request){
	file,_,err:=r.FormFile("upload")
	if err==nil{
		data,err:=ioutil.ReadAll(file)
		if err==nil{
			fmt.Fprint(w,string(data))
		}
	}
}


func main(){
	server:=http.Server{
		Addr:"127.0.0.1:8000",
	}
	http.HandleFunc("/upload/",upload2)
	server.ListenAndServe()
}