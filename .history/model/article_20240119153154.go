package model

import (
	"fmt"
	"pro6/db"
)

type ArticleLists struct {
}

type Article struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	Creatime string `json:"creatime"`
}

func Article_Search(name string) {
	// var userID int
	// name :=
	str := fmt.Sprintf("SELECT id FROM userInfo WHERE username = '%v'", name)
	db.QueryDB(str)

	// // 插入文章数据
	// _, err = db.Exec("INSERT INTO posts (title, content, user_id) VALUES ('First Post', 'Content of the first post', ?)", userID)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}

func Article_Insert() {

}
