package db

import (
	"fmt"
	"os"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

var db orm.Ormer

func GetDb() orm.Ormer {
	if db == nil {
		db_name := os.Getenv("DB_NAME")
		db_user := os.Getenv("DB_USERNAME")
		db_password := os.Getenv("DB_PASSWORD")
		
		if db_name == "" {
			fmt.Println("Environment variable DB_NAME is null.")
			return nil;
		}
		if db_user == "" {
			fmt.Println("Environment variable DB_USERNAME is null.")
			return nil;
		}
		if db_password == "" {
			fmt.Println("Environment variable DB_PASSWORD is null.")
			return nil;
		}
		fmt.Println(db_name)
		orm.RegisterDriver("postgres", orm.DRPostgres)
		orm.RegisterDataBase("default", "postgres", "postgres://"+db_user+":"+db_password+"@localhost:5432/"+db_name+"?sslmode=disable")

		// This is for auto generating tables
		orm.RunSyncdb("default", false, true)

		db = orm.NewOrm()
		db.Using("default")

	}
	return db
}
