package main

import (
	"fmt"
	"bytes"
)

func printints(valuse [] int) string{
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i,d:=range valuse{
		if i>0{
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf,"%d",d)
	}
	buf.WriteByte(']')
	return buf.String()
}

func main(){
	a:=[]int{1,2,3}
	s:=printints(a)
	fmt.Println(s)6d
}