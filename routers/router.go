package routers

import (
	"GO-CRUD/controllers"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	// log.Print("entered router")
	web.Router("/", &controllers.AppController{}, "get:GetAll")
	web.Router("/students", &controllers.AppController{}, "get:GetAllStudent")
	web.Router("/students/:id", &controllers.AppController{}, "get:GetStudent")
	web.Router("/student", &controllers.AppController{}, "post:CreateStudent")
	web.Router("/student/:id", &controllers.AppController{}, "post:UpdateStudent")
	web.Router("/student/:id", &controllers.AppController{}, "delete:DeleteStudent")
}