package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

// Article : It model of article
type Article struct {
	ID          int64  `json:"id";orm:"pk;auto"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedAt   time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt   time.Time `orm:"auto_now;type(datetime)"`
}

// TableName : It will returns table name
func (a *Article) TableName() string {
	return "articles"
}

func init() {
	orm.RegisterModel(new(Article))
}
