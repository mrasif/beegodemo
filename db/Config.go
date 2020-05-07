package db

import (
	"fmt"
	"os"

	"github.com/astaxie/beego/orm"
)

var db orm.Ormer

// GetDb : It will returns orm object
func GetDb() orm.Ormer {
	if db == nil {
		dbHost := os.Getenv("DB_HOST")
		dbName := os.Getenv("DB_NAME")
		dbUser := os.Getenv("DB_USERNAME")
		dbPassword := os.Getenv("DB_PASSWORD")

		if dbHost == "" {
			fmt.Println("Environment variable DB_HOST is null.")
			return nil
		}
		if dbName == "" {
			fmt.Println("Environment variable DB_NAME is null.")
			return nil
		}
		if dbUser == "" {
			fmt.Println("Environment variable DB_USERNAME is null.")
			return nil
		}
		if dbPassword == "" {
			fmt.Println("Environment variable DB_PASSWORD is null.")
			return nil
		}

		orm.RegisterDriver("postgres", orm.DRPostgres)
		orm.RegisterDataBase("default", "postgres", "postgres://"+dbUser+":"+dbPassword+"@"+dbHost+":5432/"+dbName+"?sslmode=disable")

		// This is for auto generating tables
		orm.RunSyncdb("default", false, true)

		db = orm.NewOrm()
		db.Using("default")

	}
	return db
}
