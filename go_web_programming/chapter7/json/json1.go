package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	// "io/ioutil"
)

type Post struct {
	Id       int       `json:"id"`
	Content  string    `json:"content"`
	Author   Author    `json:"author"`
	Comments []Comment `json:"comments"`
}

type Comment struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// func main(){
// 	jsonFile,err:=os.Open("example1.json")
// 	if err!=nil{
// 		panic(err)
// 	}
// 	defer jsonFile.Close()
// 	jsonData,err:=ioutil.ReadAll(jsonFile)
// 	if err!=nil{
// 		panic(err)
// 	}
// 	var post Post
// 	err=json.Unmarshal(jsonData,&post)
// 	if err!=nil{
// 		panic(err)
// 	}
// 	fmt.Println(post.Author.Name)
// }

func main() {
	jsonFile, err := os.Open("example1.json")
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	decoder := json.NewDecoder(jsonFile)
	for {
		var post Post
		err := decoder.Decode(&post)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Println(post)
	}
}
