package main

import (
	"testing"

	"github.com/go-xorm/xorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func BenchmarkInsert(b *testing.B) {
	db, err := xorm.NewEngine("mysql", "root:root123456@(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	u := new(User)
	u.Username, u.Password = 1, 1
	b.ResetTimer()
	db.Sync2(new(User))
	for i := 0; i < b.N; i++ {
		u.ID = 0
		db.Insert(u)
	}
}
func BenchmarkUpdate(b *testing.B) {
	db, err := xorm.NewEngine("mysql", "root:root123456@(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	u := new(User)
	u.Username, u.Password = 1, 1
	b.ResetTimer()
	db.InsertOne(u)
	db.Sync2(new(User))
	for i := 0; i < b.N; i++ {
		db.Table(u).Update(u)
	}
}
