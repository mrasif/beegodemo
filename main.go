package main

import (
	"beegodemo/models"
	"beegodemo/repo"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Hello")

	fmt.Println(time.Now())
	article := models.Article{
		Title:       "Title 1",
		Description: "Description 1",
	}

	// Insert
	id1 := repo.Save(article)
	fmt.Printf("Id: %d\n", id1)

	// Retrive
	article1 := repo.Get(id1)
	fmt.Println(article1)

	// Update
	article1.Description = "New Description"
	count := repo.Update(article1)
	fmt.Printf("Count: %d\n", count)

	article2 := repo.Get(id1)
	fmt.Println(article2)

	// Get all
	articles := repo.GetAll()
	fmt.Println(*articles)

	// Delete
	count = repo.Delete(id1)
	fmt.Printf("Count: %d\n", count)
}
