package controllers

import (
	"achat/models"
	"github.com/astaxie/beego"
	"log"
	// "github.com/hoisie/redis"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Post() {

	username := this.GetString("username")
	password := this.GetString("password")
	log.Println("username: ", username)
	log.Println("password: ", password)

	user, err := models.NewUser(username, password)
	if err != nil {
		log.Println("User: " + username + " Login failed:" + err.Error())
		this.Ctx.Output.Body([]byte("{\"code\":\"error\"}"))
		return
	}
	log.Println("User: #" + user.Name + "# Login successful:")
	user.LoginSucc()
	// log.Println("UesrSalt: " + user.Salt + "UserCookie" + user.Cookie)

	// this.SetSecureCookie("achat", "username", user.Name)
	this.Ctx.Output.Cookie("achat", user.Cookie)
	// cookie:= this.Ctx.Input.Cookie("username")
	// log.Println(cookie)
	this.Ctx.Output.Body([]byte("{\"code\":\"OK\"}"))
}
