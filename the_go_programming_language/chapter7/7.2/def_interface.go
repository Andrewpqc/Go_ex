//接口的定义方法
package main

import(

)


//单接口
type Writer interface{
	Write(p []byte)(int,error)
}

type Reader interface{
	Read(p []byte)(int,error)
}

type Closer interface{
	Close()  error
}


//嵌入式接口
type ReadWriter interface{
	Writer
	Reader
}

type ReadWriteCloser interface{
	Reader
	Writer
	Closer
}


//可以不用嵌入式来定义，像下面这样
// type ReadWriter interface{
// 	Read(p []byte)(int,error)
// 	Write(p []byte)(int,error)
// }

//混合式
// type ReadWriter interface{
// 	Write(p []byte)(int,error)
// 	Reader
// }
