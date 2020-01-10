package models

import (
	_ "github.com/go-sql-driver/mysql"
    "github.com/astaxie/beego/orm"
)

type User struct {
    Id          int
    Name        string
}


func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
    orm.RegisterDataBase("default", "mysql", "root:root123456@tcp(172.17.33.196:3306)/orm_test?charset=utf8")
    // 需要在init中注册定义的model
    orm.RegisterModel(new(User), new(Tgo))
    orm.RunSyncdb("default", false, true)
}

func AddUser(user *User) (int64, error) {
    o := orm.NewOrm()
    id, err := o.Insert(user)
    return id, err
}

