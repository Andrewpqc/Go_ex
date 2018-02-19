package main

import (
	"fmt"
	// "io/ioutil"
	"encoding/csv"
	"os"
	"strconv"
)

type Post struct {
	ID      int
	Content string
	Author  string
}

func main() {
	csvFile, err := os.Create("post.csv")
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	allposts := []Post{
		Post{ID: 1, Content: "aaaaa", Author: "A"},
		Post{ID: 2, Content: "aaaaa", Author: "B"},
		Post{ID: 3, Content: "aaaaa", Author: "C"},
		Post{ID: 4, Content: "aaaaa", Author: "D"},
	}

	writer := csv.NewWriter(csvFile)
	for _, post := range allposts {
		line := []string{strconv.Itoa(post.ID), post.Content, post.Author}
		err := writer.Write(line)
		if err != nil {
			panic(err)
		}
	}
	writer.Flush()

	file, err := os.Open("post.csv")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	var posts []Post
	for _, record := range records {
		id, _ := strconv.ParseInt(record[0], 0, 0)
		post := Post{ID: int(id), Content: record[1], Author: record[2]}
		posts = append(posts, post)
	}

	fmt.Println(posts[0].ID, posts[0].Content, posts[0].Author)
}
