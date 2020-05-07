package repo

import (
	"beegodemo/db"
	"beegodemo/models"
	"fmt"
	"time"
)

func Save(article models.Article) int64 {
	article.CreatedAt = time.Now().UnixNano() / int64(time.Millisecond)
	article.UpdatedAt = time.Now().UnixNano() / int64(time.Millisecond)
	id, _ := db.GetDb().Insert(&article)

	fmt.Println("Article inserted")
	return id

}

func Get(id int64) models.Article {
	article := models.Article{
		Id: id,
	}
	db.GetDb().Read(&article)
	return article
}

func Update(article models.Article) int64 {
	article.UpdatedAt = time.Now().UnixNano() / int64(time.Millisecond)
	count, _ := db.GetDb().Update(&article)
	return count
}
