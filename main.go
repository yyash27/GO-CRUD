package main

import (
	"log"

	_ "GO-CRUD/routers"

	"github.com/beego/beego/v2/server/web"
)

func main() {
	log.Print("main run")
	web.Run()
}