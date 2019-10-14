package main

import (
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

func initDB() {
	var err error

	engine, err = xorm.NewEngine("mysql", "root:root123456@(192.168.199.248:3306)/gorm?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	engine.Sync2(new(User))
	engine.ShowSQL(true)
}
