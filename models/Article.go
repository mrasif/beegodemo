package models

import (
	"github.com/astaxie/beego/orm"
)

// Article : It model of article
type Article struct {
	ID          int64  `json:"id";orm:"pk;auto"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"createdAt"`
	UpdatedAt   int64  `json:"updatedAt"`
}

// TableName : It will returns table name
func (u *Article) TableName() string {
	return "articles"
}

func init() {
	orm.RegisterModel(new(Article))
}
