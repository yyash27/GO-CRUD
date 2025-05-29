package routers

import (
	"GO-CRUD/controllers"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	// log.Print("entered router")
	web.Router("/", &controllers.AppController{}, "get:GetAll")

	// web.Router("/students", &controllers.StudentController{}, "post:Post")
	// web.Router("/students", &controllers.StudentController{}, "get:GetAll")
	// web.Router("/students/:id", &controllers.StudentController{}, "put:Put")
	// web.Router("/students/:id", &controllers.StudentController{}, "delete:Delete")
}