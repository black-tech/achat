package controllers

import (
	"funx/models"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	wp := models.NewWebPage("home")
	wp.IncrViewCount()

	this.Data["PageTitle"] = "FANX-HOME"
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.Data["StaticHost"] = beego.AppConfig.String("static_host")
	this.TplNames = "index.tpl"
}
