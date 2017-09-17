package main

import "fmt"

//go语言的映射的使用:映射是一个存储键值对的无序集合

//使用make创建一个映射，键的类型为string,值的类型为int
// map1:=make(map[string]int)

// //使用映射字面量创建一个映射
// map2:=map[string]string{
// 	"abc":"nnn",
// 	"def":"mmm"
// }

// //创建映射时最常用的方式就是使用映射字面量，映射的初始长度会根据初始化时键值对的数量来确定
// /*
// 映射的键可以是任意值。这个值的类型可以是内置的类型，也可以是结构类型，
// 只要这个值可以使用==做比价。切片，函数以及包含切片的结构类型这些类型由于具有引用语义，不能作为映射的键
// */

// //映射字面量声明一个空的映射
// map3:=map[string]string{}

// //不能将切片作为键
// //map4:=map[[]string]string{}  //错误

// //但是可以将切片作为值
// map5:=map[string][]int{}


//使用切片
// colors:=map[string]string{}
// colors["red"]="this is red's vaule"

/*
可以通过声明一个未初始化的映射来创建一个值为nil的映射（称为nil映射）。nil映射不能用于存储键值对
*/
func init(){

	// var map6 map[string]string
	// map6["jian"]="zhi"      //错误
	// fmt.Println(map6["jian"])

	//从映射获取值并且判断键是否存在

	map1:=map[string]string{"aaa":"this is aaa's vaule","bbb":"this is bbb's vaule"}
	
	//从映射中获取值，并且判断键是否存在
	vaule,exists:=map1["ccc"]
	if exists{
		fmt.Println(vaule)
	}else{
		fmt.Println("oh no")
	}

	//另一种方法，只获取值，判断值是否为空
	vaule1:=map1["bbb"]
	if vaule1!=""{
		fmt.Println(vaule1)
	}else{
		fmt.Println("oh no")
	}

/*
在go语言中，通过键来索引映射的时候，即使键不存也会返回一个值。在这种情况下
返回的是该值对应的类型的零值
*/
//for和range迭代映射
map2:=map[string]string{
	"aaa":"AAA",
	"bbb":"BBB",
	"ccc":"CCC",
	"ddd":"DDD",
}

for key,vaule:=range map2{
	fmt.Println(key,vaule)
}

//把一个键值对从映射中删除，使用内置的delete函数
delete(map2,"aaa")//删除aaa的那一项
for key,vaule:=range map2{
	fmt.Println(key,vaule)
}


}

func main (){
	// do something here
}