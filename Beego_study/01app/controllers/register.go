package controllers

import (
	"01app/models"
	"github.com/astaxie/beego"
)

type RegisterController struct {
	beego.Controller
}

func (c *RegisterController) Register() {
	c.TplName = "register.html"
}

func (c *RegisterController) Post() {
	user_info := models.UserInfo{}
	if err := c.ParseForm(&user_info); err!=nil {
		c.TplName = "fail.html"
	}
	if user_info.Upassword == "" && user_info.Uname == "" {
		c.TplName = "fail.html"
	} else {
		if _, err := models.AddUser(&user_info); err != nil{
			c.TplName = "fail.html"
		}else {
			c.TplName = "success.html"
		}
	}
}
