package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type Demo struct {
	Id int `json:"id" orm:"auto"`
	Name string `json:"name"`
}

func init() {
	orm.RegisterModel(new(Demo))
}