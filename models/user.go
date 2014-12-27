package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/hoisie/redis"
	"net/websocket"
)

type User struct {
	Id       int
	Name     string
	Pw       string
	Cookie   string
	rediscli redis.Client
	ws       websocket.Conn
}

func NewUser(username string, password string) (u *User, err error) {

	u = &User{}
	u.Name = username
	u.Pw = password
	u.rediscli.Addr = beego.AppConfig.String("redis_addr")
	if u.Pw = beego.AppConfig.String("redis_auth"); "" != u.Pw {
		u.rediscli.Auth(u.Pw)
	}
	var userps []byte
	userps, err = u.rediscli.Hget("userlist", u.Name)
	if err != nil || u.Pw != fmt.Sprintf("%s", userps) {
		err = errors.New("Username Or Password Error")
	}

	return
}

// func (u *User) Auth(name string, pw string) (id int, ok bool) {

// 	redis_addr := beego.AppConfig.String("redis_addr")
// 	rediscli := redis.Client{
// 		Addr: redis_addr,
// 		Db:   0,
// 		// Password:    "pink",
// 		MaxPoolSize: 10000,
// 	}
// 	return 0, false
// }
