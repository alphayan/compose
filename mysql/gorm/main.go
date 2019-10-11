package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Username int
	Password int
}

func main() {
	db, err := gorm.Open("mysql", "root:root123456@(192.168.199.248:3306)/gorm?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(new(User))
	db.LogMode(true)
	defer db.Close()
}
