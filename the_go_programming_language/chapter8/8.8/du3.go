package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
	"sync"
)

var wg sync.WaitGroup 


func walkDir(dir string, fileSize chan<- int64,n *sync.WaitGroup) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			wg.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, fileSize,n)
		} else {
			fileSize <- entry.Size()
		}
	}
}

//限制并发数量
var ch= make(chan struct{},20)
func dirents(dir string) []os.FileInfo {
	ch<-struct{}{}
	defer func(){<-ch}()
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
	var wg sync.WaitGroup
	wg.Add(1)
	go walkDir("/home/", ch,&wg)
	go func(){
		wg.Wait()
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
