/*
interface和方法集实现多态
*/
package main

import (
	"fmt"
)

type Dog struct{
	name string
	favorite string
	age int
}
func (d *Dog) eat(){
	fmt.Printf("dog %s love %s\n",d.name,d.favorite)
}
type Cat struct{
	name string
	favorite string
}
func (c *Cat) eat(){
	fmt.Printf("cat %s love eat %s\n",c.name,c.favorite)
}

//方法集
type eeeaaattt interface{
	eat()
}

func AnimalEat(u eeeaaattt){
	u.eat()
}


//实现了多态
func main(){
	d1:=Dog{"hhhdog","骨头",5}
	c1:=Cat{"jjjcat","鱼"}
	
	//注意这里一定要传地址
	AnimalEat(&d1)
	AnimalEat(&c1)
}