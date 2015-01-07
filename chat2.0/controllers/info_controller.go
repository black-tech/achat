package controllers

import (
	"funx/models"
	"github.com/astaxie/beego"
	"strconv"
)

type InfoController struct {
	beego.Controller
}

func (this *InfoController) Get() {
	wp := models.NewWebPage("info")

	this.Data["PageTitle"] = "FANX-INFO"
	this.Data["StaticHost"] = beego.AppConfig.String("static_host")
	this.Data["ViewCount"] = strconv.Itoa(wp.IncrViewCount())
	this.Data["PageMap"] = wp.GetWebPages()
	this.Data["PageCount"] = wp.GetWebPageCount()

	this.TplNames = "info.tpl"
}
