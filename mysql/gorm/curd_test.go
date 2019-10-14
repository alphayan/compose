package main

import (
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func BenchmarkCreate(b *testing.B) {
	db, err := gorm.Open("mysql", "root:root123456@(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")
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
func BenchmarkUpdate(b *testing.B) {
	db, err := gorm.Open("mysql", "root:root123456@(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	u := new(User)
	db.AutoMigrate(new(User))
	u.Username, u.Password = 1, 1
	db.Create(u)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := db.Model(u).Update(u).Error
		if err != nil {
			panic(err)
		}
	}
}
