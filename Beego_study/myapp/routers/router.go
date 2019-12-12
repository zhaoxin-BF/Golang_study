package routers

import (
	"myapp/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{}, "get:Mainfunc")

	beego.Router("/test", &controllers.TestMainController{},"get:Get;post:Post")
	beego.Router("/testinput", &controllers.TestInputMainController{},"get:Get;post:Post")//input_form表单
	beego.Router("/test_login", &controllers.TestLoginMainController{},"get:Login;post:Post")//cookies
	beego.Router("/test_login_Session", &controllers.TestLoginSessionMainController{},"get:Login;post:Post")//session
	beego.Router("/test_model", &controllers.TestModelController{},"get:Get;post:Post")
	beego.Router("/test_view", &controllers.TestViewController{},"get:Get;post:Post")
	beego.Router("/test_httplib", &controllers.TestHttplibController{},"get:Get;post:Post")
	beego.Router("/test_context", &controllers.TestContextController{},"get:Get;post:Post")
}