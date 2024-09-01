package main

import (
	"fmt"
	"gotodo-app/db"
	"gotodo-app/model"
)

func main() {
	dbConn := db.NewDB() //成功したらDBインスタンスのアドレスが返ってくる
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Task{})
}