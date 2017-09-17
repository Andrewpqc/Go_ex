package main

import (
	"fmt"
)

//居于已有类型定义自己的数据类型
//虽然myint基于int，两者并不等同，如下面：
//将int变量赋值给myint变量会报错
//
type myint int


func main(){
	var a myint
	fmt.Println(a)
	var i int
	i=3
	a=i
	fmt.Println(a)




}