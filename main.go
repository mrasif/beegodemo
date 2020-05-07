package main

import (
	"beegodemo/models"
	"beegodemo/repo"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello")

	fmt.Println(time.Now())
	article := models.Article{
		Title:       "Title 1",
		Description: "Description 1",
	}

	x := repo.Insert(article)
	fmt.Println(x)
	repo.Insert(article)

}
