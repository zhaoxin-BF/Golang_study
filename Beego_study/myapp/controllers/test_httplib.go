package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
)

type TestHttplibController struct {
	beego.Controller
}

func (c *TestHttplibController) Get() {
	req := httplib.Get("http://douban.com")
	str, err := req.String()

	if err != nil{
		panic(err)
	}

	c.Ctx.WriteString(str)
}
