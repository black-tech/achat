package controllers

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/hoisie/redis"
	"log"
)

type LoginController struct {
	beego.Controller
}

type User struct {
	id   int
	name string
	pw   string
}

func NewUser(username string, password string) (u *User, err error) {

	u = &User{}
	u.name = username
	u.pw = password

	redis_addr := beego.AppConfig.String("redis_addr")
	rediscli := redis.Client{
		Addr: redis_addr,
		Db:   0,
		// Password:    "pink",
		MaxPoolSize: 10000,
	}
	var userps []byte
	userps, err = rediscli.Hget("userlist", u.name)
	if err != nil || u.pw != fmt.Sprintf("%s", userps) {
		err = errors.New("Username Or Password Error")
	}

	return
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

	user, err := NewUser(username, password)
	if err != nil {
		log.Println("User: " + username + " Login failed:" + err.Error())
		this.Ctx.Output.Body([]byte("{\"code\":\"error\"}"))
		return
	}
	log.Println("User: #" + user.name + "# Login successful:")

	this.SetSecureCookie("achat", "username", user.name)
	// this.Ctx.Output.Cookie("username", user.name)
	// cookie:= this.Ctx.Input.Cookie("username")
	// log.Println(cookie)
	this.Ctx.Output.Body([]byte("{\"code\":\"OK\",\"password\":\"" + password + "\"}"))
}
