package db

import (
	"database/sql"
)

func NewDb() {

	db, err := sql.Open("sqlite3", "./blog.db")
	if err != nil {

		return
	}

}
