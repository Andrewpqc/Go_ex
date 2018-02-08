/**
*自定义结构类型
*/
package main

import (
	"fmt"
)

//自定义一个学生类型三个成员变量
type Student struct{
	name string
	id string
	score int
}

//Student类型的方法，指针接收者
func (s Student) printAll(){
	fmt.Printf("%s--%s--%d\n",s.name,s.id,s.score)
}

//Student类型的方法，值接收者
func (s *Student) changeScore(newScore int){
	s.score=newScore
}
func main(){
	//jack变量中的属性被初始化为对应类型的零值
	var jack Student
	//初始化的第二种方法
	lisa:=Student{"lisa","123456",86}
	//初始化的第三种方法
	bob:=&Student{name:"bob",id:"1111111",score:56}
	bob.printAll()
	//go代码在背后执行了：(*bob).printAll()
	lisa.printAll()
	lisa.changeScore(85)
	//go代码在背后执行了：(&lisa).printAll()
	bob.changeScore(85)
	bob.printAll()
	lisa.printAll()
	jack.name="jack"
	jack.printAll()

	/**
	*注意在这里如果不打分号的话，每条语句后面什么都不能有
	*不能有注释，甚至不能有空格
	*/


	/**
	*如果需要在类型的方法中做出的操作能够影响变量本身的话
	*那么就需要使用指针接收者的方式，在声明变量的时候不需要考虑
	*声明指针变量还是直接声明对象，因为go在背后为我们做好了这种转换
	*不必严格符合接收者的类型
	*/
}

/*
go函数
func [(u [*]TYPE)] funtionName([args]) [(returnItem type)]{
	//function body
	[]中的内容为可选内容
}



*/