package controllers

import (
	"github.com/astaxie/beego"
	"myapp/models"
)

type TestViewController struct {
	beego.Controller
}

func (c *TestViewController) Get() {
	//c.Data["Title"] = "Hello"
	//c.Data["Content"] = "北峰科技园"
	//c.Data["IsDisplay"] = false //1, 0 //true, false
	//c.Data["Content2"] = "bores.zhao city"
	//c.TplName = "test_view.tpl"

	var users []models.Users
	models.ReadUserInfo(&users)//改变内容还是要传地址

	c.Data["Users"] = users
	c.TplName = "test_view.tpl"
}