package controllers

import (
	"achat/models"
	"github.com/astaxie/beego"
)

var online_count int

type ChatController struct {
	beego.Controller
}

func (this *ChatController) Get() {

	ch := models.NewChatHome()

	this.Data["cookie"] = this.Ctx.Input.Cookie("username")
	this.Data["online_count"] = ch.GetOnlineCount()
	this.Data["max_online_count"] = ch.MAX_ONLINE_COUNT

	var u User
	u.name = "user1"
	this.Data["username"] = u.name
	b, _ := ch.IsOnline(u.name)
	this.Data["testData"] = b

	this.TplNames = "chat.tpl"
}
