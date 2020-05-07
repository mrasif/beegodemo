package repo

import (
	"beegodemo/db"
	"beegodemo/models"
	"fmt"
	"time"
)

// Save : It will save artice
func Save(article models.Article) int64 {
	article.CreatedAt = time.Now().UnixNano() / int64(time.Millisecond)
	article.UpdatedAt = time.Now().UnixNano() / int64(time.Millisecond)
	id, _ := db.GetDb().Insert(&article)

	fmt.Println("Article inserted")
	return id

}

// Get : It will Retrive Article
func Get(id int64) models.Article {
	article := models.Article{
		ID: id,
	}
	db.GetDb().Read(&article)
	return article
}

// Update : It will update article
func Update(article models.Article) int64 {
	article.UpdatedAt = time.Now().UnixNano() / int64(time.Millisecond)
	count, _ := db.GetDb().Update(&article)
	return count
}

// Delete : It will delete article
func Delete(id int64) int64 {
	article := models.Article{
		ID: id,
	}
	count, _ := db.GetDb().Delete(&article)
	return count
}
