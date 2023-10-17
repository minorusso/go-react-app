package main //package migrateにはなぜならない？-> main関数を使用するためmainパッケージに所属する必要がある

import (
	"fmt"
	"go-rest-api/db"
	"go-rest-api/model"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	// 引数にはDBに反映させたいモデル構造を追加
	dbConn.AutoMigrate(&model.User{}, &model.Task{})
}
