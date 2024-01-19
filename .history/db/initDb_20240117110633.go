package db

import (
	"database/sql"
	"os"

	"github.com/mbndr/logo"
)

var log = logo.NewSimpleLogger(os.Stderr, logo.DEBUG, "", true)

func NewDb() {

	db, err := sql.Open("sqlite3", "./blog.db")
	if err != nil {

		log.Erro("数据创建失败")
		return
	}
}
