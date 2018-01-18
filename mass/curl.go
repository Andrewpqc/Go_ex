package main
import(
	"fmt"
	"os"
	"io"
	//"net/http"
	"bytes"

)

// func init(){
// 	//先处理输入的命令
// 	if len(os.Args)!=2{
// 		for _,vaule:=range os.Args{
// 			fmt.Println(vaule)
// 		}
// 		fmt.Println("命令错误。。。")
// 		os.Exit(-1)
// 	}
// }



// func main(){
// 	res,err:=http.Get(os.Args[1])
// 	if err!=nil{
// 		fmt.Println("get error")
// 		fmt.Println(err)
// 		return
// 	}
// 	io.Copy(os.Stdout,res.Body)
// 	err1:=res.Body.Close()
// 	if err1!=nil{
// 		fmt.Println(err1)
// 	}


// }

// func main(){
// 	var b bytes.Buffer
// 	b.Write([]byte("hello"))

// 	fmt.Fprintf(&b,"world!")
// 	io.Copy(os.Stdout,&b)
// }