package main

import (
	"testing"
	"net/http/httptest"
	"net/http"
)

func TestHandleGet(t *testing.T){
	mux:=http.NewServeMux()
	mux.HandleFunc("/post/",handleRequest)

	writer:=httptest.NewRecorder()
	request,_:=http.NewRequest("GET","/post/1",nil)
	mux.ServeHTTP(writer,request)

	if writer.Code!=200{
		t.Errorf("response code error,want 200,got %d",writer.Code)
	}
	if writer.Body.String()!="GET"{
		t.Errorf("response body error,want 'GET',got %s",writer.Body.String())
	}
} 


func TestHandlePost(t *testing.T){
	mux:=http.NewServeMux()
	mux.HandleFunc("/post/",handleRequest)

	writer:=httptest.NewRecorder()
	request,_:=http.NewRequest("POST","/post/1",nil)
	mux.ServeHTTP(writer,request)

	if writer.Code!=200{
		t.Errorf("response code error,want 200,got %d",writer.Code)
	}
	if writer.Body.String()!="POST"{
		t.Errorf("response body error,want 'POST',got %s",writer.Body.String())
	}
} 


func TestHandlePut(t *testing.T){
	mux:=http.NewServeMux()
	mux.HandleFunc("/post/",handleRequest)

	writer:=httptest.NewRecorder()
	request,_:=http.NewRequest("PUT","/post/1",nil)
	mux.ServeHTTP(writer,request)

	if writer.Code!=200{
		t.Errorf("response code error,want 200,got %d",writer.Code)
	}
	if writer.Body.String()!="PUT"{
		t.Errorf("response body error,want 'PUT',got %s",writer.Body.String())
	}
} 


func TestHandleDelete(t *testing.T){
	mux:=http.NewServeMux()
	mux.HandleFunc("/post/",handleRequest)

	writer:=httptest.NewRecorder()
	request,_:=http.NewRequest("DELETE","/post/1",nil)
	mux.ServeHTTP(writer,request)

	if writer.Code!=200{
		t.Errorf("response code error,want 200,got %d",writer.Code)
	}
	if writer.Body.String()!="DELETE"{
		t.Errorf("response body error,want 'DELETE',got %s",writer.Body.String())
	}
} 
