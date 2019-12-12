package controllers

import (
	"github.com/astaxie/beego"
	"strconv"
)

type TestContextController struct {
	beego.Controller
}

func (c *TestContextController) Get() {      //获取访问的IP + Port
	c.Ctx.WriteString(c.Ctx.Input.IP() + ":" + strconv.Itoa(c.Ctx.Input.Port())) //端口是整型

	c.Ctx.WriteString(c.Ctx.Input.Query("name"))//获取Get Url中的数据并输出

	m := make(map[string]float64)
	m["zhangsan"] = 98.7
	c.Ctx.Output.JSON(m, false, false)
}
