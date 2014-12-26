package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/hoisie/redis"
)

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
