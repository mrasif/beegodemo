package repo

import (
	"beegodemo/db"
	"beegodemo/models"
	"fmt"
)

// Save : It will save artice
func Save(article models.Article) int64 {
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

// GetAll : It will returns all Artcles
func GetAll() *[]models.Article {
	qs := db.GetDb().QueryTable("articles")
	articles := new([]models.Article)
	qs.All(articles)
	return articles
}

// Update : It will update article
func Update(article models.Article) int64 {
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
