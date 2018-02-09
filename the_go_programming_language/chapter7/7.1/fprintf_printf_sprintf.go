/*
fmt包中的Fprintf,Printf,Sprintf的实现
*/


package fmt

import(
	"io"
	"os"
	"bytes"
)

type Writer interface{
	Write(p[]byte)(int,error)
}


func Fprintf(w io.Writer,format string,args ...interface{})(int,error){
	return 3,nil
}

func Printf(format string,args ...interface{})(int,error){
	return Fprintf(os.Stdout,format,args...)
}

func Sprintf(format string,args ... interface{})string{
	var buf bytes.Buffer
	Fprintf(&buf,format,args...)
	return buf.String()
}
