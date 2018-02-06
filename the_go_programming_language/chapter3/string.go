package main

import(
	"fmt"
	"unicode/utf8"
)


func main(){
	a:="abcdfr"
	b:=utf8.RuneCountInString(a)
	fmt.Println(b)
	
}