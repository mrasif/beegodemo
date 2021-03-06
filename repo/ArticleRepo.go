package repo

import (
	"beegodemo/db"
	"beegodemo/models"
	"fmt"
	"log"
)

// Save : It will save artice
func Save(article models.Article) int64 {
	id, error := db.GetDb().Insert(&article)

	if error != nil {
		log.Fatal(error)
	}

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

// All : It will returns all Artcles
func All() *[]models.Article {
	qs := db.GetDb().QueryTable("articles")
	articles := new([]models.Article)
	qs.All(articles)
	return articles
}

// Update : It will update article
func Update(article models.Article) int64 {
	count, error := db.GetDb().Update(&article)
	if error != nil {
		log.Fatal(error)
	}
	return count
}

// Delete : It will delete article
func Delete(id int64) int64 {
	article := models.Article{
		ID: id,
	}
	count, error := db.GetDb().Delete(&article)
	if error != nil {
		log.Fatal(error)
	}
	return count
}
