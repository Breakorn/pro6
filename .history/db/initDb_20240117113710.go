package db

import (
	"database/sql"
	"pro6/logo"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func NewDb() {
	db1, err := sql.Open("sqlite3", "./blog.db")
	if err != nil {
		logo.WARN("打开数据库失败")
		return
	}
	db = db1
	createUser()
}

func createUser() {
	str := `
	CREATE TABLE IF NOT EXISTS userInfo(
		name           CHAR(50)   NOT NULL,
		password       CHAR(50)      NOT NULL
	 )`
	db.Exec(str)
}
