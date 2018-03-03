package controllers

import (
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	c.Ctx.WriteString("this is a url to use controller!,method get")
}

func (c *UserController) Post(){
	c.Ctx.WriteString("hello,from post!")
}
