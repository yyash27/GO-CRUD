package main

import (
	"log"

	_ "GO-CRUD/routers"
	_ "GO-CRUD/models"

	"github.com/beego/beego/v2/server/web"
)

func main() {
	log.Print("main run")
	web.Run()
}