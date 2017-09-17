/*
嵌入类型

*/
package main

import (
	"fmt"
)

type Birthday struct{
	year int
	month int
	day int
}


func (b *Birthday) changeBirthday(newYear int ,newMonth int,newday int){
	b.year=newYear
	b.month=newMonth
	b.day=newday
}
type Student struct{
	name string
	score int
	//这个就是嵌入类型
	Birthday
}
func (s Student) printAll1(){
	fmt.Printf("%s--%d--%d--%d--%d\n",s.name,s.score,s.Birthday.year,s.Birthday.month,s.Birthday.day)
}

func (s Student) printAll2(){
	fmt.Printf("%s--%d--%d--%d--%d\n",s.name,s.score,s.year,s.month,s.day)
}


func main(){
	s1:=Student{
		name:"xiaoming",
		score:88,
		Birthday:Birthday{
			year:1998,
			month:5,
			day:18,
		},
	}
	s1.printAll1()
	//下面两种调用内部类型方法的方式都可以
	// s1.Birthday.changeBirthday(1999,3,15)
	s1.changeBirthday(1997,8,23)
	s1.printAll2()
/*
注意上面的s1的初始化方式和printAll1和printAll2的定义的不同
既可以通过内部类型的来访问其方法与属性，如上面的：
s1.Birthday.changeBirthday(1999,3,15)
s.Birthday.year

也可以通过外部类型变量直接访问内部类型的属性与方法：
s1.changeBirthday(1997,8,23)
s.year


*/

}