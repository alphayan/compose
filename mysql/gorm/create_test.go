package main

import (
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func BenchmarkCreate(b *testing.B) {
	db, err := gorm.Open("mysql", "root:root123456@(192.168.199.248:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	u := new(User)
	db.AutoMigrate(new(User))
	u.Username, u.Password = 1, 1
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		u.ID = 0
		err := db.Create(u).Error
		if err != nil {
			panic(err)
		}
	}
}
