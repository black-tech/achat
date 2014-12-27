package controllers

import (
	"achat/models"
	"github.com/astaxie/beego"
	// "github.com/hoisie/redis"
	"log"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Post() {
	uname := this.Ctx.Input.Cookie("username")
	// if uname != nil {
	log.Println("cookie :", uname)
	// }

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

	this.SetSecureCookie("achat", "username", user.Name)
	// this.Ctx.Output.Cookie("username", user.name)
	// cookie:= this.Ctx.Input.Cookie("username")
	// log.Println(cookie)
	this.Ctx.Output.Body([]byte("{\"code\":\"OK\"}"))

}
