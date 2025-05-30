package models

import "github.com/beego/beego/v2/client/orm"

type Student struct {
	Id            int    `json:"id" orm:"auto"`
	Name          string `json:"name" orm:"type(varchar)"`
	Address       string `json:"address" orm:"type(varchar)"`
	Class         string `json:"class" orm:"type(varchar)"`
	StudentSchool string `json:"student_school" orm:"type(varchar); column(school)"`
}

func init() {
	orm.RegisterModel(new(Student))
}