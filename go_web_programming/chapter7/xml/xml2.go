package main

import (
	"io/ioutil"
	"encoding/xml"
	// "fmt"
	// "io"
	// "io/ioutil"
	"os"
)

type Post struct {
	XMLName  xml.Name  `xml:"post"`
	Id       string    `xml:"id,attr"`
	Content  string    `xml:"content"`
	Author   Author    `xml:"author"`
	Comments []Comment `xml:"comments>comment"`
	// Xml      string    `xml:",innerxml"`
}

type Author struct {
	Id   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

type Comment struct {
	Id      string `xml:"id,attr"`
	Content string `xml:"content"`
	Author  Author `xml:"author"`
}

// func main() {
// 	xmlFile, err := os.Open("example2.xml")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer xmlFile.Close()
// 	xmlData, err := ioutil.ReadAll(xmlFile)
// 	if err != nil {
// 		panic(err)
// 	}
// 	var post Post
// 	xml.Unmarshal(xmlData, &post)
// 	fmt.Println(post)
// }

// func main() {
// 	xmlFile, err := os.Open("example2.xml")
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer xmlFile.Close()

// 	decoder := xml.NewDecoder(xmlFile)
// 	var comments []Comment
// 	for {
// 		t, err := decoder.Token()
// 		if err == io.EOF {
// 			break
// 		}

// 		if err != nil {
// 			panic(err)
// 		}

// 		switch se := t.(type) {
// 		case xml.StartElement:
// 			if se.Name.Local == "comment" {
// 				var comment Comment
// 				decoder.DecodeElement(&comment, &se)
// 				// fmt.Println(comment)
// 				comments = append(comments, comment)
// 			}
// 		}
// 	}
// 	fmt.Println(comments)
// }



func main() {
	xmlFile, err := os.Open("example2.xml")
	if err != nil {
		panic(err)
	}
	defer xmlFile.Close()
	xmlData, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		panic(err)
	}
	var post Post
	xml.Unmarshal(xmlData, &post)
	// fmt.Println(post)

	output,err:=xml.MarshalIndent(&post,"","\t")
	if err!=nil{
		panic(err)
	}

	err=ioutil.WriteFile("output.xml",[]byte(xml.Header+string(output)),0644)
	if err!=nil{
		panic(err)
	}
}
