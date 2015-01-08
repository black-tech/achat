package controllers

import (
	"funx/models"
	"github.com/astaxie/beego"
)

type ChatController struct {
	beego.Controller
}

func (this *ChatController) Get() {
	wp := models.NewWebPage("chat")
	wp.IncrViewCount()

	this.Data["PageTitle"] = "FANX-CHAT"
	this.Data["StaticHost"] = beego.AppConfig.String("static_host")

	// testStr := `{"Code":"Y29kZXN0cmluZw==","Data":"aGVyb2hlcm9oZXJv","Desc":"ZGVzY2Rlc2NkZXNj","Supp":["dGVzdDExMQ==","dGVzdDIyMjI=","dGVzdDMzMzM="]}`
	// //msg := models.DecodeJson(testStr)
	// _ = testStr
	// msg := models.NewMsg()
	// msg.Code = "codestring"
	// msg.Data = "datastring"
	// msg.Supp = []string{"asdfasdf", "fdsafdas"}
	// str := msg.ToString()
	// this.Data["TestStr"] = str

	this.TplNames = "chat.tpl"
}
