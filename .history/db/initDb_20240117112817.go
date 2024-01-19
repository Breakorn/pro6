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
	CREATE TABLE DEPARTMENT(
		ID INT PRIMARY KEY      NOT NULL,
		DEPT           CHAR(50) NOT NULL,
		EMP_ID         INT      NOT NULL
	 )`

	db.Exec(str)
}
