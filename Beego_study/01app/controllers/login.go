package controllers

import (
	"01app/models"
	"fmt"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Login() {
	c.TplName = "login.html"
}

func (c *LoginController) Post() {
	user := models.UserInfo{}
	if err := c.ParseForm(&user); err!=nil {
		c.TplName = "fail.html"
	}
	fmt.Println(user)
	user_info := models.UserInfo{Uname:user.Uname}
	if err := models.GetUser(&user_info); err!=nil{
		fmt.Println(user_info)
		c.TplName = "fail.html"
	}else {
		fmt.Println(user_info)
		fmt.Println(user)
		if user_info.Uname == user.Uname && user_info.Upassword == user.Upassword {
			c.TplName = "success.html"
		}else {
			c.TplName = "fail.html"
		}
	}
}