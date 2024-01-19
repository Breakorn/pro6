package db

import (
	"database/sql"
	"pro6/logo"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func NewDb() {
	db1, err := sql.Open("sqlite3", "./db/blog.db")
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
		name           text    NOT NULL,
		password       text    NOT NULL,
		creatime       timestamp/datetime,
		avatar         text 
	 )`
	ModifyDB(str)
}

// 操作数据库
func ModifyDB(sql string, args ...interface{}) (int64, error) {
	result, err := db.Exec(sql, args...)
	if err != nil {
		logo.WARN(err)
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		logo.WARN(err)
		return 0, err
	}
	return count, nil
}

// 查询
func DB_SEARCH(sql string) *sql.Rows {
	rows, err := db.Query(sql)
	if err != nil {
		logo.WARN(err)
		return nil
	}
	defer rows.Close()
	return rows

}
