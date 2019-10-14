package main

import "github.com/jinzhu/gorm"

var db *gorm.DB

func initDB() {
	var err error

	db, err = gorm.Open("mysql", "root:root123456@(192.168.199.248:3306)/gorm?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(new(User))
	db.LogMode(true)

}
