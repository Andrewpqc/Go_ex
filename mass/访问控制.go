package main

//引用包里面的函数
import "./testpackage"

func main(){
	a:=testpackage.Test(3)
	testpackage.TTT()
	println(a)

	s1:=testpackage.Student{"xiaoming",25}
	println(s1.Name)
}

/*
这里展示了怎样定义包和怎样在包里面引用大写开头的函数和变量



*/