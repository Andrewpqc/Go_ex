package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	count := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countlines(count, os.Stdin)
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "err:%v", err)
				continue
			}
			countlines(count, f)
			f.Close()
		}
	}

	for line, num := range count {
		fmt.Println(line, ":", num)
	}
}

func countlines(count map[string]int, f *os.File) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		count[input.Text()]++
	}
}
