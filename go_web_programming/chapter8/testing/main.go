package main

import (
	"encoding/json"
	"fmt"
	"os"
	"io/ioutil"
)

type Post struct {
	Id       int       `json:"id"`
	Content  string    `json:"content"`
	Author   Author    `json:"author"`
	Comments []Comment `json:"comments"`
}

type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func decode(filename string) (post Post, err error) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("error happend when open json file!")
		return
	}
	defer jsonFile.Close()
	decoder := json.NewDecoder(jsonFile)
	err = decoder.Decode(&post)
	if err != nil {
		fmt.Println("error happend when decoding!")
		return
	}
	return
}


func unmarshal(filename string)(post Post,err error){
	jsonFile,err:=os.Open(filename)
	if err!=nil{
		fmt.Println("error happened when open json file!")
		return
	}
	defer jsonFile.Close()

	jsonData,err:=ioutil.ReadAll(jsonFile)
	if err!=nil{
		fmt.Println("error happened when read data from json file!")
		return
	}
	err=json.Unmarshal(jsonData,&post)
	if err!=nil{
		fmt.Println("error happened when unmarshal!")
		return
	}
	return
}

func main() {
	post, err := decode("post.json")
	if err != nil {
		//错误处理
	}

	fmt.Println(post)

	post2,err:=unmarshal("post.json")
	fmt.Println(post2)
}
