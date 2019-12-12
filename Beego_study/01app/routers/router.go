package routers

import (
	"01app/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{},"get:Login;post:Post")
	beego.Router("/register", &controllers.RegisterController{},"get:Register;post:Post")
	//beego.Router("/success", &controllers.RegisterController{})
	//beego.Router("/fail", &controllers.RegisterController{})

}
