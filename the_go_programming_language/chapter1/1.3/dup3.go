package main

import(
	"fmt"
	"os"
	"io/ioutil"
	"strings"
)

func main(){
	count:=make(map[string]int)
	data,err:=ioutil.ReadFile("dup.go")
	if err!=nil{
		fmt.Fprintf(os.Stderr,"error:%v",nil)
	}
	// fmt.Println(strings.Split(string(data),"\n"))
	for _,str:=range strings.Split(string(data),"\n"){
		count[str]++
	}

	for str,size:=range count{
		fmt.Println(str,size)
	}
	
}