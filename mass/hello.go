//go语言中的数组
// package main

// import (
// 	"fmt"
// 	// "io/ioutil"
// 	// "os"
// )

//每个程序可以写任意多个init函数，
//他们会在main函数之前执行
// func init() {
// 	//接收命令行参数
// 	filename:=os.Args[1]
// 	content,err:=ioutil.ReadFile(filename)
// 	if err!=nil{
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println(content)
// 	text:=string(content)

// 	fmt.Println(text)
// 	fmt.Println("读取文件完毕")
// 	fmt.Println("---------------1----------------")
// }

// func init() {
// 	var shuzu0 [5]int
// 	//shuzu0初始化为[0,0,0,0,0]

// 	//go的数组的定义与初始化
// 	shuzu1 := [3]int{1, 2, 3}
// 	fmt.Println(shuzu1[0])

// 	//go的数组定义与初始化的另外一种方式,循环数组
// 	var shuzu2 = [...]int{1, 2, 3, 4, 5, 6}
// 	//循环go数组
// 	for _, content := range shuzu2 {
// 		fmt.Printf("%d\n", content)
// 	}

// 	//初始化数组并指定特定元素的值
// 	shuzu3 := [5]int{1: 2, 3: 5}
// 	//此时数组被初始化为[0,2,0,5,0]
// 	for index, content := range shuzu3 {
// 		//此处返回的index是索引，content是实际数据值
// 		fmt.Printf("%d:%d\n", index, content)
// 	}

// 	//将一个数组赋值给一个同类型的数组(长度和存储的数据的类型相同)
// 	shuzu0=shuzu3
// 	for index, content := range shuzu0 {
// 		//此处返回的index是索引，content是实际数据值
// 		fmt.Printf("%d:%d\n", index, content)
// 	}

// 	//容纳指针的数组
// 	shuzu4:=[3]*string{new(string),new(string),new(string)}
// 	*shuzu4[0]="aaaa"
// 	*shuzu4[1]="bbbb"
// 	*shuzu4[2]="cccc"
// 	fmt.Println(*shuzu4[0])
// }

// func init(){
// 	s1:= [5]string{"red","blue","white","black","green"}
// 	//以下标的方式访问元素
// 	fmt.Println(s1[2])
// }
// func init() {
	//多维数组
	//空的二维数组，各个位置的值均为０
	// var s0 [3][3]int
	
	//声明一个二维数组并且初始化
	// s1 := [4][2]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}
	// fmt.Println(s1[1][1])
	//声明一个二维数组，并且初始化指定的位置,其它的位置初始化为０
	// s2 := [3][3]int{0: {1, 2, 3}, 2: {1, 5, 6}}

	//声明一个二维数组，并且初始化指定的内层数组和外层数组,其它的位置初始化为０
// 	s3 := [2][5]int{0: {1: 5, 3: 6}}

// 	fmt.Println(s3[0][1])
// }

// func main() {
// 	//do nothing
// }
