package main

import "github.com/jinzhu/gorm"

func create(db *gorm.DB, u *User) {
	db.Create(u)
}
