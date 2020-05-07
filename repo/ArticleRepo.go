package repo

import (
	"beegodemo/db"
	"beegodemo/models"
	"fmt"
	"time"
)

func Insert(article models.Article) int64 {
	fmt.Println(article)
	article.CreatedAt = time.Now().UnixNano() / int64(time.Millisecond)
	article.UpdatedAt = time.Now().UnixNano() / int64(time.Millisecond)
	id, _ := db.GetDb().Insert(&article)

	fmt.Println("Article inserted")
	return id

}
