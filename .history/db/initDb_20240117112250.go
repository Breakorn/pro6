package db

import (
	"database/sql"
	"pro6/logo"

	_ "github.com/mattn/go-sqlite3"
)

var bDb *sql.DB

func NewDb() {
	db, err := sql.Open("sqlite3", "./blog.db")
	if err != nil {
		logo.WARN("打开数据库失败")
		return
	}
	bDb = db

}

func createUser() {

}
