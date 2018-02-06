package main

import (
	"strings"
	"fmt"
	"os"
	"net/url"

)

func main(){
	q:=url.QueryEscape(strings.Join(os.Args[1:]," "))
	fmt.Println(q)
}