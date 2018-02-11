package main

import (
	"fmt"
	"os"
	"path/filepath"
	"io/ioutil"
)


func walkDir(dir string,fileSize chan<-int64){
	for _,entry :=range dirents(dir){
		if entry.IsDir(){
			subdir:=filepath.Join(dir,entry.Name())
			walkDir(subdir,fileSize)
		}else{
			fileSize<-entry.Size()
		}
	}
}


func dirents(dir string)[]os.FileInfo{
	entries,err:=ioutil.ReadDir(dir)
	if err!=nil{
		fmt.Fprintf(os.Stderr,"%v",err)
		return nil
	}
	return entries
	
}

func main(){

	ch:=make(chan int64)
	go func(){
		walkDir("/home/",ch)
		close(ch)
	}() 
	var sumSize int64
	for {
		size,ok:=<-ch
		if !ok{
			break
		}
		sumSize+=size
	}
	fmt.Printf("%.1f GB\n",float64(sumSize)/1e9)


}
