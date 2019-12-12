package controllers

import (
	"github.com/astaxie/beego"
)
type UserInfo struct{
	Username string
	Password string
}
type TestLoginMainController struct {//首字母必须要大写
	beego.Controller
}

func (c *TestLoginMainController) Login() {
	name := c.Ctx.GetCookie("name")
	password :=c.Ctx.GetCookie("password")

	//do verify work
	if name != ""{
		c.Ctx.WriteString("Username:" + name + "Password" + password)
	}else{
		c.Ctx.WriteString(`<html><form action="http://127.0.0.1:8080/test_login" method="post">
								用户名：<input type="text" name="Username"/></br>
								密  码：<input type="password" name="Password"/></br>
								<input type="submit" value="提交">
							</form></html>`)
	}
}

func (c *TestLoginMainController) Post() {
	u := UserInfo{}
	if err := c.ParseForm(&u);err!=nil{//引用传出数据//接受表单
		//process error
		c.Ctx.WriteString("Error")//输出屏幕
	}
	c.Ctx.SetCookie("name",u.Username,100,"/")
	c.Ctx.SetCookie("password",u.Password,100,"/")

	c.Ctx.WriteString("Username:"+u.Username+"Password"+u.Password)//输出屏幕
}