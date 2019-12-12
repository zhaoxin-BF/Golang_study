package controllers

import (
	"github.com/astaxie/beego"
)

type TestMainController struct {//必须要大写
	beego.Controller
}

func (c *TestMainController) Get() {
	c.Ctx.WriteString("Hello world!")//直接输出
}

func (c *TestMainController) Post() {
	c.Ctx.WriteString("Hello Go!")//直接输出
}