package main

import (
	"fmt"
	"net/http"
)

func set_cookie(w http.ResponseWriter,r *http.Request){
	cookie1:=http.Cookie{
		Name:"cookie",
		Value:"cookie1's value",
		HttpOnly:true,
	}
	cookie2:=http.Cookie{
		Name:"cookie2",
		Value:"cookie2's value",
		HttpOnly:true,
	}

	w.Header().Set("Set-Cookie",cookie1.String())
	w.Header().Add("Set-Cookie",cookie2.String())
	
	//上面两行等价于下面两行
	// http.SetCookie(w,&cookie1)
	// http.SetCookie(w,&cookie2)
}


func get_cookie(w http.ResponseWriter,r *http.Request){
	cookie,err:=r.Cookie("cookie")
	if err!=nil{
		fmt.Fprint(w,"can't get cookie")
	}
	cookies:=r.Cookies()
	fmt.Fprint(w,cookie)
	fmt.Fprintln(w)
	fmt.Fprint(w,cookies)
}

func main(){
	server:=http.Server{
		Addr:"127.0.0.1:8000",
	}

	http.HandleFunc("/cookie/",set_cookie)
	http.HandleFunc("/getcookie/",get_cookie)
	server.ListenAndServe()
}