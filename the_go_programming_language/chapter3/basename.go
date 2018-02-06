package main

import (
	"strings"
	"os"
	"fmt"
)

func main() {
	if (len(os.Args) != 2) {
		return
	}
	s := os.Args[1]
	a := strings.LastIndex(s, "/")
	if a != -1 {
		s = s[a+1:]
	}

	b := strings.LastIndex(s, ".")
	if b != -1 {
		s = s[:b]
	}

	fmt.Println(s)
}
