package main

import (
	_ "achat/routers"
	"github.com/astaxie/beego"
	"log"
)

func main() {
	log.Println("hello bee")
	beego.Run()
}
