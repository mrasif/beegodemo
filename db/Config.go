package db

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

var db orm.Ormer

func GetDb() orm.Ormer {
	if db == nil {
		orm.RegisterDriver("postgres", orm.DRPostgres)
		orm.RegisterDataBase("default", "postgres", "postgres://asif@localhost:5432/beegodemo?sslmode=disable")

		// This is for auto generating tables
		orm.RunSyncdb("default", false, true)

		db = orm.NewOrm()
		db.Using("default")

	}
	return db
}
