package model

import (
	"GinHello/initDB"
	"fmt"
	"log"
)

type ArticleModel struct {
	Id      int    `json:"id"`
	Type    string `json:"type"`
	Content string `json:"content"`
}

func (article *ArticleModel) Insert() int {
	fmt.Println("\n", article)
	result, e := initDB.Db.Exec("insert into ginhello.article (type, content) values (?, ?);", article.Type, article.Content)
	if e != nil {
		log.Panicln("文章添加失败", e.Error())
	}
	i, _ := result.LastInsertId()
	return int(i)
}
