package main

import(
	"fmt"
)

type MyWriterType int

func (w *MyWriterType)Write(p []byte)(int, error){
	*w+=MyWriterType(len(p))
	return len(p),nil
}

func main(){
	var w MyWriterType
	fmt.Fprintf(&w,"Ihhh,%s","jjj")
	fmt.Println(w)
}