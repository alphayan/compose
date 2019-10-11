package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"xorm.io/core"
)

type User struct {
	ID       uint `xorm:"id,pk"`
	Username int
	Password int
}

func main() {
	engine, err := xorm.NewEngine("mysql", "root:root123456@(192.168.199.248:3306)/xorm?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	engine.Sync2(new(User))
	engine.Logger().SetLevel(core.LOG_DEBUG)
}
