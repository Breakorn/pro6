package model

import (
	"fmt"
	"pro6/db"
	"pro6/logo"
	"time"

	"github.com/jehiah/go-strftime"
)

type ArticleLists struct {
}

type Article struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	Creatime string `json:"creatime"`
}

func Article_Search(name string) ([]*Article, error) {
	var id int
	// name :=
	str := fmt.Sprintf("SELECT id FROM userInfo WHERE username = '%v'", name)
	rows := db.QueryRowDB(str)
	if err := rows.Scan(&id); err != nil {
		logo.ERROR(err)
		return nil, err
	}

	str1 := fmt.Sprintf("SELECT title, content, creatime FROM articles WHERE userId = '%v'", id)
	rows2, err := db.QueryDB(str1)
	if err != nil {
		return nil, err
	}
	articles := make([]*Article, 0)
	for rows2.Next() {
		var (
			title    string
			content  string
			creatime string
		)
		var post = new(Article)
		err := rows2.Scan(&title, &content, &creatime)

		post.Title = title
		post.Content = content
		post.Creatime = creatime

		if err != nil {
			logo.ERROR(err)
		}
		articles = append(articles, post)

	}
	return articles, nil
}

func Article_Insert(name string, article *Article) bool {
	var id int
	// name :=
	str := fmt.Sprintf("SELECT id FROM userInfo WHERE username = '%v'", name)
	rows := db.QueryRowDB(str)
	if err := rows.Scan(&id); err != nil {
		logo.ERROR(err)
		return false
	}
	currentTime := time.Now()
	// 使用 strftime 格式化时间
	customFormat := strftime.Format("%Y-%m-%d %H:%M:%S", currentTime)
	count, err := db.ModifyDB("insert into articles(title,content,creatime,userId ) values (?,?,?,?)", article.Title, article.Content, customFormat, id)
	if err != nil {
		logo.ERROR(err)
		return false
	}
	fmt.Printf("count: %v\n", count)
	return true
}
