package controllers

import (
	// "log"

	"github.com/beego/beego/v2/server/web"
)

type AppController struct {
	web.Controller
}

func (c *AppController) GetAll() {
// 	log.Print("Hi from Get All")
	c.Data["json"] = map[string]string{
		"Id" : "1",
		"Name" : "Yash",
		"Full Name" : "Yash Jain",
	}
	c.ServeJSON()
}

