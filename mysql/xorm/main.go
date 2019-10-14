package main

import (
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID       uint `xorm:"id,pk"`
	Username int
	Password int
}

func main() {
	initDB()
}
