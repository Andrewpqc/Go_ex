package main

import (
	"fmt"
)

type Post struct {
	ID      int
	Content string
	Author  string
}

var PostByID map[int]*Post
var PostByAuthor map[string][]*Post

func store(p *Post) {
	PostByID[p.ID] = p
	PostByAuthor[p.Author] = append(PostByAuthor[p.Author], p)
}

func main() {
	PostByID = make(map[int]*Post)
	PostByAuthor = make(map[string][]*Post)
	post1 := Post{
		ID:      1,
		Content: "it very interesting",
		Author:  "A",
	}
	post2 := Post{
		ID:      2,
		Content: "it very interesting2",
		Author:  "B",
	}
	post3 := Post{
		ID:      3,
		Content: "it very interesting3",
		Author:  "C",
	}
	store(&post1)
	store(&post2)
	store(&post3)

	fmt.Println(PostByID[1])
	fmt.Println(PostByID[3])

}
