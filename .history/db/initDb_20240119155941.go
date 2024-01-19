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
	str := `CREATE TABLE IF NOT EXISTS userInfo(
		id INTEGER PRIMARY KEY,
		username text NOT NULL,
		password text NOT NULL,
		creatime text,
		avatar text
	 );`
	ModifyDB(str)
}

func createArticle() {
	str := `CREATE TABLE IF NOT EXISTS articles(
		id INTEGER PRIMARY KEY,
		title text,
		content text,
		creatime text,
		userId INTEGER,
		FOREIGN KEY(userId) REFERENCES userInfo(id)
	);`
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
// 查询
func QueryRowDB(sql string) *sql.Row {
	return db.QueryRow(sql)
}

func QueryDB(sql string) (*sql.Rows, error) {
	return db.Query(sql)
}
