package controllers

import (
"github.com/astaxie/beego"
)

type TestInputMainController struct {//必须要大写
	beego.Controller
}

type User struct{
	Username string
	Password string
}




func (c *TestInputMainController) Get() {
	//id := c.GetString("id")
	//c.Ctx.WriteString("Id = "+id)//直接输出
	//
	//name := c.Input().Get("name")
	//c.Ctx.WriteString("name = "+ name)

	name := c.GetSession("name")
	password := c.GetSession("password")

	if nameString, ok := name.(string);ok && nameString != ""{
		c.Ctx.WriteString("Name:"+name.(string)+"Password:"+password.(string))//转换为string类型
	}else{
		c.Ctx.WriteString(`<html><form action="http://127.0.0.1:8080/testinput" method="post">
								<input type="text" name="Username"/>
								<input type="password" name="Password"/>
								<input type="submit" value="提交">
					</form></html>`)
	}
}

func (c *TestInputMainController) Post() {
	u := User{}
	if err := c.ParseForm(&u);err!=nil{
		//process error
	}
	c.Ctx.WriteString("输出成功")
}