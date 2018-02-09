package main

import (
	"bufio"
	"fmt"
	"strings"
)

type LineCounter int
type WordCounter int

func (l *LineCounter) Write(p []byte) (int, error) {
	s := string(p)
	var count int
	for _, ch := range s {
		if ch == 10 { //10对应的是换行符
			count++
		}
	}
	*l = LineCounter(count + 1)
	return count + 1, nil
}

func (w *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return count, err
	}
	*w = WordCounter(count)
	return count, nil
}

func main() {
	var l LineCounter
	var w WordCounter
	fmt.Fprintf(&l, "dsfg\nsadfg\nassdf\ng%d", 3)
	fmt.Println("line:", l)
	fmt.Fprintf(&w, "I love you, how are about you? %s", "hh hh")
	fmt.Println("words:", w)
}
