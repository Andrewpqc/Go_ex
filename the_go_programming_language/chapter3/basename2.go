package main

import (
	"os"
	"fmt"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	s := os.Args[1]

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	fmt.Println(s)
}
