package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

func walkDir(dir string, fileSize chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, fileSize)
		} else {
			fileSize <- entry.Size()
		}
	}
}

func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		return nil
	}
	return entries
}

func main() {

	ch := make(chan int64)
	tick := time.Tick(500 * time.Millisecond)
	go func() {
		walkDir("/home/", ch)
		close(ch)
	}()
	var sumSize int64
loop:
	for {
		select {
		case size, ok := <-ch:
			if !ok {
				break loop
			}
			sumSize += size
		case <-tick:
			fmt.Printf("%f GB\n", float64(sumSize)/1e9)
		}
	}
	fmt.Println("sum:", float64(sumSize)/1e9, "G")
}
