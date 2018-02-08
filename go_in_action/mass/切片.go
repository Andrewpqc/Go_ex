//go语言中的切片
package main

import (
	"fmt"
)


func init(){
// //内置的make函数来创建切片

// 	//创建了一个长度和容量都为５的装string类型变量的切片slice1,里面的数据被初始化为空字符串
// 	slice1:=make([]string,5)
// 	//创建了一个长度为３，容量为５的装int类型数据的切片slice2，里面的数据被初始化为０
// 	slice2:=make([]int,3,5)
// 	//注意不能创建容量比长度小的切片

// //切片字面量创建并初始化切片
// 	slice3:=[]int{1,2,3} //长度和容量均为３
// 	slice4:=[]string{"hhh","sss"}　//长度和容量均为２
// 	//注意此处与创建数组非常的相似，创建数组需要在中括号中指定数组的长度，而创建
// 	//切片则不需要在中括号中填写任何东西

// 	//在使用切片字面量的时候，可以设置初始的长度和容量,要做
// 	//的就是给出所需长度和容量作为索引，如下就声明了一个长度和容量均为１００的切片
// 	slice5:=[]int{99:0}
// //nil和空切片

// 	var slicenil1 []int

// 	//使用make创建空的整型切片
// 	slicenil2:=make([]int,0)

// 	//使用切片字面量创建空的整型切片
// 	slicenil3:=[]int{}




//使用切片
// slice:=[]int{1,2,3,4,5,6}

// for _,content:=range slice{
// 	fmt.Println(content)
// }
// fmt.Println("\n")

// slice[0]=100
// for _,content:=range slice{
// 	fmt.Println(content)
// }
// fmt.Println("\n")
// newSlice:=slice[0:3]
// //对底层数组容量为k的切片slice[i:j]来说
// //它的长度为j-i,容量为k-i
// for _,content:=range newSlice{
// 	fmt.Println(content)
// }
// newSlice[1]=99999
// //这里对新切片做出了修改，那么底层的数组同步做出了修改
// for _,content:=range newSlice{
// 	fmt.Println(content)
// }
// for _,content:=range slice{
// 	fmt.Println(content)
// }


//append增加切片的长度

// slice9:=[]int{1,2,3}
// slice10:=[]int{5,6,7}
// slice9=append(slice9,slice10...)

// //for和range版本的for循环
// for _, vaule:=range slice9{
// 	fmt.Println(vaule)
// }

// //传统的for循环
// for index:=0;index<len(slice9);index++{
// 	fmt.Println(slice9[index])
	
// }
// fmt.Println(cap(slice9))


// }
// func main(){

// 	//
// }