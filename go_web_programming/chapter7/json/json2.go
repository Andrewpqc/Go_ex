package main

import (
	"encoding/json"
	// "fmt"
	// "io/ioutil"
	"os"	

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

// func main() {
// 	post := Post{
// 		Id:      1,
// 		Content: "Hello,world!",
// 		Author: Author{
// 			Id:   2,
// 			Name: "andrew",
// 		},
// 		Comments: []Comment{
// 			Comment{
// 				Id:      3,
// 				Content: "good post!",
// 				Author:  "aaa",
// 			},
// 			Comment{
// 				Id:      4,
// 				Content: "have a nice day!",
// 				Author:  "bbb",
// 			},
// 		},
// 	}

// 	output,err:=json.MarshalIndent(&post,"","\t")
// 	if err!=nil{
// 		panic(err)
// 	}
// 	err=ioutil.WriteFile("output.json",output,0644)
// 	if err!=nil{
// 		panic(err)
// 	}
// }



func main(){
	post := Post{
		Id:      1,
		Content: "Hello,world!",
		Author: Author{
			Id:   2,
			Name: "andrew",
		},
		Comments: []Comment{
			Comment{
				Id:      3,
				Content: "good post!",
				Author:  "aaa",
			},
			Comment{
				Id:      4,
				Content: "have a nice day!",
				Author:  "bbb",
			},
		},
	}

	jsonFile,err:=os.Create("output2.json")
	if err!=nil{
		panic(err)
	}

	encoder:=json.NewEncoder(jsonFile)
	encoder.SetIndent("","\t")
	err=encoder.Encode(&post)
	if err!=nil{
		panic(err)
	}
	

}