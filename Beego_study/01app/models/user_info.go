package models

import (
	_"github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
)

var (
	db orm.Ormer //定义全局变量，数据库操作句柄
)

type UserInfo struct {
	Id int `orm:"column(id);pk"`//设置主键
	Uname string
	Upassword string
}

func init() {
	orm.Debug = true //开启调试功能，会打印SQL语句
	orm.RegisterDataBase("default","mysql","root:Zhaoxin..521@tcp(39.96.179.159:3306)/login?charset=utf8",30)
	orm.RegisterModel(new(UserInfo))
	db = orm.NewOrm()
}

func AddUser(user_info *UserInfo)(int64, error) {
	id, err := db.Insert(user_info)
	return id, err
}

func GetUser(user_info *UserInfo) error{//获取一个用户信息
	qb, _ := orm.NewQueryBuilder("mysql")//某数据库驱动
	qb.Select("*").From("user_info").Where("uname=\""+user_info.Uname+"\"")//拼接字符串

	sql := qb.String()
	err:= db.Raw(sql).QueryRow(user_info)//返回的是一条数据
	return err
}
