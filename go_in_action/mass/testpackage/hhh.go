package testpackage

func Test(i int) int{
	sum:=0
	for j:=1;j<=i;j++{
		sum+=j
	}
	return sum
}

type Student struct{
	Name string
	Score int
}


/*
这里的变量名开头是大写的时候外界才可见
开头为小写的时候外界是不可见的
*/