package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"myapp/models"
	//_"github.com/go-sql-driver/mysql"//下划线表示只引入这个包的Init方法，进行初始化操作，初始化连接数据库
	//"github.com/astaxie/beego/orm"
)

//由于Model这个名字叫 Users 那么操作的表其实是 users
//type Users struct{
//	Id int `orm:"column(id);pk"`//设置主键
//	Uname string
//	Upassword string
//}

type TestModelController struct {
	beego.Controller
}

func (c *TestModelController) Get() {
	//orm.Debug = true //开启调试功能,会打印SQL语句
	//orm.RegisterDataBase("default","mysql","root:Zhaoxin..521@tcp(39.96.179.159:3306)/beego?charset=utf8",30)
	//orm.RegisterModel(new(Users))

	//o := orm.NewOrm()//生成一个Orm对象,类比C++ 的操作数据库句柄


	//user := Users{Uname:"boreas.zhao",Upassword:"123456"}//1、2共用

	////////////////////////////////////////////////
	//1、model查询
	//id, err := o.Insert(&user)//插入

	//user.Id = 1
	//user.Uname = "beifeng" //更新
	//id, err := o.Update(&user)

	//user.Id = 2
	//id, err := o.Delete(&user)//删除

	//user.Id = 1
	//o.Read(&user)//查询

	//c.Ctx.WriteString(fmt.Sprintf("user info: %v", user))


	////////////////////////////////////////////////
	////2、SQl原生读取,直接查询、一般在脚本中用，因为容易产生注入
	//var maps []orm.Params //重点
	//o.Raw("select * from users").Values(&maps)//全部返回，通过values 转换为数组
	//
	//for _, v := range maps{
	//	c.Ctx.WriteString(fmt.Sprintf("users info:%v", v))
	//}
	//
	////一行返回
	//var users []Users
	//o.Raw("select * from users").QueryRows(&users)//全部返回, 直接查找为数组，存入对象数组


	////////////////////////////////////////////////////////////
	//3、采用queryBuilder方式进行读取////实际开发较多的事情
	//var users []Users
	//
	//qb, _ := orm.NewQueryBuilder("mysql")//orm开头
	//
	//qb.Select("upassword").From("users").Where("uname = \"beifeng\"")//拼接字符串
	//
	//sql := qb.String()
	//o.Raw(sql).QueryRows(&users)//o开头
	//
	//c.Ctx.WriteString(fmt.Sprintf("user info:%v", users))

	///////////////////////////////////////////////////////
	//models 的使用操作
	user := models.Users{Uname:"boreas.zhao", Upassword:"123456"}
	id, err := models.AddUser(&user)
	if err != nil{
		c.Ctx.WriteString(fmt.Sprintf("err =%v", err))
	}
	c.Ctx.WriteString(fmt.Sprintf("id = %v", id))//id返回的是它的自增Id
}
