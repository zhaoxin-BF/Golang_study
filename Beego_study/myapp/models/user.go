package models

import (
	_"github.com/go-sql-driver/mysql"//下划线表示只引入这个包的Init方法，进行初始化操作，初始化连接数据库
	"github.com/astaxie/beego/orm"
)

var (
	db orm.Ormer//定义全局变量，数据库操作句柄
)

type Users struct{
	Id int `orm:"column(id);pk"`//设置主键
	Uname string
	Upassword string
}

func init() {
	orm.Debug = true //开启调试功能,会打印SQL语句
	orm.RegisterDataBase("default","mysql","root:Zhaoxin..521@tcp(39.96.179.159:3306)/beego?charset=utf8",30)
	orm.RegisterModel(new(Users))//连接到哪个表上users表，必须大写，对应数据库中的某表的小写
	db = orm.NewOrm()//
}

//裸查询
func AddUser(user_info *Users)(int64, error){//注意int64

	id, err := db.Insert(user_info)
	return id, err
}

//BuilderQuery查询
func ReadUserInfo(user_info *[]Users) {//还是得传地址
	qb, _ :=orm.NewQueryBuilder("mysql")//某数据库驱动
	qb.Select("*").From("users")//某张特定的表

	sql := qb.String()
	db.Raw(sql).QueryRows(user_info)//内部不需要再次传地址了
}