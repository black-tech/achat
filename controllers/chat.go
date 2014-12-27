package controllers

import (
	"github.com/astaxie/beego"
)

type ChatController struct {
	beego.Controller
}

func (this *ChatController) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplNames = "chat.tpl"
}
