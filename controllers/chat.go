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

	u, _ := models.NewUser("user1", "hero")
	this.Data["username"] = u.Name
	b, _ := ch.IsOnline(u.Name)
	this.Data["testData"] = b

	this.TplNames = "chat.tpl"
}
