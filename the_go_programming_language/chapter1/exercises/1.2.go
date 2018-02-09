package main

import(
	"fmt"
	"os"
)

func main(){
	for i,d:=range os.Args[1:]{
		fmt.Println(i,d)
	}
}